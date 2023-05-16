package anxcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_core_location",
			Category:         "Data Sources",
			ShortDescription: `Retrieves a location identified by it's identifier or human-readable code as selectable in the Engine. This data source can be used to lookup a locations identifier required by other resources and data sources available in this provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_core_locations",
			Category:         "Data Sources",
			ShortDescription: `Provides available locations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_cpu_performance_types",
			Category:         "Data Sources",
			ShortDescription: `Provides available cpu performance types. This information can be used to provision virtual servers using the anxcloud_virtual_server resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_disk_types",
			Category:         "Data Sources",
			ShortDescription: `Provides available disk types for a specified location. This information can be used to provision virtual servers using the anxcloud_virtual_server resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_dns_records",
			Category:         "Data Sources",
			ShortDescription: `Provides DNS records for a specified zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_dns_zones",
			Category:         "Data Sources",
			ShortDescription: `Provides available DNS zones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_ip_address",
			Category:         "Data Sources",
			ShortDescription: `Retrieves an IP address. Known limitations When using the address argument, only IP addresses unique to the scope of your access token for Anexia Cloud can be retrieved. You can however get a unique result by specifying the related VLAN or network prefix.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_ip_addresses",
			Category:         "Data Sources",
			ShortDescription: `Provides available IP addresses.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_kubernetes_cluster",
			Category:         "Data Sources",
			ShortDescription: `Retrieves a Kubernetes cluster resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_nic_types",
			Category:         "Data Sources",
			ShortDescription: `Provides available network interface card types. This information can be used to provision virtual servers using the anxcloud_virtual_server resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_tags",
			Category:         "Data Sources",
			ShortDescription: `Provides available service tags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_template",
			Category:         "Data Sources",
			ShortDescription: `Provides available templates for specified location. This information can be used to provision virtual servers using the anxcloud_virtual_server resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "anxcloud_vlans",
			Category:         "Data Sources",
			ShortDescription: `Provides available VLANs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"anxcloud_core_location":         0,
		"anxcloud_core_locations":        1,
		"anxcloud_cpu_performance_types": 2,
		"anxcloud_disk_types":            3,
		"anxcloud_dns_records":           4,
		"anxcloud_dns_zones":             5,
		"anxcloud_ip_address":            6,
		"anxcloud_ip_addresses":          7,
		"anxcloud_kubernetes_cluster":    8,
		"anxcloud_nic_types":             9,
		"anxcloud_tags":                  10,
		"anxcloud_template":              11,
		"anxcloud_vlans":                 12,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
