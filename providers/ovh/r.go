package ovh

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description associated with the user.`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM`,
				},
				resource.Attribute{
					Name:        "plan_option",
					Description: `(Optional) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the order Id. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `project access right for the identity that trigger the terraform script.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Project description`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about the order that was used to create the public cloud project`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id, the same as the ` + "`" + `id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `openstack project name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `openstack project id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `project status ## Import Cloud project can be imported using the ` + "`" + `order_id` + "`" + ` that can be retrieved in the [order page](https://www.ovh.com/manager/#/dedicated/billing/orders/orders) at the creation time of the Public Cloud project. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project.my_cloud_project order_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "access",
					Description: `project access right for the identity that trigger the terraform script.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Project description`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about the order that was used to create the public cloud project`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id, the same as the ` + "`" + `id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `openstack project name`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `openstack project id`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `project status ## Import Cloud project can be imported using the ` + "`" + `order_id` + "`" + ` that can be retrieved in the [order page](https://www.ovh.com/manager/#/dedicated/billing/orders/orders) at the creation time of the Public Cloud project. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project.my_cloud_project order_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_containerregistry",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Registry name`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region of the registry`,
				},
				resource.Attribute{
					Name:        "plan_id",
					Description: `Plan ID of the registry ## Attributes Reference The following attributes are exported:`,
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
					Name:        "plan",
					Description: `Plan of the registry`,
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
					Name:        "plan",
					Description: `Plan of the registry`,
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
			Type:             "ovh_cloud_project_containerregistry_user",
			Category:         "Resources",
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
					Name:        "email",
					Description: `User email`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `User ID`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) User password`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `User name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `User email`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `User ID`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) User password`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `User name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Small description of the database service.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required, Forces new resource) The database engine you want to deploy. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) A valid OVHcloud public cloud database flavor name in which the nodes will be started. Ex: "db1-7". Changing this value upgrade the nodes with the new flavor. You can find the list of flavor names: https://www.ovhcloud.com/fr/public-cloud/prices/`,
				},
				resource.Attribute{
					Name:        "kafka_rest_api",
					Description: `(Optional) Defines whether the REST API is enabled on a kafka cluster`,
				},
				resource.Attribute{
					Name:        "nodes",
					Description: `(Required, Minimum Items: 1) List of nodes object. Multi region cluster are not yet available, all node should be identical.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional, Forces new resource) Private network id in which the node should be deployed. It's the regional openstackId of the private network`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required, Forces new resource) Public cloud region in which the node should be deployed. Ex: "GRA'.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional, Forces new resource) Private subnet ID in which the node is.`,
				},
				resource.Attribute{
					Name:        "opensearch_acls_enabled",
					Description: `(Optional) Defines whether the ACLs are enabled on an OpenSearch cluster`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `(Optional) The disk size (in GB) of the database service.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `(Optional) Advanced configuration key / value.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Plan of the cluster. Enum: "essential", "business", "enterprise".`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the engine in which the service should be deployed ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Public Cloud Database Service ID`,
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
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kafka_rest_api",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "opensearch_acls_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Defines the disk type of the database service.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Advanced configuration key / value. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database" "db" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 40m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `id` + "`" + ` of the cluster, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database.my_database_cluster service_name/engine/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Public Cloud Database Service ID`,
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
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "kafka_rest_api",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "opensearch_acls_enabled",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Defines the disk type of the database service.`,
				},
				resource.Attribute{
					Name:        "advanced_configuration",
					Description: `Advanced configuration key / value. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database" "db" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 40m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `id` + "`" + ` of the cluster, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database.my_database_cluster service_name/engine/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_database",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required, Forces new resource) The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the database. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
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
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_database" "database" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters databases can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the database, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_database.my_database service_name/engine/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
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
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_database" "database" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters databases can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the database, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_database.my_database service_name/engine/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_integration",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required, Forces new resource) The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept ` + "`" + `mongodb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `(Required, Forces new resource) ID of the destination service.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `(Optional, Forces new resource) Parameters for the integration.`,
				},
				resource.Attribute{
					Name:        "source_service_id",
					Description: `(Required, Forces new resource) ID of the source service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional, Forces new resource) Type of the integration. Available types:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `- ID of the integration.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_user.my_user service_name/engine/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "destination_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `- ID of the integration.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "source_service_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the integration.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_user.my_user service_name/engine/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_ip_restriction",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required, Forces new resource) The engine of the database cluster you want to add an IP restriction. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required, Forces new resource) Authorized IP.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the IP restriction. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the IP restriction. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_ip_restriction" "iprestriction" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database cluster IP restrictions can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and the ` + "`" + `ip` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_ip_restriction.my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the IP restriction. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_ip_restriction" "iprestriction" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database cluster IP restrictions can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and the ` + "`" + `ip` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_ip_restriction.my_ip_restriction service_name/engine/cluster_id/178.97.6.0/24 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_acl",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required, Forces new resource) Permission to give to this username on this topic. Available permissions:`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `(Required, Forces new resource) Topic affected by this ACL.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required, Forces new resource) Username affected by this ACL. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_kafka_acl" "acl" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed kafka clusters ACLs can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the acl, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_kafka_acl.my_acl service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "topic",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_kafka_acl" "acl" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed kafka clusters ACLs can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the acl, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_kafka_acl.my_acl service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_kafka_topic",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the topic. No spaces allowed.`,
				},
				resource.Attribute{
					Name:        "min_insync_replicas",
					Description: `(Optional, Forces new resource) Minimum insync replica accepted for this topic. Should be superior to 0`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `(Optional, Forces new resource) Number of partitions for this topic. Should be superior to 0`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `(Optional, Forces new resource) Number of replication for this topic. Should be superior to 1`,
				},
				resource.Attribute{
					Name:        "retention_bytes",
					Description: `(Optional, Forces new resource) Number of bytes for the retention of the data for this topic. Inferior to 0 means unlimited`,
				},
				resource.Attribute{
					Name:        "retention_hours",
					Description: `(Optional, Forces new resource) Number of hours for the retention of the data for this topic. Should be superior to -2. Inferior to 0 means unlimited ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the topic.`,
				},
				resource.Attribute{
					Name:        "min_insync_replicas",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_bytes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_hours",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_kafka_topic" "topic" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed kafka clusters topics can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the topic, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_kafka_topic.my_topic service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the topic.`,
				},
				resource.Attribute{
					Name:        "min_insync_replicas",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "replication",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_bytes",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_hours",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_kafka_topic" "topic" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed kafka clusters topics can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the topic, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_kafka_topic.my_topic service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_m3db_namespace",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the namespace.`,
				},
				resource.Attribute{
					Name:        "resolution",
					Description: `(Optional) Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.`,
				},
				resource.Attribute{
					Name:        "retention_block_data_expiration_duration",
					Description: `(Optional) Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.`,
				},
				resource.Attribute{
					Name:        "retention_block_size_duration",
					Description: `(Optional) Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_future_duration",
					Description: `(Optional) Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_past_duration",
					Description: `(Optional) Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.`,
				},
				resource.Attribute{
					Name:        "retention_period_duration",
					Description: `(Required) Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.`,
				},
				resource.Attribute{
					Name:        "snapshot_enabled",
					Description: `(Optional) Defines whether M3DB will create snapshot files for this namespace.`,
				},
				resource.Attribute{
					Name:        "writes_to_commit_log_enabled",
					Description: `(Optional) Defines whether M3DB will include writes to this namespace in the commit log. ## Attributes Reference The following attributes are exported:`,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_block_data_expiration_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_block_size_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_future_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_past_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_period_duration",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_m3db_namespace" "namespace" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed M3DB clusters namespaces can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the namespace, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_m3db_namespace.my_namespace service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_block_data_expiration_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_block_size_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_future_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_buffer_past_duration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "retention_period_duration",
					Description: `See Argument Reference above.`,
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
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_m3db_namespace" "namespace" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed M3DB clusters namespaces can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the namespace, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_m3db_namespace.my_namespace service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_m3db_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) Group of the user:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the user. A user named "avnadmin" is map with already created admin user instead of create a new user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Arbitrary string to change to trigger a password update. Use the ` + "`" + `terraform refresh` + "`" + ` command after executing ` + "`" + `terraform apply` + "`" + ` to update the output with the new password. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_m3db_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed M3DB clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_m3db_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_m3db_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed M3DB clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_m3db_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_mongodb_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the user.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional: if omit, default role) Roles the user belongs to. Available roles:`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Arbitrary string to change to trigger a password update. Use the ` + "`" + `terraform refresh` + "`" + ` command after executing ` + "`" + `terraform apply` + "`" + ` to update the output with the new password. ## Attributes Reference The following attributes are exported:`,
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
					Description: `Name of the user with the authentication database in the format name@authDB`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_mongodb_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed MongoDB clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_mongodb_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Name of the user with the authentication database in the format name@authDB`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_mongodb_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed MongoDB clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_mongodb_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_opensearch_pattern",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "max_index_count",
					Description: `(Optional, Forces new resource) Maximum number of index for this pattern.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required, Forces new resource) Pattern format. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the pattern.`,
				},
				resource.Attribute{
					Name:        "max_index_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_opensearch_pattern" "pattern" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed opensearch clusters patterns can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the pattern, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_opensearch_pattern.my_pattern service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the pattern.`,
				},
				resource.Attribute{
					Name:        "max_index_count",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_opensearch_pattern" "pattern" { # ... timeouts { create = "1h" delete = "45m" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed opensearch clusters patterns can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the pattern, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_opensearch_pattern.my_pattern service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_opensearch_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `(Optional) Acls of the user.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `(Required) Pattern of the ACL.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `(Required) Permission of the ACL Available permission:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Username affected by this acl. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Arbitrary string to change to trigger a password update. Use the ` + "`" + `terraform refresh` + "`" + ` command after executing ` + "`" + `terraform apply` + "`" + ` to update the output with the new password. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "acls",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `See Argument Reference above.`,
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
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_opensearch_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed OpenSearch clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_opensearch_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "acls",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "permission",
					Description: `See Argument Reference above.`,
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
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_opensearch_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed OpenSearch clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_opensearch_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_postgresql_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `(Optional: if omit, default role) Roles the user belongs to. Available roles:`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Arbitrary string to change to trigger a password update. Use the ` + "`" + `terraform refresh` + "`" + ` command after executing ` + "`" + `terraform apply` + "`" + ` to update the output with the new password. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `Arbitrary string to change to trigger a password update.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Roles the user belongs to.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_postgresql_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed PostgreSQL clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_postgresql_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `Arbitrary string to change to trigger a password update.`,
				},
				resource.Attribute{
					Name:        "roles",
					Description: `Roles the user belongs to.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_postgresql_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed PostgreSQL clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_postgresql_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_redis_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Optional) Categories of the user.`,
				},
				resource.Attribute{
					Name:        "channels",
					Description: `(Optional: if omit, all channels) Channels of the user.`,
				},
				resource.Attribute{
					Name:        "commands",
					Description: `(Optional) Commands of the user.`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `(Optional) Keys of the user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Arbitrary string to change to trigger a password update. Use the ` + "`" + `terraform refresh` + "`" + ` command after executing ` + "`" + `terraform apply` + "`" + ` to update the output with the new password. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "channels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "commands",
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
					Name:        "keys",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `Arbitrary string to change to trigger a password update.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_redis_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed Redis clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_redis_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "categories",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "channels",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "commands",
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
					Name:        "keys",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `Arbitrary string to change to trigger a password update.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_redis_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed Redis clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_redis_user.my_user service_name/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_database_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required, Forces new resource) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `(Required, Forces new resource) The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). Available engines:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required, Forces new resource) Cluster ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required, Forces new resource) Name of the user. A user named "avnadmin" is map with already created admin user and reset his password instead of create a new user. The "Grafana" engine only allows the "avnadmin" mapping.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `(Optional) Arbitrary string to change to trigger a password update. Use the ` + "`" + `terraform refresh` + "`" + ` command after executing ` + "`" + `terraform apply` + "`" + ` to update the output with the new password. ## Attributes Reference The following attributes are exported:`,
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
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `Arbitrary string to change to trigger a password update.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_user.my_user service_name/engine/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Sensitive) Password of the user.`,
				},
				resource.Attribute{
					Name:        "password_reset",
					Description: `Arbitrary string to change to trigger a password update.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the user. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_database_user" "user" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 20m) ## Import OVHcloud Managed database clusters users can be imported using the ` + "`" + `service_name` + "`" + `, ` + "`" + `engine` + "`" + `, ` + "`" + `cluster_id` + "`" + ` and ` + "`" + `id` + "`" + ` of the user, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_database_user.my_user service_name/engine/cluster_id/id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_failover_ip_attach",
			Category:         "Resources",
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
					Description: `The failover ip address to attach`,
				},
				resource.Attribute{
					Name:        "routed_to",
					Description: `The GUID of an instance to which the failover IP address is be attached ## Attributes Reference`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.`,
				},
				resource.Attribute{
					Name:        "kube_proxy_mode",
					Description: `(Optional) Selected mode for kube-proxy.`,
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
					Description: `(Optional) Kubernetes API server admission plugins customization`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Array of admission plugins enabled, default is ["NodeRestriction","AlwaysPulImages"] and only these admission plugins can be enabled at this time.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.`,
				},
				resource.Attribute{
					Name:        "customization_kube_proxy",
					Description: `Kubernetes kube-proxy customization`,
				},
				resource.Attribute{
					Name:        "iptables",
					Description: `(Optional) Kubernetes cluster kube-proxy customization of iptables specific config (durations format is RFC3339 duration, e.g. ` + "`" + `PT60S` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "sync_period",
					Description: `(Optional) Minimum period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. ` + "`" + `PT60S` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "min_sync_period",
					Description: `(Optional) Period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. ` + "`" + `PT60S` + "`" + `). Must be greater than 0.`,
				},
				resource.Attribute{
					Name:        "ipvs",
					Description: `(Optional) Kubernetes cluster kube-proxy customization of IPVS specific config (durations format is [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration, e.g. ` + "`" + `PT60S` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "sync_period",
					Description: `(Optional) Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. ` + "`" + `PT60S` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "min_sync_period",
					Description: `(Optional) Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration (e.g. ` + "`" + `PT60S` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "scheduler",
					Description: `(Optional) IPVS scheduler.`,
				},
				resource.Attribute{
					Name:        "tcp_timeout",
					Description: `(Optional) Timeout value used for idle IPVS TCP sessions in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration (e.g. ` + "`" + `PT60S` + "`" + `). The default value is ` + "`" + `PT0S` + "`" + `, which preserves the current timeout value on the system.`,
				},
				resource.Attribute{
					Name:        "tcp_fin_timeout",
					Description: `(Optional) Timeout value used for IPVS TCP sessions after receiving a FIN in RFC3339 duration (e.g. ` + "`" + `PT60S` + "`" + `). The default value is ` + "`" + `PT0S` + "`" + `, which preserves the current timeout value on the system.`,
				},
				resource.Attribute{
					Name:        "udp_timeout",
					Description: `(Optional) timeout value used for IPVS UDP packets in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration (e.g. ` + "`" + `PT60S` + "`" + `). The default value is ` + "`" + `PT0S` + "`" + `, which preserves the current timeout value on the system.`,
				},
				resource.Attribute{
					Name:        "private_network_id",
					Description: `(Optional) OpenStack private network (or vRack) ID to use.`,
				},
				resource.Attribute{
					Name:        "private_network_configuration",
					Description: `(Optional) The private network configuration`,
				},
				resource.Attribute{
					Name:        "default_vrack_gateway",
					Description: `If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.`,
				},
				resource.Attribute{
					Name:        "private_network_routing_as_default",
					Description: `Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE]. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "control_plane_is_up_to_date",
					Description: `True if control-plane is up-to-date.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Managed Kubernetes Service ID`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `True if all nodes and control-plane are up-to-date.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The kubeconfig file. Use this file to connect to your kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "kubeconfig_attributes",
					Description: `The kubeconfig file attributes.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The kubernetes API server URL.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `The kubernetes API server CA certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `The kubernetes API server client certificate.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `The kubernetes API server client key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
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
					Name:        "private_network_configuration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status. Should be normally set to 'READY'.`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Management URL of your cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "customization_apiserver",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "customization_kube_proxy",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_kube" "my_kube_cluster" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 10m) ## Import OVHcloud Managed Kubernetes Service clusters can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `id` + "`" + ` of the cluster, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube.my_kube_cluster service_name/kube_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "control_plane_is_up_to_date",
					Description: `True if control-plane is up-to-date.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Managed Kubernetes Service ID`,
				},
				resource.Attribute{
					Name:        "is_up_to_date",
					Description: `True if all nodes and control-plane are up-to-date.`,
				},
				resource.Attribute{
					Name:        "kubeconfig",
					Description: `The kubeconfig file. Use this file to connect to your kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "kubeconfig_attributes",
					Description: `The kubeconfig file attributes.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The kubernetes API server URL.`,
				},
				resource.Attribute{
					Name:        "cluster_ca_certificate",
					Description: `The kubernetes API server CA certificate.`,
				},
				resource.Attribute{
					Name:        "client_certificate",
					Description: `The kubernetes API server client certificate.`,
				},
				resource.Attribute{
					Name:        "client_key",
					Description: `The kubernetes API server client key.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
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
					Name:        "private_network_configuration",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "private_network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Cluster status. Should be normally set to 'READY'.`,
				},
				resource.Attribute{
					Name:        "update_policy",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Management URL of your cluster.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "customization_apiserver",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "customization_kube_proxy",
					Description: `See Argument Reference above. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_kube" "my_kube_cluster" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 10m) ## Import OVHcloud Managed Kubernetes Service clusters can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `id` + "`" + ` of the cluster, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube.my_kube_cluster service_name/kube_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_iprestrictions",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The id of the managed Kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `List of CIDR authorized to interact with the managed Kubernetes cluster. ## Attributes Reference No additional attributes than the ones provided are exported. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_kube_iprestrictions" "vrack_only" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 5m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 5m) ## Import OVHcloud Managed Kubernetes Service cluster IP restrictions can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `id` + "`" + ` of the cluster, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube_iprestrictions.iprestrictions service_name/kube_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "create",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 5m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 5m) ## Import OVHcloud Managed Kubernetes Service cluster IP restrictions can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `id` + "`" + ` of the cluster, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube_iprestrictions.iprestrictions service_name/kube_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_nodepool",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The id of the managed kubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the nodepool. Warning: ` + "`" + `_` + "`" + ` char is not allowed!`,
				},
				resource.Attribute{
					Name:        "flavor_name",
					Description: `a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: "b2-7". You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/.`,
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
					Name:        "template",
					Description: `(Optional) Managed Kubernetes nodepool template, which is a complex object constituted by two main nested objects:`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) Metadata of each node in the pool`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Optional) Annotations to apply to each node`,
				},
				resource.Attribute{
					Name:        "finalizers",
					Description: `(Optional) Finalizers to apply to each node`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) Labels to apply to each node`,
				},
				resource.Attribute{
					Name:        "spec",
					Description: `(Optional) Spec of each node in the pool`,
				},
				resource.Attribute{
					Name:        "taints",
					Description: `(Optional) Taints to apply to each node`,
				},
				resource.Attribute{
					Name:        "unschedulable",
					Description: `(Optional) If true, set nodes as un-schedulable ## Attributes Reference In addition, the following attributes are exported:`,
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
					Description: `Last update date ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_kube_nodepool" "pool" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 10m) ## Import OVHcloud Managed Kubernetes Service cluster node pool can be imported using the ` + "`" + `service_name` + "`" + `, the ` + "`" + `id` + "`" + ` of the cluster, and the ` + "`" + `id` + "`" + ` of the nodepool separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube_nodepool.pool service_name/kube_id/poolid ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `Last update date ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_kube_nodepool" "pool" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 20m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 10m) ## Import OVHcloud Managed Kubernetes Service cluster node pool can be imported using the ` + "`" + `service_name` + "`" + `, the ` + "`" + `id` + "`" + ` of the cluster, and the ` + "`" + `id` + "`" + ` of the nodepool separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube_nodepool.pool service_name/kube_id/poolid ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_kube_oidc",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The ID of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "kube_id",
					Description: `The ID of the managed kubernetes cluster.`,
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
					Description: `JWT claim to use as the username. By default, sub, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as email or name, depending on their provider. However, claims other than email will be prefixed with the issuer URL to prevent naming clashes with other plugins.`,
				},
				resource.Attribute{
					Name:        "oidcUsernamePrefix",
					Description: `Prefix prepended to username claims to prevent clashes with existing names (such as ` + "`" + `system:users` + "`" + `). For example, the value ` + "`" + `oidc:` + "`" + ` will create usernames like ` + "`" + `oidc:jane.doe` + "`" + `. If this field isn't set and ` + "`" + `oidcUsernameClaim` + "`" + ` is a value other than email the prefix defaults to ` + "`" + `issuer_url` + "`" + ` where ` + "`" + `issuer_url` + "`" + ` is the value of ` + "`" + `oidcIssuerUrl.` + "`" + ` The value - can be used to disable all prefixing.`,
				},
				resource.Attribute{
					Name:        "oidcGroupsClaim",
					Description: `Array of JWT claim to use as the user's group. If the claim is present it must be an array of strings.`,
				},
				resource.Attribute{
					Name:        "oidcGroupsPrefix",
					Description: `Prefix prepended to group claims to prevent clashes with existing names (such as ` + "`" + `system:groups` + "`" + `). For example, the value ` + "`" + `oidc:` + "`" + ` will create group names like ` + "`" + `oidc:engineering` + "`" + ` and ` + "`" + `oidc:infra` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oidcRequiredClaim",
					Description: `Array of ` + "`" + `key=value` + "`" + ` pairs that describe required claims in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value."`,
				},
				resource.Attribute{
					Name:        "oidcSigningAlgs",
					Description: `Array of signing algorithms accepted. Default is ` + "`" + `RS256` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oidcCaContent",
					Description: `Content of the certificate for the CA, in Base64 format, that signed your identity provider's web certificate. Defaults to the host's root CAs. ## Timeouts ` + "`" + `` + "`" + `` + "`" + `hcl resource "ovh_cloud_project_kube_oidc" "oidc" { # ... timeouts { create = "1h" update = "45m" delete = "50s" } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "update",
					Description: `(Default 10m)`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Default 10m) ## Import OVHcloud Managed Kubernetes Service cluster OIDC can be imported using the tenant ` + "`" + `service_name` + "`" + ` and cluster id ` + "`" + `kube_id` + "`" + ` separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_kube_oidc.my-oidc service_name/kube_id ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_network_private",
			Category:         "Resources",
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
					Description: `(Required) The name of the network.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `a vlan id to associate with the network. Changing this value recreates the resource. Defaults to 0.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `an array of valid OVHcloud public cloud region ID in which the network will be available. Ex.: "GRA1". Defaults to all public cloud regions. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the network`,
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
					Name:        "regions_attributes",
					Description: `A map representing information about the region.`,
				},
				resource.Attribute{
					Name:        "regions_attributes/region",
					Description: `The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_attributes/status",
					Description: `The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "regions_attributes/openstackid",
					Description: `The private network id in the region.`,
				},
				resource.Attribute{
					Name:        "regions_status",
					Description: `(Deprecated) A map representing the status of the network per region.`,
				},
				resource.Attribute{
					Name:        "regions_status/region",
					Description: `(Deprecated) The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_status/status",
					Description: `(Deprecated) The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the network. should be normally set to 'ACTIVE'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the network. Either 'private' or 'public'. ## Import Private network in a public cloud project can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `network_id` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_network_private.mynet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the network`,
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
					Name:        "regions_attributes",
					Description: `A map representing information about the region.`,
				},
				resource.Attribute{
					Name:        "regions_attributes/region",
					Description: `The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_attributes/status",
					Description: `The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "regions_attributes/openstackid",
					Description: `The private network id in the region.`,
				},
				resource.Attribute{
					Name:        "regions_status",
					Description: `(Deprecated) A map representing the status of the network per region.`,
				},
				resource.Attribute{
					Name:        "regions_status/region",
					Description: `(Deprecated) The id of the region.`,
				},
				resource.Attribute{
					Name:        "regions_status/status",
					Description: `(Deprecated) The status of the network in the region.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `the status of the network. should be normally set to 'ACTIVE'.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the network. Either 'private' or 'public'. ## Import Private network in a public cloud project can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `network_id` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_network_private.mynet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_network_private_subnet",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
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
					Description: `First ip for this region. ## Import Subnet in a private network of a public cloud project can be imported using the ` + "`" + `service_name` + "`" + ` , the ` + "`" + `network_id` + "`" + ` and the ` + "`" + `subnet_id` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_network_private_subnet.mysubnet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/0f0b73a4-403b-45e4-86d0-b438f1291909 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `First ip for this region. ## Import Subnet in a private network of a public cloud project can be imported using the ` + "`" + `service_name` + "`" + ` , the ` + "`" + `network_id` + "`" + ` and the ` + "`" + `subnet_id` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_cloud_project_network_private_subnet.mysubnet ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/0f0b73a4-403b-45e4-86d0-b438f1291909 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_region_storage_presign",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Required) The region in which your storage is located. Ex.: "GRA".`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of your S3 storage container/bucket.`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `(Required) Define, in seconds, for how long your URL will be valid.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The method you want to use to interact with your object. Can be either 'GET' or 'PUT'.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `(Required) The name of the object in your S3 bucket. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Computed URL result.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "object",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `Computed URL result.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description associated with the user.`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "role_name",
					Description: `The name of a role. See ` + "`" + `role_names` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role_names",
					Description: `A list of role names. Values can be: - administrator, - ai_training_operator - authentication - backup_operator - compute_operator - image_operator - infrastructure_supervisor - network_operator - network_security_operator - objectstore_operator - volume_operator ## Attributes Reference The following attributes are exported:`,
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
			Type:             "ovh_cloud_project_user_s3_credential",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_user_s3_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_cloud_project_workflow_backup",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Optional) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "region_name",
					Description: `(Mandatory) The name of the openstack region.`,
				},
				resource.Attribute{
					Name:        "cron",
					Description: `(Mandatory) The cron periodicity at which the backup workflow is scheduled`,
				},
				resource.Attribute{
					Name:        "max_execution_count",
					Description: `(Optional) The number of times the worflow is run. Default value is ` + "`" + `0` + "`" + ` which means that the workflow will be scheduled continously until its deletion`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Mandatory) The worflow name that is used in the UI`,
				},
				resource.Attribute{
					Name:        "backup_name",
					Description: `(Optional) The name of the backup files that are created. If empty, the ` + "`" + `name` + "`" + ` attribute is used.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dbaas_logs_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "archive_allowed_networks",
					Description: `List of IP blocks`,
				},
				resource.Attribute{
					Name:        "direct_input_allowed_networks",
					Description: `List of IP blocks`,
				},
				resource.Attribute{
					Name:        "query_allowed_networks",
					Description: `List of IP blocks ## Attributes Reference Id is set to the input Id. In addition, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dbaas_logs_graylog_output_stream",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Stream description`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Stream description`,
				},
				resource.Attribute{
					Name:        "parent_stream_id",
					Description: `Parent stream ID`,
				},
				resource.Attribute{
					Name:        "retention_id",
					Description: `Retention ID`,
				},
				resource.Attribute{
					Name:        "cold_storage_compression",
					Description: `Cold storage compression method. One of "LZMA", "GZIP", "DEFLATED", "ZSTD"`,
				},
				resource.Attribute{
					Name:        "cold_storage_content",
					Description: `ColdStorage content. One of "ALL", "GLEF", "PLAIN"`,
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
					Description: `ColdStorage destination. One of "PCA", "PCS"`,
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
					Name:        "pause_indexing_on_max_size",
					Description: `If set, pause indexing when maximum size is reach`,
				},
				resource.Attribute{
					Name:        "web_socket_enabled",
					Description: `Enable Websocket ## Attributes Reference Id is set to the output stream Id. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "can_alert",
					Description: `Indicates if the current user can create alert on the stream`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Stream creation`,
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
					Description: `Number of coldstored archivesr`,
				},
				resource.Attribute{
					Name:        "stream_id",
					Description: `Stream ID`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Stream last updater`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "can_alert",
					Description: `Indicates if the current user can create alert on the stream`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Stream creation`,
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
					Description: `Number of coldstored archivesr`,
				},
				resource.Attribute{
					Name:        "stream_id",
					Description: `Stream ID`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Stream last updater`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dbaas_logs_input",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allowed_networks",
					Description: `List of IP blocks`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Required) Input configuration`,
				},
				resource.Attribute{
					Name:        "flowgger",
					Description: `(Optional) Flowgger configuration`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Type of format to decode. One of "RFC5424", "LTSV", "GELF", "CAPNP"`,
				},
				resource.Attribute{
					Name:        "log_framing",
					Description: `Indicates how messages are delimited. One of "LINE", "NUL", "SYSLEN", "CAPNP"`,
				},
				resource.Attribute{
					Name:        "logstash",
					Description: `(Optional) Logstash configuration`,
				},
				resource.Attribute{
					Name:        "filter_section",
					Description: `(Optional) The filter section of logstash.conf`,
				},
				resource.Attribute{
					Name:        "input_section",
					Description: `(Required) The filter section of logstash.conf`,
				},
				resource.Attribute{
					Name:        "pattern_section",
					Description: `(Optional) The list of customs Grok patterns`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Input description`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `(Required) Input engine ID`,
				},
				resource.Attribute{
					Name:        "exposed_port",
					Description: `Port`,
				},
				resource.Attribute{
					Name:        "nb_instance",
					Description: `Number of instance running`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) service name`,
				},
				resource.Attribute{
					Name:        "stream_id",
					Description: `(Required) Associated Graylog stream`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) Input title ## Attributes Reference Id is set to the input Id. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Input creation`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname`,
				},
				resource.Attribute{
					Name:        "input_id",
					Description: `Input ID`,
				},
				resource.Attribute{
					Name:        "is_restart_required",
					Description: `Indicate if input need to be restarted`,
				},
				resource.Attribute{
					Name:        "public_address",
					Description: `Input IP address`,
				},
				resource.Attribute{
					Name:        "ssl_certificate",
					Description: `Input SSL certificate`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `init: configuration required, pending: ready to start, running: available`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Input last update`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Input creation`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Hostname`,
				},
				resource.Attribute{
					Name:        "input_id",
					Description: `Input ID`,
				},
				resource.Attribute{
					Name:        "is_restart_required",
					Description: `Indicate if input need to be restarted`,
				},
				resource.Attribute{
					Name:        "public_address",
					Description: `Input IP address`,
				},
				resource.Attribute{
					Name:        "ssl_certificate",
					Description: `Input SSL certificate`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `init: configuration required, pending: ready to start, running: available`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Input last update`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_ceph_acl",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "ovh_dedicated_nasha_partition",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) name of the partition`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) size of the partition in GB`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Required) one of "NFS", "CIFS" or "NFS_CIFS"`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A brief description of the partition ## Attributes Reference The following attributes are exported:`,
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
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Percentage of partition space used in %`,
				},
				resource.Attribute{
					Name:        "used_by_snapshots",
					Description: `Percentage of partition space used by snapshots in % ## Import HA-NAS can be imported using the ` + "`" + `{service_name}/{name}` + "`" + `, e.g. ` + "`" + `$ terraform import ovh_dedicated_nasha_partition.my-partition zpool-12345/my-partition` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Percentage of partition space used in %`,
				},
				resource.Attribute{
					Name:        "used_by_snapshots",
					Description: `Percentage of partition space used by snapshots in % ## Import HA-NAS can be imported using the ` + "`" + `{service_name}/{name}` + "`" + `, e.g. ` + "`" + `$ terraform import ovh_dedicated_nasha_partition.my-partition zpool-12345/my-partition` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_nasha_partition_access",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `(Required) name of the partition`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) ip block in x.x.x.x/x format`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) one of "readwrite", "readonly" ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. ## Import HA-NAS partition access can be imported using the ` + "`" + `{service_name}/{partition_name}/{ip}` + "`" + `, e.g. ` + "`" + `$ terraform import ovh_dedicated_nasha_partition_access.my-partition zpool-12345/my-partition/123.123.123.123%2F32` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. ## Import HA-NAS partition access can be imported using the ` + "`" + `{service_name}/{partition_name}/{ip}` + "`" + `, e.g. ` + "`" + `$ terraform import ovh_dedicated_nasha_partition_access.my-partition zpool-12345/my-partition/123.123.123.123%2F32` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_nasha_partition_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your HA-NAS (it has to be ordered via OVHcloud interface)`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `(Required) name of the partition`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Snapshot interval, allowed : day-1, day-2, day-3, day-7, hour-1, hour-6 ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. ## Import HA-NAS partition snapshot can be imported using the ` + "`" + `{service_name}/{partition_name}/{type}` + "`" + `, e.g. ` + "`" + `$ terraform import ovh_dedicated_nasha_partition_snapshot.my-partition zpool-12345/my-partition/day-3` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `See Argument Reference above. ## Import HA-NAS partition snapshot can be imported using the ` + "`" + `{service_name}/{partition_name}/{type}` + "`" + `, e.g. ` + "`" + `$ terraform import ovh_dedicated_nasha_partition_snapshot.my-partition zpool-12345/my-partition/day-3` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_dedicated_server_install_task",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server.`,
				},
				resource.Attribute{
					Name:        "partition_scheme_name",
					Description: `Partition scheme name.`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `(Required) Template name.`,
				},
				resource.Attribute{
					Name:        "bootid_on_destroy",
					Description: `If set, reboot the server on the specified boot id during destroy phase.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `see ` + "`" + `details` + "`" + ` block below. The ` + "`" + `details` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "change_log",
					Description: `Template change log details.`,
				},
				resource.Attribute{
					Name:        "custom_hostname",
					Description: `Set up the server using the provided hostname instead of the default hostname.`,
				},
				resource.Attribute{
					Name:        "disk_group_id",
					Description: `Disk group id.`,
				},
				resource.Attribute{
					Name:        "install_rtm",
					Description: `set to true to install RTM.`,
				},
				resource.Attribute{
					Name:        "install_sql_server",
					Description: `set to true to install sql server (Windows template only).`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `language.`,
				},
				resource.Attribute{
					Name:        "no_raid",
					Description: `set to true to disable RAID.`,
				},
				resource.Attribute{
					Name:        "post_installation_script_link",
					Description: `Indicate the URL where your postinstall customisation script is located.`,
				},
				resource.Attribute{
					Name:        "post_installation_script_return",
					Description: `Indicate the string returned by your postinstall customisation script on successful execution. Advice: your script should return a unique validation string in case of succes. A good example is 'loh1Xee7eo OK OK OK UGh8Ang1Gu'.`,
				},
				resource.Attribute{
					Name:        "reset_hw_raid",
					Description: `set to true to make a hardware raid reset.`,
				},
				resource.Attribute{
					Name:        "soft_raid_devices",
					Description: `soft raid devices.`,
				},
				resource.Attribute{
					Name:        "ssh_key_name",
					Description: `Name of the ssh key that should be installed. Password login will be disabled.`,
				},
				resource.Attribute{
					Name:        "use_spla",
					Description: `set to true to use SPLA.`,
				},
				resource.Attribute{
					Name:        "use_distrib_kernel",
					Description: `Use the distribution's native kernel instead of the recommended OVHcloud Kernel. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The task id`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Details of this task. (should be ` + "`" + `Install asked` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "done_date",
					Description: `Completion date in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `Function name (should be ` + "`" + `hardInstall` + "`" + `).`,
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
					Description: `Details of this task. (should be ` + "`" + `Install asked` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "done_date",
					Description: `Completion date in RFC3339 format.`,
				},
				resource.Attribute{
					Name:        "function",
					Description: `Function name (should be ` + "`" + `hardInstall` + "`" + `).`,
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
			Type:             "ovh_dedicated_server_reboot_task",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service_name of your dedicated server.`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `List of values tracked to trigger reboot, used also to form implicit dependencies. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "ovh_hosting_privatedatabase",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Custom description on your privatedatabase order.`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration.`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code.`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM Plan order valid values can be found on OVHcloud [APIv6](https://api.ovh.com/console/#/hosting/privateDatabase/availableOrderCapacities~GET) ## Attributes Reference The following attributes are exported:`,
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
					Name:        "id",
					Description: `Private database id`,
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
					Name:        "order",
					Description: `Details about your Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Product Plan`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Service duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
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
					Name:        "id",
					Description: `Private database id`,
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
					Name:        "order",
					Description: `Details about your Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `Product Plan`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `Service duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `Pricing model identifier`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_database",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Name of your new database ## Attributes Reference The id is set to the value of ` + "`" + `service_name` + "`" + `/` + "`" + `database_name` + "`" + `. ## Import OVHcloud Webhosting database can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `database_name` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ovh_hosting_privatedatabase_database.database service_name/database_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_user",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the new user (alphanumeric, minimum one number and 8 characters minimum)`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) User name used to connect on your databases ## Attributes Reference The id is set to the value of ` + "`" + `service_name` + "`" + `/` + "`" + `user_name` + "`" + `. ## Import OVHcloud database user can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `user_name` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ovh_hosting_privatedatabase_user.user service_name/user_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_user_grant",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) User name used to connect on your databases.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Required) Database name where add grant.`,
				},
				resource.Attribute{
					Name:        "grant",
					Description: `(Required) Database name where add grant. Values can be: - admin - none - ro - rw ## Attributes Reference The id is set to the value of ` + "`" + `service_name` + "`" + `/` + "`" + `user_name` + "`" + `/` + "`" + `database_name` + "`" + `/` + "`" + `grant` + "`" + ` . ## Import OVHcloud database user's grant can be imported using the ` + "`" + `service_name` + "`" + `, the ` + "`" + `user_name` + "`" + `, the ` + "`" + `database_name` + "`" + ` and the ` + "`" + `grant` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ovh_hosting_privatedatabase_user_grant.user service_name/user_name/database_name/grant ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_hosting_privatedatabase_whitelist",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your private database.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Required) The whitelisted IP in your instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Custom name for your Whitelisted IP.`,
				},
				resource.Attribute{
					Name:        "service",
					Description: `(Required) Authorize this IP to access service port. Values can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `(Required) Authorize this IP to access SFTP port. Values can be ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + ` ## Attributes Reference The id is set to the value of ` + "`" + `service_name` + "`" + `/` + "`" + `ip` + "`" + `. ## Import OVHcloud database whitelist can be imported using the ` + "`" + `service_name` + "`" + ` and the ` + "`" + `ip` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ovh_hosting_privatedatabase_whitelist.ip service_name/ip ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_ip_reverse",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "ip_reverse",
					Description: `(Required) The IP to set the reverse of ## Attributes Reference The id is set to the value of ip_reverse. ## Import The resource can be imported using the ` + "`" + `ip` + "`" + `, ` + "`" + `ip_reverse` + "`" + ` of the address, separated by "|" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_ip_reverse.my_reverse '2001:0db8:c0ff:ee::/64|2001:0db8:c0ff:ee::42' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_ip_service",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Custom description on your ip.`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM`,
				},
				resource.Attribute{
					Name:        "plan_option",
					Description: `(Optional) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM ## Attributes Reference Id is set to the order Id. In addition, the following attributes are exported:`,
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
					Name:        "ip",
					Description: `ip block`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
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
					Description: `Possible values for ip type`,
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
					Name:        "ip",
					Description: `ip block`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
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
					Description: `Possible values for ip type`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `Set the name displayed in ManagerV6 for your iplb (max 50 chars)`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM`,
				},
				resource.Attribute{
					Name:        "plan_option",
					Description: `(Optional) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM`,
				},
				resource.Attribute{
					Name:        "ssl_configuration",
					Description: `Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern". ## Attributes Reference Id is set to the order Id. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `Your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPV4 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPV6 associated to your IP load balancing. DEPRECATED.`,
				},
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The offer of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
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
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of your IP`,
				},
				resource.Attribute{
					Name:        "vrack_eligibility",
					Description: `Vrack eligibility`,
				},
				resource.Attribute{
					Name:        "vrack_name",
					Description: `Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Location where your service is`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `Your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The IPV4 associated to your IP load balancing`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The IPV6 associated to your IP load balancing. DEPRECATED.`,
				},
				resource.Attribute{
					Name:        "metrics_token",
					Description: `The metrics token associated with your IP load balancing`,
				},
				resource.Attribute{
					Name:        "offer",
					Description: `The offer of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
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
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of your IP`,
				},
				resource.Attribute{
					Name:        "vrack_eligibility",
					Description: `Vrack eligibility`,
				},
				resource.Attribute{
					Name:        "vrack_name",
					Description: `Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Location where your service is`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_http_farm",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) Valid values : ` + "`" + `http` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `mysql` + "`" + `, ` + "`" + `oco` + "`" + `, ` + "`" + `pgsql` + "`" + `, ` + "`" + `smtp` + "`" + `, ` + "`" + `tcp` + "`" + ``,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `probe interval, Value between 30 and 3600 seconds, default 30`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `What to match ` + "`" + `pattern` + "`" + ` against (` + "`" + `contains` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `matches` + "`" + `, ` + "`" + `status` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port for backends to receive traffic on.`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `Address of the backend server (IP from either internal or OVHcloud network)`,
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
					Description: `version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (` + "`" + `v1` + "`" + `, ` + "`" + `v2` + "`" + `, ` + "`" + `v2-ssl` + "`" + `, ` + "`" + `v2-ssl-cn` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "on_marked_down",
					Description: `enable action when backend marked down. (` + "`" + `shutdown-sessions` + "`" + `)`,
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
					Name:        "on_marked_down",
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
					Name:        "on_marked_down",
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Name:        "hsts",
					Description: `HTTP Strict Transport Security. Default: 'false'`,
				},
				resource.Attribute{
					Name:        "redirect_location",
					Description: `Redirection HTTP'`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `HTTP headers to add to the frontend. List of string. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of your frontend`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `See Argument Reference above.`,
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
				resource.Attribute{
					Name:        "hsts",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of your frontend`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "http_header",
					Description: `See Argument Reference above.`,
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
				resource.Attribute{
					Name:        "hsts",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_http_route",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action triggered when all rules match`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `HTTP status code for "redirect" and "reject" actions`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Farm ID for "farm" action type or URL template for "redirect" action. You may use ${uri}, ${protocol}, ${host}, ${port} and ${path} variables in redirect target`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Action to trigger if all the rules of this route matches`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your route, this field is for you`,
				},
				resource.Attribute{
					Name:        "frontend_id",
					Description: `Route traffic for this frontend`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action ## Attributes Reference In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Route status. Routes in "ok" state are ready to operate`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of rules to match to trigger action`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `Matching operator. Not all operators are available for all fields. See "availableRules"`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Value to match against this match. Interpretation if this field depends on the match and field`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of your rule`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `Name of sub-field, if applicable. This may be a Cookie or Header name for instance`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Route status. Routes in "ok" state are ready to operate`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of rules to match to trigger action`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `Matching operator. Not all operators are available for all fields. See "availableRules"`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Value to match against this match. Interpretation if this field depends on the match and field`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of your rule`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `Name of sub-field, if applicable. This may be a Cookie or Header name for instance`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_http_route_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "keepers",
					Description: `List of values tracked to trigger refresh, used also to form implicit dependencies ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Required) Valid values : ` + "`" + `http` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `mysql` + "`" + `, ` + "`" + `oco` + "`" + `, ` + "`" + `pgsql` + "`" + `, ` + "`" + `smtp` + "`" + `, ` + "`" + `tcp` + "`" + ``,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `probe interval, Value between 30 and 3600 seconds, default 30`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `What to match ` + "`" + `pattern` + "`" + ` against (` + "`" + `contains` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `internal` + "`" + `, ` + "`" + `matches` + "`" + `, ` + "`" + `status` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port for backends to receive traffic on.`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `Address of the backend server (IP from either internal or OVHcloud network)`,
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
					Description: `version of the PROXY protocol used to pass origin connection information from loadbalancer to receiving service (` + "`" + `v1` + "`" + `, ` + "`" + `v2` + "`" + `, ` + "`" + `v2-ssl` + "`" + `, ` + "`" + `v2-ssl-cn` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "on_marked_down",
					Description: `enable action when backend marked down. (` + "`" + `shutdown-sessions` + "`" + `)`,
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
					Name:        "on_marked_down",
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
					Name:        "on_marked_down",
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "ovh_iploadbalancing_tcp_route",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "action",
					Description: `(Required) Action triggered when all rules match`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `Farm ID for "farm" action type, empty for others.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Action to trigger if all the rules of this route matches`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Human readable name for your route, this field is for you`,
				},
				resource.Attribute{
					Name:        "frontend_id",
					Description: `Route traffic for this frontend`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your IP load balancing`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action ## Attributes Reference In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Route status. Routes in "ok" state are ready to operate`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of rules to match to trigger action`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `Matching operator. Not all operators are available for all fields. See "availableRules"`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Value to match against this match. Interpretation if this field depends on the match and field`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of your rule`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `Name of sub-field, if applicable. This may be a Cookie or Header name for instance`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Route status. Routes in "ok" state are ready to operate`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of rules to match to trigger action`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Name of the field to match like "protocol" or "host" "/ipLoadbalancing/{serviceName}/route/availableRules" for a list of available rules`,
				},
				resource.Attribute{
					Name:        "match",
					Description: `Matching operator. Not all operators are available for all fields. See "availableRules"`,
				},
				resource.Attribute{
					Name:        "pattern",
					Description: `Value to match against this match. Interpretation if this field depends on the match and field`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Id of your rule`,
				},
				resource.Attribute{
					Name:        "sub_field",
					Description: `Name of sub-field, if applicable. This may be a Cookie or Header name for instance`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_tcp_route_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `Name of sub-field, if applicable. This may be a Cookie or Header name for instance ## Attributes Reference No additional attribute is exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_iploadbalancing_vrack_network",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `This attribute is there for documentation purpose only and isnt passed to the OVHcloud API as it may conflicts with http/tcp farms ` + "`" + `vrack_network_id` + "`" + ` attribute`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template_partition_scheme",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template_partition_scheme_hardware_raid",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_installation_template_partition_scheme_partition",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_me_ipxe_script",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `For documentation purpose only. This attribute is not passed to the OVHcloud API as it cannot be retrieved back. Instead a fake description ('$name auto description') is passed at creation time.`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "ovh_dedicated_server_networking",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Operation description.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `status of the networking configuration (should be ` + "`" + `active` + "`" + `). ## Import A dedicated server networking configuration can be imported using the ` + "`" + `service_name` + "`" + `. ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_dedicated_server_networking.server service_name ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM`,
				},
				resource.Attribute{
					Name:        "plan_option",
					Description: `(Optional) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM ## Attributes Reference Id is set to the order Id. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Zone name`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "dnssec_supported",
					Description: `Is DNSSEC supported by this zone`,
				},
				resource.Attribute{
					Name:        "has_dns_anycast",
					Description: `hasDnsAnycast flag of the DNS zone`,
				},
				resource.Attribute{
					Name:        "last_update",
					Description: `Last update date of the DNS zone`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Zone name`,
				},
				resource.Attribute{
					Name:        "name_servers",
					Description: `Name servers that host the DNS zone`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone_record",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
					Description: `(Optional) The TTL of the record, it shall be >= to 60. ## Attributes Reference The following attributes are exported:`,
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
					Description: `The TTL of the record ## Import OVHcloud domain zone record can be imported using the ` + "`" + `id` + "`" + `, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the ` + "`" + `zone` + "`" + `, separated by "." E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_domain_zone_record.test id.zone ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The TTL of the record ## Import OVHcloud domain zone record can be imported using the ` + "`" + `id` + "`" + `, which can be retrieved by using [OVH API portal](https://api.ovh.com/console/#/domain/zone/%7BzoneName%7D/record~GET), and the ` + "`" + `zone` + "`" + `, separated by "." E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_domain_zone_record.test id.zone ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_domain_zone_redirection",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "ovh_vrack",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `yourvrackdescription`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `yourvrackname`,
				},
				resource.Attribute{
					Name:        "ovh_subsidiary",
					Description: `(Required) OVHcloud Subsidiary`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM`,
				},
				resource.Attribute{
					Name:        "plan_option",
					Description: `(Optional) Product Plan to order`,
				},
				resource.Attribute{
					Name:        "duration",
					Description: `(Required) duration`,
				},
				resource.Attribute{
					Name:        "plan_code",
					Description: `(Required) Plan code`,
				},
				resource.Attribute{
					Name:        "pricing_mode",
					Description: `(Required) Pricing model identifier`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `Catalog name`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) Representation of a configuration item for personalizing product`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) Identifier of the resource`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Path to the resource in API.OVH.COM ## Attributes Reference Id is set to the order Id. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your vrack`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "order",
					Description: `Details about an Order`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `date`,
				},
				resource.Attribute{
					Name:        "order_id",
					Description: `order id`,
				},
				resource.Attribute{
					Name:        "expiration_date",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Information about a Bill entry`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description`,
				},
				resource.Attribute{
					Name:        "order_detail_id",
					Description: `order detail id`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `expiration date`,
				},
				resource.Attribute{
					Name:        "quantity",
					Description: `quantity`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `The internal name of your vrack`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_cloudproject",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name of the vrack. If omitted, the ` + "`" + `OVH_VRACK_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the public cloud project. If omitted, the ` + "`" + `OVH_CLOUD_PROJECT_SERVICE` + "`" + ` environment variable is used. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Import Attachment of a public cloud project and a VRack can be imported using the ` + "`" + `project_id` + "`" + `, the ` + "`" + `service_name` + "`" + ` (vRackID) and the ` + "`" + `attach_id` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_vrack_cloudproject.myattach ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/vrack_pn-12345678-cloudproject_ookie9mee8Shaeghaeleeju7Xeghohv6e-attach ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above. ## Import Attachment of a public cloud project and a VRack can be imported using the ` + "`" + `project_id` + "`" + `, the ` + "`" + `service_name` + "`" + ` (vRackID) and the ` + "`" + `attach_id` + "`" + `, separated by "/" E.g., ` + "`" + `` + "`" + `` + "`" + `bash $ terraform import ovh_vrack_cloudproject.myattach ookie9mee8Shaeghaeleeju7Xeghohv6e/pn-12345678/vrack_pn-12345678-cloudproject_ookie9mee8Shaeghaeleeju7Xeghohv6e-attach ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_dedicated_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name of the vrack. If omitted, the ` + "`" + `OVH_VRACK_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) The id of the dedicated server. ## Attributes Reference The following attributes are exported:`,
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
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the vrack. If omitted, the ` + "`" + `OVH_VRACK_SERVICE` + "`" + ` environment variable is used.`,
				},
				resource.Attribute{
					Name:        "interface_id",
					Description: `(Required) The id of dedicated server network interface. ## Attributes Reference The following attributes are exported:`,
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
			Type:             "ovh_vrack_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The internal name of your vrack`,
				},
				resource.Attribute{
					Name:        "block",
					Description: `(Required) Your IP block. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Your gateway`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Your IP block`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Where you want your block announced on the network`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "gateway",
					Description: `Your gateway`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Your IP block`,
				},
				resource.Attribute{
					Name:        "zone",
					Description: `Where you want your block announced on the network`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ovh_vrack_iploadbalancing",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The id of the vrack.`,
				},
				resource.Attribute{
					Name:        "ip_loadbalancing",
					Description: `(Required) The id of the IP Load Balancing. ## Attributes Reference The following attributes are exported:`,
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

		"ovh_cloud_project":                                           0,
		"ovh_cloud_project_containerregistry":                         1,
		"ovh_cloud_project_containerregistry_user":                    2,
		"ovh_cloud_project_database":                                  3,
		"ovh_cloud_project_database_database":                         4,
		"ovh_cloud_project_database_integration":                      5,
		"ovh_cloud_project_database_ip_restriction":                   6,
		"ovh_cloud_project_database_kafka_acl":                        7,
		"ovh_cloud_project_database_kafka_topic":                      8,
		"ovh_cloud_project_database_m3db_namespace":                   9,
		"ovh_cloud_project_database_m3db_user":                        10,
		"ovh_cloud_project_database_mongodb_user":                     11,
		"ovh_cloud_project_database_opensearch_pattern":               12,
		"ovh_cloud_project_database_opensearch_user":                  13,
		"ovh_cloud_project_database_postgresql_user":                  14,
		"ovh_cloud_project_database_redis_user":                       15,
		"ovh_cloud_project_database_user":                             16,
		"ovh_cloud_project_failover_ip_attach":                        17,
		"ovh_cloud_project_kube":                                      18,
		"ovh_cloud_project_kube_iprestrictions":                       19,
		"ovh_cloud_project_kube_nodepool":                             20,
		"ovh_cloud_project_kube_oidc":                                 21,
		"ovh_cloud_project_network_private":                           22,
		"ovh_cloud_project_network_private_subnet":                    23,
		"ovh_cloud_project_region_storage_presign":                    24,
		"ovh_cloud_project_user":                                      25,
		"ovh_cloud_project_user_s3_credential":                        26,
		"ovh_cloud_project_user_s3_policy":                            27,
		"ovh_cloud_project_workflow_backup":                           28,
		"ovh_dbaas_logs_cluster":                                      29,
		"ovh_dbaas_logs_graylog_output_stream":                        30,
		"ovh_dbaas_logs_input":                                        31,
		"ovh_dedicated_ceph_acl":                                      32,
		"ovh_dedicated_nasha_partition":                               33,
		"ovh_dedicated_nasha_partition_access":                        34,
		"ovh_dedicated_nasha_partition_snapshot":                      35,
		"ovh_dedicated_server_install_task":                           36,
		"ovh_dedicated_server_reboot_task":                            37,
		"ovh_dedicated_server_update":                                 38,
		"ovh_hosting_privatedatabase":                                 39,
		"ovh_hosting_privatedatabase_database":                        40,
		"ovh_hosting_privatedatabase_user":                            41,
		"ovh_hosting_privatedatabase_user_grant":                      42,
		"ovh_hosting_privatedatabase_whitelist":                       43,
		"ovh_ip_reverse":                                              44,
		"ovh_ip_service":                                              45,
		"ovh_iploadbalancing":                                         46,
		"ovh_iploadbalancing_http_farm":                               47,
		"ovh_iploadbalancing_http_farm_server":                        48,
		"ovh_iploadbalancing_http_frontend":                           49,
		"ovh_iploadbalancing_http_route":                              50,
		"ovh_iploadbalancing_http_route_rule":                         51,
		"ovh_iploadbalancing_refresh":                                 52,
		"ovh_iploadbalancing_tcp_farm":                                53,
		"ovh_iploadbalancing_tcp_farm_server":                         54,
		"ovh_iploadbalancing_tcp_frontend":                            55,
		"ovh_iploadbalancing_tcp_route":                               56,
		"ovh_iploadbalancing_tcp_route_rule":                          57,
		"ovh_iploadbalancing_vrack_network":                           58,
		"ovh_me_identity_user":                                        59,
		"ovh_me_installation_template":                                60,
		"ovh_me_installation_template_partition_scheme":               61,
		"ovh_me_installation_template_partition_scheme_hardware_raid": 62,
		"ovh_me_installation_template_partition_scheme_partition":     63,
		"ovh_me_ipxe_script":                                          64,
		"ovh_me_ssh_key":                                              65,
		"ovh_dedicated_server_networking":                             66,
		"ovh_domain_zone":                                             67,
		"ovh_domain_zone_record":                                      68,
		"ovh_domain_zone_redirection":                                 69,
		"ovh_vrack":                                                   70,
		"ovh_vrack_cloudproject":                                      71,
		"ovh_vrack_dedicated_server":                                  72,
		"ovh_vrack_dedicated_server_interface":                        73,
		"ovh_vrack_ip":                                                74,
		"ovh_vrack_iploadbalancing":                                   75,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
