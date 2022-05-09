package vmc

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vmc_cluster",
			Category:         "Resources",
			ShortDescription: `Provides a resource to manage clusters.`,
			Description:      ``,
			Keywords: []string{
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sddc_id",
					Description: `(Required) SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "num_hosts",
					Description: `(Required) Number of hosts in the cluster. The number of hosts must be between 3 - 16 hosts for a cluster.`,
				},
				resource.Attribute{
					Name:        "host_cpu_cores_count",
					Description: `(Optional) Customize CPU cores on hosts in a cluster. Specify number of cores to be enabled on hosts in a cluster.`,
				},
				resource.Attribute{
					Name:        "host_instance_type",
					Description: `(Optional) The instance type for the esx hosts added to this cluster. Possible values are: I3_METAL, I3EN_METAL and R5_METAL. Default value: I3_METAL.`,
				},
				resource.Attribute{
					Name:        "storage_capacity",
					Description: `(Optional) For EBS-backed instances i.e: for R5_METAL only, the requested storage capacity in GiB.`,
				},
				resource.Attribute{
					Name:        "microsoft_licensing_config",
					Description: `(Optional) Indicates the desired licensing support, if any, of Microsoft software. ## Attributes Reference In addition to arguments listed above, the following attributes are exported after cluster creation:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Cluster identifier.`,
				},
				resource.Attribute{
					Name:        "cluster_info",
					Description: `Information about cluster like name, state, host instance type and cluster identifier. ## Import Cluster resource can be imported using the ` + "`" + `id` + "`" + ` and ` + "`" + `sddc_id` + "`" + `, e.g. ` + "`" + `$ terraform import vmc_cluster.cluster_1 id,sddc_id` + "`" + ` - id = Cluster Identifier - sddc_id = SDDC Identifier ` + "`" + `$ terraform import vmc_cluster.cluster_1 afe7a0fd-3f0a-48b2-9ddb-0489c22732ae,45495963-d24d-469b-830a-9003bfe132b5` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Cluster identifier.`,
				},
				resource.Attribute{
					Name:        "cluster_info",
					Description: `Information about cluster like name, state, host instance type and cluster identifier. ## Import Cluster resource can be imported using the ` + "`" + `id` + "`" + ` and ` + "`" + `sddc_id` + "`" + `, e.g. ` + "`" + `$ terraform import vmc_cluster.cluster_1 id,sddc_id` + "`" + ` - id = Cluster Identifier - sddc_id = SDDC Identifier ` + "`" + `$ terraform import vmc_cluster.cluster_1 afe7a0fd-3f0a-48b2-9ddb-0489c22732ae,45495963-d24d-469b-830a-9003bfe132b5` + "`" + ``,
				},
			},
		},
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
					Description: `Display name for public IP. ## Import Public IP resource can be imported using the ` + "`" + `nsxt_reverse_proxy_url` + "`" + ` and ` + "`" + `id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_public_ip.public_ip_1 nsxt_reverse_proxy_url,id` + "`" + ` - nsxt_reverse_proxy_url = NSX API public endpoint url used for public IP resource management - id = Public IP Identifier ` + "`" + `$ terraform import vmc_public_ip.public_ip_1 'https://nsx-44-228-76-55.rp.vmwarevmc.com/vmc/reverse-proxy/api/orgs/{orgI}/sddcs/afe7a0fd-3f0a-48b2-9ddb-0489c22732ae/sks-nsxt-manager,8d730ad4-aa6b-4f9f-9679-ec17beeaceaf' ` + "`" + ``,
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
					Description: `Display name for public IP. ## Import Public IP resource can be imported using the ` + "`" + `nsxt_reverse_proxy_url` + "`" + ` and ` + "`" + `id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_public_ip.public_ip_1 nsxt_reverse_proxy_url,id` + "`" + ` - nsxt_reverse_proxy_url = NSX API public endpoint url used for public IP resource management - id = Public IP Identifier ` + "`" + `$ terraform import vmc_public_ip.public_ip_1 'https://nsx-44-228-76-55.rp.vmwarevmc.com/vmc/reverse-proxy/api/orgs/{orgI}/sddcs/afe7a0fd-3f0a-48b2-9ddb-0489c22732ae/sks-nsxt-manager,8d730ad4-aa6b-4f9f-9679-ec17beeaceaf' ` + "`" + ``,
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
					Description: `(Optional) The storage capacity value to be requested for the SDDC primary cluster. This variable is only for R5_METAL. Possible values are 15TB, 20TB, 25TB, 30TB, 35TB per host.`,
				},
				resource.Attribute{
					Name:        "num_host",
					Description: `(Required) The number of hosts.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the vCenter and NSX appliances. 'large' or 'LARGE' SDDC size corresponds to a large vCenter appliance and large NSX appliance. 'medium' or 'MEDIUM' SDDC size corresponds to medium vCenter appliance and medium NSX appliance. Default : 'medium'.`,
				},
				resource.Attribute{
					Name:        "account_link_sddc_config",
					Description: `(Optional) The account linking configuration object.`,
				},
				resource.Attribute{
					Name:        "host_instance_type",
					Description: `(Optional) The instance type for the esx hosts in the primary cluster of the SDDC. Possible values : I3_METAL, I3EN_METAL and R5_METAL. Default value : I3_METAL. Currently I3EN_METAL host_instance_type does not support 1NODE and 2 node SDDC deployment.`,
				},
				resource.Attribute{
					Name:        "vpc_cidr",
					Description: `(Optional) SDDC management network CIDR. Only prefix of 16, 20 and 23 are supported. Note : Specify a private subnet range (RFC 1918) to be used for vCenter Server, NSX Manager, and ESXi hosts. Choose a range that will not conflict with other networks you will connect to this SDDC. Minimum CIDR sizes : /23 for up to 27 hosts, /20 for up to 251 hosts, /16 for up to 4091 hosts. Reserved CIDRs : 10.0.0.0/15, 172.31.0.0/16.`,
				},
				resource.Attribute{
					Name:        "sddc_type",
					Description: `(Optional) Denotes the sddc type , if the value is null or empty, the type is considered as default.`,
				},
				resource.Attribute{
					Name:        "vxlan_subnet",
					Description: `(Optional) A logical network segment that will be created with the SDDC under the compute gateway.`,
				},
				resource.Attribute{
					Name:        "delay_account_link",
					Description: `(Optional) Boolean flag identifying whether account linking should be delayed or not for the SDDC.`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `(Optional) Determines what additional properties are available based on cloud provider. Default value : AWS`,
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
					Description: `(Optional) Denotes if request is for a SingleAZ or a MultiAZ SDDC. Default : SingleAZ.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) Cluster identifier.`,
				},
				resource.Attribute{
					Name:        "microsoft_licensing_config",
					Description: `(Optional) Indicates the desired licensing support, if any, of Microsoft software. ## Attributes Reference In addition to arguments listed above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "cluster_info",
					Description: `Information about cluster like id, name, state, host instance type.`,
				},
				resource.Attribute{
					Name:        "sddc_size",
					Description: `Size information of vCenter appliance and NSX appliance.`,
				},
				resource.Attribute{
					Name:        "intranet_uplink_mtu",
					Description: `Uplink MTU of direct connect, sddc-grouping and outposts traffic in edge tier-0 router port. This field can be updated only after an SDDC is created. Range : 1500 - 8900. Default : 1500.`,
				},
				resource.Attribute{
					Name:        "nsxt_reverse_proxy_url",
					Description: `NSXT reverse proxy url for managing public IP.`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudadmin",
					Description: `the NSXT userID admin for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudadmin_password",
					Description: `the NSXT userID admin password for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudaudit",
					Description: `the NSXT userID audit for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudaudit_password",
					Description: `the NSXT userID audit password for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_private_url",
					Description: `for example "https://nsxManager.sddc-54-213-170-7.vmwarevmc.com/login.jsp"`,
				},
				resource.Attribute{
					Name:        "nsxt_public_url",
					Description: `same as reverse proxy ## Import SDDC resource can be imported using the ` + "`" + `id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_sddc.sddc_1 id` + "`" + ` For this example: - id = SDDC Identifier ` + "`" + `$ terraform import vmc_sddc.sddc_1 afe7a0fd-3f0a-48b2-9ddb-0489c22732ae` + "`" + ` Note : Running plan/apply after importing an SDDC causes the SDDC to be re-created. This is due to a limitation in the current Get and Update SDDC APIs. Hence, the import functionality is only partially supported.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "cluster_info",
					Description: `Information about cluster like id, name, state, host instance type.`,
				},
				resource.Attribute{
					Name:        "sddc_size",
					Description: `Size information of vCenter appliance and NSX appliance.`,
				},
				resource.Attribute{
					Name:        "intranet_uplink_mtu",
					Description: `Uplink MTU of direct connect, sddc-grouping and outposts traffic in edge tier-0 router port. This field can be updated only after an SDDC is created. Range : 1500 - 8900. Default : 1500.`,
				},
				resource.Attribute{
					Name:        "nsxt_reverse_proxy_url",
					Description: `NSXT reverse proxy url for managing public IP.`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudadmin",
					Description: `the NSXT userID admin for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudadmin_password",
					Description: `the NSXT userID admin password for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudaudit",
					Description: `the NSXT userID audit for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_cloudaudit_password",
					Description: `the NSXT userID audit password for direct NSXT access`,
				},
				resource.Attribute{
					Name:        "nsxt_private_url",
					Description: `for example "https://nsxManager.sddc-54-213-170-7.vmwarevmc.com/login.jsp"`,
				},
				resource.Attribute{
					Name:        "nsxt_public_url",
					Description: `same as reverse proxy ## Import SDDC resource can be imported using the ` + "`" + `id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_sddc.sddc_1 id` + "`" + ` For this example: - id = SDDC Identifier ` + "`" + `$ terraform import vmc_sddc.sddc_1 afe7a0fd-3f0a-48b2-9ddb-0489c22732ae` + "`" + ` Note : Running plan/apply after importing an SDDC causes the SDDC to be re-created. This is due to a limitation in the current Get and Update SDDC APIs. Hence, the import functionality is only partially supported.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vmc_site_recovery",
			Category:         "Resources",
			ShortDescription: `Provides a resource to activate and deactivate site recovery for SDDC.`,
			Description:      ``,
			Keywords: []string{
				"site",
				"recovery",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sddc_id",
					Description: `(Required) SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "srm_node_extension_key_suffix",
					Description: `(Optional) Custom extension key suffix for SRM. If not specified, default extension key will be used. The custom extension suffix must contain 13 characters or less, be composed of letters, numbers, ., - characters. The extension suffix must begin and end with a letter or number. The suffix is appended to com.vmware.vcDr- to form the full extension key. ## Attributes Reference In addition to arguments listed above, the following attributes are exported after site recovery activation:`,
				},
				resource.Attribute{
					Name:        "site_recovery_state",
					Description: `Site recovery state. Possible values are: ACTIVATED, ACTIVATING, CANCELED, DEACTIVATED, DEACTIVATING, DELETED, FAILED.`,
				},
				resource.Attribute{
					Name:        "srm_node",
					Description: `Site recovery node created after site recovery activation.`,
				},
				resource.Attribute{
					Name:        "vr_node",
					Description: `VR node created after site recovery activation. ## Import Site recovery resource can be imported using the ` + "`" + `sddc_id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_site_recovery.site_recovery_1 sddc_id` + "`" + ` - sddc_id = SDDC Identifier ` + "`" + `$ terraform import vmc_site_recovery.site_recovery_1 afe7a0fd-3f0a-48b2-9ddb-0489c22732ae` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "site_recovery_state",
					Description: `Site recovery state. Possible values are: ACTIVATED, ACTIVATING, CANCELED, DEACTIVATED, DEACTIVATING, DELETED, FAILED.`,
				},
				resource.Attribute{
					Name:        "srm_node",
					Description: `Site recovery node created after site recovery activation.`,
				},
				resource.Attribute{
					Name:        "vr_node",
					Description: `VR node created after site recovery activation. ## Import Site recovery resource can be imported using the ` + "`" + `sddc_id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_site_recovery.site_recovery_1 sddc_id` + "`" + ` - sddc_id = SDDC Identifier ` + "`" + `$ terraform import vmc_site_recovery.site_recovery_1 afe7a0fd-3f0a-48b2-9ddb-0489c22732ae` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vmc_srm_node",
			Category:         "Resources",
			ShortDescription: `Provides a resource to add an instance to SDDC after site recovery has been activated.`,
			Description:      ``,
			Keywords: []string{
				"srm",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sddc_id",
					Description: `(Required) SDDC identifier.`,
				},
				resource.Attribute{
					Name:        "srm_node_extension_key_suffix",
					Description: `(Required) Custom extension key suffix for SRM. If not specified, default extension key will be used. The custom extension suffix must contain 13 characters or less, be composed of letters, numbers, ., - characters. The extension suffix must begin and end with a letter or number. The suffix is appended to com.vmware.vcDr- to form the full extension key. ## Attributes Reference In addition to arguments listed above, the following attributes are exported after site recovery activation:`,
				},
				resource.Attribute{
					Name:        "site_recovery_state",
					Description: `Site recovery state. Possible values are: ACTIVATED, ACTIVATING, CANCELED, DEACTIVATED, DEACTIVATING, DELETED, FAILED.`,
				},
				resource.Attribute{
					Name:        "srm_node",
					Description: `Site recovery node information.`,
				},
				resource.Attribute{
					Name:        "vr_node",
					Description: `VR node information. ## Import SRM node resource can be imported using the ` + "`" + `id` + "`" + ` and ` + "`" + `sddc_id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_srm_node.srm_node_1 id,sddc_id` + "`" + ` - id = SRM Node Identifier - sddc_id = SDDC Identifier ` + "`" + `$ terraform import vmc_srm_node.srm_node_1 7aad97e9-9a4f-4e43-8817-5c8d8c0e87a5,afe7a0fd-3f0a-48b2-9ddb-0489c22732ae` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "site_recovery_state",
					Description: `Site recovery state. Possible values are: ACTIVATED, ACTIVATING, CANCELED, DEACTIVATED, DEACTIVATING, DELETED, FAILED.`,
				},
				resource.Attribute{
					Name:        "srm_node",
					Description: `Site recovery node information.`,
				},
				resource.Attribute{
					Name:        "vr_node",
					Description: `VR node information. ## Import SRM node resource can be imported using the ` + "`" + `id` + "`" + ` and ` + "`" + `sddc_id` + "`" + ` , e.g. ` + "`" + `$ terraform import vmc_srm_node.srm_node_1 id,sddc_id` + "`" + ` - id = SRM Node Identifier - sddc_id = SDDC Identifier ` + "`" + `$ terraform import vmc_srm_node.srm_node_1 7aad97e9-9a4f-4e43-8817-5c8d8c0e87a5,afe7a0fd-3f0a-48b2-9ddb-0489c22732ae` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vmc_cluster":       0,
		"vmc_public_ip":     1,
		"vmc_sddc":          2,
		"vmc_site_recovery": 3,
		"vmc_srm_node":      4,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
