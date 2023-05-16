package cockroach

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cockroach_cluster",
			Category:         "Data Sources",
			ShortDescription: `Cluster Data Source`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_cluster_cert",
			Category:         "Data Sources",
			ShortDescription: `TLS certificate for the specified CockroachDB cluster. Certificates for dedicated clusters should be written to $HOME/Library/CockroachCloud/certs/<cluster name>-ca.crt on MacOS or Linux, or $env:appdata\CockroachCloud\certs\<cluster name>-ca.crt on Windows. Serverless clusters use the root PostgreSQL CA cert. If it isn't already installed, the certificate can be appended to $HOME/.postgresql/root.crt on MacOS or Linux, or $env:appdata\postgresql\root.crt on Windows.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_connection_string",
			Category:         "Data Sources",
			ShortDescription: `Generic connection string for a given cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_organization",
			Category:         "Data Sources",
			ShortDescription: `Information about the organization associated with the user's API key`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_person_user",
			Category:         "Data Sources",
			ShortDescription: `Information about a person user`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"cockroach_cluster":           0,
		"cockroach_cluster_cert":      1,
		"cockroach_connection_string": 2,
		"cockroach_organization":      3,
		"cockroach_person_user":       4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
