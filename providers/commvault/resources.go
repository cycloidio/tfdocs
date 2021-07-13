package commvault

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "commvault_amazon_hypervisor",
			Category:         "Virtualization",
			ShortDescription: `Use the commvault_amazon_hypervisor resource type to create or delete an Amazon hypervisor in the CommCell environment.`,
			Description:      ``,
			Keywords: []string{
				"virtualization",
				"amazon",
				"hypervisor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_aws_storage",
			Category:         "Storage",
			ShortDescription: `Use the commvault_aws_storage resource type to create or delete a AWS Cloud Storage in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"aws",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_azure_hypervisor",
			Category:         "Virtualization",
			ShortDescription: `Use the commvault_azure_hypervisor resource type to create or delete an Azure hypervisor in the CommCell environment.`,
			Description:      ``,
			Keywords: []string{
				"virtualization",
				"azure",
				"hypervisor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_azure_storage",
			Category:         "Storage",
			ShortDescription: `Use the commvault_azure_storage resource type to create or delete a Azure Cloud Storage in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"azure",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_company",
			Category:         "Company",
			ShortDescription: `Use the commvault_company resource type to create or delete a Company in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"company",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_disk_storage",
			Category:         "Storage",
			ShortDescription: `Use the commvault_disk_storage resource type to create or delete a Disk Storage in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"disk",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_google_storage",
			Category:         "Storage",
			ShortDescription: `Use the commvault_google_storage resource type to create or delete a Google Cloud Storage in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"google",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_install_ma",
			Category:         "Install",
			ShortDescription: `Use the commvault_install_ma resource type to Install or Uninstall a Media Agent in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"install",
				"ma",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_plan",
			Category:         "Plans",
			ShortDescription: `Use the commvault_plan resource type to create or delete a Plans in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"plans",
				"plan",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_plan_to_vm",
			Category:         "Virtualization",
			ShortDescription: `Use the commvault_plan_to_vm resource type to associate plan to a VM in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"virtualization",
				"plan",
				"to",
				"vm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_security_association",
			Category:         "Security",
			ShortDescription: `Use the commvault_security_association resource type to Security Associations in the Commcell environment.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"association",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_user",
			Category:         "Security",
			ShortDescription: `Use the commvault_user resource type to create or delete an User in the CommCell environment.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_vm_group",
			Category:         "Virtualization",
			ShortDescription: `Use the commvault_vm_group resource type to create and delete a VM group in the CommCell environment.`,
			Description:      ``,
			Keywords: []string{
				"virtualization",
				"vm",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "commvault_vmware_hypervisor",
			Category:         "Virtualization",
			ShortDescription: `Use the commvault_vmware_hypervisor resource type to create and delete a VMware hypervisor in the CommCell environment.`,
			Description:      ``,
			Keywords: []string{
				"virtualization",
				"vmware",
				"hypervisor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"commvault_amazon_hypervisor":    0,
		"commvault_aws_storage":          1,
		"commvault_azure_hypervisor":     2,
		"commvault_azure_storage":        3,
		"commvault_company":              4,
		"commvault_disk_storage":         5,
		"commvault_google_storage":       6,
		"commvault_install_ma":           7,
		"commvault_plan":                 8,
		"commvault_plan_to_vm":           9,
		"commvault_security_association": 10,
		"commvault_user":                 11,
		"commvault_vm_group":             12,
		"commvault_vmware_hypervisor":    13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
