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
	}

	resourcesMap = map[string]int{

		"elasticstack_elasticsearch_cluster_settings":    0,
		"elasticstack_elasticsearch_component_template":  1,
		"elasticstack_elasticsearch_data_stream":         2,
		"elasticstack_elasticsearch_index":               3,
		"elasticstack_elasticsearch_index_lifecycle":     4,
		"elasticstack_elasticsearch_index_template":      5,
		"elasticstack_elasticsearch_ingest_pipeline":     6,
		"elasticstack_elasticsearch_security_role":       7,
		"elasticstack_elasticsearch_security_user":       8,
		"elasticstack_elasticsearch_snapshot_lifecycle":  9,
		"elasticstack_elasticsearch_snapshot_repository": 10,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
