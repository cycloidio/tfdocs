package ovh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_private_network (deprecated)",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a private network in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"private",
				"network (deprecated)",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_cloud_private_network_subnet (deprecated)",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a subnet in a private network of a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"private",
				"network",
				"subnet (deprecated)",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_cloud_project_network_private",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a private network in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"project",
				"network",
				"private",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Deprecated. The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `project_id` + "`" + `.`,
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
					Name:        "service_name",
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
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
			Type:             "ovh_cloud_project_network_private_subnet",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a subnet in a private network of a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"project",
				"network",
				"private",
				"subnet",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Deprecated. The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `project_id` + "`" + `.`,
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
					Name:        "service_name",
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
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
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
			Type:             "ovh_cloud_project_user",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a user in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"project",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description associated with the user.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) Deprecated. The id of the public cloud project. If omitted, the ` + "`" + `OVH_PROJECT_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `project_id` + "`" + ` is required. Conflits with ` + "`" + `project_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The name of a role. See ` + "`" + `role_names` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `A list of role names. Values can be: - administrator, - ai_training_operator - authentication - backup_operator - compute_operator - image_operator - infrastructure_supervisor - network_operator - network_security_operator - objectstore_operator - volume_operator`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. Conflicts with ` + "`" + `project_id` + "`" + `. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `the date the user was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "openstack_rc",
					Description: `a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) the password generated for the user. The password can be used with the Openstack API. This attribute is sensitive and will only be retrieve once during creation.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles associated with the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the role`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the role`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the role`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `list of permissions associated with the role`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the user. should be normally set to 'ok'.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `the username generated for the user. This username can be used with the Openstack API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `the date the user was created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "openstack_rc",
					Description: `a convenient map representing an openstack_rc file. Note: no password nor sensitive token is set in this map.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) the password generated for the user. The password can be used with the Openstack API. This attribute is sensitive and will only be retrieve once during creation.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `A list of roles associated with the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the role`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the role`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the role`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `list of permissions associated with the role`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the user. should be normally set to 'ok'.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `the username generated for the user. This username can be used with the Openstack API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_user (deprecated)",
			Category:         "Cloud Resources",
			ShortDescription: `Creates a user in a public cloud project.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"user (deprecated)",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_dedicated_ceph_acl",
			Category:         "Dedicated CEPH",
			ShortDescription: `Creates a new ACL for dedicated CEPH cluster.`,
			Description:      ``,
			Keywords: []string{
				"dedicated",
				"ceph",
				"acl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your dedicated CEPH`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required) The network IP to authorize`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `(Required) The network mask to apply ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP family. ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `IP family. ` + "`" + `IPv4` + "`" + ` or ` + "`" + `IPv6` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server_reboot_task",
			Category:         "Dedicated Server",
			ShortDescription: `Reboot your Dedicated Server`,
			Description:      ``,
			Keywords: []string{
				"dedicated",
				"server",
				"reboot",
				"task",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `List of values traccked to trigger reboot, used also to form implicit dependencies ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The task id`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Details of this task. (should be ` + "`" + `Reboot asked` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "done_date",
					Description: `Completion date in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `Function name (should be ` + "`" + `hardReboot` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `Task creation date in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Task status (should be ` + "`" + `done` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The task id`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Details of this task. (should be ` + "`" + `Reboot asked` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "done_date",
					Description: `Completion date in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `Function name (should be ` + "`" + `hardReboot` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "start_date",
					Description: `Task creation date in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Task status (should be ` + "`" + `done` + "`" + `)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server_update",
			Category:         "Dedicated Server",
			ShortDescription: `Update various properties of your Dedicated Server`,
			Description:      ``,
			Keywords: []string{
				"dedicated",
				"server",
				"update",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server.`,
				},
				resource.Attribute{
					Name:        "boot_id",
					Description: `boot id of the server`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Icmp monitoring state`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `error, hacked, hackedBlocked, ok ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "boot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "boot_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `See Argument Reference above.`,
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_iploadbalancing_http_farm",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Creates a backend server group (farm).`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"http",
				"farm",
			},
			Arguments: []resource.Attribute{
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
					Description: `Stickiness type. No stickiness if null (` + "`" + `sourceIp` + "`" + `, ` + "`" + `cookie` + "`" + `)`,
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_iploadbalancing_http_farm_server",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Creates a backend server entry linked to farm.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"http",
				"farm",
				"server",
			},
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_iploadbalancing_http_frontend",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Creates a frontend for an IP Load balancing service.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"http",
				"frontend",
			},
			Arguments: []resource.Attribute{
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
					Description: `SSL deciphering. Default: 'false'`,
				},
				resource.Attribute{
					Name:        "redirect_location",
					Description: `Redirection HTTP' ## Attributes Reference The following attributes are exported:`,
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
			Attributes: []resource.Attribute{
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_iploadbalancing_refresh",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Applies changes from other ovh_iploadbalancing_* resourcesto the production configuration of loadbalancers.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"refresh",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `List of values traccked to trigger refresh, used also to form implicit dependencies ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "keepers",
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_iploadbalancing_vrack_network",
			Category:         "IP Load Balancing Resources",
			ShortDescription: `Manage a vrack network for your IP Loadbalancing service.`,
			Description:      ``,
			Keywords: []string{
				"ip",
				"load",
				"balancing",
				"iploadbalancing",
				"vrack",
				"network",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your vrack network`,
				},
				resource.Attribute{
					Name:        "farm_id",
					Description: `This attribute is there for documentation purpose only and isnt passed to the OVH API as it may conflicts with http/tcp farms ` + "`" + `vrack_network_id` + "`" + ` attribute`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `(Required) An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) IP block of the private network in the vRack`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN of the private network in the vRack. 0 if the private network is not in a VLAN ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `(Required) Internal Load Balancer identifier of the vRack private network`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `(Required) Internal Load Balancer identifier of the vRack private network`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_identity_user",
			Category:         "Me Resources",
			ShortDescription: `Creates an identity user.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"identity",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `User description.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `User's email.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `User's group.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `User's login suffix.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `User's password. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `Creation date of this user.`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update of this user.`,
				},
				resource.Attribute{
					Name:        "password_last_update",
					Description: `When the user changed his password for the last time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current user's status.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation",
					Description: `Creation date of this user.`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update of this user.`,
				},
				resource.Attribute{
					Name:        "password_last_update",
					Description: `When the user changed his password for the last time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current user's status.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template",
			Category:         "Me Resources",
			ShortDescription: `Creates a custom installation template available for dedicated servers.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"installation",
				"template",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template_partition_scheme",
			Category:         "Me Resources",
			ShortDescription: `Creates a partition scheme for a custom installation template available for dedicated servers.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"installation",
				"template",
				"partition",
				"scheme",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template_partition_scheme_hardware_raid",
			Category:         "Me Resources",
			ShortDescription: `Creates a hardware raid group in the partition scheme of a custom installation template available for dedicated servers.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"installation",
				"template",
				"partition",
				"scheme",
				"hardware",
				"raid",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template_partition_scheme_partition",
			Category:         "Me Resources",
			ShortDescription: `Creates a partition in the partition scheme of a custom installation template available for dedicated servers.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"installation",
				"template",
				"partition",
				"scheme",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ipxe_script",
			Category:         "Me Resources",
			ShortDescription: `Creates an IPXE Script.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"ipxe",
				"script",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `For documentation purpose only. This attribute is not passed to the OVH API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the IPXE Script.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `(Required) The content of the script. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ssh_key",
			Category:         "Me Resources",
			ShortDescription: `Creates an SSH Key.`,
			Description:      ``,
			Keywords: []string{
				"me",
				"ssh",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The friendly name of this SSH key.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The content of the public key in the form "ssh-algo content", e.g. "ssh-ed25519 AAAAC3...".`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `True when this public SSH key is used for rescue mode and reinstallations. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Arguments: []resource.Attribute{
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
			Attributes: []resource.Attribute{
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
			Type:             "ovh_vrack_cloudproject",
			Category:         "vRack Resources",
			ShortDescription: `Attach a Public Cloud Project to a VRack.`,
			Description:      ``,
			Keywords: []string{
				"vrack",
				"cloudproject",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `(Optional) Deprecated. The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `vrack_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `vrack_id` + "`" + ` is required. Conflits with ` + "`" + `vrack_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_dedicated_server",
			Category:         "vRack Resources",
			ShortDescription: `Attach a Dedicated Server to a VRack.`,
			Description:      ``,
			Keywords: []string{
				"vrack",
				"dedicated",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `(Optional) Deprecated. The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `vrack_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `vrack_id` + "`" + ` is required. Conflits with ` + "`" + `vrack_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) The id of the dedicated server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_dedicated_server_interface",
			Category:         "vRack Resources",
			ShortDescription: `Attach a Dedicated Server Network Interface to a VRack.`,
			Description:      ``,
			Keywords: []string{
				"vrack",
				"dedicated",
				"server",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `(Optional) Deprecated. The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_ID` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `vrack_id` + "`" + ` is required. Conflits with ` + "`" + `service_name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_SERVICE` + "`" + ` environment variable is used. One of ` + "`" + `service_name` + "`" + ` or ` + "`" + `vrack_id` + "`" + ` is required. Conflits with ` + "`" + `vrack_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interface_id",
					Description: `(Required) The id of dedicated server network interface. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vrack_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "interface_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_iploadbalancing",
			Category:         "vRack Resources",
			ShortDescription: `Attach a IP Loadbalanging to a VRack.`,
			Description:      ``,
			Keywords: []string{
				"vrack",
				"iploadbalancing",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the vrack.`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `(Required) The id of the ip loadbalancing. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `See Argument Reference above.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ovh_cloud_private_network (deprecated)":                      0,
		"ovh_cloud_private_network_subnet (deprecated)":               1,
		"ovh_cloud_project_network_private":                           2,
		"ovh_cloud_project_network_private_subnet":                    3,
		"ovh_cloud_project_user":                                      4,
		"ovh_cloud_user (deprecated)":                                 5,
		"ovh_dedicated_ceph_acl":                                      6,
		"ovh_dedicated_server_reboot_task":                            7,
		"ovh_dedicated_server_update":                                 8,
		"ovh_ip_reverse":                                              9,
		"ovh_iploadbalancing_http_farm":                               10,
		"ovh_iploadbalancing_http_farm_server":                        11,
		"ovh_iploadbalancing_http_frontend":                           12,
		"ovh_iploadbalancing_http_route":                              13,
		"ovh_iploadbalancing_http_route_rule":                         14,
		"ovh_iploadbalancing_refresh":                                 15,
		"ovh_iploadbalancing_tcp_farm":                                16,
		"ovh_iploadbalancing_tcp_farm_server":                         17,
		"ovh_iploadbalancing_tcp_frontend":                            18,
		"ovh_iploadbalancing_vrack_network":                           19,
		"ovh_me_identity_user":                                        20,
		"ovh_me_installation_template":                                21,
		"ovh_me_installation_template_partition_scheme":               22,
		"ovh_me_installation_template_partition_scheme_hardware_raid": 23,
		"ovh_me_installation_template_partition_scheme_partition":     24,
		"ovh_me_ipxe_script":                                          25,
		"ovh_me_ssh_key":                                              26,
		"ovh_domain_zone_record":                                      27,
		"ovh_domain_zone_redirection":                                 28,
		"ovh_vrack_cloudproject":                                      29,
		"ovh_vrack_dedicated_server":                                  30,
		"ovh_vrack_dedicated_server_interface":                        31,
		"ovh_vrack_iploadbalancing":                                   32,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
