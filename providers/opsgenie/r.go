package aws

import (
	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*Resource{

		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_team",
			Category:         "Resources",
			ShortDescription: `Manages a Team within OpsGenie.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name associated with this team. OpsGenie defines that this must not be longer than 100 characters.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for this team.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `(Optional) A Member block as documented below. ` + "`" + `member` + "`" + ` supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username for the member to add to this Team.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The role for the user within the Team - can be either 'Admin' or 'User'. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the OpsGenie User. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opsgenie_team.team1 812be1a1-32c8-4666-a7fb-03ecc385106c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the OpsGenie User. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opsgenie_team.team1 812be1a1-32c8-4666-a7fb-03ecc385106c ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "opsgenie_user",
			Category:         "Resources",
			ShortDescription: `Manages a User within OpsGenie.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Argument{
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The email address associated with this user. OpsGenie defines that this must not be longer than 100 characters.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Required) The Full Name of the User.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `(Required) The Role assigned to the User. Either a built-in such as 'Owner', 'Admin' or 'User' - or the name of a custom role.`,
				},
				resource.Attribute{
					Name:        "locale",
					Description: `(Optional) Location information for the user. Please look at [Supported Locale Ids](https://www.opsgenie.com/docs/miscellaneous/supported-locales) for available locales - Defaults to "en_US".`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Timezone information of the user. Please look at [Supported Timezone Ids](https://www.opsgenie.com/docs/miscellaneous/supported-timezone-ids) for available timezones - Defaults to "America/New_York". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the OpsGenie User. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opsgenie_user.user da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Argument{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the OpsGenie User. ## Import Users can be imported using the ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import opsgenie_user.user da4faf16-5546-41e4-8330-4d0002b74048 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"opsgenie_team": 0,
		"opsgenie_user": 1,
	}
)

func GetResource(r string) (*resouce.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs]
}
