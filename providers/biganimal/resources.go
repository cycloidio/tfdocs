package biganimal

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "biganimal_aws_connection",
			Category:         "Resources",
			ShortDescription: `The awsconnection resource is used to make connection between your project and AWS. o obtain the necessary input parameters, please refer to [Connect CSP](https://www.enterprisedb.com/docs/biganimal/latest/gettingstarted/02connectingtoyourcloud/connecting_aws/).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "biganimal_azure_connection",
			Category:         "Resources",
			ShortDescription: `The azureconnection resource is used to make connection between your project and Azure. To obtain the necessary input parameters, please refer to [Connect CSP](https://www.enterprisedb.com/docs/biganimal/latest/gettingstarted/02connectingtoyourcloud/connecting_azure/).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "biganimal_cluster",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "biganimal_faraway_replica",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "biganimal_project",
			Category:         "Resources",
			ShortDescription: `The project resource is used to manage projects in your organization. See Managing projects https://www.enterprisedb.com/docs/biganimal/latest/administering_cluster/projects/ for more details. Newly created projects are not automatically connected to your cloud providers. Please visit Connecting your cloud https://www.enterprisedb.com/docs/biganimal/latest/getting_started/02_connecting_to_your_cloud/ for more details.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "biganimal_region",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"biganimal_aws_connection":   0,
		"biganimal_azure_connection": 1,
		"biganimal_cluster":          2,
		"biganimal_faraway_replica":  3,
		"biganimal_project":          4,
		"biganimal_region":           5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
