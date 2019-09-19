package panos

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "panos_dhcp_interface_info",
			Category:         "Data Sources",
			ShortDescription: `Gets DHCP client information on the given data interface.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "interface",
					Description: `(Required) The data interface to get DHCP information for. These attributes are exported once the data source refreshes:`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The interface's state.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `DHCP IP address.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The default gateway assigned.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `The DHCP server IP`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `DHCP server ID`,
				},
				resource.Attribute{
					Name:        "primary_dns",
					Description: `Primary DNS server`,
				},
				resource.Attribute{
					Name:        "secondary_dns",
					Description: `Secondary DNS server`,
				},
				resource.Attribute{
					Name:        "primary_wins",
					Description: `Primary WINS server`,
				},
				resource.Attribute{
					Name:        "secondary_wins",
					Description: `Secondary WINS`,
				},
				resource.Attribute{
					Name:        "primary_nis",
					Description: `Primary NIS`,
				},
				resource.Attribute{
					Name:        "secondary_nis",
					Description: `Secondary NIS`,
				},
				resource.Attribute{
					Name:        "primary_ntp",
					Description: `Primary NTP`,
				},
				resource.Attribute{
					Name:        "secondary_ntp",
					Description: `Secondary NTP`,
				},
				resource.Attribute{
					Name:        "pop3_server",
					Description: `POP3 Server`,
				},
				resource.Attribute{
					Name:        "smtp_server",
					Description: `SMTP Server`,
				},
				resource.Attribute{
					Name:        "dns_suffix",
					Description: `DNS Suffix`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_panorama_plugin",
			Category:         "Data Sources",
			ShortDescription: `Gets Panorama plugin info.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "installed",
					Description: `A list of installed plugins.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `(int) Total number of plugins, installed or not.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `A list of maps (see below). The following attributes are present in each ` + "`" + `details` + "`" + ` entry:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version number.`,
				},
				resource.Attribute{
					Name:        "release_date",
					Description: `Release date.`,
				},
				resource.Attribute{
					Name:        "release_note_url",
					Description: `Release note URL.`,
				},
				resource.Attribute{
					Name:        "package_file",
					Description: `The package file.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size.`,
				},
				resource.Attribute{
					Name:        "platform",
					Description: `Platform.`,
				},
				resource.Attribute{
					Name:        "installed",
					Description: `If the package is installed or not.`,
				},
				resource.Attribute{
					Name:        "downloaded",
					Description: `If the package is downloaded or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "panos_system_info",
			Category:         "Data Sources",
			ShortDescription: `Gets system info from the firewall.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "info",
					Description: `a map containing the contents of ` + "`" + `show system info` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version_major",
					Description: `Major version number.`,
				},
				resource.Attribute{
					Name:        "version_minor",
					Description: `Minor version number.`,
				},
				resource.Attribute{
					Name:        "version_patch",
					Description: `Patch version number.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"panos_dhcp_interface_info": 0,
		"panos_panorama_plugin":     1,
		"panos_system_info":         2,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
