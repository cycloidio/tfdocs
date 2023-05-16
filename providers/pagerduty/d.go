package pagerduty

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_business_service",
			Category:         "Data Sources",
			ShortDescription: `Get information about a business service that you have created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The business service name to use to find a business service in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found business service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found business service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of object. The value returned will be ` + "`" + `business_service` + "`" + `. Can be used for passing to a service dependency. [1]: https://api-reference.pagerduty.com/#!/Business_Services/get_business_services`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found business service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found business service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of object. The value returned will be ` + "`" + `business_service` + "`" + `. Can be used for passing to a service dependency. [1]: https://api-reference.pagerduty.com/#!/Business_Services/get_business_services`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_escalation_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Escalation Policy. This data source can be helpful when an escalation policy is handled outside Terraform, but you still want to reference it in other resources.`,
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
					Description: `The short name of the found escalation policy. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODEyNA-list-escalation-policies`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found escalation policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found escalation policy. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODEyNA-list-escalation-policies`,
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
					Description: `The generic service type for this extension vendor. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODEzMA-list-extension-schemas`,
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
					Description: `The generic service type for this extension vendor. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODEzMA-list-extension-schemas`,
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
					Description: `A description of the found priority. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE2NA-list-priorities`,
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
					Description: `A description of the found priority. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE2NA-list-priorities`,
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
					Description: `The name of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "routing_keys",
					Description: `Routing keys routed to this ruleset. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE3MQ-list-rulesets [2]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE3Ng-list-event-rules`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the found ruleset.`,
				},
				resource.Attribute{
					Name:        "routing_keys",
					Description: `Routing keys routed to this ruleset. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE3MQ-list-rulesets [2]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE3Ng-list-event-rules`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_schedule",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a Schedule. This data source can be helpful when a schedule is handled outside Terraform, but you still want to reference it in other resources.`,
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
					Description: `The short name of the found schedule. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE4MQ-list-schedules`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found schedule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found schedule. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODE4MQ-list-schedules`,
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
					Description: `The short name of the found service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of object. The value returned will be ` + "`" + `service` + "`" + `. Can be used for passing to a service dependency.`,
				},
				resource.Attribute{
					Name:        "auto_resolve_timeout",
					Description: `Time in seconds that an incident is automatically resolved if left open for that long. Value is null if the feature is disabled. Value must not be negative. Setting this field to 0, null (or unset) will disable the feature.`,
				},
				resource.Attribute{
					Name:        "acknowledgement_timeout",
					Description: `Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value is null if the feature is disabled. Value must not be negative. Setting this field to 0, null (or unset) will disable the feature.`,
				},
				resource.Attribute{
					Name:        "alert_creation",
					Description: `Whether a service creates only incidents, or both alerts and incidents. A service must create alerts in order to enable incident merging.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The user-provided description of the service.`,
				},
				resource.Attribute{
					Name:        "escalation_policy",
					Description: `The escalation policy associated with this service.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `The set of teams associated with the service. [1]: https://api-reference.pagerduty.com/#!/Services/get_services`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found service.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found service.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of object. The value returned will be ` + "`" + `service` + "`" + `. Can be used for passing to a service dependency.`,
				},
				resource.Attribute{
					Name:        "auto_resolve_timeout",
					Description: `Time in seconds that an incident is automatically resolved if left open for that long. Value is null if the feature is disabled. Value must not be negative. Setting this field to 0, null (or unset) will disable the feature.`,
				},
				resource.Attribute{
					Name:        "acknowledgement_timeout",
					Description: `Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value is null if the feature is disabled. Value must not be negative. Setting this field to 0, null (or unset) will disable the feature.`,
				},
				resource.Attribute{
					Name:        "alert_creation",
					Description: `Whether a service creates only incidents, or both alerts and incidents. A service must create alerts in order to enable incident merging.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The user-provided description of the service.`,
				},
				resource.Attribute{
					Name:        "escalation_policy",
					Description: `The escalation policy associated with this service.`,
				},
				resource.Attribute{
					Name:        "teams",
					Description: `The set of teams associated with the service. [1]: https://api-reference.pagerduty.com/#!/Services/get_services`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_service_integration",
			Category:         "Data Sources",
			ShortDescription: `Get information about a service integration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_name",
					Description: `(Required) The service name to use to find a service in the PagerDuty API.`,
				},
				resource.Attribute{
					Name:        "integration_summary",
					Description: `(Required) The integration summary used to find the desired integration on the service. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "integration_key",
					Description: `The integration key for the integration. This can be used to configure alerts.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "integration_key",
					Description: `The integration key for the integration. This can be used to configure alerts.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_tag",
			Category:         "Data Sources",
			ShortDescription: `Get information about a tag that you can use to assign to users, teams, and escalation_policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label of the tag to find in the PagerDuty API. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found team. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIxNw-list-tags`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found team. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIxNw-list-tags`,
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
					Description: `A description of the found team.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `ID of the parent team. This is available to accounts with the Team Hierarchy feature enabled. Please contact your account manager for more information. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIyMw-list-teams`,
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
					Description: `A description of the found team.`,
				},
				resource.Attribute{
					Name:        "parent",
					Description: `ID of the parent team. This is available to accounts with the Team Hierarchy feature enabled. Please contact your account manager for more information. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIyMw-list-teams`,
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
					Description: `The short name of the found user. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzMw-list-users`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The short name of the found user. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzMw-list-users`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_user_contact_method",
			Category:         "Data Sources",
			ShortDescription: `Get information about a contact method of a PagerDuty user (email, phone, SMS or push notification).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "user_id",
					Description: `(Required) The ID of the user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The contact method type. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label (e.g., "Work", "Mobile", "Ashley's iPhone", etc.). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the found contact method. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_short_email",
					Description: `Send an abbreviated email message instead of the standard email output. (Email contact method only.)`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The 1-to-3 digit country calling code. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label (e.g., "Work", "Mobile", "Ashley's iPhone", etc.).`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The "address" to deliver to: ` + "`" + `email` + "`" + `, ` + "`" + `phone number` + "`" + `, etc., depending on the type.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Either ` + "`" + `ios` + "`" + ` or ` + "`" + `android` + "`" + `, depending on the type of the device receiving notifications. (Push notification contact method only.) [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzOQ-list-a-user-s-contact-methods [2]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzMw-list-users`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the found user.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of the found contact method. May be (` + "`" + `email_contact_method` + "`" + `, ` + "`" + `phone_contact_method` + "`" + `, ` + "`" + `sms_contact_method` + "`" + `, ` + "`" + `push_notification_contact_method` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "send_short_email",
					Description: `Send an abbreviated email message instead of the standard email output. (Email contact method only.)`,
				},
				resource.Attribute{
					Name:        "country_code",
					Description: `The 1-to-3 digit country calling code. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `The label (e.g., "Work", "Mobile", "Ashley's iPhone", etc.).`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The "address" to deliver to: ` + "`" + `email` + "`" + `, ` + "`" + `phone number` + "`" + `, etc., depending on the type.`,
				},
				resource.Attribute{
					Name:        "blacklisted",
					Description: `If true, this phone has been blacklisted by PagerDuty and no messages will be sent to it. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `If true, this phone is capable of receiving SMS messages. (Phone and SMS contact methods only.)`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Either ` + "`" + `ios` + "`" + ` or ` + "`" + `android` + "`" + `, depending on the type of the device receiving notifications. (Push notification contact method only.) [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzOQ-list-a-user-s-contact-methods [2]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODIzMw-list-users`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "pagerduty_vendor",
			Category:         "Data Sources",
			ShortDescription: `Get information about a vendor that you can use for a service integration (e.g. Amazon Cloudwatch, Splunk, Datadog).`,
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
					Description: `The generic service type for this vendor. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODI1OQ-list-vendors [2]: https://support.pagerduty.com/docs/change-events [3]: https://support.pagerduty.com/docs/aiops`,
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
					Description: `The generic service type for this vendor. [1]: https://developer.pagerduty.com/api-reference/b3A6Mjc0ODI1OQ-list-vendors [2]: https://support.pagerduty.com/docs/change-events [3]: https://support.pagerduty.com/docs/aiops`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"pagerduty_business_service":    0,
		"pagerduty_escalation_policy":   1,
		"pagerduty_extension_schema":    2,
		"pagerduty_priority":            3,
		"pagerduty_ruleset":             4,
		"pagerduty_schedule":            5,
		"pagerduty_service":             6,
		"pagerduty_service_integration": 7,
		"pagerduty_tag":                 8,
		"pagerduty_team":                9,
		"pagerduty_user":                10,
		"pagerduty_user_contact_method": 11,
		"pagerduty_vendor":              12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
