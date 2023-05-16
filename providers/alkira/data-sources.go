package alkira

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alkira_billing_tag",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing billing tag.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_byoip",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing BYOIP Prefix.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_aruba_edge",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Aruba Edge connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_aws_vpc",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing AWS VPC connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_azure_expressroute",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Azure ExpressRoute connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_azure_vnet",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Azure Vnet connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_cisco_sdwan",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Cisco Sdwan connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_gcp_vpc",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Gcp Vpc connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_internet_exit",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Internet Exit connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_ipsec",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing ipsec connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_oci_vcn",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing OCI VCN connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get information on an existing credential.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_group",
			Category:         "Data Sources",
			ShortDescription: `This data source allows to retrieve an existing group by its name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_group_user",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing user group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_list_as_path",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing As Path List.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_list_community",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing community list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_list_extended_community",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing extended community list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_list_global_cidr",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing Global CIDR List.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_nat_rule",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing NAT policy rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_prefix_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing policy prefix list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_rule",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing policy rule.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_rule_list",
			Category:         "Data Sources",
			ShortDescription: `Use this data source to get an existing policy rule list.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_segment",
			Category:         "Data Sources",
			ShortDescription: `The segment data source allows a segment to be retrieved by its name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"alkira_billing_tag":                  0,
		"alkira_byoip":                        1,
		"alkira_connector_aruba_edge":         2,
		"alkira_connector_aws_vpc":            3,
		"alkira_connector_azure_expressroute": 4,
		"alkira_connector_azure_vnet":         5,
		"alkira_connector_cisco_sdwan":        6,
		"alkira_connector_gcp_vpc":            7,
		"alkira_connector_internet_exit":      8,
		"alkira_connector_ipsec":              9,
		"alkira_connector_oci_vcn":            10,
		"alkira_credential":                   11,
		"alkira_group":                        12,
		"alkira_group_user":                   13,
		"alkira_list_as_path":                 14,
		"alkira_list_community":               15,
		"alkira_list_extended_community":      16,
		"alkira_list_global_cidr":             17,
		"alkira_policy":                       18,
		"alkira_policy_nat_rule":              19,
		"alkira_policy_prefix_list":           20,
		"alkira_policy_rule":                  21,
		"alkira_policy_rule_list":             22,
		"alkira_segment":                      23,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
