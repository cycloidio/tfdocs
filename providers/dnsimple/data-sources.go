package dnsimple

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_certificate",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain of the SSL Certificate`,
				},
				resource.Attribute{
					Name:        "certificate_id",
					Description: `(Required) The ID of the SSL Certificate ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "server_certificate",
					Description: `The SSL Certificate`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `The Root Certificate of the issuing CA`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `A list of certificates that make up the chain`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The corresponding Private Key for the SSL Certificate`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "server_certificate",
					Description: `The SSL Certificate`,
				},
				resource.Attribute{
					Name:        "root_certificate",
					Description: `The Root Certificate of the issuing CA`,
				},
				resource.Attribute{
					Name:        "certificate_chain",
					Description: `A list of certificates that make up the chain`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The corresponding Private Key for the SSL Certificate`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "dnsimple_zone",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"dnsimple_certificate": 0,
		"dnsimple_zone":        1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
