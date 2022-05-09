package ionoscloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_backup_unit",
			Category:         "Managed Backup",
			ShortDescription: `Get Information on a IonosCloud Backup Unit`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"backup",
				"unit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing backup unit that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the backup unit you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Backup Unit.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Backup Unit.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address you want assigned to the backup unit.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The login associated with the backup unit. Derived from the contract number.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the Backup Unit.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Backup Unit.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address you want assigned to the backup unit.`,
				},
				resource.Attribute{
					Name:        "login",
					Description: `The login associated with the backup unit. Derived from the contract number.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_datacenter",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a IonosCloud Data Centers`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"datacenter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Id of an existing Virtual Data Center that you want to search for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing Virtual Data Center that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing Virtual Data Center's location. Either ` + "`" + `name` + "`" + `, ` + "`" + `location` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The regional location where the Virtual Data Center will be created`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of that Data Center. Gets incremented with every change`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `List of features supported by the location this data center is part of`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `Boolean value representing if the data center requires extra protection e.g. two factor protection`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Array of features and CPU families available in a location`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `A valid CPU family name`,
				},
				resource.Attribute{
					Name:        "max_cores",
					Description: `The maximum number of cores available`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum number of RAM in MB`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `A valid CPU vendor name`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The regional location where the Virtual Data Center will be created`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description for the Virtual Data Center`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `The version of that Data Center. Gets incremented with every change`,
				},
				resource.Attribute{
					Name:        "features",
					Description: `List of features supported by the location this data center is part of`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `Boolean value representing if the data center requires extra protection e.g. two factor protection`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Array of features and CPU families available in a location`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `A valid CPU family name`,
				},
				resource.Attribute{
					Name:        "max_cores",
					Description: `The maximum number of cores available`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum number of RAM in MB`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `A valid CPU vendor name`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_dbaas_pgsql_backups",
			Category:         "Database as a Service - Postgres",
			ShortDescription: `Get information on DbaaS PgSql Backups`,
			Description:      ``,
			Keywords: []string{
				"database",
				"as",
				"a",
				"service",
				"postgres",
				"dbaas",
				"pgsql",
				"backups",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Required) The unique ID of the cluster. ` + "`" + `cluster_id` + "`" + ` must be provided. If it is not provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_backups",
					Description: `List of backups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The unique ID of the cluster`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The friendly name of your cluster.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata of the resource.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `The ISO 8601 creation timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_date",
					Description: `The ISO 8601 modified timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Id of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_backups",
					Description: `List of backups.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The unique ID of the resource.`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `The unique ID of the cluster`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The friendly name of your cluster.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata of the resource.`,
				},
				resource.Attribute{
					Name:        "created_date",
					Description: `The ISO 8601 creation timestamp.`,
				},
				resource.Attribute{
					Name:        "last_modified_date",
					Description: `The ISO 8601 modified timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_dbaas_pgsql_cluster",
			Category:         "Database as a Service - Postgres",
			ShortDescription: `Get information on a DbaaS PgSql Cluster`,
			Description:      ``,
			Keywords: []string{
				"database",
				"as",
				"a",
				"service",
				"postgres",
				"dbaas",
				"pgsql",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "display_name",
					Description: `(Optional) Display name or an existing cluster that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the cluster you want to search for. Either ` + "`" + `display_name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "postgres_version",
					Description: `The PostgreSQL version of your cluster.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The total number of instances in the cluster (one master and n-1 standbys)`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of CPU cores per replica.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory per instance in megabytes.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The amount of storage per instance in MB.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The storage type used in your cluster.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Details about the network connection for your cluster.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The datacenter to connect your cluster to.`,
				},
				resource.Attribute{
					Name:        "lan_id",
					Description: `The LAN to connect your cluster to.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP and subnet for the database.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location where the cluster will be created. This will be where all of your instances live.`,
				},
				resource.Attribute{
					Name:        "backup_location",
					Description: `The S3 location where the backups will be stored.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The friendly name of your cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A weekly 4 hour-long window, during which maintenance might occur`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `Credentials for the database user to be created.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username for the initial postgres user.`,
				},
				resource.Attribute{
					Name:        "synchronization_mode",
					Description: `Represents different modes of replication.`,
				},
				resource.Attribute{
					Name:        "from_backup",
					Description: `The unique ID of the backup you want to restore.`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `The PostgreSQL version of your cluster.`,
				},
				resource.Attribute{
					Name:        "recovery_target_time",
					Description: `If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "postgres_version",
					Description: `The PostgreSQL version of your cluster.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `The total number of instances in the cluster (one master and n-1 standbys)`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The number of CPU cores per replica.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory per instance in megabytes.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The amount of storage per instance in MB.`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `The storage type used in your cluster.`,
				},
				resource.Attribute{
					Name:        "connections",
					Description: `Details about the network connection for your cluster.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The datacenter to connect your cluster to.`,
				},
				resource.Attribute{
					Name:        "lan_id",
					Description: `The LAN to connect your cluster to.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `The IP and subnet for the database.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location where the cluster will be created. This will be where all of your instances live.`,
				},
				resource.Attribute{
					Name:        "backup_location",
					Description: `The S3 location where the backups will be stored.`,
				},
				resource.Attribute{
					Name:        "display_name",
					Description: `The friendly name of your cluster.`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A weekly 4 hour-long window, during which maintenance might occur`,
				},
				resource.Attribute{
					Name:        "credentials",
					Description: `Credentials for the database user to be created.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The username for the initial postgres user.`,
				},
				resource.Attribute{
					Name:        "synchronization_mode",
					Description: `Represents different modes of replication.`,
				},
				resource.Attribute{
					Name:        "from_backup",
					Description: `The unique ID of the backup you want to restore.`,
				},
				resource.Attribute{
					Name:        "backup_id",
					Description: `The PostgreSQL version of your cluster.`,
				},
				resource.Attribute{
					Name:        "recovery_target_time",
					Description: `If this value is supplied as ISO 8601 timestamp, the backup will be replayed up until the given timestamp.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_dbaas_pgsql_versions",
			Category:         "Database as a Service - Postgres",
			ShortDescription: `Get information on DbaaS PgSql Versions`,
			Description:      ``,
			Keywords: []string{
				"database",
				"as",
				"a",
				"service",
				"postgres",
				"dbaas",
				"pgsql",
				"versions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `(Optional) The unique ID of the cluster. If ` + "`" + `cluster_id` + "`" + ` is not provided the data source will return the list of postgres version for all cluster. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Id of the cluster`,
				},
				resource.Attribute{
					Name:        "postgres_versions",
					Description: `list of PostgreSQL versions.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "cluster_id",
					Description: `Id of the cluster`,
				},
				resource.Attribute{
					Name:        "postgres_versions",
					Description: `list of PostgreSQL versions.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_firewall",
			Category:         "Compute Engine",
			ShortDescription: `Get Information on a IonosCloud Firewall`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"firewall",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing firewall rule that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the firewall rule you want to search for.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) The Virtual Data Center ID.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `(Required) The Server ID.`,
				},
				resource.Attribute{
					Name:        "nic_id",
					Description: `(Required) The NIC ID. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the firewall rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the firewall rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the rule: TCP, UDP, ICMP, ANY. This property is immutable.`,
				},
				resource.Attribute{
					Name:        "source_mac",
					Description: `Only traffic originating from the respective MAC address is allowed.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Only traffic originating from the respective IPv4 address is allowed.`,
				},
				resource.Attribute{
					Name:        "target_ip",
					Description: `Only traffic directed to the respective IP address of the NIC is allowed.`,
				},
				resource.Attribute{
					Name:        "port_range_start",
					Description: `Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen.`,
				},
				resource.Attribute{
					Name:        "port_range_end",
					Description: `Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the firewall rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the firewall rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `The protocol for the rule: TCP, UDP, ICMP, ANY. This property is immutable.`,
				},
				resource.Attribute{
					Name:        "source_mac",
					Description: `Only traffic originating from the respective MAC address is allowed.`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Only traffic originating from the respective IPv4 address is allowed.`,
				},
				resource.Attribute{
					Name:        "target_ip",
					Description: `Only traffic directed to the respective IP address of the NIC is allowed.`,
				},
				resource.Attribute{
					Name:        "port_range_start",
					Description: `Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen.`,
				},
				resource.Attribute{
					Name:        "port_range_end",
					Description: `Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen.`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen.`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_group",
			Category:         "User Management",
			ShortDescription: `Get information on a Ionos Cloud Groups`,
			Description:      ``,
			Keywords: []string{
				"user",
				"management",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing group that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the group you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A name for the group.`,
				},
				resource.Attribute{
					Name:        "create_datacenter",
					Description: `The group will be allowed to create virtual data centers.`,
				},
				resource.Attribute{
					Name:        "create_snapshot",
					Description: `The group will be allowed to create snapshots.`,
				},
				resource.Attribute{
					Name:        "reserve_ip",
					Description: `The group will be allowed to reserve IP addresses.`,
				},
				resource.Attribute{
					Name:        "access_activity_log",
					Description: `The group will be allowed to access the activity log.`,
				},
				resource.Attribute{
					Name:        "create_pcc",
					Description: `The group will be allowed to create pcc privilege.`,
				},
				resource.Attribute{
					Name:        "s3_privilege",
					Description: `The group will have S3 privilege.`,
				},
				resource.Attribute{
					Name:        "create_backup_unit",
					Description: `The group will be allowed to create backup unit privilege.`,
				},
				resource.Attribute{
					Name:        "create_internet_access",
					Description: `The group will be allowed to create internet access privilege.`,
				},
				resource.Attribute{
					Name:        "create_k8s_cluster",
					Description: `The group will be allowed to create kubernetes cluster privilege.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of users in group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A name for the group.`,
				},
				resource.Attribute{
					Name:        "create_datacenter",
					Description: `The group will be allowed to create virtual data centers.`,
				},
				resource.Attribute{
					Name:        "create_snapshot",
					Description: `The group will be allowed to create snapshots.`,
				},
				resource.Attribute{
					Name:        "reserve_ip",
					Description: `The group will be allowed to reserve IP addresses.`,
				},
				resource.Attribute{
					Name:        "access_activity_log",
					Description: `The group will be allowed to access the activity log.`,
				},
				resource.Attribute{
					Name:        "create_pcc",
					Description: `The group will be allowed to create pcc privilege.`,
				},
				resource.Attribute{
					Name:        "s3_privilege",
					Description: `The group will have S3 privilege.`,
				},
				resource.Attribute{
					Name:        "create_backup_unit",
					Description: `The group will be allowed to create backup unit privilege.`,
				},
				resource.Attribute{
					Name:        "create_internet_access",
					Description: `The group will be allowed to create internet access privilege.`,
				},
				resource.Attribute{
					Name:        "create_k8s_cluster",
					Description: `The group will be allowed to create kubernetes cluster privilege.`,
				},
				resource.Attribute{
					Name:        "users",
					Description: `List of users in group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_image",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a IonosCloud Images`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of an existing image that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Id of the existing image's location.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The image type, HDD or CD-ROM.`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `(Optional) Cloud init compatibility ("NONE" or "V1") If both "name" and "version" are provided the plugin will concatenate the two strings in this format [name]-[version]. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the image`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This indicates the type of image`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the image`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the image`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `description of the image`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "license_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `This indicates the type of image`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_ipblock",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a IonosCloud Ip Block`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"ipblock",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of Ip Block`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of Ip Block`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The regional location for this IP Block: us/las, us/ewr, de/fra, de/fkb.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The number of IP addresses to reserve for this block.`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `The list of IP addresses associated with this block.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_ipfailover",
			Category:         "Compute Engine",
			ShortDescription: `Get Information on ipfailover objects.`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"ipfailover",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) The ID of the datacenter containing the ip failover datasource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required) The id of the lan of which the IP failover belongs ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of a Data Center.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The reserved IP address to be used in the IP failover group.`,
				},
				resource.Attribute{
					Name:        "lan_id",
					Description: `The ID of a LAN.`,
				},
				resource.Attribute{
					Name:        "nicuuid",
					Description: `The ID of a NIC.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of a Data Center.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `The reserved IP address to be used in the IP failover group.`,
				},
				resource.Attribute{
					Name:        "lan_id",
					Description: `The ID of a LAN.`,
				},
				resource.Attribute{
					Name:        "nicuuid",
					Description: `The ID of a NIC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_k8s_cluster",
			Category:         "Managed Kubernetes",
			ShortDescription: `Get information on a IonosCloud K8s Cluster`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"kubernetes",
				"k8s",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name or an existing cluster that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the cluster you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the cluster`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubernetes version`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A maintenance window comprise of a day of the week and a time for maintenance to be allowed`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `A clock time in the day when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "day_of_the_week",
					Description: `Day of the week when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `A list of available versions for upgrading the cluster`,
				},
				resource.Attribute{
					Name:        "viable_node_pool_versions",
					Description: `A list of versions that may be used for node pools under this cluster`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `The indicator if the cluster is public or private. Be aware that setting it to false is currently in beta phase`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `The IP address of the gateway used by the cluster. This is mandatory when ` + "`" + `public` + "`" + ` is set to ` + "`" + `false` + "`" + ` and should not be provided otherwise.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `list of the IDs of the node pools in this cluster`,
				},
				resource.Attribute{
					Name:        "api_subnet_allow_list",
					Description: `access to the K8s API server is restricted to these CIDRs`,
				},
				resource.Attribute{
					Name:        "s3_buckets",
					Description: `list of S3 bucket configured for K8s usage`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Kubernetes configuration`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `structured kubernetes config consisting of a list with 1 item with the following fields:`,
				},
				resource.Attribute{
					Name:        "user_tokens",
					Description: `a convenience map to be search the token of a specific user - key - is the user name - value - is the token`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `cluster server (same as ` + "`" + `config[0].clusters[0].cluster.server` + "`" + ` but provided as an attribute for ease of use)`,
				},
				resource.Attribute{
					Name:        "ca_crt",
					Description: `base64 decoded cluster certificate authority data (provided as an attribute for direct use)`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the cluster`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the cluster`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `Kubernetes version`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A maintenance window comprise of a day of the week and a time for maintenance to be allowed`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `A clock time in the day when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "day_of_the_week",
					Description: `Day of the week when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `A list of available versions for upgrading the cluster`,
				},
				resource.Attribute{
					Name:        "viable_node_pool_versions",
					Description: `A list of versions that may be used for node pools under this cluster`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `The indicator if the cluster is public or private. Be aware that setting it to false is currently in beta phase`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `The IP address of the gateway used by the cluster. This is mandatory when ` + "`" + `public` + "`" + ` is set to ` + "`" + `false` + "`" + ` and should not be provided otherwise.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "node_pools",
					Description: `list of the IDs of the node pools in this cluster`,
				},
				resource.Attribute{
					Name:        "api_subnet_allow_list",
					Description: `access to the K8s API server is restricted to these CIDRs`,
				},
				resource.Attribute{
					Name:        "s3_buckets",
					Description: `list of S3 bucket configured for K8s usage`,
				},
				resource.Attribute{
					Name:        "kube_config",
					Description: `Kubernetes configuration`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `structured kubernetes config consisting of a list with 1 item with the following fields:`,
				},
				resource.Attribute{
					Name:        "user_tokens",
					Description: `a convenience map to be search the token of a specific user - key - is the user name - value - is the token`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `cluster server (same as ` + "`" + `config[0].clusters[0].cluster.server` + "`" + ` but provided as an attribute for ease of use)`,
				},
				resource.Attribute{
					Name:        "ca_crt",
					Description: `base64 decoded cluster certificate authority data (provided as an attribute for direct use)`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_k8s_node_pool",
			Category:         "Managed Kubernetes",
			ShortDescription: `Get information on a IonosCloud K8s Node Pool`,
			Description:      ``,
			Keywords: []string{
				"managed",
				"kubernetes",
				"k8s",
				"node",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing node pool that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the node pool you want to search for. ` + "`" + `k8s_cluster_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `id of the node pool`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the node pool`,
				},
				resource.Attribute{
					Name:        "k8s_cluster_id",
					Description: `ID of the cluster this node pool is part of`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The UUID of the VDC`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes in this node pool`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `CPU Family`,
				},
				resource.Attribute{
					Name:        "cores_count",
					Description: `CPU cores count`,
				},
				resource.Attribute{
					Name:        "ram_size",
					Description: `The amount of RAM in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The compute availability zone in which the nodes should exist`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `HDD or SDD`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The size of the volume in GB. The size should be greater than 10GB.`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The kubernetes version`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A maintenance window comprise of a day of the week and a time for maintenance to be allowed`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `A clock time in the day when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "day_of_the_week",
					Description: `Day of the week when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `The range defining the minimum and maximum number of worker nodes that the managed node group can scale in`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `The minimum number of worker nodes the node pool can scale down to`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `The maximum number of worker nodes that the node pool can scale to`,
				},
				resource.Attribute{
					Name:        "lans",
					Description: `A list of Local Area Networks the node pool is a part of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The LAN ID of an existing LAN at the related datacenter`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `An array of additional LANs attached to worker nodes`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `IPv4 or IPv6 CIDR to be routed via the interface`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `IPv4 or IPv6 Gateway IP for the route`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `A map of annotations in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `A list of kubernetes versions available for upgrade`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `The list of fixed IPs associated with this node pool`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `Public IP address for the gateway performing source NAT for the node pool's nodes belonging to a private cluster. Required only if the node pool belongs to a private cluster.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `id of the node pool`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `name of the node pool`,
				},
				resource.Attribute{
					Name:        "k8s_cluster_id",
					Description: `ID of the cluster this node pool is part of`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The UUID of the VDC`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `one of "AVAILABLE", "INACTIVE", "BUSY", "DEPLOYING", "ACTIVE", "FAILED", "SUSPENDED", "FAILED_SUSPENDED", "UPDATING", "FAILED_UPDATING", "DESTROYING", "FAILED_DESTROYING", "TERMINATED"`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes in this node pool`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `CPU Family`,
				},
				resource.Attribute{
					Name:        "cores_count",
					Description: `CPU cores count`,
				},
				resource.Attribute{
					Name:        "ram_size",
					Description: `The amount of RAM in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The compute availability zone in which the nodes should exist`,
				},
				resource.Attribute{
					Name:        "storage_type",
					Description: `HDD or SDD`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The size of the volume in GB. The size should be greater than 10GB.`,
				},
				resource.Attribute{
					Name:        "k8s_version",
					Description: `The kubernetes version`,
				},
				resource.Attribute{
					Name:        "maintenance_window",
					Description: `A maintenance window comprise of a day of the week and a time for maintenance to be allowed`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `A clock time in the day when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "day_of_the_week",
					Description: `Day of the week when maintenance is allowed`,
				},
				resource.Attribute{
					Name:        "auto_scaling",
					Description: `The range defining the minimum and maximum number of worker nodes that the managed node group can scale in`,
				},
				resource.Attribute{
					Name:        "min_node_count",
					Description: `The minimum number of worker nodes the node pool can scale down to`,
				},
				resource.Attribute{
					Name:        "max_node_count",
					Description: `The maximum number of worker nodes that the node pool can scale to`,
				},
				resource.Attribute{
					Name:        "lans",
					Description: `A list of Local Area Networks the node pool is a part of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The LAN ID of an existing LAN at the related datacenter`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the Kubernetes Node Pool LAN will reserve an IP using DHCP`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `An array of additional LANs attached to worker nodes`,
				},
				resource.Attribute{
					Name:        "network",
					Description: `IPv4 or IPv6 CIDR to be routed via the interface`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `IPv4 or IPv6 Gateway IP for the route`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `A map of labels in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "annotations",
					Description: `A map of annotations in the form of key -> value`,
				},
				resource.Attribute{
					Name:        "available_upgrade_versions",
					Description: `A list of kubernetes versions available for upgrade`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `The list of fixed IPs associated with this node pool`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `Public IP address for the gateway performing source NAT for the node pool's nodes belonging to a private cluster. Required only if the node pool belongs to a private cluster.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_lan",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a Ionos Cloud Lans`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"lan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) Datacenter's UUID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing lan that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the lan you want to search for. ` + "`" + `datacenter_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the LAN.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of lan's Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "ip_failover",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "pcc",
					Description: `The unique id of a ` + "`" + `ionoscloud_private_crossconnect` + "`" + ` resource, in order.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the LAN faces the public Internet (true) or not (false).`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the LAN.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of lan's Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "ip_failover",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "pcc",
					Description: `The unique id of a ` + "`" + `ionoscloud_private_crossconnect` + "`" + ` resource, in order.`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the LAN faces the public Internet (true) or not (false).`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_location",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a IonosCloud Locations`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"location",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the location to search for.`,
				},
				resource.Attribute{
					Name:        "feature",
					Description: `(Optional) A desired feature that the location must be able to provide. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the location`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Array of features and CPU families available in a location`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `A valid CPU family name.`,
				},
				resource.Attribute{
					Name:        "max_cores",
					Description: `The maximum number of cores available.`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum number of RAM in MB.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `A valid CPU vendor name.`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases available for the location`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the location`,
				},
				resource.Attribute{
					Name:        "cpu_architecture",
					Description: `Array of features and CPU families available in a location`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `A valid CPU family name.`,
				},
				resource.Attribute{
					Name:        "max_cores",
					Description: `The maximum number of cores available.`,
				},
				resource.Attribute{
					Name:        "max_ram",
					Description: `The maximum number of RAM in MB.`,
				},
				resource.Attribute{
					Name:        "vendor",
					Description: `A valid CPU vendor name.`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases available for the location`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_natgateway",
			Category:         "NAT Gateway",
			ShortDescription: `Get information on a Nat Gateway`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"gateway",
				"natgateway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) Datacenter's UUID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing network load balancer forwarding rule that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the network load balancer forwarding rule you want to search for. ` + "`" + `datacenter_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of that natgateway`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of that natgateway`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Collection of public IP addresses of the NAT gateway. Should be customer reserved IP addresses in that location`,
				},
				resource.Attribute{
					Name:        "lans",
					Description: `Collection of LANs connected to the NAT gateway. IPs must contain valid subnet mask. If user will not provide any IP then system will generate an IP with /24 subnet.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id for the LAN connected to the NAT gateway`,
				},
				resource.Attribute{
					Name:        "gateway_ips",
					Description: `Collection of gateway IP addresses of the NAT gateway. Will be auto-generated if not provided. Should ideally be an IP belonging to the same subnet as the LAN`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of that natgateway`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of that natgateway`,
				},
				resource.Attribute{
					Name:        "public_ips",
					Description: `Collection of public IP addresses of the NAT gateway. Should be customer reserved IP addresses in that location`,
				},
				resource.Attribute{
					Name:        "lans",
					Description: `Collection of LANs connected to the NAT gateway. IPs must contain valid subnet mask. If user will not provide any IP then system will generate an IP with /24 subnet.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id for the LAN connected to the NAT gateway`,
				},
				resource.Attribute{
					Name:        "gateway_ips",
					Description: `Collection of gateway IP addresses of the NAT gateway. Will be auto-generated if not provided. Should ideally be an IP belonging to the same subnet as the LAN`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_natgateway_rule",
			Category:         "NAT Gateway",
			ShortDescription: `Get information on a Nat Gateway Rule`,
			Description:      ``,
			Keywords: []string{
				"nat",
				"gateway",
				"natgateway",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) Datacenter's UUID.`,
				},
				resource.Attribute{
					Name:        "natgateway_id",
					Description: `(Required) Nat Gateway's UUID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing NAT gateway rule that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the NAT gateway rule you want to search for. Both ` + "`" + `datacenter_id` + "`" + ` and ` + "`" + `natgateway_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the NAT gateway rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway rule`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `ype of the NAT gateway rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.`,
				},
				resource.Attribute{
					Name:        "source_subnet",
					Description: `Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource`,
				},
				resource.Attribute{
					Name:        "target_subnet",
					Description: `Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.`,
				},
				resource.Attribute{
					Name:        "target_port_range",
					Description: `Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the NAT gateway rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the NAT gateway rule`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `ype of the NAT gateway rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the NAT gateway rule. Defaults to ALL. If protocol is 'ICMP' then targetPortRange start and end cannot be set.`,
				},
				resource.Attribute{
					Name:        "source_subnet",
					Description: `Source subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets source IP address.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Public IP address of the NAT gateway rule. Specifies the address used for masking outgoing packets source address field. Should be one of the customer reserved IP address already configured on the NAT gateway resource`,
				},
				resource.Attribute{
					Name:        "target_subnet",
					Description: `Target or destination subnet of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on the packets destination IP address. If none is provided, rule will match any address.`,
				},
				resource.Attribute{
					Name:        "target_port_range",
					Description: `Target port range of the NAT gateway rule. For SNAT rules it specifies which packets this translation rule applies to based on destination port. If none is provided, rule will match any port`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_networkloadbalancer",
			Category:         "Network Load Balancer",
			ShortDescription: `Get information on a Network Load Balancer`,
			Description:      ``,
			Keywords: []string{
				"network",
				"load",
				"balancer",
				"networkloadbalancer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) Datacenter's UUID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing network load balancer that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the network load balancer you want to search for. ` + "`" + `datacenter_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of that Network Load Balancer`,
				},
				resource.Attribute{
					Name:        "listener_lan",
					Description: `Id of the listening LAN. (inbound)`,
				},
				resource.Attribute{
					Name:        "target_lan",
					Description: `Id of the balanced private target LAN. (outbound)`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses of the Network Load Balancer. (inbound and outbound) IP of the listenerLan must be a customer reserved IP for the public load balancer and private IP for the private load balancer.`,
				},
				resource.Attribute{
					Name:        "lb_private_ips",
					Description: `Collection of private IP addresses with subnet mask of the Network Load Balancer. IPs must contain valid subnet mask. If user will not provide any IP then the system will generate one IP with /24 subnet.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of that Network Load Balancer`,
				},
				resource.Attribute{
					Name:        "listener_lan",
					Description: `Id of the listening LAN. (inbound)`,
				},
				resource.Attribute{
					Name:        "target_lan",
					Description: `Id of the balanced private target LAN. (outbound)`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses of the Network Load Balancer. (inbound and outbound) IP of the listenerLan must be a customer reserved IP for the public load balancer and private IP for the private load balancer.`,
				},
				resource.Attribute{
					Name:        "lb_private_ips",
					Description: `Collection of private IP addresses with subnet mask of the Network Load Balancer. IPs must contain valid subnet mask. If user will not provide any IP then the system will generate one IP with /24 subnet.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_networkloadbalancer_forwardingrule",
			Category:         "Network Load Balancer",
			ShortDescription: `Get information on a Network Load Balancer Forwarding Rule`,
			Description:      ``,
			Keywords: []string{
				"network",
				"load",
				"balancer",
				"networkloadbalancer",
				"forwardingrule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) Datacenter's UUID.`,
				},
				resource.Attribute{
					Name:        "networkloadbalancer_id",
					Description: `(Required) Network Load Balancer's UUID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing network load balancer forwarding rule that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the network load balancer forwarding rule you want to search for. Both ` + "`" + `datacenter_id` + "`" + ` and ` + "`" + `networkloadbalancer_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of that Network Load Balancer forwarding rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of that Network Load Balancer forwarding rule.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `Algorithm for the balancing.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the balancing.`,
				},
				resource.Attribute{
					Name:        "listener_ip",
					Description: `Listening IP. (inbound)`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listening port number. (inbound) (range: 1 to 65535)`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health check attributes for Network Load Balancer forwarding rule.`,
				},
				resource.Attribute{
					Name:        "client_timeout",
					Description: `ClientTimeout is expressed in milliseconds. This inactivity timeout applies when the client is expected to acknowledge or send data. If unset the default of 50 seconds will be used.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `It specifies the maximum time (in milliseconds) to wait for a connection attempt to a target VM to succeed. If unset, the default of 5 seconds will be used.`,
				},
				resource.Attribute{
					Name:        "target_timeout",
					Description: `TargetTimeout specifies the maximum inactivity time (in milliseconds) on the target VM side. If unset, the default of 50 seconds will be used.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `Retries specifies the number of retries to perform on a target VM after a connection failure. If unset, the default value of 3 will be used.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `Array of items in that collection.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of a balanced target VM.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the balanced target service. (range: 1 to 65535).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight parameter is used to adjust the target VM's weight relative to other target VMs.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health check attributes for Network Load Balancer forwarding rule target.`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `Check specifies whether the target VM's health is checked.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `CheckInterval determines the duration (in milliseconds) between consecutive health checks. If unspecified a default of 2000 ms is used.`,
				},
				resource.Attribute{
					Name:        "maintenance",
					Description: `Maintenance specifies if a target VM should be marked as down, even if it is not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of that Network Load Balancer forwarding rule.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of that Network Load Balancer forwarding rule.`,
				},
				resource.Attribute{
					Name:        "algorithm",
					Description: `Algorithm for the balancing.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol of the balancing.`,
				},
				resource.Attribute{
					Name:        "listener_ip",
					Description: `Listening IP. (inbound)`,
				},
				resource.Attribute{
					Name:        "listener_port",
					Description: `Listening port number. (inbound) (range: 1 to 65535)`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health check attributes for Network Load Balancer forwarding rule.`,
				},
				resource.Attribute{
					Name:        "client_timeout",
					Description: `ClientTimeout is expressed in milliseconds. This inactivity timeout applies when the client is expected to acknowledge or send data. If unset the default of 50 seconds will be used.`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `It specifies the maximum time (in milliseconds) to wait for a connection attempt to a target VM to succeed. If unset, the default of 5 seconds will be used.`,
				},
				resource.Attribute{
					Name:        "target_timeout",
					Description: `TargetTimeout specifies the maximum inactivity time (in milliseconds) on the target VM side. If unset, the default of 50 seconds will be used.`,
				},
				resource.Attribute{
					Name:        "retries",
					Description: `Retries specifies the number of retries to perform on a target VM after a connection failure. If unset, the default value of 3 will be used.`,
				},
				resource.Attribute{
					Name:        "targets",
					Description: `Array of items in that collection.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP of a balanced target VM.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port of the balanced target service. (range: 1 to 65535).`,
				},
				resource.Attribute{
					Name:        "weight",
					Description: `Weight parameter is used to adjust the target VM's weight relative to other target VMs.`,
				},
				resource.Attribute{
					Name:        "health_check",
					Description: `Health check attributes for Network Load Balancer forwarding rule target.`,
				},
				resource.Attribute{
					Name:        "check",
					Description: `Check specifies whether the target VM's health is checked.`,
				},
				resource.Attribute{
					Name:        "check_interval",
					Description: `CheckInterval determines the duration (in milliseconds) between consecutive health checks. If unspecified a default of 2000 ms is used.`,
				},
				resource.Attribute{
					Name:        "maintenance",
					Description: `Maintenance specifies if a target VM should be marked as down, even if it is not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_nic",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a Ionos Cloud NIC`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"nic",
			},
			Arguments: []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the NIC.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of a Virtual Data Center.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The ID of a server.`,
				},
				resource.Attribute{
					Name:        "lan",
					Description: `The LAN ID the NIC will sit on.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the LAN.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the NIC should get an IP address using DHCP (true) or not (false).`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses assigned to a nic. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.`,
				},
				resource.Attribute{
					Name:        "firewall_active",
					Description: `If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.`,
				},
				resource.Attribute{
					Name:        "firewall_type",
					Description: `The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_private_crossconnect",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a Ionos Cloud Private Crossconnects`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"private",
				"crossconnect",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing private crossconnect that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the private crossconnect you want to search for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the found private cross connect`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of private cross connect`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of private cross connect`,
				},
				resource.Attribute{
					Name:        "peers",
					Description: `Lists LAN's joined to this private cross connect`,
				},
				resource.Attribute{
					Name:        "lan_id",
					Description: `The id of the cross-connected LAN`,
				},
				resource.Attribute{
					Name:        "lan_name",
					Description: `The name of the cross-connected LAN`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The id of the cross-connected datacenter`,
				},
				resource.Attribute{
					Name:        "datacenter_name",
					Description: `The name of the cross-connected datacenter`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the cross-connected datacenter`,
				},
				resource.Attribute{
					Name:        "connectable_datacenters",
					Description: `Lists datacenters that can be joined to this private cross connect`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the connectable datacenter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the connectable datacenter`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location of the connectable datacenter`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of the found private cross connect`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of private cross connect`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of private cross connect`,
				},
				resource.Attribute{
					Name:        "peers",
					Description: `Lists LAN's joined to this private cross connect`,
				},
				resource.Attribute{
					Name:        "lan_id",
					Description: `The id of the cross-connected LAN`,
				},
				resource.Attribute{
					Name:        "lan_name",
					Description: `The name of the cross-connected LAN`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The id of the cross-connected datacenter`,
				},
				resource.Attribute{
					Name:        "datacenter_name",
					Description: `The name of the cross-connected datacenter`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The location of the cross-connected datacenter`,
				},
				resource.Attribute{
					Name:        "connectable_datacenters",
					Description: `Lists datacenters that can be joined to this private cross connect`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the connectable datacenter`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the connectable datacenter`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `The physical location of the connectable datacenter`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_resource",
			Category:         "User Management",
			ShortDescription: `Get information on a IonosCloud Resource`,
			Description:      ``,
			Keywords: []string{
				"user",
				"management",
				"resource",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The specific type of resources to retrieve information about.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) The ID of the specific resource to retrieve information about. ## Attributes Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Resource`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the Resource`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_s3_key",
			Category:         "User Management",
			ShortDescription: `Get Information on a IonosCloud s3 key`,
			Description:      ``,
			Keywords: []string{
				"user",
				"management",
				"s3",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the s3 key`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `The state of the s3 key`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user that owns the key`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Computed)The S3 Secret key.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the s3 key`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `The state of the s3 key`,
				},
				resource.Attribute{
					Name:        "user_id",
					Description: `The ID of the user that owns the key`,
				},
				resource.Attribute{
					Name:        "secret_key",
					Description: `(Computed)The S3 Secret key.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_server",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a Ionos Cloud Servers`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"server",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) Datacenter's UUID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing server that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the server you want to search for. ` + "`" + `datacenter_id` + "`" + ` and either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both of ` + "`" + `name` + "`" + ` and ` + "`" + `id` + "`" + ` are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `The UUID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of that resource`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of that resource`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server usages: ENTERPRISE or CUBE`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The id of the datacenter`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The total number of cores for the server`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `CPU architecture on which server gets provisione`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory for the server in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the server should exist`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `Status of the virtual Machine`,
				},
				resource.Attribute{
					Name:        "cdroms",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of cdrom`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of image`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached volume`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Hardware type of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the volume should exist`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Image or snapshot ID to be used as template for this volume`,
				},
				resource.Attribute{
					Name:        "image_password",
					Description: `Initial password to be set for installed OS`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The bus type of the volume`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this volume`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The Logical Unit Number of the storage volume`,
				},
				resource.Attribute{
					Name:        "pci_slot",
					Description: `The PCI slot number of the storage volume`,
				},
				resource.Attribute{
					Name:        "backup_unit_id",
					Description: `The uuid of the Backup Unit that user has access to`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The cloud-init configuration for the volume as base64 encoded string`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached nic`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached nid`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses assigned to a nic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the nic will reserve an IP using DHCP`,
				},
				resource.Attribute{
					Name:        "lan",
					Description: `The LAN ID the NIC will sit on`,
				},
				resource.Attribute{
					Name:        "firewall_active",
					Description: `Activate or deactivate the firewall`,
				},
				resource.Attribute{
					Name:        "firewall_type",
					Description: `The type of firewall rules that will be allowed on the NIC`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The Logical Unit Number (LUN) of the storage volume`,
				},
				resource.Attribute{
					Name:        "pci_slot",
					Description: `The PCI slot number of the Nic`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the firewall rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the firewall rule`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `he protocol for the rule`,
				},
				resource.Attribute{
					Name:        "source_mac",
					Description: `Only traffic originating from the respective MAC address is allowed`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs`,
				},
				resource.Attribute{
					Name:        "target_ip",
					Description: `In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `Defines the allowed code (from 0 to 254) if protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_start",
					Description: `Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_end",
					Description: `Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of firewall rule`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "template_uuid",
					Description: `The UUID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of that resource`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of that resource`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Server usages: ENTERPRISE or CUBE`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The id of the datacenter`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `The total number of cores for the server`,
				},
				resource.Attribute{
					Name:        "cpu_family",
					Description: `CPU architecture on which server gets provisione`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The amount of memory for the server in MB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the server should exist`,
				},
				resource.Attribute{
					Name:        "vm_state",
					Description: `Status of the virtual Machine`,
				},
				resource.Attribute{
					Name:        "cdroms",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the attached cdrom`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description of cdrom`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this Image`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Type of image`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "public",
					Description: `Indicates if the image is part of the public repository or not`,
				},
				resource.Attribute{
					Name:        "image_aliases",
					Description: `List of image aliases mapped for this Image`,
				},
				resource.Attribute{
					Name:        "cloud_init",
					Description: `Cloud init compatibility`,
				},
				resource.Attribute{
					Name:        "volumes",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached volume`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached volume`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Hardware type of the volume.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GB`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone in which the volume should exist`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `Image or snapshot ID to be used as template for this volume`,
				},
				resource.Attribute{
					Name:        "image_password",
					Description: `Initial password to be set for installed OS`,
				},
				resource.Attribute{
					Name:        "ssh_keys",
					Description: `Public SSH keys are set on the image as authorized keys for appropriate SSH login to the instance using the corresponding private key`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The bus type of the volume`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this volume`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The Logical Unit Number of the storage volume`,
				},
				resource.Attribute{
					Name:        "pci_slot",
					Description: `The PCI slot number of the storage volume`,
				},
				resource.Attribute{
					Name:        "backup_unit_id",
					Description: `The uuid of the Backup Unit that user has access to`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The cloud-init configuration for the volume as base64 encoded string`,
				},
				resource.Attribute{
					Name:        "nics",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the attached nic`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the attached nid`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `The MAC address of the NIC`,
				},
				resource.Attribute{
					Name:        "ips",
					Description: `Collection of IP addresses assigned to a nic`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Indicates if the nic will reserve an IP using DHCP`,
				},
				resource.Attribute{
					Name:        "lan",
					Description: `The LAN ID the NIC will sit on`,
				},
				resource.Attribute{
					Name:        "firewall_active",
					Description: `Activate or deactivate the firewall`,
				},
				resource.Attribute{
					Name:        "firewall_type",
					Description: `The type of firewall rules that will be allowed on the NIC`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The Logical Unit Number (LUN) of the storage volume`,
				},
				resource.Attribute{
					Name:        "pci_slot",
					Description: `The PCI slot number of the Nic`,
				},
				resource.Attribute{
					Name:        "firewall_rules",
					Description: `list of`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of the firewall rule`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the firewall rule`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `he protocol for the rule`,
				},
				resource.Attribute{
					Name:        "source_mac",
					Description: `Only traffic originating from the respective MAC address is allowed`,
				},
				resource.Attribute{
					Name:        "source_ip",
					Description: `Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs`,
				},
				resource.Attribute{
					Name:        "target_ip",
					Description: `In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed`,
				},
				resource.Attribute{
					Name:        "icmp_code",
					Description: `Defines the allowed code (from 0 to 254) if protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "icmp_type",
					Description: `Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_start",
					Description: `Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen`,
				},
				resource.Attribute{
					Name:        "port_range_end",
					Description: `Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `The type of firewall rule`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_share",
			Category:         "User Management",
			ShortDescription: `Get Information on share permission objects`,
			Description:      ``,
			Keywords: []string{
				"user",
				"management",
				"share",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "group_id",
					Description: `(Required)The ID of the specific group containing the resource to update.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Required)The ID of the specific resource to update.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Required)The uuid of the share object ` + "`" + `id` + "`" + `, ` + "`" + `resource_id` + "`" + ` and ` + "`" + `group_id` + "`" + ` must be provided. If any of them are missing, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the share resource.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the specific group containing the resource to update.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The ID of the specific resource to update.`,
				},
				resource.Attribute{
					Name:        "edit_privilege",
					Description: `The flag that specifies if the group has permission to edit privileges on this resource.`,
				},
				resource.Attribute{
					Name:        "share_privilege",
					Description: `The group has permission to share this resource.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the share resource.`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `The ID of the specific group containing the resource to update.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `The ID of the specific resource to update.`,
				},
				resource.Attribute{
					Name:        "edit_privilege",
					Description: `The flag that specifies if the group has permission to edit privileges on this resource.`,
				},
				resource.Attribute{
					Name:        "share_privilege",
					Description: `The group has permission to share this resource.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_snapshot",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a IonosCloud Snapshots`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) UUID of an existing snapshot that you want to search for.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing snapshot that you want to search for.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `(Optional) Existing snapshot's location.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Optional) The size of the snapshot to look for. Either ` + "`" + `name` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. Additionally, you can add ` + "`" + `location` + "`" + ` and ` + "`" + `size` + "`" + ` along with the ` + "`" + `name` + "`" + ` argument for a more refined search. If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned. When this happens, please refine your search string so that it is specific enough to return only one result. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the snapshot`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human readable description`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this Snapshot`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `Boolean value representing if the snapshot requires extra protection e.g. two factor protection`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `UUID of the snapshot`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Human readable description`,
				},
				resource.Attribute{
					Name:        "licence_type",
					Description: `OS type of this Snapshot`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `Location of that image/snapshot`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the image in GB`,
				},
				resource.Attribute{
					Name:        "sec_auth_protection",
					Description: `Boolean value representing if the snapshot requires extra protection e.g. two factor protection`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "cpu_hot_unplug",
					Description: `Is capable of CPU hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_unplug",
					Description: `Is capable of memory hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_plug",
					Description: `Is capable of SCSI drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_scsi_hot_unplug",
					Description: `Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_template",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a Ionos Cloud Template`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"template",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A name of that resource.`,
				},
				resource.Attribute{
					Name:        "cores",
					Description: `(Optional) The CPU cores count.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `(Optional) The RAM size in MB.`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `(Optional) The storage size in GB. Any of the arguments ca be provided. If none, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id of template`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of template`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The RAM size in MB`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The storage size in GB`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `Id of template`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of template`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The RAM size in MB`,
				},
				resource.Attribute{
					Name:        "storage_size",
					Description: `The storage size in GB`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_user",
			Category:         "User Management",
			ShortDescription: `Get information on a Ionos Cloud Users`,
			Description:      ``,
			Keywords: []string{
				"user",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "email",
					Description: `(Optional) Email of an existing user that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the user you want to search for. Either ` + "`" + `email` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "administrator",
					Description: `The group has permission to edit privileges on this resource.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address for the user.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name for the user.`,
				},
				resource.Attribute{
					Name:        "force_sec_auth",
					Description: `Indicates if secure (two-factor) authentication should be enabled for the user (true) or not (false).`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password for the user.`,
				},
				resource.Attribute{
					Name:        "sec_auth_active",
					Description: `Indicates if secure authentication is active for the user or not`,
				},
				resource.Attribute{
					Name:        "s3_canonical_user_id",
					Description: `Canonical (S3) id of the user for a given identity`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Indicates if the user is active`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Shows the id and name of the groups that the user is a member of`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the LAN.`,
				},
				resource.Attribute{
					Name:        "administrator",
					Description: `The group has permission to edit privileges on this resource.`,
				},
				resource.Attribute{
					Name:        "email",
					Description: `The e-mail address for the user.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `The first name for the user.`,
				},
				resource.Attribute{
					Name:        "force_sec_auth",
					Description: `Indicates if secure (two-factor) authentication should be enabled for the user (true) or not (false).`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `The last name for the user.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password for the user.`,
				},
				resource.Attribute{
					Name:        "sec_auth_active",
					Description: `Indicates if secure authentication is active for the user or not`,
				},
				resource.Attribute{
					Name:        "s3_canonical_user_id",
					Description: `Canonical (S3) id of the user for a given identity`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Indicates if the user is active`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Shows the id and name of the groups that the user is a member of`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "ionoscloud_volume",
			Category:         "Compute Engine",
			ShortDescription: `Get information on a Ionos Cloud Volume`,
			Description:      ``,
			Keywords: []string{
				"compute",
				"engine",
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of an existing volume that you want to search for.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) ID of the volume you want to search for. Either ` + "`" + `volume` + "`" + ` or ` + "`" + `id` + "`" + ` must be provided. If none, or both are provided, the datasource will return an error. ## Attributes Reference The following attributes are returned by the datasource:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The volume type: HDD or SSD.`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The bus type of the volume: VIRTIO or IDE.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GB.`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `List of paths to files containing a public SSH key that will be injected into IonosCloud provided Linux images.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `The associated public SSH key.`,
				},
				resource.Attribute{
					Name:        "image_password",
					Description: `Required if ` + "`" + `sshkey_path` + "`" + ` is not provided.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image or snapshot UUID.`,
				},
				resource.Attribute{
					Name:        "image_alias",
					Description: `The image alias.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The storage availability zone assigned to the volume: AUTO, ZONE_1, ZONE_2, or ZONE_3. This property is immutable.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on a new volume creation. This option will work only with cloud-init compatible images.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The LUN ID of the storage volume. Null for volumes not mounted to any VM`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.`,
				},
				resource.Attribute{
					Name:        "boot_server",
					Description: `The UUID of the attached server.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The id of the volume.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the volume.`,
				},
				resource.Attribute{
					Name:        "disk_type",
					Description: `The volume type: HDD or SSD.`,
				},
				resource.Attribute{
					Name:        "bus",
					Description: `The bus type of the volume: VIRTIO or IDE.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the volume in GB.`,
				},
				resource.Attribute{
					Name:        "ssh_key_path",
					Description: `List of paths to files containing a public SSH key that will be injected into IonosCloud provided Linux images.`,
				},
				resource.Attribute{
					Name:        "sshkey",
					Description: `The associated public SSH key.`,
				},
				resource.Attribute{
					Name:        "image_password",
					Description: `Required if ` + "`" + `sshkey_path` + "`" + ` is not provided.`,
				},
				resource.Attribute{
					Name:        "image",
					Description: `The image or snapshot UUID.`,
				},
				resource.Attribute{
					Name:        "image_alias",
					Description: `The image alias.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The storage availability zone assigned to the volume: AUTO, ZONE_1, ZONE_2, or ZONE_3. This property is immutable.`,
				},
				resource.Attribute{
					Name:        "user_data",
					Description: `The cloud-init configuration for the volume as base64 encoded string. The property is immutable and is only allowed to be set on a new volume creation. This option will work only with cloud-init compatible images.`,
				},
				resource.Attribute{
					Name:        "device_number",
					Description: `The LUN ID of the storage volume. Null for volumes not mounted to any VM`,
				},
				resource.Attribute{
					Name:        "cpu_hot_plug",
					Description: `Is capable of CPU hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "ram_hot_plug",
					Description: `Is capable of memory hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_plug",
					Description: `Is capable of nic hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "nic_hot_unplug",
					Description: `Is capable of nic hot unplug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_plug",
					Description: `Is capable of Virt-IO drive hot plug (no reboot required)`,
				},
				resource.Attribute{
					Name:        "disc_virtio_hot_unplug",
					Description: `Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.`,
				},
				resource.Attribute{
					Name:        "boot_server",
					Description: `The UUID of the attached server.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"ionoscloud_backup_unit":                        0,
		"ionoscloud_datacenter":                         1,
		"ionoscloud_dbaas_pgsql_backups":                2,
		"ionoscloud_dbaas_pgsql_cluster":                3,
		"ionoscloud_dbaas_pgsql_versions":               4,
		"ionoscloud_firewall":                           5,
		"ionoscloud_group":                              6,
		"ionoscloud_image":                              7,
		"ionoscloud_ipblock":                            8,
		"ionoscloud_ipfailover":                         9,
		"ionoscloud_k8s_cluster":                        10,
		"ionoscloud_k8s_node_pool":                      11,
		"ionoscloud_lan":                                12,
		"ionoscloud_location":                           13,
		"ionoscloud_natgateway":                         14,
		"ionoscloud_natgateway_rule":                    15,
		"ionoscloud_networkloadbalancer":                16,
		"ionoscloud_networkloadbalancer_forwardingrule": 17,
		"ionoscloud_nic":                                18,
		"ionoscloud_private_crossconnect":               19,
		"ionoscloud_resource":                           20,
		"ionoscloud_s3_key":                             21,
		"ionoscloud_server":                             22,
		"ionoscloud_share":                              23,
		"ionoscloud_snapshot":                           24,
		"ionoscloud_template":                           25,
		"ionoscloud_user":                               26,
		"ionoscloud_volume":                             27,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
