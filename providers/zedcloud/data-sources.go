package zedcloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_datastore",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloud_DataStore. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_edgeapp",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloud_edgeapp. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_edgeapp_instance",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloudedgeappinstance. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_edgenode",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloud_edgenode. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_image",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloud_image. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_network",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloud_network. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_network_instance",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloudnetworkinstance. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "zedcloud_volume_instance",
			Category:         "Data Sources",
			ShortDescription: `Schema for data source zedcloudvolumeinstance. Must specify id or name`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"zedcloud_datastore":        0,
		"zedcloud_edgeapp":          1,
		"zedcloud_edgeapp_instance": 2,
		"zedcloud_edgenode":         3,
		"zedcloud_image":            4,
		"zedcloud_network":          5,
		"zedcloud_network_instance": 6,
		"zedcloud_volume_instance":  7,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
