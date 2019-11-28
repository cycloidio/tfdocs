package cherryservers

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cherryservers_floating_ip",
			Category:         "Resources",
			ShortDescription: `Provides a CherryServers IP resource. This can be used to create, modify, and delete Static IP addresses. These IPs can be added to your server and have the network interfaces configured at creation time.`,
			Description:      ``,
			Keywords: []string{
				"floating",
				"ip",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Project ID to deploy the IP in.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region to deply in.`,
				},
				resource.Attribute{
					Name:        "routed_to_hostname",
					Description: `(Optional) A hostname to route network traffic to`,
				},
				resource.Attribute{
					Name:        "routed_to_ip",
					Description: `(Optional) An IP address to route network traffic to ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP Address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the IP eg. 127.0.0.1`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The Gateway address for the IP`,
				},
				resource.Attribute{
					Name:        "a_record",
					Description: `Public A record which points to the assigned IP address`,
				},
				resource.Attribute{
					Name:        "ptr_record",
					Description: `The PTR record that resolves to the IP address`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of IP`,
				},
				resource.Attribute{
					Name:        "routed_to",
					Description: `The hostname or IP that the IP routes to`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the IP ## Import IPs can be imported using the IP ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_ip.my_ip_address 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the IP Address`,
				},
				resource.Attribute{
					Name:        "address",
					Description: `The address of the IP eg. 127.0.0.1`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `The Gateway address for the IP`,
				},
				resource.Attribute{
					Name:        "a_record",
					Description: `Public A record which points to the assigned IP address`,
				},
				resource.Attribute{
					Name:        "ptr_record",
					Description: `The PTR record that resolves to the IP address`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of IP`,
				},
				resource.Attribute{
					Name:        "routed_to",
					Description: `The hostname or IP that the IP routes to`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the IP ## Import IPs can be imported using the IP ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_ip.my_ip_address 123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cherryservers_project",
			Category:         "Resources",
			ShortDescription: `Provides a CherryServers Project resource. This can be used to create, modify, and delete Projects. These projects can contain segregated Server and IPs that can be destroyed all at once by deleting the Project.`,
			Description:      ``,
			Keywords: []string{
				"project",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `(Required) The Server image ID or slug.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Server name. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "team_id",
					Description: `The ID of team that owns the Project`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The computed ID of the Project, to be used with other Resources ## Import Servers can be imported using the Project ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_project.myproject 8123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "team_id",
					Description: `The ID of team that owns the Project`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `The computed ID of the Project, to be used with other Resources ## Import Servers can be imported using the Project ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_project.myproject 8123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cherryservers_server",
			Category:         "Resources",
			ShortDescription: `Provides a CherryServers Server resource. This can be used to create, modify, and delete Servers. Servers also support provisioning.`,
			Description:      ``,
			Keywords: []string{
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "project_id",
					Description: `(Required) The Project ID to deploy the IP in.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region to deploy in. You can find these on [CherryServers API](https://api.cherryservers.com/doc/#tag/Regions).`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The Server image slug. You can find these on [CherryServers API](https://api.cherryservers.com/doc/#tag/Images/paths/~1v1~1plans~1{planId}~1images/get), after finding your plan_id.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) The Server name.`,
				},
				resource.Attribute{
					Name:        "plan_id",
					Description: `(Required) The unique slug that indentifies the type of Server. You can find a list of available slugs on [CherryServers API](https://api.cherryservers.com/doc/#tag/Plans/paths/~1v1~1teams~1{teamId}~1plans/get).`,
				},
				resource.Attribute{
					Name:        "ssh_key_ids",
					Description: `(Optional) A list of SSH IDs to enable in the format ` + "`" + `[12345, 123456]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip_addresses_ids",
					Description: `(Optional) A list of IP Address IDs to enable in the format ` + "`" + `[12345, 123456]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `(Optional) Base64 encoded User-Data blob. It should be either bash or cloud-config script.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) Key/value metadata for server tagging. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the Server`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the Server`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image of the Server`,
				},
				resource.Attribute{
					Name:        "primary_ip",
					Description: `The primary IPv4 address of the server. Servers will always have a primary ID in addition to any attached reserverd IPs`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IPv4 address of the server`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the server, such as "Pending"`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `The power state of the server, such as "Powered off"`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `The Server hourly price ## Import Servers can be imported using the Server ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_server.myserver 100823 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the Server`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `The hostname of the Server`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `The region of the Server`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image of the Server`,
				},
				resource.Attribute{
					Name:        "primary_ip",
					Description: `The primary IPv4 address of the server. Servers will always have a primary ID in addition to any attached reserverd IPs`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `The private IPv4 address of the server`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `The state of the server, such as "Pending"`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `The power state of the server, such as "Powered off"`,
				},
				resource.Attribute{
					Name:        "price",
					Description: `The Server hourly price ## Import Servers can be imported using the Server ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_server.myserver 100823 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cherryservers_ssh",
			Category:         "Resources",
			ShortDescription: `Provides a CherryServers SSH key resource. This can be used to create, and delete SSH keys associated with your account. Please note that you will not be able to add duplicate SSH keys to your account.`,
			Description:      ``,
			Keywords: []string{
				"ssh",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The Server name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key contents as a string to be added to your account ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH Key. Use this attribute when associating an SSH Key to a cherryservers\_server resource`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of your SSH Public key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date when this Key was added`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date when this Key was modified ## Import Servers can be imported using the SSH Key ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_ssh.mysshkey 900 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the SSH Key. Use this attribute when associating an SSH Key to a cherryservers\_server resource`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `The fingerprint of your SSH Public key`,
				},
				resource.Attribute{
					Name:        "created",
					Description: `The date when this Key was added`,
				},
				resource.Attribute{
					Name:        "updated",
					Description: `The date when this Key was modified ## Import Servers can be imported using the SSH Key ` + "`" + `id` + "`" + `, e.g. ` + "`" + `` + "`" + `` + "`" + ` terraform import cherryservers_ssh.mysshkey 900 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cherryservers_floating_ip": 0,
		"cherryservers_project":     1,
		"cherryservers_server":      2,
		"cherryservers_ssh":         3,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
