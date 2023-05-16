package citrixadc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_hanode",
			Category:         "High Availability",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"high",
				"availability",
				"hanode",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hanode_id",
					Description: `(Required) Number that uniquely identifies the node. For self node, it will always be 0. Peer node values can range from 1-64. ## Attribute Reference In addition to the arguments, the following attributes are available:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the hanode. It has the same value as the ` + "`" + `hanode_id` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "completedfliptime",
					Description: `To inform user whether flip time is elapsed or not.`,
				},
				resource.Attribute{
					Name:        "curflips",
					Description: `Keeps track of number of flips that have happened till now in current interval.`,
				},
				resource.Attribute{
					Name:        "deadinterval",
					Description: `Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.`,
				},
				resource.Attribute{
					Name:        "enaifaces",
					Description: `Enabled interfaces.`,
				},
				resource.Attribute{
					Name:        "failsafe",
					Description: `Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.`,
				},
				resource.Attribute{
					Name:        "haprop",
					Description: `Automatically propagate all commands from the primary to the secondary node, except the following:`,
				},
				resource.Attribute{
					Name:        "hastatus",
					Description: `The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node. STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.`,
				},
				resource.Attribute{
					Name:        "hasync",
					Description: `Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.`,
				},
				resource.Attribute{
					Name:        "hellointerval",
					Description: `Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.`,
				},
				resource.Attribute{
					Name:        "inc",
					Description: `This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.`,
				},
				resource.Attribute{
					Name:        "masterstatetime",
					Description: `Time elapsed in current master state.`,
				},
				resource.Attribute{
					Name:        "maxflips",
					Description: `Max number of flips allowed before becoming sticky primary`,
				},
				resource.Attribute{
					Name:        "maxfliptime",
					Description: `Interval after which flipping of node states can again start`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The netmask`,
				},
				resource.Attribute{
					Name:        "routemonitor",
					Description: `The IP address (IPv4 or IPv6).`,
				},
				resource.Attribute{
					Name:        "routemonitorstate",
					Description: `State for route monitor.`,
				},
				resource.Attribute{
					Name:        "ssl2",
					Description: `SSL card status.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `HA Master State.`,
				},
				resource.Attribute{
					Name:        "syncstatusstrictmode",
					Description: `strict mode flag for sync status`,
				},
				resource.Attribute{
					Name:        "syncvlan",
					Description: `Vlan on which HA related communication is sent. This include sync, propagation , connection mirroring , LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the hanode. It has the same value as the ` + "`" + `hanode_id` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "completedfliptime",
					Description: `To inform user whether flip time is elapsed or not.`,
				},
				resource.Attribute{
					Name:        "curflips",
					Description: `Keeps track of number of flips that have happened till now in current interval.`,
				},
				resource.Attribute{
					Name:        "deadinterval",
					Description: `Number of seconds after which a peer node is marked DOWN if heartbeat messages are not received from the peer node.`,
				},
				resource.Attribute{
					Name:        "enaifaces",
					Description: `Enabled interfaces.`,
				},
				resource.Attribute{
					Name:        "failsafe",
					Description: `Keep one node primary if both nodes fail the health check, so that a partially available node can back up data and handle traffic. This mode is set independently on each node.`,
				},
				resource.Attribute{
					Name:        "haprop",
					Description: `Automatically propagate all commands from the primary to the secondary node, except the following:`,
				},
				resource.Attribute{
					Name:        "hastatus",
					Description: `The HA status of the node. The HA status STAYSECONDARY is used to force the secondary device stay as secondary independent of the state of the Primary device. For example, in an existing HA setup, the Primary node has to be upgraded and this process would take few seconds. During the upgradation, it is possible that the Primary node may suffer from a downtime for a few seconds. However, the Secondary should not take over as the Primary node. Thus, the Secondary node should remain as Secondary even if there is a failure in the Primary node. STAYPRIMARY configuration keeps the node in primary state in case if it is healthy, even if the peer node was the primary node initially. If the node with STAYPRIMARY setting (and no peer node) is added to a primary node (which has this node as the peer) then this node takes over as the new primary and the older node becomes secondary. ENABLED state means normal HA operation without any constraints/preferences. DISABLED state disables the normal HA operation of the node.`,
				},
				resource.Attribute{
					Name:        "hasync",
					Description: `Automatically maintain synchronization by duplicating the configuration of the primary node on the secondary node. This setting is not propagated. Automatic synchronization requires that this setting be enabled (the default) on the current secondary node. Synchronization uses TCP port 3010.`,
				},
				resource.Attribute{
					Name:        "hellointerval",
					Description: `Interval, in milliseconds, between heartbeat messages sent to the peer node. The heartbeat messages are UDP packets sent to port 3003 of the peer node.`,
				},
				resource.Attribute{
					Name:        "inc",
					Description: `This option is required if the HA nodes reside on different networks. When this mode is enabled, the following independent network entities and configurations are neither propagated nor synced to the other node: MIPs, SNIPs, VLANs, routes (except LLB routes), route monitors, RNAT rules (except any RNAT rule with a VIP as the NAT IP), and dynamic routing configurations. They are maintained independently on each node.`,
				},
				resource.Attribute{
					Name:        "ipaddress",
					Description: `The NSIP or NSIP6 address of the node to be added for an HA configuration. This setting is neither propagated nor synchronized.`,
				},
				resource.Attribute{
					Name:        "masterstatetime",
					Description: `Time elapsed in current master state.`,
				},
				resource.Attribute{
					Name:        "maxflips",
					Description: `Max number of flips allowed before becoming sticky primary`,
				},
				resource.Attribute{
					Name:        "maxfliptime",
					Description: `Interval after which flipping of node states can again start`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `The netmask`,
				},
				resource.Attribute{
					Name:        "routemonitor",
					Description: `The IP address (IPv4 or IPv6).`,
				},
				resource.Attribute{
					Name:        "routemonitorstate",
					Description: `State for route monitor.`,
				},
				resource.Attribute{
					Name:        "ssl2",
					Description: `SSL card status.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `HA Master State.`,
				},
				resource.Attribute{
					Name:        "syncstatusstrictmode",
					Description: `strict mode flag for sync status`,
				},
				resource.Attribute{
					Name:        "syncvlan",
					Description: `Vlan on which HA related communication is sent. This include sync, propagation , connection mirroring , LB persistency config sync, persistent session sync and session state sync. However HA heartbeats can go all interfaces.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nitro_info",
			Category:         "Generic",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"generic",
				"nitro",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "workflow",
					Description: `(Optional) Dictionary containing the data that will guide the data source execution. Currently the following attributes are taken into consideration.`,
				},
				resource.Attribute{
					Name:        "primary_id",
					Description: `(Optional) Value for the primary id of the nitro endpoint.`,
				},
				resource.Attribute{
					Name:        "secondary_id",
					Description: `(Optional) Value for the secondary id of the nitro endpoint. ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "nitro_list",
					Description: `Contains the result returned by the ` + "`" + `binding_list` + "`" + ` lifecycle.`,
				},
				resource.Attribute{
					Name:        "nitro_object",
					Description: `Contains the result returned by the ` + "`" + `object_by_name` + "`" + ` lifecycle.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nitro_list",
					Description: `Contains the result returned by the ` + "`" + `binding_list` + "`" + ` lifecycle.`,
				},
				resource.Attribute{
					Name:        "nitro_object",
					Description: `Contains the result returned by the ` + "`" + `object_by_name` + "`" + ` lifecycle.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_nsversion",
			Category:         "NS",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"ns",
				"nsversion",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "citrixadc_sslcipher_sslvserver_bindings",
			Category:         "SSL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"ssl",
				"sslcipher",
				"sslvserver",
				"bindings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ciphername",
					Description: `(Required) Name of the cipher. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "bound_sslvservers",
					Description: `The list of the bound sslvservers names.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "bound_sslvservers",
					Description: `The list of the bound sslvservers names.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"citrixadc_hanode":                        0,
		"citrixadc_nitro_info":                    1,
		"citrixadc_nsversion":                     2,
		"citrixadc_sslcipher_sslvserver_bindings": 3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
