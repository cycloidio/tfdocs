package astra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "astra_access_list",
			Category:         "Data Sources",
			ShortDescription: `astra_access_list provides a datasource that lists the access lists for an Astra database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
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
			Type:             "astra_private_link_endpoints",
			Category:         "Data Sources",
			ShortDescription: `astra_private_link_endpoints provides a datasource that lists the private link endpoints for an Astra database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_private_links",
			Category:         "Data Sources",
			ShortDescription: `astra_private_links provides a datasource that lists the private links in an Astra database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "astra_role",
			Category:         "Data Sources",
			ShortDescription: `astra_role provides a datasource that lists the custom roles for an org.`,
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

		"astra_access_list":               0,
		"astra_available_regions":         1,
		"astra_database":                  2,
		"astra_databases":                 3,
		"astra_keyspace":                  4,
		"astra_keyspaces":                 5,
		"astra_private_link_endpoints":    6,
		"astra_private_links":             7,
		"astra_role":                      8,
		"astra_secure_connect_bundle_url": 9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
