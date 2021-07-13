package tfe

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tfe_agent_pool",
			Category:         "Data Sources",
			ShortDescription: `Get information on an agent pool.`,
			Description: `

Use this data source to get information about an agent pool.

~> **NOTE:** This data source requires using the provider with Terraform Cloud and a Terraform Cloud 
for Business account. 
[Learn more about Terraform Cloud pricing here](https://www.hashicorp.com/products/terraform/pricing).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the agent pool.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The agent pool ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The agent pool ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_ip_ranges",
			Category:         "Data Sources",
			ShortDescription: `Get Terraform Cloud/Enterprise's IP ranges of its services`,
			Description: `

Use this data source to retrieve a list of Terraform Cloud's IP ranges. For more information about these IP ranges, view our [documentation about Terraform Cloud IP Ranges](https://www.terraform.io/docs/cloud/architectural-details/ip-ranges.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "api",
					Description: `The list of IP ranges in CIDR notation used for connections from user site to Terraform Cloud APIs.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `The list of IP ranges in CIDR notation used for notifications.`,
				},
				resource.Attribute{
					Name:        "sentinel",
					Description: `The list of IP ranges in CIDR notation used for outbound requests from Sentinel policies.`,
				},
				resource.Attribute{
					Name:        "vcs",
					Description: `The list of IP ranges in CIDR notation used for connecting to VCS providers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api",
					Description: `The list of IP ranges in CIDR notation used for connections from user site to Terraform Cloud APIs.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `The list of IP ranges in CIDR notation used for notifications.`,
				},
				resource.Attribute{
					Name:        "sentinel",
					Description: `The list of IP ranges in CIDR notation used for outbound requests from Sentinel policies.`,
				},
				resource.Attribute{
					Name:        "vcs",
					Description: `The list of IP ranges in CIDR notation used for connecting to VCS providers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_oauth_client",
			Category:         "Data Sources",
			ShortDescription: `Get information on an OAuth client.`,
			Description: `

Use this data source to get information about an OAuth client.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "oauth_client_id",
					Description: `(Required) ID of the OAuth client. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The OAuth client ID. This will match ` + "`" + `oauth_client_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "api_url",
					Description: `The client's API URL.`,
				},
				resource.Attribute{
					Name:        "http_url",
					Description: `The client's HTTP URL.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `The ID of the OAuth token associated with the OAuth client.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The SSH key assigned to the OAuth client.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The OAuth client ID. This will match ` + "`" + `oauth_client_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "api_url",
					Description: `The client's API URL.`,
				},
				resource.Attribute{
					Name:        "http_url",
					Description: `The client's HTTP URL.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `The ID of the OAuth token associated with the OAuth client.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `The SSH key assigned to the OAuth client.`,
				},
			},
		},
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

