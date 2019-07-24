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
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert channel in New Relic. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the alert channel.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Alert channel type, either: ` + "`" + `campfire` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `hipchat` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `victorops` + "`" + `, or ` + "`" + `webhook` + "`" + `..`,
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
					Description: `Alert channel type, either: ` + "`" + `campfire` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `hipchat` + "`" + `, ` + "`" + `opsgenie` + "`" + `, ` + "`" + `pagerduty` + "`" + `, ` + "`" + `slack` + "`" + `, ` + "`" + `victorops` + "`" + `, or ` + "`" + `webhook` + "`" + `..`,
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
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert policy in New Relic. ## Attributes Reference`,
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
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application in New Relic. ## Attributes Reference`,
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
			Type:             "key_transaction",
			Category:         "Data Sources",
			ShortDescription: `Looks up the information about a key transaction in New Relic.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application in New Relic. ## Attributes Reference`,
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
			Type:             "synthetics_monitor",
			Category:         "Data Sources",
			ShortDescription: `Grabs a synthetics monitor by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the synthetics monitor in New Relic. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the synthetics monitor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the synthetics monitor.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"newrelic_alert_channel": 0,
		"newrelic_alert_policy":  1,
		"newrelic_application":   2,
		"key_transaction":        3,
		"synthetics_monitor":     4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
