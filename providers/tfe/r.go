package tfe

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "tfe_notification_configuration",
			Category:         "Resources",
			ShortDescription: `Manages notifications configurations.`,
			Description:      ``,
			Keywords: []string{
				"notification",
				"configuration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the notification configuration.`,
				},
				resource.Attribute{
					Name:        "destination_type",
					Description: `(Required) The type of notification configuration payload to send. Valid values are ` + "`" + `email` + "`" + `, ` + "`" + `generic` + "`" + ` or ` + "`" + `slack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "email_addresses",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "email_user_ids",
					Description: `(Optional) A list of user IDs. This value _must not_ be provided if ` + "`" + `destination_type` + "`" + ` is ` + "`" + `generic` + "`" + ` or ` + "`" + `slack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether the notification configuration should be enabled or not. Disabled configurations will not send any notifications. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) A write-only secure token for the notification configuration, which can be used by the receiving server to verify request authenticity when configured for notification configurations with a destination type of ` + "`" + `generic` + "`" + `. Defaults to ` + "`" + `null` + "`" + `. This value _must not_ be provided if ` + "`" + `destination_type` + "`" + ` is ` + "`" + `email` + "`" + ` or ` + "`" + `slack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Optional) The array of triggers for which this notification configuration will send notifications. Valid values are ` + "`" + `run:created` + "`" + `, ` + "`" + `run:planning` + "`" + `, ` + "`" + `run:needs_attention` + "`" + `, ` + "`" + `run:applying` + "`" + ` ` + "`" + `run:completed` + "`" + `, ` + "`" + `run:errored` + "`" + `. If omitted, no notification triggers are configured.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required if ` + "`" + `destination_type` + "`" + ` is ` + "`" + `generic` + "`" + ` or ` + "`" + `slack` + "`" + `) The HTTP or HTTPS URL of the notification configuration where notification requests will be made. This value _must not_ be provided if ` + "`" + `destination_type` + "`" + ` is ` + "`" + `email` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The id of the workspace that owns the notification configuration. This value _must not_ be provided if ` + "`" + `workspace_external_id` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the notification configuration. ## Import Notification configurations can be imported; use ` + "`" + `<NOTIFICATION CONFIGURATION ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_notification_configuration.test nc-qV9JnKRkmtMa4zcA ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the notification configuration. ## Import Notification configurations can be imported; use ` + "`" + `<NOTIFICATION CONFIGURATION ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_notification_configuration.test nc-qV9JnKRkmtMa4zcA ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_oauth_client",
			Category:         "Resources",
			ShortDescription: `Manages OAuth clients.`,
			Description:      ``,
			Keywords: []string{
				"oauth",
				"client",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "api_url",
					Description: `(Required) The base URL of your VCS provider's API (e.g. ` + "`" + `https://api.github.com` + "`" + ` or ` + "`" + `https://ghe.example.com/api/v3` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "http_url",
					Description: `(Required) The homepage of your VCS provider (e.g. ` + "`" + `https://github.com` + "`" + ` or ` + "`" + `https://ghe.example.com` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "oauth_token",
					Description: `(Required) The token string you were given by your VCS provider.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required for ` + "`" + `ado_server` + "`" + `) The text of the private key associated with your Azure DevOps Server account`,
				},
				resource.Attribute{
					Name:        "service_provider",
					Description: `(Required) The VCS provider being connected with. Valid options are ` + "`" + `ado_server` + "`" + `, ` + "`" + `ado_services` + "`" + `, ` + "`" + `github` + "`" + `, ` + "`" + `github_enterprise` + "`" + `, ` + "`" + `gitlab_hosted` + "`" + `, ` + "`" + `gitlab_community_edition` + "`" + `, or ` + "`" + `gitlab_enterprise_edition` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the OAuth client.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `The ID of the OAuth token associated with the OAuth client.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the OAuth client.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `The ID of the OAuth token associated with the OAuth client.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_organization",
			Category:         "Resources",
			ShortDescription: `Manages organizations.`,
			Description:      ``,
			Keywords: []string{
				"organization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Admin email address.`,
				},
				resource.Attribute{
					Name:        "session_timeout_minutes",
					Description: `(Optional) Session timeout after inactivity. Defaults to ` + "`" + `20160` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "session_remember_minutes",
					Description: `(Optional) Session expiration. Defaults to ` + "`" + `20160` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "collaborator_auth_policy",
					Description: `(Optional) Authentication policy (` + "`" + `password` + "`" + ` or ` + "`" + `two_factor_mandatory` + "`" + `). Defaults to ` + "`" + `password` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owners_team_saml_role_id",
					Description: `(Optional) The name of the "owners" team. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of the organization. ## Import Organizations can be imported; use ` + "`" + `<ORGANIZATION NAME>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_organization.test my-org-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of the organization. ## Import Organizations can be imported; use ` + "`" + `<ORGANIZATION NAME>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_organization.test my-org-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_organization_membership",
			Category:         "Resources",
			ShortDescription: `Add or remove a user from an organization.`,
			Description:      ``,
			Keywords: []string{
				"organization",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) Email of the user to add. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The organization membership ID.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user associated with the organization membership. Organization memberships can be imported; use ` + "`" + `<ORGANIZATION MEMBERSHIP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_organization_membership.test ou-wAs3zYmWAhYK7peR ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The organization membership ID.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user associated with the organization membership. Organization memberships can be imported; use ` + "`" + `<ORGANIZATION MEMBERSHIP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_organization_membership.test ou-wAs3zYmWAhYK7peR ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_organization_token",
			Category:         "Resources",
			ShortDescription: `Generates a new organization token, replacing any existing token.`,
			Description:      ``,
			Keywords: []string{
				"organization",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "force_regenerate",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, a new token will be generated even if a token already exists. This will invalidate the existing token! ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The generated token. ## Import Organization tokens can be imported; use ` + "`" + `<ORGANIZATION NAME>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_organization_token.test my-org-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The generated token. ## Import Organization tokens can be imported; use ` + "`" + `<ORGANIZATION NAME>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_organization_token.test my-org-name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_policy_set",
			Category:         "Resources",
			ShortDescription: `Manages policy sets.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"set",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the policy set.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the policy set's purpose.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `(Optional) Whether or not policies in this set will apply to all workspaces. Defaults to ` + "`" + `false` + "`" + `. This value _must not_ be provided if ` + "`" + `workspace_ids` + "`" + ` or ` + "`" + `workspace_external_ids` + "`" + ` are provided.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "policies_path",
					Description: `(Optional) The sub-path within the attached VCS repository to ingress when using ` + "`" + `vcs_repo` + "`" + `. All files and directories outside of this sub-path will be ignored. This option can only be supplied when ` + "`" + `vcs_repo` + "`" + ` is present. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "policy_ids",
					Description: `(Optional) A list of Sentinel policy IDs. This value _must not_ be provided if ` + "`" + `vcs_repo` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "vcs_repo",
					Description: `(Optional) Settings for the policy sets VCS repository. Forces a new resource if changed. This value _must not_ be provided if ` + "`" + `policy_ids` + "`" + ` are provided.`,
				},
				resource.Attribute{
					Name:        "workspace_ids",
					Description: `(Optional) A list of workspace IDs. This value _must not_ be provided if ` + "`" + `global` + "`" + ` or ` + "`" + `workspace_external_ids` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `(Required) A reference to your VCS repository in the format ` + "`" + `:org/:repo` + "`" + ` where ` + "`" + `:org` + "`" + ` and ` + "`" + `:repo` + "`" + ` refer to the organization and repository in your VCS provider.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Optional) The repository branch that Terraform will execute from. Default to ` + "`" + `master` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress_submodules",
					Description: `(Optional) Whether submodules should be fetched when cloning the VCS repository. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `(Required) Token ID of the VCS Connection (OAuth Connection Token) to use. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy set. ## Import Policy sets can be imported; use ` + "`" + `<POLICY SET ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_policy_set.test polset-wAs3zYmWAhYK7peR ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy set. ## Import Policy sets can be imported; use ` + "`" + `<POLICY SET ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_policy_set.test polset-wAs3zYmWAhYK7peR ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_policy_set_parameter",
			Category:         "Resources",
			ShortDescription: `Manages policy set parameters.`,
			Description:      ``,
			Keywords: []string{
				"policy",
				"set",
				"parameter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Name of the parameter.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the parameter.`,
				},
				resource.Attribute{
					Name:        "sensitive",
					Description: `(Optional) Whether the value is sensitive. If true then the parameter is written once and not visible thereafter. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_set_id",
					Description: `(Required) The ID of the policy set that owns the parameter. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the parameter. ## Import Parameters can be imported; use ` + "`" + `<POLICY SET ID>/<PARAMETER ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_policy_set_parameter.test polset-wAs3zYmWAhYK7peR/var-5rTwnSaRPogw6apb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the parameter. ## Import Parameters can be imported; use ` + "`" + `<POLICY SET ID>/<PARAMETER ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_policy_set_parameter.test polset-wAs3zYmWAhYK7peR/var-5rTwnSaRPogw6apb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_registry_module",
			Category:         "Resources",
			ShortDescription: `Manages registry modules`,
			Description:      ``,
			Keywords: []string{
				"registry",
				"module",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vcs_repo",
					Description: `(Required) Settings for the registry module's VCS repository. Forces a new resource if changed. The ` + "`" + `vcs_repo` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "display_identifier",
					Description: `(Required) The display identifier for your VCS repository. For most VCS providers outside of BitBucket Cloud, this will match the ` + "`" + `identifier` + "`" + ` string.`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `(Required) A reference to your VCS repository in the format ` + "`" + `:org/:repo` + "`" + ` where ` + "`" + `:org` + "`" + ` and ` + "`" + `:repo` + "`" + ` refer to the organization (or project key, for Bitbucket Server) and repository in your VCS provider. The format for Azure DevOps is :org/:project/_git/:repo.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `(Required) Token ID of the VCS Connection (OAuth Connection Token) to use. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the registry module.`,
				},
				resource.Attribute{
					Name:        "module_provider",
					Description: `The provider of the registry module.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of registry module.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The name of the organization associated with the registry module. ## Import Registry modules can be imported; use ` + "`" + `<ORGANIZATION NAME>/<REGISTRY MODULE NAME>/<REGISTRY MODULE PROVIDER>/<REGISTRY MODULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_registry_module.test my-org-name/name/provider/mod-qV9JnKRkmtMa4zcA ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the registry module.`,
				},
				resource.Attribute{
					Name:        "module_provider",
					Description: `The provider of the registry module.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of registry module.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `The name of the organization associated with the registry module. ## Import Registry modules can be imported; use ` + "`" + `<ORGANIZATION NAME>/<REGISTRY MODULE NAME>/<REGISTRY MODULE PROVIDER>/<REGISTRY MODULE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_registry_module.test my-org-name/name/provider/mod-qV9JnKRkmtMa4zcA ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_run_trigger",
			Category:         "Resources",
			ShortDescription: `Manages run triggers`,
			Description:      ``,
			Keywords: []string{
				"run",
				"trigger",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "workspace_id",
					Description: `The id of the workspace that owns the run trigger. This is the workspace where runs will be triggered. This value _must not_ be provided if ` + "`" + `workspace_external_id` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "sourceable_id",
					Description: `(Required) The id of the sourceable. The sourceable must be a workspace. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the run trigger. ## Import Run triggers can be imported; use ` + "`" + `<RUN TRIGGER ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_run_trigger.test rt-qV9JnKRkmtMa4zcA ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the run trigger. ## Import Run triggers can be imported; use ` + "`" + `<RUN TRIGGER ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_run_trigger.test rt-qV9JnKRkmtMa4zcA ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_sentinel_policy",
			Category:         "Resources",
			ShortDescription: `Manages Sentinel policies.`,
			Description:      ``,
			Keywords: []string{
				"sentinel",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the policy's purpose.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The actual policy itself.`,
				},
				resource.Attribute{
					Name:        "enforce_mode",
					Description: `(Required) The enforcement level of the policy. Valid values are ` + "`" + `advisory` + "`" + `, ` + "`" + `hard-mandatory` + "`" + ` and ` + "`" + `soft-mandatory` + "`" + `. Defaults to ` + "`" + `soft-mandatory` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy. ## Import Sentinel policies can be imported; use ` + "`" + `<ORGANIZATION NAME>/<POLICY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_sentinel_policy.test my-org-name/pol-wAs3zYmWAhYK7peR ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy. ## Import Sentinel policies can be imported; use ` + "`" + `<ORGANIZATION NAME>/<POLICY ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_sentinel_policy.test my-org-name/pol-wAs3zYmWAhYK7peR ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_ssh_key",
			Category:         "Resources",
			ShortDescription: `Manages SSH keys.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name to identify the SSH key.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The text of the SSH private key. ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team",
			Category:         "Resources",
			ShortDescription: `Manages teams.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the team.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the team ("secret" or "organization"). Defaults to "secret".`,
				},
				resource.Attribute{
					Name:        "organization_access",
					Description: `(Optional) Settings for the team's [organization access](https://www.terraform.io/docs/cloud/users-teams-organizations/permissions.html#organization-level-permissions). The ` + "`" + `organization_access` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "manage_policies",
					Description: `(Optional) Allows members to create, edit, and delete the organization's Sentinel policies and override soft-mandatory policy checks.`,
				},
				resource.Attribute{
					Name:        "manage_workspaces",
					Description: `(Optional) Allows members to create and administrate all workspaces within the organization.`,
				},
				resource.Attribute{
					Name:        "manage_vcs_settings",
					Description: `(Optional) Allows members to manage the organization's VCS Providers and SSH keys. ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team_access",
			Category:         "Resources",
			ShortDescription: `Associate a team to permissions on a workspace.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"access",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) ID of the team to add to the workspace.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Required) ID of the workspace to which the team will be added.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `(Optional) Type of fixed access to grant. Valid values are ` + "`" + `admin` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `plan` + "`" + `, or ` + "`" + `write` + "`" + `. To use ` + "`" + `custom` + "`" + ` permissions, use a ` + "`" + `permissions` + "`" + ` block instead. This value _must not_ be provided if ` + "`" + `permissions` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) Permissions to grant using [custom workspace permissions](https://www.terraform.io/docs/cloud/users-teams-organizations/permissions.html#custom-workspace-permissions). This value _must not_ be provided if ` + "`" + `access` + "`" + ` is provided. The ` + "`" + `permissions` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "runs",
					Description: `(Required) The permission to grant the team on the workspace's runs. Valid values are ` + "`" + `read` + "`" + `, ` + "`" + `plan` + "`" + `, or ` + "`" + `apply` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "variables",
					Description: `(Required) The permission to grant the team on the workspace's variables. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, or ` + "`" + `write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state_versions",
					Description: `(Required) The permission to grant the team on the workspace's state versions. Valid values are ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-outputs` + "`" + `, or ` + "`" + `write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sentinel_mocks",
					Description: `(Required) The permission to grant the team on the workspace's generated Sentinel mocks, Valid values are ` + "`" + `none` + "`" + ` or ` + "`" + `read` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workspace_locking",
					Description: `(Required) Boolean determining whether or not to grant the team permission to manually lock/unlock the workspace. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team_member",
			Category:         "Resources",
			ShortDescription: `Add or remove a user from a team.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) ID of the team.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Name of the user to add. ## Import A team member can be imported; use ` + "`" + `<TEAM ID>/<USERNAME>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_team_member.test team-47qC3LmA47piVan7/sander ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team_members",
			Category:         "Resources",
			ShortDescription: `Manages users in a team.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"members",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) ID of the team.`,
				},
				resource.Attribute{
					Name:        "usernames",
					Description: `(Required) Names of the users to add. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Team members can be imported; use ` + "`" + `<TEAM ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_team_members.test team-47qC3LmA47piVan7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Team members can be imported; use ` + "`" + `<TEAM ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_team_members.test team-47qC3LmA47piVan7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team_organization_member",
			Category:         "Resources",
			ShortDescription: `Add or remove a user from a team.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"organization",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) ID of the team.`,
				},
				resource.Attribute{
					Name:        "organization_membership_id",
					Description: `(Required) ID of the organization membership. ## Import A team member can be imported; use ` + "`" + `<TEAM ID>/<ORGANIZATION MEMBERSHIP ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_team_organization_member.test team-47qC3LmA47piVan7/ou-2342390sdf0jj ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_team_token",
			Category:         "Resources",
			ShortDescription: `Generates a new team token and overrides existing token if one exists.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) ID of the team.`,
				},
				resource.Attribute{
					Name:        "force_regenerate",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, a new token will be generated even if a token already exists. This will invalidate the existing token! ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The generated token. ## Import Team tokens can be imported; use ` + "`" + `<TEAM ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_team_token.test team-47qC3LmA47piVan7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the token.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `The generated token. ## Import Team tokens can be imported; use ` + "`" + `<TEAM ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_team_token.test team-47qC3LmA47piVan7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_variable",
			Category:         "Resources",
			ShortDescription: `Manages variables.`,
			Description:      ``,
			Keywords: []string{
				"variable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Name of the variable.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the variable.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Required) Whether this is a Terraform or environment variable. Valid values are ` + "`" + `terraform` + "`" + ` or ` + "`" + `env` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the variable.`,
				},
				resource.Attribute{
					Name:        "hcl",
					Description: `(Optional) Whether to evaluate the value of the variable as a string of HCL code. Has no effect for environment variables. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sensitive",
					Description: `(Optional) Whether the value is sensitive. If true then the variable is written once and not visible thereafter. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workspace_id",
					Description: `(Required) ID of the workspace that owns the variable. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the variable. ## Import Variables can be imported; use ` + "`" + `<ORGANIZATION NAME>/<WORKSPACE NAME>/<VARIABLE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_variable.test my-org-name/my-workspace-name/var-5rTwnSaRPogw6apb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the variable. ## Import Variables can be imported; use ` + "`" + `<ORGANIZATION NAME>/<WORKSPACE NAME>/<VARIABLE ID>` + "`" + ` as the import ID. For example: ` + "`" + `` + "`" + `` + "`" + `shell terraform import tfe_variable.test my-org-name/my-workspace-name/var-5rTwnSaRPogw6apb ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "tfe_workspace",
			Category:         "Resources",
			ShortDescription: `Manages workspaces.`,
			Description:      ``,
			Keywords: []string{
				"workspace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the workspace.`,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `(Required) Name of the organization.`,
				},
				resource.Attribute{
					Name:        "auto_apply",
					Description: `(Optional) Whether to automatically apply changes when a Terraform plan is successful. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "file_triggers_enabled",
					Description: `(Optional) Whether to filter runs based on the changed files in a VCS push. If enabled, the working directory and trigger prefixes describe a set of paths which must contain changes for a VCS push to trigger a run. If disabled, any push will trigger a run. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "operations",
					Description: `(Optional) Whether to use remote execution mode. When set to ` + "`" + `false` + "`" + `, the workspace will be used for state storage only. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "queue_all_runs",
					Description: `(Optional) Whether all runs should be queued. When set to ` + "`" + `false` + "`" + `, runs triggered by a VCS change will not be queued until at least one run is manually queued. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "speculative_enabled",
					Description: `(Optional) Whether this workspace allows speculative plans. Setting this to ` + "`" + `false` + "`" + ` prevents Terraform Cloud or the Terraform Enterprise instance from running plans on pull requests, which can improve security if the VCS repository is public or includes untrusted contributors. Defaults to ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssh_key_id",
					Description: `(Optional) The ID of an SSH key to assign to the workspace.`,
				},
				resource.Attribute{
					Name:        "terraform_version",
					Description: `(Optional) The version of Terraform to use for this workspace. Defaults to the latest available version.`,
				},
				resource.Attribute{
					Name:        "trigger_prefixes",
					Description: `(Optional) List of repository-root-relative paths which describe all locations to be tracked for changes.`,
				},
				resource.Attribute{
					Name:        "working_directory",
					Description: `(Optional) A relative path that Terraform will execute within. Defaults to the root of your repository.`,
				},
				resource.Attribute{
					Name:        "vcs_repo",
					Description: `(Optional) Settings for the workspace's VCS repository. The ` + "`" + `vcs_repo` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `(Required) A reference to your VCS repository in the format ` + "`" + `:org/:repo` + "`" + ` where ` + "`" + `:org` + "`" + ` and ` + "`" + `:repo` + "`" + ` refer to the organization and repository in your VCS provider.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `(Optional) The repository branch that Terraform will execute from. Default to ` + "`" + `master` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ingress_submodules",
					Description: `(Optional) Whether submodules should be fetched when cloning the VCS repository. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oauth_token_id",
					Description: `(Required) Token ID of the VCS Connection (OAuth Connection Token) to use. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The workspace ID. ## Import ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The workspace ID. ## Import ~>`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"tfe_notification_configuration": 0,
		"tfe_oauth_client":               1,
		"tfe_organization":               2,
		"tfe_organization_membership":    3,
		"tfe_organization_token":         4,
		"tfe_policy_set":                 5,
		"tfe_policy_set_parameter":       6,
		"tfe_registry_module":            7,
		"tfe_run_trigger":                8,
		"tfe_sentinel_policy":            9,
		"tfe_ssh_key":                    10,
		"tfe_team":                       11,
		"tfe_team_access":                12,
		"tfe_team_member":                13,
		"tfe_team_members":               14,
		"tfe_team_organization_member":   15,
		"tfe_team_token":                 16,
		"tfe_variable":                   17,
		"tfe_workspace":                  18,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
