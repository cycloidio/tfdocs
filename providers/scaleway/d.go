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
			Type:             "scaleway_bootscript",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Scaleway bootscript.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `(Optional) any supported Scaleway architecture, e.g. ` + "`" + `x86_64` + "`" + `, ` + "`" + `arm` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name_filter",
					Description: `(Optional) Regexp to match Bootscript name by`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Exact name of desired Bootscript ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found Bootscript. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `architecture of the Bootscript, e.g. ` + "`" + `arm` + "`" + ` or ` + "`" + `x86_64` + "`" + ``,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `uuid of the organization owning this Bootscript`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `is this a public bootscript`,
				},
				resource.Attribute{
					Name:        "boot_cmd_args",
					Description: `command line arguments used for booting`,
				},
				resource.Attribute{
					Name:        "dtb",
					Description: `path to Device Tree Blob detailing hardware information`,
				},
				resource.Attribute{
					Name:        "initrd",
					Description: `URL to initial ramdisk content`,
				},
				resource.Attribute{
					Name:        "kernel",
					Description: `URL to used kernel`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `architecture of the Bootscript, e.g. ` + "`" + `arm` + "`" + ` or ` + "`" + `x86_64` + "`" + ``,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `uuid of the organization owning this Bootscript`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `is this a public bootscript`,
				},
				resource.Attribute{
					Name:        "boot_cmd_args",
					Description: `command line arguments used for booting`,
				},
				resource.Attribute{
					Name:        "dtb",
					Description: `path to Device Tree Blob detailing hardware information`,
				},
				resource.Attribute{
					Name:        "initrd",
					Description: `URL to initial ramdisk content`,
				},
				resource.Attribute{
					Name:        "kernel",
					Description: `URL to used kernel`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_image",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Scaleway image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `(Required) any supported Scaleway architecture, e.g. ` + "`" + `x86_64` + "`" + `, ` + "`" + `arm` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name_filter",
					Description: `(Optional) Regexp to match Image name by`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Exact name of desired Image`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) Return most recent image if multiple exist. Can not be used together with name_filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found Image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "architecture",
					Description: `architecture of the Image, e.g. ` + "`" + `arm` + "`" + ` or ` + "`" + `x86_64` + "`" + ``,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `uuid of the organization owning this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `is this a public image`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `date when image was created`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "architecture",
					Description: `architecture of the Image, e.g. ` + "`" + `arm` + "`" + ` or ` + "`" + `x86_64` + "`" + ``,
				},
				resource.Attribute{
					Name:        "organization",
					Description: `uuid of the organization owning this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `is this a public image`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `date when image was created`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_security_group",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Exact name of desired Security Group ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found Image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the security group`,
				},
				resource.Attribute{
					Name:        "enable_default_security",
					Description: `have default security group rules been added to this security group?`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `description of the security group`,
				},
				resource.Attribute{
					Name:        "enable_default_security",
					Description: `have default security group rules been added to this security group?`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "scaleway_volume",
			Category:         "Data Sources",
			ShortDescription: `Gets information about a Volume.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Exact name of the Volume. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found Volume. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of volume this is, such as ` + "`" + `l_ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The ID of the Server which this Volume is currently attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size_in_gb",
					Description: `(Required) size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of volume this is, such as ` + "`" + `l_ssd` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The ID of the Server which this Volume is currently attached to.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"scaleway_account_ssh_key": 0,
		"scaleway_bootscript":      1,
		"scaleway_image":           2,
		"scaleway_security_group":  3,
		"scaleway_volume":          4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
