package victorops

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "victorops_escalation_policy",
			Category:         "Resources",
			ShortDescription: `Manage Escalation Policies in VictorOps.`,
			Description:      ``,
			Keywords: []string{
				"escalation",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this escalation policy`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The team_id of the team for which you want to create this escalation policy`,
				},
				resource.Attribute{
					Name:        "ignore_custom_paging_policies",
					Description: `(Optional) ` + "`" + `true` + "`" + `/` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "step",
					Description: `(Required) - The escalation policy step defined in the following structure ` + "`" + `` + "`" + `` + "`" + `hcl step { timeout = [time-out duration in seconds] entries = [ { type = [ rotationalGroup | targetPolicy ] slug = [ rotatioGroup slug | next escalation policy ID ] }, ] } ` + "`" + `` + "`" + `` + "`" + ` ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Import is not currently supported`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the escalation policy. ## Import Import is not currently supported`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "victorops_routing_key",
			Category:         "Resources",
			ShortDescription: `Creates and manages a routing key in VictorOps`,
			Description:      ``,
			Keywords: []string{
				"routing",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Routing key name.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `(Required) A list of escalation policy ids to route alerts to. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the routing key. ## Import Import is not currently supported`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the routing key. ## Import Import is not currently supported`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "victorops_team",
			Category:         "Resources",
			ShortDescription: `Creates and manages a team in VictorOps.`,
			Description:      ``,
			Keywords: []string{
				"team",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of this team ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Import is not currently supported`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the team. ## Import Import is not currently supported`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "victorops_team_membership",
			Category:         "Resources",
			ShortDescription: `Manages a user's team association within victorops`,
			Description:      ``,
			Keywords: []string{
				"team",
				"membership",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The id of the team we are adding this user to.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) The username of the victorops user to add to the team. ## Attributes Reference There are no additional attributes for this resource. ## Import Import is not currently supported`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "victorops_user",
			Category:         "Resources",
			ShortDescription: `Creates and manages a user in VictorOps.`,
			Description:      ``,
			Keywords: []string{
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "first_name",
					Description: `(Required) The first name of the user.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `(Required) The last name of the user.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `(Required) The username for this user.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `(Required) The user's email address.`,
				},
				resource.Attribute{
					Name:        "is_admin",
					Description: `DEPRECATED - the field and its value will be ignored if specified. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
				resource.Attribute{
					Name:        "default_email_contact_id",
					Description: `The ID for the default email contact ## Import Import is not currently supported`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the user.`,
				},
				resource.Attribute{
					Name:        "default_email_contact_id",
					Description: `The ID for the default email contact ## Import Import is not currently supported`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"victorops_escalation_policy": 0,
		"victorops_routing_key":       1,
		"victorops_team":              2,
		"victorops_team_membership":   3,
		"victorops_user":              4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
