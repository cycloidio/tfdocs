package datadog

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog dashboard resource. This can be used to create and manage Datadog dashboards.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard_json",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog dashboard JSON resource. This can be used to create and manage Datadog dashboards using the JSON definition.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_dashboard_list",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog dashboard_list resource. This can be used to create and manage Datadog Dashboard Lists and the individual dashboards within them.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_downtime",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog downtime resource. This can be used to create and manage Datadog downtimes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Amazon Web Services integration resource. This can be used to create and manage Datadog - Amazon Web Services integration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws_lambda_arn",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Amazon Web Services integration Lambda ARN resource. This can be used to create and manage the log collection Lambdas for an account. Update operations are currently not supported with datadog API so any change forces a new resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws_log_collection",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Amazon Web Services integration log collection resource. This can be used to manage which AWS services logs are collected from for an account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_aws_tag_filter",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog AWS tag filter resource. This can be used to create and manage Datadog AWS tag filters - US site’s endpoint only`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_azure",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Microsoft Azure integration resource. This can be used to create and manage the integrations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_gcp",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - Google Cloud Platform integration resource. This can be used to create and manage Datadog - Google Cloud Platform integration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_pagerduty",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog - PagerDuty resource. This can be used to create and manage Datadog - PagerDuty integration. See also PagerDuty Integration Guide https://www.pagerduty.com/docs/guides/datadog-integration-guide/.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_pagerduty_service_object",
			Category:         "Resources",
			ShortDescription: `Provides access to individual Service Objects of Datadog - PagerDuty integrations. Note that the Datadog - PagerDuty integration must be activated in the Datadog UI in order for this resource to be usable.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_slack_channel",
			Category:         "Resources",
			ShortDescription: `Resource for interacting with the Datadog Slack channel API`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_archive",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Archive API resource, which is used to create and manage Datadog logs archives.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_archive_order",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Archive API https://docs.datadoghq.com/api/v2/logs-archives/ resource, which is used to manage Datadog log archives order.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_custom_pipeline",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Pipeline API https://docs.datadoghq.com/api/v1/logs-pipelines/ resource, which is used to create and manage Datadog logs custom pipelines. Each datadog_logs_custom_pipeline resource defines a complete pipeline. The order of the pipelines is maintained in a different resource: datadog_logs_pipeline_order. When creating a new pipeline, you need to explicitly add this pipeline to the datadog_logs_pipeline_order resource to track the pipeline. Similarly, when a pipeline needs to be destroyed, remove its references from the datadog_logs_pipeline_order resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_index",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Index API resource. This can be used to create and manage Datadog logs indexes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_index_order",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Index API resource. This can be used to manage the order of Datadog logs indexes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_integration_pipeline",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Pipeline API resource to manage the integrations. Integration pipelines are the pipelines that are automatically installed for your organization when sending the logs with specific sources. You don't need to maintain or update these types of pipelines. Keeping them as resources, however, allows you to manage the order of your pipelines by referencing them in your datadog_logs_pipeline_order resource. If you don't need the pipeline_order feature, this resource declaration can be omitted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_metric",
			Category:         "Resources",
			ShortDescription: `Resource for interacting with the logs_metric API`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_logs_pipeline_order",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Logs Pipeline API resource, which is used to manage Datadog log pipelines order.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_metric_metadata",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog metric_metadata resource. This can be used to manage a metric's metadata.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_metric_tag_configuration",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog metric tag configuration resource. This can be used to modify tag configurations for metrics.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_monitor",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog monitor resource. This can be used to create and manage Datadog monitors.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_role",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog role resource. This can be used to create and manage Datadog roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_screenboard",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_security_monitoring_default_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Security Monitoring Rule API resource for default rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_security_monitoring_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Security Monitoring Rule API resource. This can be used to create and manage Datadog security monitoring rules. To change settings for a default rule use datadog_security_default_rule instead.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_service_level_objective",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog service level objective resource. This can be used to create and manage Datadog service level objectives.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_synthetics_global_variable",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog synthetics global variable resource. This can be used to create and manage Datadog synthetics global variables.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_synthetics_private_location",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog synthetics private location resource. This can be used to create and manage Datadog synthetics private locations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_synthetics_test",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog synthetics test resource. This can be used to create and manage Datadog synthetics test.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_timeboard",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_user",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog user resource. This can be used to create and manage Datadog users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"datadog_dashboard":                            0,
		"datadog_dashboard_json":                       1,
		"datadog_dashboard_list":                       2,
		"datadog_downtime":                             3,
		"datadog_integration_aws":                      4,
		"datadog_integration_aws_lambda_arn":           5,
		"datadog_integration_aws_log_collection":       6,
		"datadog_integration_aws_tag_filter":           7,
		"datadog_integration_azure":                    8,
		"datadog_integration_gcp":                      9,
		"datadog_integration_pagerduty":                10,
		"datadog_integration_pagerduty_service_object": 11,
		"datadog_integration_slack_channel":            12,
		"datadog_logs_archive":                         13,
		"datadog_logs_archive_order":                   14,
		"datadog_logs_custom_pipeline":                 15,
		"datadog_logs_index":                           16,
		"datadog_logs_index_order":                     17,
		"datadog_logs_integration_pipeline":            18,
		"datadog_logs_metric":                          19,
		"datadog_logs_pipeline_order":                  20,
		"datadog_metric_metadata":                      21,
		"datadog_metric_tag_configuration":             22,
		"datadog_monitor":                              23,
		"datadog_role":                                 24,
		"datadog_screenboard":                          25,
		"datadog_security_monitoring_default_rule":     26,
		"datadog_security_monitoring_rule":             27,
		"datadog_service_level_objective":              28,
		"datadog_synthetics_global_variable":           29,
		"datadog_synthetics_private_location":          30,
		"datadog_synthetics_test":                      31,
		"datadog_timeboard":                            32,
		"datadog_user":                                 33,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}