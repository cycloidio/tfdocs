package ah

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_cloud_server",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "image",
					Description: `(Required) The Cloud Server image ID or Slug of the desired image for the server OR a Cloud Server Snapshot / Auto Backup ID. Changing this creates a new server. See the [list of available images](https://websa.advancedhosting.com/slugs).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name for the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Required) Datacenter ID or Slug to start the Cloud Server in. See the [list of available datacenters](https://websa.advancedhosting.com/slugs).`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Optional) Cloud Server Plan ID or Slug that indentifies the desired plan type of the Cloud Server. Changing this resizes the existing server. See the [list of available products](https://websa.advancedhosting.com/slugs).`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Boolean to enable or disable backups. Defaults to false.`,
				},
				resource.Attribute{
					Name:        "use_password",
					Description: `(Optional) Boolean defining if password should be generated for the server and sent by email. Defaults to true.`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `(Optional) Array of SSH IDs or fingerprints to enable in the format ` + "`" + `[12345, 7e:ac:a8:45:83:e3:58:f5:3a:9f:dd:16:63:dc:fb:1e]` + "`" + `. Fingerprints can be found in the 'SSH keys' section of the panel.`,
				},
				resource.Attribute{
					Name:        "create_public_ip_address",
					Description: `(Optional) Boolean defining if a new public IP address should be created for the server. This public IP address will become a primary IP address for the Cloud Server. Defaults to true. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of vCPUs on the the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `RAM of the Cloud Server in MiB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size of the Cloud Server in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Array of IP address blocks to be assigned to the Cloud Server. --- The ` + "`" + `ips` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP address assigned to the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP address type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean indicating a Primary IP flag.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `Reverse DNS assigned to the IP address.`,
				},
				resource.Attribute{
					Name:        "assignment_id",
					Description: `ID of the IP Address Assignment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the server.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "vcpu",
					Description: `Number of vCPUs on the the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `RAM of the Cloud Server in MiB.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Disk size of the Cloud Server in GB.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Array of IP address blocks to be assigned to the Cloud Server. --- The ` + "`" + `ips` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `Public IP address assigned to the Cloud Server.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `IP address type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Boolean indicating a Primary IP flag.`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `Reverse DNS assigned to the IP address.`,
				},
				resource.Attribute{
					Name:        "assignment_id",
					Description: `ID of the IP Address Assignment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_cloud_server_snapshot",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `(Required) Cloud Server ID to create a Snapshot from.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the snapshot. If not set, the snapshot name is assigned automatically based on date and time of snapshot creation. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "cloud_server_name",
					Description: `Cloud Server name the Snapshot was created for.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Snapshot size, in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type. Can be ` + "`" + `snapshot` + "`" + ` (for manual snapshots) or ` + "`" + `backup` + "`" + ` (for automatic backups)`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "cloud_server_name",
					Description: `Cloud Server name the Snapshot was created for.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Snapshot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Snapshot size, in GB`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type. Can be ` + "`" + `snapshot` + "`" + ` (for manual snapshots) or ` + "`" + `backup` + "`" + ` (for automatic backups)`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_ip",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "type",
					Description: `(Required) IP type. Can be either ` + "`" + `public` + "`" + ` or ` + "`" + `anycast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) Datacenter ID or Slug to create IP addresses. Required if ` + "`" + `type="public"` + "`" + `, ignored if ` + "`" + `type="anycast"` + "`" + `. See the [list of available datacenters](https://websa.advancedhosting.com/slugs).`,
				},
				resource.Attribute{
					Name:        "reverse_dns",
					Description: `(Optional) Reverse DNS to be assigned to the IP address. If not specified, will be automatically generated. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of IP address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address value.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the IP address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of IP address.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `IP address value.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the IP address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_ip_assignment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `(Required) Cloud Server ID to assign IP addresses to.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Required) IP address ID or IP address value to assign.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `(Optional) - Boolean for the Primary IP flag. Only one of assignments can have this flag set to true. Default value is false. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the IP Address Assignment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the IP Address Assignment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_private_network",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip_range",
					Description: `(Required) Private Network IP range in CIDR format.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the Private Network. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of Private Network.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Private Network.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Private Network.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of Private Network.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Private Network.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Private Network.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_private_network_connection",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `(Required) Cloud Server ID to connect to a Private Network.`,
				},
				resource.Attribute{
					Name:        "private_network_id",
					Description: `(Required) Private Network ID to connect a server to.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `(Optional) Private network IP address of the Cloud Server within the network. If not set, IP is assigned automatically. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the Private Network Connection.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the Private Network Connection.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_ssh_key",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) SSH key name.`,
				},
				resource.Attribute{
					Name:        "public_key",
					Description: `(Required) Public key. If this is a file, it can be read using the file interpolation function. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SSH key.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the SSH key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the SSH key.`,
				},
				resource.Attribute{
					Name:        "fingerprint",
					Description: `Fingerprint of the SSH key.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the SSH key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_volume",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Volume name`,
				},
				resource.Attribute{
					Name:        "product",
					Description: `(`,
				},
				resource.Attribute{
					Name:        "plan",
					Description: `(Optional) Volume Plan ID or Slug that indentifies the desired plan type of Volume. Changing this erases and recreates the volume. See the [list of available volume products](https://websa.advancedhosting.com/slugs).`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) Desired volume size in GB. Changing allowed to a greater value only. Changing this increases the volume size, data is preserved. Required unless ` + "`" + `origin_volume_id` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "file_system",
					Description: `(Optional) File system formatting option. Can be one of: ` + "`" + `ext4` + "`" + `, ` + "`" + `btrfs` + "`" + `, ` + "`" + `xfs` + "`" + `, or empty. If empty, volume is not formatted. Default value is ` + "`" + `ext4` + "`" + `. Changing this erases and recreates the volume.`,
				},
				resource.Attribute{
					Name:        "origin_volume_id",
					Description: `(Optional) ID of the volume to copy from. Changing this erases and recreates the volume. If this argument is set, ` + "`" + `size` + "`" + ` is ignored. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Volume`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Volume.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Volume.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the Volume`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of the Volume.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `Creation datetime of the Volume.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ah_ah_volume_attachment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cloud_server_id",
					Description: `(Required) Cloud Server ID to attach a Volume to.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Required) Volume ID to attach to a Cloud Server. --- ## Attributes Reference In addition to the arguments listed above, the following computed attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the Volume Attachment.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of attachment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Unique ID of the Volume Attachment.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `Current state of attachment.`,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"ah_ah_cloud_server":               0,
		"ah_ah_cloud_server_snapshot":      1,
		"ah_ah_ip":                         2,
		"ah_ah_ip_assignment":              3,
		"ah_ah_private_network":            4,
		"ah_ah_private_network_connection": 5,
		"ah_ah_ssh_key":                    6,
		"ah_ah_volume":                     7,
		"ah_ah_volume_attachment":          8,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
