package atlas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "atlas_migration",
			Category:         "Resources",
			ShortDescription: `The resource applies pending migration files on the connected database.See https://atlasgo.io/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "atlas_schema",
			Category:         "Resources",
			ShortDescription: `Atlas database resource manages the data schema of the database, using an HCL file describing the wanted state of the database. See https://atlasgo.io/`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"atlas_migration": 0,
		"atlas_schema":    1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
