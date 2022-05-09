package clumio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_aws_connection",
			Category:         "Resources",
			ShortDescription: `Clumio AWS Connection Resource used to connect AWS accounts to Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_callback_resource",
			Category:         "Resources",
			ShortDescription: `Clumio Callback Resource used while on-boarding AWS clients. The purpose of this resource is to send a SNS event with the necessary details of the AWS connection configuration done on the client AWS account so that necessary connection post processing can be done in Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_organizational_unit",
			Category:         "Resources",
			ShortDescription: `Clumio AWS Connection Resource used to connect AWS accounts to Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_policy",
			Category:         "Resources",
			ShortDescription: `Clumio Policy Resource used to schedule backups on Clumio supported data sources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_policy_assignment",
			Category:         "Resources",
			ShortDescription: `Clumio Policy Assignment Resource used to assign (or unassign) policies. NOTE: Currently policy assignment is supported only for entity type "protection_group".`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_policy_rule",
			Category:         "Resources",
			ShortDescription: `Clumio Policy Rule Resource used to determine how a policy should be assigned to assets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_post_process_aws_connection",
			Category:         "Resources",
			ShortDescription: `Post-Process Clumio AWS Connection Resource used to post-process AWS connection to Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_protection_group",
			Category:         "Resources",
			ShortDescription: `Clumio S3 Protection Group Resource used to create and manage Protection Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "clumio_clumio_user",
			Category:         "Resources",
			ShortDescription: `Clumio User Resource to create and manage users in Clumio.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"clumio_clumio_aws_connection":              0,
		"clumio_clumio_callback_resource":           1,
		"clumio_clumio_organizational_unit":         2,
		"clumio_clumio_policy":                      3,
		"clumio_clumio_policy_assignment":           4,
		"clumio_clumio_policy_rule":                 5,
		"clumio_clumio_post_process_aws_connection": 6,
		"clumio_clumio_protection_group":            7,
		"clumio_clumio_user":                        8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
