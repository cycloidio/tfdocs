package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nsxt_certificate",
			Category:         "Data Sources",
			ShortDescription: `A certificate data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Certificate to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the Certificate to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Certificate.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Certificate.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_edge_cluster",
			Category:         "Data Sources",
			ShortDescription: `An Edge Cluster data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Edge Cluster to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Edge Cluster to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the edge cluster.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `This field could show deployment_type of members. It would return UNKNOWN if there is no members, and return VIRTUAL_MACHINE|PHYSICAL_MACHINE if all Edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE.`,
				},
				resource.Attribute{
					Name:        "member_node_type",
					Description: `An Edge cluster is homogeneous collection of NSX transport nodes used for north/south connectivity between NSX logical networking and physical networking. Hence all transport nodes of the cluster must be of same type. This field shows the type of transport node,`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the edge cluster.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `This field could show deployment_type of members. It would return UNKNOWN if there is no members, and return VIRTUAL_MACHINE|PHYSICAL_MACHINE if all Edge members are VIRTUAL_MACHINE|PHYSICAL_MACHINE.`,
				},
				resource.Attribute{
					Name:        "member_node_type",
					Description: `An Edge cluster is homogeneous collection of NSX transport nodes used for north/south connectivity between NSX logical networking and physical networking. Hence all transport nodes of the cluster must be of same type. This field shows the type of transport node,`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier0_router",
			Category:         "Data Sources",
			ShortDescription: `A logical Tier 0 router data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Logical Tier 0 Router to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Logical Tier 0 Router to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `The high availability mode of this logical router.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
				resource.Attribute{
					Name:        "high_availability_mode",
					Description: `The high availability mode of this logical router.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_logical_tier1_router",
			Category:         "Data Sources",
			ShortDescription: `A logical Tier 1 router data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Logical Tier 1 Router to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Logical Tier 1 Router to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the logical Tier 0 router.`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `The id of the Edge cluster where this logical router is placed.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_mac_pool",
			Category:         "Data Sources",
			ShortDescription: `A MAC pool data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of MAC pool to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the MAC pool to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the MAC pool.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the MAC pool.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_group",
			Category:         "Data Sources",
			ShortDescription: `A networking and security group data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of NS group to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the NS group to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS group.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_ns_service",
			Category:         "Data Sources",
			ShortDescription: `A networking and security service data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of NS service to retrieve`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the NS service to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS service.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the NS service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_switching_profile",
			Category:         "Data Sources",
			ShortDescription: `A switching profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Switching Profile to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name of the Switching Profile to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource type representing the specific type of this switching profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the switching profile.`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "resource_type",
					Description: `The resource type representing the specific type of this switching profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the switching profile.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nsxt_transport_zone",
			Category:         "Data Sources",
			ShortDescription: `A transport zone data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of Transport Zone to retrieve.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) The Display Name prefix of the Transport Zone to retrieve. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Transport Zone.`,
				},
				resource.Attribute{
					Name:        "host_switch_name",
					Description: `The name of the N-VDS (host switch) on all Transport Nodes in this Transport Zone that will be used to run NSX network traffic.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `The transport type of this transport zone (OVERLAY or VLAN).`,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the Transport Zone.`,
				},
				resource.Attribute{
					Name:        "host_switch_name",
					Description: `The name of the N-VDS (host switch) on all Transport Nodes in this Transport Zone that will be used to run NSX network traffic.`,
				},
				resource.Attribute{
					Name:        "transport_type",
					Description: `The transport type of this transport zone (OVERLAY or VLAN).`,
				},
			},
		},
	}

	dataSourcesMap = map[string]Resource{

		"nsxt_certificate":          0,
		"nsxt_edge_cluster":         1,
		"nsxt_logical_tier0_router": 2,
		"nsxt_logical_tier1_router": 3,
		"nsxt_mac_pool":             4,
		"nsxt_ns_group":             5,
		"nsxt_ns_service":           6,
		"nsxt_switching_profile":    7,
		"nsxt_transport_zone":       8,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
