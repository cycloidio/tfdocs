package cohesivenet

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_config",
			Category:         "Resources",
			ShortDescription: `Launches, licences and configures one or many VNS3 controllers in conjunction with the AWS Terraform provider.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_firewall_fwset",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_firewall_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_firewall_rules",
			Category:         "Resources",
			ShortDescription: `Creates firewall rules using the Cohesive simplified IP tables syntax.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_https_certs",
			Category:         "Resources",
			ShortDescription: `Uploads HTTPS certificates to the VNS3 controller UI`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_ipsec_ebpg_peers",
			Category:         "Resources",
			ShortDescription: `Creates eBGP peer in conjunction with vns3_ipsec_endpoints resource.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_ipsec_endpoints",
			Category:         "Resources",
			ShortDescription: `Creates IPsec endpoint`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_link",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_peers",
			Category:         "Resources",
			ShortDescription: `Sets controller peering.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_plugin_image",
			Category:         "Resources",
			ShortDescription: `Imports a plugin image in the VNS3 controller.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_plugin_images",
			Category:         "Resources",
			ShortDescription: `Imports a plugin image in the VNS3 controller.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_plugin_instance",
			Category:         "Resources",
			ShortDescription: `Launches a plugin instance from a plugin image already imported into the VNS3 controller.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_plugin_instances",
			Category:         "Resources",
			ShortDescription: `Launches a plugin instance from a plugin image already imported into the VNS3 controller.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cohesivenet_vns3_routes",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cohesivenet_vns3_config":           0,
		"cohesivenet_vns3_firewall_fwset":   1,
		"cohesivenet_vns3_firewall_rule":    2,
		"cohesivenet_vns3_firewall_rules":   3,
		"cohesivenet_vns3_https_certs":      4,
		"cohesivenet_vns3_ipsec_ebpg_peers": 5,
		"cohesivenet_vns3_ipsec_endpoints":  6,
		"cohesivenet_vns3_link":             7,
		"cohesivenet_vns3_peers":            8,
		"cohesivenet_vns3_plugin_image":     9,
		"cohesivenet_vns3_plugin_images":    10,
		"cohesivenet_vns3_plugin_instance":  11,
		"cohesivenet_vns3_plugin_instances": 12,
		"cohesivenet_vns3_routes":           13,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
