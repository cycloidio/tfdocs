package oraclepaas

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_application_container",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages an Appliction Container.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"application",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Application Container.`,
				},
				resource.Attribute{
					Name:        "manifest_file",
					Description: `(Optional) The json manifest file containing the attributes related to launching an application. Use either ` + "`" + `manifest_file` + "`" + ` or ` + "`" + `manifest_attributes` + "`" + ` when specifying launch information.`,
				},
				resource.Attribute{
					Name:        "manifest",
					Description: `(Optional) The manifest attributes related to launching an application. Use either ` + "`" + `manifest_file` + "`" + ` or ` + "`" + `manifest` + "`" + ` when specifying launch information. Manifest attributes is documented below.`,
				},
				resource.Attribute{
					Name:        "deployment_file",
					Description: `(Optional) The json deployment file containing the attributes related to deploying an application. Use either ` + "`" + `deployment_file` + "`" + ` or ` + "`" + `deployment_attributes` + "`" + ` when specifying deployment information.`,
				},
				resource.Attribute{
					Name:        "deployment",
					Description: `(Optional) The deployment attributes related to deploying an application. Use either ` + "`" + `deployment_file` + "`" + ` or ` + "`" + `deployment` + "`" + ` when specifying deployment information. Deployment attributes is documented below.`,
				},
				resource.Attribute{
					Name:        "archive_url",
					Description: `(Optional) Location of the application archive file in Oracle Storage Cloud Service, in the format app-name/file-name.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `(Optional) Uses Oracle Identity Cloud Service to control who can access your Java SE 7 or 8, Node.js, or PHP application. Allowed values are ` + "`" + `basic` + "`" + ` and ` + "`" + `oauth` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) A list of one or more datacenter locations in the OCI region. Required on OCI.`,
				},
				resource.Attribute{
					Name:        "git_repository",
					Description: `(Optional) The URL of the git repository to use the application container.`,
				},
				resource.Attribute{
					Name:        "git_username",
					Description: `(Optional) The username of a user with access to the git respository if the repository is private.`,
				},
				resource.Attribute{
					Name:        "git_password",
					Description: `(Optional) The password for the user with access to the git repository if the repository is private.`,
				},
				resource.Attribute{
					Name:        "load_balancer_subnets",
					Description: `(Optional) Two load balancer subnets. Required on OCI.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Comments about the application deployment.`,
				},
				resource.Attribute{
					Name:        "notification_email",
					Description: `(Optional) Email address to which application deployment status updates are sent.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The name of the region where the application container will be deployed. Required on OCI.`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Optional) The allowed runtime environment variables. The allowed variables are ` + "`" + `java` + "`" + `, ` + "`" + `node` + "`" + `, ` + "`" + `php` + "`" + `, ` + "`" + `python` + "`" + `, ` + "`" + `golang` + "`" + `, ` + "`" + `dotnet` + "`" + `, or ` + "`" + `ruby` + "`" + `. The default is ` + "`" + `java` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscription_type",
					Description: `(Optional) Whether the subscription type is ` + "`" + `hourly` + "`" + ` or ` + "`" + `monthly` + "`" + `. The default is ` + "`" + `hourly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) A map of tags for the application container. Manifest attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "runtime",
					Description: `(Optional) Details the availble runtime attributes. Runtime is documented below.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Determines whether the application is public or private. The default is ` + "`" + `worker` + "`" + ` (private).`,
				},
				resource.Attribute{
					Name:        "command",
					Description: `(Optional) Launch command to execute after the application has been uploaded.`,
				},
				resource.Attribute{
					Name:        "release",
					Description: `(Optional) Details the release attributes of a specific build. Release is documented below.`,
				},
				resource.Attribute{
					Name:        "startup_time",
					Description: `(Optional) The maximum time in seconds to wait for an application to start.`,
				},
				resource.Attribute{
					Name:        "shutdown_time",
					Description: `(Optional) The maximum time in seconds to wait for an application to stop.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Comments about the launch configuration.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) The restart mode for application instances when the application is restarted. The only allowed value is ` + "`" + `rolling` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clustered",
					Description: `(Optional) Boolean for whether the application instances act as a cluster with failover capability.`,
				},
				resource.Attribute{
					Name:        "home",
					Description: `(Optional) The context root of the application.`,
				},
				resource.Attribute{
					Name:        "health_check_endpoint",
					Description: `(Optional) The URL that the application uses for health checks. Deployment attributes supports the following:`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The amount of memory in gigabytes made available to the application. The default is ` + "`" + `2G` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instances",
					Description: `(Optional) The number of application instances. The default is ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notes",
					Description: `(Optional) Comments about the deployment.`,
				},
				resource.Attribute{
					Name:        "environment",
					Description: `(Optional) A map of environment variables used by the application.`,
				},
				resource.Attribute{
					Name:        "secure_environment",
					Description: `(Optional) A list of environment variables marked as secured on the user interface.`,
				},
				resource.Attribute{
					Name:        "java_system_properties",
					Description: `(Optional) A map os java system properties used by the application.`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Service bindings for connections to other Oracle Cloud services. Services is documented below. Runtime supports the following:`,
				},
				resource.Attribute{
					Name:        "major_version",
					Description: `(Required) The major version of the runtime environment. Release supports the following:`,
				},
				resource.Attribute{
					Name:        "build",
					Description: `(Optional) The value for a specific build.`,
				},
				resource.Attribute{
					Name:        "commit",
					Description: `(Optional) The value for a specific commit.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) The value for a specific version. Services supports the following:`,
				},
				resource.Attribute{
					Name:        "identifier",
					Description: `(Required) The value for the identifier`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of service. Allowed values are ` + "`" + `JAAS` + "`" + `, ` + "`" + `DBAAS` + "`" + `, ` + "`" + `MYSQLCS` + "`" + `, ` + "`" + `OEHCS` + "`" + `, ` + "`" + `OEHPCS` + "`" + `, ` + "`" + `DHCS` + "`" + `, ` + "`" + `caching` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the existing service.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) The username to connect to the service.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) The password to connect to the service. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "app_url",
					Description: `URL of the created application`,
				},
				resource.Attribute{
					Name:        "web_url",
					Description: `Web URL of the application`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_database_access_rule",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages a Database Access Rule for an Oracle Database Cloud service instance.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"database",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Access Rule`,
				},
				resource.Attribute{
					Name:        "service_instance_id",
					Description: `(Required) The name of the database service instance to attach the access rule to`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the Access Rule`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) The port or range of ports to allow traffic on`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The IP addresses and subnets from which traffic is allowed. Valid values are ` + "`" + `DB` + "`" + `, ` + "`" + `PUBLIC-INTERNET` + "`" + `, or a single IP address or comma-separated list of subnets (in CIDR format) or IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines whether the access rule is enabled. Default is ` + "`" + `true` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_database_service_instance",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages an Oracle Database Cloud Service instance on the Oracle Cloud Platform.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"database",
				"service",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Service Instance.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `(Required) Database edition for the service instance. Possible values are ` + "`" + `SE` + "`" + `, ` + "`" + `EE` + "`" + `, ` + "`" + `EE_HP` + "`" + `, or ` + "`" + `EE_EP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Required) Service level for the service instance. Possible values are ` + "`" + `BASIC` + "`" + ` or ` + "`" + `PAAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `(Required) Desired compute shape. Possible values are ` + "`" + `oc3` + "`" + `, ` + "`" + `oc4` + "`" + `, ` + "`" + `oc5` + "`" + `, ` + "`" + `oc6` + "`" + `, ` + "`" + `oc1m` + "`" + `, ` + "`" + `oc2m` + "`" + `, ` + "`" + `oc3m` + "`" + `, or ` + "`" + `oc4m` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscription_type",
					Description: `(Required) Billing unit. Possible values are ` + "`" + `HOURLY` + "`" + ` or ` + "`" + `MONTHLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Required) Oracle Database software version; one of: ` + "`" + `12.2.0.1` + "`" + `, ` + "`" + `12.1.0.2` + "`" + `, or ` + "`" + `11.2.0.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vm_public_key",
					Description: `(Required) Public key for the secure shell (SSH). This key will be used for authentication when connecting to the Database Cloud Service instance using an SSH client.`,
				},
				resource.Attribute{
					Name:        "database_configuration",
					Description: `(Required) Specifies the details on how to configure the database. Database configuration is documented below.`,
				},
				resource.Attribute{
					Name:        "default_access_rules",
					Description: `(Optional) Specifies the details on which default access rules are enable or disabled. Default Access Rules are configured below.`,
				},
				resource.Attribute{
					Name:        "desired_state",
					Description: `(Optional) Specifies the desired state of the service instance. Allowed values are ` + "`" + `start` + "`" + `, ` + "`" + `stop` + "`" + `, and ` + "`" + `restart` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "instantiate_from_backup",
					Description: `(Optional) Specify if the service instance's database should, after the instance is created, be replaced by a database stored in an existing cloud backup that was created using Oracle Database Backup Cloud Service. Instantiate from Backup is documented below.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `(Optional) This attribute is only applicable to accounts where regions are supported. The three-part name of an IP network to which the service instance is added. For example: /Compute-identity_domain/user/object`,
				},
				resource.Attribute{
					Name:        "ip_reservations",
					Description: `(Optional) Groups one or more IP reservations in use on this service instance. This attribute is only applicable to accounts where regions are supported.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Provides Cloud Storage information for how to implement service instance backups. Backups is documented below`,
				},
				resource.Attribute{
					Name:        "bring_your_own_license",
					Description: `(Optional) Specify if you want to use an existing perpetual license to Oracle Database to establish the right to use Oracle Database on the new instance. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description of the Service Instance.`,
				},
				resource.Attribute{
					Name:        "high_performance_storage",
					Description: `(Optional) Specifies whether the service instance will be provisioned with high performance storage. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hybrid_disastery_recovery",
					Description: `(Optional) Provides information about an Oracle Hybrid Disaster Recovery configuration. Hybrid Disaster Recovery is documented below.`,
				},
				resource.Attribute{
					Name:        "notification_email",
					Description: `(Optional) The email address to send notifications around successful or unsuccessful completions of the instance-creation operation.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Specifies the location where the service instance is provisioned (only for accounts where regions are supported).`,
				},
				resource.Attribute{
					Name:        "standby",
					Description: `(Optional) Specifies the configuration details of the standby database. This is only applicable in Oracle Cloud Infrastructure Regions. ` + "`" + `failover_database` + "`" + ` and ` + "`" + `disaster_recovery` + "`" + ` inside the ` + "`" + `database_configuration` + "`" + ` block must be set to ` + "`" + `true` + "`" + `. Standby is documented below.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) Name of the subnet within the region where the Oracle Database Cloud Service instance is to be provisioned. Database Configuration supports the following:`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Required) Password for Oracle Database administrator users sys and system. The password must meet the following requirements: Starts with a letter. Is between 8 and 30 characters long. Contains letters, at least one number, and optionally, any number of these special characters: dollar sign ` + "`" + `$` + "`" + `, pound sign ` + "`" + `#` + "`" + `, and underscore ` + "`" + `_` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "usable_storage",
					Description: `(Required) Storage size for data (in GB). Minimum value is ` + "`" + `15` + "`" + `. Maximum value depends on the backup destination: if ` + "`" + `BOTH` + "`" + ` is specified, the maximum value is ` + "`" + `1200` + "`" + `; if ` + "`" + `OSS` + "`" + ` or ` + "`" + `NONE` + "`" + ` is specified, the maximum value is ` + "`" + `2048` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Name of the availability domain within the region where the Oracle Database Cloud Service instance is to be provisioned.`,
				},
				resource.Attribute{
					Name:        "backup_destination",
					Description: `(Optional) Backup Destination. Possible values are ` + "`" + `BOTH` + "`" + `, ` + "`" + `OSS` + "`" + `, ` + "`" + `NONE` + "`" + `.This defaults to ` + "`" + `NONE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_storage_volume_size",
					Description: `(Optional) The size (in GB) for the backup storage volume.`,
				},
				resource.Attribute{
					Name:        "character_set",
					Description: `(Optional) Character Set for the Database Cloud Service Instance. Default value is ` + "`" + `AL32UTF8` + "`" + `. Supported values are: - ` + "`" + `AL32UTF8` + "`" + `, ` + "`" + `AR8ADOS710` + "`" + `, ` + "`" + `AR8ADOS720` + "`" + `, ` + "`" + `AR8APTEC715` + "`" + `, ` + "`" + `AR8ARABICMACS` + "`" + `, ` + "`" + `AR8ASMO8X` + "`" + `, ` + "`" + `AR8ISO8859P6` + "`" + `, ` + "`" + `AR8MSWIN1256` + "`" + `, ` + "`" + `AR8MUSSAD768` + "`" + `, ` + "`" + `AR8NAFITHA711` + "`" + `, ` + "`" + `AR8NAFITHA721` + "`" + `, ` + "`" + `AR8SAKHR706` + "`" + `, ` + "`" + `AR8SAKHR707` + "`" + `, ` + "`" + `AZ8ISO8859P9E` + "`" + `, ` + "`" + `BG8MSWIN` + "`" + `, ` + "`" + `BG8PC437S` + "`" + `, ` + "`" + `BLT8CP921` + "`" + `, ` + "`" + `BLT8ISO8859P13` + "`" + `, ` + "`" + `BLT8MSWIN1257` + "`" + `, ` + "`" + `BLT8PC775` + "`" + `, ` + "`" + `BN8BSCII` + "`" + `, ` + "`" + `CDN8PC863` + "`" + `, ` + "`" + `CEL8ISO8859P14` + "`" + `, ` + "`" + `CL8ISO8859P5` + "`" + `, ` + "`" + `CL8ISOIR111` + "`" + `, ` + "`" + `CL8KOI8R` + "`" + `, ` + "`" + `CL8KOI8U` + "`" + `, ` + "`" + `CL8MACCYRILLICS` + "`" + `, ` + "`" + `CL8MSWIN1251` + "`" + `, ` + "`" + `EE8ISO8859P2` + "`" + `, ` + "`" + `EE8MACCES` + "`" + `, ` + "`" + `EE8MACCROATIANS` + "`" + `, ` + "`" + `EE8MSWIN1250` + "`" + `, ` + "`" + `EE8PC852` + "`" + `, ` + "`" + `EL8DEC` + "`" + `, ` + "`" + `EL8ISO8859P7` + "`" + `, ` + "`" + `EL8MACGREEKS` + "`" + `, ` + "`" + `EL8MSWIN1253` + "`" + `, ` + "`" + `EL8PC437S` + "`" + `, ` + "`" + `EL8PC851` + "`" + `, ` + "`" + `EL8PC869` + "`" + `, ` + "`" + `ET8MSWIN923` + "`" + `, ` + "`" + `HU8ABMOD` + "`" + `, ` + "`" + `HU8CWI2` + "`" + `, ` + "`" + `IN8ISCII` + "`" + `, ` + "`" + `IS8PC861` + "`" + `, ` + "`" + `IW8ISO8859P8` + "`" + `, ` + "`" + `IW8MACHEBREWS` + "`" + `, ` + "`" + `IW8MSWIN1255` + "`" + `, ` + "`" + `IW8PC1507` + "`" + `, ` + "`" + `JA16EUC` + "`" + `, ` + "`" + `JA16EUCTILDE` + "`" + `, ` + "`" + `JA16SJIS` + "`" + `, ` + "`" + `JA16SJISTILDE` + "`" + `, ` + "`" + `JA16VMS` + "`" + `, ` + "`" + `KO16KSC5601` + "`" + `, ` + "`" + `KO16KSCCS` + "`" + `, ` + "`" + `KO16MSWIN949` + "`" + `, ` + "`" + `LA8ISO6937` + "`" + `, ` + "`" + `LA8PASSPORT` + "`" + `, ` + "`" + `LT8MSWIN921` + "`" + `, ` + "`" + `LT8PC772` + "`" + `, ` + "`" + `LT8PC774` + "`" + `, ` + "`" + `LV8PC1117` + "`" + `, ` + "`" + `LV8PC8LR` + "`" + `, ` + "`" + `LV8RST104090` + "`" + `, ` + "`" + `N8PC865` + "`" + `, ` + "`" + `NE8ISO8859P10` + "`" + `, ` + "`" + `NEE8ISO8859P4` + "`" + `, ` + "`" + `RU8BESTA` + "`" + `, ` + "`" + `RU8PC855` + "`" + `, ` + "`" + `RU8PC866` + "`" + `, ` + "`" + `SE8ISO8859P3` + "`" + `, ` + "`" + `TH8MACTHAIS` + "`" + `, ` + "`" + `TH8TISASCII` + "`" + `, ` + "`" + `TR8DEC` + "`" + `, ` + "`" + `TR8MACTURKISHS` + "`" + `, ` + "`" + `TR8MSWIN1254` + "`" + `, ` + "`" + `TR8PC857` + "`" + `, ` + "`" + `US7ASCII` + "`" + `, ` + "`" + `US8PC437` + "`" + `, ` + "`" + `UTF8` + "`" + `, ` + "`" + `VN8MSWIN1258` + "`" + `, ` + "`" + `VN8VN3` + "`" + `, ` + "`" + `WE8DEC` + "`" + `, ` + "`" + `WE8DG` + "`" + `, ` + "`" + `WE8ISO8859P1` + "`" + `, ` + "`" + `WE8ISO8859P15` + "`" + `, ` + "`" + `WE8ISO8859P9` + "`" + `, ` + "`" + `WE8MACROMAN8S` + "`" + `, ` + "`" + `WE8MSWIN1252` + "`" + `, ` + "`" + `WE8NCR4970` + "`" + `, ` + "`" + `WE8NEXTSTEP` + "`" + `, ` + "`" + `WE8PC850` + "`" + `, ` + "`" + `WE8PC858` + "`" + `, ` + "`" + `WE8PC860` + "`" + `, ` + "`" + `WE8ROMAN8` + "`" + `, ` + "`" + `ZHS16CGB231280` + "`" + `, ` + "`" + `ZHS16GBK` + "`" + `, ` + "`" + `ZHT16BIG5` + "`" + `, ` + "`" + `ZHT16CCDC` + "`" + `, ` + "`" + `ZHT16DBT` + "`" + `, ` + "`" + `ZHT16HKSCS` + "`" + `, ` + "`" + `ZHT16MSWIN950` + "`" + `, ` + "`" + `ZHT32EUC` + "`" + `, ` + "`" + `ZHT32SOPS` + "`" + `, ` + "`" + `ZHT32TRIS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_storage_volume_size",
					Description: `(Optional) The size (in GB) for the data storage volume.`,
				},
				resource.Attribute{
					Name:        "disaster_recovery",
					Description: `(Optional) Specify if an Oracle Data Guard configuration is created using the Disaster Recovery option or the High Availability option. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "failover_database",
					Description: `(Optional) Specify if an Oracle Data Guard configuration comprising a primary database and a standby database is created. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "golden_gate",
					Description: `(Optional) Specify if the database should be configured for use as the replication database of an Oracle GoldenGate Cloud Service instance. You cannot set ` + "`" + `goldenGate` + "`" + ` to ` + "`" + `true` + "`" + ` if either ` + "`" + `is_rac` + "`" + ` or ` + "`" + `failoverDatabase` + "`" + ` is set to ` + "`" + `true` + "`" + `. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "is_rac",
					Description: `(Optional) Specify if a cluster database using Oracle Real Application Clusters should be configured. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "national_character_set",
					Description: `(Optional) National Character Set for the Database Cloud Service instance. Valid values are ` + "`" + `AL16UTF16` + "`" + ` and ` + "`" + `UTF8` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pdb_name",
					Description: `(Optional) This attribute is valid when Database Cloud Service instance is configured with version 12c. Pluggable Database Name for the Database Cloud Service instance. Default value is ` + "`" + `pdb1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sid",
					Description: `(Optional) Database Name for the Database Cloud Service instance. Default value is ` + "`" + `ORCL` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "source_service_name",
					Description: `(Optional) Indicates that the service instance should be created as a "snapshot clone" of another service instance. Provide the name of the existing service instance whose snapshot is to be used.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) The name of the snapshot of the service instance specified by sourceServiceName that is to be used to create a "snapshot clone". This parameter is valid only if source_service_name is specified.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `(Optional) Time Zone for the Database Cloud Service instance. Default value is ` + "`" + `UTC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Component type to which the set of parameters applies. Defaults to ` + "`" + `db` + "`" + ``,
				},
				resource.Attribute{
					Name:        "db_demo",
					Description: `(Optional) Indicates whether to include the Demos PDB. Default Access Rules supports the following:`,
				},
				resource.Attribute{
					Name:        "enable_ssh",
					Description: `(Optional) Indicates whether to enable the ssh access rule.`,
				},
				resource.Attribute{
					Name:        "enable_http",
					Description: `(Optional) Indicates whether to enable the http access rule. This is only configurable with a single instance.`,
				},
				resource.Attribute{
					Name:        "enable_https",
					Description: `(Optional) Indiciates whether to enable the http with ssl access rule. This is only configurable with a single instance.`,
				},
				resource.Attribute{
					Name:        "enable_db_console",
					Description: `(Optional) Indicates whether to enable the db console access rule. This is only configurable with a single instance.`,
				},
				resource.Attribute{
					Name:        "enable_db_express",
					Description: `(Optional) Indicates whether to enable the db express access rule. This is only configurable with a single instance.`,
				},
				resource.Attribute{
					Name:        "enable_db_listener",
					Description: `(Optional) Indicates whether to enable the db listener access rule. This is only configurable with a single instance`,
				},
				resource.Attribute{
					Name:        "enable_em_console",
					Description: `(Optional) Indicates whether to enable the em console access rule. This is only configurable with a RAC instance.`,
				},
				resource.Attribute{
					Name:        "enable_rac_db_listener",
					Description: `(Optional) Indicates whether to enable the rac db listene access rule. This is only configurable with a RAC instance`,
				},
				resource.Attribute{
					Name:        "enable_scan_listener",
					Description: `(Optional) Indicates whether to enable the scan listener access rule. This is only configurable with a RAC instance`,
				},
				resource.Attribute{
					Name:        "enable_rac_ons",
					Description: `(Optional) Indicates whether to enable the rac ons access rule. This is only configurable with a RAC instance. Standby supports the following:`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Required) Name of the availability domain within the region where the standby database of the Oracle Database Cloud Service instance is to be provisioned.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Required) Name of the subnet within the region where the standby database of the Oracle Database Cloud Service instance is to be provisioned. Instantiate from Backup supports the following:`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `(Required) Name of the Oracle Storage Cloud Service container where the existing cloud backup is stored.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_username",
					Description: `(Required) Username of the Oracle Cloud user.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_password",
					Description: `(Required) Password of the Oracle Cloud user specified in ` + "`" + `ibkup_cloud_storage_user` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "database_id",
					Description: `(Required) Database id of the database from which the existing cloud backup was created.`,
				},
				resource.Attribute{
					Name:        "decryption_key",
					Description: `(Optional) Password used to create the existing, password-encrypted cloud backup. This password is used to decrypt the backup. Specify either ` + "`" + `ibkup_decryption_key` + "`" + ` or ` + "`" + `ibkup_wallet_file_content` + "`" + ` for decrypting the backup.`,
				},
				resource.Attribute{
					Name:        "on_premise",
					Description: `(Optional) Specify if the existing cloud backup being used to replace the database is from an on-premises database or another Database Cloud Service instance. The default value is false.`,
				},
				resource.Attribute{
					Name:        "service_id",
					Description: `(Optional) Oracle Database Cloud Service instance name from which the database of new Oracle Database Cloud Service instance should be created. This value is required if ` + "`" + `on_premise` + "`" + ` is set to true.`,
				},
				resource.Attribute{
					Name:        "wallet_file_content",
					Description: `(Optional) String containing the xsd:base64Binary representation of the cloud backup's wallet file. This wallet is used to decrypt the backup. Specify either ` + "`" + `ibkup_decryption_key` + "`" + ` or ` + "`" + `ibkup_wallet_file_content` + "`" + ` for decrypting the backup. Backups support the following:`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `(Required) Name of the Oracle Storage Cloud Service container used to provide storage for your service instance backups. Use the following format to specify the container name: ` + "`" + `<storageservicename>-<storageidentitydomain>/<containername>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cloud_storage_username",
					Description: `(Required) Username for the Oracle Storage Cloud Service administrator.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_password",
					Description: `(Required) Password for the Oracle Storage Cloud Service administrator.`,
				},
				resource.Attribute{
					Name:        "create_if_missing",
					Description: `(Optional) Specify if the given cloud_storage_container is to be created if it does not already exist. Default value is ` + "`" + `false` + "`" + `. Hybrid Disaster Recovery supports the following:`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `(Required) Name of the Oracle Storage Cloud Service container where the backup from on-premise instance is stored. Use the following format to specify the container name: ` + "`" + `<storageservicename>-<storageidentitydomain>/<containername>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cloud_storage_username",
					Description: `(Required) Username for the Oracle Storage Cloud Service administrator.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_password",
					Description: `(Required) Password for the Oracle Storage Cloud Service administrator. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "compute_site_name",
					Description: `The Oracle Cloud location housing the service instance.`,
				},
				resource.Attribute{
					Name:        "em_url",
					Description: `The URL to use to connect to Enterprise Manager on the service instance.`,
				},
				resource.Attribute{
					Name:        "glassfish_url",
					Description: `The URL to use to connect to the Oracle GlassFish Server Administration Console on the service instance.`,
				},
				resource.Attribute{
					Name:        "identity_domain",
					Description: `The identity domain housing the service instance.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the service instance.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Service Instance`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_java_access_rule",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages an Access Rule for an Java Cloud service instance.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"java",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Access Rule`,
				},
				resource.Attribute{
					Name:        "service_instance_id",
					Description: `(Required) The name of the java service instance to attach the access rule to`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) The description of the Access Rule`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) The port or range of ports to allow traffic on`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) Destination to which traffic is allowed. Valid values include ` + "`" + `WLS_ADMIN` + "`" + `, ` + "`" + `WLS_ADMIN_SERVER` + "`" + `, ` + "`" + `OTD_ADMIN_HOST` + "`" + `, ` + "`" + `OTD` + "`" + ``,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The IP addresses and subnets from which traffic is allowed. Valid values include ` + "`" + `WLS_ADMIN` + "`" + `, ` + "`" + `WLS_ADMIN_SERVER` + "`" + `, ` + "`" + `WLS_MANAGED_SERVER` + "`" + `, ` + "`" + `OTD_ADMIN_HOST` + "`" + `, ` + "`" + `OTD` + "`" + `, or a single IP address or comma-separated list of subnets (in CIDR format) or IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines whether the access rule is enabled. Default is ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Specifies the communication protocol. Valid values are ` + "`" + `tcp` + "`" + ` or ` + "`" + `udp` + "`" + `. Default is ` + "`" + `tcp` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_java_service_instance",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages an Oracle Java Cloud Service instance on Oracle Cloud Infrastructure and Oracle Cloud Infrastructure Classic.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"java",
				"service",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the Service Instance.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Required) The ssh key to connect to the java service instance.`,
				},
				resource.Attribute{
					Name:        "edition",
					Description: `(Required) The edition for the service instance. Possible values are ` + "`" + `SE` + "`" + `, ` + "`" + `EE` + "`" + `, or ` + "`" + `SUITE` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "service_version",
					Description: `(Required) Oracle WebLogic Server software version. Valid values are: ` + "`" + `12cRelease213` + "`" + `, ` + "`" + `12cRelease212` + "`" + `, ` + "`" + `12cR3` + "`" + `, or ` + "`" + `11gR1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Required) Provides Cloud Storage information for service instance backups. Backups is documented below`,
				},
				resource.Attribute{
					Name:        "metering_frequency",
					Description: `(Optional) Billing unit. Possible values are ` + "`" + `HOURLY` + "`" + ` or ` + "`" + `MONTHLY` + "`" + `. Default value is ` + "`" + `HOURLY` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Name of a data center location in the Oracle Cloud Infrastructure region that is specified in region. This is only available for OCI.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) Name of the snapshot to clone from.`,
				},
				resource.Attribute{
					Name:        "source_service_name",
					Description: `(Optional) Name of the existing Oracle Java Cloud Service instance that has the snapshot from which you are creating a clone.`,
				},
				resource.Attribute{
					Name:        "subnet",
					Description: `(Optional) A subdivision of a cloud network that is set up in the data center as specified in ` + "`" + `availability_domain` + "`" + `. This is only available for OCI.`,
				},
				resource.Attribute{
					Name:        "use_identity_service",
					Description: `(Optional) Flag that specifies whether to use Oracle Identity Cloud Service (true) or the local WebLogic identity store (false) for user authentication and to maintain administrators, application users, groups and roles. The default value is false.`,
				},
				resource.Attribute{
					Name:        "weblogic_server",
					Description: `(Required) The attributes required to create a WebLogic server alongside the java service instance. WebLogic Server is documented below.`,
				},
				resource.Attribute{
					Name:        "oracle_traffic_director",
					Description: `(Optional) The attributes required to create an Oracle Traffic Director (Load balancer). OTD is documented below.`,
				},
				resource.Attribute{
					Name:        "level",
					Description: `(Optional) Service level for the service instance. Possible values are ` + "`" + `BASIC` + "`" + ` or ` + "`" + `PAAS` + "`" + `. Default value is ` + "`" + `PAAS` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_destination",
					Description: `(Optional) Specifies whether to enable backups for this Oracle Java Cloud Service instance. Valid values are ` + "`" + `BOTH` + "`" + ` or ` + "`" + `NONE` + "`" + `. Defaults to ` + "`" + `BOTH` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "desired_state",
					Description: `(Optional) Specifies the desired state of the service instance. Allowed values are ` + "`" + `running` + "`" + ` or ` + "`" + `shutdown` + "`" + `. The default is ` + "`" + `running` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Provides additional on the java service instance.`,
				},
				resource.Attribute{
					Name:        "enable_admin_console",
					Description: `(Optional) Flag that specifies whether to enable (true) or disable (false) the access rules that control external communication to the WebLogic Server Administration Console, Fusion Middleware Control, and Load Balancer Console.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `(Optional) The three-part name of a custom IP network to attach this service instance to. For example: ` + "`" + `/Compute-identity_domain/user/object` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "assign_public_ip",
					Description: `(Optional) Flag that specifies whether to assign (true) or not assign (false) public IP addresses to the nodes in your service instance. The default is ` + "`" + `true` + "`" + `, which means any node added during service instance provisioning, or later added as part of a scaling operation, will have a public IP address assigned to it. This attribute is only applicable when provisioning an Oracle Java Cloud Service instance in a region on Oracle Cloud Infrastructure Classic, and a custom IP network is specified in ` + "`" + `ip_network` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) Name of the region where the Oracle Java Cloud Service instance is to be provisioned. This attribute is only applicable to accounts where regions are supported. A region name must be specified if you want to use ipReservations or ipNetwork.`,
				},
				resource.Attribute{
					Name:        "bring_your_own_license",
					Description: `(Optional) Flag that specifies whether to apply an existing on-premises license for Oracle WebLogic Server (true) to the new Oracle Java Cloud Service instance you are provisioning. The default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_delete",
					Description: `(Optional) Flag that specifies whether you want to force the removal of the service instance even if the database instance cannot be reached to delete the database schemas. Default value is ` + "`" + `true` + "`" + ` Backups supports the following:`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `(Required) Name of the Oracle Storage Cloud Service container used to provide storage for your service instance backups. Use the following format to specify the container name on OCI Object Storage ` + "`" + `https://swiftobjectstorage.<region>.oraclecloud.com/v1/<tenancy>/<bucketname>` + "`" + ` or for Oracle Cloud Infrastructure Classic user the format: ` + "`" + `<storageservicename>-<storageidentitydomain>/<containername>` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cloud_storage_username",
					Description: `(Optional) Username for the Oracle Storage Cloud Service administrator. If left unspecified, the username for Oracle Public Cloud is used.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_password",
					Description: `(Optional) Password for the Oracle Storage Cloud Service administrator. If left unspecified, the password for Oracle Public Cloud is used.`,
				},
				resource.Attribute{
					Name:        "create_if_missing",
					Description: `(Optional) Specify if the given cloud_storage_container is to be created if it does not already exist. Default value is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_oauth_for_storage",
					Description: `(Optional) Flag that specifies whether to use the default OAuth protected object storage for instance backups. WebLogic Server supports the following:`,
				},
				resource.Attribute{
					Name:        "database",
					Description: `(Required) Information about the database deployment on Oracle Database Cloud Service or Oracle Cloud Infrastructure. Database is documented below.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `(Required) Desired compute shape.`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Required) Admin information for the WebLogic Server. Admin is documented below.`,
				},
				resource.Attribute{
					Name:        "application_database",
					Description: `(Optional) Details of Database Cloud Service database deployments that host application schemas. Multiple can be specified. Application Database is specified below.`,
				},
				resource.Attribute{
					Name:        "backup_volume_size",
					Description: `(Optional) Size of the backup volume for the service. The value must be a multiple of GBs. You can specify this value in bytes or GBs. If specified in GBs, use the following format: nG, where n specifies the number of GBs. For example, you can express 10 GBs as bytes or GBs. For example: 100000000000 or 10G. This value defaults to the system configured volume size.`,
				},
				resource.Attribute{
					Name:        "cluster_name",
					Description: `(Optional) - Specifies the name of the cluster that contains the Managed Servers for the service instance.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Optional) Details the properties about one or more clusters. Cluster is documented below.`,
				},
				resource.Attribute{
					Name:        "connect_string",
					Description: `(Optional) - Connection string for the database. The connection string must be entered using one of the following formats:`,
				},
				resource.Attribute{
					Name:        "content_port",
					Description: `(Optional) - Port for accessing the deployed applications using HTTP. Default value is 8001.`,
				},
				resource.Attribute{
					Name:        "deployment_channel_port",
					Description: `Port for accessing the Administration Server using WLST. Default value is 9001.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Optional) Information about the WebLogic domain. Domain is documented below.`,
				},
				resource.Attribute{
					Name:        "ip_reservations",
					Description: `(Optional) A list of ip reservation names.`,
				},
				resource.Attribute{
					Name:        "load_balancer",
					Description: `(Optional) Information about the loadbalancer to attach to the java service instance. Load Balancer is specfied below.`,
				},
				resource.Attribute{
					Name:        "managed_servers",
					Description: `(Optional) Details information about the managed servers the java service instance will look after. Managed Servers is documented below.`,
				},
				resource.Attribute{
					Name:        "middleware_volume_size",
					Description: `(Optional) Size of the middleware home disk volume for the service (/u01/app/oracle/middleware). The value must be a multiple of GBs. You can specify this value in bytes or GBs. If specified in GBs, use the following format: nG, where n specifies the number of GBs. For example, you can express 10 GBs as bytes or GBs. For example: 100000000000 or 10G. This value defaults to the system configured volume size.`,
				},
				resource.Attribute{
					Name:        "node_manager",
					Description: `(Optional) Node Manager is a WebLogic Server utility that enables you to start, shut down, and restart Administration Server and Managed Server instances from a remote location. Node Manager is documented below.`,
				},
				resource.Attribute{
					Name:        "pdb_service_name",
					Description: `(Optional) Name of the pluggable database for Oracle Database 12c. The default value is the pluggable database name when the database was created.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Optional) A block of port specifications. Weblogic Server Ports are specified below.`,
				},
				resource.Attribute{
					Name:        "upper_stack_product_name",
					Description: `(Optional) The Oracle Fusion Middleware product installer to add to this Oracle Java Cloud Service instance. Valid values are ` + "`" + `ODI` + "`" + ` (Oracle Data Integrator) or ` + "`" + `WCP` + "`" + ` (Oracle WebCenter Portal)`,
				},
				resource.Attribute{
					Name:        "root_url",
					Description: `(Computed) The URL of the WebLogic Server Administration console. OTD supports the following:`,
				},
				resource.Attribute{
					Name:        "admin",
					Description: `(Required) Admin information for the Oracle Traffic Director. Admin is documented below.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `(Required) Desired compute shape.`,
				},
				resource.Attribute{
					Name:        "high_availability",
					Description: `(Optional) Flag that specifies whether load balancer HA is enabled. This value defaults to false (that is, HA is not enabled).`,
				},
				resource.Attribute{
					Name:        "ip_reservations",
					Description: `(Optional) A list of ip reservation names.`,
				},
				resource.Attribute{
					Name:        "listener",
					Description: `(Optional) Specifies the type and number of the listener port. Listener is documented below.`,
				},
				resource.Attribute{
					Name:        "load_balancing_policy",
					Description: `(Optional) Policy to use for routing requests to the load balancer. Valid policies include: ` + "`" + `least_connection_count` + "`" + `, ` + "`" + `least_response_time` + "`" + `, ` + "`" + `round_robin` + "`" + `. The default value is ` + "`" + `least_connection_count` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "root_url",
					Description: `(Computed) The URL of the OTD console. Database supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for the database administrator.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the database administrator.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the database on the Database Cloud Service. For use with Oracle Database Cloud Service only, Not required when deploying with Oracle Cloud Infrastructure native VM, Bare Metal or Exadata Cloud Service instance, see ` + "`" + `connect_string` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Computed) The hostname for the database. Admin supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for the WebLogic Server or Oracle Traffic Director administrator.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the WebLogic Server or Oracle Traffic Director administrator.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port for accessing the WebLogic Server or Oracle Traffic Director using HTTP. The default values are 7001 for WebLogic Server or 8989 for Oracle Traffic Director.`,
				},
				resource.Attribute{
					Name:        "secured_port",
					Description: `(Optional) Secured Port for accessing the WebLogic Server. The default value is 7002.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `(Computed) The hostname for the admin server on the WebLogic Server or OTD. Application Database supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username for the database administrator.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password for the database administrator.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the database deployment on the Database Cloud Service.`,
				},
				resource.Attribute{
					Name:        "pdb_name",
					Description: `(Optional) Name of the pluggable database for Oracle Database 12c. If not specified, the pluggable database name configured when the database was created will be used. Cluster supports the following:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the cluster to create.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) Type of cluster to create. Valid values are ` + "`" + `APPLICATION_CLUSTER` + "`" + ` or ` + "`" + `CACHING_CLUSTER` + "`" + ``,
				},
				resource.Attribute{
					Name:        "server_count",
					Description: `(Optional) Number of servers to create in this cluster. The default value is 1.`,
				},
				resource.Attribute{
					Name:        "servers_per_node",
					Description: `(Optional) Number of JVMs to start on each VM (node). The default value is 1.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `(Optional) Desired compute shape for the nodes in this cluster.`,
				},
				resource.Attribute{
					Name:        "path_prefixes",
					Description: `(Optional) A single path prefix or multiple path prefixes separated by commas. Domain supports the following:`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `(Optional) Mode of the domain. Valid values are ` + "`" + `DEVELOPMENT` + "`" + ` or ` + "`" + `PRODUCTION` + "`" + `. Default value is ` + "`" + `PRODUCTION` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Name of the WebLogic domain. By default, the domain name will be generated from the first eight characters of the Oracle Java Cloud Service instance name (serviceName), using the following format: first8charsOfServiceInstanceName_domain.`,
				},
				resource.Attribute{
					Name:        "partition_count",
					Description: `(Optional) Number of partitions to enable in the domain for WebLogic Server 12.2.1. Valid values include: 0 (no partitions), 1, 2, and 4.`,
				},
				resource.Attribute{
					Name:        "volume_size",
					Description: `(Optional) Size of the domain volume for the service. The value must be a multiple of GBs. You can specify this value in bytes or GBs. If specified in GBs, use the following format: nG, where n specifies the number of GBs. For example, you can express 10 GBs as bytes or GBs. For example: 100000000000 or 10G. Listener supports the following:`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Listener port for the load balancer for accessing deployed applications using HTTP. If left unspecified, applications on this service instance cannot be reached via http.`,
				},
				resource.Attribute{
					Name:        "secured_port",
					Description: `(Optional) Secured listener port for the load balancer for accessing deployed applications using HTTPS.`,
				},
				resource.Attribute{
					Name:        "privileged_port",
					Description: `(Optional) Privileged listener port for accessing the deployed applications using HTTP.`,
				},
				resource.Attribute{
					Name:        "privileged_secured_port",
					Description: `(Optional) Privileged listener port for accessing the deployed applications using HTTPS. Managed Server supports the following:`,
				},
				resource.Attribute{
					Name:        "server_count",
					Description: `(Optional) Number of Managed Servers in the domain. Valid values include: 1, 2, 4, and 8. The default value is 1.`,
				},
				resource.Attribute{
					Name:        "initial_heap_size",
					Description: `(Optional) Initial Java heap size for a Managed Server JVM, specified in megabytes.`,
				},
				resource.Attribute{
					Name:        "max_heap_size",
					Description: `(Optional) Maximum Java heap size for a Managed Server JVM, specified in megabytes.`,
				},
				resource.Attribute{
					Name:        "jvm_args",
					Description: `(Optional) One or more Managed Server JVM arguments separated by a space.`,
				},
				resource.Attribute{
					Name:        "initial_permanent_generation",
					Description: `(Optional) Initial Permanent Generation space in Java heap memory.`,
				},
				resource.Attribute{
					Name:        "max_permanent_generation",
					Description: `(Optional) Maximum Permanent Generation space in Java heap memory.`,
				},
				resource.Attribute{
					Name:        "overwrite_jvm_args",
					Description: `(Optional) Flag that determines whether the user defined Managed Server JVM arguments specified in msJvmArgs should replace the server start arguments (true), or append the server start arguments (false). Default is false. Node Manager supports the following:`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) User name for the Node Manager. This value defaults to the WebLogic administrator user name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password for the Node Manager. This value defaults to the WebLogic administrator password.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Optional) Port for the Node Manager. This value defaults to 5556. WebLogic Server Ports support the following:`,
				},
				resource.Attribute{
					Name:        "privileged_content_port",
					Description: `(Optional) Privileged content port for accessing the deployed applications using HTTP. To disable the privileged content port, set the value to 0. The default value is 80.`,
				},
				resource.Attribute{
					Name:        "priviliged_secured_content_port",
					Description: `(Optional) Privileged content port for accessing the deployed applications using HTTPS. To disable the privileged secured content port, set the value to 0. The default value is 443.`,
				},
				resource.Attribute{
					Name:        "deployment_channel_port",
					Description: `(Optional) Port for accessing the WebLogic Administration Server using WLST. The default value is 9001.`,
				},
				resource.Attribute{
					Name:        "content_port",
					Description: `(Optional) Port for accessing the deployed applications using HTTP. The default value is 8001. Load Balancer supports the following:`,
				},
				resource.Attribute{
					Name:        "load_balancing_policy",
					Description: `(Optional) Policy to use for routing requests to the origin servers of the Oracle managed load balancer. Valid values are ` + "`" + `LEAST_CONN` + "`" + `, ` + "`" + `IP_HASH` + "`" + `, or ` + "`" + `ROUND_ROBIN` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subnets",
					Description: `(Optional) A list of two OCI-managed load balancer ids. In addition to the above, the following values are exported:`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the service instance.`,
				},
				resource.Attribute{
					Name:        "uri",
					Description: `The Uniform Resource Identifier for the Service Instance For ` + "`" + `load_balancer` + "`" + ` the ` + "`" + `admin_url` + "`" + `, ` + "`" + `console_url` + "`" + `, and ` + "`" + `url` + "`" + ` are exported. (eg. ` + "`" + `load_balancer.0.url` + "`" + `)`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_mysql_access_rule",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages a MySQL Access Rule for an Oracle MySQL Cloud service instance.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"mysql",
				"access",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "service_instance_id",
					Description: `(Required) The name of MySQL instance to attach the access rule to.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Name of the rule.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the rule.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `(Optional) Communication protocol for the rule. For example, tcp.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `(Required) Ports for the rule. This can be a single port or a port range.`,
				},
				resource.Attribute{
					Name:        "source",
					Description: `(Required) The hosts from which traffic is allowed. For example, PUBLIC-INTERNET for any host on the Internet, a single IP address or a comma-separated list of subnets (in CIDR format) or IPv4 addresses.`,
				},
				resource.Attribute{
					Name:        "destination",
					Description: `(Required) The service component to allow traffic to. For example, mysql_MASTER.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Determines whether the access rule is enabled. Valid values are ` + "`" + `true` + "`" + ` and ` + "`" + `false` + "`" + `. The Default is ` + "`" + `true` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "oraclepaas_mysql_service_instance",
			Category:         "PaaS Resources",
			ShortDescription: `Creates and manages an Oracle MySQL Cloud Service instance on the Oracle Cloud Platform.`,
			Description:      ``,
			Keywords: []string{
				"paas",
				"mysql",
				"service",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required). The name of MySQL Cloud Service instance.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional). A description of the MySQL Instance`,
				},
				resource.Attribute{
					Name:        "ssh_public_key",
					Description: `(Required). The public key for the secure shell (SSH). This key wil be used for authentication when the user logs on to the instance over SSH.`,
				},
				resource.Attribute{
					Name:        "backup_destination",
					Description: `(Required) The destination where the database backups will be stored.`,
				},
				resource.Attribute{
					Name:        "shape",
					Description: `(Required) The desired compute shape. A shape defines the number of Oracle Compute Units (OCPUs) and amount of memory (RAM). See [About Shapes](http://www.oracle.com/pls/topic/lookup?ctx=cloud&id=OCSUG210) in _Using Oracle Compute Cloud Service_ for more information about shapes.`,
				},
				resource.Attribute{
					Name:        "metering_frequency",
					Description: `(Optional). The billing frequency of the service instance. Allowed values are ` + "`" + `MONTHLY` + "`" + ` and ` + "`" + `HOURLY` + "`" + ``,
				},
				resource.Attribute{
					Name:        "region",
					Description: `(Optional). Specifies the region where the instance will be provisioned.`,
				},
				resource.Attribute{
					Name:        "availability_domain",
					Description: `(Optional) Name of the availability domain within the region where the Oracle Database Cloud Service instance is to be provisioned. This is applicable only if you wish to provision to an OCI instance.`,
				},
				resource.Attribute{
					Name:        "notification_email",
					Description: `(Optional) The email address to send notifications around successful or unsuccessful completions of the instance-creation operation.`,
				},
				resource.Attribute{
					Name:        "ip_network",
					Description: `(Optional) This attribute is only applicable to accounts where regions are supported. The three-part name of an IP network to which the service instance is added. For example: /Compute-identity_domain/user/object`,
				},
				resource.Attribute{
					Name:        "vm_user",
					Description: `(Optional) The user name of account to be created in the VM.`,
				},
				resource.Attribute{
					Name:        "backups",
					Description: `(Optional) Provides Cloud Storage information for how to implement service instance backups. Backups is documented below`,
				},
				resource.Attribute{
					Name:        "mysql_configuration",
					Description: `(Required) Specified the detail of how to configure the MySQL database. mysql_configuration is documented below. ` + "`" + `backups` + "`" + ` support the following :`,
				},
				resource.Attribute{
					Name:        "cloud_storage_container",
					Description: `(Required). Name of the Oracle Storage Cloud container used for store the backups.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_username",
					Description: `(Required) Username for the Oracle Storage Cloud administrator.`,
				},
				resource.Attribute{
					Name:        "cloud_storage_password",
					Description: `(Required) Password for the Oracle Storage Cloud administrator.`,
				},
				resource.Attribute{
					Name:        "create_if_missing",
					Description: `(Optional) Specifies whether to create the container if it does not exist. Default value is ` + "`" + `false` + "`" + ` ` + "`" + `mysql_configuration` + "`" + ` supports the following :`,
				},
				resource.Attribute{
					Name:        "db_name",
					Description: `(Optional). The name of the database instance. Default value is ` + "`" + `mydatabase` + "`" + ``,
				},
				resource.Attribute{
					Name:        "db_storage",
					Description: `(Optional). The storage volume sice for MySQL data. The value must be between 25 to 1024. Defaults to 25 (GB)`,
				},
				resource.Attribute{
					Name:        "mysql_charset",
					Description: `(Optional) MySQL server character set. See [Supported Character Sets and Collation](http://dev.mysql.com/doc/en/charset-charsets.html). Default value is ` + "`" + `utf8mb4` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mysql_port",
					Description: `(Optional) The port number for the MySQL Server. The value must be between 3200-3399. Default value is ` + "`" + `3306` + "`" + ``,
				},
				resource.Attribute{
					Name:        "mysql_username",
					Description: `(Optional) The Administration user for connecting to the service via th MySQL protocol. Default value is ` + "`" + `root` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mysql_password",
					Description: `(Optional) The password for the MySQL Administration user.`,
				},
				resource.Attribute{
					Name:        "source_service_name",
					Description: `(Optional) When present, indicates that the service instance should be created as a "snapshot clone" of another service instance. Provide the name of the existing service instance whose snapshot is to be used. ` + "`" + `db_name` + "`" + `, ` + "`" + `mysql_charset` + "`" + `, ` + "`" + `mysql_collation` + "`" + `, ` + "`" + `enterpriseMonitor` + "`" + `, and associated MySQL server component parameters do not apply when cloning a service from a snapshot. For those parameters, the clone operation uses the values defined in the snapshot of the source service instance.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Optional) The name of the snapshot of the service instance specified by ` + "`" + `source_service_name` + "`" + ` that is to be used to create a "snapshot clone". This parameter is valid only if ` + "`" + `source_service_name` + "`" + ` is specified.`,
				},
				resource.Attribute{
					Name:        "enterprise_monitor_configuration",
					Description: `(Optional) Provides the Enterprise Monitor configuration for the MySQL Instance. If this is omitted, there will be no EM created for the MySQL Instance. ` + "`" + `enterprise_monitor_configuration` + "`" + ` is documented below. ` + "`" + `enterprise_monitor_configuration` + "`" + ` supports the following :`,
				},
				resource.Attribute{
					Name:        "em_agent_username",
					Description: `(Optional). Name for the Enterprise Monitor agent user.`,
				},
				resource.Attribute{
					Name:        "em_agent_password",
					Description: `(Optional). Password for MySQL Enterprise Monitor agent.`,
				},
				resource.Attribute{
					Name:        "em_username",
					Description: `(Optional) Name for the Enterprise Monitor Manager user.`,
				},
				resource.Attribute{
					Name:        "em_password",
					Description: `(Optional) Password for MySQL Enterprise Monitor manager.`,
				},
				resource.Attribute{
					Name:        "em_port",
					Description: `(Optional) The port number for the MySQL Enterprise Monitor instance. The default is 18443.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"oraclepaas_application_container":     0,
		"oraclepaas_database_access_rule":      1,
		"oraclepaas_database_service_instance": 2,
		"oraclepaas_java_access_rule":          3,
		"oraclepaas_java_service_instance":     4,
		"oraclepaas_mysql_access_rule":         5,
		"oraclepaas_mysql_service_instance":    6,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
