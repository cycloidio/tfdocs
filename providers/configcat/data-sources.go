package configcat

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "configcat_configs",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name_filter_regex",
					Description: `(Optional) Filter the Configs by name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "configs",
					Description: `A config [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `configs` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `The unique Config ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Config. ## Endpoints used [Get Configs](https://api.configcat.com/docs/index.html#operation/get-configs)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "configs",
					Description: `A config [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `configs` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "config_id",
					Description: `The unique Config ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Config. ## Endpoints used [Get Configs](https://api.configcat.com/docs/index.html#operation/get-configs)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_environments",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name_filter_regex",
					Description: `(Optional) Filter the Environments by name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "environments",
					Description: `An environment [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `environments` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `The unique Environment ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Environment. ## Endpoints used - [Get Environments](https://api.configcat.com/docs/index.html#operation/get-environments)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "environments",
					Description: `An environment [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `environments` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `The unique Environment ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Environment. ## Endpoints used - [Get Environments](https://api.configcat.com/docs/index.html#operation/get-environments)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_organizations",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_filter_regex",
					Description: `(Optional) Filter the Organizations by name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "organizations",
					Description: `An organization [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `organizations` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The unique Organization ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organization. ## Endpoints used - [Get Organizations](https://api.configcat.com/docs/index.html#operation/get-organizations)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "organizations",
					Description: `An organization [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `organizations` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The unique Organization ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Organization. ## Endpoints used - [Get Organizations](https://api.configcat.com/docs/index.html#operation/get-organizations)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_products",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name_filter_regex",
					Description: `(Optional) Filter the Products by name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "products",
					Description: `A product [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `products` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `The unique Product ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Product. ## Endpoints used - [Get Products](https://api.configcat.com/docs/index.html#operation/get-products)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "products",
					Description: `A product [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `products` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "product_id",
					Description: `The unique Product ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Product. ## Endpoints used - [Get Products](https://api.configcat.com/docs/index.html#operation/get-products)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_settings",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config_id",
					Description: `(Required) The ID of the Config.`,
				},
				resource.Attribute{
					Name:        "key_filter_regex",
					Description: `(Optional) Filter the Settings by key. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "settings",
					Description: `A setting [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `settings` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "setting_id",
					Description: `The unique Setting ID.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the Feature Flag/Setting.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Setting.`,
				},
				resource.Attribute{
					Name:        "hint",
					Description: `The hint of the Setting.`,
				},
				resource.Attribute{
					Name:        "setting_type",
					Description: `The Setting's type. Available values: ` + "`" + `boolean` + "`" + `|` + "`" + `string` + "`" + `|` + "`" + `int` + "`" + `|` + "`" + `double` + "`" + `. ## Endpoints used - [Get Settings](https://api.configcat.com/docs/index.html#operation/get-settings)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "settings",
					Description: `A setting [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `settings` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "setting_id",
					Description: `The unique Setting ID.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key of the Feature Flag/Setting.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Setting.`,
				},
				resource.Attribute{
					Name:        "hint",
					Description: `The hint of the Setting.`,
				},
				resource.Attribute{
					Name:        "setting_type",
					Description: `The Setting's type. Available values: ` + "`" + `boolean` + "`" + `|` + "`" + `string` + "`" + `|` + "`" + `int` + "`" + `|` + "`" + `double` + "`" + `. ## Endpoints used - [Get Settings](https://api.configcat.com/docs/index.html#operation/get-settings)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "configcat_tags",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "product_id",
					Description: `(Required) The ID of the Product.`,
				},
				resource.Attribute{
					Name:        "name_filter_regex",
					Description: `(Optional) Filter the Tags by name. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A tag [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `tags` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "tag_id",
					Description: `The unique Tag ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Tag.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `The color of the Tag. ## Endpoints used - [Get Tags](https://api.configcat.com/docs/index.html#operation/get-tags)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "tags",
					Description: `A tag [list](https://www.terraform.io/docs/configuration/types.html#list-) block defined as below. ### The ` + "`" + `tags` + "`" + ` [list](https://www.terraform.io/docs/configuration/types.html#list-) block`,
				},
				resource.Attribute{
					Name:        "tag_id",
					Description: `The unique Tag ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Tag.`,
				},
				resource.Attribute{
					Name:        "color",
					Description: `The color of the Tag. ## Endpoints used - [Get Tags](https://api.configcat.com/docs/index.html#operation/get-tags)`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"configcat_configs":       0,
		"configcat_environments":  1,
		"configcat_organizations": 2,
		"configcat_products":      3,
		"configcat_settings":      4,
		"configcat_tags":          5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
