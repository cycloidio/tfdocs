package fortianalyzer

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_dvm_cmd_add_device",
			Category:         "Device Manager",
			ShortDescription: `Add a device to the Device Manager database.`,
			Description:      ``,
			Keywords: []string{
				"device",
				"manager",
				"dvm",
				"cmd",
				"add",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fazadom",
					Description: `Name or ID of the ADOM where the command is to be executed on.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device. The structure of ` + "`" + `device` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `create_task - Create a new task in task manager database. nonblocking - The API will return immediately in for non-blocking call. This flag will be set automatically when the adding, importing, updating, and deleting a list of devices. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `create_task` + "`" + `, ` + "`" + `nonblocking` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "groups",
					Description: `Groups. The structure of ` + "`" + `groups` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.`,
				},
				resource.Attribute{
					Name:        "force_recreate",
					Description: `The argument is optional, if it is set, when the value changes, the resource will be re-created. The ` + "`" + `device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "adm_pass",
					Description: `<i>add real and promote device</i>.`,
				},
				resource.Attribute{
					Name:        "adm_usr",
					Description: `<i>add real and promote device</i>.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `<i>available for all operations</i>.`,
				},
				resource.Attribute{
					Name:        "deviceaction",
					Description: `Specify add device operations, or leave blank to add real device:<ul><li>"add_model" - add a model device.<li>"promote_unreg" - promote an unregistered device to be managed by FortiManager using information from database.</ul>`,
				},
				resource.Attribute{
					Name:        "fazquota",
					Description: `<i>available for all operations</i>.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `<i>add real device only</i>. Add device will probe with this IP using the log in credential specified.`,
				},
				resource.Attribute{
					Name:        "metafields",
					Description: `<i>add real and model device</i>.`,
				},
				resource.Attribute{
					Name:        "mgmt_mode",
					Description: `<i>add real and model device</i>. Valid values: ` + "`" + `unreg` + "`" + `, ` + "`" + `fmg` + "`" + `, ` + "`" + `faz` + "`" + `, ` + "`" + `fmgfaz` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mr",
					Description: `<i>add model device only</i>.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `<i>required for all operations</i>. Unique name for the device.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `<i>add model device only</i>. Valid values: ` + "`" + `unknown` + "`" + `, ` + "`" + `fos` + "`" + `, ` + "`" + `fsw` + "`" + `, ` + "`" + `foc` + "`" + `, ` + "`" + `fml` + "`" + `, ` + "`" + `faz` + "`" + `, ` + "`" + `fwb` + "`" + `, ` + "`" + `fch` + "`" + `, ` + "`" + `fct` + "`" + `, ` + "`" + `log` + "`" + `, ` + "`" + `fmg` + "`" + `, ` + "`" + `fsa` + "`" + `, ` + "`" + `fdd` + "`" + `, ` + "`" + `fac` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "os_ver",
					Description: `<i>add model device only</i>. Valid values: ` + "`" + `unknown` + "`" + `, ` + "`" + `0.0` + "`" + `, ` + "`" + `1.0` + "`" + `, ` + "`" + `2.0` + "`" + `, ` + "`" + `3.0` + "`" + `, ` + "`" + `4.0` + "`" + `, ` + "`" + `5.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "patch",
					Description: `<i>add model device only</i>.`,
				},
				resource.Attribute{
					Name:        "platform_str",
					Description: `<i>add model device only</i>. Required for determine the platform for VM platforms.`,
				},
				resource.Attribute{
					Name:        "sn",
					Description: `<i>add model device only</i>. This attribute will be used to determine the device platform, except for VM platforms, where <i>platform_str</i> is also required. The ` + "`" + `groups` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Dvm CmdAddDevice can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvm_cmd_add_device.labelname DvmCmdAddDevice $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ` ## Others ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Dvm CmdAddDevice can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvm_cmd_add_device.labelname DvmCmdAddDevice $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ` ## Others ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_dvm_cmd_del_device",
			Category:         "Device Manager",
			ShortDescription: `Delete a device.`,
			Description:      ``,
			Keywords: []string{
				"device",
				"manager",
				"dvm",
				"cmd",
				"del",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fazadom",
					Description: `Name or ID of the ADOM where the command is to be executed on.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Name or ID of the target device.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `create_task - Create a new task in task manager database. nonblocking - The API will return immediately in for non-blocking call. This flag will be set automatically when the adding, importing, updating, and deleting a list of devices. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `create_task` + "`" + `, ` + "`" + `nonblocking` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_recreate",
					Description: `The argument is optional, if it is set, when the value changes, the resource will be re-created. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Dvm CmdDelDevice can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvm_cmd_del_device.labelname DvmCmdDelDevice $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ` ## Others ~>`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Dvm CmdDelDevice can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvm_cmd_del_device.labelname DvmCmdDelDevice $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ` ## Others ~>`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_dvmdb_adom",
			Category:         "Device Manager",
			ShortDescription: `ADOM table, most attributes are read-only and can only be changed internally.`,
			Description:      ``,
			Keywords: []string{
				"device",
				"manager",
				"dvmdb",
				"adom",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "create_time",
					Description: `Create_Time.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `Desc.`,
				},
				resource.Attribute{
					Name:        "flags",
					Description: `Flags. Valid values: ` + "`" + `migration` + "`" + `, ` + "`" + `db_export` + "`" + `, ` + "`" + `no_vpn_console` + "`" + `, ` + "`" + `backup` + "`" + `, ` + "`" + `other_devices` + "`" + `, ` + "`" + `is_autosync` + "`" + `, ` + "`" + `per_device_wtp` + "`" + `, ` + "`" + `policy_check_on_install` + "`" + `, ` + "`" + `install_on_policy_check_fail` + "`" + `, ` + "`" + `auto_push_cfg` + "`" + `, ` + "`" + `per_device_fsw` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_db_retention_hours",
					Description: `Log_Db_Retention_Hours.`,
				},
				resource.Attribute{
					Name:        "log_disk_quota",
					Description: `Log_Disk_Quota.`,
				},
				resource.Attribute{
					Name:        "log_disk_quota_alert_thres",
					Description: `Log_Disk_Quota_Alert_Thres.`,
				},
				resource.Attribute{
					Name:        "log_disk_quota_split_ratio",
					Description: `Log_Disk_Quota_Split_Ratio.`,
				},
				resource.Attribute{
					Name:        "log_file_retention_hours",
					Description: `Log_File_Retention_Hours.`,
				},
				resource.Attribute{
					Name:        "metafields",
					Description: `Default metafields: none.`,
				},
				resource.Attribute{
					Name:        "mig_mr",
					Description: `Mig_Mr.`,
				},
				resource.Attribute{
					Name:        "mig_os_ver",
					Description: `Mig_Os_Ver. Valid values: ` + "`" + `unknown` + "`" + `, ` + "`" + `0.0` + "`" + `, ` + "`" + `1.0` + "`" + `, ` + "`" + `2.0` + "`" + `, ` + "`" + `3.0` + "`" + `, ` + "`" + `4.0` + "`" + `, ` + "`" + `5.0` + "`" + `, ` + "`" + `6.0` + "`" + `, ` + "`" + `7.0` + "`" + `, ` + "`" + `8.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `ems - (Value no longer used as of 4.3) provider - Global database. Valid values: ` + "`" + `ems` + "`" + `, ` + "`" + `gms` + "`" + `, ` + "`" + `provider` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mr",
					Description: `Mr.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "os_ver",
					Description: `Os_Ver. Valid values: ` + "`" + `unknown` + "`" + `, ` + "`" + `0.0` + "`" + `, ` + "`" + `1.0` + "`" + `, ` + "`" + `2.0` + "`" + `, ` + "`" + `3.0` + "`" + `, ` + "`" + `4.0` + "`" + `, ` + "`" + `5.0` + "`" + `, ` + "`" + `6.0` + "`" + `, ` + "`" + `7.0` + "`" + `, ` + "`" + `8.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "restricted_prds",
					Description: `Restricted_Prds. Valid values: ` + "`" + `fos` + "`" + `, ` + "`" + `foc` + "`" + `, ` + "`" + `fml` + "`" + `, ` + "`" + `fch` + "`" + `, ` + "`" + `fwb` + "`" + `, ` + "`" + `log` + "`" + `, ` + "`" + `fct` + "`" + `, ` + "`" + `faz` + "`" + `, ` + "`" + `fsa` + "`" + `, ` + "`" + `fsw` + "`" + `, ` + "`" + `fmg` + "`" + `, ` + "`" + `fdd` + "`" + `, ` + "`" + `fac` + "`" + `, ` + "`" + `fpx` + "`" + `, ` + "`" + `fna` + "`" + `, ` + "`" + `ffw` + "`" + `, ` + "`" + `fsr` + "`" + `, ` + "`" + `fad` + "`" + `, ` + "`" + `fdc` + "`" + `, ` + "`" + `fap` + "`" + `, ` + "`" + `fxt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "state",
					Description: `State.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `Uuid.`,
				},
				resource.Attribute{
					Name:        "workspace_mode",
					Description: `Workspace_Mode. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import Dvmdb Adom can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvmdb_adom.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import Dvmdb Adom can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvmdb_adom.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_dvmdb_group",
			Category:         "Device Manager",
			ShortDescription: `Device group table.`,
			Description:      ``,
			Keywords: []string{
				"device",
				"manager",
				"dvmdb",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scopetype",
					Description: `The scope of application of the resource. Valid values: ` + "`" + `inherit` + "`" + `, ` + "`" + `adom` + "`" + `. The ` + "`" + `inherit` + "`" + ` means that the scopetype of the provider will be inherited, and adom will also be inherited. The default value is ` + "`" + `inherit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Adom. This value is valid only when the ` + "`" + `scopetype` + "`" + ` is ` + "`" + `adom` + "`" + `, otherwise the value of adom in the provider will be inherited.`,
				},
				resource.Attribute{
					Name:        "desc",
					Description: `Desc.`,
				},
				resource.Attribute{
					Name:        "metafields",
					Description: `Default metafields: none.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "os_type",
					Description: `Os_Type. Valid values: ` + "`" + `unknown` + "`" + `, ` + "`" + `fos` + "`" + `, ` + "`" + `fsw` + "`" + `, ` + "`" + `foc` + "`" + `, ` + "`" + `fml` + "`" + `, ` + "`" + `faz` + "`" + `, ` + "`" + `fwb` + "`" + `, ` + "`" + `fch` + "`" + `, ` + "`" + `fct` + "`" + `, ` + "`" + `log` + "`" + `, ` + "`" + `fmg` + "`" + `, ` + "`" + `fsa` + "`" + `, ` + "`" + `fdd` + "`" + `, ` + "`" + `fac` + "`" + `, ` + "`" + `fpx` + "`" + `, ` + "`" + `fna` + "`" + `, ` + "`" + `ffw` + "`" + `, ` + "`" + `fsr` + "`" + `, ` + "`" + `fad` + "`" + `, ` + "`" + `fdc` + "`" + `, ` + "`" + `fap` + "`" + `, ` + "`" + `fxt` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type. Valid values: ` + "`" + `normal` + "`" + `, ` + "`" + `default` + "`" + `, ` + "`" + `auto` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import Dvmdb Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvmdb_group.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import Dvmdb Group can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_dvmdb_group.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ` ->`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_analyzer_virusreport",
			Category:         "Fmupdate",
			ShortDescription: `Send virus detection notification to FortiGuard.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"analyzer",
				"virusreport",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable sending virus detection notification to FortiGuard (default = enable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate AnalyzerVirusreport can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_analyzer_virusreport.labelname FmupdateAnalyzerVirusreport $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate AnalyzerVirusreport can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_analyzer_virusreport.labelname FmupdateAnalyzerVirusreport $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_avips_advancedlog",
			Category:         "Fmupdate",
			ShortDescription: `Enable/disable logging of FortiGuard antivirus and IPS update packages received by FortiManager's built-in FortiGuard.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"avips",
				"advancedlog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "log_fortigate",
					Description: `Enable/disable logging of FortiGuard antivirus and IPS service updates of FortiGate devices (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_server",
					Description: `Enable/disable logging of update packages received by the build-in FortiGuard server (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate AvIpsAdvancedLog can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_avips_advancedlog.labelname FmupdateAvIpsAdvancedLog $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate AvIpsAdvancedLog can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_avips_advancedlog.labelname FmupdateAvIpsAdvancedLog $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_avips_webproxy",
			Category:         "Fmupdate",
			ShortDescription: `Configure the web proxy for use with FortiGuard antivirus and IPS updates.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"avips",
				"webproxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `web proxy address.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Web proxy mode proxy - HTTP proxy mode tunnel - HTTP tunnel mode (default) Valid values: ` + "`" + `proxy` + "`" + `, ` + "`" + `tunnel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password for the user name used for authentication.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the web proxy (1 - 65535, default = 80).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable connections through the web proxy (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The user name used for authentication. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate AvIpsWebProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_avips_webproxy.labelname FmupdateAvIpsWebProxy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate AvIpsWebProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_avips_webproxy.labelname FmupdateAvIpsWebProxy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_customurllist",
			Category:         "Fmupdate",
			ShortDescription: `Configure the URL database for rating and filtering.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"customurllist",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_selection",
					Description: `Manage the URL database (default = both). both - Support both custom-url and FortiGuard database. custom-url - Custom imported URL list. fortiguard-db - Fortinet's Fortiguard database. Valid values: ` + "`" + `both` + "`" + `, ` + "`" + `custom-url` + "`" + `, ` + "`" + `fortiguard-db` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate CustomUrlList can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_customurllist.labelname FmupdateCustomUrlList $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate CustomUrlList can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_customurllist.labelname FmupdateCustomUrlList $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_diskquota",
			Category:         "Fmupdate",
			ShortDescription: `Configure disk space available for use by the Upgrade Manager.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"diskquota",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "value",
					Description: `Configure the size of the Upgrade Manager disk quota, in megabytes. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate DiskQuota can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_diskquota.labelname FmupdateDiskQuota $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate DiskQuota can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_diskquota.labelname FmupdateDiskQuota $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fctservices",
			Category:         "Fmupdate",
			ShortDescription: `Configure FortiGuard to provide services to FortiClient installations.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fctservices",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "port",
					Description: `Configure the port number on which the built-in FortiGuard should provide updates to FortiClient installations (1 - 65535, default = 80).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable built-in FortiGuard service to FortiClient installations (default = enable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FctServices can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fctservices.labelname FmupdateFctServices $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FctServices can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fctservices.labelname FmupdateFctServices $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fdssetting",
			Category:         "Fmupdate",
			ShortDescription: `Configure FortiGuard settings.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fdssetting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "User_Agent",
					Description: `Configure the user agent string.`,
				},
				resource.Attribute{
					Name:        "fds_clt_ssl_protocol",
					Description: `The SSL protocols version for connecting fds server (default = tlsv1.2). sslv3 - set SSLv3 as the client version. tlsv1.0 - set TLSv1.0 as the client version. tlsv1.1 - set TLSv1.1 as the client version. tlsv1.2 - set TLSv1.2 as the client version (default). tlsv1.3 - set TLSv1.3 as the client version. Valid values: ` + "`" + `sslv3` + "`" + `, ` + "`" + `tlsv1.0` + "`" + `, ` + "`" + `tlsv1.1` + "`" + `, ` + "`" + `tlsv1.2` + "`" + `, ` + "`" + `tlsv1.3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fds_ssl_protocol",
					Description: `The SSL protocols version for receiving fgt connection (default = tlsv1.2). sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version (default). tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: ` + "`" + `sslv3` + "`" + `, ` + "`" + `tlsv1.0` + "`" + `, ` + "`" + `tlsv1.1` + "`" + `, ` + "`" + `tlsv1.2` + "`" + `, ` + "`" + `tlsv1.3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmtr_log",
					Description: `fmtr log level emergency - Log level - emergency alert - Log level - alert critical - Log level - critical error - Log level - error warn - Log level - warn notice - Log level - notice info - Log level - info debug - Log level - debug disable - Disable linkd log Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warn` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `debug` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiguard_anycast",
					Description: `Enable/disable use of FortiGuard's anycast network disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiguard_anycast_source",
					Description: `Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet fortinet - Use Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. aws - Use Fortinet's AWS servers to provide FortiGuard services in FortiGuard's anycast network. Valid values: ` + "`" + `fortinet` + "`" + `, ` + "`" + `aws` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "linkd_log",
					Description: `The linkd log level (default = info). emergency - Log level - emergency alert - Log level - alert critical - Log level - critical error - Log level - error warn - Log level - warn notice - Log level - notice info - Log level - info debug - Log level - debug disable - Disable linkd log Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warn` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `debug` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_av_ips_version",
					Description: `The maximum number of downloadable, full version AV/IPS packages (1 - 1000, default = 20).`,
				},
				resource.Attribute{
					Name:        "max_work",
					Description: `The maximum number of worker processing download requests (1 - 32, default = 1).`,
				},
				resource.Attribute{
					Name:        "push_override",
					Description: `Push-Override. The structure of ` + "`" + `push_override` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "push_override_to_client",
					Description: `Push-Override-To-Client. The structure of ` + "`" + `push_override_to_client` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "send_report",
					Description: `send report/fssi to fds server. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "send_setup",
					Description: `forward setup to fds server. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_override",
					Description: `Server-Override. The structure of ` + "`" + `server_override` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "system_support_fct",
					Description: `Supported FortiClient versions. 4.x - Support version 4.x 5.0 - Support version 5.0 5.2 - Support version 5.2 5.4 - Support version 5.4 5.6 - Support version 5.6 6.0 - Support version 6.0 6.2 - Support version 6.2 6.4 - Support version 6.4 Valid values: ` + "`" + `4.x` + "`" + `, ` + "`" + `5.0` + "`" + `, ` + "`" + `5.2` + "`" + `, ` + "`" + `5.4` + "`" + `, ` + "`" + `5.6` + "`" + `, ` + "`" + `6.0` + "`" + `, ` + "`" + `6.2` + "`" + `, ` + "`" + `6.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_support_fdc",
					Description: `Supported FortiDeceptor versions. 3.x - Support version 3.x Valid values: ` + "`" + `3.x` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_support_fgt",
					Description: `Supported FortiOS versions. 5.4 - Support version 5.4 5.6 - Support version 5.6 6.0 - Support version 6.0 6.2 - Support version 6.2 6.4 - Support version 6.4 7.0 - Support version 7.0 Valid values: ` + "`" + `5.4` + "`" + `, ` + "`" + `5.6` + "`" + `, ` + "`" + `6.0` + "`" + `, ` + "`" + `6.2` + "`" + `, ` + "`" + `6.4` + "`" + `, ` + "`" + `7.0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_support_fml",
					Description: `Supported FortiMail versions. 4.x - Support version 4.x 5.x - Support version 5.x 6.x - Support version 6.x Valid values: ` + "`" + `4.x` + "`" + `, ` + "`" + `5.x` + "`" + `, ` + "`" + `6.x` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_support_fsa",
					Description: `Supported FortiSandbox versions. 1.x - Support version 1.x 2.x - Support version 2.x 3.x - Support version 3.x 4.x - Support version 4.x Valid values: ` + "`" + `1.x` + "`" + `, ` + "`" + `2.x` + "`" + `, ` + "`" + `3.x` + "`" + `, ` + "`" + `4.x` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_support_fts",
					Description: `Supported FortiTester versions. 4.x - Support version 4.x Valid values: ` + "`" + `4.x` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_support_fsw",
					Description: `Supported FortiSwitch versions. 4.x - Support version 4.x 5.0 - Support version 5.0 5.2 - Support version 5.2 5.4 - Support version 5.4 5.6 - Support version 5.6 6.0 - Support version 6.0 6.2 - Support version 6.2 6.4 - Support version 6.4 Valid values: ` + "`" + `4.x` + "`" + `, ` + "`" + `5.0` + "`" + `, ` + "`" + `5.2` + "`" + `, ` + "`" + `5.4` + "`" + `, ` + "`" + `5.6` + "`" + `, ` + "`" + `6.0` + "`" + `, ` + "`" + `6.2` + "`" + `, ` + "`" + `6.4` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "umsvc_log",
					Description: `The um_service log level (default = info). emergency - Log level - emergency alert - Log level - alert critical - Log level - critical error - Log level - error warn - Log level - warn notice - Log level - notice info - Log level - info debug - Log level - debug disable - Disable linkd log Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warn` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `debug` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unreg_dev_option",
					Description: `set the option for unregister devices ignore - Ignore all unregistered devices. svc-only - Allow update requests without adding the device. add-service - Add unregistered devices and allow update request. Valid values: ` + "`" + `ignore` + "`" + `, ` + "`" + `svc-only` + "`" + `, ` + "`" + `add-service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "update_schedule",
					Description: `Update-Schedule. The structure of ` + "`" + `update_schedule` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "wanip_query_mode",
					Description: `public ip query mode disable - Do not query public ip ipify - Get public IP through https://api.ipify.org Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `ipify` + "`" + `. The ` + "`" + `push_override` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `External or virtual IP address of the NAT device that will forward push messages to the FortiManager unit.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Receiving port number on the NAT device (1 - 65535, default = 9443).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable push updates for clients (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `push_override_to_client` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "announce_ip",
					Description: `Announce-Ip. The structure of ` + "`" + `announce_ip` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable push updates (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `announce_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the announce IP address (1 - 10).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Announce IPv4 address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Announce IP port (1 - 65535, default = 8890). The ` + "`" + `server_override` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "servlist",
					Description: `Servlist. The structure of ` + "`" + `servlist` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Override status. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `servlist` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Override server ID (1 - 10).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IPv4 address of the override server.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IPv6 address of the override server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number to use when contacting FortiGuard (1 - 65535, default = 443).`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Override service type. fct - Server override config for fct fds - Server override config for fds Valid values: ` + "`" + `fct` + "`" + `, ` + "`" + `fds` + "`" + `. The ` + "`" + `update_schedule` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "day",
					Description: `Configure the day the update will occur, if the freqnecy is weekly (Sunday - Saturday, default = Monday). Sunday - Update every Sunday. Monday - Update every Monday. Tuesday - Update every Tuesday. Wednesday - Update every Wednesday. Thursday - Update every Thursday. Friday - Update every Friday. Saturday - Update every Saturday. Valid values: ` + "`" + `Sunday` + "`" + `, ` + "`" + `Monday` + "`" + `, ` + "`" + `Tuesday` + "`" + `, ` + "`" + `Wednesday` + "`" + `, ` + "`" + `Thursday` + "`" + `, ` + "`" + `Friday` + "`" + `, ` + "`" + `Saturday` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Configure update frequency: every - time interval, daily - once a day, weekly - once a week (default = every). every - Time interval. daily - Every day. weekly - Every week. Valid values: ` + "`" + `every` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable scheduled updates. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Time interval between updates, or the hour and minute when the update occurs (hh: 0 - 23, mm: 0 - 59 or 60 = random, default = 00:10). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting.labelname FmupdateFdsSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting.labelname FmupdateFdsSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fdssetting_pushoverride",
			Category:         "Fmupdate",
			ShortDescription: `Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fdssetting",
				"pushoverride",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip",
					Description: `External or virtual IP address of the NAT device that will forward push messages to the FortiManager unit.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Receiving port number on the NAT device (1 - 65535, default = 9443).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable push updates for clients (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingPushOverride can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_pushoverride.labelname FmupdateFdsSettingPushOverride $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingPushOverride can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_pushoverride.labelname FmupdateFdsSettingPushOverride $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fdssetting_pushoverridetoclient",
			Category:         "Fmupdate",
			ShortDescription: `Enable/disable push updates, and override the default IP address and port used by FortiGuard to send antivirus and IPS push messages for clients.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fdssetting",
				"pushoverridetoclient",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "announce_ip",
					Description: `Announce-Ip. The structure of ` + "`" + `announce_ip` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable push updates (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `announce_ip` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the announce IP address (1 - 10).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Announce IPv4 address.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Announce IP port (1 - 65535, default = 8890). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingPushOverrideToClient can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_pushoverridetoclient.labelname FmupdateFdsSettingPushOverrideToClient $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingPushOverrideToClient can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_pushoverridetoclient.labelname FmupdateFdsSettingPushOverrideToClient $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fdssetting_serveroverride",
			Category:         "Fmupdate",
			ShortDescription: `Server override configure.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fdssetting",
				"serveroverride",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "servlist",
					Description: `Servlist. The structure of ` + "`" + `servlist` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Override status. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `servlist` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Override server ID (1 - 10).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IPv4 address of the override server.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IPv6 address of the override server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number to use when contacting FortiGuard (1 - 65535, default = 443).`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Override service type. fct - Server override config for fct fds - Server override config for fds Valid values: ` + "`" + `fct` + "`" + `, ` + "`" + `fds` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingServerOverride can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_serveroverride.labelname FmupdateFdsSettingServerOverride $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingServerOverride can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_serveroverride.labelname FmupdateFdsSettingServerOverride $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fdssetting_updateschedule",
			Category:         "Fmupdate",
			ShortDescription: `Configure the schedule when built-in FortiGuard retrieves antivirus and IPS updates.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fdssetting",
				"updateschedule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "day",
					Description: `Configure the day the update will occur, if the freqnecy is weekly (Sunday - Saturday, default = Monday). Sunday - Update every Sunday. Monday - Update every Monday. Tuesday - Update every Tuesday. Wednesday - Update every Wednesday. Thursday - Update every Thursday. Friday - Update every Friday. Saturday - Update every Saturday. Valid values: ` + "`" + `Sunday` + "`" + `, ` + "`" + `Monday` + "`" + `, ` + "`" + `Tuesday` + "`" + `, ` + "`" + `Wednesday` + "`" + `, ` + "`" + `Thursday` + "`" + `, ` + "`" + `Friday` + "`" + `, ` + "`" + `Saturday` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Configure update frequency: every - time interval, daily - once a day, weekly - once a week (default = every). every - Time interval. daily - Every day. weekly - Every week. Valid values: ` + "`" + `every` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable scheduled updates. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Time interval between updates, or the hour and minute when the update occurs (hh: 0 - 23, mm: 0 - 59 or 60 = random, default = 00:10). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingUpdateSchedule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_updateschedule.labelname FmupdateFdsSettingUpdateSchedule $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FdsSettingUpdateSchedule can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fdssetting_updateschedule.labelname FmupdateFdsSettingUpdateSchedule $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_fwmsetting",
			Category:         "Fmupdate",
			ShortDescription: `Configure firmware management settings.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"fwmsetting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auto_scan_fgt_disk",
					Description: `auto scan fgt disk if needed. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "check_fgt_disk",
					Description: `check fgt disk before upgrade image. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fds_failover_fmg",
					Description: `using fmg local image file is download from fds fails. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fds_image_timeout",
					Description: `timer for fgt download image from fortiguard (300-3600s default=1800)`,
				},
				resource.Attribute{
					Name:        "immx_source",
					Description: `Configure which of IMMX file to be used for choosing upgrade path. Default is file for FortiManager fmg - Use IMMX file for FortiManager fgt - Use IMMX file for FortiGate cloud - Use IMMX file for FortiCloud Valid values: ` + "`" + `fmg` + "`" + `, ` + "`" + `fgt` + "`" + `, ` + "`" + `cloud` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log",
					Description: `Configure log setting for fwm daemon fwm - FWM daemon log fwm_dm - FWM and Deployment service log fwm_dm_json - FWM and Deployment service log with JSON data between FMG-FGT Valid values: ` + "`" + `fwm` + "`" + `, ` + "`" + `fwm_dm` + "`" + `, ` + "`" + `fwm_dm_json` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "multiple_steps_interval",
					Description: `waiting time between multiple steps upgrade (30-180s, default=60) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FwmSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fwmsetting.labelname FmupdateFwmSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate FwmSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_fwmsetting.labelname FmupdateFwmSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_multilayer",
			Category:         "Fmupdate",
			ShortDescription: `Configure multilayer mode.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"multilayer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "webspam_rating",
					Description: `Enable/disable URL/Antispam rating service (default = enable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate Multilayer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_multilayer.labelname FmupdateMultilayer $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate Multilayer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_multilayer.labelname FmupdateMultilayer $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_publicnetwork",
			Category:         "Fmupdate",
			ShortDescription: `Enable/disable access to the public FortiGuard.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"publicnetwork",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable public network (default = enable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate Publicnetwork can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_publicnetwork.labelname FmupdatePublicnetwork $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate Publicnetwork can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_publicnetwork.labelname FmupdatePublicnetwork $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_serveraccesspriorities",
			Category:         "Fmupdate",
			ShortDescription: `Configure priorities for FortiGate units accessing antivirus updates and web filtering services.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"serveraccesspriorities",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_public",
					Description: `Enable/disable FortiGates to Access Public FortiGuard Servers when Private Servers are Unavailable (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "av_ips",
					Description: `Enable/disable Antivirus and IPS Update Service for Private Server(default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_server",
					Description: `Private-Server. The structure of ` + "`" + `private_server` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "web_spam",
					Description: `Enable/disable Web Filter and Email Filter Update Service for Private Server (default = enable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `private_server` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Private server ID (1 - 10).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IPv4 address of the FortiManager unit or private server.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IPv6 address of the FortiManager unit or private server.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `Time zone of the private server (-24 = local time zone, default = -24). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate ServerAccessPriorities can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_serveraccesspriorities.labelname FmupdateServerAccessPriorities $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate ServerAccessPriorities can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_serveraccesspriorities.labelname FmupdateServerAccessPriorities $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_serveroverridestatus",
			Category:         "Fmupdate",
			ShortDescription: `Configure strict/loose server override.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"serveroverridestatus",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "mode",
					Description: `Server override mode (default = loose). strict - Access override server only. loose - Allow access other servers. Valid values: ` + "`" + `strict` + "`" + `, ` + "`" + `loose` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate ServerOverrideStatus can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_serveroverridestatus.labelname FmupdateServerOverrideStatus $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate ServerOverrideStatus can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_serveroverridestatus.labelname FmupdateServerOverrideStatus $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_service",
			Category:         "Fmupdate",
			ShortDescription: `Enable/disable services provided by the built-in FortiGuard.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"service",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "avips",
					Description: `Enable/disable the built-in FortiGuard to provide FortiGuard antivirus and IPS updates (default = enable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate Service can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_service.labelname FmupdateService $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate Service can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_service.labelname FmupdateService $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_webspam_fgdsetting",
			Category:         "Fmupdate",
			ShortDescription: `Configure the FortiGuard run parameters.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"webspam",
				"fgdsetting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "as_cache",
					Description: `Antispam service maximum memory usage in megabytes (Maximum = Physical memory-1024, 0: no limit, default = 300).`,
				},
				resource.Attribute{
					Name:        "as_log",
					Description: `Antispam log setting (default = nospam). disable - Disable spam log. nospam - Log non-spam events. all - Log all spam lookups. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `nospam` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "as_preload",
					Description: `Enable/disable preloading antispam database to memory (default = disable). disable - Disable antispam database preload. enable - Enable antispam database preload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "av_cache",
					Description: `Antivirus service maximum memory usage, in megabytes (100 - 500, default = 300).`,
				},
				resource.Attribute{
					Name:        "av_log",
					Description: `Antivirus log setting (default = novirus). disable - Disable virus log. novirus - Log non-virus events. all - Log all virus lookups. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `novirus` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "av_preload",
					Description: `Enable/disable preloading antivirus database to memory (default = disable). disable - Disable antivirus database preload. enable - Enable antivirus database preload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "av2_cache",
					Description: `Antispam service maximum memory usage in megabytes (Maximum = Physical memory-1024, 0: no limit, default = 800).`,
				},
				resource.Attribute{
					Name:        "av2_log",
					Description: `Outbreak prevention log setting (default = noav2). disable - Disable av2 log. noav2 - Log non-av2 events. all - Log all av2 lookups. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `noav2` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "av2_preload",
					Description: `Enable/disable preloading outbreak prevention database to memory (default = disable). disable - Disable outbreak prevention database preload. enable - Enable outbreak prevention database preload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventlog_query",
					Description: `Enable/disable record query to event-log besides fgd-log (default = disable). disable - Record query to event-log besides fgd-log. enable - Do not log to event-log. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd_pull_interval",
					Description: `Fgd pull interval setting, in minutes (1 - 1440, default = 10).`,
				},
				resource.Attribute{
					Name:        "fq_cache",
					Description: `File query service maximum memory usage, in megabytes (100 - 500, default = 300).`,
				},
				resource.Attribute{
					Name:        "fq_log",
					Description: `File query log setting (default = nofilequery). disable - Disable file query log. nofilequery - Log non-file query events. all - Log all file query events. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `nofilequery` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fq_preload",
					Description: `Enable/disable preloading file query database to memory (default = disable). disable - Disable file query db preload. enable - Enable file query db preload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iot_cache",
					Description: `IoT service maximum memory usage, in megabytes (100 - 500, default = 300).`,
				},
				resource.Attribute{
					Name:        "iot_log",
					Description: `IoT log setting (default = nofilequery). disable - Disable IoT log. nofilequery - Log non-IoT events. all - Log all IoT events. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `nofilequery` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iot_preload",
					Description: `Enable/disable preloading IoT database to memory (default = disable). disable - Disable IoT db preload. enable - Enable IoT db preload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "linkd_log",
					Description: `Linkd log setting (default = debug). emergency - The unit is unusable. alert - Immediate action is required critical - Functionality is affected. error - Functionality is probably affected. warn - Functionality might be affected. notice - Information about normal events. info - General information. debug - Debug information. disable - Linkd logging is disabled. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warn` + "`" + `, ` + "`" + `notice` + "`" + `, ` + "`" + `info` + "`" + `, ` + "`" + `debug` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_client_worker",
					Description: `max worker for tcp client connection (0~16: 0 means use cpu number up to 4).`,
				},
				resource.Attribute{
					Name:        "max_log_quota",
					Description: `Maximum log quota setting, in megabytes (100 - 20480, default = 6144).`,
				},
				resource.Attribute{
					Name:        "max_unrated_site",
					Description: `Maximum number of unrated site in memory, in kilobytes(10 - 5120, default = 500).`,
				},
				resource.Attribute{
					Name:        "restrict_as1_dbver",
					Description: `Restrict system update to indicated antispam(1) database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_as2_dbver",
					Description: `Restrict system update to indicated antispam(2) database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_as4_dbver",
					Description: `Restrict system update to indicated antispam(4) database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_av_dbver",
					Description: `Restrict system update to indicated antivirus database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_av2_dbver",
					Description: `Restrict system update to indicated outbreak prevention database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_fq_dbver",
					Description: `Restrict system update to indicated file query database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_iots_dbver",
					Description: `Restrict system update to indicated file query database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "restrict_wf_dbver",
					Description: `Restrict system update to indicated web filter database version (character limit = 127).`,
				},
				resource.Attribute{
					Name:        "server_override",
					Description: `Server-Override. The structure of ` + "`" + `server_override` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "stat_log_interval",
					Description: `Statistic log interval setting, in minutes (1 - 1440, default = 60).`,
				},
				resource.Attribute{
					Name:        "stat_sync_interval",
					Description: `Synchronization interval for statistic of unrated site in minutes (1 - 60, default = 60).`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `FortiGuard database update wait time if not enough delta files, in hours (2 - 24, default = 6).`,
				},
				resource.Attribute{
					Name:        "update_log",
					Description: `Enable/disable update log setting (default = enable). disable - Disable update log. enable - Enable update log. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wf_cache",
					Description: `Web filter service maximum memory usage, in megabytes (maximum = Physical memory-1024, 0 = no limit, default = 600).`,
				},
				resource.Attribute{
					Name:        "wf_dn_cache_expire_time",
					Description: `Web filter DN cache expire time, in minutes (1 - 1440, 0 = never, default = 30).`,
				},
				resource.Attribute{
					Name:        "wf_dn_cache_max_number",
					Description: `Maximum number of Web filter DN cache (0 = disable, default = 10000).`,
				},
				resource.Attribute{
					Name:        "wf_log",
					Description: `Web filter log setting (default = nour1) disable - Disable URL log. nourl - Log non-URL events. all - Log all URL lookups. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `nourl` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wf_preload",
					Description: `Enable/disable preloading the web filter database into memory (default = disable). disable - Disable web filter database preload. enable - Enable web filter database preload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `server_override` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "servlist",
					Description: `Servlist. The structure of ` + "`" + `servlist` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Override status. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `servlist` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Override server ID (1 - 10).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IPv4 address of the override server.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IPv6 address of the override server.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number to use when contacting FortiGuard (1 - 65535, default = 443).`,
				},
				resource.Attribute{
					Name:        "service_type",
					Description: `Override service type. fgd - Server override config for fgd fgc - Server override config for fgc fsa - Server override config for fsa Valid values: ` + "`" + `fgd` + "`" + `, ` + "`" + `fgc` + "`" + `, ` + "`" + `fsa` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate WebSpamFgdSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_webspam_fgdsetting.labelname FmupdateWebSpamFgdSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate WebSpamFgdSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_webspam_fgdsetting.labelname FmupdateWebSpamFgdSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_fmupdate_webspam_webproxy",
			Category:         "Fmupdate",
			ShortDescription: `Configure the web proxy for use with FortiGuard antivirus and IPS updates.`,
			Description:      ``,
			Keywords: []string{
				"fmupdate",
				"webspam",
				"webproxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `web proxy address.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Web proxy mode proxy - HTTP proxy mode tunnel - HTTP tunnel mode (default) Valid values: ` + "`" + `proxy` + "`" + `, ` + "`" + `tunnel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password for the user name used for authentication.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the web proxy (1 - 65535, default = 80).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable connections through the web proxy (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The user name used for authentication. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate WebSpamWebProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_webspam_webproxy.labelname FmupdateWebSpamWebProxy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import Fmupdate WebSpamWebProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_fmupdate_webspam_webproxy.labelname FmupdateWebSpamWebProxy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_json_generic_api",
			Category:         "Generic",
			ShortDescription: `FortiAnalyzer API Generic Interface.`,
			Description:      ``,
			Keywords: []string{
				"generic",
				"json",
				"api",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "force_recreate",
					Description: `The argument is optional, if it is set, when its value changes, the resource will be re-created. It is usually used when the return value needs to be forced to update.`,
				},
				resource.Attribute{
					Name:        "json_content",
					Description: `Body data in JSON format.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Comment.`,
				},
				resource.Attribute{
					Name:        "response",
					Description: `API returns results. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The resource id.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_group",
			Category:         "System Admin",
			ShortDescription: `User group.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "member",
					Description: `Member. The structure of ` + "`" + `member` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `member` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Group member name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminGroup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_group.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminGroup can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_group.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_ldap",
			Category:         "System Admin",
			ShortDescription: `LDAP server entry configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"ldap",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fazadom",
					Description: `Adom. The structure of ` + "`" + `fazadom` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "adom_access",
					Description: `set all or specify adom access type. all - All ADOMs access. specify - Specify ADOMs access. Valid values: ` + "`" + `all` + "`" + `, ` + "`" + `specify` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom_attr",
					Description: `Attribute used to retrieve adom`,
				},
				resource.Attribute{
					Name:        "attributes",
					Description: `Attributes used for group searching.`,
				},
				resource.Attribute{
					Name:        "ca_cert",
					Description: `CA certificate name.`,
				},
				resource.Attribute{
					Name:        "cnid",
					Description: `Common Name Identifier (default = CN).`,
				},
				resource.Attribute{
					Name:        "connect_timeout",
					Description: `LDAP connection timeout (msec).`,
				},
				resource.Attribute{
					Name:        "dn",
					Description: `Distinguished Name.`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Filter used for group searching.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Full base DN used for group searching.`,
				},
				resource.Attribute{
					Name:        "memberof_attr",
					Description: `Attribute used to retrieve memeberof.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `LDAP server entry name.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for initial binding.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number of LDAP server (default = 389).`,
				},
				resource.Attribute{
					Name:        "profile_attr",
					Description: `Attribute used to retrieve admin profile.`,
				},
				resource.Attribute{
					Name:        "secondary_server",
					Description: `{<name_str|ip_str>} secondary LDAP server domain name or IP.`,
				},
				resource.Attribute{
					Name:        "secure",
					Description: `SSL connection. disable - No SSL. starttls - Use StartTLS. ldaps - Use LDAPS. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `starttls` + "`" + `, ` + "`" + `ldaps` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `{<name_str|ip_str>} LDAP server domain name or IP.`,
				},
				resource.Attribute{
					Name:        "tertiary_server",
					Description: `{<name_str|ip_str>} tertiary LDAP server domain name or IP.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of LDAP binding. simple - Simple password authentication without search. anonymous - Bind using anonymous user search. regular - Bind using username/password and then search. Valid values: ` + "`" + `simple` + "`" + `, ` + "`" + `anonymous` + "`" + `, ` + "`" + `regular` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Username (full DN) for initial binding.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `fazadom` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "adom_name",
					Description: `Admin domain names. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminLdap can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_ldap.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminLdap can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_ldap.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_profile",
			Category:         "System Admin",
			ShortDescription: `Admin profile.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "adom_lock",
					Description: `ADOM locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom_switch",
					Description: `Administrator domain. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_to_install",
					Description: `Enable/disable the restricted user to install objects to the devices. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "change_password",
					Description: `Enable/disable the user to change self password. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datamask",
					Description: `Enable/disable data masking. disable - Disable data masking. enable - Enable data masking. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datamask_custom_fields",
					Description: `Datamask-Custom-Fields. The structure of ` + "`" + `datamask_custom_fields` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "datamask_custom_priority",
					Description: `Prioritize custom fields. disable - Disable custom field search priority. enable - Enable custom field search priority. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datamask_fields",
					Description: `Data masking fields. user - User name. srcip - Source IP. srcname - Source name. srcmac - Source MAC. dstip - Destination IP. dstname - Dst name. email - Email. message - Message. domain - Domain. Valid values: ` + "`" + `user` + "`" + `, ` + "`" + `srcip` + "`" + `, ` + "`" + `srcname` + "`" + `, ` + "`" + `srcmac` + "`" + `, ` + "`" + `dstip` + "`" + `, ` + "`" + `dstname` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `message` + "`" + `, ` + "`" + `domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datamask_key",
					Description: `Data masking encryption key.`,
				},
				resource.Attribute{
					Name:        "datamask_unmasked_time",
					Description: `Time in days without data masking.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "device_ap",
					Description: `Manage AP. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_forticlient",
					Description: `Manage FortiClient. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_fortiswitch",
					Description: `Manage FortiSwitch. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_manager",
					Description: `Device manager. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_op",
					Description: `Device add/delete/edit. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_policy_package_lock",
					Description: `Device/Policy Package locking none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_wan_link_load_balance",
					Description: `Manage WAN link load balance. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event_management",
					Description: `Event management. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "execute_playbook",
					Description: `Execute playbook. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extension_access",
					Description: `Manage extension access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fabric_viewer",
					Description: `Fabric viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortirecorder_setting",
					Description: `FortiRecorder settings. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost1",
					Description: `Admin user trusted host IPv6, default ::/0 for all.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost10",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost2",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost3",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost4",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost5",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost6",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost7",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost8",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost9",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ips_baseline_ovrd",
					Description: `Enable/disable override baseline ips sensor. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_viewer",
					Description: `Log viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile ID.`,
				},
				resource.Attribute{
					Name:        "realtime_monitor",
					Description: `Realtime monitor. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report_viewer",
					Description: `Report viewer. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rpc_permit",
					Description: `Set none/read/read-write rpc-permission read-write - Read-write permission. none - No permission. read - Read-only permission. Valid values: ` + "`" + `read-write` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_report",
					Description: `Run reports. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scope",
					Description: `Scope. global - Global scope. adom - ADOM scope. Valid values: ` + "`" + `global` + "`" + `, ` + "`" + `adom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "script_access",
					Description: `Script access. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "super_user_profile",
					Description: `Enable/disable super user profile disable - Disable super user profile enable - Enable super user profile Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system_setting",
					Description: `System setting. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "triage_events",
					Description: `Triage events. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "trusthost1",
					Description: `Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.`,
				},
				resource.Attribute{
					Name:        "trusthost10",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost2",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost3",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost4",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost5",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost6",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost7",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost8",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost9",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "update_incidents",
					Description: `Create/update incidents. none - No permission. read - Read permission. read-write - Read-write permission. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `, ` + "`" + `read-write` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `datamask_custom_fields` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field_category",
					Description: `Field categories. log - Log. fortiview - FortiView. alert - Event management. ueba - UEBA. all - All. Valid values: ` + "`" + `log` + "`" + `, ` + "`" + `fortiview` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `ueba` + "`" + `, ` + "`" + `all` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "field_status",
					Description: `Field status. disable - Disable field. enable - Enable field. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "field_type",
					Description: `Field type. string - String. ip - IP. mac - MAC address. email - Email address. unknown - Unknown. Valid values: ` + "`" + `string` + "`" + `, ` + "`" + `ip` + "`" + `, ` + "`" + `mac` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `unknown` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{profileid}}. ## Import System AdminProfile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_profile.labelname {{profileid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{profileid}}. ## Import System AdminProfile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_profile.labelname {{profileid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_radius",
			Category:         "System Admin",
			ShortDescription: `Configure radius.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"radius",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_type",
					Description: `Authentication protocol. any - Use any supported authentication protocol. pap - PAP. chap - CHAP. mschap2 - MSCHAPv2. Valid values: ` + "`" + `any` + "`" + `, ` + "`" + `pap` + "`" + `, ` + "`" + `chap` + "`" + `, ` + "`" + `mschap2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "nas_ip",
					Description: `NAS IP address and called station ID.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Server port.`,
				},
				resource.Attribute{
					Name:        "secondary_secret",
					Description: `Secondary server secret.`,
				},
				resource.Attribute{
					Name:        "secondary_server",
					Description: `Secondary server name/IP.`,
				},
				resource.Attribute{
					Name:        "secret",
					Description: `Server secret.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Server name/IP. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminRadius can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_radius.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminRadius can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_radius.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_setting",
			Category:         "System Admin",
			ShortDescription: `Admin setting.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "access_banner",
					Description: `Enable/disable access banner. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "admin_https_redirect",
					Description: `Enable/disable redirection of HTTP admin traffic to HTTPS. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "admin_login_max",
					Description: `Maximum number admin users logged in at one time (1 - 256).`,
				},
				resource.Attribute{
					Name:        "admin_server_cert",
					Description: `HTTPS & Web Service server certificate.`,
				},
				resource.Attribute{
					Name:        "auth_addr",
					Description: `IP which is used by FGT to authorize FAZ.`,
				},
				resource.Attribute{
					Name:        "auth_port",
					Description: `Port which is used by FGT to authorize FAZ.`,
				},
				resource.Attribute{
					Name:        "banner_message",
					Description: `Banner message.`,
				},
				resource.Attribute{
					Name:        "gui_theme",
					Description: `Color scheme to use for the administration GUI. blue - Blueberry green - Kiwi red - Cherry melongene - Plum spring - Spring summer - Summer autumn - Autumn winter - Winter circuit-board - Circuit Board calla-lily - Calla Lily binary-tunnel - Binary Tunnel mars - Mars blue-sea - Blue Sea technology - Technology landscape - Landscape twilight - Twilight canyon - Canyon northern-light - Northern Light astronomy - Astronomy fish - Fish penguin - Penguin mountain - Mountain panda - Panda parrot - Parrot cave - Cave zebra - Zebra contrast-dark - High Contrast Dark Valid values: ` + "`" + `blue` + "`" + `, ` + "`" + `green` + "`" + `, ` + "`" + `red` + "`" + `, ` + "`" + `melongene` + "`" + `, ` + "`" + `spring` + "`" + `, ` + "`" + `summer` + "`" + `, ` + "`" + `autumn` + "`" + `, ` + "`" + `winter` + "`" + `, ` + "`" + `circuit-board` + "`" + `, ` + "`" + `calla-lily` + "`" + `, ` + "`" + `binary-tunnel` + "`" + `, ` + "`" + `mars` + "`" + `, ` + "`" + `blue-sea` + "`" + `, ` + "`" + `technology` + "`" + `, ` + "`" + `landscape` + "`" + `, ` + "`" + `twilight` + "`" + `, ` + "`" + `canyon` + "`" + `, ` + "`" + `northern-light` + "`" + `, ` + "`" + `astronomy` + "`" + `, ` + "`" + `fish` + "`" + `, ` + "`" + `penguin` + "`" + `, ` + "`" + `mountain` + "`" + `, ` + "`" + `panda` + "`" + `, ` + "`" + `parrot` + "`" + `, ` + "`" + `cave` + "`" + `, ` + "`" + `zebra` + "`" + `, ` + "`" + `contrast-dark` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "http_port",
					Description: `HTTP port.`,
				},
				resource.Attribute{
					Name:        "https_port",
					Description: `HTTPS port.`,
				},
				resource.Attribute{
					Name:        "idle_timeout",
					Description: `Idle timeout (60 - 28800 sec).`,
				},
				resource.Attribute{
					Name:        "idle_timeout_api",
					Description: `Idle timeout for API sessions (1 - 28800 sec).`,
				},
				resource.Attribute{
					Name:        "idle_timeout_gui",
					Description: `Idle timeout for GUI sessions (60 - 28800 sec).`,
				},
				resource.Attribute{
					Name:        "idle_timeout_sso",
					Description: `Idle timeout for SSO sessions (60 - 28800 sec).`,
				},
				resource.Attribute{
					Name:        "objects_force_deletion",
					Description: `Enable/disable used objects force deletion. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "preferred_fgfm_intf",
					Description: `Preferred interface for FGFM connection.`,
				},
				resource.Attribute{
					Name:        "shell_access",
					Description: `Enable/disable shell access. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shell_password",
					Description: `Password for shell access.`,
				},
				resource.Attribute{
					Name:        "show_add_multiple",
					Description: `Show add multiple button. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_checkbox_in_table",
					Description: `Show checkboxs in tables on GUI. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_device_import_export",
					Description: `Enable/disable import/export of ADOM, device, and group lists. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_fct_manager",
					Description: `Enable/disable FCT manager. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_hostname",
					Description: `Enable/disable hostname display in the GUI login page. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "show_log_forwarding",
					Description: `Show log forwarding tab in regular mode. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "unreg_dev_opt",
					Description: `Action to take when unregistered device connects to FortiAnalyzer. add_no_service - Add unregistered devices but deny service requests. ignore - Ignore unregistered devices. add_allow_service - Add unregistered devices and allow service requests. Valid values: ` + "`" + `add_no_service` + "`" + `, ` + "`" + `ignore` + "`" + `, ` + "`" + `add_allow_service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webadmin_language",
					Description: `Web admin language. auto_detect - Automatically detect language. english - English. simplified_chinese - Simplified Chinese. traditional_chinese - Traditional Chinese. japanese - Japanese. korean - Korean. spanish - Spanish. Valid values: ` + "`" + `auto_detect` + "`" + `, ` + "`" + `english` + "`" + `, ` + "`" + `simplified_chinese` + "`" + `, ` + "`" + `traditional_chinese` + "`" + `, ` + "`" + `japanese` + "`" + `, ` + "`" + `korean` + "`" + `, ` + "`" + `spanish` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AdminSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_setting.labelname SystemAdminSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AdminSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_setting.labelname SystemAdminSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_tacacs",
			Category:         "System Admin",
			ShortDescription: `TACACS+ server entry configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"tacacs",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "authen_type",
					Description: `Authentication type. auto - Use PAP, MSCHAP, and CHAP (in that order). ascii - ASCII. pap - PAP. chap - CHAP. mschap - MSCHAP. Valid values: ` + "`" + `auto` + "`" + `, ` + "`" + `ascii` + "`" + `, ` + "`" + `pap` + "`" + `, ` + "`" + `chap` + "`" + `, ` + "`" + `mschap` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authorization",
					Description: `Enable/disable TACACS+ authorization. disable - Disable TACACS+ authorization. enable - Enable TACACS+ authorization (service = fortigate). Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `<password_str> key to access server.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `TACACS+ server entry name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Port number of TACACS+ server.`,
				},
				resource.Attribute{
					Name:        "secondary_key",
					Description: `<password_str> key to access secondary server.`,
				},
				resource.Attribute{
					Name:        "secondary_server",
					Description: `{<name_str|ip_str>} secondary server domain name or IP.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `{<name_str|ip_str>} server domain name or IP.`,
				},
				resource.Attribute{
					Name:        "tertiary_key",
					Description: `<password_str> key to access tertiary server.`,
				},
				resource.Attribute{
					Name:        "tertiary_server",
					Description: `{<name_str|ip_str>} tertiary server domain name or IP. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminTacacs can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_tacacs.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AdminTacacs can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_tacacs.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_admin_user",
			Category:         "System Admin",
			ShortDescription: `Admin user.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"admin",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fazadom",
					Description: `Adom. The structure of ` + "`" + `fazadom` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "adom_access",
					Description: `set all/specify/exclude adom access mode. all - All ADOMs access. specify - Specify ADOMs access. exclude - Exclude ADOMs access. Valid values: ` + "`" + `all` + "`" + `, ` + "`" + `specify` + "`" + `, ` + "`" + `exclude` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom_exclude",
					Description: `Adom-Exclude. The structure of ` + "`" + `adom_exclude` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "avatar",
					Description: `Image file for avatar (maximum 4K base64 encoded).`,
				},
				resource.Attribute{
					Name:        "ca",
					Description: `PKI user certificate CA (CA name in local).`,
				},
				resource.Attribute{
					Name:        "change_password",
					Description: `Enable/disable restricted user to change self password. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dashboard",
					Description: `Dashboard. The structure of ` + "`" + `dashboard` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "dashboard_tabs",
					Description: `Dashboard-Tabs. The structure of ` + "`" + `dashboard_tabs` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "dev_group",
					Description: `device group.`,
				},
				resource.Attribute{
					Name:        "email_address",
					Description: `Email address.`,
				},
				resource.Attribute{
					Name:        "ext_auth_accprofile_override",
					Description: `Allow to use the access profile provided by the remote authentication server. disable - Disable access profile override. enable - Enable access profile override. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ext_auth_adom_override",
					Description: `Allow to use the ADOM provided by the remote authentication server. disable - Disable ADOM override. enable - Enable ADOM override. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ext_auth_group_match",
					Description: `Only administrators belonging to this group can login.`,
				},
				resource.Attribute{
					Name:        "first_name",
					Description: `First name.`,
				},
				resource.Attribute{
					Name:        "force_password_change",
					Description: `Enable/disable force password change on next login. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `Group name.`,
				},
				resource.Attribute{
					Name:        "hidden",
					Description: `Hidden administrator.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost1",
					Description: `Admin user trusted host IPv6, default ::/0 for all.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost10",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost2",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost3",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost4",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost5",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost6",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost7",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost8",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "ipv6_trusthost9",
					Description: `Admin user trusted host IPv6, default ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff/128 for none.`,
				},
				resource.Attribute{
					Name:        "last_name",
					Description: `Last name.`,
				},
				resource.Attribute{
					Name:        "ldap_server",
					Description: `LDAP server name.`,
				},
				resource.Attribute{
					Name:        "login_max",
					Description: `Max login session for this user.`,
				},
				resource.Attribute{
					Name:        "meta_data",
					Description: `Meta-Data. The structure of ` + "`" + `meta_data` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "mobile_number",
					Description: `Mobile number.`,
				},
				resource.Attribute{
					Name:        "pager_number",
					Description: `Pager number.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password.`,
				},
				resource.Attribute{
					Name:        "password_expire",
					Description: `Password expire time in GMT.`,
				},
				resource.Attribute{
					Name:        "phone_number",
					Description: `Phone number.`,
				},
				resource.Attribute{
					Name:        "policy_package",
					Description: `Policy-Package. The structure of ` + "`" + `policy_package` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "profileid",
					Description: `Profile ID.`,
				},
				resource.Attribute{
					Name:        "radius_server",
					Description: `RADIUS server name.`,
				},
				resource.Attribute{
					Name:        "rpc_permit",
					Description: `set none/read/read-write rpc-permission. read-write - Read-write permission. none - No permission. read - Read-only permission. Valid values: ` + "`" + `read-write` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `read` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key1",
					Description: `SSH public key 1.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key2",
					Description: `SSH public key 2.`,
				},
				resource.Attribute{
					Name:        "ssh_public_key3",
					Description: `SSH public key 3.`,
				},
				resource.Attribute{
					Name:        "subject",
					Description: `PKI user certificate name constraints.`,
				},
				resource.Attribute{
					Name:        "tacacs_plus_server",
					Description: `TACACS+ server name.`,
				},
				resource.Attribute{
					Name:        "th_from_profile",
					Description: `Internal use only: trusthostX from-profile flag`,
				},
				resource.Attribute{
					Name:        "th6_from_profile",
					Description: `Internal use only: ipv6_trusthostX from-profile flag`,
				},
				resource.Attribute{
					Name:        "trusthost1",
					Description: `Admin user trusted host IP, default 0.0.0.0 0.0.0.0 for all.`,
				},
				resource.Attribute{
					Name:        "trusthost10",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost2",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost3",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost4",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost5",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost6",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost7",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost8",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "trusthost9",
					Description: `Admin user trusted host IP, default 255.255.255.255 255.255.255.255 for none.`,
				},
				resource.Attribute{
					Name:        "two_factor_auth",
					Description: `Enable 2-factor authentication (certificate + password). disable - Disable 2-factor authentication. enable - Enable 2-factor authentication. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_global_theme",
					Description: `Enable/disble global theme for administration GUI. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_theme",
					Description: `Color scheme to use for the admin user GUI. blue - Blueberry green - Kiwi red - Cherry melongene - Plum spring - Spring summer - Summer autumn - Autumn winter - Winter circuit-board - Circuit Board calla-lily - Calla Lily binary-tunnel - Binary Tunnel mars - Mars blue-sea - Blue Sea technology - Technology landscape - Landscape twilight - Twilight canyon - Canyon northern-light - Northern Light astronomy - Astronomy fish - Fish penguin - Penguin mountain - Mountain panda - Panda parrot - Parrot cave - Cave zebra - Zebra contrast-dark - High Contrast Dark Valid values: ` + "`" + `blue` + "`" + `, ` + "`" + `green` + "`" + `, ` + "`" + `red` + "`" + `, ` + "`" + `melongene` + "`" + `, ` + "`" + `spring` + "`" + `, ` + "`" + `summer` + "`" + `, ` + "`" + `autumn` + "`" + `, ` + "`" + `winter` + "`" + `, ` + "`" + `circuit-board` + "`" + `, ` + "`" + `calla-lily` + "`" + `, ` + "`" + `binary-tunnel` + "`" + `, ` + "`" + `mars` + "`" + `, ` + "`" + `blue-sea` + "`" + `, ` + "`" + `technology` + "`" + `, ` + "`" + `landscape` + "`" + `, ` + "`" + `twilight` + "`" + `, ` + "`" + `canyon` + "`" + `, ` + "`" + `northern-light` + "`" + `, ` + "`" + `astronomy` + "`" + `, ` + "`" + `fish` + "`" + `, ` + "`" + `penguin` + "`" + `, ` + "`" + `mountain` + "`" + `, ` + "`" + `panda` + "`" + `, ` + "`" + `parrot` + "`" + `, ` + "`" + `cave` + "`" + `, ` + "`" + `zebra` + "`" + `, ` + "`" + `contrast-dark` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_type",
					Description: `User type. local - Local user. radius - RADIUS user. ldap - LDAP user. tacacs-plus - TACACS+ user. pki-auth - PKI user. group - Group user. sso - SSO user. Valid values: ` + "`" + `local` + "`" + `, ` + "`" + `radius` + "`" + `, ` + "`" + `ldap` + "`" + `, ` + "`" + `tacacs-plus` + "`" + `, ` + "`" + `pki-auth` + "`" + `, ` + "`" + `group` + "`" + `, ` + "`" + `sso` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "userid",
					Description: `User name.`,
				},
				resource.Attribute{
					Name:        "wildcard",
					Description: `Enable/disable wildcard remote authentication. disable - Disable username wildcard. enable - Enable username wildcard. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `fazadom` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "adom_name",
					Description: `Admin domain names. The ` + "`" + `adom_exclude` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "adom_name",
					Description: `Admin domain names. The ` + "`" + `dashboard` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "column",
					Description: `Widget's column ID.`,
				},
				resource.Attribute{
					Name:        "diskio_content_type",
					Description: `Disk I/O Monitor widget's chart type. util - bandwidth utilization. iops - the number of I/O requests. blks - the amount of data of I/O requests. Valid values: ` + "`" + `util` + "`" + `, ` + "`" + `iops` + "`" + `, ` + "`" + `blks` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskio_period",
					Description: `Disk I/O Monitor widget's data period. 1hour - 1 hour. 8hour - 8 hour. 24hour - 24 hour. Valid values: ` + "`" + `1hour` + "`" + `, ` + "`" + `8hour` + "`" + `, ` + "`" + `24hour` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_rate_period",
					Description: `Log receive monitor widget's data period. 2min - 2 minutes. 1hour - 1 hour. 6hours - 6 hours. Valid values: ` + "`" + `2min ` + "`" + `, ` + "`" + `1hour` + "`" + `, ` + "`" + `6hours` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_rate_topn",
					Description: `Log receive monitor widget's number of top items to display. 1 - Top 1. 2 - Top 2. 3 - Top 3. 4 - Top 4. 5 - Top 5. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_rate_type",
					Description: `Log receive monitor widget's statistics breakdown options. log - Show log rates for each log type. device - Show log rates for each device. Valid values: ` + "`" + `log` + "`" + `, ` + "`" + `device` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "moduleid",
					Description: `Widget ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Widget name.`,
				},
				resource.Attribute{
					Name:        "num_entries",
					Description: `Number of entries.`,
				},
				resource.Attribute{
					Name:        "refresh_interval",
					Description: `Widget's refresh interval.`,
				},
				resource.Attribute{
					Name:        "res_cpu_display",
					Description: `Widget's CPU display type. average - Average usage of CPU. each - Each usage of CPU. Valid values: ` + "`" + `average ` + "`" + `, ` + "`" + `each` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "res_period",
					Description: `Widget's data period. 10min - Last 10 minutes. hour - Last hour. day - Last day. Valid values: ` + "`" + `10min ` + "`" + `, ` + "`" + `hour` + "`" + `, ` + "`" + `day` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "res_view_type",
					Description: `Widget's data view type. real-time - Real-time view. history - History view. Valid values: ` + "`" + `real-time ` + "`" + `, ` + "`" + `history` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Widget's opened/closed state. close - Widget closed. open - Widget opened. Valid values: ` + "`" + `close` + "`" + `, ` + "`" + `open` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tabid",
					Description: `ID of tab where widget is displayed.`,
				},
				resource.Attribute{
					Name:        "time_period",
					Description: `Log Database Monitor widget's data period. 1hour - 1 hour. 8hour - 8 hour. 24hour - 24 hour. Valid values: ` + "`" + `1hour` + "`" + `, ` + "`" + `8hour` + "`" + `, ` + "`" + `24hour` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "widget_type",
					Description: `Widget type. top-lograte - Log Receive Monitor. sysres - System resources. sysinfo - System Information. licinfo - License Information. jsconsole - CLI Console. sysop - Unit Operation. alert - Alert Message Console. statistics - Statistics. rpteng - Report Engine. raid - Disk Monitor. logrecv - Logs/Data Received. devsummary - Device Summary. logdb-perf - Log Database Performance Monitor. logdb-lag - Log Database Lag Time. disk-io - Disk I/O. log-rcvd-fwd - Log receive and forwarding Monitor. Valid values: ` + "`" + `top-lograte` + "`" + `, ` + "`" + `sysres` + "`" + `, ` + "`" + `sysinfo` + "`" + `, ` + "`" + `licinfo` + "`" + `, ` + "`" + `jsconsole` + "`" + `, ` + "`" + `sysop` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `statistics` + "`" + `, ` + "`" + `rpteng` + "`" + `, ` + "`" + `raid` + "`" + `, ` + "`" + `logrecv` + "`" + `, ` + "`" + `devsummary` + "`" + `, ` + "`" + `logdb-perf` + "`" + `, ` + "`" + `logdb-lag` + "`" + `, ` + "`" + `disk-io` + "`" + `, ` + "`" + `log-rcvd-fwd` + "`" + `. The ` + "`" + `dashboard_tabs` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Tab name.`,
				},
				resource.Attribute{
					Name:        "tabid",
					Description: `Tab ID. The ` + "`" + `meta_data` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "fieldlength",
					Description: `Field length.`,
				},
				resource.Attribute{
					Name:        "fieldname",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "fieldvalue",
					Description: `Field value.`,
				},
				resource.Attribute{
					Name:        "importance",
					Description: `Importance. optional - This field is optional. required - This field is required. Valid values: ` + "`" + `optional` + "`" + `, ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status. disabled - This field is disabled. enabled - This field is enabled. Valid values: ` + "`" + `disabled` + "`" + `, ` + "`" + `enabled` + "`" + `. The ` + "`" + `policy_package` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "policy_package_name",
					Description: `Policy package names. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{userid}}. ## Import System AdminUser can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_user.labelname {{userid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{userid}}. ## Import System AdminUser can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_admin_user.labelname {{userid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_alertconsole",
			Category:         "System Alert",
			ShortDescription: `Alert console.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"alert",
				"alertconsole",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "period",
					Description: `Alert console keeps alerts for this period. 1 - 1 day. 2 - 2 days. 3 - 3 days. 4 - 4 days. 5 - 5 days. 6 - 6 days. 7 - 7 days. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `2` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `4` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_level",
					Description: `Alert console keeps alerts of this and higher severity. debug - Debug level. information - Information level. notify - Notify level. warning - Warning level. error - Error level. critical - Critical level. alert - Alert level. emergency - Emergency level. Valid values: ` + "`" + `debug` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `notify` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `emergency` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AlertConsole can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_alertconsole.labelname SystemAlertConsole $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AlertConsole can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_alertconsole.labelname SystemAlertConsole $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_alertemail",
			Category:         "System Alert",
			ShortDescription: `Configure alertemail.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"alert",
				"alertemail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable/disable authentication. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fromaddress",
					Description: `SMTP from address.`,
				},
				resource.Attribute{
					Name:        "fromname",
					Description: `SMTP from user.`,
				},
				resource.Attribute{
					Name:        "smtppassword",
					Description: `SMTP server password.`,
				},
				resource.Attribute{
					Name:        "smtpport",
					Description: `SMTP server port.`,
				},
				resource.Attribute{
					Name:        "smtpserver",
					Description: `SMTP server address.`,
				},
				resource.Attribute{
					Name:        "smtpuser",
					Description: `SMTP server user. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Alertemail can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_alertemail.labelname SystemAlertemail $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Alertemail can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_alertemail.labelname SystemAlertemail $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_alertevent",
			Category:         "System Alert",
			ShortDescription: `Alert events.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"alert",
				"alertevent",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alert_destination",
					Description: `Alert-Destination. The structure of ` + "`" + `alert_destination` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "enable_generic_text",
					Description: `Enable/disable generic text match. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_severity_filter",
					Description: `Enable/disable alert severity filter. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event_time_period",
					Description: `Time period (hours). 0.5 - 30 minutes. 1 - 1 hour. 3 - 3 hours. 6 - 6 hours. 12 - 12 hours. 24 - 1 day. 72 - 3 days. 168 - 1 week. Valid values: ` + "`" + `0.5` + "`" + `, ` + "`" + `1` + "`" + `, ` + "`" + `3` + "`" + `, ` + "`" + `6` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `72` + "`" + `, ` + "`" + `168` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "generic_text",
					Description: `Text that must be contained in a log to trigger alert.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Alert name.`,
				},
				resource.Attribute{
					Name:        "num_events",
					Description: `Minimum number of events required within time period. 1 - 1 event. 5 - 5 events. 10 - 10 events. 50 - 50 events. 100 - 100 events. Valid values: ` + "`" + `1` + "`" + `, ` + "`" + `5` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_filter",
					Description: `Required log severity to trigger alert. high - High level alert. medium-high - Medium-high level alert. medium - Medium level alert. medium-low - Medium-low level alert. low - Low level alert. Valid values: ` + "`" + `high` + "`" + `, ` + "`" + `medium-high` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `medium-low` + "`" + `, ` + "`" + `low` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_level_comp",
					Description: `Log severity threshold comparison criterion. Valid values: ` + "`" + `>=` + "`" + `, ` + "`" + `=` + "`" + `, ` + "`" + `<=` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity_level_logs",
					Description: `Log severity threshold level. no-check - Do not check severity level for this log type. information - Information level. notify - Notify level. warning - Warning level. error - Error level. critical - Critical level. alert - Alert level. emergency - Emergency level. Valid values: ` + "`" + `no-check` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `notify` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `emergency` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `alert_destination` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "from",
					Description: `Sender email address to use in alert emails.`,
				},
				resource.Attribute{
					Name:        "smtp_name",
					Description: `SMTP server name.`,
				},
				resource.Attribute{
					Name:        "snmp_name",
					Description: `SNMP trap name.`,
				},
				resource.Attribute{
					Name:        "syslog_name",
					Description: `Syslog server name.`,
				},
				resource.Attribute{
					Name:        "to",
					Description: `Recipient email address to use in alert emails.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Destination type. mail - Send email alert. snmp - Send SNMP trap. syslog - Send syslog message. Valid values: ` + "`" + `mail` + "`" + `, ` + "`" + `snmp` + "`" + `, ` + "`" + `syslog` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AlertEvent can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_alertevent.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System AlertEvent can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_alertevent.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_autodelete",
			Category:         "System AutoDelete",
			ShortDescription: `Automatic deletion policy for logs, reports, archived, and quarantined files.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"autodelete",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dlp_files_auto_deletion",
					Description: `Dlp-Files-Auto-Deletion. The structure of ` + "`" + `dlp_files_auto_deletion` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "log_auto_deletion",
					Description: `Log-Auto-Deletion. The structure of ` + "`" + `log_auto_deletion` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "quarantine_files_auto_deletion",
					Description: `Quarantine-Files-Auto-Deletion. The structure of ` + "`" + `quarantine_files_auto_deletion` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "report_auto_deletion",
					Description: `Report-Auto-Deletion. The structure of ` + "`" + `report_auto_deletion` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status_fake",
					Description: `Fake value for the menu to work. The ` + "`" + `dlp_files_auto_deletion` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. The ` + "`" + `log_auto_deletion` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. The ` + "`" + `quarantine_files_auto_deletion` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. The ` + "`" + `report_auto_deletion` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDelete can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete.labelname SystemAutoDelete $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDelete can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete.labelname SystemAutoDelete $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_autodelete_dlpfilesautodeletion",
			Category:         "System AutoDelete",
			ShortDescription: `Automatic deletion policy for DLP archives.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"autodelete",
				"dlpfilesautodeletion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteDlpFilesAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_dlpfilesautodeletion.labelname SystemAutoDeleteDlpFilesAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteDlpFilesAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_dlpfilesautodeletion.labelname SystemAutoDeleteDlpFilesAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_autodelete_logautodeletion",
			Category:         "System AutoDelete",
			ShortDescription: `Automatic deletion policy for device logs.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"autodelete",
				"logautodeletion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteLogAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_logautodeletion.labelname SystemAutoDeleteLogAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteLogAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_logautodeletion.labelname SystemAutoDeleteLogAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_autodelete_quarantinefilesautodeletion",
			Category:         "System AutoDelete",
			ShortDescription: `Automatic deletion policy for quarantined files.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"autodelete",
				"quarantinefilesautodeletion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteQuarantineFilesAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_quarantinefilesautodeletion.labelname SystemAutoDeleteQuarantineFilesAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteQuarantineFilesAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_quarantinefilesautodeletion.labelname SystemAutoDeleteQuarantineFilesAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_autodelete_reportautodeletion",
			Category:         "System AutoDelete",
			ShortDescription: `Automatic deletion policy for reports.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"autodelete",
				"reportautodeletion",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "retention",
					Description: `Automatic deletion in days, weeks, or months. days - Auto-delete data older than <value> days. weeks - Auto-delete data older than <value> weeks. months - Auto-delete data older than <value> months. Valid values: ` + "`" + `days` + "`" + `, ` + "`" + `weeks` + "`" + `, ` + "`" + `months` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "runat",
					Description: `Automatic deletion run at (0 - 23) o'clock.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable automatic deletion. disable - Disable automatic deletion. enable - Enable automatic deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Automatic deletion in x days, weeks, or months. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteReportAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_reportautodeletion.labelname SystemAutoDeleteReportAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System AutoDeleteReportAutoDeletion can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_autodelete_reportautodeletion.labelname SystemAutoDeleteReportAutoDeletion $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_backup_allsettings",
			Category:         "System Others",
			ShortDescription: `Scheduled backup settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"backup",
				"allsettings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `SSH certificate for authentication.`,
				},
				resource.Attribute{
					Name:        "crptpasswd",
					Description: `Optional password to protect backup content.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Directory in which file will be stored on backup server.`,
				},
				resource.Attribute{
					Name:        "passwd",
					Description: `Backup server login user password.`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Protocol used to backup. sftp - SFTP. ftp - FTP. scp - SCP. Valid values: ` + "`" + `sftp` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Backup server name/IP and port.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable schedule backup. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Time to backup.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Backup server login user.`,
				},
				resource.Attribute{
					Name:        "week_days",
					Description: `Week days to backup. monday - Monday. tuesday - Tuesday. wednesday - Wednesday. thursday - Thursday. friday - Friday. saturday - Saturday. sunday - Sunday. Valid values: ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `, ` + "`" + `sunday` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System BackupAllSettings can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_backup_allsettings.labelname SystemBackupAllSettings $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System BackupAllSettings can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_backup_allsettings.labelname SystemBackupAllSettings $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_centralmanagement",
			Category:         "No Category",
			ShortDescription: `Central management configuration.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"centralmanagement",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow_monitor",
					Description: `Enable/disable remote monitor of device. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "authorized_manager_only",
					Description: `Enable/disable restricted to authorized manager only. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `SSL communication encryption algorithms. default - SSL communication with high and medium encryption algorithms low - SSL communication with low encryption algorithms high - SSL communication with high encryption algorithms Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `high` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmg",
					Description: `Address of Fortimanager (IP or FQDN).`,
				},
				resource.Attribute{
					Name:        "mgmtid",
					Description: `Management ID.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Type of management server. fortimanager - FortiManager. Valid values: ` + "`" + `fortimanager` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System CentralManagement can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_centralmanagement.labelname SystemCentralManagement $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System CentralManagement can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_centralmanagement.labelname SystemCentralManagement $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_certificate_ca",
			Category:         "System Certificate",
			ShortDescription: `CA certificate.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"certificate",
				"ca",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ca",
					Description: `CA certificate.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `CA certificate comment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateCa can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_ca.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateCa can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_ca.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_certificate_crl",
			Category:         "System Certificate",
			ShortDescription: `Certificate Revocation List.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"certificate",
				"crl",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "comment",
					Description: `Comment of this Certificate Revocation List.`,
				},
				resource.Attribute{
					Name:        "crl",
					Description: `Certificate Revocation List.`,
				},
				resource.Attribute{
					Name:        "http_url",
					Description: `HTTP server URL for CRL auto-update`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "update_interval",
					Description: `CRL auto-update interval (in minutes) ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateCrl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_crl.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateCrl can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_crl.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_certificate_local",
			Category:         "System Certificate",
			ShortDescription: `Local keys and certificates.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"certificate",
				"local",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `Certificate.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Local certificate comment.`,
				},
				resource.Attribute{
					Name:        "csr",
					Description: `Certificate Signing Request.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of local certificate.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Local certificate password.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `Local certificate private-key. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateLocal can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_local.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateLocal can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_local.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_certificate_oftp",
			Category:         "System Certificate",
			ShortDescription: `OFTP certificates and keys.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"certificate",
				"oftp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `PEM format certificate.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `OFTP certificate comment.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `Choose from a local certificates.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Mode of certificates used by oftpd. default - Default mode. custom - Use custom certificate. local - Use a local certificate. Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `custom` + "`" + `, ` + "`" + `local` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for encrypted 'private-key', unset for non-encrypted.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `PEM format private key. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System CertificateOftp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_oftp.labelname SystemCertificateOftp $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System CertificateOftp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_oftp.labelname SystemCertificateOftp $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_certificate_remote",
			Category:         "System Certificate",
			ShortDescription: `Remote certificate.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"certificate",
				"remote",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `Remote certificate.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `Remote certificate comment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateRemote can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_remote.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateRemote can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_remote.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_certificate_ssh",
			Category:         "System Certificate",
			ShortDescription: `SSH certificates and keys.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"certificate",
				"ssh",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "certificate",
					Description: `SSH certificate.`,
				},
				resource.Attribute{
					Name:        "comment",
					Description: `SSH certificate comment.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of SSH certificate.`,
				},
				resource.Attribute{
					Name:        "private_key",
					Description: `SSH private-key ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateSsh can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_ssh.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System CertificateSsh can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_certificate_ssh.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_connector",
			Category:         "System Others",
			ShortDescription: `Configure connector.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"connector",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "conn_refresh_interval",
					Description: `connector refresh interval (60 - 1800 seconds).`,
				},
				resource.Attribute{
					Name:        "fsso_refresh_interval",
					Description: `FSSO refresh interval (60 - 600 seconds).`,
				},
				resource.Attribute{
					Name:        "fsso_sess_timeout",
					Description: `FSSO session timeout (60 - 600 seconds).`,
				},
				resource.Attribute{
					Name:        "px_refresh_interval",
					Description: `pxGrid refresh interval (60 - 1800 seconds).`,
				},
				resource.Attribute{
					Name:        "px_svr_timeout",
					Description: `pxGrid server timeout (30 - 600 seconds). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Connector can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_connector.labelname SystemConnector $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Connector can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_connector.labelname SystemConnector $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_dns",
			Category:         "System Global",
			ShortDescription: `DNS configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"global",
				"dns",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip6_primary",
					Description: `IPv6 primary DNS IP.`,
				},
				resource.Attribute{
					Name:        "ip6_secondary",
					Description: `IPv6 secondary DNS IP.`,
				},
				resource.Attribute{
					Name:        "primary",
					Description: `Primary DNS IP.`,
				},
				resource.Attribute{
					Name:        "secondary",
					Description: `Secondary DNS IP. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Dns can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_dns.labelname SystemDns $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Dns can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_dns.labelname SystemDns $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_docker",
			Category:         "System Others",
			ShortDescription: `Docker host.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"docker",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cpu",
					Description: `Cpu.`,
				},
				resource.Attribute{
					Name:        "default_address_pool_base",
					Description: `Set default-address-pool CIDR.`,
				},
				resource.Attribute{
					Name:        "default_address_pool_size",
					Description: `Set default-address-pool size.`,
				},
				resource.Attribute{
					Name:        "docker_user_login_max",
					Description: `Max login session for docker users.`,
				},
				resource.Attribute{
					Name:        "fortiaiops",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiauthenticator",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiportal",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortisigconverter",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortisoar",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiwlm",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fsmcollector",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mem",
					Description: `Max % RAM usage.`,
				},
				resource.Attribute{
					Name:        "sdwancontroller",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable and set registry. disable - Disable docker host service. enable - Enable production registry. qa - Enable QA test registry. dev - Enable QA test registry (without signature). Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `, ` + "`" + `qa` + "`" + `, ` + "`" + `dev` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "universalconnector",
					Description: `Enable/disable container. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Docker can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_docker.labelname SystemDocker $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Docker can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_docker.labelname SystemDocker $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_fips",
			Category:         "System Others",
			ShortDescription: `Settings for FIPS-CC mode.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"fips",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "entropy_token",
					Description: `Enable/disable entropy token when switching to FIPS mode. enable - Enable entropy token. disable - Disable entropy token. dynamic - Dynamically detect entropy token during bootup. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `, ` + "`" + `dynamic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "re_seed_interval",
					Description: `Kernel FIPS-compliant PRNG re-seed interval (0 to 1440 minutes)`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable FIPS-CC mode. disable - Disable FIPS-CC mode. enable - Enable FIPS-CC mode. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Fips can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_fips.labelname SystemFips $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Fips can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_fips.labelname SystemFips $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_fortiview_autocache",
			Category:         "System FortiView",
			ShortDescription: `FortiView auto-cache settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"fortiview",
				"autocache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aggressive_fortiview",
					Description: `Enable/disable auto-cache on fortiview aggressively. disable - Disable the aggressive fortiview auto-cache. enable - Enable the aggressive fortiview auto-cache. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "interval",
					Description: `The time interval in hours for fortiview auto-cache.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable fortiview auto-cache. disable - Disable the fortiview auto-cache. enable - Enable the fortiview auto-cache. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System FortiviewAutoCache can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_fortiview_autocache.labelname SystemFortiviewAutoCache $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System FortiviewAutoCache can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_fortiview_autocache.labelname SystemFortiviewAutoCache $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_fortiview_setting",
			Category:         "System FortiView",
			ShortDescription: `FortiView settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"fortiview",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "data_source",
					Description: `Data soure of the fortiview query. auto - Data from hcache, and from logs in a flexible way. cache-only - Data from hcache only. log-and-cache - Data from logs and hcache. Valid values: ` + "`" + `auto` + "`" + `, ` + "`" + `cache-only` + "`" + `, ` + "`" + `log-and-cache` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "not_scanned_apps",
					Description: `Include/Exclude 'Not.Scanned' applications in FortiView. Set as 'exclude' if you want to filter out never scanned applications. exclude - Exclude 'Not.Scanned' applications in FortiView. include - Include 'Not.Scanned' applications in FortiView. Valid values: ` + "`" + `exclude` + "`" + `, ` + "`" + `include` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "resolve_ip",
					Description: `Enable or disable resolving IP address to hostname in FortiView. disable - Disable resolving IP address to hostname. enable - Enable resolving IP address to hostname. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System FortiviewSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_fortiview_setting.labelname SystemFortiviewSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System FortiviewSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_fortiview_setting.labelname SystemFortiviewSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_global",
			Category:         "System Global",
			ShortDescription: `Global range attributes.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"global",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_lockout_duration",
					Description: `Lockout duration(sec) for administration.`,
				},
				resource.Attribute{
					Name:        "admin_lockout_threshold",
					Description: `Lockout threshold for administration.`,
				},
				resource.Attribute{
					Name:        "adom_mode",
					Description: `ADOM mode. normal - Normal ADOM mode. advanced - Advanced ADOM mode. Valid values: ` + "`" + `normal` + "`" + `, ` + "`" + `advanced` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom_select",
					Description: `Enable/disable select ADOM after login. disable - Disable select ADOM after login. enable - Enable select ADOM after login. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom_status",
					Description: `ADOM status. disable - Disable ADOM mode. enable - Enable ADOM mode. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_compression",
					Description: `Compression level. none - No compression. low - Low compression (fastest). normal - Normal compression. high - Best compression (slowest). Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "backup_to_subfolders",
					Description: `Enable/disable creation of subfolders on server for backup storage. disable - Disable creation of subfolders on server for backup storage. enable - Enable creation of subfolders on server for backup storage. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clone_name_option",
					Description: `set the clone object names option. default - Add a prefix of 'Clone of' to the clone name. keep - Keep the original name for user to edit. Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `keep` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "clt_cert_req",
					Description: `Require client certificate for GUI login. disable - Disable setting. enable - Require client certificate for GUI login. optional - Optional client certificate for GUI login. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `, ` + "`" + `optional` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "console_output",
					Description: `Console output mode. standard - Standard output. more - More page output. Valid values: ` + "`" + `standard` + "`" + `, ` + "`" + `more` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "country_flag",
					Description: `Country flag Status. disable - Disable country flag icon beside ip address. enable - Enable country flag icon beside ip address. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "create_revision",
					Description: `Enable/disable create revision by default. disable - Disable create revision by default. enable - Enable create revision by default. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "daylightsavetime",
					Description: `Enable/disable daylight saving time. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_logview_auto_completion",
					Description: `Enable/disable log view filter auto-completion. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "default_search_mode",
					Description: `Set the default search mode of log view. filter-based - Filter based search mode. advanced - Advanced search mode. Valid values: ` + "`" + `filter-based` + "`" + `, ` + "`" + `advanced` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "detect_unregistered_log_device",
					Description: `Detect unregistered logging device from log message. disable - Disable attribute function. enable - Enable attribute function. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_view_mode",
					Description: `Set devices/groups view mode. regular - Regular view mode. tree - Tree view mode. Valid values: ` + "`" + `regular` + "`" + `, ` + "`" + `tree` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dh_params",
					Description: `Minimum size of Diffie-Hellman prime for SSH/HTTPS (bits). 1024 - 1024 bits. 1536 - 1536 bits. 2048 - 2048 bits. 3072 - 3072 bits. 4096 - 4096 bits. 6144 - 6144 bits. 8192 - 8192 bits. Valid values: ` + "`" + `1024` + "`" + `, ` + "`" + `1536` + "`" + `, ` + "`" + `2048` + "`" + `, ` + "`" + `3072` + "`" + `, ` + "`" + `4096` + "`" + `, ` + "`" + `6144` + "`" + `, ` + "`" + `8192` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disable_module",
					Description: `Disable module list. fortiview-noc - FortiView/NOC-SOC module. fortirecorder - FortiRecorder module. siem - SIEM module. soc - SOC module. ai - AI module. Valid values: ` + "`" + `fortiview-noc` + "`" + `, ` + "`" + `fortirecorder` + "`" + `, ` + "`" + `siem` + "`" + `, ` + "`" + `soc` + "`" + `, ` + "`" + `ai` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enc_algorithm",
					Description: `SSL communication encryption algorithms. low - SSL communication using all available encryption algorithms. medium - SSL communication using high and medium encryption algorithms. high - SSL communication using high encryption algorithms. Valid values: ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm_ca_cert",
					Description: `set the extra fgfm CA certificates.`,
				},
				resource.Attribute{
					Name:        "fgfm_local_cert",
					Description: `set the fgfm local certificate.`,
				},
				resource.Attribute{
					Name:        "fgfm_ssl_protocol",
					Description: `set the lowest SSL protocols for fgfmsd. sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version (default). tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: ` + "`" + `sslv3` + "`" + `, ` + "`" + `tlsv1.0` + "`" + `, ` + "`" + `tlsv1.1` + "`" + `, ` + "`" + `tlsv1.2` + "`" + `, ` + "`" + `tlsv1.3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_member_auto_grouping",
					Description: `Enable/disable automatically group HA members feature disable - Disable automatically grouping HA members feature. enable - Enable automatically grouping HA members only when group name is unique in your network. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hostname",
					Description: `System hostname.`,
				},
				resource.Attribute{
					Name:        "language",
					Description: `System global language. english - English simch - Simplified Chinese japanese - Japanese korean - Korean spanish - Spanish trach - Traditional Chinese Valid values: ` + "`" + `english` + "`" + `, ` + "`" + `simch` + "`" + `, ` + "`" + `japanese` + "`" + `, ` + "`" + `korean` + "`" + `, ` + "`" + `spanish` + "`" + `, ` + "`" + `trach` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "latitude",
					Description: `fmg location latitude`,
				},
				resource.Attribute{
					Name:        "ldap_cache_timeout",
					Description: `LDAP browser cache timeout (seconds).`,
				},
				resource.Attribute{
					Name:        "ldapconntimeout",
					Description: `LDAP connection timeout (msec).`,
				},
				resource.Attribute{
					Name:        "lock_preempt",
					Description: `Enable/disable ADOM lock override. disable - Disable lock preempt. enable - Enable lock preempt. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_checksum",
					Description: `Record log file hash value, timestamp, and authentication code at transmission or rolling. none - No record log file checksum. md5 - Record log file's MD5 hash value only. md5-auth - Record log file's MD5 hash value and authentication code. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `md5` + "`" + `, ` + "`" + `md5-auth` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_forward_cache_size",
					Description: `Log forwarding disk cache size (GB).`,
				},
				resource.Attribute{
					Name:        "log_mode",
					Description: `Log system operation mode. analyzer - Operation mode is Analyzer collector - Operation mode is Collector Valid values: ` + "`" + `analyzer` + "`" + `, ` + "`" + `collector` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "longitude",
					Description: `fmg location longitude`,
				},
				resource.Attribute{
					Name:        "max_aggregation_tasks",
					Description: `Maximum number of concurrent tasks of a log aggregation session.`,
				},
				resource.Attribute{
					Name:        "max_log_forward",
					Description: `Maximum number of log-forward and aggregation settings.`,
				},
				resource.Attribute{
					Name:        "max_running_reports",
					Description: `Maximum number of reports generating at one time.`,
				},
				resource.Attribute{
					Name:        "multiple_steps_upgrade_in_autolink",
					Description: `Enable/disable multiple steps upgade in autolink process disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "normalized_intf_zone_only",
					Description: `allow normalized interface to be zone only. disable - Disable SSL low-grade encryption. enable - Enable SSL low-grade encryption. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_revision_db_max",
					Description: `Maximum revisions for a single database (10,000-1,000,000 default 100,000).`,
				},
				resource.Attribute{
					Name:        "object_revision_mandatory_note",
					Description: `Enable/disable mandatory note when create revision. disable - Disable object revision. enable - Enable object revision. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "object_revision_object_max",
					Description: `Maximum revisions for a single object (10-1000 default 100).`,
				},
				resource.Attribute{
					Name:        "object_revision_status",
					Description: `Enable/disable create revision when modify objects. disable - Disable object revision. enable - Enable object revision. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "oftp_ssl_protocol",
					Description: `set the lowest SSL protocols for oftpd. sslv3 - set SSLv3 as the lowest version. tlsv1.0 - set TLSv1.0 as the lowest version. tlsv1.1 - set TLSv1.1 as the lowest version. tlsv1.2 - set TLSv1.2 as the lowest version (default). tlsv1.3 - set TLSv1.3 as the lowest version. Valid values: ` + "`" + `sslv3` + "`" + `, ` + "`" + `tlsv1.0` + "`" + `, ` + "`" + `tlsv1.1` + "`" + `, ` + "`" + `tlsv1.2` + "`" + `, ` + "`" + `tlsv1.3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_object_icon",
					Description: `show icons of policy objects. disable - Disable icon of policy objects. enable - Enable icon of policy objects. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "policy_object_in_dual_pane",
					Description: `show policies and objects in dual pane. disable - Disable polices and objects in dual pane. enable - Enable polices and objects in dual pane. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_login_banner",
					Description: `Enable/disable pre-login banner. disable - Disable pre-login banner. enable - Enable pre-login banner. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pre_login_banner_message",
					Description: `Pre-login banner message.`,
				},
				resource.Attribute{
					Name:        "private_data_encryption",
					Description: `Enable/disable private data encryption using an AES 128-bit key. disable - Disable private data encryption using an AES 128-bit key. enable - Enable private data encryption using an AES 128-bit key. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "remoteauthtimeout",
					Description: `Remote authentication (RADIUS/LDAP) timeout (sec).`,
				},
				resource.Attribute{
					Name:        "search_all_adoms",
					Description: `Enable/Disable Search all ADOMs for where-used query. disable - Disable search all ADOMs for where-used queries. enable - Enable search all ADOMs for where-used queries. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_cipher_suites",
					Description: `Ssl-Cipher-Suites. The structure of ` + "`" + `ssl_cipher_suites` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "ssl_low_encryption",
					Description: `SSL low-grade encryption. disable - Disable SSL low-grade encryption. enable - Enable SSL low-grade encryption. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_protocol",
					Description: `SSL protocols. tlsv1.3 - Enable TLSv1.3. tlsv1.2 - Enable TLSv1.2. tlsv1.1 - Enable TLSv1.1. tlsv1.0 - Enable TLSv1.0. sslv3 - Enable SSLv3. Valid values: ` + "`" + `tlsv1.3` + "`" + `, ` + "`" + `tlsv1.2` + "`" + `, ` + "`" + `tlsv1.1` + "`" + `, ` + "`" + `tlsv1.0` + "`" + `, ` + "`" + `sslv3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ssl_static_key_ciphers",
					Description: `Enable/disable SSL static key ciphers. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "task_list_size",
					Description: `Maximum number of completed tasks to keep.`,
				},
				resource.Attribute{
					Name:        "tftp",
					Description: `Enable/disable TFTP in ` + "`" + `exec restore image` + "`" + ` command (disabled by default in FIPS mode) disable - Disable TFTP enable - Enable TFTP Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timezone",
					Description: `Time zone. 00 - (GMT-12:00) Eniwetak, Kwajalein. 01 - (GMT-11:00) Midway Island, Samoa. 02 - (GMT-10:00) Hawaii. 03 - (GMT-9:00) Alaska. 04 - (GMT-8:00) Pacific Time (US & Canada). 05 - (GMT-7:00) Arizona. 06 - (GMT-7:00) Mountain Time (US & Canada). 07 - (GMT-6:00) Central America. 08 - (GMT-6:00) Central Time (US & Canada). 09 - (GMT-6:00) Mexico City. 10 - (GMT-6:00) Saskatchewan. 11 - (GMT-5:00) Bogota, Lima, Quito. 12 - (GMT-5:00) Eastern Time (US & Canada). 13 - (GMT-5:00) Indiana (East). 14 - (GMT-4:00) Atlantic Time (Canada). 15 - (GMT-4:00) La Paz. 16 - (GMT-4:00) Santiago. 17 - (GMT-3:30) Newfoundland. 18 - (GMT-3:00) Brasilia. 19 - (GMT-3:00) Buenos Aires, Georgetown. 20 - (GMT-3:00) Nuuk (Greenland). 21 - (GMT-2:00) Mid-Atlantic (Deprecated). 22 - (GMT-1:00) Azores. 23 - (GMT-1:00) Cape Verde Is. 24 - (GMT) Monrovia. 25 - (GMT) London, Edinburgh. 26 - (GMT+1:00) Amsterdam, Berlin, Bern, Rome, Stockholm, Vienna. 27 - (GMT+1:00) Belgrade, Bratislava, Budapest, Ljubljana, Prague. 28 - (GMT+1:00) Brussels, Copenhagen, Madrid, Paris. 29 - (GMT+1:00) Sarajevo, Skopje, Warsaw, Zagreb. 30 - (GMT+1:00) West Central Africa. 31 - (GMT+2:00) Athens, Sofia, Vilnius. 32 - (GMT+2:00) Bucharest. 33 - (GMT+2:00) Cairo. 34 - (GMT+2:00) Harare, Pretoria. 35 - (GMT+2:00) Helsinki, Riga,Tallinn. 36 - (GMT+2:00) Jerusalem. 37 - (GMT+3:00) Baghdad. 38 - (GMT+3:00) Kuwait, Riyadh. 39 - (GMT+3:00) St.Petersburg, Volgograd. 40 - (GMT+3:00) Nairobi. 41 - (GMT+3:30) Tehran. 42 - (GMT+4:00) Abu Dhabi, Muscat. 43 - (GMT+4:00) Baku. 44 - (GMT+4:30) Kabul. 45 - (GMT+5:00) Ekaterinburg. 46 - (GMT+5:00) Islamabad, Karachi, Tashkent. 47 - (GMT+5:30) Calcutta, Chennai, Mumbai, New Delhi. 48 - (GMT+5:45) Kathmandu. 49 - (GMT+6:00) Almaty, Novosibirsk. 50 - (GMT+6:00) Astana, Dhaka. 51 - (GMT+5:30) Sri Jayawardenepura. 52 - (GMT+6:30) Rangoon. 53 - (GMT+7:00) Bangkok, Hanoi, Jakarta. 54 - (GMT+7:00) Krasnoyarsk. 55 - (GMT+8:00) Beijing, ChongQing, HongKong, Urumqi. 56 - (GMT+8:00) Irkutsk, Ulaanbaatar. 57 - (GMT+8:00) Kuala Lumpur, Singapore. 58 - (GMT+8:00) Perth. 59 - (GMT+8:00) Taipei. 60 - (GMT+9:00) Osaka, Sapporo, Tokyo, Seoul. 61 - (GMT+9:00) Yakutsk. 62 - (GMT+9:30) Adelaide. 63 - (GMT+9:30) Darwin. 64 - (GMT+10:00) Brisbane. 65 - (GMT+10:00) Canberra, Melbourne, Sydney. 66 - (GMT+10:00) Guam, Port Moresby. 67 - (GMT+10:00) Hobart. 68 - (GMT+10:00) Vladivostok. 69 - (GMT+11:00) Magadan. 70 - (GMT+11:00) Solomon Is., New Caledonia. 71 - (GMT+12:00) Auckland, Wellington. 72 - (GMT+12:00) Fiji, Kamchatka, Marshall Is. 73 - (GMT+13:00) Nuku'alofa. 74 - (GMT-4:30) Caracas. 75 - (GMT+1:00) Namibia. 76 - (GMT-5:00) Brazil-Acre. 77 - (GMT-4:00) Brazil-West. 78 - (GMT-3:00) Brazil-East. 79 - (GMT-2:00) Brazil-DeNoronha. 80 - (GMT+14:00) Kiritimati. 81 - (GMT-7:00) Baja California Sur, Chihuahua. 82 - (GMT+12:45) Chatham Islands. 83 - (GMT+3:00) Minsk. 84 - (GMT+13:00) Samoa. 85 - (GMT+3:00) Istanbul. 86 - (GMT-4:00) Paraguay. 87 - (GMT) Casablanca. 88 - (GMT+3:00) Moscow. 89 - (GMT) Greenwich Mean Time. 90 - (GMT) Dublin. 91 - (GMT) Lisbon. Valid values: ` + "`" + `00` + "`" + `, ` + "`" + `01` + "`" + `, ` + "`" + `02` + "`" + `, ` + "`" + `03` + "`" + `, ` + "`" + `04` + "`" + `, ` + "`" + `05` + "`" + `, ` + "`" + `06` + "`" + `, ` + "`" + `07` + "`" + `, ` + "`" + `08` + "`" + `, ` + "`" + `09` + "`" + `, ` + "`" + `10` + "`" + `, ` + "`" + `11` + "`" + `, ` + "`" + `12` + "`" + `, ` + "`" + `13` + "`" + `, ` + "`" + `14` + "`" + `, ` + "`" + `15` + "`" + `, ` + "`" + `16` + "`" + `, ` + "`" + `17` + "`" + `, ` + "`" + `18` + "`" + `, ` + "`" + `19` + "`" + `, ` + "`" + `20` + "`" + `, ` + "`" + `21` + "`" + `, ` + "`" + `22` + "`" + `, ` + "`" + `23` + "`" + `, ` + "`" + `24` + "`" + `, ` + "`" + `25` + "`" + `, ` + "`" + `26` + "`" + `, ` + "`" + `27` + "`" + `, ` + "`" + `28` + "`" + `, ` + "`" + `29` + "`" + `, ` + "`" + `30` + "`" + `, ` + "`" + `31` + "`" + `, ` + "`" + `32` + "`" + `, ` + "`" + `33` + "`" + `, ` + "`" + `34` + "`" + `, ` + "`" + `35` + "`" + `, ` + "`" + `36` + "`" + `, ` + "`" + `37` + "`" + `, ` + "`" + `38` + "`" + `, ` + "`" + `39` + "`" + `, ` + "`" + `40` + "`" + `, ` + "`" + `41` + "`" + `, ` + "`" + `42` + "`" + `, ` + "`" + `43` + "`" + `, ` + "`" + `44` + "`" + `, ` + "`" + `45` + "`" + `, ` + "`" + `46` + "`" + `, ` + "`" + `47` + "`" + `, ` + "`" + `48` + "`" + `, ` + "`" + `49` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `51` + "`" + `, ` + "`" + `52` + "`" + `, ` + "`" + `53` + "`" + `, ` + "`" + `54` + "`" + `, ` + "`" + `55` + "`" + `, ` + "`" + `56` + "`" + `, ` + "`" + `57` + "`" + `, ` + "`" + `58` + "`" + `, ` + "`" + `59` + "`" + `, ` + "`" + `60` + "`" + `, ` + "`" + `61` + "`" + `, ` + "`" + `62` + "`" + `, ` + "`" + `63` + "`" + `, ` + "`" + `64` + "`" + `, ` + "`" + `65` + "`" + `, ` + "`" + `66` + "`" + `, ` + "`" + `67` + "`" + `, ` + "`" + `68` + "`" + `, ` + "`" + `69` + "`" + `, ` + "`" + `70` + "`" + `, ` + "`" + `71` + "`" + `, ` + "`" + `72` + "`" + `, ` + "`" + `73` + "`" + `, ` + "`" + `74` + "`" + `, ` + "`" + `75` + "`" + `, ` + "`" + `76` + "`" + `, ` + "`" + `77` + "`" + `, ` + "`" + `78` + "`" + `, ` + "`" + `79` + "`" + `, ` + "`" + `80` + "`" + `, ` + "`" + `81` + "`" + `, ` + "`" + `82` + "`" + `, ` + "`" + `83` + "`" + `, ` + "`" + `84` + "`" + `, ` + "`" + `85` + "`" + `, ` + "`" + `86` + "`" + `, ` + "`" + `87` + "`" + `, ` + "`" + `88` + "`" + `, ` + "`" + `89` + "`" + `, ` + "`" + `90` + "`" + `, ` + "`" + `91` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tunnel_mtu",
					Description: `Maximum transportation unit(68 - 9000).`,
				},
				resource.Attribute{
					Name:        "usg",
					Description: `Enable/disable Fortiguard server restriction. disable - Contact any Fortiguard server enable - Contact Fortiguard server in USA only Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webservice_proto",
					Description: `Web Service connection support SSL protocols. tlsv1.3 - Web Service connection using TLSv1.3 protocol. tlsv1.2 - Web Service connection using TLSv1.2 protocol. tlsv1.1 - Web Service connection using TLSv1.1 protocol. tlsv1.0 - Web Service connection using TLSv1.0 protocol. sslv3 - Web Service connection using SSLv3 protocol. sslv2 - Web Service connection using SSLv2 protocol. Valid values: ` + "`" + `tlsv1.3` + "`" + `, ` + "`" + `tlsv1.2` + "`" + `, ` + "`" + `tlsv1.1` + "`" + `, ` + "`" + `tlsv1.0` + "`" + `, ` + "`" + `sslv3` + "`" + `, ` + "`" + `sslv2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "workflow_max_sessions",
					Description: `Maximum number of workflow sessions per ADOM (minimum 100).`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `ssl_cipher_suites` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "cipher",
					Description: `Cipher name`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `SSL/TLS cipher suites priority.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `SSL/TLS version the cipher suite can be used with. tls1.2-or-below - TLS 1.2 or below. tls1.3 - TLS 1.3 Valid values: ` + "`" + `tls1.2-or-below` + "`" + `, ` + "`" + `tls1.3` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Global can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_global.labelname SystemGlobal $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Global can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_global.labelname SystemGlobal $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_global_sslciphersuites",
			Category:         "No Category",
			ShortDescription: `Configure preferred SSL/TLS cipher suites`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"global",
				"sslciphersuites",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cipher",
					Description: `Cipher name`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `SSL/TLS cipher suites priority.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `SSL/TLS version the cipher suite can be used with. tls1.2-or-below - TLS 1.2 or below. tls1.3 - TLS 1.3 Valid values: ` + "`" + `tls1.2-or-below` + "`" + `, ` + "`" + `tls1.3` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System GlobalSslCipherSuites can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_global_sslciphersuites.labelname SystemGlobalSslCipherSuites $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System GlobalSslCipherSuites can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_global_sslciphersuites.labelname SystemGlobalSslCipherSuites $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_guiact",
			Category:         "System Others",
			ShortDescription: `System settings through GUI.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"guiact",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "backup_all",
					Description: `Backup all settings.`,
				},
				resource.Attribute{
					Name:        "backup_conf",
					Description: `Backup config file.`,
				},
				resource.Attribute{
					Name:        "eventlog_msg",
					Description: `Write event log.`,
				},
				resource.Attribute{
					Name:        "eventlog_path",
					Description: `Event log path.`,
				},
				resource.Attribute{
					Name:        "reboot",
					Description: `Reboot system.`,
				},
				resource.Attribute{
					Name:        "reset2default",
					Description: `Reset to factory default.`,
				},
				resource.Attribute{
					Name:        "restore_all",
					Description: `Restore all settings.`,
				},
				resource.Attribute{
					Name:        "restore_conf",
					Description: `Restore config file.`,
				},
				resource.Attribute{
					Name:        "time",
					Description: `Time. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Guiact can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_guiact.labelname SystemGuiact $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Guiact can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_guiact.labelname SystemGuiact $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_ha",
			Category:         "System Global",
			ShortDescription: `HA configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"global",
				"ha",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cfg_sync_hb_interval",
					Description: `Config sync heartbeat interval (1 - 80).`,
				},
				resource.Attribute{
					Name:        "group_id",
					Description: `HA group ID (1 - 255).`,
				},
				resource.Attribute{
					Name:        "group_name",
					Description: `HA group name.`,
				},
				resource.Attribute{
					Name:        "hb_interface",
					Description: `Interface for heartbeat.`,
				},
				resource.Attribute{
					Name:        "hb_interval",
					Description: `Heartbeat interval (1 - 20).`,
				},
				resource.Attribute{
					Name:        "healthcheck",
					Description: `Healthchecking options. DB - Check database is running fault-test - temp fault test Valid values: ` + "`" + `DB` + "`" + `, ` + "`" + `fault-test` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "initial_sync",
					Description: `Need to sync data from primary before up as an HA member. false - False. true - True. Valid values: ` + "`" + `false` + "`" + `, ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "initial_sync_threads",
					Description: `Number of threads used for initial sync (1-15).`,
				},
				resource.Attribute{
					Name:        "load_balance",
					Description: `Load balance to secondaries. disable - Disable load-balance to secondaries. round-robin - Round-Robin mode. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `round-robin` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_sync",
					Description: `Sync logs to secondary FortiAnalyzers. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Standalone or HA (a-p) mode standalone - Standalone mode. a-p - Active-Passive mode. Valid values: ` + "`" + `standalone` + "`" + `, ` + "`" + `a-p` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `HA group password.`,
				},
				resource.Attribute{
					Name:        "peer",
					Description: `Peer. The structure of ` + "`" + `peer` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "preferred_role",
					Description: `Preferred role, runtime role may be different. secondary - Prefer secondary mode, FAZ can only become primary after data-sync is done. primary - Prefer primary mode, FAZ can become primary if there's no existing primary. Valid values: ` + "`" + `secondary` + "`" + `, ` + "`" + `primary` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priority",
					Description: `Set the runtime priority between 80 (lowest) - 120 (highest).`,
				},
				resource.Attribute{
					Name:        "private_clusterid",
					Description: `Cluster ID range (1 - 64).`,
				},
				resource.Attribute{
					Name:        "private_file_quota",
					Description: `File quota in MB (2048 - 20480).`,
				},
				resource.Attribute{
					Name:        "private_hb_interval",
					Description: `Heartbeat interval (1 - 255).`,
				},
				resource.Attribute{
					Name:        "private_hb_lost_threshold",
					Description: `Heartbeat lost threshold (1 - 255).`,
				},
				resource.Attribute{
					Name:        "private_mode",
					Description: `Mode. standalone - Standalone. primary - Primary. secondary - Secondary. Valid values: ` + "`" + `standalone` + "`" + `, ` + "`" + `primary` + "`" + `, ` + "`" + `secondary` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "private_password",
					Description: `Group password.`,
				},
				resource.Attribute{
					Name:        "private_peer",
					Description: `Private-Peer. The structure of ` + "`" + `private_peer` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "unicast",
					Description: `Use unicast for HA heartbeat. disable - HA heartbeat through multicast. enable - HA heartbeat through unicast. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vip",
					Description: `Virtual IP address for the HA`,
				},
				resource.Attribute{
					Name:        "vip_interface",
					Description: `Interface for configuring virtual IP address`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `peer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of peer for management and data.`,
				},
				resource.Attribute{
					Name:        "ip_hb",
					Description: `IP address of peer's VIP interface for heartbeat, set if different from ip. (needed only when using unicast)`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number of peer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Peer enabled status. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `private_peer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Id.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of peer.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IP address (V6) of peer.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number of peer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Peer admin status. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Ha can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ha.labelname SystemHa $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Ha can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ha.labelname SystemHa $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_ha_peer",
			Category:         "System Global",
			ShortDescription: `Peers.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"global",
				"ha",
				"peer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fosid",
					Description: `Id.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of peer for management and data.`,
				},
				resource.Attribute{
					Name:        "ip_hb",
					Description: `IP address of peer's VIP interface for heartbeat, set if different from ip. (needed only when using unicast)`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number of peer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Peer enabled status. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System HaPeer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ha_peer.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System HaPeer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ha_peer.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_ha_privatepeer",
			Category:         "No Category",
			ShortDescription: `Peer.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"ha",
				"privatepeer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fosid",
					Description: `Id.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of peer.`,
				},
				resource.Attribute{
					Name:        "ip6",
					Description: `IP address (V6) of peer.`,
				},
				resource.Attribute{
					Name:        "serial_number",
					Description: `Serial number of peer.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Peer admin status. disable - Disable. enable - Enable. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System HaPrivatePeer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ha_privatepeer.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System HaPrivatePeer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ha_privatepeer.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_interface",
			Category:         "System Global",
			ShortDescription: `Interface configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"global",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aggregate",
					Description: `Aggregate interface.`,
				},
				resource.Attribute{
					Name:        "alias",
					Description: `Alias.`,
				},
				resource.Attribute{
					Name:        "allowaccess",
					Description: `Allow management access to interface. ping - PING access. https - HTTPS access. ssh - SSH access. snmp - SNMP access. http - HTTP access. webservice - Web service access. fgfm - FortiManager access. https-logging - Logging over HTTPS access. soc-fabric - SOC Fabric access. Valid values: ` + "`" + `ping` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `snmp` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `webservice` + "`" + `, ` + "`" + `fgfm` + "`" + `, ` + "`" + `https-logging` + "`" + `, ` + "`" + `soc-fabric` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Description.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `IP address of interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Ipv6. The structure of ` + "`" + `ipv6` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "lacp_mode",
					Description: `LACP mode. active - Actively use LACP to negotiate 802.3ad aggregation. Valid values: ` + "`" + `active` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lacp_speed",
					Description: `How often the interface sends LACP messages. slow - Send LACP message every 30 seconds. fast - Send LACP message every second. Valid values: ` + "`" + `slow` + "`" + `, ` + "`" + `fast` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_up_delay",
					Description: `Number of milliseconds to wait before considering a link is up.`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member. The structure of ` + "`" + `member` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "min_links",
					Description: `Minimum number of aggregated ports that must be up.`,
				},
				resource.Attribute{
					Name:        "min_links_down",
					Description: `Action to take when less than the configured minimum number of links are active. operational - Set the aggregate operationally down. administrative - Set the aggregate administratively down. Valid values: ` + "`" + `operational` + "`" + `, ` + "`" + `administrative` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `Maximum transportation unit(68 - 9000).`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Interface name.`,
				},
				resource.Attribute{
					Name:        "speed",
					Description: `Speed. auto - Auto adjust speed. 10full - 10M full-duplex. 10half - 10M half-duplex. 100full - 100M full-duplex. 100half - 100M half-duplex. 1000full - 1000M full-duplex. 10000full - 10000M full-duplex. Valid values: ` + "`" + `auto` + "`" + `, ` + "`" + `10full` + "`" + `, ` + "`" + `10half` + "`" + `, ` + "`" + `100full` + "`" + `, ` + "`" + `100half` + "`" + `, ` + "`" + `1000full` + "`" + `, ` + "`" + `10000full` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Interface status. down - Interface down. up - Interface up. Valid values: ` + "`" + `down` + "`" + `, ` + "`" + `up` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Set type of interface (physical/aggregate). physical - Physical interface. aggregate - Aggregate interface. Valid values: ` + "`" + `physical` + "`" + `, ` + "`" + `aggregate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "autogenerated",
					Description: `Indicates whether the interface is automatically created by FortiAnalyzer, for example, created during the VPN creation process. If it is, set it to "auto", else keep it empty.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Indicates whether the type of the interface is physical. If it is, set it to "physical", else keep it empty.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `ipv6` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "ip6_address",
					Description: `IPv6 address/prefix of interface.`,
				},
				resource.Attribute{
					Name:        "ip6_allowaccess",
					Description: `Allow management access to interface. ping - PING access. https - HTTPS access. ssh - SSH access. snmp - SNMP access. http - HTTP access. webservice - Web service access. fgfm - FortiManager access. https-logging - Logging over HTTPS access. Valid values: ` + "`" + `ping` + "`" + `, ` + "`" + `https` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `snmp` + "`" + `, ` + "`" + `http` + "`" + `, ` + "`" + `webservice` + "`" + `, ` + "`" + `fgfm` + "`" + `, ` + "`" + `https-logging` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ip6_autoconf",
					Description: `Enable/disable address auto config (SLAAC). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `member` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "interface_name",
					Description: `Physical interface name. -> If you want to modify the configuration of the interface which is automatically created by FortiAnalyzer (for example, the interface created during the VPN creation process), please set the ` + "`" + `autogenerated` + "`" + ` argument to "auto". ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System Interface can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_interface.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System Interface can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_interface.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_disk_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for disk logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"disk",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogDiskFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_disk_filter.labelname SystemLocallogDiskFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogDiskFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_disk_filter.labelname SystemLocallogDiskFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_disk_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for local disk logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"disk",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "diskfull",
					Description: `Policy to apply when disk is full. overwrite - Overwrite oldest log when disk is full. nolog - Stop logging when disk is full. Valid values: ` + "`" + `overwrite` + "`" + `, ` + "`" + `nolog` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_disk_full_percentage",
					Description: `Consider log disk as full at this usage percentage.`,
				},
				resource.Attribute{
					Name:        "log_disk_quota",
					Description: `Quota for controlling local log size.`,
				},
				resource.Attribute{
					Name:        "max_log_file_num",
					Description: `Maximum number of log files before rolling.`,
				},
				resource.Attribute{
					Name:        "max_log_file_size",
					Description: `Maximum log file size before rolling.`,
				},
				resource.Attribute{
					Name:        "roll_day",
					Description: `Days of week to roll logs. sunday - Sunday. monday - Monday. tuesday - Tuesday. wednesday - Wednesday. thursday - Thursday. friday - Friday. saturday - Saturday. Valid values: ` + "`" + `sunday` + "`" + `, ` + "`" + `monday` + "`" + `, ` + "`" + `tuesday` + "`" + `, ` + "`" + `wednesday` + "`" + `, ` + "`" + `thursday` + "`" + `, ` + "`" + `friday` + "`" + `, ` + "`" + `saturday` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "roll_schedule",
					Description: `Frequency to check log file for rolling. none - Not scheduled. daily - Every day. weekly - Every week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "roll_time",
					Description: `Time to roll logs (hh:mm).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Server type. FTP - Upload via FTP. SFTP - Upload via SFTP. SCP - Upload via SCP. Valid values: ` + "`" + `FTP` + "`" + `, ` + "`" + `SFTP` + "`" + `, ` + "`" + `SCP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable local disk log. disable - Do not log to local disk. enable - Log to local disk. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Upload log file when rolling. disable - Disable uploading when rolling log file. enable - Enable uploading when rolling log file. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_delete_files",
					Description: `Delete log files after uploading (default = enable). disable - Do not delete log files after uploading. enable - Delete log files after uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_time",
					Description: `Time to upload logs (hh:mm).`,
				},
				resource.Attribute{
					Name:        "uploaddir",
					Description: `Log file upload remote directory.`,
				},
				resource.Attribute{
					Name:        "uploadip",
					Description: `IP address of log uploading server.`,
				},
				resource.Attribute{
					Name:        "uploadpass",
					Description: `Password of user account in upload server.`,
				},
				resource.Attribute{
					Name:        "uploadport",
					Description: `Server port (0 = default protocol port).`,
				},
				resource.Attribute{
					Name:        "uploadsched",
					Description: `Scheduled upload (disable = upload when rolling). disable - Upload when rolling. enable - Scheduled upload. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "uploadtype",
					Description: `Types of log files that need to be uploaded. event - Upload event log. Valid values: ` + "`" + `event` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "uploaduser",
					Description: `User account in upload server.`,
				},
				resource.Attribute{
					Name:        "uploadzip",
					Description: `Compress upload logs. disable - Upload log files as plain text. enable - Upload log files compressed. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogDiskSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_disk_setting.labelname SystemLocallogDiskSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogDiskSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_disk_setting.labelname SystemLocallogDiskSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_fortianalyzer2_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for FortiAnalyzer2 logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"fortianalyzer2",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer2Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer2_filter.labelname SystemLocallogFortianalyzer2Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer2Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer2_filter.labelname SystemLocallogFortianalyzer2Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_fortianalyzer2_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for locallog to fortianalyzer.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"fortianalyzer2",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_cert_cn",
					Description: `Certificate common name of remote fortianalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Remote FortiAnalyzer server FQDN, hostname, or IP address`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Log to FortiAnalyzer status. disable - Log to FortiAnalyzer disabled. realtime - Log to FortiAnalyzer in realtime. upload - Log to FortiAnalyzer at schedule time. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `realtime` + "`" + `, ` + "`" + `upload` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_time",
					Description: `Time to upload local log files (hh:mm). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer2Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer2_setting.labelname SystemLocallogFortianalyzer2Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer2Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer2_setting.labelname SystemLocallogFortianalyzer2Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_fortianalyzer3_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for FortiAnalyzer3 logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"fortianalyzer3",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer3Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer3_filter.labelname SystemLocallogFortianalyzer3Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer3Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer3_filter.labelname SystemLocallogFortianalyzer3Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_fortianalyzer3_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for locallog to fortianalyzer.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"fortianalyzer3",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_cert_cn",
					Description: `Certificate common name of remote fortianalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Remote FortiAnalyzer server FQDN, hostname, or IP address`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Log to FortiAnalyzer status. disable - Log to FortiAnalyzer disabled. realtime - Log to FortiAnalyzer in realtime. upload - Log to FortiAnalyzer at schedule time. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `realtime` + "`" + `, ` + "`" + `upload` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_time",
					Description: `Time to upload local log files (hh:mm). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer3Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer3_setting.labelname SystemLocallogFortianalyzer3Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzer3Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer3_setting.labelname SystemLocallogFortianalyzer3Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_fortianalyzer_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for FortiAnalyzer logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzerFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer_filter.labelname SystemLocallogFortianalyzerFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzerFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer_filter.labelname SystemLocallogFortianalyzerFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_fortianalyzer_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for locallog to fortianalyzer.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "peer_cert_cn",
					Description: `Certificate common name of remote fortianalyzer.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Remote FortiAnalyzer server FQDN, hostname, or IP address`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Log to FortiAnalyzer status. disable - Log to FortiAnalyzer disabled. realtime - Log to FortiAnalyzer in realtime. upload - Log to FortiAnalyzer at schedule time. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `realtime` + "`" + `, ` + "`" + `upload` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_time",
					Description: `Time to upload local log files (hh:mm). ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzerSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer_setting.labelname SystemLocallogFortianalyzerSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogFortianalyzerSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_fortianalyzer_setting.labelname SystemLocallogFortianalyzerSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_memory_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for memory logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"memory",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogMemoryFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_memory_filter.labelname SystemLocallogMemoryFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogMemoryFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_memory_filter.labelname SystemLocallogMemoryFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_memory_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for memory buffer.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"memory",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "diskfull",
					Description: `Action upon disk full. overwrite - Overwrite oldest log when disk is full. nolog - Stop logging when disk is full. Valid values: ` + "`" + `overwrite` + "`" + `, ` + "`" + `nolog` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable memory buffer log. disable - Do not log to memory buffer. enable - Log to memory buffer. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogMemorySetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_memory_setting.labelname SystemLocallogMemorySetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogMemorySetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_memory_setting.labelname SystemLocallogMemorySetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for locallog logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "log_interval_dev_no_logging",
					Description: `Interval in minute for logging the event of no logs received from a device.`,
				},
				resource.Attribute{
					Name:        "log_interval_disk_full",
					Description: `Interval in minute for logging the event of disk full.`,
				},
				resource.Attribute{
					Name:        "log_interval_gbday_exceeded",
					Description: `Interval in minute for logging the event of the GB/Day license exceeded. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_setting.labelname SystemLocallogSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_setting.labelname SystemLocallogSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_syslogd2_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for syslog logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"syslogd2",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd2Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd2_filter.labelname SystemLocallogSyslogd2Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd2Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd2_filter.labelname SystemLocallogSyslogd2Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_syslogd2_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for remote syslog server.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"syslogd2",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `Select local certificate used for secure connection.`,
				},
				resource.Attribute{
					Name:        "csv",
					Description: `CSV format. disable - Disable CSV format. enable - Enable CSV format. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility. kernel - Kernel messages. user - Random user-level messages. ntp - NTP daemon. audit - Log audit. alert - Log alert. clock - Clock daemon. mail - Mail system. daemon - System daemons. auth - Security/authorization messages. syslog - Messages generated internally by syslog daemon. lpr - Line printer subsystem. news - Network news subsystem. uucp - Network news subsystem. cron - Clock daemon. authpriv - Security/authorization messages (private). ftp - FTP daemon. local0 - Reserved for local use. local1 - Reserved for local use. local2 - Reserved for local use. local3 - Reserved for local use. local4 - Reserved for local use. local5 - Reserved for local use. local6 - Reserved for local use. local7 - Reserved for local use. Valid values: ` + "`" + `kernel` + "`" + `, ` + "`" + `user` + "`" + `, ` + "`" + `ntp` + "`" + `, ` + "`" + `audit` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `clock` + "`" + `, ` + "`" + `mail` + "`" + `, ` + "`" + `daemon` + "`" + `, ` + "`" + `auth` + "`" + `, ` + "`" + `syslog` + "`" + `, ` + "`" + `lpr` + "`" + `, ` + "`" + `news` + "`" + `, ` + "`" + `uucp` + "`" + `, ` + "`" + `cron` + "`" + `, ` + "`" + `authpriv` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `local0` + "`" + `, ` + "`" + `local1` + "`" + `, ` + "`" + `local2` + "`" + `, ` + "`" + `local3` + "`" + `, ` + "`" + `local4` + "`" + `, ` + "`" + `local5` + "`" + `, ` + "`" + `local6` + "`" + `, ` + "`" + `local7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Remote syslog log. disable - Do not log to remote syslog server. enable - Log to remote syslog server. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "syslog_name",
					Description: `Remote syslog server name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd2Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd2_setting.labelname SystemLocallogSyslogd2Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd2Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd2_setting.labelname SystemLocallogSyslogd2Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_syslogd3_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for syslog logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"syslogd3",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd3Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd3_filter.labelname SystemLocallogSyslogd3Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd3Filter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd3_filter.labelname SystemLocallogSyslogd3Filter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_syslogd3_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for remote syslog server.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"syslogd3",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `Select local certificate used for secure connection.`,
				},
				resource.Attribute{
					Name:        "csv",
					Description: `CSV format. disable - Disable CSV format. enable - Enable CSV format. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility. kernel - Kernel messages. user - Random user-level messages. ntp - NTP daemon. audit - Log audit. alert - Log alert. clock - Clock daemon. mail - Mail system. daemon - System daemons. auth - Security/authorization messages. syslog - Messages generated internally by syslog daemon. lpr - Line printer subsystem. news - Network news subsystem. uucp - Network news subsystem. cron - Clock daemon. authpriv - Security/authorization messages (private). ftp - FTP daemon. local0 - Reserved for local use. local1 - Reserved for local use. local2 - Reserved for local use. local3 - Reserved for local use. local4 - Reserved for local use. local5 - Reserved for local use. local6 - Reserved for local use. local7 - Reserved for local use. Valid values: ` + "`" + `kernel` + "`" + `, ` + "`" + `user` + "`" + `, ` + "`" + `ntp` + "`" + `, ` + "`" + `audit` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `clock` + "`" + `, ` + "`" + `mail` + "`" + `, ` + "`" + `daemon` + "`" + `, ` + "`" + `auth` + "`" + `, ` + "`" + `syslog` + "`" + `, ` + "`" + `lpr` + "`" + `, ` + "`" + `news` + "`" + `, ` + "`" + `uucp` + "`" + `, ` + "`" + `cron` + "`" + `, ` + "`" + `authpriv` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `local0` + "`" + `, ` + "`" + `local1` + "`" + `, ` + "`" + `local2` + "`" + `, ` + "`" + `local3` + "`" + `, ` + "`" + `local4` + "`" + `, ` + "`" + `local5` + "`" + `, ` + "`" + `local6` + "`" + `, ` + "`" + `local7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Remote syslog log. disable - Do not log to remote syslog server. enable - Log to remote syslog server. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "syslog_name",
					Description: `Remote syslog server name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd3Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd3_setting.labelname SystemLocallogSyslogd3Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogd3Setting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd3_setting.labelname SystemLocallogSyslogd3Setting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_syslogd_filter",
			Category:         "System LocalLog",
			ShortDescription: `Filter for syslog logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"syslogd",
				"filter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aid",
					Description: `Log aid messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devcfg",
					Description: `Log device configuration message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "devops",
					Description: `Managered devices operations messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "diskquota",
					Description: `Log Fortianalyzer disk quota messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dm",
					Description: `Log deployment manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "docker",
					Description: `Docker application generic messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dvm",
					Description: `Log device manager messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ediscovery",
					Description: `Log Fortianalyzer ediscovery messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "epmgr",
					Description: `Log endpoint manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event",
					Description: `Log event messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "eventmgmt",
					Description: `Log Fortianalyzer event handler messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "faz",
					Description: `Log Fortianalyzer messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazha",
					Description: `Log Fortianalyzer HA messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fazsys",
					Description: `Log Fortianalyzer system messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgd",
					Description: `Log FortiGuard service message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fgfm",
					Description: `Log FGFM protocol message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fips",
					Description: `Whether to log fips messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmgws",
					Description: `Log web service messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmlmgr",
					Description: `Log FortiMail manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fmwmgr",
					Description: `Log firmware manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fortiview",
					Description: `Log Fortianalyzer FortiView messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "glbcfg",
					Description: `Log global database message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha",
					Description: `Log HA message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache",
					Description: `Log Fortianalyzer hcache messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "incident",
					Description: `Log Fortianalyzer incident messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "iolog",
					Description: `Log debug IO log message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logd",
					Description: `Log the status of log daemon. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdb",
					Description: `Log Fortianalyzer log DB messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logdev",
					Description: `Log Fortianalyzer log device messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logfile",
					Description: `Log Fortianalyzer log file messages. enable - Enable setting. disable - Disable setting. Valid values: ` + "`" + `enable` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "logging",
					Description: `Log Fortianalyzer logging messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lrmgr",
					Description: `Log log and report manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "objcfg",
					Description: `Log object configuration change message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "report",
					Description: `Log Fortianalyzer report messages. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rev",
					Description: `Log revision history message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rtmon",
					Description: `Log real-time monitor message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scfw",
					Description: `Log firewall objects message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scply",
					Description: `Log policy console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scrmgr",
					Description: `Log script manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scvpn",
					Description: `Log VPN console message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "system",
					Description: `Log system manager message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "webport",
					Description: `Log web portal message. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogdFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd_filter.labelname SystemLocallogSyslogdFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogdFilter can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd_filter.labelname SystemLocallogSyslogdFilter $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_locallog_syslogd_setting",
			Category:         "System LocalLog",
			ShortDescription: `Settings for remote syslog server.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"locallog",
				"syslogd",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cert",
					Description: `Select local certificate used for secure connection.`,
				},
				resource.Attribute{
					Name:        "csv",
					Description: `CSV format. disable - Disable CSV format. enable - Enable CSV format. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "facility",
					Description: `Remote syslog facility. kernel - Kernel messages. user - Random user-level messages. ntp - NTP daemon. audit - Log audit. alert - Log alert. clock - Clock daemon. mail - Mail system. daemon - System daemons. auth - Security/authorization messages. syslog - Messages generated internally by syslog daemon. lpr - Line printer subsystem. news - Network news subsystem. uucp - Network news subsystem. cron - Clock daemon. authpriv - Security/authorization messages (private). ftp - FTP daemon. local0 - Reserved for local use. local1 - Reserved for local use. local2 - Reserved for local use. local3 - Reserved for local use. local4 - Reserved for local use. local5 - Reserved for local use. local6 - Reserved for local use. local7 - Reserved for local use. Valid values: ` + "`" + `kernel` + "`" + `, ` + "`" + `user` + "`" + `, ` + "`" + `ntp` + "`" + `, ` + "`" + `audit` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `clock` + "`" + `, ` + "`" + `mail` + "`" + `, ` + "`" + `daemon` + "`" + `, ` + "`" + `auth` + "`" + `, ` + "`" + `syslog` + "`" + `, ` + "`" + `lpr` + "`" + `, ` + "`" + `news` + "`" + `, ` + "`" + `uucp` + "`" + `, ` + "`" + `cron` + "`" + `, ` + "`" + `authpriv` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `local0` + "`" + `, ` + "`" + `local1` + "`" + `, ` + "`" + `local2` + "`" + `, ` + "`" + `local3` + "`" + `, ` + "`" + `local4` + "`" + `, ` + "`" + `local5` + "`" + `, ` + "`" + `local6` + "`" + `, ` + "`" + `local7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "reliable",
					Description: `Enable/disable reliable realtime logging. disable - Disable reliable realtime logging. enable - Enable reliable realtime logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/disable connection secured by TLS/SSL. disable - Disable SSL connection. enable - Enable SSL connection. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "severity",
					Description: `Least severity level to log. emergency - Emergency level. alert - Alert level. critical - Critical level. error - Error level. warning - Warning level. notification - Notification level. information - Information level. debug - Debug level. Valid values: ` + "`" + `emergency` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `critical` + "`" + `, ` + "`" + `error` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `notification` + "`" + `, ` + "`" + `information` + "`" + `, ` + "`" + `debug` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Remote syslog log. disable - Do not log to remote syslog server. enable - Log to remote syslog server. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "syslog_name",
					Description: `Remote syslog server name. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogdSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd_setting.labelname SystemLocallogSyslogdSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LocallogSyslogdSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_locallog_syslogd_setting.labelname SystemLocallogSyslogdSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_alert",
			Category:         "System Log",
			ShortDescription: `Log based alert settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"alert",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_alert_count",
					Description: `Maximum number of alerts supported. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogAlert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_alert.labelname SystemLogAlert $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogAlert can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_alert.labelname SystemLogAlert $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_devicedisable",
			Category:         "System Log",
			ShortDescription: `Disable client device logging.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"devicedisable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "TTL",
					Description: `Time to Live`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device to be disabled logging`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `ID of device logging disable entry. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogDeviceDisable can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_devicedisable.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogDeviceDisable can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_devicedisable.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_fospolicystats",
			Category:         "No Category",
			ShortDescription: `FortiOS policy statistics settings.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"log",
				"fospolicystats",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "retention_days",
					Description: `Number of days for FortiOS policy stats data storage.`,
				},
				resource.Attribute{
					Name:        "sampling_interval",
					Description: `Interval to request policy stats data from FortiOS in minutes.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Disable/Enable FortiOS policy statistics feature. disable - Disable querying FortiOS policy stats. enable - Enable querying FortiOS policy stats. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogFosPolicyStats can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_fospolicystats.labelname SystemLogFosPolicyStats $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogFosPolicyStats can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_fospolicystats.labelname SystemLogFosPolicyStats $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_interfacestats",
			Category:         "System Log",
			ShortDescription: `Interface statistics settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"interfacestats",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "billing_report",
					Description: `Disable/Enable billing report feature. disable - Disable billing report. enable - Enable billing report. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "retention_days",
					Description: `Number of days for interface data storage.`,
				},
				resource.Attribute{
					Name:        "sampling_interval",
					Description: `Interval of receiving interface data from FortiGates in seconds.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Disable/Enable interface statistics feature. disable - Disable querying FortiGate interface stats. enable - Enable querying FortiGate interface stats. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogInterfaceStats can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_interfacestats.labelname SystemLogInterfaceStats $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogInterfaceStats can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_interfacestats.labelname SystemLogInterfaceStats $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_ioc",
			Category:         "System Log",
			ShortDescription: `IoC settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"ioc",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "notification",
					Description: `Disable/Enable IoC notification. disable - Disable IoC feature. enable - Enable IoC feature. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notification_throttle",
					Description: `Minute value for throttling the rate of IoC notifications.`,
				},
				resource.Attribute{
					Name:        "rescan_max_runner",
					Description: `Max count of cocurrent runner of IoC rescan.`,
				},
				resource.Attribute{
					Name:        "rescan_run_at",
					Description: `When to run IoC rescan.`,
				},
				resource.Attribute{
					Name:        "rescan_status",
					Description: `Disable/Enable IoC rescan. disable - Disable IoC feature. enable - Enable IoC feature. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Disable/Enable IoC feature. disable - Disable IoC feature. enable - Enable IoC feature. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogIoc can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ioc.labelname SystemLogIoc $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogIoc can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ioc.labelname SystemLogIoc $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_maildomain",
			Category:         "System Log",
			ShortDescription: `FortiMail domain setting.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"maildomain",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "devices",
					Description: `Devices for domain to vdom mapping`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `FortiMail domain`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `ID of FortiMail domain.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Virtual Domain name mapping to FortiMail domain ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogMailDomain can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_maildomain.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogMailDomain can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_maildomain.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_ratelimit",
			Category:         "System Others",
			ShortDescription: `Logging rate limit.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"log",
				"ratelimit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device",
					Description: `Device. The structure of ` + "`" + `device` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "device_ratelimit_default",
					Description: `Default maximum device log rate limit.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Logging rate limit mode. disable - Logging rate limit function disabled. manual - System rate limit and device rate limit both configurable, no limit if not configured. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `manual` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ratelimits",
					Description: `Ratelimits. The structure of ` + "`" + `ratelimits` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "system_ratelimit",
					Description: `Maximum system log rate limit.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `device` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device(s) filter according to filter-type setting, wildcard expression supported.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `Device filter type. devid - Device ID. Valid values: ` + "`" + `devid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Device filter ID.`,
				},
				resource.Attribute{
					Name:        "ratelimit",
					Description: `Maximum device log rate limit. The ` + "`" + `ratelimits` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "filter",
					Description: `Device or ADOM filter according to filter-type setting, wildcard expression supported.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `Device filter type. devid - Device ID. adom - ADOM name. Valid values: ` + "`" + `devid` + "`" + `, ` + "`" + `adom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Filter ID.`,
				},
				resource.Attribute{
					Name:        "ratelimit",
					Description: `Maximum log rate limit. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogRatelimit can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ratelimit.labelname SystemLogRatelimit $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogRatelimit can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ratelimit.labelname SystemLogRatelimit $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_ratelimit_device",
			Category:         "System Others",
			ShortDescription: `Device log rate limit.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"log",
				"ratelimit",
				"device",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device",
					Description: `Device(s) filter according to filter-type setting, wildcard expression supported.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `Device filter type. devid - Device ID. Valid values: ` + "`" + `devid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Device filter ID.`,
				},
				resource.Attribute{
					Name:        "ratelimit",
					Description: `Maximum device log rate limit. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogRatelimitDevice can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ratelimit_device.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogRatelimitDevice can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ratelimit_device.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_ratelimit_ratelimits",
			Category:         "No Category",
			ShortDescription: `Per device or ADOM log rate limits.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"log",
				"ratelimit",
				"ratelimits",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "filter",
					Description: `Device or ADOM filter according to filter-type setting, wildcard expression supported.`,
				},
				resource.Attribute{
					Name:        "filter_type",
					Description: `Device filter type. devid - Device ID. adom - ADOM name. Valid values: ` + "`" + `devid` + "`" + `, ` + "`" + `adom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Filter ID.`,
				},
				resource.Attribute{
					Name:        "ratelimit",
					Description: `Maximum log rate limit. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogRatelimitRatelimits can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ratelimit_ratelimits.labelname SystemLogRatelimitRatelimits $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogRatelimitRatelimits can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_ratelimit_ratelimits.labelname SystemLogRatelimitRatelimits $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_settings",
			Category:         "System Log",
			ShortDescription: `Log settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"settings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "FAC_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FAZ_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FCH_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FCT_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FDD_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FGT_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FMG_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FML_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FPX_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FSA_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "FWB_custom_field1",
					Description: `Name of custom log field to index.`,
				},
				resource.Attribute{
					Name:        "browse_max_logfiles",
					Description: `Maximum number of log files for each log browse attempt for each Adom.`,
				},
				resource.Attribute{
					Name:        "dns_resolve_dstip",
					Description: `Enable/Disable resolving destination IP by DNS. disable - Disable resolving destination IP by DNS. enable - Enable resolving destination IP by DNS. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "download_max_logs",
					Description: `Maximum number of logs for each log download attempt.`,
				},
				resource.Attribute{
					Name:        "ha_auto_migrate",
					Description: `Enabled/Disable automatically merging HA member's logs to HA cluster. disable - Disable automatically merging HA member's logs to HA cluster. enable - Enable automatically merging HA member's logs to HA cluster. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "import_max_logfiles",
					Description: `Maximum number of log files for each log import attempt.`,
				},
				resource.Attribute{
					Name:        "keep_dev_logs",
					Description: `Enable/Disable keeping the dev logs after the device has been deleted. disable - Disable keeping the dev logs after the device has been deleted. enable - Enable keeping the dev logs after the device has been deleted. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_file_archive_name",
					Description: `Log file name format for archiving, such as backup, upload or download. basic - Basic format for log archive file name, e.g. FGT20C0000000001.tlog.1417797247.log. extended - Extended format for log archive file name, e.g. FGT20C0000000001.2014-12-05-08:34:58.tlog.1417797247.log. Valid values: ` + "`" + `basic` + "`" + `, ` + "`" + `extended` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rolling_analyzer",
					Description: `Rolling-Analyzer. The structure of ` + "`" + `rolling_analyzer` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "rolling_local",
					Description: `Rolling-Local. The structure of ` + "`" + `rolling_local` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "rolling_regular",
					Description: `Rolling-Regular. The structure of ` + "`" + `rolling_regular` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sync_search_timeout",
					Description: `Maximum number of seconds for running a log search session in synchronous mode. The ` + "`" + `rolling_analyzer` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `, ` + "`" + `tue` + "`" + `, ` + "`" + `wed` + "`" + `, ` + "`" + `thu` + "`" + `, ` + "`" + `fri` + "`" + `, ` + "`" + `sat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "del_files",
					Description: `Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Upload server directory, for Unix server, use absolute`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `Roll log files when they reach this size (MB).`,
				},
				resource.Attribute{
					Name:        "gzip_format",
					Description: `Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Log files rolling schedule (hour).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Upload server IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `Upload server IP2 address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `Upload server IP3 address.`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: ` + "`" + `native` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `csv` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Log files rolling schedule (minutes).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Upload server login password.`,
				},
				resource.Attribute{
					Name:        "password2",
					Description: `Upload server login password2.`,
				},
				resource.Attribute{
					Name:        "password3",
					Description: `Upload server login password3.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Upload server IP1 port number.`,
				},
				resource.Attribute{
					Name:        "port2",
					Description: `Upload server IP2 port number.`,
				},
				resource.Attribute{
					Name:        "port3",
					Description: `Upload server IP3 port number.`,
				},
				resource.Attribute{
					Name:        "rolling_upgrade_status",
					Description: `rolling upgrade status (1|0).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: ` + "`" + `ftp` + "`" + `, ` + "`" + `sftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_hour",
					Description: `Log files upload schedule (hour).`,
				},
				resource.Attribute{
					Name:        "upload_mode",
					Description: `Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: ` + "`" + `backup` + "`" + `, ` + "`" + `mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_trigger",
					Description: `Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: ` + "`" + `on-roll` + "`" + `, ` + "`" + `on-schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Upload server login username.`,
				},
				resource.Attribute{
					Name:        "username2",
					Description: `Upload server login username2.`,
				},
				resource.Attribute{
					Name:        "username3",
					Description: `Upload server login username3.`,
				},
				resource.Attribute{
					Name:        "when",
					Description: `Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `. The ` + "`" + `rolling_local` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `, ` + "`" + `tue` + "`" + `, ` + "`" + `wed` + "`" + `, ` + "`" + `thu` + "`" + `, ` + "`" + `fri` + "`" + `, ` + "`" + `sat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "del_files",
					Description: `Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Upload server directory, for Unix server, use absolute`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `Roll log files when they reach this size (MB).`,
				},
				resource.Attribute{
					Name:        "gzip_format",
					Description: `Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Log files rolling schedule (hour).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Upload server IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `Upload server IP2 address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `Upload server IP3 address.`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: ` + "`" + `native` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `csv` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Log files rolling schedule (minutes).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Upload server login password.`,
				},
				resource.Attribute{
					Name:        "password2",
					Description: `Upload server login password2.`,
				},
				resource.Attribute{
					Name:        "password3",
					Description: `Upload server login password3.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Upload server IP1 port number.`,
				},
				resource.Attribute{
					Name:        "port2",
					Description: `Upload server IP2 port number.`,
				},
				resource.Attribute{
					Name:        "port3",
					Description: `Upload server IP3 port number.`,
				},
				resource.Attribute{
					Name:        "rolling_upgrade_status",
					Description: `rolling upgrade status (1|0).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: ` + "`" + `ftp` + "`" + `, ` + "`" + `sftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_hour",
					Description: `Log files upload schedule (hour).`,
				},
				resource.Attribute{
					Name:        "upload_mode",
					Description: `Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: ` + "`" + `backup` + "`" + `, ` + "`" + `mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_trigger",
					Description: `Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: ` + "`" + `on-roll` + "`" + `, ` + "`" + `on-schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Upload server login username.`,
				},
				resource.Attribute{
					Name:        "username2",
					Description: `Upload server login username2.`,
				},
				resource.Attribute{
					Name:        "username3",
					Description: `Upload server login username3.`,
				},
				resource.Attribute{
					Name:        "when",
					Description: `Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `. The ` + "`" + `rolling_regular` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "days",
					Description: `Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `, ` + "`" + `tue` + "`" + `, ` + "`" + `wed` + "`" + `, ` + "`" + `thu` + "`" + `, ` + "`" + `fri` + "`" + `, ` + "`" + `sat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "del_files",
					Description: `Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Upload server directory, for Unix server, use absolute`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `Roll log files when they reach this size (MB).`,
				},
				resource.Attribute{
					Name:        "gzip_format",
					Description: `Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Log files rolling schedule (hour).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Upload server IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `Upload server IP2 address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `Upload server IP3 address.`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: ` + "`" + `native` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `csv` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Log files rolling schedule (minutes).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Upload server login password.`,
				},
				resource.Attribute{
					Name:        "password2",
					Description: `Upload server login password2.`,
				},
				resource.Attribute{
					Name:        "password3",
					Description: `Upload server login password3.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Upload server IP1 port number.`,
				},
				resource.Attribute{
					Name:        "port2",
					Description: `Upload server IP2 port number.`,
				},
				resource.Attribute{
					Name:        "port3",
					Description: `Upload server IP3 port number.`,
				},
				resource.Attribute{
					Name:        "rolling_upgrade_status",
					Description: `rolling upgrade status (1|0).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: ` + "`" + `ftp` + "`" + `, ` + "`" + `sftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_hour",
					Description: `Log files upload schedule (hour).`,
				},
				resource.Attribute{
					Name:        "upload_mode",
					Description: `Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: ` + "`" + `backup` + "`" + `, ` + "`" + `mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_trigger",
					Description: `Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: ` + "`" + `on-roll` + "`" + `, ` + "`" + `on-schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Upload server login username.`,
				},
				resource.Attribute{
					Name:        "username2",
					Description: `Upload server login username2.`,
				},
				resource.Attribute{
					Name:        "username3",
					Description: `Upload server login username3.`,
				},
				resource.Attribute{
					Name:        "when",
					Description: `Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettings can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings.labelname SystemLogSettings $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettings can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings.labelname SystemLogSettings $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_settings_rollinganalyzer",
			Category:         "System Log",
			ShortDescription: `Log rolling policy for Network Analyzer logs.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"settings",
				"rollinganalyzer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "days",
					Description: `Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `, ` + "`" + `tue` + "`" + `, ` + "`" + `wed` + "`" + `, ` + "`" + `thu` + "`" + `, ` + "`" + `fri` + "`" + `, ` + "`" + `sat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "del_files",
					Description: `Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Upload server directory, for Unix server, use absolute`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `Roll log files when they reach this size (MB).`,
				},
				resource.Attribute{
					Name:        "gzip_format",
					Description: `Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Log files rolling schedule (hour).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Upload server IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `Upload server IP2 address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `Upload server IP3 address.`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: ` + "`" + `native` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `csv` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Log files rolling schedule (minutes).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Upload server login password.`,
				},
				resource.Attribute{
					Name:        "password2",
					Description: `Upload server login password2.`,
				},
				resource.Attribute{
					Name:        "password3",
					Description: `Upload server login password3.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Upload server IP1 port number.`,
				},
				resource.Attribute{
					Name:        "port2",
					Description: `Upload server IP2 port number.`,
				},
				resource.Attribute{
					Name:        "port3",
					Description: `Upload server IP3 port number.`,
				},
				resource.Attribute{
					Name:        "rolling_upgrade_status",
					Description: `rolling upgrade status (1|0).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: ` + "`" + `ftp` + "`" + `, ` + "`" + `sftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_hour",
					Description: `Log files upload schedule (hour).`,
				},
				resource.Attribute{
					Name:        "upload_mode",
					Description: `Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: ` + "`" + `backup` + "`" + `, ` + "`" + `mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_trigger",
					Description: `Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: ` + "`" + `on-roll` + "`" + `, ` + "`" + `on-schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Upload server login username.`,
				},
				resource.Attribute{
					Name:        "username2",
					Description: `Upload server login username2.`,
				},
				resource.Attribute{
					Name:        "username3",
					Description: `Upload server login username3.`,
				},
				resource.Attribute{
					Name:        "when",
					Description: `Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettingsRollingAnalyzer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings_rollinganalyzer.labelname SystemLogSettingsRollingAnalyzer $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettingsRollingAnalyzer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings_rollinganalyzer.labelname SystemLogSettingsRollingAnalyzer $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_settings_rollinglocal",
			Category:         "System Log",
			ShortDescription: `Log rolling policy for local logs.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"settings",
				"rollinglocal",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "days",
					Description: `Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `, ` + "`" + `tue` + "`" + `, ` + "`" + `wed` + "`" + `, ` + "`" + `thu` + "`" + `, ` + "`" + `fri` + "`" + `, ` + "`" + `sat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "del_files",
					Description: `Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Upload server directory, for Unix server, use absolute`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `Roll log files when they reach this size (MB).`,
				},
				resource.Attribute{
					Name:        "gzip_format",
					Description: `Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Log files rolling schedule (hour).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Upload server IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `Upload server IP2 address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `Upload server IP3 address.`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: ` + "`" + `native` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `csv` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Log files rolling schedule (minutes).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Upload server login password.`,
				},
				resource.Attribute{
					Name:        "password2",
					Description: `Upload server login password2.`,
				},
				resource.Attribute{
					Name:        "password3",
					Description: `Upload server login password3.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Upload server IP1 port number.`,
				},
				resource.Attribute{
					Name:        "port2",
					Description: `Upload server IP2 port number.`,
				},
				resource.Attribute{
					Name:        "port3",
					Description: `Upload server IP3 port number.`,
				},
				resource.Attribute{
					Name:        "rolling_upgrade_status",
					Description: `rolling upgrade status (1|0).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: ` + "`" + `ftp` + "`" + `, ` + "`" + `sftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_hour",
					Description: `Log files upload schedule (hour).`,
				},
				resource.Attribute{
					Name:        "upload_mode",
					Description: `Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: ` + "`" + `backup` + "`" + `, ` + "`" + `mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_trigger",
					Description: `Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: ` + "`" + `on-roll` + "`" + `, ` + "`" + `on-schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Upload server login username.`,
				},
				resource.Attribute{
					Name:        "username2",
					Description: `Upload server login username2.`,
				},
				resource.Attribute{
					Name:        "username3",
					Description: `Upload server login username3.`,
				},
				resource.Attribute{
					Name:        "when",
					Description: `Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettingsRollingLocal can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings_rollinglocal.labelname SystemLogSettingsRollingLocal $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettingsRollingLocal can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings_rollinglocal.labelname SystemLogSettingsRollingLocal $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_settings_rollingregular",
			Category:         "System Log",
			ShortDescription: `Log rolling policy for device logs.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"settings",
				"rollingregular",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "days",
					Description: `Log files rolling schedule (days of week). sun - Sunday. mon - Monday. tue - Tuesday. wed - Wednesday. thu - Thursday. fri - Friday. sat - Saturday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `, ` + "`" + `tue` + "`" + `, ` + "`" + `wed` + "`" + `, ` + "`" + `thu` + "`" + `, ` + "`" + `fri` + "`" + `, ` + "`" + `sat` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "del_files",
					Description: `Enable/disable log file deletion after uploading. disable - Disable log file deletion. enable - Enable log file deletion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "directory",
					Description: `Upload server directory, for Unix server, use absolute`,
				},
				resource.Attribute{
					Name:        "file_size",
					Description: `Roll log files when they reach this size (MB).`,
				},
				resource.Attribute{
					Name:        "gzip_format",
					Description: `Enable/disable compression of uploaded log files. disable - Disable compression. enable - Enable compression. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hour",
					Description: `Log files rolling schedule (hour).`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Upload server IP address.`,
				},
				resource.Attribute{
					Name:        "ip2",
					Description: `Upload server IP2 address.`,
				},
				resource.Attribute{
					Name:        "ip3",
					Description: `Upload server IP3 address.`,
				},
				resource.Attribute{
					Name:        "log_format",
					Description: `Format of uploaded log files. native - Native format (text or compact). text - Text format (convert if necessary). csv - CSV (comma-separated value) format. Valid values: ` + "`" + `native` + "`" + `, ` + "`" + `text` + "`" + `, ` + "`" + `csv` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "min",
					Description: `Log files rolling schedule (minutes).`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Upload server login password.`,
				},
				resource.Attribute{
					Name:        "password2",
					Description: `Upload server login password2.`,
				},
				resource.Attribute{
					Name:        "password3",
					Description: `Upload server login password3.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Upload server IP1 port number.`,
				},
				resource.Attribute{
					Name:        "port2",
					Description: `Upload server IP2 port number.`,
				},
				resource.Attribute{
					Name:        "port3",
					Description: `Upload server IP3 port number.`,
				},
				resource.Attribute{
					Name:        "rolling_upgrade_status",
					Description: `rolling upgrade status (1|0).`,
				},
				resource.Attribute{
					Name:        "server_type",
					Description: `Upload server type. ftp - Upload via FTP. sftp - Upload via SFTP. scp - Upload via SCP. Valid values: ` + "`" + `ftp` + "`" + `, ` + "`" + `sftp` + "`" + `, ` + "`" + `scp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload",
					Description: `Enable/disable log file uploads. disable - Disable log files uploading. enable - Enable log files uploading. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_hour",
					Description: `Log files upload schedule (hour).`,
				},
				resource.Attribute{
					Name:        "upload_mode",
					Description: `Upload mode with multiple servers. backup - Servers are attempted and used one after the other upon failure to connect. mirror - All configured servers are attempted and used. Valid values: ` + "`" + `backup` + "`" + `, ` + "`" + `mirror` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "upload_trigger",
					Description: `Event triggering log files upload. on-roll - Upload log files after they are rolled. on-schedule - Upload log files daily. Valid values: ` + "`" + `on-roll` + "`" + `, ` + "`" + `on-schedule` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `Upload server login username.`,
				},
				resource.Attribute{
					Name:        "username2",
					Description: `Upload server login username2.`,
				},
				resource.Attribute{
					Name:        "username3",
					Description: `Upload server login username3.`,
				},
				resource.Attribute{
					Name:        "when",
					Description: `Roll log files periodically. none - Do not roll log files periodically. daily - Roll log files daily. weekly - Roll log files on certain days of week. Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `daily` + "`" + `, ` + "`" + `weekly` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettingsRollingRegular can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings_rollingregular.labelname SystemLogSettingsRollingRegular $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogSettingsRollingRegular can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_settings_rollingregular.labelname SystemLogSettingsRollingRegular $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_log_topology",
			Category:         "No Category",
			ShortDescription: `Logging topology settings.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"log",
				"topology",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_depth",
					Description: `Maximum descend levels below this device.`,
				},
				resource.Attribute{
					Name:        "max_depth_share",
					Description: `Maximum descend levels below this device to share with upstream. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogTopology can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_topology.labelname SystemLogTopology $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogTopology can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_log_topology.labelname SystemLogTopology $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_logfetch_clientprofile",
			Category:         "System Log",
			ShortDescription: `Log-fetch client profile settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"logfetch",
				"clientprofile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "client_adom",
					Description: `Log-fetch client side's adom name.`,
				},
				resource.Attribute{
					Name:        "data_range",
					Description: `Data-range for fetched logs. custom - Specify some other date and time range. Valid values: ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "data_range_value",
					Description: `Last n days or hours.`,
				},
				resource.Attribute{
					Name:        "device_filter",
					Description: `Device-Filter. The structure of ` + "`" + `device_filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `End date and time of the data-range <hh:mm yyyy/mm/dd>.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Log-fetch client profile ID.`,
				},
				resource.Attribute{
					Name:        "index_fetch_logs",
					Description: `Enable/Disable indexing logs automatically after fetching logs. disable - Disable attribute function. enable - Enable attribute function. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_filter",
					Description: `Log-Filter. The structure of ` + "`" + `log_filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "log_filter_logic",
					Description: `And/Or logic for log-filters. and - Logic And. or - Logic Or. Valid values: ` + "`" + `and` + "`" + `, ` + "`" + `or` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_filter_status",
					Description: `Enable/Disable log-filter. disable - Disable attribute function. enable - Enable attribute function. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of log-fetch client profile.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Log-fetch server login password.`,
				},
				resource.Attribute{
					Name:        "peer_cert_cn",
					Description: `Certificate common name of log-fetch server.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable/Disable protecting log-fetch connection with TLS/SSL. disable - Disable attribute function. enable - Enable attribute function. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_adom",
					Description: `Log-fetch server side's adom name.`,
				},
				resource.Attribute{
					Name:        "server_ip",
					Description: `Log-fetch server IP address.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start date and time of the data-range <hh:mm yyyy/mm/dd>.`,
				},
				resource.Attribute{
					Name:        "sync_adom_config",
					Description: `Enable/Disable sync adom related config. disable - Disable attribute function. enable - Enable attribute function. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Log-fetch server login username.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `device_filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Adom name.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device name or Serial number.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Add or edit a device filter.`,
				},
				resource.Attribute{
					Name:        "vdom",
					Description: `Vdom filters. The ` + "`" + `log_filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Log filter ID.`,
				},
				resource.Attribute{
					Name:        "oper",
					Description: `Field filter operator. &lt; - =Less than or equal to &gt; - =Greater than or equal to contain - Contain not-contain - Not contain match - Match (expression) Valid values: ` + "`" + `=` + "`" + `, ` + "`" + `!=` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `contain` + "`" + `, ` + "`" + `not-contain` + "`" + `, ` + "`" + `match` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Field filter operand or free-text matching expression. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogFetchClientProfile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logfetch_clientprofile.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogFetchClientProfile can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logfetch_clientprofile.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_logfetch_serversettings",
			Category:         "System Log",
			ShortDescription: `Log-fetch server settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"log",
				"logfetch",
				"serversettings",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_conn_per_session",
					Description: `Max concurrent file download connections per session.`,
				},
				resource.Attribute{
					Name:        "max_sessions",
					Description: `Max concurrent fetch sessions.`,
				},
				resource.Attribute{
					Name:        "session_timeout",
					Description: `Fetch session timeout in minute. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogFetchServerSettings can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logfetch_serversettings.labelname SystemLogFetchServerSettings $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogFetchServerSettings can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logfetch_serversettings.labelname SystemLogFetchServerSettings $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_logforward",
			Category:         "No Category",
			ShortDescription: `Log forwarding.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"logforward",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "agg_archive_types",
					Description: `Archive types. Web_Archive - Secure_Web_Archive - Email_Archive - File_Transfer_Archive - IM_Archive - MMS_Archive - AV_Quarantine - IPS_Packets - Valid values: ` + "`" + `Web_Archive` + "`" + `, ` + "`" + `Secure_Web_Archive` + "`" + `, ` + "`" + `Email_Archive` + "`" + `, ` + "`" + `File_Transfer_Archive` + "`" + `, ` + "`" + `IM_Archive` + "`" + `, ` + "`" + `MMS_Archive` + "`" + `, ` + "`" + `AV_Quarantine` + "`" + `, ` + "`" + `IPS_Packets` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agg_data_end_time",
					Description: `End date and time of the data-range &lt;hh:mm yyyy/mm/dd&gt;.`,
				},
				resource.Attribute{
					Name:        "agg_data_start_time",
					Description: `Start date and time of the data-range &lt;hh:mm yyyy/mm/dd&gt;.`,
				},
				resource.Attribute{
					Name:        "agg_logtypes",
					Description: `Log types. none - none app-ctrl - attack - content - dlp - emailfilter - event - generic - history - traffic - virus - webfilter - netscan - fct-event - fct-traffic - fct-netscan - waf - gtp - dns - ssh - ssl - file-filter - asset - protocol - siem - Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `fct-event` + "`" + `, ` + "`" + `fct-traffic` + "`" + `, ` + "`" + `fct-netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `siem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agg_password",
					Description: `Log aggregation access password for server.`,
				},
				resource.Attribute{
					Name:        "agg_schedule",
					Description: `Schedule log aggregation mode. daily - Run daily log aggregation on-demand - Run log aggregation on demand Valid values: ` + "`" + `daily` + "`" + `, ` + "`" + `on-demand` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "agg_time",
					Description: `Daily at.`,
				},
				resource.Attribute{
					Name:        "agg_user",
					Description: `Log aggregation access user name for server.`,
				},
				resource.Attribute{
					Name:        "device_filter",
					Description: `Device-Filter. The structure of ` + "`" + `device_filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "fwd_archive_types",
					Description: `forwarding archive types. Web_Archive - Email_Archive - IM_Archive - File_Transfer_Archive - MMS_Archive - AV_Quarantine - IPS_Packets - EDISC_Archive - Valid values: ` + "`" + `Web_Archive` + "`" + `, ` + "`" + `Email_Archive` + "`" + `, ` + "`" + `IM_Archive` + "`" + `, ` + "`" + `File_Transfer_Archive` + "`" + `, ` + "`" + `MMS_Archive` + "`" + `, ` + "`" + `AV_Quarantine` + "`" + `, ` + "`" + `IPS_Packets` + "`" + `, ` + "`" + `EDISC_Archive` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_archives",
					Description: `Enable/disable forwarding archives. disable - Disable forwarding archives. enable - Enable forwarding archives. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_compression",
					Description: `Enable/disable compression for better bandwidth efficiency. disable - Disable compression of messages. enable - Enable compression of messages. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_facility",
					Description: `Facility for remote syslog. kernel - kernel messages user - random user level messages mail - Mail system. daemon - System daemons. auth - Security/authorization messages. syslog - Messages generated internally by syslog daemon. lpr - Line printer subsystem. news - Network news subsystem. uucp - Network news subsystem. clock - Clock daemon. authpriv - Security/authorization messages (private). ftp - FTP daemon. ntp - NTP daemon. audit - Log audit. alert - Log alert. cron - Clock daemon. local0 - Reserved for local use. local1 - Reserved for local use. local2 - Reserved for local use. local3 - Reserved for local use. local4 - Reserved for local use. local5 - Reserved for local use. local6 - Reserved for local use. local7 - Reserved for local use. Valid values: ` + "`" + `kernel` + "`" + `, ` + "`" + `user` + "`" + `, ` + "`" + `mail` + "`" + `, ` + "`" + `daemon` + "`" + `, ` + "`" + `auth` + "`" + `, ` + "`" + `syslog` + "`" + `, ` + "`" + `lpr` + "`" + `, ` + "`" + `news` + "`" + `, ` + "`" + `uucp` + "`" + `, ` + "`" + `clock` + "`" + `, ` + "`" + `authpriv` + "`" + `, ` + "`" + `ftp` + "`" + `, ` + "`" + `ntp` + "`" + `, ` + "`" + `audit` + "`" + `, ` + "`" + `alert` + "`" + `, ` + "`" + `cron` + "`" + `, ` + "`" + `local0` + "`" + `, ` + "`" + `local1` + "`" + `, ` + "`" + `local2` + "`" + `, ` + "`" + `local3` + "`" + `, ` + "`" + `local4` + "`" + `, ` + "`" + `local5` + "`" + `, ` + "`" + `local6` + "`" + `, ` + "`" + `local7` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_ha_bind_vip",
					Description: `When HA is enabled, always use vip as forwarding port disable - Disable bind forwarding to vip interface. enable - Enable bind forwarding to vip interface. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_log_source_ip",
					Description: `Log's source IP address (no effect for reliable forwarding). local_ip - Use FAZVM64 local ip. original_ip - Use original source ip. Valid values: ` + "`" + `local_ip` + "`" + `, ` + "`" + `original_ip` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_max_delay",
					Description: `Max delay for near realtime log forwarding. realtime - Realtime forwarding, no delay. 1min - Near realtime forwarding with up to one miniute delay. 5min - Near realtime forwarding with up to five miniutes delay. Valid values: ` + "`" + `realtime` + "`" + `, ` + "`" + `1min` + "`" + `, ` + "`" + `5min` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_reliable",
					Description: `Enable/disable reliable logging. disable - Disable reliable logging. enable - Enable reliable logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_secure",
					Description: `Enable/disable TLS/SSL secured reliable logging. disable - Disable TLS/SSL secured reliable logging. enable - Enable TLS/SSL secured reliable logging. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_server_type",
					Description: `Forwarding all logs to syslog server or FortiAnalyzer. syslog - Forward logs to generic syslog server. fortianalyzer - Forward logs to FortiAnalyzer. cef - Forward logs to a CEF (Common Event Format) server. syslog-pack - Forward logs to a FortiAnalyzer which support packed syslog message. Valid values: ` + "`" + `syslog` + "`" + `, ` + "`" + `fortianalyzer` + "`" + `, ` + "`" + `cef` + "`" + `, ` + "`" + `syslog-pack` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fwd_syslog_format",
					Description: `Forwarding format for syslog. fgt - fgt syslog format rfc-5424 - rfc-5424 syslog format Valid values: ` + "`" + `fgt` + "`" + `, ` + "`" + `rfc-5424` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Log forwarding ID.`,
				},
				resource.Attribute{
					Name:        "log_field_exclusion",
					Description: `Log-Field-Exclusion. The structure of ` + "`" + `log_field_exclusion` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "log_field_exclusion_status",
					Description: `Enable or disable log field exclusion. disable - Disable log field exclusion. enable - Enable log field exclusion. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_filter",
					Description: `Log-Filter. The structure of ` + "`" + `log_filter` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "log_filter_logic",
					Description: `Logic operator used to connect filters. and - Conjunctive filters. or - Disjunctive filters. Valid values: ` + "`" + `and` + "`" + `, ` + "`" + `or` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_filter_status",
					Description: `Enable or disable log filtering. disable - Disable log filtering. enable - Enable log filtering. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_masking_custom",
					Description: `Log-Masking-Custom. The structure of ` + "`" + `log_masking_custom` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "log_masking_custom_priority",
					Description: `Prioritize custom fields. disable - Disable custom field search priority. - Prioritize custom fields. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_masking_fields",
					Description: `Log field masking fields. user - User name. srcip - Source IP. srcname - Source name. srcmac - Source MAC. dstip - Destination IP. dstname - Dst name. email - Email. message - Message. domain - Domain. Valid values: ` + "`" + `user` + "`" + `, ` + "`" + `srcip` + "`" + `, ` + "`" + `srcname` + "`" + `, ` + "`" + `srcmac` + "`" + `, ` + "`" + `dstip` + "`" + `, ` + "`" + `dstname` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `message` + "`" + `, ` + "`" + `domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "log_masking_key",
					Description: `Log field masking key.`,
				},
				resource.Attribute{
					Name:        "log_masking_status",
					Description: `Enable or disable log field masking. disable - Disable log field masking. enable - Enable log field masking. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Log forwarding mode. forwarding - Realtime or near realtime forwarding logs to servers. aggregation - Aggregate logs and archives to Analyzer. disable - Do not forward or aggregate logs. Valid values: ` + "`" + `forwarding` + "`" + `, ` + "`" + `aggregation` + "`" + `, ` + "`" + `disable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "pcapurl_domain_ip",
					Description: `The domain name or ip for forming a pcapurl. This pcapurl will be appended to applicable forwarded logs for downloading a pcap file.`,
				},
				resource.Attribute{
					Name:        "pcapurl_enrich",
					Description: `Enable/disable enriching pcapurl. disable - Disable enriching pcapurl. enable - Enable enriching pcapurl. It will append a &apos;pcapurl&apos; field to the forwarded syslogs. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "peer_cert_cn",
					Description: `Certificate common name of log-forward server.`,
				},
				resource.Attribute{
					Name:        "proxy_service",
					Description: `Enable/disable proxy service under collector mode. disable - Disable proxy service. enable - Enable proxy service. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "proxy_service_priority",
					Description: `Proxy service priority from 1 (lowest) to 20 (highest).`,
				},
				resource.Attribute{
					Name:        "server_addr",
					Description: `Remote server address.`,
				},
				resource.Attribute{
					Name:        "server_device",
					Description: `Log forwarding server device ID.`,
				},
				resource.Attribute{
					Name:        "server_name",
					Description: `Log forwarding server name.`,
				},
				resource.Attribute{
					Name:        "server_port",
					Description: `Server listen port (1 - 65535).`,
				},
				resource.Attribute{
					Name:        "signature",
					Description: `Aggregation cfg hash token.`,
				},
				resource.Attribute{
					Name:        "sync_metadata",
					Description: `Synchronizing meta data types. sf-topology - Security Fabric topology interface-role - Interface Role device - Device information endusr-avatar - End-user avatar Valid values: ` + "`" + `sf-topology` + "`" + `, ` + "`" + `interface-role` + "`" + `, ` + "`" + `device` + "`" + `, ` + "`" + `endusr-avatar` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `device_filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "action",
					Description: `Include or exclude the specified device. include - Include specified device. exclude - Exclude specified device. include-like - Include specified device matching the given wildcard expression. exclude-like - Exclude specified device matching the given wildcard expression. Valid values: ` + "`" + `include` + "`" + `, ` + "`" + `exclude` + "`" + `, ` + "`" + `include-like` + "`" + `, ` + "`" + `exclude-like` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "adom",
					Description: `Adom name or (null) for all adoms, or a wildcard expression matching adom(s) if action is a like action.`,
				},
				resource.Attribute{
					Name:        "device",
					Description: `Device ID of log client device, or a wildcard expression matching log client device(s) if action is a like action.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Device filter ID. The ` + "`" + `log_field_exclusion` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "dev_type",
					Description: `Device type. FortiGate - FortiGate Device FortiManager - FortiManager Device Syslog - Syslog Device FortiMail - FortiMail Device FortiWeb - FortiWeb Device FortiCache - FortiCache Device FortiAnalyzer - FortiAnalyzer Device FortiSandbox - FortiSandbox Device FortiDDoS - FortiDDoS Device FortiNAC - FortiNAC Device FortiDeceptor - FortiDeceptor Device FortiADC - FortiADC Device FortiFirewall - FortiFirewall Device Valid values: ` + "`" + `FortiGate` + "`" + `, ` + "`" + `FortiManager` + "`" + `, ` + "`" + `Syslog` + "`" + `, ` + "`" + `FortiMail` + "`" + `, ` + "`" + `FortiWeb` + "`" + `, ` + "`" + `FortiCache` + "`" + `, ` + "`" + `FortiAnalyzer` + "`" + `, ` + "`" + `FortiSandbox` + "`" + `, ` + "`" + `FortiDDoS` + "`" + `, ` + "`" + `FortiNAC` + "`" + `, ` + "`" + `FortiDeceptor` + "`" + `, ` + "`" + `FortiADC` + "`" + `, ` + "`" + `FortiFirewall` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "field_list",
					Description: `List of fields to be excluded.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Log field exclusion ID.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log type. app-ctrl - Application Control appevent - APPEVENT attack - Attack content - DLP Archive dlp - Data Leak Prevention emailfilter - Email Filter event - Event generic - Generic history - Mail Statistics traffic - Traffic virus - Virus voip - VoIP webfilter - Web Filter netscan - Network Scan waf - WAF gtp - GTP dns - Domain Name System ssh - SSH ssl - SSL file-filter - FFLT Asset - Asset protocol - PROTOCOL ANY-TYPE - Any log type Valid values: ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `appevent` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `voip` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `Asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `ANY-TYPE` + "`" + `. The ` + "`" + `log_filter` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field",
					Description: `Field name. type - Log type logid - Log ID level - Level devid - Device ID vd - Vdom ID srcip - Source IP srcintf - Source Interface dstip - Destination IP dstintf - Destination Interface dstport - Destination Port user - User group - Group free-text - General free-text filter Valid values: ` + "`" + `type` + "`" + `, ` + "`" + `logid` + "`" + `, ` + "`" + `level` + "`" + `, ` + "`" + `devid` + "`" + `, ` + "`" + `vd` + "`" + `, ` + "`" + `srcip` + "`" + `, ` + "`" + `srcintf` + "`" + `, ` + "`" + `dstip` + "`" + `, ` + "`" + `dstintf` + "`" + `, ` + "`" + `dstport` + "`" + `, ` + "`" + `user` + "`" + `, ` + "`" + `group` + "`" + `, ` + "`" + `free-text` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Log filter ID.`,
				},
				resource.Attribute{
					Name:        "oper",
					Description: `Field filter operator. &lt; - =Less than or equal to &gt; - =Greater than or equal to contain - Contain not-contain - Not contain match - Match (expression) Valid values: ` + "`" + `=` + "`" + `, ` + "`" + `!=` + "`" + `, ` + "`" + `<` + "`" + `, ` + "`" + `>` + "`" + `, ` + "`" + `<=` + "`" + `, ` + "`" + `>=` + "`" + `, ` + "`" + `contain` + "`" + `, ` + "`" + `not-contain` + "`" + `, ` + "`" + `match` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Field filter operand or free-text matching expression. The ` + "`" + `log_masking_custom` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "field_name",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "field_type",
					Description: `Field type. string - String. ip - IP. mac - MAC address. email - Email address. unknown - Unknown. Valid values: ` + "`" + `string` + "`" + `, ` + "`" + `ip` + "`" + `, ` + "`" + `mac` + "`" + `, ` + "`" + `email` + "`" + `, ` + "`" + `unknown` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Field masking id. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogForward can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logforward.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System LogForward can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logforward.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_logforwardservice",
			Category:         "No Category",
			ShortDescription: `Log forwarding service.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"logforwardservice",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "accept_aggregation",
					Description: `Enable/disable accept log aggregation option. disable - Disable attribute function enable - Enable attribute function Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "aggregation_disk_quota",
					Description: `Aggregated device disk quota (MB) on server. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogForwardService can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logforwardservice.labelname SystemLogForwardService $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System LogForwardService can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_logforwardservice.labelname SystemLogForwardService $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_mail",
			Category:         "System Others",
			ShortDescription: `Alert emails.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"mail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth",
					Description: `Enable authentication. disable - Disable authentication. enable - Enable authentication. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_type",
					Description: `SMTP authentication type. psk - Use username and password to authenticate. certificate - Use local certificate to authenticate. Valid values: ` + "`" + `psk` + "`" + `, ` + "`" + `certificate` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Mail Service ID.`,
				},
				resource.Attribute{
					Name:        "local_cert",
					Description: `SMTP local certificate.`,
				},
				resource.Attribute{
					Name:        "passwd",
					Description: `SMTP account password.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `SMTP server port.`,
				},
				resource.Attribute{
					Name:        "secure_option",
					Description: `Communication secure option. default - Try STARTTLS, proceed as plain text communication otherwise. none - Communication will be in plain text format. smtps - Communication will be protected by SMTPS. starttls - Communication will be protected by STARTTLS. Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `smtps` + "`" + `, ` + "`" + `starttls` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `SMTP server.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `SMTP account username. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System Mail can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_mail.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System Mail can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_mail.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_metadata_admins",
			Category:         "System Others",
			ShortDescription: `Configure admins.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"metadata",
				"admins",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fieldlength",
					Description: `Field length. 20 - Field length of 20. 50 - Field length of 50. 255 - Field length of 255. Valid values: ` + "`" + `20` + "`" + `, ` + "`" + `50` + "`" + `, ` + "`" + `255` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fieldname",
					Description: `Field name.`,
				},
				resource.Attribute{
					Name:        "importance",
					Description: `Field importance. optional - This field is optional. required - This field is required. Valid values: ` + "`" + `optional` + "`" + `, ` + "`" + `required` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Field status. disabled - This field is disabled. enabled - This field is enabled. Valid values: ` + "`" + `disabled` + "`" + `, ` + "`" + `enabled` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fieldname}}. ## Import System MetadataAdmins can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_metadata_admins.labelname {{fieldname}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fieldname}}. ## Import System MetadataAdmins can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_metadata_admins.labelname {{fieldname}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_ntp",
			Category:         "System NTP",
			ShortDescription: `NTP settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"ntp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ntpserver",
					Description: `Ntpserver. The structure of ` + "`" + `ntpserver` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable NTP. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sync_interval",
					Description: `NTP sync interval (min).`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `ntpserver` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable/disable MD5 authentication. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Time server ID.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key for authentication.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID for authentication.`,
				},
				resource.Attribute{
					Name:        "maxpoll",
					Description: `Maximum poll interval in seconds as power of 2 (e.g. 6 means 64 seconds).`,
				},
				resource.Attribute{
					Name:        "minpoll",
					Description: `Minimum poll interval in seconds as power of 2 (e.g. 6 means 64 seconds).`,
				},
				resource.Attribute{
					Name:        "ntpv3",
					Description: `Enable/disable NTPv3. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `IP address/hostname of NTP Server. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Ntp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ntp.labelname SystemNtp $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Ntp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ntp.labelname SystemNtp $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_ntp_ntpserver",
			Category:         "System NTP",
			ShortDescription: `NTP server.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"ntp",
				"ntpserver",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "authentication",
					Description: `Enable/disable MD5 authentication. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Time server ID.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Key for authentication.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `Key ID for authentication.`,
				},
				resource.Attribute{
					Name:        "maxpoll",
					Description: `Maximum poll interval in seconds as power of 2 (e.g. 6 means 64 seconds).`,
				},
				resource.Attribute{
					Name:        "minpoll",
					Description: `Minimum poll interval in seconds as power of 2 (e.g. 6 means 64 seconds).`,
				},
				resource.Attribute{
					Name:        "ntpv3",
					Description: `Enable/disable NTPv3. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `IP address/hostname of NTP Server. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System NtpNtpserver can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ntp_ntpserver.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System NtpNtpserver can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_ntp_ntpserver.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_passwordpolicy",
			Category:         "System Others",
			ShortDescription: `Password policy.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"passwordpolicy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "change_4_characters",
					Description: `Enable/disable changing at least 4 characters for new password. disable - Disable changing at least 4 characters for new password. enable - Enable changing at least 4 characters for new password. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "expire",
					Description: `Number of days after which admin users' password will expire (0 - 3650, 0 = never expire).`,
				},
				resource.Attribute{
					Name:        "minimum_length",
					Description: `Minimum password length.`,
				},
				resource.Attribute{
					Name:        "must_contain",
					Description: `Password character requirements. upper-case-letter - Require password to contain upper case letter. lower-case-letter - Require password to contain lower case letter. number - Require password to contain number. non-alphanumeric - Require password to contain non-alphanumeric characters. Valid values: ` + "`" + `upper-case-letter` + "`" + `, ` + "`" + `lower-case-letter` + "`" + `, ` + "`" + `number` + "`" + `, ` + "`" + `non-alphanumeric` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable password policy. disable - Disable password policy. enable - Enable password policy. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System PasswordPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_passwordpolicy.labelname SystemPasswordPolicy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System PasswordPolicy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_passwordpolicy.labelname SystemPasswordPolicy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_report_autocache",
			Category:         "System Report",
			ShortDescription: `Report auto-cache settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"report",
				"autocache",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aggressive_schedule",
					Description: `Enable/disable auto-cache on schedule reports aggressively. disable - Disable the aggressive schedule auto-cache. enable - Enable the aggressive schedule auto-cache. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "order",
					Description: `The order of which SQL log table is processed first. oldest-first - The oldest SQL log table is processed first. Valid values: ` + "`" + `oldest-first` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable sql report auto cache. disable - Disable the sql report auto-cache. enable - Enable the sql report auto-cache. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System ReportAutoCache can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_report_autocache.labelname SystemReportAutoCache $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System ReportAutoCache can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_report_autocache.labelname SystemReportAutoCache $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_report_estbrowsetime",
			Category:         "System Report",
			ShortDescription: `Report estimated browse time settings`,
			Description:      ``,
			Keywords: []string{
				"system",
				"report",
				"estbrowsetime",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "max_read_time",
					Description: `Read time threshold for each page view.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Estimate browse time status. disable - Disable estimating browse time. enable - Enable estimating browse time. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System ReportEstBrowseTime can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_report_estbrowsetime.labelname SystemReportEstBrowseTime $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System ReportEstBrowseTime can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_report_estbrowsetime.labelname SystemReportEstBrowseTime $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_report_setting",
			Category:         "System Report",
			ShortDescription: `Report settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"report",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "aggregate_report",
					Description: `Enable/disable including a group report along with the per-device reports. disable - Exclude a group report along with the per-device reports. enable - Include a group report along with the per-device reports. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "capwap_port",
					Description: `Exclude capwap traffic by port.`,
				},
				resource.Attribute{
					Name:        "capwap_service",
					Description: `Exclude capwap traffic by service.`,
				},
				resource.Attribute{
					Name:        "exclude_capwap",
					Description: `Exclude capwap traffic. disable - Disable. by-port - By port. by-service - By service. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `by-port` + "`" + `, ` + "`" + `by-service` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hcache_lossless",
					Description: `Usableness of ready-with-loss hcaches. disable - Use ready-with-loss hcaches. enable - Do not use ready-with-loss hcaches. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ldap_cache_timeout",
					Description: `LDAP cache timeout in minutes, default 60, 0 means not use cache.`,
				},
				resource.Attribute{
					Name:        "max_table_rows",
					Description: `Maximum number of rows can be generated in a single table.`,
				},
				resource.Attribute{
					Name:        "report_priority",
					Description: `Priority of sql report. high - High low - Low auto - Auto Valid values: ` + "`" + `high` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `auto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "template_auto_install",
					Description: `The language used for new ADOMs (default = default). default - Default. english - English. Valid values: ` + "`" + `default` + "`" + `, ` + "`" + `english` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "week_start",
					Description: `Day of the week on which the week starts. sun - Sunday. mon - Monday. Valid values: ` + "`" + `sun` + "`" + `, ` + "`" + `mon` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System ReportSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_report_setting.labelname SystemReportSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System ReportSetting can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_report_setting.labelname SystemReportSetting $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_route",
			Category:         "System Route",
			ShortDescription: `Routing table configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"route",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "seq_num",
					Description: `Entry number. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{seq_num}}. ## Import System Route can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_route.labelname {{seq_num}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{seq_num}}. ## Import System Route can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_route.labelname {{seq_num}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_route6",
			Category:         "System Route",
			ShortDescription: `Routing table configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"route",
				"route6",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device",
					Description: `Gateway out interface.`,
				},
				resource.Attribute{
					Name:        "dst",
					Description: `Destination IP and mask for this route.`,
				},
				resource.Attribute{
					Name:        "gateway",
					Description: `Gateway IP for this route.`,
				},
				resource.Attribute{
					Name:        "prio",
					Description: `Entry number. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{prio}}. ## Import System Route6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_route6.labelname {{prio}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{prio}}. ## Import System Route6 can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_route6.labelname {{prio}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_saml",
			Category:         "System Saml",
			ShortDescription: `Global settings for SAML authentication.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"saml",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "acs_url",
					Description: `SP ACS(login) URL.`,
				},
				resource.Attribute{
					Name:        "cert",
					Description: `Certificate name.`,
				},
				resource.Attribute{
					Name:        "default_profile",
					Description: `Default Profile Name.`,
				},
				resource.Attribute{
					Name:        "entity_id",
					Description: `SP entity ID.`,
				},
				resource.Attribute{
					Name:        "fabric_idp",
					Description: `Fabric-Idp. The structure of ` + "`" + `fabric_idp` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "forticloud_sso",
					Description: `Enable/disable FortiCloud SSO (default = disable). disable - Disable Forticloud SSO. enable - Enabld Forticloud SSO. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "idp_cert",
					Description: `IDP Certificate name.`,
				},
				resource.Attribute{
					Name:        "idp_entity_id",
					Description: `IDP entity ID.`,
				},
				resource.Attribute{
					Name:        "idp_single_logout_url",
					Description: `IDP single logout url.`,
				},
				resource.Attribute{
					Name:        "idp_single_sign_on_url",
					Description: `IDP single sign-on URL.`,
				},
				resource.Attribute{
					Name:        "login_auto_redirect",
					Description: `Enable/Disable auto redirect to IDP login page. disable - Disable auto redirect to IDP Login Page. enable - Enable auto redirect to IDP Login Page. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `SAML role. IDP - IDentiy Provider. SP - Service Provider. FAB-SP - Fabric Service Provider. Valid values: ` + "`" + `IDP` + "`" + `, ` + "`" + `SP` + "`" + `, ` + "`" + `FAB-SP` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "server_address",
					Description: `server address.`,
				},
				resource.Attribute{
					Name:        "service_providers",
					Description: `Service-Providers. The structure of ` + "`" + `service_providers` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "sls_url",
					Description: `SP SLS(logout) URL.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable SAML authentication (default = disable). disable - Disable SAML authentication. enable - Enabld SAML authentication. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "user_auto_create",
					Description: `Enable/disable user auto creation (default = disable). disable - Disable auto create user. enable - Enable auto create user. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `fabric_idp` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "dev_id",
					Description: `IDP Device ID.`,
				},
				resource.Attribute{
					Name:        "idp_cert",
					Description: `IDP Certificate name.`,
				},
				resource.Attribute{
					Name:        "idp_entity_id",
					Description: `IDP entity ID.`,
				},
				resource.Attribute{
					Name:        "idp_single_logout_url",
					Description: `IDP single logout url.`,
				},
				resource.Attribute{
					Name:        "idp_single_sign_on_url",
					Description: `IDP single sign-on URL.`,
				},
				resource.Attribute{
					Name:        "idp_status",
					Description: `Enable/disable SAML authentication (default = disable). disable - Disable SAML authentication. enable - Enabld SAML authentication. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. The ` + "`" + `service_providers` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "idp_entity_id",
					Description: `IDP Entity ID.`,
				},
				resource.Attribute{
					Name:        "idp_single_logout_url",
					Description: `IDP single logout url.`,
				},
				resource.Attribute{
					Name:        "idp_single_sign_on_url",
					Description: `IDP single sign-on URL.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Prefix.`,
				},
				resource.Attribute{
					Name:        "sp_cert",
					Description: `SP certificate name.`,
				},
				resource.Attribute{
					Name:        "sp_entity_id",
					Description: `SP Entity ID.`,
				},
				resource.Attribute{
					Name:        "sp_single_logout_url",
					Description: `SP single logout URL.`,
				},
				resource.Attribute{
					Name:        "sp_single_sign_on_url",
					Description: `SP single sign-on URL. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Saml can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_saml.labelname SystemSaml $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Saml can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_saml.labelname SystemSaml $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_saml_fabricidp",
			Category:         "System Saml",
			ShortDescription: `Authorized identity providers.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"saml",
				"fabricidp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dev_id",
					Description: `IDP Device ID.`,
				},
				resource.Attribute{
					Name:        "idp_cert",
					Description: `IDP Certificate name.`,
				},
				resource.Attribute{
					Name:        "idp_entity_id",
					Description: `IDP entity ID.`,
				},
				resource.Attribute{
					Name:        "idp_single_logout_url",
					Description: `IDP single logout url.`,
				},
				resource.Attribute{
					Name:        "idp_single_sign_on_url",
					Description: `IDP single sign-on URL.`,
				},
				resource.Attribute{
					Name:        "idp_status",
					Description: `Enable/disable SAML authentication (default = disable). disable - Disable SAML authentication. enable - Enabld SAML authentication. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{dev_id}}. ## Import System SamlFabricIdp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_saml_fabricidp.labelname {{dev_id}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{dev_id}}. ## Import System SamlFabricIdp can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_saml_fabricidp.labelname {{dev_id}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_saml_serviceproviders",
			Category:         "System Saml",
			ShortDescription: `Authorized service providers.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"saml",
				"serviceproviders",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "idp_entity_id",
					Description: `IDP Entity ID.`,
				},
				resource.Attribute{
					Name:        "idp_single_logout_url",
					Description: `IDP single logout url.`,
				},
				resource.Attribute{
					Name:        "idp_single_sign_on_url",
					Description: `IDP single sign-on URL.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name.`,
				},
				resource.Attribute{
					Name:        "prefix",
					Description: `Prefix.`,
				},
				resource.Attribute{
					Name:        "sp_cert",
					Description: `SP certificate name.`,
				},
				resource.Attribute{
					Name:        "sp_entity_id",
					Description: `SP Entity ID.`,
				},
				resource.Attribute{
					Name:        "sp_single_logout_url",
					Description: `SP single logout URL.`,
				},
				resource.Attribute{
					Name:        "sp_single_sign_on_url",
					Description: `SP single sign-on URL. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System SamlServiceProviders can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_saml_serviceproviders.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System SamlServiceProviders can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_saml_serviceproviders.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_sniffer",
			Category:         "System Others",
			ShortDescription: `Interface sniffer.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"sniffer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "host",
					Description: `Hosts to filter for in sniffer traffic`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Sniffer ID.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Interface.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `Enable/disable sniffing IPv6 packets. disable - Disable sniffer for IPv6 packets. enable - Enable sniffer for IPv6 packets. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_packet_count",
					Description: `Maximum packet count (1000000, default = 4000).`,
				},
				resource.Attribute{
					Name:        "non_ip",
					Description: `Enable/disable sniffing non-IP packets. disable - Disable sniffer for non-IP packets. enable - Enable sniffer for non-IP packets. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Ports to sniff (Format examples: 10, 100-200).`,
				},
				resource.Attribute{
					Name:        "protocol",
					Description: `Integer value for the protocol type as defined by IANA (0 - 255).`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `List of VLANs to sniff. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System Sniffer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sniffer.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System Sniffer can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sniffer.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_snmp_community",
			Category:         "System SNMP",
			ShortDescription: `SNMP community configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"snmp",
				"community",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "events",
					Description: `SNMP trap events. disk_low - Disk usage too high. intf_ip_chg - Interface IP address changed. sys_reboot - System reboot. cpu_high - CPU usage too high. mem_low - Available memory is low. log-alert - Log base alert message. log-rate - High incoming log rate detected. log-data-rate - High incoming log data rate detected. lic-gbday - High licensed log GB/day detected. lic-dev-quota - High licensed device quota detected. cpu-high-exclude-nice - CPU usage exclude NICE threshold. Valid values: ` + "`" + `disk_low` + "`" + `, ` + "`" + `intf_ip_chg` + "`" + `, ` + "`" + `sys_reboot` + "`" + `, ` + "`" + `cpu_high` + "`" + `, ` + "`" + `mem_low` + "`" + `, ` + "`" + `log-alert` + "`" + `, ` + "`" + `log-rate` + "`" + `, ` + "`" + `log-data-rate` + "`" + `, ` + "`" + `lic-gbday` + "`" + `, ` + "`" + `lic-dev-quota` + "`" + `, ` + "`" + `cpu-high-exclude-nice` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hosts",
					Description: `Hosts. The structure of ` + "`" + `hosts` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "hosts6",
					Description: `Hosts6. The structure of ` + "`" + `hosts6` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Community ID.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Community name.`,
				},
				resource.Attribute{
					Name:        "query_v1_port",
					Description: `SNMP v1 query port.`,
				},
				resource.Attribute{
					Name:        "query_v1_status",
					Description: `Enable/disable SNMP v1 query. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_v2c_port",
					Description: `SNMP v2c query port.`,
				},
				resource.Attribute{
					Name:        "query_v2c_status",
					Description: `Enable/disable SNMP v2c query. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable community. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "trap_v1_rport",
					Description: `SNMP v1 trap remote port.`,
				},
				resource.Attribute{
					Name:        "trap_v1_status",
					Description: `Enable/disable SNMP v1 trap. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "trap_v2c_rport",
					Description: `SNMP v2c trap remote port.`,
				},
				resource.Attribute{
					Name:        "trap_v2c_status",
					Description: `Enable/disable SNMP v2c trap. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `hosts` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Host entry ID.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Allow interface name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Allow host IP address. The ` + "`" + `hosts6` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Host entry ID.`,
				},
				resource.Attribute{
					Name:        "interface",
					Description: `Allow interface name.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Allow host IP address. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System SnmpCommunity can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_snmp_community.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System SnmpCommunity can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_snmp_community.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_snmp_sysinfo",
			Category:         "System SNMP",
			ShortDescription: `SNMP configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"snmp",
				"sysinfo",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "contact_info",
					Description: `Contact information.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `System description.`,
				},
				resource.Attribute{
					Name:        "engine_id",
					Description: `Local SNMP engineID string (maximum 24 characters).`,
				},
				resource.Attribute{
					Name:        "fortianalyzer_legacy_sysoid",
					Description: `Enable legacy FortiAnalyzer sysObjectOID. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "location",
					Description: `System location.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable SNMP. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "trap_cpu_high_exclude_nice_threshold",
					Description: `SNMP trap for CPU usage threshold (exclude NICE processes).`,
				},
				resource.Attribute{
					Name:        "trap_high_cpu_threshold",
					Description: `SNMP trap for CPU usage threshold.`,
				},
				resource.Attribute{
					Name:        "trap_low_memory_threshold",
					Description: `SNMP trap for memory usage threshold. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System SnmpSysinfo can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_snmp_sysinfo.labelname SystemSnmpSysinfo $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System SnmpSysinfo can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_snmp_sysinfo.labelname SystemSnmpSysinfo $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_snmp_user",
			Category:         "System SNMP",
			ShortDescription: `SNMP user configuration.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"snmp",
				"user",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "auth_proto",
					Description: `Authentication protocol. md5 - HMAC-MD5-96 authentication protocol. sha - HMAC-SHA-96 authentication protocol. Valid values: ` + "`" + `md5` + "`" + `, ` + "`" + `sha` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auth_pwd",
					Description: `Password for authentication protocol.`,
				},
				resource.Attribute{
					Name:        "events",
					Description: `SNMP notifications (traps) to send. disk_low - Disk usage too high. intf_ip_chg - Interface IP address changed. sys_reboot - System reboot. cpu_high - CPU usage too high. mem_low - Available memory is low. log-alert - Log base alert message. log-rate - High incoming log rate detected. log-data-rate - High incoming log data rate detected. lic-gbday - High licensed log GB/day detected. lic-dev-quota - High licensed device quota detected. cpu-high-exclude-nice - CPU usage exclude NICE threshold. Valid values: ` + "`" + `disk_low` + "`" + `, ` + "`" + `intf_ip_chg` + "`" + `, ` + "`" + `sys_reboot` + "`" + `, ` + "`" + `cpu_high` + "`" + `, ` + "`" + `mem_low` + "`" + `, ` + "`" + `log-alert` + "`" + `, ` + "`" + `log-rate` + "`" + `, ` + "`" + `log-data-rate` + "`" + `, ` + "`" + `lic-gbday` + "`" + `, ` + "`" + `lic-dev-quota` + "`" + `, ` + "`" + `cpu-high-exclude-nice` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `SNMP user name.`,
				},
				resource.Attribute{
					Name:        "notify_hosts",
					Description: `Hosts to send notifications (traps) to.`,
				},
				resource.Attribute{
					Name:        "notify_hosts6",
					Description: `IPv6 hosts to send notifications (traps) to.`,
				},
				resource.Attribute{
					Name:        "priv_proto",
					Description: `Privacy (encryption) protocol. aes - CFB128-AES-128 symmetric encryption protocol. des - CBC-DES symmetric encryption protocol. Valid values: ` + "`" + `aes` + "`" + `, ` + "`" + `des` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "priv_pwd",
					Description: `Password for privacy (encryption) protocol.`,
				},
				resource.Attribute{
					Name:        "queries",
					Description: `Enable/disable queries for this user. disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "query_port",
					Description: `SNMPv3 query port.`,
				},
				resource.Attribute{
					Name:        "security_level",
					Description: `Security level for message authentication and encryption. no-auth-no-priv - Message with no authentication and no privacy (encryption). auth-no-priv - Message with authentication but no privacy (encryption). auth-priv - Message with authentication and privacy (encryption). Valid values: ` + "`" + `no-auth-no-priv` + "`" + `, ` + "`" + `auth-no-priv` + "`" + `, ` + "`" + `auth-priv` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System SnmpUser can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_snmp_user.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System SnmpUser can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_snmp_user.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_socfabric",
			Category:         "System Others",
			ShortDescription: `SOC Fabric.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"socfabric",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `Fabric name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `communication port (1 - 65535).`,
				},
				resource.Attribute{
					Name:        "psk",
					Description: `Fabric auth pwd.`,
				},
				resource.Attribute{
					Name:        "role",
					Description: `Enable or Disable SOC Fabric. member - SOC Fabric member. supervisor - SOC Fabric supervisor. Valid values: ` + "`" + `member` + "`" + `, ` + "`" + `supervisor` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "secure_connection",
					Description: `Enable or Disable SSL/TLS. disable - Disable SSL/TLS. enable - Enable SSL/TLS. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable or Disable SOC Fabric. disable - Disable SOC Fabric. enable - Enable SOC Fabric. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "supervisor",
					Description: `IP/FQDN of supervisor. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System SocFabric can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_socfabric.labelname SystemSocFabric $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System SocFabric can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_socfabric.labelname SystemSocFabric $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_sql",
			Category:         "System SQL",
			ShortDescription: `SQL settings.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sql",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "background_rebuild",
					Description: `Disable/Enable rebuild SQL database in the background. disable - Rebuild SQL database not in the background. enable - Rebuild SQL database in the background. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "compress_table_min_age",
					Description: `Minimum age in days for SQL tables to be compressed.`,
				},
				resource.Attribute{
					Name:        "custom_index",
					Description: `Custom-Index. The structure of ` + "`" + `custom_index` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "custom_skipidx",
					Description: `Custom-Skipidx. The structure of ` + "`" + `custom_skipidx` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "database_name",
					Description: `Database name.`,
				},
				resource.Attribute{
					Name:        "database_type",
					Description: `Database type. mysql - MySQL database. postgres - PostgreSQL local database. Valid values: ` + "`" + `mysql` + "`" + `, ` + "`" + `postgres` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_count_high",
					Description: `Must set to enable if the count of registered devices is greater than 8000. disable - Set to disable if device count is less than 8000. enable - Set to enable if device count is equal to or greater than 8000. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "event_table_partition_time",
					Description: `Maximum SQL database table partitioning time range in minute (0 for unlimited) for event logs.`,
				},
				resource.Attribute{
					Name:        "fct_table_partition_time",
					Description: `Maximum SQL database table partitioning time range in minute (0 for unlimited) for FortiClient logs.`,
				},
				resource.Attribute{
					Name:        "logtype",
					Description: `Log type. none - None. app-ctrl - attack - content - dlp - emailfilter - event - generic - history - traffic - virus - voip - webfilter - netscan - fct-event - fct-traffic - fct-netscan - waf - gtp - dns - ssh - ssl - file-filter - asset - protocol - siem - Valid values: ` + "`" + `none` + "`" + `, ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `voip` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `fct-event` + "`" + `, ` + "`" + `fct-traffic` + "`" + `, ` + "`" + `fct-netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `siem` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `Password for login remote database.`,
				},
				resource.Attribute{
					Name:        "prompt_sql_upgrade",
					Description: `Prompt to convert log database into SQL database at start time on GUI. disable - Do not prompt to upgrade log database to SQL database at start time on GUI. enable - Prompt to upgrade log database to SQL database at start time on GUI. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rebuild_event",
					Description: `Disable/Enable rebuild event during SQL database rebuilding. disable - Do not rebuild event during SQL database rebuilding. enable - Rebuild event during SQL database rebuilding. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "rebuild_event_start_time",
					Description: `Rebuild event starting date and time <hh:mm yyyy/mm/dd>.`,
				},
				resource.Attribute{
					Name:        "server",
					Description: `Database IP or hostname.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Start date and time <hh:mm yyyy/mm/dd>.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `SQL database status. disable - Disable SQL database. local - Enable local database. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `local` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "text_search_index",
					Description: `Disable/Enable text search index. disable - Do not create text search index. enable - Create text search index. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "traffic_table_partition_time",
					Description: `Maximum SQL database table partitioning time range in minute (0 for unlimited) for traffic logs.`,
				},
				resource.Attribute{
					Name:        "ts_index_field",
					Description: `Ts-Index-Field. The structure of ` + "`" + `ts_index_field` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `User name for login remote database.`,
				},
				resource.Attribute{
					Name:        "utm_table_partition_time",
					Description: `Maximum SQL database table partitioning time range in minute (0 for unlimited) for UTM logs.`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `custom_index` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `Disable/Enable case sensitive index. disable - Build a case insensitive index. enable - Build a case sensitive index. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Device type. FortiGate - Device type to FortiGate. FortiMail - Device type to FortiMail. FortiWeb - Device type to FortiWeb. Valid values: ` + "`" + `FortiGate` + "`" + `, ` + "`" + `FortiMail` + "`" + `, ` + "`" + `FortiWeb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Add or Edit log index fields.`,
				},
				resource.Attribute{
					Name:        "index_field",
					Description: `Log field name to be indexed.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log type. app-ctrl - attack - content - dlp - emailfilter - event - generic - history - traffic - virus - voip - webfilter - netscan - fct-event - fct-traffic - fct-netscan - waf - gtp - dns - ssh - ssl - file-filter - asset - protocol - siem - Valid values: ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `voip` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `fct-event` + "`" + `, ` + "`" + `fct-traffic` + "`" + `, ` + "`" + `fct-netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `siem` + "`" + `. The ` + "`" + `custom_skipidx` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Device type. FortiGate - Set device type to FortiGate. FortiManager - Set device type to FortiManager FortiClient - Set device type to FortiClient. FortiMail - Set device type to FortiMail. FortiWeb - Set device type to FortiWeb. FortiSandbox - Set device type to FortiSandbox FortiProxy - Set device type to FortiProxy Valid values: ` + "`" + `FortiGate` + "`" + `, ` + "`" + `FortiManager` + "`" + `, ` + "`" + `FortiClient` + "`" + `, ` + "`" + `FortiMail` + "`" + `, ` + "`" + `FortiWeb` + "`" + `, ` + "`" + `FortiSandbox` + "`" + `, ` + "`" + `FortiProxy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Add or Edit log index fields.`,
				},
				resource.Attribute{
					Name:        "index_field",
					Description: `Field to be added to skip index.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log type. app-ctrl - attack - content - dlp - emailfilter - event - generic - history - traffic - virus - voip - webfilter - netscan - fct-event - fct-traffic - fct-netscan - waf - gtp - dns - ssh - ssl - file-filter - asset - protocol - siem - Valid values: ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `voip` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `fct-event` + "`" + `, ` + "`" + `fct-traffic` + "`" + `, ` + "`" + `fct-netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `siem` + "`" + `. The ` + "`" + `ts_index_field` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "category",
					Description: `Category of text search index fields.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Fields of text search index. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Sql can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql.labelname SystemSql $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System Sql can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql.labelname SystemSql $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_sql_customindex",
			Category:         "System SQL",
			ShortDescription: `List of SQL index fields.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sql",
				"customindex",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "case_sensitive",
					Description: `Disable/Enable case sensitive index. disable - Build a case insensitive index. enable - Build a case sensitive index. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "device_type",
					Description: `Device type. FortiGate - Device type to FortiGate. FortiMail - Device type to FortiMail. FortiWeb - Device type to FortiWeb. Valid values: ` + "`" + `FortiGate` + "`" + `, ` + "`" + `FortiMail` + "`" + `, ` + "`" + `FortiWeb` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Add or Edit log index fields.`,
				},
				resource.Attribute{
					Name:        "index_field",
					Description: `Log field name to be indexed.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log type. app-ctrl - attack - content - dlp - emailfilter - event - generic - history - traffic - virus - voip - webfilter - netscan - fct-event - fct-traffic - fct-netscan - waf - gtp - dns - ssh - ssl - file-filter - asset - protocol - siem - Valid values: ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `voip` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `fct-event` + "`" + `, ` + "`" + `fct-traffic` + "`" + `, ` + "`" + `fct-netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `siem` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System SqlCustomIndex can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql_customindex.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System SqlCustomIndex can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql_customindex.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_sql_customskipidx",
			Category:         "System SQL",
			ShortDescription: `List of aditional SQL skip index fields.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sql",
				"customskipidx",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "device_type",
					Description: `Device type. FortiGate - Set device type to FortiGate. FortiManager - Set device type to FortiManager FortiClient - Set device type to FortiClient. FortiMail - Set device type to FortiMail. FortiWeb - Set device type to FortiWeb. FortiSandbox - Set device type to FortiSandbox FortiProxy - Set device type to FortiProxy Valid values: ` + "`" + `FortiGate` + "`" + `, ` + "`" + `FortiManager` + "`" + `, ` + "`" + `FortiClient` + "`" + `, ` + "`" + `FortiMail` + "`" + `, ` + "`" + `FortiWeb` + "`" + `, ` + "`" + `FortiSandbox` + "`" + `, ` + "`" + `FortiProxy` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "fosid",
					Description: `Add or Edit log index fields.`,
				},
				resource.Attribute{
					Name:        "index_field",
					Description: `Field to be added to skip index.`,
				},
				resource.Attribute{
					Name:        "log_type",
					Description: `Log type. app-ctrl - attack - content - dlp - emailfilter - event - generic - history - traffic - virus - voip - webfilter - netscan - fct-event - fct-traffic - fct-netscan - waf - gtp - dns - ssh - ssl - file-filter - asset - protocol - siem - Valid values: ` + "`" + `app-ctrl` + "`" + `, ` + "`" + `attack` + "`" + `, ` + "`" + `content` + "`" + `, ` + "`" + `dlp` + "`" + `, ` + "`" + `emailfilter` + "`" + `, ` + "`" + `event` + "`" + `, ` + "`" + `generic` + "`" + `, ` + "`" + `history` + "`" + `, ` + "`" + `traffic` + "`" + `, ` + "`" + `virus` + "`" + `, ` + "`" + `voip` + "`" + `, ` + "`" + `webfilter` + "`" + `, ` + "`" + `netscan` + "`" + `, ` + "`" + `fct-event` + "`" + `, ` + "`" + `fct-traffic` + "`" + `, ` + "`" + `fct-netscan` + "`" + `, ` + "`" + `waf` + "`" + `, ` + "`" + `gtp` + "`" + `, ` + "`" + `dns` + "`" + `, ` + "`" + `ssh` + "`" + `, ` + "`" + `ssl` + "`" + `, ` + "`" + `file-filter` + "`" + `, ` + "`" + `asset` + "`" + `, ` + "`" + `protocol` + "`" + `, ` + "`" + `siem` + "`" + `. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System SqlCustomSkipidx can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql_customskipidx.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{fosid}}. ## Import System SqlCustomSkipidx can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql_customskipidx.labelname {{fosid}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_sql_tsindexfield",
			Category:         "System SQL",
			ShortDescription: `List of SQL text search index fields.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"sql",
				"tsindexfield",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "category",
					Description: `Category of text search index fields.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Fields of text search index. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{category}}. ## Import System SqlTsIndexField can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql_tsindexfield.labelname {{category}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{category}}. ## Import System SqlTsIndexField can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_sql_tsindexfield.labelname {{category}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_syslog",
			Category:         "System Others",
			ShortDescription: `Syslog servers.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"syslog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ip",
					Description: `Syslog server IP address or hostname.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Syslog server name.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Syslog server port. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System Syslog can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_syslog.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{name}}. ## Import System Syslog can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_syslog.labelname {{name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_webproxy",
			Category:         "No Category",
			ShortDescription: `Configure system web proxy.`,
			Description:      ``,
			Keywords: []string{
				"no",
				"category",
				"system",
				"webproxy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "address",
					Description: `web proxy address.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `Web proxy mode (default = tunnel) proxy - proxy mode tunnel - tunnel mode (default) Valid values: ` + "`" + `proxy` + "`" + `, ` + "`" + `tunnel` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `The password for the user name used for authentication.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `The port number of the web proxy (1 - 65535, default = 1080).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Enable/disable system web proxy (default = disable). disable - Disable setting. enable - Enable setting. Valid values: ` + "`" + `disable` + "`" + `, ` + "`" + `enable` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `The user name used for authentication. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System WebProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_webproxy.labelname SystemWebProxy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource. ## Import System WebProxy can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_webproxy.labelname SystemWebProxy $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "fortianalyzer_fortianalyzer_system_workflow_approvalmatrix",
			Category:         "System Others",
			ShortDescription: `workflow approval matrix.`,
			Description:      ``,
			Keywords: []string{
				"system",
				"others",
				"workflow",
				"approvalmatrix",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "adom_name",
					Description: `Adom Name`,
				},
				resource.Attribute{
					Name:        "approver",
					Description: `Approver. The structure of ` + "`" + `approver` + "`" + ` block is documented below.`,
				},
				resource.Attribute{
					Name:        "mail_server",
					Description: `Notify mail server id.`,
				},
				resource.Attribute{
					Name:        "notify",
					Description: `Notify users`,
				},
				resource.Attribute{
					Name:        "dynamic_sort_subtable",
					Description: `true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables. The ` + "`" + `approver` + "`" + ` block supports:`,
				},
				resource.Attribute{
					Name:        "member",
					Description: `Member of approver.`,
				},
				resource.Attribute{
					Name:        "seq_num",
					Description: `Entry number. ## Attribute Reference In addition to all the above arguments, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{adom_name}}. ## Import System WorkflowApprovalMatrix can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_workflow_approvalmatrix.labelname {{adom_name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `an identifier for the resource with format {{adom_name}}. ## Import System WorkflowApprovalMatrix can be imported using any of these accepted formats: ` + "`" + `` + "`" + `` + "`" + ` $ export "FORTIANALYZER_IMPORT_TABLE"="true" $ terraform import fortianalyzer_system_workflow_approvalmatrix.labelname {{adom_name}} $ unset "FORTIANALYZER_IMPORT_TABLE" ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
	}

	resourcesMap = map[string]int{

		"fortianalyzer_fortianalyzer_dvm_cmd_add_device":                            0,
		"fortianalyzer_fortianalyzer_dvm_cmd_del_device":                            1,
		"fortianalyzer_fortianalyzer_dvmdb_adom":                                    2,
		"fortianalyzer_fortianalyzer_dvmdb_group":                                   3,
		"fortianalyzer_fortianalyzer_fmupdate_analyzer_virusreport":                 4,
		"fortianalyzer_fortianalyzer_fmupdate_avips_advancedlog":                    5,
		"fortianalyzer_fortianalyzer_fmupdate_avips_webproxy":                       6,
		"fortianalyzer_fortianalyzer_fmupdate_customurllist":                        7,
		"fortianalyzer_fortianalyzer_fmupdate_diskquota":                            8,
		"fortianalyzer_fortianalyzer_fmupdate_fctservices":                          9,
		"fortianalyzer_fortianalyzer_fmupdate_fdssetting":                           10,
		"fortianalyzer_fortianalyzer_fmupdate_fdssetting_pushoverride":              11,
		"fortianalyzer_fortianalyzer_fmupdate_fdssetting_pushoverridetoclient":      12,
		"fortianalyzer_fortianalyzer_fmupdate_fdssetting_serveroverride":            13,
		"fortianalyzer_fortianalyzer_fmupdate_fdssetting_updateschedule":            14,
		"fortianalyzer_fortianalyzer_fmupdate_fwmsetting":                           15,
		"fortianalyzer_fortianalyzer_fmupdate_multilayer":                           16,
		"fortianalyzer_fortianalyzer_fmupdate_publicnetwork":                        17,
		"fortianalyzer_fortianalyzer_fmupdate_serveraccesspriorities":               18,
		"fortianalyzer_fortianalyzer_fmupdate_serveroverridestatus":                 19,
		"fortianalyzer_fortianalyzer_fmupdate_service":                              20,
		"fortianalyzer_fortianalyzer_fmupdate_webspam_fgdsetting":                   21,
		"fortianalyzer_fortianalyzer_fmupdate_webspam_webproxy":                     22,
		"fortianalyzer_fortianalyzer_json_generic_api":                              23,
		"fortianalyzer_fortianalyzer_system_admin_group":                            24,
		"fortianalyzer_fortianalyzer_system_admin_ldap":                             25,
		"fortianalyzer_fortianalyzer_system_admin_profile":                          26,
		"fortianalyzer_fortianalyzer_system_admin_radius":                           27,
		"fortianalyzer_fortianalyzer_system_admin_setting":                          28,
		"fortianalyzer_fortianalyzer_system_admin_tacacs":                           29,
		"fortianalyzer_fortianalyzer_system_admin_user":                             30,
		"fortianalyzer_fortianalyzer_system_alertconsole":                           31,
		"fortianalyzer_fortianalyzer_system_alertemail":                             32,
		"fortianalyzer_fortianalyzer_system_alertevent":                             33,
		"fortianalyzer_fortianalyzer_system_autodelete":                             34,
		"fortianalyzer_fortianalyzer_system_autodelete_dlpfilesautodeletion":        35,
		"fortianalyzer_fortianalyzer_system_autodelete_logautodeletion":             36,
		"fortianalyzer_fortianalyzer_system_autodelete_quarantinefilesautodeletion": 37,
		"fortianalyzer_fortianalyzer_system_autodelete_reportautodeletion":          38,
		"fortianalyzer_fortianalyzer_system_backup_allsettings":                     39,
		"fortianalyzer_fortianalyzer_system_centralmanagement":                      40,
		"fortianalyzer_fortianalyzer_system_certificate_ca":                         41,
		"fortianalyzer_fortianalyzer_system_certificate_crl":                        42,
		"fortianalyzer_fortianalyzer_system_certificate_local":                      43,
		"fortianalyzer_fortianalyzer_system_certificate_oftp":                       44,
		"fortianalyzer_fortianalyzer_system_certificate_remote":                     45,
		"fortianalyzer_fortianalyzer_system_certificate_ssh":                        46,
		"fortianalyzer_fortianalyzer_system_connector":                              47,
		"fortianalyzer_fortianalyzer_system_dns":                                    48,
		"fortianalyzer_fortianalyzer_system_docker":                                 49,
		"fortianalyzer_fortianalyzer_system_fips":                                   50,
		"fortianalyzer_fortianalyzer_system_fortiview_autocache":                    51,
		"fortianalyzer_fortianalyzer_system_fortiview_setting":                      52,
		"fortianalyzer_fortianalyzer_system_global":                                 53,
		"fortianalyzer_fortianalyzer_system_global_sslciphersuites":                 54,
		"fortianalyzer_fortianalyzer_system_guiact":                                 55,
		"fortianalyzer_fortianalyzer_system_ha":                                     56,
		"fortianalyzer_fortianalyzer_system_ha_peer":                                57,
		"fortianalyzer_fortianalyzer_system_ha_privatepeer":                         58,
		"fortianalyzer_fortianalyzer_system_interface":                              59,
		"fortianalyzer_fortianalyzer_system_locallog_disk_filter":                   60,
		"fortianalyzer_fortianalyzer_system_locallog_disk_setting":                  61,
		"fortianalyzer_fortianalyzer_system_locallog_fortianalyzer2_filter":         62,
		"fortianalyzer_fortianalyzer_system_locallog_fortianalyzer2_setting":        63,
		"fortianalyzer_fortianalyzer_system_locallog_fortianalyzer3_filter":         64,
		"fortianalyzer_fortianalyzer_system_locallog_fortianalyzer3_setting":        65,
		"fortianalyzer_fortianalyzer_system_locallog_fortianalyzer_filter":          66,
		"fortianalyzer_fortianalyzer_system_locallog_fortianalyzer_setting":         67,
		"fortianalyzer_fortianalyzer_system_locallog_memory_filter":                 68,
		"fortianalyzer_fortianalyzer_system_locallog_memory_setting":                69,
		"fortianalyzer_fortianalyzer_system_locallog_setting":                       70,
		"fortianalyzer_fortianalyzer_system_locallog_syslogd2_filter":               71,
		"fortianalyzer_fortianalyzer_system_locallog_syslogd2_setting":              72,
		"fortianalyzer_fortianalyzer_system_locallog_syslogd3_filter":               73,
		"fortianalyzer_fortianalyzer_system_locallog_syslogd3_setting":              74,
		"fortianalyzer_fortianalyzer_system_locallog_syslogd_filter":                75,
		"fortianalyzer_fortianalyzer_system_locallog_syslogd_setting":               76,
		"fortianalyzer_fortianalyzer_system_log_alert":                              77,
		"fortianalyzer_fortianalyzer_system_log_devicedisable":                      78,
		"fortianalyzer_fortianalyzer_system_log_fospolicystats":                     79,
		"fortianalyzer_fortianalyzer_system_log_interfacestats":                     80,
		"fortianalyzer_fortianalyzer_system_log_ioc":                                81,
		"fortianalyzer_fortianalyzer_system_log_maildomain":                         82,
		"fortianalyzer_fortianalyzer_system_log_ratelimit":                          83,
		"fortianalyzer_fortianalyzer_system_log_ratelimit_device":                   84,
		"fortianalyzer_fortianalyzer_system_log_ratelimit_ratelimits":               85,
		"fortianalyzer_fortianalyzer_system_log_settings":                           86,
		"fortianalyzer_fortianalyzer_system_log_settings_rollinganalyzer":           87,
		"fortianalyzer_fortianalyzer_system_log_settings_rollinglocal":              88,
		"fortianalyzer_fortianalyzer_system_log_settings_rollingregular":            89,
		"fortianalyzer_fortianalyzer_system_log_topology":                           90,
		"fortianalyzer_fortianalyzer_system_logfetch_clientprofile":                 91,
		"fortianalyzer_fortianalyzer_system_logfetch_serversettings":                92,
		"fortianalyzer_fortianalyzer_system_logforward":                             93,
		"fortianalyzer_fortianalyzer_system_logforwardservice":                      94,
		"fortianalyzer_fortianalyzer_system_mail":                                   95,
		"fortianalyzer_fortianalyzer_system_metadata_admins":                        96,
		"fortianalyzer_fortianalyzer_system_ntp":                                    97,
		"fortianalyzer_fortianalyzer_system_ntp_ntpserver":                          98,
		"fortianalyzer_fortianalyzer_system_passwordpolicy":                         99,
		"fortianalyzer_fortianalyzer_system_report_autocache":                       100,
		"fortianalyzer_fortianalyzer_system_report_estbrowsetime":                   101,
		"fortianalyzer_fortianalyzer_system_report_setting":                         102,
		"fortianalyzer_fortianalyzer_system_route":                                  103,
		"fortianalyzer_fortianalyzer_system_route6":                                 104,
		"fortianalyzer_fortianalyzer_system_saml":                                   105,
		"fortianalyzer_fortianalyzer_system_saml_fabricidp":                         106,
		"fortianalyzer_fortianalyzer_system_saml_serviceproviders":                  107,
		"fortianalyzer_fortianalyzer_system_sniffer":                                108,
		"fortianalyzer_fortianalyzer_system_snmp_community":                         109,
		"fortianalyzer_fortianalyzer_system_snmp_sysinfo":                           110,
		"fortianalyzer_fortianalyzer_system_snmp_user":                              111,
		"fortianalyzer_fortianalyzer_system_socfabric":                              112,
		"fortianalyzer_fortianalyzer_system_sql":                                    113,
		"fortianalyzer_fortianalyzer_system_sql_customindex":                        114,
		"fortianalyzer_fortianalyzer_system_sql_customskipidx":                      115,
		"fortianalyzer_fortianalyzer_system_sql_tsindexfield":                       116,
		"fortianalyzer_fortianalyzer_system_syslog":                                 117,
		"fortianalyzer_fortianalyzer_system_webproxy":                               118,
		"fortianalyzer_fortianalyzer_system_workflow_approvalmatrix":                119,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
