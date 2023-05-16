package wavefront

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "wavefront_alert",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Alert resource. This allows alerts to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "alert_type",
					Description: `(Optional) The type of alert in Wavefront. Either ` + "`" + `CLASSIC` + "`" + ` (default) or ` + "`" + `THRESHOLD` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Required) The number of consecutive minutes that a series matching the condition query must evaluate to "true" (non-zero value) before the alert fires.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Optional) A comma-separated list of the email address or integration endpoint (such as PagerDuty or webhook) to notify when the alert status changes. Multiple target types can be in the list. Alert target format: ({email}|pd:{pd_key}|target:{alert-target-id}).`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `(Optional) A Wavefront query that is evaluated at regular intervals (default is 1 minute). The alert fires and notifications are triggered when a data series matching this query evaluates to a non-zero value for a set number of consecutive minutes.`,
				},
				resource.Attribute{
					Name:        "conditions",
					Description: `(Optional, ` + "`" + `THRESHOLD` + "`" + ` alerts only) a string->string map of ` + "`" + `severity` + "`" + ` to ` + "`" + `condition` + "`" + ` for which this alert will trigger.`,
				},
				resource.Attribute{
					Name:        "threshold_targets",
					Description: `(Optional, ` + "`" + `THRESHOLD` + "`" + ` alerts only) Targets for severity`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `(Optional) User-supplied additional explanatory information for this alert. Useful for linking runbooks, migrations, etc.`,
				},
				resource.Attribute{
					Name:        "display_expression",
					Description: `(Optional) A second query whose results are displayed in the alert user interface instead of the condition query. This field is often used to display a version of the condition query with Boolean operators removed so that numerical values are plotted.`,
				},
				resource.Attribute{
					Name:        "resolve_after_minutes",
					Description: `(Optional) The number of consecutive minutes that a firing series matching the condition query must evaluate to "false" (zero value) before the alert resolves. When unset, this defaults to the same value as ` + "`" + `minutes` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_resend_frequency_minutes",
					Description: `(Optional) How often to re-trigger a continually failing alert. If absent or <= 0, no re-triggering occurs.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `(Optional) - Severity of the alert, valid values are ` + "`" + `INFO` + "`" + `, ` + "`" + `SMOKE` + "`" + `, ` + "`" + `WARN` + "`" + `, ` + "`" + `SEVERE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `(Optional) A list of users or groups that can view this resource.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `(Optional) A list of users or groups that can modify this resource.`,
				},
				resource.Attribute{
					Name:        "process_rate_minutes",
					Description: `(Optional) The specified query is executed every ` + "`" + `process_rate_minutes` + "`" + ` minutes. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_alert" "test_alert" { name = "Terraform Test Alert" target = "test@example.com,target:alert-target-id" condition = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80" additional_information = "This is a Terraform Test Alert" display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )" minutes = 5 resolve_after_minutes = 5 severity = "WARN" tags = [ "terraform", "test" ] can_view = [ "test@example.com", ] process_rate_minutes = 4 } resource "wavefront_alert_target" "test_target" { name = "Terraform Test Target" description = "Test target" method = "EMAIL" recipient = "test@example.com" email_subject = "This is a test" is_html_content = true template = "{}" triggers = [ "ALERT_OPENED", "ALERT_RESOLVED" ] } resource "wavefront_alert" "test_threshold_alert" { name = "Terraform Test Alert" alert_type = "THRESHOLD" additional_information = "This is a Terraform Test Alert" display_expression = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total )" minutes = 5 resolve_after_minutes = 5 conditions = { "severe" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 80" "warn" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 60" "info" = "100-ts(\"cpu.usage_idle\", environment=preprod and cpu=cpu-total ) > 50" } threshold_targets = { "severe" = "target:${wavefront_alert_target.test_target.id}" } tags = [ "terraform" ] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Alerts can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_alert_target.alert_target 1479868728473 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_alert_target",
			Category:         "Resources",
			ShortDescription: `Provides a wavefront Alert Target resource. This allows alert targets to created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the alert target as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Description describing this alert target.`,
				},
				resource.Attribute{
					Name:        "triggers",
					Description: `(Required) A list of occurrences on which this webhook will be fired. Valid values are ` + "`" + `ALERT_OPENED` + "`" + `, ` + "`" + `ALERT_UPDATED` + "`" + `, ` + "`" + `ALERT_RESOLVED` + "`" + `, ` + "`" + `ALERT_MAINTENANCE` + "`" + `, ` + "`" + `ALERT_SNOOZED` + "`" + `, ` + "`" + `ALERT_NO_DATA` + "`" + `, ` + "`" + `ALERT_NO_DATA_RESOLVED` + "`" + `, ` + "`" + `ALERT_NO_DATA_MAINTENANCE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) A mustache template that will form the body of the POST request, email, and summary of the PagerDuty.`,
				},
				resource.Attribute{
					Name:        "recipient",
					Description: `(Required) The end point for the notification Target. ` + "`" + `EMAIL` + "`" + `: email address. ` + "`" + `PAGERDUTY` + "`" + `: PagerDuty routing key. ` + "`" + `WEBHOOK` + "`" + `: URL endpoint.`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Optional) The notification method used for notification target. One of ` + "`" + `WEBHOOK` + "`" + `, ` + "`" + `EMAIL` + "`" + `, ` + "`" + `PAGERDUTY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "route",
					Description: `(Optional) List of routing targets that this alert target will notify. See [Route](#route)`,
				},
				resource.Attribute{
					Name:        "email_subject",
					Description: `(Optional) The subject title of an email notification target.`,
				},
				resource.Attribute{
					Name:        "is_html_content",
					Description: `(Optional) Determine whether the email alert content is sent as HTML or text.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `(Optional) The value of the ` + "`" + `Content-Type` + "`" + ` header of the webhook.`,
				},
				resource.Attribute{
					Name:        "custom_headers",
					Description: `(Optional) A ` + "`" + `string->string` + "`" + ` map specifying the custom HTTP header key/value pairs that will be sent in the requests with a method of ` + "`" + `WEBHOOK` + "`" + `. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "target_id",
					Description: `The target ID prefixed with ` + "`" + `target:` + "`" + ` for interpolating into a Wavefront Alert. ### Route The ` + "`" + `route` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The notification method used for notification target. One of ` + "`" + `WEBHOOK` + "`" + `, ` + "`" + `EMAIL` + "`" + `, ` + "`" + `PAGERDUTY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The endpoint for the alert route. ` + "`" + `EMAIL` + "`" + `: email address. ` + "`" + `PAGERDUTY` + "`" + `: PagerDuty routing key. ` + "`" + `WEBHOOK` + "`" + `: URL endpoint.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) String that filters the route. Space delimited. Currently only allows a single key value pair. (e.g. ` + "`" + `env prod` + "`" + `) ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_alert_target" "test_target" { name = "Terraform Test Target" description = "Test target" method = "EMAIL" recipient = "test@example.com" email_subject = "This is a test" is_html_content = true template = "{}" triggers = [ "ALERT_OPENED", "ALERT_RESOLVED" ] } resource "wavefront_alert_target" "test_target" { name = "Terraform Test Target" description = "Test target" method = "WEBHOOK" recipient = "https://hooks.slack.com/services/test/me" content_type = "application/json" custom_headers = { "Testing" = "true" } template = "{}" triggers = [ "ALERT_OPENED", "ALERT_RESOLVED", ] route { method = "WEBHOOK" target = "https://hooks.slack.com/services/test/me/prod" filter = { key = "env" value = "prod" } } route { method = "WEBHOOK" target = "https://hooks.slack.com/services/test/me/dev" filter = { key = "env" value = "dev" } } } ` + "`" + `` + "`" + `` + "`" + ` ## Import Alert Targets can be imported using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_alert_target.alert_target abcdEFGhijKLMNO ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "target_id",
					Description: `The target ID prefixed with ` + "`" + `target:` + "`" + ` for interpolating into a Wavefront Alert. ### Route The ` + "`" + `route` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "method",
					Description: `(Required) The notification method used for notification target. One of ` + "`" + `WEBHOOK` + "`" + `, ` + "`" + `EMAIL` + "`" + `, ` + "`" + `PAGERDUTY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The endpoint for the alert route. ` + "`" + `EMAIL` + "`" + `: email address. ` + "`" + `PAGERDUTY` + "`" + `: PagerDuty routing key. ` + "`" + `WEBHOOK` + "`" + `: URL endpoint.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) String that filters the route. Space delimited. Currently only allows a single key value pair. (e.g. ` + "`" + `env prod` + "`" + `) ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_alert_target" "test_target" { name = "Terraform Test Target" description = "Test target" method = "EMAIL" recipient = "test@example.com" email_subject = "This is a test" is_html_content = true template = "{}" triggers = [ "ALERT_OPENED", "ALERT_RESOLVED" ] } resource "wavefront_alert_target" "test_target" { name = "Terraform Test Target" description = "Test target" method = "WEBHOOK" recipient = "https://hooks.slack.com/services/test/me" content_type = "application/json" custom_headers = { "Testing" = "true" } template = "{}" triggers = [ "ALERT_OPENED", "ALERT_RESOLVED", ] route { method = "WEBHOOK" target = "https://hooks.slack.com/services/test/me/prod" filter = { key = "env" value = "prod" } } route { method = "WEBHOOK" target = "https://hooks.slack.com/services/test/me/dev" filter = { key = "env" value = "dev" } } } ` + "`" + `` + "`" + `` + "`" + ` ## Import Alert Targets can be imported using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_alert_target.alert_target abcdEFGhijKLMNO ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_app_dynamics",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for App Dynamics. This allows app dynamics cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) Username is a combination of userName and the account name.`,
				},
				resource.Attribute{
					Name:        "controller_name",
					Description: `(Required) Name of the SaaS controller.`,
				},
				resource.Attribute{
					Name:        "encrypted_password",
					Description: `(Required) Password for AppDynamics user.`,
				},
				resource.Attribute{
					Name:        "app_filter_regex",
					Description: `(Optional) List of regular expressions that an application name must match (case-insensitively) in order to be ingested.`,
				},
				resource.Attribute{
					Name:        "enable_rollup",
					Description: `(Optional) Set this to ` + "`" + `false` + "`" + ` to get separate results for all values within the time range, by default it is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_error_metrics",
					Description: `(Optional) Boolean flag to control Error metric ingestion.`,
				},
				resource.Attribute{
					Name:        "enable_business_trx_metrics",
					Description: `(Optional) Boolean flag to control Business Transaction metric ingestion.`,
				},
				resource.Attribute{
					Name:        "enable_backend_metrics",
					Description: `(Optional) Boolean flag to control Backend metric ingestion.`,
				},
				resource.Attribute{
					Name:        "enable_overall_perf_metrics",
					Description: `(Optional) Boolean flag to control Overall Performance metric ingestion.`,
				},
				resource.Attribute{
					Name:        "enable_individual_node_metrics",
					Description: `(Optional) Boolean flag to control Individual Node metric ingestion.`,
				},
				resource.Attribute{
					Name:        "enable_app_infra_metrics",
					Description: `(Optional) Boolean flag to control Application Infrastructure metric ingestion.`,
				},
				resource.Attribute{
					Name:        "enable_service_endpoint_metrics",
					Description: `(Optional) Boolean flag to control Service End point metric ingestion. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_app_dynamics" "app_dynamics" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } user_name = "example2" controller_name = "exampleController2" encrypted_password = "encryptedPassword" enable_rollup = false enable_error_metrics = true enable_business_trx_metrics = true enable_backend_metrics = true enable_overall_perf_metrics = true enable_individual_node_metrics = true enable_app_infra_metrics = true enable_service_endpoint_metrics = true } ` + "`" + `` + "`" + `` + "`" + ` ## Import AppDynamic Cloud Integrations can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_app_dynamics.app_dynamics a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_aws_external_id",
			Category:         "Resources",
			ShortDescription: `Provides an External ID for use in Amazon Web Services IAM Roles. This allows External IDs to be created and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The External ID created in Wavefront. ## Import External IDs can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_aws_external_id.external_id uGJdkH3k ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_azure",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for Microsoft Azure. This allows azure cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required) Client secret for an Azure service account within your project.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Client ID for an Azure service account within your project.`,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required) Tenant ID for an Azure service account within your project.`,
				},
				resource.Attribute{
					Name:        "resource_group_filter",
					Description: `(Optional) A list of Azure resource groups from which to pull metrics.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `(Optional) A regular expression that a metric name must match (case-insensitively) in order to be ingested.`,
				},
				resource.Attribute{
					Name:        "category_filter",
					Description: `(Optional) A list of Azure Activity Log categories. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" "tag3" = "value3" } category_filter = ["ADMINISTRATIVE", "SERVICEHEALTH"] client_id = "client-id2" client_secret = "client-secret2" tenant = "my-tenant2" service_refresh_rate_in_minutes = 10 } ` + "`" + `` + "`" + `` + "`" + ` ## Import Azure Cloud Integrations can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_azure.azure a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_azure_activity_log",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for Azure Activity Logs. This allows Azure activity log cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "client_secret",
					Description: `(Required) Client secret for an Azure service account within your project.`,
				},
				resource.Attribute{
					Name:        "client_id",
					Description: `(Required) Client ID for an Azure service account within your project.`,
				},
				resource.Attribute{
					Name:        "tenant",
					Description: `(Required) Tenant ID for an Azure service account within your project.`,
				},
				resource.Attribute{
					Name:        "category_filter",
					Description: `(Optional) A list of Azure services (such as Microsoft.Compute/virtualMachines) from which to pull metrics. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_azure_activity_log" "azure_activity_log" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" "tag3" = "value3" } category_filter = ["ADMINISTRATIVE", "SERVICEHEALTH"] client_id = "client-id2" client_secret = "client-secret2" tenant = "my-tenant2" service_refresh_rate_in_minutes = 10 } ` + "`" + `` + "`" + `` + "`" + ` ## Import Azure Activity Log Cloud Integrations can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_azure_activity_log.azure_al a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_cloudtrail",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for CloudTrail. This allows CloudTrail cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) The external ID corresponding to the Role ARN.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The AWS region of the S3 bucket where CloudTrail logs are stored.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Required) Name of the S3 bucket where CloudTrail logs are stored.`,
				},
				resource.Attribute{
					Name:        "filter_rule",
					Description: `(Optional) Rule to filter CloudTrail log event.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `(Optional) The common prefix, if any, appended to all CloudTrail log files. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_aws_external_id" "ext_id" { } resource "wavefront_cloud_integration_cloudtrail" "cloudtrail" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } region = "us-west-1" bucket_name = "example-s3-bucket2" filter_rule = "someFilterRule" role_arn = "arn:aws::1234567:role/example-arn" external_id = wavefront_cloud_integration_aws_external_id.ext_id.id service_refresh_rate_in_minutes = 10 } ` + "`" + `` + "`" + `` + "`" + ` ## Import CloudTrail Cloud Integrations can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_cloudtrail.cloudtrail a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_cloudwatch",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for CloudWatch. This allows CloudWatch cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) The external ID corresponding to the Role ARN.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.`,
				},
				resource.Attribute{
					Name:        "point_tag_filter_regex",
					Description: `(Optional) A regular expression that AWS tag key name must match (case-insensitively) in order to be ingested.`,
				},
				resource.Attribute{
					Name:        "volume_selection_tags",
					Description: `(Optional) A string->string map of allow list of volume tag-value pairs (in AWS). If the volume's AWS tags match this allow list, CloudWatch data about this volume is ingested. Multiple entries are OR'ed.`,
				},
				resource.Attribute{
					Name:        "instance_selection_tags",
					Description: `(Optional) A string->string map allow list of instance tag-value pairs (in AWS). If the instance's AWS tags match this allow list, CloudWatch data about this instance is ingested. Multiple entries are OR'ed.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `(Optional) A regular expression that a CloudWatch metric name must match (case-insensitively) in order to be ingested.`,
				},
				resource.Attribute{
					Name:        "namespaces",
					Description: `(Optional) A list of namespaces that limit what we query from CloudWatch. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_aws_external_id" "ext_id" { } resource "wavefront_cloud_integration_cloudwatch" "cloudwatch" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } role_arn = "arn:aws::1234567:role/example-arn" external_id = wavefront_cloud_integration_aws_external_id.ext_id.id namespaces = ["ec2", "elb", "route53"] instance_selection_tags = { "env" = "dev" "mirror" = "b" } volume_selection_tags = { "env" = "dev" } point_tag_filter_regex = "^dev" metric_filter_regex = "^.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_ec2",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for EC2. This allows EC2 cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "role_arn",
					Description: `(Required) The external ID corresponding to the Role ARN.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `(Required) The Role ARN that the customer has created in AWS IAM to allow access to Wavefront.`,
				},
				resource.Attribute{
					Name:        "hostname_tags",
					Description: `(Optional) A list of AWS instance tags to use as the ` + "`" + `source` + "`" + ` name in a series. Default is ` + "`" + `["hostname", "host", "name"]` + "`" + `. If no tag in the list is found, the series source is set to the instance id. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_aws_external_id" "ext_id" { } resource "wavefront_cloud_integration_ec2" "ec2" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } role_arn = "arn:aws::1234567:role/example-arn" external_id = wavefront_cloud_integration_aws_external_id.ext_id.id hostname_tags = ["host", "source", "name"] service_refresh_rate_in_minutes = 10 } ` + "`" + `` + "`" + `` + "`" + ` ## Import EC2 Cloud Integrations can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_ec2.ec2 a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_gcp",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for Google Cloud Platform. This allows Google Cloud Platform cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Google Cloud Platform (GCP) Project ID.`,
				},
				resource.Attribute{
					Name:        "json_key",
					Description: `(Required) Private key for a Google Cloud Platform (GCP) service account within your project. The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP.`,
				},
				resource.Attribute{
					Name:        "categories",
					Description: `(Optional) A list of Google Cloud Platform (GCP) services. Valid values are ` + "`" + `APPENGINE` + "`" + `, ` + "`" + `BIGQUERY` + "`" + `, ` + "`" + `BIGTABLE` + "`" + `, ` + "`" + `CLOUDFUNCTIONS` + "`" + `, ` + "`" + `CLOUDIOT` + "`" + `, ` + "`" + `CLOUDSQL` + "`" + `, ` + "`" + `CLOUDTASKS` + "`" + `, ` + "`" + `COMPUTE` + "`" + `, ` + "`" + `CONTAINER` + "`" + `, ` + "`" + `DATAFLOW` + "`" + `, ` + "`" + `DATAPROC` + "`" + `, ` + "`" + `DATASTORE` + "`" + `, ` + "`" + `FIREBASEDATABASE` + "`" + `, ` + "`" + `FIREBASEHOSTING` + "`" + `, ` + "`" + `FIRESTORE` + "`" + `, ` + "`" + `INTERCONNECT` + "`" + `, ` + "`" + `LOADBALANCING` + "`" + `, ` + "`" + `LOGGING` + "`" + `, ` + "`" + `ML` + "`" + `, ` + "`" + `MONITORING` + "`" + `, ` + "`" + `PUBSUB` + "`" + `, ` + "`" + `REDIS` + "`" + `, ` + "`" + `ROUTER` + "`" + `, ` + "`" + `SERVICERUNTIME` + "`" + `, ` + "`" + `SPANNER` + "`" + `, ` + "`" + `STORAGE` + "`" + `, ` + "`" + `TPU` + "`" + `, and ` + "`" + `VPN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `(Optional) A regular expression that a metric name must match (case-insensitively) in order to be ingested. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_gcp" "gcp" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } metric_filter_regex = "^(exampleMetricRegex).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_gcp_billing",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for Google Cloud Billing. This allows GCP Billing cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Google Cloud Platform (GCP) Project ID.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) API key for Google Cloud Platform (GCP).`,
				},
				resource.Attribute{
					Name:        "json_key",
					Description: `(Required) Private key for a Google Cloud Platform (GCP) service account within your project. The account must have at least Viewer permissions. This key must be in the JSON format generated by GCP. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_gcp_billing" "gcp_billing" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } project_id = "example-gcp-project" api_key = "example-api-key" json_key = <<EOF {...your gcp key ...} EOF service_refresh_rate_in_minutes = 10 } ` + "`" + `` + "`" + `` + "`" + ` ## Import GCP Billing Cloud Integrations can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_cloud_integration_gcp_billing.gcp_billing a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_cloud_integration_newrelic",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Cloud Integration for New Relic. This allows New Relic cloud integrations to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service",
					Description: `(Required) A value denoting which cloud service this service integrates with.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The human-readable name of this integration.`,
				},
				resource.Attribute{
					Name:        "additional_tags",
					Description: `(Optional) A list of point tag key-values to add to every point ingested using this integration.`,
				},
				resource.Attribute{
					Name:        "force_save",
					Description: `(Optional) Forces this resource to save, even if errors are present.`,
				},
				resource.Attribute{
					Name:        "service_refresh_rate_in_minutes",
					Description: `(Optional) How often, in minutes, to refresh the service.`,
				},
				resource.Attribute{
					Name:        "api_key",
					Description: `(Required) New Relic REST API key.`,
				},
				resource.Attribute{
					Name:        "app_filter_regex",
					Description: `(Optional) A regular expression that an application name must match (case-insensitively) in order to collect metrics.`,
				},
				resource.Attribute{
					Name:        "host_filter_regex",
					Description: `(Optional) A regular expression that a host name must match (case-insensitively) in order to collect metrics.`,
				},
				resource.Attribute{
					Name:        "metric_filter",
					Description: `(Optional) See [Metric Filter](#metric-filter). ### Metric Filter The ` + "`" + `metric_filter` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "app_name",
					Description: `(Required) The name of a NewRelic App.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `(Required) A regular expression that a metric name must match (case-insensitively) in order to be ingested. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_cloud_integration_newrelic" "newrelic" { name = "Test Integration" force_save = true additional_tags = { "tag1" = "value1" "tag2" = "value2" } api_key = "example-api-key" app_filter_regex = "^someApp.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_dashboard",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Dashboard resource. This allows dashboards to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the dashboard.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Human-readable description of the dashboard.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `(Required) Unique identifier, also a URL slug of the dashboard.`,
				},
				resource.Attribute{
					Name:        "section",
					Description: `(Required) Dashboard chart sections. See [dashboard sections](#dashboard-sections).`,
				},
				resource.Attribute{
					Name:        "display_query_parameters",
					Description: `(Optional) Whether the dashboard parameters section is opened by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "parameter_details",
					Description: `(Optional) The current JSON representation of dashboard parameters. See [parameter details](#parameter-details).`,
				},
				resource.Attribute{
					Name:        "display_section_table_of_contents",
					Description: `(Optional) Whether the "pills" quick-linked the sections of the dashboard are displayed by default when the dashboard is shown.`,
				},
				resource.Attribute{
					Name:        "can_modify",
					Description: `(Optional) A list of users that have modify ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "can_view",
					Description: `(Optional) A list of users that have view ACL access to the dashboard.`,
				},
				resource.Attribute{
					Name:        "event_filter_type",
					Description: `(Optional) How charts belonging to this dashboard should display events. ` + "`" + `BYCHART` + "`" + ` is default if unspecified. Valid options are: ` + "`" + `BYCHART` + "`" + `, ` + "`" + `AUTOMATIC` + "`" + `, ` + "`" + `ALL` + "`" + `, ` + "`" + `NONE` + "`" + `, ` + "`" + `BYDASHBOARD` + "`" + `, and ` + "`" + `BYCHARTANDDASHBOARD` + "`" + `. ### Dashboard Sections The ` + "`" + `section` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of this section.`,
				},
				resource.Attribute{
					Name:        "row",
					Description: `(Required) See [dashboard section rows](#dashboard-section-rows). ### Dashboard Section Rows The ` + "`" + `row` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "chart",
					Description: `(Required) Charts in this section. See [dashboard chart](#dashboard-chart). ### Dashboard Chart The ` + "`" + `chart` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) Query expression to plot on the chart. See [chart source queries](#chart-source-queries).`,
				},
				resource.Attribute{
					Name:        "chart_setting",
					Description: `(Required) Chart settings. See [chart settings](#chart-settings).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the source.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `(Required) String to label the units of the chart on the Y-Axis.`,
				},
				resource.Attribute{
					Name:        "summarization",
					Description: `(Required) Summarization strategy for the chart. MEAN is default. Valid options are, ` + "`" + `MEAN` + "`" + `, ` + "`" + `MEDIAN` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, ` + "`" + `SUM` + "`" + `, ` + "`" + `COUNT` + "`" + `, ` + "`" + `LAST` + "`" + `, ` + "`" + `FIRST` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the chart.`,
				},
				resource.Attribute{
					Name:        "base",
					Description: `(Optional) The base of logarithmic scale charts. Omit or set to 0 for the default linear scale. Usually set to 10 for the traditional logarithmic scale. ### Chart Source Queries The ` + "`" + `source` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the source.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) Query expression to plot on the chart.`,
				},
				resource.Attribute{
					Name:        "source_description",
					Description: `(Optional) A description for the purpose of this source.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `(Optional) Whether the source is disabled.`,
				},
				resource.Attribute{
					Name:        "scatter_plot_source",
					Description: `(Optional) For scatter plots, does this query source the X-axis or the Y-axis, ` + "`" + `X` + "`" + `, or ` + "`" + `Y` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_builder_enabled",
					Description: `(Optional) Whether or not this source line should have the query builder enabled. ### Chart Settings The ` + "`" + `chart_setting` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Chart Type. ` + "`" + `line` + "`" + ` refers to the Line Plot, ` + "`" + `scatter` + "`" + ` to the Point Plot, ` + "`" + `stacked-area` + "`" + ` to the Stacked Area plot, ` + "`" + `table` + "`" + ` to the Tabular View, ` + "`" + `scatterplot-xy` + "`" + ` to Scatter Plot, ` + "`" + `markdown-widget` + "`" + ` to the Markdown display, and ` + "`" + `sparkline` + "`" + ` to the Single Stat view. Valid options are ` + "`" + `line` + "`" + `, ` + "`" + `scatterplot` + "`" + `, ` + "`" + `stacked-area` + "`" + `, ` + "`" + `stacked-column` + "`" + `, ` + "`" + `table` + "`" + `, ` + "`" + `scatterplot-xy` + "`" + `, ` + "`" + `markdown-widget` + "`" + `, ` + "`" + `sparkline` + "`" + `, ` + "`" + `globe` + "`" + `, ` + "`" + `nodemap` + "`" + `, ` + "`" + `top-k` + "`" + `, ` + "`" + `status-list` + "`" + `, and ` + "`" + `histogram` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max",
					Description: `(Optional) Max value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "line_type",
					Description: `(Optional) Plot interpolation type. ` + "`" + `linear` + "`" + ` is default. Valid options are ` + "`" + `linear` + "`" + `, ` + "`" + `step-before` + "`" + `, ` + "`" + `step-after` + "`" + `, ` + "`" + `basis` + "`" + `, ` + "`" + `cardinal` + "`" + `, and ` + "`" + `monotone` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "stack_type",
					Description: `(Optional) Type of stacked chart (applicable only if chart type is ` + "`" + `stacked` + "`" + `). ` + "`" + `zero` + "`" + ` (default) means stacked from y=0. ` + "`" + `expand` + "`" + ` means normalized from 0 to 1. ` + "`" + `wiggle` + "`" + ` means minimize weighted changes. ` + "`" + `silhouette` + "`" + ` means to center the stream. Valid options are ` + "`" + `zero` + "`" + `, ` + "`" + `expand` + "`" + `, ` + "`" + `wiggle` + "`" + `, ` + "`" + `silhouette` + "`" + `, and ` + "`" + `bars` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "windowing",
					Description: `(Optional) For the tabular view, whether to use the full time window for the query or the last X minutes. Valid options are ` + "`" + `full` + "`" + ` or ` + "`" + `last` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "window_size",
					Description: `(Optional) Width, in minutes, of the time window to use for ` + "`" + `last` + "`" + ` windowing.`,
				},
				resource.Attribute{
					Name:        "show_hosts",
					Description: `(Optional) For the tabular view, whether to display sources. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_labels",
					Description: `(Optional) For the tabular view, whether to display labels. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_raw_values",
					Description: `(Optional) For the tabular view, whether to display raw values. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_column_tags",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "column_tags",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "tag_mode",
					Description: `(Optional) For the tabular view, which mode to use to determine which point tags to display. Valid options are ` + "`" + `all` + "`" + `, ` + "`" + `top` + "`" + `, or ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_tags",
					Description: `(Optional) For the tabular view defines how many point tags to display.`,
				},
				resource.Attribute{
					Name:        "custom_tags",
					Description: `(Optional) For the tabular view, a list of point tags to display when using the ` + "`" + `custom` + "`" + ` tag display mode.`,
				},
				resource.Attribute{
					Name:        "group_by_source",
					Description: `(Optional) For the tabular view, whether to group multi metrics into a single row by a common source. If ` + "`" + `false` + "`" + `, each source is displayed in its own row. if ` + "`" + `true` + "`" + `, multiple metrics for the same host are displayed as different columns in the same row.`,
				},
				resource.Attribute{
					Name:        "sort_values_descending",
					Description: `(Optional) For the tabular view, whether to display values in descending order. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "y1max",
					Description: `(Optional) For plots with multiple Y-axes, max value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1min",
					Description: `(Optional) For plots with multiple Y-axes, min value for the right side Y-axis. Set null for auto.`,
				},
				resource.Attribute{
					Name:        "y1_units",
					Description: `(Optional) For plots with multiple Y-axes, units for right side Y-axis.`,
				},
				resource.Attribute{
					Name:        "y0_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for left Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y1_scale_si_by_1024",
					Description: `(Optional) Whether to scale numerical magnitude labels for right Y-axis by 1024 in the IEC/Binary manner (instead of by 1000 like SI).`,
				},
				resource.Attribute{
					Name:        "y0_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the left Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "y1_unit_autoscaling",
					Description: `(Optional) Whether to automatically adjust magnitude labels and units for the right Y-axis to favor smaller magnitudes and larger units.`,
				},
				resource.Attribute{
					Name:        "invert_dynamic_legend_hover_control",
					Description: `(Optional) Whether to disable the display of the floating legend (but reenable it when the ctrl-key is pressed).`,
				},
				resource.Attribute{
					Name:        "fixed_legend_enabled",
					Description: `(Optional) Whether to enable a fixed tabular legend adjacent to the chart.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_use_raw_stats",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the legend uses non-summarized stats instead of summarized.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_position",
					Description: `(Optional) Where the fixed legend should be displayed with respect to the chart. Valid options are ` + "`" + `RIGHT` + "`" + `, ` + "`" + `TOP` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_display_stats",
					Description: `(Optional) For a chart with a fixed legend, a list of statistics to display in the legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_sort",
					Description: `(Optional) Whether to display ` + "`" + `TOP` + "`" + ` or ` + "`" + `BOTTOM` + "`" + ` ranked series in a fixed legend. Valid options are ` + "`" + `TOP` + "`" + `, and ` + "`" + `BOTTOM` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_limit",
					Description: `(Optional) Number of series to include in the fixed legend.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_filter_field",
					Description: `(Optional) Statistic to use for determining whether a series is displayed on the fixed legend. Valid options are ` + "`" + `CURRENT` + "`" + `, ` + "`" + `MEAN` + "`" + `, ` + "`" + `MEDIAN` + "`" + `, ` + "`" + `SUM` + "`" + `, ` + "`" + `MIN` + "`" + `, ` + "`" + `MAX` + "`" + `, and ` + "`" + `COUNT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fixed_legend_hide_label",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "xmax",
					Description: `(Optional) For x-y scatterplots, max value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymax",
					Description: `(Optional) For x-y scatterplots, max value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "xmin",
					Description: `(Optional) For x-y scatterplots, min value for the X-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "ymin",
					Description: `(Optional) For x-y scatterplots, min value for the Y-axis. Set to null for auto.`,
				},
				resource.Attribute{
					Name:        "time_based_coloring",
					Description: `(Optional) For x-y scatterplots, whether to color more recent points as darker than older points.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_value_type",
					Description: `(Optional) For the single stat view, where to display the name of the query or the value of the query. Valid options are ` + "`" + `VALUE` + "`" + ` or ` + "`" + `LABEL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_color",
					Description: `(Optional) For the single stat view, the color of the displayed text (when not dynamically determined). Values should be in ` + "`" + `rgba(,,,,)` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_vertical_position",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_horizontal_position",
					Description: `(Optional) For the single stat view, the horizontal position of the displayed text. Valid options are ` + "`" + `MIDDLE` + "`" + `, ` + "`" + `LEFT` + "`" + `, ` + "`" + `RIGHT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_font_size",
					Description: `(Optional) For the single stat view, the font size of the displayed text, in percent.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_prefix",
					Description: `(Optional) For the single stat view, a string to add before the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_display_postfix",
					Description: `(Optional) For the single stat view, a string to append to the displayed text.`,
				},
				resource.Attribute{
					Name:        "sparkline_size",
					Description: `(Optional) For the single stat view, this determines whether the sparkline of the statistic is displayed in the chart. Valid options are ` + "`" + `BACKGROUND` + "`" + `, ` + "`" + `BOTTOM` + "`" + `, ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_line_color",
					Description: `(Optional) For the single stat view, the color of the line. Values should be in ` + "`" + `rgba(,,,,)` + "`" + ` format.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `(Optional) Min value of the Y-axis. Set to null or leave blank for auto.`,
				},
				resource.Attribute{
					Name:        "plain_markdown_content",
					Description: `(Optional) The markdown content for a Markdown display, in plain text.`,
				},
				resource.Attribute{
					Name:        "sparkline_fill_color",
					Description: `(Optional) For the single stat view, the color of the background fill. Values should be in ` + "`" + `rgba(,,,,)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_colors",
					Description: `(Optional) For the single stat view, A list of colors that differing query values map to. Must contain one more element than ` + "`" + `sparkline_value_color_map_values_v2` + "`" + `. Values should be in ` + "`" + `rgba(,,,,)` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values_v2",
					Description: `(Optional) For the single stat view, a list of boundaries for mapping different query values to colors. Must contain one element less than ` + "`" + `sparkline_value_color_map_colors` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_values",
					Description: `(Optional) This setting is deprecated.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_color_map_apply_to",
					Description: `(Optional) For the single stat view, whether to apply dynamic color settings to the displayed ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `. Valid options are ` + "`" + `TEXT` + "`" + ` or ` + "`" + `BACKGROUND` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_decimal_precision",
					Description: `(Optional) For the single stat view, the decimal precision of the displayed number.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_text",
					Description: `(Optional) For the single stat view, a list of display text values that different query values map to. Must contain one more element than ` + "`" + `sparkline_value_text_map_thresholds` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sparkline_value_text_map_thresholds",
					Description: `(Optional) For the single stat view, a list of threshold boundaries for mapping different query values to display text. Must contain one element less than ` + "`" + `sparkline_value_text_map_text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expected_data_spacing",
					Description: `(Optional) Threshold (in seconds) for time delta between consecutive points in a series above which a dotted line will replace a solid in in line plots. Default is 60. ### Parameter Details The ` + "`" + `parameter_details` + "`" + ` mapping supports the following:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) The label for the parameter.`,
				},
				resource.Attribute{
					Name:        "values_to_readable_strings",
					Description: `(Required) A string->string map. At least one of the keys must match the value of ` + "`" + `default_value` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the parameters.`,
				},
				resource.Attribute{
					Name:        "default_value",
					Description: `(Required) The default value of the parameter.`,
				},
				resource.Attribute{
					Name:        "hide_from_view",
					Description: `(Required) If ` + "`" + `true` + "`" + ` the parameter will only be shown on the edit view of the dashboard.`,
				},
				resource.Attribute{
					Name:        "parameter_type",
					Description: `(Required) The type of the parameter. ` + "`" + `SIMPLE` + "`" + `, ` + "`" + `LIST` + "`" + `, or ` + "`" + `DYNAMIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag_key",
					Description: `(Optional) for ` + "`" + `TAG_KEY` + "`" + ` dynamic field types, the tag key to return.`,
				},
				resource.Attribute{
					Name:        "query_value",
					Description: `(Optional) For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the query to execute to return values.`,
				},
				resource.Attribute{
					Name:        "dynamic_field_type",
					Description: `(Optional) For ` + "`" + `DYNAMIC` + "`" + ` parameter types, the type of the field. Valid options are ` + "`" + `SOURCE` + "`" + `, ` + "`" + `SOURCE_TAG` + "`" + `, ` + "`" + `METRIC_NAME` + "`" + `, ` + "`" + `TAG_KEY` + "`" + `, and ` + "`" + `MATCHING_SOURCE_TAG` + "`" + `. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_dashboard" "chart_settings_dash" { name = "Terraform Chart Settings" description = "testing, testing" url = "tftestcreate" section { name = "section 1" row { chart { name = "chart 1" description = "chart number 1" units = "something per unit" source { name = "source name" query = "ts()" } summarization = "MEAN" chart_setting { auto_column_tags = false column_tags = "deprecated" custom_tags = ["tag1", "tag2"] expected_data_spacing = 120 fixed_legend_display_stats = ["stat1", "stat2"] fixed_legend_enabled = true fixed_legend_filter_field = "CURRENT" fixed_legend_filter_limit = 1 fixed_legend_filter_sort = "TOP" fixed_legend_hide_label = false fixed_legend_position = "RIGHT" fixed_legend_use_raw_stats = true line_type = "linear" max = 100 min = 0 sparkline_decimal_precision = 1 sparkline_display_color = "rgba(1,1,1,1)" sparkline_display_font_size = 14 sparkline_display_horizontal_position = "LEFT" sparkline_display_postfix = "postfix" sparkline_display_prefix = "prefix" sparkline_display_value_type = "VALUE" sparkline_display_vertical_position = "deprecated" sparkline_fill_color = "rgba(1,1,1,1)" sparkline_line_color = "rgba(1,1,1,1)" sparkline_size = "BOTTOM" sparkline_value_color_map_apply_to = "TEXT" sparkline_value_color_map_colors = ["rgba(1,1,1,1)", "rgba(2,2,2,2)", "rgba(3,3,3,3)"] sparkline_value_color_map_values = [ 1, 2 ] sparkline_value_text_map_text = ["a"] sparkline_value_text_map_thresholds = [1] type = "line" y0_scale_si_by_1024 = true y0_unit_autoscaling = true y1max = 100 y1_scale_si_by_1024 = true y1_unit_autoscaling = true y1_units = "units" } } } } parameter_details { name = "param1" label = "param1" default_value = "defaultQuery" hide_from_view = false parameter_type = "DYNAMIC" values_to_readable_strings = { defaultQuery = "dev-elasticsearch" } dynamic_field_type = "MATCHING_SOURCE_TAG" query_value = "ts(aws.ec2.diskwritebytes.average)" } tags = [ "terraform", "test" ] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Dashboards can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_dashboard.dashboard tftestimport ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_dashboard_json",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Dashboard JSON resource. This allows dashboards to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dashboard_json",
					Description: `(Required) See the [Wavefront API Documentation](https://docs.wavefront.com/wavefront_api.html#api-documentation-wavefront-instance) for instructions on how to get to your API documentation for more details. ## Import Dashboard JSON can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_dashboard_json.dashboard_json tftestimport ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_derived_metric",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Derived Metric Resource. This allows derived metrics to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Derived Metric in Wavefront.`,
				},
				resource.Attribute{
					Name:        "query",
					Description: `(Required) A Wavefront query that is evaluated at regular intervals (default is 1 minute).`,
				},
				resource.Attribute{
					Name:        "minutes",
					Description: `(Required) How frequently the query generating the derived metric is run.`,
				},
				resource.Attribute{
					Name:        "additional_information",
					Description: `(Optional) User-supplied additional explanatory information for the derived metric.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags to assign to this resource. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_derived_metric" "derived" { name = "dummy derived metric" minutes = 10 query = "aliasMetric(mavg(5m, ts(cpu.idle)), \"cpu.5m-avg\")" additional_information = "this is a dummy derived metric" tags = [ "cpu.5m.avg", "derived" ] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Derived Metrics can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_derived_metric.derived_metric 1577102900578 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_event",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Event resource. This allows events to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the event as it is displayed in Wavefront.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tags to assign to this resource.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `(Required) The annotations associated with the event.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `(Optional) The end time of the event in epoch milliseconds. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_event" "event" { name = "terraform-test" annotations = { severity = "info" type = "event type" details = "description" } tag = [ "eventTag1" ] start_time = 1490000000000 end_time = 1490000000001 } ` + "`" + `` + "`" + `` + "`" + ` ## Import You can import events by using the id, for example: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_event.event 1479868728473 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_external_link",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront External Link Resource. This allows external links to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the external link.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) Human-readable description for this link.`,
				},
				resource.Attribute{
					Name:        "template",
					Description: `(Required) The mustache template for this link. The template must expand to a full URL, including scheme, origin, etc.`,
				},
				resource.Attribute{
					Name:        "metric_filter_regex",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. If present, the metric name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "source_filter_regex",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. If present, the source name of the highlighted series must match this regular expression in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "point_tag_filter_regexes",
					Description: `(Optional) Controls whether a link is displayed in the context menu of a highlighted series. This is a map from string to regular expression. The highlighted series must contain point tags whose keys are present in the keys of this map and whose values match the regular expressions associated with those keys in order for the link to be displayed.`,
				},
				resource.Attribute{
					Name:        "is_log_integration",
					Description: `Whether this is a "Log Integration" subType of external link. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_external_link" "basic" { name = "External Link" description = "An external link description" template = "https://example.com/source={{{source}}}&startTime={{startEpochMillis}}" metric_filter_regex = "^metric.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_ingestion_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Ingestion Policy Resource. This allows ingestion policies to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ingestion policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the ingestion policy. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_ingestion_policy" "basic" { name = "test_ingestion" description = "An ingestion policy for testing" } ` + "`" + `` + "`" + `` + "`" + ` ## Import ingestion policies can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_ingestion_policy.basic test_ingestion-1611946841064 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_maintenance_window",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Maintenance Window Resource. This allows maintenance windows to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "reason",
					Description: `(Required) The reason for the maintenance window.`,
				},
				resource.Attribute{
					Name:        "title",
					Description: `(Required) The title of the maintenance window.`,
				},
				resource.Attribute{
					Name:        "start_time_in_seconds",
					Description: `(Required) start time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "end_time_in_seconds",
					Description: `(Required) end time in seconds after 1 Jan 1970 GMT.`,
				},
				resource.Attribute{
					Name:        "relevant_customer_tags",
					Description: `List of alert tags whose matching alerts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags",
					Description: `List of source/host tags whose matching sources/hosts will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_names",
					Description: `List of source/host names that will be put into maintenance because of this maintenance window. At least one of ` + "`" + `relevant_customer_tags` + "`" + `, ` + "`" + `relevant_host_tags` + "`" + `, or ` + "`" + `relevant_host_names` + "`" + ` is required.`,
				},
				resource.Attribute{
					Name:        "relevant_host_tags_anded",
					Description: `(Optional) Whether to AND source/host tags listed in ` + "`" + `relevant_host_tags` + "`" + `. If ` + "`" + `true` + "`" + `, a source/host must contain all tags in order for the maintenance window to apply. If ` + "`" + `false` + "`" + `, the tags are OR'ed, and a source/host must contain one of the tags. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_tag_group_host_names_group_anded",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, a source/host must be in ` + "`" + `relevant_host_names` + "`" + ` and have tags matching the specification formed by ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + ` in order for this maintenance window to apply. If ` + "`" + `false` + "`" + `, a source/host must either be in ` + "`" + `relevant_host_names` + "`" + ` or match ` + "`" + `relevant_host_tags` + "`" + ` and ` + "`" + `relevant_host_tags_anded` + "`" + `. Default: ` + "`" + `false` + "`" + `. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_maintenance_window" "basic" { reason = "Routine maintenance for 2020" title = "Routine maintenance" start_time_in_seconds = 1600123456 end_time_in_seconds = 1601123456 relevant_host_tags = ["my_host_tag1", "my_host_tag2"] relevant_host_tags_anded = true } ` + "`" + `` + "`" + `` + "`" + ` ## Import Maintenance windows can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_maintenance_window.basic 1600383357095 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_metrics_policy",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Metrics Policy Resource. This allows management of Metrics Policy to control access to time series, histograms, and delta counters`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "policy_rules",
					Description: `(Required) List of Metrics Policies, must have at least one entry.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name identifier for a Metrics Policy. The name is visible on the Metrics Security Policy page.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A detailed description of the Metrics Policy. The description is visible only when you edit the rule.`,
				},
				resource.Attribute{
					Name:        "account_ids",
					Description: `(Optional) List of account ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.`,
				},
				resource.Attribute{
					Name:        "user_group_ids",
					Description: `(Optional) List of user group ids to apply Metrics Policy to. Must have at least one associated account_id, user_group_id, or role_id.`,
				},
				resource.Attribute{
					Name:        "prefixes",
					Description: `(Required) List of prefixes to match metrics on. You can specify the full metric name or use a wildcard character in metric names, sources, or point tags. The wildcard character alone (`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) List of Key/Value tags to select target metrics for policy.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The tag's key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The tag's value.`,
				},
				resource.Attribute{
					Name:        "tags_anded",
					Description: `(Required) Bool where ` + "`" + `true` + "`" + ` require all tags are met by selected metrics, else ` + "`" + `false` + "`" + ` select metrics that match any give tag.`,
				},
				resource.Attribute{
					Name:        "updater_id",
					Description: `The account_id who applied the current policy.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `When the policy was applied in epoch_millis. ## Data Source Provides a Wavefront Metrics Policy Data Source. This allows looking up the current policy and associated rules. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl data "wavefront_metrics_policy" "policy" {} output "policy" { value = data.wavefront_metrics_policy.policy } ` + "`" + `` + "`" + `` + "`" + ` ## Import Users can be imported by using the ` + "`" + `updated_epoch_millis` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_metrics_policy.some_metrics_policy 1651846476678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "updater_id",
					Description: `The account_id who applied the current policy.`,
				},
				resource.Attribute{
					Name:        "updated_epoch_millis",
					Description: `When the policy was applied in epoch_millis. ## Data Source Provides a Wavefront Metrics Policy Data Source. This allows looking up the current policy and associated rules. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl data "wavefront_metrics_policy" "policy" {} output "policy" { value = data.wavefront_metrics_policy.policy } ` + "`" + `` + "`" + `` + "`" + ` ## Import Users can be imported by using the ` + "`" + `updated_epoch_millis` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_metrics_policy.some_metrics_policy 1651846476678 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_role",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Role Resource. This allows roles to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A short description of the role.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) A list of permissions to assign to this role. Valid options are ` + "`" + `agent_management` + "`" + `, ` + "`" + `alerts_management` + "`" + `, ` + "`" + `dashboard_management` + "`" + `, ` + "`" + `embedded_charts` + "`" + `, ` + "`" + `events_management` + "`" + `, ` + "`" + `external_links_management` + "`" + `, ` + "`" + `host_tag_management` + "`" + `, ` + "`" + `metrics_management` + "`" + `, and ` + "`" + `user_management` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "assignees",
					Description: `(Optional) A list of user groups or accounts to assign to this role. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_user_group" "agents_group" { name = "Test Group" description = "Test Group" } resource "wavefront_user" "basic" { email = "test+tftesting@example.com" user_groups = [ wavefront_user_group.test_group.id ] } resource "wavefront_role" "agent_management" { name = "Agent Management Role" description = "Agent Management Role for Testing" permissions = [ "agent_management" ] assignees = [ wavefront_user_group.test_group.id ] } ` + "`" + `` + "`" + `` + "`" + ` ## Import Roles can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_role.some_role a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_service_account",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront Service Account Resource. This allows service accounts to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "identifier",
					Description: `(Required) The unique identifier of the service account to create. Must have the prefix ` + "`" + `sa::` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Required) Whether or not the service account is active.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The description of the service account.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) List of permission to grant to this service account. Valid options are ` + "`" + `agent_management` + "`" + `, ` + "`" + `alerts_management` + "`" + `, ` + "`" + `dashboard_management` + "`" + `, ` + "`" + `embedded_charts` + "`" + `, ` + "`" + `events_management` + "`" + `, ` + "`" + `external_links_management` + "`" + `, ` + "`" + `host_tag_management` + "`" + `, ` + "`" + `metrics_management` + "`" + `, and ` + "`" + `user_management` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_groups",
					Description: `(Optional) List of user groups for this service account.`,
				},
				resource.Attribute{
					Name:        "ingestion_policy",
					Description: `(Optional) ID of ingestion policy. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_user_group" "test_group" { name = "Test Group" description = "Test Group" } resource "wavefront_ingestion_policy" "test_ingestion" { name = "test_ingestion" description = "An ingestion policy for testing" } resource "wavefront_service_account" "basic" { identifier = "sa::tftesting" active = true description = "A service account description" permissions = [ "agent_management", "events_management", ] user_groups = [ wavefront_user_group.test_group.id ] ingestion_policy = wavefront_ingestion_policy.test_ingestion.id } ` + "`" + `` + "`" + `` + "`" + ` ## Import Service accounts can be imported by using ` + "`" + `identifier` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_service_account.basic sa::tftesting ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_user",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront User Resource. This allows user accounts to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The unique identifier of the user account to create. Must be a valid email address.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) List of permission to grant to this user. Valid options are ` + "`" + `agent_management` + "`" + `, ` + "`" + `alerts_management` + "`" + `, ` + "`" + `dashboard_management` + "`" + `, ` + "`" + `embedded_charts` + "`" + `, ` + "`" + `events_management` + "`" + `, ` + "`" + `external_links_management` + "`" + `, ` + "`" + `host_tag_management` + "`" + `, ` + "`" + `metrics_management` + "`" + `, and ` + "`" + `user_management` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_groups",
					Description: `(Optional) List of user groups to this user. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_user_group" "test_group" { name = "Test Group" description = "Test Group" } resource "wavefront_user" "basic" { email = "test+tftesting@example.com" permissions = [ "agent_management", "events_management", ] user_groups = [ wavefront_user_group.test_group.id ] } ` + "`" + `` + "`" + `` + "`" + ` ## Attribute Reference`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "wavefront_user_group",
			Category:         "Resources",
			ShortDescription: `Provides a Wavefront User Group Resource. This allows user groups to be created, updated, and deleted.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the user group.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A short description of the user group. ### Example ` + "`" + `` + "`" + `` + "`" + `hcl resource "wavefront_user_group" "basic" { name = "Basic User Group" description = "Basic User Group for Unit Tests" } ` + "`" + `` + "`" + `` + "`" + ` ## Import User Groups can be imported by using the ` + "`" + `id` + "`" + `, e.g.: ` + "`" + `` + "`" + `` + "`" + ` $ terraform import wavefront_user_group.some_group a411c16b-3cf7-4f03-bf11-8ca05aab898d ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"wavefront_alert":                                0,
		"wavefront_alert_target":                         1,
		"wavefront_cloud_integration_app_dynamics":       2,
		"wavefront_cloud_integration_aws_external_id":    3,
		"wavefront_cloud_integration_azure":              4,
		"wavefront_cloud_integration_azure_activity_log": 5,
		"wavefront_cloud_integration_cloudtrail":         6,
		"wavefront_cloud_integration_cloudwatch":         7,
		"wavefront_cloud_integration_ec2":                8,
		"wavefront_cloud_integration_gcp":                9,
		"wavefront_cloud_integration_gcp_billing":        10,
		"wavefront_cloud_integration_newrelic":           11,
		"wavefront_dashboard":                            12,
		"wavefront_dashboard_json":                       13,
		"wavefront_derived_metric":                       14,
		"wavefront_event":                                15,
		"wavefront_external_link":                        16,
		"wavefront_ingestion_policy":                     17,
		"wavefront_maintenance_window":                   18,
		"wavefront_metrics_policy":                       19,
		"wavefront_role":                                 20,
		"wavefront_service_account":                      21,
		"wavefront_user":                                 22,
		"wavefront_user_group":                           23,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
