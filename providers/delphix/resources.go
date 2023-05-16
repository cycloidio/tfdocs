package delphix

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "delphix_environment",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine_id",
					Description: `(Required) The DCT ID of the Engine on which to create the environment. This ID can be obtained by querying the DCT engines API. A Delphix Engine must be registered with DCT first for it to create an Engine ID.`,
				},
				resource.Attribute{
					Name:        "os_name",
					Description: `(Required) Operating system type of the environment. Valid values are ` + "`" + `[UNIX, WINDOWS]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) Host Name or IP Address of the host that being added to Delphix.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the environment.`,
				},
				resource.Attribute{
					Name:        "is_cluster",
					Description: `(Optional) Whether the environment to be created is a cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_home",
					Description: `(Optional) Absolute path to cluster home drectory. This parameter is (Required) for UNIX cluster environments.`,
				},
				resource.Attribute{
					Name:        "staging_environment",
					Description: `(Optional) Id of the environment where Delphix Connector is installed. This is a (Required) parameter when creating Windows source environments.`,
				},
				resource.Attribute{
					Name:        "connector_port",
					Description: `(Optional) Specify port on which Delphix connector will run. This is a (Required) parameter when creating Windows target environments.`,
				},
				resource.Attribute{
					Name:        "is_target",
					Description: `(Optional) Whether the environment to be created is a target cluster environment. This property is used only when creating Windows cluster environments.`,
				},
				resource.Attribute{
					Name:        "ssh_port",
					Description: `(Optional) ssh port of the environment.`,
				},
				resource.Attribute{
					Name:        "toolkit_path",
					Description: `(Optional) The path where Delphix toolkit can be pushed.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) OS username for Delphix.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) OS user's password.`,
				},
				resource.Attribute{
					Name:        "vault",
					Description: `(Optional) The name or reference of the vault from which to read the host credentials.`,
				},
				resource.Attribute{
					Name:        "hashicorp_vault_engine",
					Description: `(Optional) Vault engine name where the credential is stored.`,
				},
				resource.Attribute{
					Name:        "hashicorp_vault_secret_path",
					Description: `(Optional) Path in the vault engine where the credential is stored.`,
				},
				resource.Attribute{
					Name:        "hashicorp_vault_username_key",
					Description: `(Optional) Key for the username in the key-value store.`,
				},
				resource.Attribute{
					Name:        "hashicorp_vault_secret_key",
					Description: `(Optional) Key for the password in the key-value store.`,
				},
				resource.Attribute{
					Name:        "cyberark_vault_query_string",
					Description: `(Optional) Query to find a credential in the CyberArk vault.`,
				},
				resource.Attribute{
					Name:        "use_kerberos_authentication",
					Description: `(Optional) Whether to use kerberos authentication.`,
				},
				resource.Attribute{
					Name:        "use_engine_public_key",
					Description: `(Optional) Whether to use public key authentication.`,
				},
				resource.Attribute{
					Name:        "nfs_addresses",
					Description: `(Optional) Array of ip address or hostnames. Valid values are a list of addresses. For eg: ` + "`" + `["192.168.10.2"]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ase_db_username",
					Description: `(Optional) Username for the SAP ASE database.`,
				},
				resource.Attribute{
					Name:        "ase_db_password",
					Description: `(Optional) Password for the SAP ASE database.`,
				},
				resource.Attribute{
					Name:        "ase_db_vault",
					Description: `(Optional) The name or reference of the vault from which to read the ASE database credentials.`,
				},
				resource.Attribute{
					Name:        "ase_db_hashicorp_vault_engine",
					Description: `(Optional) Vault engine name where the credential is stored.`,
				},
				resource.Attribute{
					Name:        "ase_db_hashicorp_vault_secret_path",
					Description: `(Optional) Path in the vault engine where the credential is stored.`,
				},
				resource.Attribute{
					Name:        "ase_db_hashicorp_vault_username_key",
					Description: `(Optional) Key for the username in the key-value store.`,
				},
				resource.Attribute{
					Name:        "ase_db_hashicorp_vault_secret_key",
					Description: `(Optional) Key for the password in the key-value store.`,
				},
				resource.Attribute{
					Name:        "ase_db_cyberark_vault_query_string",
					Description: `(Optional) Query to find a credential in the CyberArk vault.`,
				},
				resource.Attribute{
					Name:        "ase_db_use_kerberos_authentication",
					Description: `(Optional) Whether to use kerberos authentication for ASE DB discovery.`,
				},
				resource.Attribute{
					Name:        "java_home",
					Description: `(Optional) The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used.`,
				},
				resource.Attribute{
					Name:        "dsp_keystore_path",
					Description: `(Optional) DSP keystore path.`,
				},
				resource.Attribute{
					Name:        "dsp_keystore_password",
					Description: `(Optional) DSP keystore password.`,
				},
				resource.Attribute{
					Name:        "dsp_keystore_alias",
					Description: `(Optional) DSP keystore alias.`,
				},
				resource.Attribute{
					Name:        "dsp_truststore_path",
					Description: `(Optional) DSP truststore path.`,
				},
				resource.Attribute{
					Name:        "dsp_truststore_password",
					Description: `(Optional) DSP truststore password.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) The environment description.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags to be created for this environment. This is a map of 2 parameters:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Key of the tag`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the tag ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of this environment for replicated and restored objects.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `A reference to the Engine that this Environment connection is associated with.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if this environment is enabled.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `The hosts that are part of this environment.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "namespace",
					Description: `The namespace of this environment for replicated and restored objects.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `A reference to the Engine that this Environment connection is associated with.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `True if this environment is enabled.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `The hosts that are part of this environment.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "delphix_vdb",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_data_id",
					Description: `(Optional) The ID or name of the source object (dSource or VDB) to provision from. All other objects referenced by the parameters must live on the same engine as the source.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `(Optional) The ID or name of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.`,
				},
				resource.Attribute{
					Name:        "target_group_id",
					Description: `(Optional) The ID of the group into which the VDB will be provisioned. If unset, a group is selected randomly on the Engine.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The unique name of the provisioned VDB within a group. If unset, a name is randomly generated.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `(Optional) The name of the database on the target environment. Defaults to name.`,
				},
				resource.Attribute{
					Name:        "cdb_id",
					Description: `(Optional) The ID of the container database (CDB) to provision an Oracle Multitenant database into. When this is not set, a new vCDB will be provisioned.`,
				},
				resource.Attribute{
					Name:        "cluster_node_ids",
					Description: `(Optional) The cluster node ids, name or addresses for this provision operation (Oracle RAC Only).`,
				},
				resource.Attribute{
					Name:        "truncate_log_on_checkpoint",
					Description: `(Optional) Whether to truncate log on checkpoint (ASE only).`,
				},
				resource.Attribute{
					Name:        "os_username",
					Description: `(Optional) The name of the privileged user to run the provision operation (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "os_password",
					Description: `(Optional) The password of the privileged user to run the provision operation (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "db_username",
					Description: `(Optional) [Updatable] The username of the database user (Oracle, ASE Only). Only for update.`,
				},
				resource.Attribute{
					Name:        "db_password",
					Description: `(Optional) [Updatable] The password of the database user (Oracle, ASE Only). Only for update.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `(Optional) The ID or name of the target environment where to provision the VDB. If repository_id unambigously identifies a repository, this is unnecessary and ignored. Otherwise, a compatible repository is randomly selected on the environment.`,
				},
				resource.Attribute{
					Name:        "environment_user_id",
					Description: `(Optional)[Updatable] The environment user ID to use to connect to the target environment.`,
				},
				resource.Attribute{
					Name:        "repository_id",
					Description: `(Optional) The ID of the target repository where to provision the VDB. A repository typically corresponds to a database installation (Oracle home, database instance, ...). Setting this attribute implicitly determines the environment where to provision the VDB.`,
				},
				resource.Attribute{
					Name:        "auto_select_repository",
					Description: `(Optional) Option to automatically select a compatible environment and repository. Mutually exclusive with repository_id.`,
				},
				resource.Attribute{
					Name:        "pre_refresh",
					Description: `(Optional) The commands to execute on the target environment before refreshing the VDB. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "post_refresh",
					Description: `(Optional) The commands to execute on the target environment after refreshing the VDB. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "pre_rollback",
					Description: `(Optional) The commands to execute on the target environment before rewinding the VDB. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "post_rollback",
					Description: `(Optional) The commands to execute on the target environment after rewinding the VDB. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "configure_clone",
					Description: `(Optional) The commands to execute on the target environment when the VDB is created or refreshed. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "pre_snapshot",
					Description: `(Optional) The commands to execute on the target environment before snapshotting a virtual source. These commands can quiesce any data prior to snapshotting. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "post_snapshot",
					Description: `(Optional) The commands to execute on the target environment after snapshotting a virtual source. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "pre_start",
					Description: `(Optional) The commands to execute on the target environment before starting a virtual source. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "post_start",
					Description: `(Optional) The commands to execute on the target environment after starting a virtual source. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "pre_stop",
					Description: `(Optional) The commands to execute on the target environment before stopping a virtual source. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "post_stop",
					Description: `(Optional) The commands to execute on the target environment after stopping a virtual source. This is a map of 5 parameters:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of the hook`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Required)Command to be executed`,
				},
				resource.Attribute{
					Name:        "shell",
					Description: `Type of shell. Valid values are ` + "`" + `[bash, shell, expect, ps, psd]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "element_id",
					Description: `Element ID for the hook`,
				},
				resource.Attribute{
					Name:        "has_credentials",
					Description: `Flag to indicate if it has credentials`,
				},
				resource.Attribute{
					Name:        "vdb_restart",
					Description: `(Optional) [Updatable] Indicates whether the Engine should automatically restart this virtual source when target host reboot is detected.`,
				},
				resource.Attribute{
					Name:        "auxiliary_template_id",
					Description: `(Optional) The ID of the configuration template to apply to the auxiliary container database. This is only relevant when provisioning a Multitenant pluggable database into an existing CDB, i.e when the cdb_id property is set. (Oracle Only)`,
				},
				resource.Attribute{
					Name:        "template_id",
					Description: `(Optional) [Updatable] The ID of the target VDB Template (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "file_mapping_rules",
					Description: `(Optional) Target VDB file mapping rules (Oracle Only). Rules must be line separated (\n or \r) and each line must have the format "pattern:replacement". Lines are applied in order.`,
				},
				resource.Attribute{
					Name:        "oracle_instance_name",
					Description: `(Optional) Target VDB SID name (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "unique_name",
					Description: `(Optional) Target VDB db_unique_name (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "vcdb_name",
					Description: `(Optional) When provisioning an Oracle Multitenant vCDB (when the cdb_id property is not set), the name of the provisioned vCDB (Oracle Multitenant Only).`,
				},
				resource.Attribute{
					Name:        "vcdb_database_name",
					Description: `(Optional) When provisioning an Oracle Multitenant vCDB (when the cdb_id property is not set), the database name of the provisioned vCDB. Defaults to the value of the vcdb_name property. (Oracle Multitenant Only).`,
				},
				resource.Attribute{
					Name:        "mount_point",
					Description: `(Optional) Mount point for the VDB (Oracle, ASE Only).`,
				},
				resource.Attribute{
					Name:        "open_reset_logs",
					Description: `(Optional) Whether to open the database after provision (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "snapshot_policy_id",
					Description: `(Optional) The ID of the snapshot policy for the VDB.`,
				},
				resource.Attribute{
					Name:        "retention_policy_id",
					Description: `(Optional) The ID of the retention policy for the VDB.`,
				},
				resource.Attribute{
					Name:        "recovery_model",
					Description: `(Optional) Recovery model of the source database (MSSql Only). Valid values are ` + "`" + `[ FULL, SIMPLE, BULK_LOGGED ]` + "`" + ``,
				},
				resource.Attribute{
					Name:        "pre_script",
					Description: `(Optional) [Updatable] PowerShell script or executable to run prior to provisioning (MSSql Only).`,
				},
				resource.Attribute{
					Name:        "post_script",
					Description: `(Optional) [Updatable] PowerShell script or executable to run after provisioning (MSSql Only).`,
				},
				resource.Attribute{
					Name:        "cdc_on_provision",
					Description: `(Optional) [Updatable] Option to enable change data capture (CDC) on both the provisioned VDB and subsequent snapshot-related operations (e.g. refresh, rewind) (MSSql Only).`,
				},
				resource.Attribute{
					Name:        "online_log_size",
					Description: `(Optional) Online log size in MB (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "online_log_groups",
					Description: `(Optional) Number of online log groups (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "archive_log",
					Description: `(Optional) Option to create a VDB in archivelog mode (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "new_dbid",
					Description: `(Optional) [Updatable] Option to generate a new DB ID for the created VDB (Oracle Only).`,
				},
				resource.Attribute{
					Name:        "listener_ids",
					Description: `(Optional) [Updatable] The listener IDs for this provision operation (Oracle Only). This is a list of listener ids. For eg: [ "listener-123", "listener-456" ]`,
				},
				resource.Attribute{
					Name:        "custom_env_vars",
					Description: `(Optional) Environment variable to be set when the engine creates a VDB. See the Engine documentation for the list of allowed/denied environment variables and rules about substitution. This is an ordered map of key-value pairs. For eg: { "MY_ENV_VAR1": "$ORACLE_HOME", "MY_ENV_VAR2": "$CRS_HOME/after" }`,
				},
				resource.Attribute{
					Name:        "custom_env_files",
					Description: `(Optional) Environment files to be sourced when the Engine creates a VDB. This path can be followed by parameters. Paths and parameters are separated by spaces. Valid values are a list of env_files. For eg: [ "/export/home/env_file_1", "/export/home/env_file_2" ]`,
				},
				resource.Attribute{
					Name:        "timestamp",
					Description: `(Optional) The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point.`,
				},
				resource.Attribute{
					Name:        "timestamp_in_database_timezone",
					Description: `(Optional) The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) The ID or name of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot.`,
				},
				resource.Attribute{
					Name:        "bookmark_id",
					Description: `(Optional) The ID or name of the bookmark from which to execute the operation. The bookmark must contain only one VDB.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The tags to be created for VDB. This is a map of 2 parameters:`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Key of the tag`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `(Required) Value of the tag ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The VDB object entity ID.`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `The database type of this VDB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The container name of this VDB.`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `The database version of this VDB.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `A reference to the Engine that this VDB belongs to.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `A reference to the Environment that hosts this VDB.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the VDB's host.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the VDB's host.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `A reference to the parent dataset of this VDB.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the group containing this VDB.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of key value pair.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date this VDB was created.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The VDB object entity ID.`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `The database type of this VDB.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The container name of this VDB.`,
				},
				resource.Attribute{
					Name:        "database_version",
					Description: `The database version of this VDB.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `A reference to the Engine that this VDB belongs to.`,
				},
				resource.Attribute{
					Name:        "environment_id",
					Description: `A reference to the Environment that hosts this VDB.`,
				},
				resource.Attribute{
					Name:        "ip_address",
					Description: `The IP address of the VDB's host.`,
				},
				resource.Attribute{
					Name:        "fqdn",
					Description: `The FQDN of the VDB's host.`,
				},
				resource.Attribute{
					Name:        "parent_id",
					Description: `A reference to the parent dataset of this VDB.`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `The name of the group containing this VDB.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `A list of key value pair.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `The date this VDB was created.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "delphix_vdb_group",
			Category:         "Resources",
			ShortDescription: ``,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `A unique identifier for the entity.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `A unique name for the entity.`,
				},
				resource.Attribute{
					Name:        "vdb_ids",
					Description: `The list of VDB IDs in this VDBGroup. ## Attribute Reference This resource exports same attributes as the arguments.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"delphix_environment": 0,
		"delphix_vdb":         1,
		"delphix_vdb_group":   2,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
