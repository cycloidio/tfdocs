package vra7

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vra7_deployment",
			Category:         "Deployment Resources",
			ShortDescription: `Provides a VMware vRA7 deployment resource. This can be used to deploy vRA7 catalog items.`,
			Description:      ``,
			Keywords: []string{
				"deployment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "businessgroup_id",
					Description: `(Optional) The id of the vRA business group to use for this deployment. Either businessgroup_id or businessgroup_name is required.`,
				},
				resource.Attribute{
					Name:        "businessgroup_name",
					Description: `(Optional) The name of the vRA business group to use for this deployment. Either businessgroup_id or businessgroup_name is required.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `(Optional) The id of the catalog item to deploy into vRA. Either catalog_item_id or catalog_item_name is required.`,
				},
				resource.Attribute{
					Name:        "catalog_item_name",
					Description: `(Optional) The name of the catalog item to deploy into vRA. Either catalog_item_id or catalog_item_name is required.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the deployment.`,
				},
				resource.Attribute{
					Name:        "reasons",
					Description: `(Optional) Reasons for requesting the deployment.`,
				},
				resource.Attribute{
					Name:        "deployment_configuration",
					Description: `(Optional) The configuration of the deployment from the catalog item. All blueprint custom properties including property groups can be added to this block. This property is discussed in detail below.`,
				},
				resource.Attribute{
					Name:        "resource_configuration",
					Description: `(Optional) The configuration of the individual components from the catalog item. This property is discussed in detail below.`,
				},
				resource.Attribute{
					Name:        "lease_days",
					Description: `(Optional) Number of lease days remaining for the deployment. NOTE: lease_days 0 means the lease never expires.`,
				},
				resource.Attribute{
					Name:        "wait_timeout",
					Description: `(Optional) Wait time out for the request. If the request is not completed within the timeout period, do a terraform refresh later to check the status of the request. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `The resource id of the deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the deployment.`,
				},
				resource.Attribute{
					Name:        "lease_start",
					Description: `Start date of the lease.`,
				},
				resource.Attribute{
					Name:        "lease_end",
					Description: `End date of the lease.`,
				},
				resource.Attribute{
					Name:        "request_status",
					Description: `The status of the catalog item request.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date when the deployment was created.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The date when the deployment was last updated after Day 2 operations.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The id of the tenant.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `The owners of the deployment. ## Nested Blocks ### resource_configuration ### This is a list of blocks that contains the machine resource level properties including the custom properties. Each resource_configuration block corresponds to a component in the blueprint/catalog_item. The sample blueprint has one vSphere machine resource/component called vSphereVM1. Properties of this machine can be specified in the config as shown in the example above. The properties like cpu, memory, storage, etc are generic machine properties and their is a custom property as well, called machine_property in the sample blueprint which is required at request time. There can be any number of machines and same format has to be followed to specify properties of other machines as well.All the properties that are required during request, must be specified in the config file. The following arguments for resource_configuration block are supported: #### Argument Reference`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) The name of the component/machine resource as in the blueprint/catalog_item`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) The machine resource level properties like cpu, memory, storage, custom properties, etc. can be added here. When fetching the state of the machine, this will be populated with a lot of information in the state file. NOTE: To add an array property, refer to the security_tag value in example above.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Optional) Cluster size for this machine resource #### Attribute Reference`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the machine resource`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the resource`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the resource`,
				},
				resource.Attribute{
					Name:        "parent_resource_id",
					Description: `ID of the deployment of which this machine is a part of`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the machine`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `ID of the catalog item request`,
				},
				resource.Attribute{
					Name:        "request_state",
					Description: `Current state of the request. It can be FAILED, IN_PROGRESS, SUCCESSFUL, etc.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of resource. It can be a machine resource type (Infrastructure.Virtual) or a deployment type (composition.resource.type.deployment), etc.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the machine. It can be On, Off, Unprovisioned, etc.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date when the resource was created.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The date when the resource was last updated after Day 2 operations. ### deployment_configuration ### This block contains the deployment level properties including the custom properties and proprty groups. These are not a fixed set of properties but referred from the blueprint. From the example of the BasicSingleMachine blueprint, their is one custom property, called deployment_property which is required at request time. All the properties that are required during request, must be specified in the config file.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "deployment_id",
					Description: `The resource id of the deployment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the deployment.`,
				},
				resource.Attribute{
					Name:        "lease_start",
					Description: `Start date of the lease.`,
				},
				resource.Attribute{
					Name:        "lease_end",
					Description: `End date of the lease.`,
				},
				resource.Attribute{
					Name:        "request_status",
					Description: `The status of the catalog item request.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date when the deployment was created.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The date when the deployment was last updated after Day 2 operations.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `The id of the tenant.`,
				},
				resource.Attribute{
					Name:        "owners",
					Description: `The owners of the deployment. ## Nested Blocks ### resource_configuration ### This is a list of blocks that contains the machine resource level properties including the custom properties. Each resource_configuration block corresponds to a component in the blueprint/catalog_item. The sample blueprint has one vSphere machine resource/component called vSphereVM1. Properties of this machine can be specified in the config as shown in the example above. The properties like cpu, memory, storage, etc are generic machine properties and their is a custom property as well, called machine_property in the sample blueprint which is required at request time. There can be any number of machines and same format has to be followed to specify properties of other machines as well.All the properties that are required during request, must be specified in the config file. The following arguments for resource_configuration block are supported: #### Argument Reference`,
				},
				resource.Attribute{
					Name:        "component_name",
					Description: `(Required) The name of the component/machine resource as in the blueprint/catalog_item`,
				},
				resource.Attribute{
					Name:        "configuration",
					Description: `(Optional) The machine resource level properties like cpu, memory, storage, custom properties, etc. can be added here. When fetching the state of the machine, this will be populated with a lot of information in the state file. NOTE: To add an array property, refer to the security_tag value in example above.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Optional) Cluster size for this machine resource #### Attribute Reference`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `ID of the machine resource`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the resource`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the resource`,
				},
				resource.Attribute{
					Name:        "parent_resource_id",
					Description: `ID of the deployment of which this machine is a part of`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address of the machine`,
				},
				resource.Attribute{
					Name:        "request_id",
					Description: `ID of the catalog item request`,
				},
				resource.Attribute{
					Name:        "request_state",
					Description: `Current state of the request. It can be FAILED, IN_PROGRESS, SUCCESSFUL, etc.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `Type of resource. It can be a machine resource type (Infrastructure.Virtual) or a deployment type (composition.resource.type.deployment), etc.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the machine. It can be On, Off, Unprovisioned, etc.`,
				},
				resource.Attribute{
					Name:        "date_created",
					Description: `The date when the resource was created.`,
				},
				resource.Attribute{
					Name:        "last_updated",
					Description: `The date when the resource was last updated after Day 2 operations. ### deployment_configuration ### This block contains the deployment level properties including the custom properties and proprty groups. These are not a fixed set of properties but referred from the blueprint. From the example of the BasicSingleMachine blueprint, their is one custom property, called deployment_property which is required at request time. All the properties that are required during request, must be specified in the config file.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vra7_deployment": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
