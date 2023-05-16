package brightbox

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "brightbox_database_snapshot",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image based upon the ` + "`" + `created_at` + "`" + ` time.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A regex string to apply to the Database Snapshot list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A regex string to apply to the Database Snapshot list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of database partition in megabytes`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state the image is in. Usually ` + "`" + `available` + "`" + `, or ` + "`" + `deleted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time and date the image was created/registered (UTC)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `true if image has been set as locked and can not be deleted`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "size",
					Description: `The size of database partition in megabytes`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state the image is in. Usually ` + "`" + `available` + "`" + `, or ` + "`" + `deleted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time and date the image was created/registered (UTC)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `true if image has been set as locked and can not be deleted`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_database_type",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A regex string to apply to the Database Type list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A regex string to apply to the Database Type list returned by Brightbox Cloud. ~>`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the database server for this type`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The memory size of the database server for this type`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Is this the default database type?`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the database server for this type`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The memory size of the database server for this type`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Is this the default database type?`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_image",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image based upon the ` + "`" + `created_at` + "`" + ` time.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A regex string to apply to the Image list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A regex string to apply to the Image list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `(Optional) Either ` + "`" + `upload` + "`" + ` or ` + "`" + `snapshot` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Name of the source for this image. Matches exactly.`,
				},
				resource.Attribute{
					Name:        "source_trigger",
					Description: `(Optional) Either ` + "`" + `manual` + "`" + ` or ` + "`" + `schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The account id that owns the image. Matches exactly.`,
				},
				resource.Attribute{
					Name:        "arch",
					Description: `(Optional) The architecture of the image: either ` + "`" + `x86_64` + "`" + ` or ` + "`" + `i686` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `(Optional) Boolean to select a public image.`,
				},
				resource.Attribute{
					Name:        "official",
					Description: `(Optional) Boolean to select an official image.`,
				},
				resource.Attribute{
					Name:        "compatibility_mode",
					Description: `(Optional) Boolean to match the compatibility mode flag.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) The username used to logon to the image. Matches exactly.`,
				},
				resource.Attribute{
					Name:        "ancestor_id",
					Description: `(Optional) The image id of the parent of the image you are looking for.`,
				},
				resource.Attribute{
					Name:        "licence_name",
					Description: `(Optional) The name of the licence for the image. Matches exactly.`,
				},
				resource.Attribute{
					Name:        "min_ram",
					Description: `(Optional) The actual size of the data within the image in megabytes. Matches exactly. ~>`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The state the image is in. Usually ` + "`" + `available` + "`" + `, ` + "`" + `deprecated` + "`" + ` or ` + "`" + `deleted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time and date the image was created/registered (UTC)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `true if image has been set as locked and can not be deleted`,
				},
				resource.Attribute{
					Name:        "virtual_size",
					Description: `The virtual size of the disk image "container" in MB`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The actual size of the data within the Image in MB`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The state the image is in. Usually ` + "`" + `available` + "`" + `, ` + "`" + `deprecated` + "`" + ` or ` + "`" + `deleted` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The time and date the image was created/registered (UTC)`,
				},
				resource.Attribute{
					Name:        "locked",
					Description: `true if image has been set as locked and can not be deleted`,
				},
				resource.Attribute{
					Name:        "virtual_size",
					Description: `The virtual size of the disk image "container" in MB`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The actual size of the data within the Image in MB`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_server_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A regex string to apply to the Server Group list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A regex string to apply to the Server Group list returned by Brightbox Cloud. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Is this the default server group?`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name of the server group`,
				},
				resource.Attribute{
					Name:        "firewall_policy",
					Description: `The ID of the Firewall Policy associated with this group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Is this the default server group?`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The Fully Qualified Domain Name of the server group`,
				},
				resource.Attribute{
					Name:        "firewall_policy",
					Description: `The ID of the Firewall Policy associated with this group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "brightbox_server_type",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A regex string to apply to the Server Type list returned by Brightbox Cloud.`,
				},
				resource.Attribute{
					Name:        "handle",
					Description: `(Optional) A regex string to apply to the Server Type list returned by Brightbox Cloud. ~>`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the server for this type. Will be zero for network storage types.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The memory size of the server for this type.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The memory size of the server for this type.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `'experimental', 'available' or 'deprecated'.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `the type of block storage available with this type: 'network' or 'local'.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "disk_size",
					Description: `The disk size of the server for this type. Will be zero for network storage types.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The memory size of the server for this type.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The memory size of the server for this type.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `'experimental', 'available' or 'deprecated'.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `the type of block storage available with this type: 'network' or 'local'.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"brightbox_database_snapshot": 0,
		"brightbox_database_type":     1,
		"brightbox_image":             2,
		"brightbox_server_group":      3,
		"brightbox_server_type":       4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
