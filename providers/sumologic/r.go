package sumologic

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sumologic_cloudsyslog_source",
			Category:         "Resources",
			ShortDescription: `Provides a Sumo Logic Cloud Syslog source.`,
			Description:      ``,
			Keywords: []string{
				"cloudsyslog",
				"source",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_collector",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic (Hosted) Collector.`,
			Description:      ``,
			Keywords: []string{
				"collector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the collector. This is required, and has to be unique. Changing this will force recreation the collector.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the collector.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) The default source category for any source attached to this collector. Can be overridden in the configuration of said sources.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) The time zone to use for this collector. The value follows the [tzdata][2] naming convention.`,
				},
				resource.Attribute{
					Name:        "fields",
					Description: `(Optional) Map containing [key/value pairs][3]. The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The internal ID of the collector. This can be used to attach sources to the collector. ## Import Collectors can be imported using the collector id, e.g.: ` + "`" + `` + "`" + `` + "`" + `hcl terraform import sumologic_collector.test 1234567890 ` + "`" + `` + "`" + `` + "`" + ` Collectors can also be imported using the collector name, which is unique per Sumo Logic account, e.g.: ` + "`" + `` + "`" + `` + "`" + `hcl terraform import sumologic_collector.test my_test_collector ` + "`" + `` + "`" + `` + "`" + ` [1]: https://help.sumologic.com/Send_Data/Hosted_Collectors [2]: https://en.wikipedia.org/wiki/Tz_database [3]: https://help.sumologic.com/Manage/Fields [4]: https://www.terraform.io/docs/configuration/resources.html#prevent_destroy`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_connection",
			Category:         "Resources",
			ShortDescription: `Provides the ability to create, read, delete, update connections.`,
			Description:      ``,
			Keywords: []string{
				"connection",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_content",
			Category:         "Resources",
			ShortDescription: `Provides a way to interact with Sumologic Content`,
			Description:      ``,
			Keywords: []string{
				"content",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_extraction_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic Field Extraction Rule`,
			Description:      ``,
			Keywords: []string{
				"extraction",
				"rule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_folder",
			Category:         "Resources",
			ShortDescription: `Provides the ability to create, read, delete, update, and manage of folders.`,
			Description:      ``,
			Keywords: []string{
				"folder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_http_source",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic HTTP source`,
			Description:      ``,
			Keywords: []string{
				"http",
				"source",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_ingest_budget",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic Ingest Budget`,
			Description:      ``,
			Keywords: []string{
				"ingest",
				"budget",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Display name of the ingest budget. This must be unique across all of the ingest budgets`,
				},
				resource.Attribute{
					Name:        "field_value",
					Description: `(Required) Custom field value that is used to assign Collectors to the ingest budget.`,
				},
				resource.Attribute{
					Name:        "capacity_bytes",
					Description: `(Required) Capacity of the ingest budget, in bytes.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the collector.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) The time zone to use for this collector. The value follows the [tzdata][2] naming convention. Defaults to ` + "`" + `Etc/UTC` + "`" + ``,
				},
				resource.Attribute{
					Name:        "reset_time",
					Description: `(Optional) Reset time of the ingest budget in HH:MM format. Defaults to ` + "`" + `00:00` + "`" + ``,
				},
				resource.Attribute{
					Name:        "reset_time",
					Description: `(Optional) Reset time of the ingest budget in HH:MM format. Defaults to ` + "`" + `00:00` + "`" + ``,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the ingest budget.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `(Optional) Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are ` + "`" + `stopCollecting` + "`" + ` and ` + "`" + `keepCollecting` + "`" + `. The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The internal ID of the ingest budget. This can be used to assign collectors to the ingest budget. ## Import Ingest budgets can be imported using the name, e.g.: ` + "`" + `` + "`" + `` + "`" + `hcl terraform import sumologic_ingest_budget.budget budgetName ` + "`" + `` + "`" + `` + "`" + ` [1]: https://help.sumologic.com/Manage/Ingestion-and-Volume/Ingest_Budgets [2]: https://en.wikipedia.org/wiki/Tz_database`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_partition",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic Partition`,
			Description:      ``,
			Keywords: []string{
				"partition",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_polling_source",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic Polling source. This source is used to import data from AWS S3 buckets.`,
			Description:      ``,
			Keywords: []string{
				"polling",
				"source",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_role",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic Role`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_scheduled_view",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic Scheduled View`,
			Description:      ``,
			Keywords: []string{
				"scheduled",
				"view",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sumologic_user",
			Category:         "Resources",
			ShortDescription: `Provides a Sumologic User`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"sumologic_cloudsyslog_source": 0,
		"sumologic_collector":          1,
		"sumologic_connection":         2,
		"sumologic_content":            3,
		"sumologic_extraction_rule":    4,
		"sumologic_folder":             5,
		"sumologic_http_source":        6,
		"sumologic_ingest_budget":      7,
		"sumologic_partition":          8,
		"sumologic_polling_source":     9,
		"sumologic_role":               10,
		"sumologic_scheduled_view":     11,
		"sumologic_user":               12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
