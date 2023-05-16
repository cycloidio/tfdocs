package steampipecloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_connection",
			Category:         "Resources",
			ShortDescription: `Provides a Steampipe Connection resource. The ` + "`" + `Steampipe Cloud Connection` + "`" + ` represents a set of tables for a single data source. Each connection is represented as a distinct Postgres schema. In order to query data, you'll need at least one connection. Connections are defined at the user account or organization level, and they can be shared by multiple workspaces within the account or organization.`,
			Description: `

Manages a connection, which is defined at the user account or organization level.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_organization",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Organization` + "`" + ` includes multiple users and is intended for organizations to collaborate and share workspaces and connections.`,
			Description: `

Manages an organization.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_organization_member",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Organization Member` + "`" + ` provides the members of an organization who can collaborate and share workspaces and connections. This resource allows you to add/remove users from your organization. When applied, an invitation will be sent to the user to become part of the organization. When destroyed, either the invitation will be cancelled or the user will be removed.`,
			Description: `

Manages an organization membership.

This resource allows you to add/remove users from your organization. When
applied, an invitation will be sent to the user to become part of the
organization. When destroyed, either the invitation will be cancelled or the
user will be removed.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_organization_workspace_member",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Organization Workspace Member` + "`" + ` provides the members of a workspace belonging to an organization who can collaborate, run queries and snapshots. This resource allows you to add/remove users from a workspace of your organization. When applied, an invitation will be sent to the user to become part of the workspace. When destroyed, either the invitation will be cancelled or the user will be removed.`,
			Description: `

Manages the membership of a workspace in an organization.

This resource allows you to add/remove users from a workspace of your 
organization. When applied, the user whose handle is passed will be
added to the workspace with the role as specified in the configuration.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_user_preferences",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud User Preferences` + "`" + ` represents the preferences settings for a user.`,
			Description: `

Allows a user to manage various preferences related to their steampipe cloud profile e.g. email preferences.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace` + "`" + ` provides a bounded context for managing, operating, and securing Steampipe resources. A workspace comprises a single Steampipe database instance as well as a directory of mod resources such as queries, benchmarks, and controls. Workspaces allow you to separate your Steampipe instances for security, operational, or organizational purposes.`,
			Description: `

Manages a workspace, which is defined at the user account or organization level.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace_aggregator",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace Aggregator` + "`" + ` represents aggregators configured in a workspace.`,
			Description: `

Manages a workspace aggregator.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace_connection",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace Connection` + "`" + ` represents a set of connections that are currently attached to the workspace. This resource can be used multiple times with the same connection for non-overlapping workspaces.`,
			Description: `

Manages a workspace connection association.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace_mod",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace Mod` + "`" + ` represents mods that are currently installed in the workspace.`,
			Description: `

Manages a workspace mod.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace_mod_variable",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace Mod Variable` + "`" + ` represents variables for a mod installed in the workspace.`,
			Description: `

Manages a workspace mod variable.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace_pipeline",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace Pipeline` + "`" + ` represents pipelines configured in a workspace.`,
			Description: `

Manages a workspace pipeline.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "steampipecloud_workspace_snapshot",
			Category:         "Resources",
			ShortDescription: `The ` + "`" + `Steampipe Cloud Workspace Snapshot` + "`" + ` represents snapshots captured in a workspace.`,
			Description: `

Manages a workspace snapshot.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"steampipecloud_connection":                    0,
		"steampipecloud_organization":                  1,
		"steampipecloud_organization_member":           2,
		"steampipecloud_organization_workspace_member": 3,
		"steampipecloud_user_preferences":              4,
		"steampipecloud_workspace":                     5,
		"steampipecloud_workspace_aggregator":          6,
		"steampipecloud_workspace_connection":          7,
		"steampipecloud_workspace_mod":                 8,
		"steampipecloud_workspace_mod_variable":        9,
		"steampipecloud_workspace_pipeline":            10,
		"steampipecloud_workspace_snapshot":            11,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
