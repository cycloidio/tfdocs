package consul

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
					Description: `(Optional) The datacenters of the policy. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "description",
					Description: `(Optional) The description of the token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `(Optional) The list of policies attached to the token.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `(Optional) The flag to set the token local to the current datacenter. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
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
					Name:        "description",
					Description: `The description of the token.`,
				},
				resource.Attribute{
					Name:        "policies",
					Description: `The list of policies attached to the token.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `The flag to set the token local to the current datacenter. ## Import ` + "`" + `consul_acl_token` + "`" + ` can be imported. This is especially useful to manage the anonymous and the master token with Terraform: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import consul_acl_token.anonymous 00000000-0000-0000-0000-000000000002 $ terraform import consul_acl_token.master-token 624d94ca-bc5c-f960-4e83-0a609cf588be ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "destination_name",
					Description: `(Required, string) The name of the destination service for the intention. This service does not have to exist.`,
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
					Description: `(Optional) A subkey to add. Supported values documented below. Multiple blocks supported. The ` + "`" + `subkey` + "`" + ` block supports the following:`,
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
					Description: `(Required) Specifies a key in Consul to be written. Supported values documented below. The ` + "`" + `key` + "`" + ` block supports the following:`,
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
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup. The following attributes are available for each health-check:`,
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
					Description: `(Optional, string) The initial health-check status. Defaults to ` + "`" + `critical` + "`" + `.`,
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
					Description: `(Required, string) The time after which the service is automatically deregistered when in the ` + "`" + `critical` + "`" + ` state. Each ` + "`" + `header` + "`" + ` must have the following attributes:`,
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
			},
		},
	}

	resourcesMap = map[string]int{

		"consul_acl_policy":       0,
		"consul_acl_token":        1,
		"consul_agent_service":    2,
		"consul_autopilot_config": 3,
		"consul_catalog_entry":    4,
		"consul_intention":        5,
		"consul_key_prefix":       6,
		"consul_keys":             7,
		"consul_node":             8,
		"consul_prepared_query":   9,
		"consul_service":          10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
