package sematext

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "sematext_sematext_app_akka",
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
			Type:             "sematext_sematext_app_apache",
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
			Type:             "sematext_sematext_app_awsebs",
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
			Type:             "sematext_sematext_app_awsec2",
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
			Type:             "sematext_sematext_app_awselb",
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
			Type:             "sematext_sematext_app_cassandra",
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
			Type:             "sematext_sematext_app_clickhouse",
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
			Type:             "sematext_sematext_app_elasticsearch",
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
			Type:             "sematext_sematext_app_hadoopmrv1",
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
			Type:             "sematext_sematext_app_hadoopyarn",
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
			Type:             "sematext_sematext_app_haproxy",
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
			Type:             "sematext_sematext_app_hbase",
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
			Type:             "sematext_sematext_app_infra",
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
			Type:             "sematext_sematext_app_jvm",
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
			Type:             "sematext_sematext_app_kafka",
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
			Type:             "sematext_sematext_app_logsene",
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
			Type:             "sematext_sematext_app_mobilelogs",
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
			Type:             "sematext_sematext_app_mongodb",
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
			Type:             "sematext_sematext_app_mysql",
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
			Type:             "sematext_sematext_app_nginx",
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
			Type:             "sematext_sematext_app_nginxplus",
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
			Type:             "sematext_sematext_app_nodejs",
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
			Type:             "sematext_sematext_app_postgresql",
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
			Type:             "sematext_sematext_app_rabbitmq",
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
			Type:             "sematext_sematext_app_redis",
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
			Type:             "sematext_sematext_app_solr",
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
			Type:             "sematext_sematext_app_solrcloud",
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
			Type:             "sematext_sematext_app_spark",
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
			Type:             "sematext_sematext_app_storm",
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
			Type:             "sematext_sematext_app_tomcat",
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
			Type:             "sematext_sematext_app_zookeeper",
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

		"sematext_sematext_app_akka":          0,
		"sematext_sematext_app_apache":        1,
		"sematext_sematext_app_awsebs":        2,
		"sematext_sematext_app_awsec2":        3,
		"sematext_sematext_app_awselb":        4,
		"sematext_sematext_app_cassandra":     5,
		"sematext_sematext_app_clickhouse":    6,
		"sematext_sematext_app_elasticsearch": 7,
		"sematext_sematext_app_hadoopmrv1":    8,
		"sematext_sematext_app_hadoopyarn":    9,
		"sematext_sematext_app_haproxy":       10,
		"sematext_sematext_app_hbase":         11,
		"sematext_sematext_app_infra":         12,
		"sematext_sematext_app_jvm":           13,
		"sematext_sematext_app_kafka":         14,
		"sematext_sematext_app_logsene":       15,
		"sematext_sematext_app_mobilelogs":    16,
		"sematext_sematext_app_mongodb":       17,
		"sematext_sematext_app_mysql":         18,
		"sematext_sematext_app_nginx":         19,
		"sematext_sematext_app_nginxplus":     20,
		"sematext_sematext_app_nodejs":        21,
		"sematext_sematext_app_postgresql":    22,
		"sematext_sematext_app_rabbitmq":      23,
		"sematext_sematext_app_redis":         24,
		"sematext_sematext_app_solr":          25,
		"sematext_sematext_app_solrcloud":     26,
		"sematext_sematext_app_spark":         27,
		"sematext_sematext_app_storm":         28,
		"sematext_sematext_app_tomcat":        29,
		"sematext_sematext_app_zookeeper":     30,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
