package launchdarkly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_custom_role",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly custom roles.`,
			Description:      ``,
			Keywords: []string{
				"custom",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_environment",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly environments.`,
			Description:      ``,
			Keywords: []string{
				"environment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_feature_flag",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly feature flags.`,
			Description:      ``,
			Keywords: []string{
				"feature",
				"flag",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_feature_flag_environment",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly environment-specific feature flag attributes.`,
			Description:      ``,
			Keywords: []string{
				"feature",
				"flag",
				"environment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_project",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly projects.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_segment",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly segments.`,
			Description:      ``,
			Keywords: []string{
				"segment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_team_member",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly team members.`,
			Description:      ``,
			Keywords: []string{
				"team",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_webhook",
			Category:         "Resources",
			ShortDescription: `Create and manage LaunchDarkly webhooks.`,
			Description:      ``,
			Keywords: []string{
				"webhook",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"launchdarkly_custom_role":              0,
		"launchdarkly_environment":              1,
		"launchdarkly_feature_flag":             2,
		"launchdarkly_feature_flag_environment": 3,
		"launchdarkly_project":                  4,
		"launchdarkly_segment":                  5,
		"launchdarkly_team_member":              6,
		"launchdarkly_webhook":                  7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
