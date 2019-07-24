package circonus

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "circonus_account",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Circonus Account.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The Circonus ID of a given account.`,
				},
				resource.Attribute{
					Name:        "current",
					Description: `(Optional) Automatically use the current Circonus Account attached to the API token making the request. At least one of the above attributes should be provided when searching for a account. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "address1",
					Description: `The first line of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "address2",
					Description: `The second line of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "cc_email",
					Description: `An optionally specified email address used in the CC line of invoices.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Circonus ID of the selected Account.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city part of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `A list of IDs for each contact group in the account.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the user's address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the account.`,
				},
				resource.Attribute{
					Name:        "invites",
					Description: `An list of users invited to use the platform. Each element in the list has both an ` + "`" + `email` + "`" + ` and ` + "`" + `role` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the account.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The Circonus ID of the user who owns this account.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state or province of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The timezone that events will be displayed in the web interface for this account.`,
				},
				resource.Attribute{
					Name:        "ui_base_url",
					Description: `The base URL of this account.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `A list of account usage limits. Each element in the list will have a ` + "`" + `limit` + "`" + ` attribute, a limit ` + "`" + `type` + "`" + `, and a ` + "`" + `used` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users who have access to this account. Each element in the list has both an ` + "`" + `id` + "`" + ` and a ` + "`" + `role` + "`" + `. The ` + "`" + `id` + "`" + ` is a Circonus ID referencing the user.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "address1",
					Description: `The first line of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "address2",
					Description: `The second line of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "cc_email",
					Description: `An optionally specified email address used in the CC line of invoices.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Circonus ID of the selected Account.`,
				},
				resource.Attribute{
					Name:        "city",
					Description: `The city part of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "contact_groups",
					Description: `A list of IDs for each contact group in the account.`,
				},
				resource.Attribute{
					Name:        "country",
					Description: `The country of the user's address.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of the account.`,
				},
				resource.Attribute{
					Name:        "invites",
					Description: `An list of users invited to use the platform. Each element in the list has both an ` + "`" + `email` + "`" + ` and ` + "`" + `role` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the account.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `The Circonus ID of the user who owns this account.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state or province of the address associated with the account.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `The timezone that events will be displayed in the web interface for this account.`,
				},
				resource.Attribute{
					Name:        "ui_base_url",
					Description: `The base URL of this account.`,
				},
				resource.Attribute{
					Name:        "usage",
					Description: `A list of account usage limits. Each element in the list will have a ` + "`" + `limit` + "`" + ` attribute, a limit ` + "`" + `type` + "`" + `, and a ` + "`" + `used` + "`" + ` attribute.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `A list of users who have access to this account. Each element in the list has both an ` + "`" + `id` + "`" + ` and a ` + "`" + `role` + "`" + `. The ` + "`" + `id` + "`" + ` is a Circonus ID referencing the user.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "circonus_collector",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Circonus Collector.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The Circonus ID of a given collector. At least one of the above attributes should be provided when searching for a collector. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The Circonus ID of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `A list of details about the individual Collector instances that make up the group of collectors. See below for a list of attributes within each collector.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `The latitude of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `The longitude of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags assigned to the selected Collector.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The of the selected Collector. This value is either ` + "`" + `circonus` + "`" + ` for a Circonus-managed, public Collector, or ` + "`" + `enterprise` + "`" + ` for a private collector that is private to an account. ## Collector Details`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The CN of an individual Collector in the Collector Group.`,
				},
				resource.Attribute{
					Name:        "external_host",
					Description: `The external host information for an individual Collector in the Collector Group. This is useful or important when talking with a Collector through a NAT'ing firewall.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `The external port number for an individual Collector in the Collector Group. This is useful or important when talking with a Collector through a NAT'ing firewall.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP address of an individual Collector in the Collector Group. This is the IP address of the interface listening on the network.`,
				},
				resource.Attribute{
					Name:        "min_version",
					Description: `??`,
				},
				resource.Attribute{
					Name:        "modules",
					Description: `A list of what modules (types of checks) this collector supports.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port the collector responds to the Circonus HTTPS REST wire protocol on.`,
				},
				resource.Attribute{
					Name:        "skew",
					Description: `The clock drift between this collector and the Circonus server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this particular collector. A string containing either ` + "`" + `active` + "`" + `, ` + "`" + `unprovisioned` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `provisioned` + "`" + `, or ` + "`" + `retired` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the collector software the collector is running.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The Circonus ID of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "details",
					Description: `A list of details about the individual Collector instances that make up the group of collectors. See below for a list of attributes within each collector.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `The latitude of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `The longitude of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the selected Collector.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of tags assigned to the selected Collector.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The of the selected Collector. This value is either ` + "`" + `circonus` + "`" + ` for a Circonus-managed, public Collector, or ` + "`" + `enterprise` + "`" + ` for a private collector that is private to an account. ## Collector Details`,
				},
				resource.Attribute{
					Name:        "cn",
					Description: `The CN of an individual Collector in the Collector Group.`,
				},
				resource.Attribute{
					Name:        "external_host",
					Description: `The external host information for an individual Collector in the Collector Group. This is useful or important when talking with a Collector through a NAT'ing firewall.`,
				},
				resource.Attribute{
					Name:        "external_port",
					Description: `The external port number for an individual Collector in the Collector Group. This is useful or important when talking with a Collector through a NAT'ing firewall.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The IP address of an individual Collector in the Collector Group. This is the IP address of the interface listening on the network.`,
				},
				resource.Attribute{
					Name:        "min_version",
					Description: `??`,
				},
				resource.Attribute{
					Name:        "modules",
					Description: `A list of what modules (types of checks) this collector supports.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port the collector responds to the Circonus HTTPS REST wire protocol on.`,
				},
				resource.Attribute{
					Name:        "skew",
					Description: `The clock drift between this collector and the Circonus server.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of this particular collector. A string containing either ` + "`" + `active` + "`" + `, ` + "`" + `unprovisioned` + "`" + `, ` + "`" + `pending` + "`" + `, ` + "`" + `provisioned` + "`" + `, or ` + "`" + `retired` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of the collector software the collector is running.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"circonus_account":   0,
		"circonus_collector": 1,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
