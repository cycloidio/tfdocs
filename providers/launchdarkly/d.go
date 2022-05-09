package launchdarkly

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_audit_log_subscription",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly audit log subscriptions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_environment",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly environments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_feature_flag",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly feature flags.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_feature_flag_environment",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly environment-specific feature flag configurations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_flag_trigger",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly flag trigers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_metric",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly metrics.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_project",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly projects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_segment",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly segments.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_team_member",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly team members.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_team_members",
			Category:         "Data Sources",
			ShortDescription: `Get information about multiple LaunchDarkly team members.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "launchdarkly_webhook",
			Category:         "Data Sources",
			ShortDescription: `Get information about LaunchDarkly webhooks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"launchdarkly_audit_log_subscription":   0,
		"launchdarkly_environment":              1,
		"launchdarkly_feature_flag":             2,
		"launchdarkly_feature_flag_environment": 3,
		"launchdarkly_flag_trigger":             4,
		"launchdarkly_metric":                   5,
		"launchdarkly_project":                  6,
		"launchdarkly_segment":                  7,
		"launchdarkly_team_member":              8,
		"launchdarkly_team_members":             9,
		"launchdarkly_webhook":                  10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
