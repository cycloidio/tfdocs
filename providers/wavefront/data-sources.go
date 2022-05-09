package wavefront

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "wavefront_default_user_group",
			Category:         "Data Sources",
			ShortDescription: `Get the default user group ` + "`" + `Everyone` + "`" + ` from Wavefront`,
			Description: `

Use this data source to get the Group ID of the ` + "`" + `Everyone` + "`" + ` group in Wavefront. 

`,
			Keywords:  []string{},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `Set to the Group ID of the ` + "`" + `Everyone` + "`" + ` group, suitable for referencing in other resources that support group memberships.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"wavefront_default_user_group": 0,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
