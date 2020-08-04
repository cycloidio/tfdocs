package tfe

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tfe_organization_membership",
			Category:         "Data Sources",
			ShortDescription: `Get information on an organization membership.`,
			Description: `

Use this data source to get information about an organization membership.

~> **NOTE:** This data source requires using the provider with Terraform Cloud or
an instance of Terraform Enterprise at least as recent as v202004-1.

~> **NOTE:** If a user updates their email address, configurations using the email address should
be updated manually.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email of the user. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The organization membership ID.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user associated with the organization membership.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The organization membership ID.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user associated with the organization membership.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a SSH key.`,
			Description: `

Use this data source to get information about a SSH key.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the SSH key.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team",
			Category:         "Data Sources",
			ShortDescription: `Get information on a team.`,
			Description: `

Use this data source to get information about a team.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the team.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team_access",
			Category:         "Data Sources",
			ShortDescription: `Get information on team permissions on a workspace.`,
			Description: `

Use this data source to get information about team permissions for a workspace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) ID of the team.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Required) ID of the workspace. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `The type of access granted to the team on the workspace.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `The permissions granted to the team on the workspaces for each whatever. The ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "runs",
					Description: `The permission granted to runs. Valid values are ` + "`" + `read` + "`" + `, ` + "`" + `plan` + "`" + `, or ` + "`" + `apply` + "`" + ``,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `The permissions granted to variables. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, or ` + "`" + `write` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state_versions",
					Description: `The permissions granted to state versions. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `read-outputs` + "`" + `, ` + "`" + `read` + "`" + `, or ` + "`" + `write` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sentinel_mocks",
					Description: `The permissions granted to Sentinel mocks. Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `read` + "`" + ``,
				},
				resource.Attribute{
					Name:        "workspace_locking",
					Description: `Whether permission is granted to manually lock the workspace or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access",
					Description: `The type of access granted to the team on the workspace.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `The permissions granted to the team on the workspaces for each whatever. The ` + "`" + `permissions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "runs",
					Description: `The permission granted to runs. Valid values are ` + "`" + `read` + "`" + `, ` + "`" + `plan` + "`" + `, or ` + "`" + `apply` + "`" + ``,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `The permissions granted to variables. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, or ` + "`" + `write` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state_versions",
					Description: `The permissions granted to state versions. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `read-outputs` + "`" + `, ` + "`" + `read` + "`" + `, or ` + "`" + `write` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sentinel_mocks",
					Description: `The permissions granted to Sentinel mocks. Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `read` + "`" + ``,
				},
				resource.Attribute{
					Name:        "workspace_locking",
					Description: `Whether permission is granted to manually lock the workspace or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_workspace",
			Category:         "Data Sources",
			ShortDescription: `Get information on a workspace.`,
			Description: `

Use this data source to get information about a workspace.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the workspace.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The workspace ID.`,
				},
				resource.Attribute{
					Name:        "auto_apply",
					Description: `Indicates whether to automatically apply changes when a Terraform plan is successful.`,
				},
				resource.Attribute{
					Name:        "file_triggers_enabled",
					Description: `Indicates whether runs are triggered based on the changed files in a VCS push (if ` + "`" + `true` + "`" + `) or always triggered on every push (if ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `Indicates whether the workspace is using remote execution mode. Set to ` + "`" + `false` + "`" + ` to switch execution mode to local. ` + "`" + `true` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "queue_all_runs",
					Description: `Indicates whether all runs should be queued.`,
				},
				resource.Attribute{
					Name:        "ssh_key_id",
					Description: `The ID of an SSH key assigned to the workspace.`,
				},
				resource.Attribute{
					Name:        "terraform_version",
					Description: `The version of Terraform used for this workspace.`,
				},
				resource.Attribute{
					Name:        "trigger_prefixes",
					Description: `List of repository-root-relative paths which describe all locations to be tracked for changes.`,
				},
				resource.Attribute{
					Name:        "vcs_repo",
					Description: `Settings for the workspace's VCS repository.`,
				},
				resource.Attribute{
					Name:        "working_directory",
					Description: `A relative path that Terraform will execute within. The ` + "`" + `vcs_repo` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `A reference to your VCS repository in the format ` + "`" + `:org/:repo` + "`" + ` where ` + "`" + `:org` + "`" + ` and ` + "`" + `:repo` + "`" + ` refer to the organization and repository in your VCS provider.`,
				},
				resource.Attribute{
					Name:        "ingress_submodules",
					Description: `Indicates whether submodules should be fetched when cloning the VCS repository.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `OAuth token ID of the configured VCS connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The workspace ID.`,
				},
				resource.Attribute{
					Name:        "auto_apply",
					Description: `Indicates whether to automatically apply changes when a Terraform plan is successful.`,
				},
				resource.Attribute{
					Name:        "file_triggers_enabled",
					Description: `Indicates whether runs are triggered based on the changed files in a VCS push (if ` + "`" + `true` + "`" + `) or always triggered on every push (if ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `Indicates whether the workspace is using remote execution mode. Set to ` + "`" + `false` + "`" + ` to switch execution mode to local. ` + "`" + `true` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "queue_all_runs",
					Description: `Indicates whether all runs should be queued.`,
				},
				resource.Attribute{
					Name:        "ssh_key_id",
					Description: `The ID of an SSH key assigned to the workspace.`,
				},
				resource.Attribute{
					Name:        "terraform_version",
					Description: `The version of Terraform used for this workspace.`,
				},
				resource.Attribute{
					Name:        "trigger_prefixes",
					Description: `List of repository-root-relative paths which describe all locations to be tracked for changes.`,
				},
				resource.Attribute{
					Name:        "vcs_repo",
					Description: `Settings for the workspace's VCS repository.`,
				},
				resource.Attribute{
					Name:        "working_directory",
					Description: `A relative path that Terraform will execute within. The ` + "`" + `vcs_repo` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `A reference to your VCS repository in the format ` + "`" + `:org/:repo` + "`" + ` where ` + "`" + `:org` + "`" + ` and ` + "`" + `:repo` + "`" + ` refer to the organization and repository in your VCS provider.`,
				},
				resource.Attribute{
					Name:        "ingress_submodules",
					Description: `Indicates whether submodules should be fetched when cloning the VCS repository.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `OAuth token ID of the configured VCS connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_workspace_ids",
			Category:         "Data Sources",
			ShortDescription: `Get information on (external) workspace IDs.`,
			Description: `

Use this data source to get a map of (external) workspace IDs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `(Required) A list of workspace names to search for. Names that don't match a real workspace will be omitted from the results, but are not an error. To select _all_ workspaces for an organization, provide a list with a single asterisk, like ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported: ~>`,
				},
				resource.Attribute{
					Name:        "full_names",
					Description: `A map of workspace names and their full names, which look like ` + "`" + `<ORGANIZATION>/<WORKSPACE>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_ids",
					Description: `A map of workspace names and their opaque, immutable IDs, which look like ` + "`" + `ws-<RANDOM STRING>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_names",
					Description: `A map of workspace names and their full names, which look like ` + "`" + `<ORGANIZATION>/<WORKSPACE>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "external_ids",
					Description: `A map of workspace names and their opaque, immutable IDs, which look like ` + "`" + `ws-<RANDOM STRING>` + "`" + `.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"tfe_organization_membership": 0,
		"tfe_ssh_key":                 1,
		"tfe_team":                    2,
		"tfe_team_access":             3,
		"tfe_workspace":               4,
		"tfe_workspace_ids":           5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
