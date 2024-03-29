package shoreline

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "shoreline_action",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_alarm",
			Category:         "Resources",
			ShortDescription: `Alarms are fully-customizable Metric or status checks that automatically trigger remediation Actions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_bot",
			Category:         "Resources",
			ShortDescription: `Alarms use Bots to execute automated Actions.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_circuit_breaker",
			Category:         "Resources",
			ShortDescription: `Shoreline circuit_breaker. An automatic rate limit on actions. See the Shoreline CircuitBreakers Documentation https://docs.shoreline.io/circuit_breakers for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_file",
			Category:         "Resources",
			ShortDescription: `Automatically distribute artifacts across your Shoreline Resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_metric",
			Category:         "Resources",
			ShortDescription: `Shoreline metric. A periodic measurement of a system property.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_notebook",
			Category:         "Resources",
			ShortDescription: `Notebooks replace static runbooks by capturing interactive debug and remediation sessions in a convenient UI.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_principal",
			Category:         "Resources",
			ShortDescription: `Shoreline principal. An authorization group (e.g. Okta groups). Note: Admin privilege (in Shoreline) to create principal objects.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_resource",
			Category:         "Resources",
			ShortDescription: `Shoreline resource. A server or compute resource in the system (e.g. host, pod, container).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"shoreline_action":          0,
		"shoreline_alarm":           1,
		"shoreline_bot":             2,
		"shoreline_circuit_breaker": 3,
		"shoreline_file":            4,
		"shoreline_metric":          5,
		"shoreline_notebook":        6,
		"shoreline_principal":       7,
		"shoreline_resource":        8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
