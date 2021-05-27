package astra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "astra_database",
			Category:         "Resources",
			ShortDescription: `astra_database provides an Astra Serverless Database resource. You can create and delete databases. Note: Classic Tier databases are not supported by the Terraform provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_keyspace",
			Category:         "Resources",
			ShortDescription: `astra_keyspace provides a keyspace resource. Keyspaces are groupings of tables for Cassandra. astra_keyspace resources are associated with a database id. You can have multiple keyspaces per DB in addition to the default keyspace provided in the astra_database resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"astra_database": 0,
		"astra_keyspace": 1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
