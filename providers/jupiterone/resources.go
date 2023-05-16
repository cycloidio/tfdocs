package jupiterone

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "jupiterone_framework",
			Category:         "Resources",
			ShortDescription: `A custom compliance standard, benchmark, or questionnaire.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "jupiterone_frameworkitem",
			Category:         "Resources",
			ShortDescription: `A FrameworkItem (Requirement in the web UI) is a control that is part of a framework`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "jupiterone_group",
			Category:         "Resources",
			ShortDescription: `A compliance group is a child of a framework. Referred to as "Section" in the web UI. Refer to the resource_framework docs for example usage`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "jupiterone_libraryitem",
			Category:         "Resources",
			ShortDescription: `A Library Item (Control in the web UI) is a control that is independent of a framework, but can be associated with framework items (requirements). Refer to the resource_framework docs for example usage`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "jupiterone_question",
			Category:         "Resources",
			ShortDescription: `A saved JupiterOne Question`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "jupiterone_rule",
			Category:         "Resources",
			ShortDescription: `A saved JupiterOne question based alert`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"jupiterone_framework":     0,
		"jupiterone_frameworkitem": 1,
		"jupiterone_group":         2,
		"jupiterone_libraryitem":   3,
		"jupiterone_question":      4,
		"jupiterone_rule":          5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
