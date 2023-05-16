package vcd

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog",
			Category:         "Data Sources",
			ShortDescription: `Provides a catalog data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Catalog name (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Catalog description.`,
				},
				resource.Attribute{
					Name:        "publish_enabled",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "cache_enabled",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "preserve_identity_information",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated;`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to this Catalog. See [Metadata](#metadata) section for details.`,
				},
				resource.Attribute{
					Name:        "catalog_version",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "number_of_vapp_templates",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "number_of_media",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vapp_template_list",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "media_item_list",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "is_local",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "publish_subscription_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "publish_subscription_url",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `. ## Filter arguments (Supported in provider`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Catalog description.`,
				},
				resource.Attribute{
					Name:        "publish_enabled",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "cache_enabled",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "preserve_identity_information",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated;`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to this Catalog. See [Metadata](#metadata) section for details.`,
				},
				resource.Attribute{
					Name:        "catalog_version",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "number_of_vapp_templates",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "number_of_media",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vapp_template_list",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "media_item_list",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "is_local",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "publish_subscription_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "publish_subscription_url",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `. ## Filter arguments (Supported in provider`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_item",
			Category:         "Data Sources",
			ShortDescription: `Provides a catalog item data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Required) Catalog name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Catalog Item name (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Catalog item description.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Key value map of metadata assigned to the associated vApp template.`,
				},
				resource.Attribute{
					Name:        "catalog_item_metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key value map of metadata assigned to the catalog item.`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to the catalog item. See [Metadata](#metadata) section for details. <a id="metadata"></a> ## Metadata The ` + "`" + `metadata_entry` + "`" + ` (`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `. ## Filter arguments (Supported in provider`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) is an expression starting with an operator (` + "`" + `>` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `==` + "`" + `), followed by a date, with optional spaces in between. For example: ` + "`" + `> 2020-02-01 12:35:00.523Z` + "`" + ` The filter recognizes several formats, but one of ` + "`" + `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` + "`" + ` or ` + "`" + `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]` + "`" + ` is recommended. Comparison with equality operator (` + "`" + `==` + "`" + `) need to define the date to the microseconds.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the latest item among the ones matching other parameters. If no other parameters are set, it retrieves the newest item.`,
				},
				resource.Attribute{
					Name:        "earliest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the earliest item among the ones matching other parameters. If no other parameters are set, it retrieves the oldest item.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Catalog item description.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Key value map of metadata assigned to the associated vApp template.`,
				},
				resource.Attribute{
					Name:        "catalog_item_metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key value map of metadata assigned to the catalog item.`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to the catalog item. See [Metadata](#metadata) section for details. <a id="metadata"></a> ## Metadata The ` + "`" + `metadata_entry` + "`" + ` (`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `. ## Filter arguments (Supported in provider`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) is an expression starting with an operator (` + "`" + `>` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `==` + "`" + `), followed by a date, with optional spaces in between. For example: ` + "`" + `> 2020-02-01 12:35:00.523Z` + "`" + ` The filter recognizes several formats, but one of ` + "`" + `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` + "`" + ` or ` + "`" + `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]` + "`" + ` is recommended. Comparison with equality operator (` + "`" + `==` + "`" + `) need to define the date to the microseconds.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the latest item among the ones matching other parameters. If no other parameters are set, it retrieves the newest item.`,
				},
				resource.Attribute{
					Name:        "earliest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the earliest item among the ones matching other parameters. If no other parameters are set, it retrieves the oldest item.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_media",
			Category:         "Data Sources",
			ShortDescription: `Provides a catalog media data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "catalog",
					Description: `(Optional; Deprecated) The name of the catalog to which media file belongs. It's mandatory if ` + "`" + `catalog_id` + "`" + ` is not used.`,
				},
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Media name in catalog (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) is an expression starting with an operator (` + "`" + `>` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `==` + "`" + `), followed by a date, with optional spaces in between. For example: ` + "`" + `> 2020-02-01 12:35:00.523Z` + "`" + ` The filter recognizes several formats, but one of ` + "`" + `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` + "`" + ` or ` + "`" + `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]` + "`" + ` is recommended. Comparison with equality operator (` + "`" + `==` + "`" + `) need to define the date to the microseconds.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the latest item among the ones matching other parameters. If no other parameters are set, it retrieves the newest item.`,
				},
				resource.Attribute{
					Name:        "earliest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the earliest item among the ones matching other parameters. If no other parameters are set, it retrieves the oldest item.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_catalog_vapp_template",
			Category:         "Data Sources",
			ShortDescription: `Provides a vApp Template data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "catalog_id",
					Description: `(Required) ID of the catalog containing the vApp Template. Can't be used if a specific VDC identifier is set (` + "`" + `vdc_id` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Required) ID of the VDC to which the vApp Template belongs. Can't be used if a specific Catalog is set (` + "`" + `catalog_id` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) vApp Template name (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `vApp Template description`,
				},
				resource.Attribute{
					Name:        "vm_names",
					Description: `Set of VM names within the vApp template`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key/value map of metadata for the associated vApp template.`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to this vApp Template. See [Metadata](/providers/vmware/vcd/latest/docs/resources/catalog_vapp_template#metadata) section for details. ## Filter arguments`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) is an expression starting with an operator (` + "`" + `>` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `==` + "`" + `), followed by a date, with optional spaces in between. For example: ` + "`" + `> 2020-02-01 12:35:00.523Z` + "`" + ` The filter recognizes several formats, but one of ` + "`" + `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` + "`" + ` or ` + "`" + `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]` + "`" + ` is recommended. Comparison with equality operator (` + "`" + `==` + "`" + `) need to define the date to the microseconds.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the latest item among the ones matching other parameters. If no other parameters are set, it retrieves the newest item.`,
				},
				resource.Attribute{
					Name:        "earliest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the earliest item among the ones matching other parameters. If no other parameters are set, it retrieves the oldest item.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `vApp Template description`,
				},
				resource.Attribute{
					Name:        "vm_names",
					Description: `Set of VM names within the vApp template`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key/value map of metadata for the associated vApp template.`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to this vApp Template. See [Metadata](/providers/vmware/vcd/latest/docs/resources/catalog_vapp_template#metadata) section for details. ## Filter arguments`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `(Optional) is an expression starting with an operator (` + "`" + `>` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `==` + "`" + `), followed by a date, with optional spaces in between. For example: ` + "`" + `> 2020-02-01 12:35:00.523Z` + "`" + ` The filter recognizes several formats, but one of ` + "`" + `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` + "`" + ` or ` + "`" + `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]` + "`" + ` is recommended. Comparison with equality operator (` + "`" + `==` + "`" + `) need to define the date to the microseconds.`,
				},
				resource.Attribute{
					Name:        "latest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the latest item among the ones matching other parameters. If no other parameters are set, it retrieves the newest item.`,
				},
				resource.Attribute{
					Name:        "earliest",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, retrieve the earliest item among the ones matching other parameters. If no other parameters are set, it retrieves the oldest item.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_library_certificate",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read certificate in System or Org library.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alias",
					Description: `(Optional) - alias (name) of certificate`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) - ID of certificate ` + "`" + `alias` + "`" + ` or ` + "`" + `id` + "`" + ` is required field. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_library_certificate` + "`" + `](/providers/vmware/vcd/latest/docs/resources/certificate_library) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_edgegateway",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX-V edge gateway data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the edge gateway (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC that owns the edge gateway. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network",
			Category:         "Data Sources",
			ShortDescription: `Provides an external network data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) external network name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `A list of IP scopes for the network. See [IP Scope](/providers/vmware/vcd/latest/docs/resources/external_network#ipscope) for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system. See [vSphere Network](/providers/vmware/vcd/latest/docs/resources/external_network#vspherenetwork) for details.`,
				},
				resource.Attribute{
					Name:        "retain_net_info_across_deployments",
					Description: `Specifies whether the network resources such as IP/MAC of router will be retained across deployments.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Network friendly description`,
				},
				resource.Attribute{
					Name:        "ip_scope",
					Description: `A list of IP scopes for the network. See [IP Scope](/providers/vmware/vcd/latest/docs/resources/external_network#ipscope) for details.`,
				},
				resource.Attribute{
					Name:        "vsphere_network",
					Description: `A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system. See [vSphere Network](/providers/vmware/vcd/latest/docs/resources/external_network#vspherenetwork) for details.`,
				},
				resource.Attribute{
					Name:        "retain_net_info_across_deployments",
					Description: `Specifies whether the network resources such as IP/MAC of router will be retained across deployments.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_external_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director External Network data source (version 2). New version of this data source uses new VCD API and is capable of creating NSX-T backed external networks as well as port group backed ones.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) external network name ## Attribute Reference All properties defined in [vcd_external_network_v2](/providers/vmware/vcd/latest/docs/resources/external_network_v2) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_global_role",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director global role data source . This can be used to read global roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the global role. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the global role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the global role to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `List of tenants to which this global role gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this global role is read-only ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how global roles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the global role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `List of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the global role to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `List of tenants to which this global role gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this global role is read-only ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how global roles and rights work together.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_independent_disk",
			Category:         "Data Sources",
			ShortDescription: `Provides a independent disk data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Disk id or name is required. If both provided - Id is used. Id can be found by using import function [Listing independent disk IDs](/providers/vmware/vcd/latest/docs/resources/independent_disk#listing-independent-disk-ids)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Disk name.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_app_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer application profile data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Application profile name for identifying the exact application profile ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_app_profile` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_app_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer application rule data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Application rule name for identifying the exact application rule ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_app_rule` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_server_pool",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer server pool data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the server pool is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Server Pool name for identifying the exact server pool ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_server_pool` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_service_monitor",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer service monitor data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the service monitor is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Service Monitor name for identifying the exact service monitor ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_service_monitor` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_lb_virtual_server",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway load balancer virtual server data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which the virtual server is defined`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for identifying the exact virtual server ## Attribute Reference All the attributes defined in ` + "`" + `vcd_lb_virtual_server` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_direct",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC Network attached to an external one. This can be used to reference internal networks for vApps to connect.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "external_network",
					Description: `The name of the external network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Defines if this network is shared between multiple vDCs in the vOrg. ## Filter arguments (Supported in provider`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "external_network",
					Description: `The name of the external network.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Defines if this network is shared between multiple vDCs in the vOrg. ## Filter arguments (Supported in provider`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC isolated Network. This can be used to reference internal networks for vApps to connect.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_isolated_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC isolated Network data source to read data or reference existing network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `(Optional) VDC or VDC Group ID. Always takes precedence over ` + "`" + `vdc` + "`" + ` fields (in resource and inherited from provider configuration)`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Deprecated; Optional) The name of VDC to use.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters.`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC routed Network. This can be used to reference internal networks for vApps to connect.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Optional) One or more parameters that will match metadata contents. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_network_routed_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC routed Network data source to read data or reference existing network (backed by NSX-T or NSX-V).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Deprecated; Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Parent VDC or VDC Group ID. All attributes defined in [routed network v2 resource](/providers/vmware/vcd/latest/docs/resources/network_routed_v2#attribute-reference) are supported. ## Filter arguments`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_cloud",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T ALB Clouds for Providers. An NSX-T Cloud is a service provider-level construct that consists of an NSX-T Manager and an NSX-T Data Center transport zone.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of NSX-T ALB Cloud ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_alb_cloud` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_alb_cloud) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_controller",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T ALB Controller for Providers. It helps to integrate VMware Cloud Director with NSX-T Advanced Load Balancer deployment. Controller instances are registered with VMware Cloud Director instance. Controller instances serve as a central control plane for the load-balancing services provided by NSX-T Advanced Load Balancer.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing NSX-T ALB Controller. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_alb_controller` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_edgegateway_service_engine_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a datasource to read NSX-T ALB Service Engine Group assignment to NSX-T Edge Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Optional) An ID of NSX-T Edge Gateway. Can be looked up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source`,
				},
				resource.Attribute{
					Name:        "service_engine_group_id",
					Description: `(Required) An ID of NSX-T Service Engine Group. Can be looked up using [vcd_nsxt_alb_service_engine_group](/providers/vmware/vcd/latest/docs/data-sources/nsxt_alb_service_engine_group) data source.`,
				},
				resource.Attribute{
					Name:        "service_engine_group_name",
					Description: `(Optional) A Name of NSX-T Service Engine Group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_importable_cloud",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to reference existing NSX-T ALB Importable Clouds. An NSX-T Importable Cloud is a reference to a Cloud configured in ALB Controller.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of NSX-T ALB Importable Cloud`,
				},
				resource.Attribute{
					Name:        "controller_id",
					Description: `(Required) - NSX-T ALB Controller ID ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "already_imported",
					Description: `boolean value which displays if the ALB Importable Cloud is already consumed`,
				},
				resource.Attribute{
					Name:        "network_pool_id",
					Description: `backing network pool ID`,
				},
				resource.Attribute{
					Name:        "network_pool_name",
					Description: `backing network pool ID`,
				},
				resource.Attribute{
					Name:        "transport_zone_name",
					Description: `backing transport zone name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "already_imported",
					Description: `boolean value which displays if the ALB Importable Cloud is already consumed`,
				},
				resource.Attribute{
					Name:        "network_pool_id",
					Description: `backing network pool ID`,
				},
				resource.Attribute{
					Name:        "network_pool_name",
					Description: `backing network pool ID`,
				},
				resource.Attribute{
					Name:        "transport_zone_name",
					Description: `backing transport zone name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_pool",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T ALB Pools for particular NSX-T Edge Gateway. Pools maintain the list of servers assigned to them and perform health monitoring, load balancing, persistence. A pool may only be used or referenced by only one virtual service at a time.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) An ID of NSX-T Edge Gateway. Can be looked up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of existing NSX-T ALB Pool. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_alb_pool` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_alb_pool) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_service_engine_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T ALB Service Engine Groups. A Service Engine Group is an isolation domain that also defines shared service engine properties, such as size, network access, and failover. Resources in a service engine group can be used for different virtual services, depending on your tenant needs. These resources cannot be shared between different service engine groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of Service Engine Group.`,
				},
				resource.Attribute{
					Name:        "sync_on_refresh",
					Description: `(Optional) - A special argument that is not passed to VCD, but alters behaviour of this resource so that it performs a Sync operation on every Terraform refresh.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_settings",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T ALB General Settings for particular NSX-T Edge Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) An ID of NSX-T Edge Gateway. Can be looked up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_alb_settings` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_alb_virtual_service",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T ALB Virtual services for particular NSX-T Edge Gateway. A virtual service advertises an IP address and ports to the external world and listens for client traffic. When a virtual service receives traffic, it directs it to members in ALB Pool.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) An ID of NSX-T Edge Gateway. Can be looked up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of ALB Virtual Service ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_alb_virtual_service` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_alb_virtual_service) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_app_port_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Application Port Profiles. Application Port Profiles include a combination of a protocol and a port, or a group of ports, that is used for Firewall and NAT services on the Edge Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Deprecated; Optional) The name of VDC to use, optional if defined at provider level. Deprecated and replaced by ` + "`" + `context_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "context_id",
					Description: `(Optional) ID of NSX-T Manager, VDC or VDC Group. Replaces deprecated field ` + "`" + `vdc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing Security Group.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Required) - ` + "`" + `SYSTEM` + "`" + `, ` + "`" + `PROVIDER` + "`" + `, or ` + "`" + `TENANT` + "`" + `. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_app_port_profile` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_app_port_profile) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_distributed_firewall",
			Category:         "Data Sources",
			ShortDescription: `The Distributed Firewall data source reads all defined rules for a particular VDC Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization in which Distributed Firewall is located. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc_group_id",
					Description: `(Required) The ID of a VDC Group ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_distributed_firewall` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edge_cluster",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available NSX-T Edge Clusters.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which edge cluster belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional, Deprecated) The name of VDC that owns the edge cluster. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "vdc_group_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "provider_vdc_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T Edge Cluster name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Edge Cluster description in NSX-T manager.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `Number of nodes in Edge Cluster.`,
				},
				resource.Attribute{
					Name:        "node_type",
					Description: `Type of nodes in Edge Cluster.`,
				},
				resource.Attribute{
					Name:        "deployment_type",
					Description: `Deployment type of Edge Cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director NSX-T edge gateway data source. This can be used to read NSX-T edge gateway configurations.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the NSX-T Edge Gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional)`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `(Optional,`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T Edge Gateway name. ## Attribute reference All properties defined in [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway_bgp_configuration",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read BGP configuration on NSX-T Edge Gateway that has a dedicated Tier-0 Gateway or VRF.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) An ID of NSX-T Edge Gateway. Can be lookup up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_edgegateway_bgp_configuration` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_configuration) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway_bgp_ip_prefix_list",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to manage NSX-T Edge Gateway BGP IP Prefix Lists. IP prefix lists can contain single or multiple IP addresses and can be used to assign BGP neighbors with access permissions for route advertisement.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) An ID of NSX-T Edge Gateway. Can be lookup up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name of existing BGP IP Prefix List in specified Edge Gateway ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_edgegateway_bgp_ip_prefix_list` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_ip_prefix_list) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway_bgp_neighbor",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Edge Gateway BGP Neighbors and their configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the edge gateway belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) An ID of NSX-T Edge Gateway. Can be looked up using [vcd_nsxt_edgegateway](/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) An IP Address (IPv4 or IPv6) of existing BGP Neighbor in specified Edge Gateway ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_edgegateway_bgp_neighbor` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_bgp_neighbor) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway_qos_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Edge Gateway QoS profiles, which can be used to modify NSX-T Edge Gateway Rate Limiting (QoS) configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nsxt_manager_id",
					Description: `(Required) NSX-T Manager where the QoS profile is defined in (can be looked up using ` + "`" + `vcd_nsxt_manager` + "`" + ` data source)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) QoS Profile Name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of QoS Profile`,
				},
				resource.Attribute{
					Name:        "committed_bandwidth",
					Description: `Committed bandwith specificd in Mb/s`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `Burst size defined in bytes`,
				},
				resource.Attribute{
					Name:        "excess_action",
					Description: `Excess action defines action on traffic exceeding bandwidth`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Description of QoS Profile`,
				},
				resource.Attribute{
					Name:        "committed_bandwidth",
					Description: `Committed bandwith specificd in Mb/s`,
				},
				resource.Attribute{
					Name:        "burst_size",
					Description: `Burst size defined in bytes`,
				},
				resource.Attribute{
					Name:        "excess_action",
					Description: `Excess action defines action on traffic exceeding bandwidth`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_edgegateway_rate_limiting",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Edge Gateway Rate Limiting (QoS) configuration.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Required) Org in which the NSX-T Edge Gateway is located`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) NSX-T Edge Gateway ID ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_edgegateway_rate_limiting` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway_rate_limiting) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_firewall",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Firewall configuration of an Edge Gateway. Firewalls allow user to control the incoming and outgoing network traffic to and from an NSX-T Data Center Edge Gateway.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using ` + "`" + `vcd_nsxt_edgegateway` + "`" + ` data source ## Attribute reference All properties defined in [vcd_nsxt_firewall](/providers/vmware/vcd/latest/docs/resources/nsxt_firewall) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_ip_set",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T IP Set. IP Sets are groups of objects to which the firewall rules apply. Combining multiple objects into IP Sets helps reduce the total number of firewall rules to be created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing IP Set. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Parent VDC or VDC Group ID. All the arguments and attributes defined in [` + "`" + `vcd_nsxt_ip_set` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_ip_set) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "owner_id",
					Description: `Parent VDC or VDC Group ID. All the arguments and attributes defined in [` + "`" + `vcd_nsxt_ip_set` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_ip_set) resource are available.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_ipsec_vpn_tunnel",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers, or VPN gateways that support IPSec.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using ` + "`" + `vcd_nsxt_edgegateway` + "`" + ` data source`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of existing IPsec VPN Tunnel ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_ipsec_vpn_tunnel` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_ipsec_vpn_tunnel) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_manager",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available NSX-T manager.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T manager name ## Attribute reference Only ID is set to be able and reference in other resources or data sources.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_nat_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T NAT rules. Source NAT (SNAT) rules change the source IP address from a private to a public IP address. Destination NAT (DNAT) rules change the destination IP address from a public to a private IP address.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Name of existing NAT Rule. -> Name uniqueness is not enforced in NSX-T NAT rules, but for this data source to work properly names should be unique so that they can be distinguished. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_nat_rule` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_nat_rule) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_dynamic_security_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for NSX-T Network Context Profile lookup to later be used in Distributed Firewall.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "context_id",
					Description: `(Required) Context ID specifies the context for Network Context Profile look up. This ID can be one of ` + "`" + `VDC Group ID` + "`" + ` (data source ` + "`" + `vcd_vdc_group` + "`" + `), ` + "`" + `Org VDC ID` + "`" + ` (data source ` + "`" + `vcd_org_vdc` + "`" + `), or ` + "`" + `NSX-T Manager ID` + "`" + ` (data source ` + "`" + `vcd_nsxt_manager` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `(Optional) Can be one of ` + "`" + `SYSTEM` + "`" + `, ` + "`" + `TENANT` + "`" + `, ` + "`" + `PROVIDER` + "`" + `. (default ` + "`" + `SYSTEM` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of Network Context Profile ## Attribute Reference ` + "`" + `id` + "`" + ` - can be used in ` + "`" + `vcd_nsxt_distributed_firewall` + "`" + ` resource field ` + "`" + `network_context_profile_ids` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_dhcp",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read DHCP pools for NSX-T Org VDC networks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "org_network_id",
					Description: `(Required) ID of parent Org VDC Routed network ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxt_network_dhcp` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_network_dhcp) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_dhcp_binding",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read NSX-T Org VDC network DHCP bindings.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization. Optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "org_network_id",
					Description: `(Required) The ID of an Org VDC network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name of DHCP binding ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_nsxt_network_dhcp_binding` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_network_dhcp_binding) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_network_imported",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Org VDC NSX-T Imported Network data source to read data or reference existing network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `(Optional) VDC or VDC Group ID. Always takes precedence over ` + "`" + `vdc` + "`" + ` fields (in resource and inherited from provider configuration)`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Deprecated; Optional) The name of VDC to use.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the network (optional when ` + "`" + `filter` + "`" + ` is used)`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters ## Attribute reference All attributes defined in [imported network resource](/providers/vmware/vcd/latest/docs/resources/nsxt_network_imported#attribute-reference) are supported. ## Filter arguments`,
				},
				resource.Attribute{
					Name:        "name_regex",
					Description: `(Optional) matches the name using a regular expression.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `(Optional) matches the IP of the resource using a regular expression. See [Filters reference](/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_security_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to access NSX-T Security Group configuration. Security Groups are groups of data center group networks to which distributed firewall rules apply. Grouping networks helps you to reduce the total number of distributed firewall rules to be created.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Deprecated; Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway_id",
					Description: `(Required) The ID of the Edge Gateway (NSX-T only). Can be looked up using`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) - Unique name of existing Security Group. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `Parent VDC or VDC Group ID. All the arguments and attributes defined in [` + "`" + `vcd_nsxt_security_group` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_security_group) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "owner_id",
					Description: `Parent VDC or VDC Group ID. All the arguments and attributes defined in [` + "`" + `vcd_nsxt_security_group` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxt_security_group) resource are available.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxt_tier0_router",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available NSX-T Tier-0 routers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) NSX-T Tier-0 router name.`,
				},
				resource.Attribute{
					Name:        "nsxt_manager_id",
					Description: `(Required) NSX-T manager should be referenced. ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "is_assigned",
					Description: `Boolean value reflecting if Tier-0 router is already consumed by external network.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_application",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director data source for reading NSX-V distributed firewall applications`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Required) The ID of VDC to use`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The identifier of the application`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used by the application`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `The ports used by the application. Could be a number, a list of numbers, or a range`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `The source port used by this application. Not all applications provide one`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `The destination port used by this application. Not all applications provide one`,
				},
				resource.Attribute{
					Name:        "app_guid",
					Description: `The application identifier, when available`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The identifier of the application`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol used by the application`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `The ports used by the application. Could be a number, a list of numbers, or a range`,
				},
				resource.Attribute{
					Name:        "source_port",
					Description: `The source port used by this application. Not all applications provide one`,
				},
				resource.Attribute{
					Name:        "destination_port",
					Description: `The destination port used by this application. Not all applications provide one`,
				},
				resource.Attribute{
					Name:        "app_guid",
					Description: `The application identifier, when available`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_application_finder",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director data source for searching NSX-V distributed firewall applications and application groups`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Required) The ID of VDC to use`,
				},
				resource.Attribute{
					Name:        "search_expression",
					Description: `(Required) The regular expression that will be used to search the applications. See [Search Expressions](#search-expressions) below`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) What kind of application we seek. One of ` + "`" + `application` + "`" + `, ` + "`" + `application_group` + "`" + ``,
				},
				resource.Attribute{
					Name:        "objects",
					Description: `A list of objects found by the search expression. Each one contains the following properties:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the object`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the object (` + "`" + `Application` + "`" + ` or ` + "`" + `ApplicationGroup` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The identifier of the object ## Search expressions To search for an application or application group, we can use simple or complex [regular expressions](https://en.wikipedia.org/wiki/Regular_expression). The expressions in this data source follow the [PCRE](https://en.wikipedia.org/wiki/Perl_Compatible_Regular_Expressions) standard. A`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "objects",
					Description: `A list of objects found by the search expression. Each one contains the following properties:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the object`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `the type of the object (` + "`" + `Application` + "`" + ` or ` + "`" + `ApplicationGroup` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The identifier of the object ## Search expressions To search for an application or application group, we can use simple or complex [regular expressions](https://en.wikipedia.org/wiki/Regular_expression). The expressions in this data source follow the [PCRE](https://en.wikipedia.org/wiki/Perl_Compatible_Regular_Expressions) standard. A`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_application_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director data source for reading NSX-V Distributed Firewall application groups`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Required) The ID of VDC to use`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the application group ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The identifier of the application groups`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `The list of the applications belonging to this group. For each one we get the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the application`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The identifier of the application`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The identifier of the application groups`,
				},
				resource.Attribute{
					Name:        "applications",
					Description: `The list of the applications belonging to this group. For each one we get the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the application`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `The identifier of the application`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_dhcp_relay",
			Category:         "Data Sources",
			ShortDescription: `Provides an NSX edge gateway DHCP relay configuration data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which DHCP relay is to be configured. ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_dhcp_relay` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxv_dhcp_relay) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_distributed_firewall",
			Category:         "Data Sources",
			ShortDescription: `The NSX-V Distributed Firewall data source reads all defined rules for a particular VDC`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Required) The ID of VDC to manage the Distributed Firewall in. Can be looked up using a ` + "`" + `vcd_org_vdc` + "`" + ` data source ## Attributes reference All the arguments and attributes defined for the ` + "`" + `vcd_nsxv_distributed_firewall` + "`" + ` resource are reported as attributes for this data source.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_dnat",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director DNAT data source for advanced edge gateways (NSX-V). This can be used to read existing rule by ID and use its attributes in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the DNAT rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) ID of DNAT rule as shown in the UI. ## Attribute Reference All the attributes defined in ` + "`" + `vcd_nsxv_dnat` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_firewall_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director firewall rule data source for advanced edge gateways (NSX-V). This can be used to read existing rules by ID and use its attributes in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the DNAT rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) ID of firewall rule (not UI number). See more information about firewall rule ID in ` + "`" + `vcd_nsxv_firewall_rule` + "`" + ` [import section](/providers/vmware/vcd/latest/docs/resources/nsxv_firewall_rule#listing-real-firewall-rule-ids). ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_firewall_rule` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxv_firewall_rule) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_ip_set",
			Category:         "Data Sources",
			ShortDescription: `Provides an IP set data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) IP set name for identifying the exact IP set ## Attribute Reference All the attributes defined in [` + "`" + `vcd_nsxv_ip_set` + "`" + `](/providers/vmware/vcd/latest/docs/resources/nsxv_ip_set) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_nsxv_snat",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to read existing rule by ID and use its attributes in other resources.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "edge_gateway",
					Description: `(Required) The name of the edge gateway on which to apply the SNAT rule.`,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `(Required) ID of SNAT rule as shown in the UI. ## Attribute Reference All the attributes defined in ` + "`" + `vcd_nsxv_snat` + "`" + ` resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org",
			Category:         "Data Sources",
			ShortDescription: `Provides an organization data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Org name ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `Org full name`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True if this organization is enabled (allows login and all other operations).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Org description.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "can_publish_catalogs",
					Description: `True if this organization is allowed to share catalogs.`,
				},
				resource.Attribute{
					Name:        "can_publish_external_catalogs",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "can_subscribe_external_catalogs",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "delay_after_power_on_seconds",
					Description: `Specifies this organization's default for virtual machine boot delay after power on.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated;`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to the organization. See [Metadata](#metadata) section for details.`,
				},
				resource.Attribute{
					Name:        "vapp_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vapp_template_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "maximum_runtime_lease_in_sec",
					Description: `How long vApps can run before they are automatically stopped (in seconds)`,
				},
				resource.Attribute{
					Name:        "power_off_on_runtime_lease_expiration",
					Description: `When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long stopped vApps are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted. <a id="vapp-template-lease"></a> ## vApp Template Lease The ` + "`" + `vapp_template_lease` + "`" + ` section contains lease parameters for vApp templates created in the current organization, as defined below:`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long vApp templates are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted <a id="metadata"></a> ## Metadata The ` + "`" + `metadata_entry` + "`" + ` (`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "full_name",
					Description: `Org full name`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True if this organization is enabled (allows login and all other operations).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Org description.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.`,
				},
				resource.Attribute{
					Name:        "can_publish_catalogs",
					Description: `True if this organization is allowed to share catalogs.`,
				},
				resource.Attribute{
					Name:        "can_publish_external_catalogs",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "can_subscribe_external_catalogs",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "delay_after_power_on_seconds",
					Description: `Specifies this organization's default for virtual machine boot delay after power on.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated;`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to the organization. See [Metadata](#metadata) section for details.`,
				},
				resource.Attribute{
					Name:        "vapp_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "vapp_template_lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "maximum_runtime_lease_in_sec",
					Description: `How long vApps can run before they are automatically stopped (in seconds)`,
				},
				resource.Attribute{
					Name:        "power_off_on_runtime_lease_expiration",
					Description: `When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long stopped vApps are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted. <a id="vapp-template-lease"></a> ## vApp Template Lease The ` + "`" + `vapp_template_lease` + "`" + ` section contains lease parameters for vApp templates created in the current organization, as defined below:`,
				},
				resource.Attribute{
					Name:        "maximum_storage_lease_in_sec",
					Description: `How long vApp templates are available before being automatically cleaned up (in seconds)`,
				},
				resource.Attribute{
					Name:        "delete_on_storage_lease_expiration",
					Description: `If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted <a id="metadata"></a> ## Metadata The ` + "`" + `metadata_entry` + "`" + ` (`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for VMware Cloud Director Organization Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the VDC belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the group. ## Attribute reference All attributes defined in [org_group](/providers/vmware/vcd/latest/docs/resources/org_group#attribute-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_ldap",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read LDAP configuration for an organization.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org_name",
					Description: `(Required) - Name of the organization containing the LDAP settings ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_org_ldap` + "`" + `](/providers/vmware/vcd/latest/docs/resources/org_ldap) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_user",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Organization user data source. This can be used to read organization users.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to which the user belongs. Optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the user. Required if ` + "`" + `user_id` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `(Optional) The ID of the user. Required if ` + "`" + `name` + "`" + ` is not set. ## Attribute reference ~>`,
				},
				resource.Attribute{
					Name:        "provider_type",
					Description: `Identity provider type for this user.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `The role of the user.`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `The full name of the user.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of the user.`,
				},
				resource.Attribute{
					Name:        "telephone",
					Description: `The Org User telephone number.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `The Org User email address.`,
				},
				resource.Attribute{
					Name:        "instant_messaging",
					Description: `The Org User instant messaging.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if the user is enabled and can log in.`,
				},
				resource.Attribute{
					Name:        "is_group_role",
					Description: `True if this user has a group role.`,
				},
				resource.Attribute{
					Name:        "is_locked",
					Description: `If the user account has been locked due to too many invalid login attempts, the value will be true.`,
				},
				resource.Attribute{
					Name:        "is_external",
					Description: `If the user account was imported from an external resource, like an LDAP.`,
				},
				resource.Attribute{
					Name:        "deployed_vm_quota",
					Description: `Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.`,
				},
				resource.Attribute{
					Name:        "stored_vm_quota",
					Description: `Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Organization user`,
				},
				resource.Attribute{
					Name:        "group_names",
					Description: `The set of group names to which this user belongs. It's only populated if the users are created after the group (with this user having a ` + "`" + `depends_on` + "`" + ` of the given group).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_org_vdc",
			Category:         "Data Sources",
			ShortDescription: `Provides an organization VDC data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional, but required if not set at provider level) Org name`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Organization VDC name ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "edge_cluster_id",
					Description: `(`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_portgroup",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for available vCenter Port Groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Organization VDC name`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) ` + "`" + `NETWORK` + "`" + ` for vSwitch port group or ` + "`" + `DV_PORTGROUP` + "`" + ` for distributed port group. ## Attribute reference Only ID is set to be able and reference in other resources or data sources.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_provider_vdc",
			Category:         "Data Sources",
			ShortDescription: `Provides a Provider VDC data source.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Provider VDC name ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Optional description of the Provider VDC.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the Provider VDC, it can be -1 (creation failed), 0 (not ready), 1 (ready), 2 (unknown) or 3 (unrecognized).`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `Set of virtual hardware versions supported by this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "compute_capacity",
					Description: `An indicator of CPU and memory capacity. See [Compute Capacity](#compute-capacity) below for details.`,
				},
				resource.Attribute{
					Name:        "compute_provider_scope",
					Description: `Represents the compute fault domain for this Provider VDC. This value is a tenant-facing tag that is shown to tenants when viewing fault domains of the child Organization VDCs (for example, a VDC Group).`,
				},
				resource.Attribute{
					Name:        "highest_supported_hardware_version",
					Description: `The highest virtual hardware version supported by this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "nsxt_manager_id",
					Description: `ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "storage_containers_ids",
					Description: `Set of IDs of the vSphere Datastores backing this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "external_network_ids",
					Description: `Set of IDs of External Networks.`,
				},
				resource.Attribute{
					Name:        "storage_profile_ids",
					Description: `Set of IDs to the Storage Profiles available to this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "resource_pool_ids",
					Description: `Set of IDs of the Resource Pools backing this provider VDC.`,
				},
				resource.Attribute{
					Name:        "network_pool_ids",
					Description: `Set IDs of the Network Pools used by this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "universal_network_pool_id",
					Description: `ID of the universal network reference.`,
				},
				resource.Attribute{
					Name:        "host_ids",
					Description: `Set with all the hosts which are connected to VC server.`,
				},
				resource.Attribute{
					Name:        "vcenter_id",
					Description: `ID of the vCenter Server that provides the Resource Pools and Datastores.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key and value pairs for Provider VDC Metadata.`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to the Provider VDC. See [Metadata](#metadata) section for details. <a id="compute-capacity"></a> ## Compute Capacity The ` + "`" + `compute_capacity` + "`" + ` attribute is a list with a single item which has the following nested attributes:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `An indicator of CPU. See [CPU and memory](#cpu-and-memory) below.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `An indicator of memory. See [CPU and memory](#cpu-and-memory) below.`,
				},
				resource.Attribute{
					Name:        "is_elastic",
					Description: `True if compute capacity can grow or shrink based on demand.`,
				},
				resource.Attribute{
					Name:        "is_ha",
					Description: `True if compute capacity is highly available. <a id="cpu-and-memory"></a> ### CPU and memory The ` + "`" + `cpu` + "`" + ` and ` + "`" + `memory` + "`" + ` indicators have the following nested attributes:`,
				},
				resource.Attribute{
					Name:        "allocation",
					Description: `Allocated CPU/Memory for this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "overhead",
					Description: `CPU/Memory overhead for this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "reserved",
					Description: `Reserved CPU/Memory for this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `Total CPU/Memory for this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `Units for the CPU/Memory of this Provider VDC.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `Used CPU/Memory in this Provider VDC. <a id="metadata"></a> ## Metadata The ` + "`" + `metadata_entry` + "`" + ` (`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_rde",
			Category:         "Data Sources",
			ShortDescription: `Provides the capability of reading an existing Runtime Defined Entity in VMware Cloud Director.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) Name of the [Organization](/providers/vmware/vcd/latest/docs/data-sources/org) that owns the RDE, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "rde_type_id",
					Description: `(Required) The ID of the [RDE Type](/providers/vmware/vcd/latest/docs/data-sources/rde_type) of the RDE to fetch.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Runtime Defined Entity. ## Attribute Reference The following attributes are supported:`,
				},
				resource.Attribute{
					Name:        "entity",
					Description: `The entity JSON.`,
				},
				resource.Attribute{
					Name:        "owner_user_id",
					Description: `The ID of the [Organization user](/providers/vmware/vcd/latest/docs/resources/org_user) that owns this Runtime Defined Entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the [Organization](/providers/vmware/vcd/latest/docs/resources/org) to which the Runtime Defined Entity belongs.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `It can be ` + "`" + `RESOLVED` + "`" + `, ` + "`" + `RESOLUTION_ERROR` + "`" + ` or ` + "`" + `PRE_CREATED` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "entity",
					Description: `The entity JSON.`,
				},
				resource.Attribute{
					Name:        "owner_user_id",
					Description: `The ID of the [Organization user](/providers/vmware/vcd/latest/docs/resources/org_user) that owns this Runtime Defined Entity.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The ID of the [Organization](/providers/vmware/vcd/latest/docs/resources/org) to which the Runtime Defined Entity belongs.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `It can be ` + "`" + `RESOLVED` + "`" + `, ` + "`" + `RESOLUTION_ERROR` + "`" + ` or ` + "`" + `PRE_CREATED` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_rde_interface",
			Category:         "Data Sources",
			ShortDescription: `Provides the capability of fetching an existing Runtime Defined Entity Interface from VMware Cloud Director.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `(Required) The vendor of the RDE Interface.`,
				},
				resource.Attribute{
					Name:        "nss",
					Description: `(Required) A unique namespace associated with the RDE Interface.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax. ## Attribute Reference All the supported attributes are defined in the [Defined Interface resource](/providers/vmware/vcd/latest/docs/resources/rde_interface#argument-reference).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_rde_type",
			Category:         "Data Sources",
			ShortDescription: `Provides the capability of fetching an existing Runtime Defined Entity Type from VMware Cloud Director.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vendor",
					Description: `(Required) The vendor of the Runtime Defined Entity Type.`,
				},
				resource.Attribute{
					Name:        "nss",
					Description: `(Required) A unique namespace associated with the Runtime Defined Entity Type.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) The version of the Runtime Defined Entity Type. Must follow [semantic versioning](https://semver.org/) syntax. ## Attribute Reference All the supported attributes are defined in the [Runtime Defined Entity Type resource](/providers/vmware/vcd/latest/docs/resources/rde_type#argument-reference).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_resource_list",
			Category:         "Data Sources",
			ShortDescription: `Provides lists of VCD resources`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) An unique name to identify the data source`,
				},
				resource.Attribute{
					Name:        "list",
					Description: `(Computed) The list of requested resources in the chosen format.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "list",
					Description: `(Computed) The list of requested resources in the chosen format.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_resource_schema",
			Category:         "Data Sources",
			ShortDescription: `Provides information about a vCD resource structure`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) An unique name to identify the data source`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Required) Which resource we want to list. It needs to use the full name of the resource (i.e. "vcd_org", not simply "org") ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) The list of attributes for the resource. Each attribute is made of:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attribute`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `an optional description of the attribute`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `whether the attribute is required`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `whether the attribute is optional`,
				},
				resource.Attribute{
					Name:        "computed",
					Description: `whether the attribute is computed`,
				},
				resource.Attribute{
					Name:        "sensitive",
					Description: `whether the attribute is sensitive`,
				},
				resource.Attribute{
					Name:        "block_attributes",
					Description: `(Computed) The list of compound attributes Each bock attribute is made of:`,
				},
				resource.Attribute{
					Name:        "nesting_type",
					Description: `(Computed) How the block is organized (one of ` + "`" + `NestingSet` + "`" + `, ` + "`" + `NestingList` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) Same composition of the simple ` + "`" + `attributes` + "`" + ` above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) The list of attributes for the resource. Each attribute is made of:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attribute`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `an optional description of the attribute`,
				},
				resource.Attribute{
					Name:        "required",
					Description: `whether the attribute is required`,
				},
				resource.Attribute{
					Name:        "optional",
					Description: `whether the attribute is optional`,
				},
				resource.Attribute{
					Name:        "computed",
					Description: `whether the attribute is computed`,
				},
				resource.Attribute{
					Name:        "sensitive",
					Description: `whether the attribute is sensitive`,
				},
				resource.Attribute{
					Name:        "block_attributes",
					Description: `(Computed) The list of compound attributes Each bock attribute is made of:`,
				},
				resource.Attribute{
					Name:        "nesting_type",
					Description: `(Computed) How the block is organized (one of ` + "`" + `NestingSet` + "`" + `, ` + "`" + `NestingList` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `(Computed) Same composition of the simple ` + "`" + `attributes` + "`" + ` above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_right",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Organization Right data source. This can be used to read existing rights`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the right. ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the right`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `The ID of the category for this right`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization`,
				},
				resource.Attribute{
					Name:        "right type",
					Description: `Type of the right (VIEW or MODIFY)`,
				},
				resource.Attribute{
					Name:        "implied_rights",
					Description: `List of rights that are implied with this one ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_rights_bundle",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director rights bundle data source. This can be used to read rights bundles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rights bundle. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the rights bundle`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `Set of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the rights bundle to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `Set of tenants to which this rights bundle gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this rights bundle is read-only ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how rights bundles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `A description of the rights bundle`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `Set of rights assigned to this role`,
				},
				resource.Attribute{
					Name:        "publish_to_all_tenants",
					Description: `When true, publishes the rights bundle to all tenants`,
				},
				resource.Attribute{
					Name:        "tenants",
					Description: `Set of tenants to which this rights bundle gets published. Ignored if ` + "`" + `publish_to_all_tenants` + "`" + ` is true.`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this rights bundle is read-only ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how rights bundles and rights work together.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_role",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director role. This can be used to read roles.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this role is read-only`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `Set of rights assigned to this role ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and rights work together.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "read_only",
					Description: `Whether this role is read-only`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `A description of the role`,
				},
				resource.Attribute{
					Name:        "bundle_key",
					Description: `Key used for internationalization.`,
				},
				resource.Attribute{
					Name:        "rights",
					Description: `Set of rights assigned to this role ## More information See [Roles management](/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and rights work together.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_storage_profile",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for VDC storage profile.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Storage profile name. ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `storage profile identifier (Supported in provider`,
				},
				resource.Attribute{
					Name:        "limit",
					Description: `Maximum number of storage bytes (scaled by 'units' field) allocated for this profile. ` + "`" + `0` + "`" + ` means ` + "`" + `maximum possible` + "`" + ``,
				},
				resource.Attribute{
					Name:        "used_storage",
					Description: `Storage used, by the storage profile (in Megabytes)`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `True if this is default storage profile for this VDC. The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if this storage profile is enabled for use in the VDC`,
				},
				resource.Attribute{
					Name:        "iops_allocated",
					Description: `Total IOPS currently allocated to this storage profile`,
				},
				resource.Attribute{
					Name:        "units",
					Description: `Scale used to define Limit`,
				},
				resource.Attribute{
					Name:        "iops_settings",
					Description: `A block providing IOPS settings. See [IOPS settings](#iopsSettings) below for details. <a id="iopsSettings"></a> ## IOPS settings (Supported from VCD`,
				},
				resource.Attribute{
					Name:        "iops_limiting_enabled",
					Description: `True if this storage profile is IOPS-based placement enabled`,
				},
				resource.Attribute{
					Name:        "maximum_disk_iops",
					Description: `The maximum IOPS value that this storage profile is permitted to deliver. Value of 0 means this max setting is disabled and there is no max disk IOPS restriction`,
				},
				resource.Attribute{
					Name:        "default_disk_iops",
					Description: `Value of 0 for disk IOPS means that no IOPS would be reserved or provisioned for that virtual disk`,
				},
				resource.Attribute{
					Name:        "disk_iops_per_gb_max",
					Description: `The maximum disk IOPs per GB value that this storage profile is permitted to deliver. A value of 0 means there is no per GB IOPS restriction`,
				},
				resource.Attribute{
					Name:        "iops_limit",
					Description: `Maximum number of IOPs that can be allocated for this profile. ` + "`" + `0` + "`" + ` means ` + "`" + `maximum possible` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_subscribed_catalog",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director subscribed catalog data source. This can be used to read a subscribed catalog.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Catalog name. Required if ` + "`" + `filter` + "`" + ` is not set.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `(Optional) Retrieves the data source using one or more filter parameters ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "storage_profile_id",
					Description: `Allows to set specific storage profile to be used for catalog.`,
				},
				resource.Attribute{
					Name:        "subscription_url",
					Description: `The URL to which this catalog is subscribed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the catalog. This is inherited from the publishing catalog`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Optional metadata of the catalog. This is inherited from the publishing catalog`,
				},
				resource.Attribute{
					Name:        "catalog_version",
					Description: `Version number from this catalog.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner of the catalog.`,
				},
				resource.Attribute{
					Name:        "number_of_vapp_templates",
					Description: `Number of vApp templates available in this catalog.`,
				},
				resource.Attribute{
					Name:        "number_of_media",
					Description: `Number of media items available in this catalog.`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `Indicates if the catalog is shared (` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `Indicates if this catalog is available for subscription. (Always return ` + "`" + `false` + "`" + ` for this data source)`,
				},
				resource.Attribute{
					Name:        "is_local",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "publish_subscription_type",
					Description: `Shows if the catalog is published, if it is a subscription from another one or none of those. (Always returns ` + "`" + `SUBSCRIBED` + "`" + ` for this data source)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `the catalog's Hyper reference.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Date and time of catalog creation.`,
				},
				resource.Attribute{
					Name:        "running_tasks",
					Description: `List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.`,
				},
				resource.Attribute{
					Name:        "failed_tasks",
					Description: `List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items. ## Filter arguments`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "storage_profile_id",
					Description: `Allows to set specific storage profile to be used for catalog.`,
				},
				resource.Attribute{
					Name:        "subscription_url",
					Description: `The URL to which this catalog is subscribed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the catalog. This is inherited from the publishing catalog`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Optional metadata of the catalog. This is inherited from the publishing catalog`,
				},
				resource.Attribute{
					Name:        "catalog_version",
					Description: `Version number from this catalog.`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `Owner of the catalog.`,
				},
				resource.Attribute{
					Name:        "number_of_vapp_templates",
					Description: `Number of vApp templates available in this catalog.`,
				},
				resource.Attribute{
					Name:        "number_of_media",
					Description: `Number of media items available in this catalog.`,
				},
				resource.Attribute{
					Name:        "is_shared",
					Description: `Indicates if the catalog is shared (` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_published",
					Description: `Indicates if this catalog is available for subscription. (Always return ` + "`" + `false` + "`" + ` for this data source)`,
				},
				resource.Attribute{
					Name:        "is_local",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "publish_subscription_type",
					Description: `Shows if the catalog is published, if it is a subscription from another one or none of those. (Always returns ` + "`" + `SUBSCRIBED` + "`" + ` for this data source)`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `the catalog's Hyper reference.`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `Date and time of catalog creation.`,
				},
				resource.Attribute{
					Name:        "running_tasks",
					Description: `List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.`,
				},
				resource.Attribute{
					Name:        "failed_tasks",
					Description: `List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items. ## Filter arguments`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_task",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director Organization Task data source. This can be used to read existing tasks.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The ID of the task ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The URI of the task.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the task. May not be unique. Defines the general operation being performed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `An optional description of the task.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of the task.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The execution status of the task. One of queued, preRunning, running, success, error, aborted.`,
				},
				resource.Attribute{
					Name:        "operation",
					Description: `A message describing the operation that is tracked by this task.`,
				},
				resource.Attribute{
					Name:        "operation_name",
					Description: `The short name of the operation that is tracked by this task.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The date and time the system started executing the task. May not be present if the task has not been executed yet.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The date and time that processing of the task was completed. May not be present if the task is still being executed.`,
				},
				resource.Attribute{
					Name:        "expiry_time",
					Description: `The date and time at which the task resource will be destroyed and no longer available for retrieval. May not be present if the task has not been executed or is still being executed.`,
				},
				resource.Attribute{
					Name:        "cancel_requested",
					Description: `Whether user has requested this processing to be canceled (` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "owner_name",
					Description: `The name of the task owner. This is typically the object that the task is creating or updating.`,
				},
				resource.Attribute{
					Name:        "owner_type",
					Description: `The type of the task owner.`,
				},
				resource.Attribute{
					Name:        "owner_id",
					Description: `The unique identifier of the task owner.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `error information from a failed task.`,
				},
				resource.Attribute{
					Name:        "user_name",
					Description: `The name of the user who started the task.`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The unique identifier of the task user.`,
				},
				resource.Attribute{
					Name:        "org_name",
					Description: `The name of the org to which the user belongs.`,
				},
				resource.Attribute{
					Name:        "org_id",
					Description: `The unique identifier of the user org.`,
				},
				resource.Attribute{
					Name:        "progress",
					Description: `Indicator of task progress as an approximate percentage between 0 and 100. Not available for all tasks.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director vApp data source. This can be used to reference vApps.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A unique name for the vApp`,
				},
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "href",
					Description: `The vApp Hyper Reference`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key value map of metadata assigned to this vApp. Key and value can be any string.`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to this vApp. See [Metadata](#metadata) section for details.`,
				},
				resource.Attribute{
					Name:        "guest_properties",
					Description: `Key value map of vApp guest properties.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The vApp status as a numeric code`,
				},
				resource.Attribute{
					Name:        "status_text",
					Description: `The vApp status as text.`,
				},
				resource.Attribute{
					Name:        "lease",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "runtime_lease_in_sec",
					Description: `How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires.`,
				},
				resource.Attribute{
					Name:        "storage_lease_in_sec",
					Description: `How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires. <a id="metadata"></a> ## Metadata The ` + "`" + `metadata_entry` + "`" + ` (`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director vApp network data source. This can be used to access a vApp network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp name.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the vApp network, unique within the vApp ## Attribute reference All attributes defined in [` + "`" + `vcd_vapp_network` + "`" + `](/providers/vmware/vcd/latest/docs/resources/vapp_network#attribute-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_org_network",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for VMware Cloud Director Org network attached to vApp. This can be used to access vApp Org network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp name.`,
				},
				resource.Attribute{
					Name:        "org_network_name",
					Description: `(Required) A name for the vApp Org network, unique within the vApp. ## Attribute reference All attributes defined in [` + "`" + `vcd_vapp_org_network` + "`" + `](/providers/vmware/vcd/latest/docs/resources/vapp_org_network#attribute-reference) are supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vapp_vm",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM data source. This can be used to access VMs within a vApp.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "vapp_name",
					Description: `(Required) The vApp this VM belongs to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name for the VM, unique within the vApp`,
				},
				resource.Attribute{
					Name:        "network_dhcp_wait_seconds",
					Description: `(Optional;`,
				},
				resource.Attribute{
					Name:        "vm_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "computer_name",
					Description: `Computer name to assign to this virtual machine.`,
				},
				resource.Attribute{
					Name:        "catalog_name",
					Description: `The catalog name in which to find the given vApp Template`,
				},
				resource.Attribute{
					Name:        "template_name",
					Description: `The name of the vApp Template to use`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of RAM (in MB) allocated to the VM`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `The amount of RAM (in MB) reservation on the underlying virtualization infrastructure`,
				},
				resource.Attribute{
					Name:        "memory_priority",
					Description: `Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload. Values can be: ` + "`" + `LOW` + "`" + `, ` + "`" + `NORMAL` + "`" + `, ` + "`" + `HIGH` + "`" + ` and ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_shares",
					Description: `Custom priority for the resource in MB`,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `The limit (in MB) for how much of memory can be consumed on the underlying virtualization infrastructure. ` + "`" + `-1` + "`" + ` value for unlimited.`,
				},
				resource.Attribute{
					Name:        "cpus",
					Description: `The number of virtual CPUs allocated to the VM`,
				},
				resource.Attribute{
					Name:        "cpu_cores",
					Description: `The number of cores per socket`,
				},
				resource.Attribute{
					Name:        "cpu_reservation",
					Description: `The amount of MHz reservation on the underlying virtualization infrastructure`,
				},
				resource.Attribute{
					Name:        "cpu_priority",
					Description: `Pre-determined relative priorities according to which the non-reserved portion of this resource is made available to the virtualized workload. Values can be: ` + "`" + `LOW` + "`" + `, ` + "`" + `NORMAL` + "`" + `, ` + "`" + `HIGH` + "`" + ` and ` + "`" + `CUSTOM` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_shares",
					Description: `Custom priority for the resource in MHz`,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `The limit (in MHz) for how much of CPU can be consumed on the underlying virtualization infrastructure. ` + "`" + `-1` + "`" + ` value for unlimited.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `(Deprecated) Use ` + "`" + `metadata_entry` + "`" + ` instead. Key value map of metadata assigned to this VM`,
				},
				resource.Attribute{
					Name:        "metadata_entry",
					Description: `A set of metadata entries assigned to this VM. See [Metadata](#metadata) section for details`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Independent disk attachment configuration.`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `A block defining a network interface. Multiple can be used.`,
				},
				resource.Attribute{
					Name:        "guest_properties",
					Description: `Key value map of guest properties`,
				},
				resource.Attribute{
					Name:        "expose_hardware_virtualization",
					Description: `Expose hardware-assisted CPU virtualization to guest OS`,
				},
				resource.Attribute{
					Name:        "internal_disk",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "hardware_version",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "sizing_policy_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "placement_policy_id",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "status_text",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "security_tags",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Value of this metadata entry.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of this metadata entry. One of: ` + "`" + `MetadataStringValue` + "`" + `, ` + "`" + `MetadataNumberValue` + "`" + `, ` + "`" + `MetadataDateTimeValue` + "`" + `, ` + "`" + `MetadataBooleanValue` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_access",
					Description: `User access level for this metadata entry. One of: ` + "`" + `PRIVATE` + "`" + ` (hidden), ` + "`" + `READONLY` + "`" + ` (read only), ` + "`" + `READWRITE` + "`" + ` (read/write).`,
				},
				resource.Attribute{
					Name:        "is_system",
					Description: `Domain for this metadata entry. true if it belongs to ` + "`" + `SYSTEM` + "`" + `, false if it belongs to ` + "`" + `GENERAL` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vcenter",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source for vCenter server attached to VCD.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) vCenter name ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "connection_status",
					Description: `vCenter connection status (e.g. ` + "`" + `CONNECTED` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "is_enabled",
					Description: `Boolean value if vCenter is enabled.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `vCenter status (e.g. ` + "`" + `READY` + "`" + `).`,
				},
				resource.Attribute{
					Name:        "vcenter_host",
					Description: `Hostname of configured vCenter.`,
				},
				resource.Attribute{
					Name:        "vcenter_version",
					Description: `vCenter version (e.g. ` + "`" + `6.7.0` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vdc_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a data source to read VDC groups.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of VDC group`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) - ID of VDC group Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be used. If both are missing, an error arises. ## Attribute Reference All the arguments and attributes defined in [` + "`" + `vcd_vdc_group` + "`" + `](/providers/vmware/vcd/latest/docs/resources/vdc_group) resource are available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM data source. This can be used to access standalone VMs.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `(Optional) The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `(Optional) The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) A name or ID for the standalone VM in VDC ## Attributes reference This data source provides all attributes available for [` + "`" + `vcd_vapp_vm` + "`" + `](/providers/vmware/vcd/latest/docs/data-sources/vapp_vm).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_affinity_rule",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM affinity rule data source. This can be used to read VM affinity and anti-affinity rules.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "org",
					Description: `The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations`,
				},
				resource.Attribute{
					Name:        "vdc",
					Description: `The name of VDC to use, optional if defined at provider level`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of VM affinity rule. Needed if we don't provide ` + "`" + `rule_id` + "`" + ``,
				},
				resource.Attribute{
					Name:        "rule_id",
					Description: `Is the ID of the affinity rule. It's the preferred way to retrieve the affinity rule, especially if the rule name could have duplicates ## Attribute reference`,
				},
				resource.Attribute{
					Name:        "polarity",
					Description: `One of ` + "`" + `Affinity` + "`" + ` or ` + "`" + `Anti-Affinity` + "`" + `. This property cannot be changed. Once created, if we need to change polarity, we need to remove the rule and create a new one.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_group",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM Group data source. This can be used to fetch vSphere VM Groups and create VM Placement Policies with them.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of VM Group to fetch from vSphere.`,
				},
				resource.Attribute{
					Name:        "provider_vdc_id",
					Description: `(Required) The ID of [Provider VDC](/providers/vmware/vcd/latest/docs/data-sources/provider_vdc) to which the VM Group belongs. ## Attributes reference`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `Name of the vSphere cluster associated to this VM Group.`,
				},
				resource.Attribute{
					Name:        "named_vm_group_id",
					Description: `ID of the named VM Group. Used to create Logical VM Groups.`,
				},
				resource.Attribute{
					Name:        "vcenter_id",
					Description: `ID of the vCenter server.`,
				},
				resource.Attribute{
					Name:        "cluster_moref",
					Description: `Managed object reference of the vSphere cluster associated to this VM Group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_placement_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM Placement Policy data source. This can be used to read a VM placement policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name VM Placement Policy.`,
				},
				resource.Attribute{
					Name:        "provider_vdc_id",
					Description: `(Required for System administrator users) The ID of the [Provider VDC](/providers/vmware/vcd/latest/docs/data-sources/provider_vdc) to which the VM Placement Policy belongs.`,
				},
				resource.Attribute{
					Name:        "vdc_id",
					Description: `(Required for Org users;`,
				},
				resource.Attribute{
					Name:        "vm_group_ids",
					Description: `This attribute can't be retrieved if the data source is used by a tenant user when fetching by ` + "`" + `vdc_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logical_vm_group_ids",
					Description: `This attribute can't be retrieved if the data source is used by a tenant user when fetching by ` + "`" + `vdc_id` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "vm_group_ids",
					Description: `This attribute can't be retrieved if the data source is used by a tenant user when fetching by ` + "`" + `vdc_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logical_vm_group_ids",
					Description: `This attribute can't be retrieved if the data source is used by a tenant user when fetching by ` + "`" + `vdc_id` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vcd_vm_sizing_policy",
			Category:         "Data Sources",
			ShortDescription: `Provides a VMware Cloud Director VM sizing policy data source. This can be used to read VM sizing policy.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name VM sizing policy ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"vcd_catalog":                                   0,
		"vcd_catalog_item":                              1,
		"vcd_catalog_media":                             2,
		"vcd_catalog_vapp_template":                     3,
		"vcd_library_certificate":                       4,
		"vcd_edgegateway":                               5,
		"vcd_external_network":                          6,
		"vcd_external_network_v2":                       7,
		"vcd_global_role":                               8,
		"vcd_independent_disk":                          9,
		"vcd_lb_app_profile":                            10,
		"vcd_lb_app_rule":                               11,
		"vcd_lb_server_pool":                            12,
		"vcd_lb_service_monitor":                        13,
		"vcd_lb_virtual_server":                         14,
		"vcd_network_direct":                            15,
		"vcd_network_isolated":                          16,
		"vcd_network_isolated_v2":                       17,
		"vcd_network_routed":                            18,
		"vcd_network_routed_v2":                         19,
		"vcd_nsxt_alb_cloud":                            20,
		"vcd_nsxt_alb_controller":                       21,
		"vcd_nsxt_alb_edgegateway_service_engine_group": 22,
		"vcd_nsxt_alb_importable_cloud":                 23,
		"vcd_nsxt_alb_pool":                             24,
		"vcd_nsxt_alb_service_engine_group":             25,
		"vcd_nsxt_alb_settings":                         26,
		"vcd_nsxt_alb_virtual_service":                  27,
		"vcd_nsxt_app_port_profile":                     28,
		"vcd_nsxt_distributed_firewall":                 29,
		"vcd_nsxt_edge_cluster":                         30,
		"vcd_nsxt_edgegateway":                          31,
		"vcd_nsxt_edgegateway_bgp_configuration":        32,
		"vcd_nsxt_edgegateway_bgp_ip_prefix_list":       33,
		"vcd_nsxt_edgegateway_bgp_neighbor":             34,
		"vcd_nsxt_edgegateway_qos_profile":              35,
		"vcd_nsxt_edgegateway_rate_limiting":            36,
		"vcd_nsxt_firewall":                             37,
		"vcd_nsxt_ip_set":                               38,
		"vcd_nsxt_ipsec_vpn_tunnel":                     39,
		"vcd_nsxt_manager":                              40,
		"vcd_nsxt_nat_rule":                             41,
		"vcd_nsxt_dynamic_security_group":               42,
		"vcd_nsxt_network_dhcp":                         43,
		"vcd_nsxt_network_dhcp_binding":                 44,
		"vcd_nsxt_network_imported":                     45,
		"vcd_nsxt_security_group":                       46,
		"vcd_nsxt_tier0_router":                         47,
		"vcd_nsxv_application":                          48,
		"vcd_nsxv_application_finder":                   49,
		"vcd_nsxv_application_group":                    50,
		"vcd_nsxv_dhcp_relay":                           51,
		"vcd_nsxv_distributed_firewall":                 52,
		"vcd_nsxv_dnat":                                 53,
		"vcd_nsxv_firewall_rule":                        54,
		"vcd_nsxv_ip_set":                               55,
		"vcd_nsxv_snat":                                 56,
		"vcd_org":                                       57,
		"vcd_org_group":                                 58,
		"vcd_org_ldap":                                  59,
		"vcd_org_user":                                  60,
		"vcd_org_vdc":                                   61,
		"vcd_portgroup":                                 62,
		"vcd_provider_vdc":                              63,
		"vcd_rde":                                       64,
		"vcd_rde_interface":                             65,
		"vcd_rde_type":                                  66,
		"vcd_resource_list":                             67,
		"vcd_resource_schema":                           68,
		"vcd_right":                                     69,
		"vcd_rights_bundle":                             70,
		"vcd_role":                                      71,
		"vcd_storage_profile":                           72,
		"vcd_subscribed_catalog":                        73,
		"vcd_task":                                      74,
		"vcd_vapp":                                      75,
		"vcd_vapp_network":                              76,
		"vcd_vapp_org_network":                          77,
		"vcd_vapp_vm":                                   78,
		"vcd_vcenter":                                   79,
		"vcd_vdc_group":                                 80,
		"vcd_vm":                                        81,
		"vcd_vm_affinity_rule":                          82,
		"vcd_vm_group":                                  83,
		"vcd_vm_placement_policy":                       84,
		"vcd_vm_sizing_policy":                          85,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
