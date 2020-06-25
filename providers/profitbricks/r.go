package profitbricks

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_datacenter",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Virtual Data Center.`,
			Description:      ``,
			Keywords: []string{
				"datacenter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required)[string] The name of the Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required)[string] The regional location where the Virtual Data Center will be created.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)[string] Description for the Virtual Data Center. ## Import Resource Datacenter can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import profitbricks_datacenter.mydc {datacenter uuid} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_firewall",
			Category:         "Resources",
			ShortDescription: `Creates and manages Firewall Rules.`,
			Description:      ``,
			Keywords: []string{
				"firewall",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_group",
			Category:         "Resources",
			ShortDescription: `Creates and manages group objects.`,
			Description:      ``,
			Keywords: []string{
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_ipblock",
			Category:         "Resources",
			ShortDescription: `Creates and manages IP Block objects.`,
			Description:      ``,
			Keywords: []string{
				"ipblock",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_ipfailover",
			Category:         "Resources",
			ShortDescription: `Creates and manages ipfailover objects.`,
			Description:      ``,
			Keywords: []string{
				"ipfailover",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_k8s_cluster",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Kubernetes Clusters.`,
			Description:      ``,
			Keywords: []string{
				"k8s",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_k8s_node_pool",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Kubernetes Node Pools.`,
			Description:      ``,
			Keywords: []string{
				"k8s",
				"node",
				"pool",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_lan",
			Category:         "Resources",
			ShortDescription: `Creates and manages LAN objects.`,
			Description:      ``,
			Keywords: []string{
				"lan",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Creates and manages Load Balancers`,
			Description:      ``,
			Keywords: []string{
				"loadbalancer",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_nic",
			Category:         "Resources",
			ShortDescription: `Creates and manages Network Interface objects.`,
			Description:      ``,
			Keywords: []string{
				"nic",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_server",
			Category:         "Resources",
			ShortDescription: `Creates and manages ProfitBricks Server objects.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_share",
			Category:         "Resources",
			ShortDescription: `Creates and manages share objects.`,
			Description:      ``,
			Keywords: []string{
				"share",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_snapshot",
			Category:         "Resources",
			ShortDescription: `Creates and manages snapshot objects.`,
			Description:      ``,
			Keywords: []string{
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages user objects.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_volume",
			Category:         "Resources",
			ShortDescription: `Creates and manages ProfitBricks Volume objects.`,
			Description:      ``,
			Keywords: []string{
				"volume",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"profitbricks_datacenter":    0,
		"profitbricks_firewall":      1,
		"profitbricks_group":         2,
		"profitbricks_ipblock":       3,
		"profitbricks_ipfailover":    4,
		"profitbricks_k8s_cluster":   5,
		"profitbricks_k8s_node_pool": 6,
		"profitbricks_lan":           7,
		"profitbricks_loadbalancer":  8,
		"profitbricks_nic":           9,
		"profitbricks_server":        10,
		"profitbricks_share":         11,
		"profitbricks_snapshot":      12,
		"profitbricks_user":          13,
		"profitbricks_volume":        14,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
