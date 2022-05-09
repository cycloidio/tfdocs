package upcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "upcloud_hosts",
			Category:         "Data Sources",
			ShortDescription: `Returns a list of available UpCloud hosts. A host identifies the host server that virtual machines are run on. Only hosts on private cloud to which the calling account has access to are available through this resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_ip_addresses",
			Category:         "Data Sources",
			ShortDescription: `Returns a set of IP Addresses that are associated with the UpCloud account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_networks",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get the available UpCloud networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_storage",
			Category:         "Data Sources",
			ShortDescription: `Returns storage resource information based on defined arguments. Data object can be used to map storage to other resource based on the ID or just to read some other storage property like zone information.Storage types are: normal, backup, cdrom, template`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_tags",
			Category:         "Data Sources",
			ShortDescription: `Data-source is deprecated.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_zone",
			Category:         "Data Sources",
			ShortDescription: `Data-source is deprecated.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "upcloud_zones",
			Category:         "Data Sources",
			ShortDescription: `Data-source is deprecated.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"upcloud_hosts":        0,
		"upcloud_ip_addresses": 1,
		"upcloud_networks":     2,
		"upcloud_storage":      3,
		"upcloud_tags":         4,
		"upcloud_zone":         5,
		"upcloud_zones":        6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
