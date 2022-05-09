package coralogix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "coralogix_alert",
			Category:         "Alerts",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"alerts",
				"alert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alert_id",
					Description: `(Required) Alert ID. ## Attribute Reference The result is an object containing the following attributes.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Alert name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Alert description.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Alert severity.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Alert state.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Alert type.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Alert filter.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `Alert condition.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `List of fields attached to alert notification.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Configuration of period when alert triggering will be allowed.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `Alert notifications.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Alert name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Alert description.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Alert severity.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Alert state.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Alert type.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Alert filter.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `Alert condition.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `List of fields attached to alert notification.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `Configuration of period when alert triggering will be allowed.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `Alert notifications.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_rule",
			Category:         "Rules",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"rules",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) Rule ID.`,
				},
				resource.Attribute{
					Name:        "rules_group_id",
					Description: `(Required) Rules Group ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Rule type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rule description.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Rule order number.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Rule state.`,
				},
				resource.Attribute{
					Name:        "rule_matcher",
					Description: `A ` + "`" + `rule_matcher` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `Rule expression.`,
				},
				resource.Attribute{
					Name:        "source_field",
					Description: `Rule source field.`,
				},
				resource.Attribute{
					Name:        "destination_field",
					Description: `Rule destination field.`,
				},
				resource.Attribute{
					Name:        "replace_value",
					Description: `Rule replace value. --- Each ` + "`" + `rule_matcher` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Rule Matcher field.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `Rule Matcher constraint.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Rule type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rule description.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Rule order number.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Rule state.`,
				},
				resource.Attribute{
					Name:        "rule_matcher",
					Description: `A ` + "`" + `rule_matcher` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `Rule expression.`,
				},
				resource.Attribute{
					Name:        "source_field",
					Description: `Rule source field.`,
				},
				resource.Attribute{
					Name:        "destination_field",
					Description: `Rule destination field.`,
				},
				resource.Attribute{
					Name:        "replace_value",
					Description: `Rule replace value. --- Each ` + "`" + `rule_matcher` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Rule Matcher field.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `Rule Matcher constraint.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "coralogix_rules_group",
			Category:         "Rules",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"rules",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "rules_group_id",
					Description: `(Required) Rules Group ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Rules Group name.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Rules Group order number.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Rules Group state.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rules Group description.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `Rules Group creator.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `Rules Group rules list.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Rules Group name.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Rules Group order number.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Rules Group state.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Rules Group description.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `Rules Group creator.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `Rules Group rules list.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"coralogix_alert":       0,
		"coralogix_rule":        1,
		"coralogix_rules_group": 2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
