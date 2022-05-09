package azapi

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azapi_azapi_resource",
			Category:         "Data Sources",
			ShortDescription: `Gets information from an existing azure resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the azure resource.`,
				},
				resource.Attribute{
					Name:        "identity",
					Description: `An ` + "`" + `identity` + "`" + ` block as defined below, which contains the Managed Service Identity information for this azure resource.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The Azure Region where the azure resource should exist.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `The output json containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here're some examples to decode json and extract the value. ` + "`" + `` + "`" + `` + "`" + ` // it will output "registry1.azurecr.io" output "login_server" { value = jsondecode(azapi_resource.example.output).properties.loginServer } // it will output "disabled" output "quarantine_policy" { value = jsondecode(azapi_resource.example.output).properties.policies.quarantinePolicy.status } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A mapping of tags which should be assigned to the azure resource. --- A ` + "`" + `identity` + "`" + ` block exports the following:`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The Type of Identity which should be used for this azure resource. Possible values are ` + "`" + `SystemAssigned` + "`" + `, ` + "`" + `UserAssigned` + "`" + ` and ` + "`" + `SystemAssigned,UserAssigned` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "identity_ids",
					Description: `A list of User Managed Identity ID's which should be assigned to the azure resource.`,
				},
				resource.Attribute{
					Name:        "principal_id",
					Description: `The Principal ID for the Service Principal associated with the Managed Service Identity of this azure resource.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The Tenant ID for the Service Principal associated with the Managed Service Identity of this azure resource. ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the azure resource.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"azapi_azapi_resource": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
