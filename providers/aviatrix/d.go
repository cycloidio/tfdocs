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
					Description: `(Required) Select PAN. Valid values: "Generic", "Palo Alto VM Series", "Palo Alto VM Panorama", "Aviatrix FQDN Gateway".`,
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
	}

	dataSourcesMap = map[string]int{

		"aviatrix_account":                    0,
		"aviatrix_caller_identity":            1,
		"aviatrix_firenet_vendor_integration": 2,
		"aviatrix_gateway":                    3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
