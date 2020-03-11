package quorum

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_account",
			Category:         "Resources",
			ShortDescription: `Use this resource to create a new Ethereum account`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_data_dir",
			Category:         "Resources",
			ShortDescription: `Use this resource to create a data dir locally. This equivalent to execute ` + "`" + `geth init` + "`" + `.`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
				"data",
				"dir",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_istanbul_extradata",
			Category:         "Resources",
			ShortDescription: `Use this resource to construct ` + "`" + `extradata` + "`" + ` field used in the genesis file. ` + "`" + `istanbul_address` + "`" + ` can be referenced from ` + "`" + `quorum_bootstrap_node_key` + "`" + ` data source or newly created from ` + "`" + `quorum_bootstrap_node_key` + "`" + ` resources.`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
				"istanbul",
				"extradata",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_keystore",
			Category:         "Resources",
			ShortDescription: `Use this resource to create a keystore which maintains multiple Ethereum accounts.`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
				"keystore",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_network",
			Category:         "Resources",
			ShortDescription: `Use this resource to create a new directory that represents a new Quorum network. Bootstraping data will be kept in this folder.`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
				"network",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_bootstrap_node_key",
			Category:         "Resources",
			ShortDescription: `Use this resource to create a node key for a new Quorum node. Node key encodes a private key that defines an identity of a Quorum node in the network. It is primarily used in P2P networking.`,
			Description:      ``,
			Keywords: []string{
				"bootstrap",
				"node",
				"key",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "quorum_transaction_manager_keypair",
			Category:         "Resources",
			ShortDescription: `Use this resource to create a key pair used in a transaction manager. This key pair provides attributes which are useful when building the configuration for a transaction manager.`,
			Description:      ``,
			Keywords: []string{
				"transaction",
				"manager",
				"keypair",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"quorum_bootstrap_account":            0,
		"quorum_bootstrap_data_dir":           1,
		"quorum_bootstrap_istanbul_extradata": 2,
		"quorum_bootstrap_keystore":           3,
		"quorum_bootstrap_network":            4,
		"quorum_bootstrap_node_key":           5,
		"quorum_transaction_manager_keypair":  6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
