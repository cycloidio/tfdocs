package honeycombio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_column",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

The ` + "`" + `honeycombio_column` + "`" + ` data source retrieves the details of a single column in a dataset.

-> **Note** Terraform will fail unless a column is returned by the search. Ensure that your search is specific enough to return a column.
If you want to match multiple columns, use the ` + "`" + `honeycombio_columns` + "`" + ` data source instead.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this column is associated with`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the column ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the column (string, integer, float, or boolean)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description of the column`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `whether or not the column is hidden from the query builder and results`,
				},
				resource.Attribute{
					Name:        "last_written_at",
					Description: `the ISO8601 formatted time that the column last received data`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the ISO8601 formatted time when the column was created`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the ISO8601 formatted time when the column's metadata (type, description, etc) was last changed`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `the type of the column (string, integer, float, or boolean)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description of the column`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `whether or not the column is hidden from the query builder and results`,
				},
				resource.Attribute{
					Name:        "last_written_at",
					Description: `the ISO8601 formatted time that the column last received data`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `the ISO8601 formatted time when the column was created`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `the ISO8601 formatted time when the column's metadata (type, description, etc) was last changed`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_columns",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

The columns data source allows the columns of a dataset to be retrieved.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset to retrieve the columns list from`,
				},
				resource.Attribute{
					Name:        "starts_with",
					Description: `(Optional) Only return columns starting with the given value. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `a list of all the column names found in the dataset`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `a list of all the column names found in the dataset`,
				},
			},
		},
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
			Type:             "honeycombio_derived_column",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

The ` + "`" + `honeycombio_derived_column` + "`" + ` data source retrieves the details of a single derived column.

-> **Note** Terraform will fail unless a derived column is returned by the search. Ensure that your search is specific enough to return a derived column.
If you want to match multiple derived columns, use the ` + "`" + `honeycombio_derived_columns` + "`" + ` data source instead.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this derived column is associated with. Use ` + "`" + `__all__` + "`" + ` for Environment-wide derived columns.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) The alias of the column ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the derived column.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description of the derived column`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `the expression of the derived column`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `the ID of the derived column.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `the description of the derived column`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `the expression of the derived column`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_derived_columns",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

The ` + "`" + `honeycombio_derived_columns` + "`" + ` data source allows the derived columns of a dataset to be retrieved.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset to retrieve the columns list from. Use ` + "`" + `__all__` + "`" + ` for Environment-wide derived columns.`,
				},
				resource.Attribute{
					Name:        "starts_with",
					Description: `(Optional) Only return derived columns starting with the given value. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `a list of all the derived column names found in the dataset`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `a list of all the derived column names found in the dataset`,
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
					Name:        "query_json",
					Description: `(Required) A JSON object describing the query according to the Query Specification. While the JSON can be constructed manually, it is easiest to use the honeycombio_query_specification data source. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "query_id",
					Description: `The ID of the Query created and executed to obtain the result.`,
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
					Name:        "query_id",
					Description: `The ID of the Query created and executed to obtain the result.`,
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
			Type:             "honeycombio_recipient",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

` + "`" + `honeycombio_recipient` + "`" + ` data source provides details about a specific recipient in the Team.

The ID of an existing recipient can be used when adding recipients to triggers or burn alerts.

-> **Note** Terraform will fail unless exactly one recipient is returned by the search. Ensure that your search is specific enough to return a single recipient ID only.
If you want to match multiple recipients, use the ` + "`" + `honeycombio_recipients` + "`" + ` data source instead.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of recipient, allowed types are ` + "`" + `email` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + ` and ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Optional) Deprecated: recipients are now a Team-level construct. Any provided value will be ignored.`,
				},
				resource.Attribute{
					Name:        "detail_filter",
					Description: `(Optional) a block to further filter recipients as described below.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Deprecated: use ` + "`" + `detail_filter` + "`" + ` instead. The target of the recipient, this has another meaning depending on the type of recipient (see the table below). Type | Target ----------|------------------------- email | an email address marker | name of the marker pagerduty | _N/A_ slack | name of the channel webhook | name of the webhook To further filter the recipient results, a ` + "`" + `detail_filter` + "`" + ` block can be provided which accepts the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the detail field to filter by. Allowed values are ` + "`" + `address` + "`" + `, ` + "`" + `channel` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `integration_name` + "`" + `, and ` + "`" + `url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the detail field to match on.`,
				},
				resource.Attribute{
					Name:        "value_regex",
					Description: `(Optional) A regular expression string to apply to the value of the detail field to match on. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the recipient.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The email recipient's address -- if of type ` + "`" + `email` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `The Slack recipient's channel -- if of type ` + "`" + `slack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The webhook recipient's name -- if of type ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Sensitive) The webhook recipient's secret -- if of type ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The webhook recipient's URL - if of type ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Sensitive) The PagerDuty recipient's integration key -- if of type ` + "`" + `pagerduty` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "integration_name",
					Description: `The PagerDuty recipient's inregration name -- if of type ` + "`" + `pagerduty` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the recipient.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The email recipient's address -- if of type ` + "`" + `email` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "channel",
					Description: `The Slack recipient's channel -- if of type ` + "`" + `slack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The webhook recipient's name -- if of type ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Sensitive) The webhook recipient's secret -- if of type ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The webhook recipient's URL - if of type ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Sensitive) The PagerDuty recipient's integration key -- if of type ` + "`" + `pagerduty` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "integration_name",
					Description: `The PagerDuty recipient's inregration name -- if of type ` + "`" + `pagerduty` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_recipients",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

` + "`" + `honeycombio_recipients` + "`" + ` data source provides recipient IDs of recipients matching a set of criteria.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of recipient, allowed types are ` + "`" + `email` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + ` and ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detail_filter",
					Description: `(Optional) a block to further filter recipients as described below. ` + "`" + `type` + "`" + ` must be set when providing a filter. To further filter the recipient results, a ` + "`" + `detail_filter` + "`" + ` block can be provided which accepts the following arguments:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the detail field to filter by. Allowed values are ` + "`" + `address` + "`" + `, ` + "`" + `channel` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `integration_name` + "`" + `, and ` + "`" + `url` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The value of the detail field to match on.`,
				},
				resource.Attribute{
					Name:        "value_regex",
					Description: `(Optional) A regular expression string to apply to the value of the detail field to match on. ~>`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the recipient IDs found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the recipient IDs found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_trigger_recipient",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `

Search the triggers of a dataset for a trigger recipient. The ID of the existing trigger recipient can be used when adding trigger recipients to new triggers.

-> **Deprecated** Use the ` + "`" + `honeycombio_recipient` + "`" + ` data source instead.

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

		"honeycombio_column":              0,
		"honeycombio_columns":             1,
		"honeycombio_datasets":            2,
		"honeycombio_derived_column":      3,
		"honeycombio_derived_columns":     4,
		"honeycombio_query_result":        5,
		"honeycombio_query_specification": 6,
		"honeycombio_recipient":           7,
		"honeycombio_recipients":          8,
		"honeycombio_trigger_recipient":   9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
