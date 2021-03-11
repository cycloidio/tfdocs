package intersight

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "intersight_aaa_audit_record",
			Category:         "aaa",
			ShortDescription: `AuditRecord presents the configuration changes made by the user per transaction.`,
			Description: `
AuditRecord presents the configuration changes made by the user per transaction.
`,
			Keywords: []string{
				"aaa",
				"audit",
				"record",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_access_policy",
			Category:         "access",
			ShortDescription: `Policy to configure server or chassis management options.`,
			Description: `
Policy to configure server or chassis management options.
`,
			Keywords: []string{
				"access",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_config_policy",
			Category:         "adapter",
			ShortDescription: `An Adapter Configuration Policy configures the Ethernet and Fibre-Channel settings for the VIC adapter.`,
			Description: `
An Adapter Configuration Policy configures the Ethernet and Fibre-Channel settings for the VIC adapter.
`,
			Keywords: []string{
				"adapter",
				"config",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_ext_eth_interface",
			Category:         "adapter",
			ShortDescription: `Physical port of a virtual interface card.`,
			Description: `
Physical port of a virtual interface card.
`,
			Keywords: []string{
				"adapter",
				"ext",
				"eth",
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_host_eth_interface",
			Category:         "adapter",
			ShortDescription: `Physical / Virtual port of an adapter as seen by the host.`,
			Description: `
Physical / Virtual port of an adapter as seen by the host.
`,
			Keywords: []string{
				"adapter",
				"host",
				"eth",
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_host_fc_interface",
			Category:         "adapter",
			ShortDescription: `Host facing fibre channel interface on a server adapter.`,
			Description: `
Host facing fibre channel interface on a server adapter.
`,
			Keywords: []string{
				"adapter",
				"host",
				"fc",
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_host_iscsi_interface",
			Category:         "adapter",
			ShortDescription: `Iscsi interface on a server adapter.`,
			Description: `
Iscsi interface on a server adapter.
`,
			Keywords: []string{
				"adapter",
				"host",
				"iscsi",
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_unit",
			Category:         "adapter",
			ShortDescription: `The physical adapter present on a server.`,
			Description: `
The physical adapter present on a server.
`,
			Keywords: []string{
				"adapter",
				"unit",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_adapter_unit_expander",
			Category:         "adapter",
			ShortDescription: `The adapter unit extension card present on a server.`,
			Description: `
The adapter unit extension card present on a server.
`,
			Keywords: []string{
				"adapter",
				"unit",
				"expander",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_app_status",
			Category:         "appliance",
			ShortDescription: `Status of an application.`,
			Description: `
Status of an application.
`,
			Keywords: []string{
				"appliance",
				"app",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Operational status of the Intersight Appliance entity is Unknown.`,
				},
				resource.Attribute{
					Name:        "Operational",
					Description: `Operational status of the Intersight Appliance entity is Operational.`,
				},
				resource.Attribute{
					Name:        "Impaired",
					Description: `Operational status of the Intersight Appliance entity is Impaired.`,
				},
				resource.Attribute{
					Name:        "AttentionNeeded",
					Description: `Operational status of the Intersight Appliance entity is AttentionNeeded.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_backup",
			Category:         "appliance",
			ShortDescription: `Backup tracks all backup requests to create a full system backup of the Intersight Appliance. There will be only one Backup managed object with a 'Started' state at any time. All other Backup managed objects will be in terminal states.`,
			Description: `
Backup tracks all backup requests to create a full system backup of the Intersight
Appliance. There will be only one Backup managed object with a 'Started' state at
any time. All other Backup managed objects will be in terminal states.
`,
			Keywords: []string{
				"appliance",
				"backup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scp",
					Description: `Secure Copy Protocol (SCP) to access the file server.`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `SSH File Transfer Protocol (SFTP) to access file server.`,
				},
				resource.Attribute{
					Name:        "Started",
					Description: `Backup or restore process has started.`,
				},
				resource.Attribute{
					Name:        "Created",
					Description: `Backup or restore is in created state.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Backup or restore process has failed.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Backup or restore process has completed.`,
				},
				resource.Attribute{
					Name:        "Copied",
					Description: `Backup file has been copied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_backup_policy",
			Category:         "appliance",
			ShortDescription: `BackupPolicy stores the Intersight Appliance's backup policy. There will be only one BackupPolicy managed object in the Intersight Appliance. Default backup policy managed object is created during the Intersight Appliance setup, and it is configured in the manual backup mode.`,
			Description: `
BackupPolicy stores the Intersight Appliance's backup policy. There will be only
one BackupPolicy managed object in the Intersight Appliance. Default backup policy
managed object is created during the Intersight Appliance setup, and it is configured
in the manual backup mode.
`,
			Keywords: []string{
				"appliance",
				"backup",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scp",
					Description: `Secure Copy Protocol (SCP) to access the file server.`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `SSH File Transfer Protocol (SFTP) to access file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_certificate_setting",
			Category:         "appliance",
			ShortDescription: `Certificate the appliance uses for browser traffic.`,
			Description: `
Certificate the appliance uses for browser traffic.
`,
			Keywords: []string{
				"appliance",
				"certificate",
				"setting",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_data_export_policy",
			Category:         "appliance",
			ShortDescription: `Data Export Policy is a category-based data collection policy that enables or disables data export (data collection) from the Intersight Appliance to the Intersight. The Data Export Policy configuration is organized hierarchically as follows. Global: Inventory: Network Storage TechSupport When the DataExportPolicy for a category is enabled/disabled, all the sub-category configurations are enabled/disabled as well. For example, if you enable/disable Inventory, all its sub-category configurations (ie. Network and Storage) are also enabled/disabled.`,
			Description: `
Data Export Policy is a category-based data collection policy that enables or disables
data export (data collection) from the Intersight Appliance to the Intersight. The Data
Export Policy configuration is organized hierarchically as follows.
  Global:
     Inventory:
        Network
        Storage
     TechSupport
When the DataExportPolicy for a category is enabled/disabled, all the sub-category configurations
are enabled/disabled as well. For example, if you enable/disable Inventory, all its sub-category
configurations (ie. Network and Storage) are also enabled/disabled.
`,
			Keywords: []string{
				"appliance",
				"data",
				"export",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_device_certificate",
			Category:         "appliance",
			ShortDescription: `DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.`,
			Description: `
DeviceCertificate managed object stores the CA Certificate used by device connector and it allow tracks it renewal.
`,
			Keywords: []string{
				"appliance",
				"device",
				"certificate",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_device_claim",
			Category:         "appliance",
			ShortDescription: `DeviceClaim managed object represents a user initiated claim request for claiming an endpoint device. There can be many DeviceClaim managed object for a given endpoint device when users claim and unclaim devices repeatedly. Claiming an endpoint device is a multi-step operation. The Intersight Appliance starts a workflow with multiple tasks to process the device claim request. The status of the device claim operation can be obtained from the claim workflow.`,
			Description: `
DeviceClaim managed object represents a user initiated claim request for claiming
an endpoint device. There can be many DeviceClaim managed object for a given endpoint
device when users claim and unclaim devices repeatedly.
Claiming an endpoint device is a multi-step operation. The Intersight Appliance
starts a workflow with multiple tasks to process the device claim request. The status
of the device claim operation can be obtained from the claim workflow.
`,
			Keywords: []string{
				"appliance",
				"device",
				"claim",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "started",
					Description: `Device claim operation has started.`,
				},
				resource.Attribute{
					Name:        "failed",
					Description: `Device claim operation has failed.`,
				},
				resource.Attribute{
					Name:        "completed",
					Description: `Device claim operation has completed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_diag_setting",
			Category:         "appliance",
			ShortDescription: `DiagSetting model is used for changing the password of the operating system's diagnostic user account. The diagnostic user account can be used to login to the Intersight Appliance virtual machine. The diagnostic user account is protected by two separate authentication mechanisms: user's password and Cisco CT-engine generated key. Only the Intersight Appliance's local account administrator has the privileges to use this REST API.`,
			Description: `
DiagSetting model is used for changing the password of the operating system's diagnostic
user account. The diagnostic user account can be used to login to the Intersight Appliance
virtual machine.
The diagnostic user account is protected by two separate authentication mechanisms: user's
password and Cisco CT-engine generated key. Only the Intersight Appliance's local account
administrator has the privileges to use this REST API.
`,
			Keywords: []string{
				"appliance",
				"diag",
				"setting",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_external_syslog_setting",
			Category:         "appliance",
			ShortDescription: `Configure External Syslog Server.`,
			Description: `
Configure External Syslog Server.
`,
			Keywords: []string{
				"appliance",
				"external",
				"syslog",
				"setting",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_file_system_status",
			Category:         "appliance",
			ShortDescription: `Status of a file system on an Intersight Appliance node.`,
			Description: `
Status of a file system on an Intersight Appliance node.
`,
			Keywords: []string{
				"appliance",
				"file",
				"system",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Operational status of the Intersight Appliance entity is Unknown.`,
				},
				resource.Attribute{
					Name:        "Operational",
					Description: `Operational status of the Intersight Appliance entity is Operational.`,
				},
				resource.Attribute{
					Name:        "Impaired",
					Description: `Operational status of the Intersight Appliance entity is Impaired.`,
				},
				resource.Attribute{
					Name:        "AttentionNeeded",
					Description: `Operational status of the Intersight Appliance entity is AttentionNeeded.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_group_status",
			Category:         "appliance",
			ShortDescription: `Status of a group of applications.`,
			Description: `
Status of a group of applications.
`,
			Keywords: []string{
				"appliance",
				"group",
				"status",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_image_bundle",
			Category:         "appliance",
			ShortDescription: `ImageBundle keeps track of all the software bundles installed in the Intersight Appliances. Each ImageBundle managed object is derived from a software upgrade manifest. ImageBundle has additional properties computed during the manifest processing. Additional properties are the dynamic attributes of the software packages declared in the software manifest. For example, SHA256 values of the software packages are computed during the software manifest processing. An ImageBundle managed object named 'current' is always present in the Intersight Appliance. The software upgrade service creates another ImageBundle managed object named 'pending' when there is a pending software upgrade. The upgrade service renames the 'pending' bundle to the 'current' bundle after the software upgrade is successful.`,
			Description: `
ImageBundle keeps track of all the software bundles installed in the Intersight Appliances.
Each ImageBundle managed object is derived from a software upgrade manifest. ImageBundle has
additional properties computed during the manifest processing. Additional properties are the
dynamic attributes of the software packages declared in the software manifest. For example,
SHA256 values of the software packages are computed during the software manifest processing.
An ImageBundle managed object named 'current' is always present in the Intersight Appliance.
The software upgrade service creates another ImageBundle managed object named 'pending'
when there is a pending software upgrade. The upgrade service renames the 'pending' bundle
to the 'current' bundle after the software upgrade is successful.
`,
			Keywords: []string{
				"appliance",
				"image",
				"bundle",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Normal",
					Description: `Normal upgrade priority is used for all the software upgrades except for the critical security updates. The upgrade service of Intersight Appliance uses the Software Upgrade Policy settings to start the upgrade process.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Critical upgrade priority is used for critical updates such as security patches. The upgrade service of the Intersight Appliance starts the upgrade as specified by the upgrade properties in the software manifest file. The upgrade service will not use the settings specified in the Software Upgrade Policy.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The upgrade has no effect on the system.`,
				},
				resource.Attribute{
					Name:        "Disruptive",
					Description: `The services will not be functional during the upgrade.`,
				},
				resource.Attribute{
					Name:        "Disruptive-reboot",
					Description: `The upgrade needs a reboot.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_node_info",
			Category:         "appliance",
			ShortDescription: `NodeInfo managed object stores the Intersight Appliance's cluster node information. NodeInfo managed objects are created during the Intersight Appliance setup. The Intersight Appliance updates the NodeInfo managed objects with status information periodically.`,
			Description: `
NodeInfo managed object stores the Intersight Appliance's cluster node information.
NodeInfo managed objects are created during the Intersight Appliance setup. The
Intersight Appliance updates the NodeInfo managed objects with status information
periodically.
`,
			Keywords: []string{
				"appliance",
				"node",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Operational status of the Intersight Appliance entity is Unknown.`,
				},
				resource.Attribute{
					Name:        "Operational",
					Description: `Operational status of the Intersight Appliance entity is Operational.`,
				},
				resource.Attribute{
					Name:        "Impaired",
					Description: `Operational status of the Intersight Appliance entity is Impaired.`,
				},
				resource.Attribute{
					Name:        "AttentionNeeded",
					Description: `Operational status of the Intersight Appliance entity is AttentionNeeded.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_node_status",
			Category:         "appliance",
			ShortDescription: `Status of an Intersight Appliance node.`,
			Description: `
Status of an Intersight Appliance node.
`,
			Keywords: []string{
				"appliance",
				"node",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Down",
					Description: `The node is yet to come up and join as a member of theKubernetes cluster.`,
				},
				resource.Attribute{
					Name:        "Preparing",
					Description: `The node has come up and joined the Kubernetes cluster,preparing to host Kubernetes pods.`,
				},
				resource.Attribute{
					Name:        "Ready",
					Description: `The node is ready to host Kubernetes pods.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Operational status of the Intersight Appliance entity is Unknown.`,
				},
				resource.Attribute{
					Name:        "Operational",
					Description: `Operational status of the Intersight Appliance entity is Operational.`,
				},
				resource.Attribute{
					Name:        "Impaired",
					Description: `Operational status of the Intersight Appliance entity is Impaired.`,
				},
				resource.Attribute{
					Name:        "AttentionNeeded",
					Description: `Operational status of the Intersight Appliance entity is AttentionNeeded.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_release_note",
			Category:         "appliance",
			ShortDescription: `ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.`,
			Description: `
ReleaseUpdate managed the object provides the information preview (new features and bug fixes) for one pending upgrade.
`,
			Keywords: []string{
				"appliance",
				"release",
				"note",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_remote_file_import",
			Category:         "appliance",
			ShortDescription: `Trigger a remote file import request by configuring this mo.`,
			Description: `
Trigger a remote file import request by configuring this mo.
`,
			Keywords: []string{
				"appliance",
				"remote",
				"file",
				"import",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scp",
					Description: `Secure Copy Protocol (SCP) to access the file server.`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `SSH File Transfer Protocol (SFTP) to access file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_restore",
			Category:         "appliance",
			ShortDescription: `Restore tracks requests to restore the Intersight Appliance. There will be only one Restore managed object with a 'Started' state at any time. All other Restore managed objects will be in terminal states.`,
			Description: `
Restore tracks requests to restore the Intersight Appliance. There will be only
one Restore managed object with a 'Started' state at any time. All other Restore
managed objects will be in terminal states.
`,
			Keywords: []string{
				"appliance",
				"restore",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "scp",
					Description: `Secure Copy Protocol (SCP) to access the file server.`,
				},
				resource.Attribute{
					Name:        "sftp",
					Description: `SSH File Transfer Protocol (SFTP) to access file server.`,
				},
				resource.Attribute{
					Name:        "Started",
					Description: `Backup or restore process has started.`,
				},
				resource.Attribute{
					Name:        "Created",
					Description: `Backup or restore is in created state.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Backup or restore process has failed.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Backup or restore process has completed.`,
				},
				resource.Attribute{
					Name:        "Copied",
					Description: `Backup file has been copied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_setup_info",
			Category:         "appliance",
			ShortDescription: `SetupInfo will have only one managed object. SetupInfo managed object is to keep track of the Intersight Appliance's setup information and guide the UI through the initial configuration of the Intersight Appliance. The SetupInfo managed object is created during the Intersight Appliance setup. The Intersight UI uses this object to store the initial configuration states that the user has completed. If the user closes the Intersight UI without finishing all the initial configuration, then the Intersight UI will use this managed object to display the next configuration that the user needs to complete when the user uses the Intersight Appliance next time.`,
			Description: `
SetupInfo will have only one managed object. SetupInfo managed object is to keep
track of the Intersight Appliance's setup information and guide the UI through
the initial configuration of the Intersight Appliance.
The SetupInfo managed object is created during the Intersight Appliance setup.
The Intersight UI uses this object to store the initial configuration states
that the user has completed. If the user closes the Intersight UI without
finishing all the initial configuration, then the Intersight UI will use this
managed object to display the next configuration that the user needs to complete
when the user uses the Intersight Appliance next time.
`,
			Keywords: []string{
				"appliance",
				"setup",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Connected",
					Description: `In connected mode, Intersight Appliance connects to Intersight SaaS and other cisco.com services.`,
				},
				resource.Attribute{
					Name:        "Private",
					Description: `In private mode, Intersight Appliance does not connect to Intersight SaaS or any cisco.com services.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_system_info",
			Category:         "appliance",
			ShortDescription: `The Intersight Appliance's system information. SystemInfo is a singleton managed object created during the Intersight Appliance setup. The Intersight Appliance updates the SystemInfo managed object with up to date cluster status information periodically.`,
			Description: `
The Intersight Appliance's system information. SystemInfo is a singleton managed object
created during the Intersight Appliance setup. The Intersight Appliance updates the
SystemInfo managed object with up to date cluster status information periodically.
`,
			Keywords: []string{
				"appliance",
				"system",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Connected",
					Description: `Intersight is able to establish a connection to the target and initiate management activities.`,
				},
				resource.Attribute{
					Name:        "NotConnected",
					Description: `Intersight is unable to establish a connection to the target.`,
				},
				resource.Attribute{
					Name:        "ClaimInProgress",
					Description: `Claim of the target is in progress. A connection to the target has not been fully established.`,
				},
				resource.Attribute{
					Name:        "Unclaimed",
					Description: `The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.`,
				},
				resource.Attribute{
					Name:        "Claimed",
					Description: `Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Operational status of the Intersight Appliance entity is Unknown.`,
				},
				resource.Attribute{
					Name:        "Operational",
					Description: `Operational status of the Intersight Appliance entity is Operational.`,
				},
				resource.Attribute{
					Name:        "Impaired",
					Description: `Operational status of the Intersight Appliance entity is Impaired.`,
				},
				resource.Attribute{
					Name:        "AttentionNeeded",
					Description: `Operational status of the Intersight Appliance entity is AttentionNeeded.`,
				},
				resource.Attribute{
					Name:        "Pacific/Kiritimati",
					Description: ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_system_status",
			Category:         "appliance",
			ShortDescription: `Status of the Intersight Appliance.`,
			Description: `
Status of the Intersight Appliance.
`,
			Keywords: []string{
				"appliance",
				"system",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Operational status of the Intersight Appliance entity is Unknown.`,
				},
				resource.Attribute{
					Name:        "Operational",
					Description: `Operational status of the Intersight Appliance entity is Operational.`,
				},
				resource.Attribute{
					Name:        "Impaired",
					Description: `Operational status of the Intersight Appliance entity is Impaired.`,
				},
				resource.Attribute{
					Name:        "AttentionNeeded",
					Description: `Operational status of the Intersight Appliance entity is AttentionNeeded.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_upgrade",
			Category:         "appliance",
			ShortDescription: `Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's upgrade service automatically creates an Upgrade managed object when there is a pending software upgrade. User may modify an active Upgrade managed object to reset the software upgrade start time. However, user may not be able to postpone an upgrade beyond the limit enforced by the upgrade grace period set in the software manifest.`,
			Description: `
Upgrade tracks the Intersight Appliance's software upgrades. Intersight Appliance's
upgrade service automatically creates an Upgrade managed object when there is a
pending software upgrade. User may modify an active Upgrade managed object to reset
the software upgrade start time. However, user may not be able to postpone an upgrade
beyond the limit enforced by the upgrade grace period set in the software manifest.
`,
			Keywords: []string{
				"appliance",
				"upgrade",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_appliance_upgrade_policy",
			Category:         "appliance",
			ShortDescription: `UpgradePolicy stores the Intersight Appliance's software upgrade policy. UpgradePolicy is a sinlgeton managed object. A default upgrade policy is created during the Intersight Appliance setup, and it is configured with an automatic upgrade policy. Automatic upgrade policy lets the system start software upgrade after a pre-defined number of seconds set in the software manifest. Scheduled upgrade policy lets the user start software upgrade at a specified schedule. However, scheduled time cannot be beyond the time limit enforced by the upgrade grace period set in the software manifest.`,
			Description: `
UpgradePolicy stores the Intersight Appliance's software upgrade policy. UpgradePolicy
is a sinlgeton managed object. A default upgrade policy is created during the Intersight
Appliance setup, and it is configured with an automatic upgrade policy.
Automatic upgrade policy lets the system start software upgrade after a pre-defined
number of seconds set in the software manifest.
Scheduled upgrade policy lets the user start software upgrade at a specified schedule.
However, scheduled time cannot be beyond the time limit enforced by the upgrade grace
period set in the software manifest.
`,
			Keywords: []string{
				"appliance",
				"upgrade",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_cluster_member",
			Category:         "asset",
			ShortDescription: `A node within a cluster of device connectors. A Device Registration may contain multiple ClusterMembers with each holding the connection details of the device connector as well as the nodes current leadership within the cluster.`,
			Description: `
A node within a cluster of device connectors. A Device Registration may contain multiple ClusterMembers with each holding the connection details of the device connector as well as the nodes current leadership within the cluster.
`,
			Keywords: []string{
				"asset",
				"cluster",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Connected",
					Description: `Intersight is able to establish a connection to the target and initiate management activities.`,
				},
				resource.Attribute{
					Name:        "NotConnected",
					Description: `Intersight is unable to establish a connection to the target.`,
				},
				resource.Attribute{
					Name:        "ClaimInProgress",
					Description: `Claim of the target is in progress. A connection to the target has not been fully established.`,
				},
				resource.Attribute{
					Name:        "Unclaimed",
					Description: `The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.`,
				},
				resource.Attribute{
					Name:        "Claimed",
					Description: `Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The node is unable to complete election or determine the current state. If the device has been registered before and the node has access to the current credentials it will establish a connection to Intersight with limited capabilities that can be used to debug the HA failure from Intersight.`,
				},
				resource.Attribute{
					Name:        "Primary",
					Description: `The node has been elected as the primary and will establish a connection to the Intersight service and accept all message types enabled for a primary node. There can only be one primary in a given cluster, while the underlying platform may be active-active only one connector will assume the primary role.`,
				},
				resource.Attribute{
					Name:        "Secondary",
					Description: `The node has been elected as a secondary node in the cluster. The device connector will establish a connection to the Intersight service with limited capabilities. e.g. file upload will be enabled, but requests to the underlying platform management will be disabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_device_configuration",
			Category:         "asset",
			ShortDescription: `The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.`,
			Description: `
The configuration of a device connector. Configuration properties may be changed by a Intersight user or by a device administrator using the connector's API exposed through the platforms management interface.
`,
			Keywords: []string{
				"asset",
				"device",
				"configuration",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_device_connector_manager",
			Category:         "asset",
			ShortDescription: `Information pertaining to a Registered Intersight Assist Appliance Device in Intersight.`,
			Description: `
Information pertaining to a Registered Intersight Assist Appliance Device in Intersight.
`,
			Keywords: []string{
				"asset",
				"device",
				"connector",
				"manager",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_device_contract_information",
			Category:         "asset",
			ShortDescription: `Contains information about the Cisco device identified by a unique identifier like serial number. It also contains information about warranty, contract status, validity of the device. In future this object could be expanded to store Case, RMA, device topology details. We observe new asset.DeviceRegisteration and use it as a trigger for creating an instance of this object. Currently the data is restricted to Cisco Standalone IMC servers and Fabric Interconnects. Support for more product lines will be added in future.`,
			Description: `
Contains information about the Cisco device identified by a unique identifier like serial number. It also contains information about warranty, contract status, validity of the device. In future this object could be expanded to store Case, RMA, device topology details. We observe new asset.DeviceRegisteration and use it as a trigger for creating an instance of this object. Currently the data is restricted to Cisco Standalone IMC servers and Fabric Interconnects. Support for more product lines will be added in future.
`,
			Keywords: []string{
				"asset",
				"device",
				"contract",
				"information",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `The device's contract status cannot be determined.`,
				},
				resource.Attribute{
					Name:        "Not Covered",
					Description: `The Cisco device does not have a valid support contract.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `The Cisco device is covered under a active support contract.`,
				},
				resource.Attribute{
					Name:        "Expiring Soon",
					Description: `The contract for this Cisco device is going to expire in the next 30 days.`,
				},
				resource.Attribute{
					Name:        "Line Item Expired",
					Description: `The Cisco device does not have a valid support contract, it has expired.`,
				},
				resource.Attribute{
					Name:        "Line Item Terminated",
					Description: `The Cisco device does not have a valid support contract, it has been terminated.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `A default value to catch cases where device type is not correctly detected.`,
				},
				resource.Attribute{
					Name:        "CiscoUcsServer",
					Description: `A device of type server. It includes Cisco IMC and UCS Managed servers.`,
				},
				resource.Attribute{
					Name:        "CiscoUcsFI",
					Description: `A device of type Fabric Interconnect. It includes the various types of Cisco Fabric Interconnects supported by Cisco Intersight.`,
				},
				resource.Attribute{
					Name:        "CiscoUcsChassis",
					Description: `A device of type Chassis. It includes various UCS chassis supported by Cisco Intersight.`,
				},
				resource.Attribute{
					Name:        "CiscoNexusSwitch",
					Description: `A device of type Nexus switch. It includes various Nexus switches supported by Cisco Intersight.`,
				},
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "Update",
					Description: `Sn2Info/Contract information needs to be updated.`,
				},
				resource.Attribute{
					Name:        "OK",
					Description: `Sn2Info/Contract information was fetched succcessfuly and updated.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Sn2Info/Contract information was not available or failed while fetching.`,
				},
				resource.Attribute{
					Name:        "Retry",
					Description: `Sn2Info/Contract information update failed and will be retried later.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_device_registration",
			Category:         "asset",
			ShortDescription: `DeviceRegistration represents a device connector enabled endpoint which has registered with Intersight.`,
			Description: `
DeviceRegistration represents a device connector enabled endpoint which has registered with Intersight.
`,
			Keywords: []string{
				"asset",
				"device",
				"registration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Connected",
					Description: `Intersight is able to establish a connection to the target and initiate management activities.`,
				},
				resource.Attribute{
					Name:        "NotConnected",
					Description: `Intersight is unable to establish a connection to the target.`,
				},
				resource.Attribute{
					Name:        "ClaimInProgress",
					Description: `Claim of the target is in progress. A connection to the target has not been fully established.`,
				},
				resource.Attribute{
					Name:        "Unclaimed",
					Description: `The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.`,
				},
				resource.Attribute{
					Name:        "Claimed",
					Description: `Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.`,
				},
				resource.Attribute{
					Name:        "Normal",
					Description: `The device connector is running in normal mode, i.e. it is not a simulation.`,
				},
				resource.Attribute{
					Name:        "Emulator",
					Description: `The device connector is running in simulation mode inside an emulated device.`,
				},
				resource.Attribute{
					Name:        "ContainerEmulator",
					Description: `The device connector is running in simulation mode inside a containerized emulated device.`,
				},
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_subscription_device_contract_information",
			Category:         "asset",
			ShortDescription: `Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status and customer address, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.`,
			Description: `
Contains information about Cisco devices associated with consumption-based subscriptions. In addition to device installation status and customer address, information about returns and replacements is also recorded here. We listen to messages sent by Cisco Install Base and create/update an instance of this object.
`,
			Keywords: []string{
				"asset",
				"subscription",
				"device",
				"contract",
				"information",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_asset_target",
			Category:         "asset",
			ShortDescription: `Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.`,
			Description: `
Target represents an entity which can be managed by Intersight. This includes physical entities like UCS and HyperFlex servers and software entities like VMware vCenter and Microsoft Azure cloud accounts.
`,
			Keywords: []string{
				"asset",
				"target",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Connected",
					Description: `Intersight is able to establish a connection to the target and initiate management activities.`,
				},
				resource.Attribute{
					Name:        "NotConnected",
					Description: `Intersight is unable to establish a connection to the target.`,
				},
				resource.Attribute{
					Name:        "ClaimInProgress",
					Description: `Claim of the target is in progress. A connection to the target has not been fully established.`,
				},
				resource.Attribute{
					Name:        "Unclaimed",
					Description: `The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.`,
				},
				resource.Attribute{
					Name:        "Claimed",
					Description: `Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.`,
				},
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_bios_boot_device",
			Category:         "bios",
			ShortDescription: `Actual boot devices of the system as enumerated by BIOS.`,
			Description: `
Actual boot devices of the system as enumerated by BIOS.
`,
			Keywords: []string{
				"bios",
				"boot",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_bios_boot_mode",
			Category:         "bios",
			ShortDescription: `The mode through which bios has booted.`,
			Description: `
The mode through which bios has booted.
`,
			Keywords: []string{
				"bios",
				"boot",
				"mode",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_bios_policy",
			Category:         "bios",
			ShortDescription: `Policy for setting BIOS tokens on the endpoint.`,
			Description: `
Policy for setting BIOS tokens on the endpoint.
`,
			Keywords: []string{
				"bios",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1500-m",
					Description: `Value - 1500-m for configuring Altitude token.`,
				},
				resource.Attribute{
					Name:        "300-m",
					Description: `Value - 300-m for configuring Altitude token.`,
				},
				resource.Attribute{
					Name:        "3000-m",
					Description: `Value - 3000-m for configuring Altitude token.`,
				},
				resource.Attribute{
					Name:        "900-m",
					Description: `Value - 900-m for configuring Altitude token.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring Altitude token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring AspmSupport token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring AspmSupport token.`,
				},
				resource.Attribute{
					Name:        "Force L0s",
					Description: `Value - Force L0s for configuring AspmSupport token.`,
				},
				resource.Attribute{
					Name:        "L1 Only",
					Description: `Value - L1 Only for configuring AspmSupport token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "115200",
					Description: `Value - 115200 for configuring BaudRate token.`,
				},
				resource.Attribute{
					Name:        "19200",
					Description: `Value - 19200 for configuring BaudRate token.`,
				},
				resource.Attribute{
					Name:        "38400",
					Description: `Value - 38400 for configuring BaudRate token.`,
				},
				resource.Attribute{
					Name:        "57600",
					Description: `Value - 57600 for configuring BaudRate token.`,
				},
				resource.Attribute{
					Name:        "9600",
					Description: `Value - 9600 for configuring BaudRate token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "13",
					Description: `Value - 13 for configuring BootOptionNumRetry token.`,
				},
				resource.Attribute{
					Name:        "5",
					Description: `Value - 5 for configuring BootOptionNumRetry token.`,
				},
				resource.Attribute{
					Name:        "Infinite",
					Description: `Value - Infinite for configuring BootOptionNumRetry token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "15",
					Description: `Value - 15 for configuring BootOptionReCoolDown token.`,
				},
				resource.Attribute{
					Name:        "45",
					Description: `Value - 45 for configuring BootOptionReCoolDown token.`,
				},
				resource.Attribute{
					Name:        "90",
					Description: `Value - 90 for configuring BootOptionReCoolDown token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Max Efficient",
					Description: `Value - Max Efficient for configuring BootPerformanceMode token.`,
				},
				resource.Attribute{
					Name:        "Max Performance",
					Description: `Value - Max Performance for configuring BootPerformanceMode token.`,
				},
				resource.Attribute{
					Name:        "Set by Intel NM",
					Description: `Value - Set by Intel NM for configuring BootPerformanceMode token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnCpuCpb token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnCpuCpb token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "FOUR (2 + 2)",
					Description: `Value - FOUR (2 + 2) for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "FOUR (4 + 0)",
					Description: `Value - FOUR (4 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "SIX (3 + 3)",
					Description: `Value - SIX (3 + 3) for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "THREE (3 + 0)",
					Description: `Value - THREE (3 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "TWO (1 + 1)",
					Description: `Value - TWO (1 + 1) for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "TWO (2 + 0)",
					Description: `Value - TWO (2 + 0) for configuring CbsCmnCpuGenDowncoreCtrl token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnCpuGlobalCstateCtrl token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnCpuGlobalCstateCtrl token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring CbsCmnCpuGlobalCstateCtrl token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnCpuL1streamHwPrefetcher token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnCpuL1streamHwPrefetcher token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring CbsCmnCpuL1streamHwPrefetcher token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnCpuL2streamHwPrefetcher token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnCpuL2streamHwPrefetcher token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring CbsCmnCpuL2streamHwPrefetcher token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnDeterminismSlider token.`,
				},
				resource.Attribute{
					Name:        "Performance",
					Description: `Value - Performance for configuring CbsCmnDeterminismSlider token.`,
				},
				resource.Attribute{
					Name:        "Power",
					Description: `Value - Power for configuring CbsCmnDeterminismSlider token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnGnbNbIommu token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnGnbNbIommu token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring CbsCmnGnbNbIommu token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring CbsCmnMemCtrlBankGroupSwapDdr4 token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmnMemMapBankInterleaveDdr4 token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CbsCmnMemMapBankInterleaveDdr4 token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsCmncTdpCtl token.`,
				},
				resource.Attribute{
					Name:        "Manual",
					Description: `Value - Manual for configuring CbsCmncTdpCtl token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsDfCmnMemIntlv token.`,
				},
				resource.Attribute{
					Name:        "Channel",
					Description: `Value - Channel for configuring CbsDfCmnMemIntlv token.`,
				},
				resource.Attribute{
					Name:        "Die",
					Description: `Value - Die for configuring CbsDfCmnMemIntlv token.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Value - None for configuring CbsDfCmnMemIntlv token.`,
				},
				resource.Attribute{
					Name:        "Socket",
					Description: `Value - Socket for configuring CbsDfCmnMemIntlv token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1 KB",
					Description: `Value - 1 KiB for configuring CbsDfCmnMemIntlvSize token.`,
				},
				resource.Attribute{
					Name:        "2 KB",
					Description: `Value - 2 KiB for configuring CbsDfCmnMemIntlvSize token.`,
				},
				resource.Attribute{
					Name:        "256 Bytes",
					Description: `Value - 256 Bytes for configuring CbsDfCmnMemIntlvSize token.`,
				},
				resource.Attribute{
					Name:        "512 Bytes",
					Description: `Value - 512 Bytes for configuring CbsDfCmnMemIntlvSize token.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CbsDfCmnMemIntlvSize token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CdnSupport token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring CdnSupport token.`,
				},
				resource.Attribute{
					Name:        "LOMs Only",
					Description: `Value - LOMs Only for configuring CdnSupport token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1-way",
					Description: `Value - 1-way for configuring ChannelInterLeave token.`,
				},
				resource.Attribute{
					Name:        "2-way",
					Description: `Value - 2-way for configuring ChannelInterLeave token.`,
				},
				resource.Attribute{
					Name:        "3-way",
					Description: `Value - 3-way for configuring ChannelInterLeave token.`,
				},
				resource.Attribute{
					Name:        "4-way",
					Description: `Value - 4-way for configuring ChannelInterLeave token.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring ChannelInterLeave token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Maximum",
					Description: `Value - Maximum for configuring CiscoDebugLevel token.`,
				},
				resource.Attribute{
					Name:        "Minimum",
					Description: `Value - Minimum for configuring CiscoDebugLevel token.`,
				},
				resource.Attribute{
					Name:        "Normal",
					Description: `Value - Normal for configuring CiscoDebugLevel token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring CkeLowPolicy token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CkeLowPolicy token.`,
				},
				resource.Attribute{
					Name:        "fast",
					Description: `Value - fast for configuring CkeLowPolicy token.`,
				},
				resource.Attribute{
					Name:        "slow",
					Description: `Value - slow for configuring CkeLowPolicy token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Level 1",
					Description: `Value - Level 1 for configuring ConfigTdpLevel token.`,
				},
				resource.Attribute{
					Name:        "Level 2",
					Description: `Value - Level 2 for configuring ConfigTdpLevel token.`,
				},
				resource.Attribute{
					Name:        "Normal",
					Description: `Value - Normal for configuring ConfigTdpLevel token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "com-0",
					Description: `Value - com-0 for configuring ConsoleRedirection token.`,
				},
				resource.Attribute{
					Name:        "com-1",
					Description: `Value - com-1 for configuring ConsoleRedirection token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring ConsoleRedirection token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring ConsoleRedirection token.`,
				},
				resource.Attribute{
					Name:        "serial-port-a",
					Description: `Value - serial-port-a for configuring ConsoleRedirection token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1",
					Description: `Value - 1 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "10",
					Description: `Value - 10 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "11",
					Description: `Value - 11 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "12",
					Description: `Value - 12 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "13",
					Description: `Value - 13 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "14",
					Description: `Value - 14 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "15",
					Description: `Value - 15 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "16",
					Description: `Value - 16 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "17",
					Description: `Value - 17 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "18",
					Description: `Value - 18 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "19",
					Description: `Value - 19 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "2",
					Description: `Value - 2 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "20",
					Description: `Value - 20 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "21",
					Description: `Value - 21 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "22",
					Description: `Value - 22 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "23",
					Description: `Value - 23 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "24",
					Description: `Value - 24 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "25",
					Description: `Value - 25 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "26",
					Description: `Value - 26 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "27",
					Description: `Value - 27 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "28",
					Description: `Value - 28 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "3",
					Description: `Value - 3 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "4",
					Description: `Value - 4 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "5",
					Description: `Value - 5 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "6",
					Description: `Value - 6 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "7",
					Description: `Value - 7 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "8",
					Description: `Value - 8 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "9",
					Description: `Value - 9 for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "all",
					Description: `Value - all for configuring CoreMultiProcessing token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "balanced-energy",
					Description: `Value - balanced-energy for configuring CpuEnergyPerformance token.`,
				},
				resource.Attribute{
					Name:        "balanced-performance",
					Description: `Value - balanced-performance for configuring CpuEnergyPerformance token.`,
				},
				resource.Attribute{
					Name:        "balanced-power",
					Description: `Value - balanced-power for configuring CpuEnergyPerformance token.`,
				},
				resource.Attribute{
					Name:        "energy-efficient",
					Description: `Value - energy-efficient for configuring CpuEnergyPerformance token.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `Value - performance for configuring CpuEnergyPerformance token.`,
				},
				resource.Attribute{
					Name:        "power",
					Description: `Value - power for configuring CpuEnergyPerformance token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `Value - custom for configuring CpuPerformance token.`,
				},
				resource.Attribute{
					Name:        "enterprise",
					Description: `Value - enterprise for configuring CpuPerformance token.`,
				},
				resource.Attribute{
					Name:        "high-throughput",
					Description: `Value - high-throughput for configuring CpuPerformance token.`,
				},
				resource.Attribute{
					Name:        "hpc",
					Description: `Value - hpc for configuring CpuPerformance token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `Value - custom for configuring CpuPowerManagement token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring CpuPowerManagement token.`,
				},
				resource.Attribute{
					Name:        "energy-efficient",
					Description: `Value - energy-efficient for configuring CpuPowerManagement token.`,
				},
				resource.Attribute{
					Name:        "performance",
					Description: `Value - performance for configuring CpuPowerManagement token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring CrQos token.`,
				},
				resource.Attribute{
					Name:        "Recipe 1",
					Description: `Value - Recipe 1 for configuring CrQos token.`,
				},
				resource.Attribute{
					Name:        "Recipe 2",
					Description: `Value - Recipe 2 for configuring CrQos token.`,
				},
				resource.Attribute{
					Name:        "Recipe 3",
					Description: `Value - Recipe 3 for configuring CrQos token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "Default",
					Description: `Value - Default for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "Option 1",
					Description: `Value - Option 1 for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "Option 2",
					Description: `Value - Option 2 for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "Option 3",
					Description: `Value - Option 3 for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "Option 4",
					Description: `Value - Option 4 for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "Option 5",
					Description: `Value - Option 5 for configuring CrfastgoConfig token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring DirectCacheAccess token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring DirectCacheAccess token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring DirectCacheAccess token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring DramClockThrottling token.`,
				},
				resource.Attribute{
					Name:        "Balanced",
					Description: `Value - Balanced for configuring DramClockThrottling token.`,
				},
				resource.Attribute{
					Name:        "Energy Efficient",
					Description: `Value - Energy Efficient for configuring DramClockThrottling token.`,
				},
				resource.Attribute{
					Name:        "Performance",
					Description: `Value - Performance for configuring DramClockThrottling token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1x",
					Description: `Value - 1x for configuring DramRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "2x",
					Description: `Value - 2x for configuring DramRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "3x",
					Description: `Value - 3x for configuring DramRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "4x",
					Description: `Value - 4x for configuring DramRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring DramRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring DramSwThermalThrottling token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring DramSwThermalThrottling token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `Value - BIOS for configuring EngPerfTuning token.`,
				},
				resource.Attribute{
					Name:        "OS",
					Description: `Value - OS for configuring EngPerfTuning token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Balanced Performance",
					Description: `Value - Balanced Performance for configuring EppProfile token.`,
				},
				resource.Attribute{
					Name:        "Balanced Power",
					Description: `Value - Balanced Power for configuring EppProfile token.`,
				},
				resource.Attribute{
					Name:        "Performance",
					Description: `Value - Performance for configuring EppProfile token.`,
				},
				resource.Attribute{
					Name:        "Power",
					Description: `Value - Power for configuring EppProfile token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring ExtendedApic token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring ExtendedApic token.`,
				},
				resource.Attribute{
					Name:        "X2APIC",
					Description: `Value - X2APIC for configuring ExtendedApic token.`,
				},
				resource.Attribute{
					Name:        "XAPIC",
					Description: `Value - XAPIC for configuring ExtendedApic token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `Value - none for configuring FlowControl token.`,
				},
				resource.Attribute{
					Name:        "rts-cts",
					Description: `Value - rts-cts for configuring FlowControl token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring HwpmEnable token.`,
				},
				resource.Attribute{
					Name:        "HWPM Native Mode",
					Description: `Value - HWPM Native Mode for configuring HwpmEnable token.`,
				},
				resource.Attribute{
					Name:        "HWPM OOB Mode",
					Description: `Value - HWPM OOB Mode for configuring HwpmEnable token.`,
				},
				resource.Attribute{
					Name:        "NATIVE MODE",
					Description: `Value - NATIVE MODE for configuring HwpmEnable token.`,
				},
				resource.Attribute{
					Name:        "Native Mode with no Legacy",
					Description: `Value - Native Mode with no Legacy for configuring HwpmEnable token.`,
				},
				resource.Attribute{
					Name:        "OOB MODE",
					Description: `Value - OOB MODE for configuring HwpmEnable token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1-way Interleave",
					Description: `Value - 1-way Interleave for configuring ImcInterleave token.`,
				},
				resource.Attribute{
					Name:        "2-way Interleave",
					Description: `Value - 2-way Interleave for configuring ImcInterleave token.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring ImcInterleave token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Base",
					Description: `Value - Base for configuring IntelSpeedSelect token.`,
				},
				resource.Attribute{
					Name:        "Config 1",
					Description: `Value - Config 1 for configuring IntelSpeedSelect token.`,
				},
				resource.Attribute{
					Name:        "Config 2",
					Description: `Value - Config 2 for configuring IntelSpeedSelect token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "No",
					Description: `Value - No for configuring IohErrorEnable token.`,
				},
				resource.Attribute{
					Name:        "Yes",
					Description: `Value - Yes for configuring IohErrorEnable token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "IOH0 24k IOH1 40k",
					Description: `Value - IOH0 24k IOH1 40k for configuring IohResource token.`,
				},
				resource.Attribute{
					Name:        "IOH0 32k IOH1 32k",
					Description: `Value - IOH0 32k IOH1 32k for configuring IohResource token.`,
				},
				resource.Attribute{
					Name:        "IOH0 40k IOH1 24k",
					Description: `Value - IOH0 40k IOH1 24k for configuring IohResource token.`,
				},
				resource.Attribute{
					Name:        "IOH0 48k IOH1 16k",
					Description: `Value - IOH0 48k IOH1 16k for configuring IohResource token.`,
				},
				resource.Attribute{
					Name:        "IOH0 56k IOH1 8k",
					Description: `Value - IOH0 56k IOH1 8k for configuring IohResource token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring LegacyUsbSupport token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring LegacyUsbSupport token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring LegacyUsbSupport token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring LomPort0state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring LomPort0state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring LomPort0state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring LomPort0state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring LomPort1state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring LomPort1state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring LomPort1state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring LomPort1state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring LomPort2state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring LomPort2state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring LomPort2state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring LomPort2state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring LomPort3state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring LomPort3state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring LomPort3state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring LomPort3state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring LvDdrMode token.`,
				},
				resource.Attribute{
					Name:        "performance-mode",
					Description: `Value - performance-mode for configuring LvDdrMode token.`,
				},
				resource.Attribute{
					Name:        "power-saving-mode",
					Description: `Value - power-saving-mode for configuring LvDdrMode token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1 Way Node Interleave",
					Description: `Value - 1 Way Node Interleave for configuring MemoryInterLeave token.`,
				},
				resource.Attribute{
					Name:        "2 Way Node Interleave",
					Description: `Value - 2 Way Node Interleave for configuring MemoryInterLeave token.`,
				},
				resource.Attribute{
					Name:        "4 Way Node Interleave",
					Description: `Value - 4 Way Node Interleave for configuring MemoryInterLeave token.`,
				},
				resource.Attribute{
					Name:        "8 Way Node Interleave",
					Description: `Value - 8 Way Node Interleave for configuring MemoryInterLeave token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring MemoryInterLeave token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring MemoryInterLeave token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1x Refresh",
					Description: `Value - 1x Refresh for configuring MemoryRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "2x Refresh",
					Description: `Value - 2x Refresh for configuring MemoryRefreshRate token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "CLTT with PECI",
					Description: `Value - CLTT with PECI for configuring MemoryThermalThrottling token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring MemoryThermalThrottling token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "inter-socket",
					Description: `Value - inter-socket for configuring MirroringMode token.`,
				},
				resource.Attribute{
					Name:        "intra-socket",
					Description: `Value - intra-socket for configuring MirroringMode token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1 GB",
					Description: `Value - 1 GiB for configuring MmcfgBase token.`,
				},
				resource.Attribute{
					Name:        "2 GB",
					Description: `Value - 2 GiB for configuring MmcfgBase token.`,
				},
				resource.Attribute{
					Name:        "2.5 GB",
					Description: `Value - 2.5 GiB for configuring MmcfgBase token.`,
				},
				resource.Attribute{
					Name:        "3 GB",
					Description: `Value - 3 GiB for configuring MmcfgBase token.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring MmcfgBase token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "BW Optimized",
					Description: `Value - BW Optimized for configuring NvmdimmPerformConfig token.`,
				},
				resource.Attribute{
					Name:        "Balanced Profile",
					Description: `Value - Balanced Profile for configuring NvmdimmPerformConfig token.`,
				},
				resource.Attribute{
					Name:        "Latency Optimized",
					Description: `Value - Latency Optimized for configuring NvmdimmPerformConfig token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Intel RSTe",
					Description: `Value - Intel RSTe for configuring OnboardScuStorageSwStack token.`,
				},
				resource.Attribute{
					Name:        "LSI SW RAID",
					Description: `Value - LSI SW RAID for configuring OnboardScuStorageSwStack token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "do-nothing",
					Description: `Value - do-nothing for configuring OsBootWatchdogTimerPolicy token.`,
				},
				resource.Attribute{
					Name:        "power-off",
					Description: `Value - power-off for configuring OsBootWatchdogTimerPolicy token.`,
				},
				resource.Attribute{
					Name:        "reset",
					Description: `Value - reset for configuring OsBootWatchdogTimerPolicy token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "10-minutes",
					Description: `Value - 10-minutes for configuring OsBootWatchdogTimerTimeout token.`,
				},
				resource.Attribute{
					Name:        "15-minutes",
					Description: `Value - 15-minutes for configuring OsBootWatchdogTimerTimeout token.`,
				},
				resource.Attribute{
					Name:        "20-minutes",
					Description: `Value - 20-minutes for configuring OsBootWatchdogTimerTimeout token.`,
				},
				resource.Attribute{
					Name:        "5-minutes",
					Description: `Value - 5-minutes for configuring OsBootWatchdogTimerTimeout token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "C0 C1 State",
					Description: `Value - C0 C1 State for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "C0/C1",
					Description: `Value - C0/C1 for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "C2",
					Description: `Value - C2 for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "C6 Non Retention",
					Description: `Value - C6 Non Retention for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "C6 Retention",
					Description: `Value - C6 Retention for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "No Limit",
					Description: `Value - No Limit for configuring PackageCstateLimit token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "High",
					Description: `Value - High for configuring PanicHighWatermark token.`,
				},
				resource.Attribute{
					Name:        "Low",
					Description: `Value - Low for configuring PanicHighWatermark token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring PartialMirrorModeConfig token.`,
				},
				resource.Attribute{
					Name:        "Percentage",
					Description: `Value - Percentage for configuring PartialMirrorModeConfig token.`,
				},
				resource.Attribute{
					Name:        "Value in GB",
					Description: `Value - Value in GiB for configuring PartialMirrorModeConfig token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring PciOptionRoMs token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring PciOptionRoMs token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring PciOptionRoMs token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring PciOptionRoMs token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieAriSupport token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring PcieAriSupport token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring PcieAriSupport token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PciePllSsc token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PciePllSsc token.`,
				},
				resource.Attribute{
					Name:        "ZeroPointFive",
					Description: `Value - ZeroPointFive for configuring PciePllSsc token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieSlotNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PcieSlotNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring PcieSlotNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring PcieSlotNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring PcieSlotNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieSlotNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PcieSlotNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring PcieSlotNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring PcieSlotNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring PcieSlotNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieSlotNvme3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PcieSlotNvme3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring PcieSlotNvme3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring PcieSlotNvme3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring PcieSlotNvme3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieSlotNvme4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PcieSlotNvme4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring PcieSlotNvme4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring PcieSlotNvme4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring PcieSlotNvme4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieSlotNvme5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PcieSlotNvme5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring PcieSlotNvme5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring PcieSlotNvme5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring PcieSlotNvme5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring PcieSlotNvme6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring PcieSlotNvme6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring PcieSlotNvme6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring PcieSlotNvme6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring PcieSlotNvme6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "AHCI",
					Description: `Value - AHCI for configuring Psata token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Psata token.`,
				},
				resource.Attribute{
					Name:        "LSI SW RAID",
					Description: `Value - LSI SW RAID for configuring Psata token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "HW ALL",
					Description: `Value - HW ALL for configuring PstateCoordType token.`,
				},
				resource.Attribute{
					Name:        "SW ALL",
					Description: `Value - SW ALL for configuring PstateCoordType token.`,
				},
				resource.Attribute{
					Name:        "SW ANY",
					Description: `Value - SW ANY for configuring PstateCoordType token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "ESCN",
					Description: `Value - ESCN for configuring PuttyKeyPad token.`,
				},
				resource.Attribute{
					Name:        "LINUX",
					Description: `Value - LINUX for configuring PuttyKeyPad token.`,
				},
				resource.Attribute{
					Name:        "SCO",
					Description: `Value - SCO for configuring PuttyKeyPad token.`,
				},
				resource.Attribute{
					Name:        "VT100",
					Description: `Value - VT100 for configuring PuttyKeyPad token.`,
				},
				resource.Attribute{
					Name:        "VT400",
					Description: `Value - VT400 for configuring PuttyKeyPad token.`,
				},
				resource.Attribute{
					Name:        "XTERMR6",
					Description: `Value - XTERMR6 for configuring PuttyKeyPad token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "bios",
					Description: `Value - BIOS for configuring PwrPerfTuning token.`,
				},
				resource.Attribute{
					Name:        "os",
					Description: `Value - os for configuring PwrPerfTuning token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "6.4-gt/s",
					Description: `Value - 6.4-gt/s for configuring QpiLinkFrequency token.`,
				},
				resource.Attribute{
					Name:        "7.2-gt/s",
					Description: `Value - 7.2-gt/s for configuring QpiLinkFrequency token.`,
				},
				resource.Attribute{
					Name:        "8.0-gt/s",
					Description: `Value - 8.0-gt/s for configuring QpiLinkFrequency token.`,
				},
				resource.Attribute{
					Name:        "9.6-gt/s",
					Description: `Value - 9.6-gt/s for configuring QpiLinkFrequency token.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring QpiLinkFrequency token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "10.4GT/s",
					Description: `Value - 10.4GT/s for configuring QpiLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "9.6GT/s",
					Description: `Value - 9.6GT/s for configuring QpiLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring QpiLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring QpiSnoopMode token.`,
				},
				resource.Attribute{
					Name:        "cluster-on-die",
					Description: `Value - cluster-on-die for configuring QpiSnoopMode token.`,
				},
				resource.Attribute{
					Name:        "early-snoop",
					Description: `Value - early-snoop for configuring QpiSnoopMode token.`,
				},
				resource.Attribute{
					Name:        "home-directory-snoop",
					Description: `Value - home-directory-snoop for configuring QpiSnoopMode token.`,
				},
				resource.Attribute{
					Name:        "home-directory-snoop-with-osb",
					Description: `Value - home-directory-snoop-with-osb for configuring QpiSnoopMode token.`,
				},
				resource.Attribute{
					Name:        "home-snoop",
					Description: `Value - home-snoop for configuring QpiSnoopMode token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "1-way",
					Description: `Value - 1-way for configuring RankInterLeave token.`,
				},
				resource.Attribute{
					Name:        "2-way",
					Description: `Value - 2-way for configuring RankInterLeave token.`,
				},
				resource.Attribute{
					Name:        "4-way",
					Description: `Value - 4-way for configuring RankInterLeave token.`,
				},
				resource.Attribute{
					Name:        "8-way",
					Description: `Value - 8-way for configuring RankInterLeave token.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `Value - auto for configuring RankInterLeave token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Always Enable",
					Description: `Value - Always Enable for configuring RedirectionAfterPost token.`,
				},
				resource.Attribute{
					Name:        "Bootloader",
					Description: `Value - Bootloader for configuring RedirectionAfterPost token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "AHCI",
					Description: `Value - AHCI for configuring SataModeSelect token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SataModeSelect token.`,
				},
				resource.Attribute{
					Name:        "LSI SW RAID",
					Description: `Value - LSI SW RAID for configuring SataModeSelect token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "adddc-sparing",
					Description: `Value - adddc-sparing for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "lockstep",
					Description: `Value - lockstep for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "maximum-performance",
					Description: `Value - maximum-performance for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "mirror-mode-1lm",
					Description: `Value - mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "mirroring",
					Description: `Value - mirroring for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "partial-mirror-mode-1lm",
					Description: `Value - partial-mirror-mode-1lm for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "sparing",
					Description: `Value - sparing for configuring SelectMemoryRasConfiguration token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SelectPprType token.`,
				},
				resource.Attribute{
					Name:        "Hard PPR",
					Description: `Value - Hard PPR for configuring SelectPprType token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "253 ASIDs",
					Description: `Value - 253 ASIDs for configuring Sev token.`,
				},
				resource.Attribute{
					Name:        "509 ASIDs",
					Description: `Value - 509 ASIDs for configuring Sev token.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Sev token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "No",
					Description: `Value - No for configuring SinglePctlEnable token.`,
				},
				resource.Attribute{
					Name:        "Yes",
					Description: `Value - Yes for configuring SinglePctlEnable token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot10linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot10linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot10linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot10linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot10linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot10state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot10state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot10state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot10state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot11linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot11linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot11linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot11linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot11linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot12linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot12linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot12linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot12linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot12linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot1state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot1state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot1state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot1state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot2state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot2state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot2state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot2state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot3state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot3state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot3state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot3state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot4state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot4state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot4state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot4state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot5state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot5state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot5state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot5state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot6state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot6state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot6state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot6state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot7linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot7linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot7linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot7linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot7linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot7state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot7state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot7state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot7state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot8linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot8linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot8linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot8linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot8linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot8state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot8state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot8state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot8state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Slot9linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring Slot9linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring Slot9linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring Slot9linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring Slot9linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Slot9state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Slot9state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring Slot9state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring Slot9state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotFlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotFlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotFlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotFlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotFlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotFrontNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotFrontNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotFrontNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotFrontNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotFrontNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotFrontNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotFrontNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotFrontNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotFrontNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotFrontNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotFrontSlot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotFrontSlot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotFrontSlot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotFrontSlot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotFrontSlot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotFrontSlot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotFrontSlot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotFrontSlot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotFrontSlot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotFrontSlot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotHbaLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotHbaLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotHbaLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotHbaLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotHbaLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SlotHbaState token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring SlotHbaState token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring SlotHbaState token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring SlotHbaState token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SlotMezzState token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring SlotMezzState token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring SlotMezzState token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring SlotMezzState token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotMlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotMlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotMlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotMlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotMlomLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SlotMlomState token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring SlotMlomState token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring SlotMlomState token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring SlotMlomState token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotMraidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotMraidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotMraidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotMraidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotMraidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SlotN1state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring SlotN1state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring SlotN1state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring SlotN1state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SlotN2state token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring SlotN2state token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring SlotN2state token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring SlotN2state token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRaidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRaidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRaidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRaidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRaidLinkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRearNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRearNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRearNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRearNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRearNvme1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRearNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRearNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRearNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRearNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRearNvme2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser1slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser1slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser1slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser1slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser1slot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser1slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser1slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser1slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser1slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser1slot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser1slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser1slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser1slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser1slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser1slot3linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser2slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser2slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser2slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser2slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser2slot4linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser2slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser2slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser2slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser2slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser2slot5linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotRiser2slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotRiser2slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotRiser2slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotRiser2slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotRiser2slot6linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring SlotSasState token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring SlotSasState token.`,
				},
				resource.Attribute{
					Name:        "Legacy Only",
					Description: `Value - Legacy Only for configuring SlotSasState token.`,
				},
				resource.Attribute{
					Name:        "UEFI Only",
					Description: `Value - UEFI Only for configuring SlotSasState token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotSsdSlot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotSsdSlot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotSsdSlot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotSsdSlot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotSsdSlot1linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SlotSsdSlot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Value - Disabled for configuring SlotSsdSlot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN1",
					Description: `Value - GEN1 for configuring SlotSsdSlot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN2",
					Description: `Value - GEN2 for configuring SlotSsdSlot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "GEN3",
					Description: `Value - GEN3 for configuring SlotSsdSlot2linkSpeed token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring SmtMode token.`,
				},
				resource.Attribute{
					Name:        "Off",
					Description: `Value - Off for configuring SmtMode token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Snc token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Snc token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Snc token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "dimm-sparing",
					Description: `Value - dimm-sparing for configuring SparingMode token.`,
				},
				resource.Attribute{
					Name:        "rank-sparing",
					Description: `Value - rank-sparing for configuring SparingMode token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "pc-ansi",
					Description: `Value - pc-ansi for configuring TerminalType token.`,
				},
				resource.Attribute{
					Name:        "vt-utf8",
					Description: `Value - vt-utf8 for configuring TerminalType token.`,
				},
				resource.Attribute{
					Name:        "vt100",
					Description: `Value - vt100 for configuring TerminalType token.`,
				},
				resource.Attribute{
					Name:        "vt100-plus",
					Description: `Value - vt100-plus for configuring TerminalType token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring Tsme token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring Tsme token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring Tsme token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Loose",
					Description: `Value - Loose for configuring UcsmBootOrderRule token.`,
				},
				resource.Attribute{
					Name:        "Strict",
					Description: `Value - Strict for configuring UcsmBootOrderRule token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Offboard",
					Description: `Value - Offboard for configuring VgaPriority token.`,
				},
				resource.Attribute{
					Name:        "Onboard",
					Description: `Value - Onboard for configuring VgaPriority token.`,
				},
				resource.Attribute{
					Name:        "Onboard VGA Disabled",
					Description: `Value - Onboard VGA Disabled for configuring VgaPriority token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Enables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Disables the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Balanced",
					Description: `Value - Balanced for configuring WorkLoadConfig token.`,
				},
				resource.Attribute{
					Name:        "I/O Sensitive",
					Description: `Value - I/O Sensitive for configuring WorkLoadConfig token.`,
				},
				resource.Attribute{
					Name:        "NUMA",
					Description: `Value - NUMA for configuring WorkLoadConfig token.`,
				},
				resource.Attribute{
					Name:        "UMA",
					Description: `Value - UMA for configuring WorkLoadConfig token.`,
				},
				resource.Attribute{
					Name:        "platform-default",
					Description: `Default value used by the platform for the BIOS setting.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Value - Auto for configuring XptPrefetch token.`,
				},
				resource.Attribute{
					Name:        "disabled",
					Description: `Value - disabled for configuring XptPrefetch token.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Value - enabled for configuring XptPrefetch token.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_bios_system_boot_order",
			Category:         "bios",
			ShortDescription: `Actual Boot Order of the system.`,
			Description: `
Actual Boot Order of the system.
`,
			Keywords: []string{
				"bios",
				"system",
				"boot",
				"order",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Legacy",
					Description: `Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.`,
				},
				resource.Attribute{
					Name:        "Uefi",
					Description: `UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Set the state of Secure Boot to Not Available.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Set the state of Secure Boot to Disabled.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Set the state of Secure Boot to Enabled.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_bios_unit",
			Category:         "bios",
			ShortDescription: `The BIOS Unit present on a server.`,
			Description: `
The BIOS Unit present on a server.
`,
			Keywords: []string{
				"bios",
				"unit",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_cdd_device",
			Category:         "boot",
			ShortDescription: `Cdd Boot Device configured on the server.`,
			Description: `
Cdd Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"cdd",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_device_boot_mode",
			Category:         "boot",
			ShortDescription: `Boot mode of the devices that BIOS uses to boot them.`,
			Description: `
Boot mode of the devices that BIOS uses to boot them.
`,
			Keywords: []string{
				"boot",
				"device",
				"mode",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_device_boot_security",
			Category:         "boot",
			ShortDescription: `Boot Security of the devices that BIOS uses to boot them.`,
			Description: `
Boot Security of the devices that BIOS uses to boot them.
`,
			Keywords: []string{
				"boot",
				"device",
				"security",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_hdd_device",
			Category:         "boot",
			ShortDescription: `Local Disk Boot Device configured on the server.`,
			Description: `
Local Disk Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"hdd",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_iscsi_device",
			Category:         "boot",
			ShortDescription: `Iscsi Boot Device configured on the server.`,
			Description: `
Iscsi Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"iscsi",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_nvme_device",
			Category:         "boot",
			ShortDescription: `Nvme Boot Device configured on the server.`,
			Description: `
Nvme Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"nvme",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_pch_storage_device",
			Category:         "boot",
			ShortDescription: `Pch Storage Boot Device configured on the server.`,
			Description: `
Pch Storage Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"pch",
				"storage",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_precision_policy",
			Category:         "boot",
			ShortDescription: `Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.`,
			Description: `
Boot order policy models a reusable boot order configuration that can be applied to multiple servers via profile association. It supports advanced boot order configuration on Cisco CIMC servers.
`,
			Keywords: []string{
				"boot",
				"precision",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Legacy",
					Description: `Legacy mode refers to the traditional process of booting from BIOS. Legacy mode uses the Master Boot Record (MBR) to locate the bootloader.`,
				},
				resource.Attribute{
					Name:        "Uefi",
					Description: `UEFI mode uses the GUID Partition Table (GPT) to locate EFI Service Partitions to boot from.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_pxe_device",
			Category:         "boot",
			ShortDescription: `Pxe Boot Device configured on the server.`,
			Description: `
Pxe Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"pxe",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_san_device",
			Category:         "boot",
			ShortDescription: `San Boot Device configured on the server.`,
			Description: `
San Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"san",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_sd_device",
			Category:         "boot",
			ShortDescription: `Sd Boot Device configured on the server.`,
			Description: `
Sd Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"sd",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_uefi_shell_device",
			Category:         "boot",
			ShortDescription: `UefiShell Boot Device configured on the server.`,
			Description: `
UefiShell Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"uefi",
				"shell",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_usb_device",
			Category:         "boot",
			ShortDescription: `Usb Boot Device configured on the server.`,
			Description: `
Usb Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"usb",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_boot_vmedia_device",
			Category:         "boot",
			ShortDescription: `Virtual Media Boot Device configured on the server.`,
			Description: `
Virtual Media Boot Device configured on the server.
`,
			Keywords: []string{
				"boot",
				"vmedia",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_adapter_unit_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies an adaptor.`,
			Description: `
Descriptor that uniquely identifies an adaptor.
`,
			Keywords: []string{
				"capability",
				"adapter",
				"unit",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_catalog",
			Category:         "capability",
			ShortDescription: `Container for capability information of managed systems. This catalog will be managed by devops using a specific role in the Catalog Admin account.`,
			Description: `
Container for capability information of managed systems.
This catalog will be managed by devops using a specific role in the Catalog Admin account.
`,
			Keywords: []string{
				"capability",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_chassis_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies an chassis enclosure.`,
			Description: `
Descriptor that uniquely identifies an chassis enclosure.
`,
			Keywords: []string{
				"capability",
				"chassis",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_chassis_manufacturing_def",
			Category:         "capability",
			ShortDescription: `Chassis enclosure manufacturing def properties.`,
			Description: `
Chassis enclosure manufacturing def properties.
`,
			Keywords: []string{
				"capability",
				"chassis",
				"manufacturing",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_cimc_firmware_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that identifies the server's redfish integration capability using cimc firmware info.`,
			Description: `
Descriptor that identifies the server's redfish integration capability using cimc firmware info.
`,
			Keywords: []string{
				"capability",
				"cimc",
				"firmware",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_equipment_physical_def",
			Category:         "capability",
			ShortDescription: `Type to represent additional switch specific capabilities.`,
			Description: `
Type to represent additional switch specific capabilities.
`,
			Keywords: []string{
				"capability",
				"equipment",
				"physical",
				"def",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UCS-FI-6454",
					Description: `The standard 4th generation UCS Fabric Interconnect with 54 ports.`,
				},
				resource.Attribute{
					Name:        "UCS-FI-64108",
					Description: `The expanded 4th generation UCS Fabric Interconnect with 108 ports.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `Unknown device type, usage is TBD.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_equipment_slot_array",
			Category:         "capability",
			ShortDescription: `Type to represent additional switch specific capabilities.`,
			Description: `
Type to represent additional switch specific capabilities.
`,
			Keywords: []string{
				"capability",
				"equipment",
				"slot",
				"array",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UCS-FI-6454",
					Description: `The standard 4th generation UCS Fabric Interconnect with 54 ports.`,
				},
				resource.Attribute{
					Name:        "UCS-FI-64108",
					Description: `The expanded 4th generation UCS Fabric Interconnect with 108 ports.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `Unknown device type, usage is TBD.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_fan_module_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies a fan module.`,
			Description: `
Descriptor that uniquely identifies a fan module.
`,
			Keywords: []string{
				"capability",
				"fan",
				"module",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_fan_module_manufacturing_def",
			Category:         "capability",
			ShortDescription: `Fan module unit that contains multiple fans.`,
			Description: `
Fan module unit that contains multiple fans.
`,
			Keywords: []string{
				"capability",
				"fan",
				"module",
				"manufacturing",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_io_card_capability_def",
			Category:         "capability",
			ShortDescription: `Chassis Iocard module capabilities.`,
			Description: `
Chassis Iocard module capabilities.
`,
			Keywords: []string{
				"capability",
				"io",
				"card",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_io_card_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies an IO card module.`,
			Description: `
Descriptor that uniquely identifies an IO card module.
`,
			Keywords: []string{
				"capability",
				"io",
				"card",
				"descriptor",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "inline",
					Description: `UIF uplink ports and IOM ports are connected inline.`,
				},
				resource.Attribute{
					Name:        "cross-connected",
					Description: `UIF uplink ports and IOM ports are cross-connected, a case in washington chassis.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_io_card_manufacturing_def",
			Category:         "capability",
			ShortDescription: `Chassis Iocard module properties.`,
			Description: `
Chassis Iocard module properties.
`,
			Keywords: []string{
				"capability",
				"io",
				"card",
				"manufacturing",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_port_group_aggregation_def",
			Category:         "capability",
			ShortDescription: `FEX/IOCARD module port group aggregation capabilities.`,
			Description: `
FEX/IOCARD module port group aggregation capabilities.
`,
			Keywords: []string{
				"capability",
				"port",
				"group",
				"aggregation",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_psu_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies a power supply.`,
			Description: `
Descriptor that uniquely identifies a power supply.
`,
			Keywords: []string{
				"capability",
				"psu",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_psu_manufacturing_def",
			Category:         "capability",
			ShortDescription: `Power supply unit properties.`,
			Description: `
Power supply unit properties.
`,
			Keywords: []string{
				"capability",
				"psu",
				"manufacturing",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_sioc_module_capability_def",
			Category:         "capability",
			ShortDescription: `Chassis SIOC module capabilities.`,
			Description: `
Chassis SIOC module capabilities.
`,
			Keywords: []string{
				"capability",
				"sioc",
				"module",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_sioc_module_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies an SIOC module.`,
			Description: `
Descriptor that uniquely identifies an SIOC module.
`,
			Keywords: []string{
				"capability",
				"sioc",
				"module",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_sioc_module_manufacturing_def",
			Category:         "capability",
			ShortDescription: `Chassis SIOC module properties.`,
			Description: `
Chassis SIOC module properties.
`,
			Keywords: []string{
				"capability",
				"sioc",
				"module",
				"manufacturing",
				"def",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_switch_capability",
			Category:         "capability",
			ShortDescription: `Type to represent additional switch specific capabilities.`,
			Description: `
Type to represent additional switch specific capabilities.
`,
			Keywords: []string{
				"capability",
				"switch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UCS-FI-6454",
					Description: `The standard 4th generation UCS Fabric Interconnect with 54 ports.`,
				},
				resource.Attribute{
					Name:        "UCS-FI-64108",
					Description: `The expanded 4th generation UCS Fabric Interconnect with 108 ports.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `Unknown device type, usage is TBD.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_switch_descriptor",
			Category:         "capability",
			ShortDescription: `Descriptor that uniquely identifies a Fabric interconnect.`,
			Description: `
Descriptor that uniquely identifies a Fabric interconnect.
`,
			Keywords: []string{
				"capability",
				"switch",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_capability_switch_manufacturing_def",
			Category:         "capability",
			ShortDescription: `Switch/Fabric-Interconnect manufacturing def properties.`,
			Description: `
Switch/Fabric-Interconnect manufacturing def properties.
`,
			Keywords: []string{
				"capability",
				"switch",
				"manufacturing",
				"def",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UCS-FI-6454",
					Description: `The standard 4th generation UCS Fabric Interconnect with 54 ports.`,
				},
				resource.Attribute{
					Name:        "UCS-FI-64108",
					Description: `The expanded 4th generation UCS Fabric Interconnect with 108 ports.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `Unknown device type, usage is TBD.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_certificatemanagement_policy",
			Category:         "certificatemanagement",
			ShortDescription: `Certificate Management policy models a reusable certificate and private key configuration that can be applied to multiple servers via profile association.`,
			Description: `
Certificate Management policy models a reusable certificate and private key configuration that can be applied to multiple servers via profile association.
`,
			Keywords: []string{
				"certificatemanagement",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_chassis_config_change_detail",
			Category:         "chassis",
			ShortDescription: `The configuration change details are captured here.`,
			Description: `
The configuration change details are captured here.
`,
			Keywords: []string{
				"chassis",
				"config",
				"change",
				"detail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Pending-changes",
					Description: `Config change flag represents changes are due to not deployed changes from Intersight.`,
				},
				resource.Attribute{
					Name:        "Drift-changes",
					Description: `Config change flag represents changes are due to endpoint configuration changes.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration.`,
				},
				resource.Attribute{
					Name:        "Created",
					Description: `The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet.`,
				},
				resource.Attribute{
					Name:        "Modified",
					Description: `The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully.`,
				},
				resource.Attribute{
					Name:        "Deleted",
					Description: `The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_chassis_config_import",
			Category:         "chassis",
			ShortDescription: `Configuration import action will import the existing configuration from chassis and generate Intersight policies profile from it. At end of successful import, chassis will be assigned to the generated profile which has policies associated with it. No chassis profile or policies will be generated if configuration import fails.`,
			Description: `
Configuration import action will import the existing configuration from chassis and generate Intersight policies profile from it. At end of successful import, chassis will be assigned to the generated profile which has policies associated with it. No chassis profile or policies will be generated if configuration import fails.
`,
			Keywords: []string{
				"chassis",
				"config",
				"import",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_chassis_config_result",
			Category:         "chassis",
			ShortDescription: `The profile configuration (deploy, validation) results with the overall state and detailed result messages.`,
			Description: `
The profile configuration (deploy, validation) results with the overall state and detailed result messages.
`,
			Keywords: []string{
				"chassis",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_chassis_config_result_entry",
			Category:         "chassis",
			ShortDescription: `The profile configuration (deploy, validation) results details information.`,
			Description: `
The profile configuration (deploy, validation) results details information.
`,
			Keywords: []string{
				"chassis",
				"config",
				"result",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_chassis_iom_profile",
			Category:         "chassis",
			ShortDescription: `A profile specifying configuration settings for IOM.`,
			Description: `
A profile specifying configuration settings for IOM.
`,
			Keywords: []string{
				"chassis",
				"iom",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IOMA",
					Description: `IOM on left side of chassis.`,
				},
				resource.Attribute{
					Name:        "IOMB",
					Description: `IOM on right side of chassis.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_chassis_profile",
			Category:         "chassis",
			ShortDescription: `A profile specifying configuration settings for a chassis.`,
			Description: `
A profile specifying configuration settings for a chassis.
`,
			Keywords: []string{
				"chassis",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "FIAttached",
					Description: `Chassis which are connected to a Fabric Interconnect that is managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cloud_regions",
			Category:         "cloud",
			ShortDescription: `The geographic location where a clouds resources are located. It has details such as cloud name, region name, region identifier, list of zones, region endpoint etc.`,
			Description: `
The geographic location where a clouds resources are located. It has details such as cloud name, region name, region identifier, list of zones, region endpoint etc.
`,
			Keywords: []string{
				"cloud",
				"regions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cloud_sku_container_type",
			Category:         "cloud",
			ShortDescription: `Stores hardware attribute information for a container.`,
			Description: `
Stores hardware attribute information for a container.
`,
			Keywords: []string{
				"cloud",
				"sku",
				"container",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "VIRTUAL_CPU",
					Description: `The CPU unit used for virtual machines.`,
				},
				resource.Attribute{
					Name:        "MILLI_CPU",
					Description: `The CPU unit used by containers.`,
				},
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "Compute",
					Description: `Compute service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `Storage service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Database",
					Description: `Database service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Network",
					Description: `Network service offered by cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cloud_sku_database_type",
			Category:         "cloud",
			ShortDescription: `Stores details of instance type which handle databases.`,
			Description: `
Stores details of instance type which handle databases.
`,
			Keywords: []string{
				"cloud",
				"sku",
				"database",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "Compute",
					Description: `Compute service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `Storage service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Database",
					Description: `Database service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Network",
					Description: `Network service offered by cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cloud_sku_instance_type",
			Category:         "cloud",
			ShortDescription: `Details for an instance type.`,
			Description: `
Details for an instance type.
`,
			Keywords: []string{
				"cloud",
				"sku",
				"instance",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "64Bit",
					Description: `Enum to indicate that the instance type suppports only 64 bit architecture.`,
				},
				resource.Attribute{
					Name:        "32Bit",
					Description: `Enum to indicate that the instance type supports only 32 bit architecture.`,
				},
				resource.Attribute{
					Name:        "both",
					Description: `Enum to indicate that the instance type supports both 32 and 64 bit architecture.`,
				},
				resource.Attribute{
					Name:        "VIRTUAL_CPU",
					Description: `The CPU unit used for virtual machines.`,
				},
				resource.Attribute{
					Name:        "MILLI_CPU",
					Description: `The CPU unit used by containers.`,
				},
				resource.Attribute{
					Name:        "MB",
					Description: `Enum to represent mega byte storage unit.`,
				},
				resource.Attribute{
					Name:        "GB",
					Description: `Enum to represent Giga byte storage unit.`,
				},
				resource.Attribute{
					Name:        "TB",
					Description: `Enum to represent Tera byte storage unit.`,
				},
				resource.Attribute{
					Name:        "MB",
					Description: `Enum to represent mega byte storage unit.`,
				},
				resource.Attribute{
					Name:        "GB",
					Description: `Enum to represent Giga byte storage unit.`,
				},
				resource.Attribute{
					Name:        "TB",
					Description: `Enum to represent Tera byte storage unit.`,
				},
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "Compute",
					Description: `Compute service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `Storage service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Database",
					Description: `Database service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Network",
					Description: `Network service offered by cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cloud_sku_network_type",
			Category:         "cloud",
			ShortDescription: `Model to represent network attributes.`,
			Description: `
Model to represent network attributes.
`,
			Keywords: []string{
				"cloud",
				"sku",
				"network",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "Compute",
					Description: `Compute service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `Storage service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Database",
					Description: `Database service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Network",
					Description: `Network service offered by cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cloud_sku_volume_type",
			Category:         "cloud",
			ShortDescription: `Stores information about the volume types.`,
			Description: `
Stores information about the volume types.
`,
			Keywords: []string{
				"cloud",
				"sku",
				"volume",
				"type",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
				resource.Attribute{
					Name:        "Compute",
					Description: `Compute service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `Storage service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Database",
					Description: `Database service offered by cloud provider.`,
				},
				resource.Attribute{
					Name:        "Network",
					Description: `Network service offered by cloud provider.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_comm_http_proxy_policy",
			Category:         "comm",
			ShortDescription: `A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.`,
			Description: `
A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
`,
			Keywords: []string{
				"comm",
				"http",
				"proxy",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_blade",
			Category:         "compute",
			ShortDescription: `Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.`,
			Description: `
Server which is housed in a chassis and shares some of the hardware with other servers in the chassis.
`,
			Keywords: []string{
				"compute",
				"blade",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IntersightStandalone",
					Description: `Intersight Standalone mode of operation.`,
				},
				resource.Attribute{
					Name:        "UCSM",
					Description: `Unified Computing System Manager mode of operation.`,
				},
				resource.Attribute{
					Name:        "Intersight",
					Description: `Intersight managed mode of operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_blade_identity",
			Category:         "compute",
			ShortDescription: `Identity object that uniquely represents a blade server object under a DR.`,
			Description: `
Identity object that uniquely represents a blade server object under a DR.
`,
			Keywords: []string{
				"compute",
				"blade",
				"identity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation value for maintenance actions on an equipment.`,
				},
				resource.Attribute{
					Name:        "Decommission",
					Description: `Decommission the equipment and temporarily remove it from being managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "Recommission",
					Description: `Recommission the equipment.`,
				},
				resource.Attribute{
					Name:        "Reack",
					Description: `Reacknowledge the equipment and discover it again.`,
				},
				resource.Attribute{
					Name:        "Remove",
					Description: `Remove the equipment permanently from Intersight management.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The running firmware version is unknown.`,
				},
				resource.Attribute{
					Name:        "Supported",
					Description: `The running firmware version is known and supports IMM mode.`,
				},
				resource.Attribute{
					Name:        "NotSupported",
					Description: `The running firmware version is known and does not support IMM mode.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Default state of an equipment. This should be an initial state when no state is defined for an equipment.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `Default Lifecycle State for a physical entity.`,
				},
				resource.Attribute{
					Name:        "Decommissioned",
					Description: `Decommission Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DecommissionInProgress",
					Description: `Decommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RecommissionInProgress",
					Description: `Recommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "OperationFailed",
					Description: `Failed Operation Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "ReackInProgress",
					Description: `ReackInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RemoveInProgress",
					Description: `RemoveInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "Discovered",
					Description: `Discovered Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryInProgress",
					Description: `DiscoveryInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryFailed",
					Description: `DiscoveryFailed Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "FirmwareUpgradeInProgress",
					Description: `Firmware upgrade is in progress on given physical entity.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The default presence state.`,
				},
				resource.Attribute{
					Name:        "Equipped",
					Description: `The server is equipped in the slot.`,
				},
				resource.Attribute{
					Name:        "EquippedMismatch",
					Description: `The slot is equipped, but there is another server currently inventoried in the slot.`,
				},
				resource.Attribute{
					Name:        "Missing",
					Description: `The server is not present in the given slot.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_board",
			Category:         "compute",
			ShortDescription: `Mother board of a server.`,
			Description: `
Mother board of a server.
`,
			Keywords: []string{
				"compute",
				"board",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_physical_summary",
			Category:         "compute",
			ShortDescription: `Consolidated view of Blades and RackUnits.`,
			Description: `
Consolidated view of Blades and RackUnits.
`,
			Keywords: []string{
				"compute",
				"physical",
				"summary",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IntersightStandalone",
					Description: `Intersight Standalone mode of operation.`,
				},
				resource.Attribute{
					Name:        "UCSM",
					Description: `Unified Computing System Manager mode of operation.`,
				},
				resource.Attribute{
					Name:        "Intersight",
					Description: `Intersight managed mode of operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_rack_unit",
			Category:         "compute",
			ShortDescription: `Describes a standalone or FI-attached Rack-mounted server.`,
			Description: `
Describes a standalone or FI-attached Rack-mounted server.
`,
			Keywords: []string{
				"compute",
				"rack",
				"unit",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IntersightStandalone",
					Description: `Intersight Standalone mode of operation.`,
				},
				resource.Attribute{
					Name:        "UCSM",
					Description: `Unified Computing System Manager mode of operation.`,
				},
				resource.Attribute{
					Name:        "Intersight",
					Description: `Intersight managed mode of operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_rack_unit_identity",
			Category:         "compute",
			ShortDescription: `Identity object that uniquely represents a server object under a DR.`,
			Description: `
Identity object that uniquely represents a server object under a DR.
`,
			Keywords: []string{
				"compute",
				"rack",
				"unit",
				"identity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation value for maintenance actions on an equipment.`,
				},
				resource.Attribute{
					Name:        "Decommission",
					Description: `Decommission the equipment and temporarily remove it from being managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "Recommission",
					Description: `Recommission the equipment.`,
				},
				resource.Attribute{
					Name:        "Reack",
					Description: `Reacknowledge the equipment and discover it again.`,
				},
				resource.Attribute{
					Name:        "Remove",
					Description: `Remove the equipment permanently from Intersight management.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Default state of an equipment. This should be an initial state when no state is defined for an equipment.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `Default Lifecycle State for a physical entity.`,
				},
				resource.Attribute{
					Name:        "Decommissioned",
					Description: `Decommission Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DecommissionInProgress",
					Description: `Decommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RecommissionInProgress",
					Description: `Recommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "OperationFailed",
					Description: `Failed Operation Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "ReackInProgress",
					Description: `ReackInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RemoveInProgress",
					Description: `RemoveInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "Discovered",
					Description: `Discovered Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryInProgress",
					Description: `DiscoveryInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryFailed",
					Description: `DiscoveryFailed Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "FirmwareUpgradeInProgress",
					Description: `Firmware upgrade is in progress on given physical entity.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_server_setting",
			Category:         "compute",
			ShortDescription: `Models the configurable properties of a server in Intersight.`,
			Description: `
Models the configurable properties of a server in Intersight.
`,
			Keywords: []string{
				"compute",
				"server",
				"setting",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation property for locator led.`,
				},
				resource.Attribute{
					Name:        "On",
					Description: `The Locator Led is turned on.`,
				},
				resource.Attribute{
					Name:        "Off",
					Description: `The Locator Led is turned off.`,
				},
				resource.Attribute{
					Name:        "Policy",
					Description: `Power state is set to the default value in the policy.`,
				},
				resource.Attribute{
					Name:        "PowerOn",
					Description: `Power state of the server is set to On.`,
				},
				resource.Attribute{
					Name:        "PowerOff",
					Description: `Power state is the server set to Off.`,
				},
				resource.Attribute{
					Name:        "PowerCycle",
					Description: `Power state the server is reset.`,
				},
				resource.Attribute{
					Name:        "HardReset",
					Description: `Power state the server is hard reset.`,
				},
				resource.Attribute{
					Name:        "Shutdown",
					Description: `Operating system on the server is shut down.`,
				},
				resource.Attribute{
					Name:        "Reboot",
					Description: `Power state of IMC is rebooted.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_compute_vmedia",
			Category:         "compute",
			ShortDescription: `Inventory of Virtual Media configuration and images uploaded.`,
			Description: `
Inventory of Virtual Media configuration and images uploaded.
`,
			Keywords: []string{
				"compute",
				"vmedia",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cond_alarm",
			Category:         "cond",
			ShortDescription: `A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.`,
			Description: `
A state-full entity representing a found problem. Alarms can be reported by the managed system itself or can be determined by Intersight.
`,
			Keywords: []string{
				"cond",
				"alarm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `The Enum value None represents that the alarm is not acknowledged and is included as part of health status and overall alarm count.`,
				},
				resource.Attribute{
					Name:        "Acknowledge",
					Description: `The Enum value Acknowledge represents that the alarm is acknowledged by user. The alarm will be ignored from the health status and overall alarm count.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The Enum value None represents that there is no severity.`,
				},
				resource.Attribute{
					Name:        "Info",
					Description: `The Enum value Info represents the Informational level of severity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `The Enum value Critical represents the Critical level of severity.`,
				},
				resource.Attribute{
					Name:        "Warning",
					Description: `The Enum value Warning represents the Warning level of severity.`,
				},
				resource.Attribute{
					Name:        "Cleared",
					Description: `The Enum value Cleared represents that the alarm severity has been cleared.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The Enum value None represents that there is no severity.`,
				},
				resource.Attribute{
					Name:        "Info",
					Description: `The Enum value Info represents the Informational level of severity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `The Enum value Critical represents the Critical level of severity.`,
				},
				resource.Attribute{
					Name:        "Warning",
					Description: `The Enum value Warning represents the Warning level of severity.`,
				},
				resource.Attribute{
					Name:        "Cleared",
					Description: `The Enum value Cleared represents that the alarm severity has been cleared.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cond_alarm_aggregation",
			Category:         "cond",
			ShortDescription: `Object which represents alarm aggregation for a managed end point.`,
			Description: `
Object which represents alarm aggregation for a managed end point.
`,
			Keywords: []string{
				"cond",
				"alarm",
				"aggregation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `The Enum value None represents that there is no severity.`,
				},
				resource.Attribute{
					Name:        "Info",
					Description: `The Enum value Info represents the Informational level of severity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `The Enum value Critical represents the Critical level of severity.`,
				},
				resource.Attribute{
					Name:        "Warning",
					Description: `The Enum value Warning represents the Warning level of severity.`,
				},
				resource.Attribute{
					Name:        "Cleared",
					Description: `The Enum value Cleared represents that the alarm severity has been cleared.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cond_hcl_status",
			Category:         "cond",
			ShortDescription: `The HCL status of a managed object after we have validated the managed object components' firmware and drivers against the HCL.`,
			Description: `
The HCL status of a managed object after we have validated the managed object components' firmware and drivers against the HCL.
`,
			Keywords: []string{
				"cond",
				"hcl",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Incomplete",
					Description: `This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Not-Found",
					Description: `At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Listed",
					Description: `At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Validated",
					Description: `At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.`,
				},
				resource.Attribute{
					Name:        "Incomplete",
					Description: `This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Not-Found",
					Description: `At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Listed",
					Description: `At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Validated",
					Description: `At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.`,
				},
				resource.Attribute{
					Name:        "Missing-Os-Info",
					Description: `This means the HclStatus for the sever failed HCL validation because we have missing os information. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Components",
					Description: `This means the HclStatus for the sever failed HCL validation because one or more of its components failed validation. To see why components failed check the related HclStatusDetails.`,
				},
				resource.Attribute{
					Name:        "Compatible",
					Description: `This means the HclStatus for the sever has passed HCL validation for all of its related components.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `This means the HclStatus for the sever has not been evaluated because it is exempted.`,
				},
				resource.Attribute{
					Name:        "Missing-Os-Driver-Info",
					Description: `The validation failed becaue the given server has no OS driver information available in the inventory. Either install UCS Tools VIB on the host ESXi or use OS Discovery Tool scripts to provide proper OS information.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Server",
					Description: `The validation failed for this server because the server's model was not listed in the HCL.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Processor",
					Description: `The validation failed because the given processor was not listed for the given server model.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Os-Info",
					Description: `The validation failed because the given OS vendor or version was not listed in HCL for the server PID and processor combination.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware",
					Description: `The validation failed because the given server firmware was not listed in the HCL for the given server PID, processor, OS vendor and version.`,
				},
				resource.Attribute{
					Name:        "Service-Unavailable",
					Description: `The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.`,
				},
				resource.Attribute{
					Name:        "Service-Error",
					Description: `The validation has failed because the HCL data service has returned a service error or unrecognized result.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `This means the HclStatus for the sever has not been evaluated because it is exempted.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Components",
					Description: `The validation has failed for this server because one or more components have \ Not-Listed\ status.`,
				},
				resource.Attribute{
					Name:        "Compatible",
					Description: `The validation has passed for this server's model, processor, OS vendor and version.`,
				},
				resource.Attribute{
					Name:        "Incomplete",
					Description: `This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Not-Found",
					Description: `At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Listed",
					Description: `At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Validated",
					Description: `At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.`,
				},
				resource.Attribute{
					Name:        "Incomplete",
					Description: `This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Not-Found",
					Description: `At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Listed",
					Description: `At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Validated",
					Description: `At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cond_hcl_status_detail",
			Category:         "cond",
			ShortDescription: `The HCL status detail for each component firmware and driver.`,
			Description: `
The HCL status detail for each component firmware and driver.
`,
			Keywords: []string{
				"cond",
				"hcl",
				"status",
				"detail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Missing-Os-Driver-Info",
					Description: `The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Server-With-Component",
					Description: `The validation failed for this component because he server model and component model combination was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Processor",
					Description: `The validation failed because the given processor was not found for the given server PID.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Os-Info",
					Description: `The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Component-Model",
					Description: `The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware",
					Description: `The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Driver",
					Description: `The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware-Driver",
					Description: `The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.`,
				},
				resource.Attribute{
					Name:        "Service-Unavailable",
					Description: `The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.`,
				},
				resource.Attribute{
					Name:        "Service-Error",
					Description: `The validation has failed because the HCL data service has return a service error or unrecognized result.`,
				},
				resource.Attribute{
					Name:        "Unrecognized-Protocol",
					Description: `The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.`,
				},
				resource.Attribute{
					Name:        "Compatible",
					Description: `The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.`,
				},
				resource.Attribute{
					Name:        "Missing-Os-Driver-Info",
					Description: `The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Server-With-Component",
					Description: `The validation failed for this component because he server model and component model combination was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Processor",
					Description: `The validation failed because the given processor was not found for the given server PID.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Os-Info",
					Description: `The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Component-Model",
					Description: `The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware",
					Description: `The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Driver",
					Description: `The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware-Driver",
					Description: `The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.`,
				},
				resource.Attribute{
					Name:        "Service-Unavailable",
					Description: `The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.`,
				},
				resource.Attribute{
					Name:        "Service-Error",
					Description: `The validation has failed because the HCL data service has return a service error or unrecognized result.`,
				},
				resource.Attribute{
					Name:        "Unrecognized-Protocol",
					Description: `The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.`,
				},
				resource.Attribute{
					Name:        "Compatible",
					Description: `The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.`,
				},
				resource.Attribute{
					Name:        "Missing-Os-Driver-Info",
					Description: `The validation failed becaue the given server has no OS driver information available in the inventory. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Server-With-Component",
					Description: `The validation failed for this component because he server model and component model combination was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Processor",
					Description: `The validation failed because the given processor was not found for the given server PID.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Os-Info",
					Description: `The validation failed because the given OS vendor and version was not found in HCL for the server PID and processor combination.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Component-Model",
					Description: `The validation failed because the given Component model was not found in the HCL for the given server PID, processor, server Firmware and OS vendor and version.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware",
					Description: `The validation failed because the given server firmware or adapter firmware was not found in the HCL for the given server PID, processor, OS vendor and version and component model.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Driver",
					Description: `The validation failed because the given driver version was not found in the HCL for the given Server PID, processor, OS vendor and version, server firmware and component firmware.`,
				},
				resource.Attribute{
					Name:        "Incompatible-Firmware-Driver",
					Description: `The validation failed because the given component firmware and driver version was not found in the HCL for the given Server PID, processor, OS vendor and version and server firmware.`,
				},
				resource.Attribute{
					Name:        "Service-Unavailable",
					Description: `The validation has failed because HCL data service is temporarily not available. The server will be re-evaluated once HCL data service is back online or finished importing new HCL data.`,
				},
				resource.Attribute{
					Name:        "Service-Error",
					Description: `The validation has failed because the HCL data service has return a service error or unrecognized result.`,
				},
				resource.Attribute{
					Name:        "Unrecognized-Protocol",
					Description: `The validation has failed for the HCL component because the HCL data service has return a validation reason that is unknown to this service. This reason is used as a default failure reason reason in case we cannot map the error reason received from the HCL data service unto one of the other enum values.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `The validation for the hardware or software HCL status was not yet evaluated because some previous validation had failed. For example if a server's hardware profile fails to validate with HCL, then the server's software status will not be evaluated.`,
				},
				resource.Attribute{
					Name:        "Compatible",
					Description: `The validation has passed for this server PID, processor, OS vendor and version, component model, component firmware and driver version.`,
				},
				resource.Attribute{
					Name:        "Incomplete",
					Description: `This means we do not have os information in Intersight for this server. Either install ucstools vib or use power shell scripts to tag proper OS information.`,
				},
				resource.Attribute{
					Name:        "Not-Found",
					Description: `At HclStatus level, this means that one of the components has failed validation. At HclStatusDetail level, this means that his component's hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Listed",
					Description: `At the HclStatus level, this means that some part of the HCL validation has failed. This could be that either the server's hardware or software profile was not listed in the HCL or one of the components' hardware or software profile was not found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Validated",
					Description: `At the HclStatus level, this means that all of the components have passed validation. At HclStatusDetail level, this means that the component's hardware or software profile was found in the HCL.`,
				},
				resource.Attribute{
					Name:        "Not-Evaluated",
					Description: `At the HclStatus level this means that this means that SW or Component status has not been evaluated as the previous evaluation step has not passed yet. At the HclStatusDetail level this means that either HW or SW status has not been evaluted because a previous evaluation step has not passed yet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_cond_hcl_status_job",
			Category:         "cond",
			ShortDescription: `An HCLStatusJob is used to batch mo inventory notifications and process the evaluation of HCLStatus. When we receive a notification for an inventory MO, we will create a HCLStatusJob and inserted into the DB if it doesn't already exist. Then based on a timer we process the jobs in the DB and clear them.`,
			Description: `
An HCLStatusJob is used to batch mo inventory notifications and process the evaluation of HCLStatus. When we receive a notification for an inventory MO, we will create a HCLStatusJob and inserted into the DB if it doesn't already exist. Then based on a timer we process the jobs in the DB and clear them.
`,
			Keywords: []string{
				"cond",
				"hcl",
				"status",
				"job",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_config_exported_item",
			Category:         "config",
			ShortDescription: `A single managed object that is being exported.`,
			Description: `
A single managed object that is being exported.
`,
			Keywords: []string{
				"config",
				"exported",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "InProgress",
					Description: `The operation is in progress.`,
				},
				resource.Attribute{
					Name:        "Success",
					Description: `The operation has succeeded.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The operation has failed.`,
				},
				resource.Attribute{
					Name:        "RollBackInitiated",
					Description: `The rollback has been inititiated for import failure.`,
				},
				resource.Attribute{
					Name:        "RollBackFailed",
					Description: `The rollback has failed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackCompleted",
					Description: `The rollback has completed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackAborted",
					Description: `The rollback has been aborted for import failure.`,
				},
				resource.Attribute{
					Name:        "OperationTimedOut",
					Description: `The operation has timed out.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_config_exporter",
			Category:         "config",
			ShortDescription: `All export operations are captured as Exporter instances. Users shall use this Exporter mo to track the export operation progress.`,
			Description: `
All export operations are captured as Exporter instances. Users shall use this Exporter
mo to track the export operation progress.
`,
			Keywords: []string{
				"config",
				"exporter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "InProgress",
					Description: `The operation is in progress.`,
				},
				resource.Attribute{
					Name:        "Success",
					Description: `The operation has succeeded.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The operation has failed.`,
				},
				resource.Attribute{
					Name:        "RollBackInitiated",
					Description: `The rollback has been inititiated for import failure.`,
				},
				resource.Attribute{
					Name:        "RollBackFailed",
					Description: `The rollback has failed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackCompleted",
					Description: `The rollback has completed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackAborted",
					Description: `The rollback has been aborted for import failure.`,
				},
				resource.Attribute{
					Name:        "OperationTimedOut",
					Description: `The operation has timed out.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_config_imported_item",
			Category:         "config",
			ShortDescription: `A single managed object that is being imported.`,
			Description: `
A single managed object that is being imported.
`,
			Keywords: []string{
				"config",
				"imported",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "InProgress",
					Description: `The operation is in progress.`,
				},
				resource.Attribute{
					Name:        "Success",
					Description: `The operation has succeeded.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The operation has failed.`,
				},
				resource.Attribute{
					Name:        "RollBackInitiated",
					Description: `The rollback has been inititiated for import failure.`,
				},
				resource.Attribute{
					Name:        "RollBackFailed",
					Description: `The rollback has failed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackCompleted",
					Description: `The rollback has completed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackAborted",
					Description: `The rollback has been aborted for import failure.`,
				},
				resource.Attribute{
					Name:        "OperationTimedOut",
					Description: `The operation has timed out.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_config_importer",
			Category:         "config",
			ShortDescription: `All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.`,
			Description: `
All import operations are captured as Importer instances. Users shall use this Importer mo to track the import operation progress.
`,
			Keywords: []string{
				"config",
				"importer",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ImageRepo",
					Description: `The 'ImageRepo' source if the source of exporter archive is image repository.`,
				},
				resource.Attribute{
					Name:        "URL",
					Description: `The 'URL' source if the source of exported archive is a URL.`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `The operation is in progress.`,
				},
				resource.Attribute{
					Name:        "Success",
					Description: `The operation has succeeded.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The operation has failed.`,
				},
				resource.Attribute{
					Name:        "RollBackInitiated",
					Description: `The rollback has been inititiated for import failure.`,
				},
				resource.Attribute{
					Name:        "RollBackFailed",
					Description: `The rollback has failed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackCompleted",
					Description: `The rollback has completed for import failure.`,
				},
				resource.Attribute{
					Name:        "RollbackAborted",
					Description: `The rollback has been aborted for import failure.`,
				},
				resource.Attribute{
					Name:        "OperationTimedOut",
					Description: `The operation has timed out.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_connectorpack_connector_pack_upgrade",
			Category:         "connectorpack",
			ShortDescription: `Used to download or install connector packs on the target device.`,
			Description: `
Used to download or install connector packs on the target device.
`,
			Keywords: []string{
				"connectorpack",
				"connector",
				"pack",
				"upgrade",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Install",
					Description: `Installs the requisite connector packs on UCS Director.`,
				},
				resource.Attribute{
					Name:        "Push",
					Description: `Pushes the requisite connector packs to UCS Director.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_connectorpack_upgrade_impact",
			Category:         "connectorpack",
			ShortDescription: `Used to determine the list of connector packs to be installed on a target UCS Director in its next upgrade cycle. Accepts the moid of the target UcsdInfo as part of the filter query. Given below is a sample url :- https://{{target}}/api/v1/connectorpack/UpgradeImpacts?$filter= ( UcsdInfo.Moid eq <<MoId>> ).`,
			Description: `
Used to determine the list of connector packs to be installed on a target UCS Director in its next upgrade cycle. Accepts the moid of the target UcsdInfo as part of the filter query. Given below is a sample url :- https://{{target}}/api/v1/connectorpack/UpgradeImpacts?$filter= ( UcsdInfo.Moid eq <<MoId>> ).
`,
			Keywords: []string{
				"connectorpack",
				"upgrade",
				"impact",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_deviceconnector_policy",
			Category:         "deviceconnector",
			ShortDescription: `Policy to control configuration changes allowed from Cisco IMC.`,
			Description: `
Policy to control configuration changes allowed from Cisco IMC.
`,
			Keywords: []string{
				"deviceconnector",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_chassis",
			Category:         "equipment",
			ShortDescription: `A physical holder housing blade servers.`,
			Description: `
A physical holder housing blade servers.
`,
			Keywords: []string{
				"equipment",
				"chassis",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IntersightStandalone",
					Description: `Intersight Standalone mode of operation.`,
				},
				resource.Attribute{
					Name:        "UCSM",
					Description: `Unified Computing System Manager mode of operation.`,
				},
				resource.Attribute{
					Name:        "Intersight",
					Description: `Intersight managed mode of operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_chassis_identity",
			Category:         "equipment",
			ShortDescription: `ChassisIdentity Object conatains connectivity information about IOMs of the chassis. ChassisID is uniquely allocated for the combination of vendor, model and serial number of the chassis.`,
			Description: `
ChassisIdentity Object conatains connectivity information about IOMs of the chassis. ChassisID is uniquely allocated for the combination of vendor, model and serial number of the chassis.
`,
			Keywords: []string{
				"equipment",
				"chassis",
				"identity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation value for maintenance actions on an equipment.`,
				},
				resource.Attribute{
					Name:        "Decommission",
					Description: `Decommission the equipment and temporarily remove it from being managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "Recommission",
					Description: `Recommission the equipment.`,
				},
				resource.Attribute{
					Name:        "Reack",
					Description: `Reacknowledge the equipment and discover it again.`,
				},
				resource.Attribute{
					Name:        "Remove",
					Description: `Remove the equipment permanently from Intersight management.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Default state of an equipment. This should be an initial state when no state is defined for an equipment.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `Default Lifecycle State for a physical entity.`,
				},
				resource.Attribute{
					Name:        "Decommissioned",
					Description: `Decommission Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DecommissionInProgress",
					Description: `Decommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RecommissionInProgress",
					Description: `Recommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "OperationFailed",
					Description: `Failed Operation Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "ReackInProgress",
					Description: `ReackInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RemoveInProgress",
					Description: `RemoveInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "Discovered",
					Description: `Discovered Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryInProgress",
					Description: `DiscoveryInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryFailed",
					Description: `DiscoveryFailed Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "FirmwareUpgradeInProgress",
					Description: `Firmware upgrade is in progress on given physical entity.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_chassis_operation",
			Category:         "equipment",
			ShortDescription: `Models the configurable properties of Chassis.`,
			Description: `
Models the configurable properties of Chassis.
`,
			Keywords: []string{
				"equipment",
				"chassis",
				"operation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation action for the Locator Led of an equipment.`,
				},
				resource.Attribute{
					Name:        "TurnOn",
					Description: `Turn on the Locator Led of an equipment.`,
				},
				resource.Attribute{
					Name:        "TurnOff",
					Description: `Turn off the Locator Led of an equipment.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_device_summary",
			Category:         "equipment",
			ShortDescription: `Aggregation of properties pertaining to different inventory MOs.`,
			Description: `
Aggregation of properties pertaining to different inventory MOs.
`,
			Keywords: []string{
				"equipment",
				"device",
				"summary",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_fan",
			Category:         "equipment",
			ShortDescription: `Fan in a Fabric Interconnect / Chassis / RackUnit.`,
			Description: `
Fan in a Fabric Interconnect / Chassis / RackUnit.
`,
			Keywords: []string{
				"equipment",
				"fan",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_fan_module",
			Category:         "equipment",
			ShortDescription: `This represents Fan module housing multiple fans for chassis/server.`,
			Description: `
This represents Fan module housing multiple fans for chassis/server.
`,
			Keywords: []string{
				"equipment",
				"fan",
				"module",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_fex",
			Category:         "equipment",
			ShortDescription: `Fabric Extender which can mutiplex traffic from the host facing ports.`,
			Description: `
Fabric Extender which can mutiplex traffic from the host facing ports.
`,
			Keywords: []string{
				"equipment",
				"fex",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_fex_identity",
			Category:         "equipment",
			ShortDescription: `FexIdentity Object conatains basic information of fabric extender. moduleId is uniquely allocated for the combination of vendor, model and serial number of the chassis.`,
			Description: `
FexIdentity Object conatains basic information of fabric extender. moduleId is uniquely allocated for the combination of vendor, model and serial number of the chassis.
`,
			Keywords: []string{
				"equipment",
				"fex",
				"identity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation value for maintenance actions on an equipment.`,
				},
				resource.Attribute{
					Name:        "Decommission",
					Description: `Decommission the equipment and temporarily remove it from being managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "Recommission",
					Description: `Recommission the equipment.`,
				},
				resource.Attribute{
					Name:        "Reack",
					Description: `Reacknowledge the equipment and discover it again.`,
				},
				resource.Attribute{
					Name:        "Remove",
					Description: `Remove the equipment permanently from Intersight management.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Default state of an equipment. This should be an initial state when no state is defined for an equipment.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `Default Lifecycle State for a physical entity.`,
				},
				resource.Attribute{
					Name:        "Decommissioned",
					Description: `Decommission Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DecommissionInProgress",
					Description: `Decommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RecommissionInProgress",
					Description: `Recommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "OperationFailed",
					Description: `Failed Operation Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "ReackInProgress",
					Description: `ReackInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RemoveInProgress",
					Description: `RemoveInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "Discovered",
					Description: `Discovered Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryInProgress",
					Description: `DiscoveryInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryFailed",
					Description: `DiscoveryFailed Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "FirmwareUpgradeInProgress",
					Description: `Firmware upgrade is in progress on given physical entity.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_fex_operation",
			Category:         "equipment",
			ShortDescription: `Models the configuration states of a FEX in Intersight.`,
			Description: `
Models the configuration states of a FEX in Intersight.
`,
			Keywords: []string{
				"equipment",
				"fex",
				"operation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation action for the Locator Led of an equipment.`,
				},
				resource.Attribute{
					Name:        "TurnOn",
					Description: `Turn on the Locator Led of an equipment.`,
				},
				resource.Attribute{
					Name:        "TurnOff",
					Description: `Turn off the Locator Led of an equipment.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_identity_summary",
			Category:         "equipment",
			ShortDescription: `Consolidated view of all equipment identities.`,
			Description: `
Consolidated view of all equipment identities.
`,
			Keywords: []string{
				"equipment",
				"identity",
				"summary",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No operation value for maintenance actions on an equipment.`,
				},
				resource.Attribute{
					Name:        "Decommission",
					Description: `Decommission the equipment and temporarily remove it from being managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "Recommission",
					Description: `Recommission the equipment.`,
				},
				resource.Attribute{
					Name:        "Reack",
					Description: `Reacknowledge the equipment and discover it again.`,
				},
				resource.Attribute{
					Name:        "Remove",
					Description: `Remove the equipment permanently from Intersight management.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The running firmware version is unknown.`,
				},
				resource.Attribute{
					Name:        "Supported",
					Description: `The running firmware version is known and supports IMM mode.`,
				},
				resource.Attribute{
					Name:        "NotSupported",
					Description: `The running firmware version is known and does not support IMM mode.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Default state of an equipment. This should be an initial state when no state is defined for an equipment.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `Default Lifecycle State for a physical entity.`,
				},
				resource.Attribute{
					Name:        "Decommissioned",
					Description: `Decommission Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DecommissionInProgress",
					Description: `Decommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RecommissionInProgress",
					Description: `Recommission Inprogress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "OperationFailed",
					Description: `Failed Operation Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "ReackInProgress",
					Description: `ReackInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "RemoveInProgress",
					Description: `RemoveInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "Discovered",
					Description: `Discovered Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryInProgress",
					Description: `DiscoveryInProgress Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "DiscoveryFailed",
					Description: `DiscoveryFailed Lifecycle state.`,
				},
				resource.Attribute{
					Name:        "FirmwareUpgradeInProgress",
					Description: `Firmware upgrade is in progress on given physical entity.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The default presence state.`,
				},
				resource.Attribute{
					Name:        "Equipped",
					Description: `The server is equipped in the slot.`,
				},
				resource.Attribute{
					Name:        "EquippedMismatch",
					Description: `The slot is equipped, but there is another server currently inventoried in the slot.`,
				},
				resource.Attribute{
					Name:        "Missing",
					Description: `The server is not present in the given slot.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_io_card",
			Category:         "equipment",
			ShortDescription: `I/O module on a chassis which multiplexes traffic from blade servers.`,
			Description: `
I/O module on a chassis which multiplexes traffic from blade servers.
`,
			Keywords: []string{
				"equipment",
				"io",
				"card",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_io_card_operation",
			Category:         "equipment",
			ShortDescription: `Models the configurable properties of a iomodule in Intersight.`,
			Description: `
Models the configurable properties of a iomodule in Intersight.
`,
			Keywords: []string{
				"equipment",
				"io",
				"card",
				"operation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `Placeholder default value for iom power state property.`,
				},
				resource.Attribute{
					Name:        "Reboot",
					Description: `IO Module reboot state property value.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_io_expander",
			Category:         "equipment",
			ShortDescription: `I/O expander card which is used as an extension for servers in a Colusa Chassis.`,
			Description: `
I/O expander card which is used as an extension for servers in a Colusa Chassis.
`,
			Keywords: []string{
				"equipment",
				"io",
				"expander",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_locator_led",
			Category:         "equipment",
			ShortDescription: `Locator Led of an Equipment like Rack, Disk etc.`,
			Description: `
Locator Led of an Equipment like Rack, Disk etc.
`,
			Keywords: []string{
				"equipment",
				"locator",
				"led",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_psu",
			Category:         "equipment",
			ShortDescription: `This represents power supply unit for chassis/server.`,
			Description: `
This represents power supply unit for chassis/server.
`,
			Keywords: []string{
				"equipment",
				"psu",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_psu_control",
			Category:         "equipment",
			ShortDescription: `This represents the power states of an equipment.`,
			Description: `
This represents the power states of an equipment.
`,
			Keywords: []string{
				"equipment",
				"psu",
				"control",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_rack_enclosure",
			Category:         "equipment",
			ShortDescription: `A physical holder housing rack servers.`,
			Description: `
A physical holder housing rack servers.
`,
			Keywords: []string{
				"equipment",
				"rack",
				"enclosure",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_rack_enclosure_slot",
			Category:         "equipment",
			ShortDescription: `Rack Server Slot in a RackEnclosure.`,
			Description: `
Rack Server Slot in a RackEnclosure.
`,
			Keywords: []string{
				"equipment",
				"rack",
				"enclosure",
				"slot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_shared_io_module",
			Category:         "equipment",
			ShortDescription: `I/O Controller present inside SIOC to provide data path from S-series server to FI.`,
			Description: `
I/O Controller present inside SIOC to provide data path from S-series server to FI.
`,
			Keywords: []string{
				"equipment",
				"shared",
				"io",
				"module",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_switch_card",
			Category:         "equipment",
			ShortDescription: `Fixed / Removable module on a Fabric Interconnect / Switch.`,
			Description: `
Fixed / Removable module on a Fabric Interconnect / Switch.
`,
			Keywords: []string{
				"equipment",
				"switch",
				"card",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.`,
				},
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `The default state of the sensor (in case no data is received).`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `State of the sensor indicating the sensor's temperature range is okay.`,
				},
				resource.Attribute{
					Name:        "upper-non-recoverable",
					Description: `State of the sensor indicating that the temperature is extremely high above normal range.`,
				},
				resource.Attribute{
					Name:        "upper-critical",
					Description: `State of the sensor indicating that the temperature is above normal range.`,
				},
				resource.Attribute{
					Name:        "upper-non-critical",
					Description: `State of the sensor indicating that the temperature is a little above the normal range.`,
				},
				resource.Attribute{
					Name:        "lower-non-critical",
					Description: `State of the sensor indicating that the temperature is a little below the normal range.`,
				},
				resource.Attribute{
					Name:        "lower-critical",
					Description: `State of the sensor indicating that the temperature is below normal range.`,
				},
				resource.Attribute{
					Name:        "lower-non-recoverable",
					Description: `State of the sensor indicating that the temperature is extremely below normal range.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_system_io_controller",
			Category:         "equipment",
			ShortDescription: `I/O Controller on a chassis which provides the data path to S-series server.`,
			Description: `
I/O Controller on a chassis which provides the data path to S-series server.
`,
			Keywords: []string{
				"equipment",
				"system",
				"io",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_tpm",
			Category:         "equipment",
			ShortDescription: `TPM security chip on server board.`,
			Description: `
TPM security chip on server board.
`,
			Keywords: []string{
				"equipment",
				"tpm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_equipment_transceiver",
			Category:         "equipment",
			ShortDescription: `Transceiver information on the chassis.`,
			Description: `
Transceiver information on the chassis.
`,
			Keywords: []string{
				"equipment",
				"transceiver",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ether_host_port",
			Category:         "ether",
			ShortDescription: `Model object contains the details of host port available on IO card or fabric extender.`,
			Description: `
Model object contains the details of host port available on IO card or fabric extender.
`,
			Keywords: []string{
				"ether",
				"host",
				"port",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ether_network_port",
			Category:         "ether",
			ShortDescription: `Model contains the details of the ethernet port connected to the FI side.`,
			Description: `
Model contains the details of the ethernet port connected to the FI side.
`,
			Keywords: []string{
				"ether",
				"network",
				"port",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ether_physical_port",
			Category:         "ether",
			ShortDescription: `Physical ethernet port present on a FI.`,
			Description: `
Physical ethernet port present on a FI.
`,
			Keywords: []string{
				"ether",
				"physical",
				"port",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ether_port_channel",
			Category:         "ether",
			ShortDescription: `Model contains the details of the ethernet port-channels configured on the FI.`,
			Description: `
Model contains the details of the ethernet port-channels configured on the FI.
`,
			Keywords: []string{
				"ether",
				"port",
				"channel",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_externalsite_authorization",
			Category:         "externalsite",
			ShortDescription: `An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.`,
			Description: `
An authentication request that will be used to get authorized from external repository like cisco.com in order to download the image. This MO creation is a one time configuration before calling firmware.Upgrade API. This MO has to be modified with updated details whenever the user has updated the credentials in external repository.
`,
			Keywords: []string{
				"externalsite",
				"authorization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cisco",
					Description: `Cisco as an external site from where the resources like image will be downloaded.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_appliance_pc_role",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create an appliance port channel.`,
			Description: `
Configuration object sent by user to create an appliance port channel.
`,
			Keywords: []string{
				"fabric",
				"appliance",
				"pc",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "1Gbps",
					Description: `Admin configurable speed 1Gbps.`,
				},
				resource.Attribute{
					Name:        "10Gbps",
					Description: `Admin configurable speed 10Gbps.`,
				},
				resource.Attribute{
					Name:        "25Gbps",
					Description: `Admin configurable speed 25Gbps.`,
				},
				resource.Attribute{
					Name:        "40Gbps",
					Description: `Admin configurable speed 40Gbps.`,
				},
				resource.Attribute{
					Name:        "100Gbps",
					Description: `Admin configurable speed 100Gbps.`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk Mode Switch Port Type.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access Mode Switch Port Type.`,
				},
				resource.Attribute{
					Name:        "Best Effort",
					Description: `QoS Priority for Best-effort traffic.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `QoS Priority for FC traffic.`,
				},
				resource.Attribute{
					Name:        "Platinum",
					Description: `QoS Priority for Platinum traffic.`,
				},
				resource.Attribute{
					Name:        "Gold",
					Description: `QoS Priority for Gold traffic.`,
				},
				resource.Attribute{
					Name:        "Silver",
					Description: `QoS Priority for Silver traffic.`,
				},
				resource.Attribute{
					Name:        "Bronze",
					Description: `QoS Priority for Bronze traffic.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_appliance_role",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create an appliance port.`,
			Description: `
Configuration object sent by user to create an appliance port.
`,
			Keywords: []string{
				"fabric",
				"appliance",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "1Gbps",
					Description: `Admin configurable speed 1Gbps.`,
				},
				resource.Attribute{
					Name:        "10Gbps",
					Description: `Admin configurable speed 10Gbps.`,
				},
				resource.Attribute{
					Name:        "25Gbps",
					Description: `Admin configurable speed 25Gbps.`,
				},
				resource.Attribute{
					Name:        "40Gbps",
					Description: `Admin configurable speed 40Gbps.`,
				},
				resource.Attribute{
					Name:        "100Gbps",
					Description: `Admin configurable speed 100Gbps.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Forward error correction option 'Auto'.`,
				},
				resource.Attribute{
					Name:        "Cl91",
					Description: `Forward error correction option 'cl91'.`,
				},
				resource.Attribute{
					Name:        "Cl74",
					Description: `Forward error correction option 'cl74'.`,
				},
				resource.Attribute{
					Name:        "trunk",
					Description: `Trunk Mode Switch Port Type.`,
				},
				resource.Attribute{
					Name:        "access",
					Description: `Access Mode Switch Port Type.`,
				},
				resource.Attribute{
					Name:        "Best Effort",
					Description: `QoS Priority for Best-effort traffic.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `QoS Priority for FC traffic.`,
				},
				resource.Attribute{
					Name:        "Platinum",
					Description: `QoS Priority for Platinum traffic.`,
				},
				resource.Attribute{
					Name:        "Gold",
					Description: `QoS Priority for Gold traffic.`,
				},
				resource.Attribute{
					Name:        "Silver",
					Description: `QoS Priority for Silver traffic.`,
				},
				resource.Attribute{
					Name:        "Bronze",
					Description: `QoS Priority for Bronze traffic.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_config_change_detail",
			Category:         "fabric",
			ShortDescription: `This captures details of configuration changes.`,
			Description: `
This captures details of configuration changes.
`,
			Keywords: []string{
				"fabric",
				"config",
				"change",
				"detail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Pending-changes",
					Description: `Config change flag represents changes are due to not deployed changes from Intersight.`,
				},
				resource.Attribute{
					Name:        "Drift-changes",
					Description: `Config change flag represents changes are due to endpoint configuration changes.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration.`,
				},
				resource.Attribute{
					Name:        "Created",
					Description: `The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet.`,
				},
				resource.Attribute{
					Name:        "Modified",
					Description: `The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully.`,
				},
				resource.Attribute{
					Name:        "Deleted",
					Description: `The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_config_result",
			Category:         "fabric",
			ShortDescription: `This provides overall state and detailed information for the deploy and validation profile configuration results.`,
			Description: `
This provides overall state and detailed information for the deploy and validation profile configuration results.
`,
			Keywords: []string{
				"fabric",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_config_result_entry",
			Category:         "fabric",
			ShortDescription: `This provides detailed information for the deploy and validation profile configuration results.`,
			Description: `
This provides detailed information for the deploy and validation profile configuration results.
`,
			Keywords: []string{
				"fabric",
				"config",
				"result",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_eth_network_control_policy",
			Category:         "fabric",
			ShortDescription: `The features that are applied on a vethernet that is connected to the vNIC.`,
			Description: `
The features that are applied on a vethernet that is connected to the vNIC.
`,
			Keywords: []string{
				"fabric",
				"eth",
				"network",
				"control",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "allow",
					Description: `Allows mac forging on an interface.`,
				},
				resource.Attribute{
					Name:        "deny",
					Description: `Denies mac forging on an interface.`,
				},
				resource.Attribute{
					Name:        "nativeVlanOnly",
					Description: `Register only the MAC addresses learnt on the native VLAN.`,
				},
				resource.Attribute{
					Name:        "allVlans",
					Description: `Register all the MAC addresses learnt on all the allowed VLANs.`,
				},
				resource.Attribute{
					Name:        "linkDown",
					Description: `The vethernet will go down in case a suitable uplink is not pinned.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `The vethernet will remain up even if a suitable uplink is not pinned.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_eth_network_group_policy",
			Category:         "fabric",
			ShortDescription: `The allowed VLAN/s on an interface.`,
			Description: `
The allowed VLAN/s on an interface.
`,
			Keywords: []string{
				"fabric",
				"eth",
				"network",
				"group",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_eth_network_policy",
			Category:         "fabric",
			ShortDescription: `A policy for all the Virtual LAN networks to be deployed on the Fabric Interconnect.`,
			Description: `
A policy for all the Virtual LAN networks to be deployed on the Fabric Interconnect.
`,
			Keywords: []string{
				"fabric",
				"eth",
				"network",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_fc_network_policy",
			Category:         "fabric",
			ShortDescription: `A policy for all the Virtual SAN networks to be deployed on the Fabric Interconnect.`,
			Description: `
A policy for all the Virtual SAN networks to be deployed on the Fabric Interconnect.
`,
			Keywords: []string{
				"fabric",
				"fc",
				"network",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_fc_uplink_pc_role",
			Category:         "fabric",
			ShortDescription: `Object sent by user to configure a fc uplink port-channel on the collection of ports.`,
			Description: `
Object sent by user to configure a fc uplink port-channel on the collection of ports.
`,
			Keywords: []string{
				"fabric",
				"fc",
				"uplink",
				"pc",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "8Gbps",
					Description: `Admin configurable speed 8Gbps.`,
				},
				resource.Attribute{
					Name:        "16Gbps",
					Description: `Admin configurable speed 16Gbps.`,
				},
				resource.Attribute{
					Name:        "32Gbps",
					Description: `Admin configurable speed 32Gbps.`,
				},
				resource.Attribute{
					Name:        "Idle",
					Description: `Fc Fill Pattern type Idle.`,
				},
				resource.Attribute{
					Name:        "Arbff",
					Description: `Fc Fill Pattern type Arbff.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_fc_uplink_role",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create a fc uplink port.`,
			Description: `
Configuration object sent by user to create a fc uplink port.
`,
			Keywords: []string{
				"fabric",
				"fc",
				"uplink",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "8Gbps",
					Description: `Admin configurable speed 8Gbps.`,
				},
				resource.Attribute{
					Name:        "16Gbps",
					Description: `Admin configurable speed 16Gbps.`,
				},
				resource.Attribute{
					Name:        "32Gbps",
					Description: `Admin configurable speed 32Gbps.`,
				},
				resource.Attribute{
					Name:        "Idle",
					Description: `Fc Fill Pattern type Idle.`,
				},
				resource.Attribute{
					Name:        "Arbff",
					Description: `Fc Fill Pattern type Arbff.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_fcoe_uplink_pc_role",
			Category:         "fabric",
			ShortDescription: `Object sent by user to configure a fcoe uplink port-channel on the collection of ports.`,
			Description: `
Object sent by user to configure a fcoe uplink port-channel on the collection of ports.
`,
			Keywords: []string{
				"fabric",
				"fcoe",
				"uplink",
				"pc",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "1Gbps",
					Description: `Admin configurable speed 1Gbps.`,
				},
				resource.Attribute{
					Name:        "10Gbps",
					Description: `Admin configurable speed 10Gbps.`,
				},
				resource.Attribute{
					Name:        "25Gbps",
					Description: `Admin configurable speed 25Gbps.`,
				},
				resource.Attribute{
					Name:        "40Gbps",
					Description: `Admin configurable speed 40Gbps.`,
				},
				resource.Attribute{
					Name:        "100Gbps",
					Description: `Admin configurable speed 100Gbps.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_fcoe_uplink_role",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create a fcoe uplink port.`,
			Description: `
Configuration object sent by user to create a fcoe uplink port.
`,
			Keywords: []string{
				"fabric",
				"fcoe",
				"uplink",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "1Gbps",
					Description: `Admin configurable speed 1Gbps.`,
				},
				resource.Attribute{
					Name:        "10Gbps",
					Description: `Admin configurable speed 10Gbps.`,
				},
				resource.Attribute{
					Name:        "25Gbps",
					Description: `Admin configurable speed 25Gbps.`,
				},
				resource.Attribute{
					Name:        "40Gbps",
					Description: `Admin configurable speed 40Gbps.`,
				},
				resource.Attribute{
					Name:        "100Gbps",
					Description: `Admin configurable speed 100Gbps.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Forward error correction option 'Auto'.`,
				},
				resource.Attribute{
					Name:        "Cl91",
					Description: `Forward error correction option 'cl91'.`,
				},
				resource.Attribute{
					Name:        "Cl74",
					Description: `Forward error correction option 'cl74'.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_multicast_policy",
			Category:         "fabric",
			ShortDescription: `A policy to configure Multicast settings for all the Virtual LAN networks.`,
			Description: `
A policy to configure Multicast settings for all the Virtual LAN networks.
`,
			Keywords: []string{
				"fabric",
				"multicast",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Disabled",
					Description: `IGMP Querier Disabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `IGMP Querier Enabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `IGMP Snooping Enabled State.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `IGMP Snooping Disabled State.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_pc_member",
			Category:         "fabric",
			ShortDescription: `PcMember object is to establish the relationship between port parameters and pcId.`,
			Description: `
PcMember object is to establish the relationship between port parameters and pcId.
`,
			Keywords: []string{
				"fabric",
				"pc",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_pc_operation",
			Category:         "fabric",
			ShortDescription: `PcOperation objects allows the user to alter the state of the port channel.`,
			Description: `
PcOperation objects allows the user to alter the state of the port channel.
`,
			Keywords: []string{
				"fabric",
				"pc",
				"operation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_port_mode",
			Category:         "fabric",
			ShortDescription: `Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.`,
			Description: `
Object sent by user to configure range of unified ports as FC/Ethernet or ports as breakout.
`,
			Keywords: []string{
				"fabric",
				"port",
				"mode",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "FibreChannel",
					Description: `Fibre Channel Port Types.`,
				},
				resource.Attribute{
					Name:        "BreakoutEthernet10G",
					Description: `Breakout Ethernet 10G Port Type.`,
				},
				resource.Attribute{
					Name:        "BreakoutEthernet25G",
					Description: `Breakout Ethernet 25G Port Type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_port_operation",
			Category:         "fabric",
			ShortDescription: `PortOperation objects allows the user to alter the state of the port.`,
			Description: `
PortOperation objects allows the user to alter the state of the port.
`,
			Keywords: []string{
				"fabric",
				"port",
				"operation",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Nil value when no action has been triggered by the user.`,
				},
				resource.Attribute{
					Name:        "Applied",
					Description: `User configured settings are in applied state.`,
				},
				resource.Attribute{
					Name:        "Applying",
					Description: `User settings are being applied on the target server.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `User configured settings could not be applied.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_port_policy",
			Category:         "fabric",
			ShortDescription: `A policy for all the physical ports of the Fabric Interconnect.`,
			Description: `
A policy for all the physical ports of the Fabric Interconnect.
`,
			Keywords: []string{
				"fabric",
				"port",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UCS-FI-6454",
					Description: `The standard 4th generation UCS Fabric Interconnect with 54 ports.`,
				},
				resource.Attribute{
					Name:        "UCS-FI-64108",
					Description: `The expanded 4th generation UCS Fabric Interconnect with 108 ports.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `Unknown device type, usage is TBD.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_server_role",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create a server port.`,
			Description: `
Configuration object sent by user to create a server port.
`,
			Keywords: []string{
				"fabric",
				"server",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_switch_cluster_profile",
			Category:         "fabric",
			ShortDescription: `This specifies the configuration policies for a cluster of switches.`,
			Description: `
This specifies the configuration policies for a cluster of switches.
`,
			Keywords: []string{
				"fabric",
				"switch",
				"cluster",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_switch_control_policy",
			Category:         "fabric",
			ShortDescription: `A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.`,
			Description: `
A policy to configure the Switching Mode, Port VLAN Optimization, MAC Aging Time, Reserved VLAN Range of the FI.
`,
			Keywords: []string{
				"fabric",
				"switch",
				"control",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_switch_profile",
			Category:         "fabric",
			ShortDescription: `This specifies configuration policies for a managed network switch.`,
			Description: `
This specifies configuration policies for a managed network switch.
`,
			Keywords: []string{
				"fabric",
				"switch",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_system_qos_policy",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to setup Quality of Service (QoS) for this switch.`,
			Description: `
Configuration object sent by user to setup Quality of Service (QoS) for this switch.
`,
			Keywords: []string{
				"fabric",
				"system",
				"qos",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_uplink_pc_role",
			Category:         "fabric",
			ShortDescription: `Object sent by user to configure a ethernet uplink port-channel on the collection of ports.`,
			Description: `
Object sent by user to configure a ethernet uplink port-channel on the collection of ports.
`,
			Keywords: []string{
				"fabric",
				"uplink",
				"pc",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "1Gbps",
					Description: `Admin configurable speed 1Gbps.`,
				},
				resource.Attribute{
					Name:        "10Gbps",
					Description: `Admin configurable speed 10Gbps.`,
				},
				resource.Attribute{
					Name:        "25Gbps",
					Description: `Admin configurable speed 25Gbps.`,
				},
				resource.Attribute{
					Name:        "40Gbps",
					Description: `Admin configurable speed 40Gbps.`,
				},
				resource.Attribute{
					Name:        "100Gbps",
					Description: `Admin configurable speed 100Gbps.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_uplink_role",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create a uplink port.`,
			Description: `
Configuration object sent by user to create a uplink port.
`,
			Keywords: []string{
				"fabric",
				"uplink",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Auto",
					Description: `Admin configurable speed AUTO ( default ).`,
				},
				resource.Attribute{
					Name:        "1Gbps",
					Description: `Admin configurable speed 1Gbps.`,
				},
				resource.Attribute{
					Name:        "10Gbps",
					Description: `Admin configurable speed 10Gbps.`,
				},
				resource.Attribute{
					Name:        "25Gbps",
					Description: `Admin configurable speed 25Gbps.`,
				},
				resource.Attribute{
					Name:        "40Gbps",
					Description: `Admin configurable speed 40Gbps.`,
				},
				resource.Attribute{
					Name:        "100Gbps",
					Description: `Admin configurable speed 100Gbps.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Forward error correction option 'Auto'.`,
				},
				resource.Attribute{
					Name:        "Cl91",
					Description: `Forward error correction option 'cl91'.`,
				},
				resource.Attribute{
					Name:        "Cl74",
					Description: `Forward error correction option 'cl74'.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_vlan",
			Category:         "fabric",
			ShortDescription: `Configuration object for Virtual LAN.`,
			Description: `
Configuration object for Virtual LAN.
`,
			Keywords: []string{
				"fabric",
				"vlan",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fabric_vsan",
			Category:         "fabric",
			ShortDescription: `Configuration object sent by user to create VSAN configurations.`,
			Description: `
Configuration object sent by user to create VSAN configurations.
`,
			Keywords: []string{
				"fabric",
				"vsan",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fault_instance",
			Category:         "fault",
			ShortDescription: `An endpoint anomaly is represented by this object.`,
			Description: `
An endpoint anomaly is represented by this object.
`,
			Keywords: []string{
				"fault",
				"instance",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fc_physical_port",
			Category:         "fc",
			ShortDescription: `Physical fibre channel port present on a FI.`,
			Description: `
Physical fibre channel port present on a FI.
`,
			Keywords: []string{
				"fc",
				"physical",
				"port",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fc_port_channel",
			Category:         "fc",
			ShortDescription: `Model contains the details of the ethernet port-channels configured on the FI.`,
			Description: `
Model contains the details of the ethernet port-channels configured on the FI.
`,
			Keywords: []string{
				"fc",
				"port",
				"channel",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fcpool_fc_block",
			Category:         "fcpool",
			ShortDescription: `A block of contiguous WWN addresses that are part of a pool.`,
			Description: `
A block of contiguous WWN addresses that are part of a pool.
`,
			Keywords: []string{
				"fcpool",
				"fc",
				"block",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fcpool_lease",
			Category:         "fcpool",
			ShortDescription: `Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.`,
			Description: `
Lease represents a single WWN ID that is part of the universe, allocated either from a pool or through static assignment.
`,
			Keywords: []string{
				"fcpool",
				"lease",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dynamic",
					Description: `Identifiers to be allocated by system.`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Identifiers are assigned by the user.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fcpool_pool",
			Category:         "fcpool",
			ShortDescription: `Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.`,
			Description: `
Pool represents a collection of WWN addresses that can be allocated to VHBAs of a server profile.
`,
			Keywords: []string{
				"fcpool",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sequential",
					Description: `Identifiers are assigned in a sequential order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Assignment order is decided by the system.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fcpool_pool_member",
			Category:         "fcpool",
			ShortDescription: `PoolMember represents a single WWN ID that is part of a pool.`,
			Description: `
PoolMember represents a single WWN ID that is part of a pool.
`,
			Keywords: []string{
				"fcpool",
				"pool",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_fcpool_universe",
			Category:         "fcpool",
			ShortDescription: `Universe represents a book keeping container to keep track of all IDs for a given account and pool type.`,
			Description: `
Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
`,
			Keywords: []string{
				"fcpool",
				"universe",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_bios_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a BIOS.`,
			Description: `
Descriptor to uniquely identify a BIOS.
`,
			Keywords: []string{
				"firmware",
				"bios",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_board_controller_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a board controller.`,
			Description: `
Descriptor to uniquely identify a board controller.
`,
			Keywords: []string{
				"firmware",
				"board",
				"controller",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_chassis_upgrade",
			Category:         "firmware",
			ShortDescription: `Firmware upgrade operation for chassis that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.`,
			Description: `
Firmware upgrade operation for chassis that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
`,
			Keywords: []string{
				"firmware",
				"chassis",
				"upgrade",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NONE",
					Description: `Upgrade status is not populated.`,
				},
				resource.Attribute{
					Name:        "IN_PROGRESS",
					Description: `The upgrade is in progress.`,
				},
				resource.Attribute{
					Name:        "SUCCESSFUL",
					Description: `The upgrade successfully completed.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `The upgrade shows failed status.`,
				},
				resource.Attribute{
					Name:        "TERMINATED",
					Description: `The upgrade has been terminated.`,
				},
				resource.Attribute{
					Name:        "direct_upgrade",
					Description: `Upgrade mode is direct download.`,
				},
				resource.Attribute{
					Name:        "network_upgrade",
					Description: `Upgrade mode is network upgrade.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_cimc_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a Cisco IMC.`,
			Description: `
Descriptor to uniquely identify a Cisco IMC.
`,
			Keywords: []string{
				"firmware",
				"cimc",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_dimm_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a DIMM.`,
			Description: `
Descriptor to uniquely identify a DIMM.
`,
			Keywords: []string{
				"firmware",
				"dimm",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_distributable",
			Category:         "firmware",
			ShortDescription: `An image distributed by Cisco.`,
			Description: `
An image distributed by Cisco.
`,
			Keywords: []string{
				"firmware",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
				resource.Attribute{
					Name:        "System",
					Description: `The distributable has been created by the System.`,
				},
				resource.Attribute{
					Name:        "User",
					Description: `The distributable has been created by the User.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_distributable_meta",
			Category:         "firmware",
			ShortDescription: `Meta information for various firmware images stored in the database. Gives information on the particular category for a product.`,
			Description: `
Meta information for various firmware images stored in the database. Gives information on the particular category for a product.
`,
			Keywords: []string{
				"firmware",
				"distributable",
				"meta",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Distributable",
					Description: `Stores firmware host utility images and fabric images.`,
				},
				resource.Attribute{
					Name:        "DriverDistributable",
					Description: `Stores driver distributable images.`,
				},
				resource.Attribute{
					Name:        "ServerConfigurationUtilityDistributable",
					Description: `Stores server configuration utility images.`,
				},
				resource.Attribute{
					Name:        "OperatingSystemFile",
					Description: `Stores operating system iso images.`,
				},
				resource.Attribute{
					Name:        "HyperflexDistributable",
					Description: `It stores HyperFlex images.`,
				},
				resource.Attribute{
					Name:        "Cisco",
					Description: `External repository hosted on cisco.com.`,
				},
				resource.Attribute{
					Name:        "IntersightCloud",
					Description: `Repository hosted by the Intersight Cloud.`,
				},
				resource.Attribute{
					Name:        "LocalMachine",
					Description: `The file is available on the local client machine. Used as an upload source type.`,
				},
				resource.Attribute{
					Name:        "NetworkShare",
					Description: `External repository in the customer datacenter. This will typically be a file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_drive_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a Drive.`,
			Description: `
Descriptor to uniquely identify a Drive.
`,
			Keywords: []string{
				"firmware",
				"drive",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_driver_distributable",
			Category:         "firmware",
			ShortDescription: `A device driver image distributed by Cisco.`,
			Description: `
A device driver image distributed by Cisco.
`,
			Keywords: []string{
				"firmware",
				"driver",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_eula",
			Category:         "firmware",
			ShortDescription: `End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.`,
			Description: `
End User License Agreement (EULA) acceptance status for an account to access cisco.com and download software.
`,
			Keywords: []string{
				"firmware",
				"eula",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_firmware_summary",
			Category:         "firmware",
			ShortDescription: `Update inventory that contains the details for the firmware running on each component in the compute.Physical object.`,
			Description: `
Update inventory that contains the details for the firmware running on each component in the compute.Physical object.
`,
			Keywords: []string{
				"firmware",
				"summary",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_gpu_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a GPU component.`,
			Description: `
Descriptor to uniquely identify a GPU component.
`,
			Keywords: []string{
				"firmware",
				"gpu",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_hba_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a HBA component.`,
			Description: `
Descriptor to uniquely identify a HBA component.
`,
			Keywords: []string{
				"firmware",
				"hba",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_iom_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a IOM component.`,
			Description: `
Descriptor to uniquely identify a IOM component.
`,
			Keywords: []string{
				"firmware",
				"iom",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_mswitch_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a mSwitch component.`,
			Description: `
Descriptor to uniquely identify a mSwitch component.
`,
			Keywords: []string{
				"firmware",
				"mswitch",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_nxos_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a Fabric interconnect.`,
			Description: `
Descriptor to uniquely identify a Fabric interconnect.
`,
			Keywords: []string{
				"firmware",
				"nxos",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_pcie_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a PCIE component.`,
			Description: `
Descriptor to uniquely identify a PCIE component.
`,
			Keywords: []string{
				"firmware",
				"pcie",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_psu_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a PSU component.`,
			Description: `
Descriptor to uniquely identify a PSU component.
`,
			Keywords: []string{
				"firmware",
				"psu",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_running_firmware",
			Category:         "firmware",
			ShortDescription: `Running Firmware on an endpoint.`,
			Description: `
Running Firmware on an endpoint.
`,
			Keywords: []string{
				"firmware",
				"running",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_sas_expander_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a SasExpandar component.`,
			Description: `
Descriptor to uniquely identify a SasExpandar component.
`,
			Keywords: []string{
				"firmware",
				"sas",
				"expander",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_server_configuration_utility_distributable",
			Category:         "firmware",
			ShortDescription: `A server configuration utility image distributed by Cisco.`,
			Description: `
A server configuration utility image distributed by Cisco.
`,
			Keywords: []string{
				"firmware",
				"server",
				"configuration",
				"utility",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_storage_controller_descriptor",
			Category:         "firmware",
			ShortDescription: `Descriptor to uniquely identify a storage controller.`,
			Description: `
Descriptor to uniquely identify a storage controller.
`,
			Keywords: []string{
				"firmware",
				"storage",
				"controller",
				"descriptor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_switch_upgrade",
			Category:         "firmware",
			ShortDescription: `Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.`,
			Description: `
Firmware upgrade operation for Fabric Interconnects that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
`,
			Keywords: []string{
				"firmware",
				"switch",
				"upgrade",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NONE",
					Description: `Upgrade status is not populated.`,
				},
				resource.Attribute{
					Name:        "IN_PROGRESS",
					Description: `The upgrade is in progress.`,
				},
				resource.Attribute{
					Name:        "SUCCESSFUL",
					Description: `The upgrade successfully completed.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `The upgrade shows failed status.`,
				},
				resource.Attribute{
					Name:        "TERMINATED",
					Description: `The upgrade has been terminated.`,
				},
				resource.Attribute{
					Name:        "direct_upgrade",
					Description: `Upgrade mode is direct download.`,
				},
				resource.Attribute{
					Name:        "network_upgrade",
					Description: `Upgrade mode is network upgrade.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_unsupported_version_upgrade",
			Category:         "firmware",
			ShortDescription: `This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.`,
			Description: `
This represents an operation managed object used for upgrading equipment that cannot be discovered due to unsupported firmware. Currently, it only supports blade upgrades.
`,
			Keywords: []string{
				"firmware",
				"unsupported",
				"version",
				"upgrade",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `Upgrade status is none when upgrade is in progress.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Upgrade completed successfully.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Upgrade status is failed when upgrade has failed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_upgrade",
			Category:         "firmware",
			ShortDescription: `Firmware upgrade operation for rack and blade servers that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.`,
			Description: `
Firmware upgrade operation for rack and blade servers that downloads the image located at Cisco/appliance/user provided HTTP repository or uses the image from a network share and upgrade. Direct download is used for upgrade that uses the image from a Cisco repository or an appliance repository. Network share is used for upgrade that use the image from a network share from your data center.
`,
			Keywords: []string{
				"firmware",
				"upgrade",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NONE",
					Description: `Upgrade status is not populated.`,
				},
				resource.Attribute{
					Name:        "IN_PROGRESS",
					Description: `The upgrade is in progress.`,
				},
				resource.Attribute{
					Name:        "SUCCESSFUL",
					Description: `The upgrade successfully completed.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `The upgrade shows failed status.`,
				},
				resource.Attribute{
					Name:        "TERMINATED",
					Description: `The upgrade has been terminated.`,
				},
				resource.Attribute{
					Name:        "direct_upgrade",
					Description: `Upgrade mode is direct download.`,
				},
				resource.Attribute{
					Name:        "network_upgrade",
					Description: `Upgrade mode is network upgrade.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_upgrade_impact_status",
			Category:         "firmware",
			ShortDescription: `Captures the impact for an upgrade.`,
			Description: `
Captures the impact for an upgrade.
`,
			Keywords: []string{
				"firmware",
				"upgrade",
				"impact",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Inprogress",
					Description: `Upgrade impact calculation is in progress.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Upgrade impact calculation is completed.`,
				},
				resource.Attribute{
					Name:        "Unavailable",
					Description: `Upgrade impact is not available since image is not present in FI.`,
				},
				resource.Attribute{
					Name:        "Basic",
					Description: `Summary of a single instance involved in the upgrade operation.`,
				},
				resource.Attribute{
					Name:        "Detail",
					Description: `Summary of the collection of single instances for a complex component involved in the upgrade operation. For example, in case of a server upgrade, a detailed summary indicates impact of all the single instances which are part of the server, such as storage controller, drives, and BIOS.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_firmware_upgrade_status",
			Category:         "firmware",
			ShortDescription: `The status for the upgrade operation to include the status for the download and upgrade stages.`,
			Description: `
The status for the upgrade operation to include the status for the download and upgrade stages.
`,
			Keywords: []string{
				"firmware",
				"upgrade",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "none",
					Description: `Server power status is none.`,
				},
				resource.Attribute{
					Name:        "powered on",
					Description: `Server power status is powered on.`,
				},
				resource.Attribute{
					Name:        "powered off",
					Description: `Server power status is powered off.`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `Upgrade stage is no upgrade stage.`,
				},
				resource.Attribute{
					Name:        "started",
					Description: `Upgrade stage is started.`,
				},
				resource.Attribute{
					Name:        "prepare initiating",
					Description: `Upgrade configuration is being prepared.`,
				},
				resource.Attribute{
					Name:        "prepare initiated",
					Description: `Upgrade configuration is initiated.`,
				},
				resource.Attribute{
					Name:        "prepared",
					Description: `Upgrade configuration is prepared.`,
				},
				resource.Attribute{
					Name:        "download initiating",
					Description: `Upgrade stage is download initiating.`,
				},
				resource.Attribute{
					Name:        "download initiated",
					Description: `Upgrade stage is download initiated.`,
				},
				resource.Attribute{
					Name:        "downloading",
					Description: `Upgrade stage is downloading.`,
				},
				resource.Attribute{
					Name:        "downloaded",
					Description: `Upgrade stage is downloaded.`,
				},
				resource.Attribute{
					Name:        "upgrade initiating on fabric A",
					Description: `Upgrade stage is in upgrade initiating when upgrade is being started in endopint.`,
				},
				resource.Attribute{
					Name:        "upgrade initiated on fabric A",
					Description: `Upgrade stage is in upgrade initiated when the upgrade has started in endpoint.`,
				},
				resource.Attribute{
					Name:        "upgrading fabric A",
					Description: `Upgrade stage is in upgrading when the upgrade requires reboot to complete.`,
				},
				resource.Attribute{
					Name:        "rebooting fabric A",
					Description: `Upgrade is in rebooting when the endpoint is being rebooted.`,
				},
				resource.Attribute{
					Name:        "upgraded fabric A",
					Description: `Upgrade stage is in upgraded when the corresponding endpoint has completed.`,
				},
				resource.Attribute{
					Name:        "upgrade initiating on fabric B",
					Description: `Upgrade stage is in upgrade initiating when upgrade is being started in endopint.`,
				},
				resource.Attribute{
					Name:        "upgrade initiated on fabric B",
					Description: `Upgrade stage is in upgrade initiated when upgrade has started in endpoint.`,
				},
				resource.Attribute{
					Name:        "upgrading fabric B",
					Description: `Upgrade stage is in upgrading when the upgrade requires reboot to complete.`,
				},
				resource.Attribute{
					Name:        "rebooting fabric B",
					Description: `Upgrade is in rebooting when the endpoint is being rebooted.`,
				},
				resource.Attribute{
					Name:        "upgraded fabric B",
					Description: `Upgrade stage is in upgraded when the corresponding endpoint has completed.`,
				},
				resource.Attribute{
					Name:        "upgrade initiating",
					Description: `Upgrade stage is upgrade initiating.`,
				},
				resource.Attribute{
					Name:        "upgrade initiated",
					Description: `Upgrade stage is upgrade initiated.`,
				},
				resource.Attribute{
					Name:        "upgrading",
					Description: `Upgrade stage is upgrading.`,
				},
				resource.Attribute{
					Name:        "oob images staging",
					Description: `Out-of-band component images staging.`,
				},
				resource.Attribute{
					Name:        "oob images staged",
					Description: `Out-of-band component images staged.`,
				},
				resource.Attribute{
					Name:        "rebooting",
					Description: `Upgrade is rebooting the endpoint.`,
				},
				resource.Attribute{
					Name:        "upgraded",
					Description: `Upgrade stage is upgraded.`,
				},
				resource.Attribute{
					Name:        "success",
					Description: `Upgrade stage is success.`,
				},
				resource.Attribute{
					Name:        "failed",
					Description: `Upgrade stage is upgrade failed.`,
				},
				resource.Attribute{
					Name:        "terminated",
					Description: `Upgrade stage is terminated.`,
				},
				resource.Attribute{
					Name:        "pending",
					Description: `Upgrade stage is pending.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `The image will be cached into Intersight Appliance or an endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `The image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `The image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `Upgrade pending reason is none.`,
				},
				resource.Attribute{
					Name:        "pending for next reboot",
					Description: `Upgrade pending reason is pending for next reboot.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_forecast_catalog",
			Category:         "forecast",
			ShortDescription: `A catalog for managing forecast settings.`,
			Description: `
A catalog for managing forecast settings.
`,
			Keywords: []string{
				"forecast",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_forecast_definition",
			Category:         "forecast",
			ShortDescription: `Definition for forecast metric settings.`,
			Description: `
Definition for forecast metric settings.
`,
			Keywords: []string{
				"forecast",
				"definition",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_forecast_instance",
			Category:         "forecast",
			ShortDescription: `Entity representing forecast result for instance of managed object, ie, data source.`,
			Description: `
Entity representing forecast result for instance of managed object, ie, data source.
`,
			Keywords: []string{
				"forecast",
				"instance",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_graphics_card",
			Category:         "graphics",
			ShortDescription: `Graphics Card present in a server.`,
			Description: `
Graphics Card present in a server.
`,
			Keywords: []string{
				"graphics",
				"card",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_graphics_controller",
			Category:         "graphics",
			ShortDescription: `Controller for a Graphics Card.`,
			Description: `
Controller for a Graphics Card.
`,
			Keywords: []string{
				"graphics",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hcl_driver_image",
			Category:         "hcl",
			ShortDescription: `Collection used to store the driver ISO urls for each server based on how it is managed.`,
			Description: `
Collection used to store the driver ISO urls for each server based on how it is managed.
`,
			Keywords: []string{
				"hcl",
				"driver",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UCSM",
					Description: `The server is managed by UCS Manager.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `The server is standalone managed by CIMC.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hcl_exempted_catalog",
			Category:         "hcl",
			ShortDescription: `Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.`,
			Description: `
Collection used to store exempted products (ie. adapters, storage controllers, etc). These products should be ignored for HCL validation purposes.
`,
			Keywords: []string{
				"hcl",
				"exempted",
				"catalog",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Adapter",
					Description: `Represents network adapter cards.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `Represents storage controllers.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `Represents graphics cards.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hcl_hyperflex_software_compatibility_info",
			Category:         "hcl",
			ShortDescription: `Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.`,
			Description: `
Lists software compatibility information between different HperFlex component versions like HyperFlex Data Platform, Hypervisor, Drive Firmware, etc.
`,
			Keywords: []string{
				"hcl",
				"hyperflex",
				"software",
				"compatibility",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hcl_operating_system",
			Category:         "hcl",
			ShortDescription: `Collection used to store operating system details.`,
			Description: `
Collection used to store operating system details.
`,
			Keywords: []string{
				"hcl",
				"operating",
				"system",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hcl_operating_system_vendor",
			Category:         "hcl",
			ShortDescription: `Collection used to store operating system vendors details.`,
			Description: `
Collection used to store operating system vendors details.
`,
			Keywords: []string{
				"hcl",
				"operating",
				"system",
				"vendor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_alarm",
			Category:         "hyperflex",
			ShortDescription: `An alarm representing a fault in the HyperFlex cluster configuration or hardware.`,
			Description: `
An alarm representing a fault in the HyperFlex cluster configuration or hardware.
`,
			Keywords: []string{
				"hyperflex",
				"alarm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The type of entity is not known.`,
				},
				resource.Attribute{
					Name:        "DISK",
					Description: `The entity is a physical storage device.`,
				},
				resource.Attribute{
					Name:        "NODE",
					Description: `The entity is a HyperFlex cluster node.`,
				},
				resource.Attribute{
					Name:        "CLUSTER",
					Description: `The entity is the HyperFlex cluster itself.`,
				},
				resource.Attribute{
					Name:        "DATASTORE",
					Description: `The entity is a logical datastore configured on the HyperFlex cluster.`,
				},
				resource.Attribute{
					Name:        "ZONE",
					Description: `The entity is a logical or physical zone configured on the HyperFlex cluster.`,
				},
				resource.Attribute{
					Name:        "VIRTUALMACHINE",
					Description: `The entity is a virtual machine running on the HyperFlex cluster.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The alarm status is not known.`,
				},
				resource.Attribute{
					Name:        "CLEARED",
					Description: `The event that triggered the alarm has been remedied and no longer requires the user's attention.`,
				},
				resource.Attribute{
					Name:        "INFO",
					Description: `The alarm represents a message that does not require the user's immediate attention.`,
				},
				resource.Attribute{
					Name:        "WARNING",
					Description: `The alarm represents a moderate fault.`,
				},
				resource.Attribute{
					Name:        "CRITICAL",
					Description: `The alarm represents a critical fault.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_app_catalog",
			Category:         "hyperflex",
			ShortDescription: `A catalog for managing application settings for HyperFlex cluster configuration service.`,
			Description: `
A catalog for managing application settings for HyperFlex cluster configuration service.
`,
			Keywords: []string{
				"hyperflex",
				"app",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_auto_support_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying the configuration to automatically generate support tickets with Cisco TAC.`,
			Description: `
A policy specifying the configuration to automatically generate support tickets with Cisco TAC.
`,
			Keywords: []string{
				"hyperflex",
				"auto",
				"support",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_backup_cluster",
			Category:         "hyperflex",
			ShortDescription: `BackupCluster object associated with Hyperflex cluster describing the backup related information.`,
			Description: `
BackupCluster object associated with Hyperflex cluster describing the backup related information.
`,
			Keywords: []string{
				"hyperflex",
				"backup",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_capability_info",
			Category:         "hyperflex",
			ShortDescription: `A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, minUcsVersion for HyperFlex version 4.0.1a corresponds to 3.2.3 or minHxdpVersion for HyperFlex Upgrade operation is 4.0.1a etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.`,
			Description: `
A capabilityInfo is like a feature set and/or feature limit for different components of a HyperFlex Cluster. A set of constraints defines the rules, and the corresponding value either determines if the feature would work on a HyperFlex cluster with specific component set, or corresponds to a limit for a set of HyperFlex components. For example, "minUcsVersion" for HyperFlex version "4.0.1a" corresponds to "3.2.3" or "minHxdpVersion" for HyperFlex Upgrade operation is "4.0.1a" etc. This data can be captured as a capability and at run-time, decision can be made to proceed with the intended operation or not, or proceed with the intended operation with a value catered to specific feature sets.
`,
			Keywords: []string{
				"hyperflex",
				"capability",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cisco_hypervisor_manager",
			Category:         "hyperflex",
			ShortDescription: `A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.`,
			Description: `
A hypervisor manager to manage Cisco HyperFlex compute clusters and is available per user account.
`,
			Keywords: []string{
				"hyperflex",
				"cisco",
				"hypervisor",
				"manager",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster",
			Category:         "hyperflex",
			ShortDescription: `A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes of the cluster.`,
			Description: `
A HyperFlex cluster. Contains inventory information concerning the health, software versions, storage, and nodes
of the cluster.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NA",
					Description: `The deployment type of the HyperFlex cluster is not available.`,
				},
				resource.Attribute{
					Name:        "Datacenter",
					Description: `The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes on the same site.`,
				},
				resource.Attribute{
					Name:        "Stretched Cluster",
					Description: `The deployment type of a HyperFlex cluster consisting of UCS Fabric Interconnect-attached nodes across different sites.`,
				},
				resource.Attribute{
					Name:        "Edge",
					Description: `The deployment type of a HyperFlex cluster consisting of 2 or more standalone nodes.`,
				},
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_backup_policy",
			Category:         "hyperflex",
			ShortDescription: `Specifies cluster backup configuration for a HyperFlex Cluster.`,
			Description: `
Specifies cluster backup configuration for a HyperFlex Cluster.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"backup",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_backup_policy_deployment",
			Category:         "hyperflex",
			ShortDescription: `Record of HyperFlex Cluster backup policy deployment.`,
			Description: `
Record of HyperFlex Cluster backup policy deployment.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"backup",
				"policy",
				"deployment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_health_check_execution_snapshot",
			Category:         "hyperflex",
			ShortDescription: `Health check execution snapshot of the HyperFlex cluster.`,
			Description: `
Health check execution snapshot of the HyperFlex cluster.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"health",
				"check",
				"execution",
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_network_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying VLANs for management, VM migration, and VM traffic.`,
			Description: `
A policy specifying VLANs for management, VM migration, and VM traffic.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"network",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "default",
					Description: `Current default value set on the hardware platform.`,
				},
				resource.Attribute{
					Name:        "1G",
					Description: `A link speed of 1 gigabit per second.`,
				},
				resource.Attribute{
					Name:        "10G",
					Description: `A link speed of 10 gigabits per second or above.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_profile",
			Category:         "hyperflex",
			ShortDescription: `A profile specifying configuration settings for a HyperFlex cluster.`,
			Description: `
A profile specifying configuration settings for a HyperFlex cluster.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "FI",
					Description: `The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.`,
				},
				resource.Attribute{
					Name:        "EDGE",
					Description: `The host servers used in the cluster deployment are standalone severs.`,
				},
				resource.Attribute{
					Name:        "HyperFlexDp",
					Description: `The type of storage is HyperFlex Data Platform.`,
				},
				resource.Attribute{
					Name:        "ThirdParty",
					Description: `The type of storage is 3rd Party Storage (PureStorage, etc..).`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_replication_network_policy",
			Category:         "hyperflex",
			ShortDescription: `Specifies all replication network parameters for the cluster.`,
			Description: `
Specifies all replication network parameters for the cluster.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"replication",
				"network",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_replication_network_policy_deployment",
			Category:         "hyperflex",
			ShortDescription: `Record of HyperFlex Cluster replication network policy deployment.`,
			Description: `
Record of HyperFlex Cluster replication network policy deployment.
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"replication",
				"network",
				"policy",
				"deployment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_cluster_storage_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying HyperFlex cluster storage settings (optional).`,
			Description: `
A policy specifying HyperFlex cluster storage settings (optional).
`,
			Keywords: []string{
				"hyperflex",
				"cluster",
				"storage",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_config_result",
			Category:         "hyperflex",
			ShortDescription: `Profile configuration (deploy, validation) results with the overall state and detailed result messages.`,
			Description: `
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
`,
			Keywords: []string{
				"hyperflex",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_config_result_entry",
			Category:         "hyperflex",
			ShortDescription: `An entry that describes the result of a cluster validation or deployment operation.`,
			Description: `
An entry that describes the result of a cluster validation or deployment operation.
`,
			Keywords: []string{
				"hyperflex",
				"config",
				"result",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_device_package_download_state",
			Category:         "hyperflex",
			ShortDescription: `HyperFlex Device Package Download State.`,
			Description: `
HyperFlex Device Package Download State.
`,
			Keywords: []string{
				"hyperflex",
				"device",
				"package",
				"download",
				"state",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_ext_fc_storage_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying external storage connectivity information via Fabric attached FC storage.`,
			Description: `
A policy specifying external storage connectivity information via Fabric attached FC storage.
`,
			Keywords: []string{
				"hyperflex",
				"ext",
				"fc",
				"storage",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_ext_iscsi_storage_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying external storage connectivity information via Fabric attached FCoE storage.`,
			Description: `
A policy specifying external storage connectivity information via Fabric attached FCoE storage.
`,
			Keywords: []string{
				"hyperflex",
				"ext",
				"iscsi",
				"storage",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_feature_limit_external",
			Category:         "hyperflex",
			ShortDescription: `The HyperFlex feature limits that are available to end users.`,
			Description: `
The HyperFlex feature limits that are available to end users.
`,
			Keywords: []string{
				"hyperflex",
				"feature",
				"limit",
				"external",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_feature_limit_internal",
			Category:         "hyperflex",
			ShortDescription: `The HyperFlex installer feature limits for internal system use.`,
			Description: `
The HyperFlex installer feature limits for internal system use.
`,
			Keywords: []string{
				"hyperflex",
				"feature",
				"limit",
				"internal",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_health",
			Category:         "hyperflex",
			ShortDescription: `The data health status and ability of the HyperFlex storage cluster to tolerate failures.`,
			Description: `
The data health status and ability of the HyperFlex storage cluster to tolerate failures.
`,
			Keywords: []string{
				"hyperflex",
				"health",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NOT_AVAILABLE",
					Description: `The cluster does not require a connection to the arbitration service.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The cluster's connection state to the arbitration service cannot be determined.`,
				},
				resource.Attribute{
					Name:        "ONLINE",
					Description: `The cluster is connected to the arbitration service.`,
				},
				resource.Attribute{
					Name:        "OFFLINE",
					Description: `The cluster is disconnected from the arbitration service.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The replication compliance of the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "COMPLIANT",
					Description: `The HyperFlex cluster is compliant with the replication policy. All data on the cluster is replicated according to the configured replication factor.`,
				},
				resource.Attribute{
					Name:        "NON_COMPLIANT",
					Description: `The HyperFlex cluster is not compliant with the replication policy. Some data on the cluster is not replicated in accordance with the configured replication factor.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The operational status of the cluster cannot be determined.`,
				},
				resource.Attribute{
					Name:        "ONLINE",
					Description: `The HyperFlex cluster is online and is performing IO operations.`,
				},
				resource.Attribute{
					Name:        "OFFLINE",
					Description: `The HyperFlex cluster is offline and is not ready to perform IO operations.`,
				},
				resource.Attribute{
					Name:        "ENOSPACE",
					Description: `The HyperFlex cluster is out of available storage capacity and cannot perform write transactions.`,
				},
				resource.Attribute{
					Name:        "READONLY",
					Description: `The HyperFlex cluster is not accepting write transactions, but can still display static cluster information.`,
				},
				resource.Attribute{
					Name:        "NOT_AVAILABLE",
					Description: `The operational status of the ZK ensemble is not provided by the HyperFlex cluster.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The operational status of the ZK ensemble cannot be determined.`,
				},
				resource.Attribute{
					Name:        "ONLINE",
					Description: `The ZK ensemble is online and operational.`,
				},
				resource.Attribute{
					Name:        "OFFLINE",
					Description: `The ZK ensemble is offline and not operational.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_health_check_definition",
			Category:         "hyperflex",
			ShortDescription: `HyperFlex health check definition metadata.`,
			Description: `
HyperFlex health check definition metadata.
`,
			Keywords: []string{
				"hyperflex",
				"health",
				"check",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ON_DEMAND",
					Description: `Execute the health check on-demand.`,
				},
				resource.Attribute{
					Name:        "SCHEDULED",
					Description: `Execute the health check on a scheduled interval.`,
				},
				resource.Attribute{
					Name:        "EXECUTE_ON_LEADER_NODE",
					Description: `Execute the health check script only on the HyperFlex cluster's leader node.`,
				},
				resource.Attribute{
					Name:        "EXECUTE_ON_ALL_NODES",
					Description: `Execute health check on all nodes and aggregate the results.`,
				},
				resource.Attribute{
					Name:        "EXECUTE_ON_ALL_NODES_AND_AGGREGATE",
					Description: `Execute the health check on all Nodes and perform custom aggregation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_health_check_execution",
			Category:         "hyperflex",
			ShortDescription: `Health check execution result for a health check definition on a HyperFlex device.`,
			Description: `
Health check execution result for a health check definition on a HyperFlex device.
`,
			Keywords: []string{
				"hyperflex",
				"health",
				"check",
				"execution",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `Indicates that the health heck execution results are unknown.`,
				},
				resource.Attribute{
					Name:        "SUCCEEDED",
					Description: `Indicates that the health check execution succeeded.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Indicates that the health check execution failed.`,
				},
				resource.Attribute{
					Name:        "TIMED_OUT",
					Description: `Indicates that the health check execution timed out before completion.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `Indicates that the health check results could not be determined.`,
				},
				resource.Attribute{
					Name:        "PASS",
					Description: `Indicates that the health check passed.`,
				},
				resource.Attribute{
					Name:        "FAIL",
					Description: `Indicates that the health check failed.`,
				},
				resource.Attribute{
					Name:        "WARN",
					Description: `Indicates that the health check completed with a warning.`,
				},
				resource.Attribute{
					Name:        "NOT_APPLICABLE",
					Description: `Indicates that the health check is either unsupported, or not applicable on the Cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_health_check_execution_snapshot",
			Category:         "hyperflex",
			ShortDescription: `Last known health check execution results of a health check Definition.`,
			Description: `
Last known health check execution results of a health check Definition.
`,
			Keywords: []string{
				"hyperflex",
				"health",
				"check",
				"execution",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `Indicates that the health heck execution results are unknown.`,
				},
				resource.Attribute{
					Name:        "SUCCEEDED",
					Description: `Indicates that the health check execution succeeded.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Indicates that the health check execution failed.`,
				},
				resource.Attribute{
					Name:        "TIMED_OUT",
					Description: `Indicates that the health check execution timed out before completion.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `Indicates that the health check results could not be determined.`,
				},
				resource.Attribute{
					Name:        "PASS",
					Description: `Indicates that the health check passed.`,
				},
				resource.Attribute{
					Name:        "FAIL",
					Description: `Indicates that the health check failed.`,
				},
				resource.Attribute{
					Name:        "WARN",
					Description: `Indicates that the health check completed with a warning.`,
				},
				resource.Attribute{
					Name:        "NOT_APPLICABLE",
					Description: `Indicates that the health check is either unsupported, or not applicable on the Cluster.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_health_check_package_checksum",
			Category:         "hyperflex",
			ShortDescription: `HyperFlex health check Debian Package SHA512 checksum.`,
			Description: `
HyperFlex health check Debian Package SHA512 checksum.
`,
			Keywords: []string{
				"hyperflex",
				"health",
				"check",
				"package",
				"checksum",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxap_cluster",
			Category:         "hyperflex",
			ShortDescription: `A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.`,
			Description: `
A HyperFlex Application Platform compute cluster. Contains inventory information concerning the health, version and ip address of the cluster. The cluster has a name assigned by user in Intersight.
`,
			Keywords: []string{
				"hyperflex",
				"hxap",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxap_datacenter",
			Category:         "hyperflex",
			ShortDescription: `A datacenter object in HyperFlex Application Platform. It is a pre-defined object created internally by the system which acts as a container (logically) for all other objects (Host, VirtualMachine, Volume etc).`,
			Description: `
A datacenter object in HyperFlex Application Platform. It is a pre-defined object created internally by the system which acts as a container (logically) for all other objects (Host, VirtualMachine, Volume etc).
`,
			Keywords: []string{
				"hyperflex",
				"hxap",
				"datacenter",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxap_event",
			Category:         "hyperflex",
			ShortDescription: `Events associated with HyperFlex Application Platform compute cluster.`,
			Description: `
Events associated with HyperFlex Application Platform compute cluster.
`,
			Keywords: []string{
				"hyperflex",
				"hxap",
				"event",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Value is Unknown if the target type is unidentified.`,
				},
				resource.Attribute{
					Name:        "Cluster",
					Description: `Cluster refers to HyperFlex AP Cluster.`,
				},
				resource.Attribute{
					Name:        "Host",
					Description: `Host refers to server node which is part of HyperFlex AP Cluster.`,
				},
				resource.Attribute{
					Name:        "VM",
					Description: `VM refers to Virtual machine available on a HyperFlex AP Node.`,
				},
				resource.Attribute{
					Name:        "Disk",
					Description: `Disk refers to Virtual Disk available on a HyperFlex AP Cluster.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The Enum value None represents that there is no severity.`,
				},
				resource.Attribute{
					Name:        "Info",
					Description: `The Enum value Info represents the Informational level of severity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `The Enum value Critical represents the Critical level of severity.`,
				},
				resource.Attribute{
					Name:        "Warning",
					Description: `The Enum value Warning represents the Warning level of severity.`,
				},
				resource.Attribute{
					Name:        "Cleared",
					Description: `The Enum value Cleared represents that the alarm severity has been cleared.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxap_host",
			Category:         "hyperflex",
			ShortDescription: `A HyperFlex Application Platform compute host entity that is part of HyperFlex compute cluster and probably runs VMs.`,
			Description: `
A HyperFlex Application Platform compute host entity that is part of HyperFlex compute cluster and probably runs VMs.
`,
			Keywords: []string{
				"hyperflex",
				"hxap",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `The entity's power state is unknown.`,
				},
				resource.Attribute{
					Name:        "PoweredOn",
					Description: `The entity is powered on.`,
				},
				resource.Attribute{
					Name:        "PoweredOff",
					Description: `The entity is powered down.`,
				},
				resource.Attribute{
					Name:        "StandBy",
					Description: `The entity is in standby mode.`,
				},
				resource.Attribute{
					Name:        "Paused",
					Description: `The entity is in pause state.`,
				},
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxap_virtual_disk",
			Category:         "hyperflex",
			ShortDescription: `The Virtual disk created on HyperFlex Application Platform compute cluster.`,
			Description: `
The Virtual disk created on HyperFlex Application Platform compute cluster.
`,
			Keywords: []string{
				"hyperflex",
				"hxap",
				"virtual",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ReadWriteOnce",
					Description: `Read write permisisons to a Virtual disk by a single virtual machine.`,
				},
				resource.Attribute{
					Name:        "ReadWriteMany",
					Description: `Read write permisisons to a Virtual disk by multiple virtual machines.`,
				},
				resource.Attribute{
					Name:        "ReadOnlyMany",
					Description: `Read only permisisons to a Virtual disk by multiple virtual machines.`,
				},
				resource.Attribute{
					Name:        "Block",
					Description: `It is a Block virtual disk.`,
				},
				resource.Attribute{
					Name:        "Filesystem",
					Description: `It is a File system virtual disk.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxap_virtual_machine",
			Category:         "hyperflex",
			ShortDescription: `The Virtual machine that runs on a Hyperflex Application platform compute host.`,
			Description: `
The Virtual machine that runs on a Hyperflex Application platform compute host.
`,
			Keywords: []string{
				"hyperflex",
				"hxap",
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The entity's power state is unknown.`,
				},
				resource.Attribute{
					Name:        "PoweredOn",
					Description: `The entity is powered on.`,
				},
				resource.Attribute{
					Name:        "PoweredOff",
					Description: `The entity is powered down.`,
				},
				resource.Attribute{
					Name:        "StandBy",
					Description: `The entity is in standby mode.`,
				},
				resource.Attribute{
					Name:        "Paused",
					Description: `The entity is in pause state.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Virtual machine state is not available.`,
				},
				resource.Attribute{
					Name:        "Running",
					Description: `Virtual machine is running normally.`,
				},
				resource.Attribute{
					Name:        "Stopped",
					Description: `Virtual machine has been stopped.`,
				},
				resource.Attribute{
					Name:        "WaitForLaunch",
					Description: `Virtual machine is wating to be launched.`,
				},
				resource.Attribute{
					Name:        "Paused",
					Description: `Virtual machine is currently paused.`,
				},
				resource.Attribute{
					Name:        "ImportInProgress",
					Description: `Virtual machine image is being imported into the platform.`,
				},
				resource.Attribute{
					Name:        "ImportFailed",
					Description: `Virtual machine image import operation failed.`,
				},
				resource.Attribute{
					Name:        "DiskCloneInProgress",
					Description: `Disk clone operation for the virtual machine is in progress.`,
				},
				resource.Attribute{
					Name:        "DiskCloneFailed",
					Description: `Disk clone operation for the virtual machine failed.`,
				},
				resource.Attribute{
					Name:        "Processing",
					Description: `Virtual machine is being created.`,
				},
				resource.Attribute{
					Name:        "UnSchedulable",
					Description: `Virtual machine cannot be scheduled to run, either due to insufficient resources or failure to match affinity specifications.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Some virtual machine operation has failed. More information is available as part of the results of the operation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_hxdp_version",
			Category:         "hyperflex",
			ShortDescription: `A HyperFlex Data Platform version.`,
			Description: `
A HyperFlex Data Platform version.
`,
			Keywords: []string{
				"hyperflex",
				"hxdp",
				"version",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_local_credential_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying credentials for HyperFlex cluster such as controller VM password, hypervisor username, and password.`,
			Description: `
A policy specifying credentials for HyperFlex cluster such as controller VM password, hypervisor username, and password.
`,
			Keywords: []string{
				"hyperflex",
				"local",
				"credential",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_node",
			Category:         "hyperflex",
			ShortDescription: `A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.`,
			Description: `
A host participating in the cluster. The host consists of a hypervisor installed on a node that manages virtual machines.
`,
			Keywords: []string{
				"hyperflex",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The role of the HyperFlex cluster node is not known.`,
				},
				resource.Attribute{
					Name:        "STORAGE",
					Description: `The HyperFlex cluster node provides both storage and compute resources for the cluster.`,
				},
				resource.Attribute{
					Name:        "COMPUTE",
					Description: `The HyperFlex cluster node provides compute resources for the cluster.`,
				},
				resource.Attribute{
					Name:        "UNKNOWN",
					Description: `The host status cannot be determined.`,
				},
				resource.Attribute{
					Name:        "ONLINE",
					Description: `The host is online and operational.`,
				},
				resource.Attribute{
					Name:        "OFFLINE",
					Description: `The host is offline and is currently not participating in the HyperFlex cluster.`,
				},
				resource.Attribute{
					Name:        "INMAINTENANCE",
					Description: `The host is not participating in the HyperFlex cluster because of a maintenance operation, such as firmware or data platform upgrade.`,
				},
				resource.Attribute{
					Name:        "DEGRADED",
					Description: `The host is degraded and may not be performing in its full operational capacity.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_node_config_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.`,
			Description: `
A policy specifying node details such as management and storage data IP ranges. For HyperFlex Edge, storage data IP range is pre-defined.
`,
			Keywords: []string{
				"hyperflex",
				"node",
				"config",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_node_profile",
			Category:         "hyperflex",
			ShortDescription: `A configuration profile per node in the HyperFlex cluster. It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.`,
			Description: `
A configuration profile per node in the HyperFlex cluster.
It defines node settings such as IP address configuration for hypervisor management network, storage data network, HyperFlex management network, and the assigned physical server.
`,
			Keywords: []string{
				"hyperflex",
				"node",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_proxy_setting_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.`,
			Description: `
A policy specifying the HTTP proxy settings to be used by the HyperFlex installation process and HyperFlex storage controller VMs. This policy is required when the internet access of your servers including CIMC and HyperFlex storage controller VMs is secured by a HTTP proxy.
`,
			Keywords: []string{
				"hyperflex",
				"proxy",
				"setting",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_server_firmware_version",
			Category:         "hyperflex",
			ShortDescription: `The server firmware version represents the UCS server firmware details.`,
			Description: `
The server firmware version represents the UCS server firmware details.
`,
			Keywords: []string{
				"hyperflex",
				"server",
				"firmware",
				"version",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_server_firmware_version_entry",
			Category:         "hyperflex",
			ShortDescription: `An entry specifying supported server firmware version in regex format.`,
			Description: `
An entry specifying supported server firmware version in regex format.
`,
			Keywords: []string{
				"hyperflex",
				"server",
				"firmware",
				"version",
				"entry",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "M5",
					Description: `M5 generation of UCS server.`,
				},
				resource.Attribute{
					Name:        "M4",
					Description: `M4 generation of UCS server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_server_model",
			Category:         "hyperflex",
			ShortDescription: `A supported server model.`,
			Description: `
A supported server model.
`,
			Keywords: []string{
				"hyperflex",
				"server",
				"model",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_software_version_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.`,
			Description: `
A policy capturing software versions for different HyperFlex Cluster compatible components ( like HyperFlex Data Platform, Hypervisor, etc. ), that the user wishes to apply on the HyperFlex cluster.
`,
			Keywords: []string{
				"hyperflex",
				"software",
				"version",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_sys_config_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.`,
			Description: `
A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.
`,
			Keywords: []string{
				"hyperflex",
				"sys",
				"config",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Pacific/Kiritimati",
					Description: ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_ucsm_config_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.`,
			Description: `
A policy specifying UCS Manager settings such as service profile org, KVM IP addresses, and MAC prefix for server configuration in Fabric Interconnect environment.
`,
			Keywords: []string{
				"hyperflex",
				"ucsm",
				"config",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_vcenter_config_policy",
			Category:         "hyperflex",
			ShortDescription: `A policy specifying vCenter configuration.`,
			Description: `
A policy specifying vCenter configuration.
`,
			Keywords: []string{
				"hyperflex",
				"vcenter",
				"config",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_vm_backup_info",
			Category:         "hyperflex",
			ShortDescription: `Virtual Machine backup information.`,
			Description: `
Virtual Machine backup information.
`,
			Keywords: []string{
				"hyperflex",
				"vm",
				"backup",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "InitializingProtection",
					Description: `Protection has started, but not completed.`,
				},
				resource.Attribute{
					Name:        "Protected",
					Description: `Protection has completed successfully.`,
				},
				resource.Attribute{
					Name:        "ExceedsInterval",
					Description: `Protection has not completed successfully in over two times the backup interval.`,
				},
				resource.Attribute{
					Name:        "PREPARE_FAILOVER_STARTED",
					Description: `The protection status is prepare failover started.`,
				},
				resource.Attribute{
					Name:        "PREPARE_FAILOVER_FAILED",
					Description: `The protection status is prepare failover failed.`,
				},
				resource.Attribute{
					Name:        "PREPARE_FAILOVER_COMPLETED",
					Description: `The protection status is prepaire failover completed.`,
				},
				resource.Attribute{
					Name:        "FAILOVER_STARTED",
					Description: `The protection status is failover started.`,
				},
				resource.Attribute{
					Name:        "FAILOVER_FAILED",
					Description: `The protection status is failover failed.`,
				},
				resource.Attribute{
					Name:        "FAILOVER_COMPLETED",
					Description: `The protection status is failover completed.`,
				},
				resource.Attribute{
					Name:        "PREPARE_REVERSEPROTECT_STARTED",
					Description: `The protection status is prepare reverse protect started.`,
				},
				resource.Attribute{
					Name:        "PREPARE_REVERSEPROTECT_FAILED",
					Description: `The protection status is prepare reverse protect failed.`,
				},
				resource.Attribute{
					Name:        "PREPARE_REVERSEPROTECT_COMPLETED",
					Description: `The protection status is prepair reverse protect completed.`,
				},
				resource.Attribute{
					Name:        "REVERSEPROTECT_STARTED",
					Description: `The protection status is reverse protect started.`,
				},
				resource.Attribute{
					Name:        "REVERSEPROTECT_FAILED",
					Description: `The protection status is reverse protect failed.`,
				},
				resource.Attribute{
					Name:        "ACTIVE",
					Description: `The protection status is active.`,
				},
				resource.Attribute{
					Name:        "CREATION_IN_PROGRESS",
					Description: `The protection status is failover in progress.`,
				},
				resource.Attribute{
					Name:        "CREATION_FAILED",
					Description: `The protection status is creation failed.`,
				},
				resource.Attribute{
					Name:        "LOCAL_RESTORE_STARTED",
					Description: `The protection status is local restore started.`,
				},
				resource.Attribute{
					Name:        "LOCAL_RESTORE_FAILED",
					Description: `The protection status is local restore failed.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_vm_restore_operation",
			Category:         "hyperflex",
			ShortDescription: `Invoke Virtual Machine restore operation.`,
			Description: `
Invoke Virtual Machine restore operation.
`,
			Keywords: []string{
				"hyperflex",
				"vm",
				"restore",
				"operation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_hyperflex_vm_snapshot_info",
			Category:         "hyperflex",
			ShortDescription: `Virtual Machine Snapshot information like replication status, snapshot point and status.`,
			Description: `
Virtual Machine Snapshot information like replication status, snapshot point and status.
`,
			Keywords: []string{
				"hyperflex",
				"vm",
				"snapshot",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NONE",
					Description: `The snapshot quiesce mode is none.`,
				},
				resource.Attribute{
					Name:        "CRASH",
					Description: `The snapshot quiesce mode is crash.`,
				},
				resource.Attribute{
					Name:        "VMTOOLS",
					Description: `The snapshot quiesce mode is VMTOOLS.`,
				},
				resource.Attribute{
					Name:        "APP_CONSISTENT",
					Description: `The snapshot quiesce mode is app consistent.`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `Snapshot replication state is none.`,
				},
				resource.Attribute{
					Name:        "SUCCESS",
					Description: `Snapshot completed successfully.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `Snapshot failed replication status code.`,
				},
				resource.Attribute{
					Name:        "PAUSED",
					Description: `Snapshot replication paused status code.`,
				},
				resource.Attribute{
					Name:        "IN_USE",
					Description: `Snapshot replica in use status code.`,
				},
				resource.Attribute{
					Name:        "STARTING",
					Description: `Snapshot replication starting.`,
				},
				resource.Attribute{
					Name:        "REPLICATING",
					Description: `Snapshot replication in progress.`,
				},
				resource.Attribute{
					Name:        "SUCCESS",
					Description: `This snapshot status code is success.`,
				},
				resource.Attribute{
					Name:        "FAILED",
					Description: `This snapshot status code is failed.`,
				},
				resource.Attribute{
					Name:        "IN_PROGRESS",
					Description: `This snapshot status code is in progress.`,
				},
				resource.Attribute{
					Name:        "DELETING",
					Description: `This snapshot status code is deleting.`,
				},
				resource.Attribute{
					Name:        "DELETED",
					Description: `This snapshot status code is deleted.`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `This snapshot status code is none.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_connector_pack",
			Category:         "iaas",
			ShortDescription: `Describes about all the connector pack versions running currently in UCSD.`,
			Description: `
Describes about all the connector pack versions running currently in UCSD.
`,
			Keywords: []string{
				"iaas",
				"connector",
				"pack",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_device_status",
			Category:         "iaas",
			ShortDescription: `List of infra accounts managed by UCSD.`,
			Description: `
List of infra accounts managed by UCSD.
`,
			Keywords: []string{
				"iaas",
				"device",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `If UCS Director managed account claim status information is unknown.`,
				},
				resource.Attribute{
					Name:        "Yes",
					Description: `If UCS Director managed account is claimed in Intersight.`,
				},
				resource.Attribute{
					Name:        "No",
					Description: `If UCS Director managed account is not claimed in Intersight.`,
				},
				resource.Attribute{
					Name:        "Not Applicable",
					Description: `If UCS Director managed account is not capable of providing claim status information.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_diagnostic_messages",
			Category:         "iaas",
			ShortDescription: `Gets diagnostics messages from UCSD.`,
			Description: `
Gets diagnostics messages from UCSD.
`,
			Keywords: []string{
				"iaas",
				"diagnostic",
				"messages",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_license_info",
			Category:         "iaas",
			ShortDescription: `Describes about license info currently available in UCSD.`,
			Description: `
Describes about license info currently available in UCSD.
`,
			Keywords: []string{
				"iaas",
				"license",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_most_run_tasks",
			Category:         "iaas",
			ShortDescription: `Describes most run workflow tasks within UCSD.`,
			Description: `
Describes most run workflow tasks within UCSD.
`,
			Keywords: []string{
				"iaas",
				"most",
				"run",
				"tasks",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_service_request",
			Category:         "iaas",
			ShortDescription: `Gets last six months Service Requests from UCSD.`,
			Description: `
Gets last six months Service Requests from UCSD.
`,
			Keywords: []string{
				"iaas",
				"service",
				"request",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_ucsd_info",
			Category:         "iaas",
			ShortDescription: `UCS Director accounts managed by Intersight.`,
			Description: `
UCS Director accounts managed by Intersight.
`,
			Keywords: []string{
				"iaas",
				"ucsd",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_ucsd_managed_infra",
			Category:         "iaas",
			ShortDescription: `Describes about UCSD Managed infrastructure statistics.`,
			Description: `
Describes about UCSD Managed infrastructure statistics.
`,
			Keywords: []string{
				"iaas",
				"ucsd",
				"managed",
				"infra",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iaas_ucsd_messages",
			Category:         "iaas",
			ShortDescription: `Gets ucsd messages from UCSD.`,
			Description: `
Gets ucsd messages from UCSD.
`,
			Keywords: []string{
				"iaas",
				"ucsd",
				"messages",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_account",
			Category:         "iam",
			ShortDescription: `The Intersight Account used to access Intersight.`,
			Description: `
The Intersight Account used to access Intersight.
`,
			Keywords: []string{
				"iam",
				"account",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_account_experience",
			Category:         "iam",
			ShortDescription: `The beta features enabled for the specified account.`,
			Description: `
The beta features enabled for the specified account.
`,
			Keywords: []string{
				"iam",
				"account",
				"experience",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_api_key",
			Category:         "iam",
			ShortDescription: `An API key is used to authenticate and authorize API requests sent by a client using the HTTP signature scheme. API keys can be used by unattended, daemon clients that need to send requests to Intersight programmatically. API keys are based on public key cryptography. To create an API key, the user must specify: 1. The purpose (description) of the API key, 2. The cryptographic hash algorithm, which is used to compute the digest of the body of HTTP requests, 3. The cryptographic parameters to generate a private/public key pair, e.g. RSA, ECDSA, EDDSA, key modulus, and 4. The signing algorithm, e.g. RSA PKCS v1.5, RSA PSS, ECDSA, EDDSA. The generated private key and public key are encoded in PEM format. The client owns the private key and is responsible for maintaining the confidentiality of the private key. The server holds the public key. The client must have a cryptographic provider compatible with the cryptographic parameters specified in the API key. For example, if you use the powershell SDK to write the client, make sure the appropriate cryptographic providers are installed on the local system. If you create an RSA key pair with modulus set to 2048, the client must support 2048-bit private keys. A maximum of 3 API keys per user is allowed. API keys are used to sign HTTP requests as follows: 1. A cryptographic digest of the body of the HTTP request is calculated using one of the supported cryptographic hash algorithms. 2. The value of the digest is base-64 encoded in the ` + "`" + `Digest` + "`" + ` HTTP header. 3. A signature is calculated as specified in the HTTP signature scheme, and the signature is added to the ` + "`" + `Authorization` + "`" + ` HTTP request header. All published Intersight SDKs support API keys.`,
			Description: `
An API key is used to authenticate and authorize API requests sent by a client using the HTTP signature scheme. API keys can be used by unattended, daemon clients that need to send requests to Intersight programmatically. API keys are based on public key cryptography.
To create an API key, the user must specify: 1. The purpose (description) of the API key, 2. The cryptographic hash algorithm, which is used to compute the digest of the body of HTTP requests, 3. The cryptographic parameters to generate a private/public key pair, e.g. RSA, ECDSA, EDDSA, key modulus, and 4. The signing algorithm, e.g. RSA PKCS v1.5, RSA PSS, ECDSA, EDDSA. The generated private key and public key are encoded in PEM format.
The client owns the private key and is responsible for maintaining the confidentiality of the private key. The server holds the public key.
The client must have a cryptographic provider compatible with the cryptographic parameters specified in the API key. For example, if you use the powershell SDK to write the client, make sure the appropriate cryptographic providers are installed on the local system. If you create an RSA key pair with modulus set to 2048, the client must support 2048-bit private keys. A maximum of 3 API keys per user is allowed.
API keys are used to sign HTTP requests as follows: 1. A cryptographic digest of the body of the HTTP request is calculated using one of the supported cryptographic hash algorithms. 2. The value of the digest is base-64 encoded in the ` + "`" + `Digest` + "`" + ` HTTP header. 3. A signature is calculated as specified in the HTTP signature scheme, and the signature is added to the ` + "`" + `Authorization` + "`" + ` HTTP request header.
All published Intersight SDKs support API keys.
`,
			Keywords: []string{
				"iam",
				"api",
				"key",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "SHA256",
					Description: `The SHA-256 cryptographic hash, as defined by NIST in FIPS 180-4.`,
				},
				resource.Attribute{
					Name:        "SHA384",
					Description: `The SHA-384 cryptographic hash, as defined by NIST in FIPS 180-4.`,
				},
				resource.Attribute{
					Name:        "SHA512",
					Description: `The SHA-512 cryptographic hash, as defined by NIST in FIPS 180-4.`,
				},
				resource.Attribute{
					Name:        "SHA512_224",
					Description: `The SHA-512/224 cryptographic hash, as defined by NIST in FIPS 180-4.`,
				},
				resource.Attribute{
					Name:        "SHA512_256",
					Description: `The SHA-512/256 cryptographic hash, as defined by NIST in FIPS 180-4.`,
				},
				resource.Attribute{
					Name:        "RSASSA-PKCS1-v1_5",
					Description: `RSASSA-PKCS1-v1_5 is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).RSASSA-PKCS1-v1_5 is included only for compatibility with existing applications.`,
				},
				resource.Attribute{
					Name:        "RSASSA-PSS",
					Description: `RSASSA-PSS is a RSA signature scheme specified in [RFC 8017](https://tools.ietf.org/html/rfc8017).It combines the RSASP1 and RSAVP1 primitives with the EMSA-PSS encoding method.In the interest of increased robustness, RSASSA-PSS is required in new applications.`,
				},
				resource.Attribute{
					Name:        "Ed25519",
					Description: `The Ed25519 signature algorithm, as specified in [RFC 8032](https://tools.ietf.org/html/rfc8032).Ed25519 is a public-key signature system with several attractive features, includingfast single-signature verification, very fast signing, fast key generation and high security level.`,
				},
				resource.Attribute{
					Name:        "Ecdsa",
					Description: `The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is encoded as a ASN.1 DER SEQUENCE with two INTEGERs (r and s), as defined in RFC3279.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.`,
				},
				resource.Attribute{
					Name:        "EcdsaP1363Format",
					Description: `The Elliptic Curve Digital Signature Standard (ECDSA), as defined by NIST in FIPS 186-4 and ANSI X9.62.The signature is the raw concatenation of r and s, as defined in the ISO/IEC 7816-8 IEEE P.1363 standard.In that format, r and s are represented as unsigned, big endian numbers.Extra padding bytes (of value 0x00) is applied so that both r and s encodings have the same size.When using ECDSA signatures, configure the client to use the same signature encoding as specified on the server side.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_app_registration",
			Category:         "iam",
			ShortDescription: ``,
			Description: `
AppRegistration encapsulates the meta-data values of a registered OAuth2 client application, as described in
https://tools.ietf.org/html/rfc7591#section-2.
Registered client applications have a set of metadata values associated with their client identifier
at the Intersight authorization server, including the list of valid redirection URIs or a display name.
The meta-data is used to specify how a client application can retrieve a OAuth2 Access Token and subsequently
invoke Intersight API on behalf of this AppRegistration.
To register an OAuth2 application, the following information must be provided.
1) Application name
2) An icon for the application
3) URL to the application's home page
4) A short description of the application
5) A list of redirect URLs
When an AppRegistration is created, a unique OAuth2 clientId is generated and returned in the HTTP response.
`,
			Keywords: []string{
				"iam",
				"app",
				"registration",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "public",
					Description: `Clients incapable of maintaining the confidentiality of their credentials.This includes clients executing on the device used by the resource owner,such as mobile applications, installed native application or a webbrowser-based application.`,
				},
				resource.Attribute{
					Name:        "confidential",
					Description: `Clients capable of maintaining the confidentiality of their credentials.For example, this could be a client implemented on a secure server withrestricted access to the client credentials.To maintain the confidentiality of the OAuth2 credentials, two use cases areconsidered.1) The application is running as a service within Intersight. The application automatically obtains the OAuth2 credentials when the application starts and the credentials are not exposed to the end-user. Because end-users (even account administrators) do not have access the OAuth2 credentials, they cannot take the credentials with them when they leave their organization.2) The application is under the control of a \ trusted\ end-user. For example, the end-user may create a native application running outside Intersight. The application uses OAuth2 credentials to interact with the Intersight API. In that case, the Intersight account administrator may generate OAuth2 credentials with a registered application using \ client_credentials\ grant type. In that case, the end-user is responsible for maintaining the confidentiality of the OAuth2 credentials. If the end-user leaves the organization, you should revoke the credentials and issue new Oauth2 credentials.Here is a possible workflow for handling OAuth2 tokens.1) User Alice (Intersight Account Administrator) logins to Intersight and deploys an Intersight application that requires an OAuth2 token.2) Intersight automatically deploys the application. The application is assigned a OAuth2 token, possibly linked to Alice. The application must NOT expose the OAuth2 secret to Alice, otherwise Alice would be able to use the token after she leaves the company.3) The application can make API calls to Intersight using its assigned OAuth2 token. For example, the application could make weekly scheduled API calls to Intersight.4) Separately, Alice may also get OAuth2 tokens that she can use to make API calls from the Intersight SDK through the northbound API. In that case, Alice will get the associated OAuth2 secrets, but not the one assigned in step #2.5) Alice leaves the organization. The OAuth2 tokens assigned in step #2 must retain their validity even after Alice has left the organization. Because the OAuth2 secrets were never shared with Alice, there is no risk Alice can reuse the OAuth2 secrets. On the other hand, the OAuth2 tokens assigned in step #4 must be invalidated because Alice had the OAuth2 tokens in her possession.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_banner_message",
			Category:         "iam",
			ShortDescription: `Configuration for the banner message, including title, contents, and display toggle.`,
			Description: `
Configuration for the banner message, including title, contents, and display toggle.
`,
			Keywords: []string{
				"iam",
				"banner",
				"message",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_certificate",
			Category:         "iam",
			ShortDescription: `Holds a certificate, signed by a CAcert.`,
			Description: `
Holds a certificate, signed by a CAcert.
`,
			Keywords: []string{
				"iam",
				"certificate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "PendingValidation",
					Description: `The certificate has not been validated.`,
				},
				resource.Attribute{
					Name:        "Valid",
					Description: `The certificate is valid.`,
				},
				resource.Attribute{
					Name:        "Invalid",
					Description: `Ther certificate is invalid.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_certificate_request",
			Category:         "iam",
			ShortDescription: `The information required to generate a certificate signing request (CSR), which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.`,
			Description: `
The information required to generate a certificate signing request (CSR),
which is a block of encoded text that is given to a Certificate Authority when applying for an SSL Certificate.
`,
			Keywords: []string{
				"iam",
				"certificate",
				"request",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_domain_group",
			Category:         "iam",
			ShortDescription: `Intersight services are mapped to three different categories of services for scaling purpose. Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with a specific number of partitions. For each cloud environment these numbers will be different.`,
			Description: `
Intersight services are mapped to three different categories of services for scaling purpose.
Three categories are defined: Partition1/Partition2/Partition3. Topics for each category are created with
a specific number of partitions. For each cloud environment these numbers will be different.
`,
			Keywords: []string{
				"iam",
				"domain",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_end_point_privilege",
			Category:         "iam",
			ShortDescription: `The privilege defined at the end point which can be assigned to a user.`,
			Description: `
The privilege defined at the end point which can be assigned to a user.
`,
			Keywords: []string{
				"iam",
				"end",
				"point",
				"privilege",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_end_point_role",
			Category:         "iam",
			ShortDescription: `The role defined in the end point which can be assigned to a user.`,
			Description: `
The role defined in the end point which can be assigned to a user.
`,
			Keywords: []string{
				"iam",
				"end",
				"point",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_end_point_user",
			Category:         "iam",
			ShortDescription: `Endpoint User or Local User.`,
			Description: `
Endpoint User or Local User.
`,
			Keywords: []string{
				"iam",
				"end",
				"point",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_end_point_user_policy",
			Category:         "iam",
			ShortDescription: `Enables creation of local users on endpoints.`,
			Description: `
Enables creation of local users on endpoints.
`,
			Keywords: []string{
				"iam",
				"end",
				"point",
				"user",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_end_point_user_role",
			Category:         "iam",
			ShortDescription: `Mapping of endpoint user to endpoint roles.`,
			Description: `
Mapping of endpoint user to endpoint roles.
`,
			Keywords: []string{
				"iam",
				"end",
				"point",
				"user",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_idp",
			Category:         "iam",
			ShortDescription: `The SAML identity provider such as Cisco, that has been used to log in to Intersight.`,
			Description: `
The SAML identity provider such as Cisco, that has been used to log in to Intersight.
`,
			Keywords: []string{
				"iam",
				"idp",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "saml",
					Description: `Use SAML as the authentication protocol for sign-on.`,
				},
				resource.Attribute{
					Name:        "oidc",
					Description: `Open ID connect to be used as an authentication protocol for sign-on.`,
				},
				resource.Attribute{
					Name:        "local",
					Description: `The local authentication method to be used for sign-on. Local type is set to default for the Intersight Appliance IdP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_idp_reference",
			Category:         "iam",
			ShortDescription: `Default Cisco IdP for authentication.`,
			Description: `
Default Cisco IdP for authentication.
`,
			Keywords: []string{
				"iam",
				"idp",
				"reference",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_ip_access_management",
			Category:         "iam",
			ShortDescription: `The access management based on IP address.`,
			Description: `
The access management based on IP address.
`,
			Keywords: []string{
				"iam",
				"ip",
				"access",
				"management",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_ip_address",
			Category:         "iam",
			ShortDescription: `Add an IP address to enable IP address based access management.`,
			Description: `
Add an IP address to enable IP address based access management.
`,
			Keywords: []string{
				"iam",
				"ip",
				"address",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_ldap_group",
			Category:         "iam",
			ShortDescription: `Mapping of LDAP Group to EndPointRoles.`,
			Description: `
Mapping of LDAP Group to EndPointRoles.
`,
			Keywords: []string{
				"iam",
				"ldap",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_ldap_policy",
			Category:         "iam",
			ShortDescription: `LDAP Policy configurations.`,
			Description: `
LDAP Policy configurations.
`,
			Keywords: []string{
				"iam",
				"ldap",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "LocalUserDb",
					Description: `Precedence is given to local user database while searching.`,
				},
				resource.Attribute{
					Name:        "LDAPUserDb",
					Description: `Precedence is given to LADP user database while searching.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_ldap_provider",
			Category:         "iam",
			ShortDescription: `LDAP Provider or LDAP Server for user authentication.`,
			Description: `
LDAP Provider or LDAP Server for user authentication.
`,
			Keywords: []string{
				"iam",
				"ldap",
				"provider",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_local_user_password_policy",
			Category:         "iam",
			ShortDescription: `Configuration for LocalUserPasswordPolicy.`,
			Description: `
Configuration for LocalUserPasswordPolicy.
`,
			Keywords: []string{
				"iam",
				"local",
				"user",
				"password",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_o_auth_token",
			Category:         "iam",
			ShortDescription: `The meta data for generating OAuth2 token of a user. It is created when user logged in via OAuth2 using Authorization Code grant and deleted upon logout, expiration timeout or manual deletion.`,
			Description: `
The meta data for generating OAuth2 token of a user.
It is created when user logged in via OAuth2 using Authorization Code grant
and deleted upon logout, expiration timeout or manual deletion.
`,
			Keywords: []string{
				"iam",
				"o",
				"auth",
				"token",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_permission",
			Category:         "iam",
			ShortDescription: `Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.`,
			Description: `
Permission provides a way to assign roles to a user or user group to perform operations on object hierarchy.
`,
			Keywords: []string{
				"iam",
				"permission",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_private_key_spec",
			Category:         "iam",
			ShortDescription: `Parameters used to generate a private key.`,
			Description: `
Parameters used to generate a private key.
`,
			Keywords: []string{
				"iam",
				"private",
				"key",
				"spec",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_privilege",
			Category:         "iam",
			ShortDescription: `Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.`,
			Description: `
Privilege represents an action which can be performed in Intersight such as creating server profile, deleting a user etc.
`,
			Keywords: []string{
				"iam",
				"privilege",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_privilege_set",
			Category:         "iam",
			ShortDescription: `A collection of privileges.`,
			Description: `
A collection of privileges.
`,
			Keywords: []string{
				"iam",
				"privilege",
				"set",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_qualifier",
			Category:         "iam",
			ShortDescription: `The qualifier defines how a user qualifies to be part of a user group.`,
			Description: `
The qualifier defines how a user qualifies to be part of a user group.
`,
			Keywords: []string{
				"iam",
				"qualifier",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_resource_limits",
			Category:         "iam",
			ShortDescription: `The resource limits used to limit resources such as User accounts.`,
			Description: `
The resource limits used to limit resources such as User accounts.
`,
			Keywords: []string{
				"iam",
				"resource",
				"limits",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_resource_permission",
			Category:         "iam",
			ShortDescription: `ResourcePermission represents the permissions defined on a resource like organization.`,
			Description: `
ResourcePermission represents the permissions defined on a resource like organization.
`,
			Keywords: []string{
				"iam",
				"resource",
				"permission",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_resource_roles",
			Category:         "iam",
			ShortDescription: `ResourceRoles provides a way to specify the roles associated with a resource like organization in a permission which can be assigned to a user or user group.`,
			Description: `
ResourceRoles provides a way to specify the roles associated with a resource like organization in a permission which can be assigned to a user or user group.
`,
			Keywords: []string{
				"iam",
				"resource",
				"roles",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_role",
			Category:         "iam",
			ShortDescription: `A role is a collection of privilege sets that are assigned to a user using a permission object.`,
			Description: `
A role is a collection of privilege sets that are assigned to a user using a permission object.
`,
			Keywords: []string{
				"iam",
				"role",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_security_holder",
			Category:         "iam",
			ShortDescription: `Holder for organization aggregated permissions and global account permissions. User configures permissions for entire account or subset of organizations and specifies associated roles with each organization. Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.`,
			Description: `
Holder for organization aggregated permissions and global account permissions.
User configures permissions for entire account or subset of organizations and specifies associated roles with each organization.
Intersight aggregates all the permissions and stores per organization aggregate permissions in iam.ResourcePermission object.
`,
			Keywords: []string{
				"iam",
				"security",
				"holder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_service_provider",
			Category:         "iam",
			ShortDescription: `SAML Service provider related information in Intersight.`,
			Description: `
SAML Service provider related information in Intersight.
`,
			Keywords: []string{
				"iam",
				"service",
				"provider",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_session",
			Category:         "iam",
			ShortDescription: `The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.`,
			Description: `
The web session of a user. After a user logs into Intersight, a session object is created. Session object is deleted upon logout, idle timeout, expiry timeout, or manual deletion.
`,
			Keywords: []string{
				"iam",
				"session",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_session_limits",
			Category:         "iam",
			ShortDescription: `The session related configuration limits.`,
			Description: `
The session related configuration limits.
`,
			Keywords: []string{
				"iam",
				"session",
				"limits",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_system",
			Category:         "iam",
			ShortDescription: `System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.`,
			Description: `
System is the top level object in the Intersight. All other objects which can be accessed globally are added to system object, like privilege sets and privileges can be shared by multiple roles and privilege sets.
`,
			Keywords: []string{
				"iam",
				"system",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_trust_point",
			Category:         "iam",
			ShortDescription: `To affirm the identity of trusted source. Allows import of third-party CA certificates in X.509 (CER) format. It can be a root CA or an trust chain that leads to a root CA.`,
			Description: `
To affirm the identity of trusted source.
Allows import of third-party CA certificates in X.509 (CER) format.
It can be a root CA or an trust chain that leads to a root CA.
`,
			Keywords: []string{
				"iam",
				"trust",
				"point",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_user",
			Category:         "iam",
			ShortDescription: `The Intersight account user.`,
			Description: `
The Intersight account user.
`,
			Keywords: []string{
				"iam",
				"user",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_user_group",
			Category:         "iam",
			ShortDescription: `User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.`,
			Description: `
User Group provides a way to assign permissions to a group of users based on the IdP attributes received after authentication.
`,
			Keywords: []string{
				"iam",
				"user",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iam_user_preference",
			Category:         "iam",
			ShortDescription: `Holder for UI preferences such as theme, columns.`,
			Description: `
Holder for UI preferences such as theme, columns.
`,
			Keywords: []string{
				"iam",
				"user",
				"preference",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_inventory_device_info",
			Category:         "inventory",
			ShortDescription: `Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.`,
			Description: `
Information pertaining to a Registered Device in starship which is pertinent to Inventory Microservice.
`,
			Keywords: []string{
				"inventory",
				"device",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_inventory_dn_mo_binding",
			Category:         "inventory",
			ShortDescription: `DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.`,
			Description: `
DnMoBinding provides a binding between a Intersight MO and a UCSM MO which has a DN.
`,
			Keywords: []string{
				"inventory",
				"dn",
				"mo",
				"binding",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_inventory_generic_inventory",
			Category:         "inventory",
			ShortDescription: `Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.`,
			Description: `
Any inventory which is represented as a key / value pair. Example - moInvKv in UCSM representing OS tools running on ESX.
`,
			Keywords: []string{
				"inventory",
				"generic",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_inventory_generic_inventory_holder",
			Category:         "inventory",
			ShortDescription: `A container class for generic inventory.`,
			Description: `
A container class for generic inventory.
`,
			Keywords: []string{
				"inventory",
				"generic",
				"holder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ipmioverlan_policy",
			Category:         "ipmioverlan",
			ShortDescription: `Intelligent Platform Management Interface Over LAN Policy.`,
			Description: `
Intelligent Platform Management Interface Over LAN Policy.
`,
			Keywords: []string{
				"ipmioverlan",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "admin",
					Description: `Privilege to perform all actions available through IPMI.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Privilege to perform some functions through IPMI but restriction on performing administrative tasks.`,
				},
				resource.Attribute{
					Name:        "read-only",
					Description: `Privilege to view information throught IPMI but restriction on making any changes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_block_lease",
			Category:         "ippool",
			ShortDescription: `BlockLease represents an IP address that is allocated from a pool to a specific entity like server profile.`,
			Description: `
BlockLease represents an IP address that is allocated from a pool to a specific entity like server profile.
`,
			Keywords: []string{
				"ippool",
				"block",
				"lease",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IPv4",
					Description: `IP V4 address type requested.`,
				},
				resource.Attribute{
					Name:        "IPv6",
					Description: `IP V6 address type requested.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_ip_lease",
			Category:         "ippool",
			ShortDescription: `IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.`,
			Description: `
IpLease represents an IP address that is allocated from a pool to a specific entity like server profile.
`,
			Keywords: []string{
				"ippool",
				"ip",
				"lease",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dynamic",
					Description: `Identifiers to be allocated by system.`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Identifiers are assigned by the user.`,
				},
				resource.Attribute{
					Name:        "IPv4",
					Description: `IP V4 address type requested.`,
				},
				resource.Attribute{
					Name:        "IPv6",
					Description: `IP V6 address type requested.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_pool",
			Category:         "ippool",
			ShortDescription: `Pool represents a collection of IPv4 and/or IPv6 addresses that can be allocated to other configuration entities like server profiles.`,
			Description: `
Pool represents a collection of IPv4 and/or IPv6 addresses that can be allocated to other configuration entities like server profiles.
`,
			Keywords: []string{
				"ippool",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sequential",
					Description: `Identifiers are assigned in a sequential order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Assignment order is decided by the system.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_pool_member",
			Category:         "ippool",
			ShortDescription: `PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.`,
			Description: `
PoolMember represents a single IPv4 and or IPv6 address that is part of a pool.
`,
			Keywords: []string{
				"ippool",
				"pool",
				"member",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IPv4",
					Description: `IP V4 address type requested.`,
				},
				resource.Attribute{
					Name:        "IPv6",
					Description: `IP V6 address type requested.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_shadow_block",
			Category:         "ippool",
			ShortDescription: `A block of Contiguous IP addresses that are part of a shadow pool.`,
			Description: `
A block of Contiguous IP addresses that are part of a shadow pool.
`,
			Keywords: []string{
				"ippool",
				"shadow",
				"block",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IPv4",
					Description: `IP V4 address type requested.`,
				},
				resource.Attribute{
					Name:        "IPv6",
					Description: `IP V6 address type requested.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_shadow_pool",
			Category:         "ippool",
			ShortDescription: `Shadow Pool is a tracking object created on behalf of an IP pool, for each VRF.`,
			Description: `
Shadow Pool is a tracking object created on behalf of an IP pool, for each VRF.
`,
			Keywords: []string{
				"ippool",
				"shadow",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sequential",
					Description: `Identifiers are assigned in a sequential order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Assignment order is decided by the system.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ippool_universe",
			Category:         "ippool",
			ShortDescription: `Universe represents a book keeping container to keep track of all IP Addresses for a given VRF.`,
			Description: `
Universe represents a book keeping container to keep track of all IP Addresses for a given VRF.
`,
			Keywords: []string{
				"ippool",
				"universe",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iqnpool_block",
			Category:         "iqnpool",
			ShortDescription: `A block of contiguous IQNs that are part of a pool.`,
			Description: `
A block of contiguous IQNs that are part of a pool.
`,
			Keywords: []string{
				"iqnpool",
				"block",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iqnpool_lease",
			Category:         "iqnpool",
			ShortDescription: `Lease represents a single IQN address that is part of the universe, allocated either from a pool or through static assignment.`,
			Description: `
Lease represents a single IQN address that is part of the universe, allocated either from a pool or through static assignment.
`,
			Keywords: []string{
				"iqnpool",
				"lease",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dynamic",
					Description: `Identifiers to be allocated by system.`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Identifiers are assigned by the user.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iqnpool_pool",
			Category:         "iqnpool",
			ShortDescription: `Pool represents a collection of iSCSI Qualified Names (IQNs) for use as initiator identifiers by iSCSI vNICs.`,
			Description: `
Pool represents a collection of iSCSI Qualified Names (IQNs) for use as initiator identifiers by iSCSI vNICs.
`,
			Keywords: []string{
				"iqnpool",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sequential",
					Description: `Identifiers are assigned in a sequential order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Assignment order is decided by the system.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iqnpool_pool_member",
			Category:         "iqnpool",
			ShortDescription: `PoolMember represents a single IQN address that is part of a pool.`,
			Description: `
PoolMember represents a single IQN address that is part of a pool.
`,
			Keywords: []string{
				"iqnpool",
				"pool",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iqnpool_universe",
			Category:         "iqnpool",
			ShortDescription: `Universe represents a book keeping container to keep track of all IDs for a given account and pool type.`,
			Description: `
Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
`,
			Keywords: []string{
				"iqnpool",
				"universe",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_iwotenant_tenant_status",
			Category:         "iwotenant",
			ShortDescription: `When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.`,
			Description: `
When an Intersight customer activates IWO license, kubernetes resources need to configured, vault policies need to be setup, etc. to run IWO services as pods. The provisioning of these resources takes a few minutes. Any feature dependent on the IWO APIs will fail till these services are healthy. TenantStatus MO provides status and health of the tenants, so user can be notified when tenant creation is in progress or has failed.
`,
			Keywords: []string{
				"iwotenant",
				"tenant",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NotStarted",
					Description: `The workflow to deploy the tenant cluster has not yet started.`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `The workflow to deploy the tenant cluster in progress. All the tasks required for thesuccessful tenant creation are running.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `The workflow to deploy the tenant cluster has completed and health checks have passed.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The workflow to deploy the tenant cluster has failed. Detailed reason for the failure isprovided from Tenant.deployStatus.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_aci_cni_apic",
			Category:         "kubernetes",
			ShortDescription: `Internally generated object of claimed APIC device known to Razor.`,
			Description: `
Internally generated object of claimed APIC device known to Razor.
`,
			Keywords: []string{
				"kubernetes",
				"aci",
				"cni",
				"apic",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_aci_cni_profile",
			Category:         "kubernetes",
			ShortDescription: `Configuration for an ACI CNI profile.`,
			Description: `
Configuration for an ACI CNI profile.
`,
			Keywords: []string{
				"kubernetes",
				"aci",
				"cni",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_aci_cni_tenant_cluster_allocation",
			Category:         "kubernetes",
			ShortDescription: `Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.`,
			Description: `
Internally generated parameter allocations for a k8s cluster using a particular ACI CNI profile.
`,
			Keywords: []string{
				"kubernetes",
				"aci",
				"cni",
				"tenant",
				"cluster",
				"allocation",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_addon",
			Category:         "kubernetes",
			ShortDescription: `An object that describes an instance of an addon definition with install/upgrade strategies and optional overrides.`,
			Description: `
An object that describes an instance of an addon definition with install/upgrade strategies and optional overrides.
`,
			Keywords: []string{
				"kubernetes",
				"addon",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "InstallOnly",
					Description: `Only install in green field. No action in case of failure or removal.`,
				},
				resource.Attribute{
					Name:        "NoAction",
					Description: `No install action performed.`,
				},
				resource.Attribute{
					Name:        "Always",
					Description: `Attempt install if chart is not already installed.`,
				},
				resource.Attribute{
					Name:        "UpgradeOnly",
					Description: `Attempt upgrade if chart or overrides options change, no action on upgrade failure.`,
				},
				resource.Attribute{
					Name:        "NoAction",
					Description: `This choice enables No upgrades to be performed.`,
				},
				resource.Attribute{
					Name:        "ReinstallOnFailure",
					Description: `Attempt upgrade first. Remove and install on upgrade failure.`,
				},
				resource.Attribute{
					Name:        "AlwaysReinstall",
					Description: `Always remove older release and reinstall.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_addon_definition",
			Category:         "kubernetes",
			ShortDescription: `An addon that can be added to any Kubernetes cluster.`,
			Description: `
An addon that can be added to any Kubernetes cluster.
`,
			Keywords: []string{
				"kubernetes",
				"addon",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "InstallOnly",
					Description: `Only install in green field. No action in case of failure or removal.`,
				},
				resource.Attribute{
					Name:        "NoAction",
					Description: `No install action performed.`,
				},
				resource.Attribute{
					Name:        "Always",
					Description: `Attempt install if chart is not already installed.`,
				},
				resource.Attribute{
					Name:        "UpgradeOnly",
					Description: `Attempt upgrade if chart or overrides options change, no action on upgrade failure.`,
				},
				resource.Attribute{
					Name:        "NoAction",
					Description: `This choice enables No upgrades to be performed.`,
				},
				resource.Attribute{
					Name:        "ReinstallOnFailure",
					Description: `Attempt upgrade first. Remove and install on upgrade failure.`,
				},
				resource.Attribute{
					Name:        "AlwaysReinstall",
					Description: `Always remove older release and reinstall.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_addon_policy",
			Category:         "kubernetes",
			ShortDescription: `A policy that defines which AddonDefinitions to use.`,
			Description: `
A policy that defines which AddonDefinitions to use.
`,
			Keywords: []string{
				"kubernetes",
				"addon",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_addon_repository",
			Category:         "kubernetes",
			ShortDescription: `Docker registry or helm repository which hosts helm charts and docker images.`,
			Description: `
Docker registry or helm repository which hosts helm charts and docker images.
`,
			Keywords: []string{
				"kubernetes",
				"addon",
				"repository",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_catalog",
			Category:         "kubernetes",
			ShortDescription: `A catalog to hold the Kubernetes related items such as versions and addons for Intersight Kubernetes Services.`,
			Description: `
A catalog to hold the Kubernetes related items such as versions and addons for Intersight Kubernetes Services.
`,
			Keywords: []string{
				"kubernetes",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_cluster",
			Category:         "kubernetes",
			ShortDescription: `Inventories a Kubernetes cluster state. A Cluster object is automatically created when a Kubernetes API server is configured for a cluster.`,
			Description: `
Inventories a Kubernetes cluster state. A Cluster object is automatically created when a Kubernetes API server is configured for a cluster.
`,
			Keywords: []string{
				"kubernetes",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Connected",
					Description: `Intersight is able to establish a connection to the target and initiate management activities.`,
				},
				resource.Attribute{
					Name:        "NotConnected",
					Description: `Intersight is unable to establish a connection to the target.`,
				},
				resource.Attribute{
					Name:        "ClaimInProgress",
					Description: `Claim of the target is in progress. A connection to the target has not been fully established.`,
				},
				resource.Attribute{
					Name:        "Unclaimed",
					Description: `The device was un-claimed from the users account by an Administrator of the device. Also indicates the failure to claim Targets of type HTTP Endpoint in Intersight.`,
				},
				resource.Attribute{
					Name:        "Claimed",
					Description: `Target of type HTTP Endpoint is successfully claimed in Intersight. Currently no validation is performed to verify the Target connectivity from Intersight at the time of creation. However invoking API from Intersight Orchestrator fails if this Target is not reachable from Intersight or if Target API credentials are incorrect.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_cluster_profile",
			Category:         "kubernetes",
			ShortDescription: `Cluster profile specifies the config profile for a Kubernetes cluster. It also depicts operations to control the life cycle of a Kubernetes cluster.`,
			Description: `
Cluster profile specifies the config profile for a Kubernetes cluster. It also depicts operations to control the life cycle of a Kubernetes cluster.
`,
			Keywords: []string{
				"kubernetes",
				"cluster",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Provided",
					Description: `Cluster management entities and endpoints are provided by the infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "Managed",
					Description: `Cluster management entities and endpoints are provisioned and managed by Intersight kubernetes service.`,
				},
				resource.Attribute{
					Name:        "Configuring",
					Description: `The cluster is being configured.`,
				},
				resource.Attribute{
					Name:        "Deploying",
					Description: `The cluster is being deployed.`,
				},
				resource.Attribute{
					Name:        "Undeploying",
					Description: `The cluster is being undeployed.`,
				},
				resource.Attribute{
					Name:        "DeployFailed",
					Description: `The cluster deployment failed.`,
				},
				resource.Attribute{
					Name:        "Upgrading",
					Description: `The cluster is being upgraded.`,
				},
				resource.Attribute{
					Name:        "Deleting",
					Description: `The cluster is being deleted.`,
				},
				resource.Attribute{
					Name:        "DeleteFailed",
					Description: `The cluster delete failed.`,
				},
				resource.Attribute{
					Name:        "Ready",
					Description: `The cluster is ready for use.`,
				},
				resource.Attribute{
					Name:        "Active",
					Description: `The cluster is being active.`,
				},
				resource.Attribute{
					Name:        "Shutdown",
					Description: `All the nodes in the cluster are powered off.`,
				},
				resource.Attribute{
					Name:        "Terminated",
					Description: `The cluster is terminated.`,
				},
				resource.Attribute{
					Name:        "Deployed",
					Description: `The cluster is deployed. The cluster may not yet be ready for use.`,
				},
				resource.Attribute{
					Name:        "UndeployFailed",
					Description: `The cluster undeploy action failed.`,
				},
				resource.Attribute{
					Name:        "NotReady",
					Description: `The cluster is created and some nodes are not ready.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_container_runtime_policy",
			Category:         "kubernetes",
			ShortDescription: `A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.`,
			Description: `
A policy specifying container runtime configuration, such as docker proxy, no proxy and bridge network IP.
`,
			Keywords: []string{
				"kubernetes",
				"container",
				"runtime",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_daemon_set",
			Category:         "kubernetes",
			ShortDescription: `A DaemonSet represents a Kubernetes DaemonSet. In Kubernetes, a DaemonSet ensures that all (or some) Nodes run a copy of a Pod. As nodes are added to the cluster, Pods are added to them. As nodes are removed from the cluster, those Pods are garbage collected. Deleting a DaemonSet will clean up the Pods it created.`,
			Description: `
A DaemonSet represents a Kubernetes DaemonSet. In Kubernetes, a DaemonSet ensures that all (or some) Nodes run a copy of a Pod. As nodes are added to the cluster, Pods are added to them. As nodes are removed from the cluster, those Pods are garbage collected. Deleting a DaemonSet will clean up the Pods it created.
`,
			Keywords: []string{
				"kubernetes",
				"daemon",
				"set",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_deployment",
			Category:         "kubernetes",
			ShortDescription: `Deployment inventory represents a Kubernetes Deployment. A Kubernetes Deployment provides declarative for Pods and ReplicaSets. For the Deployment in Intersight, it is an read only object that represent the Deployment. In Kubernetes world, you can describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.`,
			Description: `
Deployment inventory represents a Kubernetes Deployment. A Kubernetes Deployment provides declarative for Pods and ReplicaSets. For the Deployment in Intersight, it is an read only object that represent the Deployment. In Kubernetes world, you can describe a desired state in a Deployment, and the Deployment Controller changes the actual state to the desired state at a controlled rate. You can define Deployments to create new ReplicaSets, or to remove existing Deployments and adopt all their resources with new Deployments.
`,
			Keywords: []string{
				"kubernetes",
				"deployment",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_ingress",
			Category:         "kubernetes",
			ShortDescription: `Ingress inventory represent a Kubernetes Ingress. In Kubernetes, Ingress exposes HTTP and HTTPS routes from outside the cluster to services within the cluster. Traffic routing is controlled by rules defined on the Ingress resource.`,
			Description: `
Ingress inventory represent a Kubernetes Ingress. In Kubernetes, Ingress exposes HTTP and HTTPS routes from outside the cluster to services within the cluster. Traffic routing is controlled by rules defined on the Ingress resource.
`,
			Keywords: []string{
				"kubernetes",
				"ingress",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_network_policy",
			Category:         "kubernetes",
			ShortDescription: `A policy specifying the CIDR for internal networks in a Kubernetes cluster like POD network, and Service network.`,
			Description: `
A policy specifying the CIDR for internal networks in a Kubernetes cluster like POD network, and Service network.
`,
			Keywords: []string{
				"kubernetes",
				"network",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Calico",
					Description: `Calico CNI plugin as described in https://github.com/projectcalico/cni-plugin.`,
				},
				resource.Attribute{
					Name:        "Aci",
					Description: `Cisco ACI Container Network Interface plugin.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_node",
			Category:         "kubernetes",
			ShortDescription: `Node inventory represents the Kubernetes Node. One node represent one Kubernetes worker or master. A Kubernetes node is a worker machine in Kubernetes. A node may be a virtual machine or physical machine. Each node contains the services necessary to run pods and is managed by the master components.`,
			Description: `
Node inventory represents the Kubernetes Node. One node represent one Kubernetes worker or master. A Kubernetes node is a worker machine in Kubernetes. A node may be a virtual machine or physical machine. Each node contains the services necessary to run pods and is managed by the master components.
`,
			Keywords: []string{
				"kubernetes",
				"node",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_node_group_profile",
			Category:         "kubernetes",
			ShortDescription: `A configuration profile for a node group in a Kubernetes cluster.`,
			Description: `
A configuration profile for a node group in a Kubernetes cluster.
`,
			Keywords: []string{
				"kubernetes",
				"node",
				"group",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Worker",
					Description: `Node will be marked as a worker node.`,
				},
				resource.Attribute{
					Name:        "Master",
					Description: `Node will be marked as a master node.`,
				},
				resource.Attribute{
					Name:        "EmbeddedMaster",
					Description: `Node will be both a master and a worker.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_pod",
			Category:         "kubernetes",
			ShortDescription: `Pod represents a Kubernetes Pod. In Intersight, it is a read only model that represent Kubernetes Pod information. In Kubernetes, A Pod is the basic execution unit of a Kubernetes application the smallest and simplest unit in the Kubernetes object model that you create or deploy. A Pod represents processes running on your Cluster.`,
			Description: `
Pod represents a Kubernetes Pod. In Intersight, it is a read only model that represent Kubernetes Pod information. In Kubernetes, A Pod is the basic execution unit of a Kubernetes application the smallest and simplest unit in the Kubernetes object model that you create or deploy. A Pod represents processes running on your Cluster.
`,
			Keywords: []string{
				"kubernetes",
				"pod",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_service",
			Category:         "kubernetes",
			ShortDescription: `Service Inventory represent a Kubernetes Service. In Kubernetes, a Service is an abstraction which defines a logical set of Pods and a policy by which to access them (sometimes this pattern is called a micro-service).`,
			Description: `
Service Inventory represent a Kubernetes Service. In Kubernetes, a Service is an abstraction which defines a logical set of Pods and a policy by which to access them (sometimes this pattern is called a micro-service).
`,
			Keywords: []string{
				"kubernetes",
				"service",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_stateful_set",
			Category:         "kubernetes",
			ShortDescription: `StatefulSet is the workload API object used to manage stateful applications. It manages the deployment and scaling of a set of Pods, and provides guarantees about the ordering and uniqueness of these Pods.`,
			Description: `
StatefulSet is the workload API object used to manage stateful applications. It manages the deployment and scaling of a set of Pods, and provides guarantees about the ordering and uniqueness of these Pods.
`,
			Keywords: []string{
				"kubernetes",
				"stateful",
				"set",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_sys_config_policy",
			Category:         "kubernetes",
			ShortDescription: `A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.`,
			Description: `
A policy specifying system configuration such as timezone, DNS servers, and NTP Servers.
`,
			Keywords: []string{
				"kubernetes",
				"sys",
				"config",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Pacific/Kiritimati",
					Description: ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_trusted_registries_policy",
			Category:         "kubernetes",
			ShortDescription: `A policy specifying self signed docker registries and CA certificates to trust.`,
			Description: `
A policy specifying self signed docker registries and CA certificates to trust.
`,
			Keywords: []string{
				"kubernetes",
				"trusted",
				"registries",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_version",
			Category:         "kubernetes",
			ShortDescription: `A policy capturing software versions for a Kubernetes cluster.`,
			Description: `
A policy capturing software versions for a Kubernetes cluster.
`,
			Keywords: []string{
				"kubernetes",
				"version",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_version_policy",
			Category:         "kubernetes",
			ShortDescription: `Policy that defines which kubernetes version to use.`,
			Description: `
Policy that defines which kubernetes version to use.
`,
			Keywords: []string{
				"kubernetes",
				"version",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_virtual_machine_infrastructure_provider",
			Category:         "kubernetes",
			ShortDescription: `Infrastructure backend for providing virtual machines from a target.`,
			Description: `
Infrastructure backend for providing virtual machines from a target.
`,
			Keywords: []string{
				"kubernetes",
				"virtual",
				"machine",
				"infrastructure",
				"provider",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_virtual_machine_instance_type",
			Category:         "kubernetes",
			ShortDescription: `A policy specifying node configuration for a Virtual Machine.`,
			Description: `
A policy specifying node configuration for a Virtual Machine.
`,
			Keywords: []string{
				"kubernetes",
				"virtual",
				"machine",
				"instance",
				"type",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kubernetes_virtual_machine_node_profile",
			Category:         "kubernetes",
			ShortDescription: `A profile specifying configuration settings for a Virtual Machine node. It is auto-generated based on the NodeGroupProfile and VirtualMachineNodePolicy configuration. Users can do operations like Drain, Cordon, Rebuild on a node.`,
			Description: `
A profile specifying configuration settings for a Virtual Machine node. It is auto-generated based on the NodeGroupProfile and VirtualMachineNodePolicy configuration. Users can do operations like Drain, Cordon, Rebuild on a node.
`,
			Keywords: []string{
				"kubernetes",
				"virtual",
				"machine",
				"node",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "noProvider",
					Description: `Enables the use of no cloud provider.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Out of tree cloud provider, e.g. CPI for vsphere.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kvm_policy",
			Category:         "kvm",
			ShortDescription: `Policy to configure KVM Launch settings.`,
			Description: `
Policy to configure KVM Launch settings.
`,
			Keywords: []string{
				"kvm",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kvm_session",
			Category:         "kvm",
			ShortDescription: `Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.`,
			Description: `
Virtual KVM Session that provides Single Sign-On access to the vKVM console of the server. The vKVM access can be direct or can be tunneled by specifying the tunnel to be used for the access.
`,
			Keywords: []string{
				"kvm",
				"session",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kvm_tunnel",
			Category:         "kvm",
			ShortDescription: `Tunneled Virtual KVM access to the vKVM console of a server. This must be specified while creating the vKVM session to gain tunneled access.`,
			Description: `
Tunneled Virtual KVM access to the vKVM console of a server.
This must be specified while creating the vKVM session to gain tunneled access.
`,
			Keywords: []string{
				"kvm",
				"tunnel",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_kvm_vm_console",
			Category:         "kvm",
			ShortDescription: `API to launch the virtual machine console.`,
			Description: `
API to launch the virtual machine console.
`,
			Keywords: []string{
				"kvm",
				"vm",
				"console",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_account_license_data",
			Category:         "license",
			ShortDescription: `License information for an account.`,
			Description: `
License information for an account.
`,
			Keywords: []string{
				"license",
				"account",
				"data",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_customer_op",
			Category:         "license",
			ShortDescription: `Customer operation object to refresh the registration or re-authenticate, pre-created.`,
			Description: `
Customer operation object to refresh the registration or re-authenticate, pre-created.
`,
			Keywords: []string{
				"license",
				"customer",
				"op",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_iwo_customer_op",
			Category:         "license",
			ShortDescription: `Customer operation object to refresh the registration or re-authenticate, pre-created.`,
			Description: `
Customer operation object to refresh the registration or re-authenticate, pre-created.
`,
			Keywords: []string{
				"license",
				"iwo",
				"customer",
				"op",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_iwo_license_count",
			Category:         "license",
			ShortDescription: `Customer operation object to request reservation code.`,
			Description: `
Customer operation object to request reservation code.
`,
			Keywords: []string{
				"license",
				"iwo",
				"count",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_license_info",
			Category:         "license",
			ShortDescription: `License state information for a specific license entitlement. Essentials license entitlement is supported currently. licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance, the feature set defined for the license entitlement is granted to the customer.`,
			Description: `
License state information for a specific license entitlement. Essentials license entitlement is supported currently.
licenseState attribute is used for license enforcement. When license state is one of TrialPeriod, Compliance, or OutOfCompliance,
the feature set defined for the license entitlement is granted to the customer.
`,
			Keywords: []string{
				"license",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NotLicensed",
					Description: `The license token is neither activated nor registered.`,
				},
				resource.Attribute{
					Name:        "GraceExpired",
					Description: `The license grace period has expired.`,
				},
				resource.Attribute{
					Name:        "TrialPeriod",
					Description: `The 90 days of trial period.`,
				},
				resource.Attribute{
					Name:        "OutOfCompliance",
					Description: `The license is out of compliance.`,
				},
				resource.Attribute{
					Name:        "Compliance",
					Description: `The license is in compliance.`,
				},
				resource.Attribute{
					Name:        "TrialExpired",
					Description: `The trial period of 90 days has expired.`,
				},
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_license_reservation_op",
			Category:         "license",
			ShortDescription: `Customer operation object to request reservation code.`,
			Description: `
Customer operation object to request reservation code.
`,
			Keywords: []string{
				"license",
				"reservation",
				"op",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_license_smartlicense_token",
			Category:         "license",
			ShortDescription: `SmartlicenseToken collection stores license registration tokens.`,
			Description: `
SmartlicenseToken collection stores license registration tokens.
`,
			Keywords: []string{
				"license",
				"smartlicense",
				"token",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ls_service_profile",
			Category:         "ls",
			ShortDescription: `Logical Profile that can be associated to a physical server.`,
			Description: `
Logical Profile that can be associated to a physical server.
`,
			Keywords: []string{
				"ls",
				"service",
				"profile",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_macpool_id_block",
			Category:         "macpool",
			ShortDescription: `A block of contiguous MAC addresses that are part of a pool.`,
			Description: `
A block of contiguous MAC addresses that are part of a pool.
`,
			Keywords: []string{
				"macpool",
				"id",
				"block",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_macpool_lease",
			Category:         "macpool",
			ShortDescription: `Lease represents a single MAC address that is part of the universe, allocated either from a pool or through static assignment.`,
			Description: `
Lease represents a single MAC address that is part of the universe, allocated either from a pool or through static assignment.
`,
			Keywords: []string{
				"macpool",
				"lease",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "dynamic",
					Description: `Identifiers to be allocated by system.`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `Identifiers are assigned by the user.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_macpool_pool",
			Category:         "macpool",
			ShortDescription: `Pool represents a collection of MAC addresses that can be allocated to VNICs of a server profile.`,
			Description: `
Pool represents a collection of MAC addresses that can be allocated to VNICs of a server profile.
`,
			Keywords: []string{
				"macpool",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sequential",
					Description: `Identifiers are assigned in a sequential order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Assignment order is decided by the system.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_macpool_pool_member",
			Category:         "macpool",
			ShortDescription: `PoolMember represents a single MAC address that is part of a pool.`,
			Description: `
PoolMember represents a single MAC address that is part of a pool.
`,
			Keywords: []string{
				"macpool",
				"pool",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_macpool_universe",
			Category:         "macpool",
			ShortDescription: `Universe represents a book keeping container to keep track of all IDs for a given account and pool type.`,
			Description: `
Universe represents a book keeping container to keep track of all IDs for a given account and pool type.
`,
			Keywords: []string{
				"macpool",
				"universe",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_management_controller",
			Category:         "management",
			ShortDescription: `A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.`,
			Description: `
A specialized service processor that monitors the physical state of a server, using sensors and communicating with the system administrator through an independent connection.
`,
			Keywords: []string{
				"management",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_management_entity",
			Category:         "management",
			ShortDescription: `Logical representation that captures the role of each Fabric Interconnect in UCS Manager.`,
			Description: `
Logical representation that captures the role of each Fabric Interconnect in UCS Manager.
`,
			Keywords: []string{
				"management",
				"entity",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_management_interface",
			Category:         "management",
			ShortDescription: `Interface that provides access to the management controller.`,
			Description: `
Interface that provides access to the management controller.
`,
			Keywords: []string{
				"management",
				"interface",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_array",
			Category:         "memory",
			ShortDescription: `Holder housing multiple memory units.`,
			Description: `
Holder housing multiple memory units.
`,
			Keywords: []string{
				"memory",
				"array",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_config_result",
			Category:         "memory",
			ShortDescription: `Result of a previously applied Persistent Memory configuration on a server.`,
			Description: `
Result of a previously applied Persistent Memory configuration on a server.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_configuration",
			Category:         "memory",
			ShortDescription: `Persistent Memory configuration on all the Persistent Memory Modules on a server.`,
			Description: `
Persistent Memory configuration on all the Persistent Memory Modules on a server.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"configuration",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_namespace",
			Category:         "memory",
			ShortDescription: `Persistent Memory Namespace configured within a Persistent Memory Region on a server.`,
			Description: `
Persistent Memory Namespace configured within a Persistent Memory Region on a server.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"namespace",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_namespace_config_result",
			Category:         "memory",
			ShortDescription: `Result of a previously configured Persistent Memory Namespace on a server.`,
			Description: `
Result of a previously configured Persistent Memory Namespace on a server.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"namespace",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_policy",
			Category:         "memory",
			ShortDescription: `The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.`,
			Description: `
The Persistent Memory policy defines the reusable Persistent Memory related configuration that can be applied on many servers. This policy allows the application of Persistent Memory Goals and creation of Persistent Memory Regions and Namespaces. The encryption of the Persistent Memory Modules can be enabled through this policy by providing a passphrase.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "configured-from-intersight",
					Description: `The Persistent Memory Modules are configured from Intersight thorugh Persistent Memory policy.`,
				},
				resource.Attribute{
					Name:        "configured-from-operating-system",
					Description: `The Persistent Memory Modules are configured from operating system thorugh OS tools.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_region",
			Category:         "memory",
			ShortDescription: `Persistent Memory Region configured on the Persistent Memory Modules on a server.`,
			Description: `
Persistent Memory Region configured on the Persistent Memory Modules on a server.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"region",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_persistent_memory_unit",
			Category:         "memory",
			ShortDescription: `Persistent Memory Module on a server.`,
			Description: `
Persistent Memory Module on a server.
`,
			Keywords: []string{
				"memory",
				"persistent",
				"unit",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_memory_unit",
			Category:         "memory",
			ShortDescription: `This represents a memory DIMM on a server.`,
			Description: `
This represents a memory DIMM on a server.
`,
			Keywords: []string{
				"memory",
				"unit",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_meta_definition",
			Category:         "meta",
			ShortDescription: `The meta-data of managed objects and complex types.`,
			Description: `
The meta-data of managed objects and complex types.
`,
			Keywords: []string{
				"meta",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ManagedObject",
					Description: `The meta.Definition object describes a managed object.`,
				},
				resource.Attribute{
					Name:        "ComplexType",
					Description: `The meta.Definition object describes a nested complex type within a managed object.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_network_element",
			Category:         "network",
			ShortDescription: `A Unified Computing Systems (UCS) Fabric Interconnect.`,
			Description: `
A Unified Computing Systems (UCS) Fabric Interconnect.
`,
			Keywords: []string{
				"network",
				"element",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.`,
				},
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.`,
				},
				resource.Attribute{
					Name:        "IntersightStandalone",
					Description: `Intersight Standalone mode of operation.`,
				},
				resource.Attribute{
					Name:        "UCSM",
					Description: `Unified Computing System Manager mode of operation.`,
				},
				resource.Attribute{
					Name:        "Intersight",
					Description: `Intersight managed mode of operation.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `The default state of the sensor (in case no data is received).`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `State of the sensor indicating the sensor's temperature range is okay.`,
				},
				resource.Attribute{
					Name:        "upper-non-recoverable",
					Description: `State of the sensor indicating that the temperature is extremely high above normal range.`,
				},
				resource.Attribute{
					Name:        "upper-critical",
					Description: `State of the sensor indicating that the temperature is above normal range.`,
				},
				resource.Attribute{
					Name:        "upper-non-critical",
					Description: `State of the sensor indicating that the temperature is a little above the normal range.`,
				},
				resource.Attribute{
					Name:        "lower-non-critical",
					Description: `State of the sensor indicating that the temperature is a little below the normal range.`,
				},
				resource.Attribute{
					Name:        "lower-critical",
					Description: `State of the sensor indicating that the temperature is below normal range.`,
				},
				resource.Attribute{
					Name:        "lower-non-recoverable",
					Description: `State of the sensor indicating that the temperature is extremely below normal range.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_network_element_summary",
			Category:         "network",
			ShortDescription: `View MO which aggregates information pertaining to a network element from mutiple MOs.`,
			Description: `
View MO which aggregates information pertaining to a network element from mutiple MOs.
`,
			Keywords: []string{
				"network",
				"element",
				"summary",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.`,
				},
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode.`,
				},
				resource.Attribute{
					Name:        "IntersightStandalone",
					Description: `Intersight Standalone mode of operation.`,
				},
				resource.Attribute{
					Name:        "UCSM",
					Description: `Unified Computing System Manager mode of operation.`,
				},
				resource.Attribute{
					Name:        "Intersight",
					Description: `Intersight managed mode of operation.`,
				},
				resource.Attribute{
					Name:        "unknown",
					Description: `The default state of the sensor (in case no data is received).`,
				},
				resource.Attribute{
					Name:        "ok",
					Description: `State of the sensor indicating the sensor's temperature range is okay.`,
				},
				resource.Attribute{
					Name:        "upper-non-recoverable",
					Description: `State of the sensor indicating that the temperature is extremely high above normal range.`,
				},
				resource.Attribute{
					Name:        "upper-critical",
					Description: `State of the sensor indicating that the temperature is above normal range.`,
				},
				resource.Attribute{
					Name:        "upper-non-critical",
					Description: `State of the sensor indicating that the temperature is a little above the normal range.`,
				},
				resource.Attribute{
					Name:        "lower-non-critical",
					Description: `State of the sensor indicating that the temperature is a little below the normal range.`,
				},
				resource.Attribute{
					Name:        "lower-critical",
					Description: `State of the sensor indicating that the temperature is below normal range.`,
				},
				resource.Attribute{
					Name:        "lower-non-recoverable",
					Description: `State of the sensor indicating that the temperature is extremely below normal range.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_network_fc_zone_info",
			Category:         "network",
			ShortDescription: `FC Zone information of a Fabric Interconnect.`,
			Description: `
FC Zone information of a Fabric Interconnect.
`,
			Keywords: []string{
				"network",
				"fc",
				"zone",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_network_vlan_port_info",
			Category:         "network",
			ShortDescription: `Vlan Port information of a Fabric Interconnect.`,
			Description: `
Vlan Port information of a Fabric Interconnect.
`,
			Keywords: []string{
				"network",
				"vlan",
				"port",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_networkconfig_policy",
			Category:         "networkconfig",
			ShortDescription: `Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.`,
			Description: `
Enable or disable Dynamic DNS, add or update DNS settings for IPv4 and IPv6 on Cisco IMC.
`,
			Keywords: []string{
				"networkconfig",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_apic_cco_post",
			Category:         "niaapi",
			ShortDescription: `The post reporting a new release is available for APIC.`,
			Description: `
The post reporting a new release is available for APIC.
`,
			Keywords: []string{
				"niaapi",
				"apic",
				"cco",
				"post",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_apic_field_notice",
			Category:         "niaapi",
			ShortDescription: `The field notice reporting bug and related software or hardware for APIC.`,
			Description: `
The field notice reporting bug and related software or hardware for APIC.
`,
			Keywords: []string{
				"niaapi",
				"apic",
				"field",
				"notice",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_apic_hweol",
			Category:         "niaapi",
			ShortDescription: `The hardware end of life notice for APIC.`,
			Description: `
The hardware end of life notice for APIC.
`,
			Keywords: []string{
				"niaapi",
				"apic",
				"hweol",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_apic_latest_maintained_release",
			Category:         "niaapi",
			ShortDescription: `This contains the latest maintained release information for each release on APIC.`,
			Description: `
This contains the latest maintained release information for each release on APIC.
`,
			Keywords: []string{
				"niaapi",
				"apic",
				"latest",
				"maintained",
				"release",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_apic_release_recommend",
			Category:         "niaapi",
			ShortDescription: `The recommend version information for each release on APIC.`,
			Description: `
The recommend version information for each release on APIC.
`,
			Keywords: []string{
				"niaapi",
				"apic",
				"release",
				"recommend",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_apic_sweol",
			Category:         "niaapi",
			ShortDescription: `The software end of life notice for APIC.`,
			Description: `
The software end of life notice for APIC.
`,
			Keywords: []string{
				"niaapi",
				"apic",
				"sweol",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_dcnm_cco_post",
			Category:         "niaapi",
			ShortDescription: `The post reporting a new release is available for DCNM.`,
			Description: `
The post reporting a new release is available for DCNM.
`,
			Keywords: []string{
				"niaapi",
				"dcnm",
				"cco",
				"post",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_dcnm_field_notice",
			Category:         "niaapi",
			ShortDescription: `The field notice reporting bug and related software or hardware for DCNM.`,
			Description: `
The field notice reporting bug and related software or hardware for DCNM.
`,
			Keywords: []string{
				"niaapi",
				"dcnm",
				"field",
				"notice",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_dcnm_hweol",
			Category:         "niaapi",
			ShortDescription: `The hardware end of life notice for DCNM.`,
			Description: `
The hardware end of life notice for DCNM.
`,
			Keywords: []string{
				"niaapi",
				"dcnm",
				"hweol",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_dcnm_latest_maintained_release",
			Category:         "niaapi",
			ShortDescription: `This contains the latest maintained release information for each release on DCNM.`,
			Description: `
This contains the latest maintained release information for each release on DCNM.
`,
			Keywords: []string{
				"niaapi",
				"dcnm",
				"latest",
				"maintained",
				"release",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_dcnm_release_recommend",
			Category:         "niaapi",
			ShortDescription: `The recommend version information for each release on DCNM.`,
			Description: `
The recommend version information for each release on DCNM.
`,
			Keywords: []string{
				"niaapi",
				"dcnm",
				"release",
				"recommend",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_dcnm_sweol",
			Category:         "niaapi",
			ShortDescription: `The software end of life notice for DCNM.`,
			Description: `
The software end of life notice for DCNM.
`,
			Keywords: []string{
				"niaapi",
				"dcnm",
				"sweol",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_file_downloader",
			Category:         "niaapi",
			ShortDescription: `Provide a presigned url to download the metadata file from server.`,
			Description: `
Provide a presigned url to download the metadata file from server.
`,
			Keywords: []string{
				"niaapi",
				"file",
				"downloader",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_nia_metadata",
			Category:         "niaapi",
			ShortDescription: `Contains the latest Metadata available for download from server.`,
			Description: `
Contains the latest Metadata available for download from server.
`,
			Keywords: []string{
				"niaapi",
				"nia",
				"metadata",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_nib_file_downloader",
			Category:         "niaapi",
			ShortDescription: `Provide a presigned url to download the metadata file from server.`,
			Description: `
Provide a presigned url to download the metadata file from server.
`,
			Keywords: []string{
				"niaapi",
				"nib",
				"file",
				"downloader",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_nib_metadata",
			Category:         "niaapi",
			ShortDescription: `Contains the latest metadata available for download from server.`,
			Description: `
Contains the latest metadata available for download from server.
`,
			Keywords: []string{
				"niaapi",
				"nib",
				"metadata",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niaapi_version_regex",
			Category:         "niaapi",
			ShortDescription: `The regular expression pattern to recongnize the version string.`,
			Description: `
The regular expression pattern to recongnize the version string.
`,
			Keywords: []string{
				"niaapi",
				"version",
				"regex",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_apic_fan_details",
			Category:         "niatelemetry",
			ShortDescription: `Object to capture the fan details in APIC.`,
			Description: `
Object to capture the fan details in APIC.
`,
			Keywords: []string{
				"niatelemetry",
				"apic",
				"fan",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_apic_fex_details",
			Category:         "niatelemetry",
			ShortDescription: `Object to capture FEX details in APIC.`,
			Description: `
Object to capture FEX details in APIC.
`,
			Keywords: []string{
				"niatelemetry",
				"apic",
				"fex",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_apic_psu_details",
			Category:         "niatelemetry",
			ShortDescription: `Object to capture PSU details in APIC.`,
			Description: `
Object to capture PSU details in APIC.
`,
			Keywords: []string{
				"niatelemetry",
				"apic",
				"psu",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_apic_ui_page_counts",
			Category:         "niatelemetry",
			ShortDescription: `Object to capture the UI page counts in APIC.`,
			Description: `
Object to capture the UI page counts in APIC.
`,
			Keywords: []string{
				"niatelemetry",
				"apic",
				"ui",
				"page",
				"counts",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_app_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of apps installed on Nexus Dashboard.`,
			Description: `
Details of apps installed on Nexus Dashboard.
`,
			Keywords: []string{
				"niatelemetry",
				"app",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_epg",
			Category:         "niatelemetry",
			ShortDescription: `Object is available at End Point Group scope. This currently applies only to the APIC environemt.`,
			Description: `
Object is available at End Point Group scope. This currently applies only to the APIC environemt.
`,
			Keywords: []string{
				"niatelemetry",
				"epg",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_fault",
			Category:         "niatelemetry",
			ShortDescription: `Object is available at Fault scope in a fabric and provides details about a fault occurred.`,
			Description: `
Object is available at Fault scope in a fabric and provides details about a fault occurred.
`,
			Keywords: []string{
				"niatelemetry",
				"fault",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_lc",
			Category:         "niatelemetry",
			ShortDescription: `Object is available at Line Card scope.`,
			Description: `
Object is available at Line Card scope.
`,
			Keywords: []string{
				"niatelemetry",
				"lc",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_mso_schema_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of schema in Multi-Site Orchestrator.`,
			Description: `
Details of schema in Multi-Site Orchestrator.
`,
			Keywords: []string{
				"niatelemetry",
				"mso",
				"schema",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_mso_site_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of sites in Multi-Site Orchestrator.`,
			Description: `
Details of sites in Multi-Site Orchestrator.
`,
			Keywords: []string{
				"niatelemetry",
				"mso",
				"site",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_mso_tenant_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of tenant in Multi-Site Orchestrator.`,
			Description: `
Details of tenant in Multi-Site Orchestrator.
`,
			Keywords: []string{
				"niatelemetry",
				"mso",
				"tenant",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nexus_dashboard_controller_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of controller added to NexusDashboard.`,
			Description: `
Details of controller added to NexusDashboard.
`,
			Keywords: []string{
				"niatelemetry",
				"nexus",
				"dashboard",
				"controller",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nexus_dashboard_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of NexusDashboard.`,
			Description: `
Details of NexusDashboard.
`,
			Keywords: []string{
				"niatelemetry",
				"nexus",
				"dashboard",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nexus_dashboard_memory_details",
			Category:         "niatelemetry",
			ShortDescription: `Details of Nexus Dashboard's memory.`,
			Description: `
Details of Nexus Dashboard's memory.
`,
			Keywords: []string{
				"niatelemetry",
				"nexus",
				"dashboard",
				"memory",
				"details",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nexus_dashboards",
			Category:         "niatelemetry",
			ShortDescription: `Object is available for Nexus Dashboard devices.`,
			Description: `
Object is available for Nexus Dashboard devices.
`,
			Keywords: []string{
				"niatelemetry",
				"nexus",
				"dashboards",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nia_feature_usage",
			Category:         "niatelemetry",
			ShortDescription: `Object available at Device connector scope for feature and fabric information. This applies to APIC environment currently.`,
			Description: `
Object available at Device connector scope for feature and fabric information. This applies to APIC environment currently.
`,
			Keywords: []string{
				"niatelemetry",
				"nia",
				"feature",
				"usage",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nia_inventory",
			Category:         "niatelemetry",
			ShortDescription: `Inventory object available per device scope. This common object holds a device level information.`,
			Description: `
Inventory object available per device scope. This common object holds a device level information.
`,
			Keywords: []string{
				"niatelemetry",
				"nia",
				"inventory",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nia_inventory_dcnm",
			Category:         "niatelemetry",
			ShortDescription: `Inventory Object available for DCNM-scoped attributes.`,
			Description: `
Inventory Object available for DCNM-scoped attributes.
`,
			Keywords: []string{
				"niatelemetry",
				"nia",
				"inventory",
				"dcnm",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nia_inventory_fabric",
			Category:         "niatelemetry",
			ShortDescription: `Inventory Object available for Fabric-scoped attributes.`,
			Description: `
Inventory Object available for Fabric-scoped attributes.
`,
			Keywords: []string{
				"niatelemetry",
				"nia",
				"inventory",
				"fabric",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_nia_license_state",
			Category:         "niatelemetry",
			ShortDescription: `Object available at device scope for license information. This determines the usage of this attribute.`,
			Description: `
Object available at device scope for license information. This determines the usage of this attribute.
`,
			Keywords: []string{
				"niatelemetry",
				"nia",
				"license",
				"state",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_niatelemetry_tenant",
			Category:         "niatelemetry",
			ShortDescription: `Object is available at Tenant scope.`,
			Description: `
Object is available at Tenant scope.
`,
			Keywords: []string{
				"niatelemetry",
				"tenant",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ntp_policy",
			Category:         "ntp",
			ShortDescription: `Policy to configure the NTP Servers.`,
			Description: `
Policy to configure the NTP Servers.
`,
			Keywords: []string{
				"ntp",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Pacific/Kiritimati",
					Description: ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_organization_organization",
			Category:         "organization",
			ShortDescription: `Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.`,
			Description: `
Organization provides multi-tenancy within an account. Multiple organizations can be present under an account. Resources are associated to organization using resource groups. Organization can have heterogeneous resources. Resources can be shared among multiple organizations. Organizations are associated to user permissions and privileges can be specified to provide access control. User can have access to multiple organizations in same permission and with different privileges on each organization.
`,
			Keywords: []string{
				"organization",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_os_catalog",
			Category:         "os",
			ShortDescription: `A catalog of operating systems related objects such as configuration files and post install scripts. Each user account will have a local OS catalog where account users can store their private configuration files and post install scripts. Cisco provides validated answer files and post install scripts to Intersight users via shared catalogs. Intersight users will be able to read, use these files and scripts in OS installation within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.`,
			Description: `
A catalog of operating systems related objects such as configuration files and post install scripts. Each user account will have a local OS catalog where account users can store their private configuration files and post install scripts.
Cisco provides validated answer files and post install scripts to Intersight users via shared catalogs. Intersight users will be able to read, use these files and scripts in OS installation within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
`,
			Keywords: []string{
				"os",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_os_configuration_file",
			Category:         "os",
			ShortDescription: `A ConfigurationFile is an OS specific answer file that helps with the unattended installation. The file can also be a template file with placeholders instead of actual answers. Intersight supports the golang template syntax specified in https://golang.org/pkg/text/template/. The template supports placeholders for all the properties of os.Answers MO type as well as any additional user-defined properties. The values for these placeholders shall be given during OS installation in the form of os.Answers type and 'additionalProperties' in os.OsInstall object.`,
			Description: `
A ConfigurationFile is an OS specific answer file that helps with the unattended
installation.
The file can also be a template file with placeholders instead of actual answers.
Intersight supports the golang template syntax specified in https://golang.org/pkg/text/template/.
The template supports placeholders for all the properties of os.Answers MO type
as well as any additional user-defined properties. The values for these placeholders
shall be given during OS installation in the form of os.Answers type and 'additionalProperties' in
os.OsInstall object.
`,
			Keywords: []string{
				"os",
				"configuration",
				"file",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_os_distribution",
			Category:         "os",
			ShortDescription: `Intersight has the distribution details for all the Intersight supported OS distributions. There will be a Distribution object for each supported OS.`,
			Description: `
Intersight has the distribution details for all the Intersight supported OS
distributions. There will be a Distribution object for each supported OS.
`,
			Keywords: []string{
				"os",
				"distribution",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_os_install",
			Category:         "os",
			ShortDescription: `This MO models the target server, answers and other properties needed for OS installation. The OS installation can be started in the target server by doing a POST on this MO. The requests to this MO starts a OS installation workflow that can be tracked using workflow engine MO workflow.WorkflowInfo.`,
			Description: `
This MO models the target server, answers and other properties needed for
OS installation. The OS installation can be started in the target server by doing
a POST on this MO.
The requests to this MO starts a OS installation workflow that can be tracked
using workflow engine MO workflow.WorkflowInfo.
`,
			Keywords: []string{
				"os",
				"install",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vMedia",
					Description: `OS image is mounted as vMedia in target server for OS installation.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_os_supported_version",
			Category:         "os",
			ShortDescription: `The supported operating system version by SCU. The API is mainly for UI operation. In the software management page, operating system configuration will be created by providing the vendor and version. The version will be filtered based on vendor.`,
			Description: `
The supported operating system version by SCU. The API is mainly for UI operation. In the software management page,
operating system configuration will be created by providing the vendor and version. The version will be filtered
based on vendor.
`,
			Keywords: []string{
				"os",
				"supported",
				"version",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_pci_coprocessor_card",
			Category:         "pci",
			ShortDescription: `PCIe Compression and Cryptographic CPU Offload Card.`,
			Description: `
PCIe Compression and Cryptographic CPU Offload Card.
`,
			Keywords: []string{
				"pci",
				"coprocessor",
				"card",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_pci_device",
			Category:         "pci",
			ShortDescription: `PCI device present in a server.`,
			Description: `
PCI device present in a server.
`,
			Keywords: []string{
				"pci",
				"device",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_pci_link",
			Category:         "pci",
			ShortDescription: `The PCI Switch Link connected to PCIe Switch.`,
			Description: `
The PCI Switch Link connected to PCIe Switch.
`,
			Keywords: []string{
				"pci",
				"link",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_pci_switch",
			Category:         "pci",
			ShortDescription: `PCI Switch present in a server connected to two GPUs and one PCIe adapter.`,
			Description: `
PCI Switch present in a server connected to two GPUs and one PCIe adapter.
`,
			Keywords: []string{
				"pci",
				"switch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_port_group",
			Category:         "port",
			ShortDescription: `Holder for multiple ports. A switch card will have one or more port groups.`,
			Description: `
Holder for multiple ports. A switch card will have one or more port groups.
`,
			Keywords: []string{
				"port",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_port_mac_binding",
			Category:         "port",
			ShortDescription: `Establishes relationship between the ports and connected end points based on LLDP TLVs.`,
			Description: `
Establishes relationship between the ports and connected end points based on LLDP TLVs.
`,
			Keywords: []string{
				"port",
				"mac",
				"binding",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_port_sub_group",
			Category:         "port",
			ShortDescription: `Holder for multiple ports within a portGroup. SubGroup represents a break-out port group on the Fabric Interconnect.`,
			Description: `
Holder for multiple ports within a portGroup. SubGroup represents a break-out port group on the Fabric Interconnect.
`,
			Keywords: []string{
				"port",
				"sub",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_processor_unit",
			Category:         "processor",
			ShortDescription: `The CPU present on a server.`,
			Description: `
The CPU present on a server.
`,
			Keywords: []string{
				"processor",
				"unit",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recommendation_capacity_runway",
			Category:         "recommendation",
			ShortDescription: `Entity representing the new capacity runway based on recommendations.`,
			Description: `
Entity representing the new capacity runway based on recommendations.
`,
			Keywords: []string{
				"recommendation",
				"capacity",
				"runway",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "TB",
					Description: `The Enum value TB represents that the measurement unit is in terabytes.`,
				},
				resource.Attribute{
					Name:        "MB",
					Description: `The Enum value MB represents that the measurement unit is in megabytes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recommendation_physical_item",
			Category:         "recommendation",
			ShortDescription: `Entity representing the recommended physical device.`,
			Description: `
Entity representing the recommended physical device.
`,
			Keywords: []string{
				"recommendation",
				"physical",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Disk",
					Description: `The Enum value Disk represents that the item type recommended is a Physical Disk.`,
				},
				resource.Attribute{
					Name:        "Node",
					Description: `The Enum value Node represents that the item type recommended is a Storage Node.`,
				},
				resource.Attribute{
					Name:        "Cluster",
					Description: `The Enum value Cluster represents that the item type recommended is a HyperFlex Cluster.`,
				},
				resource.Attribute{
					Name:        "TB",
					Description: `The Enum value TB represents that the measurement unit is in terabytes.`,
				},
				resource.Attribute{
					Name:        "MB",
					Description: `The Enum value MB represents that the measurement unit is in megabytes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_backup_config_policy",
			Category:         "recovery",
			ShortDescription: `Backup config policy which contains all the required inputs to do backup on a local or remote server.`,
			Description: `
Backup config policy which contains all the required inputs to do backup on a local or remote server.
`,
			Keywords: []string{
				"recovery",
				"backup",
				"config",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Network Share",
					Description: `The backup is stored remotely on a separate server.`,
				},
				resource.Attribute{
					Name:        "Local Storage",
					Description: `The backup is stored locally on the endpoint.`,
				},
				resource.Attribute{
					Name:        "SCP",
					Description: `Secure Copy Protocol (SCP) to access the file server.`,
				},
				resource.Attribute{
					Name:        "SFTP",
					Description: `SSH File Transfer Protocol (SFTP) to access file server.`,
				},
				resource.Attribute{
					Name:        "FTP",
					Description: `File Transfer Protocol (FTP) to access file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_backup_profile",
			Category:         "recovery",
			ShortDescription: `Backup profile to initiate on-demand or scheduled backups at end points.`,
			Description: `
Backup profile to initiate on-demand or scheduled backups at end points.
`,
			Keywords: []string{
				"recovery",
				"backup",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_config_result",
			Category:         "recovery",
			ShortDescription: `Profile configuration (deploy, validation) results with the overall state and detailed result messages.`,
			Description: `
Profile configuration (deploy, validation) results with the overall state and detailed result messages.
`,
			Keywords: []string{
				"recovery",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_config_result_entry",
			Category:         "recovery",
			ShortDescription: `An entry that describes the result of a Backup Profile state on the end device.`,
			Description: `
An entry that describes the result of a Backup Profile state on the end device.
`,
			Keywords: []string{
				"recovery",
				"config",
				"result",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_on_demand_backup",
			Category:         "recovery",
			ShortDescription: `Handles requests for on demand backup for a given endpoint.`,
			Description: `
Handles requests for on demand backup for a given endpoint.
`,
			Keywords: []string{
				"recovery",
				"on",
				"demand",
				"backup",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Network Share",
					Description: `The backup is stored remotely on a separate server.`,
				},
				resource.Attribute{
					Name:        "Local Storage",
					Description: `The backup is stored locally on the endpoint.`,
				},
				resource.Attribute{
					Name:        "SCP",
					Description: `Secure Copy Protocol (SCP) to access the file server.`,
				},
				resource.Attribute{
					Name:        "SFTP",
					Description: `SSH File Transfer Protocol (SFTP) to access file server.`,
				},
				resource.Attribute{
					Name:        "FTP",
					Description: `File Transfer Protocol (FTP) to access file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_restore",
			Category:         "recovery",
			ShortDescription: `Triggers a restore operation on the target endpoint.`,
			Description: `
Triggers a restore operation on the target endpoint.
`,
			Keywords: []string{
				"recovery",
				"restore",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_recovery_schedule_config_policy",
			Category:         "recovery",
			ShortDescription: `Base Schedule config which contains all the required inputs to do schedule on a local or remote server.`,
			Description: `
Base Schedule config which contains all the required inputs to do schedule on a local or remote server.
`,
			Keywords: []string{
				"recovery",
				"schedule",
				"config",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_resource_group",
			Category:         "resource",
			ShortDescription: `A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.`,
			Description: `
A group of REST resources, such as a group of compute.Blade MOs. A ResourceGroup can contain static members which are specified as a set of object references, or it can contain dynamic members, which are specified through OData query filters. A Resource can be part of multiple resource groups.
`,
			Keywords: []string{
				"resource",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Allow-Selectors",
					Description: `Resources will be added to resource groups based on ODATA filter. Multiple resource group can be created to organize resources.`,
				},
				resource.Attribute{
					Name:        "Allow-All",
					Description: `All resources will become part of the Resource Group. Only one resource group can be created to organize resources.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_resource_group_member",
			Category:         "resource",
			ShortDescription: `A resolved member of a resource group.`,
			Description: `
A resolved member of a resource group.
`,
			Keywords: []string{
				"resource",
				"group",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_resource_license_resource_count",
			Category:         "resource",
			ShortDescription: `LicenseResourceCount tracks the server count info for 3 different licensing tiers.`,
			Description: `
LicenseResourceCount tracks the server count info for 3 different licensing tiers.
`,
			Keywords: []string{
				"resource",
				"license",
				"count",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_resource_membership",
			Category:         "resource",
			ShortDescription: `ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.`,
			Description: `
ResourceMembership represents a resource's associated groups, organizations and the permissions associated to this resource via organizations.
`,
			Keywords: []string{
				"resource",
				"membership",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_resource_membership_holder",
			Category:         "resource",
			ShortDescription: `A holder of REST resources and their membership.`,
			Description: `
A holder of REST resources and their membership.
`,
			Keywords: []string{
				"resource",
				"membership",
				"holder",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_sdcard_policy",
			Category:         "sdcard",
			ShortDescription: `Policy for configuring SD Card settings on endpoint.`,
			Description: `
Policy for configuring SD Card settings on endpoint.
`,
			Keywords: []string{
				"sdcard",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_sdwan_profile",
			Category:         "sdwan",
			ShortDescription: `A profile that specifies configuration settings for a SDWAN router.`,
			Description: `
A profile that specifies configuration settings for a SDWAN router.
`,
			Keywords: []string{
				"sdwan",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_sdwan_router_node",
			Category:         "sdwan",
			ShortDescription: `Configuration settings for a SDWAN vEdge router.`,
			Description: `
Configuration settings for a SDWAN vEdge router.
`,
			Keywords: []string{
				"sdwan",
				"router",
				"node",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_sdwan_router_policy",
			Category:         "sdwan",
			ShortDescription: `A policy specifying SD-WAN router details.`,
			Description: `
A policy specifying SD-WAN router details.
`,
			Keywords: []string{
				"sdwan",
				"router",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Typical",
					Description: `Typical deployment configuration with 4 vCPUs and 4GB RAM.`,
				},
				resource.Attribute{
					Name:        "Minimal",
					Description: `Minimal deployment configuration with 2 vCPUs and 4GB RAM.`,
				},
				resource.Attribute{
					Name:        "Single",
					Description: `Singly terminated WANs ar evenly distributed across SD-WAN router nodes. A given WAN connection is available only on one of the router nodes.`,
				},
				resource.Attribute{
					Name:        "Dual",
					Description: `Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_sdwan_vmanage_account_policy",
			Category:         "sdwan",
			ShortDescription: `A policy specifying vManage account details.`,
			Description: `
A policy specifying vManage account details.
`,
			Keywords: []string{
				"sdwan",
				"vmanage",
				"account",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_search_search_item",
			Category:         "search",
			ShortDescription: ``,
			Description: `
The Search service entry point to search Intersight REST resources using OData query syntax.
See [Search API query syntax](/apidocs/introduction/query/#search-api) for details
about the query syntax.
`,
			Keywords: []string{
				"search",
				"item",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_search_tag_item",
			Category:         "search",
			ShortDescription: ``,
			Description: `
The TagItems service entry point to search Tags across all Intersight REST resources using OData Query syntax.
See [Search Tags API query syntax](/apidocs/introduction/query/#search-tags-api) for details about the tag query syntax.
`,
			Keywords: []string{
				"search",
				"tag",
				"item",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_security_unit",
			Category:         "security",
			ShortDescription: `The crypto card present on a server.`,
			Description: `
The crypto card present on a server.
`,
			Keywords: []string{
				"security",
				"unit",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_server_config_change_detail",
			Category:         "server",
			ShortDescription: `The configuration change details are captured here.`,
			Description: `
The configuration change details are captured here.
`,
			Keywords: []string{
				"server",
				"config",
				"change",
				"detail",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Pending-changes",
					Description: `Config change flag represents changes are due to not deployed changes from Intersight.`,
				},
				resource.Attribute{
					Name:        "Drift-changes",
					Description: `Config change flag represents changes are due to endpoint configuration changes.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `The 'none' operation/state.Indicates a configuration profile has been deployed, and the desired configuration matches the actual device configuration.`,
				},
				resource.Attribute{
					Name:        "Created",
					Description: `The 'create' operation/state.Indicates a configuration profile has been created and associated with a device, but the configuration specified in the profilehas not been applied yet. For example, this could happen when the user creates a server profile and has not deployed the profile yet.`,
				},
				resource.Attribute{
					Name:        "Modified",
					Description: `The 'update' operation/state.Indicates some of the desired configuration changes specified in a profile have not been been applied to the associated device.This happens when the user has made changes to a profile and has not deployed the changes yet, or when the workflow to applythe configuration changes has not completed succesfully.`,
				},
				resource.Attribute{
					Name:        "Deleted",
					Description: `The 'delete' operation/state.Indicates a configuration profile has been been disassociated from a device and the user has not undeployed these changes yet.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_server_config_import",
			Category:         "server",
			ShortDescription: `Configuration import action will import the existing configuration from physical server and generate Intersight policies and server profile from it. At end of successful import, server will be assigned to the generated profile which has policies associated with it. No server profile or policies will be generated if configuration import fails.`,
			Description: `
Configuration import action will import the existing configuration from physical server and generate Intersight policies and server profile from it. At end of successful import, server will be assigned to the generated profile which has policies associated with it. No server profile or policies will be generated if configuration import fails.
`,
			Keywords: []string{
				"server",
				"config",
				"import",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_server_config_result",
			Category:         "server",
			ShortDescription: `The profile configuration (deploy, validation) results with the overall state and detailed result messages.`,
			Description: `
The profile configuration (deploy, validation) results with the overall state and detailed result messages.
`,
			Keywords: []string{
				"server",
				"config",
				"result",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_server_config_result_entry",
			Category:         "server",
			ShortDescription: `The profile configuration (deploy, validation) results details information.`,
			Description: `
The profile configuration (deploy, validation) results details information.
`,
			Keywords: []string{
				"server",
				"config",
				"result",
				"entry",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_server_profile",
			Category:         "server",
			ShortDescription: `A profile specifying configuration settings for a physical server.`,
			Description: `
A profile specifying configuration settings for a physical server.
`,
			Keywords: []string{
				"server",
				"profile",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Standalone",
					Description: `Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.`,
				},
				resource.Attribute{
					Name:        "FIAttached",
					Description: `Servers which are connected to a Fabric Interconnect that is managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "instance",
					Description: `The profile defines the configuration for a specific instance of a target.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_smtp_policy",
			Category:         "smtp",
			ShortDescription: `Name that identifies the SMTP Policy.`,
			Description: `
Name that identifies the SMTP Policy.
`,
			Keywords: []string{
				"smtp",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "critical",
					Description: `Minimum severity to report is critical.`,
				},
				resource.Attribute{
					Name:        "condition",
					Description: `Minimum severity to report is informational.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `Minimum severity to report is warning.`,
				},
				resource.Attribute{
					Name:        "minor",
					Description: `Minimum severity to report is minor.`,
				},
				resource.Attribute{
					Name:        "major",
					Description: `Minimum severity to report is major.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_snmp_policy",
			Category:         "snmp",
			ShortDescription: `Policy to configure SNMP settings on endpoint.`,
			Description: `
Policy to configure SNMP settings on endpoint.
`,
			Keywords: []string{
				"snmp",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Disabled",
					Description: `Blocks access to the information in the inventory tables.`,
				},
				resource.Attribute{
					Name:        "Limited",
					Description: `Partial access to read the information in the inventory tables.`,
				},
				resource.Attribute{
					Name:        "Full",
					Description: `Full access to read the information in the inventory tables.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_appliance_distributable",
			Category:         "software",
			ShortDescription: `Appliance image distributed by Cisco. This image is required to upgrade the on-premise Intersight Appliance. There are two use cases. In Intersight SaaS, the object represents a downloadable image, whereas on the Appliance the represents the image that is uploaded by the user and to be used for upgrade.`,
			Description: `
Appliance image distributed by Cisco. This image is required to upgrade the on-premise Intersight Appliance.
There are two use cases. In Intersight SaaS, the object represents a downloadable image, whereas on the
Appliance the represents the image that is uploaded by the user and to be used for upgrade.
`,
			Keywords: []string{
				"software",
				"appliance",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_hcl_meta",
			Category:         "software",
			ShortDescription: `A JSON file wth HCL metadata uploaded for consumption by the HCL service.`,
			Description: `
A JSON file wth HCL metadata uploaded for consumption by the HCL service.
`,
			Keywords: []string{
				"software",
				"hcl",
				"meta",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Full",
					Description: `Indicates that the JSON File does have full content for HCL metadata.`,
				},
				resource.Attribute{
					Name:        "Incremental",
					Description: `Indicates that the JSON File does have only the diff of the Hcl meta to be uploaded.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_hyperflex_bundle_distributable",
			Category:         "software",
			ShortDescription: `A HyperFlex image bundle distributed by Cisco for Private Appliance.`,
			Description: `
A HyperFlex image bundle distributed by Cisco for Private Appliance.
`,
			Keywords: []string{
				"software",
				"hyperflex",
				"bundle",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_hyperflex_distributable",
			Category:         "software",
			ShortDescription: `A HyperFlex image distributed by Cisco.`,
			Description: `
A HyperFlex image distributed by Cisco.
`,
			Keywords: []string{
				"software",
				"hyperflex",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_solution_distributable",
			Category:         "software",
			ShortDescription: `A solution image distributed by Cisco.`,
			Description: `
A solution image distributed by Cisco.
`,
			Keywords: []string{
				"software",
				"solution",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
				resource.Attribute{
					Name:        "osimage",
					Description: `The solution OS image for deployment.`,
				},
				resource.Attribute{
					Name:        "script",
					Description: `The Python script for the solution VM configuration and deployment.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_ucsd_bundle_distributable",
			Category:         "software",
			ShortDescription: `A UCSD connector pack image bundle distributed by Cisco for Private Appliance.`,
			Description: `
A UCSD connector pack image bundle distributed by Cisco for Private Appliance.
`,
			Keywords: []string{
				"software",
				"ucsd",
				"bundle",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_software_ucsd_distributable",
			Category:         "software",
			ShortDescription: `A UCSD connector pack image distributed by Cisco.`,
			Description: `
A UCSD connector pack image distributed by Cisco.
`,
			Keywords: []string{
				"software",
				"ucsd",
				"distributable",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_authorization",
			Category:         "softwarerepository",
			ShortDescription: `User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user.`,
			Description: `
User's consent for Intersight to contact an external software repository such as cisco.com, on the behalf of the user.
`,
			Keywords: []string{
				"softwarerepository",
				"authorization",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Cisco",
					Description: `External repository hosted on cisco.com.`,
				},
				resource.Attribute{
					Name:        "IntersightCloud",
					Description: `Repository hosted by the Intersight Cloud.`,
				},
				resource.Attribute{
					Name:        "LocalMachine",
					Description: `The file is available on the local client machine. Used as an upload source type.`,
				},
				resource.Attribute{
					Name:        "NetworkShare",
					Description: `External repository in the customer datacenter. This will typically be a file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_cached_image",
			Category:         "softwarerepository",
			ShortDescription: `The image cached in the customer's datacenter.`,
			Description: `
The image cached in the customer's datacenter.
`,
			Keywords: []string{
				"softwarerepository",
				"cached",
				"image",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_catalog",
			Category:         "softwarerepository",
			ShortDescription: `A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.`,
			Description: `
A container MO that holds references to the files in an account's image repository. It is internally created for each account and is used to hold information about all user uploaded files.
`,
			Keywords: []string{
				"softwarerepository",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_category_mapper",
			Category:         "softwarerepository",
			ShortDescription: `Maps a Cisco software repository image category identifier to its applicable hardware models.`,
			Description: `
Maps a Cisco software repository image category identifier to its applicable hardware models.
`,
			Keywords: []string{
				"softwarerepository",
				"category",
				"mapper",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Distributable",
					Description: `Stores firmware host utility images and fabric images.`,
				},
				resource.Attribute{
					Name:        "DriverDistributable",
					Description: `Stores driver distributable images.`,
				},
				resource.Attribute{
					Name:        "ServerConfigurationUtilityDistributable",
					Description: `Stores server configuration utility images.`,
				},
				resource.Attribute{
					Name:        "OperatingSystemFile",
					Description: `Stores operating system iso images.`,
				},
				resource.Attribute{
					Name:        "HyperflexDistributable",
					Description: `It stores HyperFlex images.`,
				},
				resource.Attribute{
					Name:        "Cisco",
					Description: `External repository hosted on cisco.com.`,
				},
				resource.Attribute{
					Name:        "IntersightCloud",
					Description: `Repository hosted by the Intersight Cloud.`,
				},
				resource.Attribute{
					Name:        "LocalMachine",
					Description: `The file is available on the local client machine. Used as an upload source type.`,
				},
				resource.Attribute{
					Name:        "NetworkShare",
					Description: `External repository in the customer datacenter. This will typically be a file server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_category_mapper_model",
			Category:         "softwarerepository",
			ShortDescription: `Maps a Cisco hardware model Series to its applicable hardware models.`,
			Description: `
Maps a Cisco hardware model Series to its applicable hardware models.
`,
			Keywords: []string{
				"softwarerepository",
				"category",
				"mapper",
				"model",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_category_support_constraint",
			Category:         "softwarerepository",
			ShortDescription: `Defines constraints for models which are supported from certain minimum image version.`,
			Description: `
Defines constraints for models which are supported from certain minimum image version.
`,
			Keywords: []string{
				"softwarerepository",
				"category",
				"support",
				"constraint",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_download_spec",
			Category:         "softwarerepository",
			ShortDescription: `The URL, certificate and other associated information required to download an image listed in an Intersight catalog.`,
			Description: `
The URL, certificate and other associated information required to download an image listed in an Intersight catalog.
`,
			Keywords: []string{
				"softwarerepository",
				"download",
				"spec",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_operating_system_file",
			Category:         "softwarerepository",
			ShortDescription: `An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.`,
			Description: `
An operating system image that resides either in an external repository or has been imported to the local repository. If the file is available in the local repository, it is marked as cached. If not, it represents a pointer to a file in an external repository.
`,
			Keywords: []string{
				"softwarerepository",
				"operating",
				"system",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action should be taken on the imported file.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedUploadUrl",
					Description: `Generate pre signed URL of file for importing into the repository.`,
				},
				resource.Attribute{
					Name:        "GeneratePreSignedDownloadUrl",
					Description: `Generate pre signed URL of file in the repository to download.`,
				},
				resource.Attribute{
					Name:        "CompleteImportProcess",
					Description: `Mark that the import process of the file into the repository is complete.`,
				},
				resource.Attribute{
					Name:        "MarkImportFailed",
					Description: `Mark to indicate that the import process of the file into the repository failed.`,
				},
				resource.Attribute{
					Name:        "PreCache",
					Description: `Cache the file into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `The cancel import process for the file into the repository.`,
				},
				resource.Attribute{
					Name:        "Extract",
					Description: `The action to extract the file in the external repository.`,
				},
				resource.Attribute{
					Name:        "Evict",
					Description: `Evict the cached file from the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "ReadyForImport",
					Description: `The image is ready to be imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Importing",
					Description: `The image is being imported into the repository.`,
				},
				resource.Attribute{
					Name:        "Imported",
					Description: `The image has been extracted and imported into the repository.`,
				},
				resource.Attribute{
					Name:        "PendingExtraction",
					Description: `Indicates that the image has been imported but not extracted in the repository.`,
				},
				resource.Attribute{
					Name:        "Extracting",
					Description: `Indicates that the image is being extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `Indicates that the image has been extracted into the repository.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The image import from an external source to the repository has failed.`,
				},
				resource.Attribute{
					Name:        "MetaOnly",
					Description: `The image is present in an external repository.`,
				},
				resource.Attribute{
					Name:        "ReadyForCache",
					Description: `The image is ready to be cached into the Intersight Appliance.`,
				},
				resource.Attribute{
					Name:        "Caching",
					Description: `Indicates that the image is being cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Indicates that the image has been cached into the Intersight Appliance or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "CachingFailed",
					Description: `Indicates that the image caching into the Intersight Appliance failed or endpoint cache.`,
				},
				resource.Attribute{
					Name:        "Corrupted",
					Description: `Indicates that the image in the local repository (or endpoint cache) has been corrupted after it was cached.`,
				},
				resource.Attribute{
					Name:        "Evicted",
					Description: `Indicates that the image has been evicted from the Intersight Appliance (or endpoint cache) to reclaim storage space.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_softwarerepository_release",
			Category:         "softwarerepository",
			ShortDescription: `A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.`,
			Description: `
A Cisco release containing one or more firmware images. Cisco releases images for rack server components or blade server components or for Fabric Interconnect components. The version for the firmware images is the same as specific Cisco release version.
`,
			Keywords: []string{
				"softwarerepository",
				"release",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "FabricSwitch",
					Description: `The images in a release that correspond to Fabric Interconnect switches.`,
				},
				resource.Attribute{
					Name:        "ComputeSystem",
					Description: `The images in a release that correspond to servers.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_sol_policy",
			Category:         "sol",
			ShortDescription: `Policy for configuring Serial Over LAN settings on endpoint.`,
			Description: `
Policy for configuring Serial Over LAN settings on endpoint.
`,
			Keywords: []string{
				"sol",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "9600",
					Description: `Use baud rate 9600 for communication.`,
				},
				resource.Attribute{
					Name:        "19200",
					Description: `Use baud rate 19200 for communication.`,
				},
				resource.Attribute{
					Name:        "38400",
					Description: `Use baud rate 38400 for communication.`,
				},
				resource.Attribute{
					Name:        "57600",
					Description: `Use baud rate 57600 for communication.`,
				},
				resource.Attribute{
					Name:        "115200",
					Description: `Use baud rate 115200 for communication.`,
				},
				resource.Attribute{
					Name:        "com0",
					Description: `Use serial port com0 for communication.`,
				},
				resource.Attribute{
					Name:        "com1",
					Description: `Use serial port com1 for communication.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ssh_policy",
			Category:         "ssh",
			ShortDescription: `Secure shell policy on the endpoint.`,
			Description: `
Secure shell policy on the endpoint.
`,
			Keywords: []string{
				"ssh",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_controller",
			Category:         "storage",
			ShortDescription: `Storage Controller present in a server.`,
			Description: `
Storage Controller present in a server.
`,
			Keywords: []string{
				"storage",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_disk_group",
			Category:         "storage",
			ShortDescription: `Group of one or more Spans to configure virtual drive.`,
			Description: `
Group of one or more Spans to configure virtual drive.
`,
			Keywords: []string{
				"storage",
				"disk",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_disk_group_policy",
			Category:         "storage",
			ShortDescription: `A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.`,
			Description: `
A reusable RAID disk group configuration that can be applied across multiple servers. Also provides options to move JBOD disks in the disk group to Unconfigured Good state before they are used in the disk group.
`,
			Keywords: []string{
				"storage",
				"disk",
				"group",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Raid0",
					Description: `RAID 0 Stripe Raid Level.`,
				},
				resource.Attribute{
					Name:        "Raid1",
					Description: `RAID 1 Mirror Raid Level.`,
				},
				resource.Attribute{
					Name:        "Raid5",
					Description: `RAID 5 Mirror Raid Level.`,
				},
				resource.Attribute{
					Name:        "Raid6",
					Description: `RAID 6 Mirror Raid Level.`,
				},
				resource.Attribute{
					Name:        "Raid10",
					Description: `RAID 10 Mirror Raid Level.`,
				},
				resource.Attribute{
					Name:        "Raid50",
					Description: `RAID 50 Mirror Raid Level.`,
				},
				resource.Attribute{
					Name:        "Raid60",
					Description: `RAID 60 Mirror Raid Level.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_enclosure",
			Category:         "storage",
			ShortDescription: `Storage Enclosure for physical disks.`,
			Description: `
Storage Enclosure for physical disks.
`,
			Keywords: []string{
				"storage",
				"enclosure",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_enclosure_disk",
			Category:         "storage",
			ShortDescription: `Physical Disk on the enclosure.`,
			Description: `
Physical Disk on the enclosure.
`,
			Keywords: []string{
				"storage",
				"enclosure",
				"disk",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_enclosure_disk_slot_ep",
			Category:         "storage",
			ShortDescription: `Physical Disk slots on the enclosure.`,
			Description: `
Physical Disk slots on the enclosure.
`,
			Keywords: []string{
				"storage",
				"enclosure",
				"disk",
				"slot",
				"ep",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_flash_controller",
			Category:         "storage",
			ShortDescription: `Storage Flex Flash Controller.`,
			Description: `
Storage Flex Flash Controller.
`,
			Keywords: []string{
				"storage",
				"flex",
				"flash",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_flash_controller_props",
			Category:         "storage",
			ShortDescription: `Flex flash controller properties.`,
			Description: `
Flex flash controller properties.
`,
			Keywords: []string{
				"storage",
				"flex",
				"flash",
				"controller",
				"props",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_flash_physical_drive",
			Category:         "storage",
			ShortDescription: `Physical Drive repersenting a SD Card.`,
			Description: `
Physical Drive repersenting a SD Card.
`,
			Keywords: []string{
				"storage",
				"flex",
				"flash",
				"physical",
				"drive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_flash_virtual_drive",
			Category:         "storage",
			ShortDescription: `Virtual Drive repersenting a SD Card.`,
			Description: `
Virtual Drive repersenting a SD Card.
`,
			Keywords: []string{
				"storage",
				"flex",
				"flash",
				"virtual",
				"drive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_util_controller",
			Category:         "storage",
			ShortDescription: `Storage Flex Util Adapter.`,
			Description: `
Storage Flex Util Adapter.
`,
			Keywords: []string{
				"storage",
				"flex",
				"util",
				"controller",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_util_physical_drive",
			Category:         "storage",
			ShortDescription: `Storage Flex Util Physical Drive.`,
			Description: `
Storage Flex Util Physical Drive.
`,
			Keywords: []string{
				"storage",
				"flex",
				"util",
				"physical",
				"drive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_flex_util_virtual_drive",
			Category:         "storage",
			ShortDescription: `Storage Flex Util Virtual Drive.`,
			Description: `
Storage Flex Util Virtual Drive.
`,
			Keywords: []string{
				"storage",
				"flex",
				"util",
				"virtual",
				"drive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_array",
			Category:         "storage",
			ShortDescription: `The details of the Hitachi storage array.`,
			Description: `
The details of the Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"array",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_controller",
			Category:         "storage",
			ShortDescription: `A storage controller entity in Hitachi storage array.`,
			Description: `
A storage controller entity in Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component operational state is unknown.`,
				},
				resource.Attribute{
					Name:        "Primary",
					Description: `Component operates in primary mode and accepts workloads.`,
				},
				resource.Attribute{
					Name:        "Secondary",
					Description: `Component is running as a secondary or standby mode.`,
				},
				resource.Attribute{
					Name:        "Maintenance",
					Description: `Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_disk",
			Category:         "storage",
			ShortDescription: `Disk entity associated with Hitachi storage array.`,
			Description: `
Disk entity associated with Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Disk protocol is unknown.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Serial Attached SCSI protocol (SAS) used in disk.`,
				},
				resource.Attribute{
					Name:        "NVMe",
					Description: `Non-volatile memory express (NVMe) protocol used in disk.`,
				},
				resource.Attribute{
					Name:        "SATA",
					Description: `Serial Advanced Technology Attachment (SATA) used in disk.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown disk type.`,
				},
				resource.Attribute{
					Name:        "SSD",
					Description: `Storage disk with Solid state disk.`,
				},
				resource.Attribute{
					Name:        "HDD",
					Description: `Storage disk with Hard disk drive.`,
				},
				resource.Attribute{
					Name:        "NVRAM",
					Description: `Storage disk with Non-volatile random-access memory type.`,
				},
				resource.Attribute{
					Name:        "SATA",
					Description: `Disk drive implementation with Serial Advanced Technology Attachment (SATA).`,
				},
				resource.Attribute{
					Name:        "BSAS",
					Description: `Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Storage disk with Fiber Channel.`,
				},
				resource.Attribute{
					Name:        "FSAS",
					Description: `Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.`,
				},
				resource.Attribute{
					Name:        "LUN",
					Description: `Logical Unit Number refers to a logical disk.`,
				},
				resource.Attribute{
					Name:        "MSATA",
					Description: `SATA disk in multi-disk carrier storage shelf.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Storage disk with serial attached SCSI.`,
				},
				resource.Attribute{
					Name:        "VMDISK",
					Description: `Virtual machine Data Disk.`,
				},
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `SAS.`,
				},
				resource.Attribute{
					Name:        "SSD(MLC)",
					Description: `SSD (MLC).`,
				},
				resource.Attribute{
					Name:        "SSD(FMC)",
					Description: `SSD (FMC).`,
				},
				resource.Attribute{
					Name:        "SSD(FMD)",
					Description: `SSD (FMD).`,
				},
				resource.Attribute{
					Name:        "SSD(SLC)",
					Description: `SSD (SLC).`,
				},
				resource.Attribute{
					Name:        "SSD",
					Description: `SSD.`,
				},
				resource.Attribute{
					Name:        "SSD(RI)",
					Description: `SSD (RI).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_host",
			Category:         "storage",
			ShortDescription: `A host group entity in Hitachi storage array. It is an abstraction used by Hitachi storage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.`,
			Description: `
A host group entity in Hitachi storage array. It is an abstraction used by Hitachi storage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "CHAP",
					Description: `CHAP-authentication mode.`,
				},
				resource.Attribute{
					Name:        "NONE",
					Description: `No-authentication mode.`,
				},
				resource.Attribute{
					Name:        "BOTH",
					Description: `Comply with Host Setting.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Port supports fibre channel protocol.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `Port supports iSCSI protocol.`,
				},
				resource.Attribute{
					Name:        "FCoE",
					Description: `Port supports fibre channel over ethernet protocol.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_host_lun",
			Category:         "storage",
			ShortDescription: `A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.`,
			Description: `
A host LUN entity in Hitachi storage array. It exists only if the volume has a connection to host group. A host lun provides public connection to all hosts associated within host group. Hitachi assign same HLU for all the host.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"host",
				"lun",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_parity_group",
			Category:         "storage",
			ShortDescription: `A parity group in Hitachi storage array.`,
			Description: `
A parity group in Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"parity",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown RAID type.`,
				},
				resource.Attribute{
					Name:        "RAID0",
					Description: `RAID0 splits (\ stripes\ ) data evenly across two or more disks, without parity information.`,
				},
				resource.Attribute{
					Name:        "RAID1",
					Description: `RAID1 requires a minimum of two disks to work, and provides data redundancy and failover.`,
				},
				resource.Attribute{
					Name:        "RAID4",
					Description: `RAID4 stripes block level data and dedicates a disk to parity.`,
				},
				resource.Attribute{
					Name:        "RAID5",
					Description: `RAID5 distributes striping and parity at a block level.`,
				},
				resource.Attribute{
					Name:        "RAID6",
					Description: `RAID6 level operates like RAID5 with distributed parity and striping.`,
				},
				resource.Attribute{
					Name:        "RAID10",
					Description: `RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy.`,
				},
				resource.Attribute{
					Name:        "RAIDDP",
					Description: `RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group.`,
				},
				resource.Attribute{
					Name:        "RAIDTEC",
					Description: `With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_pool",
			Category:         "storage",
			ShortDescription: `A pool entity in Hitachi storage array.`,
			Description: `
A pool entity in Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "Period mode",
					Description: `Period mode.`,
				},
				resource.Attribute{
					Name:        "Continuous mode",
					Description: `Continuous mode.`,
				},
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `The mode in which the monitor is started or stopped at the specified time, and the Tier range is specified by automatic calculation of the DKC (specified by using Storage Navigator).`,
				},
				resource.Attribute{
					Name:        "Manual",
					Description: `The mode in which the monitor is started or stopped by instructions from the REST API server, and the Tier range is specified by automatic calculation of the DKC.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_port",
			Category:         "storage",
			ShortDescription: `Port entity in Hitachi storage array.`,
			Description: `
Port entity in Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Port supports fibre channel protocol.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `Port supports iSCSI protocol.`,
				},
				resource.Attribute{
					Name:        "FCoE",
					Description: `Port supports fibre channel over ethernet protocol.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hitachi_volume",
			Category:         "storage",
			ShortDescription: `A volume entity in Hitachi storage array.`,
			Description: `
A volume entity in Hitachi storage array.
`,
			Keywords: []string{
				"storage",
				"hitachi",
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "Compression",
					Description: `The capacity saving function (compression) is enabled.`,
				},
				resource.Attribute{
					Name:        "Compression Deduplication",
					Description: `The capacity saving function (compression and deduplication) is enabled.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `The capacity saving function (compression and deduplication) is disabled.`,
				},
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `The capacity saving function is enabled.`,
				},
				resource.Attribute{
					Name:        "Disabled",
					Description: `The capacity saving function is disabled.`,
				},
				resource.Attribute{
					Name:        "Enabling",
					Description: `The capacity saving function is being enabled.`,
				},
				resource.Attribute{
					Name:        "Rehydrating",
					Description: `The capacity saving function is being disabled.`,
				},
				resource.Attribute{
					Name:        "Deleting",
					Description: `The volumes for which the capacity saving function is enabled are being deleted.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `An attempt to enable the capacity saving function failed.`,
				},
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "NOT DEFINED",
					Description: `The volume is not implemented.`,
				},
				resource.Attribute{
					Name:        "DEFINING",
					Description: `The volume is being created.`,
				},
				resource.Attribute{
					Name:        "REMOVING",
					Description: `The volume is being removed.`,
				},
				resource.Attribute{
					Name:        "OPEN-V",
					Description: `To be provided by Hitachi.`,
				},
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "RAID1",
					Description: `RAID1.`,
				},
				resource.Attribute{
					Name:        "RAID5",
					Description: `RAID5.`,
				},
				resource.Attribute{
					Name:        "RAID6",
					Description: `RAID6.`,
				},
				resource.Attribute{
					Name:        "N/A",
					Description: `Not available.`,
				},
				resource.Attribute{
					Name:        "NML",
					Description: `The volume is in normal status.`,
				},
				resource.Attribute{
					Name:        "BLK",
					Description: `The volume is blocked.`,
				},
				resource.Attribute{
					Name:        "BSY",
					Description: `The volume status is being changed.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The volume status is unknown (not supported).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hyper_flex_storage_container",
			Category:         "storage",
			ShortDescription: `A Storage Container (Datastore) entity.`,
			Description: `
A Storage Container (Datastore) entity.
`,
			Keywords: []string{
				"storage",
				"hyper",
				"flex",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NFS",
					Description: `Storage container created/accesed through NFS protocol.`,
				},
				resource.Attribute{
					Name:        "SMB",
					Description: `Storage container created/accessed through SMB protocol.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `Storage container created/accessed through iSCSI protocol.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_hyper_flex_volume",
			Category:         "storage",
			ShortDescription: `A HyperFlex Volume entity.`,
			Description: `
A HyperFlex Volume entity.
`,
			Keywords: []string{
				"storage",
				"hyper",
				"flex",
				"volume",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_item",
			Category:         "storage",
			ShortDescription: `FI Local Storage information.`,
			Description: `
FI Local Storage information.
`,
			Keywords: []string{
				"storage",
				"item",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_aggregate",
			Category:         "storage",
			ShortDescription: `NetApp aggregate is a collection of disks arranged into one or more RAID groups.`,
			Description: `
NetApp aggregate is a collection of disks arranged into one or more RAID groups.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"aggregate",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "HDD",
					Description: `Hard Disk Drive.`,
				},
				resource.Attribute{
					Name:        "Hybrid",
					Description: `Solid State Hard Disk Drive.`,
				},
				resource.Attribute{
					Name:        "Hybrid (Flash Pool)",
					Description: `SSHD in a flash pool.`,
				},
				resource.Attribute{
					Name:        "SSD",
					Description: `Solid State Disk.`,
				},
				resource.Attribute{
					Name:        "SSD (FabricPool)",
					Description: `SSD in a flash pool.`,
				},
				resource.Attribute{
					Name:        "VMDisk (SDS)",
					Description: `Storage disk with Hard disk drive.`,
				},
				resource.Attribute{
					Name:        "VMDisk (FabricPool)",
					Description: `Storage disk with Non-volatile random-access memory drives.`,
				},
				resource.Attribute{
					Name:        "LUN (FlexArray)",
					Description: `LUN as a disk.`,
				},
				resource.Attribute{
					Name:        "Not Mapped",
					Description: `Storage disk is not mapped.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown RAID type.`,
				},
				resource.Attribute{
					Name:        "RAID0",
					Description: `RAID0 splits (\ stripes\ ) data evenly across two or more disks, without parity information.`,
				},
				resource.Attribute{
					Name:        "RAID1",
					Description: `RAID1 requires a minimum of two disks to work, and provides data redundancy and failover.`,
				},
				resource.Attribute{
					Name:        "RAID4",
					Description: `RAID4 stripes block level data and dedicates a disk to parity.`,
				},
				resource.Attribute{
					Name:        "RAID5",
					Description: `RAID5 distributes striping and parity at a block level.`,
				},
				resource.Attribute{
					Name:        "RAID6",
					Description: `RAID6 level operates like RAID5 with distributed parity and striping.`,
				},
				resource.Attribute{
					Name:        "RAID10",
					Description: `RAID10 requires a minimum of four disks in the array. It stripes across disks for higher performance, and mirrors for redundancy.`,
				},
				resource.Attribute{
					Name:        "RAIDDP",
					Description: `RAIDDP uses up to two spare disks to replace and reconstruct the data from up to two simultaneously failed disks within the RAID group.`,
				},
				resource.Attribute{
					Name:        "RAIDTEC",
					Description: `With RAIDTEC protection, use up to three spare disks to replace and reconstruct the data from up to three simultaneously failed disks within the RAID group.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Specifies that the aggregate is discovered, but the aggregate information is not yet retrieved by the Unified Manager server.`,
				},
				resource.Attribute{
					Name:        "Online",
					Description: `Aggregate is ready and available.`,
				},
				resource.Attribute{
					Name:        "Onlining",
					Description: `Transitioning online.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Aggregate is unavailable.`,
				},
				resource.Attribute{
					Name:        "Offlining",
					Description: `Transitioning offline.`,
				},
				resource.Attribute{
					Name:        "Relocating",
					Description: `The aggregate is being relocated.`,
				},
				resource.Attribute{
					Name:        "Restricted",
					Description: `Limited operations (e.g., parity reconstruction) are allowed, but data access is not allowed.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The aggregate cannot be brought online.`,
				},
				resource.Attribute{
					Name:        "Inconsistent",
					Description: `The aggregate has been marked corrupted; contact technical support.`,
				},
				resource.Attribute{
					Name:        "Unmounted",
					Description: `The aggregate is not mounted.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_base_disk",
			Category:         "storage",
			ShortDescription: `NetApp base disk is a storage array disk.`,
			Description: `
NetApp base disk is a storage array disk.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"base",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Container is currently unknown. This is the default setting.`,
				},
				resource.Attribute{
					Name:        "Aggregate",
					Description: `Disk is used as a physical disk in an aggregate.`,
				},
				resource.Attribute{
					Name:        "Broken",
					Description: `Disk is in broken pool.`,
				},
				resource.Attribute{
					Name:        "Label Maintenance",
					Description: `Disk is in online label maintenance list.`,
				},
				resource.Attribute{
					Name:        "Foreign",
					Description: `Array LUN has been marked foreign.`,
				},
				resource.Attribute{
					Name:        "Maintenance",
					Description: `Disk is in maintenance center.`,
				},
				resource.Attribute{
					Name:        "Mediator",
					Description: `A mediator disk is a disk used on non-shared HA systems hosted by an external node which is used to communicate the viability of the storage failover between non-shared HA nodes.`,
				},
				resource.Attribute{
					Name:        "Shared",
					Description: `Disk is partitioned or in a storage pool.`,
				},
				resource.Attribute{
					Name:        "Remote",
					Description: `Disk belongs to a remote cluster.`,
				},
				resource.Attribute{
					Name:        "Spare",
					Description: `Disk is a spare disk.`,
				},
				resource.Attribute{
					Name:        "Unassigned",
					Description: `Disk ownership has not been assigned.`,
				},
				resource.Attribute{
					Name:        "Unsupported",
					Description: `Disk is not supported.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown disk type.`,
				},
				resource.Attribute{
					Name:        "SSDNVM",
					Description: `Solid state disk with Non-Volatile Memory Express protocol enabled.`,
				},
				resource.Attribute{
					Name:        "ATA",
					Description: `Advanced Technology Attachment is a type of disk drive that integrates the drive controller directly on the drive itself.`,
				},
				resource.Attribute{
					Name:        "FCAL",
					Description: `For the FC-AL disk connection type, disk shelves are connected to the controller in a loop`,
				},
				resource.Attribute{
					Name:        "BSAS",
					Description: `Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.`,
				},
				resource.Attribute{
					Name:        "FSAS",
					Description: `Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, and rotational speed of traditional enterprise-class SATA drives with the fully capable SAS interface typical for classic SAS drives.`,
				},
				resource.Attribute{
					Name:        "LUN",
					Description: `Logical Unit Number refers to a logical disk.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Storage disk with serial attached SCSI.`,
				},
				resource.Attribute{
					Name:        "MSATA",
					Description: `SATA disk in multi-disk carrier storage shelf.`,
				},
				resource.Attribute{
					Name:        "SSD",
					Description: `Storage disk with Solid state disk.`,
				},
				resource.Attribute{
					Name:        "VMDISK",
					Description: `Virtual machine Data Disk.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Disk protocol is unknown.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Serial Attached SCSI protocol (SAS) used in disk.`,
				},
				resource.Attribute{
					Name:        "NVMe",
					Description: `Non-volatile memory express (NVMe) protocol used in disk.`,
				},
				resource.Attribute{
					Name:        "SATA",
					Description: `Serial Advanced Technology Attachment (SATA) used in disk.`,
				},
				resource.Attribute{
					Name:        "Present",
					Description: `Storage disk state type is present.`,
				},
				resource.Attribute{
					Name:        "Copy",
					Description: `Storage disk state type is copy.`,
				},
				resource.Attribute{
					Name:        "Broken",
					Description: `Storage disk state type is broken.`,
				},
				resource.Attribute{
					Name:        "Maintenance",
					Description: `Storage disk state type is maintenance.`,
				},
				resource.Attribute{
					Name:        "Partner",
					Description: `Storage disk state type is partner.`,
				},
				resource.Attribute{
					Name:        "Pending",
					Description: `Storage disk state type is pending.`,
				},
				resource.Attribute{
					Name:        "Reconstructing",
					Description: `Storage disk state type is reconstructing.`,
				},
				resource.Attribute{
					Name:        "Removed",
					Description: `Storage disk state type is removed.`,
				},
				resource.Attribute{
					Name:        "Spare",
					Description: `Storage disk state type is spare.`,
				},
				resource.Attribute{
					Name:        "Unfail",
					Description: `Storage disk state type is unfail.`,
				},
				resource.Attribute{
					Name:        "Zeroing",
					Description: `Storage disk state type is zeroing.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown disk type.`,
				},
				resource.Attribute{
					Name:        "SSD",
					Description: `Storage disk with Solid state disk.`,
				},
				resource.Attribute{
					Name:        "HDD",
					Description: `Storage disk with Hard disk drive.`,
				},
				resource.Attribute{
					Name:        "NVRAM",
					Description: `Storage disk with Non-volatile random-access memory type.`,
				},
				resource.Attribute{
					Name:        "SATA",
					Description: `Disk drive implementation with Serial Advanced Technology Attachment (SATA).`,
				},
				resource.Attribute{
					Name:        "BSAS",
					Description: `Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Storage disk with Fiber Channel.`,
				},
				resource.Attribute{
					Name:        "FSAS",
					Description: `Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.`,
				},
				resource.Attribute{
					Name:        "LUN",
					Description: `Logical Unit Number refers to a logical disk.`,
				},
				resource.Attribute{
					Name:        "MSATA",
					Description: `SATA disk in multi-disk carrier storage shelf.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Storage disk with serial attached SCSI.`,
				},
				resource.Attribute{
					Name:        "VMDISK",
					Description: `Virtual machine Data Disk.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_cluster",
			Category:         "storage",
			ShortDescription: `NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.`,
			Description: `
NetApp cluster consists of one or more nodes grouped together as HA pairs to form a scalable cluster.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"cluster",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_ethernet_port",
			Category:         "storage",
			ShortDescription: `Ethernet port is a port on a node in a storage array.`,
			Description: `
Ethernet port is a port on a node in a storage array.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"ethernet",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "down",
					Description: `An inactive port is listed as Down.`,
				},
				resource.Attribute{
					Name:        "up",
					Description: `An active port is listed as Up.`,
				},
				resource.Attribute{
					Name:        "present",
					Description: `An active port is listed as present.`,
				},
				resource.Attribute{
					Name:        "LAG",
					Description: `Storage port of type lag.`,
				},
				resource.Attribute{
					Name:        "physical",
					Description: `LIFs can be configured directly on physical ports.`,
				},
				resource.Attribute{
					Name:        "VLAN",
					Description: `A logical port that receives and sends VLAN-tagged (IEEE 802.1Q standard) traffic. VLAN port characteristics include the VLAN ID for the port.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_export_policy",
			Category:         "storage",
			ShortDescription: `NetApp export policies enable client access control to volumes. Clients that match specific IP addresses and/or specific authentication types are granted access.`,
			Description: `
NetApp export policies enable client access control to volumes. Clients that match specific IP addresses and/or specific authentication types are granted access.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"export",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_fc_interface",
			Category:         "storage",
			ShortDescription: `NetApp FC Interface is a logical interface.`,
			Description: `
NetApp FC Interface is a logical interface.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"fc",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "down",
					Description: `An inactive port is listed as Down.`,
				},
				resource.Attribute{
					Name:        "up",
					Description: `An active port is listed as Up.`,
				},
				resource.Attribute{
					Name:        "present",
					Description: `An active port is listed as present.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Port supports fibre channel protocol.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `Port supports iSCSI protocol.`,
				},
				resource.Attribute{
					Name:        "FCoE",
					Description: `Port supports fibre channel over ethernet protocol.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_fc_port",
			Category:         "storage",
			ShortDescription: `Fibre Channel (FC) port is a port on a node in a storage array.`,
			Description: `
Fibre Channel (FC) port is a port on a node in a storage array.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"fc",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown port state.`,
				},
				resource.Attribute{
					Name:        "StartUp",
					Description: `The port in the storage array is booting up.`,
				},
				resource.Attribute{
					Name:        "LinkNotConnected",
					Description: `The port has finished initialization, but a link with the fabric is not established.`,
				},
				resource.Attribute{
					Name:        "Online",
					Description: `The port is initialized and a link with the fabric has been established.`,
				},
				resource.Attribute{
					Name:        "LinkDisconnected",
					Description: `The link on this port is currently not established.`,
				},
				resource.Attribute{
					Name:        "OfflineUser",
					Description: `The port is administratively disabled.`,
				},
				resource.Attribute{
					Name:        "OfflineSystem",
					Description: `The port is set to offline by the system. This happens when the port encounters too many errors.`,
				},
				resource.Attribute{
					Name:        "NodeOffline",
					Description: `The state information for the port cannot be retrieved. The node is offline or inaccessible.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Port supports fibre channel protocol.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `Port supports iSCSI protocol.`,
				},
				resource.Attribute{
					Name:        "FCoE",
					Description: `Port supports fibre channel over ethernet protocol.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_initiator_group",
			Category:         "storage",
			ShortDescription: `NetApp Initiator Group specifies host access to LUNs on the storage system.`,
			Description: `
NetApp Initiator Group specifies host access to LUNs on the storage system.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"initiator",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "FCP",
					Description: `Fibre channel initiator type which contains WWN of an HBA on the host.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `An iSCSI initiator type used by the host.`,
				},
				resource.Attribute{
					Name:        "mixed",
					Description: `For systems using both FC and iSCSI connections to the same LUN, create two igroups, one for FC and one for iSCSI. Then map the LUN to both igroups.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_ip_interface",
			Category:         "storage",
			ShortDescription: `NetApp IP interface is a logical interface.`,
			Description: `
NetApp IP interface is a logical interface.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"ip",
				"interface",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IPv4",
					Description: `IPv4 Address type.`,
				},
				resource.Attribute{
					Name:        "IPv6",
					Description: `IPv6 Address type.`,
				},
				resource.Attribute{
					Name:        "down",
					Description: `An inactive port is listed as Down.`,
				},
				resource.Attribute{
					Name:        "up",
					Description: `An active port is listed as Up.`,
				},
				resource.Attribute{
					Name:        "present",
					Description: `An active port is listed as present.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_license",
			Category:         "storage",
			ShortDescription: `NetApp licenses for NetApp Ontap.`,
			Description: `
NetApp licenses for NetApp Ontap.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"license",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_lun",
			Category:         "storage",
			ShortDescription: `NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.`,
			Description: `
NetApp LUN (logical unit number) is an identifier for a device called a logical unit addressed by a SAN protocol.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"lun",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Linux",
					Description: `Family of open source Unix-like operating systems based on the Linux kernel.`,
				},
				resource.Attribute{
					Name:        "AIX",
					Description: `Advanced Interactive Executive (AIX).`,
				},
				resource.Attribute{
					Name:        "HP-UX",
					Description: `HP-UX is implementation of the Unix operating system, based on Unix System V.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `Windows Server 2008 or Windows Server 2012 Hyper-V.`,
				},
				resource.Attribute{
					Name:        "OpenVMS",
					Description: `OpenVMS is multi-user, multiprocessing virtual memory-based operating system.`,
				},
				resource.Attribute{
					Name:        "Solaris",
					Description: `Solaris is a Unix operating system.`,
				},
				resource.Attribute{
					Name:        "NetWare",
					Description: `NetWare is a computer network operating system.`,
				},
				resource.Attribute{
					Name:        "VMware",
					Description: `An enterprise-class, type-1 hypervisor developed by VMware for deploying and serving virtual computers.`,
				},
				resource.Attribute{
					Name:        "Windows",
					Description: `Single-partition Windows disk using the Master Boot Record (MBR) partitioning style.`,
				},
				resource.Attribute{
					Name:        "Xen",
					Description: `Xen is a type-1 hypervisor, providing services that allow multiple computer operating systems to execute on the same computer hardware concurrently.`,
				},
				resource.Attribute{
					Name:        "offline",
					Description: `The LUN is administratively offline, or a more detailed offline reason is not available.`,
				},
				resource.Attribute{
					Name:        "online",
					Description: `The LUN is online.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_lun_map",
			Category:         "storage",
			ShortDescription: `NetApp LUN mapping is the process of associating a LUN with an igroup. When a LUN is mapped to an igroup, initiators in the igroup are granted. access to the LUN.`,
			Description: `
NetApp LUN mapping is the process of associating a LUN with an igroup. When a LUN is mapped to an igroup, initiators in the igroup are granted. access to the LUN.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"lun",
				"map",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_node",
			Category:         "storage",
			ShortDescription: `NetApp node is a controller in a NetApp cluster. Services and components are controlled and managed by the NetApp node.`,
			Description: `
NetApp node is a controller in a NetApp cluster. Services and components are controlled and managed by the NetApp node.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"node",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component operational state is unknown.`,
				},
				resource.Attribute{
					Name:        "Primary",
					Description: `Component operates in primary mode and accepts workloads.`,
				},
				resource.Attribute{
					Name:        "Secondary",
					Description: `Component is running as a secondary or standby mode.`,
				},
				resource.Attribute{
					Name:        "Maintenance",
					Description: `Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_storage_vm",
			Category:         "storage",
			ShortDescription: `NetApp Storage Virtual Machines contain data volumes and one or more Logical Interfaces ( LIFs ) through which they serve data to the clients.`,
			Description: `
NetApp Storage Virtual Machines contain data volumes and one or more Logical Interfaces ( LIFs ) through which they serve data to the clients.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"vm",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component state is not available.`,
				},
				resource.Attribute{
					Name:        "Starting",
					Description: `Component is being started.`,
				},
				resource.Attribute{
					Name:        "Running",
					Description: `Component is currently running.`,
				},
				resource.Attribute{
					Name:        "Stopping",
					Description: `Component is being stopped.`,
				},
				resource.Attribute{
					Name:        "Stopped",
					Description: `Component has been stopped.`,
				},
				resource.Attribute{
					Name:        "Deleting",
					Description: `Component deletion is in progress.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_volume",
			Category:         "storage",
			ShortDescription: `NetApp volume are data containers that enable you to partition and manage your data.`,
			Description: `
NetApp volume are data containers that enable you to partition and manage your data.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"volume",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "off",
					Description: `The volume will not grow or shrink in size in response to the amount of used space.`,
				},
				resource.Attribute{
					Name:        "grow",
					Description: `The volume will automatically grow when used space in the volume is above the grow threshold.`,
				},
				resource.Attribute{
					Name:        "grow_shrink",
					Description: `The volume will grow or shrink in size in response to the amount of used space.`,
				},
				resource.Attribute{
					Name:        "offline",
					Description: `Read and write access to the volume is not allowed.`,
				},
				resource.Attribute{
					Name:        "online",
					Description: `Read and write access to the volume is allowed.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Storage volume state of error type.`,
				},
				resource.Attribute{
					Name:        "mixed",
					Description: `The constituents of a FlexGroup volume are not all in the same state.`,
				},
				resource.Attribute{
					Name:        "data-protection",
					Description: `Prevents modification of the data on the Volume.`,
				},
				resource.Attribute{
					Name:        "read-write",
					Description: `Data on the Volume can be modified.`,
				},
				resource.Attribute{
					Name:        "load-sharing",
					Description: `Load Sharing.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_net_app_volume_snapshot",
			Category:         "storage",
			ShortDescription: `NetApp Volume Snapshot is a read-only image of a traditional or FlexVol volume, or an aggregate, that captures the state of the file system at a point in time.`,
			Description: `
NetApp Volume Snapshot is a read-only image of a traditional or FlexVol volume, or an aggregate, that captures the state of the file system at a point in time.
`,
			Keywords: []string{
				"storage",
				"net",
				"app",
				"volume",
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_physical_disk",
			Category:         "storage",
			ShortDescription: `Physical Disk on a server.`,
			Description: `
Physical Disk on a server.
`,
			Keywords: []string{
				"storage",
				"physical",
				"disk",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_physical_disk_extension",
			Category:         "storage",
			ShortDescription: `Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.`,
			Description: `
Information of disks as reported by controller. In certain cases like S-series servers, disk information will be reported by controller separately and this represents such information.
`,
			Keywords: []string{
				"storage",
				"physical",
				"disk",
				"extension",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_physical_disk_usage",
			Category:         "storage",
			ShortDescription: `Has usage map between physical disks and virtual drives.`,
			Description: `
Has usage map between physical disks and virtual drives.
`,
			Keywords: []string{
				"storage",
				"physical",
				"disk",
				"usage",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_array",
			Category:         "storage",
			ShortDescription: `The details of the Pure storage array.`,
			Description: `
The details of the Pure storage array.
`,
			Keywords: []string{
				"storage",
				"pure",
				"array",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_controller",
			Category:         "storage",
			ShortDescription: `A storage controller entity in Pure FlashArray.`,
			Description: `
A storage controller entity in Pure FlashArray.
`,
			Keywords: []string{
				"storage",
				"pure",
				"controller",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component operational state is unknown.`,
				},
				resource.Attribute{
					Name:        "Primary",
					Description: `Component operates in primary mode and accepts workloads.`,
				},
				resource.Attribute{
					Name:        "Secondary",
					Description: `Component is running as a secondary or standby mode.`,
				},
				resource.Attribute{
					Name:        "Maintenance",
					Description: `Component is in maintenance mode for upgrade. During maintenance mode, component does not perform any workload.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_disk",
			Category:         "storage",
			ShortDescription: `Disk entity associated with Pure FlashArray.`,
			Description: `
Disk entity associated with Pure FlashArray.
`,
			Keywords: []string{
				"storage",
				"pure",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Disk protocol is unknown.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Serial Attached SCSI protocol (SAS) used in disk.`,
				},
				resource.Attribute{
					Name:        "NVMe",
					Description: `Non-volatile memory express (NVMe) protocol used in disk.`,
				},
				resource.Attribute{
					Name:        "SATA",
					Description: `Serial Advanced Technology Attachment (SATA) used in disk.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Default unknown disk type.`,
				},
				resource.Attribute{
					Name:        "SSD",
					Description: `Storage disk with Solid state disk.`,
				},
				resource.Attribute{
					Name:        "HDD",
					Description: `Storage disk with Hard disk drive.`,
				},
				resource.Attribute{
					Name:        "NVRAM",
					Description: `Storage disk with Non-volatile random-access memory type.`,
				},
				resource.Attribute{
					Name:        "SATA",
					Description: `Disk drive implementation with Serial Advanced Technology Attachment (SATA).`,
				},
				resource.Attribute{
					Name:        "BSAS",
					Description: `Bridged SAS-SATA disks with added hardware to enable them to be plugged into a SAS-connected storage shelf.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Storage disk with Fiber Channel.`,
				},
				resource.Attribute{
					Name:        "FSAS",
					Description: `Near Line SAS. NL-SAS drives are enterprise SATA drives with a SAS interface, head, media, androtational speed of traditional enterprise-class SATA drives with the fully capable SAS interfacetypical for classic SAS drives.`,
				},
				resource.Attribute{
					Name:        "LUN",
					Description: `Logical Unit Number refers to a logical disk.`,
				},
				resource.Attribute{
					Name:        "MSATA",
					Description: `SATA disk in multi-disk carrier storage shelf.`,
				},
				resource.Attribute{
					Name:        "SAS",
					Description: `Storage disk with serial attached SCSI.`,
				},
				resource.Attribute{
					Name:        "VMDISK",
					Description: `Virtual machine Data Disk.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_host",
			Category:         "storage",
			ShortDescription: `A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.`,
			Description: `
A host entity in PureStorage FlashArray. It is an abstraction used by PureStorage to organize the storage network addresses (Fibre Channel worldwide names or iSCSI qualified names) of client computers and to control communications between clients and volumes.
`,
			Keywords: []string{
				"storage",
				"pure",
				"host",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_host_group",
			Category:         "storage",
			ShortDescription: `A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.`,
			Description: `
A host group represents a collection of hosts with common connectivity to volumes. Once a connection has been established between a host group and a volume, all of the hosts within the host group are able to access the volume through the connection. These connections are called shared connections because the connection is shared between all of the hosts within the host group.
`,
			Keywords: []string{
				"storage",
				"pure",
				"host",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_host_lun",
			Category:         "storage",
			ShortDescription: `A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.`,
			Description: `
A host LUN entity in Pure storage array. It exists only if the volume has a connection to host or host group. For host group mapping, it provides public connection to all hosts associated within host group. A volume can have private connection to host as well. It cannot have public and private connection. Pure assign same HLU for all the host in case if it is connected through host group.
`,
			Keywords: []string{
				"storage",
				"pure",
				"host",
				"lun",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_port",
			Category:         "storage",
			ShortDescription: `Port entity in Pure FlashArray.`,
			Description: `
Port entity in Pure FlashArray.
`,
			Keywords: []string{
				"storage",
				"pure",
				"port",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Component status is not available.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Component is healthy and no issues found.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `Functioning, but not at full capability due to a non-fatal failure.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Not functioning or requiring immediate attention.`,
				},
				resource.Attribute{
					Name:        "Offline",
					Description: `Component is installed, but powered off.`,
				},
				resource.Attribute{
					Name:        "Identifying",
					Description: `Component is in initialization process.`,
				},
				resource.Attribute{
					Name:        "NotAvailable",
					Description: `Component is not installed or not available.`,
				},
				resource.Attribute{
					Name:        "Updating",
					Description: `Software update is in progress.`,
				},
				resource.Attribute{
					Name:        "Unrecognized",
					Description: `Component is not recognized. It may be because the component is not installed properly or it is not supported.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `Port supports fibre channel protocol.`,
				},
				resource.Attribute{
					Name:        "iSCSI",
					Description: `Port supports iSCSI protocol.`,
				},
				resource.Attribute{
					Name:        "FCoE",
					Description: `Port supports fibre channel over ethernet protocol.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_protection_group",
			Category:         "storage",
			ShortDescription: `Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.`,
			Description: `
Protection group entity in Pure storage array. A volume can be protected by associating with protection group either directly or indirectly (either host or host group). Snapshots are created on protected volume in local array or target array or both as per scheduler configuration.
`,
			Keywords: []string{
				"storage",
				"pure",
				"protection",
				"group",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_protection_group_snapshot",
			Category:         "storage",
			ShortDescription: `Protection group snapshot entity in Pure protection group.`,
			Description: `
Protection group snapshot entity in Pure protection group.
`,
			Keywords: []string{
				"storage",
				"pure",
				"protection",
				"group",
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_replication_schedule",
			Category:         "storage",
			ShortDescription: `Pure snapshot replication schedule entity.`,
			Description: `
Pure snapshot replication schedule entity.
`,
			Keywords: []string{
				"storage",
				"pure",
				"replication",
				"schedule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_snapshot_schedule",
			Category:         "storage",
			ShortDescription: `PureStorage FlashArray snapshot schedule configuration entity.`,
			Description: `
PureStorage FlashArray snapshot schedule configuration entity.
`,
			Keywords: []string{
				"storage",
				"pure",
				"snapshot",
				"schedule",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_volume",
			Category:         "storage",
			ShortDescription: `A volume entity in PureStorage FlashArray.`,
			Description: `
A volume entity in PureStorage FlashArray.
`,
			Keywords: []string{
				"storage",
				"pure",
				"volume",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_pure_volume_snapshot",
			Category:         "storage",
			ShortDescription: `Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.`,
			Description: `
Volume snapshot entity in Pure protection group. Volume snapshots are created either on-demand or using scheduler. Snapshots are immutable and it cannot be connected to hosts or host groups, and therefore the data it contains cannot be read or written.
`,
			Keywords: []string{
				"storage",
				"pure",
				"volume",
				"snapshot",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_sas_expander",
			Category:         "storage",
			ShortDescription: `SAS Expander present in a server.`,
			Description: `
SAS Expander present in a server.
`,
			Keywords: []string{
				"storage",
				"sas",
				"expander",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_sas_port",
			Category:         "storage",
			ShortDescription: `Sas Port details of the SAS endpoint.`,
			Description: `
Sas Port details of the SAS endpoint.
`,
			Keywords: []string{
				"storage",
				"sas",
				"port",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_span",
			Category:         "storage",
			ShortDescription: `Group of disks to configure virtual drive.`,
			Description: `
Group of disks to configure virtual drive.
`,
			Keywords: []string{
				"storage",
				"span",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_storage_policy",
			Category:         "storage",
			ShortDescription: `The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.`,
			Description: `
The storage policy models the reusable storage related configuration that can be applied on many servers. This policy allows creation of RAID groups using existing disk group policies and virtual drives on the drive groups. The user has options to move all unused disks to JBOD or Unconfigured good state. The encryption of drives can be enabled through this policy using remote keys from a KMIP server.
`,
			Keywords: []string{
				"storage",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "UnconfiguredGood",
					Description: `Unconfigured good state -ready to be added in a RAID group.`,
				},
				resource.Attribute{
					Name:        "Jbod",
					Description: `JBOD state where the disks start showing up to host os.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_vd_member_ep",
			Category:         "storage",
			ShortDescription: `Reference to LocalDisk to build up a VirtualDrive.`,
			Description: `
Reference to LocalDisk to build up a VirtualDrive.
`,
			Keywords: []string{
				"storage",
				"vd",
				"member",
				"ep",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_virtual_drive",
			Category:         "storage",
			ShortDescription: `A Virtual Disk Drive or Logical Unit Number.`,
			Description: `
A Virtual Disk Drive or Logical Unit Number.
`,
			Keywords: []string{
				"storage",
				"virtual",
				"drive",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_virtual_drive_container",
			Category:         "storage",
			ShortDescription: `A Virtual Disk Drive Container.`,
			Description: `
A Virtual Disk Drive Container.
`,
			Keywords: []string{
				"storage",
				"virtual",
				"drive",
				"container",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_storage_virtual_drive_extension",
			Category:         "storage",
			ShortDescription: `Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.`,
			Description: `
Information of virtual drives as reported by a storage controller. In certain cases like S-series servers, virtual drive information will be reported by the controller separately and this represents such information.
`,
			Keywords: []string{
				"storage",
				"virtual",
				"drive",
				"extension",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_syslog_policy",
			Category:         "syslog",
			ShortDescription: `The syslog policy configure the syslog server to receive CIMC log entries.`,
			Description: `
The syslog policy configure the syslog server to receive CIMC log entries.
`,
			Keywords: []string{
				"syslog",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_tam_advisory_count",
			Category:         "tam",
			ShortDescription: `Total number of advisories currently affecting a given Account.`,
			Description: `
Total number of advisories currently affecting a given Account.
`,
			Keywords: []string{
				"tam",
				"advisory",
				"count",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_tam_advisory_definition",
			Category:         "tam",
			ShortDescription: `An Intersight Advisory. An advisory represents an identification of a potential issue and may also include a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.`,
			Description: `
An Intersight Advisory. An advisory represents an identification of a potential issue and may also include  a recommendation for resolving the said issue. Advisories may be of different kind and severity. for e.g. It could be a security vulnerability or a performance issue or a hardware issue with different recommendations for resolving them.
`,
			Keywords: []string{
				"tam",
				"advisory",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ready",
					Description: `Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.`,
				},
				resource.Attribute{
					Name:        "evaluating",
					Description: `Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation.`,
				},
				resource.Attribute{
					Name:        "securityAdvisory",
					Description: `Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).`,
				},
				resource.Attribute{
					Name:        "fieldNotice",
					Description: `Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html).`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_tam_advisory_info",
			Category:         "tam",
			ShortDescription: `State of an advisory in the context of a given account. Used to capture a given account's preferences regarding associated advisory.`,
			Description: `
State of an advisory in the context of a given account. Used to capture a given account's preferences regarding  associated advisory.
`,
			Keywords: []string{
				"tam",
				"advisory",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "active",
					Description: `Advisory is currently active and the user wants to receive updates for this advisory.`,
				},
				resource.Attribute{
					Name:        "acknowledged",
					Description: `Advisory is seen and acknowledged by the user and she no longer wants to recieve updates.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_tam_advisory_instance",
			Category:         "tam",
			ShortDescription: `Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.`,
			Description: `
Instance of an Intersight advisory applicable for an Intersight managed object. An advisory instance is created when a given advisory is found applicable for an Intersight managed object. An advisory instance is retained for some time even after being cleared for historical purposes. A 'cleared' advisory instance is deleted after the retention time is elaspsed.
`,
			Keywords: []string{
				"tam",
				"advisory",
				"instance",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "unknown",
					Description: `Intersight is unable to determine if the Advisory instance is applicable for the affected managed object.`,
				},
				resource.Attribute{
					Name:        "active",
					Description: `Advisory instance is currently active and applicable for the affected managed object.`,
				},
				resource.Attribute{
					Name:        "cleared",
					Description: `Advisory instance is no longer applicable for the affected managed object.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_tam_security_advisory",
			Category:         "tam",
			ShortDescription: `Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.`,
			Description: `
Intersight representation of a Cisco PSIRT (https://tools.cisco.com/security/center/publicationListing.x) advisory definition. It includes the description of the security advisory and a corresponding reference to the published advisory. It also includes the Intersight data sources needed to evaluate the applicability of this advisory for relevant Intersight managed objects. A PSIRT definition is evaluated against all managed object referenced using the included data sources. Only Cisco TAC and Intersight devops engineers have the ability to create PSIRT definitions in Intersight.
`,
			Keywords: []string{
				"tam",
				"security",
				"advisory",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ready",
					Description: `Advisory has been evaluated. The affected devices would be analyzed and corresponding advisory instances would be created.`,
				},
				resource.Attribute{
					Name:        "evaluating",
					Description: `Advisory is currently under evaluation. The affected devices would be analyzed but no advisory instances wouldbe created. The results of the analysis would be made available to Intersight engineering for evaluation and validation.`,
				},
				resource.Attribute{
					Name:        "interim",
					Description: `The Cisco investigation for the advisory is ongoing. Cisco will issue revisions to the advisory when additional information, including fixed software release data, becomes available.`,
				},
				resource.Attribute{
					Name:        "final",
					Description: `Cisco has completed its evaluation of the vulnerability described in the advisory. There will be no further updates unless there is a material change in the nature of the vulnerability.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_techsupportmanagement_collection_control_policy",
			Category:         "techsupportmanagement",
			ShortDescription: `Policy to control techsupport collection for a specific account.`,
			Description: `
Policy to control techsupport collection for a specific account.
`,
			Keywords: []string{
				"techsupportmanagement",
				"collection",
				"control",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `Service deployment type None.`,
				},
				resource.Attribute{
					Name:        "SaaS",
					Description: `Service deployment type SaaS.`,
				},
				resource.Attribute{
					Name:        "Appliance",
					Description: `Service deployment type Appliance.`,
				},
				resource.Attribute{
					Name:        "Enable",
					Description: `Enable techsupport collection.`,
				},
				resource.Attribute{
					Name:        "Disable",
					Description: `Disable techsupport collection.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_techsupportmanagement_download",
			Category:         "techsupportmanagement",
			ShortDescription: `Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.`,
			Description: `
Download the techsupport file. The response to this API will be the actual techsupport file. This API needs to be invoked using the download link obtained by doing GET operation on TechsupportStatus. The techsupportDownloadUrl property in the TechsupportStatus API response will have the download link. No other link can be used to download the techsupport file.
`,
			Keywords: []string{
				"techsupportmanagement",
				"download",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_techsupportmanagement_tech_support_bundle",
			Category:         "techsupportmanagement",
			ShortDescription: `A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.`,
			Description: `
A request to collect techsupport and upload it to Intersight Storage Service. The serial number, PID and/or relationship to the target resource provided by the user is used to determine the device type for techsupport collection. If the serial number, PID and target resource are specified in the request, the values must match. Valid values of device types are network.Element for fabric interconnect, compute.Blade for blade server, compute.RackUnit for rack server, equipment.Chassis for chassis, equipment.IoCard for IO Module, equipment.FEX for fabric extender and adapter.Unit for network adapter. UCSM techsupport is collected for device type network.Element. Chassis techsupport is collected for compute.Blade, equipment.Chassis, equipment.IoCard, and blade adapter.Unit. Server techsupport is collected for compute.RackUnit and rack adapter.Unit. Fabric extender techsupport is collected for device type equipment.FEX. Hyper Flex node level techsupport is collected when the request specifies the platform type (HX) and the device type is Hyperflex.Node.
`,
			Keywords: []string{
				"techsupportmanagement",
				"tech",
				"support",
				"bundle",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_techsupportmanagement_tech_support_status",
			Category:         "techsupportmanagement",
			ShortDescription: `The techsupport collection status.`,
			Description: `
The techsupport collection status.
`,
			Keywords: []string{
				"techsupportmanagement",
				"tech",
				"support",
				"status",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_terminal_audit_log",
			Category:         "terminal",
			ShortDescription: `Audit log of remote terminal user sessions.`,
			Description: `
Audit log of remote terminal user sessions.
`,
			Keywords: []string{
				"terminal",
				"audit",
				"log",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_top_system",
			Category:         "top",
			ShortDescription: `Root container for all UCSM / CIMC MOs.`,
			Description: `
Root container for all UCSM / CIMC MOs.
`,
			Keywords: []string{
				"top",
				"system",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_ucsd_backup_info",
			Category:         "ucsd",
			ShortDescription: `List of backup images available for target end device for restore operation.`,
			Description: `
List of backup images available for target end device for restore operation.
`,
			Keywords: []string{
				"ucsd",
				"backup",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_uuidpool_block",
			Category:         "uuidpool",
			ShortDescription: `A block of contiguous UUID addresses that are part of a pool.`,
			Description: `
A block of contiguous UUID addresses that are part of a pool.
`,
			Keywords: []string{
				"uuidpool",
				"block",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_uuidpool_pool",
			Category:         "uuidpool",
			ShortDescription: `Pool represents a collection of UUID items that can be allocated to server profiles.`,
			Description: `
Pool represents a collection of UUID items that can be allocated to server profiles.
`,
			Keywords: []string{
				"uuidpool",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "sequential",
					Description: `Identifiers are assigned in a sequential order.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `Assignment order is decided by the system.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_uuidpool_pool_member",
			Category:         "uuidpool",
			ShortDescription: `PoolMember represents a single UUID that is part of a pool.`,
			Description: `
PoolMember represents a single UUID that is part of a pool.
`,
			Keywords: []string{
				"uuidpool",
				"pool",
				"member",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_uuidpool_universe",
			Category:         "uuidpool",
			ShortDescription: `Universe represents a book keeping container to keep track of all UUIDs for a given account.`,
			Description: `
Universe represents a book keeping container to keep track of all UUIDs for a given account.
`,
			Keywords: []string{
				"uuidpool",
				"universe",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_uuidpool_uuid_lease",
			Category:         "uuidpool",
			ShortDescription: `UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.`,
			Description: `
UuidLease represents a single UUID that is part of the universe, allocated either from a pool or through static assignment.
`,
			Keywords: []string{
				"uuidpool",
				"uuid",
				"lease",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_host",
			Category:         "virtualization",
			ShortDescription: `Depicts operations to control the life cycle of a Hypervisor Host.`,
			Description: `
Depicts operations to control the life cycle of a Hypervisor Host.
`,
			Keywords: []string{
				"virtualization",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `A place holder for the default value.`,
				},
				resource.Attribute{
					Name:        "PowerOffStorageController",
					Description: `Power off HyperFlex storage controller node running on selected hypervisor host.`,
				},
				resource.Attribute{
					Name:        "PowerOnStorageController",
					Description: `Power on HyperFlex storage controller node running on selected hypervisor host.`,
				},
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_virtual_disk",
			Category:         "virtualization",
			ShortDescription: `Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.`,
			Description: `
Depicts disk configuration used to be create a virtual disk on a hypervisor datastore.
`,
			Keywords: []string{
				"virtualization",
				"virtual",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Block",
					Description: `It is a Block virtual disk.`,
				},
				resource.Attribute{
					Name:        "Filesystem",
					Description: `It is a File system virtual disk.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_virtual_machine",
			Category:         "virtualization",
			ShortDescription: `Depicts operations to control the life cycle of a virtual machine on a hypervisor.`,
			Description: `
Depicts operations to control the life cycle of a virtual machine on a hypervisor.
`,
			Keywords: []string{
				"virtualization",
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `A place holder for the default value.`,
				},
				resource.Attribute{
					Name:        "PowerState",
					Description: `Power action is performed on the virtual machine.`,
				},
				resource.Attribute{
					Name:        "Migrate",
					Description: `The virtual machine will be migrated from existing node to a different node in cluster. The behavior depends on the underlying hypervisor.`,
				},
				resource.Attribute{
					Name:        "Create",
					Description: `The virtual machine will be created on the specified hypervisor. This action is also useful if the virtual machine creation failed during first POST operation on VirtualMachine managed object. User can set this action to retry the virtual machine creation.`,
				},
				resource.Attribute{
					Name:        "Delete",
					Description: `The virtual machine will be deleted from the specified hypervisor. User can either set this action or can do a DELETE operation on the VirtualMachine managed object.`,
				},
				resource.Attribute{
					Name:        "linux",
					Description: `A Linux operating system.`,
				},
				resource.Attribute{
					Name:        "windows",
					Description: `A Windows operating system.`,
				},
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "PowerOff",
					Description: `The virtual machine will be powered off if it is already not in powered off state. If it is already powered off, no side-effects are expected.`,
				},
				resource.Attribute{
					Name:        "PowerOn",
					Description: `The virtual machine will be powered on if it is already not in powered on state. If it is already powered on, no side-effects are expected.`,
				},
				resource.Attribute{
					Name:        "Suspend",
					Description: `The virtual machine will be put into a suspended state.`,
				},
				resource.Attribute{
					Name:        "ShutDownGuestOS",
					Description: `The guest operating system is shut down gracefully.`,
				},
				resource.Attribute{
					Name:        "RestartGuestOS",
					Description: `It can either act as a reset switch and abruptly reset the guest operating system, or it can send a restart signal to the guest operating system so that it shuts down gracefully and restarts.`,
				},
				resource.Attribute{
					Name:        "Reset",
					Description: `Resets the virtual machine abruptly, with no consideration for work in progress.`,
				},
				resource.Attribute{
					Name:        "Restart",
					Description: `The virtual machine will be restarted only if it is in powered on state. If it is powered off, it will not be started up.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Power state of the entity is unknown.`,
				},
				resource.Attribute{
					Name:        "OVA",
					Description: `Deploy virtual machine using OVA/F file.`,
				},
				resource.Attribute{
					Name:        "Template",
					Description: `Provision virtual machine using a template file.`,
				},
				resource.Attribute{
					Name:        "Discovered",
					Description: `A virtual machine was 'discovered' and not created from Intersight. No provisioning information is available.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_vmware_cluster",
			Category:         "virtualization",
			ShortDescription: `A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.`,
			Description: `
A real cluster of resources within a data center in VMware. A cluster is a convenient grouping of resources such as Host, Datastore, etc.
`,
			Keywords: []string{
				"virtualization",
				"vmware",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_vmware_datacenter",
			Category:         "virtualization",
			ShortDescription: `Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.`,
			Description: `
Datacenter object in VMware inventory. It is the logical container for all other objects like Datastore, Host, VirtualMachine, etc.
`,
			Keywords: []string{
				"virtualization",
				"vmware",
				"datacenter",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_vmware_datastore",
			Category:         "virtualization",
			ShortDescription: `The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.`,
			Description: `
The VMware Datastore entity with its attributes. Each Datastore belongs to a Datacenter and maybe attached to VMs.
`,
			Keywords: []string{
				"virtualization",
				"vmware",
				"datastore",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The nature of the file system is unknown.`,
				},
				resource.Attribute{
					Name:        "VMFS",
					Description: `It is a Virtual Machine Filesystem.`,
				},
				resource.Attribute{
					Name:        "NFS",
					Description: `It is a Network File System.`,
				},
				resource.Attribute{
					Name:        "vSAN",
					Description: `It is a virtual Storage Area Network file system.`,
				},
				resource.Attribute{
					Name:        "VirtualVolume",
					Description: `A Virtual Volume datastore represents a storage container in a hypervisor server.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_vmware_host",
			Category:         "virtualization",
			ShortDescription: `The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.`,
			Description: `
The VMware Host entity with its attributes. Every Host belongs to a Datacenter and may run VMs.
`,
			Keywords: []string{
				"virtualization",
				"vmware",
				"host",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `The entity's power state is unknown.`,
				},
				resource.Attribute{
					Name:        "PoweredOn",
					Description: `The entity is powered on.`,
				},
				resource.Attribute{
					Name:        "PoweredOff",
					Description: `The entity is powered down.`,
				},
				resource.Attribute{
					Name:        "StandBy",
					Description: `The entity is in standby mode.`,
				},
				resource.Attribute{
					Name:        "Paused",
					Description: `The entity is in pause state.`,
				},
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `Entity status is unknown.`,
				},
				resource.Attribute{
					Name:        "Degraded",
					Description: `State is degraded, and might impact normal operation of the entity.`,
				},
				resource.Attribute{
					Name:        "Critical",
					Description: `Entity is in a critical state, impacting operations.`,
				},
				resource.Attribute{
					Name:        "Ok",
					Description: `Entity status is in a stable state, operating normally.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_vmware_vcenter",
			Category:         "virtualization",
			ShortDescription: `VMware vCenter entity. The vCenter has a name assigned by user in Intersight.`,
			Description: `
VMware vCenter entity. The vCenter has a name assigned by user in Intersight.
`,
			Keywords: []string{
				"virtualization",
				"vmware",
				"vcenter",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_virtualization_vmware_virtual_machine",
			Category:         "virtualization",
			ShortDescription: `The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.`,
			Description: `
The VMware Virtual machine. It has details such as power state, IP address, resource consumption, etc. Basic elements come from the base class and VMware specific details are provided here.
`,
			Keywords: []string{
				"virtualization",
				"vmware",
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Unknown",
					Description: `Indicates that the guest OS state cannot be determined.`,
				},
				resource.Attribute{
					Name:        "NotRunning",
					Description: `Indicates that the guest OS is not running.`,
				},
				resource.Attribute{
					Name:        "Resetting",
					Description: `Indicates that the guest OS is resetting.`,
				},
				resource.Attribute{
					Name:        "Running",
					Description: `Indicates that the guest OS is running normally.`,
				},
				resource.Attribute{
					Name:        "ShuttingDown",
					Description: `Indicates that the guest OS is shutting down.`,
				},
				resource.Attribute{
					Name:        "Standby",
					Description: `Indicates that the guest OS is in standby mode.`,
				},
				resource.Attribute{
					Name:        "ESXi",
					Description: `The hypervisor running on the HyperFlex cluster is a Vmware ESXi hypervisor of any version.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAp",
					Description: `The hypervisor running on the HyperFlex cluster is Cisco HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "Hyper-V",
					Description: `The hypervisor running on the HyperFlex cluster is Microsoft Hyper-V.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The hypervisor running on the HyperFlex cluster is not known.`,
				},
				resource.Attribute{
					Name:        "Unknown",
					Description: `The entity's power state is unknown.`,
				},
				resource.Attribute{
					Name:        "PoweredOn",
					Description: `The entity is powered on.`,
				},
				resource.Attribute{
					Name:        "PoweredOff",
					Description: `The entity is powered down.`,
				},
				resource.Attribute{
					Name:        "StandBy",
					Description: `The entity is in standby mode.`,
				},
				resource.Attribute{
					Name:        "Paused",
					Description: `The entity is in pause state.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vmedia_policy",
			Category:         "vmedia",
			ShortDescription: `Policy to configure virtual media settings on endpoint.`,
			Description: `
Policy to configure virtual media settings on endpoint.
`,
			Keywords: []string{
				"vmedia",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vmrc_console",
			Category:         "vmrc",
			ShortDescription: `API to launch VMRC console to a VMware virtual machine.`,
			Description: `
API to launch VMRC console to a VMware virtual machine.
`,
			Keywords: []string{
				"vmrc",
				"console",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_eth_adapter_policy",
			Category:         "vnic",
			ShortDescription: `An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.`,
			Description: `
An Ethernet adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. For each VIC Virtual Ethernet Interface various features like VXLAN, NVGRE, ARFS, Interrupt settings, and TCP Offload settings can be configured.
`,
			Keywords: []string{
				"vnic",
				"eth",
				"adapter",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_eth_if",
			Category:         "vnic",
			ShortDescription: `Virtual Ethernet Interface.`,
			Description: `
Virtual Ethernet Interface.
`,
			Keywords: []string{
				"vnic",
				"eth",
				"if",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `Type indicates that there is no IP associated to an vnic.`,
				},
				resource.Attribute{
					Name:        "DHCP",
					Description: `The IP address is assigned using DHCP, if available.`,
				},
				resource.Attribute{
					Name:        "Static",
					Description: `Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.`,
				},
				resource.Attribute{
					Name:        "Pool",
					Description: `An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.`,
				},
				resource.Attribute{
					Name:        "POOL",
					Description: `The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.`,
				},
				resource.Attribute{
					Name:        "STATIC",
					Description: `The user assigns a static mac/wwn address for the Virtual Interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_eth_network_policy",
			Category:         "vnic",
			ShortDescription: `An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.`,
			Description: `
An Ethernet Network policy determines if the port can carry single VLAN (Access) or multiple VLANs (Trunk) traffic. You can specify the VLAN to be associated with an Ethernet packet if no tag is found.
`,
			Keywords: []string{
				"vnic",
				"eth",
				"network",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Standalone",
					Description: `Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.`,
				},
				resource.Attribute{
					Name:        "FIAttached",
					Description: `Servers which are connected to a Fabric Interconnect that is managed by Intersight.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_eth_qos_policy",
			Category:         "vnic",
			ShortDescription: `An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.`,
			Description: `
An Ethernet Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vNIC. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can be specified like burst and rate on the outgoing traffic.
`,
			Keywords: []string{
				"vnic",
				"eth",
				"qos",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Best Effort",
					Description: `QoS Priority for Best-effort traffic.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `QoS Priority for FC traffic.`,
				},
				resource.Attribute{
					Name:        "Platinum",
					Description: `QoS Priority for Platinum traffic.`,
				},
				resource.Attribute{
					Name:        "Gold",
					Description: `QoS Priority for Gold traffic.`,
				},
				resource.Attribute{
					Name:        "Silver",
					Description: `QoS Priority for Silver traffic.`,
				},
				resource.Attribute{
					Name:        "Bronze",
					Description: `QoS Priority for Bronze traffic.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_fc_adapter_policy",
			Category:         "vnic",
			ShortDescription: `A Fibre Channel Adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. You can enable FCP Error Recovery, change the default settings of Queues and Interrupt handling for performance enhancement.`,
			Description: `
A Fibre Channel Adapter policy governs the host-side behavior of the adapter, including how the adapter handles traffic. You can enable FCP Error Recovery, change the default settings of Queues and Interrupt handling for performance enhancement.
`,
			Keywords: []string{
				"vnic",
				"fc",
				"adapter",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_fc_if",
			Category:         "vnic",
			ShortDescription: `Virtual Fibre Channel Interface.`,
			Description: `
Virtual Fibre Channel Interface.
`,
			Keywords: []string{
				"vnic",
				"fc",
				"if",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "fc-initiator",
					Description: `The default value set for vHBA Type Configuration. Fc-initiator specifies vHBA as a consumer of storage. Enables SCSI commands to transfer data and status information between host and target storage systems.`,
				},
				resource.Attribute{
					Name:        "fc-nvme-initiator",
					Description: `Fc-nvme-initiator specifies vHBA as a consumer of storage. Enables NVMe-based message commands to transfer data and status information between host and target storage systems.`,
				},
				resource.Attribute{
					Name:        "fc-nvme-target",
					Description: `Fc-nvme-target specifies vHBA as a provider of storage volumes to initiators. Enables NVMe-based message commands to transfer data and status information between host and target storage systems. Currently tech-preview, only enabled with an asynchronous driver.`,
				},
				resource.Attribute{
					Name:        "fc-target",
					Description: `Fc-target specifies vHBA as a provider of storage volumes to initiators. Enables SCSI commands to transfer data and status information between host and target storage systems. fc-target is enabled only with an asynchronous driver.`,
				},
				resource.Attribute{
					Name:        "POOL",
					Description: `The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.`,
				},
				resource.Attribute{
					Name:        "STATIC",
					Description: `The user assigns a static mac/wwn address for the Virtual Interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_fc_network_policy",
			Category:         "vnic",
			ShortDescription: `A Fibre Channel Network policy governs the vSAN configuration for the virtual interfaces.`,
			Description: `
A Fibre Channel Network policy governs the vSAN configuration for the virtual interfaces.
`,
			Keywords: []string{
				"vnic",
				"fc",
				"network",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_fc_qos_policy",
			Category:         "vnic",
			ShortDescription: `A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.`,
			Description: `
A Fibre Channel Quality of Service (QoS) policy assigns a system class to the outgoing traffic for a vHBA. This system class determines the quality of service for the outgoing traffic. For certain adapters additional controls can also be specified like burst and rate on the outgoing traffic.
`,
			Keywords: []string{
				"vnic",
				"fc",
				"qos",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Best Effort",
					Description: `QoS Priority for Best-effort traffic.`,
				},
				resource.Attribute{
					Name:        "FC",
					Description: `QoS Priority for FC traffic.`,
				},
				resource.Attribute{
					Name:        "Platinum",
					Description: `QoS Priority for Platinum traffic.`,
				},
				resource.Attribute{
					Name:        "Gold",
					Description: `QoS Priority for Gold traffic.`,
				},
				resource.Attribute{
					Name:        "Silver",
					Description: `QoS Priority for Silver traffic.`,
				},
				resource.Attribute{
					Name:        "Bronze",
					Description: `QoS Priority for Bronze traffic.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_iscsi_adapter_policy",
			Category:         "vnic",
			ShortDescription: `Set of iSCSI properties that govern the host-side behavior of the adapter.`,
			Description: `
Set of iSCSI properties that govern the host-side behavior of the adapter.
`,
			Keywords: []string{
				"vnic",
				"iscsi",
				"adapter",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_iscsi_boot_policy",
			Category:         "vnic",
			ShortDescription: `Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.`,
			Description: `
Configuration parameters to enable a server to boot its operating system from an iSCSI target machine located remotely over a network.
`,
			Keywords: []string{
				"vnic",
				"iscsi",
				"boot",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "DHCP",
					Description: `The IP address is assigned using DHCP, if available.`,
				},
				resource.Attribute{
					Name:        "Static",
					Description: `Static IPv4 address is assigned to the iSCSI boot interface based on the information entered in this area.`,
				},
				resource.Attribute{
					Name:        "Pool",
					Description: `An IPv4 address is assigned to the iSCSI boot interface from the management IP address pool.`,
				},
				resource.Attribute{
					Name:        "Static",
					Description: `Type indicates that static target interface is assigned to iSCSI boot.`,
				},
				resource.Attribute{
					Name:        "Auto",
					Description: `Type indicates that the system selects the target interface automatically during iSCSI boot.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_iscsi_static_target_policy",
			Category:         "vnic",
			ShortDescription: `Configuration parameters that defines the reachability of iSCSI Target portal.`,
			Description: `
Configuration parameters that defines the reachability of iSCSI Target portal.
`,
			Keywords: []string{
				"vnic",
				"iscsi",
				"static",
				"target",
				"policy",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_lan_connectivity_policy",
			Category:         "vnic",
			ShortDescription: `A LAN Connectivity Policy determines the network resources and the connections between the server and the LAN on the network. This policy uses Consistent Device Naming to configure the vNIC. You can configure a usNIC or VMQ connection for the vNIC to improve network performance.`,
			Description: `
A LAN Connectivity Policy determines the network resources and the connections between the server and the LAN on the network. This policy uses Consistent Device Naming to configure the vNIC. You can configure a usNIC or VMQ connection for the vNIC to improve network performance.
`,
			Keywords: []string{
				"vnic",
				"lan",
				"connectivity",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `Type indicates that there is no IQN associated to an interface.`,
				},
				resource.Attribute{
					Name:        "Static",
					Description: `Type represents that static IQN is associated to an interface.`,
				},
				resource.Attribute{
					Name:        "Pool",
					Description: `Type indicates that IQN value is sourced from an associated pool.`,
				},
				resource.Attribute{
					Name:        "custom",
					Description: `The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.`,
				},
				resource.Attribute{
					Name:        "Standalone",
					Description: `Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.`,
				},
				resource.Attribute{
					Name:        "FIAttached",
					Description: `Servers which are connected to a Fabric Interconnect that is managed by Intersight.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_lcp_status",
			Category:         "vnic",
			ShortDescription: `An internal MO to check if a LCP can be deployed or not on a specific Server Profile.`,
			Description: `
An internal MO to check if a LCP can be deployed or not on a specific Server Profile.
`,
			Keywords: []string{
				"vnic",
				"lcp",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ok",
					Description: `No issues with the LCP/SCP/VIF.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `The LCP/SCP/VIF cannot be deployed due to error.`,
				},
				resource.Attribute{
					Name:        "validating",
					Description: `Validation in progress for the LCP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_san_connectivity_policy",
			Category:         "vnic",
			ShortDescription: `SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.`,
			Description: `
SAN connectivity policy determines the network storage resources and the connections between the server and the SAN on the network. This policy enables configuration of vHBAs that the servers use to communicate with the storage network.
`,
			Keywords: []string{
				"vnic",
				"san",
				"connectivity",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "custom",
					Description: `The placement of the vNICs / vHBAs on network adapters is manually chosen by the user.`,
				},
				resource.Attribute{
					Name:        "auto",
					Description: `The placement of the vNICs / vHBAs on network adapters is automatically determined by the system.`,
				},
				resource.Attribute{
					Name:        "Standalone",
					Description: `Servers which are operating in standalone mode i.e. not connected to a Fabric Interconnected.`,
				},
				resource.Attribute{
					Name:        "FIAttached",
					Description: `Servers which are connected to a Fabric Interconnect that is managed by Intersight.`,
				},
				resource.Attribute{
					Name:        "POOL",
					Description: `The user selects a pool from which the mac/wwn address will be leased for the Virtual Interface.`,
				},
				resource.Attribute{
					Name:        "STATIC",
					Description: `The user assigns a static mac/wwn address for the Virtual Interface.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vnic_scp_status",
			Category:         "vnic",
			ShortDescription: `An internal MO to check if a SCP can be deployed or not on a specific Server Profile.`,
			Description: `
An internal MO to check if a SCP can be deployed or not on a specific Server Profile.
`,
			Keywords: []string{
				"vnic",
				"scp",
				"status",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ok",
					Description: `No issues with the LCP/SCP/VIF.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `The LCP/SCP/VIF cannot be deployed due to error.`,
				},
				resource.Attribute{
					Name:        "validating",
					Description: `Validation in progress for the LCP.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_vrf_vrf",
			Category:         "vrf",
			ShortDescription: `Virtual Routing and Forwarding (VRF) is a networking technology that implements an isolated Layer 3 domain.`,
			Description: `
Virtual Routing and Forwarding (VRF) is a networking technology that implements an isolated Layer 3 domain.
`,
			Keywords: []string{
				"vrf",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_batch_api_executor",
			Category:         "workflow",
			ShortDescription: `Intersight allows generic API tasks to be created by taking the API request body and a response parser specification in the form of content.Grammar object. Batch API associates the list of API requests to be executed as part of single task execution. Each API request takes the request body and a response parser specification.`,
			Description: `
Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.
Batch API associates the list of API requests to be executed as part of single
task execution. Each API request takes the request body and a response parser
specification.
`,
			Keywords: []string{
				"workflow",
				"batch",
				"api",
				"executor",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_build_task_meta",
			Category:         "workflow",
			ShortDescription: `Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.`,
			Description: `
Contains relationship for tasks within a workflow. It is used to dynamically generate a workflow.
`,
			Keywords: []string{
				"workflow",
				"build",
				"task",
				"meta",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_build_task_meta_owner",
			Category:         "workflow",
			ShortDescription: `Contains the list of dynamic workflow types that a microservice supports.`,
			Description: `
Contains the list of dynamic workflow types that a microservice supports.
`,
			Keywords: []string{
				"workflow",
				"build",
				"task",
				"meta",
				"owner",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_catalog",
			Category:         "workflow",
			ShortDescription: `A catalog of workflow related objects such as workflow and task definitions. Each user account will have a local workflow catalog where account users can store their private workflow and task definitions. Cisco provides validated workflows and tasks to Intersight users via shared catalogs. Intersight users will be able to read, run these workflows and tasks within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.`,
			Description: `
A catalog of workflow related objects such as workflow and task definitions. Each user account will have a local workflow catalog where account users can store their private workflow and task definitions.
Cisco provides validated workflows and tasks to Intersight users via shared catalogs. Intersight users will be able to read, run these workflows and tasks within their account context. The shared catalogs will be managed entirely by Cisco. Contributions to shared catalogs will need to be provided to Cisco who will publish them at their own discretion.
`,
			Keywords: []string{
				"workflow",
				"catalog",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_custom_data_type_definition",
			Category:         "workflow",
			ShortDescription: `Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.`,
			Description: `
Captures a customized data type definition that can be used for task or workflow input/output. This can be reused across multiple tasks and workflow definitions.
`,
			Keywords: []string{
				"workflow",
				"custom",
				"data",
				"type",
				"definition",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_error_response_handler",
			Category:         "workflow",
			ShortDescription: `Intersight allows generic API tasks to be created by taking the API request body and a response parser specification in the form of content.Grammar object. Error Response Handler allows to create a generic error response specification which can be used by multiple Batch API. The parameters provided in the Error Response Handler may be used to parse error responses from an API request, if the response specification provided for the API request does not define error parameters.`,
			Description: `
Intersight allows generic API tasks to be created by taking the API request
body and a response parser specification in the form of content.Grammar object.
Error Response Handler allows to create a generic error response specification
which can be used by multiple Batch API. The parameters provided in the Error
Response Handler may be used to parse error responses from an API request, if
the response specification provided for the API request does not define
error parameters.
`,
			Keywords: []string{
				"workflow",
				"error",
				"response",
				"handler",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "APIC",
					Description: `An Application Policy Infrastructure Controller cluster.`,
				},
				resource.Attribute{
					Name:        "DCNM",
					Description: `A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center.`,
				},
				resource.Attribute{
					Name:        "UCSFI",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM).`,
				},
				resource.Attribute{
					Name:        "UCSFIISM",
					Description: `A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight.`,
				},
				resource.Attribute{
					Name:        "IMC",
					Description: `A standalone UCS Server Integrated Management Controller.`,
				},
				resource.Attribute{
					Name:        "IMCM4",
					Description: `A standalone UCS M4 Server.`,
				},
				resource.Attribute{
					Name:        "IMCM5",
					Description: `A standalone UCS M5 server.`,
				},
				resource.Attribute{
					Name:        "UCSIOM",
					Description: `An UCS Chassis IO module.`,
				},
				resource.Attribute{
					Name:        "HX",
					Description: `A HyperFlex storage controller.`,
				},
				resource.Attribute{
					Name:        "HyperFlexAP",
					Description: `A HyperFlex Application Platform.`,
				},
				resource.Attribute{
					Name:        "UCSD",
					Description: `A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware.`,
				},
				resource.Attribute{
					Name:        "IntersightAppliance",
					Description: `A Cisco Intersight Connected Virtual Appliance.`,
				},
				resource.Attribute{
					Name:        "IntersightAssist",
					Description: `A Cisco Intersight Assist.`,
				},
				resource.Attribute{
					Name:        "PureStorageFlashArray",
					Description: `A Pure Storage FlashArray device.`,
				},
				resource.Attribute{
					Name:        "NetAppOntap",
					Description: `A NetApp ONTAP storage system.`,
				},
				resource.Attribute{
					Name:        "NetAppActiveIqUnifiedManager",
					Description: `A NetApp Active IQ Unified Manager.`,
				},
				resource.Attribute{
					Name:        "EmcScaleIo",
					Description: `An EMC ScaleIO storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVmax",
					Description: `An EMC VMAX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcVplex",
					Description: `An EMC VPLEX storage system.`,
				},
				resource.Attribute{
					Name:        "EmcXtremIo",
					Description: `An EMC XtremIO storage system.`,
				},
				resource.Attribute{
					Name:        "VmwareVcenter",
					Description: `A VMware vCenter device that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "MicrosoftHyperV",
					Description: `A Microsoft HyperV system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "AppDynamics",
					Description: `An AppDynamics controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "Dynatrace",
					Description: `A Dynatrace controller that monitors applications.`,
				},
				resource.Attribute{
					Name:        "MicrosoftSqlServer",
					Description: `A Microsoft SQL database server.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `A Kubernetes cluster that runs containerized applications.`,
				},
				resource.Attribute{
					Name:        "AmazonWebService",
					Description: `A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost.`,
				},
				resource.Attribute{
					Name:        "AmazonWebServiceBilling",
					Description: `A Amazon web service billing target to retrieve billing information stored in S3 bucket.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureServicePrincipal",
					Description: `A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions.`,
				},
				resource.Attribute{
					Name:        "MicrosoftAzureEnterpriseAgreement",
					Description: `A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs.`,
				},
				resource.Attribute{
					Name:        "DellCompellent",
					Description: `A Dell Compellent storage system.`,
				},
				resource.Attribute{
					Name:        "HPE3Par",
					Description: `A HPE 3PAR storage system.`,
				},
				resource.Attribute{
					Name:        "RedHatEnterpriseVirtualization",
					Description: `A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines.`,
				},
				resource.Attribute{
					Name:        "NutanixAcropolis",
					Description: `A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform.`,
				},
				resource.Attribute{
					Name:        "HPEOneView",
					Description: `A HPE Oneview management system that manages compute, storage, and networking.`,
				},
				resource.Attribute{
					Name:        "ServiceEngine",
					Description: `Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications.`,
				},
				resource.Attribute{
					Name:        "HitachiVirtualStoragePlatform",
					Description: `A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers.`,
				},
				resource.Attribute{
					Name:        "IMCBlade",
					Description: `An Intersight managed UCS Blade Server.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `A Terraform Cloud account.`,
				},
				resource.Attribute{
					Name:        "TerraformAgent",
					Description: `A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent.`,
				},
				resource.Attribute{
					Name:        "CustomTarget",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "HTTPEndpoint",
					Description: `An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic.`,
				},
				resource.Attribute{
					Name:        "CiscoCatalyst",
					Description: `A Cisco Catalyst networking switch device.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_pending_dynamic_workflow_info",
			Category:         "workflow",
			ShortDescription: `Information for a pending Dynamic Workflow Instance before it is run. Validation needs to be done on the dynamic workflow tasks before starting. After it begins, it will be tracked with regular WorkflowInstance.`,
			Description: `
Information for a pending Dynamic Workflow Instance before it is run.  Validation needs to be done on the dynamic workflow tasks before starting.  After it begins, it will be tracked with regular WorkflowInstance.
`,
			Keywords: []string{
				"workflow",
				"pending",
				"dynamic",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "GatheringTasks",
					Description: `Dynamic workflow is gathering tasks before workflow can start execution.`,
				},
				resource.Attribute{
					Name:        "Waiting",
					Description: `Dynamic workflow is in waiting state and not yet started execution.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_rollback_workflow",
			Category:         "workflow",
			ShortDescription: `Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.`,
			Description: `
Rollback workflow contains details about the workflow instance, tasks to be rollback along with the status and workflow instances.
`,
			Keywords: []string{
				"workflow",
				"rollback",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `If no action is set, then the default value is set to none for the action field.`,
				},
				resource.Attribute{
					Name:        "Create",
					Description: `Create rollback workflow data for the execution of the rollback workflow.`,
				},
				resource.Attribute{
					Name:        "Start",
					Description: `Start a new execution of the rollback workflow.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `If no status is set, then the default value is set none for the status field.`,
				},
				resource.Attribute{
					Name:        "Created",
					Description: `Status of the rollback workflow when it identifies the eligible tasks for rollback.`,
				},
				resource.Attribute{
					Name:        "Running",
					Description: `Status of the rollback workflow when it is in-progress.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Status of the rollback workflow after execution is successful.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Status of the rollback workflow after execution results in failure.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_task_definition",
			Category:         "workflow",
			ShortDescription: `Used to define a task which can be included within a workflow. Task definition conveys the intent that we want to achieve with the task. We can have a standalone task definition that is bound to a single implementation for that task, or we can define an TaskDefinition that will serve as the interface task definition which is linked to multiple implementation tasks. Each implemented TaskDefinition will be bound to its own implementation so we can achieve a case where single TaskDefinition has multiple implementations.`,
			Description: `
Used to define a task which can be included within a workflow. Task definition conveys the intent that we want to achieve with the task. We can have a standalone task definition that is bound to a single implementation for that task, or we can define an TaskDefinition that will serve as the interface task definition which is linked to multiple implementation tasks. Each implemented TaskDefinition will be bound to its own implementation so we can achieve a case where single TaskDefinition has multiple implementations.
`,
			Keywords: []string{
				"workflow",
				"task",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_task_info",
			Category:         "workflow",
			ShortDescription: `Task instance which represents the run time instance of a task within a workflow.`,
			Description: `
Task instance which represents the run time instance of a task within a workflow.
`,
			Keywords: []string{
				"workflow",
				"task",
				"info",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_task_metadata",
			Category:         "workflow",
			ShortDescription: `Task details is a collection of properties that are common across all the versions of a task.`,
			Description: `
Task details is a collection of properties that are common across all the versions of a task.
`,
			Keywords: []string{
				"workflow",
				"task",
				"metadata",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_workflow_definition",
			Category:         "workflow",
			ShortDescription: `Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.`,
			Description: `
Workflow definition is a collection of tasks that are sequenced in a certain way using control tasks. The tasks in the workflow definition is represented as a directed acyclic graph where each node in the graph is a task and the edges in the graph are transitions from one task to another.
`,
			Keywords: []string{
				"workflow",
				"definition",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Base",
					Description: `Base as a License type. It is default license type.`,
				},
				resource.Attribute{
					Name:        "Essential",
					Description: `Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "Standard",
					Description: `Standard as a License type.`,
				},
				resource.Attribute{
					Name:        "Advantage",
					Description: `Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "Premier",
					Description: `Premier as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Essential",
					Description: `IWO-Essential as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Advantage",
					Description: `IWO-Advantage as a License type.`,
				},
				resource.Attribute{
					Name:        "IWO-Premier",
					Description: `IWO-Premier as a License type.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_workflow_info",
			Category:         "workflow",
			ShortDescription: `Contains information for a workflow execution which is a runtime instance of workflow.`,
			Description: `
Contains information for a workflow execution which is a runtime instance of workflow.
`,
			Keywords: []string{
				"workflow",
				"info",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `No action is set, this is the default value for action field.`,
				},
				resource.Attribute{
					Name:        "Create",
					Description: `Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.`,
				},
				resource.Attribute{
					Name:        "Start",
					Description: `Start a new execution of the workflow.`,
				},
				resource.Attribute{
					Name:        "Pause",
					Description: `Pause the workflow, this can only be issued on workflows that are in running state.`,
				},
				resource.Attribute{
					Name:        "Resume",
					Description: `Resume the workflow which was previously paused through pause action on the workflow.`,
				},
				resource.Attribute{
					Name:        "Retry",
					Description: `Retry the workflow that has previously reached a final state and has the retryable property set to true on the workflow. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration.`,
				},
				resource.Attribute{
					Name:        "RetryFailed",
					Description: `Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `Cancel the workflow that is in running or waiting state.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `No action is set, this is the default value for action field.`,
				},
				resource.Attribute{
					Name:        "Create",
					Description: `Create a new instance of the workflow but it does not start the execution of the workflow. Use the Start action to start execution of the workflow.`,
				},
				resource.Attribute{
					Name:        "Start",
					Description: `Start a new execution of the workflow.`,
				},
				resource.Attribute{
					Name:        "Pause",
					Description: `Pause the workflow, this can only be issued on workflows that are in running state.`,
				},
				resource.Attribute{
					Name:        "Resume",
					Description: `Resume the workflow which was previously paused through pause action on the workflow.`,
				},
				resource.Attribute{
					Name:        "Retry",
					Description: `Retry the workflow that has previously reached a final state and has the retryable property set to true on the workflow. A running or waiting workflow cannot be retried. If the property retryFromTaskName is also passed along with this action, the workflow will be started from that specific task, otherwise the workflow will be restarted. The task name must be one of the tasks that completed or failed in the previous run, you cannot retry a workflow from a task which wasn't run in the previous iteration.`,
				},
				resource.Attribute{
					Name:        "RetryFailed",
					Description: `Retry the workflow that has failed. A running or waiting workflow or a workflow that completed successfully cannot be retried. Only the tasks that failed in the previous run will be retried and the rest of workflow will be run. This action does not restart the workflow and also does not support retrying from a specific task.`,
				},
				resource.Attribute{
					Name:        "Cancel",
					Description: `Cancel the workflow that is in running or waiting state.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Pause reason is none, which indicates there is no reason for the pause state.`,
				},
				resource.Attribute{
					Name:        "TaskWithWarning",
					Description: `Pause reason indicates the workflow is in this state due to a task that has a status as completed with warnings.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Wait reason is none, which indicates there is no reason for the waiting state.`,
				},
				resource.Attribute{
					Name:        "GatherTasks",
					Description: `Wait reason is gathering tasks, which indicates the workflow is in this state in order to gather tasks.`,
				},
				resource.Attribute{
					Name:        "Duplicate",
					Description: `Wait reason is duplicate, which indicates the workflow is a duplicate of current running workflow.`,
				},
				resource.Attribute{
					Name:        "RateLimit",
					Description: `Wait reason is rate limit, which indicates the workflow is rate limited by account/instance level throttling threshold.`,
				},
				resource.Attribute{
					Name:        "WaitTask",
					Description: `Wait reason when there are one or more wait tasks in the workflow which are yet to receive a task status update.`,
				},
				resource.Attribute{
					Name:        "PendingRetryFailed",
					Description: `Wait reason when the workflow is pending a RetryFailed action.`,
				},
				resource.Attribute{
					Name:        "SystemDefined",
					Description: `System defined workflow definition.`,
				},
				resource.Attribute{
					Name:        "UserDefined",
					Description: `User defined workflow definition.`,
				},
				resource.Attribute{
					Name:        "Dynamic",
					Description: `Dynamically defined workflow definition.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_workflow_meta",
			Category:         "workflow",
			ShortDescription: `Contains a workflow definition which is a sequence of tasks to execute.`,
			Description: `
Contains a workflow definition which is a sequence of tasks to execute.
`,
			Keywords: []string{
				"workflow",
				"meta",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "SystemDefined",
					Description: `System defined workflow definition.`,
				},
				resource.Attribute{
					Name:        "UserDefined",
					Description: `User defined workflow definition.`,
				},
				resource.Attribute{
					Name:        "Dynamic",
					Description: `Dynamically defined workflow definition.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "intersight_workflow_workflow_metadata",
			Category:         "workflow",
			ShortDescription: `Workflow metadata is a collection of properties that are common across all the versions of a workflow.`,
			Description: `
Workflow metadata is a collection of properties that are common across all the versions of a workflow.
`,
			Keywords: []string{
				"workflow",
				"metadata",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
	}

	dataSourcesMap = map[string]int{

		"intersight_aaa_audit_record":                                        0,
		"intersight_access_policy":                                           1,
		"intersight_adapter_config_policy":                                   2,
		"intersight_adapter_ext_eth_interface":                               3,
		"intersight_adapter_host_eth_interface":                              4,
		"intersight_adapter_host_fc_interface":                               5,
		"intersight_adapter_host_iscsi_interface":                            6,
		"intersight_adapter_unit":                                            7,
		"intersight_adapter_unit_expander":                                   8,
		"intersight_appliance_app_status":                                    9,
		"intersight_appliance_backup":                                        10,
		"intersight_appliance_backup_policy":                                 11,
		"intersight_appliance_certificate_setting":                           12,
		"intersight_appliance_data_export_policy":                            13,
		"intersight_appliance_device_certificate":                            14,
		"intersight_appliance_device_claim":                                  15,
		"intersight_appliance_diag_setting":                                  16,
		"intersight_appliance_external_syslog_setting":                       17,
		"intersight_appliance_file_system_status":                            18,
		"intersight_appliance_group_status":                                  19,
		"intersight_appliance_image_bundle":                                  20,
		"intersight_appliance_node_info":                                     21,
		"intersight_appliance_node_status":                                   22,
		"intersight_appliance_release_note":                                  23,
		"intersight_appliance_remote_file_import":                            24,
		"intersight_appliance_restore":                                       25,
		"intersight_appliance_setup_info":                                    26,
		"intersight_appliance_system_info":                                   27,
		"intersight_appliance_system_status":                                 28,
		"intersight_appliance_upgrade":                                       29,
		"intersight_appliance_upgrade_policy":                                30,
		"intersight_asset_cluster_member":                                    31,
		"intersight_asset_device_configuration":                              32,
		"intersight_asset_device_connector_manager":                          33,
		"intersight_asset_device_contract_information":                       34,
		"intersight_asset_device_registration":                               35,
		"intersight_asset_subscription_device_contract_information":          36,
		"intersight_asset_target":                                            37,
		"intersight_bios_boot_device":                                        38,
		"intersight_bios_boot_mode":                                          39,
		"intersight_bios_policy":                                             40,
		"intersight_bios_system_boot_order":                                  41,
		"intersight_bios_unit":                                               42,
		"intersight_boot_cdd_device":                                         43,
		"intersight_boot_device_boot_mode":                                   44,
		"intersight_boot_device_boot_security":                               45,
		"intersight_boot_hdd_device":                                         46,
		"intersight_boot_iscsi_device":                                       47,
		"intersight_boot_nvme_device":                                        48,
		"intersight_boot_pch_storage_device":                                 49,
		"intersight_boot_precision_policy":                                   50,
		"intersight_boot_pxe_device":                                         51,
		"intersight_boot_san_device":                                         52,
		"intersight_boot_sd_device":                                          53,
		"intersight_boot_uefi_shell_device":                                  54,
		"intersight_boot_usb_device":                                         55,
		"intersight_boot_vmedia_device":                                      56,
		"intersight_capability_adapter_unit_descriptor":                      57,
		"intersight_capability_catalog":                                      58,
		"intersight_capability_chassis_descriptor":                           59,
		"intersight_capability_chassis_manufacturing_def":                    60,
		"intersight_capability_cimc_firmware_descriptor":                     61,
		"intersight_capability_equipment_physical_def":                       62,
		"intersight_capability_equipment_slot_array":                         63,
		"intersight_capability_fan_module_descriptor":                        64,
		"intersight_capability_fan_module_manufacturing_def":                 65,
		"intersight_capability_io_card_capability_def":                       66,
		"intersight_capability_io_card_descriptor":                           67,
		"intersight_capability_io_card_manufacturing_def":                    68,
		"intersight_capability_port_group_aggregation_def":                   69,
		"intersight_capability_psu_descriptor":                               70,
		"intersight_capability_psu_manufacturing_def":                        71,
		"intersight_capability_sioc_module_capability_def":                   72,
		"intersight_capability_sioc_module_descriptor":                       73,
		"intersight_capability_sioc_module_manufacturing_def":                74,
		"intersight_capability_switch_capability":                            75,
		"intersight_capability_switch_descriptor":                            76,
		"intersight_capability_switch_manufacturing_def":                     77,
		"intersight_certificatemanagement_policy":                            78,
		"intersight_chassis_config_change_detail":                            79,
		"intersight_chassis_config_import":                                   80,
		"intersight_chassis_config_result":                                   81,
		"intersight_chassis_config_result_entry":                             82,
		"intersight_chassis_iom_profile":                                     83,
		"intersight_chassis_profile":                                         84,
		"intersight_cloud_regions":                                           85,
		"intersight_cloud_sku_container_type":                                86,
		"intersight_cloud_sku_database_type":                                 87,
		"intersight_cloud_sku_instance_type":                                 88,
		"intersight_cloud_sku_network_type":                                  89,
		"intersight_cloud_sku_volume_type":                                   90,
		"intersight_comm_http_proxy_policy":                                  91,
		"intersight_compute_blade":                                           92,
		"intersight_compute_blade_identity":                                  93,
		"intersight_compute_board":                                           94,
		"intersight_compute_physical_summary":                                95,
		"intersight_compute_rack_unit":                                       96,
		"intersight_compute_rack_unit_identity":                              97,
		"intersight_compute_server_setting":                                  98,
		"intersight_compute_vmedia":                                          99,
		"intersight_cond_alarm":                                              100,
		"intersight_cond_alarm_aggregation":                                  101,
		"intersight_cond_hcl_status":                                         102,
		"intersight_cond_hcl_status_detail":                                  103,
		"intersight_cond_hcl_status_job":                                     104,
		"intersight_config_exported_item":                                    105,
		"intersight_config_exporter":                                         106,
		"intersight_config_imported_item":                                    107,
		"intersight_config_importer":                                         108,
		"intersight_connectorpack_connector_pack_upgrade":                    109,
		"intersight_connectorpack_upgrade_impact":                            110,
		"intersight_deviceconnector_policy":                                  111,
		"intersight_equipment_chassis":                                       112,
		"intersight_equipment_chassis_identity":                              113,
		"intersight_equipment_chassis_operation":                             114,
		"intersight_equipment_device_summary":                                115,
		"intersight_equipment_fan":                                           116,
		"intersight_equipment_fan_module":                                    117,
		"intersight_equipment_fex":                                           118,
		"intersight_equipment_fex_identity":                                  119,
		"intersight_equipment_fex_operation":                                 120,
		"intersight_equipment_identity_summary":                              121,
		"intersight_equipment_io_card":                                       122,
		"intersight_equipment_io_card_operation":                             123,
		"intersight_equipment_io_expander":                                   124,
		"intersight_equipment_locator_led":                                   125,
		"intersight_equipment_psu":                                           126,
		"intersight_equipment_psu_control":                                   127,
		"intersight_equipment_rack_enclosure":                                128,
		"intersight_equipment_rack_enclosure_slot":                           129,
		"intersight_equipment_shared_io_module":                              130,
		"intersight_equipment_switch_card":                                   131,
		"intersight_equipment_system_io_controller":                          132,
		"intersight_equipment_tpm":                                           133,
		"intersight_equipment_transceiver":                                   134,
		"intersight_ether_host_port":                                         135,
		"intersight_ether_network_port":                                      136,
		"intersight_ether_physical_port":                                     137,
		"intersight_ether_port_channel":                                      138,
		"intersight_externalsite_authorization":                              139,
		"intersight_fabric_appliance_pc_role":                                140,
		"intersight_fabric_appliance_role":                                   141,
		"intersight_fabric_config_change_detail":                             142,
		"intersight_fabric_config_result":                                    143,
		"intersight_fabric_config_result_entry":                              144,
		"intersight_fabric_eth_network_control_policy":                       145,
		"intersight_fabric_eth_network_group_policy":                         146,
		"intersight_fabric_eth_network_policy":                               147,
		"intersight_fabric_fc_network_policy":                                148,
		"intersight_fabric_fc_uplink_pc_role":                                149,
		"intersight_fabric_fc_uplink_role":                                   150,
		"intersight_fabric_fcoe_uplink_pc_role":                              151,
		"intersight_fabric_fcoe_uplink_role":                                 152,
		"intersight_fabric_multicast_policy":                                 153,
		"intersight_fabric_pc_member":                                        154,
		"intersight_fabric_pc_operation":                                     155,
		"intersight_fabric_port_mode":                                        156,
		"intersight_fabric_port_operation":                                   157,
		"intersight_fabric_port_policy":                                      158,
		"intersight_fabric_server_role":                                      159,
		"intersight_fabric_switch_cluster_profile":                           160,
		"intersight_fabric_switch_control_policy":                            161,
		"intersight_fabric_switch_profile":                                   162,
		"intersight_fabric_system_qos_policy":                                163,
		"intersight_fabric_uplink_pc_role":                                   164,
		"intersight_fabric_uplink_role":                                      165,
		"intersight_fabric_vlan":                                             166,
		"intersight_fabric_vsan":                                             167,
		"intersight_fault_instance":                                          168,
		"intersight_fc_physical_port":                                        169,
		"intersight_fc_port_channel":                                         170,
		"intersight_fcpool_fc_block":                                         171,
		"intersight_fcpool_lease":                                            172,
		"intersight_fcpool_pool":                                             173,
		"intersight_fcpool_pool_member":                                      174,
		"intersight_fcpool_universe":                                         175,
		"intersight_firmware_bios_descriptor":                                176,
		"intersight_firmware_board_controller_descriptor":                    177,
		"intersight_firmware_chassis_upgrade":                                178,
		"intersight_firmware_cimc_descriptor":                                179,
		"intersight_firmware_dimm_descriptor":                                180,
		"intersight_firmware_distributable":                                  181,
		"intersight_firmware_distributable_meta":                             182,
		"intersight_firmware_drive_descriptor":                               183,
		"intersight_firmware_driver_distributable":                           184,
		"intersight_firmware_eula":                                           185,
		"intersight_firmware_firmware_summary":                               186,
		"intersight_firmware_gpu_descriptor":                                 187,
		"intersight_firmware_hba_descriptor":                                 188,
		"intersight_firmware_iom_descriptor":                                 189,
		"intersight_firmware_mswitch_descriptor":                             190,
		"intersight_firmware_nxos_descriptor":                                191,
		"intersight_firmware_pcie_descriptor":                                192,
		"intersight_firmware_psu_descriptor":                                 193,
		"intersight_firmware_running_firmware":                               194,
		"intersight_firmware_sas_expander_descriptor":                        195,
		"intersight_firmware_server_configuration_utility_distributable":     196,
		"intersight_firmware_storage_controller_descriptor":                  197,
		"intersight_firmware_switch_upgrade":                                 198,
		"intersight_firmware_unsupported_version_upgrade":                    199,
		"intersight_firmware_upgrade":                                        200,
		"intersight_firmware_upgrade_impact_status":                          201,
		"intersight_firmware_upgrade_status":                                 202,
		"intersight_forecast_catalog":                                        203,
		"intersight_forecast_definition":                                     204,
		"intersight_forecast_instance":                                       205,
		"intersight_graphics_card":                                           206,
		"intersight_graphics_controller":                                     207,
		"intersight_hcl_driver_image":                                        208,
		"intersight_hcl_exempted_catalog":                                    209,
		"intersight_hcl_hyperflex_software_compatibility_info":               210,
		"intersight_hcl_operating_system":                                    211,
		"intersight_hcl_operating_system_vendor":                             212,
		"intersight_hyperflex_alarm":                                         213,
		"intersight_hyperflex_app_catalog":                                   214,
		"intersight_hyperflex_auto_support_policy":                           215,
		"intersight_hyperflex_backup_cluster":                                216,
		"intersight_hyperflex_capability_info":                               217,
		"intersight_hyperflex_cisco_hypervisor_manager":                      218,
		"intersight_hyperflex_cluster":                                       219,
		"intersight_hyperflex_cluster_backup_policy":                         220,
		"intersight_hyperflex_cluster_backup_policy_deployment":              221,
		"intersight_hyperflex_cluster_health_check_execution_snapshot":       222,
		"intersight_hyperflex_cluster_network_policy":                        223,
		"intersight_hyperflex_cluster_profile":                               224,
		"intersight_hyperflex_cluster_replication_network_policy":            225,
		"intersight_hyperflex_cluster_replication_network_policy_deployment": 226,
		"intersight_hyperflex_cluster_storage_policy":                        227,
		"intersight_hyperflex_config_result":                                 228,
		"intersight_hyperflex_config_result_entry":                           229,
		"intersight_hyperflex_device_package_download_state":                 230,
		"intersight_hyperflex_ext_fc_storage_policy":                         231,
		"intersight_hyperflex_ext_iscsi_storage_policy":                      232,
		"intersight_hyperflex_feature_limit_external":                        233,
		"intersight_hyperflex_feature_limit_internal":                        234,
		"intersight_hyperflex_health":                                        235,
		"intersight_hyperflex_health_check_definition":                       236,
		"intersight_hyperflex_health_check_execution":                        237,
		"intersight_hyperflex_health_check_execution_snapshot":               238,
		"intersight_hyperflex_health_check_package_checksum":                 239,
		"intersight_hyperflex_hxap_cluster":                                  240,
		"intersight_hyperflex_hxap_datacenter":                               241,
		"intersight_hyperflex_hxap_event":                                    242,
		"intersight_hyperflex_hxap_host":                                     243,
		"intersight_hyperflex_hxap_virtual_disk":                             244,
		"intersight_hyperflex_hxap_virtual_machine":                          245,
		"intersight_hyperflex_hxdp_version":                                  246,
		"intersight_hyperflex_local_credential_policy":                       247,
		"intersight_hyperflex_node":                                          248,
		"intersight_hyperflex_node_config_policy":                            249,
		"intersight_hyperflex_node_profile":                                  250,
		"intersight_hyperflex_proxy_setting_policy":                          251,
		"intersight_hyperflex_server_firmware_version":                       252,
		"intersight_hyperflex_server_firmware_version_entry":                 253,
		"intersight_hyperflex_server_model":                                  254,
		"intersight_hyperflex_software_version_policy":                       255,
		"intersight_hyperflex_sys_config_policy":                             256,
		"intersight_hyperflex_ucsm_config_policy":                            257,
		"intersight_hyperflex_vcenter_config_policy":                         258,
		"intersight_hyperflex_vm_backup_info":                                259,
		"intersight_hyperflex_vm_restore_operation":                          260,
		"intersight_hyperflex_vm_snapshot_info":                              261,
		"intersight_iaas_connector_pack":                                     262,
		"intersight_iaas_device_status":                                      263,
		"intersight_iaas_diagnostic_messages":                                264,
		"intersight_iaas_license_info":                                       265,
		"intersight_iaas_most_run_tasks":                                     266,
		"intersight_iaas_service_request":                                    267,
		"intersight_iaas_ucsd_info":                                          268,
		"intersight_iaas_ucsd_managed_infra":                                 269,
		"intersight_iaas_ucsd_messages":                                      270,
		"intersight_iam_account":                                             271,
		"intersight_iam_account_experience":                                  272,
		"intersight_iam_api_key":                                             273,
		"intersight_iam_app_registration":                                    274,
		"intersight_iam_banner_message":                                      275,
		"intersight_iam_certificate":                                         276,
		"intersight_iam_certificate_request":                                 277,
		"intersight_iam_domain_group":                                        278,
		"intersight_iam_end_point_privilege":                                 279,
		"intersight_iam_end_point_role":                                      280,
		"intersight_iam_end_point_user":                                      281,
		"intersight_iam_end_point_user_policy":                               282,
		"intersight_iam_end_point_user_role":                                 283,
		"intersight_iam_idp":                                                 284,
		"intersight_iam_idp_reference":                                       285,
		"intersight_iam_ip_access_management":                                286,
		"intersight_iam_ip_address":                                          287,
		"intersight_iam_ldap_group":                                          288,
		"intersight_iam_ldap_policy":                                         289,
		"intersight_iam_ldap_provider":                                       290,
		"intersight_iam_local_user_password_policy":                          291,
		"intersight_iam_o_auth_token":                                        292,
		"intersight_iam_permission":                                          293,
		"intersight_iam_private_key_spec":                                    294,
		"intersight_iam_privilege":                                           295,
		"intersight_iam_privilege_set":                                       296,
		"intersight_iam_qualifier":                                           297,
		"intersight_iam_resource_limits":                                     298,
		"intersight_iam_resource_permission":                                 299,
		"intersight_iam_resource_roles":                                      300,
		"intersight_iam_role":                                                301,
		"intersight_iam_security_holder":                                     302,
		"intersight_iam_service_provider":                                    303,
		"intersight_iam_session":                                             304,
		"intersight_iam_session_limits":                                      305,
		"intersight_iam_system":                                              306,
		"intersight_iam_trust_point":                                         307,
		"intersight_iam_user":                                                308,
		"intersight_iam_user_group":                                          309,
		"intersight_iam_user_preference":                                     310,
		"intersight_inventory_device_info":                                   311,
		"intersight_inventory_dn_mo_binding":                                 312,
		"intersight_inventory_generic_inventory":                             313,
		"intersight_inventory_generic_inventory_holder":                      314,
		"intersight_ipmioverlan_policy":                                      315,
		"intersight_ippool_block_lease":                                      316,
		"intersight_ippool_ip_lease":                                         317,
		"intersight_ippool_pool":                                             318,
		"intersight_ippool_pool_member":                                      319,
		"intersight_ippool_shadow_block":                                     320,
		"intersight_ippool_shadow_pool":                                      321,
		"intersight_ippool_universe":                                         322,
		"intersight_iqnpool_block":                                           323,
		"intersight_iqnpool_lease":                                           324,
		"intersight_iqnpool_pool":                                            325,
		"intersight_iqnpool_pool_member":                                     326,
		"intersight_iqnpool_universe":                                        327,
		"intersight_iwotenant_tenant_status":                                 328,
		"intersight_kubernetes_aci_cni_apic":                                 329,
		"intersight_kubernetes_aci_cni_profile":                              330,
		"intersight_kubernetes_aci_cni_tenant_cluster_allocation":            331,
		"intersight_kubernetes_addon":                                        332,
		"intersight_kubernetes_addon_definition":                             333,
		"intersight_kubernetes_addon_policy":                                 334,
		"intersight_kubernetes_addon_repository":                             335,
		"intersight_kubernetes_catalog":                                      336,
		"intersight_kubernetes_cluster":                                      337,
		"intersight_kubernetes_cluster_profile":                              338,
		"intersight_kubernetes_container_runtime_policy":                     339,
		"intersight_kubernetes_daemon_set":                                   340,
		"intersight_kubernetes_deployment":                                   341,
		"intersight_kubernetes_ingress":                                      342,
		"intersight_kubernetes_network_policy":                               343,
		"intersight_kubernetes_node":                                         344,
		"intersight_kubernetes_node_group_profile":                           345,
		"intersight_kubernetes_pod":                                          346,
		"intersight_kubernetes_service":                                      347,
		"intersight_kubernetes_stateful_set":                                 348,
		"intersight_kubernetes_sys_config_policy":                            349,
		"intersight_kubernetes_trusted_registries_policy":                    350,
		"intersight_kubernetes_version":                                      351,
		"intersight_kubernetes_version_policy":                               352,
		"intersight_kubernetes_virtual_machine_infrastructure_provider":      353,
		"intersight_kubernetes_virtual_machine_instance_type":                354,
		"intersight_kubernetes_virtual_machine_node_profile":                 355,
		"intersight_kvm_policy":                                              356,
		"intersight_kvm_session":                                             357,
		"intersight_kvm_tunnel":                                              358,
		"intersight_kvm_vm_console":                                          359,
		"intersight_license_account_license_data":                            360,
		"intersight_license_customer_op":                                     361,
		"intersight_license_iwo_customer_op":                                 362,
		"intersight_license_iwo_license_count":                               363,
		"intersight_license_license_info":                                    364,
		"intersight_license_license_reservation_op":                          365,
		"intersight_license_smartlicense_token":                              366,
		"intersight_ls_service_profile":                                      367,
		"intersight_macpool_id_block":                                        368,
		"intersight_macpool_lease":                                           369,
		"intersight_macpool_pool":                                            370,
		"intersight_macpool_pool_member":                                     371,
		"intersight_macpool_universe":                                        372,
		"intersight_management_controller":                                   373,
		"intersight_management_entity":                                       374,
		"intersight_management_interface":                                    375,
		"intersight_memory_array":                                            376,
		"intersight_memory_persistent_memory_config_result":                  377,
		"intersight_memory_persistent_memory_configuration":                  378,
		"intersight_memory_persistent_memory_namespace":                      379,
		"intersight_memory_persistent_memory_namespace_config_result":        380,
		"intersight_memory_persistent_memory_policy":                         381,
		"intersight_memory_persistent_memory_region":                         382,
		"intersight_memory_persistent_memory_unit":                           383,
		"intersight_memory_unit":                                             384,
		"intersight_meta_definition":                                         385,
		"intersight_network_element":                                         386,
		"intersight_network_element_summary":                                 387,
		"intersight_network_fc_zone_info":                                    388,
		"intersight_network_vlan_port_info":                                  389,
		"intersight_networkconfig_policy":                                    390,
		"intersight_niaapi_apic_cco_post":                                    391,
		"intersight_niaapi_apic_field_notice":                                392,
		"intersight_niaapi_apic_hweol":                                       393,
		"intersight_niaapi_apic_latest_maintained_release":                   394,
		"intersight_niaapi_apic_release_recommend":                           395,
		"intersight_niaapi_apic_sweol":                                       396,
		"intersight_niaapi_dcnm_cco_post":                                    397,
		"intersight_niaapi_dcnm_field_notice":                                398,
		"intersight_niaapi_dcnm_hweol":                                       399,
		"intersight_niaapi_dcnm_latest_maintained_release":                   400,
		"intersight_niaapi_dcnm_release_recommend":                           401,
		"intersight_niaapi_dcnm_sweol":                                       402,
		"intersight_niaapi_file_downloader":                                  403,
		"intersight_niaapi_nia_metadata":                                     404,
		"intersight_niaapi_nib_file_downloader":                              405,
		"intersight_niaapi_nib_metadata":                                     406,
		"intersight_niaapi_version_regex":                                    407,
		"intersight_niatelemetry_apic_fan_details":                           408,
		"intersight_niatelemetry_apic_fex_details":                           409,
		"intersight_niatelemetry_apic_psu_details":                           410,
		"intersight_niatelemetry_apic_ui_page_counts":                        411,
		"intersight_niatelemetry_app_details":                                412,
		"intersight_niatelemetry_epg":                                        413,
		"intersight_niatelemetry_fault":                                      414,
		"intersight_niatelemetry_lc":                                         415,
		"intersight_niatelemetry_mso_schema_details":                         416,
		"intersight_niatelemetry_mso_site_details":                           417,
		"intersight_niatelemetry_mso_tenant_details":                         418,
		"intersight_niatelemetry_nexus_dashboard_controller_details":         419,
		"intersight_niatelemetry_nexus_dashboard_details":                    420,
		"intersight_niatelemetry_nexus_dashboard_memory_details":             421,
		"intersight_niatelemetry_nexus_dashboards":                           422,
		"intersight_niatelemetry_nia_feature_usage":                          423,
		"intersight_niatelemetry_nia_inventory":                              424,
		"intersight_niatelemetry_nia_inventory_dcnm":                         425,
		"intersight_niatelemetry_nia_inventory_fabric":                       426,
		"intersight_niatelemetry_nia_license_state":                          427,
		"intersight_niatelemetry_tenant":                                     428,
		"intersight_ntp_policy":                                              429,
		"intersight_organization_organization":                               430,
		"intersight_os_catalog":                                              431,
		"intersight_os_configuration_file":                                   432,
		"intersight_os_distribution":                                         433,
		"intersight_os_install":                                              434,
		"intersight_os_supported_version":                                    435,
		"intersight_pci_coprocessor_card":                                    436,
		"intersight_pci_device":                                              437,
		"intersight_pci_link":                                                438,
		"intersight_pci_switch":                                              439,
		"intersight_port_group":                                              440,
		"intersight_port_mac_binding":                                        441,
		"intersight_port_sub_group":                                          442,
		"intersight_processor_unit":                                          443,
		"intersight_recommendation_capacity_runway":                          444,
		"intersight_recommendation_physical_item":                            445,
		"intersight_recovery_backup_config_policy":                           446,
		"intersight_recovery_backup_profile":                                 447,
		"intersight_recovery_config_result":                                  448,
		"intersight_recovery_config_result_entry":                            449,
		"intersight_recovery_on_demand_backup":                               450,
		"intersight_recovery_restore":                                        451,
		"intersight_recovery_schedule_config_policy":                         452,
		"intersight_resource_group":                                          453,
		"intersight_resource_group_member":                                   454,
		"intersight_resource_license_resource_count":                         455,
		"intersight_resource_membership":                                     456,
		"intersight_resource_membership_holder":                              457,
		"intersight_sdcard_policy":                                           458,
		"intersight_sdwan_profile":                                           459,
		"intersight_sdwan_router_node":                                       460,
		"intersight_sdwan_router_policy":                                     461,
		"intersight_sdwan_vmanage_account_policy":                            462,
		"intersight_search_search_item":                                      463,
		"intersight_search_tag_item":                                         464,
		"intersight_security_unit":                                           465,
		"intersight_server_config_change_detail":                             466,
		"intersight_server_config_import":                                    467,
		"intersight_server_config_result":                                    468,
		"intersight_server_config_result_entry":                              469,
		"intersight_server_profile":                                          470,
		"intersight_smtp_policy":                                             471,
		"intersight_snmp_policy":                                             472,
		"intersight_software_appliance_distributable":                        473,
		"intersight_software_hcl_meta":                                       474,
		"intersight_software_hyperflex_bundle_distributable":                 475,
		"intersight_software_hyperflex_distributable":                        476,
		"intersight_software_solution_distributable":                         477,
		"intersight_software_ucsd_bundle_distributable":                      478,
		"intersight_software_ucsd_distributable":                             479,
		"intersight_softwarerepository_authorization":                        480,
		"intersight_softwarerepository_cached_image":                         481,
		"intersight_softwarerepository_catalog":                              482,
		"intersight_softwarerepository_category_mapper":                      483,
		"intersight_softwarerepository_category_mapper_model":                484,
		"intersight_softwarerepository_category_support_constraint":          485,
		"intersight_softwarerepository_download_spec":                        486,
		"intersight_softwarerepository_operating_system_file":                487,
		"intersight_softwarerepository_release":                              488,
		"intersight_sol_policy":                                              489,
		"intersight_ssh_policy":                                              490,
		"intersight_storage_controller":                                      491,
		"intersight_storage_disk_group":                                      492,
		"intersight_storage_disk_group_policy":                               493,
		"intersight_storage_enclosure":                                       494,
		"intersight_storage_enclosure_disk":                                  495,
		"intersight_storage_enclosure_disk_slot_ep":                          496,
		"intersight_storage_flex_flash_controller":                           497,
		"intersight_storage_flex_flash_controller_props":                     498,
		"intersight_storage_flex_flash_physical_drive":                       499,
		"intersight_storage_flex_flash_virtual_drive":                        500,
		"intersight_storage_flex_util_controller":                            501,
		"intersight_storage_flex_util_physical_drive":                        502,
		"intersight_storage_flex_util_virtual_drive":                         503,
		"intersight_storage_hitachi_array":                                   504,
		"intersight_storage_hitachi_controller":                              505,
		"intersight_storage_hitachi_disk":                                    506,
		"intersight_storage_hitachi_host":                                    507,
		"intersight_storage_hitachi_host_lun":                                508,
		"intersight_storage_hitachi_parity_group":                            509,
		"intersight_storage_hitachi_pool":                                    510,
		"intersight_storage_hitachi_port":                                    511,
		"intersight_storage_hitachi_volume":                                  512,
		"intersight_storage_hyper_flex_storage_container":                    513,
		"intersight_storage_hyper_flex_volume":                               514,
		"intersight_storage_item":                                            515,
		"intersight_storage_net_app_aggregate":                               516,
		"intersight_storage_net_app_base_disk":                               517,
		"intersight_storage_net_app_cluster":                                 518,
		"intersight_storage_net_app_ethernet_port":                           519,
		"intersight_storage_net_app_export_policy":                           520,
		"intersight_storage_net_app_fc_interface":                            521,
		"intersight_storage_net_app_fc_port":                                 522,
		"intersight_storage_net_app_initiator_group":                         523,
		"intersight_storage_net_app_ip_interface":                            524,
		"intersight_storage_net_app_license":                                 525,
		"intersight_storage_net_app_lun":                                     526,
		"intersight_storage_net_app_lun_map":                                 527,
		"intersight_storage_net_app_node":                                    528,
		"intersight_storage_net_app_storage_vm":                              529,
		"intersight_storage_net_app_volume":                                  530,
		"intersight_storage_net_app_volume_snapshot":                         531,
		"intersight_storage_physical_disk":                                   532,
		"intersight_storage_physical_disk_extension":                         533,
		"intersight_storage_physical_disk_usage":                             534,
		"intersight_storage_pure_array":                                      535,
		"intersight_storage_pure_controller":                                 536,
		"intersight_storage_pure_disk":                                       537,
		"intersight_storage_pure_host":                                       538,
		"intersight_storage_pure_host_group":                                 539,
		"intersight_storage_pure_host_lun":                                   540,
		"intersight_storage_pure_port":                                       541,
		"intersight_storage_pure_protection_group":                           542,
		"intersight_storage_pure_protection_group_snapshot":                  543,
		"intersight_storage_pure_replication_schedule":                       544,
		"intersight_storage_pure_snapshot_schedule":                          545,
		"intersight_storage_pure_volume":                                     546,
		"intersight_storage_pure_volume_snapshot":                            547,
		"intersight_storage_sas_expander":                                    548,
		"intersight_storage_sas_port":                                        549,
		"intersight_storage_span":                                            550,
		"intersight_storage_storage_policy":                                  551,
		"intersight_storage_vd_member_ep":                                    552,
		"intersight_storage_virtual_drive":                                   553,
		"intersight_storage_virtual_drive_container":                         554,
		"intersight_storage_virtual_drive_extension":                         555,
		"intersight_syslog_policy":                                           556,
		"intersight_tam_advisory_count":                                      557,
		"intersight_tam_advisory_definition":                                 558,
		"intersight_tam_advisory_info":                                       559,
		"intersight_tam_advisory_instance":                                   560,
		"intersight_tam_security_advisory":                                   561,
		"intersight_techsupportmanagement_collection_control_policy":         562,
		"intersight_techsupportmanagement_download":                          563,
		"intersight_techsupportmanagement_tech_support_bundle":               564,
		"intersight_techsupportmanagement_tech_support_status":               565,
		"intersight_terminal_audit_log":                                      566,
		"intersight_top_system":                                              567,
		"intersight_ucsd_backup_info":                                        568,
		"intersight_uuidpool_block":                                          569,
		"intersight_uuidpool_pool":                                           570,
		"intersight_uuidpool_pool_member":                                    571,
		"intersight_uuidpool_universe":                                       572,
		"intersight_uuidpool_uuid_lease":                                     573,
		"intersight_virtualization_host":                                     574,
		"intersight_virtualization_virtual_disk":                             575,
		"intersight_virtualization_virtual_machine":                          576,
		"intersight_virtualization_vmware_cluster":                           577,
		"intersight_virtualization_vmware_datacenter":                        578,
		"intersight_virtualization_vmware_datastore":                         579,
		"intersight_virtualization_vmware_host":                              580,
		"intersight_virtualization_vmware_vcenter":                           581,
		"intersight_virtualization_vmware_virtual_machine":                   582,
		"intersight_vmedia_policy":                                           583,
		"intersight_vmrc_console":                                            584,
		"intersight_vnic_eth_adapter_policy":                                 585,
		"intersight_vnic_eth_if":                                             586,
		"intersight_vnic_eth_network_policy":                                 587,
		"intersight_vnic_eth_qos_policy":                                     588,
		"intersight_vnic_fc_adapter_policy":                                  589,
		"intersight_vnic_fc_if":                                              590,
		"intersight_vnic_fc_network_policy":                                  591,
		"intersight_vnic_fc_qos_policy":                                      592,
		"intersight_vnic_iscsi_adapter_policy":                               593,
		"intersight_vnic_iscsi_boot_policy":                                  594,
		"intersight_vnic_iscsi_static_target_policy":                         595,
		"intersight_vnic_lan_connectivity_policy":                            596,
		"intersight_vnic_lcp_status":                                         597,
		"intersight_vnic_san_connectivity_policy":                            598,
		"intersight_vnic_scp_status":                                         599,
		"intersight_vrf_vrf":                                                 600,
		"intersight_workflow_batch_api_executor":                             601,
		"intersight_workflow_build_task_meta":                                602,
		"intersight_workflow_build_task_meta_owner":                          603,
		"intersight_workflow_catalog":                                        604,
		"intersight_workflow_custom_data_type_definition":                    605,
		"intersight_workflow_error_response_handler":                         606,
		"intersight_workflow_pending_dynamic_workflow_info":                  607,
		"intersight_workflow_rollback_workflow":                              608,
		"intersight_workflow_task_definition":                                609,
		"intersight_workflow_task_info":                                      610,
		"intersight_workflow_task_metadata":                                  611,
		"intersight_workflow_workflow_definition":                            612,
		"intersight_workflow_workflow_info":                                  613,
		"intersight_workflow_workflow_meta":                                  614,
		"intersight_workflow_workflow_metadata":                              615,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
