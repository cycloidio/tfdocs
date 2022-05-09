package valtix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "valtix_address_object",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Address Object ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "address_id",
					Description: `Set to the Terraform Address Object ID of the found resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address_id",
					Description: `Set to the Terraform Address Object ID of the found resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_policy_rule_set",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Policy Ruleset ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `Set to the Terraform Policy Ruleset ID of the found resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rule_set_id",
					Description: `Set to the Terraform Policy Ruleset ID of the found resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "valtix_service_vpc",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "csp_account_name",
					Description: `(Required) Valtix friendly name for the on-boarded CSP Account where the Service VPC/VNet resource exists`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) CSP Region where the Service VPC/VNet resource exists`,
				},
				resource.Attribute{
					Name:        "vpc_id",
					Description: `(Required) The VPC/VNet ID of the Service VPC/VNet resource ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Set to the Terraform Service VPC ID of the found resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Set to the Terraform Service VPC ID of the found resource`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"valtix_address_object":  0,
		"valtix_policy_rule_set": 1,
		"valtix_service_vpc":     2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
