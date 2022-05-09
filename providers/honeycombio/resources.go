package honeycombio

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_board",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a board. For more information about boards, check out [Collaborate with Boards](https://docs.honeycomb.io/working-with-your-data/collaborating/boards/#docs-sidebar).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the board.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the board. Supports markdown.`,
				},
				resource.Attribute{
					Name:        "style",
					Description: `(Optional) How the board should be displayed in the UI, either ` + "`" + `list` + "`" + ` (the default) or ` + "`" + `visual` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Optional) Zero or more configurations blocks (described below) with the queries of the board. Each board configuration may have zero or more ` + "`" + `query` + "`" + ` blocks, which accepts the following arguments:`,
				},
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required) The ID of the Query to run.`,
				},
				resource.Attribute{
					Name:        "query_annotation_id",
					Description: `(Optional) The ID of the Query Annotation to associate with this query.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this query is associated with.`,
				},
				resource.Attribute{
					Name:        "caption",
					Description: `(Optional) A description of the query that will be displayed on the board. Supports markdown.`,
				},
				resource.Attribute{
					Name:        "query_style",
					Description: `(Optional) How the query should be displayed within the board, either ` + "`" + `graph` + "`" + ` (the default), ` + "`" + `table` + "`" + ` or ` + "`" + `combo` + "`" + `. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the board. ## Import Boards can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_board.my_board AobW9oAZX71 ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the board from the UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the board. ## Import Boards can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_board.my_board AobW9oAZX71 ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the board from the UI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_column",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a column in a dataset.

-> **Note** Destroying or replacing this resource will not delete the previously created column. This is intentional to preserve the column and its data. The API supports deleting columns, so if this is required for your use case please open an issue in [honeycombio/terraform-provider-honeycombio](https://github.com/honeycombio/terraform-provider-honeycombio).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this column is added to.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Required) The name of the column. Must be unique per dataset.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the column, allowed values are ` + "`" + `string` + "`" + `, ` + "`" + `float` + "`" + `, ` + "`" + `integer` + "`" + ` and ` + "`" + `boolean` + "`" + `. Defaults to ` + "`" + `string` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `(Optional) Whether this column should be hidden in the query builder and sample data. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description that is shown in the UI. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the column. ## Import Columns can be imported using a combination of the dataset name and their key name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_column.my_column my-dataset/duration_ms ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the column. ## Import Columns can be imported using a combination of the dataset name and their key name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_column.my_column my-dataset/duration_ms ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_dataset",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a dataset.

-> **Note** If this dataset already exists, creating this resource is a no-op.

-> **Note** Destroying or replacing this resource will not delete the created dataset. It's not possible to delete a dataset using the API.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the dataset. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the dataset.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the dataset.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_derived_column",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a derived column. For more information about derived columns, check out [Calculate with derived columns](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this derived column is added to.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `(Required) The name of the derived column. Must be unique per dataset.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) The function of the derived column. See [Derived Column Syntax](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/derived-columns/#derived-column-syntax).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description that is shown in the UI. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the derived column. ## Import Derived columns can be imported using a combination of the dataset name and their alias, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_derived_column.my_column my-dataset/duration_ms_log10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the derived column. ## Import Derived columns can be imported using a combination of the dataset name and their alias, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_derived_column.my_column my-dataset/duration_ms_log10 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_marker",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a marker. For more information about markers, check out [Annotate the timeline with Markers](https://docs.honeycomb.io/working-with-your-data/customizing-your-query/markers/).

-> **Note** Destroying or replacing this resource will not delete the previously created marker. This is intentional to preserve the markers. At this time, it is not possible to remove markers using this provider.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this marker is placed on.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `(Optional) The message on the marker.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the marker, Honeycomb.io can display markers in different colors depending on their type.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Optional) A target for the Marker. If you click on the Marker text, it will take you to this URL. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the marker.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the marker.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_query",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a query in a dataset.

Queries can be used by triggers and boards, or be executed via the [Query Data API](https://docs.honeycomb.io/api/query-results/).

-> **Note** Queries can only be created or read. Any changes will result in a new query object being created, and destroying it does nothing.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this query is added to.`,
				},
				resource.Attribute{
					Name:        "query_json",
					Description: `(Required) A JSON object describing the query according to the [Query Specification](https://docs.honeycomb.io/api/query-specification/#fields-on-a-query-specification). While the JSON can be constructed manually, it is easiest to use the [` + "`" + `honeycombio_query_specification` + "`" + `](../data-sources/query_specification.md) data source. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the query. Useful for adding it to a board and/or creating a query annotation. ## Import Queries cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the query. Useful for adding it to a board and/or creating a query annotation. ## Import Queries cannot be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_query_annotation",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a query annotation in a dataset.

-> **Note** A query annotation points to a specific query. Any change to the query will result in a new query ID and the annotation will no longer apply.
If you use the "honeycombio_query_specification" to determine the ` + "`" + `query_id` + "`" + ` parameter (as in the example below), Terraform will destroy the old query annotation and create a new one.
If this is wrong for your use case, please open an issue in [honeycombio/terraform-provider-honeycombio](https://github.com/honeycombio/terraform-provider-honeycombio).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this query is added to.`,
				},
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required) The ID of the query that the annotation will be created on. Note that a query can have more than one annotation.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the query annotation that will display in the Honeycomb UI.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description for the query annotation. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the query annotation. Useful for adding it to a board. ## Import Query annotations cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the query annotation. Useful for adding it to a board. ## Import Query annotations cannot be imported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_trigger",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a trigger. For more information about triggers, check out [Alert with Triggers](https://docs.honeycomb.io/working-with-your-data/triggers/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the trigger.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this trigger is associated with.`,
				},
				resource.Attribute{
					Name:        "query_id",
					Description: `(Required) The ID of the Query that the Trigger will execute.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) A configuration block (described below) describing the threshold of the trigger.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the trigger.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) The state of the trigger. If true, the trigger will not be run. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `(Optional) The interval (in seconds) in which to check the results of the query’s calculation against the threshold. Value must be divisible by 60 and between 60 and 86400 (between 1 minute and 1 day). Defaults to 900 (15 minutes).`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Optional) Zero or more configuration blocks (described below) with the recipients to notify when the trigger fires. ->`,
				},
				resource.Attribute{
					Name:        "op",
					Description: `(Required) The operator to apply, allowed threshold operators are ` + "`" + `>` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<` + "`" + `, and ` + "`" + `<=` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value to be used with the operator. Each trigger configuration may have zero or more ` + "`" + `recipient` + "`" + ` blocks, which each accept the following arguments. A trigger recipient block can either refer to an existing recipient (a recipient that is already present in another trigger) or a new recipient. When specifying an existing recipient, only ` + "`" + `id` + "`" + ` may be set. If you pass in a recipient without its ID and only include the type and target, Honeycomb will make a best effort to match to an existing recipient. To retrieve the ID of an existing recipient, refer to the [` + "`" + `honeycombio_trigger_recipient` + "`" + `](../data-sources/trigger_recipient.md) data source.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the trigger recipient, allowed types are ` + "`" + `email` + "`" + `, ` + "`" + `marker` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + ` and ` + "`" + `webhook` + "`" + `. Should not be used in combination with ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Target of the trigger recipient, this has another meaning depending on the type of recipient (see the table below). Should not be used in combination with ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an already existing recipient. Should not be used in combination with ` + "`" + `type` + "`" + ` and ` + "`" + `target` + "`" + `. Type | Target ----------|------------------------- email | an email address marker | name of the marker pagerduty | _N/A_ slack | name of the channel webhook | name of the webhook ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the trigger. ## Import Triggers can be imported using a combination of the dataset name and their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_trigger.my_trigger my-dataset/AeZzSoWws9G ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the trigger from the UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the trigger. ## Import Triggers can be imported using a combination of the dataset name and their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_trigger.my_trigger my-dataset/AeZzSoWws9G ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the trigger from the UI.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"honeycombio_board":            0,
		"honeycombio_column":           1,
		"honeycombio_dataset":          2,
		"honeycombio_derived_column":   3,
		"honeycombio_marker":           4,
		"honeycombio_query":            5,
		"honeycombio_query_annotation": 6,
		"honeycombio_trigger":          7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
