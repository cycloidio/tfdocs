package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_network_private",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a private network in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"network",
				"private",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `an array of valid OVH public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions_status",
					Description: `A map representing the status of the network per region.`,
				},
				resource.Attribute{
					Name:        "regions_status/region",
					Description: `The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_status/status",
					Description: `The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the network. should be normally set to 'ACTIVE'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the network. Either 'private' or 'public'.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions_status",
					Description: `A map representing the status of the network per region.`,
				},
				resource.Attribute{
					Name:        "regions_status/region",
					Description: `The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_status/status",
					Description: `The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the network. should be normally set to 'ACTIVE'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the network. Either 'private' or 'public'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_network_private_subnet",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a subnet in a private network of a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"network",
				"private",
				"subnet",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The id of the network. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `(Optional) Enable DHCP. Changing this forces a new resource to be created. Defaults to false. _`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) First ip for this region. Changing this value recreates the subnet.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) Last ip for this region. Changing this value recreates the subnet.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Global network in CIDR format. Changing this value recreates the subnet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which the network subnet will be created. Ex.: "GRA1". Changing this value recreates the resource.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `Set to true if you don't want to set a default gateway IP. Changing this value recreates the resource. Defaults to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `The IP of the gateway`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Ip Block representing the subnet cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools",
					Description: `List of ip pools allocated in the subnet.`,
				},
				resource.Attribute{
					Name:        "ip_pools/network",
					Description: `Global network with cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools/region",
					Description: `Region where this subnet is created.`,
				},
				resource.Attribute{
					Name:        "ip_pools/dhcp",
					Description: `DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "ip_pools/end",
					Description: `Last ip for this region.`,
				},
				resource.Attribute{
					Name:        "ip_pools/start",
					Description: `First ip for this region.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `The IP of the gateway`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Ip Block representing the subnet cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools",
					Description: `List of ip pools allocated in the subnet.`,
				},
				resource.Attribute{
					Name:        "ip_pools/network",
					Description: `Global network with cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools/region",
					Description: `Region where this subnet is created.`,
				},
				resource.Attribute{
					Name:        "ip_pools/dhcp",
					Description: `DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "ip_pools/end",
					Description: `Last ip for this region.`,
				},
				resource.Attribute{
					Name:        "ip_pools/start",
					Description: `First ip for this region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_user",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a user in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description associated with the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `the username generated for the user. This username can be used with the Openstack API.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) the password generated for the user. The password can be used with the Openstack API. This attribute is sensitive and will only be retrieve once during creation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the user. should be normally set to 'ok'.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `the date the user was created.`,
				},
				resource.Attribute{
					Name:        "openstack_rc",
					Description: `a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `the username generated for the user. This username can be used with the Openstack API.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) the password generated for the user. The password can be used with the Openstack API. This attribute is sensitive and will only be retrieve once during creation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the user. should be normally set to 'ok'.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `the date the user was created.`,
				},
				resource.Attribute{
					Name:        "openstack_rc",
					Description: `a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_ip_reverse",
			Category:         "IP Resources",
			ShortDescription: `Provides a OVH IP reverse resource.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"reverse",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The IP block to which the IP belongs`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `(Required) The value of the reverse`,
				},
				resource.Attribute{
					Name:        "ipreverse",
					Description: `(Optional) The IP to set the reverse of, default to ` + "`" + `ip` + "`" + ` if ` + "`" + `ip` + "`" + ` is a /32 (IPv4) or a /128 (IPv6) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ipreverse",
					Description: `The IP to set the reverse of`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The value of the reverse`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "ipreverse",
					Description: `The IP to set the reverse of`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `The value of the reverse`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_http_route",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Manage http route for a loadbalancer service.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"http",
				"route",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your route, this field is for you`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action`,
				},
				resource.Attribute{
					Name:        "action.status",
					Description: `HTTP status code for "redirect" and "reject" actions`,
				},
				resource.Attribute{
					Name:        "action.target",
					Description: `Farm ID for "farm" action type or URL template for "redirect" action. You may use ${uri}, ${protocol}, ${host}, ${port} and ${path} variables in redirect target`,
				},
				resource.Attribute{
					Name:        "action.type",
					Description: `(Required) Action to trigger if all the rules of this route matches`,
				},
				resource.Attribute{
					Name:        "frontend_id",
					Description: `Route traffic for this frontend ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action.status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action.target",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action.type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "frontend_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action.status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action.target",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "action.type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "frontend_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_http_route_rule",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Manage rules for HTTP route.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"http",
				"route",
				"rule",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `(Required) The route to apply this rule`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your rule, this field is for you`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Name of the field to match like "protocol" or "host". See "/ipLoadbalancing/{serviceName}/availableRouteRules" for a list of available rules`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `(Required) Matching operator. Not all operators are available for all fields. See "/ipLoadbalancing/{serviceName}/availableRouteRules"`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `Invert the matching operator effect`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Value to match against this match. Interpretation if this field depends on the match and field`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `Name of sub-field, if applicable. This may be a Cookie or Header name for instance ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "route_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_tcp_farm",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Creates a backend server group (farm).`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"tcp",
				"farm",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `Load balancing algorithm. ` + "`" + `roundrobin` + "`" + ` if null (` + "`" + `first` + "`" + `, ` + "`" + `leastconn` + "`" + `, ` + "`" + `roundrobin` + "`" + `, ` + "`" + `source` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Readable label for loadbalancer farm`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port attached to your farm ([1..49151]). Inherited from frontend if null`,
				},
				resource.Attribute{
					Name:        "stickiness",
					Description: `Stickiness type. No stickiness if null (` + "`" + `sourceIp` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Zone where the farm will be defined (ie. ` + "`" + `GRA` + "`" + `, ` + "`" + `BHS` + "`" + ` also supports ` + "`" + `ALL` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "probe",
					Description: `define a backend healthcheck probe`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Valid values : ` + "`" + `http` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `mysql` + "`" + `, ` + "`" + `oko` + "`" + `, ` + "`" + `pgsql` + "`" + `, ` + "`" + `smtp` + "`" + `, ` + "`" + `tcp` + "`" + ``,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `probe interval, Value between 30 and 3600 seconds, default 30`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `What to mach ` + "`" + `pattern` + "`" + ` against (` + "`" + `contains` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `matches` + "`" + `, ` + "`" + `status` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port for backends to recieve traffic on.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `Negate probe result`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Pattern to match against ` + "`" + `match` + "`" + ``,
				},
				resource.Attribute{
					Name:        "force_ssl",
					Description: `Force use of SSL (TLS)`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `URL for HTTP probe type.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `HTTP probe method (` + "`" + `GET` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `internal` + "`" + `) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "stickiness",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "probe",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "force_ssl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "balance",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "stickiness",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "probe",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "negate",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "force_ssl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_tcp_farm_server",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Creates a backend server entry linked to farm.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"tcp",
				"farm",
				"server",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "farm_id",
					Description: `ID of the farm this server is attached to`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Label for the server`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Address of the backend server (IP from either internal or OVH network)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `backend status - ` + "`" + `active` + "`" + ` or ` + "`" + `inactive` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port that backend will respond on`,
				},
				resource.Attribute{
					Name:        "proxy_protocol_version",
					Description: `version of the PROXY protocol used to pass origin connection information from loadbalancer to recieving service (` + "`" + `v1` + "`" + `, ` + "`" + `v2` + "`" + `, ` + "`" + `v2-ssl` + "`" + `, ` + "`" + `v2-ssl-cn` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `used in loadbalancing algorithm`,
				},
				resource.Attribute{
					Name:        "probe",
					Description: `defines if backend will be probed to determine health and keep as active in farm if healthy`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `is the connection ciphered with SSL (TLS)`,
				},
				resource.Attribute{
					Name:        "backup",
					Description: `is it a backup server used in case of failure of all the non-backup backends ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "farm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "probe",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `Value of the stickiness cookie used for this backend.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "farm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "proxy_protocol_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "probe",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cookie",
					Description: `Value of the stickiness cookie used for this backend.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_tcp_frontend",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Creates a frontend for an IP Load balancing service.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"tcp",
				"frontend",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your frontend, this field is for you`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port(s) attached to your frontend. Supports single port (numerical value), range (2 dash-delimited increasing ports) and comma-separated list of 'single port' and/or 'range'. Each port must be in the [1;49151] range`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) Zone where the frontend will be defined (ie. ` + "`" + `gra` + "`" + `, ` + "`" + `bhs` + "`" + ` also supports ` + "`" + `all` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "allowed_source",
					Description: `Restrict IP Load Balancing access to these ip block. No restriction if null. List of IP blocks.`,
				},
				resource.Attribute{
					Name:        "dedicated_ipfo",
					Description: `Only attach frontend on these ip. No restriction if null. List of Ip blocks.`,
				},
				resource.Attribute{
					Name:        "default_farm_id",
					Description: `Default TCP Farm of your frontend`,
				},
				resource.Attribute{
					Name:        "default_ssl_id",
					Description: `Default ssl served to your customer`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disable your frontend. Default: 'false'`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `SSL deciphering. Default: 'false' ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of your frontend`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allowed_source",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dedicated_ipfo",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_farm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_ssl_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `Id of your frontend`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "allowed_source",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dedicated_ipfo",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_farm_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_ssl_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone_record",
			Category:         "Domain Resources",
			ShortDescription: `Provides a OVH domain zone resource.`,
			Description:      ``,
			Keywords: []string{
				"domain",
				"zone",
				"record",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Required) The name of the record`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The value of the record`,
				},
				resource.Attribute{
					Name:        "fieldtype",
					Description: `(Required) The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The TTL of the record ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "subDomain",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The value of the record`,
				},
				resource.Attribute{
					Name:        "fieldType",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record ## Import OVH record can be imported using the ` + "`" + `id` + "`" + ` and the ` + "`" + `zone` + "`" + `, eg: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import ovh_domain_zone_record.test 1234OVH_ID.zone.tld ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The record ID`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The domain to add the record to`,
				},
				resource.Attribute{
					Name:        "subDomain",
					Description: `The name of the record`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The value of the record`,
				},
				resource.Attribute{
					Name:        "fieldType",
					Description: `The type of the record`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `The TTL of the record ## Import OVH record can be imported using the ` + "`" + `id` + "`" + ` and the ` + "`" + `zone` + "`" + `, eg: ` + "`" + `` + "`" + `` + "`" + `sh $ terraform import ovh_domain_zone_record.test 1234OVH_ID.zone.tld ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone_redirection",
			Category:         "Domain Resources",
			ShortDescription: `Provides a OVH domain zone resource.`,
			Description:      ``,
			Keywords: []string{
				"domain",
				"zone",
				"redirection",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "zone",
					Description: `(Required) The domain to add the redirection to`,
				},
				resource.Attribute{
					Name:        "subdomain",
					Description: `(Optional) The name of the redirection`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The value of the redirection`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the redirection, with values:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of this redirection`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `(Optional) Keywords to describe this redirection`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Optional) Title of this redirection ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The redirection ID`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The domain to add the redirection to`,
				},
				resource.Attribute{
					Name:        "subDomain",
					Description: `The name of the redirection`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The value of the redirection`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the redirection`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the redirection`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `Keywords of the redirection`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the redirection`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The redirection ID`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The domain to add the redirection to`,
				},
				resource.Attribute{
					Name:        "subDomain",
					Description: `The name of the redirection`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `The value of the redirection`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the redirection`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the redirection`,
				},
				resource.Attribute{
					Name:        "keywords",
					Description: `Keywords of the redirection`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `The title of the redirection`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_publiccloud_private_network",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a private network in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"publiccloud",
				"private",
				"network",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the network.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `an array of valid OVH public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions_status",
					Description: `A map representing the status of the network per region.`,
				},
				resource.Attribute{
					Name:        "regions_status/region",
					Description: `The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_status/status",
					Description: `The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the network. should be normally set to 'ACTIVE'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the network. Either 'private' or 'public'.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "regions_status",
					Description: `A map representing the status of the network per region.`,
				},
				resource.Attribute{
					Name:        "regions_status/region",
					Description: `The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_status/status",
					Description: `The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the network. should be normally set to 'ACTIVE'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the network. Either 'private' or 'public'.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_publiccloud_private_network_subnet",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a subnet in a private network of a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"publiccloud",
				"private",
				"network",
				"subnet",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The id of the network. Changing this forces a new resource to be created.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `(Optional) Enable DHCP. Changing this forces a new resource to be created. Defaults to false. _`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) First ip for this region. Changing this value recreates the subnet.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) Last ip for this region. Changing this value recreates the subnet.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) Global network in CIDR format. Changing this value recreates the subnet`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which the network subnet will be created. Ex.: "GRA1". Changing this value recreates the resource.`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `Set to true if you don't want to set a default gateway IP. Changing this value recreates the resource. Defaults to false. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `The IP of the gateway`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Ip Block representing the subnet cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools",
					Description: `List of ip pools allocated in the subnet.`,
				},
				resource.Attribute{
					Name:        "ip_pools/network",
					Description: `Global network with cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools/region",
					Description: `Region where this subnet is created.`,
				},
				resource.Attribute{
					Name:        "ip_pools/dhcp",
					Description: `DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "ip_pools/end",
					Description: `Last ip for this region.`,
				},
				resource.Attribute{
					Name:        "ip_pools/start",
					Description: `First ip for this region.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "dhcp_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `The IP of the gateway`,
				},
				resource.Attribute{
					Name:        "no_gateway",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `Ip Block representing the subnet cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools",
					Description: `List of ip pools allocated in the subnet.`,
				},
				resource.Attribute{
					Name:        "ip_pools/network",
					Description: `Global network with cidr.`,
				},
				resource.Attribute{
					Name:        "ip_pools/region",
					Description: `Region where this subnet is created.`,
				},
				resource.Attribute{
					Name:        "ip_pools/dhcp",
					Description: `DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "ip_pools/end",
					Description: `Last ip for this region.`,
				},
				resource.Attribute{
					Name:        "ip_pools/start",
					Description: `First ip for this region.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_publiccloud_user",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a user in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"publiccloud",
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description associated with the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `the username generated for the user. This username can be used with the Openstack API.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) the password generated for the user. The password can be used with the Openstack API. This attribute is sensitive and will only be retrieve once during creation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the user. should be normally set to 'ok'.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `the date the user was created.`,
				},
				resource.Attribute{
					Name:        "openstack_rc",
					Description: `a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `the username generated for the user. This username can be used with the Openstack API.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) the password generated for the user. The password can be used with the Openstack API. This attribute is sensitive and will only be retrieve once during creation.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the user. should be normally set to 'ok'.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `the date the user was created.`,
				},
				resource.Attribute{
					Name:        "openstack_rc",
					Description: `a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_cloudproject",
			Category:         "vRack Resources",
			ShortDescription: `Attach an existing public cloud project to an existing VRack.`,
			Description:      ``,
			Keywords: []string{
				"vrack",
				"cloudproject",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `(Required) The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Notes The vrack attachment isn't a proper resource with an ID. As such, the resource id will be forged from the vrack and project ids and there's no correct way to import the resource in terraform. When the resource is created by terraform, it first checks if the attachment already exists within OVH infrastructure; if it exists it set the resource id without modifying anything. Otherwise, it will try to attach the vrack with the public cloud project.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Notes The vrack attachment isn't a proper resource with an ID. As such, the resource id will be forged from the vrack and project ids and there's no correct way to import the resource in terraform. When the resource is created by terraform, it first checks if the attachment already exists within OVH infrastructure; if it exists it set the resource id without modifying anything. Otherwise, it will try to attach the vrack with the public cloud project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_publiccloud_attachment",
			Category:         "vRack Resources",
			ShortDescription: `Attach an existing PublicCloud project to an existing VRack.`,
			Description:      ``,
			Keywords: []string{
				"vrack",
				"publiccloud",
				"attachment",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `(Required) The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_ID` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Notes The vrack attachment isn't a proper resource with an ID. As such, the resource id will be forged from the vrack and project ids and there's no correct way to import the resource in terraform. When the resource is created by terraform, it first checks if the attachment already exists within OVH infrastructure; if it exists it set the resource id without modifying anything. Otherwise, it will try to attach the vrack with the public cloud project.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Notes The vrack attachment isn't a proper resource with an ID. As such, the resource id will be forged from the vrack and project ids and there's no correct way to import the resource in terraform. When the resource is created by terraform, it first checks if the attachment already exists within OVH infrastructure; if it exists it set the resource id without modifying anything. Otherwise, it will try to attach the vrack with the public cloud project.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ovh_cloud_network_private":              0,
		"ovh_cloud_network_private_subnet":       1,
		"ovh_cloud_user":                         2,
		"ovh_ip_reverse":                         3,
		"ovh_iploadbalancing_http_route":         4,
		"ovh_iploadbalancing_http_route_rule":    5,
		"ovh_iploadbalancing_tcp_farm":           6,
		"ovh_iploadbalancing_tcp_farm_server":    7,
		"ovh_iploadbalancing_tcp_frontend":       8,
		"ovh_domain_zone_record":                 9,
		"ovh_domain_zone_redirection":            10,
		"ovh_publiccloud_private_network":        11,
		"ovh_publiccloud_private_network_subnet": 12,
		"ovh_publiccloud_user":                   13,
		"ovh_vrack_cloudproject":                 14,
		"ovh_vrack_publiccloud_attachment":       15,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
