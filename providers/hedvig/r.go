package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "access resource",
			Category:         "Resources",
			ShortDescription: `Adds an ACL address to a controller resource of a vdisk resource.`,
			Description:      ``,
			Keywords: []string{
				"access resource",
			},
			Arguments: []resource.Argument{
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
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vdisk",
					Description: `(Required) The name of the vdisk the LUN is on.`,
				},
				resource.Attribute{
					Name:        "controller",
					Description: `(Required) The fully qualified domain name for the controller that the LUN is to attach to.`,
				},
			},
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "vdisk",
					Description: `(Required) The name of the vdisk the Mount is on.`,
				},
				resource.Attribute{
					Name:        "controller",
					Description: `(Required) The fully qualified domain name for the controller that the Mount is to attach to.`,
				},
			},
			Attributes: []resource.Argument{},
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
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to be used by the Vdisk for identification.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the disk in GB.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of the disk; can be either ` + "`" + `BLOCK` + "`" + ` or ` + "`" + `NFS` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"access resource": 0,
		"lun resource":    1,
		"mount resource":  2,
		"vdisk resource":  3,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
