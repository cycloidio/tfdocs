package infracost

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_api_gateway_rest_api",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The ID of the Rest API.`,
				},
				resource.Attribute{
					Name:        "monthly_requests",
					Description: `(Optional) The estimated monthly requests to the Rest API Gateway. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_apigatewayv2_api",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The ID of the Rest API.`,
				},
				resource.Attribute{
					Name:        "monthly_requests",
					Description: `(Optional) The estimated monthly requests to the HTTP API Gateway.`,
				},
				resource.Attribute{
					Name:        "average_request_size",
					Description: `(Optional) The estimated average request size in (KB) sent to the HTTP API Gateway. Requests are metered in 512KB increments, maximum size is 10MB.`,
				},
				resource.Attribute{
					Name:        "monthly_messages",
					Description: `(Optional) The number of messages sent to the Websocket API Gateway.`,
				},
				resource.Attribute{
					Name:        "average_message_size",
					Description: `(Optional) The average size of the messages (KB) sent to the Websocket API Gateway. Messages are metered in 32 KB increments, maximum size is 128KB. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_cloudwatch_log_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the CloudWatch logs log group to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_ingestion",
					Description: `(Optional) The estimated GB of data ingested by CloudWatch logs per month.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_storage",
					Description: `(Optional) The estimated GB of data stored by CloudWatch logs per month.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_scanned",
					Description: `(Optional) The estimated GB of data scanned by CloudWatch logs insights per month. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_codebuild_project",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the Codebuild Projects to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_build_minutes",
					Description: `(Required) The estimated total duration of monthly codebuild execution usage in minutes. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_dx_gateway_association",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the DX gateway association(s) to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_processed",
					Description: `(Optional) The estimated GB of data processed by the DX gateway association per month. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_dynamodb_table",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the DynamoDBs to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_write_request_units",
					Description: `(Optional) The estimated write request units per month in (used for on-demand DynamoDB).`,
				},
				resource.Attribute{
					Name:        "monthly_read_request_units",
					Description: `(Optional) The estimated read request units per month in (used for on-demand DynamoDB).`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_storage",
					Description: `(Optional) The estimated storage for tables per month in GBs.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_continuous_backup_storage",
					Description: `(Optional) The estimated storage for continuous backups (PITR) per month in GBs.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_on_demand_backup_storage",
					Description: `(Optional) The estimated storage for on-demand backups per month in GBs.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_restore",
					Description: `(Optional) The estimated size of restored data per month in GBs.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_in",
					Description: `(Optional) The estimated size of transferred data into DynamoDB per month in GBs.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_out",
					Description: `(Optional) The estimated size of transferred data out of DynamoDB per month in GBs.`,
				},
				resource.Attribute{
					Name:        "monthly_streams_read_request_units",
					Description: `(Optional) The estimated streams read request units per month. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_ec2_transit_gateway_vpc_attachment",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the EC2 transit gateway attachment(s) to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_processed",
					Description: `(Optional) The estimated GB of data processed by the EC2 transit gateway attachment(s) per month. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_ecr_repository",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the ECR Repository to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) The estimated monthly ECR repository storage usage. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_lambda_function",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the Lambda functions to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_requests",
					Description: `(Optional) The estimated monthly requests to the Lambda function per month. See [Usage values](#usage_values) below for details on attributes.`,
				},
				resource.Attribute{
					Name:        "average_request_duration",
					Description: `(Optional) The estimated average duration of each request in milliseconds. See [Usage values](#usage_values) below for details on attributes. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_nat_gateway",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the Nat Gateways to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_processed",
					Description: `(Optional) The estimated GB of data processed by the NAT Gateway per month. See [monthly_gb_data_processed](#monthly_gb_data_processed) below for details on attributes. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_sns_topic",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the SNS Topic to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_requests",
					Description: `(Optional) The estimated monthly requests to SNS.`,
				},
				resource.Attribute{
					Name:        "request_size",
					Description: `(Optional) The size of the requests to SNS, SNS bills in 64KB chunks. So if you process 1,000,000 requests at 128KB you pay for 2,000,000 requests. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_sns_topic_subscription",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the SNS Topic to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_requests",
					Description: `(Optional) The estimated monthly requests to SNS.`,
				},
				resource.Attribute{
					Name:        "request_size",
					Description: `(Optional) The size of the requests to SNS, SNS bills in 64KB chunks. So if you process 1,000,000 requests at 128KB you pay for 2,000,000 requests. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_sqs_queue",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the SQS Queues to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_requests",
					Description: `(Optional) The estimated monthly requests to SQS.`,
				},
				resource.Attribute{
					Name:        "request_size",
					Description: `(Optional) The size of the requests to SQS, SQS bills in 64KB chunks. So if you process 1,000,000 requests at 128KB you pay for 2,000,000 requests. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_vpc_endpoint",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the VPC endpoint(s) to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_processed",
					Description: `(Optional) The estimated GB of data processed by the VPC endpoint(s) per month. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infracost_aws_vpn_connection",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resources",
					Description: `(Required) The IDs of the VPN connections to apply the estimated usage.`,
				},
				resource.Attribute{
					Name:        "monthly_gb_data_processed",
					Description: `(Optional) The estimated monthly data processed through a transit gateway attached to your VPN Connection. ### Usage values Each of the usage value blocks currently supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Optional) The estimated value.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"infracost_aws_api_gateway_rest_api":               0,
		"infracost_aws_apigatewayv2_api":                   1,
		"infracost_aws_cloudwatch_log_group":               2,
		"infracost_aws_codebuild_project":                  3,
		"infracost_aws_dx_gateway_association":             4,
		"infracost_aws_dynamodb_table":                     5,
		"infracost_aws_ec2_transit_gateway_vpc_attachment": 6,
		"infracost_aws_ecr_repository":                     7,
		"infracost_aws_lambda_function":                    8,
		"infracost_aws_nat_gateway":                        9,
		"infracost_aws_sns_topic":                          10,
		"infracost_aws_sns_topic_subscription":             11,
		"infracost_aws_sqs_queue":                          12,
		"infracost_aws_vpc_endpoint":                       13,
		"infracost_aws_vpn_connection":                     14,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
