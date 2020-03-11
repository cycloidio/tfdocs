package oktaasa

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "oktaasa_assign_group",
			Category:         "Resources",
			ShortDescription: `The oktaasa_assign_group resource assigns control access levels in group definitions in Okta's ASA.`,
			Description:      ``,
			Keywords: []string{
				"assign",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oktaasa_create_group",
			Category:         "Resources",
			ShortDescription: `The oktaasa_create_group resource creates projects in Okta's ASA.`,
			Description:      ``,
			Keywords: []string{
				"create",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oktaasa_enrollment_token",
			Category:         "Resources",
			ShortDescription: `The oktaasa_token resource creates enrollment tokens which are base64-encoded objects with metadata that Okta's ASA Agent can configure itself from. Enrollment is the process where Okta's ASA agent configures a server to be managed by a specific project.`,
			Description:      ``,
			Keywords: []string{
				"enrollment",
				"token",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oktaasa_project",
			Category:         "Resources",
			ShortDescription: `The oktaasa_project resource creates projects in Okta's ASA.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"oktaasa_assign_group":     0,
		"oktaasa_create_group":     1,
		"oktaasa_enrollment_token": 2,
		"oktaasa_project":          3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
