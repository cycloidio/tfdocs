package rancher

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "rancher_certificate",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher certificate.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The setting name.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Required) The ID of the environment. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The certificate CN.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The certificate algorithm.`,
				},
				resource.Attribute{
					Name:        "cert_fingerprint",
					Description: `The certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `The certificate expiration date.`,
				},
				resource.Attribute{
					Name:        "issued_at",
					Description: `The certificate creation date.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The certificate issuer.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate serial number.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `The list of certificate Subject Alternative Names.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The certificate version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The certificate CN.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `The certificate algorithm.`,
				},
				resource.Attribute{
					Name:        "cert_fingerprint",
					Description: `The certificate fingerprint.`,
				},
				resource.Attribute{
					Name:        "expires_at",
					Description: `The certificate expiration date.`,
				},
				resource.Attribute{
					Name:        "issued_at",
					Description: `The certificate creation date.`,
				},
				resource.Attribute{
					Name:        "issuer",
					Description: `The certificate issuer.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `The certificate serial number.`,
				},
				resource.Attribute{
					Name:        "subject_alternative_names",
					Description: `The list of certificate Subject Alternative Names.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The certificate version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_environment",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher environment.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The setting name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The environment description.`,
				},
				resource.Attribute{
					Name:        "orchestration",
					Description: `The environment orchestration engine.`,
				},
				resource.Attribute{
					Name:        "project_template_id",
					Description: `The environment project template ID.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The environment members.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The environment description.`,
				},
				resource.Attribute{
					Name:        "orchestration",
					Description: `The environment orchestration engine.`,
				},
				resource.Attribute{
					Name:        "project_template_id",
					Description: `The environment project template ID.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `The environment members.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "rancher_setting",
			Category:         "Data Sources",
			ShortDescription: `Get information on a Rancher setting.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The setting name. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `the settting's value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `the settting's value.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"rancher_certificate": 0,
		"rancher_environment": 1,
		"rancher_setting":     2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
