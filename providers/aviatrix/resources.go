package aviatrix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_account",
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
					Description: `(Required) Type of cloud service provider. Only AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, Alibaba Cloud and Edge CSP are supported currently. Enter 1 for AWS, 4 for GCP, 8 for Azure, 16 for OCI, 32 for AzureGov, 256 for AWSGov, 1024 for AWSChina or 2048 for AzureChina, 8192 for Alibaba Cloud, 65536 for Edge CSP. ### AWS ~>`,
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
					Description: `(Optional) AWS Access Key. Required when ` + "`" + `aws_iam` + "`" + ` is "false" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(Optional) AWS Secret Key. Required when ` + "`" + `aws_iam` + "`" + ` is "false" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_role_app",
					Description: `(Optional) AWS App role ARN, this option is for UserConnect. Required when ` + "`" + `aws_iam` + "`" + ` is "true" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_role_ec2",
					Description: `(Optional) AWS EC2 role ARN, this option is for UserConnect. Required when ` + "`" + `aws_iam` + "`" + ` is "true" and when creating an account for AWS.`,
				},
				resource.Attribute{
					Name:        "aws_gateway_role_app",
					Description: `(Optional) A separate AWS App role ARN to assign to gateways created by the controller. Required when ` + "`" + `aws_gateway_role_ec2` + "`" + ` is set. Only allowed when ` + "`" + `aws_iam` + "`" + `, ` + "`" + `awsgov_iam` + "`" + `, or ` + "`" + `awschina_iam` + "`" + ` is "true" when creating an account for AWS, AWSGov or AWSChina, respectively. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "aws_gateway_role_ec2",
					Description: `(Optional) A separate AWS EC2 role ARN to assign to gateways created by the controller. Required when ` + "`" + `aws_gateway_role_app` + "`" + ` is set. Only allowed when ` + "`" + `aws_iam` + "`" + `, ` + "`" + `awsgov_iam` + "`" + `, or ` + "`" + `awschina_iam` + "`" + ` is "true" when creating an account for AWS, AWSGov or AWSChina, respectively. Available as of provider version R2.19+. ### Azure`,
				},
				resource.Attribute{
					Name:        "arm_subscription_id",
					Description: `(Optional) Azure ARM Subscription ID. Required when creating an account for Azure.`,
				},
				resource.Attribute{
					Name:        "arm_directory_id",
					Description: `(Optional) Azure ARM Directory ID. Required when creating an account for Azure.`,
				},
				resource.Attribute{
					Name:        "arm_application_id",
					Description: `(Optional) Azure ARM Application ID. Required when creating an account for Azure.`,
				},
				resource.Attribute{
					Name:        "arm_application_key",
					Description: `(Optional) Azure ARM Application key. Required when creating an account for Azure. ### Google Cloud`,
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
					Description: `(Optional) Oracle OCI Tenancy ID. Required when creating an account for OCI.`,
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
					Description: `(Optional) Oracle OCI API Private Key local file path. Required when creating an account for OCI. ### AzureGov Cloud`,
				},
				resource.Attribute{
					Name:        "azuregov_subscription_id",
					Description: `(Optional) AzureGov ARM Subscription ID. Required when creating an account for AzureGov. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "azuregov_directory_id",
					Description: `(Optional) AzureGov ARM Directory ID. Required when creating an account for AzureGov. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "azuregov_application_id",
					Description: `(Optional) AzureGov ARM Application ID. Required when creating an account for AzureGov. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "azuregov_application_key",
					Description: `(Optional) AzureGov ARM Application key. Required when creating an account for AzureGov. Available as of provider version R2.19+. ### AWSGov Cloud`,
				},
				resource.Attribute{
					Name:        "awsgov_account_number",
					Description: `(Optional) AWSGov Account number to associate with Aviatrix account. Required when creating an account for AWSGov.`,
				},
				resource.Attribute{
					Name:        "awsgov_iam",
					Description: `(Optional) AWSGov IAM-role based flag. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awsgov_access_key",
					Description: `(Optional) AWS Access Key. Required when creating an account for AWSGov.`,
				},
				resource.Attribute{
					Name:        "awsgov_secret_key",
					Description: `(Optional) AWS Secret Key. Required when creating an account for AWSGov.`,
				},
				resource.Attribute{
					Name:        "awsgov_role_app",
					Description: `(Optional) AWSGov App role ARN. Available when ` + "`" + `awsgov_iam` + "`" + ` is "true" and when creating an account for AWSGov. If left empty, the ARN will be computed. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awsgov_role_ec2",
					Description: `(Optional) AWSGov EC2 role ARN. Available when ` + "`" + `awsgov_iam` + "`" + ` is "true" and when creating an account for AWSGov. If left empty, the ARN will be computed. Available as of provider version 2.19+. ### AWSChina Cloud`,
				},
				resource.Attribute{
					Name:        "awschina_account_number",
					Description: `(Optional) AWSChina Account number to associate with Aviatrix account. Required when creating an account for AWSChina. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_iam",
					Description: `(Optional) AWSChina IAM-role based flag. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_role_app",
					Description: `(Optional) AWSChina App role ARN. Available when ` + "`" + `awschina_iam` + "`" + ` is "true" and when creating an account for AWSChina. If left empty, the ARN will be computed. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_role_ec2",
					Description: `(Optional) AWSChina EC2 role ARN. Available when ` + "`" + `awschina_iam` + "`" + ` is "true" and when creating an account for AWSChina. If left empty, the ARN will be computed. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_access_key",
					Description: `(Optional) AWSChina Access Key. Required when ` + "`" + `awschina_iam` + "`" + ` is "false" and when creating an account for AWSChina. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "awschina_secret_key",
					Description: `(Optional) AWSChina Secret Key. Required when ` + "`" + `awschina_iam` + "`" + ` is "false" and when creating an account for AWSChina. Available as of provider version 2.19+. ### AzureChina Cloud`,
				},
				resource.Attribute{
					Name:        "azurechina_subscription_id",
					Description: `(Optional) AzureChina ARM Subscription ID. Required when creating an account for AzureChina. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "azurechina_directory_id",
					Description: `(Optional) AzureChina ARM Directory ID. Required when creating an account for AzureChina. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "azurechina_application_id",
					Description: `(Optional) AzureChina ARM Application ID. Required when creating an account for AzureChina. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "azurechina_application_key",
					Description: `(Optional) AzureChina ARM Application key. Required when creating an account for AzureChina. Available as of provider version 2.19+. ### Alibaba Cloud`,
				},
				resource.Attribute{
					Name:        "alicloud_account_id",
					Description: `(Optional) Alibaba Cloud Account number to associate with Aviatrix account. Required when creating an account for Alibaba Cloud.`,
				},
				resource.Attribute{
					Name:        "alicloud_access_key",
					Description: `(Optional) Alibaba Cloud Access Key. Required when creating an account for Alibaba Cloud.`,
				},
				resource.Attribute{
					Name:        "alicloud_secret_key",
					Description: `(Optional) Alibaba Cloud Secret Key. Required when creating an account for Alibaba Cloud. ### AWS Top Secret Region`,
				},
				resource.Attribute{
					Name:        "awsts_account_number",
					Description: `(Optional) AWS Top Secret Region Account Number. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_url",
					Description: `(Optional) AWS Top Secret Region CAP Url. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_agency",
					Description: `(Optional) AWS Top Secret Region CAP Agency. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_mission",
					Description: `(Optional) AWS Top Secret Region Mission. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_role_name",
					Description: `(Optional) AWS Top Secret Region Role Name. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert",
					Description: `(Optional) AWS Top Secret Region CAP Certificate local file path. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert_key",
					Description: `(Optional) AWS Top Secret Region CAP Certificate Key local file path. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_ca_chain_cert",
					Description: `(Optional) AWS Top Secret Region Custom Certificate Authority local file path. Required when creating an account in AWS Top Secret Region. Available as of provider version R2.19.5+. ### AWS Secret Region`,
				},
				resource.Attribute{
					Name:        "awss_account_number",
					Description: `(Optional) AWS Secret Region Account Number. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_url",
					Description: `(Optional) AWS Secret Region CAP Url. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_agency",
					Description: `(Optional) AWS Secret Region CAP Agency. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_account_name",
					Description: `(Optional) AWS Secret Region Account Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_role_name",
					Description: `(Optional) AWS Secret Region Role Name. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert",
					Description: `(Optional) AWS Secret Region CAP Certificate local file path. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_key",
					Description: `(Optional) AWS Secret Region CAP Certificate Key local file path. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_ca_chain_cert",
					Description: `(Optional) AWS Secret Region Custom Certificate Authority local file path. Required when creating an account in AWS Secret Region. Available as of provider version R2.19.5+. ### Edge CSP`,
				},
				resource.Attribute{
					Name:        "edge_csp_username",
					Description: `(Optional) Edge CSP username. Required when creating an Edge CSP account.`,
				},
				resource.Attribute{
					Name:        "edge_csp_password",
					Description: `(Optional) Edge CSP password. Required when creating an Edge CSP account. ### Misc. ~>`,
				},
				resource.Attribute{
					Name:        "audit_account",
					Description: `(Optional) Specify whether to enable the audit account feature. If this feature is enabled, terraform will give a warning if there is an issue with the account credentials. Changing ` + "`" + `audit_account` + "`" + ` to "false" will not prevent the Controller from performing account audits. It will only prevent Terraform from displaying a warning. Valid values: true, false. Default: false. Available as of provider version 2.19+.`,
				},
				resource.Attribute{
					Name:        "rbac_groups",
					Description: `(Optional) A list of existing RBAC group names. This attribute should only be used when creating an account. Updating this attribute will have no effect. Available as of provider version R2.23.0+. ->`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert_path",
					Description: `(Optional) AWS Top Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert_key_path",
					Description: `(Optional) AWS Top Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "aws_ca_cert_path",
					Description: `(Optional) AWS Top Secret Region or Secret Region Custom Certificate Authority file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_path",
					Description: `(Optional) AWS Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_key_path",
					Description: `(Optional) AWS Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "awsts_cap_cert_path",
					Description: `(Optional) AWS Top Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awsts_cap_cert_key_path",
					Description: `(Optional) AWS Top Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "aws_ca_cert_path",
					Description: `(Optional) AWS Top Secret Region or Secret Region Custom Certificate Authority file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_path",
					Description: `(Optional) AWS Secret Region CAP Certificate file name on the controller. Available as of provider R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "awss_cap_cert_key_path",
					Description: `(Optional) AWS Secret Region CAP Certificate Key file name on the controller. Available as of provider R2.19.5+. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_account_user",
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
					Description: `(Required) Name of account user to be created. It can only include alphanumeric characters(lower case only), hyphens, dots or underscores. 1 to 80 in length. No spaces are allowed.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email of address of account user to be created.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Login password for the account user to be created. If password is changed, current account will be destroyed and a new account will be created. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Cloud account name of user to be created. Deprecated as of Aviatrix provider R2.13 (Controller 5.4) due to RBAC implementation. ->`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R2.13+, and an`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_guard_duty",
			Category:         "Security",
			ShortDescription: `Manage AWS GuardDuty configuration`,
			Description:      ``,
			Keywords: []string{
				"security",
				"aws",
				"guard",
				"duty",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account name.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region.`,
				},
				resource.Attribute{
					Name:        "excluded_ips",
					Description: `(Optional) Set of excluded IPs. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_peer",
			Category:         "Peering",
			ShortDescription: `Creates and manages native AWS VPC peerings`,
			Description:      ``,
			Keywords: []string{
				"peering",
				"aws",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name1",
					Description: `(Required) Account name of AWS VPC1.`,
				},
				resource.Attribute{
					Name:        "account_name2",
					Description: `(Required) Account name of AWS VPC2.`,
				},
				resource.Attribute{
					Name:        "vpc_id1",
					Description: `(Required) VPC ID of AWS VPC1. Example: AWS: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "vpc_id2",
					Description: `(Required) VPC ID of AWS VPC2. Example: AWS: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "vpc_reg1",
					Description: `(Required) Region of AWS VPC1. Example: AWS: "us-east-1".`,
				},
				resource.Attribute{
					Name:        "vpc_reg2",
					Description: `(Required) Region of AWS VPC2. Example: AWS: "us-east-1".`,
				},
				resource.Attribute{
					Name:        "rtb_list1",
					Description: `(Optional) List of Route table IDs of VPC1. Example: ["rtb-abcd1234", "rtb-wxyz5678"].`,
				},
				resource.Attribute{
					Name:        "rtb_list2",
					Description: `(Optional) List of Route table IDs of VPC2. Example: ["rtb-abcd1234", "rtb-wxyz5678"]. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw",
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
					Description: `(Required) Name of the AWS TGW to be created`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Name of the cloud account in the Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region of AWS TGW to be created in`,
				},
				resource.Attribute{
					Name:        "aws_side_as_number",
					Description: `(Required) BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001". ### Misc.`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Optional) Type of cloud service provider, requires an integer value. Supported for AWS (1) and AWSGov (256). Default value: 1.`,
				},
				resource.Attribute{
					Name:        "enable_multicast",
					Description: `(Optional) Enable multicast. Default value: false. Valid values: true, false. Available in provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "cidrs",
					Description: `(Optional) Set of TGW CIDRs. For example, ` + "`" + `cidrs = ["10.0.10.0/24", "10.1.10.0/24"]` + "`" + `. Available as of provider version R2.18.1+.`,
				},
				resource.Attribute{
					Name:        "inspection_mode",
					Description: `(Optional) Inspection mode. Valid values: "Domain-based", "Connection-based". Default value: "Domain-based". Available as of provider version R2.23+. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "tgw_id",
					Description: `TGW ID. Available as of provider version R2.19+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "security_domains",
					Description: `(Required if ` + "`" + `manage_security_domain` + "`" + ` is true) Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Required) Three default domains ("Aviatrix_Edge_Domain", "Default_Domain" and "Shared_Service_Domain") are required with AWS TGW's creation.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firewall",
					Description: `(Optional) Set to true if the security domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_egress",
					Description: `(Optional) Set to true if the security domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_firewall",
					Description: `(Optional) Set to true if the security domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.`,
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
					Description: `(Required) Region of the VPC, needs to be consistent with AWS TGW's region.`,
				},
				resource.Attribute{
					Name:        "vpc_account_name",
					Description: `(Required) Cloud account name of the VPC in the Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the VPC to be attached to the security domain`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) Advanced option. VPC subnets separated by ',' to attach to the VPC. If left blank, the Aviatrix Controller automatically selects a subnet representing each AZ for the VPC attachment. Example: "subnet-214f5646,subnet-085e8c81a89d70846".`,
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
					Description: `(Optional) Advanced option. Customized route(s) to be advertised to other VPCs that are connected to the same TGW. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "disable_local_route_propagation",
					Description: `(Optional) Advanced option. If set to true, it disables automatic route propagation of this VPC to other VPCs within the same security domain. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "attached_aviatrix_transit_gateway",
					Description: `(Optional) A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.`,
				},
				resource.Attribute{
					Name:        "manage_security_domain",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage security domains using the`,
				},
				resource.Attribute{
					Name:        "manage_transit_gateway_attachment",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the`,
				},
				resource.Attribute{
					Name:        "manage_vpc_attachment",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_id",
					Description: `TGW ID. Available as of provider version R2.19+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "security_domains",
					Description: `(Required if ` + "`" + `manage_security_domain` + "`" + ` is true) Security Domains to create together with AWS TGW's creation. Three default domains, along with the connections between them, are created automatically. These three domains can't be deleted, but the connection between any two of them can be.`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Required) Three default domains ("Aviatrix_Edge_Domain", "Default_Domain" and "Shared_Service_Domain") are required with AWS TGW's creation.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firewall",
					Description: `(Optional) Set to true if the security domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_egress",
					Description: `(Optional) Set to true if the security domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_firewall",
					Description: `(Optional) Set to true if the security domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false.`,
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
					Description: `(Required) Region of the VPC, needs to be consistent with AWS TGW's region.`,
				},
				resource.Attribute{
					Name:        "vpc_account_name",
					Description: `(Required) Cloud account name of the VPC in the Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the VPC to be attached to the security domain`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) Advanced option. VPC subnets separated by ',' to attach to the VPC. If left blank, the Aviatrix Controller automatically selects a subnet representing each AZ for the VPC attachment. Example: "subnet-214f5646,subnet-085e8c81a89d70846".`,
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
					Description: `(Optional) Advanced option. Customized route(s) to be advertised to other VPCs that are connected to the same TGW. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "disable_local_route_propagation",
					Description: `(Optional) Advanced option. If set to true, it disables automatic route propagation of this VPC to other VPCs within the same security domain. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "attached_aviatrix_transit_gateway",
					Description: `(Optional) A list of names of Aviatrix Transit Gateway(s) (transit VPCs) to attach to the Aviatrix_Edge_Domain.`,
				},
				resource.Attribute{
					Name:        "manage_security_domain",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage security domains using the`,
				},
				resource.Attribute{
					Name:        "manage_transit_gateway_attachment",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage transit gateway attachments to the TGW using the`,
				},
				resource.Attribute{
					Name:        "manage_vpc_attachment",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage VPC attachments to the TGW using the`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_connect",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix AWS TGW Connect connections`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"connect",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) AWS TGW name.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Connection name.`,
				},
				resource.Attribute{
					Name:        "transport_vpc_id",
					Description: `(Required) Transport Attachment VPC ID.`,
				},
				resource.Attribute{
					Name:        "network_domain_name",
					Description: `(Required) Network Domain name. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connect_attachment_id",
					Description: `Connect Attachment ID.`,
				},
				resource.Attribute{
					Name:        "transport_attachment_id",
					Description: `Transport Attachment ID. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Optional) Security Domain name. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connect_attachment_id",
					Description: `Connect Attachment ID.`,
				},
				resource.Attribute{
					Name:        "transport_attachment_id",
					Description: `Transport Attachment ID. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Optional) Security Domain name. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_connect_peer",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix AWS TGW Connect peers`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"connect",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) AWS TGW name.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) TGW Connect connection name.`,
				},
				resource.Attribute{
					Name:        "connect_peer_name",
					Description: `(Required) TGW Connect peer name.`,
				},
				resource.Attribute{
					Name:        "connect_attachment_id",
					Description: `(Required) Connect Attachment ID.`,
				},
				resource.Attribute{
					Name:        "peer_as_number",
					Description: `(Required) Peer AS Number.`,
				},
				resource.Attribute{
					Name:        "peer_gre_address",
					Description: `(Required) Peer GRE IP Address.`,
				},
				resource.Attribute{
					Name:        "bgp_inside_cidrs",
					Description: `(Required) Set of BGP Inside CIDR Block(s).`,
				},
				resource.Attribute{
					Name:        "tgw_gre_address",
					Description: `(Optional) AWS TGW GRE IP Address. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connect_peer_id",
					Description: `Connect Peer ID. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connect_peer_id",
					Description: `Connect Peer ID. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_directconnect",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix AWS TGW DirectConnect connections`,
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
					Name:        "allowed_prefix",
					Description: `(Required) A list of comma separated CIDRs for DXGW to advertise to remote(on-prem).`,
				},
				resource.Attribute{
					Name:        "network_domain_name",
					Description: `(Required) The name of a network domain, to which the direct connect gateway will be attached. ### Optional`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable/disable [encrypted transit approval](https://docs.aviatrix.com/HowTos/tgw_approval.html) for AWS TGW DirectConnect. Valid values: true, false. Default value: false. The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(Optional) The name of a security domain, to which the direct connect gateway will be attached. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_intra_domain_inspection",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages the intra domain inspection of network domains in an AWS TGW`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"intra",
				"domain",
				"inspection",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) The AWS TGW name.`,
				},
				resource.Attribute{
					Name:        "route_domain_name",
					Description: `(Required) The name of a network domain.`,
				},
				resource.Attribute{
					Name:        "firewall_domain_name",
					Description: `(Required) The name of a firewall network domain. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_network_domain",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix network domains`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"network",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network domain.`,
				},
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) The AWS TGW name of the network domain.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firewall",
					Description: `(Optional) Set to true if the network domain is to be used as an Aviatrix Firewall Domain for the Aviatrix Firewall Network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_egress",
					Description: `(Optional) Set to true if the network domain is to be used as a native egress domain (for non-Aviatrix Firewall Network-based central Internet bound traffic). Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "native_firewall",
					Description: `(Optional) Set to true if the network domain is to be used as a native firewall domain (for non-Aviatrix Firewall Network-based firewall traffic inspection). Valid values: true, false. Default value: false. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_peering",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix inter-region AWS TGW peerings`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name1",
					Description: `(Required) This parameter represents name of the first AWS TGW to make a peer pair.`,
				},
				resource.Attribute{
					Name:        "tgw_name2",
					Description: `(Required) This parameter represents name of the second AWS TGW to make a peer pair. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_peering_domain_conn",
			Category:         "TGW Orchestrator",
			ShortDescription: `Creates and manages Aviatrix domain connections between peered AWS TGWs`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"peering",
				"domain",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name1",
					Description: `(Required) The AWS TGW name of the source domain to make a connection.`,
				},
				resource.Attribute{
					Name:        "domain_name1",
					Description: `(Required) The name of the source domain to make a connection.`,
				},
				resource.Attribute{
					Name:        "tgw_name2",
					Description: `(Required) The AWS TGW name of the destination domain to make a connection.`,
				},
				resource.Attribute{
					Name:        "domain_name2",
					Description: `(Required) The name of the destination domain to make a connection. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_transit_gateway_attachment",
			Category:         "TGW Orchestrator",
			ShortDescription: `Manages attachment of transit gateways to the AWS TGW`,
			Description:      ``,
			Keywords: []string{
				"tgw",
				"orchestrator",
				"aws",
				"transit",
				"gateway",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tgw_name",
					Description: `(Required) Name of the AWS TGW.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS Region of the TGW.`,
				},
				resource.Attribute{
					Name:        "vpc_account_name",
					Description: `(Required) The name of the cloud account in the Aviatrix controller, which is associated with the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the VPC, where transit gateway is launched.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_name",
					Description: `(Required) Name of the transit gateway to be attached to the AWS TGW. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_vpc_attachment",
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
					Description: `(Required) AWS Region of the TGW.`,
				},
				resource.Attribute{
					Name:        "vpc_account_name",
					Description: `(Required) The name of the cloud account in the Aviatrix controller, which is associated with the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the VPC to be attached to the specified ` + "`" + `network_domain_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_domain_name",
					Description: `(Required) The name of the network domain, to which the VPC will be attached to. If changed, the VPC will be detached from the old domain, and attached to the new domain. ->`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional and ForceNew) Advanced option. VPC subnets separated by ',' to attach to the VPC. If omitted, the Aviatrix Controller automatically computes a subnet representing each AZ for the VPC attachment and Terraform will not manage this attribute. Example: "subnet-214f5646,subnet-085e8c81a89d70846".`,
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
					Description: `(Optional and ForceNew) Advanced option. Customized route(s) to be advertised to other VPCs that are connected to the same TGW. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16".`,
				},
				resource.Attribute{
					Name:        "disable_local_route_propagation",
					Description: `(Optional and ForceNew) Advanced option. If set to true, it disables automatic route propagation of this VPC to other VPCs within the same network domain. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "edge_attachment",
					Description: `(Optional) Advanced option. To allow access to the private IP of the MGMT interface of the Firewalls, set this attribute to enable Management Access From Onprem. This feature advertises the Firewalls private MGMT subnet to your Edge domain. Example: "vpn-0068bb31917ff2289". The following argument is deprecated:`,
				},
				resource.Attribute{
					Name:        "security_domain_name",
					Description: `(optional) The name of the security domain, to which the VPC will be attached to. If changed, the VPC will be detached from the old domain, and attached to the new domain. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_aws_tgw_vpn_conn",
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
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R2.11.0+, and an`,
				},
				resource.Attribute{
					Name:        "remote_as_number",
					Description: `(Optional) AWS side as a number. Integer between 1-4294967294. Example: "12".`,
				},
				resource.Attribute{
					Name:        "remote_cidr",
					Description: `(Optional) Remote CIDRs separated by ",". Example: AWS: "16.0.0.0/16,16.1.0.0/16".`,
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
					Description: `(Optional) Pre-Shared Key for Tunnel 2. A 8-64 character string with alphanumeric underscore(_) and dot(.). It cannot start with 0.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable/disable [encrypted transit approval](https://docs.aviatrix.com/HowTos/tgw_approval.html) for AWS TGW VPN connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_global_acceleration",
					Description: `(Optional) Enable Global Acceleration. Type: Boolean. Default: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpn_id",
					Description: `ID of the VPN generated by creation of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_tunnel_data",
					Description: `AWS TGW VPN tunnel data. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpn_id",
					Description: `ID of the VPN generated by creation of the connection.`,
				},
				resource.Attribute{
					Name:        "vpn_tunnel_data",
					Description: `AWS TGW VPN tunnel data. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_azure_peer",
			Category:         "Peering",
			ShortDescription: `Creates and manages of the Aviatrix peerings between Azure VNets`,
			Description:      ``,
			Keywords: []string{
				"peering",
				"azure",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name1",
					Description: `(Required) Name of the Azure cloud account in the Aviatrix controller for VNet 1.`,
				},
				resource.Attribute{
					Name:        "account_name2",
					Description: `(Required) Name of the Azure cloud account in the Aviatrix controller for VNet 2. ->`,
				},
				resource.Attribute{
					Name:        "vnet_name_resource_group1",
					Description: `(Required) Azure VNet 1's name. Example: "VNet_Name1:Resource_Group_Name1:GUID1".`,
				},
				resource.Attribute{
					Name:        "vnet_name_resource_group2",
					Description: `(Required) Azure VNet 2's name. Example: "VNet_Name2:Resource_Group_Name2:GUID2".`,
				},
				resource.Attribute{
					Name:        "vnet_reg1",
					Description: `(Required) Region of Azure VNet 1. Example: "East US 2".`,
				},
				resource.Attribute{
					Name:        "vnet_reg2",
					Description: `(Required) Region of Azure VNet 2. Example: "East US 2". ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vnet_cidr1",
					Description: `List of VNet CIDR of vnet_name_resource_group1.`,
				},
				resource.Attribute{
					Name:        "vnet_cidr2",
					Description: `List of VNet CIDR of vnet_name_resource_group2. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vnet_cidr1",
					Description: `List of VNet CIDR of vnet_name_resource_group1.`,
				},
				resource.Attribute{
					Name:        "vnet_cidr2",
					Description: `List of VNet CIDR of vnet_name_resource_group2. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_azure_spoke_native_peering",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix Azure spoke native peerings`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"azure",
				"spoke",
				"native",
				"peering",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transit_gateway_name",
					Description: `(Required) Name of an Transit FireNet-enabled Azure transit gateway.`,
				},
				resource.Attribute{
					Name:        "spoke_account_name",
					Description: `(Required) An Aviatrix account that corresponds to a subscription in Azure.`,
				},
				resource.Attribute{
					Name:        "spoke_region",
					Description: `(Required) Spoke VNet region. Example: "West US". ->`,
				},
				resource.Attribute{
					Name:        "spoke_vpc_id",
					Description: `(Required) Combination of the Spoke's VNet name, resource group and GUID. Example: "Foo_VNet:Bar_RG:GUID". ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_azure_vng_conn",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages the connection between Aviatrix Transit Gateway and Azure VNG`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"azure",
				"vng",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_gateway_name",
					Description: `(Required) Primary Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Connection name. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vng_name",
					Description: `Name of Azure VNG.`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `The status of the connection. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `VPC ID.`,
				},
				resource.Attribute{
					Name:        "vng_name",
					Description: `Name of Azure VNG.`,
				},
				resource.Attribute{
					Name:        "attached",
					Description: `The status of the connection. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_centralized_transit_firenet",
			Category:         "Firewall Network",
			ShortDescription: `Creates and manages the centralized Transit FireNet`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"centralized",
				"transit",
				"firenet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_firenet_gw_name",
					Description: `(Required) Primary FireNet gateway name.`,
				},
				resource.Attribute{
					Name:        "secondary_firenet_gw_name",
					Description: `(Required) Secondary FireNet gateway name. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_cloudn_registration",
			Category:         "CloudN",
			ShortDescription: `Register and Deregister CloudN as a Gateway`,
			Description:      ``,
			Keywords: []string{
				"cloudn",
				"registration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Gateway name to assign to the CloudN device. Type: String.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) Aviatrix CloudN's public or private IP. Type: String.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Aviatrix account username which will be used to log in to Aviatrix CloudN. Type: String.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Aviatrix account password corresponding to above username. Type: String. ### Optional`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) BGP AS Number to assign to the Transit Gateway. Type: String.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Requires local_as_number to be set. Type: List. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_cloudn_transit_gateway_attachment",
			Category:         "CloudN",
			ShortDescription: `Create and manage CloudN Transit Gateway Attachments`,
			Description:      ``,
			Keywords: []string{
				"cloudn",
				"transit",
				"gateway",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) CloudN device name. Type: String.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_name",
					Description: `(Required) Transit Gateway Name. Type: String.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Connection Name. Type: String.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_bgp_asn",
					Description: `(Required) Transit Gateway BGP AS Number. Type: String.`,
				},
				resource.Attribute{
					Name:        "cloudn_bgp_asn",
					Description: `(Required) CloudN BGP AS Number. Type: String.`,
				},
				resource.Attribute{
					Name:        "cloudn_lan_interface_neighbor_ip",
					Description: `(Required) CloudN LAN Interface Neighbor's IP Address. Type: String.`,
				},
				resource.Attribute{
					Name:        "cloudn_lan_interface_neighbor_bgp_asn",
					Description: `(Required) CloudN LAN Interface Neighbor's AS Number. Type: String. ### Optional`,
				},
				resource.Attribute{
					Name:        "enable_over_private_network",
					Description: `(Optional) Enable connection over private network. Type: Boolean. Default: true.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Enable Jumbo Frame support for the connection. Type: Boolean. Default: false.`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional) Enable Dead Peer Detection. Type: Boolean. Default: true.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Enable learned CIDRs approval. Type: Boolean. Default: false. Available as of provider version R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "approved_cidrs",
					Description: `(Optional/Computed) Set of approved CIDRs. Requires ` + "`" + `enable_learned_cidrs_approval` + "`" + ` to be true. Type: Set(String). Available as of provider version R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Requires transit_gateway_bgp_asn to be set. Type: List. Available as of provider version R2.21.0+. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_cloudwatch_agent",
			Category:         "Settings",
			ShortDescription: `Enables and disables cloudwatch_agent`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"cloudwatch",
				"agent",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of cloudwatch agent. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of cloudwatch agent. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_access_allow_list_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix Controller's Access Allow List Config`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"access",
				"allow",
				"list",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_list",
					Description: `(Required) Set of IPs allowed access to the Controller.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP address allowed access to the Controller.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the IP address. ### Optional`,
				},
				resource.Attribute{
					Name:        "enable_enforce",
					Description: `(Optional) Set to true to enable enforcement of the ` + "`" + `allow_list {}` + "`" + `'s IPs. Valid values: true, false. Default: false. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_bgp_max_as_limit_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller BGP max AS limit for transit gateways`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"bgp",
				"max",
				"as",
				"limit",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_as_limit",
					Description: `The maximum AS path limit allowed by transit gateways when handling BGP/Peering route propagation. Must be a number in the range [1-254]. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_cert_domain_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller cert domain config`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"cert",
				"domain",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert_domain",
					Description: `(Optional) Domain name that is used in FQDN for generating cert. Default value: "aviatrixnetwork.com". ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_config",
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
					Name:        "target_version",
					Description: `(Optional) The release version number to which the controller will be upgraded to. If not specified, controller will not be upgraded. If set to "latest", controller will be upgraded to the latest release. Please see the [Controller upgrade guide](https://docs.aviatrix.com/HowTos/inline_upgrade.html) for more information.`,
				},
				resource.Attribute{
					Name:        "manage_gateway_upgrades",
					Description: `(Optional) If true, aviatrix_controller_config will upgrade all gateways when target_version is set. If false, only the controller will be upgraded when target_version is set. In that case gateway upgrades should be handled in each gateway resource individually using the software_version and image_version attributes. Type: boolean. Default: true. Available as of provider version R2.20.0+. ### Security Options`,
				},
				resource.Attribute{
					Name:        "http_access",
					Description: `(Optional) Switch for HTTP access. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "fqdn_exception_rule",
					Description: `(Optional) Enable/disable packets without an SNI field to pass through gateway(s). Valid values: true, false. Default value: true. For more information on this setting, please see [here](https://docs.aviatrix.com/HowTos/FQDN_Whitelists_Ref_Design.html#exception-rule)`,
				},
				resource.Attribute{
					Name:        "aws_guard_duty_scanning_interval",
					Description: `(Optional) Configure the AWS Guard Duty scanning interval. Valid values: 5, 10, 15, 30 or 60. Default value: 60. Available as of provider version R2.18+. ### Backup`,
				},
				resource.Attribute{
					Name:        "backup_configuration",
					Description: `(Optional) Switch to enable/disable controller CloudN backup config. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "backup_cloud_type",
					Description: `(Optional) Type of cloud service provider, requires an integer value. Use 1 for AWS, 4 for GCP, 8 for Azure, 16 for OCI, and 256 for AWSGov.`,
				},
				resource.Attribute{
					Name:        "backup_account_name",
					Description: `(Optional) Name of the cloud account in the Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "backup_bucket_name",
					Description: `(Optional) Bucket Name. Required to enable configuration backup for AWS, AWSGov, GCP and OCI.`,
				},
				resource.Attribute{
					Name:        "backup_storage_name",
					Description: `(Optional) Storage name. Required to enable configuration backup for Azure.`,
				},
				resource.Attribute{
					Name:        "backup_container_name",
					Description: `(Optional) Container name. Required to enable configuration backup for Azure.`,
				},
				resource.Attribute{
					Name:        "backup_region",
					Description: `(Optional) Name of region. Required to enable configuration backup for Azure and OCI.`,
				},
				resource.Attribute{
					Name:        "multiple_backups",
					Description: `(Optional) Switch to enable the Controller to backup up to a maximum of 3 rotating backups. Valid values: true, false. Default value: false. ->`,
				},
				resource.Attribute{
					Name:        "ca_certificate_file_path",
					Description: `(Optional) File path to CA certificate. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "server_public_certificate_file_path",
					Description: `(Optional) File path to the server public certificate. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "server_private_key_file_path",
					Description: `(Optional) File path to server private key. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "ca_certificate_file",
					Description: `(Optional) CA certificate. To read certificate file from a file, please use the built-in ` + "`" + `file` + "`" + ` function. Available as of provider version R2.21.2+.`,
				},
				resource.Attribute{
					Name:        "server_public_certificate_file",
					Description: `(Optional) Server public certificate. To read certificate file from a file, please use the built-in ` + "`" + `file` + "`" + ` function. Available as of provider version R2.21.2+.`,
				},
				resource.Attribute{
					Name:        "server_private_key_file",
					Description: `(Optional) Server private key. To read the private key from a file, please use the built-in ` + "`" + `file` + "`" + ` function. Available as of provider version R2.21.2+. ### Misc.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC/VNET DNS Server for the controller. Valid values: true, false. Default value: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Current version of the controller without build number. Example: "6.5"`,
				},
				resource.Attribute{
					Name:        "previous_version",
					Description: `Previous version of the controller including the build number. Example: "6.5.123". Available as of provider version R2.20.0+.`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current version of the controller including the build number. Example: "6.5.123". Available as of provider version R2.20.0+. ~>`,
				},
				resource.Attribute{
					Name:        "sg_management_account_name",
					Description: `(Optional) Select the [primary access account](https://docs.aviatrix.com/HowTos/aviatrix_account.html#setup-primary-access-account-for-aws-cloud).`,
				},
				resource.Attribute{
					Name:        "security_group_management",
					Description: `(Optional) Enable to allow Controller to automatically manage inbound rules from gateways. Valid values: true, false. Default value: false. ## Import Instance controller_config can be imported using controller IP, e.g. controller IP is : 10.11.12.13 ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_controller_config.test 10-11-12-13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `Current version of the controller without build number. Example: "6.5"`,
				},
				resource.Attribute{
					Name:        "previous_version",
					Description: `Previous version of the controller including the build number. Example: "6.5.123". Available as of provider version R2.20.0+.`,
				},
				resource.Attribute{
					Name:        "current_version",
					Description: `Current version of the controller including the build number. Example: "6.5.123". Available as of provider version R2.20.0+. ~>`,
				},
				resource.Attribute{
					Name:        "sg_management_account_name",
					Description: `(Optional) Select the [primary access account](https://docs.aviatrix.com/HowTos/aviatrix_account.html#setup-primary-access-account-for-aws-cloud).`,
				},
				resource.Attribute{
					Name:        "security_group_management",
					Description: `(Optional) Enable to allow Controller to automatically manage inbound rules from gateways. Valid values: true, false. Default value: false. ## Import Instance controller_config can be imported using controller IP, e.g. controller IP is : 10.11.12.13 ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_controller_config.test 10-11-12-13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_email_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller email config resource`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"email",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_alert_email",
					Description: `(Required) Email to receive important account and certification information.`,
				},
				resource.Attribute{
					Name:        "critical_alert_email",
					Description: `(Required) Email to receive field notices and critical notices.`,
				},
				resource.Attribute{
					Name:        "security_event_email",
					Description: `(Required) Email to receive security and CVE (Common Vulnerabilities and Exposures) notification emails.`,
				},
				resource.Attribute{
					Name:        "status_change_email",
					Description: `(Required) Email to receive system/tunnel status notification emails.`,
				},
				resource.Attribute{
					Name:        "status_change_notification_interval",
					Description: `(Optional) Status change notification interval in seconds. Default value: 60. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_alert_email_verified",
					Description: `Whether admin alert notification email is verified.`,
				},
				resource.Attribute{
					Name:        "critical_alert_email_verified",
					Description: `Whether critical alert notification email is verified.`,
				},
				resource.Attribute{
					Name:        "security_event_email_verified",
					Description: `Whether security event notification email is verified.`,
				},
				resource.Attribute{
					Name:        "status_change_email_verified",
					Description: `Whether status change notification email is verified. ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_alert_email_verified",
					Description: `Whether admin alert notification email is verified.`,
				},
				resource.Attribute{
					Name:        "critical_alert_email_verified",
					Description: `Whether critical alert notification email is verified.`,
				},
				resource.Attribute{
					Name:        "security_event_email_verified",
					Description: `Whether security event notification email is verified.`,
				},
				resource.Attribute{
					Name:        "status_change_email_verified",
					Description: `Whether status change notification email is verified. ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_email_exception_notification_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller email exception notification config`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"email",
				"exception",
				"notification",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_email_exception_notification",
					Description: `(Optional) Enable exception email notification. When set to true, exception email will be sent to "exception@aviatrix.com", when set to false, exception email will be sent to controller's admin email. Valid values: true, false. Default value: true. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_gateway_keepalive_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix Controller Gateway Keepalive for gateways`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"gateway",
				"keepalive",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "keepalive_speed",
					Description: `The gateway keepalive template name. Must be one of "slow", "medium" or "fast". Visit [here](https://docs.aviatrix.com/HowTos/gateway.html#gateway-keepalives) for the complete documentation about the gateway keepalive configuration. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_private_mode_config",
			Category:         "Private Mode",
			ShortDescription: `Creates and manages an Aviatrix Controller's Private Mode config`,
			Description:      ``,
			Keywords: []string{
				"private",
				"mode",
				"controller",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_private_mode",
					Description: `(Required) Whether to enable Private Mode on an Aviatrix Controller. ### Optional`,
				},
				resource.Attribute{
					Name:        "copilot_instance_id",
					Description: `(Optional) Instance ID of a copilot instance to associate with an Aviatrix Controller in Private Mode. The copilot instance must be in the same VPC as the Aviatrix Controller. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_private_oob",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller private OOB config resource`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"private",
				"oob",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `(Optional) Switch to enable/disable Aviatrix controller private OOB. Valid values: true, false. Default value: false. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_controller_security_group_management_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller security group management config resource`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"controller",
				"security",
				"group",
				"management",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Select the [primary access account](https://docs.aviatrix.com/HowTos/aviatrix_account.html#setup-primary-access-account-for-aws-cloud).`,
				},
				resource.Attribute{
					Name:        "enable_security_group_management",
					Description: `(Required) Enable to allow Controller to automatically manage inbound rules from gateways. Valid values: true, false. ## Import Instance controller_security_group_management_config can be imported using controller IP, e.g. controller IP is : 10.11.12.13 ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_controller_security_group_management_config.test 10-11-12-13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_copilot_association",
			Category:         "Settings",
			ShortDescription: `Creates and manages a CoPilot Association`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"copilot",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "copilot_address",
					Description: `(Required) CoPilot instance IP Address or Hostname. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_copilot_security_group_management_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages CoPilot Security Group Management Configuration`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"copilot",
				"security",
				"group",
				"management",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_copilot_security_group_management",
					Description: `(Required) Switch to enable copilot security group management. Valid values: true, false. ### Optional`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Optional) Cloud type. The type of this attribute is Integer. Only support AWS, Azure and OCI. Required to enable copilot security group management.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) Aviatrix access account name. Required to enable copilot security group management.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC ID. Required to enable copilot security group management.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional) CoPilot instance ID. Required to enable copilot security group management.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region where CoPilot is deployed. Required and valid for AWS and Azure.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Zone where CoPilot is deployed. Required and valid for GCP. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_datadog_agent",
			Category:         "Settings",
			ShortDescription: `Enables and disables datadog agent`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"datadog",
				"agent",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of datadog agent. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of datadog agent. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_device_interface_config",
			Category:         "CloudN",
			ShortDescription: `Configures WAN primary interface and IP for a device.`,
			Description:      ``,
			Keywords: []string{
				"cloudn",
				"device",
				"interface",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) Name of the device.`,
				},
				resource.Attribute{
					Name:        "wan_primary_interface",
					Description: `(Required) Name of the WAN primary interface.`,
				},
				resource.Attribute{
					Name:        "wan_primary_interface_public_ip",
					Description: `(Required) The WAN Primary interface public IP. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_distributed_firewalling_config",
			Category:         "Secured Networking",
			ShortDescription: `Creates and manages an Aviatrix Distributed-firewalling Config`,
			Description:      ``,
			Keywords: []string{
				"secured",
				"networking",
				"distributed",
				"firewalling",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_distributed_firewalling",
					Description: `(Optional) Whether to enable Private Mode on an Aviatrix Controller. Valid values: true, false. Default value: false. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_distributed_firewalling_intra_vpc",
			Category:         "Secured Networking",
			ShortDescription: `Creates and manages an Aviatrix Distributed-firewalling Intra VPC List`,
			Description:      ``,
			Keywords: []string{
				"secured",
				"networking",
				"distributed",
				"firewalling",
				"intra",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpcs",
					Description: `(Required) List of distributed-firewalling enabled intra VPCs.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account Name of the VPC.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) vpc_id of the VPC.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of the VPC.. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_distributed_firewalling_policy_list",
			Category:         "Secured Networking",
			ShortDescription: `Creates and manages an Aviatrix Distributed-firewalling Policy List`,
			Description:      ``,
			Keywords: []string{
				"secured",
				"networking",
				"distributed",
				"firewalling",
				"policy",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) List of policies.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the policy.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action for the policy. Must be one of PERMIT or DENY.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority for the policy. Default: 0. Type: Integer.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol for the policy. Must be one of TCP, UDP, ICMP or ANY.`,
				},
				resource.Attribute{
					Name:        "src_smart_groups",
					Description: `(Required) List of Smart Group UUIDs for the source for the policy.`,
				},
				resource.Attribute{
					Name:        "dst_smart_groups",
					Description: `(Required) List of Smart Group UUIDs for the destination for the policy.`,
				},
				resource.Attribute{
					Name:        "web_groups",
					Description: `(Optional) List of Web Group UUIDs for the policy. Available as of provider version 3.1.0+.`,
				},
				resource.Attribute{
					Name:        "flow_app_requirement",
					Description: `(Optional) Flow application requirement for the policy. Available as of provider version 3.1.0+.`,
				},
				resource.Attribute{
					Name:        "decrypt_policy",
					Description: `(Optional) Decryption options for the policy. Available as of provider version 3.1.0+.`,
				},
				resource.Attribute{
					Name:        "port_ranges",
					Description: `(Optional) List of port ranges for the policy. Cannot be used when ` + "`" + `protocol` + "`" + ` is "ICMP".`,
				},
				resource.Attribute{
					Name:        "lo",
					Description: `(Required) Lower bound for the range of ports.`,
				},
				resource.Attribute{
					Name:        "hi",
					Description: `(Optional) Upper bound for the range of ports. When not set, ` + "`" + `lo` + "`" + ` is the only port that matches the policy.`,
				},
				resource.Attribute{
					Name:        "watch",
					Description: `(Optional) Whether to enforce the policy or only watch packets. If "true" packets are only watched. This allows you to observe if the traffic impacted by this rule causes any inadvertent issues (such as traffic being dropped). Type: Boolean.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `(Optional) Whether to enable logging for packets that match the policy. Type: Boolean.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Computed) UUID for the Policy. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_dns_profile",
			Category:         "Copilot",
			ShortDescription: `Creates Aviatrix DNS Profile`,
			Description:      ``,
			Keywords: []string{
				"copilot",
				"dns",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) DNS profile name. ### Optional`,
				},
				resource.Attribute{
					Name:        "global_dns_servers",
					Description: `(Optional) List of global DNS servers. Example: ["8.8.8.8", "8.8.3.4"].`,
				},
				resource.Attribute{
					Name:        "local_domain_names",
					Description: `(Optional) List of local domain names. Example: ["avx.internal.com", "avx.media.com"].`,
				},
				resource.Attribute{
					Name:        "lan_dns_servers",
					Description: `(Optional) List of LAN DNS servers. Example: ["1.2.3.4", "5.6.7.8"].`,
				},
				resource.Attribute{
					Name:        "wan_dns_servers",
					Description: `(Optional) List of WAN DNS servers. Example: ["2.3.4.5", "6.7.8.9"]. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_csp",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates Aviatrix Edge CSP`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"csp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Edge CSP account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Edge CSP name.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Site ID.`,
				},
				resource.Attribute{
					Name:        "project_uuid",
					Description: `(Required) Edge CSP project UUID.`,
				},
				resource.Attribute{
					Name:        "compute_node_uuid",
					Description: `(Required) Edge CSP compute node UUID.`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `(Required) Edge CSP template UUID.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type. Valid values: WAN, LAN, or MANAGEMENT.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The rate of data can be moved through the interface, requires an integer value. Unit is in Mb/s.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) Primary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "enable_vrrp",
					Description: `(Optional) Enable VRRP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "vrrp_virtual_ip",
					Description: `(Optional) VRRP virtual IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"].`,
				},
				resource.Attribute{
					Name:        "enable_management_over_private_network",
					Description: `(Optional) Switch to enable management over the private network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby",
					Description: `(Optional) Switch to enable Edge Active-Standby mode. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby_preemptive",
					Description: `(Optional) Switch to enable Preemptive Mode for Edge Active-Standby. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "wan_interface_names",
					Description: `(Optional) List of WAN interface names.`,
				},
				resource.Attribute{
					Name:        "lan_interface_names",
					Description: `(Optional) List of LAN interface names.`,
				},
				resource.Attribute{
					Name:        "management_interface_names",
					Description: `(Optional) List of management interface names.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) BGP AS Number to assign to Edge CSP.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) List of AS numbers to prepend gateway BGP AS_Path field. Valid only when ` + "`" + `local_as_number` + "`" + ` is set. Example: ["65023", "65023"].`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable learned CIDR approval. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `(Optional) Set of approved learned CIDRs. Valid only when ` + "`" + `enable_learned_cidrs_approval` + "`" + ` is set to true. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_manual_advertise_cidrs",
					Description: `(Optional) Set of intended CIDRs to be advertised to external BGP router. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "enable_preserve_as_path",
					Description: `(Optional) Switch to enable preserve as_path when advertising manual summary CIDRs. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `(Optional) BGP route polling time. Unit is in seconds. Valid values are between 10 and 50. Default value: 50.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `(Optional) BGP hold time. Unit is in seconds. Valid values are between 12 and 360. Default value: 180.`,
				},
				resource.Attribute{
					Name:        "enable_edge_transitive_routing",
					Description: `(Optional) Switch to enable Edge transitive routing. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Switch to enable jumbo frame. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude of Edge CSP. Valid values are between -90 and 90. Example: "47.7511".`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude of Edge CSP. Valid values are between -180 and 180. Example: "120.7401".`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Ethernet interface RX queue size. Once set, can't be deleted or disabled. Valid values: "1K", "2K", "4K".`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) VLAN configuration.`,
				},
				resource.Attribute{
					Name:        "parent_interface_name",
					Description: `(Required) Parent interface name.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) VLAN ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) LAN sub-interface IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) LAN sub-interface gateway IP.`,
				},
				resource.Attribute{
					Name:        "peer_ip_address",
					Description: `(Optional) LAN sub-interface IP address on HA gateway.`,
				},
				resource.Attribute{
					Name:        "peer_gateway_ip",
					Description: `(Optional) LAN sub-interface gateway IP on HA gateway.`,
				},
				resource.Attribute{
					Name:        "vrrp_virtual_ip",
					Description: `(Optional) LAN sub-interface virtual IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag.`,
				},
				resource.Attribute{
					Name:        "dns_profile_name",
					Description: `(Optional) DNS profile to be associated with gateway, select an existing template.`,
				},
				resource.Attribute{
					Name:        "enable_single_ip_snat",
					Description: `(Optional) Enable Single IP SNAT. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_auto_advertise_lan_cidrs",
					Description: `(Optional) Enable auto advertise LAN CIDRs. Valid values: true, false. Default value: true. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge CSP. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge CSP. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_csp_ha",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates Aviatrix Edge CSP HA`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"csp",
				"ha",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_gw_name",
					Description: `(Required) Primary Edge CSP name.`,
				},
				resource.Attribute{
					Name:        "compute_node_uuid",
					Description: `(Required) Edge CSP compute node UUID. ->`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The rate of data can be moved through the interface, requires an integer value. Unit is in Mb/s.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) Primary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"]. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Edge CSP account name. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Edge CSP account name. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_equinix",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates Aviatrix Edge Equinix`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"equinix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Edge Equinix account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Edge Equinix name.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Site ID.`,
				},
				resource.Attribute{
					Name:        "ztp_file_download_path",
					Description: `(Required) The folder path where the ZTP file will be downloaded.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type. Valid values: WAN, LAN, or MANAGEMENT.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The rate of data can be moved through the interface, requires an integer value. Unit is in Mb/s.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) Primary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "enable_vrrp",
					Description: `(Optional) Enable VRRP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "vrrp_virtual_ip",
					Description: `(Optional) VRRP virtual IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"].`,
				},
				resource.Attribute{
					Name:        "enable_management_over_private_network",
					Description: `(Optional) Switch to enable management over the private network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby",
					Description: `(Optional) Switch to enable Edge Active-Standby mode. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby_preemptive",
					Description: `(Optional) Switch to enable Preemptive Mode for Edge Active-Standby. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) BGP AS Number to assign to Edge Equinix.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) List of AS numbers to prepend gateway BGP AS_Path field. Valid only when ` + "`" + `local_as_number` + "`" + ` is set. Example: ["65023", "65023"].`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable learned CIDR approval. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `(Optional) Set of approved learned CIDRs. Valid only when ` + "`" + `enable_learned_cidrs_approval` + "`" + ` is set to true. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_manual_advertise_cidrs",
					Description: `(Optional) Set of intended CIDRs to be advertised to external BGP router. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "enable_preserve_as_path",
					Description: `(Optional) Switch to enable preserve as_path when advertising manual summary CIDRs. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `(Optional) BGP route polling time in seconds. Valid values are between 10 and 50. Default value: 50.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `(Optional) BGP hold time in seconds. Valid values are between 12 and 360. Default value: 180.`,
				},
				resource.Attribute{
					Name:        "enable_edge_transitive_routing",
					Description: `(Optional) Switch to enable Edge transitive routing. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Switch to enable jumbo frame. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude of Edge Equinix. Valid values are between -90 and 90. Example: "47.7511".`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude of Edge Equinix. Valid values are between -180 and 180. Example: "120.7401".`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Ethernet interface RX queue size. Once set, can't be deleted or disabled. Valid values: "1K", "2K", "4K".`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) VLAN configuration.`,
				},
				resource.Attribute{
					Name:        "parent_interface_name",
					Description: `(Required) Parent interface name.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) VLAN ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) LAN sub-interface IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) LAN sub-interface gateway IP.`,
				},
				resource.Attribute{
					Name:        "peer_ip_address",
					Description: `(Optional) LAN sub-interface IP address on HA gateway.`,
				},
				resource.Attribute{
					Name:        "peer_gateway_ip",
					Description: `(Optional) LAN sub-interface gateway IP on HA gateway.`,
				},
				resource.Attribute{
					Name:        "vrrp_virtual_ip",
					Description: `(Optional) LAN sub-interface virtual IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag.`,
				},
				resource.Attribute{
					Name:        "dns_profile_name",
					Description: `(Optional) DNS profile to be associated with gateway, select an existing template.`,
				},
				resource.Attribute{
					Name:        "enable_single_ip_snat",
					Description: `(Optional) Enable Single IP SNAT. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_auto_advertise_lan_cidrs",
					Description: `(Optional) Enable auto advertise LAN CIDRs. Valid values: true, false. Default value: true. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge Equinix. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge Equinix. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_equinix_ha",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates an Aviatrix Edge Equinix HA`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"equinix",
				"ha",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_gw_name",
					Description: `(Required) Primary Edge Equinix name.`,
				},
				resource.Attribute{
					Name:        "ztp_file_download_path",
					Description: `(Required) The folder path where the ZTP file will be downloaded.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The rate of data can be moved through the interface, requires an integer value. Unit is in Mb/s.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) Primary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"]. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Edge Equinix account name. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Edge Equinix account name. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_neo",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates Aviatrix Edge NEO`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"neo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Edge NEO account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Edge NEO name.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Site ID.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) Device ID.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Gateway size. Valid values: "small", "medium", "large", "x-large".`,
				},
				resource.Attribute{
					Name:        "wan_interface_names",
					Description: `(Required) List of WAN interface names. Default value: ["eth0"].`,
				},
				resource.Attribute{
					Name:        "lan_interface_names",
					Description: `(Required) List of LAN interface names. Default value: ["eth1"].`,
				},
				resource.Attribute{
					Name:        "management_interface_names",
					Description: `(Required) List of management interface names. Default value: ["eth2"].`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type. Valid values: WAN, LAN, or MANAGEMENT.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The rate of data can be moved through the interface, requires an integer value. Unit is in Mb/s.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) Primary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "enable_vrrp",
					Description: `(Optional) Enable VRRP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "vrrp_virtual_ip",
					Description: `(Optional) VRRP virtual IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"].`,
				},
				resource.Attribute{
					Name:        "enable_management_over_private_network",
					Description: `(Optional) Switch to enable management over the private network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby",
					Description: `(Optional) Switch to enable Edge Active-Standby mode. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby_preemptive",
					Description: `(Optional) Switch to enable Preemptive Mode for Edge Active-Standby. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "wan_interface_names",
					Description: `(Optional) List of WAN interface names. Default value: ["eth0"].`,
				},
				resource.Attribute{
					Name:        "lan_interface_names",
					Description: `(Optional) List of LAN interface names. Default value: ["eth1"].`,
				},
				resource.Attribute{
					Name:        "management_interface_names",
					Description: `(Optional) List of management interface names. Default value: ["eth2"].`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) BGP AS Number to assign to Edge NEO.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) List of AS numbers to prepend gateway BGP AS_Path field. Valid only when ` + "`" + `local_as_number` + "`" + ` is set. Example: ["65023", "65023"].`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable learned CIDR approval. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `(Optional) Set of approved learned CIDRs. Valid only when ` + "`" + `enable_learned_cidrs_approval` + "`" + ` is set to true. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_manual_advertise_cidrs",
					Description: `(Optional) Set of intended CIDRs to be advertised to external BGP router. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "enable_preserve_as_path",
					Description: `(Optional) Switch to enable preserve as_path when advertising manual summary CIDRs. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `(Optional) BGP route polling time in seconds. Valid values are between 10 and 50. Default value: 50.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `(Optional) BGP hold time in seconds. Valid values are between 12 and 360. Default value: 180.`,
				},
				resource.Attribute{
					Name:        "enable_edge_transitive_routing",
					Description: `(Optional) Switch to enable Edge transitive routing. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Switch to enable jumbo frame. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude of Edge NEO. Valid values are between -90 and 90. Example: "47.7511".`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude of Edge Neo. Valid values are between -180 and 180. Example: "120.7401".`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Ethernet interface RX queue size. Once set, can't be deleted or disabled. Valid values: "1K", "2K", "4K".`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Required) VLAN configuration.`,
				},
				resource.Attribute{
					Name:        "parent_interface_name",
					Description: `(Required) Parent interface name.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Required) VLAN ID.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) LAN sub-interface IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) LAN sub-interface gateway IP.`,
				},
				resource.Attribute{
					Name:        "peer_ip_address",
					Description: `(Optional) LAN sub-interface IP address on HA gateway.`,
				},
				resource.Attribute{
					Name:        "peer_gateway_ip",
					Description: `(Optional) LAN sub-interface gateway IP on HA gateway.`,
				},
				resource.Attribute{
					Name:        "vrrp_virtual_ip",
					Description: `(Optional) LAN sub-interface virtual IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag.`,
				},
				resource.Attribute{
					Name:        "dns_profile_name",
					Description: `(Optional) DNS profile to be associated with gateway, select an existing template.`,
				},
				resource.Attribute{
					Name:        "enable_single_ip_snat",
					Description: `(Optional) Enable Single IP SNAT. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_auto_advertise_lan_cidrs",
					Description: `(Optional) Enable auto advertise LAN CIDRs. Valid values: true, false. Default value: true. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge NEO. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge NEO. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_neo_device_onboarding",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Onboards an Edge NEO device`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"neo",
				"device",
				"onboarding",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Edge NEO account name.`,
				},
				resource.Attribute{
					Name:        "device_name",
					Description: `(Required) Edge NEO device name.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `(Required) Device serial number.`,
				},
				resource.Attribute{
					Name:        "hardware_model",
					Description: `(Required) Device hardware model. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"].`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Network configurations.`,
				},
				resource.Attribute{
					Name:        "interface_name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Set to true to enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "ipv4_cidr",
					Description: `(Optional) Interface IPV4 CIDR.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ips",
					Description: `(Optional) Set of DNS server IPs.`,
				},
				resource.Attribute{
					Name:        "proxy_server_ip",
					Description: `(Optional) Proxy server IP.`,
				},
				resource.Attribute{
					Name:        "download_config_file",
					Description: `(Optional) Set to true to download the Edge NEO static config file. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "config_file_download_path",
					Description: `(Optional) The location where the config file will be stored. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `Edge NEO device ID. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "device_id",
					Description: `Edge NEO device ID. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_neo_ha",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates an Aviatrix Edge NEO HA`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"neo",
				"ha",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_gw_name",
					Description: `(Required) Primary Edge NEO name.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Required) Device ID.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) The rate of data can be moved through the interface, requires an integer value. Unit is in Mb/s.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) Primary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Tag. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix. Example: ["67.207.104.16/29", "64.71.12.144/29"]. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Edge NEO account name. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Edge NEO account name. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_spoke",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates Aviatrix Edge as a Spoke`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"spoke",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Edge as a Spoke name.`,
				},
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Site ID.`,
				},
				resource.Attribute{
					Name:        "ztp_file_type",
					Description: `(Required) ZTP file type. Valid values: "iso", "cloud-init".`,
				},
				resource.Attribute{
					Name:        "ztp_file_download_path",
					Description: `(Required) The folder path where the ZTP file will be downloaded.`,
				},
				resource.Attribute{
					Name:        "interfaces",
					Description: `(Required) WAN/LAN/MANAGEMENT interfaces.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Interface name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type. Valid values: WAN, LAN, or MANAGEMENT.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `(Optional) Enable DHCP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "wan_public_ip",
					Description: `(Optional) WAN public IP.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Interface static IP address.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) Gateway IP. ### Optional`,
				},
				resource.Attribute{
					Name:        "management_egress_ip_prefix_list",
					Description: `(Optional) Set of management egress gateway IP and subnet prefix.`,
				},
				resource.Attribute{
					Name:        "enable_management_over_private_network",
					Description: `(Optional) Switch to enable management over the private network. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby",
					Description: `(Optional) Switch to enable Edge Active-Standby mode. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_edge_active_standby_preemptive",
					Description: `(Optional) Switch to enable Preemptive Mode for Edge Active-Standby. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dns_server_ip",
					Description: `(Optional) DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static".`,
				},
				resource.Attribute{
					Name:        "secondary_dns_server_ip",
					Description: `(Optional) Secondary DNS server IP. Required and valid when ` + "`" + `management_interface_config` + "`" + ` is "Static". ### Advanced Options`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) BGP AS Number to assign to Edge as a Spoke.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) List of AS numbers to prepend gateway BGP AS_Path field. Valid only when ` + "`" + `local_as_number` + "`" + ` is set. Example: ["65023", "65023"].`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable learned CIDR approval. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `(Optional) Set of approved learned CIDRs. Valid only when ` + "`" + `enable_learned_cidrs_approval` + "`" + ` is set to true. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_manual_advertise_cidrs",
					Description: `(Optional) Set of intended CIDRs to be advertised to external BGP router. Example: ["10.1.0.0/116", "10.2.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "enable_preserve_as_path",
					Description: `(Optional) Switch to enable preserve as_path when advertising manual summary CIDRs. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `(Optional) BGP route polling time. Unit is in seconds. Valid values are between 10 and 50. Default value: 50.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `(Optional) BGP hold time. Unit is in seconds. Valid values are between 12 and 360. Default value: 180.`,
				},
				resource.Attribute{
					Name:        "enable_edge_transitive_routing",
					Description: `(Optional) Switch to enable Edge transitive routing. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Switch to enable jumbo frame. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `(Optional) Latitude of Edge as a Spoke. Valid values are between -90 and 90. Example: "47.7511".`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `(Optional) Longitude of Edge as a Spoke. Valid values are between -180 and 180. Example: "120.7401".`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Ethernet interface RX queue size. Once set, can't be deleted or disabled. Valid values: "1K", "2K", "4K". ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge as a Spoke. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `State of Edge as a Spoke. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_spoke_external_device_conn",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Edge as a Spoke external device connections`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"spoke",
				"external",
				"device",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "site_id",
					Description: `(Required) Edge as a Spoke site iD.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Edge as a Spoke name.`,
				},
				resource.Attribute{
					Name:        "bgp_local_as_num",
					Description: `(Required) BGP local AS number.`,
				},
				resource.Attribute{
					Name:        "bgp_remote_as_num",
					Description: `(Required) BGP remote AS number.`,
				},
				resource.Attribute{
					Name:        "local_lan_ip",
					Description: `(Required) Local LAN IP.`,
				},
				resource.Attribute{
					Name:        "remote_lan_ip",
					Description: `(Required) Remote LAN IP. ### Optional ->`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Optional) Connection name.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Optional) Connection type. Valid value: 'bgp'. Default value: 'bgp'.`,
				},
				resource.Attribute{
					Name:        "tunnel_protocol",
					Description: `(Optional) Tunnel protocol. Valid value: 'LAN'. Default value: 'LAN'. Case insensitive.`,
				},
				resource.Attribute{
					Name:        "enable_edge_underlay",
					Description: `(Optional) Enable BGP over WAN underlay. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "remote_cloud_type",
					Description: `(Optional) Remote cloud type. Valid values: "AWS", "AZURE".`,
				},
				resource.Attribute{
					Name:        "number_of_retries",
					Description: `(Optional) Number of retries. Default value: 0.`,
				},
				resource.Attribute{
					Name:        "retry_interval",
					Description: `(Optional) Retry interval in seconds. Default value: 300.`,
				},
				resource.Attribute{
					Name:        "ha_enabled",
					Description: `(Optional) Set as true if there are two external devices.`,
				},
				resource.Attribute{
					Name:        "backup_remote_lan_ip",
					Description: `(Optional) Backup Remote LAN IP. Required for HA BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "backup_local_lan_ip",
					Description: `(Optional) Backup Local LAN IP. Required for HA BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "backup_bgp_remote_as_num",
					Description: `(Optional) Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required if HA enabled for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_edge_spoke_transit_attachment",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix Edge as a Spoke to Transit attachments`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"edge",
				"spoke",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spoke_gw_name",
					Description: `(Required) Name of the Edge as a Spoke to attach to transit network.`,
				},
				resource.Attribute{
					Name:        "transit_gw_name",
					Description: `(Required) Name of the transit gateway to attach the Edge as a Spoke to. ### Options`,
				},
				resource.Attribute{
					Name:        "enable_over_private_network",
					Description: `(Optional) Switch to enable over the private network. Valid values: true, false. Default: true.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Switch to enable jumbo frame. Valid values: true, false. Default: false.`,
				},
				resource.Attribute{
					Name:        "enable_insane_mode",
					Description: `(Optional) Switch to enable insane mode. Valid values: true, false. Default: false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_tunnel_number",
					Description: `(Optional) Insane mode tunnel number, requires an integer value. Valid Range: 2-4.`,
				},
				resource.Attribute{
					Name:        "spoke_prepend_as_path",
					Description: `(Optional) Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Edge as a Spoke.`,
				},
				resource.Attribute{
					Name:        "transit_prepend_as_path",
					Description: `(Optional) Connection based AS Path Prepend. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on the Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "number_of_retries",
					Description: `(Optional) Number of retries. Default value: 0.`,
				},
				resource.Attribute{
					Name:        "retry_interval",
					Description: `(Optional) Retry interval in seconds. Default value: 300. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_filebeat_forwarder",
			Category:         "Settings",
			ShortDescription: `Enables and disables filebeat forwarder`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"filebeat",
				"forwarder",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of filebeat forwarder. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of filebeat forwarder. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firenet",
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
					Description: `(Required) VPC ID of the Security VPC.`,
				},
				resource.Attribute{
					Name:        "inspection_enabled",
					Description: `(Optional) Enable/disable traffic inspection. Valid values: true, false. Default value: true. ->`,
				},
				resource.Attribute{
					Name:        "inspection_enabled",
					Description: `Default value is true for associating firewall instance to FireNet. Only false is supported for associating FQDN gateway to FireNet.`,
				},
				resource.Attribute{
					Name:        "egress_enabled",
					Description: `(Optional) Enable/disable egress through firewall. Valid values: true, false. Default value: false. ->`,
				},
				resource.Attribute{
					Name:        "egress_enabled",
					Description: `Default value is false for associating firewall instance to FireNet. Only true is supported for associating FQDN gateway to FireNet.`,
				},
				resource.Attribute{
					Name:        "egress_static_cidrs",
					Description: `(Optional) List of egress static CIDRs. Egress is required to be enabled. Example: ["1.171.15.184/32", "1.171.15.185/32"]. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "east_west_inspection_excluded_cidrs",
					Description: `(Optional) Network List Excluded From East-West Inspection. CIDRs to be excluded from inspection. Type: Set(String). Available as of provider version R2.19.5+.`,
				},
				resource.Attribute{
					Name:        "tgw_segmentation_for_egress_enabled",
					Description: `(Optional) Enable TGW segmentation for egress. Valid values: true or false. Default value: false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "hashing_algorithm",
					Description: `(Optional) Hashing algorithm to load balance traffic across the firewall. Valid values: "2-Tuple", "5-Tuple". Default value: "5-Tuple".`,
				},
				resource.Attribute{
					Name:        "keep_alive_via_lan_interface_enabled",
					Description: `(Optional) Enable Keep Alive via Firewall LAN Interface. Valid values: true or false. Default value: false. Available as of provider version R2.18.1+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "fail_close_enabled",
					Description: `Enable Fail Close. Removed as of provider version R2.23+.`,
				},
				resource.Attribute{
					Name:        "manage_firewall_instance_association",
					Description: `(Optional) Enable this attribute to manage firewall associations in-line. If set to true, in-line ` + "`" + `firewall_instance_association` + "`" + ` blocks can be used. If set to false, all firewall associations must be managed via standalone ` + "`" + `aviatrix_firewall_instance_association` + "`" + ` resources. Default value: true. Valid values: true or false. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "firewall_instance_association",
					Description: `(Optional) Dynamic block of firewall instance(s) to be associated with the FireNet.`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `(Required) Name of the primary FireNet gateway.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Firewall instance.`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `(Optional) Type of firewall. Valid values: "Generic", "fqdn_gateway". Default value: "Generic". Value "fqdn_gateway" is required for FQDN gateway.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `(Optional) Firewall instance name.`,
				},
				resource.Attribute{
					Name:        "management_interface",
					Description: `(Optional) Management interface ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall",
			Category:         "Security",
			ShortDescription: `Creates and manages Aviatrix Stateful Firewall Policies`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Gateway name to attach firewall policy to.`,
				},
				resource.Attribute{
					Name:        "base_policy",
					Description: `(Optional) New base policy. Valid Values: "allow-all", "deny-all". Default value: "deny-all"`,
				},
				resource.Attribute{
					Name:        "base_log_enabled",
					Description: `(Optional) Indicates whether enable logging or not. Valid Values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "manage_firewall_policies",
					Description: `(Optional) Enable to manage firewall policies via in-line rules. If false, policies must be managed using ` + "`" + `aviatrix_firewall_policy` + "`" + ` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) New access policy for the gateway. Seven fields are required for each policy item: ` + "`" + `src_ip` + "`" + `, ` + "`" + `dst_ip` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `port` + "`" + `, ` + "`" + `action` + "`" + `, ` + "`" + `log_enabled` + "`" + ` and ` + "`" + `description` + "`" + `. No duplicate rules (with same ` + "`" + `src_ip` + "`" + `, ` + "`" + `dst_ip` + "`" + `, ` + "`" + `protocol` + "`" + ` and ` + "`" + `port` + "`" + `) are allowed.`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Required) Source address, a valid IPv4 address or tag name such "HR" or "marketing" etc. Example: "10.30.0.0/16". The`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Required) Destination address, a valid IPv4 address or tag name such "HR" or "marketing" etc. Example: "10.30.0.0/16". The`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) A single port or a range of port numbers. Example: "25", "25:1024".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall_instance",
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
					Description: `(Required) VPC ID of the Security VPC. For GCP, ` + "`" + `vpc_id` + "`" + ` must be in the form vpc_id~-~gcloud_project_id.`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `(Optional) Name of the primary FireNet gateway.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `(Required) Name of the firewall instance to be created.`,
				},
				resource.Attribute{
					Name:        "firewall_image",
					Description: `(Required) One of the AWS/Azure/GCP AMIs from various vendors such as Palo Alto Networks.`,
				},
				resource.Attribute{
					Name:        "firewall_image_id",
					Description: `(Optional) Firewall image ID. Only applicable to AWS. Please use AMI ID. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "firewall_size",
					Description: `(Required) Instance size of the firewall. Example: "m5.xlarge".`,
				},
				resource.Attribute{
					Name:        "management_subnet",
					Description: `(Optional) Management Interface Subnet. Select the subnet whose name contains “gateway and firewall management”. For GCP, ` + "`" + `management_subnet` + "`" + ` must be in the form ` + "`" + `cidr~~region~~name` + "`" + `. Required for Palo Alto Networks VM-Series and OCI Check Point firewalls. Otherwise, it must be empty.`,
				},
				resource.Attribute{
					Name:        "egress_subnet",
					Description: `(Required) Egress Interface Subnet. Select the subnet whose name contains “FW-ingress-egress”. For GCP, ` + "`" + `egress_subnet` + "`" + ` must be in the form ` + "`" + `cidr~~region~~name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "firewall_image_version",
					Description: `(Optional) Version of firewall image. If not specified, Controller will automatically select the latest version available.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability Zone. Required if creating a Firewall Instance with a Native AWS GWLB-enabled VPC. Applicable to AWS, Azure, and GCP only. Available as of provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "management_vpc_id",
					Description: `(Optional) Management VPC ID. Only used for GCP firewall. Required for Palo Alto Networks VM-Series, and required to be empty for Check Point or Fortinet series. Available as of provider version R2.18.1+.`,
				},
				resource.Attribute{
					Name:        "egress_vpc_id",
					Description: `(Optional) Egress VPC ID. Required for GCP. Available as of provider version R2.18.1+.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `(Optional) Fault domain. Required and valid only for OCI. Available as of provider version R2.19.3. ->`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Optional) Applicable to Azure or AzureGov deployment only. ### Advanced Options`,
				},
				resource.Attribute{
					Name:        "iam_role",
					Description: `(Optional) Only available for AWS. In advanced mode, create an IAM Role on the AWS account that launched the FireNet gateway. Create a policy to attach to the role. The policy is to allow access to "Bootstrap Bucket".`,
				},
				resource.Attribute{
					Name:        "bootstrap_storage_name",
					Description: `(Optional) Advanced option. Bootstrap storage name. Applicable to Azure or AzureGov and Palo Alto Networks VM-Series/Fortinet Series deployment only. Available as of provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "storage_access_key",
					Description: `(Optional) Advanced option. Storage access key. Applicable to Azure or AzureGov and Palo Alto Networks VM-Series deployment only. Available as of provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "file_share_folder",
					Description: `(Optional) Advanced option. File share folder. Applicable to Azure or AzureGov and Palo Alto Networks VM-Series deployment only. Available as of provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "share_directory",
					Description: `(Optional) Advanced option. Share directory. Applicable to Azure or AzureGov and Palo Alto Networks VM-Series deployment only. Available as of provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "sic_key",
					Description: `(Optional) Advanced option. Sic key. Applicable to Check Point Series deployment only.`,
				},
				resource.Attribute{
					Name:        "container_folder",
					Description: `(Optional) Advanced option. Container folder. Applicable to Azure or AzureGov and Fortinet Series deployment only.`,
				},
				resource.Attribute{
					Name:        "sas_url_config",
					Description: `(Optional) Advanced option. SAS URL Config. Applicable to Azure or AzureGov and Fortinet Series deployment only.`,
				},
				resource.Attribute{
					Name:        "sas_url_license",
					Description: `(Optional) Advanced option. SAS URL License. Applicable to Azure or AzureGov and Fortinet Series deployment only.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Advanced option. User Data. Applicable to Check Point Series and Fortinet Series deployment only. Type: String. ### Misc.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Mapping of key value pairs of tags for a firewall instance. Only available for AWS, AWSGov, GCP and Azure firewall instances. For AWS, AWSGov and Azure allowed characters are: letters, spaces, and numbers plus the following special characters: + - = . _ : @. For GCP allowed characters are: lowercase letters, numbers, "-" and "_". Example: {"key1" = "value1", "key2" = "value2"}. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud Type.`,
				},
				resource.Attribute{
					Name:        "gcp_vpc_id",
					Description: `GCP Only. The current VPC ID. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Cloud Type.`,
				},
				resource.Attribute{
					Name:        "gcp_vpc_id",
					Description: `GCP Only. The current VPC ID. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall_instance_association",
			Category:         "Firewall Network",
			ShortDescription: `Create and manage a single firewall instance association`,
			Description:      ``,
			Keywords: []string{
				"firewall",
				"network",
				"instance",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the Security VPC. ->`,
				},
				resource.Attribute{
					Name:        "firenet_gw_name",
					Description: `(Optional) Name of the primary FireNet gateway. Required for FireNet without Native GWLB VPC.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) ID of Firewall instance. ->`,
				},
				resource.Attribute{
					Name:        "vendor_type",
					Description: `(Optional) Type of firewall. Valid values: "Generic", "fqdn_gateway". Default value: "Generic". Value "fqdn_gateway" is required for FQDN gateway.`,
				},
				resource.Attribute{
					Name:        "firewall_name",
					Description: `(Optional) Firewall instance name.`,
				},
				resource.Attribute{
					Name:        "management_interface",
					Description: `(Optional) Management interface ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall_management_access",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Manage Aviatrix Firewall Management Access`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"firewall",
				"management",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transit_firenet_gateway_name",
					Description: `(Required) Name of the Transit FireNet-enabled transit gateway. Currently supports AWS(1) and Azure(8) providers.`,
				},
				resource.Attribute{
					Name:        "management_access_resource_name",
					Description: `(Required) Name of the resource to enable Firewall Management Access. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall_policy",
			Category:         "Security",
			ShortDescription: `Manages Aviatrix Stateful Firewall Policies`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Gateway name to attach firewall policy to.`,
				},
				resource.Attribute{
					Name:        "src_ip",
					Description: `(Required) CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The`,
				},
				resource.Attribute{
					Name:        "dst_ip",
					Description: `(Required) CIDRs separated by comma or tag names such "HR" or "marketing" etc. Example: "10.30.0.0/16,10.45.0.0/20". The`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) A single port or a range of port numbers. Example: "25", "25:1024".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_firewall_tag",
			Category:         "Security",
			ShortDescription: `Creates and manages Aviatrix Stateful Firewall Tags`,
			Description:      ``,
			Keywords: []string{
				"security",
				"firewall",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "firewall_tag",
					Description: `(Required) Name of the stateful firewall tag to be created. ### Tag Rules`,
				},
				resource.Attribute{
					Name:        "cidr_list",
					Description: `(Optional) Dynamic block representing a CIDR to filter, and a name to identify it:`,
				},
				resource.Attribute{
					Name:        "cidr_tag_name",
					Description: `(Required) A name to identify the CIDR. Example: "policy1".`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) CIDR address to filter. Example: "10.88.88.88/32". ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_fqdn",
			Category:         "Security",
			ShortDescription: `Manages Aviatrix FQDN filtering for gateways`,
			Description:      ``,
			Keywords: []string{
				"security",
				"fqdn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fqdn_tag",
					Description: `(Required) FQDN Filter tag name.`,
				},
				resource.Attribute{
					Name:        "fqdn_enabled",
					Description: `(Optional) FQDN Filter tag status. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "fqdn_mode",
					Description: `(Optional) Specify FQDN mode: whitelist or blacklist. Valid values: "white", "black".`,
				},
				resource.Attribute{
					Name:        "manage_domain_names",
					Description: `(Optional) Enable to manage domain name rules in-line. If false, domain name rules must be managed using ` + "`" + `aviatrix_fqdn_tag_rule` + "`" + ` resources. Default: true. Valid values: true, false. Available in provider version R2.17+.`,
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
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base Policy' if no value provided.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_fqdn_global_config",
			Category:         "Security",
			ShortDescription: `Manages Aviatrix FQDN Global Config`,
			Description:      ``,
			Keywords: []string{
				"security",
				"fqdn",
				"global",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "enable_exception_rule",
					Description: `(Optional) If enabled, it allows packets passing through the gateway without an SNI field. Only applies to whitelist. Valid values: true, false. Default value: true.`,
				},
				resource.Attribute{
					Name:        "enable_private_network_filtering",
					Description: `(Optional) If enabled, destination FQDN names that translate to private IP address range (RFC 1918) are subject to FQDN whitelist filtering function. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_custom_network_filtering",
					Description: `(Optional) If enabled, it customizes packet destination address ranges not to be filtered by FQDN. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "configured_ips",
					Description: `(Optional) Customized packet destination address ranges not to be filtered by FQDN. Can be selected from pre-defined RFC 1918 range, or own network range. Required with ` + "`" + `enable_custom_network_filtering` + "`" + ` enabled.`,
				},
				resource.Attribute{
					Name:        "enable_caching",
					Description: `(Optional) If enabled, it caches the resolved IP address from FQDN filter. Valid values: true, false. Default value: true.`,
				},
				resource.Attribute{
					Name:        "enable_exact_match",
					Description: `(Optional) If enabled, the resolved IP address from FQDN filter is cached so that if subsequent TCP session matches the cached IP address list, FQDN domain name is not checked and the session is allowed to pass. Valid values: true, false. Default value: false. ## Import Instance fqdn_global_config can be imported using controller IP, e.g. controller IP is : 10.11.12.13 ` + "`" + `` + "`" + `` + "`" + ` $ terraform import aviatrix_fqdn_global_config.test 10-11-12-13 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_fqdn_pass_through",
			Category:         "Security",
			ShortDescription: `Manages Aviatrix FQDN filter pass-through`,
			Description:      ``,
			Keywords: []string{
				"security",
				"fqdn",
				"pass",
				"through",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Gateway name to apply pass-through rules to.`,
				},
				resource.Attribute{
					Name:        "pass_through_cidrs",
					Description: `(Required) List of origin CIDR's to allow to pass-through FQDN filtering rules. Minimum list length: 1. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_fqdn_tag_rule",
			Category:         "Security",
			ShortDescription: `Manages Aviatrix FQDN filtering domain name rule`,
			Description:      ``,
			Keywords: []string{
				"security",
				"fqdn",
				"tag",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fqdn_tag_name",
					Description: `(Required) FQDN Filter tag name.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) FQDN. Example: "facebook.com".`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Protocol. Valid values: "all", "tcp", "udp", "icmp".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Port. Example "25".`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) What action should happen to matching requests. Possible values are: 'Base Policy', 'Allow' or 'Deny'. Defaults to 'Base Policy' if no value provided.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_gateway",
			Category:         "Gateway",
			ShortDescription: `Creates and manages Aviatrix gateways`,
			Description:      ``,
			Keywords: []string{
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Cloud service provider to use to launch the gateway. Requires an integer value. Currently supports AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud (8192), AWS Top Secret (16384) and AWS Secret (32768).`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Account name. This account will be used to launch Aviatrix gateway.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the Aviatrix gateway to be created. !> When creating a Gateway with an Azure VNet created in Controller version 6.4 or earlier or with an Azure VNet created out of band, referencing ` + "`" + `vpc_id` + "`" + ` in another resource on the same apply that creates this Gateway will cause Terraform to throw an error. Please use the Gateway data source to reference the ` + "`" + `vpc_id` + "`" + ` of this Gateway in other resources. ~> As of Provider version R2.21.2+, the ` + "`" + `vpc_id` + "`" + ` of an OCI VCN has been changed from its name to its OCID. !> As of Provider version R.22.0+, the ` + "`" + `vpc_id` + "`" + ` of a GCP VPC has been updated to include the project ID, e.g. vpc_name~-~project_id. When creating a Gateway using the old format, referencing ` + "`" + `vpc_id` + "`" + ` in another resource on the same apply that creates this Gateway will cause Terraform to throw an error. Please use the Gateway data source to reference the ` + "`" + `vpc_id` + "`" + ` of this Transit Gateway in other resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID/VNet name of cloud provider. Example: AWS/AWSGov/AWSChina: "vpc-abcd1234", GCP: "vpc-gcp-test~-~project-id", Azure/AzureGov/AzureChina: "vnet_name:rg_name:resource_guid", OCI: "ocid1.vcn.oc1.iad.aaaaaaaaba3pv6wkcr4jqae5f44n2b2m2yt2j6rx32uzr4h25vqstifsfdsq".`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) VPC region the gateway will be created in. Example: AWS: "us-east-1", GCP: "us-west2-a", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-west-1", AWSChina: "cn-north-1", AzureChina: "China North", AWS Top Secret: "us-iso-east-1", AWS Secret: "us-isob-east-1".`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Size of the gateway instance. Example: AWS/AWSGov/AWSChina: "t2.large", GCP: "n1-standard-1", Azure/AzureGov/AzureChina: "Standard_B1s", OCI: "VM.Standard2.2".`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `(Optional) Fault domain. Required and valid only for OCI. Available as of provider version R2.19.3. ### HA`,
				},
				resource.Attribute{
					Name:        "peering_ha_subnet",
					Description: `(Optional) Public subnet CIDR to create Peering HA Gateway in. Required if enabling Peering HA for AWS/AWSGov/AWS Top Secret/AWS Secret/Azure/AzureGov/Alibaba Cloud. Optional if enabling Peering HA for GCP. Example: AWS: "10.0.0.0/16".`,
				},
				resource.Attribute{
					Name:        "peering_ha_zone",
					Description: `(Optional) Zone to create Peering HA Gateway in. Required if enabling Peering HA for GCP. Example: GCP: "us-west1-c". Optional for Azure. Valid values for Azure gateways are in the form "az-n". Example: "az-2". Available for Azure as of provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "peering_ha_insane_mode_az",
					Description: `(Optional) Region + Availability Zone of subnet being created for Insane Mode-enabled Peering HA Gateway. Required for AWS only if ` + "`" + `insane_mode` + "`" + ` is set and ` + "`" + `peering_ha_subnet` + "`" + ` is set. Example: AWS: "us-west-1a".`,
				},
				resource.Attribute{
					Name:        "peering_ha_eip",
					Description: `(Optional) Public IP address to be assigned to the HA peering instance. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.`,
				},
				resource.Attribute{
					Name:        "peering_ha_azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP address resource and its resource group in Azure to be assigned to the HA peering instance. Example: "IP_Name:Resource_Group_Name". Required if ` + "`" + `peering_ha_eip` + "`" + ` is set and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_size",
					Description: `(Optional) Size of the Peering HA Gateway to be created. Required if enabling Peering HA.`,
				},
				resource.Attribute{
					Name:        "peering_ha_availability_domain",
					Description: `(Optional) Peering HA gateway availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "peering_ha_fault_domain",
					Description: `(Optional) Peering HA gateway fault domain. Required and valid only for OCI. Available as of provider version R2.19.3. ### Insane Mode`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Enable [Insane Mode](https://docs.aviatrix.com/HowTos/insane_mode.html) for Gateway. Insane Mode gateway size must be at least c5 series (AWS) or Standard_D3_v2 (Azure/AzureGov). If enabled, a valid /26 CIDR segment of the VPC must be specified to create a new subnet. Only supported for AWS, AWSGov, Azure, AzureGov, AWS China, Azure China, AWS Top Secret or AWS Secret. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) Region + Availability Zone of subnet being created for Insane Mode gateway. Required for AWS, AWSGov, AWS China, AWS Top Secret or AWS Secret if ` + "`" + `insane_mode` + "`" + ` is set. Example: AWS: "us-west-1a". ### SNAT/DNAT`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `(Optional) Enable Source NAT in "single ip" mode for this gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "vpn_access",
					Description: `(Optional) Enable [user access through VPN](https://docs.aviatrix.com/HowTos/gateway.html#vpn-access) to this gateway. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "vpn_cidr",
					Description: `(Optional) VPN CIDR block for the gateway. Required if ` + "`" + `vpn_access` + "`" + ` is true. Example: "192.168.43.0/24".`,
				},
				resource.Attribute{
					Name:        "max_vpn_conn",
					Description: `(Optional) Maximum number of active VPN users allowed to be connected to this gateway. Required if ` + "`" + `vpn_access` + "`" + ` is true. Make sure the number is smaller than the VPN CIDR block. Example: 100.`,
				},
				resource.Attribute{
					Name:        "enable_elb",
					Description: `(Optional) Specify whether to enable ELB or not. Not supported for OCI gateways. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "elb_name",
					Description: `(Optional) A name for the ELB that is created. If it is not specified, a name is generated automatically.`,
				},
				resource.Attribute{
					Name:        "vpn_protocol",
					Description: `(Optional) Transport mode for VPN connection. All ` + "`" + `cloud_types` + "`" + ` support TCP with ELB, and UDP without ELB. AWS(1) additionally supports UDP with ELB. Valid values: "TCP", "UDP". If not specified, "TCP" will be used. #### Split Tunnel`,
				},
				resource.Attribute{
					Name:        "split_tunnel",
					Description: `(Optional) Enable/disable Split Tunnel Mode. Valid values: true, false. Default value: true. Please see [here](https://docs.aviatrix.com/HowTos/gateway.html#split-tunnel-mode) for more information on split tunnel.`,
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
					Description: `(Optional) A list of destination CIDR ranges that will also go through the VPN tunnel when Split Tunnel Mode is enabled. #### MFA Authentication`,
				},
				resource.Attribute{
					Name:        "otp_mode",
					Description: `(Optional) Two step authentication mode. Valid values: "2" for DUO, "3" for Okta.`,
				},
				resource.Attribute{
					Name:        "saml_enabled",
					Description: `(Optional) Enable/disable SAML. This field is available in Controller version 3.3 or later release. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_vpn_nat",
					Description: `(Optional) Enable/disable VPN NAT. Only supported for VPN gateway. Valid values: true, false. Default value: true.`,
				},
				resource.Attribute{
					Name:        "okta_token",
					Description: `(Optional) Token for Okta auth mode. Required if ` + "`" + `otp_mode` + "`" + ` is "3".`,
				},
				resource.Attribute{
					Name:        "okta_url",
					Description: `(Optional) URL for Okta auth mode. Required if ` + "`" + `otp_mode` + "`" + ` is "3".`,
				},
				resource.Attribute{
					Name:        "okta_username_suffix",
					Description: `(Optional) Username suffix for Okta auth mode. Example: "aviatrix.com".`,
				},
				resource.Attribute{
					Name:        "duo_integration_key",
					Description: `(Optional) Integration key for DUO auth mode. Required if ` + "`" + `otp_mode` + "`" + ` is "2".`,
				},
				resource.Attribute{
					Name:        "duo_secret_key",
					Description: `(Optional) Secret key for DUO auth mode. Required if ` + "`" + `otp_mode` + "`" + ` is "2".`,
				},
				resource.Attribute{
					Name:        "duo_api_hostname",
					Description: `(Optional) API hostname for DUO auth mode. Required: Yes if ` + "`" + `otp_mode` + "`" + ` is "2".`,
				},
				resource.Attribute{
					Name:        "duo_push_mode",
					Description: `(Optional) Push mode for DUO auth. Required if ` + "`" + `otp_mode` + "`" + ` is "2". Valid values: "auto", "selective" and "token".`,
				},
				resource.Attribute{
					Name:        "enable_ldap",
					Description: `(Optional) Enable/disable LDAP. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "ldap_server",
					Description: `(Optional) LDAP server address. Required if ` + "`" + `enable_ldap` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "ldap_bind_dn",
					Description: `(Optional) LDAP bind DN. Required if ` + "`" + `enable_ldap` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "ldap_password",
					Description: `(Optional) LDAP password. Required if ` + "`" + `enable_ldap` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "ldap_base_dn",
					Description: `(Optional) LDAP base DN. Required if ` + "`" + `enable_ldap` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "ldap_username_attribute",
					Description: `(Optional) LDAP user attribute. Required if ` + "`" + `enable_ldap` + "`" + ` is true. #### Modify VPN Configuration`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `(Optional) It sets the value (seconds) of the [idle timeout](https://docs.aviatrix.com/HowTos/openvpn_faq.html#how-do-i-fix-the-aviatrix-vpn-timing-out-too-quickly). This idle timeout feature is enable only if this attribute is set, otherwise it is disabled. The entered value must be an integer number greater than 300. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "renegotiation_interval",
					Description: `(Optional) It sets the value (seconds) of the [renegotiation interval](https://docs.aviatrix.com/HowTos/openvpn_faq.html#how-do-i-fix-the-aviatrix-vpn-timing-out-too-quickly). This renegotiation interval feature is enable only if this attribute is set, otherwise it is disabled. The entered value must be an integer number greater than 300. Available in provider version R2.17.1+. ### Designated Gateway`,
				},
				resource.Attribute{
					Name:        "enable_designated_gateway",
					Description: `(Optional) Enable Designated Gateway feature for Gateway. Only supported for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false. Please view documentation [here](https://docs.aviatrix.com/HowTos/gateway.html#designated-gateway) for more information on this feature.`,
				},
				resource.Attribute{
					Name:        "additional_cidrs_designated_gateway",
					Description: `(Optional) A list of CIDR ranges separated by comma to configure when "Designated Gateway" feature is enabled. Example: "10.8.0.0/16,10.9.0.0/16,10.10.0.0/16". ### Encryption`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `(Optional) Enable EBS volume encryption for the gateway. Only supported for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "customer_managed_keys",
					Description: `(Optional and Sensitive) Customer-managed key ID. ### Monitor Gateway Subnets ~>`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `(Optional) If set to true, the [Monitor Gateway Subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet) feature is enabled. Default value is false. Available in provider version R2.17.1+. ~>`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `(Optional) Set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true. Available in provider version R2.17.1+. ### FQDN Gateway`,
				},
				resource.Attribute{
					Name:        "fqdn_lan_cidr",
					Description: `(Optional) If ` + "`" + `fqdn_lan_cidr` + "`" + ` is set, the FQDN gateway will be created with an additional LAN interface using the provided CIDR. This attribute is required when enabling FQDN gateway FireNet in Azure or GCP. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "fqdn_lan_vpc_id",
					Description: `(Optional) FQDN LAN VPC ID. This attribute is required when enabling FQDN gateway FireNet in GCP. Available as of provider version R2.18.1+. ### Spot Instance`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `(Optional) Enable spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `(Optional) Price for spot instance. NOT supported for production deployment. ### Gateway Upgrade`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `(Optional/Computed) The software version of the gateway. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrade can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `(Optional/Computed) The image version of the gateway. Use ` + "`" + `aviatrix_gateway_image` + "`" + ` data source to programmatically retrieve this value for the desired ` + "`" + `software_version` + "`" + `. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "peering_ha_software_version",
					Description: `(Optional/Computed) The software version of the HA gateway. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the HA gateway upgrade can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "peering_ha_image_version",
					Description: `(Optional/Computed) The image version of the HA gateway. Use ` + "`" + `aviatrix_gateway_image` + "`" + ` data source to programmatically retrieve this value for the desired ` + "`" + `ha_software_version` + "`" + `. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0. ### Misc.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `(Optional) If set to false, use an available address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 2.7+. Valid values: true, false. Default: true.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) Specified EIP to use for gateway creation. Required when ` + "`" + `allocate_new_eip` + "`" + ` is false. Available in Controller version 3.5+. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP Address resource and its resource group in Azure to be assigned to the gateway instance. Example: "IP_Name:Resource_Group_Name". Required when ` + "`" + `allocate_new_eip` + "`" + ` is false and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC DNS Server for gateway. Currently only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, Alibaba Cloud, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability Zone. Only available for Azure (8), Azure GOV (32), Azure CHINA (2048) and Public Subnet Filtering gateway. Available for Azure as of provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Enable jumbo frames for this gateway. Default value is true.`,
				},
				resource.Attribute{
					Name:        "enable_gro_gso",
					Description: `(Optional) Enable GRO/GSO for this transit gateway. Default value is true. Available in provider R3.1.0+.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Map of tags to assign to the gateway. Only available for AWS, AWSGov, AWSChina, Azure, AzureGov, AzureChina, AWS Top Secret and AWS Secret gateways. Allowed characters vary by cloud type but always include: letters, spaces, and numbers. AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret allow the use of any character. Azure, AzureGov and AzureChina allows the following special characters: + - = . _ : @. Example: {"key1" = "value1", "key2" = "value2"}.`,
				},
				resource.Attribute{
					Name:        "tunnel_detection_time",
					Description: `(Optional) The IPsec tunnel down detection time for the Gateway in seconds. Must be a number in the range [20-600]. The default value is set by the controller (60 seconds if nothing has been changed).`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Gateway ethernet interface RX queue size. Applies on HA as well if enabled. Once set, can't be deleted or disabled. Available for AWS as of provider version R2.22+. ### Public Subnet Filtering Gateway ~>`,
				},
				resource.Attribute{
					Name:        "enable_public_subnet_filtering",
					Description: `(Optional) Create a [Public Subnet Filtering gateway](https://docs.aviatrix.com/HowTos/public_subnet_filtering_faq.html). Valid values: true or false. Default value: false. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "public_subnet_filtering_route_tables",
					Description: `(Optional) Route tables whose associated public subnets are protected. Only valid when ` + "`" + `enable_public_subnet_filtering` + "`" + ` attribute is true. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "public_subnet_filtering_ha_route_tables",
					Description: `(Optional) Route tables whose associated public subnets are protected for the HA PSF gateway. Required when ` + "`" + `enable_public_subnet_filtering` + "`" + ` and ` + "`" + `peering_ha_subnet` + "`" + ` are set. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "public_subnet_filtering_guard_duty_enforced",
					Description: `(Optional) Whether to enforce Guard Duty IP blocking. Only valid when ` + "`" + `enable_public_subnet_filtering` + "`" + ` attribute is true. Valid values: true or false. Default value: true. Available as of provider version R2.18+. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "elb_dns_name",
					Description: `ELB DNS name.`,
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
					Name:        "peering_ha_security_group_id",
					Description: `HA security group used for the gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the gateway created.`,
				},
				resource.Attribute{
					Name:        "peering_ha_cloud_instance_id",
					Description: `Cloud instance ID of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_name",
					Description: `Aviatrix gateway unique name of HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_private_ip",
					Description: `Private IP address of HA gateway.`,
				},
				resource.Attribute{
					Name:        "fqdn_lan_interface",
					Description: `The lan interface id of the of FQDN gateway with additional LAN interface. This attribute will be exported when enabling FQDN gateway firenet in Azure. Available in provider version R2.17.1+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Specify the DNS IP, only required while using a custom private DNS for the VPC.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and Azure(8).`,
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
					Name:        "cloudn_bkup_gateway_inst_id",
					Description: `Instance ID of the backup gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the gateway created.`,
				},
				resource.Attribute{
					Name:        "peering_ha_public_ip",
					Description: `Public IP address of the peering HA Gateway created.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Tag list of the gateway instance. Only available for AWS, AWSGov, AWSChina, Azure, AzureGov and AzureChina gateways. Example: ["key1:value1", "key2:value2"]. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "elb_dns_name",
					Description: `ELB DNS name.`,
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
					Name:        "peering_ha_security_group_id",
					Description: `HA security group used for the gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the gateway created.`,
				},
				resource.Attribute{
					Name:        "peering_ha_cloud_instance_id",
					Description: `Cloud instance ID of the HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_gw_name",
					Description: `Aviatrix gateway unique name of HA gateway.`,
				},
				resource.Attribute{
					Name:        "peering_ha_private_ip",
					Description: `Private IP address of HA gateway.`,
				},
				resource.Attribute{
					Name:        "fqdn_lan_interface",
					Description: `The lan interface id of the of FQDN gateway with additional LAN interface. This attribute will be exported when enabling FQDN gateway firenet in Azure. Available in provider version R2.17.1+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "dns_server",
					Description: `Specify the DNS IP, only required while using a custom private DNS for the VPC.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and Azure(8).`,
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
					Name:        "cloudn_bkup_gateway_inst_id",
					Description: `Instance ID of the backup gateway.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the gateway created.`,
				},
				resource.Attribute{
					Name:        "peering_ha_public_ip",
					Description: `Public IP address of the peering HA Gateway created.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Tag list of the gateway instance. Only available for AWS, AWSGov, AWSChina, Azure, AzureGov and AzureChina gateways. Example: ["key1:value1", "key2:value2"]. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_gateway_certificate_config",
			Category:         "Settings",
			ShortDescription: `Manages Aviatrix gateway certificate configuration`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"gateway",
				"certificate",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ca_certificate",
					Description: `(Required) CA Certificate in PEM format. To read certificate from a file please use the built-in ` + "`" + `file` + "`" + ` function.`,
				},
				resource.Attribute{
					Name:        "ca_private_key",
					Description: `(Required/Sensitive) CA Private Key. To read the private key from a file please use the built-in ` + "`" + `file` + "`" + ` function. ## Import !>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_gateway_dnat",
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
					Description: `(Required) Name of the Aviatrix gateway the custom DNAT will be configured for.`,
				},
				resource.Attribute{
					Name:        "dnat_policy",
					Description: `(Required) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1) and Azure(8).`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a source IP address range where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) This is a qualifier condition that specifies a source port that the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a destination IP address range where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) This is a qualifier condition that specifies a destination port where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) This is a qualifier condition that specifies a destination port protocol where the rule applies. When not specified, this field is not used. !> ` + "`" + `interface` + "`" + ` must be empty when using a route-based connection in a policy.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) This is a qualifier condition that specifies output interface where the rule applies. When not specified, this field is not used. Must be empty when ` + "`" + `connection` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) This is a qualifier condition that specifies output connection where the rule applies. Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) This is a rule field that specifies a tag or mark of a TCP session when all qualifier conditions meet. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dnat_ips",
					Description: `(Optional) This is a rule field that specifies the translated destination IP address when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "dnat_port",
					Description: `(Optional) This is a rule field that specifies the translated destination port when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule field must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry.`,
				},
				resource.Attribute{
					Name:        "apply_route_entry",
					Description: `(Optional) This is an option to program the route entry 'DST CIDR pointing to Aviatrix Gateway' into Cloud platform routing table. Type: Boolean. Default: True. Available as of provider version R2.19.2+. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_gateway_snat",
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
					Description: `(Required) Name of the Aviatrix gateway the custom SNAT will be configured for.`,
				},
				resource.Attribute{
					Name:        "snat_mode",
					Description: `(Optional) NAT mode. Valid values: "customized_snat". Default value: "customized_snat".`,
				},
				resource.Attribute{
					Name:        "snat_policy",
					Description: `(Required) Policy rule applied for enabling source NAT (mode: "customized_snat"). Currently only supports AWS(1) and Azure(8).`,
				},
				resource.Attribute{
					Name:        "src_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a source IP address range where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "src_port",
					Description: `(Optional) This is a qualifier condition that specifies a source port that the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_cidr",
					Description: `(Optional) This is a qualifier condition that specifies a destination IP address range where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "dst_port",
					Description: `(Optional) This is a qualifier condition that specifies a destination port where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) This is a qualifier condition that specifies a destination port protocol where the rule applies. Valid values: 'all', 'tcp', 'udp', 'icmp'. 'Default: 'all'. !> ` + "`" + `interface` + "`" + ` must be empty when using a route-based connection in a policy.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `(Optional) This is a qualifier condition that specifies output interface where the rule applies. When not specified, this field is not used. Must be empty when ` + "`" + `connection` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "connection",
					Description: `(Optional) This is a qualifier condition that specifies output connection where the rule applies. Default value: "None".`,
				},
				resource.Attribute{
					Name:        "mark",
					Description: `(Optional) This is a qualifier condition that specifies a tag or mark of a TCP session where the rule applies. When not specified, this field is not used.`,
				},
				resource.Attribute{
					Name:        "snat_ips",
					Description: `(Optional) This is a rule field that specifies the changed source IP address when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "snat_port",
					Description: `(Optional) This is a rule field that specifies the changed source port when all specified qualifier conditions meet. When not specified, this field is not used. One of the rule fields must be specified for this rule to take effect.`,
				},
				resource.Attribute{
					Name:        "exclude_rtb",
					Description: `(Optional) This field specifies which VPC private route table will not be programmed with the default route entry. !>`,
				},
				resource.Attribute{
					Name:        "apply_route_entry",
					Description: `(Optional) This is an option to program the route entry 'DST CIDR pointing to Aviatrix Gateway' into Cloud platform routing table. Type: Boolean. Default: True. Available as of provider version R2.21.0+. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_geo_vpn",
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
					Description: `(Required) List of ELB names to attach to this Geo VPN name. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_global_vpc_excluded_instance",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Manages the instance to be excluded for Aviatrix global VPC tagging`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"global",
				"vpc",
				"excluded",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Aviatrix GCP access account name.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `(Required) Name of the instance to be excluded for tagging.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of the instance. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the global VPC exclude list object. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the global VPC exclude list object. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_global_vpc_tagging_settings",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Manages how to tag newly found instances for global VPC.`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"global",
				"vpc",
				"tagging",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_state",
					Description: `(Required) Tagging service state. Valid values: "semi_automatic", "automatic", "disabled".`,
				},
				resource.Attribute{
					Name:        "enable_alert",
					Description: `(Required) Set to true to enable alert. Valid values: true, false. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_link_hierarchy",
			Category:         "Copilot",
			ShortDescription: `Creates Aviatrix Link Hierarchy`,
			Description:      ``,
			Keywords: []string{
				"copilot",
				"link",
				"hierarchy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Link hierarchy name. ### Optional`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `(Optional) List of named links.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Link name.`,
				},
				resource.Attribute{
					Name:        "wan_link",
					Description: `(Optional) Set of WAN links.`,
				},
				resource.Attribute{
					Name:        "wan_tag",
					Description: `(Optional) WAN tag. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of link hierarchy. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of link hierarchy. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_netflow_agent",
			Category:         "Settings",
			ShortDescription: `Enables and disables netflow_agent`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"netflow",
				"agent",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of netflow agent. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of netflow agent. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_periodic_ping",
			Category:         "Gateway",
			ShortDescription: `Manages periodic pings on Aviatrix gateways`,
			Description:      ``,
			Keywords: []string{
				"gateway",
				"periodic",
				"ping",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required) Interval between pings in seconds.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP Address to ping. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_private_mode_lb",
			Category:         "Private Mode",
			ShortDescription: `Creates and manages a Private Mode load balancer`,
			Description:      ``,
			Keywords: []string{
				"private",
				"mode",
				"lb",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Name of the access account.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the VPC for the load balancer.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Name of the region containing the VPC.`,
				},
				resource.Attribute{
					Name:        "lb_type",
					Description: `(Required) Type of load balancer to create. Must be one of controller or multicloud.`,
				},
				resource.Attribute{
					Name:        "multicloud_access_vpc_id",
					Description: `(Optional) ID of the VPC with a multicloud endpoint. Required when ` + "`" + `lb_type` + "`" + ` is multicloud.`,
				},
				resource.Attribute{
					Name:        "proxies",
					Description: `(Optional) List of multicloud proxies. Only valid when ` + "`" + `lb_type` + "`" + ` is multicloud.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Instance ID of the proxy.`,
				},
				resource.Attribute{
					Name:        "proxy_type",
					Description: `(Required) Type of load balancer. Must be one of controller or multicloud.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the proxy. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_private_mode_multicloud_endpoint",
			Category:         "Private Mode",
			ShortDescription: `Creates and manages a Private Mode multicloud endpoint`,
			Description:      ``,
			Keywords: []string{
				"private",
				"mode",
				"multicloud",
				"endpoint",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) Name of the access account.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) ID of the VPC to create the endpoint in.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region of the VPC.`,
				},
				resource.Attribute{
					Name:        "controller_lb_vpc_id",
					Description: `(Required) ID of the VPC containing a Private Mode controller load balancer. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dns_entry",
					Description: `DNS entry of the endpoint. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns_entry",
					Description: `DNS entry of the endpoint. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_proxy_config",
			Category:         "Settings",
			ShortDescription: `Creates and manages an Aviatrix controller proxy config resource`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"proxy",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "http_proxy",
					Description: `(Required) Http proxy URL.`,
				},
				resource.Attribute{
					Name:        "https_proxy",
					Description: `(Required) Https proxy URL.`,
				},
				resource.Attribute{
					Name:        "proxy_ca_certificate",
					Description: `(Optional) Server CA Certificate file. Use the ` + "`" + `file` + "`" + ` function to read from a file. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_qos_class",
			Category:         "Copilot",
			ShortDescription: `Creates Aviatrix QoS Class`,
			Description:      ``,
			Keywords: []string{
				"copilot",
				"qos",
				"class",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of QoS class.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Optional) Priority of QoS class, requires an integer value. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of QoS class. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of QoS class. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_qos_policy_list",
			Category:         "Copilot",
			ShortDescription: `Creates Aviatrix QoS Policy List`,
			Description:      ``,
			Keywords: []string{
				"copilot",
				"qos",
				"policy",
				"list",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) List of QoS policies.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of QoS class.`,
				},
				resource.Attribute{
					Name:        "dscp_values",
					Description: `(Required) List of DSCP values.`,
				},
				resource.Attribute{
					Name:        "qos_class_uuid",
					Description: `(Required) QoS class UUID. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_rbac_group",
			Category:         "Accounts",
			ShortDescription: `Creates and manages Aviatrix RBAC groups`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"rbac",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) This parameter represents the name of a RBAC group to be created. ### Optional`,
				},
				resource.Attribute{
					Name:        "local_login",
					Description: `(Optional) Whether to allow members of an RBAC group to bypass LDAP/MFA for Duo login . Supported values: true, false. Default value: false. Available in provider version R2.17.1+. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_rbac_group_access_account_attachment",
			Category:         "Accounts",
			ShortDescription: `Creates and manages Aviatrix RBAC group access account attachments`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"rbac",
				"group",
				"access",
				"account",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) This parameter represents the name of a RBAC group.`,
				},
				resource.Attribute{
					Name:        "access_account_name",
					Description: `(Required) Account name. This can be used for logging in to CloudN console or UserConnect controller. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_rbac_group_permission_attachment",
			Category:         "Accounts",
			ShortDescription: `Creates and manages Aviatrix RBAC group permission attachments`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"rbac",
				"group",
				"permission",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) This parameter represents the name of a RBAC group.`,
				},
				resource.Attribute{
					Name:        "permission_name",
					Description: `(Required) This parameter represents the permission to attach to the RBAC group. Valid ` + "`" + `permission_name` + "`" + ` values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_rbac_group_user_attachment",
			Category:         "Accounts",
			ShortDescription: `Creates and manages Aviatrix RBAC group user attachments`,
			Description:      ``,
			Keywords: []string{
				"accounts",
				"rbac",
				"group",
				"user",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_name",
					Description: `(Required) This parameter represents the name of a RBAC group.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Username of the account user. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_remote_syslog",
			Category:         "Settings",
			ShortDescription: `Enables and disables remote syslog`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"remote",
				"syslog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of remote syslog.`,
				},
				resource.Attribute{
					Name:        "notls",
					Description: `This attribute is true if the remote syslog is not protected by TLS. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of remote syslog.`,
				},
				resource.Attribute{
					Name:        "notls",
					Description: `This attribute is true if the remote syslog is not protected by TLS. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_saml_endpoint",
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
					Description: `(Required) The SAML endpoint name.`,
				},
				resource.Attribute{
					Name:        "idp_metadata_type",
					Description: `(Required) The IDP Metadata type. Can be either "Text" or "URL".`,
				},
				resource.Attribute{
					Name:        "idp_metadata",
					Description: `(Optional) The IDP Metadata from SAML provider. Required if ` + "`" + `idp_metadata_type` + "`" + ` is "Text" and should be unset if type is "URL". Normally the metadata is in XML format which may contain special characters. Best practice is to use the file function to read from a local Metadata XML file.`,
				},
				resource.Attribute{
					Name:        "idp_metadata_url",
					Description: `(Optional) The IDP Metadata URL from SAML provider. Required if ` + "`" + `idp_metadata_type` + "`" + ` is "URL" and should be unset if type is "Text". ->`,
				},
				resource.Attribute{
					Name:        "custom_entity_id",
					Description: `(Optional) Custom Entity ID. Required to be non-empty for 'Custom' Entity ID type, empty for 'Hostname' Entity ID type.`,
				},
				resource.Attribute{
					Name:        "custom_saml_request_template",
					Description: `(Optional) Custom SAML Request Template in string. ### Advanced`,
				},
				resource.Attribute{
					Name:        "sign_authn_request",
					Description: `(Optional) Whether to sign SAML AuthnRequests. Supported values: true, false . Default value: false. Available in provider version R2.17.1+. ### Controller Login`,
				},
				resource.Attribute{
					Name:        "controller_login",
					Description: `(Optional) Valid values: true, false. Default value: false. Set true for creating a saml endpoint for controller login.`,
				},
				resource.Attribute{
					Name:        "access_set_by",
					Description: `(Optional) Access type. Valid values: "controller", "profile_attribute". Default value: "controller".`,
				},
				resource.Attribute{
					Name:        "rbac_groups",
					Description: `(Optional) List of rbac groups. Required for controller login and "access_set_by" of "controller". ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_segmentation_network_domain",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages an Aviatrix Segmentation Network Domain`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"segmentation",
				"network",
				"domain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Required) Name of the Network Domain. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_segmentation_network_domain_association",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages an Aviatrix Segmentation Network Domain Association`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"segmentation",
				"network",
				"domain",
				"association",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "network_domain_name",
					Description: `(Required) Name of the Segmentation Network Domain.`,
				},
				resource.Attribute{
					Name:        "attachment_name",
					Description: `(Required) Name of the transit gateway attachment, Spoke or Edge, to associate with the network domain. ### Optional`,
				},
				resource.Attribute{
					Name:        "transit_gateway_name",
					Description: `(Optional) Name of the Transit Gateway. ## Import ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_segmentation_network_domain_connection_policy",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages an Aviatrix Segmentation Network Domain Connection Policy`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"segmentation",
				"network",
				"domain",
				"connection",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_name_1",
					Description: `(Required) Name of the Network Domain to connect to Domain 2.`,
				},
				resource.Attribute{
					Name:        "domain_name_2",
					Description: `(Required) Name of the Network Domain to connect to Domain 1. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_site2cloud",
			Category:         "Site2Cloud",
			ShortDescription: `Creates and manages Aviatrix Site2Cloud connections`,
			Description:      ``,
			Keywords: []string{
				"site2cloud",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the cloud gateway.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Site2Cloud connection name.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_type",
					Description: `(Required) Remote gateway type. Valid Values: "generic", "avx", "aws", "azure", "sonicwall", "oracle".`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) Connection type. Valid Values: "mapped", "unmapped".`,
				},
				resource.Attribute{
					Name:        "tunnel_type",
					Description: `(Required) Site2Cloud tunnel type. Valid Values: "policy", "route".`,
				},
				resource.Attribute{
					Name:        "primary_cloud_gateway_name",
					Description: `(Required) Primary cloud gateway name.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_ip",
					Description: `(Required) Remote gateway IP.`,
				},
				resource.Attribute{
					Name:        "remote_subnet_cidr",
					Description: `(Required) Remote subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "remote_subnet_virtual",
					Description: `Remote subnet CIDR (Virtual).`,
				},
				resource.Attribute{
					Name:        "local_subnet_cidr",
					Description: `(Optional) Local subnet CIDR.`,
				},
				resource.Attribute{
					Name:        "local_subnet_virtual",
					Description: `Local subnet CIDR (Virtual).`,
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
					Description: `(Optional) Backup Pre-Shared Key.`,
				},
				resource.Attribute{
					Name:        "local_tunnel_ip",
					Description: `(Optional) Local tunnel IP address. Only valid for route based connection. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "remote_tunnel_ip",
					Description: `(Optional) Remote tunnel IP address. Only valid for route based connection. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "backup_local_tunnel_ip",
					Description: `(Optional) Backup local tunnel IP address. Only valid when HA enabled route based connection. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "backup_remote_tunnel_ip",
					Description: `(Optional) Backup remote tunnel IP address. Only valid when HA enabled route based connection. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "enable_single_ip_ha",
					Description: `(Optional) Enable single IP HA feature. Available as of provider version 2.19+. ### Custom Algorithms`,
				},
				resource.Attribute{
					Name:        "custom_algorithms",
					Description: `(Optional) Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "phase_1_authentication",
					Description: `(Optional) Phase one Authentication. Valid values: "SHA-1", "SHA-256", "SHA-384" and "SHA-512". Default value: "SHA-256".`,
				},
				resource.Attribute{
					Name:        "phase_2_authentication",
					Description: `(Optional) Phase two Authentication. Valid values: "NO-AUTH", "HMAC-SHA-1", "HMAC-SHA-256", "HMAC-SHA-384" and "HMAC-SHA-512". Default value: "HMAC-SHA-256".`,
				},
				resource.Attribute{
					Name:        "phase_1_dh_groups",
					Description: `(Optional) Phase one DH Groups. Valid values: "1", "2", "5", "14", "15", "16", "17", "18", "19", "20" and "21". Default value: "14".`,
				},
				resource.Attribute{
					Name:        "phase_2_dh_groups",
					Description: `(Optional) Phase two DH Groups. Valid values: "1", "2", "5", "14", "15", "16", "17", "18", "19", "20" and "21". Default value: "14".`,
				},
				resource.Attribute{
					Name:        "phase_1_encryption",
					Description: `(Optional) Phase one Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", and "AES-256-GCM-128". Default value: "AES-256-CBC".`,
				},
				resource.Attribute{
					Name:        "phase_2_encryption",
					Description: `(Optional) Phase two Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", "AES-256-GCM-128" and "NULL-ENCR". Default value: "AES-256-CBC". ### Encryption over ExpressRoute/DirectConnect`,
				},
				resource.Attribute{
					Name:        "private_route_encryption",
					Description: `(Optional) Private route encryption switch. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "route_table_list",
					Description: `(Optional) Route tables to modify. Required when ` + "`" + `private_route_encryption` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_latitude",
					Description: `(Optional) Latitude of remote gateway. Required when ` + "`" + `private_route_encryption` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_longitude",
					Description: `(Optional) Longitude of remote gateway. Required when ` + "`" + `private_route_encryption` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_latitude",
					Description: `(Optional) Latitude of backup remote gateway. Required when ` + "`" + `private_route_encryption` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_longitude",
					Description: `(Optional) Longitude of backup remote gateway. Required when ` + "`" + `private_route_encryption` + "`" + ` is true. ### Custom Mapped ~>`,
				},
				resource.Attribute{
					Name:        "custom_mapped",
					Description: `(Optional) Enable custom mapped connection. Default value: false. Valid values: true/false. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "remote_source_real_cidrs",
					Description: `(Optional) List of Remote Initiated Traffic Source Real CIDRs.`,
				},
				resource.Attribute{
					Name:        "remote_source_virtual_cidrs",
					Description: `(Optional) List of Remote Initiated Traffic Source Virtual CIDRs.`,
				},
				resource.Attribute{
					Name:        "remote_destination_real_cidrs",
					Description: `(Optional) List of Remote Initiated Traffic Destination Real CIDRs.`,
				},
				resource.Attribute{
					Name:        "remote_destination_virtual_cidrs",
					Description: `(Optional) List of Remote Initiated Traffic Destination Virtual CIDRs.`,
				},
				resource.Attribute{
					Name:        "local_source_real_cidrs",
					Description: `(Optional) List of Local Initiated Traffic Source Real CIDRs.`,
				},
				resource.Attribute{
					Name:        "local_source_virtual_cidrs",
					Description: `(Optional) List of Local Initiated Traffic Source Virtual CIDRs.`,
				},
				resource.Attribute{
					Name:        "local_destination_real_cidrs",
					Description: `(Optional) List of Local Initiated Traffic Destination Real CIDRs.`,
				},
				resource.Attribute{
					Name:        "local_destination_virtual_cidrs",
					Description: `(Optional) List of Local Initiated Traffic Destination Virtual CIDRs. ### Misc.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Authentication Type. Valid values: 'PSK' and 'Cert'. Default value: 'PSK'.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) Pre-Shared Key.`,
				},
				resource.Attribute{
					Name:        "ca_cert_tag_name",
					Description: `(Optional) Name of Remote CA Certificate Tag for creating Site2Cloud tunnels. Required for Cert based authentication type.`,
				},
				resource.Attribute{
					Name:        "remote_identifier",
					Description: `(Optional) Remote identifier. Required for Cert based authentication type. Example: "gw-10-10-0-115".`,
				},
				resource.Attribute{
					Name:        "backup_remote_identifier",
					Description: `(Optional) Backup remote identifier. Required for Cert based authentication type with HA enabled. Example: "gw-10-10-0-116".`,
				},
				resource.Attribute{
					Name:        "ssl_server_pool",
					Description: `(Optional) Specify ssl_server_pool. Default value: "192.168.44.0/24".`,
				},
				resource.Attribute{
					Name:        "enable_dead_peer_detection",
					Description: `(Optional) Enable/disable Deed Peer Detection for an existing site2cloud connection. Default value: true.`,
				},
				resource.Attribute{
					Name:        "enable_active_active",
					Description: `(Optional) Enable/disable active active HA for an existing site2cloud connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_ikev2",
					Description: `(Optional) Switch to enable IKEv2. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "forward_traffic_to_transit",
					Description: `(Optional) Enable spoke gateway with mapped site2cloud configurations to forward traffic from site2cloud connection to Aviatrix Transit Gateway. Default value: false. Valid values: true or false. Available in provider version 2.17.2+.`,
				},
				resource.Attribute{
					Name:        "enable_event_triggered_ha",
					Description: `(Optional) Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+. ->`,
				},
				resource.Attribute{
					Name:        "phase1_local_identifier",
					Description: `(Optional) Phase 1 local identifier. By default, gateway’s public IP is configured as the Local Identifier. Available as of provider version R3.1.0+.`,
				},
				resource.Attribute{
					Name:        "phase1_remote_identifier",
					Description: `(Optional) List of phase 1 remote identifier of the IPsec tunnel. This can be configured as a list of any string, including empty string. Example: ["1.2.3.4"] when HA is disabled, ["1.2.3.4", "abcd"] when HA is enabled. Available as of provider version R2.19+. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "local_subnet_cidr",
					Description: `Local subnet CIDR. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "local_subnet_cidr",
					Description: `Local subnet CIDR. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_site2cloud_ca_cert_tag",
			Category:         "Site2Cloud",
			ShortDescription: `Creates and manages Aviatrix Site2Cloud CA Cert Tags`,
			Description:      ``,
			Keywords: []string{
				"site2cloud",
				"ca",
				"cert",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tag_name",
					Description: `(Required) Site2Cloud ca cert tag name.`,
				},
				resource.Attribute{
					Name:        "ca_certificates",
					Description: `(Required) A set of CA certificates.`,
				},
				resource.Attribute{
					Name:        "cert_content",
					Description: `(Required) Content of cert certificate to create only one cert. One CA cert only per file.`,
				},
				resource.Attribute{
					Name:        "common_name",
					Description: `(Computed) Common name of created cert.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `(Computed) Expiration time of created cert.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) Unique id of created cert.`,
				},
				resource.Attribute{
					Name:        "issuer_name",
					Description: `(Computed) Issuer name of created cert.`,
				},
				resource.Attribute{
					Name:        "unique_serial",
					Description: `(Computed) Unique serial of created cert. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_sla_class",
			Category:         "Copilot",
			ShortDescription: `Creates Aviatrix SLA Class`,
			Description:      ``,
			Keywords: []string{
				"copilot",
				"sla",
				"class",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of SLA class. ### Optional`,
				},
				resource.Attribute{
					Name:        "latency",
					Description: `(Optional) Latency of sla class in ms.`,
				},
				resource.Attribute{
					Name:        "jitter",
					Description: `(Optional) Jitter of sla class in ms.`,
				},
				resource.Attribute{
					Name:        "packet_drop_rate",
					Description: `(Optional) Packet drop rate of sla class. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of SLA class. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of SLA class. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_smart_group",
			Category:         "Secured Networking",
			ShortDescription: `Creates and manages an Aviatrix Smart Group`,
			Description:      ``,
			Keywords: []string{
				"secured",
				"networking",
				"smart",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Smart Group.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Required) Block containing match expressions to filter the Smart Group.`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Required) List of match expressions. The Smart Group will be a union of all resources matched by each ` + "`" + `match_expressions` + "`" + `.` + "`" + `match_expressions` + "`" + ` blocks cannot be empty.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) - CIDR block or IP Address this expression matches. ` + "`" + `cidr` + "`" + ` cannot be used with any other filters in the same ` + "`" + `match_expressions` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Optional) - FQDN address this expression matches. ` + "`" + `fqdn` + "`" + ` cannot be used with any other filters in the same ` + "`" + `match_expressions` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "site",
					Description: `(Optional) - Edge Site-ID this expression matches. ` + "`" + `site` + "`" + ` cannot be used with any other filters in the same ` + "`" + `match_expressions` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) - Type of resource this expression matches. Must be one of "vm", "vpc" or "subnet". ` + "`" + `type` + "`" + ` is required when ` + "`" + `cidr` + "`" + `, ` + "`" + `fqdn` + "`" + ` and ` + "`" + `site` + "`" + ` are all not used.`,
				},
				resource.Attribute{
					Name:        "res_id",
					Description: `(Optional) - Resource ID this expression matches.`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `(Optional) - Account ID this expression matches.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Optional) - Account name this expression matches.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name this expression matches.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) - Region this expression matches.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) - Zone this expression matches.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) - Map of tags this expression matches. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the Smart Group. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the Smart Group. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_splunk_logging",
			Category:         "Settings",
			ShortDescription: `Enables and disables splunk logging`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"splunk",
				"logging",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of splunk logging. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of splunk logging. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_external_device_conn",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix spoke external device connections`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"external",
				"device",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the Aviatrix spoke gateway.`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Spoke external device connection name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_ip",
					Description: `(Optional) Remote gateway IP.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'. ~>`,
				},
				resource.Attribute{
					Name:        "tunnel_protocol",
					Description: `(Optional) Tunnel protocol, only valid with ` + "`" + `connection_type` + "`" + ` = 'bgp'. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive.`,
				},
				resource.Attribute{
					Name:        "bgp_local_as_num",
					Description: `(Optional) BGP local ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "bgp_remote_as_num",
					Description: `(Optional) BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `(Optional) Remote CIDRs joined as a string with ','. Required for a 'static' type connection. ### HA`,
				},
				resource.Attribute{
					Name:        "ha_enabled",
					Description: `(Optional) Set as true if there are two external devices.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_ip",
					Description: `(Optional) Backup remote gateway IP. Required if HA enabled.`,
				},
				resource.Attribute{
					Name:        "backup_bgp_remote_as_num",
					Description: `(Optional) Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required if HA enabled for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "backup_pre_shared_key",
					Description: `(Optional) Backup Pre-Shared Key.`,
				},
				resource.Attribute{
					Name:        "backup_local_tunnel_cidr",
					Description: `(Optional) Source CIDR for the tunnel from the backup Aviatrix spoke gateway.`,
				},
				resource.Attribute{
					Name:        "backup_remote_tunnel_cidr",
					Description: `(Optional) Destination CIDR for the tunnel to the backup external device.`,
				},
				resource.Attribute{
					Name:        "backup_direct_connect",
					Description: `(Optional) Backup direct connect for backup external device. ### Custom Algorithms`,
				},
				resource.Attribute{
					Name:        "custom_algorithms",
					Description: `(Optional) Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "phase_1_authentication",
					Description: `(Optional) Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'. Default value: 'SHA-256'.`,
				},
				resource.Attribute{
					Name:        "phase_2_authentication",
					Description: `(Optional) Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'. Default value: 'HMAC-SHA-256'.`,
				},
				resource.Attribute{
					Name:        "phase_1_dh_groups",
					Description: `(Optional) Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.`,
				},
				resource.Attribute{
					Name:        "phase_2_dh_groups",
					Description: `(Optional) Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.`,
				},
				resource.Attribute{
					Name:        "phase_1_encryption",
					Description: `(Optional) Phase one Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", and "AES-256-GCM-128". Default value: "AES-256-CBC".`,
				},
				resource.Attribute{
					Name:        "phase_2_encryption",
					Description: `(Optional) Phase two Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", "AES-256-GCM-128" and "NULL-ENCR". Default value: "AES-256-CBC". ### BGP over LAN (Available as of provider version R2.18+) ~>`,
				},
				resource.Attribute{
					Name:        "remote_vpc_name",
					Description: `(Optional) Name of the remote VPC for a LAN BGP connection with an Azure Spoke Gateway. Required when ` + "`" + `connection_type` + "`" + ` = 'bgp' and ` + "`" + `tunnel_protocol` + "`" + ` = 'LAN' with an Azure spoke gateway. Must be in the format "<vnet-name>:<resource-group-name>:<subscription-id>". Available as of provider version R3.0.2+.`,
				},
				resource.Attribute{
					Name:        "remote_lan_ip",
					Description: `(Optional) Remote LAN IP. Required for BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "local_lan_ip",
					Description: `(Optional) Local LAN IP. Required for GCP BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "backup_remote_lan_ip",
					Description: `(Optional) Backup Remote LAN IP. Required for HA BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "backup_local_lan_ip",
					Description: `(Optional) Backup Local LAN IP. Required for GCP HA BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "enable_bgp_lan_activemesh",
					Description: `(Optional) Switch to enable BGP LAN ActiveMesh mode. Only valid for GCP and Azure with Remote Gateway HA enabled. Requires Azure Remote Gateway insane mode enabled. Valid values: true, false. Default: false. Available as of provider version R3.0.2+. ### BGP MD5 Authentication (Available as of provider version R2.21.1+) ~>`,
				},
				resource.Attribute{
					Name:        "bgp_md5_key",
					Description: `(Optional) BGP MD5 Authentication Key. Example: 'avx01,avx02'.`,
				},
				resource.Attribute{
					Name:        "backup_bgp_md5_key",
					Description: `(Optional) Backup BGP MD5 Authentication Key. Valid with HA enabled for connection. Example: 'avx03,avx04'. ### Misc.`,
				},
				resource.Attribute{
					Name:        "direct_connect",
					Description: `(Optional) Set true for private network infrastructure.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) Pre-Shared Key.`,
				},
				resource.Attribute{
					Name:        "local_tunnel_cidr",
					Description: `(Optional) Source CIDR for the tunnel from the Aviatrix spoke gateway.`,
				},
				resource.Attribute{
					Name:        "remote_tunnel_cidr",
					Description: `(Optional) Destination CIDR for the tunnel to the external device.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Enable learned CIDRs approval for the connection. Only valid with ` + "`" + `connection_type` + "`" + ` = 'bgp'. Requires the spoke_gateway's ` + "`" + `learned_cidrs_approval_mode` + "`" + ` attribute be set to 'connection'. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "approved_cidrs",
					Description: `(Optional/Computed) Set of approved CIDRs. Requires ` + "`" + `enable_learned_cidrs_approval` + "`" + ` to be true. Type: Set(String).`,
				},
				resource.Attribute{
					Name:        "enable_ikev2",
					Description: `(Optional) Set as true to enable IKEv2 protocol.`,
				},
				resource.Attribute{
					Name:        "manual_bgp_advertised_cidrs",
					Description: `(Optional) Configure manual BGP advertised CIDRs for this connection. Only valid with ` + "`" + `connection_type` + "`" + `= 'bgp'.`,
				},
				resource.Attribute{
					Name:        "enable_event_triggered_ha",
					Description: `(Optional) Enable Event Triggered HA. Default value: false. Valid values: true or false.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Enable Jumbo Frame for the transit external device connection. Only valid with 'GRE' tunnels under 'bgp' connection. Requires spoke to be jumbo frame and insane mode enabled. Valid values: true, false. Default value: false. Available as of provider version R3.0.2+. ->`,
				},
				resource.Attribute{
					Name:        "phase1_local_identifier",
					Description: `(Optional) Phase 1 local identifier. By default, gateway’s public IP is configured as the Local Identifier. Available as of provider version R3.1.0+.`,
				},
				resource.Attribute{
					Name:        "phase1_remote_identifier",
					Description: `(Optional) List of phase 1 remote identifier of the IPsec tunnel. This can be configured as a list of any string, including emtpy string. Example: ["1.2.3.4"] when HA is disabled, ["1.2.3.4", "abcd"] when HA is enabled. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_gateway",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix spoke gateways`,
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
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently, only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192), AWS Top Secret(16384) and AWS Secret (32768) are supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway which is going to be created. !> When creating a Spoke Gateway with an Azure VNet created in Controller version 6.4 or earlier or with an Azure VNet created out of band, referencing ` + "`" + `vpc_id` + "`" + ` in another resource on the same apply that creates this Spoke Gateway will cause Terraform to throw an error. Please use the Spoke Gateway data source to reference the ` + "`" + `vpc_id` + "`" + ` of this Spoke Gateway in other resources. ~> As of Provider version R2.21.2+, the ` + "`" + `vpc_id` + "`" + ` of an OCI VCN has been changed from its name to its OCID. !> As of Provider version R.22.0+, the ` + "`" + `vpc_id` + "`" + ` of a GCP VPC has been updated to include the project ID, e.g. vpc_name~-~project_id. When creating a Spoke Gateway using the old format, referencing ` + "`" + `vpc_id` + "`" + ` in another resource on the same apply that creates this Spoke Gateway will cause Terraform to throw an error. Please use the Spoke Gateway data source to reference the ` + "`" + `vpc_id` + "`" + ` of this Spoke Gateway in other resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Example: AWS/AWSGov/AWSChina: "vpc-abcd1234", GCP: "vpc-gcp-test~-~project-id", Azure/AzureGov/AzureChina: "vnet_name:rg_name:resource_guid", OCI: "ocid1.vcn.oc1.iad.aaaaaaaaba3pv6wkcr4jqae5f44n2b2m2yt2j6rx32uzr4h25vqstifsfdsq".`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west2-a", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-west-1, AWSChina: "cn-north-1", AzureChina: "China North", AWS Top Secret: "us-iso-east-1", AWS Secret: "us-isob-east-1".`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Size of the gateway instance. Example: AWS/AWSGov/AWSChina: "t2.large", Azure/AzureGov/AzureChina: "Standard_B1s", OCI: "VM.Standard2.2", GCP: "n1-standard-1".`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `(Optional) Fault domain. Required and valid only for OCI. Available as of provider version R2.19.3. ### HA`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `(Optional) HA Subnet. Required if enabling HA for AWS, AWSGov, AWSChina, Azure, AzureGov, AzureChina, OCI, Alibaba Cloud, AWS Top Secret or AWS Secret gateways. Optional for GCP. Setting to empty/unsetting will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet. Example: "10.12.0.0/24"`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `(Optional) HA Zone. Required if enabling HA for GCP gateway. Optional for Azure. For GCP, setting to empty/unsetting will disable HA and setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c". For Azure, this is an optional parameter to place the HA gateway in a specific availability zone. Valid values for Azure gateways are in the form "az-n". Example: "az-2". Available for Azure as of provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `(Optional) Public IP address that you want to assign to the HA peering instance. If no value is given, a new EIP will automatically be allocated. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.`,
				},
				resource.Attribute{
					Name:        "ha_azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP Address resource and its resource group in Azure to be assigned to the HA Spoke Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if ` + "`" + `ha_eip` + "`" + ` is set and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `(Optional) HA Gateway Size. Mandatory if enabling HA.`,
				},
				resource.Attribute{
					Name:        "ha_availability_domain",
					Description: `(Optional) HA gateway availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "ha_fault_domain",
					Description: `(Optional) HA gateway fault domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "manage_ha_gateway",
					Description: `(Optional) Enable to manage Aviatrix spoke HA gateway using the aviatrix_spoke_gateway resource. If this is set to false, spoke HA gateways must be managed using the aviatrix_spoke_ha_gateway resource. Valid values: true, false. Default value: true. Available in provider R3.0+. ->`,
				},
				resource.Attribute{
					Name:        "manage_ha_gateway",
					Description: `If you are using/upgraded to Aviatrix Terraform Provider R3.0+, and an aviatrix_spoke_gateway resource was originally created with a provider version <R3.0, you must do 'terraform refresh' to update and apply the attribute's default value (true) into the state file. Please see notes [Introduction to Gateway Group](https://registry.terraform.io/providers/AviatrixSystems/aviatrix/latest/docs/guides/introduction_to_gateway_group) for more information. ### Insane Mode`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Enable [Insane Mode](https://docs.aviatrix.com/HowTos/insane_mode.html) for Spoke Gateway. Insane Mode gateway size must be at least c5 size (AWS, AWSGov, AWS China, AWS Top Secret and AWS Secret) or Standard_D3_v2 (Azure and AzureGov); for GCP only four size are supported: "n1-highcpu-4", "n1-highcpu-8", "n1-highcpu-16" and "n1-highcpu-32". If enabled, you must specify a valid /26 CIDR segment of the VPC to create a new subnet for AWS, Azure, AzureGov, AWSGov, AWS Top Secret and AWS Secret. Only available for AWS, GCP/OCI, Azure, AzureGov, AzureChina, AWSGov, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Spoke Gateway. Required for AWS, AWSGov, AWS China, AWS Top Secret or AWS Secret if ` + "`" + `insane_mode` + "`" + ` is enabled. Example: AWS: "us-west-1a". ### SNAT/DNAT`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `(Optional) Specify whether to enable Source NAT feature in "single_ip" mode on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently, only supports AWS(1) and Azure(8). Valid values: true, false. ->`,
				},
				resource.Attribute{
					Name:        "bgp_ecmp",
					Description: `(Optional) Enable Equal Cost Multi Path (ECMP) routing for the next hop. Default value: false.`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `(Optional) BGP hold time. Unit is in seconds. Valid values are between 12 and 360. Default value: 180.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `(Optional) BGP route polling time. Unit is in seconds. Valid values are between 10 and 50. Default value: "50".`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_manual_advertise_cidrs",
					Description: `(Optional) Intended CIDR list to be advertised to external BGP router. Empty list is not valid. Example: ["10.2.0.0/16", "10.4.0.0/16"].`,
				},
				resource.Attribute{
					Name:        "enable_active_standby",
					Description: `(Optional) Enables [Active-Standby Mode](https://docs.aviatrix.com/HowTos/transit_advanced.html#active-standby). Available only with HA enabled. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby_preemptive",
					Description: `(Optional) Enables Preemptive Mode for Active-Standby. Available only with BGP enabled, HA enabled and Active-Standby enabled. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) Changes the Aviatrix Spoke Gateway ASN number before you setup Aviatrix Spoke Gateway connection configurations.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) List of AS numbers to populate BGP AS_PATH field when it advertises to VGW or peer devices.`,
				},
				resource.Attribute{
					Name:        "disable_route_propagation",
					Description: `(Optional) Disables route propagation on BGP Spoke to attached Transit Gateway. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_preserve_as_path",
					Description: `(Optional) Enable preserve as_path when advertising manual summary cidrs on BGP spoke gateway. Valid values: true, false. Default value: false. Available as of provider version R.2.22.1+ ### BGP over LAN`,
				},
				resource.Attribute{
					Name:        "enable_bgp_over_lan",
					Description: `(Optional) Pre-allocate a network interface(eth4) for "BGP over LAN" functionality. Must be enabled to create a BGP over LAN ` + "`" + `aviatrix_spoke_external_device_conn` + "`" + ` resource with this Spoke Gateway. Only valid for 8 (Azure), 32 (AzureGov) or AzureChina (2048). Valid values: true or false. Default value: false. Available as of provider version R3.0.2+. ->`,
				},
				resource.Attribute{
					Name:        "bgp_lan_interfaces_count",
					Description: `(Optional) Number of interfaces that will be created for BGP over LAN enabled Azure spoke. Applies on HA Transit as well if enabled. Available as of provider version R3.0.2+. ### Encryption`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `(Optional) Enable EBS volume encryption for Gateway. Only supports AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret providers. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "customer_managed_keys",
					Description: `(Optional and Sensitive) Customer managed key ID. ### Route Customization`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `(Optional) A list of comma separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to this spoke gateway only. Example: "10.0.0.0/116,10.2.0.0/16".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `(Optional) A list of comma separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to this spoke gateway only. Example: "10.2.0.0/116,10.3.0.0/16".`,
				},
				resource.Attribute{
					Name:        "included_advertised_spoke_routes",
					Description: `(Optional) A list of comma separated CIDRs to be advertised to on-prem as 'Included CIDR List'. When configured, it will replace all advertised routes from this VPC. Example: "10.4.0.0/116,10.5.0.0/16".`,
				},
				resource.Attribute{
					Name:        "enable_private_vpc_default_route",
					Description: `(Optional) Program default route in VPC private route table. Default: false. Valid values: true or false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "enable_skip_public_route_table_update",
					Description: `(Optional) Skip programming VPC public route table. Default: false. Valid values: true or false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "enable_auto_advertise_s2c_cidrs",
					Description: `(Optional) Auto Advertise Spoke Site2Cloud CIDRs. Default: false. Valid values: true or false. Available as of provider version R2.19+. ### [Learned CIDRs Approval for BGP Spoke Gateway](https://docs.aviatrix.com/HowTos/transit_approval.html) ->`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable/disable learned CIDR approval for BGP Spoke Gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "learned_cidrs_approval_mode",
					Description: `(Optional) Learned CIDRs approval mode. Either "gateway" (approval on a per-gateway basis) or "connection" (approval on a per-connection basis). Only "gateway" is supported for BGP SPOKE Gateway. Default value: "gateway". Available as of provider version R2.21+.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `(Optional) A set of approved learned CIDRs. Only valid when ` + "`" + `enable_learned_cidrs_approval` + "`" + ` is set to true. Example: ["10.250.0.0/16", "10.251.0.0/16"]. Available as of provider version R2.21+. ### [Monitor Gateway Subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet) ~>`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `(Optional) If set to true, the [Monitor Gateway Subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet) feature is enabled. Default value is false. Available in provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `(Optional) Set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true. Available in provider version R2.18+. ### [Private OOB](https://docs.aviatrix.com/HowTos/private_oob.html)`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `(Optional) Enable Private OOB feature. Only available for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `(Optional) OOB management subnet. Required if enabling Private OOB. Example: "11.0.2.0/24".`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `(Optional) OOB availability zone. Required if enabling Private OOB. Example: "us-west-1a".`,
				},
				resource.Attribute{
					Name:        "ha_oob_management_subnet",
					Description: `(Optional) HA OOB management subnet. Required if enabling Private OOB and HA. Example: "11.0.0.48/28".`,
				},
				resource.Attribute{
					Name:        "ha_oob_availability_zone",
					Description: `(Optional) HA OOB availability zone. Required if enabling Private OOB and HA. Example: "us-west-1b". ### Spot Instance`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `(Optional) Enable spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `(Optional) Price for spot instance. NOT supported for production deployment. ### Gateway Upgrade`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `(Optional/Computed) The software version of the gateway. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrade can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `(Optional/Computed) The image version of the gateway. Use ` + "`" + `aviatrix_gateway_image` + "`" + ` data source to programmatically retrieve this value for the desired ` + "`" + `software_version` + "`" + `. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "ha_software_version",
					Description: `(Optional/Computed) The software version of the HA gateway. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the HA gateway upgrade can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "ha_image_version",
					Description: `(Optional/Computed) The image version of the HA gateway. Use ` + "`" + `aviatrix_gateway_image` + "`" + ` data source to programmatically retrieve this value for the desired ` + "`" + `ha_software_version` + "`" + `. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0. ### Misc.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `(Optional) When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 4.7+. Valid values: true, false. Default: true.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) Required when ` + "`" + `allocate_new_eip` + "`" + ` is false. It uses the specified EIP for this gateway. Available in Controller 4.7+. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP Address resource and its resource group in Azure to be assigned to the Spoke Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if ` + "`" + `allocate_new_eip` + "`" + ` is false and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC DNS Server for Gateway. Currently only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, Alibaba Cloud, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability Zone. Only available for Azure (8), Azure GOV (32) and Azure CHINA (2048). Must be in the form 'az-n', for example, 'az-2'. Available in provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Enable jumbo frames for this spoke gateway. Default value is true.`,
				},
				resource.Attribute{
					Name:        "enable_gro_gso",
					Description: `(Optional) Enable GRO/GSO for this transit gateway. Default value is true. Available in provider R3.1.0+.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Map of tags to assign to the gateway. Only available for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret gateways. Allowed characters vary by cloud type but always include: letters, spaces, and numbers. AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret allow the use of any character. Azure, AzureGov and AzureChina allows the following special characters: + - = . _ : @. Example: {"key1" = "value1", "key2" = "value2"}.`,
				},
				resource.Attribute{
					Name:        "tunnel_detection_time",
					Description: `(Optional) The IPsec tunnel down detection time for the Spoke Gateway in seconds. Must be a number in the range [20-600]. The default value is set by the controller (60 seconds if nothing has been changed).`,
				},
				resource.Attribute{
					Name:        "enable_bgp",
					Description: `(Optional) Enable BGP for this spoke gateway. Only available for AWS and Azure. Valid values: true, false. Default value: false. Available in provider R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Gateway ethernet interface RX queue size. Applies on HA as well if enabled. Once set, can't be deleted or disabled. Available for AWS as of provider version R2.22+.`,
				},
				resource.Attribute{
					Name:        "private_mode_lb_vpc_id",
					Description: `(Optional) VPC ID of Private Mode load balancer. Required when Private Mode is enabled on the Controller. Available in provider version R2.23+.`,
				},
				resource.Attribute{
					Name:        "private_mode_subnet_zone",
					Description: `(Optional) Availability Zone of the subnet. Required when Private Mode is enabled on the Controller and ` + "`" + `cloud_type` + "`" + ` is AWS or AWSGov. Available in Provider version R2.23+.`,
				},
				resource.Attribute{
					Name:        "ha_private_mode_subnet_zone",
					Description: `(Optional) Availability Zone of the HA subnet. Required when Private Mode is enabled on the Controller and ` + "`" + `cloud_type` + "`" + ` is AWS or AWSGov with HA. Available in Provider version R2.23+. !>`,
				},
				resource.Attribute{
					Name:        "enable_globla_vpc",
					Description: `(Optional) Enable global VPC. Only supported for GCP spoke gateways. Valid values: true, false. Default value: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix spoke gateway unique name of HA spoke gateway.`,
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
					Name:        "public_ip",
					Description: `Public IP address of the Spoke Gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address of the HA Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the spoke gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_security_group_id",
					Description: `HA security group used for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of the HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for spoke external device connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for spoke external device HA connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Specify whether enabling Source NAT feature on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently only supports AWS(1), Azure(8) and AWSGov(256). Valid values: true, false.`,
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
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1), Azure(8), and AWSGov(256).`,
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
					Name:        "enable_active_mesh",
					Description: `(Optional) Switch to enable/disable [Active Mesh Mode](https://docs.aviatrix.com/HowTos/activemesh_faq.html) for Spoke Gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "manage_transit_gateway_attachment",
					Description: `(Optional) Enable to manage spoke-to-Aviatrix transit gateway attachments using the`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `(Optional) Specify the Aviatrix transit gateways to attach this spoke gateway to. Format is a comma separated list of transit gateway names. For example: "transit-gw1,transit-gw2".`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only supported for AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina. Example: ["key1:value1", "key2:value2"]. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix spoke gateway unique name of HA spoke gateway.`,
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
					Name:        "public_ip",
					Description: `Public IP address of the Spoke Gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address of the HA Spoke Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the spoke gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address of HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_security_group_id",
					Description: `HA security group used for the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the spoke gateway.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of the HA spoke gateway.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for spoke external device connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for spoke external device HA connection creation. Only supports 8 (Azure), 32 (AzureGov) or AzureChina (2048). Available as of provider version R3.0.2+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Specify whether enabling Source NAT feature on the gateway or not. Please disable AWS NAT instance before enabling this feature. Currently only supports AWS(1), Azure(8) and AWSGov(256). Valid values: true, false.`,
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
					Description: `(Optional) Policy rule applied for enabling Destination NAT (DNAT), which allows you to change the destination to a virtual address range. Currently only supports AWS(1), Azure(8), and AWSGov(256).`,
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
					Name:        "enable_active_mesh",
					Description: `(Optional) Switch to enable/disable [Active Mesh Mode](https://docs.aviatrix.com/HowTos/activemesh_faq.html) for Spoke Gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "manage_transit_gateway_attachment",
					Description: `(Optional) Enable to manage spoke-to-Aviatrix transit gateway attachments using the`,
				},
				resource.Attribute{
					Name:        "transit_gw",
					Description: `(Optional) Specify the Aviatrix transit gateways to attach this spoke gateway to. Format is a comma separated list of transit gateway names. For example: "transit-gw1,transit-gw2".`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only supported for AWS, Azure, AzureGov, AWSGov, AWSChina and AzureChina. Example: ["key1:value1", "key2:value2"]. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_gateway_subnet_group",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix spoke gateway subnet groups`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"gateway",
				"subnet",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of spoke gateway subnet group.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Aviatrix spoke gateway name.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) A set of subnets in the subnet group. The format of each subnet must be "CIDR~~subnet name". Example: ` + "`" + `["10.2.48.0/20~~subnet1", "10.2.64.0/20~~subnet2"]` + "`" + ` ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_ha_gateway",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix spoke ha gateways`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"ha",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "primary_gw_name",
					Description: `(Required) Name of the primary gateway which is already or will be created before this Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability Zone. Required for GCP gateway, example: "us-west1-c". Optional for Azure gateway in the form "az-n", example: "az-2".`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Availability domain. Required and valid only for OCI.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `(Optional) Fault domain. Required and valid only for OCI. ### Insane Mode`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Enable [Insane Mode](https://docs.aviatrix.com/HowTos/insane_mode.html) for Spoke HA Gateway. Insane Mode gateway size must be at least c5 size (AWS, AWSGov, AWS China, AWS Top Secret and AWS Secret) or Standard_D3_v2 (Azure and AzureGov); for GCP only four size are supported: "n1-highcpu-4", "n1-highcpu-8", "n1-highcpu-16" and "n1-highcpu-32". If enabled, you must specify a valid /26 CIDR segment of the VPC to create a new subnet for AWS, Azure, AzureGov, AWSGov, AWS Top Secret and AWS Secret. Only available for AWS, GCP/OCI, Azure, AzureGov, AzureChina, AWSGov, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Spoke HA Gateway. Required for AWS, AWSGov, AWS China, AWS Top Secret or AWS Secret if ` + "`" + `insane_mode` + "`" + ` is enabled. Example: AWS: "us-west-1a". ### Misc.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Optional) Name of the Spoke HA Gateway which is going to be created. If not set, controller will auto generate a name for this gateway.`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Optional) Size of the Spoke HA Gateway instance. If not set, controller will use the same value as primary gateway's. Example: AWS/AWSGov/AWSChina: "t2.large", Azure/AzureGov/AzureChina: "Standard_B1s", OCI: "VM.Standard2.2", GCP: "n1-standard-1".`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) If set, the set IP will be used for this gateway.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP Address resource and its resource group in Azure to be assigned to the Spoke Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if ` + "`" + `eip` + "`" + ` is set and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 3.0+. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `Software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region in which the Spoke HA Gateway was created.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the Spoke HA Gateway created.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Spoke HA Gateway created. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `Software version of the gateway.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `Image version of the gateway.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `Region in which the Spoke HA Gateway was created.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the Spoke HA Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the Spoke HA Gateway created.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Spoke HA Gateway created. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_spoke_transit_attachment",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix Spoke-to-Transit attachments`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"spoke",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "spoke_gw_name",
					Description: `(Required) Name of the spoke gateway to attach to transit network.`,
				},
				resource.Attribute{
					Name:        "transit_gw_name",
					Description: `(Required) Name of the transit gateway to attach the spoke gateway to. ### Advanced Options`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `(Optional) Learned routes will be propagated to these route tables. Example: ["rtb-212ff547","rtb-04539787"].`,
				},
				resource.Attribute{
					Name:        "enable_max_performance",
					Description: `(Optional) Indicates whether the maximum amount of HPE tunnels will be created. Only valid when transit and spoke gateways are each launched in Insane Mode and in the same cloud type. Default value: true. Available as of provider version R2.22.2+.`,
				},
				resource.Attribute{
					Name:        "spoke_prepend_as_path",
					Description: `(Optional) Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on spoke_gateway_name. Available as of provider version R2.23+.`,
				},
				resource.Attribute{
					Name:        "transit_prepend_as_path",
					Description: `(Optional) Connection based AS Path Prepend. Valid only for BGP connection. Can only use the gateway's own local AS number, repeated up to 25 times. Applies on transit_gateway_name. Available as of provider version R2.23+. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
				resource.Attribute{
					Name:        "spoke_bgp_enabled",
					Description: `Indicates whether the spoke gateway is BGP enabled or not. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "spoke_bgp_enabled",
					Description: `Indicates whether the spoke gateway is BGP enabled or not. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_sumologic_forwarder",
			Category:         "Settings",
			ShortDescription: `Enables and disables sumologic forwarder`,
			Description:      ``,
			Keywords: []string{
				"settings",
				"sumologic",
				"forwarder",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of sumologic forwarder. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The status of sumologic forwarder. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_traffic_classifier",
			Category:         "Copilot",
			ShortDescription: `Creates Aviatrix Traffic Classifier`,
			Description:      ``,
			Keywords: []string{
				"copilot",
				"traffic",
				"classifier",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policies",
					Description: `(Required) List of traffic classifier policies.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of traffic classifier.`,
				},
				resource.Attribute{
					Name:        "source_smart_group_uuids",
					Description: `(Required) List of source smart group UUIDs.`,
				},
				resource.Attribute{
					Name:        "destination_smart_group_uuids",
					Description: `(Required) List of destination smart group UUIDs.`,
				},
				resource.Attribute{
					Name:        "port_ranges",
					Description: `(Optional) Port ranges.`,
				},
				resource.Attribute{
					Name:        "low",
					Description: `(Optional) Low port range.`,
				},
				resource.Attribute{
					Name:        "high",
					Description: `(Optional) High port range.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol.`,
				},
				resource.Attribute{
					Name:        "link_hierarchy_uuid",
					Description: `(Optional) Link hierarchy UUID.`,
				},
				resource.Attribute{
					Name:        "sla_class_uuid",
					Description: `(Optional) SLA class UUID.`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `(Optional) Enable logging.`,
				},
				resource.Attribute{
					Name:        "route_type",
					Description: `(Optional) Route type. ## Attribute Reference In addition to all arguments above, the following attribute is exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_trans_peer",
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
					Description: `(Required) Destination CIDR. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_transit_external_device_conn",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix transit external device connections`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"external",
				"device",
				"conn",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID of the Aviatrix transit gateway. For GCP BGP over LAN connection, it is in the format of "vpc_name~-~project_name".`,
				},
				resource.Attribute{
					Name:        "connection_name",
					Description: `(Required) Transit external device connection name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Aviatrix transit gateway name.`,
				},
				resource.Attribute{
					Name:        "remote_gateway_ip",
					Description: `(Optional) Remote gateway IP. Required when ` + "`" + `tunnel_protocol` + "`" + ` != 'LAN'.`,
				},
				resource.Attribute{
					Name:        "connection_type",
					Description: `(Required) Connection type. Valid values: 'bgp', 'static'. Default value: 'bgp'. ~>`,
				},
				resource.Attribute{
					Name:        "tunnel_protocol",
					Description: `(Optional) Tunnel protocol, only valid with ` + "`" + `connection_type` + "`" + ` = 'bgp'. Valid values: 'IPsec', 'GRE' or 'LAN'. Default value: 'IPsec'. Case insensitive. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "bgp_local_as_num",
					Description: `(Optional) BGP local ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "bgp_remote_as_num",
					Description: `(Optional) BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "remote_subnet",
					Description: `(Optional) Remote CIDRs joined as a string with ','. Required for a 'static' type connection. ~>`,
				},
				resource.Attribute{
					Name:        "remote_vpc_name",
					Description: `(Optional) Name of the remote VPC for a LAN BGP connection with an Azure Transit Gateway. Required when ` + "`" + `connection_type` + "`" + ` = 'bgp' and ` + "`" + `tunnel_protocol` + "`" + ` = 'LAN' with an Azure transit gateway. Must be in the format "<vnet-name>:<resource-group-name>:<subscription-id>". Available as of provider version R2.18+. ### HA`,
				},
				resource.Attribute{
					Name:        "ha_enabled",
					Description: `(Optional) Set as true if there are two external devices.`,
				},
				resource.Attribute{
					Name:        "backup_remote_gateway_ip",
					Description: `(Optional) Backup remote gateway IP. Required if HA enabled.`,
				},
				resource.Attribute{
					Name:        "backup_bgp_remote_as_num",
					Description: `(Optional) Backup BGP remote ASN (Autonomous System Number). Integer between 1-4294967294. Required if HA enabled for 'bgp' connection.`,
				},
				resource.Attribute{
					Name:        "backup_pre_shared_key",
					Description: `(Optional) Backup Pre-Shared Key.`,
				},
				resource.Attribute{
					Name:        "backup_local_tunnel_cidr",
					Description: `(Optional) Source CIDR for the tunnel from the backup Aviatrix transit gateway.`,
				},
				resource.Attribute{
					Name:        "backup_remote_tunnel_cidr",
					Description: `(Optional) Destination CIDR for the tunnel to the backup external device.`,
				},
				resource.Attribute{
					Name:        "backup_direct_connect",
					Description: `(Optional) Backup direct connect for backup external device. ### Custom Algorithms`,
				},
				resource.Attribute{
					Name:        "custom_algorithms",
					Description: `(Optional) Switch to enable custom/non-default algorithms for IPSec Authentication/Encryption. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "phase_1_authentication",
					Description: `(Optional) Phase one Authentication. Valid values: 'SHA-1', 'SHA-256', 'SHA-384' and 'SHA-512'. Default value: 'SHA-256'.`,
				},
				resource.Attribute{
					Name:        "phase_2_authentication",
					Description: `(Optional) Phase two Authentication. Valid values: 'NO-AUTH', 'HMAC-SHA-1', 'HMAC-SHA-256', 'HMAC-SHA-384' and 'HMAC-SHA-512'. Default value: 'HMAC-SHA-256'.`,
				},
				resource.Attribute{
					Name:        "phase_1_dh_groups",
					Description: `(Optional) Phase one DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.`,
				},
				resource.Attribute{
					Name:        "phase_2_dh_groups",
					Description: `(Optional) Phase two DH Groups. Valid values: '1', '2', '5', '14', '15', '16', '17', '18', '19', '20' and '21'. Default value: '14'.`,
				},
				resource.Attribute{
					Name:        "phase_1_encryption",
					Description: `(Optional) Phase one Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", and "AES-256-GCM-128". Default value: "AES-256-CBC".`,
				},
				resource.Attribute{
					Name:        "phase_2_encryption",
					Description: `(Optional) Phase two Encryption. Valid values: "3DES", "AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "AES-128-GCM-64", "AES-128-GCM-96", "AES-128-GCM-128", "AES-256-GCM-64", "AES-256-GCM-96", "AES-256-GCM-128" and "NULL-ENCR". Default value: "AES-256-CBC". ### BGP over LAN (Available as of provider version R2.18+) ~>`,
				},
				resource.Attribute{
					Name:        "remote_lan_ip",
					Description: `(Optional) Remote LAN IP. Required for BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "local_lan_ip",
					Description: `(Optional) Local LAN IP. Required for GCP BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "backup_remote_lan_ip",
					Description: `(Optional) Backup Remote LAN IP. Required for HA BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "backup_local_lan_ip",
					Description: `(Optional) Backup Local LAN IP. Required for GCP HA BGP over LAN connection.`,
				},
				resource.Attribute{
					Name:        "enable_bgp_lan_activemesh",
					Description: `(Optional) Switch to enable BGP LAN ActiveMesh mode. Only valid for GCP and Azure with Remote Gateway HA enabled. Requires Azure Remote Gateway insane mode enabled. Valid values: true, false. Default: false. Available as of provider version R2.21+. ### BGP MD5 Authentication (Available as of provider version R2.21.1+) ~>`,
				},
				resource.Attribute{
					Name:        "bgp_md5_key",
					Description: `(Optional) BGP MD5 Authentication Key. Example: 'avx01,avx02'. For BGP LAN ActiveMesh mode disabled, example: 'avx01'.`,
				},
				resource.Attribute{
					Name:        "backup_bgp_md5_key",
					Description: `(Optional) Backup BGP MD5 Authentication Key. Valid with HA enabled for connection. Example: 'avx03,avx04'. For BGP LAN ActiveMesh mode disabled, example: 'avx03'. ### Misc.`,
				},
				resource.Attribute{
					Name:        "direct_connect",
					Description: `(Optional) Set true for private network infrastructure.`,
				},
				resource.Attribute{
					Name:        "pre_shared_key",
					Description: `(Optional) Pre-Shared Key.`,
				},
				resource.Attribute{
					Name:        "local_tunnel_cidr",
					Description: `(Optional) Source CIDR for the tunnel from the Aviatrix transit gateway.`,
				},
				resource.Attribute{
					Name:        "remote_tunnel_cidr",
					Description: `(Optional) Destination CIDR for the tunnel to the external device.`,
				},
				resource.Attribute{
					Name:        "enable_edge_segmentation",
					Description: `(Optional) Switch to allow this connection to communicate with a Network Domain via Connection Policy.`,
				},
				resource.Attribute{
					Name:        "switch_to_ha_standby_gateway",
					Description: `(Optional) Switch to HA Standby Transit Gateway connection. Only valid with Transit Gateway that has [Active-Standby Mode](https://docs.aviatrix.com/HowTos/transit_advanced.html#active-standby) enabled and for non-HA external device. Valid values: true, false. Default: false. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Enable learned CIDRs approval for the connection. Only valid with ` + "`" + `connection_type` + "`" + ` = 'bgp'. Requires the transit_gateway's ` + "`" + `learned_cidrs_approval_mode` + "`" + ` attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "approved_cidrs",
					Description: `(Optional/Computed) Set of approved CIDRs. Requires ` + "`" + `enable_learned_cidrs_approval` + "`" + ` to be true. Type: Set(String).`,
				},
				resource.Attribute{
					Name:        "enable_ikev2",
					Description: `(Optional) Set as true to enable IKEv2 protocol.`,
				},
				resource.Attribute{
					Name:        "manual_bgp_advertised_cidrs",
					Description: `(Optional) Configure manual BGP advertised CIDRs for this connection. Only valid with ` + "`" + `connection_type` + "`" + `= 'bgp'. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "enable_event_triggered_ha",
					Description: `(Optional) Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Enable Jumbo Frame for the transit external device connection. Only valid with 'GRE' tunnels under 'bgp' connection. Requires transit to be jumbo frame and insane mode enabled. Valid values: true, false. Default value: false. Available as of provider version R2.22.2+. ->`,
				},
				resource.Attribute{
					Name:        "phase1_local_identifier",
					Description: `(Optional) Phase 1 local identifier. By default, gateway’s public IP is configured as the Local Identifier. Available as of provider version R3.1.0+.`,
				},
				resource.Attribute{
					Name:        "phase1_remote_identifier",
					Description: `(Optional) List of phase 1 remote identifier of the IPsec tunnel. This can be configured as a list of any string, including emtpy string. Example: ["1.2.3.4"] when HA is disabled, ["1.2.3.4", "abcd"] when HA is enabled. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_transit_firenet_policy",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix Transit FireNet policies`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"firenet",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "transit_firenet_gateway_name",
					Description: `(Required) Name of the Transit FireNet-enabled transit gateway. Currently supports AWS and Azure.`,
				},
				resource.Attribute{
					Name:        "inspected_resource_name",
					Description: `(Required) The name of the resource which will be inspected. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_transit_gateway",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages the Aviatrix Transit Network gateways`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
				"gateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192), AWS Top Secret(16384) and AWS Secret (32768) are supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Required) Name of the gateway which is going to be created. !> When creating a Transit Gateway with an Azure VNet created in Controller version 6.4 or earlier or with an Azure VNet created out of band, referencing ` + "`" + `vpc_id` + "`" + ` in anothe resource on the same apply that creates this Transit Gateway will cause Terraform to throw an error. Please use the Transit Gateway data source to reference the ` + "`" + `vpc_id` + "`" + ` of this Transit Gateway in other resources. ~> As of Provider version R2.21.2+, the ` + "`" + `vpc_id` + "`" + ` of an OCI VCN has been changed from its name to its OCID. !> As of Provider version R.22.0+, the ` + "`" + `vpc_id` + "`" + ` of a GCP VPC has been updated to include the project ID, e.g. vpc_name~-~project_id. When creating a Transit Gateway using the old format, referencing ` + "`" + `vpc_id` + "`" + ` in another resource on the same apply that creates this Transit Gateway will cause Terraform to throw an error. Please use the Transit Gateway data source to reference the ` + "`" + `vpc_id` + "`" + ` of this Transit Gateway in other resources.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider. Example: AWS/AWSGov/AWSChina: "vpc-abcd1234", GCP: "vpc-gcp-test~-~project-id", Azure/AzureGov/AzureChina: "vnet_name:rg_name:resource_guid", OCI: "ocid1.vcn.oc1.iad.aaaaaaaaba3pv6wkcr4jqae5f44n2b2m2yt2j6rx32uzr4h25vqstifsfdsq".`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `(Required) Region of cloud provider. Example: AWS: "us-east-1", GCP: "us-west2-a", Azure: "East US 2", OCI: "us-ashburn-1", AzureGov: "USGov Arizona", AWSGov: "us-gov-west-1", AWSChina: "cn-north-1", AzureChina: "China North", AWS Top Secret: "us-iso-east-1", AWS Secret: "us-isob-east-1".`,
				},
				resource.Attribute{
					Name:        "gw_size",
					Description: `(Required) Size of the gateway instance. Example: AWS: "t2.large", Azure/AzureGov: "Standard_B1s", OCI: "VM.Standard2.2", GCP: "n1-standard-1", AWSGov: "t2.large", AWSChina: "t2.large", AzureChina: "Standard_A0".`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. Example: "172.31.0.0/20".`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "fault_domain",
					Description: `(Optional) Fault domain. Required and valid only for OCI. Available as of provider version R2.19.3. ### HA`,
				},
				resource.Attribute{
					Name:        "ha_subnet",
					Description: `(Optional) HA Subnet CIDR. Required only if enabling HA for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, OCI, Alibaba Cloud, AWS Top Secret or AWS Secret gateways. Optional for GCP. Setting to empty/unsetting will disable HA. Setting to a valid subnet CIDR will create an HA gateway on the subnet. Example: "10.12.0.0/24".`,
				},
				resource.Attribute{
					Name:        "ha_zone",
					Description: `(Optional) HA Zone. Required if enabling HA for GCP gateway. Optional if enabling HA for Azure gateway. For GCP, setting to empty/unsetting will disable HA and setting to a valid zone will create an HA gateway in the zone. Example: "us-west1-c". For Azure, this is an optional parameter to place the HA gateway in a specific availability zone. Valid values for Azure gateways are in the form "az-n". Example: "az-2". Available for Azure as of provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "ha_insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Transit HA Gateway. Required for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret if ` + "`" + `insane_mode` + "`" + ` is enabled and ` + "`" + `ha_subnet` + "`" + ` is set. Example: AWS: "us-west-1a".`,
				},
				resource.Attribute{
					Name:        "ha_eip",
					Description: `(Optional) Public IP address that you want to assign to the HA peering instance. If no value is given, a new EIP will automatically be allocated. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.`,
				},
				resource.Attribute{
					Name:        "ha_azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP Address resource and its resource group in Azure to be assigned to the HA Transit Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if ` + "`" + `ha_eip` + "`" + ` is set and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.`,
				},
				resource.Attribute{
					Name:        "ha_gw_size",
					Description: `(Optional) HA Gateway Size. Mandatory if enabling HA. Example: "t2.micro".`,
				},
				resource.Attribute{
					Name:        "ha_availability_domain",
					Description: `(Optional) HA gateway availability domain. Required and valid only for OCI. Available as of provider version R2.19.3.`,
				},
				resource.Attribute{
					Name:        "ha_fault_domain",
					Description: `(Optional) HA gateway fault domain. Required and valid only for OCI. Available as of provider version R2.19.3. ### Insane Mode`,
				},
				resource.Attribute{
					Name:        "insane_mode",
					Description: `(Optional) Specify true for [Insane Mode](https://docs.aviatrix.com/HowTos/insane_mode.html) high performance gateway. Insane Mode gateway size must be at least c5 size (AWS, AWSGov, AWS China, AWS Top Secret and AWS Secret) or Standard_D3_v2 (Azure and AzureGov); for GCP only four size are supported: "n1-highcpu-4", "n1-highcpu-8", "n1-highcpu-16" and "n1-highcpu-32". If enabled, you must specify a valid /26 CIDR segment of the VPC to create a new subnet for AWS, Azure, AzureGov, AWSGov, AWS Top Secret and AWS Secret. Only available for AWS, GCP/OCI, Azure, AzureGov, AzureChina, AWSGov, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "insane_mode_az",
					Description: `(Optional) AZ of subnet being created for Insane Mode Transit Gateway. Required for AWS, AWSGov, AWS China, AWS Top Secret or AWS Secret if ` + "`" + `insane_mode` + "`" + ` is enabled. Example: AWS: "us-west-1a". ### SNAT`,
				},
				resource.Attribute{
					Name:        "single_ip_snat",
					Description: `(Optional) Enable "single_ip" mode Source NAT for this container. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "enable_segmentation",
					Description: `(Optional) Enable transit gateway for segmentation. Valid values: true, false. Default: false. ### Advanced Options`,
				},
				resource.Attribute{
					Name:        "connected_transit",
					Description: `(Optional) Specify Connected Transit status. If enabled, it allows spokes to run traffics to other spokes via transit gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_advertise_transit_cidr",
					Description: `(Optional) Switch to enable/disable advertise transit VPC network CIDR for a VGW connection. Available as of R2.6.`,
				},
				resource.Attribute{
					Name:        "bgp_manual_spoke_advertise_cidrs",
					Description: `(Optional) Intended CIDR list to be advertised to external BGP router. Example: "10.2.0.0/16,10.4.0.0/16". Available as of R2.6.`,
				},
				resource.Attribute{
					Name:        "enable_hybrid_connection",
					Description: `(Optional) Sign of readiness for AWS TGW connection. Only supported for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret. Example: false.`,
				},
				resource.Attribute{
					Name:        "enable_firenet",
					Description: `(Optional) Set to true to use gateway for legacy [AWS TGW-based FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_transit_summarize_cidr_to_tgw",
					Description: `(Optional) Enable summarize CIDR to TGW. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby",
					Description: `(Optional) Enables [Active-Standby Mode](https://docs.aviatrix.com/HowTos/transit_advanced.html#active-standby). Available only with HA enabled. Valid values: true, false. Default value: false. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "enable_active_standby_preemptive",
					Description: `(Optional) Enables Preemptive Mode for Active-Standby. Available only with BGP enabled, HA enabled and Active-Standby enabled. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "bgp_polling_time",
					Description: `(Optional) BGP route polling time. Unit is in seconds. Valid values are between 10 and 50. Default value: "50".`,
				},
				resource.Attribute{
					Name:        "bgp_hold_time",
					Description: `(Optional) BGP hold time. Unit is in seconds. Valid values are between 12 and 360. Default value: 180.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) List of AS numbers to populate BGP AP_PATH field when it advertises to VGW or peer devices.`,
				},
				resource.Attribute{
					Name:        "local_as_number",
					Description: `(Optional) Changes the Aviatrix Transit Gateway ASN number before you setup Aviatrix Transit Gateway connection configurations.`,
				},
				resource.Attribute{
					Name:        "bgp_ecmp",
					Description: `(Optional) Enable Equal Cost Multi Path (ECMP) routing for the next hop. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_multi_tier_transit",
					Description: `(Optional) Enable Multi-tier Transit mode on transit gateway. When enabled, transit gateway will propagate routes it receives from its transit peering peer to other transit peering peers. ` + "`" + `local_as_number` + "`" + ` is required. Default value: false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "enable_s2c_rx_balancing",
					Description: `(Optional) Enable S2C receive packet CPU re-balancing on transit gateway. Valid values: true, false. Default value: false. Available in provider version R2.21.2+.`,
				},
				resource.Attribute{
					Name:        "enable_preserve_as_path",
					Description: `(Optional) Enable preserve as_path when advertising manual summary cidrs on transit gateway. Valid values: true, false. Default value: false. Available as of provider version R.2.22.1+. ->`,
				},
				resource.Attribute{
					Name:        "enable_transit_firenet",
					Description: `(Optional) Set to true to use gateway for [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) connection. Valid values: true, false. Default value: false. Available in provider version R2.12+.`,
				},
				resource.Attribute{
					Name:        "lan_vpc_id",
					Description: `(Optional) LAN VPC ID. Only valid when enabling Transit FireNet on GCP. Available as of provider version R2.18.1+.`,
				},
				resource.Attribute{
					Name:        "lan_private_subnet",
					Description: `(Optional) LAN Private Subnet. Only valid when enabling Transit FireNet on GCP. Available as of provider version R2.18.1+.`,
				},
				resource.Attribute{
					Name:        "enable_egress_transit_firenet",
					Description: `(Optional) Enable [Egress Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_workflow.html#b-enable-transit-firenet-on-aviatrix-egress-transit-gateway). Valid values: true, false. Default value: false. Available in provider version R2.16.3+. ->`,
				},
				resource.Attribute{
					Name:        "enable_gateway_load_balancer",
					Description: `(Optional) Enable FireNet interfaces with AWS Gateway Load Balancer. Only valid when ` + "`" + `enable_firenet` + "`" + ` or ` + "`" + `enable_transit_firenet` + "`" + ` are set to true and ` + "`" + `cloud_type` + "`" + ` = 1 (AWS). Currently, AWS Gateway Load Balancer is only supported in AWS regions: us-west-2, us-east-1, eu-west-1, ap-southeast-2 and sa-east-1. Valid values: true or false. Default value: false. Available as of provider version R2.18+. ### BGP over LAN`,
				},
				resource.Attribute{
					Name:        "enable_bgp_over_lan",
					Description: `(Optional) Pre-allocate a network interface(eth4) for "BGP over LAN" functionality. Must be enabled to create a BGP over LAN ` + "`" + `aviatrix_transit_external_device_conn` + "`" + ` resource with this Transit Gateway. Only valid for GCP (4), Azure (8), AzureGov (32) or AzureChina (2048). Valid values: true or false. Default value: false. Available as of provider version R2.18+. Updatable as of provider version 3.0.3+.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_interfaces",
					Description: `(Optional) Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP Transit. Each interface has the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_interfaces",
					Description: `(Optional) Interfaces to run BGP protocol on top of the ethernet interface, to connect to the onprem/remote peer. Only available for GCP Transit HA. Each interface has the following attributes:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC-ID/VNet-Name of cloud provider.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) A VPC Network address range selected from one of the available network ranges. ->`,
				},
				resource.Attribute{
					Name:        "bgp_lan_interfaces_count",
					Description: `(Optional) Number of interfaces that will be created for BGP over LAN enabled Azure transit. Applies on HA Transit as well if enabled. Valid value: 1~5 for FireNet case, 1~7 for Non-FireNet case. Available as of provider version R2.22+. Updatable as of provider version 3.0.3+. ### Encryption`,
				},
				resource.Attribute{
					Name:        "enable_encrypt_volume",
					Description: `(Optional) Enable EBS volume encryption for Gateway. Only supports AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "customer_managed_keys",
					Description: `(Optional and Sensitive) Customer managed key ID. ### Route Customization`,
				},
				resource.Attribute{
					Name:        "customized_spoke_vpc_routes",
					Description: `(Optional) A list of comma-separated CIDRs to be customized for the spoke VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. It applies to all spoke gateways attached to this transit gateway. Example: "10.0.0.0/16,10.2.0.0/16".`,
				},
				resource.Attribute{
					Name:        "filtered_spoke_vpc_routes",
					Description: `(Optional) A list of comma-separated CIDRs to be filtered from the spoke VPC route table. When configured, filtering CIDR(s) or it’s subnet will be deleted from VPC routing tables as well as from spoke gateway’s routing table. It applies to all spoke gateways attached to this transit gateway. Example: "10.2.0.0/16,10.3.0.0/16".`,
				},
				resource.Attribute{
					Name:        "excluded_advertised_spoke_routes",
					Description: `(Optional) A list of comma-separated CIDRs to be advertised to on-prem as 'Excluded CIDR List'. When configured, it inspects all the advertised CIDRs from its spoke gateways and remove those included in the 'Excluded CIDR List'. Example: "10.4.0.0/16,10.5.0.0/16".`,
				},
				resource.Attribute{
					Name:        "customized_transit_vpc_routes",
					Description: `(Optional) A list of CIDRs to be customized for the transit VPC routes. When configured, it will replace all learned routes in VPC routing tables, including RFC1918 and non-RFC1918 CIDRs. To be effective, ` + "`" + `enable_advertise_transit_cidr` + "`" + ` or firewall management access for a Transit FireNet gateway must be enabled. Example: ["10.0.0.0/16", "10.2.0.0/16"]. ### [Learned CIDRs Approval](https://docs.aviatrix.com/HowTos/transit_approval.html) ->`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Switch to enable/disable encrypted transit approval for transit gateway. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "learned_cidrs_approval_mode",
					Description: `(Optional) Learned CIDRs approval mode. Either "gateway" (approval on a per gateway basis) or "connection" (approval on a per connection basis). Default value: "gateway". Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "approved_learned_cidrs",
					Description: `(Optional) A set of approved learned CIDRs. Only valid when ` + "`" + `enable_learned_cidrs_approval` + "`" + ` is set to true. Example: ["10.250.0.0/16", "10.251.0.0/16"]. Available as of provider version R2.21+. ### [Monitor Gateway Subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet) ~>`,
				},
				resource.Attribute{
					Name:        "enable_monitor_gateway_subnets",
					Description: `(Optional) If set to true, the [Monitor Gateway Subnets](https://docs.aviatrix.com/HowTos/gateway.html#monitor-gateway-subnet) feature is enabled. Default value is false. Available in provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "monitor_exclude_list",
					Description: `(Optional) Set of monitored instance ids. Only valid when 'enable_monitor_gateway_subnets' = true. Available in provider version R2.18+. ### [Private OOB](https://docs.aviatrix.com/HowTos/private_oob.html)`,
				},
				resource.Attribute{
					Name:        "enable_private_oob",
					Description: `(Optional) Enable Private OOB feature. Only available for AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "oob_management_subnet",
					Description: `(Optional) OOB management subnet. Required if enabling Private OOB. Example: "11.0.2.0/24".`,
				},
				resource.Attribute{
					Name:        "oob_availability_zone",
					Description: `(Optional) OOB availability zone. Required if enabling Private OOB. Example: "us-west-1a".`,
				},
				resource.Attribute{
					Name:        "ha_oob_management_subnet",
					Description: `(Optional) HA OOB management subnet. Required if enabling Private OOB and HA. Example: "11.0.0.48/28".`,
				},
				resource.Attribute{
					Name:        "ha_oob_availability_zone",
					Description: `(Optional) HA OOB availability zone. Required if enabling Private OOB and HA. Example: "us-west-1b". ### Spot Instance`,
				},
				resource.Attribute{
					Name:        "enable_spot_instance",
					Description: `(Optional) Enable spot instance. NOT supported for production deployment.`,
				},
				resource.Attribute{
					Name:        "spot_price",
					Description: `(Optional) Price for spot instance. NOT supported for production deployment. ### Gateway Upgrade`,
				},
				resource.Attribute{
					Name:        "software_version",
					Description: `(Optional/Computed) The software version of the gateway. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrade can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "image_version",
					Description: `(Optional/Computed) The image version of the gateway. Use ` + "`" + `aviatrix_gateway_image` + "`" + ` data source to programmatically retrieve this value for the desired ` + "`" + `software_version` + "`" + `. If set, we will attempt to update the gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "ha_software_version",
					Description: `(Optional/Computed) The software version of the HA gateway. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the HA gateway upgrade can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "6.5.821". Available as of provider version R2.20.0.`,
				},
				resource.Attribute{
					Name:        "ha_image_version",
					Description: `(Optional/Computed) The image version of the HA gateway. Use ` + "`" + `aviatrix_gateway_image` + "`" + ` data source to programmatically retrieve this value for the desired ` + "`" + `ha_software_version` + "`" + `. If set, we will attempt to update the HA gateway to the specified version if current version is different. If left blank, the gateway upgrades can be managed with the ` + "`" + `aviatrix_controller_config` + "`" + ` resource. Type: String. Example: "hvm-cloudx-aws-022021". Available as of provider version R2.20.0. ### Misc.`,
				},
				resource.Attribute{
					Name:        "allocate_new_eip",
					Description: `(Optional) When value is false, reuse an idle address in Elastic IP pool for this gateway. Otherwise, allocate a new Elastic IP and use it for this gateway. Available in Controller 4.7+. Valid values: true, false. Default: true.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `(Optional) Required when ` + "`" + `allocate_new_eip` + "`" + ` is false. It uses the specified EIP for this gateway. Available in Controller version 4.7+. Only available for AWS, GCP, Azure, OCI, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret.`,
				},
				resource.Attribute{
					Name:        "azure_eip_name_resource_group",
					Description: `(Optional) Name of public IP Address resource and its resource group in Azure to be assigned to the Transit Gateway instance. Example: "IP_Name:Resource_Group_Name". Required if ` + "`" + `allocate_new_eip` + "`" + ` is false and ` + "`" + `cloud_type` + "`" + ` is Azure, AzureGov or AzureChina. Available as of provider version 2.20+.`,
				},
				resource.Attribute{
					Name:        "enable_vpc_dns_server",
					Description: `(Optional) Enable VPC DNS Server for Gateway. Currently only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, Alibaba Cloud, AWS Top Secret and AWS Secret gateways. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Optional) Availability Zone. Only available Azure (8), Azure GOV (32) and Azure CHINA (2048). Must be in the form 'az-n', for example, 'az-2'. Available in provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "enable_jumbo_frame",
					Description: `(Optional) Enable jumbo frames for this transit gateway. Default value is true.`,
				},
				resource.Attribute{
					Name:        "enable_gro_gso",
					Description: `(Optional) Enable GRO/GSO for this transit gateway. Default value is true. Available in provider R3.1.0+.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Map of tags to assign to the gateway. Only available for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina, AWS Top Secret and AWS Secret gateways. Allowed characters vary by cloud type but always include: letters, spaces, and numbers. AWS, AWSGov, AWSChina, AWS Top Secret and AWS Secret allow the use of any character. Azure, AzureGov and AzureChina allows the following special characters: + - = . _ : @. Example: {"key1" = "value1", "key2" = "value2"}.`,
				},
				resource.Attribute{
					Name:        "tunnel_detection_time",
					Description: `(Optional) The IPsec tunnel down detection time for the Transit Gateway in seconds. Must be a number in the range [20-600]. The default value is set by the controller (60 seconds if nothing has been changed).`,
				},
				resource.Attribute{
					Name:        "rx_queue_size",
					Description: `(Optional) Gateway ethernet interface RX queue size. Applies on HA as well if enabled. Once set, can't be deleted or disabled. Available for AWS as of provider version R2.22+.`,
				},
				resource.Attribute{
					Name:        "private_mode_lb_vpc_id",
					Description: `(Optional) VPC ID of Private Mode load balancer. Required when Private Mode is enabled on the Controller. Available in Provider version R2.23+.`,
				},
				resource.Attribute{
					Name:        "private_mode_subnet_zone",
					Description: `(Optional) Availability Zone of the subnet. Required when Private Mode is enabled on the Controller and ` + "`" + `cloud_type` + "`" + ` is AWS or AWSGov. Available in Provider version R2.23+.`,
				},
				resource.Attribute{
					Name:        "ha_private_mode_subnet_zone",
					Description: `(Optional) Availability Zone of the HA subnet. Required when Private Mode is enabled on the Controller and ` + "`" + `cloud_type` + "`" + ` is AWS or AWSGov with HA. Available in Provider version R2.23+. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix transit gateway unique name of HA transit gateway.`,
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
					Name:        "public_ip",
					Description: `Public IP address of the Transit Gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address of the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the transit gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address of the HA transit gateway created.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_security_group_id",
					Description: `HA security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of the HA transit gateway.`,
				},
				resource.Attribute{
					Name:        "lan_interface_cidr",
					Description: `LAN interface CIDR of the transit gateway created (will be used when enabling FQDN Firenet in Azure). Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "ha_lan_interface_cidr",
					Description: `LAN interface CIDR of the HA transit gateway created (will be used when enabling FQDN Firenet in Azure). Available in provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP and Azure. Available as of provider version R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP and Azure. Available as of provider version R2.21.0+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_firenet_interfaces",
					Description: `(Optional) Sign of readiness for FireNet connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina. Example: ["key1:value1","key2:value2"].`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `(Optional) Switch to enable/disable [Active Mesh Mode](https://docs.aviatrix.com/HowTos/activemesh_faq.html) for Transit Gateway. Valid values: true, false. Default value: false.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ha_gw_name",
					Description: `Aviatrix transit gateway unique name of HA transit gateway.`,
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
					Name:        "public_ip",
					Description: `Public IP address of the Transit Gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_public_ip",
					Description: `Public IP address of the HA Transit Gateway.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP address of the transit gateway created.`,
				},
				resource.Attribute{
					Name:        "ha_private_ip",
					Description: `Private IP address of the HA transit gateway created.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_security_group_id",
					Description: `HA security group used for the transit gateway.`,
				},
				resource.Attribute{
					Name:        "cloud_instance_id",
					Description: `Cloud instance ID of the transit gateway.`,
				},
				resource.Attribute{
					Name:        "ha_cloud_instance_id",
					Description: `Cloud instance ID of the HA transit gateway.`,
				},
				resource.Attribute{
					Name:        "lan_interface_cidr",
					Description: `LAN interface CIDR of the transit gateway created (will be used when enabling FQDN Firenet in Azure). Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "ha_lan_interface_cidr",
					Description: `LAN interface CIDR of the HA transit gateway created (will be used when enabling FQDN Firenet in Azure). Available in provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device connection creation. Only supports GCP and Azure. Available as of provider version R2.21.0+.`,
				},
				resource.Attribute{
					Name:        "ha_bgp_lan_ip_list",
					Description: `List of available BGP LAN interface IPs for transit external device HA connection creation. Only supports GCP and Azure. Available as of provider version R2.21.0+. The following arguments are deprecated:`,
				},
				resource.Attribute{
					Name:        "enable_firenet_interfaces",
					Description: `(Optional) Sign of readiness for FireNet connection. Valid values: true, false. Default value: false.`,
				},
				resource.Attribute{
					Name:        "enable_snat",
					Description: `(Optional) Enable Source NAT for this container. Valid values: true, false.`,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Optional) Instance tag of cloud provider. Only supported for AWS, Azure, AzureGov, AWSGov, AWSChina, AzureChina. Example: ["key1:value1","key2:value2"].`,
				},
				resource.Attribute{
					Name:        "enable_active_mesh",
					Description: `(Optional) Switch to enable/disable [Active Mesh Mode](https://docs.aviatrix.com/HowTos/activemesh_faq.html) for Transit Gateway. Valid values: true, false. Default value: false.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_transit_gateway_peering",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Creates and manages Aviatrix transit gateway peerings`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
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
					Description: `(Required) The second transit gateway name to make a peer pair. ### Optional`,
				},
				resource.Attribute{
					Name:        "gateway1_excluded_network_cidrs",
					Description: `(Optional) List of excluded network CIDRs for the first transit gateway.`,
				},
				resource.Attribute{
					Name:        "gateway2_excluded_network_cidrs",
					Description: `(Optional) List of excluded network CIDRs for the second transit gateway.`,
				},
				resource.Attribute{
					Name:        "gateway1_excluded_tgw_connections",
					Description: `(Optional) List of excluded TGW connections for the first transit gateway.`,
				},
				resource.Attribute{
					Name:        "gateway2_excluded_tgw_connections",
					Description: `(Optional) List of excluded TGW connections for the second transit gateway.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path1",
					Description: `(Optional) AS Path Prepend for BGP connection. Can only use the transit's own local AS number, repeated up to 25 times. Applies on transit_gateway_name1. Available in provider version R2.17.2+.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path2",
					Description: `(Optional) AS Path Prepend for BGP connection. Can only use the transit's own local AS number, repeated up to 25 times. Applies on transit_gateway_name2. Available in provider version R2.17.2+.`,
				},
				resource.Attribute{
					Name:        "enable_peering_over_private_network",
					Description: `(Optional) Advanced option. Enable peering over private network. Only appears and applies to when the two Multi-cloud Transit Gateways are each launched in Insane Mode and in a different cloud type. Conflicts with ` + "`" + `enable_insane_mode_encryption_over_internet` + "`" + ` and ` + "`" + `tunnel_count` + "`" + `. Type: Boolean. Default: false. Available in provider version R2.17.1+.`,
				},
				resource.Attribute{
					Name:        "enable_single_tunnel_mode",
					Description: `(Optional) Advanced option. Enable peering with Single-Tunnel mode. Only appears and applies to when the two Multi-cloud Transit Gateways are each launched in Insane Mode and in a different cloud type. Required with ` + "`" + `enable_peering_over_private_network` + "`" + `. Conflicts with ` + "`" + `enable_insane_mode_encryption_over_internet` + "`" + ` and ` + "`" + `tunnel_count` + "`" + `. Type: Boolean. Default: false. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "enable_insane_mode_encryption_over_internet",
					Description: `(Optional) Advanced option. Enable Insane Mode Encryption over Internet. Transit gateways must be in Insane Mode. Currently, only inter-cloud connections between AWS and Azure are supported. Required with valid ` + "`" + `tunnel_count` + "`" + `. Conflicts with ` + "`" + `enable_peering_over_private_network` + "`" + ` and ` + "`" + `enable_single_tunnel_mode` + "`" + `. Type: Boolean. Default: false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "tunnel_count",
					Description: `(Optional) Advanced option. Number of public tunnels. Required with ` + "`" + `enable_insane_mode_encryption_over_internet` + "`" + `. Conflicts with ` + "`" + `enable_peering_over_private_network` + "`" + ` and ` + "`" + `enable_single_tunnel_mode` + "`" + `. Type: Integer. Valid Range: 2-20. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "enable_max_performance",
					Description: `(Optional) Indicates whether the maximum amount of HPE tunnels will be created. Only valid when the two transit gateways are each launched in Insane Mode and in the same cloud type. Default value: true. Available as of provider version R2.22.2+. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_tunnel",
			Category:         "Peering",
			ShortDescription: `Creates and manages Aviatrix Encrypted Peering tunnels`,
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
					Description: `(Required) The second VPC Container name to make a peer pair. ### HA`,
				},
				resource.Attribute{
					Name:        "enable_ha",
					Description: `(Optional) Enable this attribute if peering-HA is enabled on the gateways. Valid values: true, false. Default value: false. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
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
					Description: `(Computed) Name of the peering link. ## Import`,
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
					Description: `(Computed) Name of the peering link. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vgw_conn",
			Category:         "Multi-Cloud Transit",
			ShortDescription: `Manages the connection between the Aviatrix Transit Gateway to VGW`,
			Description:      ``,
			Keywords: []string{
				"multi",
				"cloud",
				"transit",
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
					Description: `(Required) VPC ID where the Transit Gateway is located. Example: AWS: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "bgp_vgw_id",
					Description: `(Required) ID of AWS VGW that will be used for this connection. Example: "vgw-abcd1234".`,
				},
				resource.Attribute{
					Name:        "bgp_vgw_account",
					Description: `(Required) Cloud Account used to create the AWS VGW that will be used for this connection. Example: "dev-account-1".`,
				},
				resource.Attribute{
					Name:        "bgp_vgw_region",
					Description: `(Required) Region of AWS VGW that will be used for this connection. Example: "us-east-1".`,
				},
				resource.Attribute{
					Name:        "bgp_local_as_num",
					Description: `(Required) BGP Local ASN (Autonomous System Number). Integer between 1-4294967294. Example: "65001". ### Optional`,
				},
				resource.Attribute{
					Name:        "enable_learned_cidrs_approval",
					Description: `(Optional) Enable learned CIDRs approval for the connection. Requires the transit_gateway's 'learned_cidrs_approval_mode' attribute be set to 'connection'. Valid values: true, false. Default value: false. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "manual_bgp_advertised_cidrs",
					Description: `(Optional) Configure manual BGP advertised CIDRs for this connection. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "enable_event_triggered_ha",
					Description: `(Optional) Enable Event Triggered HA. Default value: false. Valid values: true or false. Available as of provider version R2.19+.`,
				},
				resource.Attribute{
					Name:        "prepend_as_path",
					Description: `(Optional) Connection AS Path Prepend customized by specifying AS PATH for a BGP connection. Available as of provider version R2.19.2. The following arguments are deprecated:`,
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
			Type:             "aviatrix_aviatrix_vpc",
			Category:         "Useful Tools",
			ShortDescription: `Creates and manages Aviatrix-created VPCs`,
			Description:      ``,
			Keywords: []string{
				"useful",
				"tools",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_type",
					Description: `(Required) Type of cloud service provider, requires an integer value. Currently only AWS(1), GCP(4), Azure(8), OCI(16), AzureGov(32), AWSGov(256), AWSChina(1024), AzureChina(2048), Alibaba Cloud(8192) are supported.`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `(Required) This parameter represents the name of a Cloud-Account in Aviatrix controller.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the VPC to be created.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Region of cloud provider.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) VPC CIDR.`,
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
					Description: `Name of this subnet. ### Advanced Options`,
				},
				resource.Attribute{
					Name:        "subnet_size",
					Description: `(Optional) Subnet size. Only supported for AWS, Azure provider. Example: 24. Available in provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "num_of_subnet_pairs",
					Description: `(Optional) Number of public subnet and private subnet pair created. Only supported for AWS, Azure provider. Example: 1. Available in provider version R2.17+.`,
				},
				resource.Attribute{
					Name:        "enable_private_oob_subnet",
					Description: `(Optional) Switch to enable private oob subnet. Only supported for AWS, AWSGov and AWSChina providers. Valid values: true, false. Default value: false. Available as of provider version R2.18+.`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Optional) The name of an existing resource group or a new resource group to be created for the Azure VNet. A new resource group will be created if left blank. Only available for Azure, AzureGov and AzureChina providers. Available as of provider version R2.19+. ### Misc.`,
				},
				resource.Attribute{
					Name:        "aviatrix_transit_vpc",
					Description: `(Optional) Specify whether it is an [Aviatrix Transit VPC](https://docs.aviatrix.com/HowTos/create_vpc.html#aviatrix-transit-vpc) to be used for [Transit Network](https://docs.aviatrix.com/HowTos/transitvpc_faq.html) or [TGW](https://docs.aviatrix.com/HowTos/tgw_faq.html) solutions.`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `(Optional) Specify whether it is an Aviatrix FireNet VPC to be used for [Aviatrix FireNet](https://docs.aviatrix.com/HowTos/firewall_network_faq.html) and [Transit FireNet](https://docs.aviatrix.com/HowTos/transit_firenet_faq.html) solutions.`,
				},
				resource.Attribute{
					Name:        "enable_native_gwlb",
					Description: `(Optional) Enable Native AWS Gateway Load Balancer for FireNet Function. Only valid with cloud_type = 1 (AWS).`,
				},
				resource.Attribute{
					Name:        "private_mode_subnets",
					Description: `(Optional) Switch to only launch private subnets. Only available when Private Mode is enabled on the Controller. Only AWS, Azure, AzureGov and AWSGov are supported. Available in Provider version R2.23+. ->`,
				},
				resource.Attribute{
					Name:        "aviatrix_firenet_vpc",
					Description: `If you are using/ upgraded to Aviatrix Terraform Provider R1.8+, and a VPC resource was originally created with a provider version <R1.8, you must do 'terraform refresh' to update and apply the attribute’s default value (false) into the state file. ## Attribute Reference In addition to all arguments above, the following attributes are exported: ->`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC to be created.`,
				},
				resource.Attribute{
					Name:        "azure_vnet_resource_id",
					Description: `Azure VNet resource ID.`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.`,
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
					Description: `ID of this subnet.`,
				},
				resource.Attribute{
					Name:        "private_subnets",
					Description: `List of private subnet of the VPC(AWS, Azure) to be created.`,
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
					Description: `ID of this subnet.`,
				},
				resource.Attribute{
					Name:        "public_subnets",
					Description: `List of public subnet of the VPC(AWS, Azure) to be created.`,
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
					Description: `ID of this subnet.`,
				},
				resource.Attribute{
					Name:        "availability_domains",
					Description: `List of OCI availability domains.`,
				},
				resource.Attribute{
					Name:        "fault_domains",
					Description: `List of OCI fault domains. ->`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `If created as a FireNet VPC, four public subnets will be created in the following order: subnet for firewall-mgmt in the first zone, subnet for ingress-egress in the first zone, subnet for firewall-mgmt in the second zone, and subnet for ingress-egress in the second zone. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vpc_id",
					Description: `ID of the VPC to be created.`,
				},
				resource.Attribute{
					Name:        "azure_vnet_resource_id",
					Description: `Azure VNet resource ID.`,
				},
				resource.Attribute{
					Name:        "route_tables",
					Description: `List of route table ids associated with this VPC. Only populated for AWS, AWSGov and Azure VPC.`,
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
					Description: `ID of this subnet.`,
				},
				resource.Attribute{
					Name:        "private_subnets",
					Description: `List of private subnet of the VPC(AWS, Azure) to be created.`,
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
					Description: `ID of this subnet.`,
				},
				resource.Attribute{
					Name:        "public_subnets",
					Description: `List of public subnet of the VPC(AWS, Azure) to be created.`,
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
					Description: `ID of this subnet.`,
				},
				resource.Attribute{
					Name:        "availability_domains",
					Description: `List of OCI availability domains.`,
				},
				resource.Attribute{
					Name:        "fault_domains",
					Description: `List of OCI fault domains. ->`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `If created as a FireNet VPC, four public subnets will be created in the following order: subnet for firewall-mgmt in the first zone, subnet for ingress-egress in the first zone, subnet for firewall-mgmt in the second zone, and subnet for ingress-egress in the second zone. ## Import`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vpn_cert_download",
			Category:         "OpenVPN",
			ShortDescription: `Creates and Manages Aviatrix VPN Users`,
			Description:      ``,
			Keywords: []string{
				"openvpn",
				"vpn",
				"cert",
				"download",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "download_enabled",
					Description: `(Optional) Whether the VPN Certificate download is enabled. Supported Values: "true", "false".`,
				},
				resource.Attribute{
					Name:        "saml_endpoints",
					Description: `(Optional) List of SAML endpoint names for which the downloading should be enabled . Currently, only a single endpoint is supported. Example: ["saml_endpoint_1"]. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vpn_profile",
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
					Description: `(Optional) Base policy rule of the profile to be added. Enter "allow_all" or "deny_all", based on whether you want a whitelist or blacklist. ### Policy Options`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Optional) New security policy for the profile. Each policy has the following attributes:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Should be the opposite of the base rule for correct behavior. Valid values for action: "allow", "deny".`,
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
					Description: `(Required) CIDR to be allowed or denied. Valid values for target: IPv4 CIDRs. Example: "10.30.0.0/16". ### Misc.`,
				},
				resource.Attribute{
					Name:        "manage_user_attachment",
					Description: `(Optional) This parameter is a switch used to determine whether or not to manage VPN user attachments to the VPN profile using this resource. If this is set to false, attachment must be managed using the`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) List of VPN users to attach to this profile. This should be set to null if ` + "`" + `manage_user_attachment` + "`" + ` is set to false. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vpn_user",
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
					Description: `(Optional) VPC ID of Aviatrix VPN gateway. Used together with ` + "`" + `gw_name` + "`" + `. Example: "vpc-abcd1234".`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `(Optional) If ELB is enabled, this will be the name of the ELB, else it will be the name of the Aviatrix VPN gateway. Used together with ` + "`" + `vpc_id` + "`" + `. Example: "gw1".`,
				},
				resource.Attribute{
					Name:        "dns_name",
					Description: `(Optional) FQDN of a DNS based VPN service such as GeoVPN or UDP load balancer. Example: "vpn.testuser.com".`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) VPN user name. Example: "user".`,
				},
				resource.Attribute{
					Name:        "user_email",
					Description: `(Optional) VPN user's email. Example: "abc@xyz.com". ### SAML`,
				},
				resource.Attribute{
					Name:        "saml_endpoint",
					Description: `(Optional) This is the name of the SAML endpoint to which the user is to be associated. This is required if adding user to a SAML gateway/LB. ### Misc.`,
				},
				resource.Attribute{
					Name:        "manage_user_attachment",
					Description: `(Optional) This parameter is a switch to determine whether or not to manage VPN user attachments to the VPN profile using this resource. If this is set to false, attachment must be managed using the`,
				},
				resource.Attribute{
					Name:        "profiles",
					Description: `(Optional) List of VPN profiles for user to attach to. This should be set to null if ` + "`" + `manage_user_attachment` + "`" + ` is set to false. ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_vpn_user_accelerator",
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
					Description: `(Required) Name of ELB to be added to VPN User Accelerator. Example: "Aviatrix-vpc-abcd2134". ## Import`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "aviatrix_aviatrix_web_group",
			Category:         "Secured Networking",
			ShortDescription: `Creates and manages an Aviatrix Web Group`,
			Description:      ``,
			Keywords: []string{
				"secured",
				"networking",
				"web",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Web Group.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Required) Block containing match expressions to filter the Web Group.`,
				},
				resource.Attribute{
					Name:        "match_expressions",
					Description: `(Required) List of match expressions. The Web Group will be a union of all resources matched by each ` + "`" + `match_expressions` + "`" + `.` + "`" + `match_expressions` + "`" + ` blocks cannot be empty.`,
				},
				resource.Attribute{
					Name:        "snifilter",
					Description: `(Optional) - Server name indicator this expression matches. ` + "`" + `snifilter` + "`" + ` cannot be used with ` + "`" + `urlfilter` + "`" + ` filters in the same ` + "`" + `match_expressions` + "`" + ` block.`,
				},
				resource.Attribute{
					Name:        "urlfilter",
					Description: `(Optional) - URL address this expression matches. ` + "`" + `urlfilter` + "`" + ` cannot be used with ` + "`" + `snifilter` + "`" + ` filters in the same ` + "`" + `match_expressions` + "`" + ` block. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the Web Group. ## Import`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the Web Group. ## Import`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"aviatrix_aviatrix_account":                                        0,
		"aviatrix_aviatrix_account_user":                                   1,
		"aviatrix_aviatrix_aws_guard_duty":                                 2,
		"aviatrix_aviatrix_aws_peer":                                       3,
		"aviatrix_aviatrix_aws_tgw":                                        4,
		"aviatrix_aviatrix_aws_tgw_connect":                                5,
		"aviatrix_aviatrix_aws_tgw_connect_peer":                           6,
		"aviatrix_aviatrix_aws_tgw_directconnect":                          7,
		"aviatrix_aviatrix_aws_tgw_intra_domain_inspection":                8,
		"aviatrix_aviatrix_aws_tgw_network_domain":                         9,
		"aviatrix_aviatrix_aws_tgw_peering":                                10,
		"aviatrix_aviatrix_aws_tgw_peering_domain_conn":                    11,
		"aviatrix_aviatrix_aws_tgw_transit_gateway_attachment":             12,
		"aviatrix_aviatrix_aws_tgw_vpc_attachment":                         13,
		"aviatrix_aviatrix_aws_tgw_vpn_conn":                               14,
		"aviatrix_aviatrix_azure_peer":                                     15,
		"aviatrix_aviatrix_azure_spoke_native_peering":                     16,
		"aviatrix_aviatrix_azure_vng_conn":                                 17,
		"aviatrix_aviatrix_centralized_transit_firenet":                    18,
		"aviatrix_aviatrix_cloudn_registration":                            19,
		"aviatrix_aviatrix_cloudn_transit_gateway_attachment":              20,
		"aviatrix_aviatrix_cloudwatch_agent":                               21,
		"aviatrix_aviatrix_controller_access_allow_list_config":            22,
		"aviatrix_aviatrix_controller_bgp_max_as_limit_config":             23,
		"aviatrix_aviatrix_controller_cert_domain_config":                  24,
		"aviatrix_aviatrix_controller_config":                              25,
		"aviatrix_aviatrix_controller_email_config":                        26,
		"aviatrix_aviatrix_controller_email_exception_notification_config": 27,
		"aviatrix_aviatrix_controller_gateway_keepalive_config":            28,
		"aviatrix_aviatrix_controller_private_mode_config":                 29,
		"aviatrix_aviatrix_controller_private_oob":                         30,
		"aviatrix_aviatrix_controller_security_group_management_config":    31,
		"aviatrix_aviatrix_copilot_association":                            32,
		"aviatrix_aviatrix_copilot_security_group_management_config":       33,
		"aviatrix_aviatrix_datadog_agent":                                  34,
		"aviatrix_aviatrix_device_interface_config":                        35,
		"aviatrix_aviatrix_distributed_firewalling_config":                 36,
		"aviatrix_aviatrix_distributed_firewalling_intra_vpc":              37,
		"aviatrix_aviatrix_distributed_firewalling_policy_list":            38,
		"aviatrix_aviatrix_dns_profile":                                    39,
		"aviatrix_aviatrix_edge_csp":                                       40,
		"aviatrix_aviatrix_edge_csp_ha":                                    41,
		"aviatrix_aviatrix_edge_equinix":                                   42,
		"aviatrix_aviatrix_edge_equinix_ha":                                43,
		"aviatrix_aviatrix_edge_neo":                                       44,
		"aviatrix_aviatrix_edge_neo_device_onboarding":                     45,
		"aviatrix_aviatrix_edge_neo_ha":                                    46,
		"aviatrix_aviatrix_edge_spoke":                                     47,
		"aviatrix_aviatrix_edge_spoke_external_device_conn":                48,
		"aviatrix_aviatrix_edge_spoke_transit_attachment":                  49,
		"aviatrix_aviatrix_filebeat_forwarder":                             50,
		"aviatrix_aviatrix_firenet":                                        51,
		"aviatrix_aviatrix_firewall":                                       52,
		"aviatrix_aviatrix_firewall_instance":                              53,
		"aviatrix_aviatrix_firewall_instance_association":                  54,
		"aviatrix_aviatrix_firewall_management_access":                     55,
		"aviatrix_aviatrix_firewall_policy":                                56,
		"aviatrix_aviatrix_firewall_tag":                                   57,
		"aviatrix_aviatrix_fqdn":                                           58,
		"aviatrix_aviatrix_fqdn_global_config":                             59,
		"aviatrix_aviatrix_fqdn_pass_through":                              60,
		"aviatrix_aviatrix_fqdn_tag_rule":                                  61,
		"aviatrix_aviatrix_gateway":                                        62,
		"aviatrix_aviatrix_gateway_certificate_config":                     63,
		"aviatrix_aviatrix_gateway_dnat":                                   64,
		"aviatrix_aviatrix_gateway_snat":                                   65,
		"aviatrix_aviatrix_geo_vpn":                                        66,
		"aviatrix_aviatrix_global_vpc_excluded_instance":                   67,
		"aviatrix_aviatrix_global_vpc_tagging_settings":                    68,
		"aviatrix_aviatrix_link_hierarchy":                                 69,
		"aviatrix_aviatrix_netflow_agent":                                  70,
		"aviatrix_aviatrix_periodic_ping":                                  71,
		"aviatrix_aviatrix_private_mode_lb":                                72,
		"aviatrix_aviatrix_private_mode_multicloud_endpoint":               73,
		"aviatrix_aviatrix_proxy_config":                                   74,
		"aviatrix_aviatrix_qos_class":                                      75,
		"aviatrix_aviatrix_qos_policy_list":                                76,
		"aviatrix_aviatrix_rbac_group":                                     77,
		"aviatrix_aviatrix_rbac_group_access_account_attachment":           78,
		"aviatrix_aviatrix_rbac_group_permission_attachment":               79,
		"aviatrix_aviatrix_rbac_group_user_attachment":                     80,
		"aviatrix_aviatrix_remote_syslog":                                  81,
		"aviatrix_aviatrix_saml_endpoint":                                  82,
		"aviatrix_aviatrix_segmentation_network_domain":                    83,
		"aviatrix_aviatrix_segmentation_network_domain_association":        84,
		"aviatrix_aviatrix_segmentation_network_domain_connection_policy":  85,
		"aviatrix_aviatrix_site2cloud":                                     86,
		"aviatrix_aviatrix_site2cloud_ca_cert_tag":                         87,
		"aviatrix_aviatrix_sla_class":                                      88,
		"aviatrix_aviatrix_smart_group":                                    89,
		"aviatrix_aviatrix_splunk_logging":                                 90,
		"aviatrix_aviatrix_spoke_external_device_conn":                     91,
		"aviatrix_aviatrix_spoke_gateway":                                  92,
		"aviatrix_aviatrix_spoke_gateway_subnet_group":                     93,
		"aviatrix_aviatrix_spoke_ha_gateway":                               94,
		"aviatrix_aviatrix_spoke_transit_attachment":                       95,
		"aviatrix_aviatrix_sumologic_forwarder":                            96,
		"aviatrix_aviatrix_traffic_classifier":                             97,
		"aviatrix_aviatrix_trans_peer":                                     98,
		"aviatrix_aviatrix_transit_external_device_conn":                   99,
		"aviatrix_aviatrix_transit_firenet_policy":                         100,
		"aviatrix_aviatrix_transit_gateway":                                101,
		"aviatrix_aviatrix_transit_gateway_peering":                        102,
		"aviatrix_aviatrix_tunnel":                                         103,
		"aviatrix_aviatrix_vgw_conn":                                       104,
		"aviatrix_aviatrix_vpc":                                            105,
		"aviatrix_aviatrix_vpn_cert_download":                              106,
		"aviatrix_aviatrix_vpn_profile":                                    107,
		"aviatrix_aviatrix_vpn_user":                                       108,
		"aviatrix_aviatrix_vpn_user_accelerator":                           109,
		"aviatrix_aviatrix_web_group":                                      110,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
