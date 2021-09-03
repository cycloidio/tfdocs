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
					Name:        "id",
					Description: `(Optional) Id of an existing Virtual Data Center that you want to search for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing Virtual Data Center that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing Virtual Data Center's location. Either ` + "`" + `name` + "`" + `, ` + "`" + `location` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The regional location where the Virtual Data Center will be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of that Data Center. Gets incremented with every change`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `List of features supported by the location this data center is part of`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `Boolean value representing if the data center requires extra protection e.g. two factor protection`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The regional location where the Virtual Data Center will be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of that Data Center. Gets incremented with every change`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `List of features supported by the location this data center is part of`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `Boolean value representing if the data center requires extra protection e.g. two factor protection`,
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
					Description: `(Optional) Image aliases`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `(Optional) Cloud init compatibility ("NONE" or "V1") If both "name" and "version" are provided the plugin will concatenate the two strings in this format [name]-[version]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the image`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This indicates the type of image`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the image`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This indicates the type of image`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot.`,
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
					Description: `(Optional) ID of the cluster you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
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
					Name:        "k8s_version",
					Description: `Kubernetes version`,
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
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubernetes version. The provider will ignore changes of patch level.`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `list of available versions for upgrading the cluster`,
				},
				resource.Attribute{
					Name:        "viable_node_pool_versions",
					Description: `list of versions that may be used for node pools under this cluster`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `list of the IDs of the node pools in this cluster`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Raw Kubernetes configuration; use ` + "`" + `yamlencode` + "`" + ` or ` + "`" + `jsonencode` + "`" + ` when dumping this to a file`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `the indicator if the cluster is public or private`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `the IP address of the gateway used by the cluster`,
				},
				resource.Attribute{
					Name:        "api_subnet_allow_list",
					Description: `access to the K8s API server is restricted to these CIDRs`,
				},
				resource.Attribute{
					Name:        "s3_buckets",
					Description: `list of S3 bucket configured for K8s usage`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `structured kubernetes config consisting of a list with 1 item with the following fields:`,
				},
				resource.Attribute{
					Name:        "user_tokens",
					Description: `a convenience map to search the token of a specific user - key - is the user name - value - is the token`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `cluster server (same as ` + "`" + `config[0].clusters[0].cluster.server` + "`" + ` but provided as an attribute for ease of use)`,
				},
				resource.Attribute{
					Name:        "ca_crt",
					Description: `base64 decoded cluster certificate authority data (provided as an attribute for direct use)`,
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
					Name:        "k8s_version",
					Description: `Kubernetes version`,
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
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubernetes version. The provider will ignore changes of patch level.`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `list of available versions for upgrading the cluster`,
				},
				resource.Attribute{
					Name:        "viable_node_pool_versions",
					Description: `list of versions that may be used for node pools under this cluster`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `list of the IDs of the node pools in this cluster`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Raw Kubernetes configuration; use ` + "`" + `yamlencode` + "`" + ` or ` + "`" + `jsonencode` + "`" + ` when dumping this to a file`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `the indicator if the cluster is public or private`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `the IP address of the gateway used by the cluster`,
				},
				resource.Attribute{
					Name:        "api_subnet_allow_list",
					Description: `access to the K8s API server is restricted to these CIDRs`,
				},
				resource.Attribute{
					Name:        "s3_buckets",
					Description: `list of S3 bucket configured for K8s usage`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `structured kubernetes config consisting of a list with 1 item with the following fields:`,
				},
				resource.Attribute{
					Name:        "user_tokens",
					Description: `a convenience map to search the token of a specific user - key - is the user name - value - is the token`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `cluster server (same as ` + "`" + `config[0].clusters[0].cluster.server` + "`" + ` but provided as an attribute for ease of use)`,
				},
				resource.Attribute{
					Name:        "ca_crt",
					Description: `base64 decoded cluster certificate authority data (provided as an attribute for direct use)`,
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
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
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
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
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
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the LAN.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of lan's Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "ip_failover",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "pcc",
					Description: `The unique id of a ` + "`" + `ionoscloud_private_crossconnect` + "`" + ` resource, in order.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the LAN faces the public Internet (true) or not (false).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the LAN.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of lan's Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "ip_failover",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "pcc",
					Description: `The unique id of a ` + "`" + `ionoscloud_private_crossconnect` + "`" + ` resource, in order.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the LAN faces the public Internet (true) or not (false).`,
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
					Description: `(Optional) A desired feature that the location must be able to provide. ## Attributes Reference The following attributes are returned by the datasource:`,
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
					Name:        "id",
					Description: `The id of the private cross-connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the private cross-connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the private cross-connection.`,
				},
				resource.Attribute{
					Name:        "peers",
					Description: `Lists LAN's joined to this private cross connect`,
				},
				resource.Attribute{
					Name:        "connectable_datacenters",
					Description: `Lists datacenters that can be joined to this private cross connect`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the private cross-connection.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the private cross-connection.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description for the private cross-connection.`,
				},
				resource.Attribute{
					Name:        "peers",
					Description: `Lists LAN's joined to this private cross connect`,
				},
				resource.Attribute{
					Name:        "connectable_datacenters",
					Description: `Lists datacenters that can be joined to this private cross connect`,
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
					Name:        "id",
					Description: `The id of the server`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `Number of server CPU cores`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory for the server in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the server should exist`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `Status of the virtual Machine`,
				},
				resource.Attribute{
					Name:        "cdroms",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of cdrom`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of image`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached volume`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Hardware type of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the volume should exist`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Image or snapshot ID to be used as template for this volume`,
				},
				resource.Attribute{
					Name:        "image_alias",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "image_password",
					Description: `Initial password to be set for installed OS`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The bus type of the volume`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this volume`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The Logical Unit Number of the storage volume`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached nic`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached nid`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses assigned to a nic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the nic will reserve an IP using DHCP`,
				},
				resource.Attribute{
					Name:        "lan",
					Description: `The LAN ID the NIC will sit on`,
				},
				resource.Attribute{
					Name:        "firewall_active",
					Description: `Activate or deactivate the firewall`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Indicates if NAT is enabled on this NIC. This is now deprecated.`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the firewall rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the firewall rule`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `he protocol for the rule`,
				},
				resource.Attribute{
					Name:        "source_mac",
					Description: `Only traffic originating from the respective MAC address is allowed`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs`,
				},
				resource.Attribute{
					Name:        "target_ip",
					Description: `In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `Defines the allowed code (from 0 to 254) if protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_start",
					Description: `Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_end",
					Description: `Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the server`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the server`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `Number of server CPU cores`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory for the server in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the server should exist`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `Status of the virtual Machine`,
				},
				resource.Attribute{
					Name:        "cdroms",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of cdrom`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of image`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached volume`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Hardware type of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the volume should exist`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Image or snapshot ID to be used as template for this volume`,
				},
				resource.Attribute{
					Name:        "image_alias",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "image_password",
					Description: `Initial password to be set for installed OS`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The bus type of the volume`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this volume`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The Logical Unit Number of the storage volume`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached nic`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached nid`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses assigned to a nic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the nic will reserve an IP using DHCP`,
				},
				resource.Attribute{
					Name:        "lan",
					Description: `The LAN ID the NIC will sit on`,
				},
				resource.Attribute{
					Name:        "firewall_active",
					Description: `Activate or deactivate the firewall`,
				},
				resource.Attribute{
					Name:        "nat",
					Description: `Indicates if NAT is enabled on this NIC. This is now deprecated.`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the firewall rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the firewall rule`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `he protocol for the rule`,
				},
				resource.Attribute{
					Name:        "source_mac",
					Description: `Only traffic originating from the respective MAC address is allowed`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs`,
				},
				resource.Attribute{
					Name:        "target_ip",
					Description: `In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `Defines the allowed code (from 0 to 254) if protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_start",
					Description: `Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_end",
					Description: `Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen`,
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
					Description: `(Optional) The size of the snapshot to look for. ## Attributes Reference The following attributes are returned by the datasource:`,
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
