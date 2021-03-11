package vra

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_aws",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_aws.`,
			Description: `\_cloud\_account\_aws

Provides a VMware vRA vra_cloud_account_aws data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this AWS cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this AWS cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "access_key",
					Description: `Access key id for Aws.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
					Name:        "access_key",
					Description: `Access key id for Aws.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
			Type:             "vra_cloud_account_azure",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_azure.`,
			Description: `\_cloud\_account\_azure

Provides a VMware vRA vra_cloud_account_azure data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this Azure cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this Azure cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "application_id",
					Description: `Azure Client Application ID.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `Azure Subscription ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
					Name:        "tenant_id",
					Description: `Azure Tenant ID.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `Azure Client Application ID.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `Azure Subscription ID.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
					Name:        "tenant_id",
					Description: `Azure Tenant ID.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_gcp",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_gcp.`,
			Description: `\_cloud\_account\_gcp

Provides a VMware vRA vra_cloud_account_gcp data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this GCP cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this GCP cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "client_email",
					Description: `GCP Client email.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Name:        "private_key_id",
					Description: `GCP Private key ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `GCP Project ID.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
					Name:        "client_email",
					Description: `GCP Client email.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Name:        "private_key_id",
					Description: `GCP Private key ID.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `GCP Project ID.`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
			Type:             "vra_cloud_account_nsxt",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_nsxt.`,
			Description: `\_cloud\_account\_nsxt

