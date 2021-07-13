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
Manages address objects that can be used as source (for egress) or a target (backend) for ingress.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the address object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the address object`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) List of IP Address or FQDN. If the backend_address value is true, then this must be a list with one entry only.`,
				},
				resource.Attribute{
					Name:        "backend_address",
					Description: `(Optional) true/false. This must be set to true if this address object runs an application that's proxied/protected by the valtix_gw. The`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Required), name of the tag`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(Required), value of the tag`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Optional) Restrict the tag_key and tag_value for only the given csp_account_name (e.g "gcp1". This is the name of the account added via valtix_cloud_account)`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(OptionalRestrict the tag_key and tag_value for only the given vpc_id`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Restrict the tag_key and tag_value for only the given region (e.g "us-east1")`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) Resource group name ## DYNAMIC_VPC`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the VPC is defined`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC Id`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) Resource group name ## DYNAMIC_SUBNET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the VPC is defined`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC Id`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Required) Subnet Id`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) Resource group name ## DYNAMIC_INSTANCE`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the VPC is defined`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC Id`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required) Instance Id`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) Resource group name ## DYNAMIC_SECURITY_GROUP`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the VPC is defined`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC Id`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `(Required) Security Group Id`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) Resource group name ## DYNAMIC_USER_DEFINED_TAG`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the VPC is defined`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC Id`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Required), name of the tag`,
				},
				resource.Attribute{
					Name:        "tag_value",
					Description: `(Required), value of the tag`,
				},
				resource.Attribute{
					Name:        "resource_group",
					Description: `(Azure only) Resource group name ## DYNAMIC_SERVICE_ENDPOINTS`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "service_endpoint_name",
					Description: `(Required) Service end point name ## GEO_IP`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Geo IP Value ## STORAGE_BUCKET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) This is the name of the account added via valtix_cloud_account that selects the csp account to get the VPC`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Bucket Name`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) One of "SlackWebHook", "PagerDutyEventApi", "ServiceNowWebHook"`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required for ServiceNow and Pagerduty) Key`,
				},
				resource.Attribute{
					Name:        "integration_url",
					Description: `(Required for ServiceNow and Slack) Webhook URL or ServiceNow Url`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) Name of the alert profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description`,
				},
				resource.Attribute{
					Name:        "alert_profile",
					Description: `(Required) Alert profile id`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) One of "Type_Inventory", "Type_SystemLogs",`,
				},
				resource.Attribute{
					Name:        "sub_type",
					Description: `(Required) For Type_Inventory, valid sub_type is "SubType_InventoryGuardRails", For Type_SystemLogs, valid sub_type is "SubType_SystemLogsGateway" or "SubType_SystemLogsAccount"`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) One of "Info", "Medium", "High". For SystemLogs "Critical" and "Warning" are other severity types`,
				},
				resource.Attribute{
					Name:        "is_active",
					Description: `(Required) true or false`,
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

Manages certificates that are presented to the end user when they access valtix gw proxy

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
					Description: `(Required) "IMPORT_CONTENTS", "AWS_KMS", "AWS_SECRET", "AZURE_KEY_VAULT_SECRET", "GCP_SECRET"`,
				},
				resource.Attribute{
					Name:        "certificate_body",
					Description: `(Required) Certificate content. Provide a string with the whole content here or if it's on a file system use ` + "`" + `file("filename")` + "`" + ``,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `(Optional) Certificate chain content. Provide a string with the whole content here or if it's on a file system use ` + "`" + `file("filename")` + "`" + `. This is optional if the certificate has a chain available separately and not part of the certificate_body ## Additional arguments based on the certificate_type ### IMPORT_CONTENTS`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required for IMPORT_CONTENTS) Certificate content. Provide a string with the whole content here or if it's on a file system use ` + "`" + `file("filename")` + "`" + ` ### AWS_KMS`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key that's stored on GCP secrets manager, provide the csp_account_name (GCP) that stores the secrets key.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the key/secret is location`,
				},
				resource.Attribute{
					Name:        "private_key_cipher_text",
					Description: `(Required) KMS encrypted cipher text ### AWS_SECRET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key that's stored on GCP secrets manager, provide the csp_account_name (GCP) that stores the secrets key`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the key/secret is location`,
				},
				resource.Attribute{
					Name:        "aws_secret_name",
					Description: `(Required) Secrets manager key name that has the private key ### GCP_SECRET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key that's stored on GCP secrets manager, provide the csp_account_name (GCP) that stores the secrets key`,
				},
				resource.Attribute{
					Name:        "secret_name",
					Description: `(Required) Secrets manager key name that has the private key`,
				},
				resource.Attribute{
					Name:        "secret_version",
					Description: `(Optional) Secrets manager key version ### AZURE_KEY_VAULT_SECRET`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) To use a private key that's stored on GCP secrets manager, provide the csp_account_name (GCP) that stores the secrets key`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the key/secret is location`,
				},
				resource.Attribute{
					Name:        "azure_key_vault_name",
					Description: `(Required) Azure key vault name`,
				},
				resource.Attribute{
					Name:        "azure_key_vault_secret_name",
					Description: `(Required) Azure secret name`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_cloud_account",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Cloud account resource defines the credentials of the cloud provider that can be accessed by the Valtix controller.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Cloud Account on the Valtix console. Must contain only alphanumeric, hyphens or underscore characters and not exceed 100 characters`,
				},
				resource.Attribute{
					Name:        "csp_type",
					Description: `(Required) Defines the Cloud Service Provider. Must be "GCP" or "AWS" or "AZURE" ### AWS Arguments`,
				},
				resource.Attribute{
					Name:        "aws_credentials_type",
					Description: `(AWS - Required) must be "AWS_IAM_ROLE"`,
				},
				resource.Attribute{
					Name:        "aws_iam_role",
					Description: `(AWS - Required) Cross IAM role ARN that Valtix assumes to manage your cloud account`,
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
					Name:        "inventory_monitoring",
					Description: `Enable inventory monitoring (can be repeated multiple times), look at [Inventory Monitoring](#inventory-monitoring) for details ### Azure Arguments`,
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
					Description: `Enable inventory monitoring (can be repeated multiple times), look at [Inventory Monitoring](#inventory-monitoring) for details ### GCP Arguments`,
				},
				resource.Attribute{
					Name:        "gcp_credentials_file",
					Description: `(GCP - Required) Service account credentials key file created for the Valtix controller access.`,
				},
				resource.Attribute{
					Name:        "inventory_monitoring",
					Description: `Enable inventory monitoring (can be repeated multiple times), look at [Inventory Monitoring](#inventory-monitoring) for details ## Inventory Monitoring`,
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
					Description: `External id that must be used in the trust settings of the cross account IAM role`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_id",
					Description: `External id that must be used in the trust settings of the cross account IAM role`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_gateway",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create a valtix gateway  

Please note that with the exclusion of description and gateway_image, any other arugment changes are prevented at this time as it would result in a service interruption to the Valtix gateway. In order to change these arguments, you must destroy and then create the valtix_gateway resource.

[valtix_cloud_account](../valtix_cloud_account), [valtix_policy_rule_set](../valtix_policy_rule_set)
must be defined before valtix_gateway can be created

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Valtix gateway.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Valtix gateway`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The CSP account where the gateway will be deployed.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Must be one of:`,
				},
				resource.Attribute{
					Name:        "gateway_image",
					Description: `(Required) Example "2.3-01". This is the Valtix image version to be deployed for this gateway. Please consult with Valtix support for recommended version.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Required) "EDGE" or "HUB". Use "EDGE" for Azure, GCP and for AWS Single VPC. Use "HUB" for AWS Transit Gateway deployment mode`,
				},
				resource.Attribute{
					Name:        "security_type",
					Description: `(Optional) "INGRESS" or "EGRESS". Default "INGRESS"`,
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
					Description: `(GCP - Required) This is the GCP gateway service account email, that provides permissions for the Valtix gateway to integrate with other GCP project resources such as Secrets Manager and storage buckets.`,
				},
				resource.Attribute{
					Name:        "aws_iam_role_firewall",
					Description: `(AWS - Required) This is the IAM role that's assigned to the Valtix firewall instances.`,
				},
				resource.Attribute{
					Name:        "azure_user_identity_id",
					Description: `(Azure - Optional) User assgined identity This is the IAM role that's assigned to the Valtix firewall instances.`,
				},
				resource.Attribute{
					Name:        "azure_resource_group",
					Description: `(Azure - Required) Azure resource group name used for all Valtix gateway resources`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Region where the Valtix gateway is deployed.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) VPC ID where the Valtix gateway is deployed and is used for data traffic to be inspected. This must be either the VPC where you apps run or the shared services VPC that's peered (or hub via transit gateway) to other spoke (app) VPCs. Please note that for HUB mode, this vpc_id must refer to a services VPC ID attribute that is exported using the [valtix_service_vpc](/terraform/valtix_service_vpc/#valtix_service_vpc) resource`,
				},
				resource.Attribute{
					Name:        "aws_gateway_lb",
					Description: `(Optional only for AWS HUB mode) "true" or "false". If attribute is "true", this will deploy AWS Gateway using AWS Gateway Load Balancer. This is only for EGRESS gateway that will support both East-West and Egress traffic. This is only available for HUB mode and used in regions that support AWS Gateway Load Balancer`,
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
					Description: `(Optional) minimum number of instances per zone (default 1)`,
				},
				resource.Attribute{
					Name:        "max_instances",
					Description: `(Optional) maximum number of instances per zone (default 1)`,
				},
				resource.Attribute{
					Name:        "health_check_port",
					Description: `(Optional) port number that NLB uses to monitor the instances (default 65534). A rule must be configured on the datapath_security_group to allow traffic to this port.`,
				},
				resource.Attribute{
					Name:        "log_profile",
					Description: `(Optional) log profile id`,
				},
				resource.Attribute{
					Name:        "packet_capture_profile",
					Description: `(Optional) packet profile id`,
				},
				resource.Attribute{
					Name:        "diagnostics_profile",
					Description: `(Optional) diagnostics profile id`,
				},
				resource.Attribute{
					Name:        "instance_details",
					Description: `(Required) contains the parameters used to deploy gateway in each availability zone. This block can be repeated multiple times for deploying gateway instances in multiple zones. Look below for the structure of this block. Atleast 1 block must be provided. ## Instance Details This section is not required for AWS HUB mode as instance details are configured from service VPC referenced in vpc_id attribute`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Required) specifies the availability zone where the Valtix gateway instance(s) are deployed`,
				},
				resource.Attribute{
					Name:        "mgmt_subnet",
					Description: `(Required) specifies the VPC subnet ID used for management traffic where the Valtix gateway instance(s) are deployed for this availability zone.`,
				},
				resource.Attribute{
					Name:        "datapath_subnet",
					Description: `(Required) specifies the VPC subnet ID used for data traffic where the Valtix gateway instance(s) are deployed for this availability zone.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rule_set",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

