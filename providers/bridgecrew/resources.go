package bridgecrew

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_complex_policy",
			Category:         "Resources",
			ShortDescription: `Create a new custom 'complex' security policy for the Bridgecrew Platform`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_policy",
			Category:         "Resources",
			ShortDescription: `Create a new custom YAML based security policy for the Bridgecrew Platform`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "bridgecrew_simple_policy",
			Category:         "Resources",
			ShortDescription: `Create a new custom 'simple' security policy for the Bridgecrew Platform`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"bridgecrew_complex_policy": 0,
		"bridgecrew_policy":         1,
		"bridgecrew_simple_policy":  2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
