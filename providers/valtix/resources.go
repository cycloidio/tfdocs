package valtix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "valtix_address_object",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing Address Objects that can be used as a Source or Destination (Src/Dest) in a Policy Ruleset Rule, or as a Target Backend Address (Reverse Proxy Target) in a Reverse Proxy Service Object.

The Address Object is used in a Policy Ruleset Rule to define the segmentation policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A list of IPs/CIDRs <br><br>For an example, see [STATIC (Source Destination) Example](#static-source-destination-example) #### STATIC (Reverse Proxy Target) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A list of one or more IPs/FQDNs`,
				},
				resource.Attribute{
					Name:        "backend_address",
					Description: `(Required) This argument must be set to ` + "`" + `true` + "`" + ` <br><br>For an example, see [STATIC (Reverse Proxy Target) Example](#static-reverse-proxy-target-example) #### DYNAMIC_APPLICATIONS (Reverse Proxy Target) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Optional) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC/VNet ID to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Region to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) The Resource Group to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Required) A single block that represents a Tag key-value pair associated with one or more Application Load Balancers. ` + "`" + `tag_list` + "`" + ` structure [defined below](#tag-list). The ` + "`" + `tag_list` + "`" + ` is used to dynamically associate the IPs or FQDNs of the set of Application Load Balancers that match the Tag key-value pair.`,
				},
				resource.Attribute{
					Name:        "backend_address",
					Description: `(Required) This argument must be set to ` + "`" + `true` + "`" + ` <br><br>For an example, see [DYNAMIC_APPLICATIONS (Reverse Proxy Target) Example](#dynamic_applications-reverse-proxy-target-example) #### DYNAMIC_VPC (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `vpc_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The Region to restrict the scope of the ` + "`" + `vpc_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) The Resource Group to restrict the scope of the ` + "`" + `vpc_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC/VNet ID is used to dynamically associate all CIDRs of the VPC/VNet <br><br>For an example, see [DYNAMIC_VPC (Source Destination) Example](#dynamic_vpc-source-destination-example) #### DYNAMIC_SECURITY_GROUP (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `security_group_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The Region to restrict the scope of the ` + "`" + `security_group_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC/VNet ID to restrict the scope of the ` + "`" + `security_group_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) The Resource Group to restrict the scope of the ` + "`" + `security_group_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) The Security Group ID used to dynamically associate all IPs of the Security Group <br><br>For an example, see [DYNAMIC_SECURITY_GROUP (Source Destination) Example](#dynamic_security_group-source-destination-example) #### DYNAMIC_INSTANCE (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `instance_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The Region to restrict the scope of the ` + "`" + `instance_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC/VNet ID to restrict the scope of the ` + "`" + `instance_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) The Resource Group to restrict the scope of the ` + "`" + `instance_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) The Instance ID used to dynamically associate all IPs of the Instance <br><br>For an example, see [DYNAMIC_INSTANCE (Source Destination) Example](#dynamic_instance-source-destination-example) #### DYNAMIC_SUBNET (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `subnet_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The Region to restrict the scope of the ` + "`" + `subnet_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC ID to restrict the scope of the ` + "`" + `subnet_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) The Resource Group to restrict the scope of the ` + "`" + `subnet_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) The Subnet ID used to dynamically associate the CIDR of the Subnet <br><br>For an example, see [DYNAMIC_SUBNET (Source Destination) Example](#dynamic_subnet-source-destination-example) #### DYNAMIC_USER_DEFINED_TAG (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Optional) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The Region to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional) The VPC ID to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) The Resource Group to restrict the scope of the ` + "`" + `tag_list` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tag_list",
					Description: `(Required) The set of one or more blocks where each block represents a Tag key-value pair and Resource type. The set of blocks is operated on by an AND operator. ` + "`" + `tag_list` + "`" + ` structure [defined below](#tag-list). The set of ` + "`" + `tag_list` + "`" + ` arguments are used to dynamically associate the IPs or CIDRs of Instances, Subnets and VPCs/Vnets that match the set of Tag key-value pairs. <br><br>For an example, see [DYNAMIC_USER_DEFINED_TAG (Source Destination) Example](#dynamic_user_defined_tag-source-destination-example) #### DYNAMIC_SERVICE_ENDPOINTS (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The name of the CSP account onboarded into Valtix to restrict the scope of the ` + "`" + `service_endpoint_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The Region to restrict the scope of the ` + "`" + `service_endpoint_name` + "`" + ``,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) The Service Endpoint used to dynamically associate all CIDRs of the Service Endpoint <br><br>For an example, see [DYNAMIC_SERVICE_ENDPOINTS (Source Destination) Example](#dynamic_service_endpoints-source-destination-example) #### GEO_IP (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) A list of Geo IPs defined by their Country name. A full list of Country names can be obtained from the [GeoNames Countries](https://www.geonames.org/countries/) site.`,
				},
				resource.Attribute{
					Name:        "match_xff",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Defines whether the IP information available in the X-Forwarded-For HTTP request header should be evaluated. When not specified, the default value is ` + "`" + `false` + "`" + `. (Even though this argument is optional, it is recommended to specify a value explicitly, as the default value may change in the future). <br><br>For an example, see [GEO_IP (Source Destination) Example](#geo_ip-source-destination-example) #### GROUP (Source Destination) Arguments`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Address Object`,
				},
				resource.Attribute{
					Name:        "address_group_ids",
					Description: `(Required) A list of ` + "`" + `address_id` + "`" + ` to be grouped together <br><br>For an example, see [GROUP (Source Destination) Example](#group-source-destination-example) ### Tag List A ` + "`" + `tag_list` + "`" + ` block representing a Tag key-value pair requires the following arguments:`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Required) The Tag Key specified in a Tag key-value pair`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(Required) The Tag Value specified in a Tag key-value pair`,
				},
				resource.Attribute{
					Name:        "address_id",
					Description: `The ID of the Address Object that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address_id",
					Description: `The ID of the Address Object that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_alert_profile",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the alert profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) One of ` + "`" + `SlackWebHook` + "`" + `, ` + "`" + `PagerDutyEventApi` + "`" + `, or ` + "`" + `ServiceNowWebHook` + "`" + ``,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required for ServiceNow and Pagerduty) Key`,
				},
				resource.Attribute{
					Name:        "integration_url",
					Description: `(Required for ServiceNow and Slack) Webhook URL or ServiceNow Url ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_alert_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Alert Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "alert_profile",
					Description: `(Required) Alert Profile ID`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) One of ` + "`" + `Type_Inventory` + "`" + `, ` + "`" + `Type_SystemLogs` + "`" + `,`,
				},
				resource.Attribute{
					Name:        "sub_type",
					Description: `(Required) For Type_Inventory, valid sub_type is ` + "`" + `SubType_InventoryGuardRails` + "`" + `. For Type_SystemLogs, valid sub_type is ` + "`" + `SubType_SystemLogsGateway` + "`" + ` or ` + "`" + `SubType_SystemLogsAccount` + "`" + ``,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) One of ` + "`" + `Info` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `High` + "`" + `. For SystemLogs ` + "`" + `Critical` + "`" + ` and ` + "`" + `Warning` + "`" + ` are other severity types.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Required) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_certificate",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing Certificates used in Decryption Profiles

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the certificate`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the certificate`,
				},
				resource.Attribute{
					Name:        "certificate_type",
					Description: `(Required) ["IMPORT_CONTENTS"](#import_contents), ["AWS_KMS"](#aws_kms), ["AWS_SECRET"](#aws_secret), ["AZURE_KEY_VAULT_SECRET"](#azure_key_vault_secret), ["GCP_SECRET"](#gcp_secret)`,
				},
				resource.Attribute{
					Name:        "certificate_body",
					Description: `(Required) Certificate content. Provide a string with the entire content or if located in a file use ` + "`" + `file("filename")` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional) Certificate chain content. Provide a string with the entire content or if it is located in a file use ` + "`" + `file("filename")` + "`" + `. This is optional if the certificate has a separate chain that is not part of the certificate body. ### IMPORT_CONTENTS`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required for IMPORT_CONTENTS) Certificate content. Provide a string with the entire content or if it is located in a file use ` + "`" + `file("filename")` + "`" + ` ### AWS_KMS`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key that is stored in AWS Key Management System (KMS), provide the csp_account_name (AWS) where the private key will be stored. This is the friendly name defined in Valtix for the on-boarded CSP account.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the private key will be stored`,
				},
				resource.Attribute{
					Name:        "private_key_cipher_text",
					Description: `(Required) AWS KMS encrypted cipher text ### AWS_SECRET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key stored in AWS Secrets Manager, provide the csp_account_name (AWS) where the private key will be stored. This is the friendly name defined in Valtix for the on-boarded CSP account.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the private key will be stored`,
				},
				resource.Attribute{
					Name:        "aws_secret_name",
					Description: `(Required) AWS Secrets Manager key name that contains the private key ### AZURE_KEY_VAULT_SECRET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key stored in Azure Key Vault, provide the csp_account_name (Azure) where the private key will be stored. This is the friendly name defined in Valtix for the on-boarded CSP account.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the private key will be stored`,
				},
				resource.Attribute{
					Name:        "azure_key_vault_name",
					Description: `(Required) Azure Key Vault name where the private key will be stored`,
				},
				resource.Attribute{
					Name:        "azure_key_vault_secret_name",
					Description: `(Required) Azure Key Vault secret name that contains the private key ### GCP_SECRET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key stored in GCP Secrets Manager, provide the csp_account_name (GCP) where the private key will be stored. This is the friendly name defined in Valtix for the on-boarded CSP account.`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) GCP Secrets Manager key name that contains the private key`,
				},
				resource.Attribute{
					Name:        "secret_version",
					Description: `(Optional) GCP Secrets Manager key version ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the certificate that can be referenced in other resources (e.g. valtix_profile_decryption)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Name of the certificate that can be referenced in other resources (e.g. valtix_profile_decryption)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_cloud_account",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing Cloud Accounts and Subscriptions on-boarded into Valtix.  The Cloud Account defines the credentials used for the Valtix Controller to discover asset inventory and traffic, and orchestrate and manage the deployment of Valtix Gateways.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud Account on the Valtix Console. Must contain only alphanumeric, hyphens or underscore characters and not exceed 100 characters`,
				},
				resource.Attribute{
					Name:        "csp_type",
					Description: `(Required) Defines the Cloud Service Provider. Must be ` + "`" + `GCP` + "`" + `, ` + "`" + `AWS` + "`" + ` or ` + "`" + `AZURE` + "`" + ` ### AWS Arguments`,
				},
				resource.Attribute{
					Name:        "aws_credentials_type",
					Description: `(AWS - Required) must be ` + "`" + `AWS_IAM_ROLE` + "`" + ``,
				},
				resource.Attribute{
					Name:        "aws_iam_role",
					Description: `(AWS - Required) IAM role ARN used to deploy and manage Valtix in your cloud account`,
				},
				resource.Attribute{
					Name:        "aws_account_number",
					Description: `(AWS - Required) AWS account number`,
				},
				resource.Attribute{
					Name:        "aws_iam_role_external_id",
					Description: `(AWS - Required) External Id for trust relationship`,
				},
				resource.Attribute{
					Name:        "aws_inventory_iam_role",
					Description: `(AWS - Optional) IAM Role ARN used by CloudWatch Event Rule to post inventory events to the Valtix Controller`,
				},
				resource.Attribute{
					Name:        "inventory_monitoring",
					Description: `Enable inventory monitoring (can be repeated multiple times). See [Inventory Monitoring](#inventory-monitoring) for details. ### Azure Arguments`,
				},
				resource.Attribute{
					Name:        "azure_directory_id",
					Description: `(Azure - Required) Azure Active Directory Id (Tenant Id)`,
				},
				resource.Attribute{
					Name:        "azure_subscription_id",
					Description: `(Azure - Required) Azure Subscription Id where the Valtix gateway instances are deployed`,
				},
				resource.Attribute{
					Name:        "azure_application_id",
					Description: `(Azure - Required) Azure Application Id that's used as credentials (along with the secret) to manage Azure account/subscription`,
				},
				resource.Attribute{
					Name:        "azure_client_secret",
					Description: `(Azure - Required) Azure client secret for the above application`,
				},
				resource.Attribute{
					Name:        "inventory_monitoring",
					Description: `Enable inventory monitoring (can be repeated multiple times). See [Inventory Monitoring](#inventory-monitoring) for details. ### GCP Arguments`,
				},
				resource.Attribute{
					Name:        "gcp_credentials_file",
					Description: `(GCP - Required) Service account credentials key file created for the Valtix controller access.`,
				},
				resource.Attribute{
					Name:        "inventory_monitoring",
					Description: `Enable inventory monitoring (can be repeated multiple times). See [Inventory Monitoring](#inventory-monitoring) for details. ## Inventory Monitoring`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions to enable and monitor inventory`,
				},
				resource.Attribute{
					Name:        "refresh_interval",
					Description: `Interval in minutes where the inventory is refreshed`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_external_id",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the External ID ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID that must be used in the trust settings of the cross account IAM role`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID that must be used in the trust settings of the cross account IAM role`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_gateway",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing Valtix Gateways

~> **Note on valtix_gateway resource creation**
[valtix_cloud_account](../valtix_cloud_account), [valtix_policy_rule_set](../valtix_policy_rule_set)
must be defined before valtix_gateway can be created

~> **Note on valtix_gateway resource management**
Except for the arguments of ` + "`" + `description` + "`" + `, ` + "`" + `gateway_image` + "`" + ` and ` + "`" + `gateway_state` + "`" + `, changes to any other arguments are prevented.  Any of these argument changes require replacing the Gateway, which would result in a service disruption.  In order to change these arguments, you must do one of the following:  1) Set the ` + "`" + `gateway_state` + "`" + ` to ` + "`" + `INACTIVE` + "`" + `, apply the change, then set the ` + "`" + `gateway_state` + "`" + ` to ` + "`" + `ACTIVE` + "`" + ` or 2) Destroy the ` + "`" + `valtix_gateway` + "`" + ` resource, apply the change, and (re)create the ` + "`" + `valtix_gateway` + "`" + ` resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Valtix Gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Valtix Gateway`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The CSP account where the Gateway will be deployed.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Must be one of:`,
				},
				resource.Attribute{
					Name:        "gateway_image",
					Description: `(Required) Example ` + "`" + `2.11-08` + "`" + `. This is the Valtix image version to be deployed for this Gateway. A list of applicable Gateway image versions is available from the Valtix Portal (ADMINISTRATION -> Management -> System -> Gateway Images). Please view the [Valtix Release Recommendation](https://docs.valtix.com/releases/recommendation/) for the recommended Gateway release or contact Valtix Support.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(AWS, Azure - Required) "EDGE" or "HUB". Look into product documentation for different deployment modes. This argument is not supported for GCP and must not be used.`,
				},
				resource.Attribute{
					Name:        "security_type",
					Description: `(Optional) ` + "`" + `INGRESS` + "`" + ` or ` + "`" + `EGRESS` + "`" + `. If not specified, the default is ` + "`" + `INGRESS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gateway_state",
					Description: `(Optional) Specifies the state of the Valtix Gateway. When set to ` + "`" + `ACTIVE` + "`" + `, the Gateway will be active and operational. When set to ` + "`" + `INACTIVE` + "`" + `, the Gateway will be disabled and not operational. If not specified, the default is ` + "`" + `ACTIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_rule_set_id",
					Description: `(Required) Rule set id of valtix_policy_rule_set.`,
				},
				resource.Attribute{
					Name:        "ssh_key_pair",
					Description: `(AWS - Required) SSH key pair name that's already in your AWS account`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Azure - Required) Contents of SSH public key`,
				},
				resource.Attribute{
					Name:        "gcp_service_account_email",
					Description: `(GCP - Required) This is the GCP Gateway service account email, that provides permissions for the Valtix Gateway to integrate with other GCP project resources such as Secrets Manager and storage buckets.`,
				},
				resource.Attribute{
					Name:        "aws_iam_role_firewall",
					Description: `(AWS - Required) This is the IAM role that's assigned to the Valtix firewall instances.`,
				},
				resource.Attribute{
					Name:        "azure_user_identity_id",
					Description: `(Azure - Optional) User assigned identity This is the IAM role that's assigned to the Valtix firewall instances.`,
				},
				resource.Attribute{
					Name:        "azure_resource_group",
					Description: `(Azure - Required) Azure resource group name used for all Valtix Gateway resources`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the Valtix Gateway is deployed.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID where the Valtix Gateway is deployed and is used for data traffic to be inspected. This must be either the VPC where you apps run or the shared services VPC that's peered (or hub via Transit Gateway) to other spoke (app) VPCs. Please note that for HUB mode, this vpc_id must refer to`,
				},
				resource.Attribute{
					Name:        "aws_gateway_lb",
					Description: `(Optional only for AWS HUB mode) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If attribute is ` + "`" + `true` + "`" + `, this will deploy AWS Gateway using AWS Gateway Load Balancer. This is only for EGRESS Gateway that will support both East-West and Egress traffic.`,
				},
				resource.Attribute{
					Name:        "mgmt_vpc_id",
					Description: `(GCP - Required) GCP VPC ID where the management interface of the Valtix Gateway is attached.`,
				},
				resource.Attribute{
					Name:        "mgmt_security_group",
					Description: `(Required except for AWS HUB mode) Security group ID for management traffic or GCP network tag to be used to define GCP firewall rules for Valtix firewall instances to communicate with the Valtix controller. This must allow all outbound access from Valtix management interface`,
				},
				resource.Attribute{
					Name:        "datapath_security_group",
					Description: `(Required except for AWS HUB mode) Security group ID for the datapath traffic (application traffic) or GCP network tag to be used to define GCP firewall rules for application traffic to pass through the Valtix Gateway. This must allow traffic to the ports that are defined as services on the Valtix controller`,
				},
				resource.Attribute{
					Name:        "min_instances",
					Description: `(Optional) Minimum number of instances per zone (default 1)`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `(Optional) Maximum number of instances per zone (default 1)`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `(Optional) Port number that NLB uses to monitor the instances (default 65534). A rule must be configured on the datapath_security_group to allow traffic to this port.`,
				},
				resource.Attribute{
					Name:        "log_profile",
					Description: `(Optional) Log Profile ID`,
				},
				resource.Attribute{
					Name:        "packet_capture_profile",
					Description: `(Optional) Packet Profile ID`,
				},
				resource.Attribute{
					Name:        "diagnostics_profile",
					Description: `(Optional) Diagnostics Profile ID`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Optional) Gateway settings. This block can be repeated multiple times. Please check [this section](#gateway-settings) for the structure.`,
				},
				resource.Attribute{
					Name:        "instance_details",
					Description: `(Required - EDGE Mode) This block is only needed when deploying a Gateway in EDGE mode. This block should not be used when deploying a Gateway in HUB mode. For EDGE mode deployment, the block can be repeated multiple times for deploying Gateway instances in multiple Availability Zones. Look below for the [structure](#instance-details) of this block. In EDGE mode, at least 1 block must be provided. ## Instance Details This section is not required for AWS/Azure HUB mode as instance details are obtained from service VPC referenced in vpc_id attribute`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the availability zone where the Valtix Gateway instance(s) are deployed`,
				},
				resource.Attribute{
					Name:        "mgmt_subnet",
					Description: `(Required) Specifies the VPC subnet ID used for management traffic where the Valtix Gateway instance(s) are deployed for this availability zone.`,
				},
				resource.Attribute{
					Name:        "datapath_subnet",
					Description: `(Required) Specifies the VPC subnet ID used for data traffic where the Valtix Gateway instance(s) are deployed for this availability zone. ## Gateway Settings Gateway settings define a list of settings that applies to the given Gateway ### To enable EBS encryption for the Gateway instances using default KMS key ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "gateway.aws.ebs.encryption.key.default" value = true } ` + "`" + `` + "`" + `` + "`" + ` ### To enable EBS encryption for the Gateway instances using specified KMS key ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "gateway.aws.ebs.encryption.key.customer_key" value = "<KMS key ID>" } ` + "`" + `` + "`" + `` + "`" + ` ### To add a list of custom tags to the Gateway instances ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "custom_tags" value = "[{\"key\":\"customer_key1\",\"value\":\"customer_value1\"},{\"key\":\"customer_key2\",\"value\":\"customer_value2\"}]" } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "gateway_gwlb_endpoints",
					Description: `AWS Gateway Load Balancer endpoints created in each of the AZs. It's in the following format ` + "`" + `` + "`" + `` + "`" + `hcl gateway_gwlb_endpoints { endpoint_id = "vpce-047c749fc6f7e0c0d" network_interface_id = "eni-017eacdb23d2ebaf4" subnet_id = "subnet-0d61750e97caafd9d" } gateway_gwlb_endpoints { endpoint_id = "vpce-0707fa3f03c5064a7" network_interface_id = "eni-020464bd838461bca" subnet_id = "subnet-0fd61e07f200224f1" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gateway_endpoint",
					Description: `For the Ingress Gateway, shows the NLB DNS/IP of the Valtix Gateway. This must be used as an endpoint for your application and Valtix proxies the traffic received on this endpoint to the target application configured`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway_gwlb_endpoints",
					Description: `AWS Gateway Load Balancer endpoints created in each of the AZs. It's in the following format ` + "`" + `` + "`" + `` + "`" + `hcl gateway_gwlb_endpoints { endpoint_id = "vpce-047c749fc6f7e0c0d" network_interface_id = "eni-017eacdb23d2ebaf4" subnet_id = "subnet-0d61750e97caafd9d" } gateway_gwlb_endpoints { endpoint_id = "vpce-0707fa3f03c5064a7" network_interface_id = "eni-020464bd838461bca" subnet_id = "subnet-0fd61e07f200224f1" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gateway_endpoint",
					Description: `For the Ingress Gateway, shows the NLB DNS/IP of the Valtix Gateway. This must be used as an endpoint for your application and Valtix proxies the traffic received on this endpoint to the target application configured`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rule_set",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Policy Ruleset.  A Policy Ruleset is a list of Policy Rules that define the segmentation and security policy for protecting traffic.  The Policy Ruleset can be applied to multiple Gateways across multiple clouds to accommodate a dynamic multi-cloud policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the rule set`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the rule set`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Deprecated) This block can be repeated multiple times. Look below for the [definition/structure](#rule) of the rule ## Rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Rule name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Rule detailed description.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `. Set the rule's state to enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) ` + "`" + `ReverseProxy` + "`" + `, ` + "`" + `ForwardProxy` + "`" + `, ` + "`" + `Forwarding` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) address_id of the valtix_address_object. Defaults to "any".`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) address_id of the valtix_address_object. Defaults to "any".`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Service id of the valtix_service_object. The service object's service_type must match the rule type`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) ` + "`" + `ALLOW_LOG` + "`" + ` (log the event), ` + "`" + `ALLOW` + "`" + ` (do not log the event), ` + "`" + `DENY` + "`" + ` (log the event), ` + "`" + `DENY_NOLOG` + "`" + ` (do not log the event). Events are viewed in the Valtix UI (Investigate -> Flow Analytics).`,
				},
				resource.Attribute{
					Name:        "network_intrusion_profile",
					Description: `(Optional) profile_id of the valtix_profile_network_intrusion.`,
				},
				resource.Attribute{
					Name:        "dlp_profile",
					Description: `(Optional) profile_id of the valtix_profile_dlp.`,
				},
				resource.Attribute{
					Name:        "web_protection_profile",
					Description: `(Optional) profile_id of the valtix_profile_web_protection.`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_profile",
					Description: `(Optional) profile_id of the valtix_profile_fqdn_filter.`,
				},
				resource.Attribute{
					Name:        "anti_virus_profile",
					Description: `(Optional) profile_id of the valtix_profile_anti_virus.`,
				},
				resource.Attribute{
					Name:        "url_filter",
					Description: `(Optional) profile_id of the valtix_profile_url_filter.`,
				},
				resource.Attribute{
					Name:        "malicious_src_profile",
					Description: `(Optional) profile_id of the valtix_profile_malicious_src_profile.`,
				},
				resource.Attribute{
					Name:        "packet_capture_enabled",
					Description: `(Optional) true/false. Capture pcap when traffic matches the rule. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `ID of the Ruleset that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `ID of the Ruleset that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rules",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing the Policy Rules that are used to define the segmentation policy and advanced security protection of a Policy Ruleset.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `(Required) The ID of the Policy Rule Set`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A list of rules, this block can be repeated multiple times. Look below for the [definition/structure](#rule) of the rule ## Rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Rule name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Rule detailed description.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `. Set the rule's state to enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) address_id of the valtix_address_object. Defaults to "any".`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) address_id of the valtix_address_object. Defaults to "any".`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Service id of the valtix_service_object. The service object's service_type must match the rule type`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) ` + "`" + `ALLOW_LOG` + "`" + ` (log the event), ` + "`" + `ALLOW` + "`" + ` (do not log the event), ` + "`" + `DENY` + "`" + ` (log the event), ` + "`" + `DENY_NOLOG` + "`" + ` (do not log the event). Events are viewed in the Valtix UI (Investigate -> Flow Analytics).`,
				},
				resource.Attribute{
					Name:        "network_intrusion_profile",
					Description: `(Optional) profile_id of the valtix_profile_network_intrusion.`,
				},
				resource.Attribute{
					Name:        "dlp_profile",
					Description: `(Optional) profile_id of the valtix_profile_dlp.`,
				},
				resource.Attribute{
					Name:        "web_protection_profile",
					Description: `(Optional) profile_id of the valtix_profile_web_protection.`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_profile",
					Description: `(Optional) profile_id of the valtix_profile_fqdn_filter.`,
				},
				resource.Attribute{
					Name:        "anti_virus_profile",
					Description: `(Optional) profile_id of the valtix_profile_anti_virus.`,
				},
				resource.Attribute{
					Name:        "url_filter",
					Description: `(Optional) profile_id of the valtix_profile_url_filter.`,
				},
				resource.Attribute{
					Name:        "malicious_src_profile",
					Description: `(Optional) profile_id of the valtix_profile_malicious_src_profile.`,
				},
				resource.Attribute{
					Name:        "packet_capture_enabled",
					Description: `(Optional) true/false. Capture pcap when traffic matches the rule.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_anti_virus",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing an Anti Virus (AV) Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "auto_update",
					Description: `(Optional) Auto update the Talos ClamAV Ruleset version with a delay specified by ` + "`" + `delay_by_days` + "`" + ` argument. Applicable values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) Number of days to delay updating the Talos ClamAV Ruleset version after it has been published by Valtix. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately). New Rulesets as published as soon as updates are available from the Vendor and validation testing is completed by Valtix.`,
				},
				resource.Attribute{
					Name:        "talos_ruleset_version",
					Description: `(Optional) Talos ClamAV Ruleset version. Applicable values can be found from within the UI. The Rulesets are published frequently. Unless a specific version is desired, Valtix recommends using Auto Update as described above. If this argument is specified, Auto Update of Talos ClamAV Ruleset is disabled and the Profile will use the specified Talos ClamAV Ruleset version.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Antivirus Network Threat is detected. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). If not specified, then the action assumed is ` + "`" + `Allow Log` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_application_threat",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Web Application Firewall (WAF) Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "auto_update_crs",
					Description: `(Optional) Auto Update of the CRS Ruleset version with a delay specified by ` + "`" + `delay_by_days_crs` + "`" + ` argument. Applicable values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days_crs",
					Description: `(Optional) Number of days to delay updating the CRS Ruleset version after it has been published by Valtix. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately). New Rulesets as published as soon as updates are available from the Vendor and validation testing is completed by Valtix.`,
				},
				resource.Attribute{
					Name:        "crs_ruleset_version",
					Description: `(Optional) CRS Ruleset version. Applicable values can be found from within the UI. The Rulesets are published frequently. Unless a specific version is desired, Valtix recommends using Auto Update as described above. If this argument is specified, Auto Update of CRS Ruleset is disabled and the Profile will use the specified CRS Ruleset version.`,
				},
				resource.Attribute{
					Name:        "auto_update_trustwave",
					Description: `(Optional) Auto update the Trustwave Ruleset version with a delay specified by ` + "`" + `delay_by_days_trustwave` + "`" + ` argument. Applicable values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days_trustwave",
					Description: `(Optional) Number of days to delay updating the Trustwave Ruleset version after it has been published by Valtix. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately). New Rulesets as published as soon as updates are available from the Vendor and validation testing is completed by Valtix.`,
				},
				resource.Attribute{
					Name:        "trustwave_ruleset_version",
					Description: `(Optional) Trustwave Ruleset version. Applicable values can be found from within the UI. The Rulesets are published frequently. Unless a specific version is desired, Valtix recommends using Auto Update as described above. If this argument is specified, Auto Update of Trustwave Ruleset is disabled and the Profile will use the specified Trustwave Ruleset version.`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(Required) An integer between ` + "`" + `1` + "`" + ` and ` + "`" + `4` + "`" + `. Higher number leads to more false positives, but also helps in detecting more attacks. Recommended value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_anamoly",
					Description: `(Optional) Request anomaly score used in the Mod Security anomaly scoring system. If not specified, the efault value is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "response_anamoly",
					Description: `(Optional) Response anomaly score used in the Mod Security anomaly scoring system. If not specified, the default value is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pcap",
					Description: `(Optional) Captures a PCAP when attacks are detected. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "api_logging",
					Description: `(Optional) Logs API requests and generates a HAR file. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "request_inspection_profile",
					Description: `(Optional) List of Request attacks detected by the CRS Ruleset`,
				},
				resource.Attribute{
					Name:        "response_inspection_profile",
					Description: `(Optional) List of Response attacks detected by the CRS Ruleset`,
				},
				resource.Attribute{
					Name:        "advanced_application_profile",
					Description: `(Optional) Advanced attacks detected by the Trustwave Ruleset`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Web Application (WAF) Application Threat is detected. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event). If not specified, then the action assumed is ` + "`" + `Deny Log` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rule_event_filter",
					Description: `(Optional) Rate Limit / Sample a set of rules. Structure is [defined below](#rule-event-filter)`,
				},
				resource.Attribute{
					Name:        "event_suppressor",
					Description: `(Optional) Suppress a given set of rule ids for traffic from certain sources. Structure is [defined below](#event-suppressor)`,
				},
				resource.Attribute{
					Name:        "profile_event_filter",
					Description: `(Optional) Similar to the rule_event_filter but applies to the whole profile instead of specific rule(s). Structure is [defined below](#profile-event-filter) ## Rule Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of Rule IDs to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) ` + "`" + `RATE` + "`" + ` or ` + "`" + `SAMPLE` + "`" + `. If ` + "`" + `RATE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` and ` + "`" + `time` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events within specified time period. If ` + "`" + `SAMPLE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events.`,
				},
				resource.Attribute{
					Name:        "number_of_events",
					Description: `(Optional) Number of times the attack must match a Rule ID before the action is applied`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Used when the ` + "`" + `type` + "`" + ` is set to ` + "`" + `RATE` + "`" + ` where the number of times the attack must match a Rule ID within a specified time period (in seconds) before the action is applied. ## Rule Suppressor`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) List of source IPs or CIDRs`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of Rule IDs to filter ## Profile Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of Rule IDs to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) ` + "`" + `RATE` + "`" + ` or ` + "`" + `SAMPLE` + "`" + `. If ` + "`" + `RATE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` and ` + "`" + `time` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events within specified time period. If ` + "`" + `SAMPLE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events.`,
				},
				resource.Attribute{
					Name:        "number_of_events",
					Description: `(Optional) Number of times the attack must match a Rule ID before the action is applied`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Used when the ` + "`" + `type` + "`" + ` is set to ` + "`" + `RATE` + "`" + ` where the number of times the attack must match a Rule ID within a specified time period (in seconds) before the action is applied. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_decryption",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Decryption Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the decryption profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the profile`,
				},
				resource.Attribute{
					Name:        "certificate_name",
					Description: `(Required) Certificate name from the valtix_certificate object ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_diagnostic",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Diagnostic Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Diagnostic Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Diagnostic Profile`,
				},
				resource.Attribute{
					Name:        "csp_account",
					Description: `(Required) Cloud account name added to the Controller (valtix_cloud_account.name)`,
				},
				resource.Attribute{
					Name:        "storage_upload_path",
					Description: `(Required - AWS) Storage bucket name`,
				},
				resource.Attribute{
					Name:        "azure_storage_accnt_name",
					Description: `(Required - Azure) Storage account name`,
				},
				resource.Attribute{
					Name:        "azure_blob_container_name",
					Description: `(Required - Azure) Storage container name`,
				},
				resource.Attribute{
					Name:        "azure_storage_access_key",
					Description: `(Optional - Azure) Storage account access key, if required ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_dlp",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Data Loss Prevention (DLP) Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "dlp_filter_list",
					Description: `(Required) A list of DLP Filters. Structure [defined below](#dlp-filter). ## DLP Filter`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description of the Filter`,
				},
				resource.Attribute{
					Name:        "patterns",
					Description: `(Optional) List of custom Perl Compatible Regular Expression (PCRE) patterns`,
				},
				resource.Attribute{
					Name:        "static_patterns",
					Description: `(Optional) List of pre-defined patterns`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) Number of times the pattern must be seen before a match is determined`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take when a Data Loss Prevention (DLP) Network Threat is detected. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_fqdn",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing an FQDN Filtering Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "no_fqdn_deny",
					Description: `(Optional) Deny traffic when no FQDN found in packet. Applicable values:`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_list",
					Description: `(Required) One or more ` + "`" + `fqdn_list` + "`" + ` resources, where each resource is a row in the FQDN Filter List (maximum of 32 resources). Structure [defined below](#fqdn-filter-list).`,
				},
				resource.Attribute{
					Name:        "uncategorized_fqdn_filter",
					Description: `(Required) Uncategorized FQDN Filter action for any FQDN that does not match the FQDNs defined in the ` + "`" + `fqdn_filter_list` + "`" + ` resource and is not represented by any vendor category (whether specified or not). Structure [defined below](#uncategorized-fqdn-filter).`,
				},
				resource.Attribute{
					Name:        "default_fqdn_filter",
					Description: `(Required) Default FQDN Filter action for any FQDN that does not match the FQDNs defined in the ` + "`" + `fqdn_filter_list` + "`" + ` resource or is not matched by the ` + "`" + `uncategorized_fqdn_filter` + "`" + ` resource (if specified). This should be the last resource specified in the list of resources. Structure [defined below](#default-fqdn-filter). ## FQDN Filter List`,
				},
				resource.Attribute{
					Name:        "fqdn_list",
					Description: `(Required) List of strings (maximum of 64 strings): Applicable values: FQDNs or Perl Compatible Regular Expression (PCRE) patterns. Structure [defined below](#fqdn-list).`,
				},
				resource.Attribute{
					Name:        "vendor_category_list",
					Description: `(Optional) List of pre-defined Vendor Categories (maximum 128 categories). Structure [defined below](#vendor-category-list).`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Action to take when an FQDN matches an entry in the ` + "`" + `fqnd_list` + "`" + ` or ` + "`" + `vendor_category_list` + "`" + `. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event).`,
				},
				resource.Attribute{
					Name:        "decryption_exception",
					Description: `(Optional) When used in conjunction with a proxy Rule (ForwardProxy, ReverseProxy), instructs the proxy engine to bypass decryption. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `. ## FQDN List ` + "`" + `` + "`" + `` + "`" + ` fqdn_list = ["www.website1.com", ".`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_l7dos",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the IPS profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the IPS profile`,
				},
				resource.Attribute{
					Name:        "request_limits",
					Description: `(Required) List of URI and limits. Structure of limit is [defined below](#request-limit) ## Request Limit`,
				},
				resource.Attribute{
					Name:        "target_uri",
					Description: `(Required) URI to which the limits are applied. E.g '/api/resource1' or '/api/login'`,
				},
				resource.Attribute{
					Name:        "requests_rate_block",
					Description: `(Required) Number of requests/sec processed by Valtix. Depending on the number of instances running in the Gateway cluster, the backend application would get more traffic than the number specified here. This limit is per client_ip/src_ip`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `(Required) The queue size for client_ip/src_ip. Once the requests fill the burst_size and are still not processed, new requests from the same client would be dropped`,
				},
				resource.Attribute{
					Name:        "rate_limited_methods",
					Description: `(Required) HTTP actions/methods to which the rate limiting is applied. Valid values are ` + "`" + `METHOD_GET` + "`" + `, ` + "`" + `METHOD_POST` + "`" + `, ` + "`" + `METHOD_PUT` + "`" + `, ` + "`" + `METHOD_DELETE` + "`" + `, ` + "`" + `METHOD_PATCH` + "`" + `, ` + "`" + `METHOD_HEAD` + "`" + `, ` + "`" + `METHOD_OPTIONS` + "`" + ` ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_log_forwarding",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the decryption profile`,
				},
				resource.Attribute{
					Name:        "siem_vendor",
					Description: `(Required) One of ` + "`" + `SPLUNK` + "`" + `, ` + "`" + `DATADOG` + "`" + `, ` + "`" + `GCPLOGGING_FROM_GATEWAY` + "`" + `, or ` + "`" + `REMOTE_SYSLOG` + "`" + ` ## Additional arguments based on ` + "`" + `siem_vendor` + "`" + ` ### SPLUNK`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required ) HTTPS URL`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required ) HTTPS auth token`,
				},
				resource.Attribute{
					Name:        "splunk_index",
					Description: `(Required) Splunk index name to store the events ### DATADOG`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required ) HTTPS URL`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) HTTPS auth token ### GCPLOGGING_FROM_GATEWAY`,
				},
				resource.Attribute{
					Name:        "log_name",
					Description: `[Optional] GCP log name to store the logs. If not specified, the default name is "valtix-gateway-logs" ### SYSLOG`,
				},
				resource.Attribute{
					Name:        "syslog_server_ip",
					Description: `(Required) Syslog Server IP`,
				},
				resource.Attribute{
					Name:        "syslog_port",
					Description: `(Required) Syslog Server port`,
				},
				resource.Attribute{
					Name:        "network_threat_severity",
					Description: `(Required) Syslog Server port`,
				},
				resource.Attribute{
					Name:        "syslog_flow_logs",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Forward flow logs to Syslog.`,
				},
				resource.Attribute{
					Name:        "syslog_firewall_events",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Forward firewall events to Syslog.`,
				},
				resource.Attribute{
					Name:        "syslog_https_logs",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Forward TLS logs to Syslog.`,
				},
				resource.Attribute{
					Name:        "network_threat_severity",
					Description: `(Optional) One of ` + "`" + `Alert` + "`" + `, ` + "`" + `Emergency` + "`" + `, ` + "`" + `Critical` + "`" + `, ` + "`" + `Error` + "`" + `, ` + "`" + `Warning` + "`" + `, ` + "`" + `Notice` + "`" + `, ` + "`" + `Info` + "`" + `, or ` + "`" + `Debug` + "`" + `. Forward Network Threat events (IPS) with the given severity only.`,
				},
				resource.Attribute{
					Name:        "web_attack_severity",
					Description: `(Optional) One of ` + "`" + `Alert` + "`" + `, ` + "`" + `Emergency` + "`" + `, ` + "`" + `Critical` + "`" + `, ` + "`" + `Error` + "`" + `, ` + "`" + `Warning` + "`" + `, ` + "`" + `Notice` + "`" + `, ` + "`" + `Info` + "`" + `, ` + "`" + `Debug` + "`" + `. Forward Web Attacks with the given severity only. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_malicious_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Malicious IP Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "ip_reputation_enabled",
					Description: `(Optional) Defines whether the Malicious IP Profile is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` (enabled) or ` + "`" + `false` + "`" + ` (disabled). If not specified, the default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_update",
					Description: `(Optional) Auto update the Trustwave Malicious IP Ruleset version. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) Number of days to delay updating the Trustwave Malicious IP Ruleset version after it has been published by Valtix. Valid values are integers from 0 to 30. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (0 days). Valtix publishes new Rulesets as soon as updates are available from the Vendor and complete testing by Valtix.`,
				},
				resource.Attribute{
					Name:        "vendor_rule_set_name",
					Description: `(Optional) Trustwave IP Reputation Ruleset version. Applicable values can be found from within the UI. The Rulesets are published everyday. Unless you want to use a specific version, Valtix recommends to use auto_update as described above. If this argument is specified, Auto Update of Trustwave IP Reputation Ruleset is disabled and the Profile will only use the specified version for Trustwave IP Reputation Ruleset.`,
				},
				resource.Attribute{
					Name:        "match_xff",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Defines whether the IP information available in the X-Forwarded-For HTTP request header should be evaluated. When not specified, the default value is ` + "`" + `false` + "`" + `. (Even though this argument is optional, it is recommended to specify a value explicitly, as the default value may change in the future). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_network_intrusion",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Network Intrusion (IPS) Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "auto_update",
					Description: `(Optional) Auto Update the Talos Network Intrusion Ruleset version. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if unspecified) is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) Number of days to delay updating the Talos Network Intrusion Ruleset version after it has been published by Valtix. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately). New Rulesets as published as soon as updates are available from the Vendor and validation testing is completed by Valtix.`,
				},
				resource.Attribute{
					Name:        "talos_ruleset_version",
					Description: `(Optional) Talos Network Intrusion Ruleset version. Applicable values can be found from within the UI. The Rulesets are published frequently. Unless a specific version is desired, Valtix recommends using Auto Update as described above. If this argument is specified, Auto Update of Talos Network Intrusion Ruleset is disabled and the Profile will use the specified Talos Network Intrusion Ruleset version.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Network Intrusion (IDS/IPS) Network Threat is detected for Rules that fall into the Profile set (the entire set or Rules defined by the configuration of the Profile). Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). If not specified, then the action assumed is ` + "`" + `Rule Default` + "`" + `, the action defined in the Policy Ruleset Rule. This action can be overridden for each Policy, [Category](#categories) and [Class](#classes) of Rules.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Pre-defined Talos Network Intrusion Ruleset to use. Applicable values: ` + "`" + `CONNECTIVITY` + "`" + `, ` + "`" + `BALANCED` + "`" + `, ` + "`" + `SECURITY` + "`" + `, ` + "`" + `MAX_DETECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_action",
					Description: `(Optional) Action to take when a Network Intrusion (IDS/IPS) Network Threat is detected for Rules that fall into the Policy set. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). This action is an override to the Profile action. If not specified, then the action assumed is the action defined for the Profile.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Optional) Pre-defined Categories. Structure [defined below](#categories). This block can be repeated multiple times.`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `(Optional) Predefined Classes. Structure [defined below](#classes). This block can be repeated multiple times.`,
				},
				resource.Attribute{
					Name:        "pcap",
					Description: `(Optional) Capture packets (PCAP) when traffic matches the Network Intrusion rules. Valid values are`,
				},
				resource.Attribute{
					Name:        "rule_event_filter",
					Description: `(Optional) Rate Limit / Sample a set of rules. Structure [defined below](#rule-event-filter). This block can be repeated multiple times.`,
				},
				resource.Attribute{
					Name:        "event_suppressor",
					Description: `(Optional) Suppress a given set of Rule IDs for traffic from certain sources. Structure [defined below](#event-suppressor). This block can be repeated multiple times.`,
				},
				resource.Attribute{
					Name:        "profile_event_filter",
					Description: `(Optional) Similar to the`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of a Talos Network Intrusion attack Category`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Network Intrusion (IDS/IPS) Network Threat is detected for Rules that fall into the Category set. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). This action is an override to the Policy action. If not specified, then the action assumed is the action defined for the Policy. ## Classes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Talos Network Intrusion attack Class`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Network Intrusion (IDS/IPS) Network Threat is detected for Rules that fall into the Class set. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). This action is an override to the Categories action. If not specified, then the action assumed is the action defined for the Category. ## Rule Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of Rule IDs to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) ` + "`" + `RATE` + "`" + ` or ` + "`" + `SAMPLE` + "`" + `. If ` + "`" + `RATE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` and ` + "`" + `time` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events within specified time period. If ` + "`" + `SAMPLE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events.`,
				},
				resource.Attribute{
					Name:        "number_of_events",
					Description: `(Optional) Number of times the attack must match a Rule ID before the action is applied`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Used when the ` + "`" + `type` + "`" + ` is set to ` + "`" + `RATE` + "`" + ` where the number of times the attack must match a Rule ID within a specified time period (in seconds) before the action is applied. ## Rule Suppressor`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) List of source IPs or CIDRs`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of Rule IDs to filter ## Profile Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of Rule IDs to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) ` + "`" + `RATE` + "`" + ` or ` + "`" + `SAMPLE` + "`" + `. If ` + "`" + `RATE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` and ` + "`" + `time` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events within specified time period. If ` + "`" + `SAMPLE` + "`" + ` is selected, the ` + "`" + `number_of_events` + "`" + ` must be specified. The action is taken once the provided Rule IDs match the number of events.`,
				},
				resource.Attribute{
					Name:        "number_of_events",
					Description: `(Optional) Number of times the attack must match a Rule ID before the action is applied`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Used when the ` + "`" + `type` + "`" + ` is set to ` + "`" + `RATE` + "`" + ` where the number of times the attack must match a Rule ID within a specified time period (in seconds) before the action is applied. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_packet_capture",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Packet Capture Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Packet Capture Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Packet Capture Profile`,
				},
				resource.Attribute{
					Name:        "csp_account",
					Description: `(Required) Cloud account name added to the Controller (valtix_cloud_account.name)`,
				},
				resource.Attribute{
					Name:        "storage_upload_path",
					Description: `(Required - AWS and GCP) Storage bucket name`,
				},
				resource.Attribute{
					Name:        "azure_storage_accnt_name",
					Description: `(Required - Azure) Storage account name`,
				},
				resource.Attribute{
					Name:        "azure_blob_container_name",
					Description: `(Required - Azure) Storage container name`,
				},
				resource.Attribute{
					Name:        "azure_storage_access_key",
					Description: `(Optional - Azure) Storage account access key, if required ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_urlfilter",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a URL Filtering Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Profile`,
				},
				resource.Attribute{
					Name:        "url_filter_list",
					Description: `(Required) One or more ` + "`" + `url_list` + "`" + ` resources, where each resource is a row in the URL Filter List (maximum of 64 resources). Structure [defined below](#url-filter-list).`,
				},
				resource.Attribute{
					Name:        "uncategorized_url_filter",
					Description: `(Required) Uncategorized URL Filter action for any URL that does not match the URLs defined in the ` + "`" + `url_filter_list` + "`" + ` resource and is not represented by any vendor category (whether specified or not). Structure [defined below](#uncategorized-url-filter).`,
				},
				resource.Attribute{
					Name:        "default_url_filter",
					Description: `(Required) Default URL Filter action for any URL that does not match the URLs defined in the ` + "`" + `url_filter_list` + "`" + ` resource or is not matched by the ` + "`" + `uncategorized_url_filter` + "`" + ` resource (if specified). This should be the last resource specified in the list of resources. Structure [defined below](#default-url-filter). ## URL Filter List`,
				},
				resource.Attribute{
					Name:        "url_list",
					Description: `(Required) List of strings (maximum of 64 strings). Applicable values: URLs or Perl Compatible Regular Expression (PCRE) patterns. Structure [defined below](#url-list).`,
				},
				resource.Attribute{
					Name:        "vendor_category_list",
					Description: `(Optional) List of pre-defined Vendor Categories (maximum of 128 categories). Structure [defined below](#vendor-category-list).`,
				},
				resource.Attribute{
					Name:        "filter_methods",
					Description: `(Optional) List of URL methods (` + "`" + `ALL` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `GET` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `PATCH` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `PUT` + "`" + `). If not specified, the default value is ` + "`" + `ALL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Action to take when a URL matches an entry in the ` + "`" + `url_list` + "`" + ` or ` + "`" + `vendor_category_list` + "`" + `. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event).`,
				},
				resource.Attribute{
					Name:        "return_status",
					Description: `(Optional) HTTP status code to return when URLs are denied. This argument only applies to resources that have a ` + "`" + `policy` + "`" + ` set to ` + "`" + `Deny Log` + "`" + ` or ` + "`" + `Deny No Log` + "`" + `. ## URL List ` + "`" + `` + "`" + `` + "`" + ` url_list = ["www.website1.com", ".`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "profile_id",
					Description: `ID of the Profile that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_object",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Service Object.  A Service Object is used in a Policy Ruleset Rule to define how the traffic will be processed by the Rule.

~> **Note on valtix_service_object ReverseProxy resource**
If multiple applications are accessed using the same port (e.g., port 443), then an SNI must be specified to differentiate the requests associated with each application.

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `ID of the Service Object that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_vpc",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Valtix Service VPC/VNet.

A Valtix Service VPC/VNet resource is used in AWS, Azure and GCP deployments to create a Service VPC/VNet as the destination for a Valtix Gateway deployment. For a Valtix Gateway deployed in AWS, Azure or GCP using HUB mode, a Service VPC created and managed by Valtix must be deployed as a pre-requisite. Gateway instances will be deployed in all Availability Zones associated with the Service VPC.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Service VPC/VNet`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The CSP Account name (configured in Valtix) where the Service VPC/VNet will be deployed`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The Region/Location where the Service VPC/VNet will be deployed`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) CIDR of the Service VPC/VNet to be deployed. For GCP only: This CIDR is used for the Datapath Subnet created within the Datapath VPC as part of the GCP Service VPC construct. When the Gateways are deployed within the GCP Service VPC construct (Management and Datapath VPCs), the datapath interfaces will be deployed in the Datapath VPC / Subnet. This CIDR block must be different than the CIDR specified in the ` + "`" + `management_cidr` + "`" + ` argument.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Required) List of Availability Zones for the Region/Location to associate with the Service VPC/VNet. Valtix Gateways deployed in this Service VPC/VNet will have instances deployed in all associated Availability Zones.`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `(Required for AWS) Transit Gateway ID for the Service VPC to attach to`,
				},
				resource.Attribute{
					Name:        "azure_resource_group",
					Description: `(Required for Azure) Resource Group Name in which the Service VNet and its resources are created`,
				},
				resource.Attribute{
					Name:        "management_cidr",
					Description: `(Required for GCP) The CIDR used for the Management Subnet created within the Management VPC as part of the GCP Service VPC construct. When the Gateways are deployed within the GCP Service VPC construct (Management and Datapath VPCs), the management interfaces will be deployed in the Management VPC / Subnet. This CIDR block must be different than the CIDR specified in the ` + "`" + `cidr` + "`" + ` argument.`,
				},
				resource.Attribute{
					Name:        "use_nat_gateway",
					Description: `(Optional for AWS) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If ` + "`" + `true` + "`" + `, enables egress communication through NAT gateway in each AZ (Default ` + "`" + `false` + "`" + `). ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Terraform resource ID of the Services VPC/VNet`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Cloud specific ID of the Services VPC/VNet`,
				},
				resource.Attribute{
					Name:        "service_vpc_id",
					Description: `Same as ` + "`" + `id` + "`" + ` (for backward compatibility)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Terraform resource ID of the Services VPC/VNet`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Cloud specific ID of the Services VPC/VNet`,
				},
				resource.Attribute{
					Name:        "service_vpc_id",
					Description: `Same as ` + "`" + `id` + "`" + ` (for backward compatibility)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_spoke_vpc",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing Spoke VPC protection.

In AWS, the Spoke VPC is attached to the Transit Gateway and appropriate routing is setup in Transit Gateway route tables, so the traffic reaches the Valtix Service VPC. (Customer has to set the default route in the Spoke VPC routing tables to point to the Transit Gateway).

In Azure, VNet peering is setup between the Service VNet and the Spoke VNet.

In GCP, VPC peering is setup between the Service VPC and the Spoke VPC.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_vpc_id",
					Description: `(Required) Service VPC/VNet ID`,
				},
				resource.Attribute{
					Name:        "spoke_vpc_id",
					Description: `(Required) Spoke VPC/VNet ID`,
				},
				resource.Attribute{
					Name:        "spoke_vpc_csp_account_name",
					Description: `(Optional - AWS, GCP) Valtix CSP Account Name where the Spoke VPC resides (if different from the Valtix CSP Account Name in which the Service VPC resides)`,
				},
				resource.Attribute{
					Name:        "spoke_vpc_region",
					Description: `(Optional - AWS) Region where the spoke vpc exists (if different from the CSP account name in which Service VPC is created)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"valtix_address_object":             0,
		"valtix_alert_profile":              1,
		"valtix_alert_rule":                 2,
		"valtix_certificate":                3,
		"valtix_cloud_account":              4,
		"valtix_external_id":                5,
		"valtix_gateway":                    6,
		"valtix_policy_rule_set":            7,
		"valtix_policy_rules":               8,
		"valtix_profile_anti_virus":         9,
		"valtix_profile_application_threat": 10,
		"valtix_profile_decryption":         11,
		"valtix_profile_diagnostic":         12,
		"valtix_profile_dlp":                13,
		"valtix_profile_fqdn":               14,
		"valtix_profile_l7dos":              15,
		"valtix_profile_log_forwarding":     16,
		"valtix_profile_malicious_ip":       17,
		"valtix_profile_network_intrusion":  18,
		"valtix_profile_packet_capture":     19,
		"valtix_profile_urlfilter":          20,
		"valtix_service_object":             21,
		"valtix_service_vpc":                22,
		"valtix_spoke_vpc":                  23,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
