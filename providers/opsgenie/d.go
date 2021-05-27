package opsgenie

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_escalation",
			Category:         "Data Sources",
			ShortDescription: `You can use this data source to map existing Escalation within Opsgenie.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the escalation. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Escalation.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `Escalation rules`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Escalation Description`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `If owner team exist the id of the team is exported`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Escalation.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `Escalation rules`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Escalation Description`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `If owner team exist the id of the team is exported`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_heartbeat",
			Category:         "Data Sources",
			ShortDescription: `Manages existing Heartbeats within Opsgenie.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the heartbeat ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of the heartbeat`,
				},
				resource.Attribute{
					Name:        "interval_unit",
					Description: `Interval specified as minutes, hours or days.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specifies how often a heartbeat message should be expected.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/disable heartbeat monitoring.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `Owner team of the heartbeat.`,
				},
				resource.Attribute{
					Name:        "alert_message",
					Description: `Specifies the alert message for heartbeat expiration alert. If this is not provided, default alert message is "HeartbeatName is expired".`,
				},
				resource.Attribute{
					Name:        "alert_priority",
					Description: `Specifies the alert priority for heartbeat expiration alert. If this is not provided, default priority is P3.`,
				},
				resource.Attribute{
					Name:        "alert_tags",
					Description: `Specifies the alert tags for heartbeat expiration alert.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of the heartbeat`,
				},
				resource.Attribute{
					Name:        "interval_unit",
					Description: `Interval specified as minutes, hours or days.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `Specifies how often a heartbeat message should be expected.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/disable heartbeat monitoring.`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `Owner team of the heartbeat.`,
				},
				resource.Attribute{
					Name:        "alert_message",
					Description: `Specifies the alert message for heartbeat expiration alert. If this is not provided, default alert message is "HeartbeatName is expired".`,
				},
				resource.Attribute{
					Name:        "alert_priority",
					Description: `Specifies the alert priority for heartbeat expiration alert. If this is not provided, default priority is P3.`,
				},
				resource.Attribute{
					Name:        "alert_tags",
					Description: `Specifies the alert tags for heartbeat expiration alert.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_schedule",
			Category:         "Data Sources",
			ShortDescription: `Manages a Schedule within Opsgenie.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the schedule. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A Member block as documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Timezone of schedule. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones - Default: ` + "`" + `America/New_York` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The description of schedule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/disable state of schedule`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `Owner team id of the schedule.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Schedule.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `A Member block as documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Timezone of schedule. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones - Default: ` + "`" + `America/New_York` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The description of schedule.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enable/disable state of schedule`,
				},
				resource.Attribute{
					Name:        "owner_team_id",
					Description: `Owner team id of the schedule.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_service",
			Category:         "Data Sources",
			ShortDescription: `Manages existing Service within Opsgenie.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the service. This field must not be longer than 100 characters. The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Service.`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `Team id of the service.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description field of the service that is generally used to provide a detailed information about the service.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_team",
			Category:         "Data Sources",
			ShortDescription: `Manages existing Team within Opsgenie.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name associated with this team. Opsgenie defines that this must not be longer than 100 characters. The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie Team.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `A Member block as documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description for this team.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_user",
			Category:         "Data Sources",
			ShortDescription: `Manages existing User within Opsgenie.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The email address associated with this user. Opsgenie defines that this must not be longer than 100 characters. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie User.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The Full Name of the User.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The Role assigned to the User. Either a built-in such as 'Owner', 'Admin' or 'User' - or the name of a custom role.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `Location information for the user. Please look at [Supported Locale Ids](https://docs.opsgenie.com/docs/supported-locales) for available locales.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Timezone information of the user. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Opsgenie User.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The Full Name of the User.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The Role assigned to the User. Either a built-in such as 'Owner', 'Admin' or 'User' - or the name of a custom role.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `Location information for the user. Please look at [Supported Locale Ids](https://docs.opsgenie.com/docs/supported-locales) for available locales.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Timezone information of the user. Please look at [Supported Timezone Ids](https://docs.opsgenie.com/docs/supported-timezone-ids) for available timezones.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"opsgenie_escalation": 0,
		"opsgenie_heartbeat":  1,
		"opsgenie_schedule":   2,
		"opsgenie_service":    3,
		"opsgenie_team":       4,
		"opsgenie_user":       5,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
