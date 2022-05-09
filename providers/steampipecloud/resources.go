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
	}

	resourcesMap = map[string]int{

		"steampipecloud_connection":           0,
		"steampipecloud_organization":         1,
		"steampipecloud_organization_member":  2,
		"steampipecloud_workspace":            3,
		"steampipecloud_workspace_connection": 4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
