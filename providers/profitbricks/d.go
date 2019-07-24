package profitbricks

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_datacenter",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ProfitBricks Data Centers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name or part of the name of an existing Virtual Data Center that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing Virtual Data Center's location. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_image",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ProfitBricks Images`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name or part of the name of an existing image that you want to search for.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Version of the image (see details below).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing image's location.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The image type, HDD or CD-ROM. If both "name" and "version" are provided the plugin will concatenate the two strings in this format [name]-[version]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_location",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ProfitBricks Locations`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name or part of the location name to search for.`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `(Optional) A desired feature that the location must be able to provide. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the location`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the location`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_resource",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ProfitBricks Resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The specific type of resources to retrieve information about.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) The ID of the specific resource to retrieve information about. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_snapshot",
			Category:         "Data Sources",
			ShortDescription: `Get information on a ProfitBricks Snapshots`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name or part of the name of an existing snapshot that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing snapshot's location.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the snapshot to look for. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the snapshot`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the snapshot`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"profitbricks_datacenter": 0,
		"profitbricks_image":      1,
		"profitbricks_location":   2,
		"profitbricks_resource":   3,
		"profitbricks_snapshot":   4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