Provides a VMware vRA vra_cloud_account_nsxt data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this NSX-T cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Host name for the NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "username",
					Description: `Username to authenticate with the cloud account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Host name for the NSX-T cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "username",
					Description: `Username to authenticate with the cloud account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_nsxv",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_nsxv.`,
			Description: `\_cloud\_account\_nsxv

Provides a VMware vRA vra_cloud_account_nsxv data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this NSX-V cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Host name for the NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "username",
					Description: `Username to authenticate with the cloud account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Host name for the NSX-V cloud account.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "username",
					Description: `Username to authenticate with the cloud account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_vmc",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_vmc.`,
			Description: `\_cloud\_account\_vmc

Provides a VMware vRA vra_cloud_account_vmc data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this vmc cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this vmc cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "nsx_hostname",
					Description: `The IP address of the NSX Manager server in the specified SDDC / FQDN.`,
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
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "region-ids",
					Description: `A set of region IDs that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `Identifier of the on-premise SDDC to be used by this cloud account. Note that NSX-V SDDCs are not supported.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "vcenter_hostname",
					Description: `The IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.`,
				},
				resource.Attribute{
					Name:        "vcenter_username",
					Description: `vCenter user name for the specified SDDC. The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collector.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "nsx_hostname",
					Description: `The IP address of the NSX Manager server in the specified SDDC / FQDN.`,
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
					Name:        "regions",
					Description: `A set of region names that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "region-ids",
					Description: `A set of region IDs that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `Identifier of the on-premise SDDC to be used by this cloud account. Note that NSX-V SDDCs are not supported.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "vcenter_hostname",
					Description: `The IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.`,
				},
				resource.Attribute{
					Name:        "vcenter_username",
					Description: `vCenter user name for the specified SDDC. The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_cloud_account_vsphere",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_cloud_account_vsphere.`,
			Description: `\_cloud\_account\_vsphere

Provides a VMware vRA vra_cloud_account_vsphere data source.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this vSphere cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of this vSphere cloud account. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the base info.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "enabled_region_ids",
					Description: `A set of region IDs that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "username",
					Description: `The vSphere username to authenticate the vsphere account.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "associated_cloud_account_ids",
					Description: `Cloud accounts associated with this cloud account.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the base info.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `Identifier of a data collector vm deployed in the on premise infrastructure.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "enabled_region_ids",
					Description: `A set of region IDs that are enabled for this account.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The IP address or FQDN of the vCenter Server. The cloud proxy belongs on this vCenter.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "vmware", "value": "provider" } ]`,
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
				resource.Attribute{
					Name:        "username",
					Description: `The vSphere username to authenticate the vsphere account.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_fabric_datastore_vsphere",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vSphere fabric datastores.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Only one of 'filter' or 'id' must be specified.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the vSphere data source. Only one of 'filter' or 'id' must be specified. ## Attribute Reference`,
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
					Description: `Id of datacenter in which the datastore is present.`,
				},
				resource.Attribute{
					Name:        "free_size_gb",
					Description: `Indicates free size available in datastore.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the datastore.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of datastore.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
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
					Description: `Id of datacenter in which the datastore is present.`,
				},
				resource.Attribute{
					Name:        "free_size_gb",
					Description: `Indicates free size available in datastore.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the datastore.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of datastore.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_fabric_network",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vRA fabric networks.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Filter query string that is supported by vRA multi-cloud IaaS API. ## Attribute Reference`,
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
					Name:        "cidr",
					Description: `Network CIDR to be used.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the vRA fabric network.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicates whether the sub-network supports public IP assignment.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether this is the default subnet for the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `State object representing a network on a external cloud provider.`,
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
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the fabric network.`,
				},
				resource.Attribute{
					Name:        "organization_id",
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
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
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
					Name:        "cidr",
					Description: `Network CIDR to be used.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the vRA fabric network.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `Indicates whether the sub-network supports public IP assignment.`,
				},
				resource.Attribute{
					Name:        "is_default",
					Description: `Indicates whether this is the default subnet for the zone.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `State object representing a network on a external cloud provider.`,
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
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the fabric network.`,
				},
				resource.Attribute{
					Name:        "organization_id",
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
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_fabric_storage_account_azure",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for fabric Azure storage account.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `Search criteria to narrow down the fabric Azure storage accounts. Only one of 'filter' or 'id' must be specified.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the fabric Azure storage account. Only one of 'filter' or 'id' must be specified. ## Attribute Reference`,
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
					Name:        "description",
					Description: `A human-friendly description of the fabric Azure storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed. example: Standard_LRS / Premium_LRS`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
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
					Name:        "description",
					Description: `A human-friendly description of the fabric Azure storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
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
					Description: `A set of tag keys and optional values that were set on this resource. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed. example: Standard_LRS / Premium_LRS`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_fabric_storage_policy_vsphere",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for fabric vSphere storage policy.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `Search criteria to narrow down the fabric vSphere storage policy. Only one of 'filter' or 'id' must be specified.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the fabric vSphere storage policy. Only one of 'filter' or 'id' must be specified. ## Attribute Reference`,
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
					Name:        "description",
					Description: `A human-friendly description of the fabric vSphere storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option. Only one of 'filter', 'id', 'name' or 'region_id' must be specified.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
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
					Name:        "description",
					Description: `A human-friendly description of the fabric vSphere storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option. Only one of 'filter', 'id', 'name' or 'region_id' must be specified.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_image",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vRA Images.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `Search criteria to narrow down Images.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Image. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description of the fabric vSphere storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Indicates whether this fabric image is private. For vSphere, private images are considered to be templates and snapshots and public are Content Library Items.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The regionId of the image.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description of the fabric vSphere storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "private",
					Description: `Indicates whether this fabric image is private. For vSphere, private images are considered to be templates and snapshots and public are Content Library Items.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The regionId of the image.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_network_domain",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for Network domain objects.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Search criteria to narrow down the network domain objects. ## Attribute Reference`,
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
					Name:        "description",
					Description: `A human-friendly description of the fabric vSphere storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the fabric network domain object.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network domain.`,
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
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
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
					Name:        "description",
					Description: `A human-friendly description of the fabric vSphere storage account.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the fabric network domain object.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the network domain.`,
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
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_region",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_region.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes:       []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_security_group",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for security groups.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `(Required) Search criteria to narrow down the Security groups. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description of the security groups.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of security rules.`,
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
					Name:        "description",
					Description: `A human-friendly description of the security groups.`,
				},
				resource.Attribute{
					Name:        "external_id",
					Description: `External entity Id on the provider side.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region for which this entity is defined.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the security group.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the security group.`,
				},
				resource.Attribute{
					Name:        "organization_id",
					Description: `ID of organization that entity belongs to.`,
				},
				resource.Attribute{
					Name:        "rules",
					Description: `List of security rules.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_block_device",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_block_device.`,
			Description: `

Provides a data lookup for a vra_block_device.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the block device.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Search criteria to filter the list of block devices.`,
				},
				resource.Attribute{
					Name:        "expand_snapshots",
					Description: `(Optional) Indicates whether the snapshots of the block-devices should be included in the state. Applicable only for first class block devices. ## Attributes Reference A block device data source supports the following attributes:`,
				},
				resource.Attribute{
					Name:        "capacity_in_gb",
					Description: `Capacity of the block device in GB.`,
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
					Name:        "custom_properties",
					Description: `Additional custom properties that may be used to extend the machine.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `The id of the deployment that is associated with this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describes machine within the scope of your organization and is not propagated to the cloud.`,
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
					Description: `The external zoneId of the resource.`,
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
					Name:        "project_id",
					Description: `The id of the project the current user belongs to.`,
				},
				resource.Attribute{
					Name:        "persistent",
					Description: `Indicates whether the block device survives a delete action.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the block device.`,
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
					Name:        "capacity_in_gb",
					Description: `Capacity of the block device in GB.`,
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
					Name:        "custom_properties",
					Description: `Additional custom properties that may be used to extend the machine.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `The id of the deployment that is associated with this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describes machine within the scope of your organization and is not propagated to the cloud.`,
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
					Description: `The external zoneId of the resource.`,
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
					Name:        "project_id",
					Description: `The id of the project the current user belongs to.`,
				},
				resource.Attribute{
					Name:        "persistent",
					Description: `Indicates whether the block device survives a delete action.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the block device.`,
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
			Type:             "vra_vra_block_device_snapshot",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_block_device_snapshots.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "var.block_device_id",
					Description: `(Required) The id of the existing block device. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Indicates whether this snapshot is the current snapshot on the block-device.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option. Only one of 'filter', 'id', 'name' or 'region_id' must be specified.`,
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
					Name:        "update_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "is_current",
					Description: `Indicates whether this snapshot is the current snapshot on the block-device.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `HATEOAS of the entity`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option. Only one of 'filter', 'id', 'name' or 'region_id' must be specified.`,
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
					Name:        "update_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_blueprint",
			Category:         "Resources",
			ShortDescription: `A blueprint data source.`,
			Description: `\_blueprint

This data source provides information about a cloud template (blueprint) in vRA.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of this cloud template. One of ` + "`" + `id` + "`" + ` or ` + "`" + `name` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the cloud template. One of ` + "`" + `id` + "`" + ` or ` + "`" + `name` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The id of the project to narrow the search while looking for cloud templates. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Blueprint YAML content.`,
				},
				resource.Attribute{
					Name:        "content_source_id",
					Description: `The id of the content source.`,
				},
				resource.Attribute{
					Name:        "content_source_path",
					Description: `Content source path.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_at",
					Description: `Content source last sync at.`,
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
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The name of the project the entity belongs to.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "request_scope_org",
					Description: `Flag to indicate whether this blueprint can be requested from any project in the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cloud template. Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
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
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template is valid.`,
				},
				resource.Attribute{
					Name:        "validation_messages",
					Description: `List of validations messages.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "content",
					Description: `Blueprint YAML content.`,
				},
				resource.Attribute{
					Name:        "content_source_id",
					Description: `The id of the content source.`,
				},
				resource.Attribute{
					Name:        "content_source_path",
					Description: `Content source path.`,
				},
				resource.Attribute{
					Name:        "content_source_sync_at",
					Description: `Content source last sync at.`,
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
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The name of the project the entity belongs to.`,
				},
				resource.Attribute{
					Name:        "self_link",
					Description: `HATEOAS of the entity.`,
				},
				resource.Attribute{
					Name:        "request_scope_org",
					Description: `Flag to indicate whether this blueprint can be requested from any project in the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cloud template. Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
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
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template is valid.`,
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
			ShortDescription: `A blueprint version data source.`,
			Description: `\_blueprint\_version

This data source provides information about a cloud template (blueprint) version in vRA.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_id",
					Description: `(Required) Name of the cloud template. One of ` + "`" + `id` + "`" + ` or ` + "`" + `name` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the cloud template version. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "blueprint_description",
					Description: `Description of the cloud template.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Blueprint YAML content.`,
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
					Name:        "description",
					Description: `(Optional) Cloud template version description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cloud template version.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The name of the project the entity belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cloud template. Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template is valid.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cloud template version.`,
				},
				resource.Attribute{
					Name:        "version_change_log",
					Description: `Cloud template version change log.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_description",
					Description: `Description of the cloud template.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `Blueprint YAML content.`,
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
					Name:        "description",
					Description: `(Optional) Cloud template version description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the cloud template version.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "project_name",
					Description: `The name of the project the entity belongs to.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the cloud template. Supported values: ` + "`" + `DRAFT` + "`" + `, ` + "`" + `VERSIONED` + "`" + `, ` + "`" + `RELEASED` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "valid",
					Description: `Flag to indicate if the current content of the cloud template is valid.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Cloud template version.`,
				},
				resource.Attribute{
					Name:        "version_change_log",
					Description: `Cloud template version change log.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_catalog_item",
			Category:         "Resources",
			ShortDescription: `A data source for a catalog item.`,
			Description: `\_catalog\_item

This data source provides information about a catalog item in vRA.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "expand_projects",
					Description: `(Optional) Flag to indicate whether to expand detailed project data for the catalog item.`,
				},
				resource.Attribute{
					Name:        "expand_versions",
					Description: `(Optional) Flag to indicate whether to expand detailed versions of the catalog item.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of catalog item. One of ` + "`" + `id` + "`" + `, or ` + "`" + `name` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the catalog item. One of ` + "`" + `id` + "`" + `, or ` + "`" + `name` + "`" + ` must be provided. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date-time when the entity was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Catalog item description.`,
				},
				resource.Attribute{
					Name:        "last_updated_at",
					Description: `Date-time when the entity was last updated.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `List of associated project IDs that can be used for requesting this catalog item.`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `List of associated projects that can be used for requesting this catalog item.`,
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
					Name:        "schema",
					Description: `Json schema describing request parameters, a simplified version of http://json-schema.org/latest/json-schema-validation.html#rfc.section.5`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `LibraryItem source ID.`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `LibraryItem source name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: ``,
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
					Name:        "versions",
					Description: `Catalog item versions.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date-time when catalog item version was created at.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the catalog item version.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date-time when the entity was created.`,
				},
				resource.Attribute{
					Name:        "created_by",
					Description: `The user the entity was created by.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Catalog item description.`,
				},
				resource.Attribute{
					Name:        "last_updated_at",
					Description: `Date-time when the entity was last updated.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user the entity was last updated by.`,
				},
				resource.Attribute{
					Name:        "project_ids",
					Description: `List of associated project IDs that can be used for requesting this catalog item.`,
				},
				resource.Attribute{
					Name:        "projects",
					Description: `List of associated projects that can be used for requesting this catalog item.`,
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
					Name:        "schema",
					Description: `Json schema describing request parameters, a simplified version of http://json-schema.org/latest/json-schema-validation.html#rfc.section.5`,
				},
				resource.Attribute{
					Name:        "source_id",
					Description: `LibraryItem source ID.`,
				},
				resource.Attribute{
					Name:        "source_name",
					Description: `LibraryItem source name.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: ``,
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
					Name:        "versions",
					Description: `Catalog item versions.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date-time when catalog item version was created at.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the catalog item version.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_catalog_source_blueprint",
			Category:         "Resources",
			ShortDescription: `A data source for catalog source of type cloud template (blueprint).`,
			Description: `\_catalog\_source\_blueprint

This data source provides information about a catalog source of type cloud template (blueprint) in vRA.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of catalog source. One of ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` or ` + "`" + `project_id` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the catalog source. One of ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` or ` + "`" + `project_id` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The id of the project. One of ` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + ` or ` + "`" + `project_id` + "`" + ` must be provided. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `Custom configuration of the catalog source as a map of key values.`,
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
					Name:        "description",
					Description: `Catalog source description.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `A flag indicating that all the items can be requested across all projects.`,
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
					Description: `Time at which the last import was completed at.`,
				},
				resource.Attribute{
					Name:        "last_import_errors",
					Description: `A list of errors seen at last time the catalog source is imported.`,
				},
				resource.Attribute{
					Name:        "last_import_started_at",
					Description: `Time at which the last import was started at.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user that last updated the catalog source.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Type of catalog source. Example: ` + "`" + `blueprint` + "`" + `, ` + "`" + `CFT` + "`" + `, etc.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "config",
					Description: `Custom configuration of the catalog source as a map of key values.`,
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
					Name:        "description",
					Description: `Catalog source description.`,
				},
				resource.Attribute{
					Name:        "global",
					Description: `A flag indicating that all the items can be requested across all projects.`,
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
					Description: `Time at which the last import was completed at.`,
				},
				resource.Attribute{
					Name:        "last_import_errors",
					Description: `A list of errors seen at last time the catalog source is imported.`,
				},
				resource.Attribute{
					Name:        "last_import_started_at",
					Description: `Time at which the last import was started at.`,
				},
				resource.Attribute{
					Name:        "last_updated_by",
					Description: `The user that last updated the catalog source.`,
				},
				resource.Attribute{
					Name:        "type_id",
					Description: `Type of catalog source. Example: ` + "`" + `blueprint` + "`" + `, ` + "`" + `CFT` + "`" + `, etc.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_catalog_source_entitlement",
			Category:         "Resources",
			ShortDescription: `A data source for catalog source entitlement.`,
			Description: `\_catalog\_source\_entitlement

This data source provides information about a catalog source entitlement in vRA.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "catalog_source_id",
					Description: `(Optional) The id of the catalog source to find the entitlement. One of ` + "`" + `catalog_source_id` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of entitlement. One of ` + "`" + `catalog_source_id` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The id of the project that this entitlement belongs to. ## Attribute Reference`,
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
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_data_collector",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for data collector data source.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Data collector name. Example: Datacollector1 ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Data collector host name. Example: dc1-lnd.mycompany.com`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IPv4 Address of the data collector VM. Example: 10.0.0.1`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the data collector. Example: ACTIVE, INACTIVE`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `Data collector host name. Example: dc1-lnd.mycompany.com`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IPv4 Address of the data collector VM. Example: 10.0.0.1`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Current status of the data collector. Example: ACTIVE, INACTIVE`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_deployment",
			Category:         "Resources",
			ShortDescription: `A deployment data source.`,
			Description: `\_deployment

This data source provides information about a deployment in vRA.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "expand_last_request",
					Description: `(Optional) Flag to indicate whether to expand last request on the deployment.`,
				},
				resource.Attribute{
					Name:        "expand_project",
					Description: `(Optional) Flag to indicate whether to expand project information.`,
				},
				resource.Attribute{
					Name:        "expand_resources",
					Description: `(Optional) Flag to indicate whether to expand resources in the deployment.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the deployment. One of ` + "`" + `id` + "`" + ` or ` + "`" + `name` + "`" + ` must be provided.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the deployment. One of ` + "`" + `id` + "`" + ` or ` + "`" + `name` + "`" + ` must be provided. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "blueprint_d",
					Description: `The id of the vRA cloud template to request the deployment. Conflicts with ` + "`" + `catalog_item_id` + "`" + ` and ` + "`" + `blueprint_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "blueprint_version",
					Description: `The version of the vRA cloud template to request the deployment. Used only when ` + "`" + `blueprint_id` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "blueprint_content",
					Description: `vRA Cloud template content. Conflicts with ` + "`" + `blueprint_id` + "`" + ` and ` + "`" + `catalog_item_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `The id of the vRA catalog item to request the deployment. Conflicts with ` + "`" + `blueprint_id` + "`" + ` and ` + "`" + `blueprint_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog_item_version",
					Description: `The version of the vRA catalog item to request the deployment. Used only when ` + "`" + `catalog_item_id` + "`" + ` is provided.`,
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
					Name:        "description",
					Description: `A human-friendly description.`,
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
					Name:        "inputs",
					Description: `Inputs provided by the user while requesting / updating the deployment.`,
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
					Name:        "org_id",
					Description: `The ID of the organization this deployment belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The user this deployment belongs to.`,
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
					Name:        "project_id",
					Description: `The id of the project this deployment belongs to.`,
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
					Description: `A description of the resource.`,
				},
				resource.Attribute{
					Name:        "expense",
					Description: `Expense incurred for this resource.`,
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
					Description: `Deployment status. Supported values are: ` + "`" + `CREATE_SUCCESSFUL` + "`" + `, ` + "`" + `CREATE_INPROGRESS` + "`" + `, ` + "`" + `CREATE_FAILED` + "`" + `, ` + "`" + `UPDATE_SUCCESSFUL` + "`" + `, ` + "`" + `UPDATE_INPROGRESS` + "`" + `, ` + "`" + `UPDATE_FAILED` + "`" + `, ` + "`" + `DELETE_SUCCESSFUL` + "`" + `, ` + "`" + `DELETE_INPROGRESS` + "`" + `, ` + "`" + `DELETE_FAILED` + "`" + `, ` + "`" + `ACTION_SUCCESSFUL` + "`" + `, ` + "`" + `ACTION_INPROGRESS` + "`" + `, ` + "`" + `ACTION_FAILED` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "blueprint_d",
					Description: `The id of the vRA cloud template to request the deployment. Conflicts with ` + "`" + `catalog_item_id` + "`" + ` and ` + "`" + `blueprint_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "blueprint_version",
					Description: `The version of the vRA cloud template to request the deployment. Used only when ` + "`" + `blueprint_id` + "`" + ` is provided.`,
				},
				resource.Attribute{
					Name:        "blueprint_content",
					Description: `vRA Cloud template content. Conflicts with ` + "`" + `blueprint_id` + "`" + ` and ` + "`" + `catalog_item_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog_item_id",
					Description: `The id of the vRA catalog item to request the deployment. Conflicts with ` + "`" + `blueprint_id` + "`" + ` and ` + "`" + `blueprint_content` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "catalog_item_version",
					Description: `The version of the vRA catalog item to request the deployment. Used only when ` + "`" + `catalog_item_id` + "`" + ` is provided.`,
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
					Name:        "description",
					Description: `A human-friendly description.`,
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
					Name:        "inputs",
					Description: `Inputs provided by the user while requesting / updating the deployment.`,
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
					Name:        "org_id",
					Description: `The ID of the organization this deployment belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The user this deployment belongs to.`,
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
					Name:        "project_id",
					Description: `The id of the project this deployment belongs to.`,
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
					Description: `A description of the resource.`,
				},
				resource.Attribute{
					Name:        "expense",
					Description: `Expense incurred for this resource.`,
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
					Description: `Deployment status. Supported values are: ` + "`" + `CREATE_SUCCESSFUL` + "`" + `, ` + "`" + `CREATE_INPROGRESS` + "`" + `, ` + "`" + `CREATE_FAILED` + "`" + `, ` + "`" + `UPDATE_SUCCESSFUL` + "`" + `, ` + "`" + `UPDATE_INPROGRESS` + "`" + `, ` + "`" + `UPDATE_FAILED` + "`" + `, ` + "`" + `DELETE_SUCCESSFUL` + "`" + `, ` + "`" + `DELETE_INPROGRESS` + "`" + `, ` + "`" + `DELETE_FAILED` + "`" + `, ` + "`" + `ACTION_SUCCESSFUL` + "`" + `, ` + "`" + `ACTION_INPROGRESS` + "`" + `, ` + "`" + `ACTION_FAILED` + "`" + `.`,
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
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `(Optional) The id of the region for which this profile is defined as in vRealize Automation(vRA). ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
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
					Name:        "description",
					Description: `A human-friendly description.`,
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
			Type:             "vra_vra_machine",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_machine.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.`,
				},
				resource.Attribute{
					Name:        "cloud_account_ids",
					Description: `Set of ids of the cloud accounts this resource belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the base resource.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment id that is associated with this resource.`,
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
					Description: `The external zoneId of the resource.`,
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
					Name:        "power_state",
					Description: `Power state of machine.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project this resource belongs to.`,
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
					Name:        "address",
					Description: `Primary address allocated or in use by this machine. The actual type of the address depends on the adapter type. Typically it is either the public or the external IP address.`,
				},
				resource.Attribute{
					Name:        "cloud_account_ids",
					Description: `Set of ids of the cloud accounts this resource belongs to.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the base resource.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment id that is associated with this resource.`,
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
					Description: `The external zoneId of the resource.`,
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
					Name:        "power_state",
					Description: `Power state of machine.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project this resource belongs to.`,
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
			Type:             "vra_vra_network",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for vra_network.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A human-friendly name used as an identifier in APIs that support this option. ## Attribute Reference`,
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
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the base resource.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment id that is associated with this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
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
					Name:        "outbound_access",
					Description: `Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project this resource belongs to.`,
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
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the base resource.`,
				},
				resource.Attribute{
					Name:        "deployment_id",
					Description: `Deployment id that is associated with this resource.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
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
					Name:        "outbound_access",
					Description: `Flag to indicate if the network needs to have outbound access or not. Default is true. This field will be ignored if there is proper input for networkType customProperty`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The id of the project this resource belongs to.`,
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
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
				},
				resource.Attribute{
					Name:        "isolated_external_fabric_network_id",
					Description: `(Optional) The Id of the fabric network used for outbound access.`,
				},
				resource.Attribute{
					Name:        "isolated_network_domain_id",
					Description: `(Optional) The Id of the network domain used for creating isolated networks. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "fabric_network_ids",
					Description: `A list of fabric network Ids which are assigned to the network profile. example:[ "6543" ]`,
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
					Name:        "isolation_type",
					Description: `Specifies the isolation type e.g. none, subnet or security group`,
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
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The id of the region for which this profile is defined as in vRealize Automation(vRA).`,
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
					Name:        "custom_properties",
					Description: `Additional properties that may be used to extend the Network Profile object that is produced from this specification. For isolationType security group, datastoreId identifies the Compute Resource Edge datastore. computeCluster and resourcePoolId identify the Compute Resource Edge cluster. For isolationType subnet, distributedLogicalRouterStateLink identifies the on-demand network distributed local router. onDemandNetworkIPAssignmentType identifies the on-demand network IP range assignment type static, dynamic, or mixed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The external regionId of the resource.`,
				},
				resource.Attribute{
					Name:        "fabric_network_ids",
					Description: `A list of fabric network Ids which are assigned to the network profile. example:[ "6543" ]`,
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
					Name:        "isolation_type",
					Description: `Specifies the isolation type e.g. none, subnet or security group`,
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
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `Email of the user that owns the entity.`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `The id of the region for which this profile is defined as in vRealize Automation(vRA).`,
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
			ShortDescription: `Provides a data lookup for vra_project.`,
			Description: `
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
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
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
					Description: `(Optional) A human-friendly name used as an identifier in APIs that support this option.`,
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
					Name:        "zone_assignments",
					Description: `(Optional) List of configurations for zone assignment to a project.`,
				},
				resource.Attribute{
					Name:        "shared_resources",
					Description: `(Optional) The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "viewer",
					Description: `(Optional) List of viewer users associated with the project.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region data source.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_account_id",
					Description: `(Optional) The Cloud Account ID.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Search criteria to narrow down Images.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of the region to find.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The specific region associated with the cloud account.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the region from the associated vCenter. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
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
					Name:        "external_region_id",
					Description: `External regionId of region.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the given region within the cloud account.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when entity was created. Date and time format is ISO 8601 and UTC.`,
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
					Name:        "external_region_id",
					Description: `External regionId of region.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the given region within the cloud account.`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when entity was updated. Date and time format is ISO 8601 and UTC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region_enumeration",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region enumeration data source.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self signed certificate when connecting to vSphere. Example: false`,
				},
				resource.Attribute{
					Name:        "dcid",
					Description: `(Optional) ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the user used to authenticate with the cloud Account`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username to authenticate with the cloud account ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region_enumeration_aws",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region enumeration for AWS cloud account.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_key",
					Description: `(Optional) Aws Access key ID.`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Required) Aws Secret Access Key. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: us-east-2, ap-northeast-1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: us-east-2, ap-northeast-1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region_enumeration_azure",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region enumeration for Azure cloud account.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "application_id",
					Description: `(Required) Azure Client Application ID`,
				},
				resource.Attribute{
					Name:        "application_key",
					Description: `(Required) Azure Client Application Secret Key`,
				},
				resource.Attribute{
					Name:        "subscription_id",
					Description: `(Required) Azure Subscribtion ID`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Required) Azure Tenant ID ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: northamerica-northeast1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: northamerica-northeast1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region_enumeration_gcp",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region enumeration for GCP cloud account.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_email",
					Description: `(Required) Client E-mail ID.`,
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
					Description: `(Required) GCP Project ID. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: northamerica-northeast1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: northamerica-northeast1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region_enumeration_vmc",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region enumeration for VMC cloud account.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self signed certificate when connecting to vSphere. Example: false`,
				},
				resource.Attribute{
					Name:        "api_token",
					Description: `(Required) API Token for the cloud account endpoint.`,
				},
				resource.Attribute{
					Name:        "dc_id",
					Description: `(Optional) ID of a data collector vm deployed in the on premise infrastructure. Refer to the data-collector API to create or list data collectors.`,
				},
				resource.Attribute{
					Name:        "sddc_name",
					Description: `(Required) Identifier of the on-premise SDDC to be used by this cloud account.`,
				},
				resource.Attribute{
					Name:        "vcenter_hostname",
					Description: `(Required) The IP address or FQDN of the vCenter Server in the specified SDDC. The cloud proxy belongs on this vCenter.`,
				},
				resource.Attribute{
					Name:        "vcenter_password",
					Description: `(Required) Password for the user used to authenticate with the cloud Account`,
				},
				resource.Attribute{
					Name:        "vcenter_username",
					Description: `(Required) vCenter user name for the specified SDDC.The specified user requires CloudAdmin credentials. The user does not require CloudGlobalAdmin credentials. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: northamerica-northeast1`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of Region names to enable provisioning on. Example: northamerica-northeast1`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vra_vra_region_enumeration_vsphere",
			Category:         "Resources",
			ShortDescription: `Provides a data lookup for region enumeration for vSphere cloud account.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_self_signed_cert",
					Description: `(Optional) Accept self signed certificate when connecting to vSphere. Example: false`,
				},
				resource.Attribute{
					Name:        "dcid",
					Description: `(Required) ID of a data collector vm deployed in the on premise infrastructure. Example: d5316b00-f3b8-4895-9e9a-c4b98649c2ca`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host name for the cloud account endpoint. Example: dc1-lnd.mycompany.com`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the user used to authenticate with the cloud Account`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username to authenticate with the cloud account ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "regions",
					Description: `A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "regions",
					Description: `A set of datacenter managed object reference identifiers to enable provisioning on. Example: Datacenter:datacenter-2`,
				},
			},
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
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance. ## Attributes Reference`,
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
					Name:        "supports_encryption",
					Description: `Indicates whether this storage profile supports encryption or not.`,
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
					Name:        "supports_encryption",
					Description: `Indicates whether this storage profile supports encryption or not.`,
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
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
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
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
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
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
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
					Name:        "region_id",
					Description: `A link to the region that is associated with the storage profile.`,
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
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
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
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
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
					Name:        "region_id",
					Description: `A link to the region that is associated with the storage profile.`,
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
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 6801 and UTC.`,
				},
				resource.Attribute{
					Name:        "data_disk_caching",
					Description: `Indicates the caching mechanism for additional disk.`,
				},
				resource.Attribute{
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.`,
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
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "os_disk_caching",
					Description: `Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `Id of a storage account where in the disk is placed.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `Indicates whether this storage policy should support encryption or not.`,
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
					Name:        "region_id",
					Description: `A link to the region that is associated with the storage profile.`,
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
					Name:        "data_disk_caching",
					Description: `Indicates the caching mechanism for additional disk.`,
				},
				resource.Attribute{
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.`,
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
					Name:        "name",
					Description: `A human-friendly name used as an identifier in APIs that support this option.`,
				},
				resource.Attribute{
					Name:        "os_disk_caching",
					Description: `Indicates the caching mechanism for OS disk. Default policy for OS disks is Read/Write.`,
				},
				resource.Attribute{
					Name:        "storage_account_id",
					Description: `Id of a storage account where in the disk is placed.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `Indicates whether this storage policy should support encryption or not.`,
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
					Name:        "region_id",
					Description: `A link to the region that is associated with the storage profile.`,
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
					Name:        "filter",
					Description: `(Optional) Filter query string that is supported by vRA multi-cloud IaaS API. Example: regionId eq '<regionId>' and cloudAccountId eq '<cloudAccountId>'.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
				},
				resource.Attribute{
					Name:        "shares_level",
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
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_mode",
					Description: `Type of mode for the disk.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "limit_iops",
					Description: `The upper bound for the I/O operations per second allocated for each virtual disk.`,
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
					Name:        "provisioning_type",
					Description: `Type of provisioning policy for the disk.`,
				},
				resource.Attribute{
					Name:        "shares",
					Description: `A specific number of shares assigned to each virtual machine.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `Indicates whether this storage policy should support encryption or not.`,
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
					Name:        "default_item",
					Description: `Indicates if this storage profile is a default profile.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "disk_mode",
					Description: `Type of mode for the disk.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `Indicates the performance tier for the storage type. Premium disks are SSD backed and Standard disks are HDD backed.`,
				},
				resource.Attribute{
					Name:        "external_region_id",
					Description: `The id of the region as seen in the cloud provider for which this profile is defined.`,
				},
				resource.Attribute{
					Name:        "limit_iops",
					Description: `The upper bound for the I/O operations per second allocated for each virtual disk.`,
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
					Name:        "provisioning_type",
					Description: `Type of provisioning policy for the disk.`,
				},
				resource.Attribute{
					Name:        "shares",
					Description: `A specific number of shares assigned to each virtual machine.`,
				},
				resource.Attribute{
					Name:        "supports_encryption",
					Description: `Indicates whether this storage policy should support encryption or not.`,
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
			ShortDescription: `Provides a data lookup for vra_zone.`,
			Description: `
`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the image profile instance.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A human-friendly name used as an identifier in APIs that support this option. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `A list of key value pair of properties that will be used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `The folder relative path to the datacenter where resources are deployed to. (only applicable for vSphere cloud zones)`,
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
					Name:        "placement_policy",
					Description: `The id of the region for which this zone is defined`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `A link to the region that is associated with the storage profile.`,
				},
				resource.Attribute{
					Name:        "shared_resources",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "tags_to_match",
					Description: `A set of tag keys and optional values for compute resource filtering. example:[ { "key" : "compliance", "value": "pci" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "created_at",
					Description: `Date when the entity was created. The date is in ISO 8601 and UTC.`,
				},
				resource.Attribute{
					Name:        "custom_properties",
					Description: `A list of key value pair of properties that will be used.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A human-friendly description.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `The folder relative path to the datacenter where resources are deployed to. (only applicable for vSphere cloud zones)`,
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
					Name:        "placement_policy",
					Description: `The id of the region for which this zone is defined`,
				},
				resource.Attribute{
					Name:        "region_id",
					Description: `A link to the region that is associated with the storage profile.`,
				},
				resource.Attribute{
					Name:        "shared_resources",
					Description: `The id of the organization this entity belongs to.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A set of tag keys and optional values that were set on this Network Profile. example:[ { "key" : "ownedBy", "value": "Rainpole" } ]`,
				},
				resource.Attribute{
					Name:        "tags_to_match",
					Description: `A set of tag keys and optional values for compute resource filtering. example:[ { "key" : "compliance", "value": "pci" } ]`,
				},
				resource.Attribute{
					Name:        "updated_at",
					Description: `Date when the entity was last updated. The date is ISO 8601 and UTC.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"vra_cloud_account_aws":              0,
		"vra_cloud_account_azure":            1,
		"vra_cloud_account_gcp":              2,
		"vra_cloud_account_nsxt":             3,
		"vra_cloud_account_nsxv":             4,
		"vra_cloud_account_vmc":              5,
		"vra_cloud_account_vsphere":          6,
		"vra_fabric_datastore_vsphere":       7,
		"vra_fabric_network":                 8,
		"vra_fabric_storage_account_azure":   9,
		"vra_fabric_storage_policy_vsphere":  10,
		"vra_image":                          11,
		"vra_network_domain":                 12,
		"vra_region":                         13,
		"vra_security_group":                 14,
		"vra_vra_block_device":               15,
		"vra_vra_block_device_snapshot":      16,
		"vra_vra_blueprint":                  17,
		"vra_vra_blueprint_version":          18,
		"vra_vra_catalog_item":               19,
		"vra_vra_catalog_source_blueprint":   20,
		"vra_vra_catalog_source_entitlement": 21,
		"vra_vra_data_collector":             22,
		"vra_vra_deployment":                 23,
		"vra_vra_image_profile":              24,
		"vra_vra_machine":                    25,
		"vra_vra_network":                    26,
		"vra_vra_network_profile":            27,
		"vra_vra_project":                    28,
		"vra_vra_region":                     29,
		"vra_vra_region_enumeration":         30,
		"vra_vra_region_enumeration_aws":     31,
		"vra_vra_region_enumeration_azure":   32,
		"vra_vra_region_enumeration_gcp":     33,
		"vra_vra_region_enumeration_vmc":     34,
		"vra_vra_region_enumeration_vsphere": 35,
		"vra_vra_storage_profile":            36,
		"vra_vra_storage_profile_aws":        37,
		"vra_vra_storage_profile_azure":      38,
		"vra_vra_storage_profile_vsphere":    39,
		"vra_vra_zone":                       40,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
