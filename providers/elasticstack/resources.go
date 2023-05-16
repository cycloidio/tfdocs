package elasticstack

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_cluster_settings",
			Category:         "Cluster",
			ShortDescription: `Updates cluster-wide settings.`,
			Description: `

Updates cluster-wide settings. If the Elasticsearch security features are enabled, you must have the manage cluster privilege to use this API. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/cluster-update-settings.html

`,
			Keywords: []string{
				"cluster",
				"elasticsearch",
				"settings",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_component_template",
			Category:         "Index",
			ShortDescription: `Creates or updates a component template.`,
			Description: `

Creates or updates a component template. Component templates are building blocks for constructing index templates that specify index mappings, settings, and aliases. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html

`,
			Keywords: []string{
				"index",
				"elasticsearch",
				"component",
				"template",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_data_stream",
			Category:         "Index",
			ShortDescription: `Manages Elasticsearch Data Streams`,
			Description: `

Manages data streams. This resource can create, delete and show the information about the created data stream. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/data-stream-apis.html

`,
			Keywords: []string{
				"index",
				"elasticsearch",
				"data",
				"stream",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_enrich_policy",
			Category:         "Enrich",
			ShortDescription: `Managing Elasticsearch enrich policies, see: https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-apis.html`,
			Description: `

Creates or updates enrich policies, see: https://www.elastic.co/guide/en/elasticsearch/reference/current/enrich-apis.html

`,
			Keywords: []string{
				"enrich",
				"elasticsearch",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_index",
			Category:         "Index",
			ShortDescription: `Creates or updates an index.`,
			Description: `

Creates or updates an index. This resource can define settings, mappings and aliases. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html

`,
			Keywords: []string{
				"index",
				"elasticsearch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_index_lifecycle",
			Category:         "Index",
			ShortDescription: `Creates or updates lifecycle policy.`,
			Description: `

Creates or updates lifecycle policy. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-put-lifecycle.html and https://www.elastic.co/guide/en/elasticsearch/reference/current/ilm-index-lifecycle.html

`,
			Keywords: []string{
				"index",
				"elasticsearch",
				"lifecycle",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_index_template",
			Category:         "Index",
			ShortDescription: `Creates or updates an index template.`,
			Description: `

Creates or updates an index template. Index templates define settings, mappings, and aliases that can be applied automatically to new indices. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html

`,
			Keywords: []string{
				"index",
				"elasticsearch",
				"template",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_ingest_pipeline",
			Category:         "Ingest",
			ShortDescription: `Manages Ingest Pipelines`,
			Description: `

Use ingest APIs to manage tasks and resources related to ingest pipelines and processors. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/ingest-apis.html

`,
			Keywords: []string{
				"ingest",
				"elasticsearch",
				"pipeline",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_logstash_pipeline",
			Category:         "Logstash",
			ShortDescription: `Creates or updates centrally managed logstash pipelines.`,
			Description: `

Creates or updates centrally managed logstash pipelines. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/logstash-apis.html

`,
			Keywords: []string{
				"logstash",
				"elasticsearch",
				"pipeline",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_script",
			Category:         "Cluster",
			ShortDescription: `Creates or updates a stored script or search template.`,
			Description: `

Creates or updates a stored script or search template. See https://www.elastic.co/guide/en/elasticsearch/reference/current/create-stored-script-api.html

`,
			Keywords: []string{
				"cluster",
				"elasticsearch",
				"script",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_security_api_key",
			Category:         "Security",
			ShortDescription: `Creates an API key for access without requiring basic authentication. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-create-api-key.html`,
			Description:      ``,
			Keywords: []string{
				"security",
				"elasticsearch",
				"api",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_security_role",
			Category:         "Security",
			ShortDescription: `Adds and updates roles in the native realm.`,
			Description: `

Adds and updates roles in the native realm. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role.html

`,
			Keywords: []string{
				"security",
				"elasticsearch",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_security_role_mapping",
			Category:         "Security",
			ShortDescription: `Manage role mappings. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html`,
			Description: `

Manage role mappings. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-role-mapping.html

`,
			Keywords: []string{
				"security",
				"elasticsearch",
				"role",
				"mapping",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_security_system_user",
			Category:         "Security",
			ShortDescription: `Updates system user's password and enablement.`,
			Description: `

Updates system user's password and enablement. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/built-in-users.html
Since this resource is to manage built-in users, destroy will not delete the underlying Elasticsearch and will only remove it from Terraform state.

`,
			Keywords: []string{
				"security",
				"elasticsearch",
				"system",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_security_user",
			Category:         "Security",
			ShortDescription: `Adds and updates users in the native realm.`,
			Description: `

Adds and updates users in the native realm. These users are commonly referred to as native users. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html

`,
			Keywords: []string{
				"security",
				"elasticsearch",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_snapshot_lifecycle",
			Category:         "Snapshot",
			ShortDescription: `Creates or updates a snapshot lifecycle policy.`,
			Description: `

Creates or updates a snapshot lifecycle policy. See, https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-put-policy.html

`,
			Keywords: []string{
				"snapshot",
				"elasticsearch",
				"lifecycle",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_snapshot_repository",
			Category:         "Snapshot",
			ShortDescription: `Registers or updates a snapshot repository.`,
			Description: `

Registers or updates a snapshot repository. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/put-snapshot-repo-api.html and https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html

`,
			Keywords: []string{
				"snapshot",
				"elasticsearch",
				"repository",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_transform",
			Category:         "Transform",
			ShortDescription: `Manages transforms. Transforms enable you to convert existing Elasticsearch indices into summarized indices.`,
			Description: `

Creates, updates, starts and stops a transform. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/transforms.html

**NOTE:** Some transform settings require a minimum Elasticsearch version. Such settings will be ignored when applied to versions below the required one (a warning will be issued in the logs).

`,
			Keywords: []string{
				"transform",
				"elasticsearch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_elasticsearch_watch",
			Category:         "Watcher",
			ShortDescription: `Adds and manages a Watch.`,
			Description: `

Adds and manages a Watch. See: https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api.html

`,
			Keywords: []string{
				"watcher",
				"elasticsearch",
				"watch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_fleet_agent_policy",
			Category:         "Fleet",
			ShortDescription: `Creates or updates a Fleet Agent Policy.`,
			Description: `

Creates or updates a Fleet Agent Policy. See https://www.elastic.co/guide/en/fleet/current/fleet-api-docs.html#create-agent-policy-api

`,
			Keywords: []string{
				"fleet",
				"agent",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_kibana_alerting_rule",
			Category:         "Kibana",
			ShortDescription: `Creates or updates a Kibana alerting rule.`,
			Description: `

Creates or updates a Kibana alerting rule. See https://www.elastic.co/guide/en/kibana/8.6/create-and-manage-rules.html

`,
			Keywords: []string{
				"kibana",
				"alerting",
				"rule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "elasticstack_kibana_space",
			Category:         "Kibana",
			ShortDescription: `Creates or updates a Kibana space.`,
			Description: `

Creates or updates a Kibana space. See https://www.elastic.co/guide/en/kibana/master/xpack-spaces.html

`,
			Keywords: []string{
				"kibana",
				"space",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"elasticstack_elasticsearch_cluster_settings":      0,
		"elasticstack_elasticsearch_component_template":    1,
		"elasticstack_elasticsearch_data_stream":           2,
		"elasticstack_elasticsearch_enrich_policy":         3,
		"elasticstack_elasticsearch_index":                 4,
		"elasticstack_elasticsearch_index_lifecycle":       5,
		"elasticstack_elasticsearch_index_template":        6,
		"elasticstack_elasticsearch_ingest_pipeline":       7,
		"elasticstack_elasticsearch_logstash_pipeline":     8,
		"elasticstack_elasticsearch_script":                9,
		"elasticstack_elasticsearch_security_api_key":      10,
		"elasticstack_elasticsearch_security_role":         11,
		"elasticstack_elasticsearch_security_role_mapping": 12,
		"elasticstack_elasticsearch_security_system_user":  13,
		"elasticstack_elasticsearch_security_user":         14,
		"elasticstack_elasticsearch_snapshot_lifecycle":    15,
		"elasticstack_elasticsearch_snapshot_repository":   16,
		"elasticstack_elasticsearch_transform":             17,
		"elasticstack_elasticsearch_watch":                 18,
		"elasticstack_fleet_agent_policy":                  19,
		"elasticstack_kibana_alerting_rule":                20,
		"elasticstack_kibana_space":                        21,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
