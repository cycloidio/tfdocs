package vra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_aws",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_aws resource.`,
			Description: `\_cloud\_account\_aws

Creates a VMware vRealize Automation AWS cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `(Required) Access key ID for AWS.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of AWS cloud account.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Set of region names enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) AWS Secret Access Key`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and values to apply to the cloud account. Example:[ { "key" : "vmware", "value": "provider" } ] ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of AWS cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the AWS cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_aws.new_aws 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of AWS cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the AWS cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_aws.new_aws 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_azure",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_azure resource.`,
			Description: `\_cloud\_account\_azure

Creates a VMware vRealize Automation Azure cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) Azure Client Application ID.`,
				},
				resource.Attribute{
					Name:        "application_key",
					Description: `(Required) Azure Client Application Secret Key.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of Azure cloud account.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Set of region names enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) Azure Subscription ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and values to apply to the cloud account. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Azure Tenant ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the Azure cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_azure.new_azure 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the Azure cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_azure.new_azure 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_gcp",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_gcp resource.`,
			Description: `\_cloud\_account\_gcp

Creates a VMware vRealize Automation GCP cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) GCP Client email.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of GCP cloud account.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Set of region names enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Required) GCP Private key.`,
				},
				resource.Attribute{
					Name:        "private_key_id",
					Description: `(Required) GCP Private key ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) GCP Project ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and values to apply to the cloud account. Example:[ { "key" : "vmware", "value": "provider" } ] ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of GCP cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the GCP cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_gcp.new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of GCP cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the GCP cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_gcp.new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_nsxt",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_nsxt resource.`,
			Description: `\_cloud\_account\_nsxt

Creates a VMware vRealize Automation NSX-T cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self-signed certificate when connecting to the cloud account.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) Identifier of a data collector VM deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host name for NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password used to authenticate to the cloud Account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and values to apply to the cloud account. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username used to authenticate to the cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with the cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the NSX-T cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_nsxt.new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with the cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the NSX-T cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_nsxt.new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_nsxv",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_nsxv resource.`,
			Description: `\_cloud\_account\_nsxv

Creates a VMware vRealize Automation NSX-V cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self-signed certificate when connecting to the cloud account.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) Identifier of a data collector VM deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host name for NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password used to authenticate to the cloud account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and values to apply to the cloud account. Example: [ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username used to authenticate with the cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated to the cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the NSX-V cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_nsxv.new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated to the cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the NSX-V cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_nsxv.new_gcp 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_vmc",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_vmc resource.`,
			Description: `\_cloud\_account\_vmc

Creates a VMware vRealize Automation VMC cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self-signed certificate when connecting to the cloud account.`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Required) VMC API access key.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) Identifier of a data collector VM deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "nsx_hostname",
					Description: `(Required) IP address of the NSX Manager server in the specified SDDC / FQDN.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) Set of region names enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `(Required) Identifier of the on-premise SDDC to be used by the cloud account. Note that NSX-V SDDCs are not supported.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and values to apply to the cloud account. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "vcenter_hostname",
					Description: `(Required) IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.`,
				},
				resource.Attribute{
					Name:        "vcenter_password",
					Description: `(Required) Password used to authenticate to the cloud Account.`,
				},
				resource.Attribute{
					Name:        "vcenter_username",
					Description: `(Required) vCenter user name for the specified SDDC. The user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VMC cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "region-ids",
					Description: `Set of region IDs enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the VMC cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_vmc.new_vmc 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VMC cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "region-ids",
					Description: `Set of region IDs enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the VMC cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_vmc.new_vmc 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_vsphere",
			Category:         "Resources",
			ShortDescription: `Creates a vra_cloud_account_vsphere resource.`,
			Description: `\_cloud\_account\_vsphere

Creates a VMware vRealize Automation vSphere cloud account resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self-signed certificate when connecting to the cloud account.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) Identifier of a data collector VM deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the vSphere cloud account.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password used to authenticate to the cloud account.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `(Optional) A set of region names that are enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tag keys and optional values to apply to the cloud account. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) vSphere username used to authenticate to the cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with the cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `A list of key value pair of properties associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the vSphere cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "region-ids",
					Description: `Set of region IDs that enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the vSphere cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_vsphere.new_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with the cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `A list of key value pair of properties associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the vSphere cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "region-ids",
					Description: `Set of region IDs that enabled for the cloud account.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC. ## Import To import the vSphere cloud account, use the ID as in the following example: ` + "`" + `$ terraform import vra_cloud_account_vsphere.new_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_fabric_network_vsphere",
			Category:         "Resources",
			ShortDescription: `Updates a fabric_network_vsphere resource.`,
			Description: `

