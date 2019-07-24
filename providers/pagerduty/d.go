package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_escalation_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Escalation Policy. This data source can be helpful when an escalation policy is handled outside Terraform but you still want to reference it in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Type:             "pagerduty_schedule",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Schedule. This data source can be helpful when a schedule is handled outside Terraform but you still want to reference it in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Type:             "pagerduty_team",
			Category:         "Data Sources",
			ShortDescription: `Get information about a team that you can use with escalation_policies, schedules etc.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{
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

	dataSourcesMap = map[string]Resource{

		"pagerduty_escalation_policy": 0,
		"pagerduty_extension_schema":  1,
		"pagerduty_schedule":          2,
		"pagerduty_team":              3,
		"pagerduty_user":              4,
		"pagerduty_vendor":            5,
	}
)

func GetDataSource(r string) (*resouce.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs]
}
