package flexibleengine

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_apig_environments",
			Category:         "API Gateway (Dedicated APIG)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"api",
				"gateway",
				"dedicated",
				"apig",
				"environments",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to query the APIG environment list. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Required, String) Specifies an ID of the APIG dedicated instance to which the API environment belongs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the API environment. The API environment name consists of 3 to 64 characters, starting with a letter. Only letters, digits and underscores (_) are allowed. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Data source ID.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `List of APIG environment details. The structure is documented below. The ` + "`" + `environments` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the APIG environment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The environment name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description about the API environment.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the APIG environment was created, in RFC-3339 format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Data source ID.`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `List of APIG environment details. The structure is documented below. The ` + "`" + `environments` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the APIG environment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The environment name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description about the API environment.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Time when the APIG environment was created, in RFC-3339 format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_availability_zones",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The ` + "`" + `region` + "`" + ` to fetch availability zones from, defaults to the provider's ` + "`" + `region` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The ` + "`" + `state` + "`" + ` of the availability zones to match, default ("available"). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_blockstorage_availability_zones_v3",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"blockstorage",
				"availability",
				"zones",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the Block Storage client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The ` + "`" + `state` + "`" + ` of the availability zones to match. Can either be ` + "`" + `available` + "`" + ` or ` + "`" + `unavailable` + "`" + `. Default is ` + "`" + `available` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_blockstorage_volume_v2",
			Category:         "Elastic Volume Service (EVS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Storage.svg",
			Keywords: []string{
				"elastic",
				"volume",
				"service",
				"evs",
				"blockstorage",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Volume client. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the volume.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the volume. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cbr_vaults",
			Category:         "Cloud Backup and Recovery (CBR)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"backup",
				"and",
				"recovery",
				"cbr",
				"vaults",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in hashcode format.`,
				},
				resource.Attribute{
					Name:        "vaults",
					Description: `List of CBR vault details. The object structure of each CBR vault is documented below. The ` + "`" + `vaults` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The vault ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The CBR vault name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The object type of the CBR vault.`,
				},
				resource.Attribute{
					Name:        "consistent_level",
					Description: `The backup specifications. The valid value is`,
				},
				resource.Attribute{
					Name:        "protection_type",
					Description: `The protection type of the CBR vault.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The vault capacity, in GB.`,
				},
				resource.Attribute{
					Name:        "auto_expand_enabled",
					Description: `Whether to enable automatic expansion of the backup protection type vault.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `The policy associated with the CBR vault.`,
				},
				resource.Attribute{
					Name:        "allocated",
					Description: `The allocated capacity of the vault, in GB.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `The used capacity, in GB.`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `The specification code.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The vault status.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `The name of the bucket for the vault.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The key/value pairs to associate with the CBR vault.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `An array of one or more resources to attach to the CBR vault. The object structure is documented below. The ` + "`" + `resources` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The ID of the ECS instance to be backed up.`,
				},
				resource.Attribute{
					Name:        "includes",
					Description: `An array of disk or SFS file system IDs which will be included in the backup.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_addon_template",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"addon",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, String) Specifies the ID of CCE cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) Specifies the add-on name. The supported addons are as follows: +`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required, String) Specifies the add-on version. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of the addon template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the add-on.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The detail configuration of the add-on template.`,
				},
				resource.Attribute{
					Name:        "stable",
					Description: `Whether the add-on template is a stable version.`,
				},
				resource.Attribute{
					Name:        "support_version/virtual_machine",
					Description: `The cluster (Virtual Machine) version that the add-on template supported.`,
				},
				resource.Attribute{
					Name:        "support_version/bare_metal",
					Description: `The cluster (Bare Metal) version that the add-on template supported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id of the addon template.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the add-on.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `The detail configuration of the add-on template.`,
				},
				resource.Attribute{
					Name:        "stable",
					Description: `Whether the add-on template is a stable version.`,
				},
				resource.Attribute{
					Name:        "support_version/virtual_machine",
					Description: `The cluster (Virtual Machine) version that the add-on template supported.`,
				},
				resource.Attribute{
					Name:        "support_version/bare_metal",
					Description: `The cluster (Bare Metal) version that the add-on template supported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_cluster_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-CCE.svg",
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"cluster",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional)The Name of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of container cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The state of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of the cluster. Possible values: VirtualMachine, BareMetal or Windows ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cluster in string format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of cluster in string format.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The cluster specification in string format.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `The container network segment.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `The container network type: overlay_l2 , underlay_ipvlan or vpc-router.`,
				},
				resource.Attribute{
					Name:        "service_network_cidr",
					Description: `The service network segment.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC used to create the node.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet used to create the node.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `The ID of the high speed network used to create bare metal nodes.`,
				},
				resource.Attribute{
					Name:        "internal_endpoint",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external_endpoint",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_apig_endpoint",
					Description: `The endpoint of the cluster to be accessed through API Gateway.`,
				},
				resource.Attribute{
					Name:        "billingMode",
					Description: `Charging mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `Authentication mode of the cluster, possible values are x509 and rbac.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `Advanced configuration of master nodes. Structure is documented below. The ` + "`" + `masters` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone (AZ) of the master node.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cluster in string format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of cluster in string format.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The cluster specification in string format.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `The container network segment.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `The container network type: overlay_l2 , underlay_ipvlan or vpc-router.`,
				},
				resource.Attribute{
					Name:        "service_network_cidr",
					Description: `The service network segment.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC used to create the node.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet used to create the node.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `The ID of the high speed network used to create bare metal nodes.`,
				},
				resource.Attribute{
					Name:        "internal_endpoint",
					Description: `The internal network address.`,
				},
				resource.Attribute{
					Name:        "external_endpoint",
					Description: `The external network address.`,
				},
				resource.Attribute{
					Name:        "external_apig_endpoint",
					Description: `The endpoint of the cluster to be accessed through API Gateway.`,
				},
				resource.Attribute{
					Name:        "billingMode",
					Description: `Charging mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `Authentication mode of the cluster, possible values are x509 and rbac.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `Advanced configuration of master nodes. Structure is documented below. The ` + "`" + `masters` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone (AZ) of the master node.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_clusters",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"clusters",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the CCE clusters. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional, String) Specifies the ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional, String) Specifies the type of the cluster. Possible values:`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String) Specifies the VPC ID to which the cluster belongs.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `(Optional, String) Specifies the enterprise project ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies the status of the cluster. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `Indicates a list of IDs of all CCE clusters found.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `Indicates a list of CCE clusters found. Structure is documented below. The ` + "`" + `clusters` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster. Possible values:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The specification of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `The charging mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `The container network segment.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `The container network type:`,
				},
				resource.Attribute{
					Name:        "eni_subnet_id",
					Description: `The ENI subnet ID.`,
				},
				resource.Attribute{
					Name:        "eni_subnet_cidr",
					Description: `The ENI network segment.`,
				},
				resource.Attribute{
					Name:        "service_network_cidr",
					Description: `The service network segment.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `The authentication mode of the cluster, possible values are x509 and rbac. Defaults to`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `The advanced configuration of master nodes.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The vpc ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet used to create the node.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `The ID of the high speed network used to create bare metal nodes.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `The enterprise project ID of the CCE cluster.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `The access addresses of kube-apiserver in the cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters",
					Description: `The certificate clusters. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_users",
					Description: `The certificate users. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "kube_config_raw",
					Description: `The raw Kubernetes config to be used by kubectl and other compatible tools. The ` + "`" + `endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the cluster access address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the cluster access address. +`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_authority_data",
					Description: `The certificate data. The ` + "`" + `certificate_users` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "client_key_data",
					Description: `The client key data.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `Indicates a list of IDs of all CCE clusters found.`,
				},
				resource.Attribute{
					Name:        "clusters",
					Description: `Indicates a list of CCE clusters found. Structure is documented below. The ` + "`" + `clusters` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `The type of the cluster. Possible values:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the cluster.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The specification of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the cluster.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `The charging mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `The container network segment.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `The container network type:`,
				},
				resource.Attribute{
					Name:        "eni_subnet_id",
					Description: `The ENI subnet ID.`,
				},
				resource.Attribute{
					Name:        "eni_subnet_cidr",
					Description: `The ENI network segment.`,
				},
				resource.Attribute{
					Name:        "service_network_cidr",
					Description: `The service network segment.`,
				},
				resource.Attribute{
					Name:        "authentication_mode",
					Description: `The authentication mode of the cluster, possible values are x509 and rbac. Defaults to`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `The advanced configuration of master nodes.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The security group ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The vpc ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet used to create the node.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `The ID of the high speed network used to create bare metal nodes.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `The enterprise project ID of the CCE cluster.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `The access addresses of kube-apiserver in the cluster. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_clusters",
					Description: `The certificate clusters. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "certificate_users",
					Description: `The certificate users. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "kube_config_raw",
					Description: `The raw Kubernetes config to be used by kubectl and other compatible tools. The ` + "`" + `endpoints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The URL of the cluster access address.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the cluster access address. +`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The cluster name.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The server IP address.`,
				},
				resource.Attribute{
					Name:        "certificate_authority_data",
					Description: `The certificate data. The ` + "`" + `certificate_users` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The user name.`,
				},
				resource.Attribute{
					Name:        "client_certificate_data",
					Description: `The client certificate data.`,
				},
				resource.Attribute{
					Name:        "client_key_data",
					Description: `The client key data.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_node_ids_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description: `

` + "`" + `flexibleengine_cce_node_ids_v3` + "`" + ` provides a list of node ids for a CCE cluster.
This data source can be useful for getting back a list of node ids for a CCE cluster.

`,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"node",
				"ids",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the node ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the node ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cce_node_v3",
			Category:         "Cloud Container Engine (CCE)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"container",
				"engine",
				"cce",
				"node",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The id of container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of the node.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) - The id of the node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) - The state of the node. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference:`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The flavor id to be used.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Available partitions where the node is located.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `Key pair name when logging in to select the key pair mode.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `Node's billing mode: The value is 0 (on demand).`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `List of existing elastic IP IDs.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The node's virtual machine ID in ECS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the node`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Elastic IP parameters of the node.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `Elastic IP address type.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `Bandwidth sharing type.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `Bandwidth (Mbit/s), in the range of [1, 2000].`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `Bandwidth billing type.`,
				},
				resource.Attribute{
					Name:        "root_volume",
					Description: `It corresponds to the system disk related configuration. + ` + "`" + `disk_size` + "`" + ` - Disk size in GB. + ` + "`" + `volume_type` + "`" + ` - Disk type.`,
				},
				resource.Attribute{
					Name:        "data_volumes",
					Description: `Represents the data disk to be created. + ` + "`" + `disk_size` + "`" + ` - Disk size in GB. + ` + "`" + `volume_type` + "`" + ` - Disk type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The flavor id to be used.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Available partitions where the node is located.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `Key pair name when logging in to select the key pair mode.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `Node's billing mode: The value is 0 (on demand).`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `List of existing elastic IP IDs.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The node's virtual machine ID in ECS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the node`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Elastic IP parameters of the node.`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `Elastic IP address type.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `Bandwidth sharing type.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `Bandwidth (Mbit/s), in the range of [1, 2000].`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `Bandwidth billing type.`,
				},
				resource.Attribute{
					Name:        "root_volume",
					Description: `It corresponds to the system disk related configuration. + ` + "`" + `disk_size` + "`" + ` - Disk size in GB. + ` + "`" + `volume_type` + "`" + ` - Disk type.`,
				},
				resource.Attribute{
					Name:        "data_volumes",
					Description: `Represents the data disk to be created. + ` + "`" + `disk_size` + "`" + ` - Disk size in GB. + ` + "`" + `volume_type` + "`" + ` - Disk type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_availability_zones_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"compute",
				"availability",
				"zones",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The ` + "`" + `region` + "`" + ` to fetch availability zones from, defaults to the provider's ` + "`" + `region` + "`" + ``,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The ` + "`" + `state` + "`" + ` of the availability zones to match, default ("available"). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_bms_flavors_v2",
			Category:         "Bare Metal Server (BMS)",
			ShortDescription: ``,
			Description: `

Use this data source to get an available BMS Flavor.

`,
			Keywords: []string{
				"bare",
				"metal",
				"server",
				"bms",
				"compute",
				"flavors",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The BMS flavor id.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The memory size (in MB) of the BMS flavor.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The disk size (GB) in the BMS flavor.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `This is a reserved attribute.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The BMS flavor id.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The memory size (in MB) of the BMS flavor.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `The disk size (GB) in the BMS flavor.`,
				},
				resource.Attribute{
					Name:        "swap",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "rx_tx_factor",
					Description: `This is a reserved attribute.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_bms_keypairs_v2",
			Category:         "Bare Metal Server (BMS)",
			ShortDescription: ``,
			Description: `

` + "`" + `flexibleengine_compute_bms_keypairs_v2` + "`" + ` used to query SSH key pairs.

`,
			Keywords: []string{
				"bare",
				"metal",
				"server",
				"bms",
				"compute",
				"keypairs",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - It is the key pair name. ## Attributes Reference All of the argument attributes are also exported as result attributes.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `It gives the information about the public key in the key pair.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `It is the fingerprint information about the key pair.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "public_key",
					Description: `It gives the information about the public key in the key pair.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `It is the fingerprint information about the key pair.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_bms_nic_v2",
			Category:         "Bare Metal Server (BMS)",
			ShortDescription: ``,
			Description: `

` + "`" + `flexibleengine_compute_bms_nic_v2` + "`" + ` used to query information about a BMS NIC based on the NIC ID.

`,
			Keywords: []string{
				"bare",
				"metal",
				"server",
				"bms",
				"compute",
				"nic",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) - This is the unique BMS id.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) - The ID of the NIC.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) - The NIC port status. ## Attributes Reference All of the argument attributes are also exported as result attributes.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `It is NIC's mac address.`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `The NIC IP address.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the network to which the NIC port belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "mac_address",
					Description: `It is NIC's mac address.`,
				},
				resource.Attribute{
					Name:        "fixed_ips",
					Description: `The NIC IP address.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the network to which the NIC port belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_bms_server_v2",
			Category:         "Bare Metal Server (BMS)",
			ShortDescription: ``,
			Description: `

` + "`" + `flexibleengine_compute_bms_server_v2` + "`" + ` used to query a BMS or BMSs details.

`,
			Keywords: []string{
				"bare",
				"metal",
				"server",
				"bms",
				"compute",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) - The unique ID of the BMS.`,
				},
				resource.Attribute{
					Name:        "host_id",
					Description: `It is the host ID of the BMS.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The BMS metadata is specified.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `It gives the BMS network address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The list of security groups to which the BMS belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Specifies the BMS tag.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `It specifies whether a BMS is locked, true: The BMS is locked, false: The BMS is not locked.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Specifies the AZ ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Provides supplementary information about the pool.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The UUID of the kernel image when the AMI image is used.`,
				},
				resource.Attribute{
					Name:        "hypervisor_hostname",
					Description: `It is the name of a host on the hypervisor.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Instance name is specified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "host_id",
					Description: `It is the host ID of the BMS.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The BMS metadata is specified.`,
				},
				resource.Attribute{
					Name:        "access_ip_v4",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "access_ip_v6",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `It gives the BMS network address.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `The list of security groups to which the BMS belongs.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Specifies the BMS tag.`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `It specifies whether a BMS is locked, true: The BMS is locked, false: The BMS is not locked.`,
				},
				resource.Attribute{
					Name:        "config_drive",
					Description: `This is a reserved attribute.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Specifies the AZ ID.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Provides supplementary information about the pool.`,
				},
				resource.Attribute{
					Name:        "kernel_id",
					Description: `The UUID of the kernel image when the AMI image is used.`,
				},
				resource.Attribute{
					Name:        "hypervisor_hostname",
					Description: `It is the name of a host on the hypervisor.`,
				},
				resource.Attribute{
					Name:        "instance_name",
					Description: `Instance name is specified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_flavors_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"flavors",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to obtain the flavors. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) Specifies the AZ name.`,
				},
				resource.Attribute{
					Name:        "performance_type",
					Description: `(Optional, String) Specifies the ECS flavor type.`,
				},
				resource.Attribute{
					Name:        "generation",
					Description: `(Optional, String) Specifies the generation of an ECS type.`,
				},
				resource.Attribute{
					Name:        "cpu_core",
					Description: `(Optional, Int) Specifies the number of vCPUs in the ECS flavor.`,
				},
				resource.Attribute{
					Name:        "memory_size",
					Description: `(Optional, Int) Specifies the memory size(GB) in the ECS flavor. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `A list of flavors.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `A list of flavors.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_instance_v2",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Computing-ECS.svg",
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"instance",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the server instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the server name, which can be queried with a regular expression.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies the IPv4 addresses of the server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional) Specifies the flavor ID. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone where the instance is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The image name of the instance.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `The flavor name of the instance.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `The key pair that is used to authenticate the instance.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `The EIP address that is associted to the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `The system disk voume ID.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The user data (information after encoding) configured during instance creation.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `An array of one or more security group names to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `An array of one or more networks to attach to the instance. The network object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "block_device",
					Description: `An array of one or more disks to attach to the instance. The block_device object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `The scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The network UUID to attach to the server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port ID corresponding to the IP address on that network.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `The fixed IPv4 address of the instance on this network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `The Fixed IPv6 address of the instance on that network. The ` + "`" + `block_device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `The volume boot index on that attachment.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size on that attachment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The volume type on that attachment.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `The volume pci address on that attachment. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The UUID of a Server Group where the instance will be placed into.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone where the instance is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The image name of the instance.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `The flavor name of the instance.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `The key pair that is used to authenticate the instance.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `The EIP address that is associted to the instance.`,
				},
				resource.Attribute{
					Name:        "system_disk_id",
					Description: `The system disk voume ID.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The user data (information after encoding) configured during instance creation.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `An array of one or more security group names to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `An array of one or more networks to attach to the instance. The network object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "block_device",
					Description: `An array of one or more disks to attach to the instance. The block_device object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `The scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The network UUID to attach to the server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port ID corresponding to the IP address on that network.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `The fixed IPv4 address of the instance on this network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `The Fixed IPv6 address of the instance on that network. The ` + "`" + `block_device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "boot_index",
					Description: `The volume boot index on that attachment.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The volume size on that attachment.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The volume type on that attachment.`,
				},
				resource.Attribute{
					Name:        "pci_address",
					Description: `The volume pci address on that attachment. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The UUID of a Server Group where the instance will be placed into.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_compute_instances",
			Category:         "Elastic Cloud Server (ECS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"cloud",
				"server",
				"ecs",
				"compute",
				"instances",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the server instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the server name, which can be queried with a regular expression.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `(Optional) Specifies the IPv4 addresses of the server.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `(Optional) Specifies the flavor ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specifies the status of the instance. The valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `(Optional) Specifies the flavor name of the instance.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `(Optional) Specifies the image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional) Specifies the availability zone where the instance is located.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `(Optional) Specifies the key pair that is used to authenticate the instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Data source ID.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `List of ECS instance details. The object structure of each ECS instance is documented below. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone where the instance is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The image name of the instance.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The flavor ID of the instance.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `The flavor name of the instance.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `The key pair that is used to authenticate the instance.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `The EIP address that is associted to the instance.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The user data (information after encoding) configured during instance creation.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `An array of one or more security group names to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `An array of one or more networks to attach to the instance. The network object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "volume_attached",
					Description: `An array of one or more disks to attach to the instance. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `The scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The network UUID to attach to the server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port ID corresponding to the IP address on that network.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `The fixed IPv4 address of the instance on this network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `The Fixed IPv6 address of the instance on that network. The ` + "`" + `volume_attached` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "is_sys_volume",
					Description: `Whether the volume is the system disk. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The UUID of a Server Group where the instance will be placed into.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Data source ID.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `List of ECS instance details. The object structure of each ECS instance is documented below. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone where the instance is located.`,
				},
				resource.Attribute{
					Name:        "image_id",
					Description: `The image ID of the instance.`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `The image name of the instance.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The flavor ID of the instance.`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `The flavor name of the instance.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `The key pair that is used to authenticate the instance.`,
				},
				resource.Attribute{
					Name:        "floating_ip",
					Description: `The EIP address that is associted to the instance.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The user data (information after encoding) configured during instance creation.`,
				},
				resource.Attribute{
					Name:        "security_groups",
					Description: `An array of one or more security group names to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `An array of one or more networks to attach to the instance. The network object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "volume_attached",
					Description: `An array of one or more disks to attach to the instance. The object structure is documented below.`,
				},
				resource.Attribute{
					Name:        "scheduler_hints",
					Description: `The scheduler with hints on how the instance should be launched. The available hints are described below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata of the instance in key/value format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the instance. The ` + "`" + `network` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The network UUID to attach to the server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port ID corresponding to the IP address on that network.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC on that network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v4",
					Description: `The fixed IPv4 address of the instance on this network.`,
				},
				resource.Attribute{
					Name:        "fixed_ip_v6",
					Description: `The Fixed IPv6 address of the instance on that network. The ` + "`" + `volume_attached` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `The volume id on that attachment.`,
				},
				resource.Attribute{
					Name:        "is_sys_volume",
					Description: `Whether the volume is the system disk. The ` + "`" + `scheduler_hints` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The UUID of a Server Group where the instance will be placed into.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_csbs_backup_policy_v1",
			Category:         "Cloud Server Backup Service (CSBS)",
			ShortDescription: ``,
			Description: `

The FlexibleEngine CSBS Backup Policy data source allows access of backup Policy resources.

`,
			Icon: "Storage-CSBS.svg",
			Keywords: []string{
				"cloud",
				"server",
				"backup",
				"service",
				"csbs",
				"policy",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Specifies the ID of backup policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the backup policy name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specifies the backup policy status. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the backup policy description.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `Provides the Backup provider ID.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Specifies the parameters of a backup policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the backup policy description.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `Provides the Backup provider ID.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Specifies the parameters of a backup policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_csbs_backup_v1",
			Category:         "Cloud Server Backup Service (CSBS)",
			ShortDescription: ``,
			Description: `

The FlexibleEngine CSBS Backup data source allows access of backup resources.

`,
			Icon: "Storage-CSBS.svg",
			Keywords: []string{
				"cloud",
				"server",
				"backup",
				"service",
				"csbs",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Specifies the ID of backup.`,
				},
				resource.Attribute{
					Name:        "backup_name",
					Description: `(Optional) Specifies the backup name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specifies the backup status.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) Specifies the backup object name.`,
				},
				resource.Attribute{
					Name:        "backup_record_id",
					Description: `(Optional) Specifies the backup record ID.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) Specifies the type of backup objects.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) Specifies the backup object ID.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Specifies the Policy Id.`,
				},
				resource.Attribute{
					Name:        "vm_ip",
					Description: `(Optional) Specifies the ip of VM. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Provides the backup description.`,
				},
				resource.Attribute{
					Name:        "auto_trigger",
					Description: `Specifies whether automatic trigger is enabled.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies average speed.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies the backup capacity.`,
				},
				resource.Attribute{
					Name:        "volume_backups",
					Description: `block supports the following arguments: + ` + "`" + `status` + "`" + ` - Status of backup Volume. + ` + "`" + `space_saving_ratio` + "`" + ` - Specifies the space saving rate. + ` + "`" + `name` + "`" + ` - It gives EVS disk backup name. + ` + "`" + `bootable` + "`" + ` - Specifies whether the disk is bootable. + ` + "`" + `average_speed` + "`" + ` - Specifies the average speed. + ` + "`" + `source_volume_size` + "`" + ` - Shows source volume size in GB. + ` + "`" + `source_volume_id` + "`" + ` - It specifies source volume ID. + ` + "`" + `incremental` + "`" + ` - Shows whether incremental backup is used. + ` + "`" + `snapshot_id` + "`" + ` - ID of snapshot. + ` + "`" + `source_volume_name` + "`" + ` - Specifies source volume name. + ` + "`" + `image_type` + "`" + ` - It specifies backup. The default value is backup. + ` + "`" + `id` + "`" + ` - Specifies Cinder backup ID. + ` + "`" + `size` + "`" + ` - Specifies accumulated size (MB) of backups.`,
				},
				resource.Attribute{
					Name:        "vm_metadata",
					Description: `block supports the following arguments: + ` + "`" + `name` + "`" + ` - Name of backup data. + ` + "`" + `eip` + "`" + ` - Specifies elastic IP address of the ECS. + ` + "`" + `cloud_service_type` + "`" + ` - Specifies ECS type. + ` + "`" + `ram` + "`" + ` - Specifies memory size of the ECS, in MB. + ` + "`" + `vcpus` + "`" + ` - Specifies CPU cores corresponding to the ECS. + ` + "`" + `private_ip` + "`" + ` - It specifies internal IP address of the ECS. + ` + "`" + `disk` + "`" + ` - Shows system disk size corresponding to the ECS specifications. + ` + "`" + `image_type` + "`" + ` - Specifies image type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Provides the backup description.`,
				},
				resource.Attribute{
					Name:        "auto_trigger",
					Description: `Specifies whether automatic trigger is enabled.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies average speed.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies the backup capacity.`,
				},
				resource.Attribute{
					Name:        "volume_backups",
					Description: `block supports the following arguments: + ` + "`" + `status` + "`" + ` - Status of backup Volume. + ` + "`" + `space_saving_ratio` + "`" + ` - Specifies the space saving rate. + ` + "`" + `name` + "`" + ` - It gives EVS disk backup name. + ` + "`" + `bootable` + "`" + ` - Specifies whether the disk is bootable. + ` + "`" + `average_speed` + "`" + ` - Specifies the average speed. + ` + "`" + `source_volume_size` + "`" + ` - Shows source volume size in GB. + ` + "`" + `source_volume_id` + "`" + ` - It specifies source volume ID. + ` + "`" + `incremental` + "`" + ` - Shows whether incremental backup is used. + ` + "`" + `snapshot_id` + "`" + ` - ID of snapshot. + ` + "`" + `source_volume_name` + "`" + ` - Specifies source volume name. + ` + "`" + `image_type` + "`" + ` - It specifies backup. The default value is backup. + ` + "`" + `id` + "`" + ` - Specifies Cinder backup ID. + ` + "`" + `size` + "`" + ` - Specifies accumulated size (MB) of backups.`,
				},
				resource.Attribute{
					Name:        "vm_metadata",
					Description: `block supports the following arguments: + ` + "`" + `name` + "`" + ` - Name of backup data. + ` + "`" + `eip` + "`" + ` - Specifies elastic IP address of the ECS. + ` + "`" + `cloud_service_type` + "`" + ` - Specifies ECS type. + ` + "`" + `ram` + "`" + ` - Specifies memory size of the ECS, in MB. + ` + "`" + `vcpus` + "`" + ` - Specifies CPU cores corresponding to the ECS. + ` + "`" + `private_ip` + "`" + ` - It specifies internal IP address of the ECS. + ` + "`" + `disk` + "`" + ` - Shows system disk size corresponding to the ECS specifications. + ` + "`" + `image_type` + "`" + ` - Specifies image type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_cts_tracker_v1",
			Category:         "Cloud Trace Service (CTS)",
			ShortDescription: ``,
			Description: `

CTS Tracker data source allows access of Cloud Tracker.

`,
			Icon: "Management&Deployment-CTS.svg",
			Keywords: []string{
				"cloud",
				"trace",
				"service",
				"cts",
				"tracker",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tracker_name",
					Description: `(Optional) The tracker name.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Optional) The OBS bucket name for a tracker.`,
				},
				resource.Attribute{
					Name:        "file_prefix_name",
					Description: `(Optional) The prefix of a log that needs to be stored in an OBS bucket.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of a tracker. ## Attributes Reference All above argument parameters can be exported as attribute parameters.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dcs_az_v1",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"dcs",
				"az",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Indicates the name of an AZ.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Optional) Indicates the code of an AZ.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Indicates the port number of an AZ. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found az. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dcs_maintainwindow_v1",
			Category:         "Distributed Cache Service (DCS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"cache",
				"service",
				"dcs",
				"maintainwindow",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default",
					Description: `(Optional) Specifies whether a maintenance time window is set to the default time segment.`,
				},
				resource.Attribute{
					Name:        "seq",
					Description: `(Optional) Specifies the sequential number of a maintenance time window.`,
				},
				resource.Attribute{
					Name:        "begin",
					Description: `(Optional) Specifies the time at which a maintenance time window starts.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Optional) Specifies the time at which a maintenance time window ends. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found maintainwindow. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
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
			Type:             "flexibleengine_dcs_product_v1",
			Category:         "Distributed Cache Service (DCS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"cache",
				"service",
				"dcs",
				"product",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional, String) The engine of the cache instance. Valid values are`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, String) The version of a cache engine. It is valid when the engine is`,
				},
				resource.Attribute{
					Name:        "cache_mode",
					Description: `(Optional, String) The mode of a cache engine. The valid values are as follows: + ` + "`" + `single` + "`" + ` - Single-node. + ` + "`" + `ha` + "`" + ` - Master/Standby. + ` + "`" + `cluster` + "`" + ` - Redis Cluster, it is valid when the engine is`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `(Optional, Float) The total memory of the cache, in GB. It is valid when the engine is redis 4.0/5.0. + Single-node and Master/Standby instances support: ` + "`" + `0.125` + "`" + `, ` + "`" + `0.25` + "`" + `, ` + "`" + `0.5` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `48` + "`" + ` and ` + "`" + `64` + "`" + `. + Redis Cluster and Proxy Cluster instances support: ` + "`" + `4` + "`" + `, ` + "`" + `8` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `48` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `96` + "`" + `, ` + "`" + `128` + "`" + `, ` + "`" + `192` + "`" + `, ` + "`" + `256` + "`" + `, ` + "`" + `384` + "`" + `, ` + "`" + `512` + "`" + `, ` + "`" + `768` + "`" + ` and ` + "`" + `1024` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "replica_count",
					Description: `(Optional, Int) The number of replicas includes the master. It is valid when the engine is redis 4.0/5.0 with`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `(Optional, String) Specifies the DCS instance specification code. You can log in to the DCS console, click`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The found product ID.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `The CPU architecture of DCS instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The found product ID.`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `The CPU architecture of DCS instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dds_flavor_v3",
			Category:         "Document Database Service (DDS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"document",
				"database",
				"service",
				"dds",
				"flavor",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V3 dds client.`,
				},
				resource.Attribute{
					Name:        "engine_name",
					Description: `(Optional) The engine name of the dds, now only DDS-Community is supported.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `(Optional) The spec code of a dds flavor. ## Available value for attributes engine_name | type | vcpus | ram | speccode ---- | --- | --- DDS-Community | mongos | 1 | 4 | dds.mongodb.s3.medium.4.mongos DDS-Community | mongos | 2 | 8 | dds.mongodb.s3.large.4.mongos DDS-Community | mongos | 4 | 16 | dds.mongodb.s3.xlarge.4.mongos DDS-Community | mongos | 8 | 32 | dds.mongodb.s3.2xlarge.4.mongos DDS-Community | mongos | 16 | 64 | dds.mongodb.s3.4xlarge.4.mongos DDS-Community | shard | 1 | 4 | dds.mongodb.s3.medium.4.shard DDS-Community | shard | 2 | 8 | dds.mongodb.s3.large.4.shard DDS-Community | shard | 4 | 16 | dds.mongodb.s3.xlarge.4.shard DDS-Community | shard | 8 | 32 | dds.mongodb.s3.2xlarge.4.shard DDS-Community | shard | 16 | 64 | dds.mongodb.s3.4xlarge.4.shard DDS-Community | config | 2 | 4 | dds.mongodb.s3.large.2.config DDS-Community | replica | 1 | 4 | dds.mongodb.s3.medium.4.repset DDS-Community | replica | 2 | 8 | dds.mongodb.s3.large.4.repset DDS-Community | replica | 4 | 16 | dds.mongodb.s3.xlarge.4.repset DDS-Community | replica | 8 | 32 | dds.mongodb.s3.2xlarge.4.repset DDS-Community | replica | 16 | 64 | dds.mongodb.s3.4xlarge.4.repset ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The vcpus of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The ram of the dds flavor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The vcpus of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The ram of the dds flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dds_flavors_v3",
			Category:         "Document Database Service (DDS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"document",
				"database",
				"service",
				"dds",
				"flavors",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_name",
					Description: `(Optional, String) Specifies the engine name of the dds, the default value is "DDS-Community".`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String) Specifies the type of the dds falvor. "mongos", "shard", "config", "replica" and "single" are supported.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional, String) Specifies the vcpus of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional, String) Specifies the ram of the dds flavor in GB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `The name of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See ` + "`" + `type` + "`" + ` above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See ` + "`" + `vcpus` + "`" + ` above.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `See ` + "`" + `memory` + "`" + ` above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `The name of the dds flavor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See ` + "`" + `type` + "`" + ` above.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `See ` + "`" + `vcpus` + "`" + ` above.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `See ` + "`" + `memory` + "`" + ` above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dms_kafka_instances",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"kafka",
				"instances",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to query the kafka instance list. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, String) Specifies the kafka instance ID to match exactly.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the kafka instance name for data-source queries.`,
				},
				resource.Attribute{
					Name:        "fuzzy_match",
					Description: `(Optional, Bool) Specifies whether to match the instance name fuzzily, the default is a exact match (` + "`" + `flase` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies the kafka instance status for data-source queries.`,
				},
				resource.Attribute{
					Name:        "include_failure",
					Description: `(Optional, Bool) Specifies whether the query results contain instances that failed to create. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The result of the query's list of kafka instances. The structure is documented below. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The instance type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance description.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The list of AZ names.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `The product ID used by the instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The kafka engine version.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `The storage I/O specification.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `The message storage capacity, in GB unit.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The subnet ID to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The security group ID associated with the instance.`,
				},
				resource.Attribute{
					Name:        "manager_user",
					Description: `The username for logging in to the Kafka Manager.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `The access username.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `The time at which a maintenance time window starts, the format is ` + "`" + `HH:mm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `The time at which a maintenance time window ends, the format is ` + "`" + `HH:mm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_public_ip",
					Description: `Whether public access to the instance is enabled.`,
				},
				resource.Attribute{
					Name:        "public_ip_ids",
					Description: `The IDs of the elastic IP address (EIP).`,
				},
				resource.Attribute{
					Name:        "public_conn_addresses",
					Description: `The instance public access address. The format of each connection address is ` + "`" + `{IP address}:{port}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_policy",
					Description: `The action to be taken when the memory usage reaches the disk capacity threshold.`,
				},
				resource.Attribute{
					Name:        "dumping",
					Description: `Whether to dumping is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_auto_topic",
					Description: `Whether to enable automatic topic creation.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `The maximum number of topics in the DMS kafka instance.`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `Whether the Kafka SASL_SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `The used message storage space, in GB unit.`,
				},
				resource.Attribute{
					Name:        "connect_address",
					Description: `The IP address for instance connection.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The instance status.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `The resource specifications identifier.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user ID who created the instance.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The username who created the instance.`,
				},
				resource.Attribute{
					Name:        "management_connect_address",
					Description: `The connection address of the Kafka manager of an instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The key/value pairs to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "cross_vpc_accesses",
					Description: `Indicates the Access information of cross-VPC. The structure is documented below. The ` + "`" + `cross_vpc_accesses` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "listener_ip",
					Description: `The listener IP address.`,
				},
				resource.Attribute{
					Name:        "advertised_ip",
					Description: `The advertised IP Address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `The port ID associated with the address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The result of the query's list of kafka instances. The structure is documented below. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The instance ID.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The instance type.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The instance name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The instance description.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The list of AZ names.`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `The product ID used by the instance.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The kafka engine version.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `The storage I/O specification.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `The message storage capacity, in GB unit.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC ID to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The subnet ID to which the instance belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The security group ID associated with the instance.`,
				},
				resource.Attribute{
					Name:        "manager_user",
					Description: `The username for logging in to the Kafka Manager.`,
				},
				resource.Attribute{
					Name:        "access_user",
					Description: `The access username.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `The time at which a maintenance time window starts, the format is ` + "`" + `HH:mm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `The time at which a maintenance time window ends, the format is ` + "`" + `HH:mm` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_public_ip",
					Description: `Whether public access to the instance is enabled.`,
				},
				resource.Attribute{
					Name:        "public_ip_ids",
					Description: `The IDs of the elastic IP address (EIP).`,
				},
				resource.Attribute{
					Name:        "public_conn_addresses",
					Description: `The instance public access address. The format of each connection address is ` + "`" + `{IP address}:{port}` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_policy",
					Description: `The action to be taken when the memory usage reaches the disk capacity threshold.`,
				},
				resource.Attribute{
					Name:        "dumping",
					Description: `Whether to dumping is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_auto_topic",
					Description: `Whether to enable automatic topic creation.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `The maximum number of topics in the DMS kafka instance.`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `Whether the Kafka SASL_SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `The used message storage space, in GB unit.`,
				},
				resource.Attribute{
					Name:        "connect_address",
					Description: `The IP address for instance connection.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The instance status.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `The resource specifications identifier.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The user ID who created the instance.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The username who created the instance.`,
				},
				resource.Attribute{
					Name:        "management_connect_address",
					Description: `The connection address of the Kafka manager of an instance.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The key/value pairs to associate with the instance.`,
				},
				resource.Attribute{
					Name:        "cross_vpc_accesses",
					Description: `Indicates the Access information of cross-VPC. The structure is documented below. The ` + "`" + `cross_vpc_accesses` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "listener_ip",
					Description: `The listener IP address.`,
				},
				resource.Attribute{
					Name:        "advertised_ip",
					Description: `The advertised IP Address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `The port ID associated with the address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dms_product",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"product",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the DMS products. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Required, String) Specifies the bandwidth of a DMS instance. The valid values are`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional, String) Specifies the name of a message engine. Only`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional, String) Specifies the version of a message engine. The default value is`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The DMS product ID.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The list of availability zones where there are available resources.`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `The DMS product specification, for example, dms.instance.kafka.cluster.c3.small.2.`,
				},
				resource.Attribute{
					Name:        "cpu_arch",
					Description: `The CPU architecture of a DMS instance.`,
				},
				resource.Attribute{
					Name:        "ecs_flavor_id",
					Description: `The flavor of the corresponding ECS.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `The maximum number of topics in a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `The minimum storage capacity of the DMS product.`,
				},
				resource.Attribute{
					Name:        "storage_spec_codes",
					Description: `The list of supported storage specification. The item of the list can be one of`,
				},
				resource.Attribute{
					Name:        "max_tps",
					Description: `The maximum number of messages per unit time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The DMS product ID.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The list of availability zones where there are available resources.`,
				},
				resource.Attribute{
					Name:        "spec_code",
					Description: `The DMS product specification, for example, dms.instance.kafka.cluster.c3.small.2.`,
				},
				resource.Attribute{
					Name:        "cpu_arch",
					Description: `The CPU architecture of a DMS instance.`,
				},
				resource.Attribute{
					Name:        "ecs_flavor_id",
					Description: `The flavor of the corresponding ECS.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `The maximum number of topics in a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `The minimum storage capacity of the DMS product.`,
				},
				resource.Attribute{
					Name:        "storage_spec_codes",
					Description: `The list of supported storage specification. The item of the list can be one of`,
				},
				resource.Attribute{
					Name:        "max_tps",
					Description: `The maximum number of messages per unit time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dms_rocketmq_broker",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"rocketmq",
				"broker",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to query the data source. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, String) Specifies the ID of the rocketMQ instance. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID.`,
				},
				resource.Attribute{
					Name:        "brokers",
					Description: `Indicates the list of the brokers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID.`,
				},
				resource.Attribute{
					Name:        "brokers",
					Description: `Indicates the list of the brokers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dms_rocketmq_instances",
			Category:         "Distributed Message Service (DMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"distributed",
				"message",
				"service",
				"dms",
				"rocketmq",
				"instances",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to query the data source. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `(Optional, String) Specifies the ID of the RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies the status of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "exact_match_name",
					Description: `(Optional, String) Specifies whether to search for the instance that precisely matches a specified instance name. Value options:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `Indicates the list of DMS RocketMQ instances. The [Instance](#DmsRocketMQInstances_Instance) structure is documented below. <a name="DmsRocketMQInstances_Instance"></a> The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the name of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the description of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DMS RocketMQ instance type.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Indicates the instance specification. For a cluster DMS RocketMQ instance, VM specifications and the number of nodes are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Indicates the version of the RocketMQ engine.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the ID of a VPC.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `Indicates a product ID.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the ID of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `Indicates the list of availability zone names, where instance brokers reside and which has available resources.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `Indicates the time at which the maintenance window starts. The format is HH:mm:ss.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `Indicates the time at which the maintenance window ends. The format is HH:mm:ss.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `Indicates the message storage capacity. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `Indicates the used message storage space. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "enable_publicip",
					Description: `Indicates whether to enable public access.`,
				},
				resource.Attribute{
					Name:        "publicip_id",
					Description: `Indicates the ID of the EIP bound to the instance. Use commas (,) to separate multiple EIP IDs. This parameter is mandatory if public access is enabled (that is, enable_publicip is set to true).`,
				},
				resource.Attribute{
					Name:        "publicip_address",
					Description: `Indicates the public IP address.`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `Indicates whether the RocketMQ SASL_SSL is enabled. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "cross_vpc_accesses",
					Description: `Indicates the Cross-VPC access information. The [CrossVpc](#DmsRocketMQInstances_InstanceCrossVpc) structure is documented below.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `Indicates the storage I/O specification.`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Indicates whether to support IPv6. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Indicates the node quantity.`,
				},
				resource.Attribute{
					Name:        "new_spec_billing_enable",
					Description: `Indicates the whether billing based on new specifications is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_acl",
					Description: `Indicates whether access control is enabled.`,
				},
				resource.Attribute{
					Name:        "broker_num",
					Description: `Specifies the broker numbers. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "namesrv_address",
					Description: `Indicates the metadata address.`,
				},
				resource.Attribute{
					Name:        "broker_address",
					Description: `Indicates the service data address.`,
				},
				resource.Attribute{
					Name:        "public_namesrv_address",
					Description: `Indicates the public network metadata address.`,
				},
				resource.Attribute{
					Name:        "public_broker_address",
					Description: `Indicates the public network service data address.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Indicates the resource specifications. <a name="DmsRocketMQInstances_InstanceCrossVpc"></a> The ` + "`" + `cross_vpc_accesses` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "lisenter_ip",
					Description: `Indicates the IP of the listener.`,
				},
				resource.Attribute{
					Name:        "advertised_ip",
					Description: `Indicates the advertised IP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the port.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Indicates the port ID associated with the address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `Indicates the list of DMS RocketMQ instances. The [Instance](#DmsRocketMQInstances_Instance) structure is documented below. <a name="DmsRocketMQInstances_Instance"></a> The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the name of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the status of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the description of the DMS RocketMQ instance.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the DMS RocketMQ instance type.`,
				},
				resource.Attribute{
					Name:        "specification",
					Description: `Indicates the instance specification. For a cluster DMS RocketMQ instance, VM specifications and the number of nodes are returned.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `Indicates the version of the RocketMQ engine.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the ID of a VPC.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `Indicates a product ID.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the ID of a security group.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `Indicates the list of availability zone names, where instance brokers reside and which has available resources.`,
				},
				resource.Attribute{
					Name:        "maintain_begin",
					Description: `Indicates the time at which the maintenance window starts. The format is HH:mm:ss.`,
				},
				resource.Attribute{
					Name:        "maintain_end",
					Description: `Indicates the time at which the maintenance window ends. The format is HH:mm:ss.`,
				},
				resource.Attribute{
					Name:        "storage_space",
					Description: `Indicates the message storage capacity. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "used_storage_space",
					Description: `Indicates the used message storage space. Unit: GB.`,
				},
				resource.Attribute{
					Name:        "enable_publicip",
					Description: `Indicates whether to enable public access.`,
				},
				resource.Attribute{
					Name:        "publicip_id",
					Description: `Indicates the ID of the EIP bound to the instance. Use commas (,) to separate multiple EIP IDs. This parameter is mandatory if public access is enabled (that is, enable_publicip is set to true).`,
				},
				resource.Attribute{
					Name:        "publicip_address",
					Description: `Indicates the public IP address.`,
				},
				resource.Attribute{
					Name:        "ssl_enable",
					Description: `Indicates whether the RocketMQ SASL_SSL is enabled. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "cross_vpc_accesses",
					Description: `Indicates the Cross-VPC access information. The [CrossVpc](#DmsRocketMQInstances_InstanceCrossVpc) structure is documented below.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `Indicates the storage I/O specification.`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Indicates whether to support IPv6. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Indicates the node quantity.`,
				},
				resource.Attribute{
					Name:        "new_spec_billing_enable",
					Description: `Indicates the whether billing based on new specifications is enabled.`,
				},
				resource.Attribute{
					Name:        "enable_acl",
					Description: `Indicates whether access control is enabled.`,
				},
				resource.Attribute{
					Name:        "broker_num",
					Description: `Specifies the broker numbers. Defaults to 1.`,
				},
				resource.Attribute{
					Name:        "namesrv_address",
					Description: `Indicates the metadata address.`,
				},
				resource.Attribute{
					Name:        "broker_address",
					Description: `Indicates the service data address.`,
				},
				resource.Attribute{
					Name:        "public_namesrv_address",
					Description: `Indicates the public network metadata address.`,
				},
				resource.Attribute{
					Name:        "public_broker_address",
					Description: `Indicates the public network service data address.`,
				},
				resource.Attribute{
					Name:        "resource_spec_code",
					Description: `Indicates the resource specifications. <a name="DmsRocketMQInstances_InstanceCrossVpc"></a> The ` + "`" + `cross_vpc_accesses` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "lisenter_ip",
					Description: `Indicates the IP of the listener.`,
				},
				resource.Attribute{
					Name:        "advertised_ip",
					Description: `Indicates the advertised IP.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the port.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `Indicates the port ID associated with the address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dns_zone_v2",
			Category:         "Domain Name Service (DNS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "zone.svg",
			Keywords: []string{
				"domain",
				"name",
				"service",
				"dns",
				"zone",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 DNS client. A DNS client is needed to retrieve zone ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the zone.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) The email contact for the zone record.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The zone's status.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `(Optional) The time to live (TTL) of the zone.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `(Optional) The type of the zone. Can either be ` + "`" + `public` + "`" + ` or ` + "`" + `private` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found zone. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number of the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `The ID of the pool hosting the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID that owns the zone.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ttl",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "masters",
					Description: `An array of master DNS servers.`,
				},
				resource.Attribute{
					Name:        "serial",
					Description: `The serial number of the zone.`,
				},
				resource.Attribute{
					Name:        "pool_id",
					Description: `The ID of the pool hosting the zone.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The project ID that owns the zone.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_dws_flavors",
			Category:         "Data Warehouse Service (DWS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"data",
				"warehouse",
				"service",
				"dws",
				"flavors",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the DWS cluster client. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) Specifies the availability zone name.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional, String) Specifies the vcpus of the DWS node flavor.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional, String) Specifies the ram of the DWS node flavor in GB. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The name of the DWS node flavor. It is referenced by`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Indicates the vcpus of the DWS node flavor.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Indicates the ram of the DWS node flavor in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `Indicates Disk type.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Indicates the Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the availability zone where the node resides.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The name of the DWS node flavor. It is referenced by`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Indicates the vcpus of the DWS node flavor.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Indicates the ram of the DWS node flavor in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `Indicates Disk type.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Indicates the Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the availability zone where the node resides.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_certificate",
			Category:         "Elastic Load Balance (Dedicated ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"dedicated",
				"elb",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to obtain the Dedicated ELB certificate. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, String) The name of certificate. The value is case sensitive and does not supports fuzzy matching. ->`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain of the certificate. This parameter is valid only when ` + "`" + `type` + "`" + ` is "server".`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the certificate type. The value can be one of the following: + ` + "`" + `server` + "`" + `: indicates the server certificate. + ` + "`" + `client` + "`" + `: indicates the CA certificate.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description for the certificate.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Indicates the time when the certificate expires.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The certificate ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `The domain of the certificate. This parameter is valid only when ` + "`" + `type` + "`" + ` is "server".`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Specifies the certificate type. The value can be one of the following: + ` + "`" + `server` + "`" + `: indicates the server certificate. + ` + "`" + `client` + "`" + `: indicates the CA certificate.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-readable description for the certificate.`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `Indicates the time when the certificate expires.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_elb_flavors",
			Category:         "Elastic Load Balance (Dedicated ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"dedicated",
				"elb",
				"flavors",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to obtain the flavors. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String) Specifies the flavor type. Valid values are`,
				},
				resource.Attribute{
					Name:        "max_connections",
					Description: `(Optional, Int) Specifies the maximum connections in the flavor.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional, Int) Specifies the bandwidth size(Mbit/s) in the flavor.`,
				},
				resource.Attribute{
					Name:        "cps",
					Description: `(Optional, Int) Specifies the cps in the flavor.`,
				},
				resource.Attribute{
					Name:        "qps",
					Description: `(Optional, Int) Specifies the qps in the L7 flavor. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of flavor IDs.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `A list of flavors. Each element contains the following attributes: + ` + "`" + `id` + "`" + ` - ID of the flavor. + ` + "`" + `name` + "`" + ` - Name of the flavor. + ` + "`" + `type` + "`" + ` - Type of the flavor. + ` + "`" + `max_connections` + "`" + ` - Maximum connections of the flavor. + ` + "`" + `cps` + "`" + ` - Cps of the flavor. + ` + "`" + `qps` + "`" + ` - Qps of the L7 flavor. + ` + "`" + `bandwidth` + "`" + ` - Bandwidth size(Mbit/s) of the flavor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of flavor IDs.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `A list of flavors. Each element contains the following attributes: + ` + "`" + `id` + "`" + ` - ID of the flavor. + ` + "`" + `name` + "`" + ` - Name of the flavor. + ` + "`" + `type` + "`" + ` - Type of the flavor. + ` + "`" + `max_connections` + "`" + ` - Maximum connections of the flavor. + ` + "`" + `cps` + "`" + ` - Cps of the flavor. + ` + "`" + `qps` + "`" + ` - Qps of the L7 flavor. + ` + "`" + `bandwidth` + "`" + ` - Bandwidth size(Mbit/s) of the flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_enterprise_project",
			Category:         "Enterprise Project Management Service (EPS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"enterprise",
				"project",
				"management",
				"service",
				"eps",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the enterprise project name. Fuzzy search is supported.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional, String) Specifies the ID of an enterprise project. The value "0" indicates the default enterprise project.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, Int) Specifies the status of an enterprise project. + 1 indicates Enabled. + 2 indicates Disabled. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The supplementary information about the enterprise project.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The UTC time when the enterprise project was created. Example: 2018-05-18T06:49:06Z`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The UTC time when the enterprise project was modified. Example: 2018-05-28T02:21:36Z`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The supplementary information about the enterprise project.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The UTC time when the enterprise project was created. Example: 2018-05-18T06:49:06Z`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The UTC time when the enterprise project was modified. Example: 2018-05-28T02:21:36Z`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_fgs_dependencies",
			Category:         "FunctionGraph",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"functiongraph",
				"fgs",
				"dependencies",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the dependent packages. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, String) Specifies the dependent package type to match. Valid values:`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Optional, String) Specifies the dependent package runtime to match. Valid values:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the dependent package runtime to match. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A data source ID.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `All dependent packages that match. + ` + "`" + `id` + "`" + ` - Dependent package ID. + ` + "`" + `name` + "`" + ` - Dependent package name. + ` + "`" + `owner` + "`" + ` - Dependent package owner. + ` + "`" + `link` + "`" + ` - URL of the dependent package in the OBS console. + ` + "`" + `etag` + "`" + ` - Unique ID of the dependent package. + ` + "`" + `size` + "`" + ` - Dependent package size. + ` + "`" + `file_name` + "`" + ` - File name of the Dependent package. + ` + "`" + `runtime` + "`" + ` - Dependent package runtime.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A data source ID.`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `All dependent packages that match. + ` + "`" + `id` + "`" + ` - Dependent package ID. + ` + "`" + `name` + "`" + ` - Dependent package name. + ` + "`" + `owner` + "`" + ` - Dependent package owner. + ` + "`" + `link` + "`" + ` - URL of the dependent package in the OBS console. + ` + "`" + `etag` + "`" + ` - Unique ID of the dependent package. + ` + "`" + `size` + "`" + ` - Dependent package size. + ` + "`" + `file_name` + "`" + ` - File name of the Dependent package. + ` + "`" + `runtime` + "`" + ` - Dependent package runtime.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_gaussdb_cassandra_instances",
			Category:         "GaussDB NoSQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"gaussdb",
				"nosql",
				"cassandra",
				"instances",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to obtain the instance. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String) Specifies the VPC ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, String) Specifies the network ID of a subnet. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the data source.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `An array of available instances. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the name of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the VPC ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the network ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Indicates the instance mode.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Indicates the instance specifications.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the security group ID.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `Indicates the enterprise project id.`,
				},
				resource.Attribute{
					Name:        "db_user_name",
					Description: `Indicates the default username.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the instance availability zone.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Indicates the count of the nodes.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Indicates the size of the volume.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the list of private IP address of the nodes.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `Indicates the database information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `Indicates the advanced backup policy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Indicates the key/value tags of the instance. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Indicates the database engine.`,
				},
				resource.Attribute{
					Name:        "storage_engine",
					Description: `Indicates the database storage engine.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the database version. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Indicates the backup time window.`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `Indicates the number of days to retain the generated The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the node ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Indicates the private IP address of a node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the node status.`,
				},
				resource.Attribute{
					Name:        "support_reduce",
					Description: `Indicates whether the node support reduce.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the availability zone where the node resides.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the data source.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `An array of available instances. The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the name of the instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `Indicates the VPC ID.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Indicates the network ID of a subnet.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the DB instance status.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Indicates the instance mode.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Indicates the instance specifications.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `Indicates the security group ID.`,
				},
				resource.Attribute{
					Name:        "enterprise_project_id",
					Description: `Indicates the enterprise project id.`,
				},
				resource.Attribute{
					Name:        "db_user_name",
					Description: `Indicates the default username.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the instance availability zone.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Indicates the database port.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `Indicates the count of the nodes.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `Indicates the size of the volume.`,
				},
				resource.Attribute{
					Name:        "private_ips",
					Description: `Indicates the list of private IP address of the nodes.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `Indicates the database information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "backup_strategy",
					Description: `Indicates the advanced backup policy. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `Indicates the instance nodes information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Indicates the key/value tags of the instance. The ` + "`" + `datastore` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `Indicates the database engine.`,
				},
				resource.Attribute{
					Name:        "storage_engine",
					Description: `Indicates the database storage engine.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Indicates the database version. The ` + "`" + `backup_strategy` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Indicates the backup time window.`,
				},
				resource.Attribute{
					Name:        "keep_days",
					Description: `Indicates the number of days to retain the generated The ` + "`" + `nodes` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the node ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the node name.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Indicates the private IP address of a node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Indicates the node status.`,
				},
				resource.Attribute{
					Name:        "support_reduce",
					Description: `Indicates whether the node support reduce.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Indicates the availability zone where the node resides.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_gaussdb_nosql_flavors",
			Category:         "GaussDB NoSQL",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"gaussdb",
				"nosql",
				"flavors",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the GaussDB specifications. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Optional, String) Specifies the type of the database engine. The valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `(Optional, String) Specifies the version of the database engine.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional, Int) Specifies the number of vCPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional, Int) Specifies the memory size in gigabytes (GB).`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) Specifies the availability zone (AZ) of the GaussDB specifications. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `The information of the GaussDB specifications. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The spec code of the flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of vCPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory size, in GB.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The type of the database engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version of the database engine.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `All available zones (on sale) for current flavor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `The information of the GaussDB specifications. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The spec code of the flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The number of vCPUs.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory size, in GB.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `The type of the database engine.`,
				},
				resource.Attribute{
					Name:        "engine_version",
					Description: `The version of the database engine.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `All available zones (on sale) for current flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_custom_role_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"custom",
				"role",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the custom policy.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the custom policy.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain the policy belongs to.`,
				},
				resource.Attribute{
					Name:        "references",
					Description: `(Optional) The number of citations for the custom policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the custom policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Display mode. Valid options are AX: Account level and XA: Project level. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `Document of the custom policy.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `The catalog of the custom policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "policy",
					Description: `Document of the custom policy.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `The catalog of the custom policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_group",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the identity group.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional, String) Specifies the ID of the identity group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the identity group. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `Indicates the domain the group belongs to.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Indicates the users the group contains. Structure is documented below. The ` + "`" + `users` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the IAM user name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the User.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the description of the IAM user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates the whether the IAM user is enabled.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `Indicates the time when the password will expire. Null indicates that the password has unlimited validity.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Indicates the password status. True means that the password needs to be changed, and false means that the password is normal.`,
				},
				resource.Attribute{
					Name:        "password_strength",
					Description: `Indicates the password strength. The value can be high, mid, or low.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "domain_id",
					Description: `Indicates the domain the group belongs to.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Indicates the users the group contains. Structure is documented below. The ` + "`" + `users` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the IAM user name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the User.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the description of the IAM user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates the whether the IAM user is enabled.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `Indicates the time when the password will expire. Null indicates that the password has unlimited validity.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Indicates the password status. True means that the password needs to be changed, and false means that the password is normal.`,
				},
				resource.Attribute{
					Name:        "password_strength",
					Description: `Indicates the password strength. The value can be high, mid, or low.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_project_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"project",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the project.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain this project belongs to.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `(Optional) The parent of this project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether the project is available.`,
				},
				resource.Attribute{
					Name:        "is_domain",
					Description: `Whether this project is a domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the project.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether the project is available.`,
				},
				resource.Attribute{
					Name:        "is_domain",
					Description: `Whether this project is a domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_role_v3",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"role",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain the role belongs to. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the role displayed on the console.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `The service catalog of the policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The display mode of the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The content of the policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The display name of the role displayed on the console.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the policy.`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `The service catalog of the policy.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The display mode of the policy.`,
				},
				resource.Attribute{
					Name:        "policy",
					Description: `The content of the policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_identity_users",
			Category:         "Identity and Access Management (IAM)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"identity",
				"and",
				"access",
				"management",
				"iam",
				"users",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the IAM user name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional, String) Specifies the status of the IAM user, the default value is`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The details of the queried IAM users. The structure is documented below. The ` + "`" + `users` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the User.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the IAM user name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the description of the IAM user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates the whether the IAM user is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Indicates the user groups to which an IAM user belongs.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `Indicates the time when the password will expire. Null indicates that the password has unlimited validity.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Indicates the password status. True means that the password needs to be changed, and false means that the password is normal.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `The details of the queried IAM users. The structure is documented below. The ` + "`" + `users` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates the ID of the User.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Indicates the IAM user name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Indicates the description of the IAM user.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Indicates the whether the IAM user is enabled.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Indicates the user groups to which an IAM user belongs.`,
				},
				resource.Attribute{
					Name:        "password_expires_at",
					Description: `Indicates the time when the password will expire. Null indicates that the password has unlimited validity.`,
				},
				resource.Attribute{
					Name:        "password_status",
					Description: `Indicates the password status. True means that the password needs to be changed, and false means that the password is normal.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_images_image_v2",
			Category:         "Image Management Service (IMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"image",
				"management",
				"service",
				"ims",
				"images",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the image. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the image. Exact matching is used.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The owner (UUID) of the image.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `(Optional) The minimum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "size_max",
					Description: `(Optional) The maximum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort images based on a certain key. Defaults to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Search for images with a specific tag.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared".`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `The URL for uploading and downloading the image file.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `The URL for uploading and downloading the image file.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The date the image was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_kms_data_key_v1",
			Category:         "Key Management Service (KMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"data",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) The globally unique identifier for the key. Changing this gets the new data encryption key.`,
				},
				resource.Attribute{
					Name:        "datakey_length",
					Description: `(Required) Number of bits in the length of a DEK (data encryption keys). The maximum number is 512. Changing this gets the new data encryption key.`,
				},
				resource.Attribute{
					Name:        "encryption_context",
					Description: `(Optional) The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"} ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the date of the found data key. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "plain_text",
					Description: `The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
				resource.Attribute{
					Name:        "cipher_text",
					Description: `The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plain_text",
					Description: `The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
				resource.Attribute{
					Name:        "cipher_text",
					Description: `The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_kms_key_v1",
			Category:         "Key Management Service (KMS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Security-KMS.svg",
			Keywords: []string{
				"key",
				"management",
				"service",
				"kms",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The globally unique identifier for the key. Changing this gets the new key.`,
				},
				resource.Attribute{
					Name:        "key_alias",
					Description: `(Optional) The alias in which to create the key. It is required when we create a new key. Changing this gets the new key.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `(Optional) The description of the key as viewed in FlexibleEngine console. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `(Optional) The state of a key. "2" indicates that the key is enabled. "3" indicates that the key is disabled. "4" indicates that the key is scheduled for deletion. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `(Optional) Identification of a Master Key. The value "1" indicates a Default Master Key, and the value "0" indicates a key. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) ID of a user domain for the key. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "origin",
					Description: `(Optional) Origin of a key. such as: kms. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Region where a key resides. Changing this gets a new key. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "rotation_enabled",
					Description: `Indicates whether the key rotation is enabled or not.`,
				},
				resource.Attribute{
					Name:        "rotation_interval",
					Description: `The key rotation interval. It's valid when rotation is enabled.`,
				},
				resource.Attribute{
					Name:        "rotation_number",
					Description: `The total number of key rotations. It's valid when rotation is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "rotation_enabled",
					Description: `Indicates whether the key rotation is enabled or not.`,
				},
				resource.Attribute{
					Name:        "rotation_interval",
					Description: `The key rotation interval. It's valid when rotation is enabled.`,
				},
				resource.Attribute{
					Name:        "rotation_number",
					Description: `The total number of key rotations. It's valid when rotation is enabled.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_certificate_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"certificate",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific Certificate to retrieve.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-readable name for the Certificate. Does not have to be unique.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-readable description for the Certificate.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) The domain of the Certificate. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The public encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The private encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "certificate",
					Description: `The public encrypted key of the Certificate, PEM format.`,
				},
				resource.Attribute{
					Name:        "update_time",
					Description: `Indicates the update time.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Indicates the creation time.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_lb_loadbalancer_v2",
			Category:         "Elastic Load Balance (ELB)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-ELB.svg",
			Keywords: []string{
				"elastic",
				"load",
				"balance",
				"elb",
				"lb",
				"loadbalancer",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the load balancer.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional, String) Specifies the data source ID of the load balancer in UUID format.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the supplementary information about the load balancer.`,
				},
				resource.Attribute{
					Name:        "vip_subnet_id",
					Description: `(Optional, String) Specifies the ID of the subnet where the load balancer works.`,
				},
				resource.Attribute{
					Name:        "vip_address",
					Description: `(Optional, String) Specifies the private IP address of the load balancer. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The ID of the port bound to the private IP address of the load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated with the load balancer.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vip_port_id",
					Description: `The ID of the port bound to the private IP address of the load balancer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The operating status of the load balancer.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The tags associated with the load balancer.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_modelarts_dataset_versions",
			Category:         "AI Development Platform (ModelArts)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"ai",
				"development",
				"platform",
				"modelarts",
				"dataset",
				"versions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain dataset versions. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "dataset_id",
					Description: `(Required, String) Specifies the ID of dataset.`,
				},
				resource.Attribute{
					Name:        "split_ratio",
					Description: `(Optional, String) Specifies the range of splitting ratio which randomly divides a labeled sample into a training set and a validation set. Separate the minimum and maximum split ratios with commas, for example: "0.0,1.0".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the dataset version. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `Indicates a list of all dataset versions found. Structure is documented below. The ` + "`" + `versions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dataset version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the dataset version.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dataset version.`,
				},
				resource.Attribute{
					Name:        "split_ratio",
					Description: `The ratio of splitting which randomly divides a labeled sample into a training set and a validation set.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dataset version status. Valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `The total number of samples.`,
				},
				resource.Attribute{
					Name:        "storage_path",
					Description: `The path to save the manifest file of the version.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Whether this version is current version.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time, in UTC format.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The last update time, in UTC format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `Indicates a list of all dataset versions found. Structure is documented below. The ` + "`" + `versions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dataset version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the dataset version.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dataset version.`,
				},
				resource.Attribute{
					Name:        "split_ratio",
					Description: `The ratio of splitting which randomly divides a labeled sample into a training set and a validation set.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dataset version status. Valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `The total number of samples.`,
				},
				resource.Attribute{
					Name:        "storage_path",
					Description: `The path to save the manifest file of the version.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Whether this version is current version.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The creation time, in UTC format.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The last update time, in UTC format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_modelarts_datasets",
			Category:         "AI Development Platform (ModelArts)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"ai",
				"development",
				"platform",
				"modelarts",
				"datasets",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain datasets. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of datasets.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Int) Specifies the type of datasets. The options are: +`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "datasets",
					Description: `Indicates a list of all datasets found. Structure is documented below. The ` + "`" + `datasets` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dataset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the dataset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the dataset.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dataset.`,
				},
				resource.Attribute{
					Name:        "output_path",
					Description: `The OBS path for storing output files such as labeled files.`,
				},
				resource.Attribute{
					Name:        "data_source",
					Description: `The data sources which is used to imported the source data (such as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "schemas",
					Description: `The schema information of source data when ` + "`" + `type` + "`" + ` is ` + "`" + `400` + "`" + `(Table Type). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The labels information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `The dataset format. Valid values include: ` + "`" + `Default` + "`" + `, ` + "`" + `CarbonData` + "`" + `: Carbon format(Supported only for table type dataset.).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The dataset creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dataset status. Valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `The type of data source. Valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The OBS path when ` + "`" + `data_type` + "`" + ` is ` + "`" + `0` + "`" + `(OBS) or the HDFS path when ` + "`" + `data_type` + "`" + ` is ` + "`" + `4` + "`" + `(MRS). All the file in this directory and subdirectories will be which be imported to the dataset.`,
				},
				resource.Attribute{
					Name:        "with_column_header",
					Description: `Whether the data contains table header when the type of dataset is ` + "`" + `400` + "`" + `(Table type). The ` + "`" + `schemas` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The field type. Valid values include: ` + "`" + `String` + "`" + `, ` + "`" + `Short` + "`" + `, ` + "`" + `Int` + "`" + `, ` + "`" + `Long` + "`" + `, ` + "`" + `Double` + "`" + `, ` + "`" + `Float` + "`" + `, ` + "`" + `Byte` + "`" + `, ` + "`" + `Date` + "`" + `, ` + "`" + `Timestamp` + "`" + `, ` + "`" + `Bool` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The field name. The ` + "`" + `labels` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of label.`,
				},
				resource.Attribute{
					Name:        "property_color",
					Description: `The color of label.`,
				},
				resource.Attribute{
					Name:        "property_shape",
					Description: `The shape of label. Valid values include: ` + "`" + `bndbox` + "`" + `, ` + "`" + `polygon` + "`" + `, ` + "`" + `circle` + "`" + `, ` + "`" + `line` + "`" + `, ` + "`" + `dashed` + "`" + `, ` + "`" + `point` + "`" + `, ` + "`" + `polyline` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "property_shortcut",
					Description: `The shortcut of label.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Indicates a data source ID.`,
				},
				resource.Attribute{
					Name:        "datasets",
					Description: `Indicates a list of all datasets found. Structure is documented below. The ` + "`" + `datasets` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the dataset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the dataset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the dataset.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the dataset.`,
				},
				resource.Attribute{
					Name:        "output_path",
					Description: `The OBS path for storing output files such as labeled files.`,
				},
				resource.Attribute{
					Name:        "data_source",
					Description: `The data sources which is used to imported the source data (such as pictures/files/audio, etc.) in this directory and subdirectories to the dataset. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "schemas",
					Description: `The schema information of source data when ` + "`" + `type` + "`" + ` is ` + "`" + `400` + "`" + `(Table Type). Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The labels information. Structure is documented below.`,
				},
				resource.Attribute{
					Name:        "data_format",
					Description: `The dataset format. Valid values include: ` + "`" + `Default` + "`" + `, ` + "`" + `CarbonData` + "`" + `: Carbon format(Supported only for table type dataset.).`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The dataset creation time.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Dataset status. Valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "data_type",
					Description: `The type of data source. Valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `The OBS path when ` + "`" + `data_type` + "`" + ` is ` + "`" + `0` + "`" + `(OBS) or the HDFS path when ` + "`" + `data_type` + "`" + ` is ` + "`" + `4` + "`" + `(MRS). All the file in this directory and subdirectories will be which be imported to the dataset.`,
				},
				resource.Attribute{
					Name:        "with_column_header",
					Description: `Whether the data contains table header when the type of dataset is ` + "`" + `400` + "`" + `(Table type). The ` + "`" + `schemas` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The field type. Valid values include: ` + "`" + `String` + "`" + `, ` + "`" + `Short` + "`" + `, ` + "`" + `Int` + "`" + `, ` + "`" + `Long` + "`" + `, ` + "`" + `Double` + "`" + `, ` + "`" + `Float` + "`" + `, ` + "`" + `Byte` + "`" + `, ` + "`" + `Date` + "`" + `, ` + "`" + `Timestamp` + "`" + `, ` + "`" + `Bool` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The field name. The ` + "`" + `labels` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of label.`,
				},
				resource.Attribute{
					Name:        "property_color",
					Description: `The color of label.`,
				},
				resource.Attribute{
					Name:        "property_shape",
					Description: `The shape of label. Valid values include: ` + "`" + `bndbox` + "`" + `, ` + "`" + `polygon` + "`" + `, ` + "`" + `circle` + "`" + `, ` + "`" + `line` + "`" + `, ` + "`" + `dashed` + "`" + `, ` + "`" + `point` + "`" + `, ` + "`" + `polyline` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "property_shortcut",
					Description: `The shortcut of label.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_nat_gateway_v2",
			Category:         "NAT Gateway (NAT)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "nat-gateway.svg",
			Keywords: []string{
				"nat",
				"gateway",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional, String) Specifies the ID of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the NAT gateway name. The name can contain only digits, letters, underscores (_), and hyphens(-).`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String) Specifies the ID of the VPC this NAT gateway belongs to.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, String) Specifies the ID of the VPC Subnet of the downstream interface (the next hop of the DVR) of the NAT gateway.`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional, String) The NAT gateway type. The value can be: + ` + "`" + `1` + "`" + `: small type, which supports up to 10,000 SNAT connections. + ` + "`" + `2` + "`" + `: medium type, which supports up to 50,000 SNAT connections. + ` + "`" + `3` + "`" + `: large type, which supports up to 200,000 SNAT connections. + ` + "`" + `4` + "`" + `: extra-large type, which supports up to 1,000,000 SNAT connections.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional, String) Specifies the description of the NAT gateway. The value contains 0 to 255 characters, and angle brackets (<) and (>) are not allowed.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies the status of the NAT gateway.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_network_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network.svg",
			Keywords: []string{
				"deprecated",
				"networking",
				"network",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve networks ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `(Optional) The CIDR of a subnet within the network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found network. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_port",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the port. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional, String) Specifies the ID of the port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional, String) Specifies the ID of the network the port belongs to.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional, String) Specifies the port IP address filter.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional, String) Specifies the MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies the status of the port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional, String) The list of port security group IDs to filter. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `The device owner of the port.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `The ID of the device the port belongs to.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of security group IDs applied on the port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `The device owner of the port.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `The ID of the device the port belongs to.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The collection of security group IDs applied on the port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_networking_secgroup_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "security-group.svg",
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"networking",
				"secgroup",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve security groups ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "secgroup_id",
					Description: `(Optional) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found security group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_flavors_v1",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"deprecated",
				"rds",
				"flavors",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region in which to obtain the V1 rds client.`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `(Required) The datastore name of the rds.`,
				},
				resource.Attribute{
					Name:        "datastore_version",
					Description: `(Required) The datastore version of the rds.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `(Optional) The spec code of a rds flavor. ## Available value for attributes <!-- markdownlint-disable MD033 --> datastore_name | datastore_version | speccode ---- | --- | --- PostgreSQL | 9.5.5 <br> 9.6.3 <br> 9.6.5| rds.pg.s1.xlarge rds.pg.m1.2xlarge rds.pg.c2.xlarge rds.pg.s1.medium rds.pg.c2.medium rds.pg.s1.large rds.pg.c2.large rds.pg.m1.large rds.pg.s1.2xlarge rds.pg.m1.xlarge MySQL| 5.6.33 <br>5.6.30 <br>5.6.34 <br>5.6.35 <br>5.7.17 | rds.mysql.s1.medium rds.mysql.s1.large rds.mysql.s1.xlarge rds.mysql.s1.2xlarge rds.mysql.m1.2xlarge rds.mysql.c2.medium rds.mysql.c2.large rds.mysql.c2.xlarge rds.mysql.m1.large rds.mysql.m1.xlarge SQLServer| 2014 SP2 SE | rds.mssql.s1.xlarge rds.mssql.m1.2xlarge rds.mssql.c2.xlarge rds.mssql.s1.2xlarge rds.mssql.m1.xlarge ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found rds flavor. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The name of the rds flavor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The name of the rds flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rds_flavors_v3",
			Category:         "Relational Database Service (RDS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"relational",
				"database",
				"service",
				"rds",
				"flavors",
				"v3",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_type",
					Description: `(Required, String) Specifies the DB engine. Value: MySQL, PostgreSQL, SQLServer.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required, String) Specifies the database version. MySQL databases support MySQL 5.6 and 5.7. PostgreSQL databases support PostgreSQL 9.5 and 9.6. Microsoft SQL Server databases support 2014_SE, 2016_SE, and 2016_EE.`,
				},
				resource.Attribute{
					Name:        "instance_mode",
					Description: `(Optional, String) The mode of instance. Value:`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `(Optional, Int) Specifies the number of vCPUs in the RDS flavor.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional, Int) Specifies the memory size(GB) in the RDS flavor.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `(Optional, String) Specifies the performance specification, the valid values are as follows: +`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `(Optional, String) Specifies the availability zone which the RDS flavor belongs to. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The CPU size.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory size in GB.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The performance specification.`,
				},
				resource.Attribute{
					Name:        "instance_mode",
					Description: `The mode of instance.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The availability zones which the RDS flavor belongs to.`,
				},
				resource.Attribute{
					Name:        "db_versions",
					Description: `The Available versions of the database.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The CPU size.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory size in GB.`,
				},
				resource.Attribute{
					Name:        "group_type",
					Description: `The performance specification.`,
				},
				resource.Attribute{
					Name:        "instance_mode",
					Description: `The mode of instance.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `The availability zones which the RDS flavor belongs to.`,
				},
				resource.Attribute{
					Name:        "db_versions",
					Description: `The Available versions of the database.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rts_software_config",
			Category:         "Resource Template Service (RTS)",
			ShortDescription: ``,
			Description: `

The RTS Software Config data source provides details about a specific RTS Software Config.

`,
			Keywords: []string{
				"resource",
				"template",
				"service",
				"rts",
				"software",
				"config",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the software configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the software configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The namespace that groups this software configuration by when it is delivered to a server.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `A list of software configuration inputs.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of software configuration outputs.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The software configuration code.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `The software configuration options.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `The namespace that groups this software configuration by when it is delivered to a server.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `A list of software configuration inputs.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of software configuration outputs.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The software configuration code.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `The software configuration options.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rts_stack_resource_v1",
			Category:         "Resource Template Service (RTS)",
			ShortDescription: ``,
			Description: `

The FlexibleEngine RTS Stack Resource data source allows access to stack resource metadata.

`,
			Keywords: []string{
				"resource",
				"template",
				"service",
				"rts",
				"stack",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "stack_name",
					Description: `(Required) The unique stack name.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The name of a resource in the stack.`,
				},
				resource.Attribute{
					Name:        "physical_resource_id",
					Description: `(Optional) The physical resource ID.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The resource type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "logical_resource_id",
					Description: `The logical resource ID.`,
				},
				resource.Attribute{
					Name:        "resource_status",
					Description: `The status of the resource.`,
				},
				resource.Attribute{
					Name:        "resource_status_reason",
					Description: `The resource operation reason.`,
				},
				resource.Attribute{
					Name:        "required_by",
					Description: `Specifies the resource dependency.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_resource_id",
					Description: `The logical resource ID.`,
				},
				resource.Attribute{
					Name:        "resource_status",
					Description: `The status of the resource.`,
				},
				resource.Attribute{
					Name:        "resource_status_reason",
					Description: `The resource operation reason.`,
				},
				resource.Attribute{
					Name:        "required_by",
					Description: `Specifies the resource dependency.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_rts_stack_v1",
			Category:         "Resource Template Service (RTS)",
			ShortDescription: ``,
			Description: `

The FlexibleEngine RTS Stack data source allows access to stack outputs and other useful data including the template body.

`,
			Keywords: []string{
				"resource",
				"template",
				"service",
				"rts",
				"stack",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the stack. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier of the stack.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `Whether the rollback of the stack is disabled when stack creation fails.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of stack outputs.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `A map of parameters that specify input parameters for the stack.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `Structure containing the template body.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `Specifies the timeout duration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier of the stack.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `Whether the rollback of the stack is disabled when stack creation fails.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of stack outputs.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `A map of parameters that specify input parameters for the stack.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `Structure containing the template body.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `Specifies the timeout duration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_s3_bucket_object",
			Category:         "Object Storage Service (OBS)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "bucket.svg",
			Keywords: []string{
				"object",
				"storage",
				"service",
				"obs",
				"s3",
				"bucket",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "range",
					Description: `(Optional) Obtains the specified range bytes of an object. The value is a range starting from 0 to maximum object length minus one. If the range is invalid, all object data is returned.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured, the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. S3 stores the value of this header in the object metadata.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured, the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. S3 stores the value of this header in the object metadata.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sdrs_domain_v1",
			Category:         "Storage Disaster Recovery Service (SDRS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disaster",
				"recovery",
				"service",
				"sdrs",
				"domain",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the name of an active-active domain. Currently only support SDRS_HypeDomain01. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the active-active domain. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the description of an active-active domain.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the description of an active-active domain.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sfs_file_system_v2",
			Category:         "Scalable File Service (SFS)",
			ShortDescription: ``,
			Description: `

Provides information about an Shared File System (SFS).

`,
			Icon: "Storage-SFS.svg",
			Keywords: []string{
				"scalable",
				"file",
				"service",
				"sfs",
				"system",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the shared file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the shared file system. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size (GB) of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type for the shared file system, such as high-performance storage (composed of SSDs) or large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `The level of visibility for the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The protocol for sharing file systems.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key and value pairs as a dictionary of strings.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The path for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "export_locations",
					Description: `The list of mount locations.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `The level of the access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `The type of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the access rule.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `The UUID of the mount location of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_instance_id",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `Identifies which mount locations are most efficient and are used preferentially when multiple mount locations exist.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size (GB) of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type for the shared file system, such as high-performance storage (composed of SSDs) or large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `The level of visibility for the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The protocol for sharing file systems.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key and value pairs as a dictionary of strings.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The path for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "export_locations",
					Description: `The list of mount locations.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `The level of the access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `The type of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The status of the access rule.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `The UUID of the mount location of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_instance_id",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `Identifies which mount locations are most efficient and are used preferentially when multiple mount locations exist.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sfs_turbos",
			Category:         "Scalable File Service (SFS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"scalable",
				"file",
				"service",
				"sfs",
				"turbos",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the SFS turbo file systems. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional, Int) Specifies the capacity of the SFS turbo file system, in GB. The value ranges from ` + "`" + `500` + "`" + ` to ` + "`" + `32,768` + "`" + `, and must be large than ` + "`" + `10,240` + "`" + ` for an enhanced file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `(Optional, String) Specifies the protocol of the SFS turbo file system. The valid value is`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `(Optional, String) Specifies the type of the SFS turbo file system. The valid values are`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "turbos",
					Description: `The list of the SFS turbo file systems. The [object](#turbo) structure is documented below. <a name="turbo"></a> The ` + "`" + `turbos` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The capacity of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The protocol of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The type of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "enhanced",
					Description: `Whether the SFS turbo file system is enhanced.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone where the SFS turbo file system is located.`,
				},
				resource.Attribute{
					Name:        "available_capacity",
					Description: `The available capacity of the SFS turbo file system, in GB.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The mount point of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "crypt_key_id",
					Description: `The ID of a KMS key to encrypt the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC to which the SFS turbo belongs.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet to which the SFS turbo belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group to which the SFS turbo belongs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "turbos",
					Description: `The list of the SFS turbo file systems. The [object](#turbo) structure is documented below. <a name="turbo"></a> The ` + "`" + `turbos` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource ID of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The capacity of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The protocol of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The type of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "enhanced",
					Description: `Whether the SFS turbo file system is enhanced.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone where the SFS turbo file system is located.`,
				},
				resource.Attribute{
					Name:        "available_capacity",
					Description: `The available capacity of the SFS turbo file system, in GB.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The mount point of the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "crypt_key_id",
					Description: `The ID of a KMS key to encrypt the SFS turbo file system.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC to which the SFS turbo belongs.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet to which the SFS turbo belongs.`,
				},
				resource.Attribute{
					Name:        "security_group_id",
					Description: `The ID of the security group to which the SFS turbo belongs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_smn_topics",
			Category:         "Simple Message Notification (SMN)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"simple",
				"message",
				"notification",
				"smn",
				"topics",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) Specifies the region in which to obtain the SMN topics. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the topic.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `(Optional, String) Specifies the topic URN.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional, String) Specifies the topic display name. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `An array of SMN topics found. Structure is documented below. The ` + "`" + `topics` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The topic ID. The value is the topic URN.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `The topic URN.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The topic display name.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `Message pushing policy. +`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "topics",
					Description: `An array of SMN topics found. Structure is documented below. The ` + "`" + `topics` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the topic.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The topic ID. The value is the topic URN.`,
				},
				resource.Attribute{
					Name:        "topic_urn",
					Description: `The topic URN.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The topic display name.`,
				},
				resource.Attribute{
					Name:        "push_policy",
					Description: `Message pushing policy. +`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_sms_source_servers",
			Category:         "Server Migration Service (SMS)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"server",
				"migration",
				"service",
				"sms",
				"source",
				"servers",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional, String) Specifies the ID of the source server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) Specifies the name of the source server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional, String) Specifies the status of the source server.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional, String) Specifies the IP address of the source server. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `An array of SMS source servers found. Structure is documented below. The ` + "`" + `servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the source server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the source server.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP address of the source server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the source server.`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `Whether the source server is properly connected to SMS.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The OS type of the source server. The value can be`,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `The OS version of the source server, for example, UBUNTU_20_4_64BIT.`,
				},
				resource.Attribute{
					Name:        "registered_time",
					Description: `The UTC time when the source server is registered.`,
				},
				resource.Attribute{
					Name:        "agent_version",
					Description: `The version of Agent installed on the source server.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The vcpus count of the source server.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory size in MB.`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `The disk information of the source server. Structure is documented below. The ` + "`" + `disks` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The disk name, for example, /dev/vda.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The disk size in MB.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `The disk type. The value can be`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID.`,
				},
				resource.Attribute{
					Name:        "servers",
					Description: `An array of SMS source servers found. Structure is documented below. The ` + "`" + `servers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the source server.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the source server.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP address of the source server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the source server.`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `Whether the source server is properly connected to SMS.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `The OS type of the source server. The value can be`,
				},
				resource.Attribute{
					Name:        "os_version",
					Description: `The OS version of the source server, for example, UBUNTU_20_4_64BIT.`,
				},
				resource.Attribute{
					Name:        "registered_time",
					Description: `The UTC time when the source server is registered.`,
				},
				resource.Attribute{
					Name:        "agent_version",
					Description: `The version of Agent installed on the source server.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `The vcpus count of the source server.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The memory size in MB.`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `The disk information of the source server. Structure is documented below. The ` + "`" + `disks` + "`" + ` blocks support:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The disk name, for example, /dev/vda.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The disk size in MB.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `The disk type. The value can be`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vbs_backup_policy_v2",
			Category:         "Volume Backup Service (VBS)",
			ShortDescription: ``,
			Description: `

The VBS Backup Policy data source provides details about a specific VBS backup policy.

`,
			Keywords: []string{
				"volume",
				"backup",
				"service",
				"vbs",
				"policy",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Specifies the start time of the backup job.The value is in the HH:mm format.`,
				},
				resource.Attribute{
					Name:        "retain_first_backup",
					Description: `Specifies whether to retain the first backup in the current month.`,
				},
				resource.Attribute{
					Name:        "rentention_num",
					Description: `Specifies number of retained backups.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Specifies the backup interval. The value is in the range of 1 to 14 days.`,
				},
				resource.Attribute{
					Name:        "policy_resource_count",
					Description: `Specifies the number of volumes associated with the backup policy.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Specifies the start time of the backup job.The value is in the HH:mm format.`,
				},
				resource.Attribute{
					Name:        "retain_first_backup",
					Description: `Specifies whether to retain the first backup in the current month.`,
				},
				resource.Attribute{
					Name:        "rentention_num",
					Description: `Specifies number of retained backups.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Specifies the backup interval. The value is in the range of 1 to 14 days.`,
				},
				resource.Attribute{
					Name:        "policy_resource_count",
					Description: `Specifies the number of volumes associated with the backup policy.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vbs_backup_v2",
			Category:         "Volume Backup Service (VBS)",
			ShortDescription: ``,
			Description: `

The VBS Backup data source provides details about a specific VBS Backup.

`,
			Keywords: []string{
				"volume",
				"backup",
				"service",
				"vbs",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) The source volume ID of the backup.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) ID of the snapshot associated with the backup.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the VBS backup. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ where the backup resides.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `The container of the backup.`,
				},
				resource.Attribute{
					Name:        "service_metadata",
					Description: `The metadata of the vbs backup.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ where the backup resides.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `The container of the backup.`,
				},
				resource.Attribute{
					Name:        "service_metadata",
					Description: `The metadata of the vbs backup.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_eip",
			Category:         "Elastic IP (EIP)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"elastic",
				"ip",
				"eip",
				"vpc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "public_ip",
					Description: `(Optional, String) The public ip address of the EIP.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional, String) The port id of the EIP. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EIP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip of the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth_id",
					Description: `The bandwidth id of the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `The bandwidth size of the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth_share_type",
					Description: `The bandwidth share type of the EIP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the EIP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the EIP.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private ip of the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth_id",
					Description: `The bandwidth id of the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `The bandwidth size of the EIP.`,
				},
				resource.Attribute{
					Name:        "bandwidth_share_type",
					Description: `The bandwidth share type of the EIP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_peering_connection_v2",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description: `

The VPC Peering Connection data source provides details about a specific VPC peering connection.

`,
			Icon: "peer link.svg",
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"peering",
				"connection",
				"v2",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_route_ids_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description: `

!> **WARNING:** It has been deprecated, use ` + "`" + `flexibleengine_vpc_route_table` + "`" + ` to get the route details.

` + "`" + `flexibleengine_vpc_route_ids_v2` + "`" + ` provides a list of route ids for a vpc_id.

This resource can be useful for getting back a list of route ids for a vpc.

`,
			Keywords: []string{
				"deprecated",
				"vpc",
				"route",
				"ids",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the route ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the route ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_route_table",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"route",
				"table",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to query the vpc route table. If omitted, the provider-level region will be used.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_route_v2",
			Category:         "Deprecated",
			ShortDescription: ``,
			Description: `

!> **WARNING:** It has been deprecated, use ` + "`" + `flexibleengine_vpc_route_table` + "`" + ` to get the route details.

` + "`" + `flexibleengine_vpc_route_v2` + "`" + ` provides details about a specific VPC route.

`,
			Icon: "Network-VPC.svg",
			Keywords: []string{
				"deprecated",
				"vpc",
				"route",
				"v2",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nexthop",
					Description: `The next hop of the route. If the route type is peering, it will provide VPC peering connection ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nexthop",
					Description: `The next hop of the route. If the route type is peering, it will provide VPC peering connection ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_subnet_ids_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"subnet",
				"ids",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the subnet ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the subnet ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_subnet_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "subnet.svg",
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"subnet",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) - Specifies the ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "dns_list",
					Description: `The IP address list of DNS servers on the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enable",
					Description: `DHCP function for the subnet.`,
				},
				resource.Attribute{
					Name:        "ipv4_subnet_id",
					Description: `The ID of the IPv4 subnet (Native OpenStack API).`,
				},
				resource.Attribute{
					Name:        "ipv6_enable",
					Description: `Whether the IPv6 is enabled.`,
				},
				resource.Attribute{
					Name:        "ipv6_subnet_id",
					Description: `The ID of the IPv6 subnet (Native OpenStack API).`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr",
					Description: `The IPv6 subnet CIDR block.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `The IPv6 subnet gateway.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpc_v1",
			Category:         "Virtual Private Cloud (VPC)",
			ShortDescription: ``,
			Description:      ``,
			Icon:             "Network-VPC.svg",
			Keywords: []string{
				"virtual",
				"private",
				"cloud",
				"vpc",
				"v1",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 VPC client. A VPC client is needed to retrieve VPCs. If omitted, the region argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The cidr block of the desired VPC. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The list of route information with destination and nexthop fields.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The list of route information with destination and nexthop fields.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpcep_endpoints",
			Category:         "VPC Endpoint (VPCEP)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"endpoint",
				"vpcep",
				"endpoints",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional, String) Specifies the name of the VPC endpoint service. The value is not case-sensitive and supports fuzzy match.`,
				},
				resource.Attribute{
					Name:        "endpoint_id",
					Description: `(Optional, String) Specifies the unique ID of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Optional, String) Specifies the unique ID of the vpc holding the VPC endpoint service. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `Indicates the public VPC endpoints information. Structure is documented below. The ` + "`" + `endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the public VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The connection status of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC holding the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the subnet holding the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "packet_id",
					Description: `The marker id of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "enable_dns",
					Description: `Flag indicating dns has been enabled for the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `Flag indicating access control have been enabled on this VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `List of IP or CIDR block which can access the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "private_domain_name",
					Description: `DNS name pointing to the VPC endpoint ip.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The key/value pairs to associate with the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project holding the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date of the VPC endpoint.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `Indicates the public VPC endpoints information. Structure is documented below. The ` + "`" + `endpoints` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the public VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The connection status of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `The ID of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The ID of the VPC holding the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `The ID of the subnet holding the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "packet_id",
					Description: `The marker id of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "enable_dns",
					Description: `Flag indicating dns has been enabled for the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "enable_whitelist",
					Description: `Flag indicating access control have been enabled on this VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "whitelist",
					Description: `List of IP or CIDR block which can access the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "private_domain_name",
					Description: `DNS name pointing to the VPC endpoint ip.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `The key/value pairs to associate with the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The ID of the project holding the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date of the VPC endpoint.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date of the VPC endpoint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_vpcep_public_services",
			Category:         "VPC Endpoint (VPCEP)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"vpc",
				"endpoint",
				"vpcep",
				"public",
				"services",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional, String) Specifies the name of the public VPC endpoint service. The value is not case-sensitive and supports fuzzy match.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Optional, String) Specifies the unique ID of the public VPC endpoint service. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the public VPC endpoint services.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `Indicates the public VPC endpoint services information. Structure is documented below. The ` + "`" + `services` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the public VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the public VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The owner of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "is_charge",
					Description: `Indicates whether the associated VPC endpoint carries a charge.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Specifies a data source ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region in which to obtain the public VPC endpoint services.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `Indicates the public VPC endpoint services information. Structure is documented below. The ` + "`" + `services` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the public VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The name of the public VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `The type of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The owner of the VPC endpoint service.`,
				},
				resource.Attribute{
					Name:        "is_charge",
					Description: `Indicates whether the associated VPC endpoint carries a charge.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "flexibleengine_waf_dedicated_instances",
			Category:         "Web Application Firewall (WAF)",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"web",
				"application",
				"firewall",
				"waf",
				"dedicated",
				"instances",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional, String) The region in which to query the WAF dedicated instance. If omitted, the provider-level region will be used.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional, String) The id of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional, String) The name of WAF dedicated instance. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format. The following attributes are exported: The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone names for the WAF dedicated instances.`,
				},
				resource.Attribute{
					Name:        "specification_code",
					Description: `The specification code of instance. Different specifications have different throughput. Values are: + ` + "`" + `waf.instance.professional` + "`" + ` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only). +` + "`" + `waf.instance.enterprise` + "`" + ` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `The ECS cpu architecture of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "ecs_flavor",
					Description: `The flavor of the ECS used by the WAF instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC id of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet of WAF dedicated instance VPC.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The security group of the instance. This is an array of security group ids.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The service of the instance.`,
				},
				resource.Attribute{
					Name:        "service_ip",
					Description: `The service ip of the instance.`,
				},
				resource.Attribute{
					Name:        "run_status",
					Description: `The running status of the instance. Values are: + ` + "`" + `0` + "`" + ` - Instance is creating. + ` + "`" + `1` + "`" + ` - Instance has created. + ` + "`" + `2` + "`" + ` - Instance is deleting. + ` + "`" + `3` + "`" + ` - Instance has deleted. + ` + "`" + `4` + "`" + ` - Instance create failed.`,
				},
				resource.Attribute{
					Name:        "access_status",
					Description: `The access status of the instance. ` + "`" + `0` + "`" + `: inaccessible, ` + "`" + `1` + "`" + `: accessible.`,
				},
				resource.Attribute{
					Name:        "upgradable",
					Description: `The instance is to support upgrades. ` + "`" + `0` + "`" + `: Cannot be upgraded, ` + "`" + `1` + "`" + `: Can be upgraded.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The instance group ID used by the WAF dedicated instance in ELB mode.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The data source ID in UUID format. The following attributes are exported: The ` + "`" + `instances` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "available_zone",
					Description: `The available zone names for the WAF dedicated instances.`,
				},
				resource.Attribute{
					Name:        "specification_code",
					Description: `The specification code of instance. Different specifications have different throughput. Values are: + ` + "`" + `waf.instance.professional` + "`" + ` - The professional edition, throughput: 100 Mbit/s; QPS: 2,000 (Reference only). +` + "`" + `waf.instance.enterprise` + "`" + ` - The enterprise edition, throughput: 500 Mbit/s; QPS: 10,000 (Reference only).`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `The ECS cpu architecture of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "ecs_flavor",
					Description: `The flavor of the ECS used by the WAF instance.`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `The VPC id of WAF dedicated instance.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the VPC Subnet of WAF dedicated instance VPC.`,
				},
				resource.Attribute{
					Name:        "security_group",
					Description: `The security group of the instance. This is an array of security group ids.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The service of the instance.`,
				},
				resource.Attribute{
					Name:        "service_ip",
					Description: `The service ip of the instance.`,
				},
				resource.Attribute{
					Name:        "run_status",
					Description: `The running status of the instance. Values are: + ` + "`" + `0` + "`" + ` - Instance is creating. + ` + "`" + `1` + "`" + ` - Instance has created. + ` + "`" + `2` + "`" + ` - Instance is deleting. + ` + "`" + `3` + "`" + ` - Instance has deleted. + ` + "`" + `4` + "`" + ` - Instance create failed.`,
				},
				resource.Attribute{
					Name:        "access_status",
					Description: `The access status of the instance. ` + "`" + `0` + "`" + `: inaccessible, ` + "`" + `1` + "`" + `: accessible.`,
				},
				resource.Attribute{
					Name:        "upgradable",
					Description: `The instance is to support upgrades. ` + "`" + `0` + "`" + `: Cannot be upgraded, ` + "`" + `1` + "`" + `: Can be upgraded.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The instance group ID used by the WAF dedicated instance in ELB mode.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"flexibleengine_apig_environments":                  0,
		"flexibleengine_availability_zones":                 1,
		"flexibleengine_blockstorage_availability_zones_v3": 2,
		"flexibleengine_blockstorage_volume_v2":             3,
		"flexibleengine_cbr_vaults":                         4,
		"flexibleengine_cce_addon_template":                 5,
		"flexibleengine_cce_cluster_v3":                     6,
		"flexibleengine_cce_clusters":                       7,
		"flexibleengine_cce_node_ids_v3":                    8,
		"flexibleengine_cce_node_v3":                        9,
		"flexibleengine_compute_availability_zones_v2":      10,
		"flexibleengine_compute_bms_flavors_v2":             11,
		"flexibleengine_compute_bms_keypairs_v2":            12,
		"flexibleengine_compute_bms_nic_v2":                 13,
		"flexibleengine_compute_bms_server_v2":              14,
		"flexibleengine_compute_flavors_v2":                 15,
		"flexibleengine_compute_instance_v2":                16,
		"flexibleengine_compute_instances":                  17,
		"flexibleengine_csbs_backup_policy_v1":              18,
		"flexibleengine_csbs_backup_v1":                     19,
		"flexibleengine_cts_tracker_v1":                     20,
		"flexibleengine_dcs_az_v1":                          21,
		"flexibleengine_dcs_maintainwindow_v1":              22,
		"flexibleengine_dcs_product_v1":                     23,
		"flexibleengine_dds_flavor_v3":                      24,
		"flexibleengine_dds_flavors_v3":                     25,
		"flexibleengine_dms_kafka_instances":                26,
		"flexibleengine_dms_product":                        27,
		"flexibleengine_dms_rocketmq_broker":                28,
		"flexibleengine_dms_rocketmq_instances":             29,
		"flexibleengine_dns_zone_v2":                        30,
		"flexibleengine_dws_flavors":                        31,
		"flexibleengine_elb_certificate":                    32,
		"flexibleengine_elb_flavors":                        33,
		"flexibleengine_enterprise_project":                 34,
		"flexibleengine_fgs_dependencies":                   35,
		"flexibleengine_gaussdb_cassandra_instances":        36,
		"flexibleengine_gaussdb_nosql_flavors":              37,
		"flexibleengine_identity_custom_role_v3":            38,
		"flexibleengine_identity_group":                     39,
		"flexibleengine_identity_project_v3":                40,
		"flexibleengine_identity_role_v3":                   41,
		"flexibleengine_identity_users":                     42,
		"flexibleengine_images_image_v2":                    43,
		"flexibleengine_kms_data_key_v1":                    44,
		"flexibleengine_kms_key_v1":                         45,
		"flexibleengine_lb_certificate_v2":                  46,
		"flexibleengine_lb_loadbalancer_v2":                 47,
		"flexibleengine_modelarts_dataset_versions":         48,
		"flexibleengine_modelarts_datasets":                 49,
		"flexibleengine_nat_gateway_v2":                     50,
		"flexibleengine_networking_network_v2":              51,
		"flexibleengine_networking_port":                    52,
		"flexibleengine_networking_secgroup_v2":             53,
		"flexibleengine_rds_flavors_v1":                     54,
		"flexibleengine_rds_flavors_v3":                     55,
		"flexibleengine_rts_software_config":                56,
		"flexibleengine_rts_stack_resource_v1":              57,
		"flexibleengine_rts_stack_v1":                       58,
		"flexibleengine_s3_bucket_object":                   59,
		"flexibleengine_sdrs_domain_v1":                     60,
		"flexibleengine_sfs_file_system_v2":                 61,
		"flexibleengine_sfs_turbos":                         62,
		"flexibleengine_smn_topics":                         63,
		"flexibleengine_sms_source_servers":                 64,
		"flexibleengine_vbs_backup_policy_v2":               65,
		"flexibleengine_vbs_backup_v2":                      66,
		"flexibleengine_vpc_eip":                            67,
		"flexibleengine_vpc_peering_connection_v2":          68,
		"flexibleengine_vpc_route_ids_v2":                   69,
		"flexibleengine_vpc_route_table":                    70,
		"flexibleengine_vpc_route_v2":                       71,
		"flexibleengine_vpc_subnet_ids_v1":                  72,
		"flexibleengine_vpc_subnet_v1":                      73,
		"flexibleengine_vpc_v1":                             74,
		"flexibleengine_vpcep_endpoints":                    75,
		"flexibleengine_vpcep_public_services":              76,
		"flexibleengine_waf_dedicated_instances":            77,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
