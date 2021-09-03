package satori

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "satori_access_rule",
			Category:         "Resources",
			ShortDescription: `satori_access_rule resource allows defining dataset access rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "satori_dataset",
			Category:         "Resources",
			ShortDescription: `satori_dataset resource allows defining datasets.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "satori_dataset_definition",
			Category:         "Resources",
			ShortDescription: `satori_dataset_definition resource allows defining datasets (definition tab only).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "satori_directory_group",
			Category:         "Resources",
			ShortDescription: `satori_directory_group resource allows defining directory groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "satori_self_service_access_rule",
			Category:         "Resources",
			ShortDescription: `satori_self_service_access_rule resource allows defining dataset self service access rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"satori_access_rule":              0,
		"satori_dataset":                  1,
		"satori_dataset_definition":       2,
		"satori_directory_group":          3,
		"satori_self_service_access_rule": 4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
