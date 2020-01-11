package dome9

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dome9_aws_security_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about AWS Security Groups in Dome9`,
			Description: `

Use this data source to get information about an AWS Security Group onboarded to Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dome9_security_group_name",
					Description: `Name of the Security Group.`,
				},
				resource.Attribute{
					Name:        "dome9_cloud_account_id",
					Description: `Cloud account id in Dome9.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Security Group description.`,
				},
				resource.Attribute{
					Name:        "aws_region_id",
					Description: `AWS region; in AWS format (e.g., "us-east-1").`,
				},
				resource.Attribute{
					Name:        "is_protected",
					Description: `indicates whether the Security Group is protected.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Security Group id.`,
				},
				resource.Attribute{
					Name:        "vpc_name",
					Description: `name of VPC containing the Security Group.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Security Group tags.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `Security Group services.`,
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
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_azure_security_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about Azure Security Group in Dome9`,
			Description: `

Use this data source to get information about an Azure Security Group onboarded to Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dome9_security_group_name",
					Description: `(Required) Name of the Security Group.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Security Group region.`,
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
					Description: `(Optional) Security Group description.`,
				},
				resource.Attribute{
					Name:        "is_tamper_protected",
					Description: `(Optional) Is Security Group tamper protected.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Security Group tags.`,
				},
				resource.Attribute{
					Name:        "inbound",
					Description: `(Optional) Security Group services.`,
				},
				resource.Attribute{
					Name:        "outbound",
					Description: `(Optional) Security Group services.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_aws",
			Category:         "Data Sources",
			ShortDescription: `Get information about AWS cloud account onboarded to Dome9.`,
			Description: `

Use this data source to get information about an AWS cloud account onboarded to Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The Dome9 id for the AWS account ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("AWS").`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The cloud account name in Dome9.`,
				},
				resource.Attribute{
					Name:        "external_account_number",
					Description: `The AWS account number.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Credentials error status.`,
				},
				resource.Attribute{
					Name:        "is_fetching_suspended",
					Description: `Fetching suspending status.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date account was onboarded to Dome9.`,
				},
				resource.Attribute{
					Name:        "full_protection",
					Description: `The tamper Protection mode for current security groups.`,
				},
				resource.Attribute{
					Name:        "allow_read_only",
					Description: `The AWS cloud account operation mode. true for "Manage", false for "Readonly".`,
				},
				resource.Attribute{
					Name:        "net_sec",
					Description: `The network security configuration for the AWS cloud account.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id.`,
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
					Description: `Restricted IAM safe entities which has the following:`,
				},
				resource.Attribute{
					Name:        "roles_ARNs",
					Description: `Restricted IAM safe entities roles ARNs`,
				},
				resource.Attribute{
					Name:        "users_ARNs",
					Description: `Restricted IAM safe entities users ARNs`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider ("AWS").`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The cloud account name in Dome9.`,
				},
				resource.Attribute{
					Name:        "external_account_number",
					Description: `The AWS account number.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Credentials error status.`,
				},
				resource.Attribute{
					Name:        "is_fetching_suspended",
					Description: `Fetching suspending status.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date account was onboarded to Dome9.`,
				},
				resource.Attribute{
					Name:        "full_protection",
					Description: `The tamper Protection mode for current security groups.`,
				},
				resource.Attribute{
					Name:        "allow_read_only",
					Description: `The AWS cloud account operation mode. true for "Manage", false for "Readonly".`,
				},
				resource.Attribute{
					Name:        "net_sec",
					Description: `The network security configuration for the AWS cloud account.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id.`,
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
					Description: `Restricted IAM safe entities which has the following:`,
				},
				resource.Attribute{
					Name:        "roles_ARNs",
					Description: `Restricted IAM safe entities roles ARNs`,
				},
				resource.Attribute{
					Name:        "users_ARNs",
					Description: `Restricted IAM safe entities users ARNs`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_azure",
			Category:         "Data Sources",
			ShortDescription: `Get information about Azure cloud account onboarded to Dome9.`,
			Description: `

Use this data source to get information about an Azure cloud account onboarded to Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The Dome9 id for the Azure account. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Account name (in Dome9).`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `Azure subscription id for account.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Azure tenant id.`,
				},
				resource.Attribute{
					Name:        "operation_mode",
					Description: `Dome9 operation mode for the Azure account (Read-Only or Managed).`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (Azure).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date Azure account was onboarded to a Dome9 account.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Account name (in Dome9).`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `Azure subscription id for account.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `Azure tenant id.`,
				},
				resource.Attribute{
					Name:        "operation_mode",
					Description: `Dome9 operation mode for the Azure account (Read-Only or Managed).`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (Azure).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Date Azure account was onboarded to a Dome9 account.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_gcp",
			Category:         "Data Sources",
			ShortDescription: `Get information about a GCP cloud account onboarded to Dome9.`,
			Description: `

Use this data source to get information about a GCP cloud account onboarded to Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The Dome9 id for the GCP account. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `GCP account name in Dome9.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `the Google project id (that will be onboarded).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name.`,
				},
				resource.Attribute{
					Name:        "gsuite_user",
					Description: `Gsuite user.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Azure tenant id.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (gcp).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `GCP account name in Dome9.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `the Google project id (that will be onboarded).`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_id",
					Description: `Organizational unit id.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_path",
					Description: `Organizational unit path.`,
				},
				resource.Attribute{
					Name:        "organizational_unit_name",
					Description: `Organizational unit name.`,
				},
				resource.Attribute{
					Name:        "gsuite_user",
					Description: `Gsuite user.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Domain name.`,
				},
				resource.Attribute{
					Name:        "domain_name",
					Description: `Azure tenant id.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `The cloud provider (gcp).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_continuous_compliance_notification",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Dome9 continuous compliance notification policy.`,
			Description: `

Use this data source to get information about a Dome9 continuous compliance notification policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id for the continuous compliance notification policy in Dome9. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Notification policy name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the notification.`,
				},
				resource.Attribute{
					Name:        "alerts_console",
					Description: `Include in the alerts console.`,
				},
				resource.Attribute{
					Name:        "scheduled_report",
					Description: `Scheduled report details.`,
				},
				resource.Attribute{
					Name:        "change_detection",
					Description: `Change detection options.`,
				},
				resource.Attribute{
					Name:        "gcp_security_command_center_integration",
					Description: `GCP security command center details`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Notification policy name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the notification.`,
				},
				resource.Attribute{
					Name:        "alerts_console",
					Description: `Include in the alerts console.`,
				},
				resource.Attribute{
					Name:        "scheduled_report",
					Description: `Scheduled report details.`,
				},
				resource.Attribute{
					Name:        "change_detection",
					Description: `Change detection options.`,
				},
				resource.Attribute{
					Name:        "gcp_security_command_center_integration",
					Description: `GCP security command center details`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_continuous_compliance_policy",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Dome9 continuous compliance policy.`,
			Description: `

Use this data source to get information about a Dome9 continuous compliance policy.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id for the cloud account in Dome9. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `GCP account name in Dome9.`,
				},
				resource.Attribute{
					Name:        "external_account_id",
					Description: `The account number.`,
				},
				resource.Attribute{
					Name:        "cloud_account_type",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "bundle_id",
					Description: `Organizational unit id.`,
				},
				resource.Attribute{
					Name:        "notification_ids",
					Description: `Organizational unit path.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `GCP account name in Dome9.`,
				},
				resource.Attribute{
					Name:        "external_account_id",
					Description: `The account number.`,
				},
				resource.Attribute{
					Name:        "cloud_account_type",
					Description: `creation date for project in Google.`,
				},
				resource.Attribute{
					Name:        "bundle_id",
					Description: `Organizational unit id.`,
				},
				resource.Attribute{
					Name:        "notification_ids",
					Description: `Organizational unit path.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_iplist",
			Category:         "Data Sources",
			ShortDescription: `Get information about an IP list in Dome9.`,
			Description: `

Use this data source to get information about an IP list in Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the IP list in Dome9. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `IP list name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `IP list description.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `Items (IP addresses) in the IP list.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `IP list name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `IP list description.`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `Items (IP addresses) in the IP list.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_organizational_unit",
			Category:         "Data Sources",
			ShortDescription: `Get information about an Organizational Unit in Dome9.`,
			Description: `

Use this data source to get information about a Organizational Unit in Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the organizational unit in Dome9. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organizational Unit in Dome9.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `The Organizational Unit parent ID.`,
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
					Description: `Is the parent of Organizational Unit root.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organizational Unit in Dome9.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `The Organizational Unit parent ID.`,
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
					Description: `Is the parent of Organizational Unit root.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_role",
			Category:         "Data Sources",
			ShortDescription: `Get information about a role in Dome9.`,
			Description: `

Use this data source to get information about a role in Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the role list in Dome9. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
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
			Attributes: []resource.Attribute{
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
		},
		&resource.Resource{
			Name:             "",
			Type:             "dome9_ruleset",
			Category:         "Data Sources",
			ShortDescription: `Get information about a ruleset in Dome9.`,
			Description: `

Use this data source to get information about a ruleset in Dome9.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the ruleset in Dome9. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The ruleset name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The ruleset description.`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `Cloud vendor that the ruleset is associated with.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Rule set creation time.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Rule set last update time.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of rules in the ruleset.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The ruleset name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The ruleset description.`,
				},
				resource.Attribute{
					Name:        "cloud_vendor",
					Description: `Cloud vendor that the ruleset is associated with.`,
				},
				resource.Attribute{
					Name:        "created_time",
					Description: `Rule set creation time.`,
				},
				resource.Attribute{
					Name:        "updated_time",
					Description: `Rule set last update time.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of rules in the ruleset.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"dome9_aws_security_group":                 0,
		"dome9_azure_security_group":               1,
		"dome9_cloudaccount_aws":                   2,
		"dome9_cloudaccount_azure":                 3,
		"dome9_cloudaccount_gcp":                   4,
		"dome9_continuous_compliance_notification": 5,
		"dome9_continuous_compliance_policy":       6,
		"dome9_iplist":                             7,
		"dome9_organizational_unit":                8,
		"dome9_role":                               9,
		"dome9_ruleset":                            10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
