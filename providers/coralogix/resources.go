package coralogix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
					Name:        "name",
					Description: `(Required) Alert name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Alert type, one of the following: ` + "`" + `text` + "`" + `, ` + "`" + `ratio` + "`" + `, ` + "`" + `unique_count` + "`" + `, ` + "`" + `relative_time` + "`" + `, ` + "`" + `metric` + "`" + `. For ` + "`" + `new_value` + "`" + ` alerts the value should be ` + "`" + `text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Required) Alert severity, one of the following: ` + "`" + `info` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `critical` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Required) Alert state.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) A ` + "`" + `filter` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Alert description.`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) A ` + "`" + `metric` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) A ` + "`" + `condition` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "schedule",
					Description: `(Optional) A ` + "`" + `schedule` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) An array that contains log fields to be included with the alert notification.`,
				},
				resource.Attribute{
					Name:        "notifications",
					Description: `(Optional) A ` + "`" + `notifications` + "`" + ` block as documented below. --- Each ` + "`" + `filter` + "`" + ` block should contains the following:`,
				},
				resource.Attribute{
					Name:        "text",
					Description: `(Optional) String query to be alerted on.`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `(Optional) List of application names to be alerted on.`,
				},
				resource.Attribute{
					Name:        "subsystems",
					Description: `(Optional) List of subsystem names to be alerted on.`,
				},
				resource.Attribute{
					Name:        "severities",
					Description: `(Optional) List of log severities to be alerted on, one of the following: ` + "`" + `debug` + "`" + `, ` + "`" + `verbose` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `critical` + "`" + `. Each ` + "`" + `metric` + "`" + ` block should contains the following:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) The name of the metric field to alert on.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The source of the metric. Either ` + "`" + `logs2metrics` + "`" + ` or ` + "`" + `Prometheus` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "arithmetic_operator",
					Description: `(Required) ` + "`" + `0` + "`" + ` - avg, ` + "`" + `1` + "`" + ` - min, ` + "`" + `2` + "`" + ` - max, ` + "`" + `3` + "`" + ` - sum, ` + "`" + `4` + "`" + ` - count, ` + "`" + `5` + "`" + ` - percentile (for percentile you need to supply the requested percentile in arithmetic_operator_modifier).`,
				},
				resource.Attribute{
					Name:        "arithmetic_operator_modifier",
					Description: `(Optional) For ` + "`" + `percentile(5)` + "`" + ` ` + "`" + `arithmetic_operator` + "`" + ` you need to supply the value in this property.`,
				},
				resource.Attribute{
					Name:        "sample_threshold_percentage",
					Description: `(Required) The metric value must cross the threshold within this percentage of the timeframe (sum and count arithmetic operators do not use this parameter since they aggregate over the entire requested timeframe), ` + "`" + `increments of 10` + "`" + `, ` + "`" + `0 <= value <= 90` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "non_null_percentage",
					Description: `(Required) The minimum percentage of the timeframe that should have values for this alert to trigger, ` + "`" + `increments of 10` + "`" + `, ` + "`" + `0 <= value <= 100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "swap_null_values",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, missing data will be considered as 0, otherwise, it will not be considered at all. Each ` + "`" + `condition` + "`" + ` block should contains the following:`,
				},
				resource.Attribute{
					Name:        "condition_type",
					Description: `(Required) Alert condition type, one of the following: ` + "`" + `less_than` + "`" + `, ` + "`" + `more_than` + "`" + `, ` + "`" + `more_than_usual` + "`" + `, ` + "`" + `new_value` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "threshold",
					Description: `(Required) Number of log occurrences that is needed to trigger the alert.`,
				},
				resource.Attribute{
					Name:        "timeframe",
					Description: `(Required) The bounded time frame for the threshold to be occurred within, to trigger the alert.`,
				},
				resource.Attribute{
					Name:        "group_by",
					Description: `(Optional) The field to`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `(Required) Days when alert triggering is allowed, one of the following: ` + "`" + `Mo` + "`" + `, ` + "`" + `Tu` + "`" + `, ` + "`" + `We` + "`" + `, ` + "`" + `Th` + "`" + `, ` + "`" + `Fr` + "`" + `, ` + "`" + `Sa` + "`" + `, ` + "`" + `Su` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Required) Time from which alert triggering is allowed, for example ` + "`" + `00:00:00` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) Time till which alert triggering is allowed, for example ` + "`" + `23:59:59` + "`" + `. Each ` + "`" + `notifications` + "`" + ` block should contains the following:`,
				},
				resource.Attribute{
					Name:        "emails",
					Description: `(Optional) List of email address to notify.`,
				},
				resource.Attribute{
					Name:        "integrations",
					Description: `(Optional) List of integration channels to notify. ## Import Alerts can be imported using their ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import coralogix_alert.alert <alert_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
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
					Name:        "rules_group_id",
					Description: `(Required) Rules Group ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Rule name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Rule type, one of the following: ` + "`" + `extract` + "`" + `, ` + "`" + `jsonextract` + "`" + `, ` + "`" + `parse` + "`" + `, ` + "`" + `replace` + "`" + `, ` + "`" + `allow` + "`" + `, ` + "`" + `block` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Rule description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Rule state.`,
				},
				resource.Attribute{
					Name:        "rule_matcher",
					Description: `(Optional) A ` + "`" + `rule_matcher` + "`" + ` block as documented below.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) Rule expression. Should be valid regular expression.`,
				},
				resource.Attribute{
					Name:        "source_field",
					Description: `(Optional) Rule source field.`,
				},
				resource.Attribute{
					Name:        "destination_field",
					Description: `(Optional) Rule destination field.`,
				},
				resource.Attribute{
					Name:        "replace_value",
					Description: `(Optional) Rule replace value. --- Each ` + "`" + `rule_matcher` + "`" + ` block should contains the following:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `(Required) Rule Matcher field.`,
				},
				resource.Attribute{
					Name:        "constraint",
					Description: `(Required) Rule Matcher constraint. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Rule order number. ## Import Rules can be imported using their ID and rules group ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import coralogix_rule.rule <rules_group_id>/<rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "order",
					Description: `Rule order number. ## Import Rules can be imported using their ID and rules group ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import coralogix_rule.rule <rules_group_id>/<rule_id> ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "name",
					Description: `(Required) Rules Group name.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Rules Group state.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Rules Group description.`,
				},
				resource.Attribute{
					Name:        "creator",
					Description: `(Optional) Rules Group creator. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `Rules Group order number. ## Import Rules Groups can be imported using their ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import coralogix_rules_group.rules_group <rules_group_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "order",
					Description: `Rules Group order number. ## Import Rules Groups can be imported using their ID. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import coralogix_rules_group.rules_group <rules_group_id> ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"coralogix_alert":       0,
		"coralogix_rule":        1,
		"coralogix_rules_group": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
