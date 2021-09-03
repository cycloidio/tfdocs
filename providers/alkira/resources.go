package alkira

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "alkira_billing_tag",
			Category:         "Resources",
			ShortDescription: `Manage Billing Tag.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_aws_vpc",
			Category:         "Resources",
			ShortDescription: `Manage AWS Cloud Connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_azure_vnet",
			Category:         "Resources",
			ShortDescription: `Manage Azure Cloud Connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_cisco_sdwan",
			Category:         "Resources",
			ShortDescription: `Manage Cisco SD-WAN Connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_gcp_vpc",
			Category:         "Resources",
			ShortDescription: `Manage GCP Cloud Connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_internet_exit",
			Category:         "Resources",
			ShortDescription: `Manage Internet Exit. An internet exit is an exit from the CXP to theinternet and allows the traffic fromthe various Users & Sites or Cloud Connectors toflow towards the Internet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_connector_ipsec",
			Category:         "Resources",
			ShortDescription: `Manage IPSec Connector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential_aws_vpc",
			Category:         "Resources",
			ShortDescription: `Manage AWS credential for authentication. The following methods are supported: Static credentialsEnvironment variables *** Static Credentials: Static credentials can be provided by adding an aws_access_keyand aws_secret_key in-line in the AWS provider block. *** Environment Variables: You can provide your credentials via the AWS_ACCESS_KEY_ID andAWS_SECRET_ACCESS_KEY, environment variables, representing yourAWS Access Key and AWS Secret Key, respectively.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential_azure_vnet",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential_cisco_sdwan",
			Category:         "Resources",
			ShortDescription: `Manage Cisco SD-WAN credential for authentication. The following methods are supported: Static credentialsEnvironment variables *** Static Credentials: Static credentials can be provided by adding an usernameand password in-line in the CISCO SD-WAN block. *** Environment Variables: You can provide your credentials via the CISCO_SDWAN_USERNAME andCISCO_SDWAN_PASSWORD, environment variables, representing yourCisco SD-WAN username and password, respectively.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential_gcp_vpc",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential_pan",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_credential_pan_instance",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_group",
			Category:         "Resources",
			ShortDescription: `Manage groups Groups can contain one or many connectors across different segments. This grouping of connectors can be for policy enforcement purposes or for monitoring purposes within the network. It allows for easier policy assignment by assigning policies to the entire group at the same time instead of having to assign them individually.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_internet_application",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_prefix_list",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_rule",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_policy_rule_list",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_segment",
			Category:         "Resources",
			ShortDescription: `This resource manages segments. A Segment is a section of a network isolated from one anotherto make it possible to more effectively control who has accessto what. Segmentation also allows for segregation of resourcesbetween segments for security and isolation purposes.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_service_pan",
			Category:         "Resources",
			ShortDescription: `Manage PAN firewall.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "alkira_tenant_network",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"alkira_billing_tag":             0,
		"alkira_connector_aws_vpc":       1,
		"alkira_connector_azure_vnet":    2,
		"alkira_connector_cisco_sdwan":   3,
		"alkira_connector_gcp_vpc":       4,
		"alkira_connector_internet_exit": 5,
		"alkira_connector_ipsec":         6,
		"alkira_credential_aws_vpc":      7,
		"alkira_credential_azure_vnet":   8,
		"alkira_credential_cisco_sdwan":  9,
		"alkira_credential_gcp_vpc":      10,
		"alkira_credential_pan":          11,
		"alkira_credential_pan_instance": 12,
		"alkira_group":                   13,
		"alkira_internet_application":    14,
		"alkira_policy":                  15,
		"alkira_policy_prefix_list":      16,
		"alkira_policy_rule":             17,
		"alkira_policy_rule_list":        18,
		"alkira_segment":                 19,
		"alkira_service_pan":             20,
		"alkira_tenant_network":          21,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
