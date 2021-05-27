package ionoscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_datacenter",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud Data Centers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of an existing Virtual Data Center that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing Virtual Data Center's location. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_image",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud Images`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of an existing image that you want to search for.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the image (see details below).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing image's location.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The image type, HDD or CD-ROM.`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `Image aliases`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility ("NONE" or "V1") If both "name" and "version" are provided the plugin will concatenate the two strings in this format [name]-[version]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_k8s_cluster",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud K8s Cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name or an existing cluster that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the cluster you want to search for. k Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the cluster`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubernetes version`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `list of the IDs of the node pools in this cluster`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Kubernetes configuration`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the cluster`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubernetes version`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `list of the IDs of the node pools in this cluster`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Kubernetes configuration`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_k8s_node_pool",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud K8s Node Pool`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing node pool that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the node pool you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the node pool`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the node pool`,
				},
				resource.Attribute{
					Name:        "k8s_cluster_id",
					Description: `ID of the cluster this node pool is part of`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The UUID of the VDC`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes in this node pool`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `CPU Family`,
				},
				resource.Attribute{
					Name:        "cores_count",
					Description: `CPU cores count`,
				},
				resource.Attribute{
					Name:        "ram_size",
					Description: `The amount of RAM in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The compute availability zone in which the nodes should exist`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `HDD or SDD`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Size of storage`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The kubernetes version`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A maintenance window comprise of a day of the week and a time for maintenance to be allowed`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `A clock time in the day when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "day_of_the_week",
					Description: `Day of the week when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `The range defining the minimum and maximum number of worker nodes that the managed node group can scale in`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `The minimum number of worker nodes the node pool can scale down to`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `The maximum number of worker nodes that the node pool can scale to`,
				},
				resource.Attribute{
					Name:        "lans",
					Description: `A list of Local Area Networks the node pool is a part of`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `A map of annotations in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `A list of kubernetes versions available for upgrade`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `The list of fixed IPs associated with this node pool`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the node pool`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the node pool`,
				},
				resource.Attribute{
					Name:        "k8s_cluster_id",
					Description: `ID of the cluster this node pool is part of`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The UUID of the VDC`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes in this node pool`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `CPU Family`,
				},
				resource.Attribute{
					Name:        "cores_count",
					Description: `CPU cores count`,
				},
				resource.Attribute{
					Name:        "ram_size",
					Description: `The amount of RAM in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The compute availability zone in which the nodes should exist`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `HDD or SDD`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `Size of storage`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The kubernetes version`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A maintenance window comprise of a day of the week and a time for maintenance to be allowed`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `A clock time in the day when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "day_of_the_week",
					Description: `Day of the week when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `The range defining the minimum and maximum number of worker nodes that the managed node group can scale in`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `The minimum number of worker nodes the node pool can scale down to`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `The maximum number of worker nodes that the node pool can scale to`,
				},
				resource.Attribute{
					Name:        "lans",
					Description: `A list of Local Area Networks the node pool is a part of`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `A map of annotations in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `A list of kubernetes versions available for upgrade`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `The list of fixed IPs associated with this node pool`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_lan",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Ionos Cloud Lans`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing lan that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the lan you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "ip_failover",
					Description: `list of`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_failover",
					Description: `list of`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_location",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud Locations`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the location to search for.`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `(Optional) A desired feature that the location must be able to provide. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the location`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the location`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_private_crossconnect",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Ionos Cloud Private Crossconnects`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing private crossconnect that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the private crossconnect you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "peers",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "connectable_datacenters",
					Description: `list of`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "peers",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "connectable_datacenters",
					Description: `list of`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_resource",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud Resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The specific type of resources to retrieve information about.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) The ID of the specific resource to retrieve information about. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_server",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Ionos Cloud Servers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing server that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the server you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "cdroms",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `list of`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cdroms",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `list of`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information on a IonosCloud Snapshots`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of an existing snapshot that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing snapshot's location.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the snapshot to look for. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the snapshot`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the snapshot`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ionoscloud_datacenter":           0,
		"ionoscloud_image":                1,
		"ionoscloud_k8s_cluster":          2,
		"ionoscloud_k8s_node_pool":        3,
		"ionoscloud_lan":                  4,
		"ionoscloud_location":             5,
		"ionoscloud_private_crossconnect": 6,
		"ionoscloud_resource":             7,
		"ionoscloud_server":               8,
		"ionoscloud_snapshot":             9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
