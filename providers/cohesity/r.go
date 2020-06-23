package cohesity

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cohesity_cloud_edition_cluster",
			Category:         "Resources",
			ShortDescription: `Create cloud edition cluster, apply license key and destroy cluster.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"edition",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_physical_edition_cluster",
			Category:         "Resources",
			ShortDescription: `Create physical edition cluster, apply license key and destroy cluster.`,
			Description:      ``,
			Keywords: []string{
				"physical",
				"edition",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_virtual_edition_cluster",
			Category:         "Resources",
			ShortDescription: `Create virtual edition cluster, apply license key and destroy cluster.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"edition",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cohesity_cloud_edition_cluster":    0,
		"cohesity_physical_edition_cluster": 1,
		"cohesity_virtual_edition_cluster":  2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
