package davinci

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "davinci_application",
			Category:         "Application",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"application",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "davinci_connection",
			Category:         "Connection",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"connection",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "davinci_flow",
			Category:         "Flow",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"flow",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "davinci_variable",
			Category:         "Variable",
			ShortDescription: ``,
			Description:      ``,
			Keywords: []string{
				"variable",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"davinci_application": 0,
		"davinci_connection":  1,
		"davinci_flow":        2,
		"davinci_variable":    3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
