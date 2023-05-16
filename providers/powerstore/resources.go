package powerstore

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "powerstore_host",
			Category:         "Resources",
			ShortDescription: `Host resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_hostgroup",
			Category:         "Resources",
			ShortDescription: `HostGroup resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_protectionpolicy",
			Category:         "Resources",
			ShortDescription: `ProtectionPolicy resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_snapshotrule",
			Category:         "Resources",
			ShortDescription: `SnapshotRule resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_storagecontainer",
			Category:         "Resources",
			ShortDescription: `StorageContainer resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_volume",
			Category:         "Resources",
			ShortDescription: `.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_volume_snapshot",
			Category:         "Resources",
			ShortDescription: `volume snapshot resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_volumegroup",
			Category:         "Resources",
			ShortDescription: `VolumeGroup resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "powerstore_volumegroup_snapshot",
			Category:         "Resources",
			ShortDescription: `Volume Group Snapshot resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"powerstore_host":                 0,
		"powerstore_hostgroup":            1,
		"powerstore_protectionpolicy":     2,
		"powerstore_snapshotrule":         3,
		"powerstore_storagecontainer":     4,
		"powerstore_volume":               5,
		"powerstore_volume_snapshot":      6,
		"powerstore_volumegroup":          7,
		"powerstore_volumegroup_snapshot": 8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
