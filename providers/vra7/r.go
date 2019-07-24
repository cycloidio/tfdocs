package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vra7_deployment",
			Category:         "Deployment Resources",
			ShortDescription: `Provides a VMware vRA7 deployment resource. This can be used to deploy vRA7 catalog items.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "businessgroup_id",
					Description: `(Optional) The id of the vRA business group to use for this deployment.`,
				},
				resource.Attribute{
					Name:        "businessgroup_name",
					Description: `(Optional) The name of the vRA business group to use for this deployment.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `(Optional) The id of the catalog item to deploy into vRA.`,
				},
				resource.Attribute{
					Name:        "catalog_item_name",
					Description: `(Optional) The name of the catalog item to deploy into vRA.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the deployment`,
				},
				resource.Attribute{
					Name:        "reasons",
					Description: `(Optional) Reasons for requesting the deployment`,
				},
				resource.Attribute{
					Name:        "deployment_configuration",
					Description: `(Optional) The configuration of the deployment from the catalog item`,
				},
				resource.Attribute{
					Name:        "resource_configuration",
					Description: `(Optional) The configuration of the individual components from the catalog item ## Nested Blocks ### deployment_configuration ### This block contains the deployment level properties including the custom properties. These are not a fixed set of properties but referred from the blueprint. There are generic properties like _leaseDays, _number_of_instances, etc but they are optional and from the example of the BasicSingleMachine blueprint, their is one custom property, called deployment_property which is required at request time. All the properties that are required during request, must be specified in the config file. ### resource_configuration ### This block contains the machine resource level properties including the custom properties. These are not a fixed set of properties but referred from the blueprint. The sample blueprint has one vSphere machine resource called vSphereVM1. Properties of this machine can be specified in the config in the format "vSphereVM1.property_name". The properties like cpu, memory, storage, etc are generic machine properties and their is a custom property as well, called machine_property in the sample blueprint which is required at request time. There can be any number of machines and same format has to be followed to specify properties of other machines as well. All the properties that are required during request, must be specified in the config file. ### More examples ###`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"vra7_deployment": 0,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
