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
					Description: `(Required) A list of IPs, CIDRs or FQDNs. The total number of FQDNs is limited to 200 where each FQDN can resolve to at most 400 IPs. <br><br>For an example, see [STATIC (Source Destination) Example](#static-source-destination-examples) ~>`,
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
					Description: `(Required) This argument must be set to ` + "`" + `true` + "`" + ` <br><br>For an example, see [STATIC (Reverse Proxy Target) Example](#static-reverse-proxy-target-examples) #### DYNAMIC_APPLICATIONS (Reverse Proxy Target) Arguments`,
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
					Description: `(Required) A list of ` + "`" + `address_id` + "`" + ` to be grouped together <br><br>For an example, see [GROUP (Source Destination) Example](#group-source-destination-example) #### DYNAMIC_ASG (Source Destination) Arguments`,
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
					Description: `(Required) Name of the Azure CSP account onboarded into Valtix`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) ID of Azure Region where application security group is located`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Required) Name of Azure Resource Group where application security group is located`,
				},
				resource.Attribute{
					Name:        "application_security_group_id",
					Description: `(Required) ID of Azure Application Security Group (ASG), for example : /subscriptions/<SUBSCRIPTION_ID>/resourceGroups/<RESOURCE_GROUP>/providers/Microsoft.Network/applicationSecurityGroups/<APPLICATION_SECURITY_GROUP> <br><br>For an example, see [DYNAMIC_ASG (Source Destination) Example](#dynamic_asg-source-destination-example) ### Tag List A ` + "`" + `tag_list` + "`" + ` block representing a Tag key-value pair requires the following arguments:`,
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
					Name:        "id",
					Description: `ID of the Address Object resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "address_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Address Object resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_address_object.ip_cidr_ag 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Address Object resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "address_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Address Object resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_address_object.ip_cidr_ag 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Name of the Alert Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the Alert Profile`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) One of ` + "`" + `PagerDutyEventApi` + "`" + `, ` + "`" + `ServiceNowWebHook` + "`" + `, ` + "`" + `SlackWebHook` + "`" + `, ` + "`" + `DataDog` + "`" + `, ` + "`" + `MicrosoftSentinel` + "`" + ` ### Destination-specific Arguments ### Pagerduty`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) Shared key / primary key used to authenticate with Pagerduty ### ServiceNow`,
				},
				resource.Attribute{
					Name:        "integration_url",
					Description: `(Required) HTTPS endpoint URL`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) HTTPS auth token ### Slack`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) Shared key / primary key used to authenticate with Slack ### Datadog`,
				},
				resource.Attribute{
					Name:        "integration_url",
					Description: `(Required) HTTPS endpoint URL`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) HTTPS auth token ### Microsoft Sentinel`,
				},
				resource.Attribute{
					Name:        "log_analytics_log_type",
					Description: `(Required) Name of the MS Sentinel table used to store the alerts`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace_id",
					Description: `(Required) ID of the MS Sentinel workspace`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) Shared key / primary key used to authenticate with MS Sentinel ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Alert Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Alert Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_alert_profile.slack 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Alert Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Alert Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_alert_profile.slack 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Description of the Alert Profile`,
				},
				resource.Attribute{
					Name:        "alert_profile",
					Description: `(Required) Alert Profile ID`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Valid values are ` + "`" + `Type_Inventory` + "`" + `, ` + "`" + `Type_SystemLogs` + "`" + `, or ` + "`" + `Type_AuditLogs` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sub_type",
					Description: `(Required - SystemLogs) Applicable only to ` + "`" + `type` + "`" + ` set to ` + "`" + `Type_Inventory` + "`" + ` or ` + "`" + `Type_SystemLogs` + "`" + `. For ` + "`" + `type` + "`" + ` set to ` + "`" + `Type_Inventory` + "`" + `, valid value for ` + "`" + `sub_type` + "`" + ` is ` + "`" + `SubType_InventoryGuardRails` + "`" + `. For ` + "`" + `type` + "`" + ` set to ` + "`" + `Type_SystemLogs` + "`" + `, valid values for ` + "`" + `sub_type` + "`" + ` are ` + "`" + `SubType_SystemLogsGateway` + "`" + ` or ` + "`" + `SubType_SystemLogsAccount` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) Applicable values when ` + "`" + `type` + "`" + ` is set to ` + "`" + `Type_Inventory` + "`" + ` are ` + "`" + `Info` + "`" + `, ` + "`" + `Medium` + "`" + `, or ` + "`" + `High` + "`" + `. Applicable values when ` + "`" + `type` + "`" + ` is set to ` + "`" + `Type_SystemLogs` + "`" + ` are ` + "`" + `Info` + "`" + `, ` + "`" + `Warning` + "`" + `, ` + "`" + `Medium` + "`" + `, ` + "`" + `High` + "`" + `, or ` + "`" + `Critical` + "`" + `. Alerts will be sent for all alerts that have severity equal to and higher than the specified value.`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Required) Applicable values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` ## Import Alert Rule resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_alert_rule.alert_rule1 alert-rule1 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) GCP Secrets Manager key name that contains the private key. The value must be the full path of the key name (e.g., ` + "`" + `projects/012345678910/secrets/gcp-certificate-secret` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "secret_version",
					Description: `(Optional) GCP Secrets Manager key version ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Certificate resource that can be referenced in other resources (e.g. valtix_profile_decryption) ## Import Certificate resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_certificate.import_certificate import_certificate ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Certificate resource that can be referenced in other resources (e.g. valtix_profile_decryption) ## Import Certificate resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_certificate.import_certificate import_certificate ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_cloud_account",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing Valtix Cloud Accounts consisting of on-boarded AWS Accounts, Azure Subscriptions, GCP Projects and OCI Tenants.  The Valtix Cloud Account defines the credentials used for the Valtix Controller to discover inventory assets and traffic, and orchestrate and manage the deployment of Valtix Gateways.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud Account on the Valtix Console. Must contain only alphanumeric, hyphens or underscore characters and not exceed 100 characters`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Cloud Account`,
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
					Description: `(GCP - Required) Service Account credentials key file created for the Valtix Controller access.`,
				},
				resource.Attribute{
					Name:        "inventory_monitoring",
					Description: `Enable inventory monitoring (can be repeated multiple times). See [Inventory Monitoring](#inventory-monitoring) for details.`,
				},
				resource.Attribute{
					Name:        "gcp_project_id",
					Description: `(Optional) Add the specified Project ID (name) to the Valtix Controller. If this is not specified, then the Project ID from the credentials file is used. The argument is required if a Service Account is used to manage multiple GCP Projects. ## Inventory Monitoring`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions to enable and monitor inventory`,
				},
				resource.Attribute{
					Name:        "refresh_interval",
					Description: `Interval in minutes where the inventory is refreshed ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Cloud Account resource that can be referenced in other resources ## Import Cloud Account resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_cloud_account.aws1 aws-account1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Cloud Account resource that can be referenced in other resources ## Import Cloud Account resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_cloud_account.aws1 aws-account1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_cloud_account_log_profile_association",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The name of the CSP account onboarded into Valtix`,
				},
				resource.Attribute{
					Name:        "log_profile_id",
					Description: `(Required) The ID of the Log Forwarding Profile`,
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
					Description: `External ID that must be used in the trust settings of the cross Account IAM Role`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_id",
					Description: `External ID that must be used in the trust settings of the cross Account IAM Role`,
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

~> **Note on Valtix Gateway Policy**
The [` + "`" + `valtix_cloud_account` + "`" + `](../valtix_cloud_account) and [` + "`" + `valtix_policy_rule_set` + "`" + `](../valtix_policy_rule_set)
resources must be created before the ` + "`" + `valtix_gateway` + "`" + ` resource can be created

~> **Note on Shared VPC Resources**
The Valtix Gateway resource (` + "`" + `valtix_gateway` + "`" + `) can be deployed using Subnets of a VPC that is shared from another Project.  The Gateway instances themselves are deployed in the local Project, while the interfaces use the Subnets of the Shared VPC.  When the VPC subnets are shared, the check to confirm the shared VPC is valid, and the subnets are contained within, is not possible.  Valtix has removed the check altogether.  It is important that the shared VPC ID specified is correct and valid to ensure proper Gateway deployment.<br><br>
It is also required that the Project that is sharing the VPC needs to grant *Compute Network User* access to the Valtix Controller service principal associated with Project the VPC is shared with.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Gateway`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Gateway`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The CSP account where the Gateway will be deployed`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) The instance type used when deploying the Gateway. Applicable CSP-specific values are:`,
				},
				resource.Attribute{
					Name:        "gateway_image",
					Description: `(Required) Represents the image version to be used for this Gateway. A list of applicable image versions is available from the Valtix Portal (ADMINISTRATION -> Management -> System -> Gateway Images). Please view the [Valtix Release Recommendation](https://docs.valtix.com/releases/recommendation/) for the recommended Gateway image version or contact [Valtix Support](mailto:support@valtix.com).`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required - AWS, Azure, GCP) Applicable values are ` + "`" + `EDGE` + "`" + ` or ` + "`" + `HUB` + "`" + `. Review the Valtix product documentation for information on the different deployment modes. This argument is not supported for OCI and must not be used.`,
				},
				resource.Attribute{
					Name:        "security_type",
					Description: `(Optional) Applicable values are ` + "`" + `INGRESS` + "`" + ` or ` + "`" + `EGRESS` + "`" + `. If not specified, the default is ` + "`" + `INGRESS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "gateway_state",
					Description: `(Optional) Specifies the state of the Gateway. Applicable values are ` + "`" + `ACTIVE` + "`" + ` and ` + "`" + `INACTIVE` + "`" + `. When set to ` + "`" + `ACTIVE` + "`" + `, the Gateway will be enabled and operational. When set to ` + "`" + `INACTIVE` + "`" + `, the Gateway will be disabled and not operational. If not specified, the default is ` + "`" + `ACTIVE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait_for_gateway_state",
					Description: `(Optional) Determines if Terraform should wait for the Gateway state, defined by the ` + "`" + `gateway_state` + "`" + ` argument, to be achieved before completing. Applicable values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_rule_set_id",
					Description: `(Required) Rule set ID of valtix_policy_rule_set.`,
				},
				resource.Attribute{
					Name:        "ssh_key_pair",
					Description: `(Required - AWS, Azure) Name of the SSH Key Pair created within the AWS Account or Azure Subscription. The CSP Key Pairs are Regional constructs and must be created in the same Region as specified by the ` + "`" + `region` + "`" + ` argument where the Gateway will be deployed.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Required - Azure, GCP) The SSH public key to be assigned to the Gateway instances. Must be`,
				},
				resource.Attribute{
					Name:        "azure_user_name",
					Description: `(Optional - Azure) Name to use as the user when SSH to a Azure Gateway instance. When not specified, ` + "`" + `centos` + "`" + ` is used.`,
				},
				resource.Attribute{
					Name:        "gcp_user_name",
					Description: `(Optional - GCP) Name to use as the user when SSH to a GCP Gateway instance. When not specified, ` + "`" + `centos` + "`" + ` is used.`,
				},
				resource.Attribute{
					Name:        "gcp_service_account_email",
					Description: `(Required - GCP) The GCP Service Account Email that defines the permissions for the Gateway to integrate with other GCP Project resources such as Secrets Manager and Storage Buckets.`,
				},
				resource.Attribute{
					Name:        "aws_iam_role_firewall",
					Description: `(Required - AWS) The AWS IAM role that defines the permissions for the Gateway to integrate with other AWS Account resources such as Key Pairs, Secrets Manager and Key Management Service (KMS).`,
				},
				resource.Attribute{
					Name:        "azure_user_identity_id",
					Description: `(Optional - Azure) The Azure User Assigned Identity that defines the permissions for the Gateway to integrate with other Azure Subscription resources such as Key Pairs, Key Vault and Blob Storage.`,
				},
				resource.Attribute{
					Name:        "azure_resource_group",
					Description: `(Required - Azure) Azure Resource Group name used to associate all created Gateway resources`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the Gateway will be deployed`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC/VNet ID where the Gateway will be deployed. For HUB mode deployments, the value must refer to the`,
				},
				resource.Attribute{
					Name:        "aws_gateway_lb",
					Description: `(Optional - AWS) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. This argument only applies to Gateway deployments in AWS with ` + "`" + `mode` + "`" + ` set to ` + "`" + `HUB` + "`" + ` and ` + "`" + `security_type` + "`" + ` set to ` + "`" + `EGRESS` + "`" + `. If the argument is set to ` + "`" + `true` + "`" + `, the Gateway will be deployed using an AWS Gateway Load Balancer (GWLB). If not specified, the default value is ` + "`" + `false` + "`" + ` and the Gateway will be deployed using an internal AWS Network Load Balancer (NLB), which is a legacy deployment mode prior to AWS offering the GWLB. (Even though this argument is optional, it is recommended to specify a value explicitly, as the default value may change in the future).`,
				},
				resource.Attribute{
					Name:        "azure_gateway_lb",
					Description: `(Optional - Azure)`,
				},
				resource.Attribute{
					Name:        "mgmt_vpc_id",
					Description: `(Required - GCP) GCP VPC ID where the management interface of the Gateway is attached`,
				},
				resource.Attribute{
					Name:        "mgmt_security_group",
					Description: `(Required - EDGE Mode) AWS Security Group name, Azure Network Security Group ID or GCP Network Tag name to assigned to the management interface to permit management traffic to egress the Gateway. This must allow all outbound traffic for the Gateway to communicate with the Valtix Controller, Valtix S3 Bucket and for management DNS resolution.`,
				},
				resource.Attribute{
					Name:        "datapath_security_group",
					Description: `(Required - EDGE Mode) AWS Security Group name, Azure Network Security Group ID or GCP Network Tag name to assigned to the datapath interface to permit datapath traffic to ingress or egress the Gateway. It's recommended to leave this open so that all traffic can be sent and received by the Gateway where the Gateway Policy will control whether traffic is allowed or denied.`,
				},
				resource.Attribute{
					Name:        "min_instances",
					Description: `(Optional) Minimum number of instances per availability zone. If not specified, the default value is ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `(Optional) Maximum number of instances per availability zone. If not specified, the default value is ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `(Optional) TCP Port number that the Valtix orchestrated load balancers use for health checks to the Gateway instances. If not specified, the default value is ` + "`" + `65534` + "`" + `. A rule must be configured on the ` + "`" + `datapath_security_group` + "`" + ` to allow traffic to this TCP Port.`,
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
					Name:        "ntp_profile",
					Description: `(Optional) NTP Profile ID`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `(Optional) Gateway Settings block. This block can be repeated multiple times. See [Gateway Settings](#gateway-settings) for the block structure.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) User-defined Tags. This is a map of one or more user-defined key/value pairs that will be applied to each Gateway instance. The key is an unquoted name and the value is a quoted string. See [Gateway Tags](#gateway-tags) for the block structure. The Valtix Controller will add a Tag with keys of ` + "`" + `Name` + "`" + ` and ` + "`" + `valtix_acct` + "`" + ` during Gateway orchestration. If a user-defined tag for either of those keys is specified, the user-defined values will used in place of the Controller-defined values.`,
				},
				resource.Attribute{
					Name:        "instance_details",
					Description: `(Required - EDGE Mode) Gateway Instance Details. This block is only needed when deploying a Gateway in EDGE mode. This block should not be used when deploying a Gateway in HUB mode. For EDGE mode deployment, the block can be repeated multiple times for deploying Gateway instances in multiple Availability Zones. See [Instance Details](#instance-details) for the block structure. In EDGE mode, at least 1 block must be provided.`,
				},
				resource.Attribute{
					Name:        "gateway_lb_integration",
					Description: `(Optional - AWS) AWS Global Accelerator integration. This block is used when integrating the Gateway with the AWS Global Accelerator. The block can be repeated multiple times for integrating with multiple Global Accelerators. See [Global Accelerator](#global-accelerator) for the block structure. ## Instance Details This section is not required for AWS/Azure HUB mode as instance details are obtained from service VPC referenced in vpc_id attribute`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) Specifies the Availability Zone where the Gateway instance will be deployed`,
				},
				resource.Attribute{
					Name:        "mgmt_subnet",
					Description: `(Required) Specifies the VPC/VNet Subnet ID used to deploy the Gateway management interface within the specified Availability Zone`,
				},
				resource.Attribute{
					Name:        "datapath_subnet",
					Description: `(Required) Specifies the VPC/VNet Subnet ID used to deploy the Gateway datapath interface within the specified Availability Zone ## Global Accelerator This section is used for AWS Global Accelerator integration. Applies only to an Ingress Gateway deployed in AWS.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Set to ` + "`" + `AWS_GLOBAL_ACCELERATOR` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The Valtix name of the Global Accelerator`,
				},
				resource.Attribute{
					Name:        "awsga_resource_arn",
					Description: `(Required) The AWS ARN for the Global Accelerator`,
				},
				resource.Attribute{
					Name:        "awsga_endpoint_group_arn",
					Description: `(Required) The AWS ARN for the Global Accelerator Endpoint Group`,
				},
				resource.Attribute{
					Name:        "awsga_resource_name",
					Description: `(Optional) The AWS name of the Global Accelerator`,
				},
				resource.Attribute{
					Name:        "awsga_resource_fqdn",
					Description: `(Required) The AWS FQDN of the Global Accelerator ### To integrate with an AWS Global Accelerator ` + "`" + `` + "`" + `` + "`" + `hcl gateway_lb_integration { type = "AWS_GLOBAL_ACCELERATOR" name = "listener1" awsga_resource_arn = "arn:aws:globalaccelerator::902505820678:accelerator/d0c8cd60-e90c-4bb6-814e-6783716e1149" awsga_endpoint_group_arn = "arn:aws:globalaccelerator::902505820678:accelerator/d0c8cd60-e90c-4bb6-814e-6783716e1149/listener/80dab0fc/endpoint-group/d8c9bc581002" awsga_resource_name = "ga-for-hub" awsga_resource_fqdn = "a8ba3455f59690717.awsglobalaccelerator.com" } ` + "`" + `` + "`" + `` + "`" + ` ## Gateway Settings Gateway Settings are set of configurable settings that can be specified by the user and applied to the Gateway and its corresponding Gateway Instances. If a change to the settings from their defaults are necessary, the individual setting and the configured value can be applied to the list of settings as part of the settings block of the Gateway resource. ### Gateway AWS/Azure/GCP EBS/Disk Encryption Settings #### To enable AWS EBS Encryption using default KMS Key ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "gateway.aws.ebs.encryption.key.default" value = true } ` + "`" + `` + "`" + `` + "`" + ` #### To enable AWS EBS Encryption using specified KMS Key ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "gateway.aws.ebs.encryption.key.customer_key" value = "<KMS key ID>" } settings { name = "gateway.aws.ebs.encryption.key.customer_key" value = "arn:aws:kms:us-east-1:1112223333:key/67b90d29-5083-4166-a111-bfc810340f7d" } ` + "`" + `` + "`" + `` + "`" + ` #### To enable Azure Disk Encryption using specified KMS Key ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "gateway.azure.disk.encryption.key.customer_key" value = "<Disk Encryption Set Path>" } settings { name = "gateway.azure.disk.encryption.key.customer_key" value = "/subscriptions/1111111-33f9-4f5c-86e4-222222222/resourceGroups/valtix-rg/providers/Microsoft.Compute/diskEncryptionSets/valtixDiskEncryptionSet" } ` + "`" + `` + "`" + `` + "`" + ` #### To enable GCP Disk Encryption using specified Crypto Key ` + "`" + `` + "`" + `` + "`" + `hcl settings { name = "gateway.gcp.disk.encryption.key.customer_key" value = "<Crypto Key Path>" } settings { name = "gateway.gcp.disk.encryption.key.customer_key" value = "projects/security-icon-111111/locations/us-east1/keyRings/valtix-disk-encryption/cryptoKeys/key1" } ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Gateway resource`,
				},
				resource.Attribute{
					Name:        "gateway_gwlb_endpoints",
					Description: `(AWS only) AWS Gateway Load Balancer endpoints created in each of the AZs displayed in the format as follows: ` + "`" + `` + "`" + `` + "`" + `hcl gateway_gwlb_endpoints { endpoint_id = "vpce-047c749fc6f7e0c0d" network_interface_id = "eni-017eacdb23d2ebaf4" subnet_id = "subnet-0d61750e97caafd9d" } gateway_gwlb_endpoints { endpoint_id = "vpce-0707fa3f03c5064a7" network_interface_id = "eni-020464bd838461bca" subnet_id = "subnet-0fd61e07f200224f1" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gwlb_service_name",
					Description: `(AWS only) VPC Endpoint Service Name associated with the AWS Gateway Load Balancer. This name can be used by the AWS Terraform Provider for establishing a GWLB Endpoint connection.`,
				},
				resource.Attribute{
					Name:        "gwlb_service_id",
					Description: `(AWS only) VPC Endpoint Service ID associated with the AWS Gateway Load Balancer. This ID can be used by the AWS Terraform Provider for accepting GWLB Endpoint connections and assigning Principals.`,
				},
				resource.Attribute{
					Name:        "gateway_endpoint",
					Description: `For Gateways of ` + "`" + `security_type = INGRESS` + "`" + `, this represents the NLB endpoint (FQDN, IP Address) to be used as the target for the client communicating with any application protected by the Valtix Ingress Gateway. This information is most often used in a DNS A record (IP Address) or CNAME record (FQDN) to resolve the application FQDN to the Valtix Ingress Gateway endpoint. Valtix will receive traffic on this endpoint and proxy the traffic to the appropriate backend application based on the configured policy. For the Ingress Gateway, this attribute is populated for Gateways deployed in all CSPs (AWS, Azure, GCP, OCI). For Gateways of ` + "`" + `security_type = EGRESS` + "`" + `, this represents the NLB endpoint (IP Address) to be used as a target for routing traffic from the Spoke VPC/VNet/VCN to the Valtix Egress / East-West Gateway. Valtix will receive traffic from clients, and forward or proxy the traffic to the appropriate destination based on the configured policy. For the Egress / East-West Gateway, this attribute is only populated for non-AWS Gateways (Azure, GCP, OCI). For the AWS Gateways, traffic is routed to the AWS Transit Gateway (TGW) or Gateway Load Balancer (GWLB) endpoints. ## Import Gateway resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_gateway.aws-gw1 aws-gw1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Gateway resource`,
				},
				resource.Attribute{
					Name:        "gateway_gwlb_endpoints",
					Description: `(AWS only) AWS Gateway Load Balancer endpoints created in each of the AZs displayed in the format as follows: ` + "`" + `` + "`" + `` + "`" + `hcl gateway_gwlb_endpoints { endpoint_id = "vpce-047c749fc6f7e0c0d" network_interface_id = "eni-017eacdb23d2ebaf4" subnet_id = "subnet-0d61750e97caafd9d" } gateway_gwlb_endpoints { endpoint_id = "vpce-0707fa3f03c5064a7" network_interface_id = "eni-020464bd838461bca" subnet_id = "subnet-0fd61e07f200224f1" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "gwlb_service_name",
					Description: `(AWS only) VPC Endpoint Service Name associated with the AWS Gateway Load Balancer. This name can be used by the AWS Terraform Provider for establishing a GWLB Endpoint connection.`,
				},
				resource.Attribute{
					Name:        "gwlb_service_id",
					Description: `(AWS only) VPC Endpoint Service ID associated with the AWS Gateway Load Balancer. This ID can be used by the AWS Terraform Provider for accepting GWLB Endpoint connections and assigning Principals.`,
				},
				resource.Attribute{
					Name:        "gateway_endpoint",
					Description: `For Gateways of ` + "`" + `security_type = INGRESS` + "`" + `, this represents the NLB endpoint (FQDN, IP Address) to be used as the target for the client communicating with any application protected by the Valtix Ingress Gateway. This information is most often used in a DNS A record (IP Address) or CNAME record (FQDN) to resolve the application FQDN to the Valtix Ingress Gateway endpoint. Valtix will receive traffic on this endpoint and proxy the traffic to the appropriate backend application based on the configured policy. For the Ingress Gateway, this attribute is populated for Gateways deployed in all CSPs (AWS, Azure, GCP, OCI). For Gateways of ` + "`" + `security_type = EGRESS` + "`" + `, this represents the NLB endpoint (IP Address) to be used as a target for routing traffic from the Spoke VPC/VNet/VCN to the Valtix Egress / East-West Gateway. Valtix will receive traffic from clients, and forward or proxy the traffic to the appropriate destination based on the configured policy. For the Egress / East-West Gateway, this attribute is only populated for non-AWS Gateways (Azure, GCP, OCI). For the AWS Gateways, traffic is routed to the AWS Transit Gateway (TGW) or Gateway Load Balancer (GWLB) endpoints. ## Import Gateway resources can be imported using the resource ` + "`" + `name` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_gateway.aws-gw1 aws-gw1 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rule_set",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Resource for creating and managing a Policy Rule Set.  A Policy Rule Set is a list of Policy Rules that define the segmentation and security policy for protecting traffic.  The Policy Rule Set can be applied to multiple Gateways across multiple clouds to accommodate a dynamic multi-cloud policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Policy Rule Set`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Policy Rule Set`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specifies whether the Policy Rule Set is a Group Policy Rule Set. The only applicable value is ` + "`" + `GROUP` + "`" + `. If not specified, the Policy Rule Set operates as a Standalone (non-Group) Policy Rule Set.`,
				},
				resource.Attribute{
					Name:        "group_member_ids",
					Description: `(Required - Group). Ordered list of Policy Ruleset (Standalone) IDs that make up the components of the Policy Rulset (Group). This argument only applies when ` + "`" + `type` + "`" + ` is set to ` + "`" + `GROUP` + "`" + `. The list can contain zero or more IDs and is limited to a maximum of 100 IDs. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy Rule Set resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Policy Rule Set resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_policy_rule_set.ingress_policy_standalone 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Policy Rule Set resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Policy Rule Set resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_policy_rule_set.ingress_policy_standalone 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rules",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing the segmentation and security Policy Rules associated with a Policy Rule Set.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `(Required) ID of the Policy Rule Set. (e.g.,`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `A set of Rule blocks that define the segmentation and security Policy. Each Rule is represented by a separate block. Zero or more blocks can be specified, with a maximum of 2047 blocks. Structure is [documented below](#rule). ### Rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Rule`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Rule`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Required) ` + "`" + `ENABLED` + "`" + ` or ` + "`" + `DISABLED` + "`" + `. Specifies whether the rule is enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) ` + "`" + `ReverseProxy` + "`" + `, ` + "`" + `ForwardProxy` + "`" + `, and ` + "`" + `Forwarding` + "`" + `. Specifies how traffic will be processed by the Rule.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) ID of the Address Object (valtix_address_object). If not specified, the default value is the ID of the ` + "`" + `any` + "`" + ` Address Object. To obtain the ID of the pre-defined Address Objects (` + "`" + `any` + "`" + `, ` + "`" + `internet` + "`" + `, ` + "`" + `any-private-rfc1918` + "`" + `), use the Address Object Data Source (see example usage above). (e.g.,`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Optional) ID of the Address Object (valtix_address_object). If not specified, the default value is the ID of the ` + "`" + `any` + "`" + ` Address Object. To obtain the ID of the pre-defined Address Objects (` + "`" + `any` + "`" + `, ` + "`" + `internet` + "`" + `, ` + "`" + `any-private-rfc1918` + "`" + `), use the Address Object Data Source (see example usage above). (e.g.,`,
				},
				resource.Attribute{
					Name:        "fqdn_match_profile",
					Description: `(Optional) ID of the FQDN Match Profile`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) ID of the Service Object (valtix_service_object). The Service Object ` + "`" + `service_type` + "`" + ` argument must match the Rule ` + "`" + `type` + "`" + ` argument. (e.g.,`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action to take when traffic matches the Rule. Applicable values: ` + "`" + `Allow Log` + "`" + ` (log the event), ` + "`" + `Allow No Log` + "`" + ` (do not log the event), ` + "`" + `Deny Log` + "`" + ` (log the event), ` + "`" + `Deny No Log` + "`" + ` (do not log the event). Action names were changed in release 22.06 of the Valtix Terraform Provider. The previous action names (` + "`" + `ALLOW` + "`" + ` ` + "`" + `ALLOW_LOG` + "`" + `, ` + "`" + `DENY` + "`" + `, ` + "`" + `DENY_LOG` + "`" + `) are still valid, but Terraform will interpret the previous values as configuration changes until the new values are used.`,
				},
				resource.Attribute{
					Name:        "send_deny_reset",
					Description: `(Optional) Specifies whether a TCP Reset packet is sent by the Gateway when the session traffic is denied. Applies only if the Rule ` + "`" + `action` + "`" + ` or FQDN Filter Profile ` + "`" + `policy` + "`" + ` is set to ` + "`" + `Deny Log` + "`" + ` or ` + "`" + `Deny No Log` + "`" + `, and if the Rule ` + "`" + `type` + "`" + ` is set to ` + "`" + `Forwarding` + "`" + `. Applicable values: ` + "`" + `true` + "`" + ` (send a TCP Reset) or ` + "`" + `false` + "`" + ` (do not send a TCP Reset). If not specified, the default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_intrusion_profile",
					Description: `(Optional) ID of the IDS/IPS Profile (valtix_profile_network_intrusion). (e.g.,`,
				},
				resource.Attribute{
					Name:        "dlp_profile",
					Description: `(Optional) ID of the DLP Profile (valtix_profile_dlp). (e.g.,`,
				},
				resource.Attribute{
					Name:        "web_protection_profile",
					Description: `(Optional) ID of the WAF Profile (valtix_profile_web_protection). (e.g.,`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_profile",
					Description: `(Optional) ID of the FQDN Filter Profile (valtix_profile_fqdn_filter). (e.g.,`,
				},
				resource.Attribute{
					Name:        "anti_malware_profile",
					Description: `(Optional) ID of the AM Profile (valtix_profile_anti_malware). (e.g.,`,
				},
				resource.Attribute{
					Name:        "anti_virus_profile",
					Description: `(Optional - Deprecated) ID of the AV Profile (valtix_profile_anti_virus). (e.g.,`,
				},
				resource.Attribute{
					Name:        "url_filter",
					Description: `(Optional) ID of the URL Filter Profile (valtix_profile_url_filter). (e.g.,`,
				},
				resource.Attribute{
					Name:        "malicious_ip_profile",
					Description: `(Optional) ID of the Malicious IP Profile (valtix_profile_malicious_ip_profile). (e.g.,`,
				},
				resource.Attribute{
					Name:        "packet_capture_enabled",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Specifies the Packet Capture Profile to enable packet capture for each session that is processed by the Rule. If not specified, the default value is ` + "`" + `false` + "`" + `. ## Import Policy Rules resources can be imported using the Policy Rule Set resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_policy_rules.ingress_policy_standalone 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_anti_malware",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing an Anti-Malware (AM) Profile.  This resource replaces the former Anti-Virus (AV) Profile resource (` + "`" + `valtix_profile_anti_virus` + "`" + `).

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
					Description: `(Optional) Action to take when a Anti-Malware Network Threat is detected. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). If not specified, then the action assumed is ` + "`" + `Allow Log` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Anti-Malware Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Anti-Malware (AM) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_anti_malware.am_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Anti-Malware Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Anti-Malware (AM) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_anti_malware.am_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_anti_virus",
			Category:         "Resources",
			ShortDescription: ``,
			Description: ` (deprecated)
Resource for creating and managing an Anti Virus (AV) Profile (DEPRECATED).  This resource will remain available for two release cycles at which point it will be removed.  The Anti Malware (AM) Profile resource (` + "`" + `valtix_profile_anti_malware` + "`" + `) should be used in its place.

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
					Name:        "id",
					Description: `ID of the Anti-Virus Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Anti-Virus Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_application_threat",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Application Threat (WAF) Profile

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
					Name:        "disable_trustwave_ruleset",
					Description: `(Optional) Specifies whether the Trustwave Ruleset is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if not specified) is ` + "`" + `false` + "`" + `. If disabled, the Trustwave Ruleset auto-update will not occur. At least one Ruleset must be enabled (Trustwave or Custom). Both cannot be disabled at the same time. The CRS Ruleset cannot be disabled.`,
				},
				resource.Attribute{
					Name:        "auto_update_custom",
					Description: `(Optional) Auto Update of the Custom Ruleset version with a delay specified by ` + "`" + `delay_by_days_crs` + "`" + ` argument. Applicable values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days_custom",
					Description: `(Optional) Number of days to delay updating the Custom Ruleset version after it has been imported by the user. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately).`,
				},
				resource.Attribute{
					Name:        "custom_ruleset_version",
					Description: `(Optional) Custom Ruleset version. Applicable values can be found from within the UI as specified by the user when defining and importing the Custom Rulesets. Unless a specific version is desired, Valtix recommends using Auto Update. If this argument is specified, Auto Update of Custom Ruleset is disabled and the Profile will use the specified Custom Ruleset version.`,
				},
				resource.Attribute{
					Name:        "disable_custom_ruleset",
					Description: `(Optional) Specifies whether the Custom Ruleset is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if not specified) is ` + "`" + `false` + "`" + `. If disabled, the Custom Ruleset auto-update will not occur. At least one Ruleset must be enabled (Trustwave or Custom). Both cannot be disabled at the same time. The CRS Ruleset cannot be disabled.`,
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
					Name:        "id",
					Description: `ID of the Application Threat (WAF) Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Application Threat (WAF) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_application_threat.waf_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Application Threat (WAF) Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Application Threat (WAF) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_application_threat.waf_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Name of certificate that is defined in the ` + "`" + `valtix_certificate` + "`" + ` resource that contains the desired certificate`,
				},
				resource.Attribute{
					Name:        "min_tls_version",
					Description: `(Optional) The minimum TLS version to use when establishing a secure frontend and backend connection when processing traffic via a forward or reverse proxy. Applicable values are: ` + "`" + `TLS_VERSION_1_3` + "`" + `, ` + "`" + `TLS_VERSION_1_2` + "`" + `, ` + "`" + `TLS_VERSION_1_1` + "`" + `, ` + "`" + `TLS_VERSION_1_0` + "`" + `. If not specified, the default value is ` + "`" + `TLS_VERSION_1_0` + "`" + `. IMPORTANT: It is required to use the same value for ` + "`" + `min_tls_version` + "`" + ` in all Decryption Profiles that are used by Service Objects referenced by a Policy Ruleset or Policy Ruleset Group. If different values are used, the value that will be applied cannot be predetermined. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Decryption Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Decryption Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_decryption.decryption_profile1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Decryption Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Decryption Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_decryption.decryption_profile1 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "id",
					Description: `ID of the Diagnostic Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Diagnostic Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_diagnostic.awsdiag1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Diagnostic Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Diagnostic Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_diagnostic.awsdiag1 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "id",
					Description: `ID of the Data Loss Prevention (DLP) Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Data Loss Prevention (DLP) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_dlp.dlp1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Data Loss Prevention (DLP) Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Data Loss Prevention (DLP) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_dlp.dlp1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_fqdn",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing an FQDN Profile that can be used in a Policy Ruleset for Matching and Filtering purposes.

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
					Name:        "type",
					Description: `(Optional) Specifies whether the FQDN Profile is a Group FQDN Profile. The only applicable value is ` + "`" + `GROUP` + "`" + `. If not specified, the FQDN Profile operates as Standalone (non-Group).`,
				},
				resource.Attribute{
					Name:        "group_member_ids",
					Description: `(Required - Group). Ordered list of FQDN Profile (Standalone) IDs that make up the components of the FQDN Profile (Group). This argument only applies when ` + "`" + `type` + "`" + ` is set to ` + "`" + `GROUP` + "`" + `. The list can contain zero or more IDs and is limited to a maximum of 30 IDs. The resulting aggregated FQDN Profile is limited to a total of 254 FQDN Profile Filter List blocks.`,
				},
				resource.Attribute{
					Name:        "uncategorized_fqdn_filter",
					Description: `(Required) Uncategorized FQDN Filter action for any FQDN that does not match the FQDNs defined in the ` + "`" + `fqdn_filter_list` + "`" + ` resource and is not represented by any vendor category (whether specified or not). This argument is required no matter the ` + "`" + `type` + "`" + ` specified, but only applies when the Profile operates as Standalone. When operating as part of a Group, the Group setting will apply. Structure [defined below](#uncategorized-fqdn-filter).`,
				},
				resource.Attribute{
					Name:        "default_fqdn_filter",
					Description: `(Required) Default FQDN Filter action for any FQDN that does not match the FQDNs defined in the ` + "`" + `fqdn_filter_list` + "`" + ` resource or is not matched by the ` + "`" + `uncategorized_fqdn_filter` + "`" + ` resource (if specified). This should be the last resource specified in the list of resources. This argument is required no matter the ` + "`" + `type` + "`" + ` specified, but only applies when the Profile operates as Standalone. When operating as part of a Group, the Group setting will apply. Structure [defined below](#default-fqdn-filter).`,
				},
				resource.Attribute{
					Name:        "no_fqdn_deny",
					Description: `(Optional) Deny traffic when no FQDN is found in the packet. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `false` + "`" + `. ### FQDN Profile Group Arguments (Match)`,
				},
				resource.Attribute{
					Name:        "group_member_ids",
					Description: `(Required - Group). Ordered list of FQDN Profile (Standalone) IDs that make up the components of the FQDN Profile (Group). This argument only applies when ` + "`" + `type` + "`" + ` is set to ` + "`" + `GROUP` + "`" + `. The list can contain zero or more IDs and is limited to a maximum of 30 IDs. The resulting aggregated FQDN Profile is limited to a total of 254 FQDN Profile Filter List blocks. ### FQDN Profile Standalone Arguments (Filter)`,
				},
				resource.Attribute{
					Name:        "no_fqdn_deny",
					Description: `(Optional) Deny traffic when no FQDN is found in the packet. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default value: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_list",
					Description: `(Required) One or more blocks, where each block is a row in the FQDN Filter Profile (maximum of 254 blocks). Structure [defined below](#fqdn-filter-list).`,
				},
				resource.Attribute{
					Name:        "uncategorized_fqdn_filter",
					Description: `(Required) Uncategorized FQDN Filter action for any FQDN that does not match the FQDNs defined in the ` + "`" + `fqdn_filter_list` + "`" + ` resource and is not represented by any vendor category (whether specified or not). This argument is required no matter the ` + "`" + `type` + "`" + ` specified, but only applies when the Profile operates as Standalone. When operating as part of a Group, the Group setting will apply. Structure [defined below](#uncategorized-fqdn-filter).`,
				},
				resource.Attribute{
					Name:        "default_fqdn_filter",
					Description: `(Required) Default FQDN Filter action for any FQDN that does not match the FQDNs defined in the ` + "`" + `fqdn_filter_list` + "`" + ` resource or is not matched by the ` + "`" + `uncategorized_fqdn_filter` + "`" + ` resource (if specified). This should be the last resource specified in the list of resources. This argument is required no matter the ` + "`" + `type` + "`" + ` specified, but only applies when the Profile operates as Standalone. When operating as part of a Group, the Group setting will apply. Structure [defined below](#default-fqdn-filter). ### FQDN Profile Standalone Arguments (Match)`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_list",
					Description: `(Required) One or more blocks, where each block is a row in the FQDN Filter Profile (maximum of 254 blocks). Structure [defined below](#fqdn-filter-list). ### FQDN Profile Filter List Arguments (Filter)`,
				},
				resource.Attribute{
					Name:        "fqdn_list",
					Description: `(Required) List of FQDNs (maximum of 60 FQDNs per list, combined with categories; maximum 255 characters per FQDN). Applicable values are Perl Compatible Regular Expression (PCRE) patterns representing FQDNs. When specifying a multi-level domain (e.g., ` + "`" + `www.example.com` + "`" + `), it's important to escape the ` + "`" + `.` + "`" + ` character (e.g., ` + "`" + `www\\.example\\.com` + "`" + `) otherwise it will be treated as a wildcard for any single character. Structure [defined below](#fqdn-list).`,
				},
				resource.Attribute{
					Name:        "vendor_category_list",
					Description: `(Optional) List of pre-defined Vendor Categories (maximum 60 categories per list, combined with FQDNs). Structure [defined below](#vendor-category-list).`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Action to take when an FQDN matches an entry in the ` + "`" + `fqnd_list` + "`" + ` or ` + "`" + `vendor_category_list` + "`" + `. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event).`,
				},
				resource.Attribute{
					Name:        "decryption_exception",
					Description: `(Optional) When used in conjunction with a proxy Rule (ForwardProxy, ReverseProxy), instructs the proxy engine to bypass decryption. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `. ### FQDN Profile Filter List Arguments (Match)`,
				},
				resource.Attribute{
					Name:        "fqdn_list",
					Description: `(Required) List of FQDNs (maximum of 60 FQDNs per list; maximum 255 characters per FQDN). Applicable values are Perl Compatible Regular Expression (PCRE) patterns representing FQDNs. When specifying a multi-level domain (e.g., ` + "`" + `www.example.com` + "`" + `), it's important to escape the ` + "`" + `.` + "`" + ` character (e.g., ` + "`" + `www\\.example\\.com` + "`" + `) otherwise it will be treated as a wildcard for any single character. Structure [defined below](#fqdn-list).`,
				},
				resource.Attribute{
					Name:        "decryption_exception",
					Description: `(Optional) When used in conjunction with a proxy Rule (ForwardProxy, ReverseProxy), instructs the proxy engine to bypass decryption. Applicable values: ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `. ### FQDN List ` + "`" + `` + "`" + `` + "`" + `hcl fqdn_list = [ "www\\.website1\\.com", ".`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the FQDN Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import FQDN Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_fqdn.fqdn1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the FQDN Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import FQDN Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_fqdn.fqdn1 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "id",
					Description: `ID of the L7DOS Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import L7 DOS (L7DOS) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_l7dos.l7dos1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the L7DOS Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import L7 DOS (L7DOS) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_l7dos.l7dos1 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) Name of the Log Forwarding profile`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) One of ` + "`" + `REMOTE_SYSLOG` + "`" + `, ` + "`" + `SPLUNK` + "`" + `, ` + "`" + `DATADOG` + "`" + `, ` + "`" + `GCPLOGGING_FROM_GATEWAY` + "`" + `, ` + "`" + `SUMO_LOGIC` + "`" + `, ` + "`" + `AWS_S3` + "`" + `, ` + "`" + `MS_SENTINEL` + "`" + ``,
				},
				resource.Attribute{
					Name:        "siem_vendor",
					Description: `(Deprecated) One of ` + "`" + `REMOTE_SYSLOG` + "`" + `, ` + "`" + `SPLUNK` + "`" + `, ` + "`" + `DATADOG` + "`" + `, ` + "`" + `GCPLOGGING_FROM_GATEWAY` + "`" + ``,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Specifies whether the Log Forwarding Profile is a Group Log Forwarding Profile. The only applicable value is ` + "`" + `GROUP` + "`" + `. If not specified, the Log Forwarding Profile operates as a Standalone (non-Group) Log Forwarding Profile.`,
				},
				resource.Attribute{
					Name:        "group_member_ids",
					Description: `(Required - Group). Ordered list of Log Forwarding Profile (Standalone) IDs that make up the components of the Log Forwarding Profile (Group). This argument only applies when ` + "`" + `type` + "`" + ` is set to ` + "`" + `GROUP` + "`" + `. The list can contain zero or more IDs and is limited to a maximum of 5 IDs. ### Destination-specific Arguments ### Splunk`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) HTTPS endpoint URL`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) HTTPS auth token`,
				},
				resource.Attribute{
					Name:        "splunk_index",
					Description: `(Required) Index name where the logs will be stored ### Syslog`,
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
					Name:        "syslog_flow_logs",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Forward Flow Log events.`,
				},
				resource.Attribute{
					Name:        "syslog_firewall_events",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Forward L4 Firewall (L4_FW) events.`,
				},
				resource.Attribute{
					Name:        "syslog_https_logs",
					Description: `(Optional) ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Forward TLS Log/Error (TLS_LOG, TLS_ERROR) events.`,
				},
				resource.Attribute{
					Name:        "network_threat_severity",
					Description: `(Optional) One of ` + "`" + `Alert` + "`" + `, ` + "`" + `Emergency` + "`" + `, ` + "`" + `Critical` + "`" + `, ` + "`" + `Error` + "`" + `, ` + "`" + `Warning` + "`" + `, ` + "`" + `Notice` + "`" + `, ` + "`" + `Info` + "`" + `, or ` + "`" + `Debug` + "`" + `. Forward Network Threat (IPS) events with the given severity only.`,
				},
				resource.Attribute{
					Name:        "web_attack_severity",
					Description: `(Optional) One of ` + "`" + `Alert` + "`" + `, ` + "`" + `Emergency` + "`" + `, ` + "`" + `Critical` + "`" + `, ` + "`" + `Error` + "`" + `, ` + "`" + `Warning` + "`" + `, ` + "`" + `Notice` + "`" + `, ` + "`" + `Info` + "`" + `, ` + "`" + `Debug` + "`" + `. Forward Web Attack (WAF) events with the given severity only. ### Datadog`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) HTTPS endpoint URL`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) HTTPS auth token ### GCP Logging`,
				},
				resource.Attribute{
					Name:        "log_name",
					Description: `(Optional) GCP Logging name where the logs will be stored. If not specified, the default name is ` + "`" + `valtix-gateway-logs` + "`" + `. ### Sumo Logic`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required) Collector endpoint URL ### AWS S3 Bucket`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) Friendly name of the onboarded AWS account where the S3 Bucket resides`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Globally unique name of the S3 Bucket ### MS Sentinel`,
				},
				resource.Attribute{
					Name:        "log_analytics_log_type",
					Description: `(Required) Name of the MS Sentinel table used to store the logs`,
				},
				resource.Attribute{
					Name:        "log_analytics_workspace_id",
					Description: `(Required) ID of the MS Sentinel workspace`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) Shared key / primary key used to authenticate with MS Sentinel ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Log Forwarding Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Log Forwarding Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_log_forwarding.splunk 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Log Forwarding Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Log Forwarding Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_log_forwarding.splunk 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_malicious_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Malicious IP (MIP) Profile

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
					Name:        "id",
					Description: `ID of the Malicious IP Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Malicious IP (MIP) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_malicious_ip.mips_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Malicious IP Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Malicious IP (MIP) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_malicious_ip.mips_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_network_intrusion",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Network Intrusion (IPS/IDS) Profile

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
					Description: `(Optional) Auto Update the Talos Ruleset version. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if unspecified) is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) Number of days to delay applying the Talos Ruleset version after it has been published by Valtix. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately). New Rulesets as published as soon as updates are available from the Vendor and validation testing is completed by Valtix.`,
				},
				resource.Attribute{
					Name:        "talos_ruleset_version",
					Description: `(Optional) Talos Ruleset version. Applicable values can be found from within the UI. The Rulesets are published frequently. Unless a specific version is desired, Valtix recommends using Auto Update. If this argument is specified, Auto Update of Talos Ruleset is disabled and the Profile will use the specified Talos Ruleset version.`,
				},
				resource.Attribute{
					Name:        "disable_talos_ruleset",
					Description: `(Optional) Specifies whether the Talos Ruleset is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if not specified) is ` + "`" + `false` + "`" + `. If disabled, the Talos Ruleset auto-update will not occur. At least one Ruleset must be enabled (Talos or Custom). Both cannot be disabled at the same time.`,
				},
				resource.Attribute{
					Name:        "auto_update_custom",
					Description: `(Optional) Auto Update the Custom Ruleset version. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if unspecified) is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "delay_by_days_custom",
					Description: `(Optional) Number of days to delay applying the Custom Ruleset version after it has been imported by the user. Applicable values are integers from ` + "`" + `0` + "`" + ` to ` + "`" + `30` + "`" + `. A value of ` + "`" + `0` + "`" + ` means immediate update (0 days). The default value is ` + "`" + `0` + "`" + ` (immediately).`,
				},
				resource.Attribute{
					Name:        "custom_ruleset_version",
					Description: `(Optional) Custom Ruleset version. Applicable values can be found from within the UI as specified by the user when defining and importing the Custom Rulesets. Unless a specific version is desired, Valtix recommends using Auto Update. If this argument is specified, Auto Update of Custom Ruleset is disabled and the Profile will use the specified Custom Ruleset version.`,
				},
				resource.Attribute{
					Name:        "disable_custom_ruleset",
					Description: `(Optional) Specifies whether the Custom Ruleset is enabled or disabled. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default (if not specified) is ` + "`" + `false` + "`" + `. If disabled, the Custom Ruleset auto-update will not occur. At least one Ruleset must be enabled (Talos or Custom). Both cannot be disabled at the same time.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Network Intrusion (IDS/IPS) Network Threat is detected for Rules that fall into the Profile set (the entire set or Rules defined by the configuration of the Profile). Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). If not specified, then the action assumed is ` + "`" + `Rule Default` + "`" + `, the action defined in the Policy Ruleset Rule. This action can be overridden for each Policy, [Category](#categories) and [Class](#classes) of Rules.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Pre-defined Talos Ruleset to use. Applicable values: ` + "`" + `CONNECTIVITY` + "`" + `, ` + "`" + `BALANCED` + "`" + `, ` + "`" + `SECURITY` + "`" + `, ` + "`" + `MAX_DETECT` + "`" + `.`,
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
					Description: `(Optional) Capture packets (PCAP) when traffic matches a Ruleset rule. Valid values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
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
					Description: `(Required) Name of a Talos attack Category`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when a Network Intrusion (IDS/IPS) Network Threat is detected for Rules that fall into the Category set. Applicable values: ` + "`" + `Allow Log` + "`" + ` (allow and log the event), ` + "`" + `Allow No Log` + "`" + ` (allow and do not log the event), ` + "`" + `Deny Log` + "`" + ` (deny and log the event), ` + "`" + `Deny No Log` + "`" + ` (deny and do not log the event). This action is an override to the Policy action. If not specified, then the action assumed is the action defined for the Policy. ## Classes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Talos attack Class`,
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
					Description: `(Optional) List of Rule IDs to filter`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) Generator ID (GID) necessary when suppressing Talos Ruleset preprocessor rule IDs (lower numbered rule IDs). If not specified, default value is ` + "`" + `1` + "`" + `, which is the GID for Talos Standard Text Rules (higher numbered rule IDs). ## Profile Event Filter`,
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
					Name:        "id",
					Description: `ID of the Network Intrusion (IDS/IPS) Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Network Intrusion (IDS/IPS) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_network_intrusion.ips_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Network Intrusion (IDS/IPS) Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Network Intrusion (IDS/IPS) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_network_intrusion.ips_auto 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_ntp",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing an NTP Profile. The Gateway uses NTP to synchronize its time. The NTP communcication uses the Gateway Management interface.  The NTP Profile allows a user to specify a set of custom NTP servers to use in place of the default servers.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the NTP Profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the NTP Profile`,
				},
				resource.Attribute{
					Name:        "server_list",
					Description: `(Required) List of NTP servers to be applied to the Gateway NTP configuration ### Default NTP Servers The default NTP servers that are configured on the Gateway for each CSP are as follows:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NTP Profile resource that can be referenced in other resources (e.g.,`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the NTP Profile resource that can be referenced in other resources (e.g.,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_packet_capture",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Packet Capture (PCAP) Profile

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
					Name:        "id",
					Description: `ID of the Packet Capture Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Packet Capture (PCAP) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_packet_capture.awspcap1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Packet Capture Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Packet Capture (PCAP) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_packet_capture.awspcap1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_urlfilter",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a URL Filter (URL) Profile

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
					Description: `(Required) One or more blocks, where each block is a row in the URL Filter Profile (maximum of 254 blocks). Structure [defined below](#url-list).`,
				},
				resource.Attribute{
					Name:        "uncategorized_url_filter",
					Description: `(Required) Uncategorized URL Filter action for any URL that does not match the URLs defined in the ` + "`" + `url_filter_list` + "`" + ` resource and is not represented by any vendor category (whether specified or not). Structure [defined below](#uncategorized-url-filter).`,
				},
				resource.Attribute{
					Name:        "default_url_filter",
					Description: `(Required) Default URL Filter action for any URL that does not match the URLs defined in the ` + "`" + `url_filter_list` + "`" + ` resource or is not matched by the ` + "`" + `uncategorized_url_filter` + "`" + ` resource (if specified). This should be the last resource specified in the list of resources. Structure [defined below](#default-url-filter).`,
				},
				resource.Attribute{
					Name:        "deny_response",
					Description: `(Optional) Specifies the HTTP response message to return back to the client when the URL is denied. This response is in addition to the ` + "`" + `return_status` + "`" + ` response code that is specified for each URL Filter List (` + "`" + `url_filter_list` + "`" + `) block. ## URL List`,
				},
				resource.Attribute{
					Name:        "url_list",
					Description: `(Required) List of URLs (maximum of 60 URLs per list, combined with categories; maximum of 2048 characters per URL). Applicable values are Perl Compatible Regular Expression (PCRE) patterns representing FQDNs. When specifying a multi-level domain (e.g., ` + "`" + `www.example.com` + "`" + `), it's important to escape the ` + "`" + `.` + "`" + ` character (e.g., ` + "`" + `www\\.example\\.com` + "`" + `) otherwise it will be treated as a wildcard for any single character. Structure [defined below](#url-list).`,
				},
				resource.Attribute{
					Name:        "vendor_category_list",
					Description: `(Optional) List of pre-defined Vendor Categories (maximum of 60 categories per list, combined with URLs). Structure [defined below](#vendor-category-list).`,
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
					Description: `(Optional) HTTP status code to return when URLs are denied. This argument only applies to resources that have a ` + "`" + `policy` + "`" + ` set to ` + "`" + `Deny Log` + "`" + ` or ` + "`" + `Deny No Log` + "`" + `. ## URL List ` + "`" + `` + "`" + `` + "`" + `hcl url_list = [ "https://.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the URL Filter Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import URL Filter (URL) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_urlfilter.url1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the URL Filter Profile resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "profile_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import URL Filter (URL) Profile resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_profile_urlfilter.url1 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_object",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `
Resource for creating and managing a Service Object.  A Service Object is used in a Policy Ruleset Rule to define how traffic will be matched and processed by the Rule.

~> **Note on SNI for a ReverseProxy Service Object**
  If multiple applications are accessed using the same port (e.g., port 443), then an SNI must be specified to differentiate the requests associated with each application.

~> **Note on Service Type for a ReverseProxy Service Object**
  If multiple ReverseProxy Service Objects that use the same port (e.g., port 443) are applied to the same Policy Ruleset, the Service Type (` + "`" + `service_type` + "`" + `) must be represented consistently by either an L7 proxy (` + "`" + `HTTP` + "`" + `, ` + "`" + `HTTPS` + "`" + `, ` + "`" + `WEBSOCKET` + "`" + `, ` + "`" + `WEBSOCKET_S` + "`" + `) or an L4/L5 proxy (` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + `, ` + "`" + `TLS` + "`" + `).  L7 and L4/L5 proxies cannot be mixed for the same Policy Ruleset on the same port.  For further discussion or questions, please contact [Valtix Support](mailto:support@valtix.com).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Service Object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Service Object`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) Type of Service Object. Applicable values are ` + "`" + `ReverseProxy` + "`" + `, ` + "`" + `ForwardProxy` + "`" + `, or ` + "`" + `Forwarding` + "`" + `. ### ReverseProxy Arguments`,
				},
				resource.Attribute{
					Name:        "transport_mode",
					Description: `(Required) Proxy to use for processing traffic. Applicable values for secure proxy are ` + "`" + `HTTPS` + "`" + `, ` + "`" + `TLS` + "`" + `, or ` + "`" + `WEBSOCKET_S` + "`" + `. Applicable values for non-secure proxy are ` + "`" + `HTTP` + "`" + `, ` + "`" + `TCP` + "`" + `, or ` + "`" + `WEBSOCKET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tls_profile",
					Description: `(Optional) Application certificate to issue and TLS settings to use when negotiating a secure TLS session between the client and Gateway. Applicable value is a Decryption Profile (` + "`" + `valtix_profile_decryption` + "`" + `) ID.`,
				},
				resource.Attribute{
					Name:        "client_tls_profile",
					Description: `(Optional) Certificate to use to authenticate the client certificate in a mutual TLS handshake. Applicable value is a Decryption Profile (` + "`" + `valtix_profile_decryption` + "`" + `) ID.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Destination port(s) to use for processing traffic. Multiple ports or port ranges can be specified by repeating this block. Structure is [defined below](#reverseproxy-port-arguments).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol to use for processing traffic. Applicable values are ` + "`" + `TCP` + "`" + ` or ` + "`" + `UDP` + "`" + `. If not specified, the default value is ` + "`" + `TCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sni",
					Description: `(Optional) SNIs to use as an L5 match to distinguish multiple TLS applications that use the same port. Applicable values are one or more FQDN strings that are used to match against the TLS Hello header Service Name Indication (SNI). The match will issue the corresponding Certificate defined by the ` + "`" + `tls_profile` + "`" + ` and establish a backend connection to the application defined by the ` + "`" + `backend_address_group` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "l7dos_profile",
					Description: `(Optional) L7 DOS Profile to use to rate limit transactions. Applicable value is an L7 DOS Profile (` + "`" + `valtix_profile_l7dos` + "`" + `) ID.`,
				},
				resource.Attribute{
					Name:        "backend_address_group",
					Description: `(Required) Server or load balancer to use for establishing the connection between the Gateway and server. Applicable value is a Reverse Proxy Target Address Object (` + "`" + `valtix_address_object` + "`" + `) ID. ### ReverseProxy Port Arguments This block can be specified multiple times to define a list of one or more Destination and Backend Ports. ` + "`" + `` + "`" + `` + "`" + `hcl port { destination_ports = "443" backend_ports = "443" } # If the listen ports are continuous and target ports are continuous port { destination_ports = "80-100" backend_ports = "8000-8020" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Required) Destination port number(s) as a single port or a continuous range of ports (e.g, ` + "`" + `80` + "`" + ` or ` + "`" + `80-100` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "backend_ports",
					Description: `(Required) Target backend port number(s) specified as a single port or as a continuous range of ports (e.g, ` + "`" + `80` + "`" + ` or ` + "`" + `80-100` + "`" + `). ~>`,
				},
				resource.Attribute{
					Name:        "transport_mode",
					Description: `(Required) Proxy to use for processing traffic. Applicable values for non-secure proxy are ` + "`" + `HTTP` + "`" + ` or ` + "`" + `WEBSOCKET` + "`" + `. Applicable values for secure proxy are ` + "`" + `HTTPS` + "`" + ` or ` + "`" + `WEBSOCKET_S` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Destination port(s) to use for processing traffic. Multiple ports or port ranges can be specified by repeating this block. Structure is [defined below](#forwardproxy-port-arguments).`,
				},
				resource.Attribute{
					Name:        "tls_profile",
					Description: `(Optional) Intermediate, self-signed or enterprise-signed certificate to be issued and TLS settings to be used when negotiating a secure TLS session between the client and Gateway. Applicable value is a Decryption Profile (` + "`" + `valtix_profile_decryption` + "`" + `) ID. ### ForwardProxy Port Arguments This block can be specified multiple times to define a list of Destination Ports. ` + "`" + `` + "`" + `` + "`" + `hcl port { destination_ports = "443" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Required) Destination port number(s) as a single port or a continuous range of ports (e.g, ` + "`" + `80` + "`" + ` or ` + "`" + `80-100` + "`" + `) <br><br>For examples, see [ForwardProxy Examples](#forwardproxy-examples) ### Forwarding Arguments`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Protocol to use for processing traffic. Applicable values are ` + "`" + `TCP` + "`" + `, ` + "`" + `UDP` + "`" + ` or ` + "`" + `ICMP` + "`" + `. If not specified, the default value is ` + "`" + `TCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Destination port(s) to use for processing traffic. Multiple ports or port ranges can be specified by repeating this block. Structure is [defined below](#forwarding-port-arguments).`,
				},
				resource.Attribute{
					Name:        "source_nat",
					Description: `(Optional) Specifies whether Source Network Address Translation (SNAT) should be performed on the session. Applicable values are ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. If not specified, the default value is ` + "`" + `true` + "`" + `. (Even though this argument is optional, it is recommended to specify a value explicitly, as the default value may change in the future). ### Forwarding Port Arguments This block can be specified multiple times to define a list of Destination Ports. ` + "`" + `` + "`" + `` + "`" + `hcl port { destination_ports = "443" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Required) Destination port number(s) as a single port or a continuous range of ports (e.g, ` + "`" + `80` + "`" + ` or ` + "`" + `80-100` + "`" + `) <br><br>For examples, see [Forwarding Examples](#forwarding-examples) ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Service Object resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Service Object resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_service_object.app1_svc_http 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Service Object resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Service Object resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_service_object.app1_svc_http 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `ID of the Services VPC/VNet resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Cloud specific ID of the Services VPC/VNet`,
				},
				resource.Attribute{
					Name:        "service_vpc_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Service VPC/VNet resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_service_vpc.aws_service_vpc 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Services VPC/VNet resource that can be referenced in other resources (e.g.,`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Cloud specific ID of the Services VPC/VNet`,
				},
				resource.Attribute{
					Name:        "service_vpc_id",
					Description: `(Deprecated) Same as the ` + "`" + `id` + "`" + ` attribute ## Import Service VPC/VNet resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_service_vpc.aws_service_vpc 10 ` + "`" + `` + "`" + `` + "`" + ``,
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
				resource.Attribute{
					Name:        "transit_gateway_attachment_subnets",
					Description: `(Optional - AWS) List of Subnet Ids used to attach to the Transit Gateway. By default, a random subnet in each availability zone is selected ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Spoke VPC resource ## Import Spoke VPC resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_spoke_vpc.valtix_spoke 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Spoke VPC resource ## Import Spoke VPC resources can be imported using the resource ` + "`" + `id` + "`" + `: ` + "`" + `` + "`" + `` + "`" + `hcl $ terraform import valtix_spoke_vpc.valtix_spoke 10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"valtix_address_object":                        0,
		"valtix_alert_profile":                         1,
		"valtix_alert_rule":                            2,
		"valtix_certificate":                           3,
		"valtix_cloud_account":                         4,
		"valtix_cloud_account_log_profile_association": 5,
		"valtix_external_id":                           6,
		"valtix_gateway":                               7,
		"valtix_policy_rule_set":                       8,
		"valtix_policy_rules":                          9,
		"valtix_profile_anti_malware":                  10,
		"valtix_profile_anti_virus":                    11,
		"valtix_profile_application_threat":            12,
		"valtix_profile_decryption":                    13,
		"valtix_profile_diagnostic":                    14,
		"valtix_profile_dlp":                           15,
		"valtix_profile_fqdn":                          16,
		"valtix_profile_l7dos":                         17,
		"valtix_profile_log_forwarding":                18,
		"valtix_profile_malicious_ip":                  19,
		"valtix_profile_network_intrusion":             20,
		"valtix_profile_ntp":                           21,
		"valtix_profile_packet_capture":                22,
		"valtix_profile_urlfilter":                     23,
		"valtix_service_object":                        24,
		"valtix_service_vpc":                           25,
		"valtix_spoke_vpc":                             26,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
