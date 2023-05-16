package seeweb

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "seeweb_action",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) A unique numeric ID that can be used to identify and reference an action. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the action. This can be "in-progress", "completed", or "error".`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `A unique identifier for the resource that the action is associated with`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The current action plan.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A unique identifier for the account that the action is associated with.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was created.`,
				},
				resource.Attribute{
					Name:        "started_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was initiated.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was completed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This is the type of action that the object represents.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of resource that the action is associated with.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `A value that represent the percentage of completation. [1]: https://docs.seeweb.it/ecs/api/#list-all-actions`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the action. This can be "in-progress", "completed", or "error".`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `A unique identifier for the resource that the action is associated with`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The current action plan.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A unique identifier for the account that the action is associated with.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was created.`,
				},
				resource.Attribute{
					Name:        "started_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was initiated.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was completed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This is the type of action that the object represents.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of resource that the action is associated with.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `A value that represent the percentage of completation. [1]: https://docs.seeweb.it/ecs/api/#list-all-actions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_actions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "actions",
					Description: `List of actions with their data. ### actions (` + "`" + `actions` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique number ID.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the action. This can be "in-progress", "completed", or "error".`,
				},
				resource.Attribute{
					Name:        "resource",
					Description: `A unique identifier for the resource that the action is associated with`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `The current action plan.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `A unique identifier for the account that the action is associated with.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was created.`,
				},
				resource.Attribute{
					Name:        "started_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was initiated.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `A time value given in ISO8601 combined date and time format that represents when the action was completed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This is the type of action that the object represents.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `The type of resource that the action is associated with.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `A value that represent the percentage of completation. [1]: https://docs.seeweb.it/ecs/api/#list-all-actions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_group",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) A unique number ID. ## Attributes Reference`,
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
					Description: `The current status of a group. [1]: https://docs.seeweb.it/ecs/api/#list-all-groups`,
				},
			},
			Attributes: []resource.Attribute{
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
					Description: `The current status of a group. [1]: https://docs.seeweb.it/ecs/api/#list-all-groups`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_groups",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "groups",
					Description: `List of groups with their data. ### groups (` + "`" + `groups` + "`" + `) supports the following:`,
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
					Description: `The current status of a group. [1]: https://docs.seeweb.it/ecs/api/#list-all-groups`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_plans",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plans",
					Description: `List of plans with their data. ### Plans (` + "`" + `plans` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique number ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Plan name.`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Number of CPU.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Memory sixe in MB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size in MB.`,
				},
				resource.Attribute{
					Name:        "hourly_price",
					Description: `Hourly price in Euro.`,
				},
				resource.Attribute{
					Name:        "montly_price",
					Description: `Montly price in Euro.`,
				},
				resource.Attribute{
					Name:        "windows",
					Description: `So windows family, true or false.`,
				},
				resource.Attribute{
					Name:        "available",
					Description: `Plan available, true or false.`,
				},
				resource.Attribute{
					Name:        "available_regions",
					Description: `Available regions. The ` + "`" + `available_regions` + "`" + ` block contains the following attributes:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique number ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Region location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Region description. [1]: https://docs.seeweb.it/ecs/api/#list-all-plans`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_regions",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `List of regions with their data. ### Regions (` + "`" + `regions` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique number ID.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Region location.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Region description. [1]: https://docs.seeweb.it/ecs/api/#list-all-regions`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_server",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The server name to use to find a server in the Seeweb API. ## Attributes Reference`,
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
					Name:        "plan",
					Description: `The current server plan.`,
				},
				resource.Attribute{
					Name:        "plan_size",
					Description: `The server plan configuration sizes.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `An unique identifier for the server region.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable name for this server.`,
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
					Description: `Disk size in MB. [1]: https://docs.seeweb.it/ecs/api/#list-all-servers`,
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
					Name:        "plan",
					Description: `The current server plan.`,
				},
				resource.Attribute{
					Name:        "plan_size",
					Description: `The server plan configuration sizes.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `An unique identifier for the server region.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable name for this server.`,
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
					Description: `Disk size in MB. [1]: https://docs.seeweb.it/ecs/api/#list-all-servers`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_servers",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "servers",
					Description: `List of servers with their data. ### servers (` + "`" + `servers` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The server name to use to find a server in the Seeweb API.`,
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
					Name:        "plan",
					Description: `The current server plan.`,
				},
				resource.Attribute{
					Name:        "plan_size",
					Description: `The server plan configuration sizes.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `An unique identifier for the server region.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable name for this server.`,
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
					Description: `Disk size in MB. [1]: https://docs.seeweb.it/ecs/api/#list-all-servers`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_template",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) A unique numeric ID that can be used to identify and reference a template. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A unique template name.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the template was created.`,
				},
				resource.Attribute{
					Name:        "active_flag",
					Description: `A flag that represents if a template is manageble.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the template. This can be "Creating", "Created", "Deleting" or "Deleted".`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `A unique identifier.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable friendly description. [1]: https://docs.seeweb.it/ecs/api/#list-all-templates`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `A unique template name.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the template was created.`,
				},
				resource.Attribute{
					Name:        "active_flag",
					Description: `A flag that represents if a template is manageble.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the template. This can be "Creating", "Created", "Deleting" or "Deleted".`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `A unique identifier.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable friendly description. [1]: https://docs.seeweb.it/ecs/api/#list-all-templates`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "seeweb_templates",
			Category:         "Data Sources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "templates",
					Description: `List of templates with their data. ### templates (` + "`" + `templates` + "`" + `) supports the following:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `A unique numeric ID that can be used to identify and reference a template.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `A time value given in ISO8601 combined date and time format that represents when the template was created.`,
				},
				resource.Attribute{
					Name:        "active_flag",
					Description: `A flag that represents if a template is manageble.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The current status of the template. This can be "Creating", "Created", "Deleting" or "Deleted".`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `A unique identifier.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `A human-readable friendly description. [1]: https://docs.seeweb.it/ecs/api/#list-all-templates`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"seeweb_action":    0,
		"seeweb_actions":   1,
		"seeweb_group":     2,
		"seeweb_groups":    3,
		"seeweb_plans":     4,
		"seeweb_regions":   5,
		"seeweb_server":    6,
		"seeweb_servers":   7,
		"seeweb_template":  8,
		"seeweb_templates": 9,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
