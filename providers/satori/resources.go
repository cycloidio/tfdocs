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
			Type:             "satori_custom_taxonomy_category",
			Category:         "Resources",
			ShortDescription: `satori_custom_taxonomy_category resource allows defining custom taxonomy categories.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "satori_custom_taxonomy_classifier",
			Category:         "Resources",
			ShortDescription: `Custom taxonomy classifier.`,
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
			Type:             "satori_datastore",
			Category:         "Resources",
			ShortDescription: `satori_datastore resource allows defining datastore(s)`,
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
			Type:             "satori_masking_profile",
			Category:         "Resources",
			ShortDescription: `masking_profile resource allows defining masking profiles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "satori_security_policy",
			Category:         "Resources",
			ShortDescription: `security_policy resource allows defining security policies.`,
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

		"satori_access_rule":                0,
		"satori_custom_taxonomy_category":   1,
		"satori_custom_taxonomy_classifier": 2,
		"satori_dataset":                    3,
		"satori_dataset_definition":         4,
		"satori_datastore":                  5,
		"satori_directory_group":            6,
		"satori_masking_profile":            7,
		"satori_security_policy":            8,
		"satori_self_service_access_rule":   9,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
