package turbot

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "turbot_control",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `\_control

This data source can be used to fetch information about a specific control.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the control.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of the control.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Optional) The unique identifier of the resource which the control is targeting.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the control.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Message explaining the state of the control.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Additional information regarding the control state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags set on the control.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `The state of the control.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Message explaining the state of the control.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Additional information regarding the control state.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Tags set on the control.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_policy",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `\_policy\_value

This data source can be used to fetch information about a specific policy
setting.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The unique identifier of the policy for which the value needs to be extracted.`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `(Required) The unique ID of the resource at the level of which the information needs to be fetched. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The value that the policy is set to.`,
				},
				resource.Attribute{
					Name:        "value_source",
					Description: `The values for the policy derived from the template.`,
				},
				resource.Attribute{
					Name:        "precedence",
					Description: `The priority level of the policy.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The final state of the set policy.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Message explaining the state of the set policy.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Additional information regarding the set policy.`,
				},
				resource.Attribute{
					Name:        "setting_id",
					Description: `The unique id of the the policy setting.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `The value that the policy is set to.`,
				},
				resource.Attribute{
					Name:        "value_source",
					Description: `The values for the policy derived from the template.`,
				},
				resource.Attribute{
					Name:        "precedence",
					Description: `The priority level of the policy.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The final state of the set policy.`,
				},
				resource.Attribute{
					Name:        "reason",
					Description: `Message explaining the state of the set policy.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Additional information regarding the set policy.`,
				},
				resource.Attribute{
					Name:        "setting_id",
					Description: `The unique id of the the policy setting.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "turbot_resource",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description: `
This data source can be used to fetch information about a specific resource.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The unique identifier of the resource. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "data",
					Description: `JSON representation of the details of the resource. When parsed, it must be valid for the type schema.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A set of data that describes and gives information about the data of the resource`,
				},
				resource.Attribute{
					Name:        "akas",
					Description: `A list of akas for the resource`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `User defined way of logically grouping resources.`,
				},
				resource.Attribute{
					Name:        "turbot",
					Description: `JSON representation of turbot data of the resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "data",
					Description: `JSON representation of the details of the resource. When parsed, it must be valid for the type schema.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A set of data that describes and gives information about the data of the resource`,
				},
				resource.Attribute{
					Name:        "akas",
					Description: `A list of akas for the resource`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `User defined way of logically grouping resources.`,
				},
				resource.Attribute{
					Name:        "turbot",
					Description: `JSON representation of turbot data of the resource.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"turbot_control":  0,
		"turbot_policy":   1,
		"turbot_resource": 2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
