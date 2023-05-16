package infra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "infra_access_key",
			Category:         "Resources",
			ShortDescription: `Provides an Infra access key. This resource can be used to create and manage connector access keys. -> This resource requires Infra server version 0.20.0 or higher.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infra_grant",
			Category:         "Resources",
			ShortDescription: `Provides an Infra grant. This resource can be used to assign grants to users or groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infra_group",
			Category:         "Resources",
			ShortDescription: `Groups are used in Infra to manage a collection of users. A group can then be associated with a role and cluster via a grant and all users with the group will gain that role and and corresponding access to the cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infra_group_membership",
			Category:         "Resources",
			ShortDescription: `Provides an Infra user grant. This resource can be used to assign groups to users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infra_identity_provider",
			Category:         "Resources",
			ShortDescription: `The infra identity provider resource can be used to create and manage identity providers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "infra_user",
			Category:         "Resources",
			ShortDescription: `Infra user resource creates a user with a specified name. The name must be an email address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"infra_access_key":        0,
		"infra_grant":             1,
		"infra_group":             2,
		"infra_group_membership":  3,
		"infra_identity_provider": 4,
		"infra_user":              5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
