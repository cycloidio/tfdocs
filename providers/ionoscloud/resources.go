package ionoscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_backup_unit",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud Backup Units.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_datacenter",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud Virtual Data Center.`,
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
					Description: `(Optional)[string] Description for the Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `(Optional) [bool] Boolean value representing if the data center requires extra protection e.g. two factor protection`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Computed) The version of that Data Center. Gets incremented with every change`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `(Computed) List of features supported by the location this data center is part of ## Import Resource Datacenter can be imported using the ` + "`" + `resource id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + `shell terraform import ionoscloud_datacenter.mydc {datacenter uuid} ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_firewall",
			Category:         "Resources",
			ShortDescription: `Creates and manages Firewall Rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_group",
			Category:         "Resources",
			ShortDescription: `Creates and manages group objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_ipblock",
			Category:         "Resources",
			ShortDescription: `Creates and manages IP Block objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_ipfailover",
			Category:         "Resources",
			ShortDescription: `Creates and manages ipfailover objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_k8s_cluster",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud Kubernetes Clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_k8s_node_pool",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud Kubernetes Node Pools.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_lan",
			Category:         "Resources",
			ShortDescription: `Creates and manages LAN objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_loadbalancer",
			Category:         "Resources",
			ShortDescription: `Creates and manages Load Balancers`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_nic",
			Category:         "Resources",
			ShortDescription: `Creates and manages Network Interface objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_private_crossconnect",
			Category:         "Resources",
			ShortDescription: `Creates and manages Private Cross Connections between virtual datacenters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_s3_key",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud S3 keys.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_server",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud Server objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_share",
			Category:         "Resources",
			ShortDescription: `Creates and manages share objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_snapshot",
			Category:         "Resources",
			ShortDescription: `Creates and manages snapshot objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages user objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_volume",
			Category:         "Resources",
			ShortDescription: `Creates and manages IonosCloud Volume objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ionoscloud_backup_unit":          0,
		"ionoscloud_datacenter":           1,
		"ionoscloud_firewall":             2,
		"ionoscloud_group":                3,
		"ionoscloud_ipblock":              4,
		"ionoscloud_ipfailover":           5,
		"ionoscloud_k8s_cluster":          6,
		"ionoscloud_k8s_node_pool":        7,
		"ionoscloud_lan":                  8,
		"ionoscloud_loadbalancer":         9,
		"ionoscloud_nic":                  10,
		"ionoscloud_private_crossconnect": 11,
		"ionoscloud_s3_key":               12,
		"ionoscloud_server":               13,
		"ionoscloud_share":                14,
		"ionoscloud_snapshot":             15,
		"ionoscloud_user":                 16,
		"ionoscloud_volume":               17,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
