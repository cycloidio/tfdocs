package astra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "astra_available_regions",
			Category:         "Data Sources",
			ShortDescription: `Retrieve a list of available cloud regions in Astra`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_database",
			Category:         "Data Sources",
			ShortDescription: `astra_database provides a datasource for Astra an Astra database. This can be used to select an existing database within your Astra Organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_databases",
			Category:         "Data Sources",
			ShortDescription: `astra_databases provides a datasource for a list of Astra databases. This can be used to select databases within your Astra Organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_keyspace",
			Category:         "Data Sources",
			ShortDescription: `astra_keyspace provides a datasource for a particular keyspace. See astra_keyspaces if you're looking to fetch all the keyspaces for a particular database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_keyspaces",
			Category:         "Data Sources",
			ShortDescription: `astra_keyspaces provides a datasource that lists the keyspaces in an Astra database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_secure_connect_bundle_url",
			Category:         "Data Sources",
			ShortDescription: `astra_secure_connect_bundle_url provides a datasource that generates a temporary secure connect bundle URL. This URL lasts five minutes. Secure connect bundles are used to connect to Astra using cql cassandra drivers. See the docs https://docs.datastax.com/en/astra/docs/connecting-to-database.html for more information on how to connect.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"astra_available_regions":         0,
		"astra_database":                  1,
		"astra_databases":                 2,
		"astra_keyspace":                  3,
		"astra_keyspaces":                 4,
		"astra_secure_connect_bundle_url": 5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
