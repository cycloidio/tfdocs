package dome9

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dome9_attach_iam_safe",
			Category:         "Resources",
			ShortDescription: `Attach IAM safe to AWS cloud account.`,
			Description:      ``,
			Keywords: []string{
				"attach",
				"iam",
				"safe",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_cloud_account_id",
					Description: `(Required) AWS cloud account to attach IAM safe to it.`,
				},
				resource.Attribute{
					Name:        "aws_group_arn",
					Description: `(Required) AWS group arn.`,
				},
				resource.Attribute{
					Name:        "aws_policy_arn",
					Description: `(Required) AWS policy arn. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode. ## Import Cloud account IAM safe can be imported; use ` + "`" + `<AWS CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_attach_iam_safe_re.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mode",
					Description: `Mode. ## Import Cloud account IAM safe can be imported; use ` + "`" + `<AWS CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_attach_iam_safe_re.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_aws_security_group",
			Category:         "Resources",
			ShortDescription: `Creates AWS Security Group in Dome9`,
			Description:      ``,
			Keywords: []string{
				"aws",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dome9_security_group_name",
					Description: `(Required) Name of the Security Group.`,
				},
				resource.Attribute{
					Name:        "dome9_cloud_account_id",
					Description: `(Required) Cloud account id in Dome9.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Security Group description.`,
				},
				resource.Attribute{
					Name:        "aws_region_id",
					Description: `(Optional) AWS region, in AWS format (e.g., "us-east-1"); default is us_east_1.`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `(Optional) Indicates the Security Group is in Protected mode.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) VPC id for VPC containing the Security Group.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `(Optional) Security Group VPC name.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Security Group tags.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Security Group services. ### Security Group services ` + "`" + `services` + "`" + ` has the these arguments:`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `(Required) inbound service.`,
				},
				resource.Attribute{
					Name:        "outbound",
					Description: `(Required) outbound service. The configuration of inbound and outbound is:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Service description.`,
				},
				resource.Attribute{
					Name:        "protocol_type",
					Description: `(Required) Service protocol type. Select from "ALL", "HOPOPT", "ICMP", "IGMP", "GGP", "IPV4", "ST", "TCP", "CBT", "EGP", "IGP", "BBN_RCC_MON", "NVP2", "PUP", "ARGUS", "EMCON", "XNET", "CHAOS", "UDP", "MUX", "DCN_MEAS", "HMP", "PRM", "XNS_IDP", "TRUNK1", "TRUNK2", "LEAF1", "LEAF2", "RDP", "IRTP", "ISO_TP4", "NETBLT", "MFE_NSP", "MERIT_INP", "DCCP", "ThreePC", "IDPR", "XTP", "DDP", "IDPR_CMTP", "TPplusplus", "IL", "IPV6", "SDRP", "IPV6_ROUTE", "IPV6_FRAG", "IDRP", "RSVP", "GRE", "DSR", "BNA", "ESP", "AH", "I_NLSP", "SWIPE", "NARP", "MOBILE", "TLSP", "SKIP", "ICMPV6", "IPV6_NONXT", "IPV6_OPTS", "CFTP", "SAT_EXPAK", "KRYPTOLAN", "RVD", "IPPC", "SAT_MON", "VISA", "IPCV", "CPNX", "CPHB", "WSN", "PVP", "BR_SAT_MON", "SUN_ND", "WB_MON", "WB_EXPAK", "ISO_IP", "VMTP", "SECURE_VMTP", "VINES", "TTP", "NSFNET_IGP", "DGP", "TCF", "EIGRP", "OSPFIGP", "SPRITE_RPC", "LARP", "MTP", "AX25", "IPIP", "MICP", "SCC_SP", "ETHERIP", "ENCAP", "GMTP", "IFMP", "PNNI", "PIM", "ARIS", "SCPS", "QNX", "AN", "IPCOMP", "SNP", "COMPAQ_PEER", "IPX_IN_IP", "VRRP", "PGM", "L2TP", "DDX", "IATP", "STP", "SRP", "UTI", "SMP", "SM", "PTP", "ISIS", "FIRE", "CRTP", "CRUDP", "SSCOPMCE", "IPLT", "SPS", "PIPE", "SCTP", "FC", "RSVP_E2E_IGNORE", "MOBILITY_HEADER", "UDPLITE", "MPLS_IN_IP", "MANET", "HIP", "SHIM6", "WESP" or "ROHC".`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Service type (port).`,
				},
				resource.Attribute{
					Name:        "open_for_all",
					Description: `(Optional) Is open for all.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Service scope which has the following configuration:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) scope type.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required) scope data. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_name",
					Description: `AWS cloud account name.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `Security Group external id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_name",
					Description: `AWS cloud account name.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `Security Group external id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_azure_security_group",
			Category:         "Resources",
			ShortDescription: `Creates azure security group in Dome9`,
			Description:      ``,
			Keywords: []string{
				"azure",
				"security",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dome9_security_group_name",
					Description: `(Required) Name of the security group.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region can be one of the following: ` + "`" + `centralus` + "`" + `, ` + "`" + `eastus` + "`" + `, ` + "`" + `eastus2` + "`" + `, ` + "`" + `usgovlowa` + "`" + `, ` + "`" + `usgovvirginia` + "`" + `, ` + "`" + `northcentralus` + "`" + `, ` + "`" + `southcentralus` + "`" + `, ` + "`" + `westus` + "`" + `, ` + "`" + `westus2` + "`" + `, ` + "`" + `westcentralus` + "`" + `, ` + "`" + `northeurope` + "`" + `, ` + "`" + `westeurope` + "`" + `, ` + "`" + `eastasia` + "`" + `, ` + "`" + `southeastasia` + "`" + `, ` + "`" + `japaneast` + "`" + `, ` + "`" + `japanwest` + "`" + `, ` + "`" + `brazilsouth` + "`" + `, ` + "`" + `australiaeast` + "`" + `, ` + "`" + `australiasoutheast` + "`" + `, ` + "`" + `centralindia` + "`" + `, ` + "`" + `southindia` + "`" + `, ` + "`" + `westindia` + "`" + `, ` + "`" + `chinaeast` + "`" + `, ` + "`" + `chinanorth` + "`" + `, ` + "`" + `canadacentral` + "`" + `, ` + "`" + `canadaeast` + "`" + `, ` + "`" + `germanycentral` + "`" + `, ` + "`" + `germanynortheast` + "`" + `, ` + "`" + `koreacentral` + "`" + `, ` + "`" + `uksouth` + "`" + `, ` + "`" + `ukwest` + "`" + `, ` + "`" + `koreasout` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Required) Azure resource group name.`,
				},
				resource.Attribute{
					Name:        "dome9_cloud_account_id",
					Description: `(Required) Cloud account id in Dome9.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Security group description.`,
				},
				resource.Attribute{
					Name:        "is_tamper_protected",
					Description: `(Optional) Is security group tamper protected.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Security group tags list of ` + "`" + `key` + "`" + `, ` + "`" + `value` + "`" + `:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag value.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `(Optional) Security group services.`,
				},
				resource.Attribute{
					Name:        "outbound",
					Description: `(Optional) Security group services. The configuration of inbound and outbound is:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Service description.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `(Required) Service priority (a number between 100 and 4096)`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Service access (Allow / Deny).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) Service protocol (UDP / TCP / ANY).`,
				},
				resource.Attribute{
					Name:        "source_port_ranges",
					Description: `(Required) Source port ranges.`,
				},
				resource.Attribute{
					Name:        "destination_port_ranges",
					Description: `(Required) Destination port ranges.`,
				},
				resource.Attribute{
					Name:        "source_scopes",
					Description: `(Required) List of source scopes for the service (CIDR / IP List / Tag):`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) scope type.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required) scope data.`,
				},
				resource.Attribute{
					Name:        "destination_scopes",
					Description: `(Required) List of destination scopes for the service (CIDR / IP List / Tag).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) scope type.`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `(Required) scope data.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Gets or sets the default security rules of network security group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "external_security_group_id",
					Description: `Azure external security group id.`,
				},
				resource.Attribute{
					Name:        "cloud_account_name",
					Description: `Azure cloud account name.`,
				},
				resource.Attribute{
					Name:        "last_updated_by_dome9",
					Description: `Last updated by dome9. ## Import The security group can be imported; use ` + "`" + `<SESCURITY GROUP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_azure_security_group.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_security_group_id",
					Description: `Azure external security group id.`,
				},
				resource.Attribute{
					Name:        "cloud_account_name",
					Description: `Azure cloud account name.`,
				},
				resource.Attribute{
					Name:        "last_updated_by_dome9",
					Description: `Last updated by dome9. ## Import The security group can be imported; use ` + "`" + `<SESCURITY GROUP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_azure_security_group.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_aws",
			Category:         "Resources",
			ShortDescription: `Onboard AWS cloud account`,
			Description:      ``,
			Keywords: []string{
				"cloudaccount",
				"aws",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of AWS account in Dome9`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The information needed for Dome9 System in order to connect to the AWS cloud account`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `(Optional) The Organizational Unit that this cloud account will be attached to ### Credentials ` + "`" + `credentials` + "`" + ` has the following arguments:`,
				},
				resource.Attribute{
					Name:        "ARN",
					Description: `(Required) AWS Role ARN (to be assumed by Dome9)`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) The AWS role External ID (Dome9 will have to use this secret in order to assume the role)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The cloud account onboarding method. Set to "RoleBased". ### Network security configuration ` + "`" + `net_sec` + "`" + ` has the these arguments:`,
				},
				resource.Attribute{
					Name:        "Regions",
					Description: `(Required) list of the supported regions, and their configuration:`,
				},
				resource.Attribute{
					Name:        "new_group_behavior",
					Description: `(Required) The network security configuration. Select "ReadOnly", "FullManage", or "Reset".`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region, in AWS format (e.g., "us-east-1") ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("AWS").`,
				},
				resource.Attribute{
					Name:        "external_account_number",
					Description: `The AWS account number.`,
				},
				resource.Attribute{
					Name:        "is_fetching_suspended",
					Description: `Fetching suspending status.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9.`,
				},
				resource.Attribute{
					Name:        "full_protection",
					Description: `The protection mode for existing security groups in the account.`,
				},
				resource.Attribute{
					Name:        "allow_read_only",
					Description: `The AWS cloud account operation mode. true for "Full-Manage", false for "Readonly".`,
				},
				resource.Attribute{
					Name:        "net_sec",
					Description: `The network security configuration for the AWS cloud account. If not given, sets to default value.`,
				},
				resource.Attribute{
					Name:        "IAM_safe",
					Description: `IAM safe entity details`,
				},
				resource.Attribute{
					Name:        "AWS_group_ARN",
					Description: `AWS group ARN`,
				},
				resource.Attribute{
					Name:        "AWS_policy_ARN",
					Description: `AWS policy ARN`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode`,
				},
				resource.Attribute{
					Name:        "restricted_IAM_entities",
					Description: `Restricted IAM safe entities, which have the following fields:`,
				},
				resource.Attribute{
					Name:        "roles_ARNs",
					Description: `Restricted IAM safe entities roles ARNs`,
				},
				resource.Attribute{
					Name:        "users_ARNs",
					Description: `Restricted IAM safe entities users ARNs ## Import AWS cloud account can be imported; use ` + "`" + `<AWS CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_AWS.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("AWS").`,
				},
				resource.Attribute{
					Name:        "external_account_number",
					Description: `The AWS account number.`,
				},
				resource.Attribute{
					Name:        "is_fetching_suspended",
					Description: `Fetching suspending status.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9.`,
				},
				resource.Attribute{
					Name:        "full_protection",
					Description: `The protection mode for existing security groups in the account.`,
				},
				resource.Attribute{
					Name:        "allow_read_only",
					Description: `The AWS cloud account operation mode. true for "Full-Manage", false for "Readonly".`,
				},
				resource.Attribute{
					Name:        "net_sec",
					Description: `The network security configuration for the AWS cloud account. If not given, sets to default value.`,
				},
				resource.Attribute{
					Name:        "IAM_safe",
					Description: `IAM safe entity details`,
				},
				resource.Attribute{
					Name:        "AWS_group_ARN",
					Description: `AWS group ARN`,
				},
				resource.Attribute{
					Name:        "AWS_policy_ARN",
					Description: `AWS policy ARN`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode`,
				},
				resource.Attribute{
					Name:        "restricted_IAM_entities",
					Description: `Restricted IAM safe entities, which have the following fields:`,
				},
				resource.Attribute{
					Name:        "roles_ARNs",
					Description: `Restricted IAM safe entities roles ARNs`,
				},
				resource.Attribute{
					Name:        "users_ARNs",
					Description: `Restricted IAM safe entities users ARNs ## Import AWS cloud account can be imported; use ` + "`" + `<AWS CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_AWS.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_azure",
			Category:         "Resources",
			ShortDescription: `Onboard Azure cloud account`,
			Description:      ``,
			Keywords: []string{
				"cloudaccount",
				"azure",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Azure account in Dome9`,
				},
				resource.Attribute{
					Name:        "operation_mode",
					Description: `(Required) Dome9 operation mode for the Azure account ("Read-Only" or "Managed")`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) The Azure subscription id for account`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The Azure tenant id`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Azure account id`,
				},
				resource.Attribute{
					Name:        "client_password",
					Description: `(Required) Password for account`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `(Optional) Organizational Unit that this cloud account will be attached to ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Azure cloud account`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("Azure")`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name ## Import Azure cloud account can be imported; use ` + "`" + `<Azure CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_Azure.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Azure cloud account`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("Azure")`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date the account was onboarded to Dome9`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name ## Import Azure cloud account can be imported; use ` + "`" + `<Azure CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_Azure.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_gcp",
			Category:         "Resources",
			ShortDescription: `Onboard GCP cloud account`,
			Description:      ``,
			Keywords: []string{
				"cloudaccount",
				"gcp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Google account name in Dome9`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) Project ID`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) Private key ID`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) Private key`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) GCP client email`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Client id`,
				},
				resource.Attribute{
					Name:        "client_x509_cert_url",
					Description: `(Required) client x509 certificate URL`,
				},
				resource.Attribute{
					Name:        "gsuite_user",
					Description: `(Optional) The Gsuite user`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `(Optional) The domain name`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `(Optional) Organizational Unit that this cloud account will be attached to ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the GCP cloud account`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (gcp).`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name. ## Import GCP cloud account can be imported; use ` + "`" + `<GCP CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_gcp.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the GCP cloud account`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (gcp).`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name. ## Import GCP cloud account can be imported; use ` + "`" + `<GCP CLOUD ACCOUNT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_cloudaccount_gcp.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_continuous_compliance_notification",
			Category:         "Resources",
			ShortDescription: `Creates continuous compliance notification in Dome9`,
			Description:      ``,
			Keywords: []string{
				"continuous",
				"compliance",
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The cloud account id in Dome9.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the notification. at least one of ` + "`" + `alerts_console` + "`" + `, ` + "`" + `scheduled_report` + "`" + `, or ` + "`" + `change_detection` + "`" + ` must be included`,
				},
				resource.Attribute{
					Name:        "alerts_console",
					Description: `(Optional) send findings (also) to the Dome9 web app alerts console (Boolean); default is False.`,
				},
				resource.Attribute{
					Name:        "scheduled_report",
					Description: `Scheduled email report notification block:`,
				},
				resource.Attribute{
					Name:        "email_sending_state",
					Description: `send schedule report of findings by email; can be "Enabled" or "Disabled". if ` + "`" + `email_sending_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "schedule_data",
					Description: `Schedule details:`,
				},
				resource.Attribute{
					Name:        "cron_expression",
					Description: `the schedule to issue the email report (in cron expression format)`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of report; can be "Detailed", "Summary", "FullCsv" or "FullCsvZip"`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `comma-separated list of email recipients`,
				},
				resource.Attribute{
					Name:        "change_detection",
					Description: `Send changes in findings options:`,
				},
				resource.Attribute{
					Name:        "email_sending_stat",
					Description: `send email report of changes in findings; can be "Enabled" or "Disabled". if ` + "`" + `email_sending_stat` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "email_data",
					Description: `Email notification details:`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `comma-separated list of email recipients`,
				},
				resource.Attribute{
					Name:        "email_per_finding_sending_state",
					Description: `send separate email notification for each finding; can be "Enabled" or "Disabled" if ` + "`" + `email_per_finding_sending_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "email_per_finding_data",
					Description: `Email per finding notification details:`,
				},
				resource.Attribute{
					Name:        "recipients",
					Description: `comma-separated list of email recipients`,
				},
				resource.Attribute{
					Name:        "notification_output_format",
					Description: `(Required) format of JSON block for finding; can be "JsonWithFullEntity", "JsonWithBasicEntity", or "PlainText".`,
				},
				resource.Attribute{
					Name:        "sns_sending_state",
					Description: `send by AWS SNS for each new finding; can be "Enabled" or "Disabled". if ` + "`" + `sns_sending_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "sns_data",
					Description: `SNS notification details:`,
				},
				resource.Attribute{
					Name:        "sns_topic_arn",
					Description: `SNS topic ARN`,
				},
				resource.Attribute{
					Name:        "sns_output_format",
					Description: `SNS output format; can be "JsonWithFullEntity", "JsonWithBasicEntity", or "PlainText".`,
				},
				resource.Attribute{
					Name:        "external_ticket_creating_state",
					Description: `send each finding to an external ticketing system; can be "Enabled" or "Disabled". if ` + "`" + `external_ticket_creating_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "ticketing_system_data",
					Description: `Ticketing system details:`,
				},
				resource.Attribute{
					Name:        "system_type",
					Description: `system type; can be "ServiceOne", "Jira", or "PagerDuty"`,
				},
				resource.Attribute{
					Name:        "should_close_tickets",
					Description: `ticketing system should close tickets when resolved (bool)`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `ServiceNow domain name (ServiceNow only)`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `user name (ServiceNow only)`,
				},
				resource.Attribute{
					Name:        "pass",
					Description: `password (ServiceNow only)`,
				},
				resource.Attribute{
					Name:        "project_key",
					Description: `project key (Jira) or API Key (PagerDuty)`,
				},
				resource.Attribute{
					Name:        "issue_type",
					Description: `issue type (Jira)`,
				},
				resource.Attribute{
					Name:        "webhook_integration_state",
					Description: `send findings to an HTTP endpoint (webhook); can be "Enabled" or "Disabled". if ` + "`" + `webhook_integration_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "webhook_data",
					Description: `Webhook data block supports:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `HTTP endpoint URL`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `HTTP method, "Post" by default.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `authentication method; "NoAuth" by default`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `username in endpoint system`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `password in endpoint system`,
				},
				resource.Attribute{
					Name:        "format_type",
					Description: `format for JSON block for finding; can be "Basic" or "ServiceNow"`,
				},
				resource.Attribute{
					Name:        "aws_security_hub_integration_state",
					Description: `send findings to AWS Secure Hub; can be "Enabled" or "Disabled". if ` + "`" + `aws_security_hub_integration_state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "aws_security_hub_integration",
					Description: `AWS security hub integration block supports:`,
				},
				resource.Attribute{
					Name:        "external_account_id",
					Description: `(Required) external account id`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) AWS region ` + "`" + `gcp_security_command_center_integration` + "`" + ` is a change_detection option`,
				},
				resource.Attribute{
					Name:        "gcp_security_command_center_integration",
					Description: `GCP security command center details`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `send findings to the GCP Security Command Center; can be "Enabled" or "Disabled" if ` + "`" + `state` + "`" + ` is Enabled, the following must be included:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `GCP Project id`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `GCP Source id ## Import The notification can be imported; use ` + "`" + `<NOTIFICATION ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_continuouscompliance_notification.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_continuous_compliance_policy",
			Category:         "Resources",
			ShortDescription: `Creates continuous compliance policies in Dome9`,
			Description:      ``,
			Keywords: []string{
				"continuous",
				"compliance",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `(Required) The cloud account id.`,
				},
				resource.Attribute{
					Name:        "external_account_id",
					Description: `(Required) The account number.`,
				},
				resource.Attribute{
					Name:        "bundle_id",
					Description: `(Required) The bundle id for the bundle that will be used in the policy.`,
				},
				resource.Attribute{
					Name:        "cloud_account_type",
					Description: `(Required) The cloud account provider ("AWS", "Azure", "Google").`,
				},
				resource.Attribute{
					Name:        "notification_ids",
					Description: `(Required) The notification policy id's for the policy [list]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the compliance policy. ## Import The policy can be imported; use ` + "`" + `<POLICY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_continuous_compliance_policy.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the compliance policy. ## Import The policy can be imported; use ` + "`" + `<POLICY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_continuous_compliance_policy.test 00000000-0000-0000-0000-000000000000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_iam_safe_entity",
			Category:         "Resources",
			ShortDescription: `Protect cloud accounts that are managed by Dome9. Control access to them with targeted short-term authorizations (involving the Dome9 mobile app).`,
			Description:      ``,
			Keywords: []string{
				"iam",
				"safe",
				"entity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "protection_mode",
					Description: `(Required) Protection mode; can be "Protect", "ProtectWithElevation".`,
				},
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) Entity type to protect; can be "User", "Role".`,
				},
				resource.Attribute{
					Name:        "aws_cloud_account_id",
					Description: `(Required) AWS cloud account id to protect.`,
				},
				resource.Attribute{
					Name:        "entity_name",
					Description: `(Required) AWS IAM user or role name to protect.`,
				},
				resource.Attribute{
					Name:        "dome9_users_id_to_protect",
					Description: `(Optional) When ProtectWithElevation mode selected, dome9 users ids must be provided.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Can be one of the following: ` + "`" + `Unattached` + "`" + `, ` + "`" + `Attached` + "`" + ` or ` + "`" + `Restricted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "attached_dome9_users",
					Description: `List of users in protect with elevation mode.`,
				},
				resource.Attribute{
					Name:        "exists_in_aws",
					Description: `Is exist in aws.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Role or User arn.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `Can be one of the following: ` + "`" + `Unattached` + "`" + `, ` + "`" + `Attached` + "`" + ` or ` + "`" + `Restricted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "attached_dome9_users",
					Description: `List of users in protect with elevation mode.`,
				},
				resource.Attribute{
					Name:        "exists_in_aws",
					Description: `Is exist in aws.`,
				},
				resource.Attribute{
					Name:        "arn",
					Description: `Role or User arn.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_iplist",
			Category:         "Resources",
			ShortDescription: `Create IP lists in Dome9`,
			Description:      ``,
			Keywords: []string{
				"iplist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the IP list in Dome9`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the list (what it represents); defaults to empty string`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `(Optional) the individual IP addresses for the list; defaults to empty list ### Items The ` + "`" + `items` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) IP address`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `(Optional) Comment ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `IP list Id ## Import IP list can be imported; use ` + "`" + `<IP LIST ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_iplist.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `IP list Id ## Import IP list can be imported; use ` + "`" + `<IP LIST ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_iplist.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_organizational_unit",
			Category:         "Resources",
			ShortDescription: `Create organizational unit in Dome9`,
			Description:      ``,
			Keywords: []string{
				"organizational",
				"unit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the organizational unit in Dome9.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The organizational unit parent ID. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Organizational unit Id`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Dome9 internal account ID.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Organizational Unit full path (IDs).`,
				},
				resource.Attribute{
					Name:        "path_str",
					Description: `Organizational Unit full path (names).`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Organizational Unit creation time.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Organizational Unit update time.`,
				},
				resource.Attribute{
					Name:        "aws_cloud_accounts_count",
					Description: `Number of AWS cloud accounts in the Organizational Unit.`,
				},
				resource.Attribute{
					Name:        "azure_cloud_accounts_count",
					Description: `Number of Azure cloud accounts in the Organizational Unit.`,
				},
				resource.Attribute{
					Name:        "google_cloud_accounts_count",
					Description: `Number of GCP cloud accounts in the Organizational Unit.`,
				},
				resource.Attribute{
					Name:        "aws_aggregated_cloud_accounts_count",
					Description: `Number of AWS cloud accounts in the Organizational Unit and its children.`,
				},
				resource.Attribute{
					Name:        "azure_aggregate_cloud_accounts_count",
					Description: `Number of Azure cloud accounts in the Organizational Unit and its children.`,
				},
				resource.Attribute{
					Name:        "google_aggregate_cloud_accounts_count",
					Description: `Number of GCP cloud accounts in the Organizational Unit and its children.`,
				},
				resource.Attribute{
					Name:        "is_root",
					Description: `Is Organizational Unit root.`,
				},
				resource.Attribute{
					Name:        "is_parent_root",
					Description: `Is the parent of Organizational Unit root. ## Import Organizational unit can be imported; use ` + "`" + `<ORGANIZATIONAL UNIT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_organizational_unit.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Organizational unit Id`,
				},
				resource.Attribute{
					Name:        "account_id",
					Description: `Dome9 internal account ID.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Organizational Unit full path (IDs).`,
				},
				resource.Attribute{
					Name:        "path_str",
					Description: `Organizational Unit full path (names).`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Organizational Unit creation time.`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `Organizational Unit update time.`,
				},
				resource.Attribute{
					Name:        "aws_cloud_accounts_count",
					Description: `Number of AWS cloud accounts in the Organizational Unit.`,
				},
				resource.Attribute{
					Name:        "azure_cloud_accounts_count",
					Description: `Number of Azure cloud accounts in the Organizational Unit.`,
				},
				resource.Attribute{
					Name:        "google_cloud_accounts_count",
					Description: `Number of GCP cloud accounts in the Organizational Unit.`,
				},
				resource.Attribute{
					Name:        "aws_aggregated_cloud_accounts_count",
					Description: `Number of AWS cloud accounts in the Organizational Unit and its children.`,
				},
				resource.Attribute{
					Name:        "azure_aggregate_cloud_accounts_count",
					Description: `Number of Azure cloud accounts in the Organizational Unit and its children.`,
				},
				resource.Attribute{
					Name:        "google_aggregate_cloud_accounts_count",
					Description: `Number of GCP cloud accounts in the Organizational Unit and its children.`,
				},
				resource.Attribute{
					Name:        "is_root",
					Description: `Is Organizational Unit root.`,
				},
				resource.Attribute{
					Name:        "is_parent_root",
					Description: `Is the parent of Organizational Unit root. ## Import Organizational unit can be imported; use ` + "`" + `<ORGANIZATIONAL UNIT ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_organizational_unit.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_role",
			Category:         "Resources",
			ShortDescription: `Create role in Dome9`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Dome9 role name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Dome9 role description.`,
				},
				resource.Attribute{
					Name:        "permit_rulesets",
					Description: `Is permitted permit rulesets (Optional) .`,
				},
				resource.Attribute{
					Name:        "permit_notifications",
					Description: `Is permitted permit notifications (Optional) .`,
				},
				resource.Attribute{
					Name:        "permit_policies",
					Description: `Is permitted permit policies (Optional) .`,
				},
				resource.Attribute{
					Name:        "permit_alert_actions",
					Description: `Is permitted permit alert actions (Optional) .`,
				},
				resource.Attribute{
					Name:        "permit_on_boarding",
					Description: `Is permitted permit on boarding (Optional) .`,
				},
				resource.Attribute{
					Name:        "cross_account_access",
					Description: `(Optional) Cross account access.`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Optional) Create permission list.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Access permission list ([SRL](#SRL) Type).`,
				},
				resource.Attribute{
					Name:        "view",
					Description: `(Optional) View permission list ([SRL](#SRL) Type).`,
				},
				resource.Attribute{
					Name:        "manage",
					Description: `(Optional) Manage permission list ([SRL](#SRL) Type). ### SRL`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Accepted values: AWS, Azure, GCP, OrganizationalUnit.`,
				},
				resource.Attribute{
					Name:        "main_id",
					Description: `(Optional) Cloud Account or Organizational Unit ID.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Accepted values: "us_east_1", "us_west_1", "eu_west_1", "ap_southeast_1", "ap_northeast_1", "us_west_2", "sa_east_1", "ap_southeast_2", "eu_central_1", "ap_northeast_2", "ap_south_1", "us_east_2", "ca_central_1", "eu_west_2", "eu_west_3", "eu_north_1".`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Optional) AWS Security Group ID.`,
				},
				resource.Attribute{
					Name:        "traffic",
					Description: `(Optional) Accepted values: "All Traffic", "All Services".`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_ruleset",
			Category:         "Resources",
			ShortDescription: `Create ruleset in Dome9`,
			Description:      ``,
			Keywords: []string{
				"ruleset",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ruleset in Dome9.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the ruleset (what it represents); defaults to empty string.`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `(Required) Cloud vendor that the ruleset is associated with.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `(Optional) Language of the rules; defaults to 'en' (English). ### Rules The ` + "`" + `rules` + "`" + ` supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Rule name`,
				},
				resource.Attribute{
					Name:        "logic",
					Description: `(Optional) Rule GSL logic. This is the text of the rule, using Dome9 GSL syntax`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) Rule severity`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Rule description`,
				},
				resource.Attribute{
					Name:        "compliance_tag",
					Description: `(Optional) A reference to a compliance standard ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Ruleset Id ## Import Ruleset can be imported; use ` + "`" + `<RULE SET ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_rule_set.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Ruleset Id ## Import Ruleset can be imported; use ` + "`" + `<RULE SET ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import dome9_rule_set.test 00000 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"dome9_attach_iam_safe":                    0,
		"dome9_aws_security_group":                 1,
		"dome9_azure_security_group":               2,
		"dome9_cloudaccount_aws":                   3,
		"dome9_cloudaccount_azure":                 4,
		"dome9_cloudaccount_gcp":                   5,
		"dome9_continuous_compliance_notification": 6,
		"dome9_continuous_compliance_policy":       7,
		"dome9_iam_safe_entity":                    8,
		"dome9_iplist":                             9,
		"dome9_organizational_unit":                10,
		"dome9_role":                               11,
		"dome9_ruleset":                            12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
