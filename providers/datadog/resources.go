package datadog

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "datadog_api_key",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog API Key resource. This can be used to create and manage Datadog API Keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_application_key",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Application Key resource. This can be used to create and manage Datadog Application Keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_authn_mapping",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog AuthN Mappings resource. This feature lets you automatically assign roles to users based on their SAML attributes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_child_organization",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Child Organization resource. This can be used to create Datadog Child Organizations. To manage created organization use datadog_organization_settings.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_cloud_configuration_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Cloud Configuration Rule resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_cloud_workload_security_agent_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Cloud Workload Security Agent Rule API resource for agent rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
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
			ShortDescription: `Provides a Datadog AWS tag filter resource. This can be used to create and manage Datadog AWS tag filters.`,
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
			Type:             "datadog_integration_cloudflare_account",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog IntegrationCloudflareAccount resource. This can be used to create and manage Datadog integrationcloudflareaccount.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_confluent_account",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog IntegrationConfluentAccount resource. This can be used to create and manage Datadog integrationconfluentaccount.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_confluent_resource",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog IntegrationConfluentResource resource. This can be used to create and manage Datadog integrationconfluentresource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_fastly_account",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog IntegrationFastlyAccount resource. This can be used to create and manage Datadog integrationfastlyaccount.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_integration_fastly_service",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog IntegrationFastlyService resource. This can be used to create and manage Datadog integrationfastlyservice.`,
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
			Type:             "datadog_integration_opsgenie_service_object",
			Category:         "Resources",
			ShortDescription: `Resource for interacting with Datadog Opsgenie Service API.`,
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
			Type:             "datadog_ip_allowlist",
			Category:         "Resources",
			ShortDescription: `Provides the Datadog IP allowlist resource. This can be used to manage the Datadog IP allowlist`,
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
			ShortDescription: `Provides a Datadog Logs Index API resource. This can be used to create and manage Datadog logs indexes.Note: It is not possible to delete logs indexes through Terraform, so an index remains in your account after the resource is removed from your terraform config. Reach out to support to delete a logs index.`,
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
			Type:             "datadog_monitor_config_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog monitor config policy resource. This can be used to create and manage Datadog monitor config policies.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_monitor_json",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog monitor JSON resource. This can be used to create and manage Datadog monitors using the JSON definition.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_organization_settings",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Organization resource. This can be used to manage your Datadog organization's settings.`,
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
			Type:             "datadog_rum_application",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog RUM application resource. This can be used to create and manage Datadog RUM applications.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_security_monitoring_default_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Security Monitoring Rule API resource for default rules. It can only be imported, you can't create a default rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_security_monitoring_filter",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Security Monitoring Rule API resource for security filters.`,
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
			Type:             "datadog_sensitive_data_scanner_group",
			Category:         "Resources",
			ShortDescription: `Provides a Sensitive Data Scanner group resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_sensitive_data_scanner_group_order",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog Sensitive Data Scanner Group Order API resource. This can be used to manage the order of Datadog Sensitive Data Scanner Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_sensitive_data_scanner_rule",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog SensitiveDataScannerRule resource. This can be used to create and manage Datadog sensitivedatascanner_rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_service_account",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog service account resource. This can be used to create and manage Datadog service accounts.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_service_definition_yaml",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog service definition resource. This can be used to create and manage Datadog service definitions in the service catalog using the YAML/JSON definition.`,
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
			Type:             "datadog_slo_correction",
			Category:         "Resources",
			ShortDescription: `Resource for interacting with the slo_correction API.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_spans_metric",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog SpansMetric resource. This can be used to create and manage Datadog spans_metric.`,
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
			Type:             "datadog_user",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog user resource. This can be used to create and manage Datadog users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_webhook",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog webhook resource. This can be used to create and manage Datadog webhooks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "datadog_webhook_custom_variable",
			Category:         "Resources",
			ShortDescription: `Provides a Datadog webhooks custom variable resource. This can be used to create and manage Datadog webhooks custom variables.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"datadog_api_key":                              0,
		"datadog_application_key":                      1,
		"datadog_authn_mapping":                        2,
		"datadog_child_organization":                   3,
		"datadog_cloud_configuration_rule":             4,
		"datadog_cloud_workload_security_agent_rule":   5,
		"datadog_dashboard":                            6,
		"datadog_dashboard_json":                       7,
		"datadog_dashboard_list":                       8,
		"datadog_downtime":                             9,
		"datadog_integration_aws":                      10,
		"datadog_integration_aws_lambda_arn":           11,
		"datadog_integration_aws_log_collection":       12,
		"datadog_integration_aws_tag_filter":           13,
		"datadog_integration_azure":                    14,
		"datadog_integration_cloudflare_account":       15,
		"datadog_integration_confluent_account":        16,
		"datadog_integration_confluent_resource":       17,
		"datadog_integration_fastly_account":           18,
		"datadog_integration_fastly_service":           19,
		"datadog_integration_gcp":                      20,
		"datadog_integration_opsgenie_service_object":  21,
		"datadog_integration_pagerduty":                22,
		"datadog_integration_pagerduty_service_object": 23,
		"datadog_integration_slack_channel":            24,
		"datadog_ip_allowlist":                         25,
		"datadog_logs_archive":                         26,
		"datadog_logs_archive_order":                   27,
		"datadog_logs_custom_pipeline":                 28,
		"datadog_logs_index":                           29,
		"datadog_logs_index_order":                     30,
		"datadog_logs_integration_pipeline":            31,
		"datadog_logs_metric":                          32,
		"datadog_logs_pipeline_order":                  33,
		"datadog_metric_metadata":                      34,
		"datadog_metric_tag_configuration":             35,
		"datadog_monitor":                              36,
		"datadog_monitor_config_policy":                37,
		"datadog_monitor_json":                         38,
		"datadog_organization_settings":                39,
		"datadog_role":                                 40,
		"datadog_rum_application":                      41,
		"datadog_security_monitoring_default_rule":     42,
		"datadog_security_monitoring_filter":           43,
		"datadog_security_monitoring_rule":             44,
		"datadog_sensitive_data_scanner_group":         45,
		"datadog_sensitive_data_scanner_group_order":   46,
		"datadog_sensitive_data_scanner_rule":          47,
		"datadog_service_account":                      48,
		"datadog_service_definition_yaml":              49,
		"datadog_service_level_objective":              50,
		"datadog_slo_correction":                       51,
		"datadog_spans_metric":                         52,
		"datadog_synthetics_global_variable":           53,
		"datadog_synthetics_private_location":          54,
		"datadog_synthetics_test":                      55,
		"datadog_user":                                 56,
		"datadog_webhook":                              57,
		"datadog_webhook_custom_variable":              58,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
