package scaleway

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "scaleway_account_ssh_key",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Scaleway SSH key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_baremetal_offer",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a baremetal offer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an Instance Image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_security_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_server",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an Instance Server.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_instance_volume",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an instance volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_k8s_cluster",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Kubernetes Cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_k8s_pool",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Kubernetes Cluster's Pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_lb_ip",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Load Balancer IP.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_marketplace_image",
			Category:         "Data Sources",
			ShortDescription: `Gets local image ID of an image from its label name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_rdb_acl",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the RDB instance network Access Control List.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_rdb_database",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an RDB database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_rdb_instance",
			Category:         "Data Sources",
			ShortDescription: `Gets information about an RDB instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_rdb_privilege",
			Category:         "Data Sources",
			ShortDescription: `Gets information about the privilege on a RDB database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_registry_image",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a registry image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_registry_namespace",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a registry namespace.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_vpc_private_network",
			Category:         "Data Sources",
			ShortDescription: `Get information about Scaleway VPC Private Networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Exact name of the private network. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found private network. Addition attributes are exported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"scaleway_account_ssh_key":         0,
		"scaleway_baremetal_offer":         1,
		"scaleway_instance_image":          2,
		"scaleway_instance_security_group": 3,
		"scaleway_instance_server":         4,
		"scaleway_instance_volume":         5,
		"scaleway_k8s_cluster":             6,
		"scaleway_k8s_pool":                7,
		"scaleway_lb_ip":                   8,
		"scaleway_marketplace_image":       9,
		"scaleway_rdb_acl":                 10,
		"scaleway_rdb_database":            11,
		"scaleway_rdb_instance":            12,
		"scaleway_rdb_privilege":           13,
		"scaleway_registry_image":          14,
		"scaleway_registry_namespace":      15,
		"scaleway_vpc_private_network":     16,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
