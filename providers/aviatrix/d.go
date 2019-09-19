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
					Description: `(Required) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference`,
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
					Description: `(Computed) Aviatrix caller identity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cid",
					Description: `(Computed) Aviatrix caller identity.`,
				},
			},
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
					Description: `(Optional) Account name. This can be used for logging in to CloudN console or UserConnect controller. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix gateway name. ## The following arguments are computed - please do not edit in the resource file:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider. (Only AWS is supported currently. Value of 1 for AWS.)`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `AWS VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `AWS VPC Region.`,
				},
				resource.Attribute{
					Name:        "vpc_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "account_name",
					Description: `Aviatrix account name.`,
				},
				resource.Attribute{
					Name:        "gw_name",
					Description: `Aviatrix gateway name. ## The following arguments are computed - please do not edit in the resource file:`,
				},
				resource.Attribute{
					Name:        "cloud_type",
					Description: `Type of cloud service provider. (Only AWS is supported currently. Value of 1 for AWS.)`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `AWS VPC ID.`,
				},
				resource.Attribute{
					Name:        "vpc_reg",
					Description: `AWS VPC Region.`,
				},
				resource.Attribute{
					Name:        "vpc_size",
					Description: `Instance type.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the Gateway created`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"aviatrix_account":         0,
		"aviatrix_caller_identity": 1,
		"aviatrix_gateway":         2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
