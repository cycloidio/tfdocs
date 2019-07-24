package rancher2

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rancher2_project",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) ID of the Rancher 2 cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The project name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster-wide unique ID of the Rancher 2 project.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the project as stored by Rancher 2.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The project's description.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `Annotations of the rancher2 project (map).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels of the rancher2 project (map).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster-wide unique ID of the Rancher 2 project.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `UUID of the project as stored by Rancher 2.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The project's description.`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `Annotations of the rancher2 project (map).`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `Labels of the rancher2 project (map).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher2_setting",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher v2 setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The setting name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `the settting's value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `the settting's value.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"rancher2_project": 0,
		"rancher2_setting": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
