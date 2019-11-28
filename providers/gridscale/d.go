package gridscale

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "gridscale_ip",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of an ip.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the ip.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Defines the IP Address (v4 or v6) the ip.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `The IP prefix of the ip.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter an object belongs to.`,
				},
				resource.Attribute{
					Name:        "failover",
					Description: `failover mode of this ip. If true, then this IP is no longer available for DHCP and can no longer be related to any server..`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the ip.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `The reverse DNS of the ip.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country of the ip.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location of the ip.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time the ip was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last ip change.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the ip is administratively blocked.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the ip has been running.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_network",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of a network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The UUID of the network.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location, that helps to identify which datacenter the network belongs to.`,
				},
				resource.Attribute{
					Name:        "l2security",
					Description: `Defines information about MAC spoofing protection.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the network.`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The type of the network.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the network locates.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code, which works as a location identifier.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The uman-readable name of the location where the network locates.`,
				},
				resource.Attribute{
					Name:        "delete_block",
					Description: `Defines if the network is administratively blocked.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the network was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last network change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_sshkey",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of an sshkey.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the sshkey.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the sshkey.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `The OpenSSH public key string of the sshkey.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the sshkey.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `The date and time of the sshkey was initially created.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `The date and time of the last sshkey change.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_storage",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of a storage.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the storage.`,
				},
				resource.Attribute{
					Name:        "change_time",
					Description: `Defines the date and time of the last storage change.`,
				},
				resource.Attribute{
					Name:        "location_iata",
					Description: `The IATA airport code of the location where storage locates.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the storage.`,
				},
				resource.Attribute{
					Name:        "license_product_no",
					Description: `The license key (e.g. Windows Servers), if the template used by the storage requires.`,
				},
				resource.Attribute{
					Name:        "location_country",
					Description: `The human-readable name of the country where the storage locates.`,
				},
				resource.Attribute{
					Name:        "usage_in_minutes",
					Description: `Total minutes the the storage has been running.`,
				},
				resource.Attribute{
					Name:        "last_used_template",
					Description: `The UUID of the last used template on the storage.`,
				},
				resource.Attribute{
					Name:        "current_price",
					Description: `The price for the current period since the last bill.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `The capacity (GB) of the storage.`,
				},
				resource.Attribute{
					Name:        "location_uuid",
					Description: `The UUID of the location where the storage locates.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The type of the storage.`,
				},
				resource.Attribute{
					Name:        "parent_uuid",
					Description: `The UUID of the parent of the storage.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The human-readable name of the storage.`,
				},
				resource.Attribute{
					Name:        "location_name",
					Description: `The human-readable name of the location where the storage locates.`,
				},
				resource.Attribute{
					Name:        "create_time",
					Description: `Defines the date and time the storage was initially created.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `The list of labels.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "gridscale_template",
			Category:         "Data Sources",
			ShortDescription: `Gets the id of a template by name.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The exact name of the template as show in [the expert panel of gridscale](https://my.gridscale.io/Expert/Template). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the template.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"gridscale_ip":       0,
		"gridscale_network":  1,
		"gridscale_sshkey":   2,
		"gridscale_storage":  3,
		"gridscale_template": 4,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
