package quorum

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_genesis_mixhash",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to obtain ` + "`" + `MixHash` + "`" + ` value being used in the genesis file. Especially when using Istanbul consensus algorithm which defines a constant digest value.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_node_key",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to obtain various information parsed from an existing node key in hex format. Node key encodes a private key that defines an identity of a Quorum node in the network. It is primarily used in P2P networking.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"quorum_bootstrap_genesis_mixhash": 0,
		"quorum_bootstrap_node_key":        1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
