package hedvig

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "access resource",
			Category:         "Resources",
			ShortDescription: `Adds an ACL address to a controller resource of a vdisk resource.`,
			Description:      ``,
			Keywords: []string{
				"access resource",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdisk",
					Description: `(Required) The name of the Vdisk that this Access is associated with.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Required) The fully qualified domain name of the controller this Access is associated with.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `(Required) The actual address that this Access is providing access to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of address provided in ` + "`" + `address` + "`" + `. Can be ` + "`" + `host` + "`" + `, ` + "`" + `ip` + "`" + ` or ` + "`" + `iqn` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lun resource",
			Category:         "Resources",
			ShortDescription: `Adds a virtual disk as a LUN to a controller.`,
			Description:      ``,
			Keywords: []string{
				"lun resource",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdisk",
					Description: `(Required) The name of the vdisk the LUN is on.`,
				},
				resource.Attribute{
					Name:        "controller",
					Description: `(Required) The fully qualified domain name for the controller that the LUN is to attach to.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "mount resource",
			Category:         "Resources",
			ShortDescription: `Mounts vdisk resource with a particular controller.`,
			Description:      ``,
			Keywords: []string{
				"mount resource",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdisk",
					Description: `(Required) The name of the vdisk the Mount is on.`,
				},
				resource.Attribute{
					Name:        "controller",
					Description: `(Required) The fully qualified domain name for the controller that the Mount is to attach to.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vdisk resource",
			Category:         "Resources",
			ShortDescription: `Manages a Vdisk resource on a Hedvig cluster.`,
			Description:      ``,
			Keywords: []string{
				"vdisk resource",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be used by the Vdisk for identification.`,
				},
				resource.Attribute{
					Name:        "residence",
					Description: `(Optional) Disk residence; can be either ` + "`" + `HDD` + "`" + ` or ` + "`" + `Flash` + "`" + ``,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the disk in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the disk; can be either ` + "`" + `BLOCK` + "`" + ` or ` + "`" + `NFS` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scsi3pr",
					Description: `(Required, defaults to false) Enables SCSI-3 Persistent Reservations for use wwith Clustered Shared Volumes (CSV)`,
				},
				resource.Attribute{
					Name:        "blocksize",
					Description: `(Optional, defaults to 4096)`,
				},
				resource.Attribute{
					Name:        "cacheenabled",
					Description: `(Optional, defaults to false) Enables client-side caching support for virtual disk blocks, to cache to local SSD or PCIe devices at the application compute tier`,
				},
				resource.Attribute{
					Name:        "clusteredfilesystem",
					Description: `(Optional, defaults to false) Formats a clustered file system on top of a virtual disk to be presented to multiple hosts`,
				},
				resource.Attribute{
					Name:        "compressed",
					Description: `(Optional, defaults to false)`,
				},
				resource.Attribute{
					Name:        "deduplication",
					Description: `(Optional, defaults to false)`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "encryption",
					Description: `(Optional, defaults to false)`,
				},
				resource.Attribute{
					Name:        "replicationfactor",
					Description: `(Optional, defaults to 3) Can be any integer 1 - 6`,
				},
				resource.Attribute{
					Name:        "replicationpolicy",
					Description: `(Optional, defaults to Agnostic) Can be RackAware, DataCenterAware, or Agnostic (RackUnaware)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"access resource": 0,
		"lun resource":    1,
		"mount resource":  2,
		"vdisk resource":  3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
