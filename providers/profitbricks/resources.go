package profitbricks

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_backup_unit",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Backup Units.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_datacenter",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Virtual Data Center.`,
			Description:      ``,
			Keywords:         []string{},
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
			Type:             "profitbricks_profitbricks_firewall",
			Category:         "Resources",
			ShortDescription: `Creates and manages Firewall Rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_group",
			Category:         "Resources",
			ShortDescription: `Creates and manages group objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_ipblock",
			Category:         "Resources",
			ShortDescription: `Creates and manages IP Block objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_ipfailover",
			Category:         "Resources",
			ShortDescription: `Creates and manages ipfailover objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_k8s_cluster",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Kubernetes Clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_k8s_node_pool",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks Kubernetes Node Pools.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_lan",
			Category:         "Resources",
			ShortDescription: `Creates and manages LAN objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Creates and manages Load Balancers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_nic",
			Category:         "Resources",
			ShortDescription: `Creates and manages Network Interface objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_private_crossconnect",
			Category:         "Resources",
			ShortDescription: `Creates and manages Private Cross Connections between virtual datacenters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_s3_key",
			Category:         "Resources",
			ShortDescription: `Creates and manages Profitbricks S3 keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_server",
			Category:         "Resources",
			ShortDescription: `Creates and manages ProfitBricks Server objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_share",
			Category:         "Resources",
			ShortDescription: `Creates and manages share objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_snapshot",
			Category:         "Resources",
			ShortDescription: `Creates and manages snapshot objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages user objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "profitbricks_profitbricks_volume",
			Category:         "Resources",
			ShortDescription: `Creates and manages ProfitBricks Volume objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"profitbricks_profitbricks_backup_unit":          0,
		"profitbricks_profitbricks_datacenter":           1,
		"profitbricks_profitbricks_firewall":             2,
		"profitbricks_profitbricks_group":                3,
		"profitbricks_profitbricks_ipblock":              4,
		"profitbricks_profitbricks_ipfailover":           5,
		"profitbricks_profitbricks_k8s_cluster":          6,
		"profitbricks_profitbricks_k8s_node_pool":        7,
		"profitbricks_profitbricks_lan":                  8,
		"profitbricks_profitbricks_loadbalancer":         9,
		"profitbricks_profitbricks_nic":                  10,
		"profitbricks_profitbricks_private_crossconnect": 11,
		"profitbricks_profitbricks_s3_key":               12,
		"profitbricks_profitbricks_server":               13,
		"profitbricks_profitbricks_share":                14,
		"profitbricks_profitbricks_snapshot":             15,
		"profitbricks_profitbricks_user":                 16,
		"profitbricks_profitbricks_volume":               17,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
