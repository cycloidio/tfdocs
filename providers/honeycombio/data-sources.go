package honeycombio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_datasets",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

The datasets data source allows the datasets of an account to be retrieved.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "starts_with",
					Description: `(Optional) Only return datasets starting with the given value. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `a list of all the dataset names.`,
				},
				resource.Attribute{
					Name:        "slugs",
					Description: `a list of all the dataset slugs.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `a list of all the dataset names.`,
				},
				resource.Attribute{
					Name:        "slugs",
					Description: `a list of all the dataset slugs.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_query_result",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

The ` + "`" + `query_result` + "`" + ` data source allows you to execute Honeycomb queries via the [Query Data API](https://docs.honeycomb.io/api/query-results/).
As this data source is a wrapper around the Query Data API all of its [documented restrictions](https://docs.honeycomb.io/api/query-results/#api-restrictions) apply.

~> **NOTE:** Use of this data source requires a Honeycomb Enterprise plan.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this query is associated with.`,
				},
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required) The ID of the query that will be executed to obtain the result. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "query_url",
					Description: `The permalink to the executed query's results.`,
				},
				resource.Attribute{
					Name:        "query_image_url",
					Description: `The permalink to the visualization of the executed query's results.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `The results of the executed query. This will be a list of maps, with each map's keys set to the breakdowns and calculations of the query. Due to a limitation of the Terraform Plugin SDK, all values are transformed into strings.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "query_url",
					Description: `The permalink to the executed query's results.`,
				},
				resource.Attribute{
					Name:        "query_image_url",
					Description: `The permalink to the visualization of the executed query's results.`,
				},
				resource.Attribute{
					Name:        "results",
					Description: `The results of the executed query. This will be a list of maps, with each map's keys set to the breakdowns and calculations of the query. Due to a limitation of the Terraform Plugin SDK, all values are transformed into strings.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_query_specification",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Generates a [Query Specificaiton](https://docs.honeycomb.io/api/query-specification/) in JSON format.

This is a data source which can be used to construct a JSON representation of a Honeycomb [Query Specification](https://docs.honeycomb.io/api/query-specification/). The ` + "`" + `json` + "`" + ` attribute contains a serialized JSON representation which can be passed to the ` + "`" + `query_json` + "`" + ` field of the [` + "`" + `honeycombio_query` + "`" + `](../resources/query.md) resource for use in boards and triggers.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "calculation",
					Description: `(Optional) Zero or more configuration blocks (described below) with the calculations that should be displayed. If no calculations are specified, ` + "`" + `COUNT` + "`" + ` will be used.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Zero or more configuration blocks (described below) with the filters that should be applied.`,
				},
				resource.Attribute{
					Name:        "filter_combination",
					Description: `(Optional) How to combine multiple filters, either ` + "`" + `AND` + "`" + ` (default) or ` + "`" + `OR` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "breakdowns",
					Description: `(Optional) A list of fields to group by.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) Zero or more configuration blocks (described below) describing how to order the query results. Each term must appear in either ` + "`" + `calculation` + "`" + ` or ` + "`" + `breakdowns` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "having",
					Description: `(Optional) Zero or more filters used to restrict returned groups in the query result.`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `(Optional) The maximum number of query results, must be between 1 and 1000.`,
				},
				resource.Attribute{
					Name:        "time_range",
					Description: `(Optional) The time range of the query in seconds, defaults to two hours.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `(Optional) The absolute start time of the query in Unix Time (= seconds since epoch).`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The absolute end time of the query in Unix Time (= seconds since epoch).`,
				},
				resource.Attribute{
					Name:        "granularity",
					Description: `(Optional) The time resolution of the query’s graph, in seconds. Valid values must be in between the query’s time range /10 at maximum, and /1000 at minimum. ~>`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Required) The operator to apply, see the supported list of calculation operators at [Calculation Operators](https://docs.honeycomb.io/api/query-specification/#calculation-operators).`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) The column to apply the operator to, not needed with ` + "`" + `COUNT` + "`" + ` or ` + "`" + `CONCURRENCY` + "`" + `. Each query configuration may have zero or more ` + "`" + `filter` + "`" + ` blocks, which each accept the following arguments:`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Required) The column to apply the filter to.`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Required) The operator to apply, see the supported list of filter operators at [Filter Operators](https://docs.honeycomb.io/api/query-specification/#filter-operators). Not all operators require a value.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value used for the filter. Not needed if op is ` + "`" + `exists` + "`" + `, ` + "`" + `not-exists` + "`" + `, ` + "`" + `in` + "`" + ` or ` + "`" + `not-in` + "`" + `. Mutually exclusive with the other ` + "`" + `value_`,
				},
				resource.Attribute{
					Name:        "value_string",
					Description: `(Optional) Deprecated: use 'value' instead. The value used for the filter when the column is a string. Mutually exclusive with ` + "`" + `value` + "`" + ` and the other ` + "`" + `value_`,
				},
				resource.Attribute{
					Name:        "value_integer",
					Description: `(Optional) Deprecated: use 'value' instead. The value used for the filter when the column is an integer. Mutually exclusive with ` + "`" + `value` + "`" + ` and the other ` + "`" + `value_`,
				},
				resource.Attribute{
					Name:        "value_float",
					Description: `(Optional) Deprecated: use 'value' instead. The value used for the filter when the column is a float. Mutually exclusive with ` + "`" + `value` + "`" + ` and the other ` + "`" + `value_`,
				},
				resource.Attribute{
					Name:        "value_boolean",
					Description: `(Optional) Deprecated: use 'value' instead. The value used for the filter when the column is a boolean. Mutually exclusive with ` + "`" + `value` + "`" + ` and the other ` + "`" + `value_`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Optional) The calculation operator to apply, see the supported list of calculation operators at [Calculation Operators](https://docs.honeycomb.io/api/query-specification/#calculation-operators).`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `(Optional) Either a column present in ` + "`" + `breakdown` + "`" + ` or a column to ` + "`" + `op` + "`" + ` applies to.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `(Optional) The sort direction, if set must be ` + "`" + `ascending` + "`" + ` or ` + "`" + `descending` + "`" + `. Each query configuration may have zero or more ` + "`" + `having` + "`" + ` blocks, which each accept the following arguments.`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `The operator to apply to filter the query results. One of ` + "`" + `=` + "`" + `, ` + "`" + `!=` + "`" + `, ` + "`" + `>` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<` + "`" + `, or ` + "`" + `<=` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "calculate_op",
					Description: `The calculation operator to apply, supports all of the [Calculation Operators](https://docs.honeycomb.io/api/query-specification/#calculation-operators) with the exception of ` + "`" + `HEATMAP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `The column to apply the ` + "`" + `calculate_op` + "`" + ` to, not needed with ` + "`" + `COUNT` + "`" + ` or ` + "`" + `CONCURRENCY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value used with ` + "`" + `op` + "`" + `. Currently assumed to be a number. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the trigger.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `JSON representation of the query according to the [Query Specification](https://docs.honeycomb.io/api/query-specification/#fields-on-a-query-specification), can be used as input for other resources.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the trigger.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `JSON representation of the query according to the [Query Specification](https://docs.honeycomb.io/api/query-specification/#fields-on-a-query-specification), can be used as input for other resources.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_trigger_recipient",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Search the triggers of a dataset for a trigger recipient. The ID of the existing trigger recipient can be used when adding trigger recipients to new triggers. Specifying a trigger recipient by its ID is necessary when adding Slack recipients, since these can not be created using the API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) Search through all triggers linked to this dataset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of recipient, allowed types are ` + "`" + `email` + "`" + `, ` + "`" + `marker` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + ` and ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Target of the trigger, this has another meaning depending on the type of recipient (see the table below). Type | Target ----------|------------------------- email | an email address marker | name of the marker pagerduty | _N/A_ slack | name of the channel webhook | name of the webhook ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the trigger recipient.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the trigger recipient.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"honeycombio_datasets":            0,
		"honeycombio_query_result":        1,
		"honeycombio_query_specification": 2,
		"honeycombio_trigger_recipient":   3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
