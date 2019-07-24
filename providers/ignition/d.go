package ignition

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ignition_config",
			Category:         "Data Sources",
			ShortDescription: `Renders an ignition configuration as JSON`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "disks",
					Description: `(Optional) The list of disks to be configured and their options.`,
				},
				resource.Attribute{
					Name:        "arrays",
					Description: `(Optional) The list of RAID arrays to be configured.`,
				},
				resource.Attribute{
					Name:        "filesystems",
					Description: `(Optional) The list of filesystems to be configured and/or used in the ` + "`" + `ignition_file` + "`" + `, ` + "`" + `ignition_directory` + "`" + `, and ` + "`" + `ignition_link` + "`" + ` resources.`,
				},
				resource.Attribute{
					Name:        "files",
					Description: `(Optional) The list of files to be written.`,
				},
				resource.Attribute{
					Name:        "directories",
					Description: `(Optional) The list of directories to be created.`,
				},
				resource.Attribute{
					Name:        "links",
					Description: `(Optional) The list of links to be created.`,
				},
				resource.Attribute{
					Name:        "systemd",
					Description: `(Optional) The list of systemd units. Describes the desired state of the systemd units.`,
				},
				resource.Attribute{
					Name:        "networkd",
					Description: `(Optional) The list of networkd units. Describes the desired state of the networkd files.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `(Optional) The list of accounts to be added.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) The list of groups to be added.`,
				},
				resource.Attribute{
					Name:        "append",
					Description: `(Optional) Any number of blocks with the configs to be appended to the current config.`,
				},
				resource.Attribute{
					Name:        "replace",
					Description: `(Optional) A block with config that will replace the current. The ` + "`" + `append` + "`" + ` and ` + "`" + `replace` + "`" + ` blocks supports:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The URL of the config. Supported schemes are http, https, tftp, s3, and data. When using http, it is advisable to use the verification option to ensure the contents haven't been modified.`,
				},
				resource.Attribute{
					Name:        "verification",
					Description: `(Optional) The hash of the config, in the form _\<type\>-\<value\>_ where type is sha512. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered template.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "rendered",
					Description: `The final rendered template.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_directory",
			Category:         "Data Sources",
			ShortDescription: `Describes a directory to be created in a particular filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filesystem",
					Description: `(Required) The internal identifier of the filesystem. This matches the last filesystem with the given identifier. This should be a valid name from a _ignition\_filesystem_ resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path to the directory.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The directory's permission mode. Note that the mode can be specified as a decimal value (i.e. 0755 -> 493) or an octal value(i.e 0755).`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) The user ID of the owner.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) The group ID of the owner. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_disk",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired state of a system’s disk.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device",
					Description: `(Required) The absolute path to the device. Devices are typically referenced by the _/dev/disk/by-`,
				},
				resource.Attribute{
					Name:        "wipe_table",
					Description: `(Optional) Whether or not the partition tables shall be wiped. When true, the partition tables are erased before any further manipulation. Otherwise, the existing entries are left intact.`,
				},
				resource.Attribute{
					Name:        "partition",
					Description: `(Optional) The list of partitions and their configuration for this particular disk.. The ` + "`" + `partition` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The PARTLABEL for the partition.`,
				},
				resource.Attribute{
					Name:        "number",
					Description: `(Optional) The partition number, which dictates it’s position in the partition table (one-indexed). If zero, use the next available partition slot.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the partition (in sectors). If zero, the partition will fill the remainder of the disk.`,
				},
				resource.Attribute{
					Name:        "start",
					Description: `(Optional) The start of the partition (in sectors). If zero, the partition will be positioned at the earliest available part of the disk.`,
				},
				resource.Attribute{
					Name:        "type_guid",
					Description: `(Optional) The GPT [partition type GUID](http://en.wikipedia.org/wiki/GUID_Partition_Table#Partition_type_GUIDs). If omitted, the default will be _0FC63DAF-8483-4772-8E79-3D69D8477DE4_ (Linux filesystem data). ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_file",
			Category:         "Data Sources",
			ShortDescription: `Describes a file to be written in a particular filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filesystem",
					Description: `(Required) The internal identifier of the filesystem. This matches the last filesystem with the given identifier. This should be a valid name from a _ignition\_filesystem_ resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path to the file.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) Block to provide the file content inline.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Optional) Block to retrieve the file content from a remote location. __Note__: ` + "`" + `content` + "`" + ` and ` + "`" + `source` + "`" + ` are mutually exclusive.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The file's permission mode. The mode must be properly specified as a decimal value (i.e. 0644 -> 420).`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) The user ID of the owner.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) The group ID of the owner. The ` + "`" + `content` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "mime",
					Description: `(Required) MIME format of the content (default _text/plain_).`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) Content of the file. The ` + "`" + `source` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The URL of the file contents. Supported schemes are http, https, tftp, s3, and [data][rfc2397]. When using http, it is advisable to use the verification option to ensure the contents haven't been modified.`,
				},
				resource.Attribute{
					Name:        "compression",
					Description: `(Optional) The type of compression used on the contents (null or gzip). Compression cannot be used with S3.`,
				},
				resource.Attribute{
					Name:        "verification",
					Description: `(Optional) The hash of the config, in the form _\<type\>-\<value\>_ where type is sha512. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_. [rfc2397]: https://tools.ietf.org/html/rfc2397`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_. [rfc2397]: https://tools.ietf.org/html/rfc2397`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_filesystem",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired state of a system’s filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The identifier for the filesystem, internal to Ignition. This is only required if the filesystem needs to be referenced in the a _ignition\_files_ resource.`,
				},
				resource.Attribute{
					Name:        "mount",
					Description: `(Optional) Contains the set of mount and formatting options for the filesystem. A non-null entry indicates that the filesystem should be mounted before it is used by Ignition.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The mount-point of the filesystem. A non-null entry indicates that the filesystem has already been mounted by the system at the specified path. This is really only useful for _/sysroot_. The ` + "`" + `mount` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `(Required) The absolute path to the device. Devices are typically referenced by the _/dev/disk/by-`,
				},
				resource.Attribute{
					Name:        "format",
					Description: `(Required) The filesystem format (ext4, btrfs, xfs, vfat, or swap).`,
				},
				resource.Attribute{
					Name:        "wipe_filesystem",
					Description: `(Optional) Whether or not to wipe the device before filesystem creation.`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Optional) The label of the filesystem.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `(Optional) The uuid of the filesystem.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `(Optional) Any additional options to be passed to the format-specific mkfs utility. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_group",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired group additions to the passwd database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The groupname for the account.`,
				},
				resource.Attribute{
					Name:        "password_hash",
					Description: `(Optional) The encrypted password for the account.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) The group ID of the new account. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_link",
			Category:         "Data Sources",
			ShortDescription: `Describes a link to be created in a particular filesystem.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filesystem",
					Description: `(Required) The internal identifier of the filesystem. This matches the last filesystem with the given identifier. This should be a valid name from a _ignition\_filesystem_ resource.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The absolute path to the link.`,
				},
				resource.Attribute{
					Name:        "target",
					Description: `(Required) The target path of the link.`,
				},
				resource.Attribute{
					Name:        "hard",
					Description: `(Optional) A symbolic link is created if this is false, a hard one if this is true.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) The user ID of the owner.`,
				},
				resource.Attribute{
					Name:        "gid",
					Description: `(Optional) The group ID of the owner. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_networkd_unit",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired state of the networkd units.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the file. This must be suffixed with a valid unit type (e.g. _00-eth0.network_).`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Required) The contents of the networkd file. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_raid",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired state of the system’s RAID.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name to use for the resulting md device.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Required) The redundancy level of the array (e.g. linear, raid1, raid5, etc.).`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Required) The list of devices (referenced by their absolute path) in the array.`,
				},
				resource.Attribute{
					Name:        "spares",
					Description: `(Optional) The number of spares (if applicable) in the array. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_systemd_unit",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired state of the systemd units.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the unit. This must be suffixed with a valid unit type (e.g. _thing.service_).`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Whether or not the service shall be enabled. When true, the service is enabled. In order for this to have any effect, the unit must have an install section. (default true)`,
				},
				resource.Attribute{
					Name:        "mask",
					Description: `(Optional) Whether or not the service shall be masked. When true, the service is masked by symlinking it to _/dev/null_.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The contents of the unit.`,
				},
				resource.Attribute{
					Name:        "dropin",
					Description: `(Optional) The list of drop-ins for the unit. The ` + "`" + `dropin` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the drop-in. This must be suffixed with _.conf_.`,
				},
				resource.Attribute{
					Name:        "content",
					Description: `(Optional) The contents of the drop-in. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ignition_user",
			Category:         "Data Sources",
			ShortDescription: `Describes the desired user additions to the passwd database.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The username for the account.`,
				},
				resource.Attribute{
					Name:        "password_hash",
					Description: `(Optional) The encrypted password for the account.`,
				},
				resource.Attribute{
					Name:        "ssh_authorized_keys",
					Description: `(Optional) A list of SSH keys to be added to the user’s authorized_keys.`,
				},
				resource.Attribute{
					Name:        "uid",
					Description: `(Optional) The user ID of the new account.`,
				},
				resource.Attribute{
					Name:        "gecos",
					Description: `(Optional) The GECOS field of the new account.`,
				},
				resource.Attribute{
					Name:        "home_dir",
					Description: `(Optional) The home directory of the new account.`,
				},
				resource.Attribute{
					Name:        "no_create_home",
					Description: `(Optional) Whether or not to create the user’s home directory.`,
				},
				resource.Attribute{
					Name:        "primary_group",
					Description: `(Optional) The name or ID of the primary group of the new account.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `(Optional) The list of supplementary groups of the new account.`,
				},
				resource.Attribute{
					Name:        "no_user_group",
					Description: `(Optional) Whether or not to create a group with the same name as the user.`,
				},
				resource.Attribute{
					Name:        "no_log_init",
					Description: `(Optional) Whether or not to add the user to the lastlog and faillog databases.`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `(Optional) The login shell of the new account.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `(Optional) Whether or not to make the account a system account. This only has an effect if the account doesn't exist yet. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID used to reference this resource in _ignition_config_.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ignition_config":        0,
		"ignition_directory":     1,
		"ignition_disk":          2,
		"ignition_file":          3,
		"ignition_filesystem":    4,
		"ignition_group":         5,
		"ignition_link":          6,
		"ignition_networkd_unit": 7,
		"ignition_raid":          8,
		"ignition_systemd_unit":  9,
		"ignition_user":          10,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