A policy rule set is a list of firewall rules. The rule set can be applied to multiple gateways to achieve identical security posture. It is recommended to create an empty policy rule set with this resource and manage the rules using the valtix_policy_rules resource.

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
					Description: `(Required) "ENABLED" or "DISABLED". Set the rule's state to enabled or disabled.`,
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
					Description: `(Required) "ALLOW_LOG", "ALLOW" (does not log the flow), "DENY_NOLOG" (does not log the flow), "DENY" (log the flow)`,
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
			Type:             "valtix_policy_rules",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Policy Rules is a list of firewall rules that are specified for a Policy Rule Set. 

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
					Description: `(Required) "ENABLED" or "DISABLED". Set the rule's state to enabled or disabled.`,
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
					Description: `(Required) "ALLOW_LOG", "ALLOW" (does not log the flow), "DENY_NOLOG" (does not log the flow), "DENY" (log the flow)`,
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

Create Anti Virus Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Anti Virus profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Anti Virus profile`,
				},
				resource.Attribute{
					Name:        "auto_update",
					Description: `(Optional) Auto update the Talos ClamAV Ruleset version daily with a delay specified by ` + "`" + `delay_by_days` + "`" + ` parameter. The valid values are true/false and it is true by default..`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) How many days before we use a Talos ClamAV Ruleset version after it has been published by Valtix. The default for this argument is 7 days, meaning that after the Jan 1st ruleset is published by Valtix, it is applied to the profile, and hence all the gateways using the profile, on Jan 8th. Valtix publishes new rulesets every day except when our internal testing fails.`,
				},
				resource.Attribute{
					Name:        "talos_ruleset_version",
					Description: `(Optional) Talos ClamAV Ruleset version. Find the values from the UI. The rulesets are published everyday. Unless you want to use a specific version, Valtix recommends to use auto_update as described above If this argument is specified, Auto Update of Talos ClamAV Ruleset is disabled and the profile will only use this version for Talos ClamAV Ruleset.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Default action for all the attacks. Valid values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_application_threat",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create Web Application Firewall (WAF) profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the profile`,
				},
				resource.Attribute{
					Name:        "auto_update_crs",
					Description: `(Optional) Auto update the CRS (Core Rule Set) version daily with a delay specified by ` + "`" + `delay_by_days_crs` + "`" + ` parameter. The valid values are true/false and it is true by default..`,
				},
				resource.Attribute{
					Name:        "delay_by_days_crs",
					Description: `(Optional) How many days before we use a CRS ruleset version after it has been published by Valtix. The default for this argument is 7 days, meaning that after the Jan 1st ruleset is published by Valtix, it is applied to the profile, and hence all the gateways using the profile, on Jan 8th. Valtix publishes new rulesets every day except when our internal testing fails.`,
				},
				resource.Attribute{
					Name:        "crs_ruleset_version",
					Description: `(Optional) CRS (Core Rule Set) version. Find the values from the UI. The rulesets are published everyday. Unless you want to use a specific version, Valtix recommends to use auto_update as described above If this argument is specified, Auto Update of CRS ruleset is disabled and the profile will only use this version for CRS ruleset.`,
				},
				resource.Attribute{
					Name:        "auto_update_trustwave",
					Description: `(Optional) Auto update the Trustwave Rule Set version daily with a delay specified by ` + "`" + `delay_by_days_trustwave` + "`" + ` parameter. The valid values are true/false and it is true by default..`,
				},
				resource.Attribute{
					Name:        "delay_by_days_trustwave",
					Description: `(Optional) How many days before we use a Trustwave Rule Set version after it has been published by Valtix. The default for this argument is 7 days, meaning that after the Jan 1st ruleset is published by Valtix, it is applied to the profile, and hence all the gateways using the profile, on Jan 8th. Valtix publishes new rulesets every day except when our internal testing fails.`,
				},
				resource.Attribute{
					Name:        "trustwave_ruleset_version",
					Description: `(Optional) Trustwave Rule Set version. Valid values are (as of this publication of this document):`,
				},
				resource.Attribute{
					Name:        "paranoia_level",
					Description: `(Required) An integer between 1 and 4. Higher number leads to more false positives but also helps in detecting more attacks. Recommended value is 1`,
				},
				resource.Attribute{
					Name:        "request_anamoly",
					Description: `(Optional) Request anomaly score used in the mod_security anomaly scoring system. Default value is 3`,
				},
				resource.Attribute{
					Name:        "response_anamoly",
					Description: `(Optional) Response anomaly score used in the mod_security anomaly scoring system. Default value is 3`,
				},
				resource.Attribute{
					Name:        "pcap",
					Description: `(Optional) true/false. Capture pcap when attacks are detected`,
				},
				resource.Attribute{
					Name:        "api_logging",
					Description: `(Optional) true/false. Log API requests`,
				},
				resource.Attribute{
					Name:        "request_inspection_profile",
					Description: `(Optional) List of Request attacks from the Core Rule Set`,
				},
				resource.Attribute{
					Name:        "response_inspection_profile",
					Description: `(Optional) List of Response attacks from the Core Rule Set`,
				},
				resource.Attribute{
					Name:        "advanced_application_profile",
					Description: `(Optional) Advanced attacks from the trustwave rule set`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Default "BLOCK". Can be "DETECT"`,
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
					Description: `(Optional) List of rule ids to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) "RATE" or "SAMPLE". When "RATE" is selected, number_of_events and time must be provided. action is applied once the provided rule_ids match the given count in the given time. If the type is "SAMPLE", the action is applied once the count of the events matces`,
				},
				resource.Attribute{
					Name:        "number_of_events",
					Description: `(Optional) Number of times the rule id attack must match before the action is applied`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Used when the type is "RATE", the number_of_events must match in the given time (in seconds) ## Event Suppressor`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) List of source ips or CIDRs`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of rule ids to filter ## Profile Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of rule ids to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) "RATE" or "SAMPLE". When "RATE" is selected, number_of_events and time must be provided. action is applied once the provided rule_ids match the given count in the given time. If the type is "SAMPLE", the action is applied once the count of the events matces`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_decryption",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create a decryption profile that can be used in a service.

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
					Description: `(Required) Certificate name from the valtix_certificate object`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_dlp",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Data loss prevention profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the profile`,
				},
				resource.Attribute{
					Name:        "dlp_filter_list",
					Description: `(Required) A list of dlp filters. Structure is [defined below](#dlp-filter) ## DLP Filter`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Descrption of the filter`,
				},
				resource.Attribute{
					Name:        "patterns",
					Description: `(Optional) List of custom regular expression patterns`,
				},
				resource.Attribute{
					Name:        "static_patterns",
					Description: `(Optional) List of predefined patterns`,
				},
				resource.Attribute{
					Name:        "count",
					Description: `(Required) Number of times the pattern must be seen before the data loss prevention functionality takes action`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required) action upon detecting the data loss prevention. Valid values:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_fqdn",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create FQDN filtering profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the profile`,
				},
				resource.Attribute{
					Name:        "fqdn_filter_list",
					Description: `(Required) List of fqdn_filter resources. Structure [defined below](#fqdn-filter)`,
				},
				resource.Attribute{
					Name:        "default_fqdn_filter",
					Description: `(Optional) Default FQDN filter behavious. Structure [defined below](#fqdn-filter) ## FQDN Filter`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `(Required) String or regular expression or a predefined Category`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Action to take on the matching url (and method) "ALLOW_LOG", "ALLOW" (does not log the flow), "DENY_NOLOG" (does not log the flow), "DENY" (log the flow).`,
				},
				resource.Attribute{
					Name:        "decryption_exception",
					Description: `(Optional) true/false. In the proxy mode disable decryption and inspection of packets`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) HTTP actions/methods to which the rate limiting is applied. Valid values are GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS`,
				},
			},
			Attributes: []resource.Attribute{},
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
					Description: `(Required) One of`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required ) https URL.`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required ) https auth token`,
				},
				resource.Attribute{
					Name:        "splunk_index",
					Description: `(Required) splunk index name to store the events ### DATADOG`,
				},
				resource.Attribute{
					Name:        "endpoint",
					Description: `(Required ) https URL.`,
				},
				resource.Attribute{
					Name:        "auth_token",
					Description: `(Required) https auth token ### GCPLOGGING_FROM_GATEWAY`,
				},
				resource.Attribute{
					Name:        "log_name",
					Description: `[Optional] gcp log name to store the logs. If not specified, the default name is "valtix-gateway-logs" ### SYSLOG`,
				},
				resource.Attribute{
					Name:        "syslog_server_ip",
					Description: `(Required) syslog server ip`,
				},
				resource.Attribute{
					Name:        "syslog_server_port",
					Description: `(Required) syslog server port`,
				},
				resource.Attribute{
					Name:        "network_threat_severity",
					Description: `(Required) syslog server port`,
				},
				resource.Attribute{
					Name:        "syslog_flow_logs",
					Description: `(Optional) true/false. forward flow logs to syslog`,
				},
				resource.Attribute{
					Name:        "syslog_firewall_events",
					Description: `(Optional) true/false. forward firewall events to syslog`,
				},
				resource.Attribute{
					Name:        "syslog_https_logs",
					Description: `(Optional) true/false. forward tls logs to syslog`,
				},
				resource.Attribute{
					Name:        "network_threat_severity",
					Description: `(Optional) One of "Alert", "Emergency", "Critical", "Error", "Warning", "Notice", "Info", "Debug". Forward network threat events (ips) with the given severity only`,
				},
				resource.Attribute{
					Name:        "web_attack_severity",
					Description: `(Optional) One of "Alert", "Emergency", "Critical", "Error", "Warning", "Notice", "Info", "Debug". Forward web attacks with the given severity only`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_malicious_source",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create Malicious Sources Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Malicious Sources profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the Malicious Sources profile`,
				},
				resource.Attribute{
					Name:        "ip_reputation_enabled",
					Description: `(Optional) true/false`,
				},
				resource.Attribute{
					Name:        "auto_update",
					Description: `(Optional) Auto update the Trustwave IP Reputation Ruleset version daily with a delay specified by ` + "`" + `delay_by_days` + "`" + ` parameter. The valid values are true/false and it is true by default.`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) How many days before we use a Trustwave IP Reputation Ruleset version after it has been published by Valtix. The default for this argument is 7 days, meaning that after the Jan 1st ruleset is published by Valtix, it is applied to the profile, and hence all the gateways using the profile, on Jan 8th. Valtix publishes new rulesets every day except when our internal testing fails.`,
				},
				resource.Attribute{
					Name:        "vendor_rule_set_name",
					Description: `(Optional) Vendor rule set name/version. Find the values from the UI. The rulesets are published everyday. Unless you want to use a specific version, Valtix recommends to use auto_update as described above If this argument is specified, Auto Update of Trustwave IP Reputation Ruleset is disabled and the profile will only use this version for Trustwave IP Reputation Ruleset.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_network_intrusion",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create IPS (Network Intrusion) Profile

