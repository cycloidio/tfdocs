package vmc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vmc_public_ip",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage public IPs.`,
			Description:      ``,
			Keywords: []string{
				"public",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nsxt_reverse_proxy_url",
					Description: `(Required) NSXT reverse proxy url for managing public IP. Computed after SDDC creation.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name for public IP. ## Attributes Reference In addition to arguments listed above, the following attributes are exported after public IP creation:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Public IP identifier.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for public IP.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Public IP identifier.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Public IP.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `Display name for public IP.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vmc_sddc",
			Category:         "Resources",
			ShortDescription: `Provides a resource to provision SDDC.`,
			Description:      ``,
			Keywords: []string{
				"sddc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_id",
					Description: `(Required) Organization identifier.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The AWS specific (e.g us-west-2) or VMC specific region (e.g US_WEST_2) of the cloud resources to work in.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `(Required) Name of the SDDC.`,
				},
				resource.Attribute{
					Name:        "storage_capacity",
					Description: `(Optional) The storage capacity value to be requested for the sddc primary cluster, in GiBs. If provided, instead of using the direct-attached storage, a capacity value amount of separable storage will be used.`,
				},
				resource.Attribute{
					Name:        "num_host",
					Description: `(Required) The number of hosts.`,
				},
				resource.Attribute{
					Name:        "account_link_sddc_config",
					Description: `(Optional) The account linking configuration object.`,
				},
				resource.Attribute{
					Name:        "host_instance_type",
					Description: `(Optional) The instance type for the esx hosts in the primary cluster of the SDDC.`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `(Optional) AWS VPC IP range. Only prefix of 16 or 20 is currently supported.`,
				},
				resource.Attribute{
					Name:        "sddc_type",
					Description: `(Optional) Denotes the sddc type , if the value is null or empty, the type is considered as default.`,
				},
				resource.Attribute{
					Name:        "vxlan_subnet",
					Description: `(Optional) VXLAN IP subnet in CIDR for compute gateway.`,
				},
				resource.Attribute{
					Name:        "delay_account_link",
					Description: `(Optional) Boolean flag identifying whether account linking should be delayed or not for the SDDC.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) Determines what additional properties are available based on cloud provider. Acceptable values are "ZEROCLOUD" and "AWS" with AWS as the default value.`,
				},
				resource.Attribute{
					Name:        "skip_creating_vxlan",
					Description: `(Optional) Boolean value to skip creating vxlan for compute gateway for SDDC provisioning.`,
				},
				resource.Attribute{
					Name:        "sso_domain",
					Description: `(Optional) The SSO domain name to use for vSphere users. If not specified, vmc.local will be used.`,
				},
				resource.Attribute{
					Name:        "sddc_template_id",
					Description: `(Optional) If provided, configuration from the template will applied to the provisioned SDDC.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `(Optional) Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default is SingleAZ.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Cluster identifier. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `SDDC identifier.`,
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
					Name:        "nsxt_reverse_proxy_url",
					Description: `NSXT reverse proxy url for managing public IP.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vmc_public_ip": 0,
		"vmc_sddc":      1,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
