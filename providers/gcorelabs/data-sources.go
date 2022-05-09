package gcorelabs

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_floatingip",
			Category:         "Data Sources",
			ShortDescription: `A floating IP is a static IP address that points to one of your Instances. It allows you to redirect network traffic to any of your Instances in the same datacenter.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_image",
			Category:         "Data Sources",
			ShortDescription: `Represent image data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_instance",
			Category:         "Data Sources",
			ShortDescription: `Represent instance. Could be used with baremetal also`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_k8s",
			Category:         "Data Sources",
			ShortDescription: `Represent k8s cluster with one default pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_k8s_pool",
			Category:         "Data Sources",
			ShortDescription: `Represent k8s cluster's pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_lblistener",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_lbpool",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_loadbalancer",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_loadbalancerv2",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_network",
			Category:         "Data Sources",
			ShortDescription: `Represent network. A network is a software-defined network in a cloud computing infrastructure`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_project",
			Category:         "Data Sources",
			ShortDescription: `Represent project data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_region",
			Category:         "Data Sources",
			ShortDescription: `Represent region data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_reservedfixedip",
			Category:         "Data Sources",
			ShortDescription: `Represent reserved ips`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_router",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_secret",
			Category:         "Data Sources",
			ShortDescription: `Represent secret`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_securitygroup",
			Category:         "Data Sources",
			ShortDescription: `Represent SecurityGroups(Firewall)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_servergroup",
			Category:         "Data Sources",
			ShortDescription: `Represent server group data`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_storage_s3",
			Category:         "Data Sources",
			ShortDescription: `Represent s3 storage resource. https://storage.gcorelabs.com/storage/list`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_storage_s3_bucket",
			Category:         "Data Sources",
			ShortDescription: `Represent storage s3 bucket resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_storage_sftp",
			Category:         "Data Sources",
			ShortDescription: `Represent sftp storage resource. https://storage.gcorelabs.com/storage/list`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_storage_sftp_key",
			Category:         "Data Sources",
			ShortDescription: `Represent storage key resource. https://storage.gcorelabs.com/ssh-key/list`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_subnet",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gcorelabs_gcore_volume",
			Category:         "Data Sources",
			ShortDescription: `Represent volume. A volume is a file storage which is similar to SSD and HDD hard disks`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"gcorelabs_gcore_floatingip":        0,
		"gcorelabs_gcore_image":             1,
		"gcorelabs_gcore_instance":          2,
		"gcorelabs_gcore_k8s":               3,
		"gcorelabs_gcore_k8s_pool":          4,
		"gcorelabs_gcore_lblistener":        5,
		"gcorelabs_gcore_lbpool":            6,
		"gcorelabs_gcore_loadbalancer":      7,
		"gcorelabs_gcore_loadbalancerv2":    8,
		"gcorelabs_gcore_network":           9,
		"gcorelabs_gcore_project":           10,
		"gcorelabs_gcore_region":            11,
		"gcorelabs_gcore_reservedfixedip":   12,
		"gcorelabs_gcore_router":            13,
		"gcorelabs_gcore_secret":            14,
		"gcorelabs_gcore_securitygroup":     15,
		"gcorelabs_gcore_servergroup":       16,
		"gcorelabs_gcore_storage_s3":        17,
		"gcorelabs_gcore_storage_s3_bucket": 18,
		"gcorelabs_gcore_storage_sftp":      19,
		"gcorelabs_gcore_storage_sftp_key":  20,
		"gcorelabs_gcore_subnet":            21,
		"gcorelabs_gcore_volume":            22,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
