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
			ShortDescription: `Shoreline action. A command that can be run. See the Shoreline Actions Documentation https://docs.shoreline.io/actions for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_alarm",
			Category:         "Resources",
			ShortDescription: `Shoreline alarm. A condition that triggers Alerts or Actions. See the Shoreline Alarms Documentation https://docs.shoreline.io/alarms for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_bot",
			Category:         "Resources",
			ShortDescription: `Shoreline bot. An automation that ties an Action to an Alert. See the Shoreline Bots Documentation https://docs.shoreline.io/bots for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_file",
			Category:         "Resources",
			ShortDescription: `Shoreline file. A datafile that is automatically copied/distributed to defined Resources. See the Shoreline OpCp Documentation https://docs.shoreline.io/op/commands/cp for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_metric",
			Category:         "Resources",
			ShortDescription: `Shoreline metric. A periodic measurement of a system property. See the Shoreline Metrics Documentation https://docs.shoreline.io/metrics for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "shoreline_resource",
			Category:         "Resources",
			ShortDescription: `Shoreline resource. A server or compute resource in the system (e.g. host, pod, container). See the Shoreline Resources Documentation https://docs.shoreline.io/platform/resources for more info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"shoreline_action":   0,
		"shoreline_alarm":    1,
		"shoreline_bot":      2,
		"shoreline_file":     3,
		"shoreline_metric":   4,
		"shoreline_resource": 5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
