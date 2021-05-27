package vmc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vmc_connected_accounts",
			Category:         "Data Sources",
			ShortDescription: `A connected accounts data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Computed) Organization identifier.`,
				},
				resource.Attribute{
					Name:        "account_number",
					Description: `(Required) AWS account number. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The corresponding connected (customer) account UUID this connection is attached to.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The corresponding connected (customer) account UUID this connection is attached to.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vmc_customer_subnets",
			Category:         "Data Sources",
			ShortDescription: `A customer subnets data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Computed) Organization identifier.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The AWS specific (e.g us-west-2) or VMC specific region (e.g US_WEST_2) of the cloud resources to work in.`,
				},
				resource.Attribute{
					Name:        "num_hosts",
					Description: `(Optional) The number of hosts.`,
				},
				resource.Attribute{
					Name:        "connected_account_id",
					Description: `(Required) The linked connected account identifier.`,
				},
				resource.Attribute{
					Name:        "sddc_id",
					Description: `(Optional) SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "force_refresh",
					Description: `(Optional) When true, forces the mappings for datacenters to be refreshed for the connected account.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Optional) The server instance type to be used.`,
				},
				resource.Attribute{
					Name:        "sddc_type",
					Description: `(Optional) The sddc type to be used. (1NODE, SingleAZ, MultiAZ) ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "customer_available_zones",
					Description: `A list of AWS availability zones.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of AWS subnet IDs to create links to in the customer's account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "customer_available_zones",
					Description: `A list of AWS availability zones.`,
				},
				resource.Attribute{
					Name:        "ids",
					Description: `A list of AWS subnet IDs to create links to in the customer's account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vmc_org",
			Category:         "Data Sources",
			ShortDescription: `A organization data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Computed) ID of the organization.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vmc_sddc",
			Category:         "Data Sources",
			ShortDescription: `An sddc data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Organization identifier.`,
				},
				resource.Attribute{
					Name:        "sddc_id",
					Description: `(Required) ID of the SDDC. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The AWS specific (e.g us-west-2) or VMC specific region (e.g US_WEST_2) of the cloud resources to work in.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `Name of the SDDC.`,
				},
				resource.Attribute{
					Name:        "num_host",
					Description: `The number of hosts.`,
				},
				resource.Attribute{
					Name:        "sddc_type",
					Description: `Denotes the sddc type.`,
				},
				resource.Attribute{
					Name:        "sddc_state",
					Description: `Denotes the sddc state.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `Possible values are "ZEROCLOUD" and "AWS".`,
				},
				resource.Attribute{
					Name:        "skip_creating_vxlan",
					Description: `Boolean value to skip creating vxlan for compute gateway for SDDC provisioning.`,
				},
				resource.Attribute{
					Name:        "sso_domain",
					Description: `The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `Availability Zones.`,
				},
				resource.Attribute{
					Name:        "vc_url",
					Description: `VC URL.`,
				},
				resource.Attribute{
					Name:        "cloud_username",
					Description: `Cloud user name.`,
				},
				resource.Attribute{
					Name:        "nsxt_reverse_proxy_url",
					Description: `NSXT reverse proxy url for managing public IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The AWS specific (e.g us-west-2) or VMC specific region (e.g US_WEST_2) of the cloud resources to work in.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `Name of the SDDC.`,
				},
				resource.Attribute{
					Name:        "num_host",
					Description: `The number of hosts.`,
				},
				resource.Attribute{
					Name:        "sddc_type",
					Description: `Denotes the sddc type.`,
				},
				resource.Attribute{
					Name:        "sddc_state",
					Description: `Denotes the sddc state.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `Possible values are "ZEROCLOUD" and "AWS".`,
				},
				resource.Attribute{
					Name:        "skip_creating_vxlan",
					Description: `Boolean value to skip creating vxlan for compute gateway for SDDC provisioning.`,
				},
				resource.Attribute{
					Name:        "sso_domain",
					Description: `The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.`,
				},
				resource.Attribute{
					Name:        "availability_zones",
					Description: `Availability Zones.`,
				},
				resource.Attribute{
					Name:        "vc_url",
					Description: `VC URL.`,
				},
				resource.Attribute{
					Name:        "cloud_username",
					Description: `Cloud user name.`,
				},
				resource.Attribute{
					Name:        "nsxt_reverse_proxy_url",
					Description: `NSXT reverse proxy url for managing public IP.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"vmc_connected_accounts": 0,
		"vmc_customer_subnets":   1,
		"vmc_org":                2,
		"vmc_sddc":               3,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
