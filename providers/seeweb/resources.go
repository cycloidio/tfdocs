package seeweb

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "seeweb_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "notes",
					Description: `(Required) A human-readable friendly description. (`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) A private password used to make a subview of the infrastructure. (`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique number ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A unique group name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable friendly description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `The current status of a group. ## Import groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import seeweb_group.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ` [1]: https://docs.seeweb.it/ecs/api/#create-a-new-group`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique number ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A unique group name.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable friendly description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `The current status of a group. ## Import groups can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import seeweb_group.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ` [1]: https://docs.seeweb.it/ecs/api/#create-a-new-group`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "plan",
					Description: `(Required) the server plan. (`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) location identifier. (`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The image name of a public or private template. (`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Required) The human-readable string you wish to use to display server name.`,
				},
				resource.Attribute{
					Name:        "ssh_key",
					Description: `(Optional) Public key label. (`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `(Optional) The group name where the server belongs. If it set, it can not be blank and to dissociate a group from a server the keyword`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `The server public IPv4.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The server public IPv6.`,
				},
				resource.Attribute{
					Name:        "plan_size",
					Description: `The server plan configuration sizes.`,
				},
				resource.Attribute{
					Name:        "so",
					Description: `The template image used to deploy this server.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the server was created.`,
				},
				resource.Attribute{
					Name:        "deletion_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the server was deleted.`,
				},
				resource.Attribute{
					Name:        "active_flag",
					Description: `A flag value that shows if the server is active.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The server status: Booted, Booting, Deleting, Deleted.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The server API version.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `The server account username.`,
				},
				resource.Attribute{
					Name:        "virttype",
					Description: `The virtualization engine name. ### Plan Size (` + "`" + `plan_size` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "core",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Memory size in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size in MB. ## Import Servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import seeweb_server.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ` [1]: https://docs.seeweb.it/ecs/api/#create-new-server`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ipv4",
					Description: `The server public IPv4.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `The server public IPv6.`,
				},
				resource.Attribute{
					Name:        "plan_size",
					Description: `The server plan configuration sizes.`,
				},
				resource.Attribute{
					Name:        "so",
					Description: `The template image used to deploy this server.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the server was created.`,
				},
				resource.Attribute{
					Name:        "deletion_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the server was deleted.`,
				},
				resource.Attribute{
					Name:        "active_flag",
					Description: `A flag value that shows if the server is active.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The server status: Booted, Booting, Deleting, Deleted.`,
				},
				resource.Attribute{
					Name:        "api_version",
					Description: `The server API version.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `The server account username.`,
				},
				resource.Attribute{
					Name:        "virttype",
					Description: `The virtualization engine name. ### Plan Size (` + "`" + `plan_size` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "core",
					Description: `Number of CPU cores.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Memory size in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size in MB. ## Import Servers can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import seeweb_server.main PLBP09X ` + "`" + `` + "`" + `` + "`" + ` [1]: https://docs.seeweb.it/ecs/api/#create-new-server`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"seeweb_group":  0,
		"seeweb_server": 1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
