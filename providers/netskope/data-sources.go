package netskope

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "netskope_ipsec_pops",
			Category:         "Data Sources",
			ShortDescription: `IPSec PoPs are not yet supported.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netskope_ipsec_tunnels",
			Category:         "Data Sources",
			ShortDescription: `IPSec Tunnels are not yet supported.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netskope_privateapps",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about Private Apps deployed in a Netskope tenant. ---`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "netskope_publishers",
			Category:         "Data Sources",
			ShortDescription: `Retrieves information about Publishers deployed in a Netskope tenant.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"netskope_ipsec_pops":    0,
		"netskope_ipsec_tunnels": 1,
		"netskope_privateapps":   2,
		"netskope_publishers":    3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
