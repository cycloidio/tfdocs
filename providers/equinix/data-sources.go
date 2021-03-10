package equinix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "equinix_ecx_l2_sellerprofile",
			Category:         "Data Sources",
			ShortDescription: `Provides Equinix Fabric Layer2 Seller Profile resource`,
			Description: `

Use this data source to get details of Equinix Fabric layer 2
seller profile with a given name and / or organization.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_ecx_l2_sellerprofiles",
			Category:         "Data Sources",
			ShortDescription: `Provides Equinix Fabric Layer2 Seller Profile resource`,
			Description: `

Use this data source to get details of available Equinix Fabric layer 2
seller profiles. It is possible to apply filtering criteria for
returned list of profiles.

`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_ecx_port",
			Category:         "Data Sources",
			ShortDescription: `Get information on Equinix Fabric port`,
			Description: `

Use this data source to get details of Equinix Fabric port with a given name.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the port ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the port`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the port`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Port location region`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Port location Equinix Business Exchange (IBX)`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `Port location metro code`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the device (primary / secondary) where the port resides`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `The VLAN encapsulation of the port (Dot1q or QinQ)`,
				},
				resource.Attribute{
					Name:        "buyout",
					Description: `Boolean value that indicates whether the port supports unlimited connections. If "false", the port is a standard port with limited connections. If "true", the port is an "unlimited connections" port that allows multiple connections at no additional charge.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Port Bandwidth in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status that indicates whether a port has been assigned or is ready for connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "uuid",
					Description: `Unique identifier of the port`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of the port`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `Port location region`,
				},
				resource.Attribute{
					Name:        "ibx",
					Description: `Port location Equinix Business Exchange (IBX)`,
				},
				resource.Attribute{
					Name:        "metro_code",
					Description: `Port location metro code`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `The priority of the device (primary / secondary) where the port resides`,
				},
				resource.Attribute{
					Name:        "encapsulation",
					Description: `The VLAN encapsulation of the port (Dot1q or QinQ)`,
				},
				resource.Attribute{
					Name:        "buyout",
					Description: `Boolean value that indicates whether the port supports unlimited connections. If "false", the port is a standard port with limited connections. If "true", the port is an "unlimited connections" port that allows multiple connections at no additional charge.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `Port Bandwidth in bytes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Port status that indicates whether a port has been assigned or is ready for connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_account",
			Category:         "Data Sources",
			ShortDescription: `Get information on Equinix Network Edge billing account`,
			Description: `

Use this data source to get number and identifier of Equinix Network Edge
billing account in a given metro location.

Billing account reference is required to create Network Edge virtual device
in corresponding metro location.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "metro_code",
					Description: `(Required) Account location metro code`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Account name for filtering`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Account status for filtering. Possible values are "Active", "Processing", "Submitted", "Staged" ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `Account unique number`,
				},
				resource.Attribute{
					Name:        "ucm_id",
					Description: `Account unique identifier`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "number",
					Description: `Account unique number`,
				},
				resource.Attribute{
					Name:        "ucm_id",
					Description: `Account unique identifier`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_device_platform",
			Category:         "Data Sources",
			ShortDescription: `Get information on Equinix Network Edge device platform configuration`,
			Description: `

Use this data source to get Equinix Network Edge device platform configuration details
for a given device type.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Device type code`,
				},
				resource.Attribute{
					Name:        "flavor",
					Description: `(Optional) Device platform flavor that determines number of CPU cores and memory. Supported values:`,
				},
				resource.Attribute{
					Name:        "core_count",
					Description: `(Optional) Number of CPU cores used to limit platform search results`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `(Optional) List of software package codes to limit platform search results`,
				},
				resource.Attribute{
					Name:        "management_types",
					Description: `(Optional) List of device management types to limit platform search results. Supported values:`,
				},
				resource.Attribute{
					Name:        "license_options",
					Description: `(Optional) List of device licensing options to limit platform search result. Supported values:`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of memory provided by device platform`,
				},
				resource.Attribute{
					Name:        "memory_unit",
					Description: `Unit of memory provider by device platform`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "memory",
					Description: `The amount of memory provided by device platform`,
				},
				resource.Attribute{
					Name:        "memory_unit",
					Description: `Unit of memory provider by device platform`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_device_software",
			Category:         "Data Sources",
			ShortDescription: `Get information on Equinix Network Edge device software`,
			Description: `

Use this data source to get Equinix Network Edge device software details for a given
device type.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_type",
					Description: `(Required) Code of a device type`,
				},
				resource.Attribute{
					Name:        "version_regex",
					Description: `(Optional) A regex string to apply on returned versions and filter search results`,
				},
				resource.Attribute{
					Name:        "stable",
					Description: `(Optional) Boolean value to limit query results to stable versions only`,
				},
				resource.Attribute{
					Name:        "packages",
					Description: `(Optional) Limits returned versions to those that are supported by given software package codes`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) Boolean value to indicate that most recent version should be used`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `Version number`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Software image name`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Version release date`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Version status`,
				},
				resource.Attribute{
					Name:        "release_notes_link",
					Description: `Link to version release notes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `Version number`,
				},
				resource.Attribute{
					Name:        "image_name",
					Description: `Software image name`,
				},
				resource.Attribute{
					Name:        "date",
					Description: `Version release date`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Version status`,
				},
				resource.Attribute{
					Name:        "release_notes_link",
					Description: `Link to version release notes`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "equinix_network_device_type",
			Category:         "Data Sources",
			ShortDescription: `Get information on Equinix Network Edge device type`,
			Description: `

Use this data source to get Equinix Network Edge device type details.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Device type name`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `(Optional) Device type vendor i.e.`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `(Optional) Device type category, one of:`,
				},
				resource.Attribute{
					Name:        "metro_codes",
					Description: `(Optional) List of metro codes where device type has to be available ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `Device type short code, unique identifier of a network device type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Device type textual description`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "code",
					Description: `Device type short code, unique identifier of a network device type`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Device type textual description`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"equinix_ecx_l2_sellerprofile":    0,
		"equinix_ecx_l2_sellerprofiles":   1,
		"equinix_ecx_port":                2,
		"equinix_network_account":         3,
		"equinix_network_device_platform": 4,
		"equinix_network_device_software": 5,
		"equinix_network_device_type":     6,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
