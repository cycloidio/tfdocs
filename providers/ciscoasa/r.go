package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_access_in_rules",
			Category:         "Resources",
			ShortDescription: `Provides Cisco ASA Inbound Access Rules.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"in",
				"rules",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "interface",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) One or more ` + "`" + `rule` + "`" + ` elements as defined below.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) Default ` + "`" + `false` + "`" + `. ### ` + "`" + `rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destination_service",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permit",
					Description: `(Optional) Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_service",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed)`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_access_out_rules",
			Category:         "Resources",
			ShortDescription: `Provides a Cisco ASA Outbound Access Rule.`,
			Description:      ``,
			Keywords: []string{
				"access",
				"out",
				"rules",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "interface",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) One or more ` + "`" + `rule` + "`" + ` elements as defined below.`,
				},
				resource.Attribute{
					Name:        "managed",
					Description: `(Optional) Default ` + "`" + `false` + "`" + `. ### ` + "`" + `rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destination_service",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permit",
					Description: `(Optional) Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_service",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed)`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_acl",
			Category:         "Resources",
			ShortDescription: `Provides a Cisco ASA ACL resource.`,
			Description:      ``,
			Keywords: []string{
				"acl",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL.`,
				},
				resource.Attribute{
					Name:        "rule",
					Description: `(Required) One or more ` + "`" + `rule` + "`" + ` elements as defined below. ### ` + "`" + `rule` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "destination_service",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `(Optional) Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_interval",
					Description: `(Optional) Default ` + "`" + `300` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_status",
					Description: `(Optional) Must be one of ` + "`" + `Default` + "`" + `, ` + "`" + `Debugging` + "`" + `, ` + "`" + `Disabled` + "`" + `, ` + "`" + `Notifications` + "`" + `, ` + "`" + `Critical` + "`" + `, ` + "`" + `Emergencies` + "`" + `, ` + "`" + `Warnings` + "`" + `, ` + "`" + `Errors` + "`" + `, ` + "`" + `Informational` + "`" + `, ` + "`" + `Alerts` + "`" + `. Default ` + "`" + `Default` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "permit",
					Description: `(Optional) Default ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remarks",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "source_service",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Computed)`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_network_object",
			Category:         "Resources",
			ShortDescription: `Provides a Cisco ASA Network Object resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"object",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) The value representing the object. This can be a single host, a range of hosts (` + "`" + `<ip>-<ip>` + "`" + `), or a CIDR.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_network_object_group",
			Category:         "Resources",
			ShortDescription: `Provides a Cisco ASA Network Object Group resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"object",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) The list of the group members.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_network_service_group",
			Category:         "Resources",
			ShortDescription: `Provides a Cisco ASA Network Service Group resource.`,
			Description:      ``,
			Keywords: []string{
				"network",
				"service",
				"group",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the group.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Required) The list of the group members.`,
				},
			},
			Attributes: []resource.Argument{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ciscoasa_static_route",
			Category:         "Resources",
			ShortDescription: `Provides a Cisco ASA Static Route resource.`,
			Description:      ``,
			Keywords: []string{
				"static",
				"route",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The name of the interface.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `(Required)`,
				},
				resource.Attribute{
					Name:        "metric",
					Description: `(Optional) Default ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tracked",
					Description: `(Optional) Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tunneled",
					Description: `(Optional) Default ` + "`" + `false` + "`" + `.`,
				},
			},
			Attributes: []resource.Argument{},
		},
	}

	resourcesMap = map[string]int{

		"ciscoasa_access_in_rules":       0,
		"ciscoasa_access_out_rules":      1,
		"ciscoasa_acl":                   2,
		"ciscoasa_network_object":        3,
		"ciscoasa_network_object_group":  4,
		"ciscoasa_network_service_group": 5,
		"ciscoasa_static_route":          6,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
