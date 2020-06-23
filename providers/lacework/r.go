package lacework

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_aws_cfg",
			Category:         "Resources",
			ShortDescription: `Create an manage AWS Config integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"aws",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework AWS Config integration name.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_aws_ct",
			Category:         "Resources",
			ShortDescription: `Create an manage AWS CloudTrail integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"aws",
				"ct",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework AWS CloudTrail integration name.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `(Required) The SQS Queue URL.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_azure_al",
			Category:         "Resources",
			ShortDescription: `Create an manage Azure Cloud Activity Log integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"azure",
				"al",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework Azure Config integration name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The directory tenant ID.`,
				},
				resource.Attribute{
					Name:        "queue_url",
					Description: `(Required) The storage queue URL.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_azure_cfg",
			Category:         "Resources",
			ShortDescription: `Create an manage Azure Cloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"azure",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework Azure Config integration name.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) The directory tenant ID.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_gcp_at",
			Category:         "Resources",
			ShortDescription: `Create an manage Google Cloud Audit Trail integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"gcp",
				"at",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework GCP Audit Trail integration name.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The organization or project ID.`,
				},
				resource.Attribute{
					Name:        "subscription",
					Description: `(Required) The subscription queue name.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "resource_level",
					Description: `(Optional) The integration level. Must be one of ` + "`" + `PROJECT` + "`" + ` or ` + "`" + `ORGANIZATION` + "`" + `. Defaults to ` + "`" + `PROJECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "lacework_integration_gcp_cfg",
			Category:         "Resources",
			ShortDescription: `Create an manage Google Cloud Config integrations`,
			Description:      ``,
			Keywords: []string{
				"integration",
				"gcp",
				"cfg",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Lacework GCP Config integration name.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required) The organization or project ID.`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `(Required) The credentials needed by the integration. See [Credentials](#credentials) below for details.`,
				},
				resource.Attribute{
					Name:        "resource_level",
					Description: `(Optional) The integration level. Must be one of ` + "`" + `PROJECT` + "`" + ` or ` + "`" + `ORGANIZATION` + "`" + `. Defaults to ` + "`" + `PROJECT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) The state of the external integration. Defaults to ` + "`" + `true` + "`" + `. ### Credentials ` + "`" + `credentials` + "`" + ` supports the following arguments:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"lacework_integration_aws_cfg":   0,
		"lacework_integration_aws_ct":    1,
		"lacework_integration_azure_al":  2,
		"lacework_integration_azure_cfg": 3,
		"lacework_integration_gcp_at":    4,
		"lacework_integration_gcp_cfg":   5,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
