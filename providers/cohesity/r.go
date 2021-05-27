package cohesity

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cohesity_cloud_edition_cluster",
			Category:         "Resources",
			ShortDescription: `Create cloud edition cluster, apply license key and destroy cluster.`,
			Description:      ``,
			Keywords: []string{
				"cloud",
				"edition",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_job_run",
			Category:         "Resources",
			ShortDescription: `Run a Protection Job on Cohesity Cluster`,
			Description:      ``,
			Keywords: []string{
				"job",
				"run",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_job_vmware",
			Category:         "Resources",
			ShortDescription: `Create a VMware Protection Job on Cohesity Cluster`,
			Description:      ``,
			Keywords: []string{
				"job",
				"vmware",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_physical_edition_cluster",
			Category:         "Resources",
			ShortDescription: `Create physical edition cluster, apply license key and destroy cluster.`,
			Description:      ``,
			Keywords: []string{
				"physical",
				"edition",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_restore_vmware_vm",
			Category:         "Resources",
			ShortDescription: `Restore a VMware VM`,
			Description:      ``,
			Keywords: []string{
				"restore",
				"vmware",
				"vm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_source_vmware",
			Category:         "Resources",
			ShortDescription: `Register a VMware source on Cohesity Cluster`,
			Description:      ``,
			Keywords: []string{
				"source",
				"vmware",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesity_virtual_edition_cluster",
			Category:         "Resources",
			ShortDescription: `Create virtual edition cluster, apply license key and destroy cluster.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"edition",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cohesity_cloud_edition_cluster":    0,
		"cohesity_job_run":                  1,
		"cohesity_job_vmware":               2,
		"cohesity_physical_edition_cluster": 3,
		"cohesity_restore_vmware_vm":        4,
		"cohesity_source_vmware":            5,
		"cohesity_virtual_edition_cluster":  6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
