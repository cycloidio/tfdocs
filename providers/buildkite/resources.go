package buildkite

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "buildkite_agent_token",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create and manage agent tokens.

Buildkite Documentation: https://buildkite.com/docs/agent/v3/tokens

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) This is the description of the agent token. -> Changing ` + "`" + `description` + "`" + ` will cause the resource to be destroyed and re-created. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Graphql ID of the created agent token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The value of the created agent token.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the token. ## Import Tokens can be imported using the ` + "`" + `GraphQL ID` + "`" + ` (not UUID), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import buildkite_agent_token.fleet QWdlbnRUb2tlbi0tLTQzNWNhZDU4LWU4MWQtNDVhZi04NjM3LWIxY2Y4MDcwMjM4ZA== ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Graphql ID of the created agent token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The value of the created agent token.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the token. ## Import Tokens can be imported using the ` + "`" + `GraphQL ID` + "`" + ` (not UUID), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import buildkite_agent_token.fleet QWdlbnRUb2tlbi0tLTQzNWNhZDU4LWU4MWQtNDVhZi04NjM3LWIxY2Y4MDcwMjM4ZA== ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buildkite_pipeline",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create and manage pipelines for repositories.

Buildkite Documentation: https://buildkite.com/docs/pipelines

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buildkite_pipeline_schedule",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create and manage pipeline schedules.

Buildkite Documentation: https://buildkite.com/docs/pipelines/scheduled-builds

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "pipeline_id",
					Description: `(Required) Terraform resource ID of a buildkite pipeline (Buildkite GraphQL ID).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Schedule label.`,
				},
				resource.Attribute{
					Name:        "cronline",
					Description: `(Required) Schedule interval (see [docs](https://buildkite.com/docs/pipelines/scheduled-builds#schedule-intervals)).`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Required) The branch to use for the build.`,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `(Optional, Default: ` + "`" + `HEAD` + "`" + `) The commit ref to use for the build.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional, Default: ` + "`" + `Scheduled build` + "`" + `) The message to use for the build.`,
				},
				resource.Attribute{
					Name:        "env",
					Description: `(Optional) A map of environment variables to use for the build.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, Default: ` + "`" + `true` + "`" + `) Whether the schedule should run. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The GraphQL ID of the pipeline schedule`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the pipeline schedule ## Import Pipeline schedules can be imported using a slug (which consists of ` + "`" + `$BUILDKITE_ORGANIZATION_SLUG/$BUILDKITE_PIPELINE_SLUG/$PIPELINE_SCHEDULE_UUID` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import buildkite_pipeline_schedule.test myorg/test/1be3e7c7-1e03-4011-accf-b2d8eec90222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The GraphQL ID of the pipeline schedule`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the pipeline schedule ## Import Pipeline schedules can be imported using a slug (which consists of ` + "`" + `$BUILDKITE_ORGANIZATION_SLUG/$BUILDKITE_PIPELINE_SLUG/$PIPELINE_SCHEDULE_UUID` + "`" + `), e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import buildkite_pipeline_schedule.test myorg/test/1be3e7c7-1e03-4011-accf-b2d8eec90222 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buildkite_team",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows you to create and manage teams.

Buildkite Documentation: https://buildkite.com/docs/pipelines/permissions

Note: You must first enable Teams on your organization.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the team.`,
				},
				resource.Attribute{
					Name:        "privacy",
					Description: `(Required) The privacy level to set the team too.`,
				},
				resource.Attribute{
					Name:        "default_team",
					Description: `(Required) Whether to assign this team to a user by default.`,
				},
				resource.Attribute{
					Name:        "default_member_role",
					Description: `(Required) Default role to assign to a team member.`,
				},
				resource.Attribute{
					Name:        "members_can_create_pipelines",
					Description: `(Optional) Whether team members can create.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description to assign to the team. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The name of the team.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID for the team.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `The name of the team.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID for the team.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buildkite_team_member",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

This resource allows manage team membership for existing organization users.

The user must already be part of the organization to which you are managing team membership. This will not invite a new user to the organization.

Buildkite Documentation: https://buildkite.com/docs/pipelines/permissions

Note: You must first enable Teams on your organization.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "role",
					Description: `(Required) Either MEMBER or MAINTAINER.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The GraphQL ID of the team to add to/remove from.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The GraphQL ID of the user to add/remove. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID for the team membership. ## Import Team members can be imported using the GraphQL ID of the membership. Note this is different to the ID of the user. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import buildkite_team_member.a_smith VGVhbU1lbWJlci0tLTVlZDEyMmY2LTM2NjQtNDI1MS04YzMwLTc4NjRiMDdiZDQ4Zg== ` + "`" + `` + "`" + `` + "`" + ` To find the ID of a team member you are trying to import you can use the GraphQL snippet below. A link to this snippet can also be found at https://buildkite.com/user/graphql/console/c6a2cc65-dc59-49df-95c6-7167b68dbd5d. You will need fo fill in the organization slug and search terms for teams and members. Both search terms work on the name of the associated object. ` + "`" + `` + "`" + `` + "`" + `graphql query { organization(slug: "") { teams(first: 2, search: "") { edges { node { members(first: 2, search: "") { edges { node { id } } } } } } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID for the team membership. ## Import Team members can be imported using the GraphQL ID of the membership. Note this is different to the ID of the user. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import buildkite_team_member.a_smith VGVhbU1lbWJlci0tLTVlZDEyMmY2LTM2NjQtNDI1MS04YzMwLTc4NjRiMDdiZDQ4Zg== ` + "`" + `` + "`" + `` + "`" + ` To find the ID of a team member you are trying to import you can use the GraphQL snippet below. A link to this snippet can also be found at https://buildkite.com/user/graphql/console/c6a2cc65-dc59-49df-95c6-7167b68dbd5d. You will need fo fill in the organization slug and search terms for teams and members. Both search terms work on the name of the associated object. ` + "`" + `` + "`" + `` + "`" + `graphql query { organization(slug: "") { teams(first: 2, search: "") { edges { node { members(first: 2, search: "") { edges { node { id } } } } } } } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"buildkite_agent_token":       0,
		"buildkite_pipeline":          1,
		"buildkite_pipeline_schedule": 2,
		"buildkite_team":              3,
		"buildkite_team_member":       4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
