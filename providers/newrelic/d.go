package newrelic

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_channel",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about an alert channel in New Relic.`,
			Description: `\_alert\_channel

Use this data source to get information about a specific alert channel in New Relic that already exists.  More information on Terraform's data sources can be found [here](https://www.terraform.io/docs/configuration/data-sources.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert channel in New Relic. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert channel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Alert channel type, either: ` + "`" + `email` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `victorops` + "`" + `, or ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_ids",
					Description: `A list of policy IDs associated with the alert channel.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert channel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Alert channel type, either: ` + "`" + `email` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `victorops` + "`" + `, or ` + "`" + `webhook` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_ids",
					Description: `A list of policy IDs associated with the alert channel.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_alert_policy",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about an alert policy in New Relic.`,
			Description: `\_alert\_policy

Use this data source to get information about a specific alert policy in New Relic that already exists.
More information on Terraform's data sources can be found [here](https://www.terraform.io/docs/configuration/data-sources.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert policy in New Relic. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert policy.`,
				},
				resource.Attribute{
					Name:        "incident_preference",
					Description: `The rollup strategy for the policy. Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET. The default is PER_POLICY.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the policy was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert policy.`,
				},
				resource.Attribute{
					Name:        "incident_preference",
					Description: `The rollup strategy for the policy. Options include: PER_POLICY, PER_CONDITION, or PER_CONDITION_AND_TARGET. The default is PER_POLICY.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time the policy was created.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `The time the policy was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_application",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about an application in New Relic.`,
			Description: `\_application

Use this data source to get information about a specific application in New Relic that already exists. More information on Terraform's data sources can be found [here](https://www.terraform.io/docs/configuration/data-sources.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application in New Relic. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the application.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `A list of instance IDs associated with the application.`,
				},
				resource.Attribute{
					Name:        "host_ids",
					Description: `A list of host IDs associated with the application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the application.`,
				},
				resource.Attribute{
					Name:        "instance_ids",
					Description: `A list of instance IDs associated with the application.`,
				},
				resource.Attribute{
					Name:        "host_ids",
					Description: `A list of host IDs associated with the application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_key_transaction",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about a key transaction in New Relic.`,
			Description: `\_key\_transaction

Use this data source to get information about a specific key transaction in New Relic that already exists.  More information on Terraform's data sources can be found [here](https://www.terraform.io/docs/configuration/data-sources.html).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application in New Relic. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the application.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the application.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_plugin",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about a plugin in New Relic.`,
			Description: `\_plugin

Use this data source to get information about a specific installed plugin in New Relic. More information on Terraform's data sources can be found [here](https://www.terraform.io/docs/configuration/data-sources.html).

Each plugin published to New Relic's Plugin Central is assigned a [GUID](https://docs.newrelic.com/docs/plugins/plugin-developer-resources/planning-your-plugin/parts-plugin#guid). Once you have installed a plugin into your account it is assigned an ID. This account-specific ID is required when creating Plugins alert conditions.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "guid",
					Description: `(Required) The GUID of the plugin in New Relic. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the installed plugin instance.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the installed plugin instance.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_plugin_component",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about a plugin component in New Relic.`,
			Description: `\_plugin\_component

Use this data source to get information about a single plugin component in New Relic that already exists.
More information on Terraform's data sources can be found [here](https://www.terraform.io/docs/configuration/data-sources.html).

Each plugin component reporting into to New Relic is assigned a unique ID. Once you have a plugin component reporting data into your account, its component ID can be used to create Plugins alert conditions.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plugin_id",
					Description: `(Required) The ID of the plugin instance this component belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the plugin component. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the plugin component.`,
				},
				resource.Attribute{
					Name:        "health_status",
					Description: `The health status of the plugin component.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the plugin component.`,
				},
				resource.Attribute{
					Name:        "health_status",
					Description: `The health status of the plugin component.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "newrelic_synthetics_monitor",
			Category:         "Data Sources",
			ShortDescription: `Grabs a synthetics monitor by name.`,
			Description: `\_synthetics\_monitor

Use this data source to get information about a specific synthetics monitor in New Relic that already exists. This can be used to set up a Synthetics alert condition.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the synthetics monitor in New Relic. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "monitor_id",
					Description: `The ID of the synthetics monitor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "monitor_id",
					Description: `The ID of the synthetics monitor.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"newrelic_alert_channel":      0,
		"newrelic_alert_policy":       1,
		"newrelic_application":        2,
		"newrelic_key_transaction":    3,
		"newrelic_plugin":             4,
		"newrelic_plugin_component":   5,
		"newrelic_synthetics_monitor": 6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