Updates a VMware vRealize Automation fabric_network_vsphere resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `Network CIDR to be used.`,
				},
				resource.Attribute{
					Name:        "default_gateway",
					Description: `IPv4 default gateway to be used.`,
				},
				resource.Attribute{
					Name:        "default_ipv6_gateway",
					Description: `IPv6 default gateway to be used.`,
				},
				resource.Attribute{
					Name:        "dns_search_domains",
					Description: `List of dns search domains for the vSphere network.`,
				},
				resource.Attribute{
					Name:        "dns_server_addresses",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `Domain for the vSphere network.`,
				},
				resource.Attribute{
					Name:        "ipv6_cidr",
					Description: `Network IPv6 CIDR to be used.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether this is the default subnet for the zone.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicates whether the sub-network supports public IP assignment. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_ids",
					Description: `Set of ids of the cloud accounts this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this network is defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the vRA fabric network.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Set of tag keys and values to apply to the resource. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC. ## Import To import the vSphere fabric network resource, use the ID as in the following example: ` + "`" + `$ terraform import fabric_network_vsphere.new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_ids",
					Description: `Set of ids of the cloud accounts this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this network is defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the vRA fabric network.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Set of tag keys and values to apply to the resource. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC. ## Import To import the vSphere fabric network resource, use the ID as in the following example: ` + "`" + `$ terraform import fabric_network_vsphere.new_fabric_network_vsphere 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_network_ip_range",
			Category:         "Resources",
			ShortDescription: `Creates a network_ip_range resource.`,
			Description: `

Creates a VMware vRealize Automation network_ip_range resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `State object representing a network on a external cloud provider.`,
				},
				resource.Attribute{
					Name:        "end_ip_address",
					Description: `End IP address of the range.`,
				},
				resource.Attribute{
					Name:        "fabric_network_id",
					Description: `Fabric network Id.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the network IP range.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `IP address version: IPv4 or IPv6. Default: IPv4.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "start_ip_address",
					Description: `Start IP address of the range. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Set of tag keys and values to apply to the resource. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC. ## Import To import the vRA Network IP range, use the ID as in the following example: ` + "`" + `$ terraform import network_ip_range.new_ip_range 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Set of tag keys and values to apply to the resource. Example:[ { "key" : "vmware", "value": "provider" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC. ## Import To import the vRA Network IP range, use the ID as in the following example: ` + "`" + `$ terraform import network_ip_range.new_ip_range 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_block_device",
			Category:         "Resources",
			ShortDescription: `Creates a vra_block_device resource.`,
			Description: `

Creates a VMware vRealize Automation block device resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "capacity_in_gb",
					Description: `(Required) Capacity of block device in GB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of project that current user belongs to.`,
				},
				resource.Attribute{
					Name:        "constraints",
					Description: `(Optional) Storage, network, and extensibility constraints to be applied when provisioning through the project.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Describes machine within the scope of your organization and is not propagated to the cloud.`,
				},
				resource.Attribute{
					Name:        "disk_content_base_64",
					Description: `(Optional) Content of a disk, base64 encoded.`,
				},
				resource.Attribute{
					Name:        "encrypted",
					Description: `(Optional) Indicates whether block device should be encrypted or not.`,
				},
				resource.Attribute{
					Name:        "expand_snapshots",
					Description: `(Optional) Indicates whether snapshots of block devices should be included in the state. Applies only to first class block devices.`,
				},
				resource.Attribute{
					Name:        "purge",
					Description: `(Optional) Indicates if the disk must be completely destroyed or should be kept in the system. Valid only for block devices with ‘persistent’ set to true. Used to destroy the resource.`,
				},
				resource.Attribute{
					Name:        "persistent",
					Description: `(Optional) Indicates whether block device survives a delete action.`,
				},
				resource.Attribute{
					Name:        "source_reference",
					Description: `(Optional) URI to use for block device. Example: ami-0d4cfd66 ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional custom properties that may be used to extend the machine.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `ID of deployment associated with resource.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity ID on provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `External regionId of resource.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `External zoneId of resource.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `Represents a machine snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describes machine within the scope of your organization and is not propagated to the cloud.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the block device snapshot.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Indicates whether snapshot on block device is current.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human-friendly name for block device snapshot.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that block device snapshot belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of block device snapshot owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of block device.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Set of tag keys and values to apply to the resource instance. Example:[ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional custom properties that may be used to extend the machine.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `ID of deployment associated with resource.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity ID on provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `External regionId of resource.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `External zoneId of resource.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "snapshots",
					Description: `Represents a machine snapshot.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describes machine within the scope of your organization and is not propagated to the cloud.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the block device snapshot.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Indicates whether snapshot on block device is current.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human-friendly name for block device snapshot.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that block device snapshot belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of block device snapshot owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of block device.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `Set of tag keys and values to apply to the resource instance. Example:[ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_block_device_snapshot",
			Category:         "Resources",
			ShortDescription: `Creates a VMware vRealize Automation vra_block_device_snapshot resource.`,
			Description: `

Creates a VMware vRealize Automation block device snapshot resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "block_device_id",
					Description: `(Required) ID of block device.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Indicates whether snapshot on block device is current.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Indicates whether snapshot on block device is current.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_blueprint",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to create a vRealize Automation cloud template, formerly know as blueprint.`,
			Description: `\_blueprint

Creates a VMware vRealize Automation (vRA) cloud template resource, formerly known as a blueprint.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Blueprint YAML content.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "request_scope_org",
					Description: `(Optional) Flag to indicate whether blueprint can be requested from any project in the organization that entity belongs to. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "content_source_id",
					Description: `ID of content source.`,
				},
				resource.Attribute{
					Name:        "content_source_path",
					Description: `Content source path.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_at",
					Description: `Date when content source was last synced. The date is in ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_messages",
					Description: `Content source last sync messages.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_status",
					Description: `Content source last sync status. Supported values: ` + "`" + `SUCCESSFUL` + "`" + `, ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_source_type",
					Description: `Content source type.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. The date is in ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created entity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of cloud template.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `Name of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of cloud template. Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "total_released_versions",
					Description: `Total number of released versions.`,
				},
				resource.Attribute{
					Name:        "total_versions",
					Description: `Total number of versions.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who last updated the entity.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template/blueprint is valid.`,
				},
				resource.Attribute{
					Name:        "validation_messages",
					Description: `List of validations messages.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content_source_id",
					Description: `ID of content source.`,
				},
				resource.Attribute{
					Name:        "content_source_path",
					Description: `Content source path.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_at",
					Description: `Date when content source was last synced. The date is in ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_messages",
					Description: `Content source last sync messages.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_status",
					Description: `Content source last sync status. Supported values: ` + "`" + `SUCCESSFUL` + "`" + `, ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "content_source_type",
					Description: `Content source type.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. The date is in ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user who created entity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of cloud template.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `Name of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of cloud template. Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "total_released_versions",
					Description: `Total number of released versions.`,
				},
				resource.Attribute{
					Name:        "total_versions",
					Description: `Total number of versions.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user who last updated the entity.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template/blueprint is valid.`,
				},
				resource.Attribute{
					Name:        "validation_messages",
					Description: `List of validations messages.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_blueprint_version",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to create a vRealize Automation cloud template version.`,
			Description: `\_blueprint\_version

Creates a VMware vRealize Automation cloud template (blueprint) version resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_id",
					Description: `(Required) ID of the cloud template (blueprint).`,
				},
				resource.Attribute{
					Name:        "change_log",
					Description: `(Optional) Cloud template (blueprint) version log.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description for the cloud template (blueprint) version.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Optional) Flag to indicate whether to release the version.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Cloud template (blueprint) version. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "blueprint_description",
					Description: `Description of cloud template (blueprint).`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Blueprint YAML content.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created the entity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of cloud template (blueprint) version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of cloud template (blueprint) version.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `Name of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cloud template (blueprint). Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `User who last updated the entity.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template (blueprint) is valid. ## Import To import the cloud template (blueprint) version, use the ID as in the following example: ` + "`" + `$ terraform import vra_blueprint_version.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_description",
					Description: `Description of cloud template (blueprint).`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Blueprint YAML content.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created the entity.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of cloud template (blueprint) version.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of cloud template (blueprint) version.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `Name of project that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cloud template (blueprint). Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `User who last updated the entity.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template (blueprint) is valid. ## Import To import the cloud template (blueprint) version, use the ID as in the following example: ` + "`" + `$ terraform import vra_blueprint_version.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_catalog_source_blueprint",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to create a vRealize Automation catalog source of type cloud template (blueprint).`,
			Description: `\_catalog\_source\_blueprint

Creates a VMware vRealize Automation catalog source resource of type cloud template, formerly known as a blueprint.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config",
					Description: `(Optional) Custom configuration of the catalog source as a map of key values.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) ID of the project this entity belongs to. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created the entity.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `Flag indicating that all items can be requested across all projects.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of catalog source.`,
				},
				resource.Attribute{
					Name:        "items_found",
					Description: `Number of items found in the catalog source.`,
				},
				resource.Attribute{
					Name:        "items_imported",
					Description: `Number of items imported from the catalog source.`,
				},
				resource.Attribute{
					Name:        "last_import_completed_at",
					Description: `Time at which the last import completed.`,
				},
				resource.Attribute{
					Name:        "last_import_errors",
					Description: `List of errors seen when the catalog source was last imported.`,
				},
				resource.Attribute{
					Name:        "last_import_started_at",
					Description: `Time at which the last import started.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who last updated the catalog source.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Type of catalog source. Example: ` + "`" + `blueprint` + "`" + `, ` + "`" + `CFT` + "`" + `, etc. ## Import To import the cloud template catalog source, use the ID as in the following example: ` + "`" + `$ terraform import vra_catalog_source_blueprint.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `User who created the entity.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `Flag indicating that all items can be requested across all projects.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of catalog source.`,
				},
				resource.Attribute{
					Name:        "items_found",
					Description: `Number of items found in the catalog source.`,
				},
				resource.Attribute{
					Name:        "items_imported",
					Description: `Number of items imported from the catalog source.`,
				},
				resource.Attribute{
					Name:        "last_import_completed_at",
					Description: `Time at which the last import completed.`,
				},
				resource.Attribute{
					Name:        "last_import_errors",
					Description: `List of errors seen when the catalog source was last imported.`,
				},
				resource.Attribute{
					Name:        "last_import_started_at",
					Description: `Time at which the last import started.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `User who last updated the catalog source.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Type of catalog source. Example: ` + "`" + `blueprint` + "`" + `, ` + "`" + `CFT` + "`" + `, etc. ## Import To import the cloud template catalog source, use the ID as in the following example: ` + "`" + `$ terraform import vra_catalog_source_blueprint.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_catalog_source_entitlement",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to create a vRealize Automation catalog source entitlement.`,
			Description: `\_catalog\_source\_entitlement

This resource provides a way to create a vRealize Automation(vRA) catalog source entitlement.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_source_id",
					Description: `(Required) The id of the catalog source to create the entitlement.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project this entity belongs to. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "definition",
					Description: `Represents a catalog item or content source that is linked to a project via an entitlement.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of either the catalog item, or the catalog source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of either the catalog source, or the catalog item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of either the catalog item, or the catalog source.`,
				},
				resource.Attribute{
					Name:        "number_of_items",
					Description: `Number of items in the associated catalog source.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Type of the catalog source.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Content definition type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of this catalog source entitlement. ## Import Catalog source entitlement can be imported using the id, e.g. ` + "`" + `$ terraform import vra_catalog_source_entitlement.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "definition",
					Description: `Represents a catalog item or content source that is linked to a project via an entitlement.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of either the catalog item, or the catalog source.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of either the catalog source, or the catalog item.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of either the catalog item, or the catalog source.`,
				},
				resource.Attribute{
					Name:        "number_of_items",
					Description: `Number of items in the associated catalog source.`,
				},
				resource.Attribute{
					Name:        "source_type",
					Description: `Type of the catalog source.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Content definition type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of this catalog source entitlement. ## Import Catalog source entitlement can be imported using the id, e.g. ` + "`" + `$ terraform import vra_catalog_source_entitlement.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_content_source",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to create a content source in vRealize Automation(vRA).`,
			Description: `\_content_source

This resource provides a way to create a content source vRealize Automation(vRA).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "config",
					Description: `(Required) Content source custom configuration.`,
				},
				resource.Attribute{
					Name:        "branch",
					Description: `Content source branch name.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `Content source type. Supported values are ` + "`" + `BLUEPRINT` + "`" + `, ` + "`" + `IMAGE` + "`" + `, ` + "`" + `ABX_SCRIPTS` + "`" + `, ` + "`" + `TERRAFORM_CONFIGURATION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "integration_id",
					Description: `Content source integration id as seen in vRA integrations.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `Path to refer to in the content source repository and branch.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `Name of the project.`,
				},
				resource.Attribute{
					Name:        "repository",
					Description: `Content source repository.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name for content source used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "sync_enabled",
					Description: `(Required) Flag indicating whether sync is enabled for this content source.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `(Required) Content Source type. Supported values are ` + "`" + `com.gitlab` + "`" + `, ` + "`" + `com.github` + "`" + `, ` + "`" + `com.vmware.marketplace` + "`" + `, ` + "`" + `org.bitbucket` + "`" + `. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of this cloud template.`,
				},
				resource.Attribute{
					Name:        "last_updated_at",
					Description: `Date when the entity was last updated. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to. ## Import Content source can be imported using the id, e.g. ` + "`" + `$ terraform import vra_content_source.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of this cloud template.`,
				},
				resource.Attribute{
					Name:        "last_updated_at",
					Description: `Date when the entity was last updated. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to. ## Import Content source can be imported using the id, e.g. ` + "`" + `$ terraform import vra_content_source.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_deployment",
			Category:         "Resources",
			ShortDescription: `A resource that can be used to create a vRealize Automation deployment.`,
			Description: `\_deployment

This resource provides a way to create a deployment in vRealize Automation(vRA) by either using a blueprint, or catalog item, or an inline blueprint.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_d",
					Description: `(Optional) The id of the vRA cloud template to request the deployment. Conflicts with ` + "`" + `catalog_item_id` + "`" + ` and ` + "`" + `blueprint_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "blueprint_version",
					Description: `(Optional) The version of the vRA cloud template to request the deployment. Used only when ` + "`" + `blueprint_id` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "blueprint_content",
					Description: `(Optional) vRA Cloud template content. Conflicts with ` + "`" + `blueprint_id` + "`" + ` and ` + "`" + `catalog_item_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `(Optional) The id of the vRA catalog item to request the deployment. Conflicts with ` + "`" + `blueprint_id` + "`" + ` and ` + "`" + `blueprint_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog_item_version",
					Description: `(Optional) The version of the vRA catalog item to request the deployment. Used only when ` + "`" + `catalog_item_id` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "expand_project",
					Description: `(Optional) Flag to indicate whether to expand project information.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `(Optional) Inputs provided by the user. For inputs including those with default values, refer to ` + "`" + `inputs_including_defaults` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `(Optional) The ID of the organization this deployment belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The user this deployment belongs to. At create, the owner is ignored but is used to update during next apply.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project this entity belongs to. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "expense",
					Description: `Expense incurred for the deployment.`,
				},
				resource.Attribute{
					Name:        "additional_expense",
					Description: `Additional expense incurred for the resource.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Expense sync message code if any.`,
				},
				resource.Attribute{
					Name:        "compute_expense",
					Description: `Compute expense of the entity.`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last expense sync time.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Expense sync message if any.`,
				},
				resource.Attribute{
					Name:        "network_expense",
					Description: `Network expense of the entity.`,
				},
				resource.Attribute{
					Name:        "storage_expense",
					Description: `Storage expense of the entity.`,
				},
				resource.Attribute{
					Name:        "total_expense",
					Description: `Total expense of the entity.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `Monetary unit.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of this entity.`,
				},
				resource.Attribute{
					Name:        "inputs_including_defaults",
					Description: `All the inputs applied during last create/update operation, including those with default values. For the list of inputs provided by the user in the configuration, refer to ` + "`" + `inputs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_request",
					Description: `Represents deployment requests.`,
				},
				resource.Attribute{
					Name:        "action_id",
					Description: `Identifier of the requested action.`,
				},
				resource.Attribute{
					Name:        "approved_at",
					Description: `Time at which the request was approved.`,
				},
				resource.Attribute{
					Name:        "blueprint_id",
					Description: `Identifier of the requested blueprint in the form ‘UUID:version’.`,
				},
				resource.Attribute{
					Name:        "cancelable",
					Description: `Indicates whether request can be canceled or not.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `Identifier of the requested catalog item in the form ‘UUID:version’.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `Time at which the request completed.`,
				},
				resource.Attribute{
					Name:        "completed_tasks",
					Description: `The number of tasks completed while fulfilling this request.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Longer user-friendly details of the request.`,
				},
				resource.Attribute{
					Name:        "dismissed",
					Description: `Indicates whether request is in dismissed state.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Request identifier.`,
				},
				resource.Attribute{
					Name:        "initialized_at",
					Description: `Time at which the request was initialized.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `List of request inputs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Short user-friendly label of the request (e.g. ‘shuting down myVM’).`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `Request outputs.`,
				},
				resource.Attribute{
					Name:        "requested_by",
					Description: `The user that initiated the request.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `Optional resource name to which the request applies to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Request overall execution status. Supported values: ` + "`" + `CREATED` + "`" + `, ` + "`" + `PENDING` + "`" + `, ` + "`" + `INITIALIZATION` + "`" + `, ` + "`" + `CHECKING_APPROVAL` + "`" + `, ` + "`" + `APPROVAL_PENDING` + "`" + `, ` + "`" + `INPROGRESS` + "`" + `, ` + "`" + `COMPLETION` + "`" + `, ` + "`" + `APPROVAL_REJECTED` + "`" + `, ` + "`" + `ABORTED` + "`" + `, ` + "`" + `SUCCESSFUL` + "`" + `, ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update time (e.g. date format ‘2019-07-13T23:16:49.310Z’).`,
				},
				resource.Attribute{
					Name:        "last_updated_at",
					Description: `Time at which the deployment was last updated.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user that last updated the deployment.`,
				},
				resource.Attribute{
					Name:        "lease_expire_at",
					Description: `Time at which the deployment lease expires.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human friendly description.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the entity.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the entity.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the entity, if applicable.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Expanded resources for the deployment. Content of this property will not be maintained backward compatible.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).`,
				},
				resource.Attribute{
					Name:        "depends_on",
					Description: `A list of other resources this resource depends on.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the resource. ` + "`" + `expense` + "`" + ` - Expense incurred for this resource.`,
				},
				resource.Attribute{
					Name:        "additional_expense",
					Description: `Additional expense incurred for the resource.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Expense sync message code if any.`,
				},
				resource.Attribute{
					Name:        "compute_expense",
					Description: `Compute expense of the entity.`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last expense sync time.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Expense sync message if any.`,
				},
				resource.Attribute{
					Name:        "network_expense",
					Description: `Network expense of the entity.`,
				},
				resource.Attribute{
					Name:        "storage_expense",
					Description: `Storage expense of the entity.`,
				},
				resource.Attribute{
					Name:        "total_expense",
					Description: `Total expense of the entity.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `Monetary unit.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the resource.`,
				},
				resource.Attribute{
					Name:        "properties_json",
					Description: `List of properties in the encoded JSON string format.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the resource. Supported values are ` + "`" + `PARTIAL` + "`" + `, ` + "`" + `TAINTED` + "`" + `, ` + "`" + `OK.` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sync_status",
					Description: `The current sync status. Supported values are ` + "`" + `SUCCESS` + "`" + `, ` + "`" + `MISSING` + "`" + `, ` + "`" + `STALE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Deployment status. Supported values are: ` + "`" + `CREATE_SUCCESSFUL` + "`" + `, ` + "`" + `CREATE_INPROGRESS` + "`" + `, ` + "`" + `CREATE_FAILED` + "`" + `, ` + "`" + `UPDATE_SUCCESSFUL` + "`" + `, ` + "`" + `UPDATE_INPROGRESS` + "`" + `, ` + "`" + `UPDATE_FAILED` + "`" + `, ` + "`" + `DELETE_SUCCESSFUL` + "`" + `, ` + "`" + `DELETE_INPROGRESS` + "`" + `, ` + "`" + `DELETE_FAILED` + "`" + `, ` + "`" + `ACTION_SUCCESSFUL` + "`" + `, ` + "`" + `ACTION_INPROGRESS` + "`" + `, ` + "`" + `ACTION_FAILED` + "`" + `. ## Import Deployment can be imported using the id, e.g. ` + "`" + `$ terraform import vra_deployment.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "expense",
					Description: `Expense incurred for the deployment.`,
				},
				resource.Attribute{
					Name:        "additional_expense",
					Description: `Additional expense incurred for the resource.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Expense sync message code if any.`,
				},
				resource.Attribute{
					Name:        "compute_expense",
					Description: `Compute expense of the entity.`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last expense sync time.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Expense sync message if any.`,
				},
				resource.Attribute{
					Name:        "network_expense",
					Description: `Network expense of the entity.`,
				},
				resource.Attribute{
					Name:        "storage_expense",
					Description: `Storage expense of the entity.`,
				},
				resource.Attribute{
					Name:        "total_expense",
					Description: `Total expense of the entity.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `Monetary unit.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of this entity.`,
				},
				resource.Attribute{
					Name:        "inputs_including_defaults",
					Description: `All the inputs applied during last create/update operation, including those with default values. For the list of inputs provided by the user in the configuration, refer to ` + "`" + `inputs` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "last_request",
					Description: `Represents deployment requests.`,
				},
				resource.Attribute{
					Name:        "action_id",
					Description: `Identifier of the requested action.`,
				},
				resource.Attribute{
					Name:        "approved_at",
					Description: `Time at which the request was approved.`,
				},
				resource.Attribute{
					Name:        "blueprint_id",
					Description: `Identifier of the requested blueprint in the form ‘UUID:version’.`,
				},
				resource.Attribute{
					Name:        "cancelable",
					Description: `Indicates whether request can be canceled or not.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `Identifier of the requested catalog item in the form ‘UUID:version’.`,
				},
				resource.Attribute{
					Name:        "completed_at",
					Description: `Time at which the request completed.`,
				},
				resource.Attribute{
					Name:        "completed_tasks",
					Description: `The number of tasks completed while fulfilling this request.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `Longer user-friendly details of the request.`,
				},
				resource.Attribute{
					Name:        "dismissed",
					Description: `Indicates whether request is in dismissed state.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Request identifier.`,
				},
				resource.Attribute{
					Name:        "initialized_at",
					Description: `Time at which the request was initialized.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `List of request inputs.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Short user-friendly label of the request (e.g. ‘shuting down myVM’).`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `Request outputs.`,
				},
				resource.Attribute{
					Name:        "requested_by",
					Description: `The user that initiated the request.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `Optional resource name to which the request applies to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Request overall execution status. Supported values: ` + "`" + `CREATED` + "`" + `, ` + "`" + `PENDING` + "`" + `, ` + "`" + `INITIALIZATION` + "`" + `, ` + "`" + `CHECKING_APPROVAL` + "`" + `, ` + "`" + `APPROVAL_PENDING` + "`" + `, ` + "`" + `INPROGRESS` + "`" + `, ` + "`" + `COMPLETION` + "`" + `, ` + "`" + `APPROVAL_REJECTED` + "`" + `, ` + "`" + `ABORTED` + "`" + `, ` + "`" + `SUCCESSFUL` + "`" + `, ` + "`" + `FAILED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Last update time (e.g. date format ‘2019-07-13T23:16:49.310Z’).`,
				},
				resource.Attribute{
					Name:        "last_updated_at",
					Description: `Time at which the deployment was last updated.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user that last updated the deployment.`,
				},
				resource.Attribute{
					Name:        "lease_expire_at",
					Description: `Time at which the deployment lease expires.`,
				},
				resource.Attribute{
					Name:        "project",
					Description: `The project this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human friendly description.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the entity.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the entity.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version of the entity, if applicable.`,
				},
				resource.Attribute{
					Name:        "resources",
					Description: `Expanded resources for the deployment. Content of this property will not be maintained backward compatible.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation time (e.g. date format ‘2019-07-13T23:16:49.310Z’).`,
				},
				resource.Attribute{
					Name:        "depends_on",
					Description: `A list of other resources this resource depends on.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the resource. ` + "`" + `expense` + "`" + ` - Expense incurred for this resource.`,
				},
				resource.Attribute{
					Name:        "additional_expense",
					Description: `Additional expense incurred for the resource.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Expense sync message code if any.`,
				},
				resource.Attribute{
					Name:        "compute_expense",
					Description: `Compute expense of the entity.`,
				},
				resource.Attribute{
					Name:        "last_update_time",
					Description: `Last expense sync time.`,
				},
				resource.Attribute{
					Name:        "message",
					Description: `Expense sync message if any.`,
				},
				resource.Attribute{
					Name:        "network_expense",
					Description: `Network expense of the entity.`,
				},
				resource.Attribute{
					Name:        "storage_expense",
					Description: `Storage expense of the entity.`,
				},
				resource.Attribute{
					Name:        "total_expense",
					Description: `Total expense of the entity.`,
				},
				resource.Attribute{
					Name:        "unit",
					Description: `Monetary unit.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique identifier of the resource.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the resource.`,
				},
				resource.Attribute{
					Name:        "properties_json",
					Description: `List of properties in the encoded JSON string format.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The current state of the resource. Supported values are ` + "`" + `PARTIAL` + "`" + `, ` + "`" + `TAINTED` + "`" + `, ` + "`" + `OK.` + "`" + ``,
				},
				resource.Attribute{
					Name:        "sync_status",
					Description: `The current sync status. Supported values are ` + "`" + `SUCCESS` + "`" + `, ` + "`" + `MISSING` + "`" + `, ` + "`" + `STALE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the resource.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Deployment status. Supported values are: ` + "`" + `CREATE_SUCCESSFUL` + "`" + `, ` + "`" + `CREATE_INPROGRESS` + "`" + `, ` + "`" + `CREATE_FAILED` + "`" + `, ` + "`" + `UPDATE_SUCCESSFUL` + "`" + `, ` + "`" + `UPDATE_INPROGRESS` + "`" + `, ` + "`" + `UPDATE_FAILED` + "`" + `, ` + "`" + `DELETE_SUCCESSFUL` + "`" + `, ` + "`" + `DELETE_INPROGRESS` + "`" + `, ` + "`" + `DELETE_FAILED` + "`" + `, ` + "`" + `ACTION_SUCCESSFUL` + "`" + `, ` + "`" + `ACTION_INPROGRESS` + "`" + `, ` + "`" + `ACTION_FAILED` + "`" + `. ## Import Deployment can be imported using the id, e.g. ` + "`" + `$ terraform import vra_deployment.this 05956583-6488-4e7d-84c9-92a7b7219a15` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_flavor_profile",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_flavor_profile.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The id of the region for which this profile is defined as in vRealize Automation(vRA).`,
				},
				resource.Attribute{
					Name:        "flavor_mapping",
					Description: `(Optional) Map between global fabric flavor keys and fabric flavor descriptions. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `The ID of the cloud account this flavor profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The ID of the external region that is associated with the flavor profile.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `The ID of the cloud account this flavor profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The ID of the external region that is associated with the flavor profile.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_image_profile",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_image_profile.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The id of the region for which this profile is defined as in vRealize Automation(vRA). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "image_mapping",
					Description: `Image mapping defined for the corresponding region.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "image_mapping",
					Description: `Image mapping defined for the corresponding region.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_load_balancer",
			Category:         "Resources",
			ShortDescription: `Provides a VMware vRA vra_load_balancer resource.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "custom_properties",
					Description: `(Optional) Additional custom properties that may be used to extend the machine.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) The id of the deployment that is associated with this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Describes machine within the scope of your organization and is not propagated to the cloud.`,
				},
				resource.Attribute{
					Name:        "internet_facing",
					Description: `(Optional) An Internet-facing load balancer has a publicly resolvable DNS name, so it can route requests from clients over the Internet to the instances that are registered with the load balancer.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Required) A set of network interface specifications for this load balancer.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project the current user belongs to.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `(Required) The load balancer route configuration regarding ports and protocols.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `Algorithm employed for load balancing.`,
				},
				resource.Attribute{
					Name:        "algorithm_parameters",
					Description: `Parameters need for load balancing algorithm.Use newline to separate multiple parameters.`,
				},
				resource.Attribute{
					Name:        "health_check_configuration",
					Description: `Load balancer health check configuration.`,
				},
				resource.Attribute{
					Name:        "healthy_threshold",
					Description: `Number of consecutive successful checks before considering a particular back-end instance as healthy.`,
				},
				resource.Attribute{
					Name:        "http_method",
					Description: `HTTP or HTTPS method to use when sending a health check request.`,
				},
				resource.Attribute{
					Name:        "interval_seconds",
					Description: `Interval (in seconds) at which the health checks will be performed.`,
				},
				resource.Attribute{
					Name:        "passive_monitor",
					Description: `Enable passive monitor mode. This setting only applies to NSX-T.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port on the back-end instance machine to use for the health check.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used for the health check.`,
				},
				resource.Attribute{
					Name:        "request_body",
					Description: `Request body. Used by HTTP, HTTPS, TCP, UDP.`,
				},
				resource.Attribute{
					Name:        "response_body",
					Description: `Expected response body. Used by HTTP, HTTPS, TCP, UDP.`,
				},
				resource.Attribute{
					Name:        "timeout_seconds",
					Description: `Timeout (in seconds) to wait for a response from the back-end instance.`,
				},
				resource.Attribute{
					Name:        "unhealthy_threashold",
					Description: `Number of consecutive check failures before considering a particular back-end instance as unhealthy.`,
				},
				resource.Attribute{
					Name:        "urlPath",
					Description: `URL path on the back-end instance against which a request will be performed for the health check. Useful when the health check protocol is HTTP/HTTPS.`,
				},
				resource.Attribute{
					Name:        "member_port",
					Description: `Member port where the traffic is routed to.`,
				},
				resource.Attribute{
					Name:        "member_protocol",
					Description: `The protocol of the member traffic.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port which the load balancer is listening to.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol of the incoming load balancer requests. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A list of links to target load balancer pool members. Links can be to either a machine or a machine's network interface.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Primary address allocated or in use by this load balancer. The address could be an in the form of a publicly resolvable DNS name or an IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource instance. example:[ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Tag’s key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Tag’s value.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `A list of links to target load balancer pool members. Links can be to either a machine or a machine's network interface.`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Primary address allocated or in use by this load balancer. The address could be an in the form of a publicly resolvable DNS name or an IP address.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource instance. example:[ { "key" : "vmware.enumeration.type", "value": "nebs_block" } ]`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Tag’s key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Tag’s value.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_machine",
			Category:         "Resources",
			ShortDescription: `Creates a vra_machine resource.`,
			Description: `

Creates a VMware vRealize Automation machine resource.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "boot_config",
					Description: `(Optional) Machine boot config that will be passed to the instance. Used to perform common automated configuration tasks and even run scripts after instance starts.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Calid cloud config data in json-escaped yaml syntax.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `(Optional) Additional properties that may be used to extend the base resource.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) Describes machine within the scope of your organization and is not propagated to the cloud.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `(Optional) Specification for attaching/detaching disks to a machine.`,
				},
				resource.Attribute{
					Name:        "block_device_id",
					Description: `(Required) ID of the existing block device.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-friendly block-device name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Required) Flavor of machine instance.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Optional) Type of image used for this machine.`,
				},
				resource.Attribute{
					Name:        "image_disk_constraints",
					Description: `(Optional) Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets. example:[{"mandatory" : "true", "expression": "environment:prod"}, {"mandatory" : "false", "expression": "pci"}]. It is nested argument with the following properties.`,
				},
				resource.Attribute{
					Name:        "expression",
					Description: `(Required) Constraint that is conveyed to the policy engine. An expression of the form "[!]tag-key[:[tag-value]]", used to indicate a constraint match on keys and values of tags.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `(Required) Indicates whether this constraint should be strictly enforced or not.`,
				},
				resource.Attribute{
					Name:        "image_ref",
					Description: `(Optional) Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `(Optional) Set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `(Optional) List of IP addresses allocated or in use by this network interface. example:[ "10.1.2.190" ]`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `(Optional) Additional properties that may be used to extend the base type.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "device_index",
					Description: `(Optional) The device index of this network interface.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) ID of the network instance that this network interface plugs into.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) List of security group ids which this network interface will be assigned to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Set of tag keys and optional values that should be set on any resource that is produced from this specification. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]. It is nested argument with the following properties.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Tag’s key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Tag’s value. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.`,
				},
				resource.Attribute{
					Name:        "constraints",
					Description: `Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets. Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "disks_list",
					Description: `List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.`,
				},
				resource.Attribute{
					Name:        "block_device_id",
					Description: `ID of existing block device.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human-friendly block-device name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity ID on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `External regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `External zoneId of the resource.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `Power state of machine.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project that resource belongs to.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.`,
				},
				resource.Attribute{
					Name:        "constraints",
					Description: `Constraints used to drive placement policies for the virtual machine produced from the specification. Constraint expressions are matched against tags on existing placement targets. Example:[{"mandatory" : "true", "expression": "environment":"prod"}, {"mandatory" : "false", "expression": "pci"}]`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "disks_list",
					Description: `List of all disks attached to a machine including boot disk, and additional block devices attached using the disks attribute.`,
				},
				resource.Attribute{
					Name:        "block_device_id",
					Description: `ID of existing block device.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human-friendly block-device name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity ID on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `External regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `External zoneId of the resource.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `Power state of machine.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `ID of project that resource belongs to.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_network",
			Category:         "Resources",
			ShortDescription: `Provides a VMware vRA vra_network resource.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "custom_properties",
					Description: `(Optional) Additional properties that may be used to extend the base resource.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `(Optional) Deployment id that is associated with this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "outbound_access",
					Description: `(Optional) Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project this resource belongs to. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `IPv4 address range of the network in CIDR format.`,
				},
				resource.Attribute{
					Name:        "constraints",
					Description: `List of storage, network and extensibility constraints to be applied when provisioning through this project.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `The external zoneId of the resource.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `Self link of this request.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cidr",
					Description: `IPv4 address range of the network in CIDR format.`,
				},
				resource.Attribute{
					Name:        "constraints",
					Description: `List of storage, network and extensibility constraints to be applied when provisioning through this project.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_zone_id",
					Description: `The external zoneId of the resource.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `Self link of this request.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_network_profile",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_network_profile.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "custom_properties",
					Description: `(Optional) Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "fabric_network_ids",
					Description: `(Optional) A list of fabric network Ids which are assigned to the network profile. example:[ "6543" ]`,
				},
				resource.Attribute{
					Name:        "isolated_external_fabric_network_id",
					Description: `(Optional) The id of the fabric network used for outbound access.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The id of the region for which this profile is defined as in vRealize Automation(vRA). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `The ID of the cloud account this flavor profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "isolated_network_cidr_prefix",
					Description: `The CIDR prefix length to be used for the isolated networks that are created with the network profile.`,
				},
				resource.Attribute{
					Name:        "isolated_network_domain_cidr",
					Description: `CIDR of the isolation network domain.`,
				},
				resource.Attribute{
					Name:        "isolated_network_domain_id",
					Description: `The id of the network domain used for creating isolated networks.`,
				},
				resource.Attribute{
					Name:        "isolation_type",
					Description: `Specifies the isolation type e.g. none, subnet or security group`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to. Deprecated, refer to org_id instead.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `A list of security group Ids which are assigned to the network profile. example:[ "6545" ]`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `The ID of the cloud account this flavor profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "isolated_network_cidr_prefix",
					Description: `The CIDR prefix length to be used for the isolated networks that are created with the network profile.`,
				},
				resource.Attribute{
					Name:        "isolated_network_domain_cidr",
					Description: `CIDR of the isolation network domain.`,
				},
				resource.Attribute{
					Name:        "isolated_network_domain_id",
					Description: `The id of the network domain used for creating isolated networks.`,
				},
				resource.Attribute{
					Name:        "isolation_type",
					Description: `Specifies the isolation type e.g. none, subnet or security group`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to. Deprecated, refer to org_id instead.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `A list of security group Ids which are assigned to the network profile. example:[ "6545" ]`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_project",
			Category:         "Resources",
			ShortDescription: `Provides a VMware vRA vra_project resource.`,
			Description: `\_project
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "administrators",
					Description: `(Optional) List of administrator users associated with the project. Only administrators can manage project's configuration.`,
				},
				resource.Attribute{
					Name:        "constraints",
					Description: `(Optional) List of storage, network and extensibility constraints to be applied when provisioning through this project.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "machine_naming_template",
					Description: `(Optional) The naming template to be used for resources provisioned in this project.`,
				},
				resource.Attribute{
					Name:        "members",
					Description: `(Optional) List of member users associated with the project.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "operation_timeout",
					Description: `(Optional) The timeout that should be used for Blueprint operations and Provisioning tasks. The timeout is in seconds.`,
				},
				resource.Attribute{
					Name:        "shared_resources",
					Description: `(Optional) Specifies whether the resources in this projects are shared or not. If not set default will be used.`,
				},
				resource.Attribute{
					Name:        "viewer",
					Description: `(Optional) List of viewer users associated with the project.`,
				},
				resource.Attribute{
					Name:        "zone_assignments",
					Description: `(Optional) List of configurations for zone assignment to a project.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_storage_profile",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_storage_profile.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_properties",
					Description: `(Optional) Map of storage properties that are to be applied on disk while provisioning.`,
				},
				resource.Attribute{
					Name:        "disk_target_properties",
					Description: `(Optional) Map of storage placements to know where the disk is provisioned.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The id of the region for which this profile is defined as in vRealize Automation(vRA).`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `(Optional) Indicates whether this storage profile supports encryption or not. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `Id of the cloud account this storage profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `Id of the cloud account this storage profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_storage_profile_aws",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_storage_profile_aws.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default_item",
					Description: `(Required) Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `(Optional) Indicates the type of storage device.`,
				},
				resource.Attribute{
					Name:        "iops",
					Description: `(Optional) Indicates maximum I/O operations per second in range(1-20,000).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) A link to the region that is associated with the storage profile.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `(Optional) Indicates whether this storage profile supports encryption or not.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `(Optional) Indicates the type of volume associated with type of storage. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_storage_profile_azure",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_storage_profile_azure.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_disk_caching",
					Description: `(Optional) Indicates the caching mechanism for additional disk.`,
				},
				resource.Attribute{
					Name:        "default_item",
					Description: `(Required) Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "os_disk_caching",
					Description: `(Optional) Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) A link to the region that is associated with the storage profile.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `(Optional) Id of a storage account where in the disk is placed.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `(Optional) Indicates whether this storage policy should support encryption or not. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_storage_profile_vsphere",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_storage_profile_vsphere.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) Id of the vSphere Datastore for placing disk and VM.`,
				},
				resource.Attribute{
					Name:        "default_item",
					Description: `(Required) Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_mode",
					Description: `(Optional) Type of mode for the disk.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `(Optional) Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.`,
				},
				resource.Attribute{
					Name:        "limit_iops",
					Description: `(Optional) The upper bound for the I/O operations per second allocated for each virtual disk.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "provisioning_type",
					Description: `(Optional) Type of provisioning policy for the disk.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) The Id of the region that is associated with the storage profile.`,
				},
				resource.Attribute{
					Name:        "shares",
					Description: `(Optional) A specific number of shares assigned to each virtual machine.`,
				},
				resource.Attribute{
					Name:        "shares_level",
					Description: `(Optional) Indicates whether this storage profile supports encryption or not.`,
				},
				resource.Attribute{
					Name:        "storage_policy_id",
					Description: `(Optional) Id of the vSphere Storage Policy to be applied.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `(Optional) Indicates whether this storage policy should support encryption or not. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `Id of the cloud account this storage profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `Id of the cloud account this storage profile belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_zone",
			Category:         "Resources",
			ShortDescription: `Provides a VMware vRA vra_zone resource.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The folder relative path to the datacenter where resources are deployed to. (only applicable for vSphere cloud zones)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "placement_policy",
					Description: `(Optional) The id of the region for which this zone is defined. Valid values are: ` + "`" + `DEFAULT` + "`" + `, ` + "`" + `SPREAD` + "`" + `, ` + "`" + `BINPACK` + "`" + `. Default is ` + "`" + `DEFAULT` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Required) A link to the region that is associated with the zone. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A set of tag keys and optional values that were set on this zone.`,
				},
				resource.Attribute{
					Name:        "tags_to_match",
					Description: `(Optional) A set of tag keys and optional values for compute resource filtering. example:[ { "key" : "compliance", "value": "pci" } ] ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `The ID of the cloud account this zone belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `A list of key value pair of properties related to the zone.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The ID of the external region that is associated with the zone.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `The ID of the cloud account this zone belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `A list of key value pair of properties related to the zone.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The ID of the external region that is associated with the zone.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of entity owner.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"vra_cloud_account_aws":              0,
		"vra_cloud_account_azure":            1,
		"vra_cloud_account_gcp":              2,
		"vra_cloud_account_nsxt":             3,
		"vra_cloud_account_nsxv":             4,
		"vra_cloud_account_vmc":              5,
		"vra_cloud_account_vsphere":          6,
		"vra_fabric_network_vsphere":         7,
		"vra_network_ip_range":               8,
		"vra_vra_block_device":               9,
		"vra_vra_block_device_snapshot":      10,
		"vra_vra_blueprint":                  11,
		"vra_vra_blueprint_version":          12,
		"vra_vra_catalog_source_blueprint":   13,
		"vra_vra_catalog_source_entitlement": 14,
		"vra_vra_content_source":             15,
		"vra_vra_deployment":                 16,
		"vra_vra_flavor_profile":             17,
		"vra_vra_image_profile":              18,
		"vra_vra_load_balancer":              19,
		"vra_vra_machine":                    20,
		"vra_vra_network":                    21,
		"vra_vra_network_profile":            22,
		"vra_vra_project":                    23,
		"vra_vra_storage_profile":            24,
		"vra_vra_storage_profile_aws":        25,
		"vra_vra_storage_profile_azure":      26,
		"vra_vra_storage_profile_vsphere":    27,
		"vra_vra_zone":                       28,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
