package nomad

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nomad_acl_policy",
			Category:         "Resources",
			ShortDescription: `Manages an ACL policy registered on the Nomad server.`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_acl_token",
			Category:         "Resources",
			ShortDescription: `Manages an ACL token in Nomad.`,
			Description:      ``,
			Keywords: []string{
				"acl",
				"token",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_job",
			Category:         "Resources",
			ShortDescription: `Manages the lifecycle of registering and deregistering Nomad jobs (applications).`,
			Description:      ``,
			Keywords: []string{
				"job",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_namespace",
			Category:         "Resources",
			ShortDescription: `Provisions a namespace within a Nomad cluster.`,
			Description:      ``,
			Keywords: []string{
				"namespace",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_quota_specification",
			Category:         "Resources",
			ShortDescription: `Manages a quota specification in a Nomad cluster.`,
			Description:      ``,
			Keywords: []string{
				"quota",
				"specification",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nomad_sentinel_policy",
			Category:         "Resources",
			ShortDescription: `Manages a Sentinel policy registered on the Nomad server.`,
			Description:      ``,
			Keywords: []string{
				"sentinel",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"nomad_acl_policy":          0,
		"nomad_acl_token":           1,
		"nomad_job":                 2,
		"nomad_namespace":           3,
		"nomad_quota_specification": 4,
		"nomad_sentinel_policy":     5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
