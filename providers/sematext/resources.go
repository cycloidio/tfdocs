package sematext

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_akka",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_apache",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_awsebs",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below;`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(optional) if not set then reads from env AWS_ACCESS_KEY_ID.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(optionl) is not present set from env AWS_SECRET_ACCESS_KEY`,
				},
				resource.Attribute{
					Name:        "aws_fetch_frequency",
					Description: `(required) one of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(optional) if not present withh set from env AWS_REGION. ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_awsec2",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below;`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(optional) if not set then reads from env AWS_ACCESS_KEY_ID.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(optionl) is not present set from env AWS_SECRET_ACCESS_KEY`,
				},
				resource.Attribute{
					Name:        "aws_fetch_frequency",
					Description: `(required) one of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(optional) if not present withh set from env AWS_REGION. ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_awselb",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below;`,
				},
				resource.Attribute{
					Name:        "aws_access_key",
					Description: `(optional) if not set then reads from env AWS_ACCESS_KEY_ID.`,
				},
				resource.Attribute{
					Name:        "aws_secret_key",
					Description: `(optionl) is not present set from env AWS_SECRET_ACCESS_KEY`,
				},
				resource.Attribute{
					Name:        "aws_fetch_frequency",
					Description: `(required) one of MINUTE|FIVE_MINUTES|FIFTEEN_MINUTES.`,
				},
				resource.Attribute{
					Name:        "aws_region",
					Description: `(optional) if not present withh set from env AWS_REGION. ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_cassandra",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_clickhouse",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_elasticsearch",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_hadoopmrv1",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_hadoopyarn",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_haproxy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_hbase",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_infra",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_jvm",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_kafka",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_logsene",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_mobilelogs",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_mongodb",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_mysql",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_nginx",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_nginxplus",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_nodejs",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_postgresql",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_rabbitmq",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_redis",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_solr",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_solrcloud",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_spark",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_storm",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_tomcat",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_monitor_zookeeper",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `List attributes that this resource exports;`,
				},
				resource.Attribute{
					Name:        "billing_plan_id",
					Description: `List attributes that this resource exports. [Refer to plan guidance for list of legal values](../guides/plans.md);`,
				},
				resource.Attribute{
					Name:        "apptoken.name",
					Description: `Refer note below; ## App-tokens Sematext Cloud Apps have one or more apptokens. These app-tokens are named and have an id issued on creation. This is so that downstream resources can be grouped, with the instances within each group using the same app-token. Within Terraform state the apptoken.id is a calculated value, issued by Sematext Cloud. When provisioning, downstream resources should set the app-token id in their configuration. To prevent redeployment of downstream resources on change to apptokens we recommend using [null resource provider](https://registry.terraform.io/providers/hashicorp/null/latest/docs/resources/resource) to handle the provisioning. ## Outputs`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The app id. Provided on creation and used in terraform destroy operations.`,
				},
				resource.Attribute{
					Name:        "apptoken.id",
					Description: `On creation of the resource an app-token is generated for re-use in configuration of APM / collectors.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"sematext_sematext_monitor_akka":          0,
		"sematext_sematext_monitor_apache":        1,
		"sematext_sematext_monitor_awsebs":        2,
		"sematext_sematext_monitor_awsec2":        3,
		"sematext_sematext_monitor_awselb":        4,
		"sematext_sematext_monitor_cassandra":     5,
		"sematext_sematext_monitor_clickhouse":    6,
		"sematext_sematext_monitor_elasticsearch": 7,
		"sematext_sematext_monitor_hadoopmrv1":    8,
		"sematext_sematext_monitor_hadoopyarn":    9,
		"sematext_sematext_monitor_haproxy":       10,
		"sematext_sematext_monitor_hbase":         11,
		"sematext_sematext_monitor_infra":         12,
		"sematext_sematext_monitor_jvm":           13,
		"sematext_sematext_monitor_kafka":         14,
		"sematext_sematext_monitor_logsene":       15,
		"sematext_sematext_monitor_mobilelogs":    16,
		"sematext_sematext_monitor_mongodb":       17,
		"sematext_sematext_monitor_mysql":         18,
		"sematext_sematext_monitor_nginx":         19,
		"sematext_sematext_monitor_nginxplus":     20,
		"sematext_sematext_monitor_nodejs":        21,
		"sematext_sematext_monitor_postgresql":    22,
		"sematext_sematext_monitor_rabbitmq":      23,
		"sematext_sematext_monitor_redis":         24,
		"sematext_sematext_monitor_solr":          25,
		"sematext_sematext_monitor_solrcloud":     26,
		"sematext_sematext_monitor_spark":         27,
		"sematext_sematext_monitor_storm":         28,
		"sematext_sematext_monitor_tomcat":        29,
		"sematext_sematext_monitor_zookeeper":     30,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
