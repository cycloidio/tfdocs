package dome9

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dome9_cloudaccount_aws",
			Category:         "Data Sources",
			ShortDescription: `Get information about AWS cloud account onboarded to Dome9.`,
			Description: `

Use this data source to get information about an  AWS cloud account onboarded to Dome9.

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
					Name:        "account_id",
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
					Name:        "account_id",
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
	}

	dataSourcesMap = map[string]int{

		"dome9_cloudaccount_aws":                   0,
		"dome9_cloudaccount_azure":                 1,
		"dome9_cloudaccount_gcp":                   2,
		"dome9_continuous_compliance_notification": 3,
		"dome9_continuous_compliance_policy":       4,
		"dome9_iplist":                             5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
