package ovh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_capabilities_containerregistry",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all capabilities plans ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `List of container registry capability for a single region`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The region name`,
				},
				resource.Attribute{
					Name:        "plans",
					Description: `Available plans in the region`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Plan code from the catalog`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Plan creation date`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features of the plan`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `Vulnerability scanning`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Plan ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Plan name`,
				},
				resource.Attribute{
					Name:        "registry_limits",
					Description: `Container registry limits`,
				},
				resource.Attribute{
					Name:        "image_storage",
					Description: `Docker image storage limits in bytes`,
				},
				resource.Attribute{
					Name:        "parallel_request",
					Description: `Parallel requests on Docker image API (/v2 Docker registry API)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Plan last update date`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `List of container registry capability for a single region`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `The region name`,
				},
				resource.Attribute{
					Name:        "plans",
					Description: `Available plans in the region`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Plan code from the catalog`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Plan creation date`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features of the plan`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `Vulnerability scanning`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Plan ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Plan name`,
				},
				resource.Attribute{
					Name:        "registry_limits",
					Description: `Container registry limits`,
				},
				resource.Attribute{
					Name:        "image_storage",
					Description: `Docker image storage limits in bytes`,
				},
				resource.Attribute{
					Name:        "parallel_request",
					Description: `Parallel requests on Docker image API (/v2 Docker registry API)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Plan last update date`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_capabilities_containerregistry_filter",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region name`,
				},
				resource.Attribute{
					Name:        "plan_name",
					Description: `The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Plan code from the catalog`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Plan creation date`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features of the plan`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `Vulnerability scanning`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Plan ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Plan name`,
				},
				resource.Attribute{
					Name:        "registry_limits",
					Description: `Container registry limits`,
				},
				resource.Attribute{
					Name:        "image_storage",
					Description: `Docker image storage limits in bytes`,
				},
				resource.Attribute{
					Name:        "parallel_request",
					Description: `Parallel requests on Docker image API (/v2 Docker registry API)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Plan last update date`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `Plan code from the catalog`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Plan creation date`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `Features of the plan`,
				},
				resource.Attribute{
					Name:        "vulnerability",
					Description: `Vulnerability scanning`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Plan ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Plan name`,
				},
				resource.Attribute{
					Name:        "registry_limits",
					Description: `Container registry limits`,
				},
				resource.Attribute{
					Name:        "image_storage",
					Description: `Docker image storage limits in bytes`,
				},
				resource.Attribute{
					Name:        "parallel_request",
					Description: `Parallel requests on Docker image API (/v2 Docker registry API)`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Plan last update date`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_containerregistries",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all registries ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The list of container registries associated with the project.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Registry creation date`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Registry ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Registry name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of your registry`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the registry`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Current size of the registry (bytes)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Registry status`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Registry last update date`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Access url of the registry`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of your registry`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of container registries associated with the project.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Registry creation date`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Registry ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Registry name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of your registry`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the registry`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Current size of the registry (bytes)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Registry status`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Registry last update date`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Access url of the registry`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of your registry`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_containerregistry",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `Registry ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Registry creation date`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Registry ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Registry name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of your registry`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the registry`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Current size of the registry (bytes)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Registry status`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Registry last update date`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Access url of the registry`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of your registry`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Registry creation date`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Registry ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Registry name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project ID of your registry`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the registry`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Current size of the registry (bytes)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Registry status`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Registry last update date`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Access url of the registry`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of your registry`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_containerregistry_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "registry_id",
					Description: `Registry ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The list of users of the container registry associated with the project.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `User ID`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `User name`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `User email`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of users of the container registry associated with the project.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `User ID`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `User name`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `User email`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The database engine you want to get information. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Cluster ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `Time on which backups start every day.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Small description of the database service.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `List of all endpoints objects of the service.`,
				},
				resource.Attribute{
					Name:        "component",
					Description: `Type of component the URI relates to.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the cluster.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection port for the endpoint.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `Scheme used to generate the URI.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `Defines whether the endpoint uses SSL.`,
				},
				resource.Attribute{
					Name:        "ssl_mode",
					Description: `SSL mode used to connect to the service if the SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `URI of the endpoint.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `A valid OVHcloud public cloud database flavor name in which the nodes will be started.`,
				},
				resource.Attribute{
					Name:        "kafka_rest_api",
					Description: `Defines whether the REST API is enabled on a kafka cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_time",
					Description: `Time on which maintenances can start every day.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Type of network of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of nodes object.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Private network id in which the node should be deployed. It's the regional openstackId of the private network`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Public cloud region in which the node should be deployed.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Private subnet ID in which the node is.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Plan of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the engine in which the service should be deployed`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size (in GB) of the database service.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The disk type of the database service.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Advanced configuration key / value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `Time on which backups start every day.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Small description of the database service.`,
				},
				resource.Attribute{
					Name:        "endpoints",
					Description: `List of all endpoints objects of the service.`,
				},
				resource.Attribute{
					Name:        "component",
					Description: `Type of component the URI relates to.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain of the cluster.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path of the endpoint.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Connection port for the endpoint.`,
				},
				resource.Attribute{
					Name:        "scheme",
					Description: `Scheme used to generate the URI.`,
				},
				resource.Attribute{
					Name:        "ssl",
					Description: `Defines whether the endpoint uses SSL.`,
				},
				resource.Attribute{
					Name:        "ssl_mode",
					Description: `SSL mode used to connect to the service if the SSL is enabled.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `URI of the endpoint.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `A valid OVHcloud public cloud database flavor name in which the nodes will be started.`,
				},
				resource.Attribute{
					Name:        "kafka_rest_api",
					Description: `Defines whether the REST API is enabled on a kafka cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_time",
					Description: `Time on which maintenances can start every day.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `Type of network of the cluster.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of nodes object.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `Private network id in which the node should be deployed. It's the regional openstackId of the private network`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Public cloud region in which the node should be deployed.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Private subnet ID in which the node is.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Plan of the cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the engine in which the service should be deployed`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size (in GB) of the database service.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The disk type of the database service.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Advanced configuration key / value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_capabilities",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to ` + "`" + `service_name` + "`" + ` value. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "engines",
					Description: `Database engines available.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `Default version used for the engine.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the engine.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Engine name.`,
				},
				resource.Attribute{
					Name:        "ssl_modes",
					Description: `SSL modes for this engine.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `Versions available for this engine.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Flavors available.`,
				},
				resource.Attribute{
					Name:        "core",
					Description: `Flavor core number.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Flavor ram size in GB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the flavor.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Flavor disk size in GB.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `Options available.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the option.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the option.`,
				},
				resource.Attribute{
					Name:        "plans",
					Description: `Plans available.`,
				},
				resource.Attribute{
					Name:        "backup_retention",
					Description: `Automatic backup retention duration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the plan.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the plan.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "engines",
					Description: `Database engines available.`,
				},
				resource.Attribute{
					Name:        "default_version",
					Description: `Default version used for the engine.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the engine.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Engine name.`,
				},
				resource.Attribute{
					Name:        "ssl_modes",
					Description: `SSL modes for this engine.`,
				},
				resource.Attribute{
					Name:        "versions",
					Description: `Versions available for this engine.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Flavors available.`,
				},
				resource.Attribute{
					Name:        "core",
					Description: `Flavor core number.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Flavor ram size in GB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the flavor.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `Flavor disk size in GB.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `Options available.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the option.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the option.`,
				},
				resource.Attribute{
					Name:        "plans",
					Description: `Plans available.`,
				},
				resource.Attribute{
					Name:        "backup_retention",
					Description: `Automatic backup retention duration.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the plan.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the plan.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_certificates",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want database information. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference The following attributes are exported: ` + "`" + `id` + "`" + ` is set to the md5 sum of the CA. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `CA certificate used for the service.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `CA certificate used for the service.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_database",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want database information. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the database. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the database.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Defines if the database has been created by default.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the database.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the database.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the database.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the database.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Defines if the database has been created by default.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the database.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the database.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the database.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_databases",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want to list databases. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all database ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "database_ids",
					Description: `The list of databases ids of the database cluster associated with the project.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "database_ids",
					Description: `The list of databases ids of the database cluster associated with the project.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_integration",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required, Forces new resource) The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept ` + "`" + `mongodb` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Integration ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `ID of the destination service.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Parameters for the integration.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_service_id",
					Description: `ID of the source service.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the integration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `ID of the destination service.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Parameters for the integration.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_service_id",
					Description: `ID of the source service.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the integration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_integrations",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept ` + "`" + `mongodb` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all integration ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "integration_ids",
					Description: `The list of integrations ids of the database cluster associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "integration_ids",
					Description: `The list of integrations ids of the database cluster associated with the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_ip_restrictions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want to list IP restrictions. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all IP restrictions. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of IP restriction of the database associated with the project.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of IP restriction of the database associated with the project.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_acl",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) ACL ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Permission to give to this username on this topic.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `Topic affected by this ACL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username affected by this ACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Permission to give to this username on this topic.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `Topic affected by this ACL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username affected by this ACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_acls",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all ACL ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "acl_ids",
					Description: `The list of ACLs ids of the kafka cluster associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "acl_ids",
					Description: `The list of ACLs ids of the kafka cluster associated with the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_topic",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Topic ID ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_insync_replicas",
					Description: `Minimum insync replica accepted for this topic.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the topic.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `Number of partitions for this topic.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `Number of replication for this topic.`,
				},
				resource.Attribute{
					Name:        "retention_bytes",
					Description: `Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited`,
				},
				resource.Attribute{
					Name:        "retention_hours",
					Description: `Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "min_insync_replicas",
					Description: `Minimum insync replica accepted for this topic.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the topic.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `Number of partitions for this topic.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `Number of replication for this topic.`,
				},
				resource.Attribute{
					Name:        "retention_bytes",
					Description: `Number of bytes for the retention of the data for this topic. Inferior to 0 mean Unlimited`,
				},
				resource.Attribute{
					Name:        "retention_hours",
					Description: `Number of hours for the retention of the data for this topic. Inferior to 0 mean Unlimited`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_topics",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all topics ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic_ids",
					Description: `The list of topics ids of the kafka cluster associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic_ids",
					Description: `The list of topics ids of the kafka cluster associated with the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_user_access",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) User ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the Cert. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `User cert.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Sensitive) User key for the cert.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `User cert.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Sensitive) User key for the cert.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_m3db_namespace",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the namespace. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "resolution",
					Description: `Resolution for an aggregated namespace.`,
				},
				resource.Attribute{
					Name:        "retention_block_data_expiration_duration",
					Description: `Controls how long we wait before expiring stale data.`,
				},
				resource.Attribute{
					Name:        "retention_block_size_duration",
					Description: `Controls how long to keep a block in memory before flushing to a fileset on disk.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_future_duration",
					Description: `Controls how far into the future writes to the namespace will be accepted.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_past_duration",
					Description: `Controls how far into the past writes to the namespace will be accepted.`,
				},
				resource.Attribute{
					Name:        "retention_period_duration",
					Description: `Controls the duration of time that M3DB will retain data for the namespace.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of namespace.`,
				},
				resource.Attribute{
					Name:        "writes_to_commit_log_enabled",
					Description: `Defines whether M3DB will include writes to this namespace in the commit log.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the namespace.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "resolution",
					Description: `Resolution for an aggregated namespace.`,
				},
				resource.Attribute{
					Name:        "retention_block_data_expiration_duration",
					Description: `Controls how long we wait before expiring stale data.`,
				},
				resource.Attribute{
					Name:        "retention_block_size_duration",
					Description: `Controls how long to keep a block in memory before flushing to a fileset on disk.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_future_duration",
					Description: `Controls how far into the future writes to the namespace will be accepted.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_past_duration",
					Description: `Controls how far into the past writes to the namespace will be accepted.`,
				},
				resource.Attribute{
					Name:        "retention_period_duration",
					Description: `Controls the duration of time that M3DB will retain data for the namespace.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of namespace.`,
				},
				resource.Attribute{
					Name:        "writes_to_commit_log_enabled",
					Description: `Defines whether M3DB will include writes to this namespace in the commit log.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_m3db_namespaces",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all namespaces ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "namespace_ids",
					Description: `The list of namespaces ids of the M3DB cluster associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "namespace_ids",
					Description: `The list of namespaces ids of the M3DB cluster associated with the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_m3db_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_mongodb_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user with the authentication database in the format name@authDB, for example: johndoe@admin ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Roles the user belongs to`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Roles the user belongs to`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_opensearch_pattern",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Pattern ID. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_index_count",
					Description: `Maximum number of index for this pattern.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Pattern format.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the pattern.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "max_index_count",
					Description: `Maximum number of index for this pattern.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Pattern format.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the pattern.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_opensearch_patterns",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all patterns ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern_ids",
					Description: `The list of patterns ids of the opensearch cluster associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern_ids",
					Description: `The list of patterns ids of the opensearch cluster associated with the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_opensearch_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `Acls of the user.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Pattern of the ACL.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Permission of the ACL.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `Acls of the user.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Pattern of the ACL.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `Permission of the ACL.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_postgresql_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Roles the user belongs to.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Roles the user belongs to.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_redis_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `Categories of the user.`,
				},
				resource.Attribute{
					Name:        "channels",
					Description: `Channels of the user.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `Commands of the user.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `Keys of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "categories",
					Description: `Categories of the user.`,
				},
				resource.Attribute{
					Name:        "channels",
					Description: `Channels of the user.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `Commands of the user.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `Keys of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want user information. To get a full list of available engine visit : [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the user. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date of the creation of the user.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the user.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The engine of the database cluster you want to list users. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) Cluster ID ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all user ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `The list of users ids of the database cluster associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "user_ids",
					Description: `The list of users ids of the database cluster associated with the project.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_databases",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) The database engine you want to list. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the md5 sum of the list of all databases ids. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_ids",
					Description: `The list of managed databases ids of the project.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_ids",
					Description: `The list of managed databases ids of the project.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_failover_ip_attach",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The failover ip address to query ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "block",
					Description: `The IP block`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `The Ip continent`,
				},
				resource.Attribute{
					Name:        "geoloc",
					Description: `The Ip location`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Ip id`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The Ip Address`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Current operation progress in percent`,
				},
				resource.Attribute{
					Name:        "routedTo",
					Description: `Instance where ip is routed to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Ip status, can be ` + "`" + `ok` + "`" + ` or ` + "`" + `operationPending` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subType",
					Description: `IP sub type, can be ` + "`" + `cloud` + "`" + ` or ` + "`" + `ovh` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "block",
					Description: `The IP block`,
				},
				resource.Attribute{
					Name:        "continentCode",
					Description: `The Ip continent`,
				},
				resource.Attribute{
					Name:        "geoloc",
					Description: `The Ip location`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Ip id`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The Ip Address`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Current operation progress in percent`,
				},
				resource.Attribute{
					Name:        "routedTo",
					Description: `Instance where ip is routed to`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Ip status, can be ` + "`" + `ok` + "`" + ` or ` + "`" + `operationPending` + "`" + ``,
				},
				resource.Attribute{
					Name:        "subType",
					Description: `IP sub type, can be ` + "`" + `cloud` + "`" + ` or ` + "`" + `ovh` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The id of the managed kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The OVHcloud public cloud region ID of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Kubernetes version of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_id",
					Description: `OpenStack private network (or vrack) ID to use.`,
				},
				resource.Attribute{
					Name:        "control_plane_is_up_to_date",
					Description: `True if control-plane is up-to-date.`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `True if all nodes and control-plane are up-to-date.`,
				},
				resource.Attribute{
					Name:        "next_upgrade_versions",
					Description: `Kubernetes versions available for upgrade.`,
				},
				resource.Attribute{
					Name:        "nodes_url",
					Description: `Cluster nodes URL.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status. Should be normally set to 'READY'.`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Management URL of your cluster.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `Selected mode for kube-proxy.`,
				},
				resource.Attribute{
					Name:        "apiserver",
					Description: `Kubernetes API server customization`,
				},
				resource.Attribute{
					Name:        "kube_proxy",
					Description: `Kubernetes kube-proxy customization`,
				},
				resource.Attribute{
					Name:        "customization_apiserver",
					Description: `Kubernetes API server customization`,
				},
				resource.Attribute{
					Name:        "admissionplugins",
					Description: `Kubernetes API server admission plugins customization`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.`,
				},
				resource.Attribute{
					Name:        "customization_kube_proxy",
					Description: `Kubernetes kube-proxy customization`,
				},
				resource.Attribute{
					Name:        "iptables",
					Description: `Kubernetes cluster kube-proxy customization of iptables specific config.`,
				},
				resource.Attribute{
					Name:        "sync_period",
					Description: `Minimum period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.`,
				},
				resource.Attribute{
					Name:        "min_sync_period",
					Description: `Period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.`,
				},
				resource.Attribute{
					Name:        "ipvs",
					Description: `Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
				resource.Attribute{
					Name:        "sync_period",
					Description: `Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.`,
				},
				resource.Attribute{
					Name:        "min_sync_period",
					Description: `Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `IPVS scheduler.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `Timeout value used for idle IPVS TCP sessions in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
				resource.Attribute{
					Name:        "tcp_fin_timeout",
					Description: `Timeout value used for IPVS TCP sessions after receiving a FIN in RFC3339 duration.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `timeout value used for IPVS UDP packets in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The OVHcloud public cloud region ID of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Kubernetes version of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "private_network_id",
					Description: `OpenStack private network (or vrack) ID to use.`,
				},
				resource.Attribute{
					Name:        "control_plane_is_up_to_date",
					Description: `True if control-plane is up-to-date.`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `True if all nodes and control-plane are up-to-date.`,
				},
				resource.Attribute{
					Name:        "next_upgrade_versions",
					Description: `Kubernetes versions available for upgrade.`,
				},
				resource.Attribute{
					Name:        "nodes_url",
					Description: `Cluster nodes URL.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status. Should be normally set to 'READY'.`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Management URL of your cluster.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `Selected mode for kube-proxy.`,
				},
				resource.Attribute{
					Name:        "apiserver",
					Description: `Kubernetes API server customization`,
				},
				resource.Attribute{
					Name:        "kube_proxy",
					Description: `Kubernetes kube-proxy customization`,
				},
				resource.Attribute{
					Name:        "customization_apiserver",
					Description: `Kubernetes API server customization`,
				},
				resource.Attribute{
					Name:        "admissionplugins",
					Description: `Kubernetes API server admission plugins customization`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.`,
				},
				resource.Attribute{
					Name:        "customization_kube_proxy",
					Description: `Kubernetes kube-proxy customization`,
				},
				resource.Attribute{
					Name:        "iptables",
					Description: `Kubernetes cluster kube-proxy customization of iptables specific config.`,
				},
				resource.Attribute{
					Name:        "sync_period",
					Description: `Minimum period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.`,
				},
				resource.Attribute{
					Name:        "min_sync_period",
					Description: `Period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.`,
				},
				resource.Attribute{
					Name:        "ipvs",
					Description: `Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
				resource.Attribute{
					Name:        "sync_period",
					Description: `Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.`,
				},
				resource.Attribute{
					Name:        "min_sync_period",
					Description: `Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `IPVS scheduler.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `Timeout value used for idle IPVS TCP sessions in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
				resource.Attribute{
					Name:        "tcp_fin_timeout",
					Description: `Timeout value used for IPVS TCP sessions after receiving a FIN in RFC3339 duration.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `timeout value used for IPVS UDP packets in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_iprestrictions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The id of the managed kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of CIDRs that restricts the access to the API server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of CIDRs that restricts the access to the API server.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_nodepool",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The id of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the node pool. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the nodepool. Changing this value recreates the resource. Warning: "_" char is not allowed!`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". Changing this value recreates the resource. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/`,
				},
				resource.Attribute{
					Name:        "desired_nodes",
					Description: `number of nodes to start.`,
				},
				resource.Attribute{
					Name:        "max_nodes",
					Description: `maximum number of nodes allowed in the pool. Setting ` + "`" + `desired_nodes` + "`" + ` over this value will raise an error.`,
				},
				resource.Attribute{
					Name:        "min_nodes",
					Description: `minimum number of nodes allowed in the pool. Setting ` + "`" + `desired_nodes` + "`" + ` under this value will raise an error.`,
				},
				resource.Attribute{
					Name:        "monthly_billed",
					Description: `(Optional) should the nodes be billed on a monthly basis. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "anti_affinity",
					Description: `(Optional) should the pool use the anti-affinity feature. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "autoscale",
					Description: `(Optional) Enable auto-scaling for the pool. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "available_nodes",
					Description: `Number of nodes which are actually ready in the pool`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date`,
				},
				resource.Attribute{
					Name:        "current_nodes",
					Description: `Number of nodes present in the pool`,
				},
				resource.Attribute{
					Name:        "desired_nodes",
					Description: `Number of nodes you desire in the pool`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Flavor name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id`,
				},
				resource.Attribute{
					Name:        "size_status",
					Description: `Status describing the state between number of nodes wanted and available ones`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status`,
				},
				resource.Attribute{
					Name:        "up_to_date_nodes",
					Description: `Number of nodes with the latest version installed in the pool`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the nodepool. Changing this value recreates the resource. Warning: "_" char is not allowed!`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". Changing this value recreates the resource. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/`,
				},
				resource.Attribute{
					Name:        "desired_nodes",
					Description: `number of nodes to start.`,
				},
				resource.Attribute{
					Name:        "max_nodes",
					Description: `maximum number of nodes allowed in the pool. Setting ` + "`" + `desired_nodes` + "`" + ` over this value will raise an error.`,
				},
				resource.Attribute{
					Name:        "min_nodes",
					Description: `minimum number of nodes allowed in the pool. Setting ` + "`" + `desired_nodes` + "`" + ` under this value will raise an error.`,
				},
				resource.Attribute{
					Name:        "monthly_billed",
					Description: `(Optional) should the nodes be billed on a monthly basis. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "anti_affinity",
					Description: `(Optional) should the pool use the anti-affinity feature. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "autoscale",
					Description: `(Optional) Enable auto-scaling for the pool. Default to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "available_nodes",
					Description: `Number of nodes which are actually ready in the pool`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date`,
				},
				resource.Attribute{
					Name:        "current_nodes",
					Description: `Number of nodes present in the pool`,
				},
				resource.Attribute{
					Name:        "desired_nodes",
					Description: `Number of nodes you desire in the pool`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Flavor name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Project id`,
				},
				resource.Attribute{
					Name:        "size_status",
					Description: `Status describing the state between number of nodes wanted and available ones`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status`,
				},
				resource.Attribute{
					Name:        "up_to_date_nodes",
					Description: `Number of nodes with the latest version installed in the pool`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_nodepool_nodes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The ID of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The ID of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the node pool from which we want the nodes. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of all nodes composing the kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date.`,
				},
				resource.Attribute{
					Name:        "deployed_at",
					Description: `(Optional) Date of the effective deployment.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Flavor name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Openstack ID of the underlying VM of the node.`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `Is the node in the target version of the cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the node.`,
				},
				resource.Attribute{
					Name:        "node_pool_id",
					Description: `Managed kubernetes node pool ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Public cloud project ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version in which the node is.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of all nodes composing the kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date.`,
				},
				resource.Attribute{
					Name:        "deployed_at",
					Description: `(Optional) Date of the effective deployment.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Flavor name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node.`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Openstack ID of the underlying VM of the node.`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `Is the node in the target version of the cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the node.`,
				},
				resource.Attribute{
					Name:        "node_pool_id",
					Description: `Managed kubernetes node pool ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Public cloud project ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version in which the node is.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_nodes",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The ID of the managed kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of all nodes composing the kubernetes cluster`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date`,
				},
				resource.Attribute{
					Name:        "deployed_at",
					Description: `(Optional) Date of the effective deployment`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Flavor name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Openstack ID of the underlying VM of the node`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `Is the node in the target version of the cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the node`,
				},
				resource.Attribute{
					Name:        "node_pool_id",
					Description: `Managed kubernetes node pool ID`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Public cloud project ID`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version in which the node is`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `List of all nodes composing the kubernetes cluster`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation date`,
				},
				resource.Attribute{
					Name:        "deployed_at",
					Description: `(Optional) Date of the effective deployment`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `Flavor name`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the node`,
				},
				resource.Attribute{
					Name:        "instance_id",
					Description: `Openstack ID of the underlying VM of the node`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `Is the node in the target version of the cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the node`,
				},
				resource.Attribute{
					Name:        "node_pool_id",
					Description: `Managed kubernetes node pool ID`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `Public cloud project ID`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update date`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version in which the node is`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_oidc",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The id of the managed kubernetes cluster. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The OIDC client ID.`,
				},
				resource.Attribute{
					Name:        "issuer_url",
					Description: `The OIDC issuer url.`,
				},
				resource.Attribute{
					Name:        "oidcUsernameClaim",
					Description: `JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.`,
				},
				resource.Attribute{
					Name:        "oidcUsernamePrefix",
					Description: `Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and oidcUsernameClaim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.`,
				},
				resource.Attribute{
					Name:        "oidcGroupsClaim",
					Description: `Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.`,
				},
				resource.Attribute{
					Name:        "oidcGroupsPrefix",
					Description: `Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.`,
				},
				resource.Attribute{
					Name:        "oidcRequiredClaim",
					Description: `Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."`,
				},
				resource.Attribute{
					Name:        "oidcSigningAlgs",
					Description: `Array of signing algorithms accepted. Default is \"RS256\".`,
				},
				resource.Attribute{
					Name:        "oidcCaContent",
					Description: `Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `The OIDC client ID.`,
				},
				resource.Attribute{
					Name:        "issuer_url",
					Description: `The OIDC issuer url.`,
				},
				resource.Attribute{
					Name:        "oidcUsernameClaim",
					Description: `JWT claim to use as the user name. By default sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.`,
				},
				resource.Attribute{
					Name:        "oidcUsernamePrefix",
					Description: `Prefix prepended to username claims to prevent clashes with existing names (such as system: users). For example, the value oidc: will create usernames like oidc:jane.doe. If this field isn't set and oidcUsernameClaim is a value other than email the prefix defaults to ( Issuer URL )# where ( Issuer URL ) is the value of oidcIssuerUrl. The value - can be used to disable all prefixing.`,
				},
				resource.Attribute{
					Name:        "oidcGroupsClaim",
					Description: `Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.`,
				},
				resource.Attribute{
					Name:        "oidcGroupsPrefix",
					Description: `Prefix prepended to group claims to prevent clashes with existing names (such as system: groups). For example, the value oidc: will create group names like oidc:engineering and oidc:infra.`,
				},
				resource.Attribute{
					Name:        "oidcRequiredClaim",
					Description: `Array of key=value pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."`,
				},
				resource.Attribute{
					Name:        "oidcSigningAlgs",
					Description: `Array of signing algorithms accepted. Default is \"RS256\".`,
				},
				resource.Attribute{
					Name:        "oidcCaContent",
					Description: `Content of the certificate for the CA, in base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_region",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the region associated with the public cloud project. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project concatenated with the name of the region. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "continent_code",
					Description: `the code of the geographic continent the region is running. E.g.: EU for Europe, US for America...`,
				},
				resource.Attribute{
					Name:        "datacenter_location",
					Description: `The location code of the datacenter. E.g.: "GRA", meaning Gravelines, for region "GRA1"`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `The list of public cloud services running within the region`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `the name of the public cloud service`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_regions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "has_services_up",
					Description: `(Optional) List of services which has to be UP in regions. Example: "image", "instance", "network", "storage", "volume", "workflow", ... If left blank, returns all regions associated with the service_name. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the project. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project, filtered by services UP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of regions associated with the project, filtered by services UP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_user_s3_credential",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_user_s3_credentials",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_user_s3_policy",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dbaas_logs_cluster",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The service name. It's the ID of your Logs Data Platform instance. ## Attributes Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dbaas_logs_input_engine",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The service name. It's the ID of your Logs Data Platform instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the logs input engine.`,
				},
				resource.Attribute{
					Name:        "is_deprecated",
					Description: `Indicates if engine will soon not be supported.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Software version ## Attributes Reference ` + "`" + `id` + "`" + ` is set to input engine ID`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dbaas_logs_output_graylog_stream",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The service name. It's the ID of your Logs Data Platform instance.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `Stream description ## Attributes Reference ` + "`" + `id` + "`" + ` is set to output graylog stream ID. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cold_storage_compression",
					Description: `Cold storage compression method`,
				},
				resource.Attribute{
					Name:        "cold_storage_content",
					Description: `ColdStorage content`,
				},
				resource.Attribute{
					Name:        "cold_storage_enabled",
					Description: `Is Cold storage enabled?`,
				},
				resource.Attribute{
					Name:        "cold_storage_notify_enabled",
					Description: `Notify on new Cold storage archive`,
				},
				resource.Attribute{
					Name:        "cold_storage_retention",
					Description: `Cold storage retention in year`,
				},
				resource.Attribute{
					Name:        "cold_storage_target",
					Description: `ColdStorage destination`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Stream creation`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Stream description`,
				},
				resource.Attribute{
					Name:        "indexing_enabled",
					Description: `Enable ES indexing`,
				},
				resource.Attribute{
					Name:        "indexing_max_size",
					Description: `Maximum indexing size (in GB)`,
				},
				resource.Attribute{
					Name:        "indexing_notify_enabled",
					Description: `If set, notify when size is near 80, 90 or 100 % of the maximum configured setting`,
				},
				resource.Attribute{
					Name:        "is_editable",
					Description: `Indicates if you are allowed to edit entry`,
				},
				resource.Attribute{
					Name:        "is_shareable",
					Description: `Indicates if you are allowed to share entry`,
				},
				resource.Attribute{
					Name:        "nb_alert_condition",
					Description: `Number of alert condition`,
				},
				resource.Attribute{
					Name:        "nb_archive",
					Description: `Number of coldstored archives`,
				},
				resource.Attribute{
					Name:        "parent_stream_id",
					Description: `Parent stream ID`,
				},
				resource.Attribute{
					Name:        "pause_indexing_on_max_size",
					Description: `If set, pause indexing when maximum size is reach`,
				},
				resource.Attribute{
					Name:        "retention_id",
					Description: `Retention ID`,
				},
				resource.Attribute{
					Name:        "stream_id",
					Description: `Stream ID`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Stream last update`,
				},
				resource.Attribute{
					Name:        "web_socket_enabled",
					Description: `Enable Websocket`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cold_storage_compression",
					Description: `Cold storage compression method`,
				},
				resource.Attribute{
					Name:        "cold_storage_content",
					Description: `ColdStorage content`,
				},
				resource.Attribute{
					Name:        "cold_storage_enabled",
					Description: `Is Cold storage enabled?`,
				},
				resource.Attribute{
					Name:        "cold_storage_notify_enabled",
					Description: `Notify on new Cold storage archive`,
				},
				resource.Attribute{
					Name:        "cold_storage_retention",
					Description: `Cold storage retention in year`,
				},
				resource.Attribute{
					Name:        "cold_storage_target",
					Description: `ColdStorage destination`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Stream creation`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Stream description`,
				},
				resource.Attribute{
					Name:        "indexing_enabled",
					Description: `Enable ES indexing`,
				},
				resource.Attribute{
					Name:        "indexing_max_size",
					Description: `Maximum indexing size (in GB)`,
				},
				resource.Attribute{
					Name:        "indexing_notify_enabled",
					Description: `If set, notify when size is near 80, 90 or 100 % of the maximum configured setting`,
				},
				resource.Attribute{
					Name:        "is_editable",
					Description: `Indicates if you are allowed to edit entry`,
				},
				resource.Attribute{
					Name:        "is_shareable",
					Description: `Indicates if you are allowed to share entry`,
				},
				resource.Attribute{
					Name:        "nb_alert_condition",
					Description: `Number of alert condition`,
				},
				resource.Attribute{
					Name:        "nb_archive",
					Description: `Number of coldstored archives`,
				},
				resource.Attribute{
					Name:        "parent_stream_id",
					Description: `Parent stream ID`,
				},
				resource.Attribute{
					Name:        "pause_indexing_on_max_size",
					Description: `If set, pause indexing when maximum size is reach`,
				},
				resource.Attribute{
					Name:        "retention_id",
					Description: `Retention ID`,
				},
				resource.Attribute{
					Name:        "stream_id",
					Description: `Stream ID`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Stream last update`,
				},
				resource.Attribute{
					Name:        "web_socket_enabled",
					Description: `Enable Websocket`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_ceph",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name of the dedicated CEPH cluster. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "ceph_mons",
					Description: `list of CEPH monitors IPs`,
				},
				resource.Attribute{
					Name:        "ceph_version",
					Description: `CEPH cluster version`,
				},
				resource.Attribute{
					Name:        "crush_tunables",
					Description: `CRUSH algorithm settings. Possible values`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `CEPH cluster label`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `cluster region`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Cluster size in TB`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state of the cluster`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ceph_mons",
					Description: `list of CEPH monitors IPs`,
				},
				resource.Attribute{
					Name:        "ceph_version",
					Description: `CEPH cluster version`,
				},
				resource.Attribute{
					Name:        "crush_tunables",
					Description: `CRUSH algorithm settings. Possible values`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `CEPH cluster label`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `cluster region`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Cluster size in TB`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state of the cluster`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the service`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_installation_templates",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of installation templates IDs available for dedicated servers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of installation templates IDs available for dedicated servers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_nasha",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated HA-NAS. ## Attributes Reference ` + "`" + `id` + "`" + ` is set with the service_name of the dedicated HA-NAS. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The storage service name`,
				},
				resource.Attribute{
					Name:        "can_create_partition",
					Description: `True, if partition creation is allowed on this HA-NAS`,
				},
				resource.Attribute{
					Name:        "custom_name",
					Description: `The name you give to the HA-NAS`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `area of HA-NAS`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `the disk type of the HA-NAS. Possible values are: ` + "`" + `hdd` + "`" + `, ` + "`" + `ssd` + "`" + `, ` + "`" + `nvme` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access IP of HA-NAS`,
				},
				resource.Attribute{
					Name:        "monitored",
					Description: `Send an email to customer if any issue is detected`,
				},
				resource.Attribute{
					Name:        "zpool_capacity",
					Description: `percentage of HA-NAS space used in %`,
				},
				resource.Attribute{
					Name:        "zpool_size",
					Description: `the size of the HA-NAS in GB`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The storage service name`,
				},
				resource.Attribute{
					Name:        "can_create_partition",
					Description: `True, if partition creation is allowed on this HA-NAS`,
				},
				resource.Attribute{
					Name:        "custom_name",
					Description: `The name you give to the HA-NAS`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `area of HA-NAS`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `the disk type of the HA-NAS. Possible values are: ` + "`" + `hdd` + "`" + `, ` + "`" + `ssd` + "`" + `, ` + "`" + `nvme` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Access IP of HA-NAS`,
				},
				resource.Attribute{
					Name:        "monitored",
					Description: `Send an email to customer if any issue is detected`,
				},
				resource.Attribute{
					Name:        "zpool_capacity",
					Description: `percentage of HA-NAS space used in %`,
				},
				resource.Attribute{
					Name:        "zpool_size",
					Description: `the size of the HA-NAS in GB`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server. ## Attributes Reference ` + "`" + `id` + "`" + ` is set with the service_name of the dedicated server. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "boot_id",
					Description: `boot id of the server`,
				},
				resource.Attribute{
					Name:        "commercial_range",
					Description: `dedicater server commercial range`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `dedicated datacenter localisation (bhs1,bhs2,...)`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `dedicated server ip (IPv4)`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `dedicated server ip blocks`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `link speed of the server`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Icmp monitoring state`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `dedicated server name`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Operating system`,
				},
				resource.Attribute{
					Name:        "professional_use",
					Description: `Does this server have professional use option`,
				},
				resource.Attribute{
					Name:        "rack",
					Description: `rack id of the server`,
				},
				resource.Attribute{
					Name:        "rescue_mail",
					Description: `rescue mail of the server`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `dedicated server reverse`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `root device of the server`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `your server id`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `error, hacked, hackedBlocked, ok`,
				},
				resource.Attribute{
					Name:        "support_level",
					Description: `Dedicated server support level (critical, fastpath, gs, pro)`,
				},
				resource.Attribute{
					Name:        "vnis",
					Description: `the list of Virtualnetworkinterface assiociated with this server`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `VirtualNetworkInterface activation state`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `VirtualNetworkInterface mode (public,vrack,vrack_aggregation)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User defined VirtualNetworkInterface name`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `VirtualNetworkInterface unique id`,
				},
				resource.Attribute{
					Name:        "vrack",
					Description: `vRack name`,
				},
				resource.Attribute{
					Name:        "ncis",
					Description: `NetworkInterfaceControllers bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_vnis",
					Description: `List of enabled vrack VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_aggregation_vnis",
					Description: `List of enabled vrack_aggregation VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_public_vnis",
					Description: `List of enabled public VNI uuids`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_id",
					Description: `boot id of the server`,
				},
				resource.Attribute{
					Name:        "commercial_range",
					Description: `dedicater server commercial range`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `dedicated datacenter localisation (bhs1,bhs2,...)`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `dedicated server ip (IPv4)`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `dedicated server ip blocks`,
				},
				resource.Attribute{
					Name:        "link_speed",
					Description: `link speed of the server`,
				},
				resource.Attribute{
					Name:        "monitoring",
					Description: `Icmp monitoring state`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `dedicated server name`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Operating system`,
				},
				resource.Attribute{
					Name:        "professional_use",
					Description: `Does this server have professional use option`,
				},
				resource.Attribute{
					Name:        "rack",
					Description: `rack id of the server`,
				},
				resource.Attribute{
					Name:        "rescue_mail",
					Description: `rescue mail of the server`,
				},
				resource.Attribute{
					Name:        "reverse",
					Description: `dedicated server reverse`,
				},
				resource.Attribute{
					Name:        "root_device",
					Description: `root device of the server`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `your server id`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `error, hacked, hackedBlocked, ok`,
				},
				resource.Attribute{
					Name:        "support_level",
					Description: `Dedicated server support level (critical, fastpath, gs, pro)`,
				},
				resource.Attribute{
					Name:        "vnis",
					Description: `the list of Virtualnetworkinterface assiociated with this server`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `VirtualNetworkInterface activation state`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `VirtualNetworkInterface mode (public,vrack,vrack_aggregation)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `User defined VirtualNetworkInterface name`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Server bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `VirtualNetworkInterface unique id`,
				},
				resource.Attribute{
					Name:        "vrack",
					Description: `vRack name`,
				},
				resource.Attribute{
					Name:        "ncis",
					Description: `NetworkInterfaceControllers bound to this VirtualNetworkInterface`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_vnis",
					Description: `List of enabled vrack VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_vrack_aggregation_vnis",
					Description: `List of enabled vrack_aggregation VNI uuids`,
				},
				resource.Attribute{
					Name:        "enabled_public_vnis",
					Description: `List of enabled public VNI uuids`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server_boots",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your dedicated server.`,
				},
				resource.Attribute{
					Name:        "boot_type",
					Description: `(Optional) Filter the value of bootType property (harddisk, rescue, ipxeCustomerScript, internal, network) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated server netboots.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated server netboots.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_servers",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated servers IDs associated with your OVHcloud Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of dedicated servers IDs associated with your OVHcloud Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the domain zone. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the domain zone name. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database ## Attributes Reference ` + "`" + `id` + "`" + ` is set to database service_name. In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of CPU on your private database`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `Datacenter where this private database is located`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name displayed in customer panel for your private database`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Private database hostname`,
				},
				resource.Attribute{
					Name:        "hostname_ftp",
					Description: `Private database FTP hostname`,
				},
				resource.Attribute{
					Name:        "infrastructure",
					Description: `Infrastructure where service was stored`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `Type of the private database offer`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Private database service port`,
				},
				resource.Attribute{
					Name:        "port_ftp",
					Description: `Private database FTP port`,
				},
				resource.Attribute{
					Name:        "quota_size",
					Description: `Space allowed (in MB) on your private database`,
				},
				resource.Attribute{
					Name:        "quota_used",
					Description: `Sapce used (in MB) on your private database`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Amount of ram (in MB) on your private database`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Private database server name`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Private database state`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Private database available versions`,
				},
				resource.Attribute{
					Name:        "version_label",
					Description: `Private database version label`,
				},
				resource.Attribute{
					Name:        "version_number",
					Description: `Private database version number`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of CPU on your private database`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `Datacenter where this private database is located`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Name displayed in customer panel for your private database`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Private database hostname`,
				},
				resource.Attribute{
					Name:        "hostname_ftp",
					Description: `Private database FTP hostname`,
				},
				resource.Attribute{
					Name:        "infrastructure",
					Description: `Infrastructure where service was stored`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `Type of the private database offer`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Private database service port`,
				},
				resource.Attribute{
					Name:        "port_ftp",
					Description: `Private database FTP port`,
				},
				resource.Attribute{
					Name:        "quota_size",
					Description: `Space allowed (in MB) on your private database`,
				},
				resource.Attribute{
					Name:        "quota_used",
					Description: `Sapce used (in MB) on your private database`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Amount of ram (in MB) on your private database`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Private database server name`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Private database state`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Private database available versions`,
				},
				resource.Attribute{
					Name:        "version_label",
					Description: `Private database version label`,
				},
				resource.Attribute{
					Name:        "version_number",
					Description: `Private database version number`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_database",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database name ## Attributes Reference ` + "`" + `id` + "`" + ` is set to ` + "`" + `service_name` + "`" + `/` + "`" + `database_name` + "`" + `. In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "backup_time",
					Description: `Time of the next backup (every day)`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "quota_used",
					Description: `Space used by the database (in MB)`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Users granted to this database`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User's name granted on this database`,
				},
				resource.Attribute{
					Name:        "grant_type",
					Description: `Grant of this user for this database`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_time",
					Description: `Time of the next backup (every day)`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "quota_used",
					Description: `Space used by the database (in MB)`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `Users granted to this database`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User's name granted on this database`,
				},
				resource.Attribute{
					Name:        "grant_type",
					Description: `Grant of this user for this database`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `User name ## Attributes Reference ` + "`" + `id` + "`" + ` is set to ` + "`" + `service_name` + "`" + `/` + "`" + `user_name` + "`" + `. In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `Users granted to this database`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database's name linked to this user`,
				},
				resource.Attribute{
					Name:        "grant_type",
					Description: `Grant of this user for this database`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "databases",
					Description: `Users granted to this database`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database's name linked to this user`,
				},
				resource.Attribute{
					Name:        "grant_type",
					Description: `Grant of this user for this database`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_user_grant",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `The database name on which grant the user`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The user name ## Attributes Reference ` + "`" + `id` + "`" + ` is set to ` + "`" + `service_name` + "`" + `/` + "`" + `user_name` + "`" + `/` + "`" + `database_name` + "`" + `/` + "`" + `grant` + "`" + `. In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "grant",
					Description: `Grant name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "grant",
					Description: `Grant name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_whitelist",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The whitelisted IP in your instance ## Attributes Reference ` + "`" + `id` + "`" + ` is set to ` + "`" + `service_name` + "`" + `/` + "`" + `ip` + "`" + `. In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `The last update date of this whitelist`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Custom name for your Whitelisted IP`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Authorize this IP to access service port`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `Authorize this IP to access SFTP port`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Whitelist status`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation date of the database`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `The last update date of this whitelist`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Custom name for your Whitelisted IP`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `Authorize this IP to access service port`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `Authorize this IP to access SFTP port`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Whitelist status`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_ip_service",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The service name ## Attributes Reference ` + "`" + `id` + "`" + ` is set to ip service ip block. In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "can_be_terminated",
					Description: `can be terminated`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `country`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Custom description on your ip`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `ip block`,
				},
				resource.Attribute{
					Name:        "organisation_id",
					Description: `IP block organisation Id`,
				},
				resource.Attribute{
					Name:        "routed_to",
					Description: `Routage information`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Service where ip is routed to`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Possible values for ip type ( "cdn", "cloud", "dedicated", "failover", "hosted_ssl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "can_be_terminated",
					Description: `can be terminated`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `country`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Custom description on your ip`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `ip block`,
				},
				resource.Attribute{
					Name:        "organisation_id",
					Description: `IP block organisation Id`,
				},
				resource.Attribute{
					Name:        "routed_to",
					Description: `Routage information`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `Service where ip is routed to`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Possible values for ip type ( "cdn", "cloud", "dedicated", "failover", "hosted_ssl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPV6 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPV4 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Location where your service is. This takes an array of values.`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The offer of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `Your IP load balancing`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of your IP. Can take any of the following value: "blacklisted", "deleted", "free", "ok", "quarantined", "suspended"`,
				},
				resource.Attribute{
					Name:        "vrack_eligibility",
					Description: `Vrack eligibility. Takes a boolean value.`,
				},
				resource.Attribute{
					Name:        "vrack_name",
					Description: `Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product`,
				},
				resource.Attribute{
					Name:        "ssl_configuration",
					Description: `Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Can take any of the following value: "intermediate", "modern"`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `the name displayed in ManagerV6 for your iplb (max 50 chars) ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the service_name of your IP load balancing In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing This attribute is sensitive.`,
				},
				resource.Attribute{
					Name:        "orderable_zone",
					Description: `Available additional zone for your Load Balancer`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The zone three letter code`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `The billing planCode for this zone`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing This attribute is sensitive.`,
				},
				resource.Attribute{
					Name:        "orderable_zone",
					Description: `Available additional zone for your Load Balancer`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The zone three letter code`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `The billing planCode for this zone`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_vrack_network",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "vrack_network_id",
					Description: `(Required) Internal Load Balancer identifier of the vRack private network ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your vrack network`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP block of the private network in the vRack`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN of the private network in the vRack. 0 if the private network is not in a VLAN`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your vrack network`,
				},
				resource.Attribute{
					Name:        "nat_ip",
					Description: `An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `IP block of the private network in the vRack`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `VLAN of the private network in the vRack. 0 if the private network is not in a VLAN`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_vrack_networks",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `Filters networks on the subnet.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `Filters networks on the vlan id. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `The list of vrack network ids.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of vrack network ids.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_identity_user",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user",
					Description: `(Required) User's login. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `User's login suffix.`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `Creation date of this user.`,
				},
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
					Name:        "login",
					Description: `User's login suffix.`,
				},
				resource.Attribute{
					Name:        "creation",
					Description: `Creation date of this user.`,
				},
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
			Type:             "ovh_me_identity_users",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `The list of the user's logins of all the identity users.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "users",
					Description: `The list of the user's logins of all the identity users.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_templates",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of custom installation templates IDs available for dedicated servers.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of custom installation templates IDs available for dedicated servers.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ipxe_script",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the IPXE Script. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The content of the script.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The content of the script.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ipxe_scripts",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of the names of all the IPXE Scripts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of the names of all the IPXE Scripts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_paymentmean_bankaccount",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regexp",
					Description: `(Optional) a regexp used to filter bank accounts on their ` + "`" + `description` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "use_default",
					Description: `(Optional) Retrieve bank account marked as default payment mean.`,
				},
				resource.Attribute{
					Name:        "use_oldest",
					Description: `(Optional) Retrieve oldest bank account. project.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) Filter bank accounts on their ` + "`" + `state` + "`" + ` attribute. Can be "blockedForIncidents", "valid", "pendingValidation" ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the bank account payment mean`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the bank account`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved bank account is marked as the default payment mean`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the bank account`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved bank account is marked as the default payment mean`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_paymentmean_creditcard",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description_regexp",
					Description: `(Optional) a regexp used to filter credit cards on their ` + "`" + `description` + "`" + ` attributes.`,
				},
				resource.Attribute{
					Name:        "use_default",
					Description: `(Optional) Retrieve credit card marked as default payment mean.`,
				},
				resource.Attribute{
					Name:        "use_last_to_expire",
					Description: `(Optional) Retrieve the credit card that will be the last to expire according to its expiration date.`,
				},
				resource.Attribute{
					Name:        "states",
					Description: `(Optional) Filter credit cards on their ` + "`" + `state` + "`" + ` attribute. Can be "expired", "valid", "tooManyFailures" ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the credit card payment mean`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved credit card is marked as the default payment mean`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `the description attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `the state attribute of the credit card`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `a boolean which tells if the retrieved credit card is marked as the default payment mean`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ssh_key",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The name of the SSH key. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The content of the public key. E.g.: "ssh-ed25519 AAAAC3..."`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `True when this public SSH key is used for rescue mode and reinstallations.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The content of the public key. E.g.: "ssh-ed25519 AAAAC3..."`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `True when this public SSH key is used for rescue mode and reinstallations.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ssh_keys",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of the names of all the SSH keys.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The list of the names of all the SSH keys.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_order_cart",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of your cart`,
				},
				resource.Attribute{
					Name:        "assign",
					Description: `Assign a shopping cart to an loggedin client. Values can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `Expiration time (format: 2006-01-02T15:04:05+00:00) ## Attributes Reference ` + "`" + `id` + "`" + ` is set to your cart ID In addition, the following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "cart_id",
					Description: `Cart identifier`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Indicates if the cart has already been validated`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `Items of your cart`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cart_id",
					Description: `Cart identifier`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Indicates if the cart has already been validated`,
				},
				resource.Attribute{
					Name:        "items",
					Description: `Items of your cart`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_order_cart_product",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cart_id",
					Description: `(Required) Cart identifier`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) product ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `products results`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `products results`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_order_cart_product_options",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cart_id",
					Description: `(Required) Cart identifier`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Product`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "result",
					Description: `products results`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `Define if options of this family are exclusive with each other`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Option family`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Define if an option of this family is mandatory`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `products results`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `Define if options of this family are exclusive with each other`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Option family`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Define if an option of this family is mandatory`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_order_cart_product_options_plan",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cart_id",
					Description: `(Required) Cart identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "options_plan_code",
					Description: `(Required) options plan code.`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Product offer identifier`,
				},
				resource.Attribute{
					Name:        "price_capacity",
					Description: `(Required) Capacity of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Product ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "selected_price",
					Description: `Selected Price according to capacity`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `Define if options of this family are exclusive with each other`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Option family`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Define if an option of this family is mandatory`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "selected_price",
					Description: `Selected Price according to capacity`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "exclusive",
					Description: `Define if options of this family are exclusive with each other`,
				},
				resource.Attribute{
					Name:        "family",
					Description: `Option family`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `Define if an option of this family is mandatory`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_order_cart_product_plan",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cart_id",
					Description: `(Required) Cart identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Product offer identifier`,
				},
				resource.Attribute{
					Name:        "price_capacity",
					Description: `(Required) Capacity of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(Required) Product ## Attributes Reference The following attributes are exported.`,
				},
				resource.Attribute{
					Name:        "selected_price",
					Description: `Selected Price according to capacity`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "selected_price",
					Description: `Selected Price according to capacity`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Product offer identifier`,
				},
				resource.Attribute{
					Name:        "product_name",
					Description: `Name of the product`,
				},
				resource.Attribute{
					Name:        "product_type",
					Description: `Product type`,
				},
				resource.Attribute{
					Name:        "prices",
					Description: `Prices of the product offer`,
				},
				resource.Attribute{
					Name:        "capacities",
					Description: `Capacities of the pricing (type of pricing)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the pricing`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Duration for ordering the product`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Interval of renewal`,
				},
				resource.Attribute{
					Name:        "maximum_quantity",
					Description: `Maximum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "maximum_repeat",
					Description: `Maximum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "minimum_quantity",
					Description: `Minimum quantity that can be ordered`,
				},
				resource.Attribute{
					Name:        "minimum_repeat",
					Description: `Minimum repeat for renewal`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `Price of the product (Price with its currency and textual representation)`,
				},
				resource.Attribute{
					Name:        "currency_code",
					Description: `Currency code`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `Textual representation`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The effective price`,
				},
				resource.Attribute{
					Name:        "price_in_ucents",
					Description: `Price of the product in micro-centims`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "pricing_type",
					Description: `Pricing type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vps",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server. ## Attribute Reference ` + "`" + `id` + "`" + ` is set with the service\_name of the vps name (ex: "vps-123456.ovh.net") In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `The OVHcloud cluster the vps is in`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter in which the vps is located`,
				},
				resource.Attribute{
					Name:        "datacenter.longname",
					Description: `The fullname of the datacenter (ex: "Strasbourg SBG1")`,
				},
				resource.Attribute{
					Name:        "datacenter.name",
					Description: `The short name of the datacenter (ex: "sbg1)`,
				},
				resource.Attribute{
					Name:        "displayname",
					Description: `The displayed name in the OVHcloud web admin`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of IPs addresses attached to the vps`,
				},
				resource.Attribute{
					Name:        "keymap",
					Description: `The keymap for the ip kvm, valid values "", "fr", "us"`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of memory in MB of the vps.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `A dict describing the type of vps.`,
				},
				resource.Attribute{
					Name:        "model.name",
					Description: `The model name (ex: model1)`,
				},
				resource.Attribute{
					Name:        "model.offer",
					Description: `The model human description (ex: "VPS 2016 SSD 1")`,
				},
				resource.Attribute{
					Name:        "model.version",
					Description: `The model version (ex: "2017v2")`,
				},
				resource.Attribute{
					Name:        "netbootmode",
					Description: `The source of the boot kernel`,
				},
				resource.Attribute{
					Name:        "offertype",
					Description: `The type of offer (ssd, cloud, classic)`,
				},
				resource.Attribute{
					Name:        "slamonitoring",
					Description: `A boolean to indicate if OVHcloud SLA monitoring is active.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the vps`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of server`,
				},
				resource.Attribute{
					Name:        "vcore",
					Description: `The number of vcore of the vps`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The OVHcloud zone where the vps is`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster",
					Description: `The OVHcloud cluster the vps is in`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `The datacenter in which the vps is located`,
				},
				resource.Attribute{
					Name:        "datacenter.longname",
					Description: `The fullname of the datacenter (ex: "Strasbourg SBG1")`,
				},
				resource.Attribute{
					Name:        "datacenter.name",
					Description: `The short name of the datacenter (ex: "sbg1)`,
				},
				resource.Attribute{
					Name:        "displayname",
					Description: `The displayed name in the OVHcloud web admin`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of IPs addresses attached to the vps`,
				},
				resource.Attribute{
					Name:        "keymap",
					Description: `The keymap for the ip kvm, valid values "", "fr", "us"`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of memory in MB of the vps.`,
				},
				resource.Attribute{
					Name:        "model",
					Description: `A dict describing the type of vps.`,
				},
				resource.Attribute{
					Name:        "model.name",
					Description: `The model name (ex: model1)`,
				},
				resource.Attribute{
					Name:        "model.offer",
					Description: `The model human description (ex: "VPS 2016 SSD 1")`,
				},
				resource.Attribute{
					Name:        "model.version",
					Description: `The model version (ex: "2017v2")`,
				},
				resource.Attribute{
					Name:        "netbootmode",
					Description: `The source of the boot kernel`,
				},
				resource.Attribute{
					Name:        "offertype",
					Description: `The type of offer (ssd, cloud, classic)`,
				},
				resource.Attribute{
					Name:        "slamonitoring",
					Description: `A boolean to indicate if OVHcloud SLA monitoring is active.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the vps`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of server`,
				},
				resource.Attribute{
					Name:        "vcore",
					Description: `The number of vcore of the vps`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `The OVHcloud zone where the vps is`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vpss",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of VPS IDs associated with your OVH Account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of VPS IDs associated with your OVH Account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vracks",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of vrack service name available for your OVHcloud account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "result",
					Description: `The list of vrack service name available for your OVHcloud account.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ovh_cloud_project_capabilities_containerregistry":        0,
		"ovh_cloud_project_capabilities_containerregistry_filter": 1,
		"ovh_cloud_project_containerregistries":                   2,
		"ovh_cloud_project_containerregistry":                     3,
		"ovh_cloud_project_containerregistry_users":               4,
		"ovh_cloud_project_database":                              5,
		"ovh_cloud_project_database_capabilities":                 6,
		"ovh_cloud_project_database_certificates":                 7,
		"ovh_cloud_project_database_database":                     8,
		"ovh_cloud_project_database_databases":                    9,
		"ovh_cloud_project_database_integration":                  10,
		"ovh_cloud_project_database_integrations":                 11,
		"ovh_cloud_project_database_ip_restrictions":              12,
		"ovh_cloud_project_database_kafka_acl":                    13,
		"ovh_cloud_project_database_kafka_acls":                   14,
		"ovh_cloud_project_database_kafka_topic":                  15,
		"ovh_cloud_project_database_kafka_topics":                 16,
		"ovh_cloud_project_database_kafka_user_access":            17,
		"ovh_cloud_project_database_m3db_namespace":               18,
		"ovh_cloud_project_database_m3db_namespaces":              19,
		"ovh_cloud_project_database_m3db_user":                    20,
		"ovh_cloud_project_database_mongodb_user":                 21,
		"ovh_cloud_project_database_opensearch_pattern":           22,
		"ovh_cloud_project_database_opensearch_patterns":          23,
		"ovh_cloud_project_database_opensearch_user":              24,
		"ovh_cloud_project_database_postgresql_user":              25,
		"ovh_cloud_project_database_redis_user":                   26,
		"ovh_cloud_project_database_user":                         27,
		"ovh_cloud_project_database_users":                        28,
		"ovh_cloud_project_databases":                             29,
		"ovh_cloud_project_failover_ip_attach":                    30,
		"ovh_cloud_project_kube":                                  31,
		"ovh_cloud_project_kube_iprestrictions":                   32,
		"ovh_cloud_project_kube_nodepool":                         33,
		"ovh_cloud_project_kube_nodepool_nodes":                   34,
		"ovh_cloud_project_kube_nodes":                            35,
		"ovh_cloud_project_kube_oidc":                             36,
		"ovh_cloud_project_region":                                37,
		"ovh_cloud_project_regions":                               38,
		"ovh_cloud_project_user":                                  39,
		"ovh_cloud_project_user_s3_credential":                    40,
		"ovh_cloud_project_user_s3_credentials":                   41,
		"ovh_cloud_project_user_s3_policy":                        42,
		"ovh_cloud_project_users":                                 43,
		"ovh_dbaas_logs_cluster":                                  44,
		"ovh_dbaas_logs_input_engine":                             45,
		"ovh_dbaas_logs_output_graylog_stream":                    46,
		"ovh_dedicated_ceph":                                      47,
		"ovh_dedicated_installation_templates":                    48,
		"ovh_dedicated_nasha":                                     49,
		"ovh_dedicated_server":                                    50,
		"ovh_dedicated_server_boots":                              51,
		"ovh_dedicated_servers":                                   52,
		"ovh_domain_zone":                                         53,
		"ovh_hosting_privatedatabase":                             54,
		"ovh_hosting_privatedatabase_database":                    55,
		"ovh_hosting_privatedatabase_user":                        56,
		"ovh_hosting_privatedatabase_user_grant":                  57,
		"ovh_hosting_privatedatabase_whitelist":                   58,
		"ovh_ip_service":                                          59,
		"ovh_iploadbalancing":                                     60,
		"ovh_iploadbalancing_vrack_network":                       61,
		"ovh_iploadbalancing_vrack_networks":                      62,
		"ovh_me":                                                  63,
		"ovh_me_identity_user":                                    64,
		"ovh_me_identity_users":                                   65,
		"ovh_me_installation_template":                            66,
		"ovh_me_installation_templates":                           67,
		"ovh_me_ipxe_script":                                      68,
		"ovh_me_ipxe_scripts":                                     69,
		"ovh_me_paymentmean_bankaccount":                          70,
		"ovh_me_paymentmean_creditcard":                           71,
		"ovh_me_ssh_key":                                          72,
		"ovh_me_ssh_keys":                                         73,
		"ovh_order_cart":                                          74,
		"ovh_order_cart_product":                                  75,
		"ovh_order_cart_product_options":                          76,
		"ovh_order_cart_product_options_plan":                     77,
		"ovh_order_cart_product_plan":                             78,
		"ovh_vps":                                                 79,
		"ovh_vpss":                                                80,
		"ovh_vracks":                                              81,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
