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
					Name:        "column_layout",
					Description: `(Optional) the number of columns to layout on the board, either ` + "`" + `multi` + "`" + ` (the default) or ` + "`" + `single` + "`" + `. Only ` + "`" + `visual` + "`" + ` style boards (see below) have a column layout.`,
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
					Description: `(Deprecated) The dataset this query is associated with.`,
				},
				resource.Attribute{
					Name:        "caption",
					Description: `(Optional) A description of the query that will be displayed on the board. Supports markdown.`,
				},
				resource.Attribute{
					Name:        "graph_settings",
					Description: `(Optional) A map of boolean toggles to manages the settings for this query's graph on the board. If a value is unspecified, it is assumed to be false. Currently supported toggles are:`,
				},
				resource.Attribute{
					Name:        "query_style",
					Description: `(Optional) How the query should be displayed within the board, either ` + "`" + `graph` + "`" + ` (the default), ` + "`" + `table` + "`" + ` or ` + "`" + `combo` + "`" + `. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the board.`,
				},
				resource.Attribute{
					Name:        "board_url",
					Description: `The URL to the board in the Honeycomb UI. ## Import Boards can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import honeycombio_board.my_board AobW9oAZX71 ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the board from the UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the board.`,
				},
				resource.Attribute{
					Name:        "board_url",
					Description: `The URL to the board in the Honeycomb UI. ## Import Boards can be imported using their ID, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import honeycombio_board.my_board AobW9oAZX71 ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the board from the UI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_burn_alert",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a burn alert. For more information about burn alerts, check out [Define Burn Alerts](https://docs.honeycomb.io/working-with-your-data/slos/slo-process/#define-burn-alerts).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "slo_id",
					Description: `(Required) ID of the SLO this burn alert is associated with.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this burn alert is associated with.`,
				},
				resource.Attribute{
					Name:        "exhaustion_minutes",
					Description: `(Required) The amount of time, in minutes, remaining before the SLO's error budget will be exhausted and the alert will fire.`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Optional) Zero or more configuration blocks (described below) with the recipients to notify when the alert fires. Each burn alert configuration may have one or more ` + "`" + `recipient` + "`" + ` blocks, which each accept the following arguments. A recipient block can either refer to an existing recipient (a recipient that is already present in another burn alert or trigger) or a new recipient. When specifying an existing recipient, only ` + "`" + `id` + "`" + ` may be set. If you pass in a recipient without its ID and only include the type and target, Honeycomb will make a best effort to match to an existing recipient. To retrieve the ID of an existing recipient, refer to the [` + "`" + `honeycombio_recipient` + "`" + `](../data-sources/recipient.md) data source.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the recipient, allowed types are ` + "`" + `email` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + ` and ` + "`" + `webhook` + "`" + `. Should not be used in combination with ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) Target of the recipient, this has another meaning depending on the type of recipient (see the table below). Should not be used in combination with ` + "`" + `id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of an already existing recipient. Should not be used in combination with ` + "`" + `type` + "`" + ` and ` + "`" + `target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_details",
					Description: `(Optional) a block of additional details to send along with the notification. The only supported option currently is ` + "`" + `pagerduty_severity` + "`" + ` which has a default value of ` + "`" + `critical` + "`" + ` but can be set to one of ` + "`" + `info` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `error` + "`" + `, or ` + "`" + `critical` + "`" + ` and must be used in combination with a PagerDuty recipient. Type | Target ----------|------------------------- email | an email address pagerduty | _N/A_ slack | name of the channel webhook | name of the webhook ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the burn alert. ## Import Burn Alerts can be imported using a combination of the dataset name and their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_burn_alert.my_alert my-dataset/bj9BwOb1uKz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the burn alert. ## Import Burn Alerts can be imported using a combination of the dataset name and their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_burn_alert.my_alert my-dataset/bj9BwOb1uKz ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_column",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a Honeycomb Column resource.
This can be used to create, update, and delete columns in a dataset.

-> **Note**: deleting a column is a destructive and irreversible operation which also removes the data in the column.

-> **Note**: prior to version 0.13.0 of the provider, columns were *not* deleted on destroy but left in place and only removed from state.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this column is added to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the column. Must be unique per dataset.`,
				},
				resource.Attribute{
					Name:        "key_name",
					Description: `(Deprecated) Please use ` + "`" + `name` + "`" + ` instead. The name of the column. Must be unique per dataset. Conficts with ` + "`" + `name` + "`" + `.`,
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
					Description: `ID of the column.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO8601 formatted time the column was created`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `ISO8601 formatted time the column was updated`,
				},
				resource.Attribute{
					Name:        "last_written_at",
					Description: `ISO8601 formatted time the column was last written to (received event data) ## Import Columns can be imported using a combination of the dataset name and their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_column.my_column my-dataset/duration_ms ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the column.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO8601 formatted time the column was created`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `ISO8601 formatted time the column was updated`,
				},
				resource.Attribute{
					Name:        "last_written_at",
					Description: `ISO8601 formatted time the column was last written to (received event data) ## Import Columns can be imported using a combination of the dataset name and their name, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_column.my_column my-dataset/duration_ms ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `(Required) The name of the dataset.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A longer description for dataset.`,
				},
				resource.Attribute{
					Name:        "expand_json_depth",
					Description: `(Optional) The maximum unpacking depth of nested JSON fields. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the dataset.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO8601 formatted time the column was created`,
				},
				resource.Attribute{
					Name:        "last_written_at",
					Description: `ISO8601 formatted time the column was last written to (received event data) ## Import Datasets can be imported by their slug, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import honeycombio_dataset.my_dataset my-dataset ` + "`" + `` + "`" + `` + "`" + ` You can find the slug in the URL bar when visiting the Dataset from the UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "slug",
					Description: `The slug of the dataset.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `ISO8601 formatted time the column was created`,
				},
				resource.Attribute{
					Name:        "last_written_at",
					Description: `ISO8601 formatted time the column was last written to (received event data) ## Import Datasets can be imported by their slug, e.g. ` + "`" + `` + "`" + `` + "`" + `shell $ terraform import honeycombio_dataset.my_dataset my-dataset ` + "`" + `` + "`" + `` + "`" + ` You can find the slug in the URL bar when visiting the Dataset from the UI.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_dataset_definitions",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Dataset Definitions define the fields in your Dataset that have special meaning.

-> **Note** Some Dataset Definitions are automatically set when a dataset is created or first receives an event.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "column_type",
					Description: `The type of the column assigned to the definition. Will be one of ` + "`" + `column` + "`" + ` or ` + "`" + `derived_column` + "`" + `. ## Import Dataset Definitions cannot be imported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "column_type",
					Description: `The type of the column assigned to the definition. Will be one of ` + "`" + `column` + "`" + ` or ` + "`" + `derived_column` + "`" + `. ## Import Dataset Definitions cannot be imported.`,
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
			Type:             "honeycombio_email_recipient",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

` + "`" + `honeycombio_email_recipient` + "`" + ` allows you to define and manage an Email recipient that can be used by Triggers or BurnAlerts notifications.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The email address to send the notification to. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import Email Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_email_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import Email Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_email_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "honeycombio_marker_setting",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a marker setting. For more information about marker settings, check out the [Marker Settings API](https://docs.honeycomb.io/api/marker-settings/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this marker setting is placed on.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the marker setting, Honeycomb.io can display markers in different colors depending on their type.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Required) The color set for the marker as a hex color code (e.g. ` + "`" + `#DF4661` + "`" + `) ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the marker setting.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp when the marker setting was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp when the marker setting was last modified.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the marker setting.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Timestamp when the marker setting was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Timestamp when the marker setting was last modified.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_pagerduty_recipient",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

` + "`" + `honeycombio_pagerduty_recipient` + "`" + ` allows you to define and manage a PagerDuty recipient that can be used by Triggers or BurnAlerts notifications.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_key",
					Description: `(Required) The key of the PagerDuty Integration to send the notification to.`,
				},
				resource.Attribute{
					Name:        "integration_name",
					Description: `(Required) The name of the PagerDuty Integration to send the notification to. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import PagerDuty Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_pagerduty_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import PagerDuty Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_pagerduty_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
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
			Type:             "honeycombio_slack_recipient",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

` + "`" + `honeycombio_slack_recipient` + "`" + ` allows you to define and manage a Slack channel or user recipient that can be used by Triggers or BurnAlerts notifications.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "channel",
					Description: `(Required) The Slack channel or username to send the notification to. Must begin with ` + "`" + `#` + "`" + ` or ` + "`" + `@` + "`" + `. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import Slack Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_slack_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import Slack Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_slack_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_slo",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Creates a service level objective (SLO). For more information about SLOs, check out [Set Service Level Objectives (SLOs)](https://docs.honeycomb.io/working-with-your-data/slos/).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SLO.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the SLO's intent and context.`,
				},
				resource.Attribute{
					Name:        "dataset",
					Description: `(Required) The dataset this SLO is created in. Must be the same dataset as the SLI unless the SLI's dataset is ` + "`" + `"__all__"` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sli",
					Description: `(Required) The alias of the Derived Column that will be used as the SLI to indicate event success. The derived column used as the SLI must be in the same dataset as the SLO. Additionally, the column evaluation should consistently return nil, true, or false, as these are the only valid values for an SLI.`,
				},
				resource.Attribute{
					Name:        "target_percentage",
					Description: `(Required) The percentage of qualified events that you expect to succeed during the ` + "`" + `time_period` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_period",
					Description: `(Required) The time period, in days, over which your SLO will be evaluated. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SLO. ## Import SLOs can be imported using a combination of the dataset name and their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_slo.my_slo my-dataset/bj9BwOb1uKz ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the SLO from the UI.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SLO. ## Import SLOs can be imported using a combination of the dataset name and their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_slo.my_slo my-dataset/bj9BwOb1uKz ` + "`" + `` + "`" + `` + "`" + ` You can find the ID in the URL bar when visiting the SLO from the UI.`,
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
					Description: `(Optional) The interval (in seconds) in which to check the results of the query’s calculation against the threshold. This value must be divisible by 60, between 60 and 86400 (between 1 minute and 1 day), and not be more than 4 times the query's duration. Defaults to 900 (15 minutes).`,
				},
				resource.Attribute{
					Name:        "alert_type",
					Description: `(Optional) The frequency for the alert to trigger. (` + "`" + `on_change` + "`" + ` is the default behavior, ` + "`" + `on_true` + "`" + ` can also be selected)`,
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
					Description: `(Required) The value to be used with the operator. Each trigger configuration may have zero or more ` + "`" + `recipient` + "`" + ` blocks, which each accept the following arguments. A trigger recipient block can either refer to an existing recipient (a recipient that is already present in another trigger) or a new recipient. When specifying an existing recipient, only ` + "`" + `id` + "`" + ` may be set. If you pass in a recipient without its ID and only include the type and target, Honeycomb will make a best effort to match to an existing recipient. To retrieve the ID of an existing recipient, refer to the [` + "`" + `honeycombio_recipient` + "`" + `](../data-sources/recipient.md) data source.`,
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
					Description: `(Optional) The ID of an already existing recipient. Should not be used in combination with ` + "`" + `type` + "`" + ` and ` + "`" + `target` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_details",
					Description: `(Optional) a block of additional details to send along with the notification. The only supported option currently is ` + "`" + `pagerduty_severity` + "`" + ` which has a default value of ` + "`" + `critical` + "`" + ` but can be set to one of ` + "`" + `info` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `error` + "`" + `, or ` + "`" + `critical` + "`" + ` and must be used in combination with a PagerDuty recipient. Type | Target ----------|------------------------- email | an email address marker | name of the marker pagerduty | _N/A_ slack | name of the channel webhook | name of the webhook ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
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
		&resource.Resource{
			Name:             "",
			Type:             "honeycombio_webhook_recipient",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

` + "`" + `honeycombio_webhook_recipient` + "`" + ` allows you to define and manage a Webhook recipient that can be used by Triggers or BurnAlerts notifications.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Webhook Integration to create.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `(Required) The secret to include when sending the notification to the webhook.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) The URL of the endpoint to send the notification to. ## Attribute Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import Webhook Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_webhook_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the recipient. ## Import Webhook Recipients can be imported by their ID, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import honeycombio_webhook_recipient.my_recipient nx2zsegA0dZ ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"honeycombio_board":               0,
		"honeycombio_burn_alert":          1,
		"honeycombio_column":              2,
		"honeycombio_dataset":             3,
		"honeycombio_dataset_definitions": 4,
		"honeycombio_derived_column":      5,
		"honeycombio_email_recipient":     6,
		"honeycombio_marker":              7,
		"honeycombio_marker_setting":      8,
		"honeycombio_pagerduty_recipient": 9,
		"honeycombio_query":               10,
		"honeycombio_query_annotation":    11,
		"honeycombio_slack_recipient":     12,
		"honeycombio_slo":                 13,
		"honeycombio_trigger":             14,
		"honeycombio_webhook_recipient":   15,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
