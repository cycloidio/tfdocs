package ably

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ably_api_key",
			Category:         "Resources",
			ShortDescription: `The ably_key resource allows you to create and manage Ably API keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_app",
			Category:         "Resources",
			ShortDescription: `The ably_app resource allows you to create and manage Ably Apps and configure Ably Push notifications. Read more about Ably Push Notifications in Ably documentation: https://ably.com/docs/general/push`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_namespace",
			Category:         "Resources",
			ShortDescription: `The ably_namespace resource allows you to manage namespaces for channel rules in Ably. Read more in the Ably documentation: https://ably.com/docs/general/channel-rules-namespaces.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_queue",
			Category:         "Resources",
			ShortDescription: `The ably_queue resource allows you to create and manage Ably queues. Read more about Ably queues in Ably documentation: https://ably.com/docs/general/queues.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_amqp",
			Category:         "Resources",
			ShortDescription: `The ably_rule_amqp resource allows you to create and manage an Ably integration rule for AMQP. Read more at https://ably.com/docs/general/firehose/amqp-rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_amqp_external",
			Category:         "Resources",
			ShortDescription: `The ably_rule_amqp_external resource allows you to create and manage an Ably integration rule for Firehose. Read more at https://ably.com/docs/general/firehose`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_azure_function",
			Category:         "Resources",
			ShortDescription: `The ably_rule_azure_function resource allows you to create and manage an Ably integration rule for Microsoft Azure Functions. Read more at https://ably.com/docs/general/webhooks/azure`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_cloudflare_worker",
			Category:         "Resources",
			ShortDescription: `The ably_rule_cloudflare_worker resource allows you to create and manage an Ably integration rule for Cloudflare workers. Read more at https://ably.com/docs/general/webhooks/cloudflare`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_google_function",
			Category:         "Resources",
			ShortDescription: `The ably_rule_google_cloud_function resource allows you to create and manage an Ably integration rule for Google cloud functions. Read more at https://ably.com/docs/general/webhooks/google-functions`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_http",
			Category:         "Resources",
			ShortDescription: `The ably_rule_http resource allows you to create and manage an Ably integration rule for HTTP. Read more at https://ably.com/docs/general/webhooks`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_ifttt",
			Category:         "Resources",
			ShortDescription: `The ably_rule_ifttt resource allows you to create and manage an Ably integration rule for IFTTT. Read more at https://ably.com/docs/general/webhooks/ifttt`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_kafka",
			Category:         "Resources",
			ShortDescription: `The ably_rule_kafka resource allows you to create and manage an Ably integration rule for Kafka. Read more at https://ably.com/docs/general/firehose/kafka-rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_kinesis",
			Category:         "Resources",
			ShortDescription: `The ably_rule_kinesis resource allows you to create and manage an Ably integration rule for AWS Kinesis. Read more at https://ably.com/docs/general/firehose/kinesis-rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_lambda",
			Category:         "Resources",
			ShortDescription: `The ably_rule_lambda resource allows you to create and manage an Ably integration rule for AWS Lambda. Read more at https://ably.com/docs/general/webhooks/aws-lambda`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_pulsar",
			Category:         "Resources",
			ShortDescription: `The ably_rule_pulsar resource allows you to create and manage an Ably integration rule for Pulsar. Read more at https://ably.com/docs/general/firehose/pulsar-rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_sqs",
			Category:         "Resources",
			ShortDescription: `The ably_rule_sqs resource allows you to create and manage an Ably integration rule for AWS SQS. Read more at https://ably.com/docs/general/firehose/sqs-rule`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ably_rule_zapier",
			Category:         "Resources",
			ShortDescription: `The ably_rule_zapier resource allows you to create and manage an Ably integration rule for Zapier. Read more at https://ably.com/docs/general/webhooks/zapier`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ably_api_key":                0,
		"ably_app":                    1,
		"ably_namespace":              2,
		"ably_queue":                  3,
		"ably_rule_amqp":              4,
		"ably_rule_amqp_external":     5,
		"ably_rule_azure_function":    6,
		"ably_rule_cloudflare_worker": 7,
		"ably_rule_google_function":   8,
		"ably_rule_http":              9,
		"ably_rule_ifttt":             10,
		"ably_rule_kafka":             11,
		"ably_rule_kinesis":           12,
		"ably_rule_lambda":            13,
		"ably_rule_pulsar":            14,
		"ably_rule_sqs":               15,
		"ably_rule_zapier":            16,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
