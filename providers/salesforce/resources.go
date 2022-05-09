package salesforce

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "salesforce_profile",
			Category:         "Resources",
			ShortDescription: `Profile Resource for the Salesforce Provider. Please note that Users must have a Profile assigned to them, Profiles cannot be deleted if a User is assigned to it, and Salesforce does not allow the deletion of Users, only deactivation. Terraform will warn after destroy of a User that it has only been deactivated and now removed from state. A common issue with this pattern is a Profile and User created in tandem will fail to delete the Profile on destroy due to the lingering assignment. Should you wish to destroy a created Profile, it's advised that an apply that moves all affected Users to a static Profile be run first, after which the Profile can be safely destroyed.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "salesforce_user",
			Category:         "Resources",
			ShortDescription: `User Resource for the Salesforce Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "salesforce_user_role",
			Category:         "Resources",
			ShortDescription: `User Role Resource for the Salesforce Provider`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"salesforce_profile":   0,
		"salesforce_user":      1,
		"salesforce_user_role": 2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
