package chef

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "chef_data_bag",
			Category:         "Resources",
			ShortDescription: `Creates and manages a data bag in Chef Server.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"bag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name to assign to the data bag. This is the name that other server clients will use to find and retrieve data from the data bag. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "api_uri",
					Description: `The URI representing this data bag in the Chef server API.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "api_uri",
					Description: `The URI representing this data bag in the Chef server API.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "chef_data_bag_item",
			Category:         "Resources",
			ShortDescription: `Creates and manages an object within a data bag in Chef Server.`,
			Description:      ``,
			Keywords: []string{
				"data",
				"bag",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_bag_name",
					Description: `(Required) The name of the data bag into which this item will be placed.`,
				},
				resource.Attribute{
					Name:        "content_json",
					Description: `(Required) A string containing a JSON object that will be the content of the item. Must at minimum contain a property called "id" that is unique within the data bag, which will become the identifier of the created item. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The value of the "id" property in the ` + "`" + `` + "`" + `content_json` + "`" + `` + "`" + ` JSON object, which can be used by clients to retrieve this item's content.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The value of the "id" property in the ` + "`" + `` + "`" + `content_json` + "`" + `` + "`" + ` JSON object, which can be used by clients to retrieve this item's content.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "chef_environment",
			Category:         "Resources",
			ShortDescription: `Creates and manages an environment in Chef Server.`,
			Description:      ``,
			Keywords: []string{
				"environment",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name to assign to the environment. This name will be used when nodes are created within the environment.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the environment. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "default_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the default attributes for the environment.`,
				},
				resource.Attribute{
					Name:        "override_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the override attributes for the environment.`,
				},
				resource.Attribute{
					Name:        "cookbook_constraints",
					Description: `(Optional) Mapping of cookbook names to cookbook version constraints that should apply for this environment. ## Attributes Reference This resource exports no further attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "chef_node",
			Category:         "Resources",
			ShortDescription: `Creates and manages a node in Chef Server.`,
			Description:      ``,
			Keywords: []string{
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name to assign to the node.`,
				},
				resource.Attribute{
					Name:        "automatic_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the automatic attributes for the node.`,
				},
				resource.Attribute{
					Name:        "normal_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the normal attributes for the node.`,
				},
				resource.Attribute{
					Name:        "default_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the default attributes for the node.`,
				},
				resource.Attribute{
					Name:        "override_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the override attributes for the node.`,
				},
				resource.Attribute{
					Name:        "run_list",
					Description: `(Optional) List of strings to set as the [run list](https://docs.chef.io/run_lists.html) for the node. ## Attributes Reference This resource exports no further attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "chef_role",
			Category:         "Resources",
			ShortDescription: `Creates and manages a role in Chef Server.`,
			Description:      ``,
			Keywords: []string{
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The unique name to assign to the role.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description of the role. If not set, a placeholder of "Managed by Terraform" will be set.`,
				},
				resource.Attribute{
					Name:        "default_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the default attributes for the role.`,
				},
				resource.Attribute{
					Name:        "override_attributes_json",
					Description: `(Optional) String containing a JSON-serialized object containing the override attributes for the role.`,
				},
				resource.Attribute{
					Name:        "run_list",
					Description: `(Optional) List of strings to set as the [run list](https://docs.chef.io/run_lists.html) for any nodes that belong to this role. ## Attributes Reference This resource exports no further attributes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"chef_data_bag":      0,
		"chef_data_bag_item": 1,
		"chef_environment":   2,
		"chef_node":          3,
		"chef_role":          4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