~> **NOTE:** Using ` + "`" + `global_remote_state` + "`" + ` or ` + "`" + `remote_state_consumer_ids` + "`" + ` requires using the provider with Terraform Cloud or an instance of Terraform Enterprise at least as recent as v202104-1.

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
					Name:        "allow_destroy_plan",
					Description: `Indicates whether destroy plans can be queued on the workspace.`,
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
					Name:        "global_remote_state",
					Description: `(Optional) Whether the workspace should allow all workspaces in the organization to access its state data during runs. If false, then only specifically approved workspaces can access its state (determined by the ` + "`" + `remote_state_consumer_ids` + "`" + ` argument).`,
				},
				resource.Attribute{
					Name:        "remote_state_consumer_ids",
					Description: `(Optional) A set of workspace IDs that will be set as the remote state consumers for the given workspace. Cannot be used if ` + "`" + `global_remote_state` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `Indicates whether the workspace is using remote execution mode. Set to ` + "`" + `false` + "`" + ` to switch execution mode to local. ` + "`" + `true` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "policy_check_failures",
					Description: `The number of policy check failures from the latest run.`,
				},
				resource.Attribute{
					Name:        "queue_all_runs",
					Description: `Indicates whether the workspace will automatically perform runs in response to webhooks immediately after its creation. If ` + "`" + `false` + "`" + `, an initial run must be manually queued to enable future automatic runs.`,
				},
				resource.Attribute{
					Name:        "resource_count",
					Description: `The number of resources managed by the workspace.`,
				},
				resource.Attribute{
					Name:        "run_failures",
					Description: `The number of run failures on the workspace.`,
				},
				resource.Attribute{
					Name:        "runs_count",
					Description: `The number of runs on the workspace.`,
				},
				resource.Attribute{
					Name:        "speculative_enabled",
					Description: `Indicates whether this workspace allows speculative plans.`,
				},
				resource.Attribute{
					Name:        "ssh_key_id",
					Description: `The ID of an SSH key assigned to the workspace.`,
				},
				resource.Attribute{
					Name:        "structured_run_output_enabled",
					Description: `Indicates whether runs in this workspace use the enhanced apply UI.`,
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
					Description: `A reference to your VCS repository in the format ` + "`" + `<organization>/<repository>` + "`" + ` where ` + "`" + `<organization>` + "`" + ` and ` + "`" + `<repository>` + "`" + ` refer to the organization and repository in your VCS provider.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The repository branch that Terraform will execute from.`,
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
					Name:        "allow_destroy_plan",
					Description: `Indicates whether destroy plans can be queued on the workspace.`,
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
					Name:        "global_remote_state",
					Description: `(Optional) Whether the workspace should allow all workspaces in the organization to access its state data during runs. If false, then only specifically approved workspaces can access its state (determined by the ` + "`" + `remote_state_consumer_ids` + "`" + ` argument).`,
				},
				resource.Attribute{
					Name:        "remote_state_consumer_ids",
					Description: `(Optional) A set of workspace IDs that will be set as the remote state consumers for the given workspace. Cannot be used if ` + "`" + `global_remote_state` + "`" + ` is set to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `Indicates whether the workspace is using remote execution mode. Set to ` + "`" + `false` + "`" + ` to switch execution mode to local. ` + "`" + `true` + "`" + ` by default.`,
				},
				resource.Attribute{
					Name:        "policy_check_failures",
					Description: `The number of policy check failures from the latest run.`,
				},
				resource.Attribute{
					Name:        "queue_all_runs",
					Description: `Indicates whether the workspace will automatically perform runs in response to webhooks immediately after its creation. If ` + "`" + `false` + "`" + `, an initial run must be manually queued to enable future automatic runs.`,
				},
				resource.Attribute{
					Name:        "resource_count",
					Description: `The number of resources managed by the workspace.`,
				},
				resource.Attribute{
					Name:        "run_failures",
					Description: `The number of run failures on the workspace.`,
				},
				resource.Attribute{
					Name:        "runs_count",
					Description: `The number of runs on the workspace.`,
				},
				resource.Attribute{
					Name:        "speculative_enabled",
					Description: `Indicates whether this workspace allows speculative plans.`,
				},
				resource.Attribute{
					Name:        "ssh_key_id",
					Description: `The ID of an SSH key assigned to the workspace.`,
				},
				resource.Attribute{
					Name:        "structured_run_output_enabled",
					Description: `Indicates whether runs in this workspace use the enhanced apply UI.`,
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
					Description: `A reference to your VCS repository in the format ` + "`" + `<organization>/<repository>` + "`" + ` where ` + "`" + `<organization>` + "`" + ` and ` + "`" + `<repository>` + "`" + ` refer to the organization and repository in your VCS provider.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `The repository branch that Terraform will execute from.`,
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
			ShortDescription: `Get information on workspace IDs.`,
			Description: `

Use this data source to get a map of workspace IDs.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `(Required) A list of workspace names to search for. Names that don't match a real workspace will be omitted from the results, but are not an error. To select _all_ workspaces for an organization, provide a list with a single asterisk, like ` + "`" + `["`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "full_names",
					Description: `A map of workspace names and their full names, which look like ` + "`" + `<ORGANIZATION>/<WORKSPACE>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A map of workspace names and their opaque, immutable IDs, which look like ` + "`" + `ws-<RANDOM STRING>` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_names",
					Description: `A map of workspace names and their full names, which look like ` + "`" + `<ORGANIZATION>/<WORKSPACE>` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A map of workspace names and their opaque, immutable IDs, which look like ` + "`" + `ws-<RANDOM STRING>` + "`" + `.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"tfe_agent_pool":              0,
		"tfe_ip_ranges":               1,
		"tfe_oauth_client":            2,
		"tfe_organization_membership": 3,
		"tfe_ssh_key":                 4,
		"tfe_team":                    5,
		"tfe_team_access":             6,
		"tfe_workspace":               7,
		"tfe_workspace_ids":           8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
