package consul

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "consul_agent_config",
			Category:         "Data Sources",
			ShortDescription: `Provides the configuration information of the local Consul agent.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the agent is running in`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The ID of the node the agent is running on`,
				},
				resource.Attribute{
					Name:        "node_name",
					Description: `The name of the node the agent is running on`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Boolean if the agent is a server or not`,
				},
				resource.Attribute{
					Name:        "revision",
					Description: `The first 9 characters of the VCS revision of the build of Consul that is running`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the build of Consul that is running`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_agent_self",
			Category:         "Data Sources",
			ShortDescription: `Provides the configuration information of the local Consul agent.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dns",
					Description: `A map of DNS configuration attributes. See below for details on the contents of the ` + "`" + `dns` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the Consul agent.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_autopilot_health",
			Category:         "Data Sources",
			ShortDescription: `Provides health information of the autopilot.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The datacenter to use. This overrides the agent's default datacenter and the datacenter in the provider setup. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `Whether all the servers in the cluster are currently healthy`,
				},
				resource.Attribute{
					Name:        "failure_tolerance",
					Description: `The number of redundant healthy servers that could fail without causing an outage`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of server health information. See below for details on the available information. ### Server health information`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Raft ID of the server`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The node name of the server`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the server`,
				},
				resource.Attribute{
					Name:        "serf_status",
					Description: `The status of the SerfHealth check of the server`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Consul version of the server`,
				},
				resource.Attribute{
					Name:        "leader",
					Description: `Whether the server is currently leader`,
				},
				resource.Attribute{
					Name:        "last_contact",
					Description: `The time elapsed since the server's last contact with the leader`,
				},
				resource.Attribute{
					Name:        "last_term",
					Description: `The server's last known Raft leader term`,
				},
				resource.Attribute{
					Name:        "last_index",
					Description: `The index of the server's last committed Raft log entry`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `Whether the server is healthy according to the current Autopilot configuration`,
				},
				resource.Attribute{
					Name:        "voter",
					Description: `Whether the server is a voting member of the Raft cluster`,
				},
				resource.Attribute{
					Name:        "stable_since",
					Description: `The time this server has been in its current ` + "`" + `` + "`" + `Healthy` + "`" + `` + "`" + ` state`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "healthy",
					Description: `Whether all the servers in the cluster are currently healthy`,
				},
				resource.Attribute{
					Name:        "failure_tolerance",
					Description: `The number of redundant healthy servers that could fail without causing an outage`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `A list of server health information. See below for details on the available information. ### Server health information`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Raft ID of the server`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The node name of the server`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the server`,
				},
				resource.Attribute{
					Name:        "serf_status",
					Description: `The status of the SerfHealth check of the server`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The Consul version of the server`,
				},
				resource.Attribute{
					Name:        "leader",
					Description: `Whether the server is currently leader`,
				},
				resource.Attribute{
					Name:        "last_contact",
					Description: `The time elapsed since the server's last contact with the leader`,
				},
				resource.Attribute{
					Name:        "last_term",
					Description: `The server's last known Raft leader term`,
				},
				resource.Attribute{
					Name:        "last_index",
					Description: `The index of the server's last committed Raft log entry`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `Whether the server is healthy according to the current Autopilot configuration`,
				},
				resource.Attribute{
					Name:        "voter",
					Description: `Whether the server is a voting member of the Raft cluster`,
				},
				resource.Attribute{
					Name:        "stable_since",
					Description: `The time this server has been in its current ` + "`" + `` + "`" + `Healthy` + "`" + `` + "`" + ` state`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_key_prefix",
			Category:         "Data Sources",
			ShortDescription: `Reads values from a "namespace" of Consul keys that share a common name prefix.`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) Specifies the common prefix shared by all keys that will be read by this data source instance. In most cases, this will end with a slash to read a "folder" of subkeys.`,
				},
				resource.Attribute{
					Name:        "subkey",
					Description: `(Optional) Specifies a subkey in Consul to be read. Supported values documented below. Multiple blocks supported. The ` + "`" + `subkey` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This is the name of the key. This value of the key is exposed as ` + "`" + `var.<name>` + "`" + `. This is not the path of the subkey in Consul.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) This is the subkey path in Consul (which will be appended to the given ` + "`" + `path_prefix` + "`" + `) to construct the full key that will be used to read the value.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) This is the default value to set for ` + "`" + `var.<name>` + "`" + ` if the key does not exist in Consul. Defaults to an empty string. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from.`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `the common prefix shared by all keys being read.`,
				},
				resource.Attribute{
					Name:        "var.<name>",
					Description: `For each name given, the corresponding attribute has the value of the key.`,
				},
				resource.Attribute{
					Name:        "subkeys",
					Description: `A map of the subkeys and values is set if no ` + "`" + `subkey` + "`" + ` block is provided.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from.`,
				},
				resource.Attribute{
					Name:        "path_prefix",
					Description: `the common prefix shared by all keys being read.`,
				},
				resource.Attribute{
					Name:        "var.<name>",
					Description: `For each name given, the corresponding attribute has the value of the key.`,
				},
				resource.Attribute{
					Name:        "subkeys",
					Description: `A map of the subkeys and values is set if no ` + "`" + `subkey` + "`" + ` block is provided.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_keys",
			Category:         "Data Sources",
			ShortDescription: `Reads values from the Consul key/value store.`,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) Specifies a key in Consul to be read. Supported values documented below. Multiple blocks supported. The ` + "`" + `key` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) This is the name of the key. This value of the key is exposed as ` + "`" + `var.<name>` + "`" + `. This is not the path of the key in Consul.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) This is the path in Consul that should be read or written to.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) This is the default value to set for ` + "`" + `var.<name>` + "`" + ` if the key does not exist in Consul. Defaults to an empty string. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from.`,
				},
				resource.Attribute{
					Name:        "var.<name>",
					Description: `For each name given, the corresponding attribute has the value of the key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from.`,
				},
				resource.Attribute{
					Name:        "var.<name>",
					Description: `For each name given, the corresponding attribute has the value of the key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_nodes",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of nodes in a given Consul datacenter.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The Consul datacenter to query. Defaults to the same value found in ` + "`" + `query_options` + "`" + ` parameter specified below, or if that is empty, the ` + "`" + `datacenter` + "`" + ` value found in the Consul agent that this provider is configured to talk to then the datacenter in the provider setup.`,
				},
				resource.Attribute{
					Name:        "query_options",
					Description: `(Optional) See below. The ` + "`" + `query_options` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_stale",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the default, allow responses from Consul servers that are followers.`,
				},
				resource.Attribute{
					Name:        "require_consistent",
					Description: `(Optional) When ` + "`" + `true` + "`" + ` force the client to perform a read on at least quorum servers and verify the result is the same. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Specify the Consul ACL token to use when performing the request. This defaults to the same API token configured by the ` + "`" + `consul` + "`" + ` provider but may be overriden if necessary.`,
				},
				resource.Attribute{
					Name:        "wait_index",
					Description: `(Optional) Index number used to enable blocking quereis.`,
				},
				resource.Attribute{
					Name:        "wait_time",
					Description: `(Optional) Max time the client should wait for a blocking query to return. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "node_ids",
					Description: `A list of the Consul node IDs.`,
				},
				resource.Attribute{
					Name:        "node_names",
					Description: `A list of the Consul node names.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes and details about each Consul agent. The list of per-node attributes is detailed below. The following is a list of the per-node attributes contained within the ` + "`" + `nodes` + "`" + ` map:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Node ID of the Consul agent.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "node_ids",
					Description: `A list of the Consul node IDs.`,
				},
				resource.Attribute{
					Name:        "node_names",
					Description: `A list of the Consul node names.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `A list of nodes and details about each Consul agent. The list of per-node attributes is detailed below. The following is a list of the per-node attributes contained within the ` + "`" + `nodes` + "`" + ` map:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Node ID of the Consul agent.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_service",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Consul service`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The Consul datacenter to query. Defaults to the same value found in ` + "`" + `query_options` + "`" + ` parameter specified below, or if that is empty, the ` + "`" + `datacenter` + "`" + ` value found in the Consul agent that this provider is configured to talk to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service name to select.`,
				},
				resource.Attribute{
					Name:        "query_options",
					Description: `(Optional) See below.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A single tag that can be used to filter the list of nodes to return based on a single matching tag.. The ` + "`" + `query_options` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_stale",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the default, allow responses from Consul servers that are followers.`,
				},
				resource.Attribute{
					Name:        "require_consistent",
					Description: `(Optional) When ` + "`" + `true` + "`" + ` force the client to perform a read on at least quorum servers and verify the result is the same. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Specify the Consul ACL token to use when performing the request. This defaults to the same API token configured by the ` + "`" + `consul` + "`" + ` provider but may be overriden if necessary.`,
				},
				resource.Attribute{
					Name:        "wait_index",
					Description: `(Optional) Index number used to enable blocking quereis.`,
				},
				resource.Attribute{
					Name:        "wait_time",
					Description: `(Optional) Max time the client should wait for a blocking query to return. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The name of the tag used to filter the list of nodes in ` + "`" + `service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `A list of nodes and details about each endpoint advertising a service. Each element in the list is a map of attributes that correspond to each individual node. The list of per-node attributes is detailed below. The following is a list of the per-node ` + "`" + `service` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The Node ID of the Consul agent advertising the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The name of the tag used to filter the list of nodes in ` + "`" + `service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `A list of nodes and details about each endpoint advertising a service. Each element in the list is a map of attributes that correspond to each individual node. The list of per-node attributes is detailed below. The following is a list of the per-node ` + "`" + `service` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `The Node ID of the Consul agent advertising the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_service_health",
			Category:         "Data Sources",
			ShortDescription: `Filter service instances based on health status`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The Consul datacenter to query.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service name to select.`,
				},
				resource.Attribute{
					Name:        "near",
					Description: `(Optional) Specifies a node name to sort the node list in ascending order based on the estimated round trip time from that node.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) A single tag that can be used to filter the list to return based on a single matching tag.`,
				},
				resource.Attribute{
					Name:        "node_meta",
					Description: `(Optional) Filter the results to nodes with the specified key/value pairs.`,
				},
				resource.Attribute{
					Name:        "passing",
					Description: `(Optional) Whether to return only nodes with all checks in the passing state. Defaults to ` + "`" + `true` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "near",
					Description: `The node to which the result must be sorted to.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The name of the tag used to filter the list.`,
				},
				resource.Attribute{
					Name:        "node_meta",
					Description: `The list of metadata to filter the nodes.`,
				},
				resource.Attribute{
					Name:        "passing",
					Description: `Whether to return only nodes with all checks in the passing state.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list of entries and details about each endpoint advertising a service. Each element in the list has three attributes: ` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + ` and ` + "`" + `checks` + "`" + `. The list of the attributes of each one is detailed below. The following is a list of the per-entry ` + "`" + `node` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Node ID of the Consul node advertising the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the node.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the node.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter in which the node is running.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Node meta data tag information, if any. The following is a list of the per-entry ` + "`" + `service` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The list of tags associated with this instance.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of this instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of this instance.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Service metadata tag information, if any. ` + "`" + `checks` + "`" + ` is a list of the health-checks associated to the entry with the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this health-check.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `The name of the node associated with this health-check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this health-check.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this health-check.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human readable description of the current state of the health-check.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `The output of the health-check.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service associated to this health-check.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service associated with this health-check.`,
				},
				resource.Attribute{
					Name:        "service_tags",
					Description: `The list of tags associated with this health-check.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "near",
					Description: `The node to which the result must be sorted to.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `The name of the tag used to filter the list.`,
				},
				resource.Attribute{
					Name:        "node_meta",
					Description: `The list of metadata to filter the nodes.`,
				},
				resource.Attribute{
					Name:        "passing",
					Description: `Whether to return only nodes with all checks in the passing state.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `A list of entries and details about each endpoint advertising a service. Each element in the list has three attributes: ` + "`" + `node` + "`" + `, ` + "`" + `service` + "`" + ` and ` + "`" + `checks` + "`" + `. The list of the attributes of each one is detailed below. The following is a list of the per-entry ` + "`" + `node` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Node ID of the Consul node advertising the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the node.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the node.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter in which the node is running.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Node meta data tag information, if any. The following is a list of the per-entry ` + "`" + `service` + "`" + ` attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The list of tags associated with this instance.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of this instance.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port of this instance.`,
				},
				resource.Attribute{
					Name:        "meta",
					Description: `Service metadata tag information, if any. ` + "`" + `checks` + "`" + ` is a list of the health-checks associated to the entry with the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of this health-check.`,
				},
				resource.Attribute{
					Name:        "node",
					Description: `The name of the node associated with this health-check.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of this health-check.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this health-check.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human readable description of the current state of the health-check.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `The output of the health-check.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the service associated to this health-check.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the service associated with this health-check.`,
				},
				resource.Attribute{
					Name:        "service_tags",
					Description: `The list of tags associated with this health-check.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "consul_services",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of services in a given Consul datacenter.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The Consul datacenter to query. Defaults to the same value found in ` + "`" + `query_options` + "`" + ` parameter specified below, or if that is empty, the ` + "`" + `datacenter` + "`" + ` value found in the Consul agent that this provider is configured to talk to.`,
				},
				resource.Attribute{
					Name:        "query_options",
					Description: `(Optional) See below. The ` + "`" + `query_options` + "`" + ` block supports the following:`,
				},
				resource.Attribute{
					Name:        "allow_stale",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, the default, allow responses from Consul servers that are followers.`,
				},
				resource.Attribute{
					Name:        "require_consistent",
					Description: `(Optional) When ` + "`" + `true` + "`" + ` force the client to perform a read on at least quorum servers and verify the result is the same. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "token",
					Description: `(Optional) Specify the Consul ACL token to use when performing the request. This defaults to the same API token configured by the ` + "`" + `consul` + "`" + ` provider but may be overriden if necessary.`,
				},
				resource.Attribute{
					Name:        "wait_index",
					Description: `(Optional) Index number used to enable blocking quereis.`,
				},
				resource.Attribute{
					Name:        "wait_time",
					Description: `(Optional) Max time the client should wait for a blocking query to return. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of the Consul services found. This will always contain the list of services found.`,
				},
				resource.Attribute{
					Name:        "services.<service>",
					Description: `For each name given, the corresponding attribute is a Terraform map of services and their tags. The value is an alphanumerically sorted, whitespace delimited set of tags associated with the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of the tags found for each service. If more than one service shares the same tag, unique service names will be joined by whitespace (this is the inverse of ` + "`" + `services` + "`" + ` and can be used to lookup the services that match a single tag).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter the keys are being read from to.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `A list of the Consul services found. This will always contain the list of services found.`,
				},
				resource.Attribute{
					Name:        "services.<service>",
					Description: `For each name given, the corresponding attribute is a Terraform map of services and their tags. The value is an alphanumerically sorted, whitespace delimited set of tags associated with the service.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A map of the tags found for each service. If more than one service shares the same tag, unique service names will be joined by whitespace (this is the inverse of ` + "`" + `services` + "`" + ` and can be used to lookup the services that match a single tag).`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"consul_agent_config":     0,
		"consul_agent_self":       1,
		"consul_autopilot_health": 2,
		"consul_key_prefix":       3,
		"consul_keys":             4,
		"consul_nodes":            5,
		"consul_service":          6,
		"consul_service_health":   7,
		"consul_services":         8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
