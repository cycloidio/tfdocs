package pnap

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "pnap_server",
			Category:         "Resources",
			ShortDescription: `Provides a PhoenixNAP server resource. This can be used to create, modify, and delete servers.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Server hostname.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server description.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `(Required) The server’s OS ID used when the server was created (e.g., ubuntu/bionic, centos/centos7). For a full list of available operating systems visit [API docs](https://developers.phoenixnap.com/docs/bmc/1).`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Server type ID. Cannot be changed once a server is created (e.g., s1.c1.small, s1.c1.medium).`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) Server Location ID. Cannot be changed once a server is created (e.g., PHX).`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Required) A list of SSH Keys that will be installed on the Linux server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Action to perform on server. Allowed actions are: reboot, reset, powered-on, powered-off, shutdown. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `A description of the machine's CPU.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Server hostname.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Server Location ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The server’s OS ID used when the server was created.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `A description of the machine's RAM.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server type ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private IP Addresses assigned to server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public IP Addresses assigned to server. Must contain at least 1 item`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `A description of the machine's CPU.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Server description.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `Server hostname.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique identifier of the server.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Server Location ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `The server’s OS ID used when the server was created.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `A description of the machine's RAM.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server type ID. Cannot be changed once a server is created.`,
				},
				resource.Attribute{
					Name:        "private_ip_addresses",
					Description: `Private IP Addresses assigned to server. Must contain at least 1 item.`,
				},
				resource.Attribute{
					Name:        "public_ip_addresses",
					Description: `Public IP Addresses assigned to server. Must contain at least 1 item`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"pnap_server": 0,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
