package powerstore

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerstore_host",
			Category:         "Data Sources",
			ShortDescription: `Host DataSource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_hostgroup",
			Category:         "Data Sources",
			ShortDescription: `HostGroup DataSource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_protectionpolicy",
			Category:         "Data Sources",
			ShortDescription: `ProtectionPolicy DataSource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_snapshotrule",
			Category:         "Data Sources",
			ShortDescription: `SnapshotRule DataSource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_volume",
			Category:         "Data Sources",
			ShortDescription: `.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_volumegroup",
			Category:         "Data Sources",
			ShortDescription: `VolumeGroup DataSource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"powerstore_host":             0,
		"powerstore_hostgroup":        1,
		"powerstore_protectionpolicy": 2,
		"powerstore_snapshotrule":     3,
		"powerstore_volume":           4,
		"powerstore_volumegroup":      5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
