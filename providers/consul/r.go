package consul

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "consul_acl_auth_method",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create an ACL auth method`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"auth",
				"method",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL auth method.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the ACL auth method.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) An optional name to use instead of the name attribute when displaying information about this auth method.`,
				},
				resource.Attribute{
					Name:        "max_token_ttl",
					Description: `(Optional) The maximum life of any token created by this auth method.`,
				},
				resource.Attribute{
					Name:        "token_locality",
					Description: `(Optional) The kind of token that this auth method produces. This can be either 'local' or 'global'.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A free form human readable description of the auth method.`,
				},
				resource.Attribute{
					Name:        "config_json",
					Description: `(Required) The raw configuration for this ACL auth method.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) The raw configuration for this ACL auth method. This attribute is deprecated and will be removed in a future version. ` + "`" + `config_json` + "`" + ` should be used instead.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the policy within.`,
				},
				resource.Attribute{
					Name:        "namespace_rule",
					Description: `(Optional, Enterprise Only) A set of rules that control which namespace tokens created via this auth method will be created within. Each ` + "`" + `namespace_rule` + "`" + ` can have the following attributes:`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) Specifies the expression used to match this namespace rule against valid identities returned from an auth method validation. Defaults to ` + "`" + `""` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bind_namespace",
					Description: `(Required) If the namespace rule's ` + "`" + `selector` + "`" + ` matches then this is used to control the namespace where the token is created. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the the auth method.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ACL auth method.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the ACL auth method.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `An optional name to use instead of the name attribute when displaying information about this auth method.`,
				},
				resource.Attribute{
					Name:        "max_token_ttl",
					Description: `The maximum life of any token created by this auth method.`,
				},
				resource.Attribute{
					Name:        "token_locality",
					Description: `The kind of token that this auth method produces. This can be either 'local' or 'global'.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free form human readable description of the auth method.`,
				},
				resource.Attribute{
					Name:        "config_json",
					Description: `The raw configuration for this ACL auth method.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The raw configuration for this ACL auth method. This attribute is deprecated and will be removed in a future version. If the configuration is too complex to be represented as a map of strings it will be blank. ` + "`" + `config_json` + "`" + ` should be used instead.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Enterprise Only) The namespace to create the policy within.`,
				},
				resource.Attribute{
					Name:        "namespace_rule",
					Description: `(Enterprise Only) A set of rules that control which namespace tokens created via this auth method will be created within.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the the auth method.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ACL auth method.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the ACL auth method.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `An optional name to use instead of the name attribute when displaying information about this auth method.`,
				},
				resource.Attribute{
					Name:        "max_token_ttl",
					Description: `The maximum life of any token created by this auth method.`,
				},
				resource.Attribute{
					Name:        "token_locality",
					Description: `The kind of token that this auth method produces. This can be either 'local' or 'global'.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free form human readable description of the auth method.`,
				},
				resource.Attribute{
					Name:        "config_json",
					Description: `The raw configuration for this ACL auth method.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The raw configuration for this ACL auth method. This attribute is deprecated and will be removed in a future version. If the configuration is too complex to be represented as a map of strings it will be blank. ` + "`" + `config_json` + "`" + ` should be used instead.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Enterprise Only) The namespace to create the policy within.`,
				},
				resource.Attribute{
					Name:        "namespace_rule",
					Description: `(Enterprise Only) A set of rules that control which namespace tokens created via this auth method will be created within.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_acl_binding_rule",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create an ACL binding rule`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"binding",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_method",
					Description: `(Required) The name of the ACL auth method this rule apply.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A free form human readable description of the binding rule.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `(Optional) The expression used to math this rule against valid identities returned from an auth method validation.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `(Required) Specifies the way the binding rule affects a token created at login.`,
				},
				resource.Attribute{
					Name:        "bind_name",
					Description: `(Required) The name to bind to a token at login-time.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the binding rule within. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the the binding rule.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The name of the ACL auth method this rule apply.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free form human readable description of the binding rule.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `The expression used to math this rule against valid identities returned from an auth method validation.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `Specifies the way the binding rule affects a token created at login.`,
				},
				resource.Attribute{
					Name:        "bind_name",
					Description: `The name to bind to a token at login-time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the the binding rule.`,
				},
				resource.Attribute{
					Name:        "auth_method",
					Description: `The name of the ACL auth method this rule apply.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free form human readable description of the binding rule.`,
				},
				resource.Attribute{
					Name:        "selector",
					Description: `The expression used to math this rule against valid identities returned from an auth method validation.`,
				},
				resource.Attribute{
					Name:        "bind_type",
					Description: `Specifies the way the binding rule affects a token created at login.`,
				},
				resource.Attribute{
					Name:        "bind_name",
					Description: `The name to bind to a token at login-time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_acl_policy",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create an ACL policy`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `(Required) The rules of the policy.`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional) The datacenters of the policy.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the policy within. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `The rules of the policy.`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `The datacenters of the policy. ## Import ` + "`" + `consul_acl_policy` + "`" + ` can be imported: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_policy.my-policy 1c90ef03-a6dd-6a8c-ac49-042ad3752896 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `The rules of the policy.`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `The datacenters of the policy. ## Import ` + "`" + `consul_acl_policy` + "`" + ` can be imported: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_policy.my-policy 1c90ef03-a6dd-6a8c-ac49-042ad3752896 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_acl_role",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create an ACL role`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A free form human readable description of the role.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) The list of policies that should be applied to the role.`,
				},
				resource.Attribute{
					Name:        "service_identities",
					Description: `(Optional) The list of service identities that should be applied to the role.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the role within. The ` + "`" + `service_identities` + "`" + ` supports:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The name of the service.`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional) The datacenters the effective policy is valid within. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ACL role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free form human readable description of the role.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of policies that should be applied to the role.`,
				},
				resource.Attribute{
					Name:        "service_identities",
					Description: `The list of service identities that should be applied to the role. ## Import ` + "`" + `consul_acl_role` + "`" + ` can be imported: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_role.read 816a195f-6cb1-2e8d-92af-3011ae706318 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the role.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ACL role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A free form human readable description of the role.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of policies that should be applied to the role.`,
				},
				resource.Attribute{
					Name:        "service_identities",
					Description: `The list of service identities that should be applied to the role. ## Import ` + "`" + `consul_acl_role` + "`" + ` can be imported: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_role.read 816a195f-6cb1-2e8d-92af-3011ae706318 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_acl_token",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create an ACL token`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"token",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accessor_id",
					Description: `(Optional) The uuid of the token. If omitted, Consul will generate a random uuid.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) The list of policies attached to the token.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional) The list of roles attached to the token.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) The flag to set the token local to the current datacenter.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the token within. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The token accessor ID.`,
				},
				resource.Attribute{
					Name:        "accessor_id",
					Description: `The token accessor ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of policies attached to the token.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `The list of roles attached to the token.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `The flag to set the token local to the current datacenter. ## Import ` + "`" + `consul_acl_token` + "`" + ` can be imported. This is especially useful to manage the anonymous and the master token with Terraform: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_token.anonymous 00000000-0000-0000-0000-000000000002 $ terraform import consul_acl_token.master-token 624d94ca-bc5c-f960-4e83-0a609cf588be ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The token accessor ID.`,
				},
				resource.Attribute{
					Name:        "accessor_id",
					Description: `The token accessor ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of policies attached to the token.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `The list of roles attached to the token.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `The flag to set the token local to the current datacenter. ## Import ` + "`" + `consul_acl_token` + "`" + ` can be imported. This is especially useful to manage the anonymous and the master token with Terraform: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_token.anonymous 00000000-0000-0000-0000-000000000002 $ terraform import consul_acl_token.master-token 624d94ca-bc5c-f960-4e83-0a609cf588be ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_acl_token_policy_attachment",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to create a link between an ACL token and a policy`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"token",
				"policy",
				"attachment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "token_id",
					Description: `(Required) The id of the token.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `(Required) The name of the policy attached to the token. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `The id of the token.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The name of the policy attached to the token. ## Import ` + "`" + `consul_acl_token_policy_attachment` + "`" + ` can be imported. This is especially useful to manage the policies attached to the anonymous and the master tokens with Terraform: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_token_policy_attachment.anonymous 00000000-0000-0000-0000-000000000002:policy_name $ terraform import consul_acl_token_policy_attachment.master-token 624d94ca-bc5c-f960-4e83-0a609cf588be:policy_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The attachment ID.`,
				},
				resource.Attribute{
					Name:        "token_id",
					Description: `The id of the token.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The name of the policy attached to the token. ## Import ` + "`" + `consul_acl_token_policy_attachment` + "`" + ` can be imported. This is especially useful to manage the policies attached to the anonymous and the master tokens with Terraform: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_token_policy_attachment.anonymous 00000000-0000-0000-0000-000000000002:policy_name $ terraform import consul_acl_token_policy_attachment.master-token 624d94ca-bc5c-f960-4e83-0a609cf588be:policy_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_agent_service",
			Category:         "Resources",
			ShortDescription: `Provides access to Agent Service data in Consul. This can be used to define a service associated with a particular agent. Currently, defining health checks for an agent service is not supported.`,
			Description:      ``,
			Keywords: []string{
				"agent",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The address of the service. Defaults to the address of the agent.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of values that are opaque to Consul, but can be used to distinguish between services or nodes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service, defaults to the value of ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service, defaults to the value of ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_autopilot_config",
			Category:         "Resources",
			ShortDescription: `Provides access to the Autopilot Configuration of Consul.`,
			Description:      ``,
			Keywords: []string{
				"autopilot",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "cleanup_dead_servers",
					Description: `(Optional) Whether to remove failing servers when a replacement comes online. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "last_contact_threshold",
					Description: `(Optional) The time after which a server is considered as unhealthy and will be removed. Defaults to ` + "`" + `"200ms"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_trailing_logs",
					Description: `(Optional) The maximum number of Raft log entries a server can trail the leader. Defaults to 250.`,
				},
				resource.Attribute{
					Name:        "server_stabilization_time",
					Description: `(Optional) The period to wait for a server to be healthy and stable before being promoted to a full, voting member. Defaults to ` + "`" + `"10s"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "redundancy_zone_tag",
					Description: `(Optional) The [redundancy zone](https://www.consul.io/docs/guides/autopilot.html#redundancy-zones) tag to use. Consul will try to keep one voting server by zone to take advantage of isolated failure domains. Defaults to an empty string.`,
				},
				resource.Attribute{
					Name:        "disable_upgrade_migration",
					Description: `(Optional) Whether to disable [upgrade migrations](https://www.consul.io/docs/guides/autopilot.html#redundancy-zones). Defaults to false.`,
				},
				resource.Attribute{
					Name:        "upgrade_version_tag",
					Description: `(Optional) The tag to override the version information used during a migration. Defaults to an empty string. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter used.`,
				},
				resource.Attribute{
					Name:        "cleanup_dead_servers",
					Description: `Whether to remove failing servers.`,
				},
				resource.Attribute{
					Name:        "last_contact_threshold",
					Description: `The time after which a server is considered as unhealthy and will be removed.`,
				},
				resource.Attribute{
					Name:        "max_trailing_logs",
					Description: `The maximum number of Raft log entries a server can trail the leader.`,
				},
				resource.Attribute{
					Name:        "server_stabilization_time",
					Description: `The period to wait for a server to be healthy and stable before being promoted to a full, voting member.`,
				},
				resource.Attribute{
					Name:        "redundancy_zone_tag",
					Description: `The [redundancy zone](https://www.consul.io/docs/guides/autopilot.html#redundancy-zones) tag used. Consul will try to keep one voting server by zone to take advantage of isolated failure domains.`,
				},
				resource.Attribute{
					Name:        "disable_upgrade_migration",
					Description: `Whether to disable [upgrade migrations](https://www.consul.io/docs/guides/autopilot.html#redundancy-zones).`,
				},
				resource.Attribute{
					Name:        "upgrade_version_tag",
					Description: `The tag to override the version information used during a migration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter used.`,
				},
				resource.Attribute{
					Name:        "cleanup_dead_servers",
					Description: `Whether to remove failing servers.`,
				},
				resource.Attribute{
					Name:        "last_contact_threshold",
					Description: `The time after which a server is considered as unhealthy and will be removed.`,
				},
				resource.Attribute{
					Name:        "max_trailing_logs",
					Description: `The maximum number of Raft log entries a server can trail the leader.`,
				},
				resource.Attribute{
					Name:        "server_stabilization_time",
					Description: `The period to wait for a server to be healthy and stable before being promoted to a full, voting member.`,
				},
				resource.Attribute{
					Name:        "redundancy_zone_tag",
					Description: `The [redundancy zone](https://www.consul.io/docs/guides/autopilot.html#redundancy-zones) tag used. Consul will try to keep one voting server by zone to take advantage of isolated failure domains.`,
				},
				resource.Attribute{
					Name:        "disable_upgrade_migration",
					Description: `Whether to disable [upgrade migrations](https://www.consul.io/docs/guides/autopilot.html#redundancy-zones).`,
				},
				resource.Attribute{
					Name:        "upgrade_version_tag",
					Description: `The tag to override the version information used during a migration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_catalog_entry",
			Category:         "Resources",
			ShortDescription: `Registers a node or service with the Consul Catalog. Currently, defining health checks is not supported.`,
			Description:      ``,
			Keywords: []string{
				"catalog",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The address of the node being added to, or referenced in the catalog.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Required) The name of the node being added to, or referenced in the catalog.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Optional) A service to optionally associated with the node. Supported values are documented below.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) ACL token. The ` + "`" + `service` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional) The address of the service. Defaults to the node address.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the service. Defaults to the ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the service`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) The port of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of values that are opaque to Consul, but can be used to distinguish between services or nodes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `The ID of the service, defaults to the value of ` + "`" + `name` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `The ID of the service, defaults to the value of ` + "`" + `name` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_certificate_authority",
			Category:         "Resources",
			ShortDescription: `A resource that manage the Consul Connect Certificate Authority`,
			Description:      ``,
			Keywords: []string{
				"certificate",
				"authority",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "connect_provider",
					Description: `(Required, string) Specifies the CA provider type to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `(Required, map) The raw configuration to use for the chosen provider. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "connect_provider",
					Description: `Specifies the CA provider type to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The raw configuration to use for the chosen provider. ## Import ` + "`" + `certificate_authority` + "`" + ` can be imported: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import certificate_authority.connect connect-ca ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "connect_provider",
					Description: `Specifies the CA provider type to use.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The raw configuration to use for the chosen provider. ## Import ` + "`" + `certificate_authority` + "`" + ` can be imported: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import certificate_authority.connect connect-ca ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_config_entry",
			Category:         "Resources",
			ShortDescription: `Registers a configuration entry in Consul.`,
			Description:      ``,
			Keywords: []string{
				"config",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "kind",
					Description: `(Required) The kind of configuration entry to register.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the configuration entry being registred.`,
				},
				resource.Attribute{
					Name:        "config_json",
					Description: `(Optional) An arbitrary map of configuration values. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the configuration entry.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of the configuration entry.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the configuration entry.`,
				},
				resource.Attribute{
					Name:        "config_json",
					Description: `A map of configuration values.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the configuration entry.`,
				},
				resource.Attribute{
					Name:        "kind",
					Description: `The kind of the configuration entry.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the configuration entry.`,
				},
				resource.Attribute{
					Name:        "config_json",
					Description: `A map of configuration values.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_intention",
			Category:         "Resources",
			ShortDescription: `A resource that can create intentions for Consul Connect.`,
			Description:      ``,
			Keywords: []string{
				"intention",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_name",
					Description: `(Required, string) The name of the source service for the intention. This service does not have to exist.`,
				},
				resource.Attribute{
					Name:        "source_namespace",
					Description: `(Optional, Enterprise Only) The source namespace of the intention.`,
				},
				resource.Attribute{
					Name:        "destination_name",
					Description: `(Required, string) The name of the destination service for the intention. This service does not have to exist.`,
				},
				resource.Attribute{
					Name:        "destination_namespace",
					Description: `(Optional, Enterprise Only) The destination namespace of the intention.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Required, string) The intention action. Must be one of ` + "`" + `allow` + "`" + ` or ` + "`" + `deny` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional, map) Key/value pairs that are opaque to Consul and are associated with the intention.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, string) Optional description that can be used by Consul tooling, but is not used internally.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the intention.`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `The source for the intention.`,
				},
				resource.Attribute{
					Name:        "destination_name",
					Description: `The destination for the intention.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the intention.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Key/value pairs associated with the intention.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the intention.`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `The source for the intention.`,
				},
				resource.Attribute{
					Name:        "destination_name",
					Description: `The destination for the intention.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the intention.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Key/value pairs associated with the intention.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_key_prefix",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to manage a namespace of Consul keys that share a common name prefix.`,
			Description:      ``,
			Keywords: []string{
				"key",
				"prefix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The ACL token to use. This overrides the token that the agent provides by default.`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `(Required) Specifies the common prefix shared by all keys that will be managed by this resource instance. In most cases this will end with a slash, to manage a "folder" of keys.`,
				},
				resource.Attribute{
					Name:        "subkeys",
					Description: `(Optional) A mapping from subkey name (which will be appended to the given ` + "`" + `path_prefix` + "`" + `) to the value that should be stored at that key. Use slashes, as shown in the above example, to create "sub-folders" under the given path prefix.`,
				},
				resource.Attribute{
					Name:        "subkey",
					Description: `(Optional) A subkey to add. Supported values documented below. Multiple blocks supported.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the keys within. The ` + "`" + `subkey` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) This is the path (which will be appended to the given ` + "`" + `path_prefix` + "`" + `) in Consul that should be written to.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to write to the given path.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Optional) An [unsigned integer value](https://www.consul.io/api/kv.html#flags-1) to attach to the key (defaults to 0). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read/written to. ## Import ` + "`" + `consul_key_prefix` + "`" + ` can be imported. This is useful when the path already and you know all keys in path must be removed. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_key_prefix.myapp_config myapp/config/ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read/written to. ## Import ` + "`" + `consul_key_prefix` + "`" + ` can be imported. This is useful when the path already and you know all keys in path must be removed. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_key_prefix.myapp_config myapp/config/ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_keys",
			Category:         "Resources",
			ShortDescription: `Writes values into the Consul key/value store.`,
			Description:      ``,
			Keywords: []string{
				"keys",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The ACL token to use. This overrides the token that the agent provides by default.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Specifies a key in Consul to be written. Supported values documented below.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the keys within. The ` + "`" + `key` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) This is the path in Consul that should be written to.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to write to the given path.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `(Optional) An [unsigned integer value](https://www.consul.io/api/kv.html#flags-1) to attach to the key (defaults to 0).`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Optional) If true, then the key will be deleted when either its configuration block is removed from the configuration or the entire resource is destroyed. Otherwise, it will be left in Consul. Defaults to false. ### Deprecated ` + "`" + `key` + "`" + ` arguments Prior to Terraform 0.7, this resource was used both to read`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being written to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being written to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_license",
			Category:         "Resources",
			ShortDescription: `Manage the Consul Enterprise license.`,
			Description:      ``,
			Keywords: []string{
				"license",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "license",
					Description: `(Required) The Consul license to use. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Whether the license is valid.`,
				},
				resource.Attribute{
					Name:        "license_id",
					Description: `The ID of the license used.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer the license is attached to.`,
				},
				resource.Attribute{
					Name:        "installation_id",
					Description: `The ID of the current installation.`,
				},
				resource.Attribute{
					Name:        "issue_time",
					Description: `The date the license was issued.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start time of the license.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `The expiration time of the license.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `The product for which the license is valid.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `The metadata attached to the license.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features for which the license is valid.`,
				},
				resource.Attribute{
					Name:        "warnings",
					Description: `A list of warning messages regarding the license validity.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "valid",
					Description: `Whether the license is valid.`,
				},
				resource.Attribute{
					Name:        "license_id",
					Description: `The ID of the license used.`,
				},
				resource.Attribute{
					Name:        "customer_id",
					Description: `The ID of the customer the license is attached to.`,
				},
				resource.Attribute{
					Name:        "installation_id",
					Description: `The ID of the current installation.`,
				},
				resource.Attribute{
					Name:        "issue_time",
					Description: `The date the license was issued.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start time of the license.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `The expiration time of the license.`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `The product for which the license is valid.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `The metadata attached to the license.`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `The features for which the license is valid.`,
				},
				resource.Attribute{
					Name:        "warnings",
					Description: `A list of warning messages regarding the license validity.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_namespace",
			Category:         "Resources",
			ShortDescription: `Manage a Consul namespace.`,
			Description:      ``,
			Keywords: []string{
				"namespace",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The namespace name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Free form namespace description.`,
				},
				resource.Attribute{
					Name:        "policy_defaults",
					Description: `(Optional) The list of default policies that should be applied to all tokens created in this namespace.`,
				},
				resource.Attribute{
					Name:        "role_defaults",
					Description: `(Optional) The list of default roles that should be applied to all tokens created in this namespace.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional) Specifies arbitrary KV metadata to associate with the namespace. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The namespace name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The namespace description.`,
				},
				resource.Attribute{
					Name:        "policy_defaults",
					Description: `The list of default policies that will be applied to all tokens created in this namespace.`,
				},
				resource.Attribute{
					Name:        "role_defaults",
					Description: `The list of default roles that will be applied to all tokens created in this namespace.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Arbitrary KV metadata associated with the namespace.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The namespace name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The namespace description.`,
				},
				resource.Attribute{
					Name:        "policy_defaults",
					Description: `The list of default policies that will be applied to all tokens created in this namespace.`,
				},
				resource.Attribute{
					Name:        "role_defaults",
					Description: `The list of default roles that will be applied to all tokens created in this namespace.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Arbitrary KV metadata associated with the namespace.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_network_area",
			Category:         "Resources",
			ShortDescription: `Manage Network Areas.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"area",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The ACL token to use. This overrides the token that the agent provides by default.`,
				},
				resource.Attribute{
					Name:        "peer_datacenter",
					Description: `(Required) The name of the Consul datacenter that will be joined to form the area.`,
				},
				resource.Attribute{
					Name:        "retry_join",
					Description: `(Optional) Specifies a list of Consul servers to attempt to join. Servers can be given as ` + "`" + `IP` + "`" + `, ` + "`" + `IP:port` + "`" + `, ` + "`" + `hostname` + "`" + `, or ` + "`" + `hostname:port` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `(Optional) Specifies whether gossip over this area should be encrypted with TLS if possible. Defaults to ` + "`" + `false` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter being queried.`,
				},
				resource.Attribute{
					Name:        "peer_datacenter",
					Description: `The name of the Consul datacenter joined to form the area.`,
				},
				resource.Attribute{
					Name:        "retry_join",
					Description: `The list of Consul servers Consul attempts to join.`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `Whether the gossip over this area should be encrypted with TLS.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter being queried.`,
				},
				resource.Attribute{
					Name:        "peer_datacenter",
					Description: `The name of the Consul datacenter joined to form the area.`,
				},
				resource.Attribute{
					Name:        "retry_join",
					Description: `The list of Consul servers Consul attempts to join.`,
				},
				resource.Attribute{
					Name:        "use_tls",
					Description: `Whether the gossip over this area should be encrypted with TLS.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_node",
			Category:         "Resources",
			ShortDescription: `Provides access to Node data in Consul. This can be used to define a node.`,
			Description:      ``,
			Keywords: []string{
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The address of the node being added to, or referenced in the catalog.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the node being added to, or referenced in the catalog.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional, map) Key/value pairs that are associated with the node. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional, map) Key/value pairs that are associated with the node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional, map) Key/value pairs that are associated with the node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_prepared_query",
			Category:         "Resources",
			ShortDescription: `Allows Terraform to manage a Consul prepared query`,
			Description:      ``,
			Keywords: []string{
				"prepared",
				"query",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) The ACL token to use when saving the prepared query. This overrides the token that the agent provides by default.`,
				},
				resource.Attribute{
					Name:        "stored_token",
					Description: `(Optional) The ACL token to store with the prepared query. This token will be used by default whenever the query is executed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the prepared query. Used to identify the prepared query during requests. Can be specified as an empty string to configure the query as a catch-all.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) The name of the service to query.`,
				},
				resource.Attribute{
					Name:        "session",
					Description: `(Optional) The name of the Consul session to tie this query's lifetime to. This is an advanced parameter that should not be used without a complete understanding of Consul sessions and the implications of their use (it is recommended to leave this blank in nearly all cases). If this parameter is omitted the query will not expire.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The list of required and/or disallowed tags. If a tag is in this list it must be present. If the tag is preceded with a "!" then it is disallowed.`,
				},
				resource.Attribute{
					Name:        "only_passing",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the prepared query will only return nodes with passing health checks in the result.`,
				},
				resource.Attribute{
					Name:        "near",
					Description: `(Optional) Allows specifying the name of a node to sort results near using Consul's distance sorting and network coordinates. The magic ` + "`" + `_agent` + "`" + ` value can be used to always sort nearest the node servicing the request.`,
				},
				resource.Attribute{
					Name:        "ignore_check_ids",
					Description: `(Optional) Specifies a list of check IDs that should be ignored when filtering unhealthy instances. This is mostly useful in an emergency or as a temporary measure when a health check is found to be unreliable. Being able to ignore it in centrally-defined queries can be simpler than de-registering the check as an interim solution until the check can be fixed.`,
				},
				resource.Attribute{
					Name:        "node_meta",
					Description: `(Optional) Specifies a list of user-defined key/value pairs that will be used for filtering the query results to nodes with the given metadata values present.`,
				},
				resource.Attribute{
					Name:        "service_meta",
					Description: `(Optional) Specifies a list of user-defined key/value pairs that will be used for filtering the query results to services with the given metadata values present.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `(Optional) Options for controlling behavior when no healthy nodes are available in the local DC.`,
				},
				resource.Attribute{
					Name:        "nearest_n",
					Description: `(Optional) Return results from this many datacenters, sorted in ascending order of estimated RTT.`,
				},
				resource.Attribute{
					Name:        "datacenters",
					Description: `(Optional) Remote datacenters to return results from.`,
				},
				resource.Attribute{
					Name:        "dns",
					Description: `(Optional) Settings for controlling the DNS response details.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL to send when returning DNS results.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Optional) Query templating options. This is used to make a single prepared query respond to many different requests.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of template matching to perform. Currently only ` + "`" + `name_prefix_match` + "`" + ` is supported.`,
				},
				resource.Attribute{
					Name:        "regexp",
					Description: `(Required) The regular expression to match with. When using ` + "`" + `name_prefix_match` + "`" + `, this regex is applied against the query name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the prepared query, generated by Consul. ## Import ` + "`" + `consul_prepared_query` + "`" + ` can be imported with the query's ID in the Consul HTTP API. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_prepared_query.my_service 71ecfb82-717a-4258-b4b6-2fb75144d856 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the prepared query, generated by Consul. ## Import ` + "`" + `consul_prepared_query` + "`" + ` can be imported with the query's ID in the Consul HTTP API. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_prepared_query.my_service 71ecfb82-717a-4258-b4b6-2fb75144d856 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_service",
			Category:         "Resources",
			ShortDescription: `A high-level resource for creating a Service in Consul in the Consul catalog.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the service.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `(Required, string) The name of the node the to register the service on.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Optional, string) The address of the service. Defaults to the address of the node.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional, int) The port of the service.`,
				},
				resource.Attribute{
					Name:        "checks",
					Description: `(Optional, list of checks) Health-checks to register to monitor the service. The list of attributes for each health-check is detailled below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional, set of strings) A list of values that are opaque to Consul, but can be used to distinguish between services or nodes.`,
				},
				resource.Attribute{
					Name:        "enable_tag_override",
					Description: `(Optional, boolean) Specifies to disable the anti-entropy feature for this service's tags. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `(Optional) A map of arbitrary KV metadata linked to the service instance.`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `(Optional, Enterprise Only) The namespace to create the service within. The following attributes are available for each health-check:`,
				},
				resource.Attribute{
					Name:        "check_id",
					Description: `(Optional, string) An ID,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the health-check.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional, string) An opaque field meant to hold human readable text.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, string) The initial health-check status.`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `(Optional, string) The TCP address and port to connect to for a TCP check.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `(Optional, string) The HTTP endpoint to call for an HTTP check.`,
				},
				resource.Attribute{
					Name:        "header",
					Description: `(Optional, set of headers) The headers to send for an HTTP check. The attributes of each header is given below.`,
				},
				resource.Attribute{
					Name:        "tls_skip_verify",
					Description: `(Optional, boolean) Whether to deactivate certificate verification for HTTP health-checks. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional, string) The method to use for HTTP health-checks. Defaults to ` + "`" + `GET` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `(Required, string) The interval to wait between each health-check invocation.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Required, string) The timeout value for HTTP checks.`,
				},
				resource.Attribute{
					Name:        "deregister_critical_service_after",
					Description: `(Optional, string) The time after which the service is automatically deregistered when in the ` + "`" + `critical` + "`" + ` state. Defaults to ` + "`" + `30s` + "`" + `. Each ` + "`" + `header` + "`" + ` must have the following attributes:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, string) The name of the header.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required, list of strings) The header's list of values. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `The node the service is registered on.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the service.`,
				},
				resource.Attribute{
					Name:        "checks",
					Description: `The list of health-checks associated with the service.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter of the service.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `A map of arbitrary KV metadata linked to the service instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the service.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `The node the service is registered on.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the service.`,
				},
				resource.Attribute{
					Name:        "checks",
					Description: `The list of health-checks associated with the service.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter of the service.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `A map of arbitrary KV metadata linked to the service instance.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"consul_acl_auth_method":             0,
		"consul_acl_binding_rule":            1,
		"consul_acl_policy":                  2,
		"consul_acl_role":                    3,
		"consul_acl_token":                   4,
		"consul_acl_token_policy_attachment": 5,
		"consul_agent_service":               6,
		"consul_autopilot_config":            7,
		"consul_catalog_entry":               8,
		"consul_certificate_authority":       9,
		"consul_config_entry":                10,
		"consul_intention":                   11,
		"consul_key_prefix":                  12,
		"consul_keys":                        13,
		"consul_license":                     14,
		"consul_namespace":                   15,
		"consul_network_area":                16,
		"consul_node":                        17,
		"consul_prepared_query":              18,
		"consul_service":                     19,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
