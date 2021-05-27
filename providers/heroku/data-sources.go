package heroku

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "heroku_addon",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Heroku Addon.`,
			Description: `

Use this data source to get information about a Heroku Addon.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The add-on name ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The add-on name`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The plan name`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `The ID of the plan provider`,
				},
				resource.Attribute{
					Name:        "config_vars",
					Description: `The Configuration variables of the add-on`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the add-on`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The add-on name`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The plan name`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `The ID of the plan provider`,
				},
				resource.Attribute{
					Name:        "config_vars",
					Description: `The Configuration variables of the add-on`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_app",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Heroku App.`,
			Description: `

Use this data source to get information about a Heroku App.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application. In Heroku, this is also the unique ID, so it must be unique and have a minimum of 3 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique UUID of the Heroku app.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the application. In Heroku, this is also the unique ID.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `The application stack is what platform to run the application in.`,
				},
				resource.Attribute{
					Name:        "buildpacks",
					Description: `A list of buildpacks that this app uses.`,
				},
				resource.Attribute{
					Name:        "space",
					Description: `The private space in which the app runs. Not present if this is a common runtime app.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which the app is deployed.`,
				},
				resource.Attribute{
					Name:        "git_url",
					Description: `The Git URL for the application. This is used for deploying new versions of the app.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `The web (HTTP) URL that the application can be accessed at by default.`,
				},
				resource.Attribute{
					Name:        "heroku_hostname",
					Description: `The hostname for the Heroku application, suitable for pointing DNS records.`,
				},
				resource.Attribute{
					Name:        "config_vars",
					Description: `A map of all configuration variables for the app.`,
				},
				resource.Attribute{
					Name:        "acm",
					Description: `True if Heroku ACM is enabled for this app, false otherwise.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The Heroku Team that owns this app.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Heroku Team (organization).`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if the app access is locked`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique UUID of the Heroku app.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique UUID of the Heroku app.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the application. In Heroku, this is also the unique ID.`,
				},
				resource.Attribute{
					Name:        "stack",
					Description: `The application stack is what platform to run the application in.`,
				},
				resource.Attribute{
					Name:        "buildpacks",
					Description: `A list of buildpacks that this app uses.`,
				},
				resource.Attribute{
					Name:        "space",
					Description: `The private space in which the app runs. Not present if this is a common runtime app.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which the app is deployed.`,
				},
				resource.Attribute{
					Name:        "git_url",
					Description: `The Git URL for the application. This is used for deploying new versions of the app.`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `The web (HTTP) URL that the application can be accessed at by default.`,
				},
				resource.Attribute{
					Name:        "heroku_hostname",
					Description: `The hostname for the Heroku application, suitable for pointing DNS records.`,
				},
				resource.Attribute{
					Name:        "config_vars",
					Description: `A map of all configuration variables for the app.`,
				},
				resource.Attribute{
					Name:        "acm",
					Description: `True if Heroku ACM is enabled for this app, false otherwise.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The Heroku Team that owns this app.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Heroku Team (organization).`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `True if the app access is locked`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The unique UUID of the Heroku app.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_pipeline",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Heroku Pipeline.`,
			Description: `

Use this data source to get information about a Heroku Pipeline.

~> **NOTE:**
This data source can only be used to fetch information regarding a pipeline that has apps already associated to it.
This is a limitation in the Heroku Platform API where it is not possible to query a pipeline without apps by its name.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The pipeline name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the pipeline`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The pipeline owner's ID`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `The pipeline owner's type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the pipeline`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The pipeline owner's ID`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `The pipeline owner's type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Heroku Private Space.`,
			Description: `

Use this data source to get information about a [Heroku Private Space](https://www.heroku.com/private-spaces).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Heroku Private Space. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Heroku Private Space. In Heroku, this is also the unique .`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the Heroku Private Space.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which the Heroku Private Space is deployed.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Heroku Private Space. Either ` + "`" + `allocating` + "`" + ` or ` + "`" + `allocated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shield",
					Description: `Whether or not the space has [Shield](https://devcenter.heroku.com/articles/private-spaces#shield-private-spaces) turned on. One of ` + "`" + `on` + "`" + ` or ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The Heroku Team that owns this space. The fields for this block are documented below.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The RFC-1918 CIDR the Private Space will use. It must be a /16 in 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16`,
				},
				resource.Attribute{
					Name:        "data_cidr",
					Description: `The RFC-1918 CIDR that the Private Space will use for the Heroku-managed peering connection that’s automatically created when using Heroku Data add-ons. It must be between a /16 and a /20`,
				},
				resource.Attribute{
					Name:        "outbound_ips",
					Description: `The space's stable outbound [NAT IPs](https://devcenter.heroku.com/articles/platform-api-reference#space-network-address-translation). The ` + "`" + `organization` + "`" + ` block supports:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Heroku Private Space. In Heroku, this is also the unique .`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the Heroku Private Space.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which the Heroku Private Space is deployed.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the Heroku Private Space. Either ` + "`" + `allocating` + "`" + ` or ` + "`" + `allocated` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shield",
					Description: `Whether or not the space has [Shield](https://devcenter.heroku.com/articles/private-spaces#shield-private-spaces) turned on. One of ` + "`" + `on` + "`" + ` or ` + "`" + `off` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The Heroku Team that owns this space. The fields for this block are documented below.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The RFC-1918 CIDR the Private Space will use. It must be a /16 in 10.0.0.0/8, 172.16.0.0/12 or 192.168.0.0/16`,
				},
				resource.Attribute{
					Name:        "data_cidr",
					Description: `The RFC-1918 CIDR that the Private Space will use for the Heroku-managed peering connection that’s automatically created when using Heroku Data add-ons. It must be between a /16 and a /20`,
				},
				resource.Attribute{
					Name:        "outbound_ips",
					Description: `The space's stable outbound [NAT IPs](https://devcenter.heroku.com/articles/platform-api-reference#space-network-address-translation). The ` + "`" + `organization` + "`" + ` block supports:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_space_peering_info",
			Category:         "Data Sources",
			ShortDescription: `Get peering information on a Heroku Private Space.`,
			Description: `

Use this data source to get peering information about a [Heroku Private Space](https://www.heroku.com/private-spaces).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Heroku Private Space. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `The AWS account ID that the Heroku Private Space runs in.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `The AWS region that the Heroku Private Space runs in.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID of the Heroku Private Space.`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `The CIDR block of the VPC ID.`,
				},
				resource.Attribute{
					Name:        "dyno_cidr_blocks",
					Description: `The CIDR blocks that the Dynos run on.`,
				},
				resource.Attribute{
					Name:        "unavailable_cidr_blocks",
					Description: `A list of unavailable CIDR blocks.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "aws_account_id",
					Description: `The AWS account ID that the Heroku Private Space runs in.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `The AWS region that the Heroku Private Space runs in.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID of the Heroku Private Space.`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `The CIDR block of the VPC ID.`,
				},
				resource.Attribute{
					Name:        "dyno_cidr_blocks",
					Description: `The CIDR blocks that the Dynos run on.`,
				},
				resource.Attribute{
					Name:        "unavailable_cidr_blocks",
					Description: `A list of unavailable CIDR blocks.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_team",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Heroku Team.`,
			Description: `

Use this data source to get information about a Heroku Team.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The team name ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Whether to use this team when none is specified`,
				},
				resource.Attribute{
					Name:        "credit_card_collections",
					Description: `Whether charges incurred by the team are paid by credit card`,
				},
				resource.Attribute{
					Name:        "membership_limit",
					Description: `Upper limit of members allowed in a team`,
				},
				resource.Attribute{
					Name:        "provisioned_licenses",
					Description: `Whether the team is provisioned licenses by Salesforce`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of team Will likely be either "enterprise" or "team"`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Whether to use this team when none is specified`,
				},
				resource.Attribute{
					Name:        "credit_card_collections",
					Description: `Whether charges incurred by the team are paid by credit card`,
				},
				resource.Attribute{
					Name:        "membership_limit",
					Description: `Upper limit of members allowed in a team`,
				},
				resource.Attribute{
					Name:        "provisioned_licenses",
					Description: `Whether the team is provisioned licenses by Salesforce`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `type of team Will likely be either "enterprise" or "team"`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "heroku_team_members",
			Category:         "Data Sources",
			ShortDescription: `Get information on members for a Heroku Team.`,
			Description: `

Use this data source to get information about members for a Heroku Team.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team",
					Description: `(Required) The team name.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Required) List of roles. Acceptable values are ` + "`" + `admin` + "`" + `, ` + "`" + `member` + "`" + `, ` + "`" + `viewer` + "`" + `, ` + "`" + `collaborator` + "`" + `, ` + "`" + `owner` + "`" + `. At least one role must be specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the team.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `All members of the team that have a specified role defined in the ` + "`" + `roles` + "`" + ` attribute above.`,
				},
				resource.Attribute{
					Name:        "team_member_id",
					Description: `Unique identifier of the team member on the team.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Unique identifier of the team member. This is the member's user ID in Heroku.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the team member.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role in the team.`,
				},
				resource.Attribute{
					Name:        "federated",
					Description: `Whether the user is federated and belongs to an Identity Provider.`,
				},
				resource.Attribute{
					Name:        "two_factor_authentication",
					Description: `Whether the Enterprise team member has two-factor authentication enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the team.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `All members of the team that have a specified role defined in the ` + "`" + `roles` + "`" + ` attribute above.`,
				},
				resource.Attribute{
					Name:        "team_member_id",
					Description: `Unique identifier of the team member on the team.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `Unique identifier of the team member. This is the member's user ID in Heroku.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `Email address of the team member.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Role in the team.`,
				},
				resource.Attribute{
					Name:        "federated",
					Description: `Whether the user is federated and belongs to an Identity Provider.`,
				},
				resource.Attribute{
					Name:        "two_factor_authentication",
					Description: `Whether the Enterprise team member has two-factor authentication enabled.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"heroku_addon":              0,
		"heroku_app":                1,
		"heroku_pipeline":           2,
		"heroku_space":              3,
		"heroku_space_peering_info": 4,
		"heroku_team":               5,
		"heroku_team_members":       6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
