package ome

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ome_configuration_baseline",
			Category:         "Resources",
			ShortDescription: `Resource for managing configuration baselines on OpenManage Enterprise. Updates are supported for all the parameters. When schedule is true, following parameters are considered: notify_on_schedule, cron, email_addresses, output_format`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ome_configuration_compliance",
			Category:         "Resources",
			ShortDescription: `Resource for managing configuration baselines remediation. Updates are supported for the following parameters: target_devices, job_retry_count, sleep_interval, run_later, cron.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ome_deployment",
			Category:         "Resources",
			ShortDescription: `Resource for managing template deployment on OpenManage Enterprise. Updates are supported for the following parameters: device_ids, device_servicetags, boot_to_network_iso, forced_shutdown, options_time_to_wait_before_shutdown, power_state_off, options_precheck_only, options_strict_checking_vlan, options_continue_on_warning, run_later, cron, device_attributes, job_retry_count, sleep_interval.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ome_template",
			Category:         "Resources",
			ShortDescription: `Resource for managing template on OpenManage Enterprise.Updates are supported for the following parameters: name, description, attributes, job_retry_count, sleep_interval, identity_pool_name, vlan.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"ome_configuration_baseline":   0,
		"ome_configuration_compliance": 1,
		"ome_deployment":               2,
		"ome_template":                 3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
