package cloudsigma

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_acl",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a CloudSigma ACL resource. Access Control Lists (ACLs) can be used to
grant permissions to another user to manage your resources. Permissions can be
granted on servers, drives, network resources, and firewall policies.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the ACL.`,
				},
				resource.Attribute{
					Name:        "permissions",
					Description: `(Optional) An array of strings containing the permissions. This may be "ATTACH", "CLONE", "EDIT", "LIST", "OPEN_VNC", "START", "STOP". ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ACL.`,
				},
				resource.Attribute{
					Name:        "onwer",
					Description: `The ownership of the ACL.`,
				},
				resource.Attribute{
					Name:        "resource_uri",
					Description: `The resource URI of the ACL.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the ACL.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the ACL.`,
				},
				resource.Attribute{
					Name:        "onwer",
					Description: `The ownership of the ACL.`,
				},
				resource.Attribute{
					Name:        "resource_uri",
					Description: `The resource URI of the ACL.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_drive",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a CloudSigma Drive resource which can be attached to a Server.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "clone_drive_id",
					Description: `(Optional) UUID of the drive that will be cloned`,
				},
				resource.Attribute{
					Name:        "media",
					Description: `(Required) Media representation type. It can be ` + "`" + `cdrom` + "`" + ` or ` + "`" + `disk` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human readable name of the drive`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of the drive in bytes`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of the tags UUIDs to be applied to the drive. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "media",
					Description: `(Required) Media representation type. It can be ` + "`" + `cdrom` + "`" + ` or ` + "`" + `disk` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human readable name of the drive`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of the drive in bytes`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "media",
					Description: `(Required) Media representation type. It can be ` + "`" + `cdrom` + "`" + ` or ` + "`" + `disk` + "`" + ``,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human readable name of the drive`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of the drive in bytes`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_firewall_policy",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the firewall policy ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_remote_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "drive",
					Description: `(Required) The UUID of the drive`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Required) The location of the remote snapshot`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the remote snapshot ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a CloudSigma Server resource. This can be used to create, modify,
and delete Servers.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `(Required) Server's CPU Clock speed measured in MHz`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) Server's RAM measured in bytes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Human readable name of server`,
				},
				resource.Attribute{
					Name:        "vnc_password",
					Description: `(Required) VNC Password to connect to server`,
				},
				resource.Attribute{
					Name:        "drive",
					Description: `(Optional) Drive attached to the server on creation - uuid - (Required) The UUID of the drive`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `(Optional) Network interface card attached to the server - ipv4_address - (Optional) The IP address reference. Only used with ` + "`" + `static` + "`" + ` type - type - (Optional) Configuration type. Valid values: ` + "`" + `dhcp` + "`" + `, ` + "`" + `static` + "`" + `, ` + "`" + `manual` + "`" + ` - vlan_uuid - (Optional) The UUID of the VLAN reference`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A list of the tags UUIDs to be applied to the server. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "cpu",
					Description: `Server's CPU Clock speed measured in MHz`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Server's RAM measured in bytes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human readable name of server`,
				},
				resource.Attribute{
					Name:        "vnc_password",
					Description: `VNC Password to connect to server`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `Server's CPU Clock speed measured in MHz`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Server's RAM measured in bytes`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Human readable name of server`,
				},
				resource.Attribute{
					Name:        "vnc_password",
					Description: `VNC Password to connect to server`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "drive",
					Description: `(Required) The UUID of the drive`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the snapshot ## Attributes Reference In addition to all above arguments, the following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_ssh_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a CloudSigma SSH key resource. to allow you to manage SSH keys. Keys
created with this resource can be referenced in your server configuration via
their IDs.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the SSH key for identification`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `(Optional) The private key. If this is a file, it can be read using the file interpolation function`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) The public key. If this is a file, it can be read using the file interpolation function ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The text of the private key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the key`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the SSH key`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `The text of the private key`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `The text of the public key`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "cloudsigma_tag",
			Category:         "Resources",
			ShortDescription: ``,
			Description: `

Provides a CloudSigma tag resource. A Tag is a label that can be applied to
a CloudSigma resource in order to better organize or facilitate the lookups
and actions on it. Tags created with this resource can be referenced in your
configurations via their ID.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the tag ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "onwer",
					Description: `The ownership of the tag.`,
				},
				resource.Attribute{
					Name:        "resource_uri",
					Description: `The resource URI of the tag.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the tag.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the tag.`,
				},
				resource.Attribute{
					Name:        "onwer",
					Description: `The ownership of the tag.`,
				},
				resource.Attribute{
					Name:        "resource_uri",
					Description: `The resource URI of the tag.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"cloudsigma_acl":             0,
		"cloudsigma_drive":           1,
		"cloudsigma_firewall_policy": 2,
		"cloudsigma_remote_snapshot": 3,
		"cloudsigma_server":          4,
		"cloudsigma_snapshot":        5,
		"cloudsigma_ssh_key":         6,
		"cloudsigma_tag":             7,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
