package clumio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "clumio_auto_user_provisioning_rule",
			Category:         "Resources",
			ShortDescription: `Clumio Auto User Provisioning Rule Resource used to determine the Role and Organizational Units to be assigned to a user based on their groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_auto_user_provisioning_setting",
			Category:         "Resources",
			ShortDescription: `Clumio Auto User Provisioning Setting Resource used to determine if Auto User Provisioning feature is enabled or not.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_aws_connection",
			Category:         "Resources",
			ShortDescription: `Clumio AWS Connection Resource used to connect AWS accounts to Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_callback_resource",
			Category:         "Resources",
			ShortDescription: `Clumio Callback Resource used while on-boarding AWS clients. The purpose of this resource is to send a SNS event with the necessary details of the AWS connection configuration done on the client AWS account so that necessary connection post processing can be done in Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_organizational_unit",
			Category:         "Resources",
			ShortDescription: `Resource for creating and managing Organizational Unit in Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_policy",
			Category:         "Resources",
			ShortDescription: `Clumio Policy Resource used to schedule backups on Clumio supported data sources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_policy_assignment",
			Category:         "Resources",
			ShortDescription: `Clumio Policy Assignment Resource used to assign (or unassign) policies. NOTE: Currently policy assignment is supported only for entity type "protection_group".`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_policy_rule",
			Category:         "Resources",
			ShortDescription: `Clumio Policy Rule Resource used to determine how a policy should be assigned to assets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_post_process_aws_connection",
			Category:         "Resources",
			ShortDescription: `Post-Process Clumio AWS Connection Resource used to post-process AWS connection to Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_post_process_kms",
			Category:         "Resources",
			ShortDescription: `Post-Process Clumio KMS Resource used to post-process KMS in Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_protection_group",
			Category:         "Resources",
			ShortDescription: `Clumio S3 Protection Group Resource used to create and manage Protection Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_user",
			Category:         "Resources",
			ShortDescription: `Clumio User Resource to create and manage users in Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_wallet",
			Category:         "Resources",
			ShortDescription: `Clumio Wallet Resource to create and manage wallets in Clumio. Wallets should be created "after" connecting an AWS account to Clumio.NOTE: To protect against accidental deletion, wallets cannot be destroyed. To remove a wallet, contact Clumio support.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"clumio_auto_user_provisioning_rule":    0,
		"clumio_auto_user_provisioning_setting": 1,
		"clumio_aws_connection":                 2,
		"clumio_callback_resource":              3,
		"clumio_organizational_unit":            4,
		"clumio_policy":                         5,
		"clumio_policy_assignment":              6,
		"clumio_policy_rule":                    7,
		"clumio_post_process_aws_connection":    8,
		"clumio_post_process_kms":               9,
		"clumio_protection_group":               10,
		"clumio_user":                           11,
		"clumio_wallet":                         12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
