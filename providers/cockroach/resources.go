package cockroach

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cockroach_allow_list",
			Category:         "Resources",
			ShortDescription: `Allow list of IP range`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_client_ca_cert",
			Category:         "Resources",
			ShortDescription: `Manages client CA certs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_cluster",
			Category:         "Resources",
			ShortDescription: `Cluster Resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_cmek",
			Category:         "Resources",
			ShortDescription: `Customer-managed encryption keys (CMEK) resource for a single cluster`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_database",
			Category:         "Resources",
			ShortDescription: `Database`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_finalize_version_upgrade",
			Category:         "Resources",
			ShortDescription: `Utility resource that represents the one-time action of finalizing a cluster's pending CockroachDB version upgrade.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_log_export_config",
			Category:         "Resources",
			ShortDescription: `Log Export Config Resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_metric_export_cloudwatch_config",
			Category:         "Resources",
			ShortDescription: `Metric Export CloudWatch Config Resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_metric_export_datadog_config",
			Category:         "Resources",
			ShortDescription: `Metric Export Datadog Config Resource`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_private_endpoint_connection",
			Category:         "Resources",
			ShortDescription: `AWS PrivateLink Endpoint Connection`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_private_endpoint_services",
			Category:         "Resources",
			ShortDescription: `PrivateEndpointServices contains services that allow for VPC communication, either via PrivateLink (AWS) or Peering (GCP)`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_sql_user",
			Category:         "Resources",
			ShortDescription: `SQL user and password`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cockroach_user_role_grants",
			Category:         "Resources",
			ShortDescription: `Role grants`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"cockroach_allow_list":                      0,
		"cockroach_client_ca_cert":                  1,
		"cockroach_cluster":                         2,
		"cockroach_cmek":                            3,
		"cockroach_database":                        4,
		"cockroach_finalize_version_upgrade":        5,
		"cockroach_log_export_config":               6,
		"cockroach_metric_export_cloudwatch_config": 7,
		"cockroach_metric_export_datadog_config":    8,
		"cockroach_private_endpoint_connection":     9,
		"cockroach_private_endpoint_services":       10,
		"cockroach_sql_user":                        11,
		"cockroach_user_role_grants":                12,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
