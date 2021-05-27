package ec

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_deployment",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud deployment resource, which allows deployments to be created, updated, and deleted.`,
			Description: `

Provides an Elastic Cloud deployment resource, which allows deployments to be created, updated, and deleted.

~> **Note on Elastic Stack versions** Using a version prior to ` + "`" + `6.6.0` + "`" + ` is not supported.

~> **Note on traffic filters** If you use ` + "`" + `traffic_filter` + "`" + ` on an ` + "`" + `ec_deployment` + "`" + `, Terraform will manage the full set of traffic rules for the deployment, and treat additional traffic filters as drift. For this reason, ` + "`" + `traffic_filter` + "`" + ` cannot be mixed with the ` + "`" + `ec_deployment_traffic_filter_association` + "`" + ` resource for a given deployment.

-> **Note on regions and deployment templates** Before you start, you might want to read about [Elastic Cloud deployments](https://www.elastic.co/guide/en/cloud/current/ec-create-deployment.html) and check the [full list](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html) of regions and deployment templates available in Elasticsearch Service (ESS).

-> **Note on Elasticsearch topology IDs** Since the addition of data tiers, each Elasticsearch topology block requires the ` + "`" + `"id"` + "`" + ` field to be set. The accepted values are set in the deployment template that you have chosen, but values are closely related to the Elasticsearch data tiers. [Learn more abut Elasticsearch data tiers](https://www.elastic.co/guide/en/elasticsearch/reference/current/data-tiers.html). For a complete list of all the supported values, refer to the deployment template definition used by your deployment.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Elasticsearch Service (ESS) region where to create the deployment. For Elastic Cloud Enterprise (ECE) installations, set ` + "`" + `"ece-region"` + "`" + `. -> If you change the ` + "`" + `region` + "`" + `, the resource will be destroyed and re-created.`,
				},
				resource.Attribute{
					Name:        "deployment_template_id",
					Description: `(Required) Deployment template identifier to create the deployment from. See the [full list](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html) of regions and deployment templates available in ESS.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Elastic Stack version to use for all the deployment resources. -> Read the [ESS stack version policy](https://www.elastic.co/guide/en/cloud/current/ec-version-policy.html#ec-version-policy-available) to understand which versions are available.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the deployment.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) Deployment alias, affects the format of the resource URLs.`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `(Optional) Request ID to set when you create the deployment. Use it only when previous attempts return an error and ` + "`" + `request_id` + "`" + ` is returned as part of the error.`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Can be set multiple times to compose complex topologies.`,
				},
				resource.Attribute{
					Name:        "ref_id",
					Description: `(Optional) Can be set on the Elasticsearch resource. The default value ` + "`" + `main-elasticsearch` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) Unique topology identifier. It generally refers to an Elasticsearch data tier, such as ` + "`" + `hot_content` + "`" + `, ` + "`" + `warm` + "`" + `, ` + "`" + `cold` + "`" + `, ` + "`" + `coordinating` + "`" + `, ` + "`" + `frozen` + "`" + `, ` + "`" + `ml` + "`" + ` or ` + "`" + `master` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Amount in Gigabytes per topology element in the ` + "`" + `"<size in GB>g"` + "`" + ` notation. When omitted, it defaults to the deployment template value.`,
				},
				resource.Attribute{
					Name:        "size_resource",
					Description: `(Optional) Type of resource to which the size is assigned. Defaults to ` + "`" + `"memory"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_count",
					Description: `(Optional) Number of zones the instance type of the Elasticsearch cluster will span. This is used to set or unset HA on an Elasticsearch node type. When omitted, it defaults to the deployment template value.`,
				},
				resource.Attribute{
					Name:        "node_type_data",
					Description: `(Optional) The node type for the Elasticsearch cluster (data node).`,
				},
				resource.Attribute{
					Name:        "node_type_master",
					Description: `(Optional) The node type for the Elasticsearch cluster (master node).`,
				},
				resource.Attribute{
					Name:        "node_type_ingest",
					Description: `(Optional) The node type for the Elasticsearch cluster (ingest node).`,
				},
				resource.Attribute{
					Name:        "node_type_ml",
					Description: `(Optional) The node type for the Elasticsearch cluster (machine learning node).`,
				},
				resource.Attribute{
					Name:        "autoscaling",
					Description: `(Optional) Autoscaling policy defining the maximum and / or minimum total size for this topology element. For more information refer to the ` + "`" + `autoscaling` + "`" + ` block. ~>`,
				},
				resource.Attribute{
					Name:        "min_size",
					Description: `(Optional) Defines the minimum size the deployment will scale down to. When set, scale down will be enabled, please note that not all the tiers support this option.`,
				},
				resource.Attribute{
					Name:        "min_size_resource",
					Description: `(Optional) Defines the resource type the scale down will use (Defaults to ` + "`" + `"memory"` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "max_size",
					Description: `(Optional) Defines the maximum size the deployment will scale up to. When set, scaling up will be enabled. All tiers should support this option.`,
				},
				resource.Attribute{
					Name:        "max_size_resource",
					Description: `(Optional) Defines the resource type the scale up will use (Defaults to ` + "`" + `"memory"` + "`" + `). -> Note that none of these settings will take effect unless ` + "`" + `elasticsearch.autoscale` + "`" + ` is set to ` + "`" + `"true"` + "`" + `. Please refer to the [Deployment Autoscaling](https://www.elastic.co/guide/en/cloud/current/ec-autoscaling.html) documentation for an updated list of the Elasticsearch tiers supporting scale up and scale down. ##### Config The optional ` + "`" + `elasticsearch.config` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "plugins",
					Description: `(Optional) List of Elasticsearch supported plugins. Check the Stack Pack version to see which plugins are supported for each version. This is currently only available from the UI and [ecctl](https://www.elastic.co/guide/en/ecctl/master/ecctl_stack_list.html).`,
				},
				resource.Attribute{
					Name:        "user_settings_json",
					Description: `(Optional) JSON-formatted user level ` + "`" + `elasticsearch.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_json",
					Description: `(Optional) JSON-formatted admin (ECE) level ` + "`" + `elasticsearch.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_yaml",
					Description: `(Optional) YAML-formatted user level ` + "`" + `elasticsearch.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_yaml",
					Description: `(Optional) YAML-formatted admin (ECE) level ` + "`" + `elasticsearch.yml` + "`" + ` setting overrides. ##### Remote Cluster The optional ` + "`" + `elasticsearch.remote_cluster` + "`" + ` block can be set multiple times. It represents one or multiple remote clusters to which the local Elasticsearch cluster connects for Cross Cluster Search and supports the following settings:`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Can be set multiple times to compose complex topologies.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_cluster_ref_id",
					Description: `(Optional) This field references the ` + "`" + `ref_id` + "`" + ` of the deployment Elasticsearch cluster. The default value ` + "`" + `main-elasticsearch` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "ref_id",
					Description: `(Optional) Can be set on the Kibana resource. The default value ` + "`" + `main-kibana` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "instance_configuration_id",
					Description: `(Optional) Default instance configuration of the deployment template. No need to change this value since Kibana has only one _instance type_.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Amount of memory (RAM) per topology element in the "<size in GB>g" notation. When omitted, it defaults to the deployment template value.`,
				},
				resource.Attribute{
					Name:        "size_resource",
					Description: `(Optional) Type of resource to which the size is assigned. Defaults to ` + "`" + `"memory"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_count",
					Description: `(Optional) Number of zones that the Kibana deployment will span. This is used to set HA. When omitted, it defaults to the deployment template value. ##### Config The optional ` + "`" + `kibana.config` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "user_settings_json",
					Description: `(Optional) JSON-formatted user level ` + "`" + `kibana.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_json",
					Description: `(Optional) JSON-formatted admin (ECE) level ` + "`" + `kibana.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_yaml",
					Description: `(Optional) YAML-formatted user level ` + "`" + `kibana.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_yaml",
					Description: `(Optional) YAML-formatted admin (ECE) level ` + "`" + `kibana.yml` + "`" + ` setting overrides. #### APM The optional ` + "`" + `apm` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Can be set multiple times to compose complex topologies.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_cluster_ref_id",
					Description: `(Optional) This field references the ` + "`" + `ref_id` + "`" + ` of the deployment Elasticsearch cluster. The default value ` + "`" + `main-elasticsearch` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "ref_id",
					Description: `(Optional) Can be set on the APM resource. The default value ` + "`" + `main-apm` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "instance_configuration_id",
					Description: `(Optional) Default instance configuration of the deployment template. No need to change this value since APM has only one _instance type_.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Amount of memory (RAM) per topology element in the "<size in GB>g" notation. When omitted, it defaults to the deployment template value.`,
				},
				resource.Attribute{
					Name:        "size_resource",
					Description: `(Optional) Type of resource to which the size is assigned. Defaults to ` + "`" + `"memory"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_count",
					Description: `(Optional) Number of zones that the APM deployment will span. This is used to set HA. When omitted, it defaults to the deployment template value. ##### Config The optional ` + "`" + `apm.config` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "debug_enabled",
					Description: `(Optional) Enable debug mode for APM servers. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_settings_json",
					Description: `(Optional) JSON-formatted user level ` + "`" + `apm.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_json",
					Description: `(Optional) JSON-formatted admin (ECE) level ` + "`" + `apm.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_yaml",
					Description: `(Optional) YAML-formatted user level ` + "`" + `apm.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_yaml",
					Description: `(Optional) YAML-formatted admin (ECE) level ` + "`" + `apm.yml` + "`" + ` setting overrides. #### Enterprise Search The optional ` + "`" + `enterprise_search` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "topology",
					Description: `(Optional) Can be set multiple times to compose complex topologies.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_cluster_ref_id",
					Description: `(Optional) This field references the ` + "`" + `ref_id` + "`" + ` of the deployment Elasticsearch cluster. The default value ` + "`" + `main-elasticsearch` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "ref_id",
					Description: `(Optional) Can be set on the Enterprise Search resource. The default value ` + "`" + `main-enterprise_search` + "`" + ` is recommended.`,
				},
				resource.Attribute{
					Name:        "instance_configuration_id",
					Description: `(Optional) Default instance configuration of the deployment template. To change it, use the [full list](https://www.elastic.co/guide/en/cloud/current/ec-regions-templates-instances.html) of regions and deployment templates available in ESS.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Amount of memory (RAM) per ` + "`" + `topology` + "`" + ` element in the "<size in GB>g" notation. When omitted, it defaults to the deployment template value.`,
				},
				resource.Attribute{
					Name:        "size_resource",
					Description: `(Optional) Type of resource to which the size is assigned. Defaults to ` + "`" + `"memory"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "zone_count",
					Description: `(Optional) Number of zones that the Enterprise Search deployment will span. This is used to set HA. When omitted, it defaults to the deployment template value. ##### Config The optional ` + "`" + `enterprise_search.config` + "`" + ` block supports the following arguments:`,
				},
				resource.Attribute{
					Name:        "user_settings_json",
					Description: `(Optional) JSON-formatted user level ` + "`" + `enterprise_search.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_json",
					Description: `(Optional) JSON-formatted admin (ECE) level ` + "`" + `enterprise_search.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_yaml",
					Description: `(Optional) YAML-formatted user level ` + "`" + `enterprise_search.yml` + "`" + ` setting overrides.`,
				},
				resource.Attribute{
					Name:        "user_settings_override_yaml",
					Description: `(Optional) YAML-formatted admin (ECE) level ` + "`" + `enterprise_search.yml` + "`" + ` setting overrides. ### Timeouts`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Deployment identifier.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_username",
					Description: `Auto-generated Elasticsearch username.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_password",
					Description: `Auto-generated Elasticsearch password.`,
				},
				resource.Attribute{
					Name:        "apm_secret_token",
					Description: `Auto-generated APM secret_token, empty unless an ` + "`" + `apm` + "`" + ` resource is specified.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.resource_id",
					Description: `Elasticsearch resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.region",
					Description: `Elasticsearch region.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.cloud_id",
					Description: `Encoded Elasticsearch credentials to use in Beats or Logstash. For more information, see [Configure Beats and Logstash with Cloud ID](https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.http_endpoint",
					Description: `Elasticsearch resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.https_endpoint",
					Description: `Elasticsearch resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.instance_configuration_id",
					Description: `instance configuration of the deployment topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_data",
					Description: `Node type (data) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_master",
					Description: `Node type (master) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ingest",
					Description: `Node type (ingest) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ml",
					Description: `Node type (machine learning) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_roles",
					Description: `List of roles for the topology element. They are inferred from the deployment template.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.policy_override_json",
					Description: `Computed policy overrides set directly via the API or other clients.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.snapshot_source.#.source_elasticsearch_cluster_id",
					Description: `ID of the Elasticsearch cluster that will be used as the source of the snapshot.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.snapshot_source.#.snapshot_name",
					Description: `Name of the snapshot to restore.`,
				},
				resource.Attribute{
					Name:        "kibana.#.resource_id",
					Description: `Kibana resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "kibana.#.region",
					Description: `Kibana region.`,
				},
				resource.Attribute{
					Name:        "kibana.#.http_endpoint",
					Description: `Kibana resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "kibana.#.https_endpoint",
					Description: `Kibana resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "apm.#.resource_id",
					Description: `APM resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "apm.#.region",
					Description: `APM region.`,
				},
				resource.Attribute{
					Name:        "apm.#.http_endpoint",
					Description: `APM resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "apm.#.https_endpoint",
					Description: `APM resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.resource_id",
					Description: `Enterprise Search resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.region",
					Description: `Enterprise Search region.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.http_endpoint",
					Description: `Enterprise Search resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.https_endpoint",
					Description: `Enterprise Search resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_appserver",
					Description: `Node type (Appserver) for the Enterprise Search topology element.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_connector",
					Description: `Node type (Connector) for the Enterprise Search topology element.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_worker",
					Description: `Node type (worker) for the Enterprise Search topology element.`,
				},
				resource.Attribute{
					Name:        "observability.#.deployment_id",
					Description: `Destination deployment ID for the shipped logs and monitoring metrics.`,
				},
				resource.Attribute{
					Name:        "observability.#.ref_id",
					Description: `(Optional) Elasticsearch resource kind ref_id of the destination deployment.`,
				},
				resource.Attribute{
					Name:        "observability.#.logs",
					Description: `Enables or disables shipping logs. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "observability.#.metrics",
					Description: `Enables or disables shipping metrics. Defaults to true. ## Import ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Deployment identifier.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_username",
					Description: `Auto-generated Elasticsearch username.`,
				},
				resource.Attribute{
					Name:        "elasticsearch_password",
					Description: `Auto-generated Elasticsearch password.`,
				},
				resource.Attribute{
					Name:        "apm_secret_token",
					Description: `Auto-generated APM secret_token, empty unless an ` + "`" + `apm` + "`" + ` resource is specified.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.resource_id",
					Description: `Elasticsearch resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.region",
					Description: `Elasticsearch region.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.cloud_id",
					Description: `Encoded Elasticsearch credentials to use in Beats or Logstash. For more information, see [Configure Beats and Logstash with Cloud ID](https://www.elastic.co/guide/en/cloud/current/ec-cloud-id.html).`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.http_endpoint",
					Description: `Elasticsearch resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.https_endpoint",
					Description: `Elasticsearch resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.instance_configuration_id",
					Description: `instance configuration of the deployment topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_data",
					Description: `Node type (data) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_master",
					Description: `Node type (master) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ingest",
					Description: `Node type (ingest) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_type_ml",
					Description: `Node type (machine learning) for the Elasticsearch topology element.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.node_roles",
					Description: `List of roles for the topology element. They are inferred from the deployment template.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.topology.#.autoscaling.#.policy_override_json",
					Description: `Computed policy overrides set directly via the API or other clients.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.snapshot_source.#.source_elasticsearch_cluster_id",
					Description: `ID of the Elasticsearch cluster that will be used as the source of the snapshot.`,
				},
				resource.Attribute{
					Name:        "elasticsearch.#.snapshot_source.#.snapshot_name",
					Description: `Name of the snapshot to restore.`,
				},
				resource.Attribute{
					Name:        "kibana.#.resource_id",
					Description: `Kibana resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "kibana.#.region",
					Description: `Kibana region.`,
				},
				resource.Attribute{
					Name:        "kibana.#.http_endpoint",
					Description: `Kibana resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "kibana.#.https_endpoint",
					Description: `Kibana resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "apm.#.resource_id",
					Description: `APM resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "apm.#.region",
					Description: `APM region.`,
				},
				resource.Attribute{
					Name:        "apm.#.http_endpoint",
					Description: `APM resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "apm.#.https_endpoint",
					Description: `APM resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.resource_id",
					Description: `Enterprise Search resource unique identifier.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.region",
					Description: `Enterprise Search region.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.http_endpoint",
					Description: `Enterprise Search resource HTTP endpoint.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.https_endpoint",
					Description: `Enterprise Search resource HTTPs endpoint.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_appserver",
					Description: `Node type (Appserver) for the Enterprise Search topology element.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_connector",
					Description: `Node type (Connector) for the Enterprise Search topology element.`,
				},
				resource.Attribute{
					Name:        "enterprise_search.#.topology.#.node_type_worker",
					Description: `Node type (worker) for the Enterprise Search topology element.`,
				},
				resource.Attribute{
					Name:        "observability.#.deployment_id",
					Description: `Destination deployment ID for the shipped logs and monitoring metrics.`,
				},
				resource.Attribute{
					Name:        "observability.#.ref_id",
					Description: `(Optional) Elasticsearch resource kind ref_id of the destination deployment.`,
				},
				resource.Attribute{
					Name:        "observability.#.logs",
					Description: `Enables or disables shipping logs. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "observability.#.metrics",
					Description: `Enables or disables shipping metrics. Defaults to true. ## Import ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_deployment_extension",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud extension resource, which allows extensions to be created, updated, and deleted.`,
			Description: `
Provides an Elastic Cloud extension resource, which allows extensions to be created, updated, and deleted.

Extensions allow users of Elastic Cloud to use custom plugins, scripts, or dictionaries to enhance the core functionality of Elasticsearch. Before you install an extension, be sure to check out the supported and official [Elasticsearch plugins](https://www.elastic.co/guide/en/elasticsearch/plugins/current/index.html) already available.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the extension.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the extension.`,
				},
				resource.Attribute{
					Name:        "extension_type",
					Description: `(Required) ` + "`" + `bundle` + "`" + ` or ` + "`" + `plugin` + "`" + ` allowed. A ` + "`" + `bundle` + "`" + ` will usually contain a dictionary or script, where a ` + "`" + `plugin` + "`" + ` is compiled from source.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Elastic stack version, a numeric version for plugins, e.g. 2.3.0 should be set. Major version e.g. 2.`,
				},
				resource.Attribute{
					Name:        "download_url",
					Description: `(Optional) The URL to download the extension archive.`,
				},
				resource.Attribute{
					Name:        "file_path",
					Description: `(Optional) File path of the extension uploaded.`,
				},
				resource.Attribute{
					Name:        "file_hash",
					Description: `(Optional) Hash value of the file. If it is changed, the file is reuploaded. ## Attributes Reference In addition to all the arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Extension identifier.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The extension URL to be used in the plan.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The datetime the extension was last modified.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The extension file size in bytes. ## Import You can import extensions using the ` + "`" + `id` + "`" + `, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ec_deployment_extension.name 320b7b540dfc967a7a649c18e2fce4ed ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Extension identifier.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The extension URL to be used in the plan.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `The datetime the extension was last modified.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The extension file size in bytes. ## Import You can import extensions using the ` + "`" + `id` + "`" + `, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ec_deployment_extension.name 320b7b540dfc967a7a649c18e2fce4ed ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_deployment_traffic_filter",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud traffic filter resource, which allows traffic filter rules to be created, updated, and deleted. Traffic filter rules are used to limit inbound traffic to deployment resources.`,
			Description: `

Provides an Elastic Cloud traffic filter resource, which allows traffic filter rules to be created, updated, and deleted. Traffic filter rules are used to limit inbound traffic to deployment resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the ruleset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of the ruleset. It can be ` + "`" + `"ip"` + "`" + ` or ` + "`" + `"vpce"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) Filter region, the ruleset can only be attached to deployments in the specific region.`,
				},
				resource.Attribute{
					Name:        "include_by_default",
					Description: `(Optional) To automatically include the ruleset in the new deployments. Defaults to ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ruleset. ### Rules The ` + "`" + `rule` + "`" + ` block supports the following configuration options:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Source type, ` + "`" + `"ip"` + "`" + ` or ` + "`" + `"vpce"` + "`" + `, from which the ruleset accepts traffic.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of this individual rule. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ruleset ID. For any ` + "`" + `rule` + "`" + ` an ` + "`" + `id` + "`" + ` is exported. For example: ` + "`" + `ec_deployment_traffic_filter.default.rule.0.id` + "`" + `. ## Import You can import traffic filters using the ` + "`" + `id` + "`" + `, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ec_deployment_traffic_filter.name 320b7b540dfc967a7a649c18e2fce4ed ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ruleset ID. For any ` + "`" + `rule` + "`" + ` an ` + "`" + `id` + "`" + ` is exported. For example: ` + "`" + `ec_deployment_traffic_filter.default.rule.0.id` + "`" + `. ## Import You can import traffic filters using the ` + "`" + `id` + "`" + `, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import ec_deployment_traffic_filter.name 320b7b540dfc967a7a649c18e2fce4ed ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ec_ec_deployment_traffic_filter_association",
			Category:         "Resources",
			ShortDescription: `Provides an Elastic Cloud traffic filter association resource, which allows traffic filter rules to be associated to an Elastic Cloud deployment outside of the control of Terraform. Associations can be created and deleted.`,
			Description: `

Provides an Elastic Cloud traffic filter association resource, which allows traffic filter rules to be associated to an Elastic Cloud deployment outside of the control of Terraform. Associations can be created and deleted.

~> **Note on traffic filters** If you use ` + "`" + `traffic_filter` + "`" + ` on an ` + "`" + `ec_deployment` + "`" + `, Terraform will manage the full set of traffic rules for the deployment, and treat additional traffic filters as drift. For this reason, ` + "`" + `traffic_filter` + "`" + ` cannot be mixed with the ` + "`" + `ec_deployment_traffic_filter_association` + "`" + ` resource for a given deployment.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "traffic_filter_id",
					Description: `(Required) Traffic filter ID of the rule to use for the attachment.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Required) Deployment ID of the deployment to which the traffic filter rule is attached. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An autogenerated ID. ## Import Import is not supported on this resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An autogenerated ID. ## Import Import is not supported on this resource.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ec_ec_deployment":                            0,
		"ec_ec_deployment_extension":                  1,
		"ec_ec_deployment_traffic_filter":             2,
		"ec_ec_deployment_traffic_filter_association": 3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
