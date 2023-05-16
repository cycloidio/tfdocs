package powerflex

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerflex_sdc",
			Category:         "Resources",
			ShortDescription: `This resource can be used to manage Storage Data Clients on a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_sdc_volumes_mapping",
			Category:         "Resources",
			ShortDescription: `This resource can be used to manage mapping of volumes to an SDC on a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_sds",
			Category:         "Resources",
			ShortDescription: `This resource can be used to manage Storage Data Servers on a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_snapshot",
			Category:         "Resources",
			ShortDescription: `This resource can be used to manage snapshots of volumes on a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_storage_pool",
			Category:         "Resources",
			ShortDescription: `This resource can be used to manage Storage Pools on a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerflex_volume",
			Category:         "Resources",
			ShortDescription: `This resource can be used to manage volumes on a PowerFlex array.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"powerflex_sdc":                 0,
		"powerflex_sdc_volumes_mapping": 1,
		"powerflex_sds":                 2,
		"powerflex_snapshot":            3,
		"powerflex_storage_pool":        4,
		"powerflex_volume":              5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
