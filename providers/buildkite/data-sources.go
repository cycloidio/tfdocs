package buildkite

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "buildkite_meta",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Use this data source to look up the source IP addresses that Buildkite may use
to send external requests, including webhooks and API calls to source control
systems (like GitHub Enterprise Server and BitBucket Server).

Buildkite Documentation: https://buildkite.com/docs/apis/rest-api/meta

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "webhook_ips",
					Description: `A list of strings, each one an IP address (x.x.x.x) or CIDR address (x.x.x.x/32) that Buildkite may use to send webhooks and other external requests.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "webhook_ips",
					Description: `A list of strings, each one an IP address (x.x.x.x) or CIDR address (x.x.x.x/32) that Buildkite may use to send webhooks and other external requests.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buildkite_pipeline",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Use this data source to look up properties on a specific pipeline. This is
particularly useful for looking up the webhook URL for each pipeline.

Buildkite Documentation: https://buildkite.com/docs/pipelines

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `(Required) The slug of the pipeline, available in the URL of the pipeline on buildkite.com ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the pipeline.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the pipeline.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `The git URL of the repository.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `The default branch to prefill when new builds are created or triggered.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the pipeline.`,
				},
				resource.Attribute{
					Name:        "default_branch",
					Description: `The default branch to prefill when new builds are created or triggered, usually main or master but can be anything.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the pipeline.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `The git URL of the repository.`,
				},
				resource.Attribute{
					Name:        "webhook_url",
					Description: `The default branch to prefill when new builds are created or triggered.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"buildkite_meta":     0,
		"buildkite_pipeline": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