`,
			Keywords: []string{},
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
					Name:        "auto_update",
					Description: `(Optional) Auto update the Talos IPS Ruleset version daily with a delay specified by ` + "`" + `delay_by_days` + "`" + ` parameter. The valid values are true/false and it is true by default..`,
				},
				resource.Attribute{
					Name:        "delay_by_days",
					Description: `(Optional) How many days before we use a Talos IPS Ruleset version after it has been published by Valtix. The default for this argument is 7 days, meaning that after the Jan 1st ruleset is published by Valtix, it is applied to the profile, and hence all the gateways using the profile, on Jan 8th. Valtix publishes new rulesets every day except when our internal testing fails.`,
				},
				resource.Attribute{
					Name:        "talos_ruleset_version",
					Description: `(Optional) Talos IPS ruleset version. Find the values from the UI. The rulesets are published everyday. Unless you want to use a specific version, Valtix recommends to use auto_update as described above If this argument is specified, Auto Update of Talos IPS Ruleset is disabled and the profile will only use this version for Talos IPS Ruleset.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Default action for all the attacks. Can be overwritten for each attack type. Default value is specified in the talos rule set. If this is not specified then whatever action that's specified in the ruleset is used (called Rule Default). Valid values:`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Pre-defined IPS ruleset to use, must be one of:`,
				},
				resource.Attribute{
					Name:        "policy_action",
					Description: `(Optional) Action to apply for the predefined policy. Look at action attribute for valid values. If this is not specified, the action defined in the 'action' attribute is used.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Optional) List of predefined categories. Structure [defined below](#categories)`,
				},
				resource.Attribute{
					Name:        "classes",
					Description: `(Optional) List of predefined classes. Structure [defined below](#classes)`,
				},
				resource.Attribute{
					Name:        "pcap",
					Description: `(Optional) true/false. Capture pcap when traffic matches the IPS rules`,
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
					Description: `(Optional) Similar to the rule_event_filter but applies to the whole profile instead of specific rule(s). Structure is [defined below](#profile-event-filter) ## Categories`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of IPS attacks categories`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Values same as action attribute. If this is not specified, then the value defined in 'action' attribute is used ## Classes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of IPS attacks classes`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Values same as action attribute. If this is not specified, then the value defined in 'action' attribute is used ## Rule Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of rule ids to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) "RATE" or "SAMPLE". When "RATE" is selected, number_of_events and time must be provided. action is applied once the provided rule_ids match the given count in the given time. If the type is "SAMPLE", the action is applied once the count of the events matces`,
				},
				resource.Attribute{
					Name:        "number_of_events",
					Description: `(Optional) Number of times the rule id attack must match before the action is applied`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `(Optional) Used when the type is "RATE", the number_of_events must match in the given time (in seconds) ## Event Suppressor`,
				},
				resource.Attribute{
					Name:        "source_ips",
					Description: `(Optional) List of source ips or CIDRs`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of rule ids to filter ## Profile Event Filter`,
				},
				resource.Attribute{
					Name:        "rule_ids",
					Description: `(Optional) List of rule ids to filter`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) "RATE" or "SAMPLE". When "RATE" is selected, number_of_events and time must be provided. action is applied once the provided rule_ids match the given count in the given time. If the type is "SAMPLE", the action is applied once the count of the events matces`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_packet_capture",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Storage location to store the pcap files

`,
			Keywords: []string{},
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
					Name:        "csp_account",
					Description: `(Required) Cloud account name added to the Controller (valtix_cloud_account.name)`,
				},
				resource.Attribute{
					Name:        "storage_upload_path",
					Description: `(Required AWS and GCP) Storage bucket name`,
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
					Description: `(Optional - Azure) Storage account access key, if required`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_profile_urlfilter",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Create URL Filtering Profile

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the profile`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the profile`,
				},
				resource.Attribute{
					Name:        "url_filter_list",
					Description: `(Required) List of url_filter resources. Structure [defined below](#url-filter)`,
				},
				resource.Attribute{
					Name:        "default_url_filter",
					Description: `(Optional) Default behavior of URL filter. Structure [defined below](#url-filter) ## URL Filter`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) String or regular expression or a predefined Category`,
				},
				resource.Attribute{
					Name:        "filter_methods",
					Description: `(Optional) URL Methods (GET, POST etc). Default all the methods`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) Action to take on the matching url (and method) "ALLOW_LOG", "ALLOW" (does not log the flow), "DENY_NOLOG" (does not log the flow), "DENY" (log the flow).`,
				},
				resource.Attribute{
					Name:        "return_status",
					Description: `(Optional) HTTP status code to return for DENY and DENY_NOLOG policy`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_object",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the service object`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) "ReverseProxy", "ForwardProxy" or "Forwarding"`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) "TCP" or "UDP". "TCP" is default. This is the listener protocol.`,
				},
				resource.Attribute{
					Name:        "transport_mode",
					Description: `(Required) "HTTP", "HTTPS", "TCP", "TLS". The protocol used by the gateway to communicate with the backend.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) This can be specified multiple times if the service runs on multiple ports. Structure is [documented below](#reverseproxy-port)`,
				},
				resource.Attribute{
					Name:        "sni",
					Description: `(Optional) List of FQDN strings that's matched by GW to find the destination (target) address group. Used to distinguish multiple TLS applications on a single port`,
				},
				resource.Attribute{
					Name:        "tls_profile",
					Description: `(Optional) Decryption profile id`,
				},
				resource.Attribute{
					Name:        "l7dos_profile",
					Description: `(Optional) L7 DOS profile id ## ReverseProxy Port port can be specified multiple times to define a list of ports that the service can listen. ` + "`" + `` + "`" + `` + "`" + `hcl port { destination_ports = "443" backend_ports = "443" } # if the listen ports are continuous and target ports are continuous port { destination_ports = "80-100" backend_ports = "8000-8020" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Required) Destination port number as a string or a continuous range of destination port numbers (e.g "80" or "80-100")`,
				},
				resource.Attribute{
					Name:        "backend_ports",
					Description: `(Required) Backend (target) port number as a string or a continuous range of target port numbers (e.g "80" or "80-100"). The range count in destination and backend must match. ## ForwardProxy`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the service object`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) "ForwardProxy"`,
				},
				resource.Attribute{
					Name:        "transport_mode",
					Description: `(Required) "HTTP", "HTTPS"`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) This can be specified multiple times if the service can run on multiple ports. Structure is [documented below](#forwardproxy-port)`,
				},
				resource.Attribute{
					Name:        "tls_profile",
					Description: `(Optional) Decryption profile id. ## Forwarding`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service object`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the service object`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `(Required) "Forwarding"`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) "TCP" or "UDP". "TCP" is default.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) This can be specified multiple times if the service can run on multiple ports. Structure is [documented below](#forwardproxy-port)`,
				},
				resource.Attribute{
					Name:        "source_nat",
					Description: `(Optional) true or false. Specifies whether source NAT (Network Address Translation) would be performed on the packet flow ## ForwardProxy Port port can be specified multiple times to define a list of ports that the service can listen. ` + "`" + `` + "`" + `` + "`" + `hcl port { destination_ports = "443" } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "destination_ports",
					Description: `(Required) Destination port number as a string or a continuous range of destination port numbers (e.g "80" or "80-100")`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_vpc",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service VPC`,
				},
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) The CSP account where the service VPC will be deployed`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The AWS region where the service VPC will be deployed`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Required) The CIDR of the service VPC to be deployed`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `(Required) The list of availability zones that the service VPC will use. HUB mode gateways deployed in this service VPC will use all availability zones defined here`,
				},
				resource.Attribute{
					Name:        "transit_gateway_id",
					Description: `(Required) Transit Gateway ID for the service VPC to attach to ## Attribute Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) VPC ID of the Services VPC that is created`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) VPC ID of the Services VPC that is created`,
				},
			},
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
		"valtix_profile_dlp":                12,
		"valtix_profile_fqdn":               13,
		"valtix_profile_l7dos":              14,
		"valtix_profile_log_forwarding":     15,
		"valtix_profile_malicious_source":   16,
		"valtix_profile_network_intrusion":  17,
		"valtix_profile_packet_capture":     18,
		"valtix_profile_urlfilter":          19,
		"valtix_service_object":             20,
		"valtix_service_vpc":                21,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
