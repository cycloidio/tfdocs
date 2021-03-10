package configcat

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "configcat_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Config. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique Config ID. ## Import Configs can be imported using the ConfigId. Get the ConfigId using the [GetConfigs API](https://api.configcat.com/docs/#operation/get-configs) for example. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_config.example 08d86d63-2726-47cd-8bfc-59608ecb91e2 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique Config ID. ## Import Configs can be imported using the ConfigId. Get the ConfigId using the [GetConfigs API](https://api.configcat.com/docs/#operation/get-configs) for example. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_config.example 08d86d63-2726-47cd-8bfc-59608ecb91e2 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_environment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Environment. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique Environment ID. ## Import Environments can be imported using the EnvironmentId. Get the EnvironmentId using the [GetEnvironments API](https://api.configcat.com/docs/#operation/get-environments) for example. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_environment.example 08d86d63-2726-47cd-8bfc-59608ecb91e2 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique Environment ID. ## Import Environments can be imported using the EnvironmentId. Get the EnvironmentId using the [GetEnvironments API](https://api.configcat.com/docs/#operation/get-environments) for example. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_environment.example 08d86d63-2726-47cd-8bfc-59608ecb91e2 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_product",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "organization_id",
					Description: `(Required) The ID of the Organization.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Product. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique Product ID. ## Import Products can be imported using the ProductId. Get the ProductId using the [GetProducts API](https://api.configcat.com/docs/#operation/get-products) for example. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_product.example 08d86d63-2726-47cd-8bfc-59608ecb91e2 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique Product ID. ## Import Products can be imported using the ProductId. Get the ProductId using the [GetProducts API](https://api.configcat.com/docs/#operation/get-products) for example. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_product.example 08d86d63-2726-47cd-8bfc-59608ecb91e2 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_setting",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the Config.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The key of the Feature Flag/Setting.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Setting.`,
				},
				resource.Attribute{
					Name:        "hint",
					Description: `(Optional) The hint of the Setting.`,
				},
				resource.Attribute{
					Name:        "setting_type",
					Description: `(Optional) Default: ` + "`" + `boolean` + "`" + `. The Setting's type. Available values: ` + "`" + `boolean` + "`" + `|` + "`" + `string` + "`" + `|` + "`" + `int` + "`" + `|` + "`" + `double` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique Setting ID. ## Import Feature Flags/Settings can be imported using the SettingId. Get the SettingId using e.g. the [GetSettings API](https://api.configcat.com/docs/#operation/get-settings). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_setting.example 1234 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique Setting ID. ## Import Feature Flags/Settings can be imported using the SettingId. Get the SettingId using e.g. the [GetSettings API](https://api.configcat.com/docs/#operation/get-settings). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_setting.example 1234 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_setting_tag",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "setting_id",
					Description: `(Required) The ID of the Feature Flag/Setting.`,
				},
				resource.Attribute{
					Name:        "tag_id",
					Description: `(Required) The ID of the Tag. ## Import Tags can be imported using a combined SettingId:TagId ID. Get the SettingId using e.g. the [GetSettings API](https://api.configcat.com/docs/#operation/get-settings). Get the TagId using e.g. the [GetTags API](https://api.configcat.com/docs/#operation/get-tags). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_setting_tag.example 1234:5678 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_setting_value",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the Environment.`,
				},
				resource.Attribute{
					Name:        "setting_id",
					Description: `(Required) The ID of the Feature Flag/Setting.`,
				},
				resource.Attribute{
					Name:        "setting_type",
					Description: `(Required) The Setting's type.`,
				},
				resource.Attribute{
					Name:        "mandatory_notes",
					Description: `(Optional) Default: "". If the Product's "Mandatory notes" preference is turned on for the Environment the Mandatory note must be passed.`,
				},
				resource.Attribute{
					Name:        "init_only",
					Description: `(Optional) Default: true. Read more below. The Feature Flag/Setting's value`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The Setting's value. Type: ` + "`" + `string` + "`" + `. It must be compatible with the ` + "`" + `setting_type` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rollout_rules",
					Description: `(Optional) A [list](https://www.terraform.io/docs/configuration/types.html#list-) to define [Rollout rules](https://configcat.com/docs/advanced/targeting/#anatomy-of-a-targeting-rule). Read more below.`,
				},
				resource.Attribute{
					Name:        "percentage_items",
					Description: `(Optional) A [list](https://www.terraform.io/docs/configuration/types.html#list-) to define [Percentage items](https://configcat.com/docs/advanced/targeting/#targeting-a-percentage-of-users). Read more below. ### ` + "`" + `rollout_rules` + "`" + ` list By adding a rule, you specify a group of your users and what feature flag - or other settings - value they should get.`,
				},
				resource.Attribute{
					Name:        "comparison_attribute",
					Description: `(Required) The [comparison attribute](https://configcat.com/docs/advanced/targeting/#attribute).`,
				},
				resource.Attribute{
					Name:        "comparator",
					Description: `(Required) The [comparator](https://configcat.com/docs/advanced/targeting/#comparator).`,
				},
				resource.Attribute{
					Name:        "comparison_value",
					Description: `(Required) The [comparison value](https://configcat.com/docs/advanced/targeting/#comparison-value).`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The exact [value](https://configcat.com/docs/advanced/targeting/#served-value) that will be served to the users who match the targeting rule. Type: ` + "`" + `string` + "`" + `. It must be compatible with the ` + "`" + `setting_type` + "`" + `. ### ` + "`" + `percentage_items` + "`" + ` list With percentage-based user targeting, you can specify a randomly selected fraction of your users whom a feature will be enabled or a different value will be served.`,
				},
				resource.Attribute{
					Name:        "percentage",
					Description: `(Required) Any [number](https://configcat.com/docs/advanced/targeting/#-value) between 0 and 100 that represents a randomly allocated fraction of your users.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The exact [value](https://configcat.com/docs/advanced/targeting/#served-value-1) that will be served to the users that fall into that fraction. Type: ` + "`" + `string` + "`" + `. It must be compatible with the ` + "`" + `setting_type` + "`" + `. ### ` + "`" + `init_only` + "`" + ` argument The main purpose of this resource to provide an initial value for the Feature Flag/Setting. The ` + "`" + `init_only` + "`" + ` argument's default value is ` + "`" + `true` + "`" + `. Meaning that the Feature Flag or Setting's`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_tag",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Tag.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `(Optional) Default: ` + "`" + `panther` + "`" + `. The color of the Tag. Valid values: ` + "`" + `panther` + "`" + `, ` + "`" + `whale` + "`" + `, ` + "`" + `salmon` + "`" + `, ` + "`" + `lizard` + "`" + `, ` + "`" + `canary` + "`" + `, ` + "`" + `koala` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique Tag ID. ## Import Tags can be imported using the TagId. Get the TagId using e.g. the [GetTags API](https://api.configcat.com/docs/#operation/get-tags). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_tag.example 1234 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique Tag ID. ## Import Tags can be imported using the TagId. Get the TagId using e.g. the [GetTags API](https://api.configcat.com/docs/#operation/get-tags). ` + "`" + `` + "`" + `` + "`" + ` $ terraform import configcat_tag.example 1234 ` + "`" + `` + "`" + `` + "`" + ` [Read more](https://learn.hashicorp.com/tutorials/terraform/state-import) about importing. ## Endpoints used`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"configcat_config":        0,
		"configcat_environment":   1,
		"configcat_product":       2,
		"configcat_setting":       3,
		"configcat_setting_tag":   4,
		"configcat_setting_value": 5,
		"configcat_tag":           6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
