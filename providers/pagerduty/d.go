package pagerduty

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_escalation_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Escalation Policy. This data source can be helpful when an escalation policy is handled outside Terraform but you still want to reference it in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to use to find an escalation policy in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found escalation policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found escalation policy. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Escalation_Policies/get_escalation_policies`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found escalation policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found escalation policy. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Escalation_Policies/get_escalation_policies`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_extension_schema",
			Category:         "Data Sources",
			ShortDescription: `Get information about an extension vendor that you can use for a service (e.g: Slack, Generic Webhook, ServiceNow).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The extension name to use to find an extension vendor in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this extension vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Extension_Schemas/get_extension_schemas`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found extension vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this extension vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Extension_Schemas/get_extension_schemas`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_priority",
			Category:         "Data Sources",
			ShortDescription: `Get information about a priority that you can use with ruleset_rules, etc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the priority to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found priority.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found priority.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found priority. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1priorities/get`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found priority.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found priority.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found priority. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1priorities/get`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_ruleset",
			Category:         "Data Sources",
			ShortDescription: `Get information about a ruleset that you have created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ruleset to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found ruleset. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets/get [2]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets~1%7Bid%7D~1rules/get`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found ruleset. [1]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets/get [2]: https://developer.pagerduty.com/api-reference/reference/REST/openapiv3.json/paths/~1rulesets~1%7Bid%7D~1rules/get`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_schedule",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Schedule. This data source can be helpful when a schedule is handled outside Terraform but you still want to reference it in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to use to find a schedule in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found schedule. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Schedules/get_schedules`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found schedule. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Schedules/get_schedules`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service",
			Category:         "Data Sources",
			ShortDescription: `Get information about a service that you have created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The service name to use to find a service in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found service. [1]: https://api-reference.pagerduty.com/#!/Services/get_services`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found service. [1]: https://api-reference.pagerduty.com/#!/Services/get_services`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_team",
			Category:         "Data Sources",
			ShortDescription: `Get information about a team that you can use with escalation_policies, schedules etc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the team to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found team. [1]: https://v1.developer.pagerduty.com/documentation/rest/teams/list`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found team.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found team.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the found team. [1]: https://v1.developer.pagerduty.com/documentation/rest/teams/list`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user",
			Category:         "Data Sources",
			ShortDescription: `Get information about a user that you can use for a service integration (e.g Amazon Cloudwatch, Splunk, Datadog).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The email to use to find a user in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found user. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Users/get_users`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found user. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Users/get_users`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_vendor",
			Category:         "Data Sources",
			ShortDescription: `Get information about a vendor that you can use for a service integration (e.g Amazon Cloudwatch, Splunk, Datadog).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The vendor name to use to find a vendor in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Vendors/get_vendors`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found vendor.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found vendor.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The generic service type for this vendor. [1]: https://v2.developer.pagerduty.com/v2/page/api-reference#!/Vendors/get_vendors`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"pagerduty_escalation_policy": 0,
		"pagerduty_extension_schema":  1,
		"pagerduty_priority":          2,
		"pagerduty_ruleset":           3,
		"pagerduty_schedule":          4,
		"pagerduty_service":           5,
		"pagerduty_team":              6,
		"pagerduty_user":              7,
		"pagerduty_vendor":            8,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
