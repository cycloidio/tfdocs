package buddy

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_group",
			Category:         "Resources",
			ShortDescription: `Create and manage a user's group Workspace administrator rights are required Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_group_member",
			Category:         "Resources",
			ShortDescription: `Create and manage a workspace group member Workspace administrator rights are required Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_integration",
			Category:         "Resources",
			ShortDescription: `Create and manage an integration Token scopes required: INTEGRATION_ADD, INTEGRATION_MANAGE, INTEGRATION_INFO`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_member",
			Category:         "Resources",
			ShortDescription: `Create and manage a workspace member Workspace administrator rights are required Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_permission",
			Category:         "Resources",
			ShortDescription: `Create and manage a workspace permission (role) Workspace administrator rights are required Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_pipeline",
			Category:         "Resources",
			ShortDescription: `Create and manage a pipeline Token scopes required: WORKSPACE, EXECUTION_MANAGE, EXECUTION_INFO`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_profile",
			Category:         "Resources",
			ShortDescription: `Manage a user profile Token scope required: USER_INFO`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_profile_email",
			Category:         "Resources",
			ShortDescription: `Create and manage a user's email Token scopes required: MANAGE_EMAILS, USER_EMAIL`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_profile_public_key",
			Category:         "Resources",
			ShortDescription: `Create and manage a user's public key Token scope required: USER_KEY`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_project",
			Category:         "Resources",
			ShortDescription: `Create and manage a workspace project Workspace administrator rights are required Token scopes required: WORKSPACE, PROJECT_DELETE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_project_group",
			Category:         "Resources",
			ShortDescription: `Manage a workspace project group permission Workspace administrator rights are required Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_project_member",
			Category:         "Resources",
			ShortDescription: `Manage a member's permission (role) in a project Workspace administrator rights are required Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_variable",
			Category:         "Resources",
			ShortDescription: `Create and manage a variable Workspace administrator rights are required Token scopes required: WORKSPACE, VARIABLE_ADD, VARIABLE_MANAGE, VARIABLE_INFO`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_variable_ssh_key",
			Category:         "Resources",
			ShortDescription: `Create and manage a variable of SSH key type Workspace administrator rights are required Token scope required: WORKSPACE, VARIABLE_ADD, VARIABLE_MANAGE, VARIABLE_INFO`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_webhook",
			Category:         "Resources",
			ShortDescription: `Create and manage a workspace webhook Workspace administrator rights are required Token scopes required: WORKSPACE, WEBHOOK_ADD, WEBHOOK_MANAGE, WEBHOOK_INFO`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "buddy_buddy_workspace",
			Category:         "Resources",
			ShortDescription: `Create and manage a workspace Invite-only token is required. Contact support@buddy.works for more details Token scope required: WORKSPACE`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"buddy_buddy_group":              0,
		"buddy_buddy_group_member":       1,
		"buddy_buddy_integration":        2,
		"buddy_buddy_member":             3,
		"buddy_buddy_permission":         4,
		"buddy_buddy_pipeline":           5,
		"buddy_buddy_profile":            6,
		"buddy_buddy_profile_email":      7,
		"buddy_buddy_profile_public_key": 8,
		"buddy_buddy_project":            9,
		"buddy_buddy_project_group":      10,
		"buddy_buddy_project_member":     11,
		"buddy_buddy_variable":           12,
		"buddy_buddy_variable_ssh_key":   13,
		"buddy_buddy_webhook":            14,
		"buddy_buddy_workspace":          15,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
