package azapi

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "azapi_azapi_resource",
			Category:         "Resources",
			ShortDescription: `Manages a Azure resource`,
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
					Name:        "output",
					Description: `The output json containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here're some examples to decode json and extract the value. ` + "`" + `` + "`" + `` + "`" + ` // it will output "registry1.azurecr.io" output "login_server" { value = jsondecode(azapi_resource.example.output).properties.loginServer } // it will output "disabled" output "quarantine_policy" { value = jsondecode(azapi_resource.example.output).properties.policies.quarantinePolicy.status } ` + "`" + `` + "`" + `` + "`" + ` --- A ` + "`" + `identity` + "`" + ` block exports the following:`,
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
					Name:        "create",
					Description: `(Defaults to 30 minutes) Used when creating the azure resource.`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the azure resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 minutes) Used when deleting the azure resource. ## Import Azure resource can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azapi_resource.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1 ` + "`" + `` + "`" + `` + "`" + ` It also supports specifying API version by using the ` + "`" + `resource id` + "`" + ` with ` + "`" + `api-version` + "`" + ` as a query parameter, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import azapi_resource.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/resGroup1/providers/Microsoft.MachineLearningServices/workspaces/workspace1/computes/cluster1?api-version=2021-07-01 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azapi_azapi_resource_action",
			Category:         "Resources",
			ShortDescription: `Perform resource action which changes an existing resource's state`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the azure resource action.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `The output json containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here are some examples to decode json and extract the value. ` + "`" + `` + "`" + `` + "`" + `hcl // it will output "nHGYNd`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 minutes) Used when creating the azure resource.`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the azure resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 minutes) Used when deleting the azure resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "azapi_azapi_update_resource",
			Category:         "Resources",
			ShortDescription: `Manages a subset of an existing azure resource's properties`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the azure resource.`,
				},
				resource.Attribute{
					Name:        "output",
					Description: `The output json containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here're some examples to decode json and extract the value. ` + "`" + `` + "`" + `` + "`" + ` // it will output "registry1.azurecr.io" output "login_server" { value = jsondecode(azapi_resource.example.output).properties.loginServer } // it will output "disabled" output "quarantine_policy" { value = jsondecode(azapi_resource.example.output).properties.policies.quarantinePolicy.status } ` + "`" + `` + "`" + `` + "`" + ` ## Timeouts The ` + "`" + `timeouts` + "`" + ` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `(Defaults to 30 minutes) Used when creating the azure resource.`,
				},
				resource.Attribute{
					Name:        "read",
					Description: `(Defaults to 5 minutes) Used when retrieving the azure resource.`,
				},
				resource.Attribute{
					Name:        "delete",
					Description: `(Defaults to 30 minutes) Used when deleting the azure resource.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"azapi_azapi_resource":        0,
		"azapi_azapi_resource_action": 1,
		"azapi_azapi_update_resource": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
