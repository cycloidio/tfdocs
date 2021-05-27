package ec

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_deployment",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about an Elastic Cloud deployment.`,
			Description: `

Use this data source to retrieve information about an existing Elastic Cloud deployment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of an existing Elastic Cloud deployment. ## Attributes Reference ~>`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Deployment alias.`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `Overall health status of the deployment.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the deployment.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region where the deployment can be found.`,
				},
				resource.Attribute{
					Name:        "deployment_template_id",
					Description: `ID of the deployment template used to create the deployment.`,
				},
				resource.Attribute{
					Name:        "traffic_filter",
					Description: `Traffic filter block, which contains a list of traffic filter rule identifiers.`,
				},
				resource.Attribute{
					Name:        "observability.#.deployment_id",
					Description: `Destination deployment ID for the shipped logs and monitoring metrics.`,
				},
				resource.Attribute{
					Name:        "observability.#.ref_id",
					Description: `Elasticsearch resource kind ref_id of the destination deployment.`,
				},
				resource.Attribute{
					Name:        "observability.#.logs",
					Description: `Defines whether logs are enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "observability.#.metrics",
					Description: `Defines whether metrics are enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `Instance configuration of the Elasticsearch resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.autoscale",
					Description: `Whether or not Elasticsearch autoscaling is enabled.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.cloud_id",
					Description: `The encoded Elasticsearch credentials to use in Beats or Logstash. See [Configure Beats and Logstash with Cloud ID](https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html) for more information.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology",
					Description: `Topology element definition.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_roles",
					Description: `Defines the list of Elasticsearch node roles assigned to the topology element (>=7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_data",
					Description: `Defines whether this node can hold data (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_master",
					Description: `Defines whether this node can be elected master (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ingest",
					Description: `Defines whether this node can run an ingest pipeline (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ml",
					Description: `Defines whether this node can run ML jobs (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.max_size",
					Description: `The maximum size for the scale up policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.max_size_resource",
					Description: `The maximum size resource for the scale up policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.min_size",
					Description: `The minimum size for the scale down policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.min_size_resource",
					Description: `The minimum size for the scale down policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.policy_override_json",
					Description: `The advanced policy overrides for the autoscaling policy.`,
				},
				resource.Attribute{
					Name:        "kibana",
					Description: `Instance configuration of the Kibana type.`,
				},
				resource.Attribute{
					Name:        "kibana.#.elasticsearch_cluster_ref_id",
					Description: `The user-specified ID of the Elasticsearch cluster to which this resource kind will link.`,
				},
				resource.Attribute{
					Name:        "kibana.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "kibana.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "kibana.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "kibana.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "kibana.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "kibana.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "kibana.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology",
					Description: `Node topology element definition.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "apm",
					Description: `Instance configuration of the APM type.`,
				},
				resource.Attribute{
					Name:        "apm.#.elasticsearch_cluster_ref_id",
					Description: `The user-specified ID of the Elasticsearch cluster to which this resource kind will link.`,
				},
				resource.Attribute{
					Name:        "apm.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "apm.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "apm.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "apm.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "apm.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "apm.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "apm.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology",
					Description: `Node topology element definition.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "enterprise_search",
					Description: `Instance configuration of the Enterprise Search type.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.elasticsearch_cluster_ref_id",
					Description: `The user-specified ID of the Elasticsearch cluster to which this resource kind will link.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology",
					Description: `Node topology element definition.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_appserver",
					Description: `Defines whether this instance should run as application/API server.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_connector",
					Description: `Defines whether this instance should run as connector.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_worker",
					Description: `Defines whether this instance should run as background worker.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "alias",
					Description: `Deployment alias.`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `Overall health status of the deployment.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the deployment.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Region where the deployment can be found.`,
				},
				resource.Attribute{
					Name:        "deployment_template_id",
					Description: `ID of the deployment template used to create the deployment.`,
				},
				resource.Attribute{
					Name:        "traffic_filter",
					Description: `Traffic filter block, which contains a list of traffic filter rule identifiers.`,
				},
				resource.Attribute{
					Name:        "observability.#.deployment_id",
					Description: `Destination deployment ID for the shipped logs and monitoring metrics.`,
				},
				resource.Attribute{
					Name:        "observability.#.ref_id",
					Description: `Elasticsearch resource kind ref_id of the destination deployment.`,
				},
				resource.Attribute{
					Name:        "observability.#.logs",
					Description: `Defines whether logs are enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "observability.#.metrics",
					Description: `Defines whether metrics are enabled or disabled.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `Instance configuration of the Elasticsearch resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.autoscale",
					Description: `Whether or not Elasticsearch autoscaling is enabled.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.cloud_id",
					Description: `The encoded Elasticsearch credentials to use in Beats or Logstash. See [Configure Beats and Logstash with Cloud ID](https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html) for more information.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology",
					Description: `Topology element definition.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_roles",
					Description: `Defines the list of Elasticsearch node roles assigned to the topology element (>=7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_data",
					Description: `Defines whether this node can hold data (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_master",
					Description: `Defines whether this node can be elected master (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ingest",
					Description: `Defines whether this node can run an ingest pipeline (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ml",
					Description: `Defines whether this node can run ML jobs (<7.10.0).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.max_size",
					Description: `The maximum size for the scale up policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.max_size_resource",
					Description: `The maximum size resource for the scale up policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.min_size",
					Description: `The minimum size for the scale down policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.min_size_resource",
					Description: `The minimum size for the scale down policy.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.policy_override_json",
					Description: `The advanced policy overrides for the autoscaling policy.`,
				},
				resource.Attribute{
					Name:        "kibana",
					Description: `Instance configuration of the Kibana type.`,
				},
				resource.Attribute{
					Name:        "kibana.#.elasticsearch_cluster_ref_id",
					Description: `The user-specified ID of the Elasticsearch cluster to which this resource kind will link.`,
				},
				resource.Attribute{
					Name:        "kibana.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "kibana.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "kibana.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "kibana.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "kibana.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "kibana.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "kibana.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology",
					Description: `Node topology element definition.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "kibana.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "apm",
					Description: `Instance configuration of the APM type.`,
				},
				resource.Attribute{
					Name:        "apm.#.elasticsearch_cluster_ref_id",
					Description: `The user-specified ID of the Elasticsearch cluster to which this resource kind will link.`,
				},
				resource.Attribute{
					Name:        "apm.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "apm.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "apm.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "apm.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "apm.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "apm.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "apm.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology",
					Description: `Node topology element definition.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "apm.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "enterprise_search",
					Description: `Instance configuration of the Enterprise Search type.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.elasticsearch_cluster_ref_id",
					Description: `The user-specified ID of the Elasticsearch cluster to which this resource kind will link.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.healthy",
					Description: `Resource kind health status.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.http_endpoint",
					Description: `HTTP endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.https_endpoint",
					Description: `HTTPS endpoint for the resource kind.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.ref_id",
					Description: `User specified ref_id for the resource kind.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.resource_id",
					Description: `The resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.status",
					Description: `Resource kind status (for example, "started", "stopped", etc).`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology",
					Description: `Node topology element definition.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.instance_configuration_id",
					Description: `Controls the allocation of this topology element as well as allowed sizes and node_types. It needs to match the ID of an existing instance configuration.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.size",
					Description: `Amount of memory (RAM) per topology element in the "<size in GB>g" notation.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.zone_count",
					Description: `Number of zones in which nodes will be placed.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_appserver",
					Description: `Defines whether this instance should run as application/API server.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_connector",
					Description: `Defines whether this instance should run as connector.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_worker",
					Description: `Defines whether this instance should run as background worker.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_deployments",
			Category:         "Data Sources",
			ShortDescription: `Returns a list of deployments that match the specified query.`,
			Description: `

Use this data source to retrieve a list of IDs for the deployment and resource kinds, based on the specified query.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_prefix",
					Description: `Prefix that one or several deployment names have in common.`,
				},
				resource.Attribute{
					Name:        "deployment_template_id",
					Description: `ID of the deployment template used to create the deployment.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The maximum number of deployments to return. Defaults to ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Key value map of arbitrary string tags for the deployment.`,
				},
				resource.Attribute{
					Name:        "healthy",
					Description: `Overall health status of the deployment.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `Filter by Elasticsearch resource kind status or configuration.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.status",
					Description: `Resource kind status (Available statuses are: initializing, stopping, stopped, rebooting, restarting, reconfiguring, and started).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.healthy",
					Description: `Overall health status of the Elasticsearch instances.`,
				},
				resource.Attribute{
					Name:        "kibana",
					Description: `Filter by Kibana resource kind status or configuration.`,
				},
				resource.Attribute{
					Name:        "kibana.#.status",
					Description: `Resource kind status (Available statuses are: initializing, stopping, stopped, rebooting, restarting, reconfiguring, and started).`,
				},
				resource.Attribute{
					Name:        "kibana.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "kibana.#.healthy",
					Description: `Overall health status of the Kibana instances.`,
				},
				resource.Attribute{
					Name:        "apm",
					Description: `Filter by APM resource kind status or configuration.`,
				},
				resource.Attribute{
					Name:        "apm.#.status",
					Description: `Resource kind status (Available statuses are: initializing, stopping, stopped, rebooting, restarting, reconfiguring, and started).`,
				},
				resource.Attribute{
					Name:        "apm.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "apm.#.healthy",
					Description: `Overall health status of the APM instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search",
					Description: `Filter by Enterprise Search resource kind status or configuration.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.status",
					Description: `Resource kind status (Available statuses are: initializing, stopping, stopped, rebooting, restarting, reconfiguring, and started).`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.version",
					Description: `Elastic stack version.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.healthy",
					Description: `Overall health status of the Enterprise Search instances. ## Attributes Reference ~>`,
				},
				resource.Attribute{
					Name:        "deployments",
					Description: `List of deployments which match the specified query.`,
				},
				resource.Attribute{
					Name:        "deployments.#.deployment_id",
					Description: `The deployment unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.elasticsearch_resource_id",
					Description: `The Elasticsearch resource unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.kibana_resource_id",
					Description: `The Kibana resource unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.apm_resource_id",
					Description: `The APM resource unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.enterprise_search_resource_id",
					Description: `The Enterprise Search resource unique ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "deployments",
					Description: `List of deployments which match the specified query.`,
				},
				resource.Attribute{
					Name:        "deployments.#.deployment_id",
					Description: `The deployment unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.elasticsearch_resource_id",
					Description: `The Elasticsearch resource unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.kibana_resource_id",
					Description: `The Kibana resource unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.apm_resource_id",
					Description: `The APM resource unique ID.`,
				},
				resource.Attribute{
					Name:        "deployments.#.enterprise_search_resource_id",
					Description: `The Enterprise Search resource unique ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_stack",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about an Elastic Cloud stack.`,
			Description: `

Use this data source to retrieve information about an existing Elastic Cloud stack.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `The stack version.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `To have this version accessible/not accessible by the calling user. This is only relevant for Elasticsearch Service (ESS), not for ECE.`,
				},
				resource.Attribute{
					Name:        "min_upgradable_from",
					Description: `The minimum stack version recommended.`,
				},
				resource.Attribute{
					Name:        "upgradable_to",
					Description: `The stack version you can upgrade to.`,
				},
				resource.Attribute{
					Name:        "allowlisted",
					Description: `To include/not include this version in the ` + "`" + `allowlist` + "`" + `. This is only relevant for Elasticsearch Service (ESS), not for ECE.`,
				},
				resource.Attribute{
					Name:        "apm",
					Description: `Information for APM workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "apm.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "apm.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "apm.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "apm.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "apm.#.docker_image",
					Description: `Docker image to use for the APM instance.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `Information for Elasticsearch workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.default_plugins",
					Description: `List of default plugins which are included in all Elasticsearch cluster instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.docker_image",
					Description: `Docker image to use for the Elasticsearch cluster instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.plugins",
					Description: `List of available plugins to be specified by users in Elasticsearch cluster instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search",
					Description: `Information for Enterprise Search workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.docker_image",
					Description: `Docker image to use for the Enterprise Search instance.`,
				},
				resource.Attribute{
					Name:        "kibana",
					Description: `Information for Kibana workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "kibana.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "kibana.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "kibana.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "kibana.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "kibana.#.docker_image",
					Description: `Docker image to use for the Kibana instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `The stack version.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `To have this version accessible/not accessible by the calling user. This is only relevant for Elasticsearch Service (ESS), not for ECE.`,
				},
				resource.Attribute{
					Name:        "min_upgradable_from",
					Description: `The minimum stack version recommended.`,
				},
				resource.Attribute{
					Name:        "upgradable_to",
					Description: `The stack version you can upgrade to.`,
				},
				resource.Attribute{
					Name:        "allowlisted",
					Description: `To include/not include this version in the ` + "`" + `allowlist` + "`" + `. This is only relevant for Elasticsearch Service (ESS), not for ECE.`,
				},
				resource.Attribute{
					Name:        "apm",
					Description: `Information for APM workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "apm.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "apm.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "apm.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "apm.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "apm.#.docker_image",
					Description: `Docker image to use for the APM instance.`,
				},
				resource.Attribute{
					Name:        "elasticsearch",
					Description: `Information for Elasticsearch workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.default_plugins",
					Description: `List of default plugins which are included in all Elasticsearch cluster instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.docker_image",
					Description: `Docker image to use for the Elasticsearch cluster instances.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.plugins",
					Description: `List of available plugins to be specified by users in Elasticsearch cluster instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search",
					Description: `Information for Enterprise Search workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.docker_image",
					Description: `Docker image to use for the Enterprise Search instance.`,
				},
				resource.Attribute{
					Name:        "kibana",
					Description: `Information for Kibana workloads on this stack version.`,
				},
				resource.Attribute{
					Name:        "kibana.#.denylist",
					Description: `List of configuration options that cannot be overridden by user settings.`,
				},
				resource.Attribute{
					Name:        "kibana.#.capacity_constraints_min",
					Description: `Minimum size of the instances.`,
				},
				resource.Attribute{
					Name:        "kibana.#.capacity_constraints_max",
					Description: `Maximum size of the instances.`,
				},
				resource.Attribute{
					Name:        "kibana.#.compatible_node_types",
					Description: `List of node types compatible with this one.`,
				},
				resource.Attribute{
					Name:        "kibana.#.docker_image",
					Description: `Docker image to use for the Kibana instance.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ec_ec_deployment":  0,
		"ec_ec_deployments": 1,
		"ec_ec_stack":       2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
