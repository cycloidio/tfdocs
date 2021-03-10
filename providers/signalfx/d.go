package signalfx

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "signalfx_aws_services",
			Category:         "Data Sources",
			ShortDescription: `Provides a list AWS service names.`,
			Description: `

Use this data source to get a list of AWS service names.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_azure_services",
			Category:         "Data Sources",
			ShortDescription: `Provides a list Azure service names.`,
			Description: `

Use this data source to get a list of Azure service names.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_dimension_values",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of dimension values given a query`,
			Description: `

Use this data source to get a list of dimension values matching the provided query.

~> **NOTE** This data source only allows 1000 values, as it's kinda nuts to make anything with ` + "`" + `for_each` + "`" + ` that big in SignalFx. This is negotiable.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_gcp_services",
			Category:         "Data Sources",
			ShortDescription: `Provides a list GCP service names.`,
			Description: `

Use this data source to get a list of GCP service names.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "signalfx_pagerduty_integration",
			Category:         "Data Sources",
			ShortDescription: `Provides information on an existing PagerDuty integration.`,
			Description: `

Use this data source to get information on an existing PagerDuty integration.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Specify the exact name of the desired PagerDuty integration ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the integration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether the integration is enabled.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the integration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the integration.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Whether the integration is enabled.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"signalfx_aws_services":          0,
		"signalfx_azure_services":        1,
		"signalfx_dimension_values":      2,
		"signalfx_gcp_services":          3,
		"signalfx_pagerduty_integration": 4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
