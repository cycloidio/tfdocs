package icinga2

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "icinga2_checkcommand",
			Category:         "Resources",
			ShortDescription: `Configures a checkcommand resource. This allows checkcommands to be configured, updated and deleted.`,
			Description:      ``,
			Keywords: []string{
				"checkcommand",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "arguments",
					Description: `(Optional) A mapping of arguments to include with the command.`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required) Path to the command te be executed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name by which to reference the checkcommand`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `(Optional) A list of Icinga2 templates to assign to the host.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "icinga2_host",
			Category:         "Resources",
			ShortDescription: `Configures a host resource. This allows hosts to be configured, updated and deleted.`,
			Description:      ``,
			Keywords: []string{
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "check_command",
					Description: `(Required) The name of an existing Icinga2 CheckCommand object that is used to determine if the host is available or not.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) The groups of the host.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The hostname of the host.`,
				},
				resource.Attribute{
					Name:        "templates",
					Description: `(Optional) A list of Icinga2 templates to assign to the host.`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `(Optional) A mapping of variables to assign to the host.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "icinga2_hostgroup",
			Category:         "Resources",
			ShortDescription: `Configures a hostgroup resource. This allows hostgroup to be configured, updated and deleted.`,
			Description:      ``,
			Keywords: []string{
				"hostgroup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the hostgroup.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Required) The name of the hostgroup to display in the Icinga2 interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "icinga2_notification",
			Category:         "Resources",
			ShortDescription: `Configures a notifiation resource. This allows notifications to be configured, updated and deleted.`,
			Description:      ``,
			Keywords: []string{
				"notification",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servicename",
					Description: `(Optional) Service to send notification for.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "icinga2_service",
			Category:         "Resources",
			ShortDescription: `Configures a service resource. This allows service to be configured, updated and deleted.`,
			Description:      ``,
			Keywords: []string{
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Service object.`,
				},
				resource.Attribute{
					Name:        "check_command",
					Description: `(Required) The name of an existing Icinga2 CheckCommand object that is used to determine if the service is available on the host.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The host to check the service's status on`,
				},
				resource.Attribute{
					Name:        "vars",
					Description: `(Optional) A mapping of variables to assign to the service.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "icinga2_user",
			Category:         "Resources",
			ShortDescription: `Configures an user resource. This allows users to be configured, updated and deleted.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) An email string for this user. Useful for notification commands.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"icinga2_checkcommand": 0,
		"icinga2_host":         1,
		"icinga2_hostgroup":    2,
		"icinga2_notification": 3,
		"icinga2_service":      4,
		"icinga2_user":         5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
