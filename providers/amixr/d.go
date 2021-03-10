package amixr

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "amixr_action",
			Category:         "Data Sources",
			ShortDescription: `Get information about an action (outgoing webhook).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The action name.`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `(Required) The ID of the integration. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found action.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found action.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_schedule",
			Category:         "Data Sources",
			ShortDescription: `Get information about a schedule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The schedule's name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the found schedule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the found schedule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_slack_channel",
			Category:         "Data Sources",
			ShortDescription: `Get information about a Slack channel.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Slack channel name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "slack_id",
					Description: `The Slack ID of the found Slack channel.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slack_id",
					Description: `The Slack ID of the found Slack channel.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_user",
			Category:         "Data Sources",
			ShortDescription: `Get information about a user.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) User's email. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `User's role in team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The ID of current team.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `User's role in team.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The ID of current team.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "amixr_user_group",
			Category:         "Data Sources",
			ShortDescription: `Get information about a User Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slack_handle",
					Description: `(Required) The User Group Slack handle. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found User Group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found User Group.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"amixr_action":        0,
		"amixr_schedule":      1,
		"amixr_slack_channel": 2,
		"amixr_user":          3,
		"amixr_user_group":    4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
