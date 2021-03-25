package nutanixkps

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_application",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Kubernetes Application`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_node",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Service Domain Node. A Service Domain Node is the baremetal/virtual machine that is being managed by Karbon Platform Services.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_project",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Project.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_servicebinding",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Service Binding.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_servicedomain",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Service Domain.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_serviceinstance",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Service Instance.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_storageprofile",
			Category:         "Resources",
			ShortDescription: `Describes a Karbon Platform Services Storage Profile.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "nutanixkps_vm_cloud_config",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"nutanixkps_application":     0,
		"nutanixkps_node":            1,
		"nutanixkps_project":         2,
		"nutanixkps_servicebinding":  3,
		"nutanixkps_servicedomain":   4,
		"nutanixkps_serviceinstance": 5,
		"nutanixkps_storageprofile":  6,
		"nutanixkps_vm_cloud_config": 7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
