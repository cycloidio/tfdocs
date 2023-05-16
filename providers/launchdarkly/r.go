package launchdarkly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_audit_log_subscription",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly integration audit log subscriptions.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"audit",
				"log",
				"subscription",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_custom_role",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly custom roles.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"custom",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_destination",
			Category:         "Guides",
			ShortDescription: `Interact with the LaunchDarkly data export destinations API.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"destination",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_environment",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly environments.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"environment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_feature_flag",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly feature flags.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"feature",
				"flag",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_feature_flag_environment",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly environment-specific feature flag attributes.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"feature",
				"flag",
				"environment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_flag_trigger",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly flag triggers.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"flag",
				"trigger",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_metric",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly metrics.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"metric",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_project",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly projects.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"project",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_segment",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly segments.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"segment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_team",
			Category:         "Guides",
			ShortDescription: `Create and manage a LaunchDarkly team.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"team",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_team_member",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly team members.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"team",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_webhook",
			Category:         "Guides",
			ShortDescription: `Create and manage LaunchDarkly webhooks.`,
			Description:      ``,
			Keywords: []string{
				"guides",
				"webhook",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"launchdarkly_audit_log_subscription":   0,
		"launchdarkly_custom_role":              1,
		"launchdarkly_destination":              2,
		"launchdarkly_environment":              3,
		"launchdarkly_feature_flag":             4,
		"launchdarkly_feature_flag_environment": 5,
		"launchdarkly_flag_trigger":             6,
		"launchdarkly_metric":                   7,
		"launchdarkly_project":                  8,
		"launchdarkly_segment":                  9,
		"launchdarkly_team":                     10,
		"launchdarkly_team_member":              11,
		"launchdarkly_webhook":                  12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
