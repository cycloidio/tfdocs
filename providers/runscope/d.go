package runscope

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "runscope_bucket",
			Category:         "Data Sources",
			ShortDescription: `Get information about a single runscope bucket.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The unique key of the bucket. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique key of the found bucket.`,
				},
				resource.Attribute{
					Name:        "team_uuid",
					Description: `The team unique identifier that owns the bucket.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Type name of the bucket.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique key of the found bucket.`,
				},
				resource.Attribute{
					Name:        "team_uuid",
					Description: `The team unique identifier that owns the bucket.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Type name of the bucket.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_buckets",
			Category:         "Data Sources",
			ShortDescription: `Get information about runscope buckets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter to reduce the list of buckets returned. Variables (` + "`" + `filter` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field to filter on, currently either: ` + "`" + `key` + "`" + `, ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `The list of values to match against ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "keys",
					Description: `A list of the keys of matching buckets.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "keys",
					Description: `A list of the keys of matching buckets.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_integration",
			Category:         "Data Sources",
			ShortDescription: `Get information about runscope integrations enabled on for your team.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_uuid",
					Description: `(Required) Your team unique identifier.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of integration to lookup i.e. pagerduty ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the found integration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the found integration.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "runscope_integrations",
			Category:         "Data Sources",
			ShortDescription: `Get information about runscope integrations enabled on for your team.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter to reduce the list of integrations returned. Variables (` + "`" + `filter` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the field to filter on, currently either: ` + "`" + `id` + "`" + `, ` + "`" + `type` + "`" + ` or ` + "`" + `description` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `The list of values to match against ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the found integration.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the found integration.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"runscope_bucket":       0,
		"runscope_buckets":      1,
		"runscope_integration":  2,
		"runscope_integrations": 3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
