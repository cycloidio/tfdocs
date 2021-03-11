package intersight

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "cl91",
					Description: `Use cl91 standard as FEC mode setting. 'Clause 91' aka RS-FEC ('ReedSolomon' FEC) offers better error protection against bursty and random errors but adds latency.`,
				},
				resource.Attribute{
					Name:        "cl74",
					Description: `Use cl74 standard as FEC mode setting. 'Clause 74' aka FC-FEC ('FireCode' FEC) offers simple, low-latency protection against 1 burst/sparse bit error, but it is not good for random errors.`,
				},
				resource.Attribute{
					Name:        "Off",
					Description: `Disable FEC mode on the DCE Interface. + ` + "`" + `interface_id` + "`" + `:(int) DCE interface id on which settings needs to be configured. Supported values are (0-3). + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `eth_settings` + "`" + `:(HashMap) - Global Ethernet settings for this adapter. This complex property has following sub-properties: + ` + "`" + `lldp_enabled` + "`" + `:(bool) Status of LLDP protocol on the adapter interfaces. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `fc_settings` + "`" + `:(HashMap) - Global Fibre Channel settings for this adapter. This complex property has following sub-properties: + ` + "`" + `fip_enabled` + "`" + `:(bool) Status of FIP protocol on the adapter interfaces. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `port_channel_settings` + "`" + `:(HashMap) - Port Channel settings for this adapter. This complex property has following sub-properties: + ` + "`" + `enabled` + "`" + `:(bool) When Port Channel is enabled, two vNICs and two vHBAs are available for use on the adapter card. When disabled, four vNICs and four vHBAs are available for use on the adapter card. Disabling port channel reboots the server. Port Channel is supported only for Cisco VIC 1455/1457 adapters. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `slot_id` + "`" + `:(string) PCIe slot where the VIC adapter is installed. Supported values are (1-15) and MLOM.`,
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
				resource.Attribute{
					Name:        "Pacific/Kiritimati",
					Description: `+ ` + "`" + `week_of_month` + "`" + `:(int) Schedule a task on a specific week of the month. Valid values are 1 through 5. Value of 5 means last week of the month. WeekOfMonth may not be set when dayOfMonth is specified.`,
				},
			},
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
					Description: `A Cisco Catalyst networking switch device. ## Import ` + "`" + `intersight_asset_target` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_asset_target.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ` ## Allowed Types in ` + "`" + `AdditionalProperties` + "`" + ` ### [asset.OrchestrationService](#argument-reference) OrchestrationService provides the necessary configuration details to enable Intersight Orchestration on the selected managed target. Subject to licensing. ### [asset.TerraformIntegrationService](#argument-reference) TerraformIntegrationService provides the necessary configuration details to enable Intersight Cloud Region on the selected Terraform Cloud. ### [asset.WorkloadOptimizerService](#argument-reference) WorkloadOptimizerService provides the necessary configuration details to enable Intersight Workflow Optimizer on the selected managed target. Subject to licensing. ### [asset.CloudConnection](#argument-reference) CloudConnection provides the necessary details for Intersight to connect to and authenticate with a target at a well-known service address. The service address is inferred based upon the target type. For example Amazon Web Services. ### [asset.HttpConnection](#argument-reference) HttpConnection provides the necessary details for Intersight to connect to and authenticate with a managed target over an HTTP connection.`,
				},
			},
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
					Description: `Value - enabled for configuring XptPrefetch token. ## Import ` + "`" + `intersight_bios_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_bios_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
				resource.Attribute{
					Name:        "name",
					Description: `Use interface name to select virtual ethernet interface.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `Use MAC address to select virtual ethernet interface.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `Use port to select virtual ethernet interface.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Default value if IpType is not specified.`,
				},
				resource.Attribute{
					Name:        "IPv4",
					Description: `The IPv4 address family type.`,
				},
				resource.Attribute{
					Name:        "IPv6",
					Description: `The IPv6 address family type.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `No sub type for SD card boot device.`,
				},
				resource.Attribute{
					Name:        "flex-util",
					Description: `Use of FlexUtil (microSD) card as sub type for SD card boot device.`,
				},
				resource.Attribute{
					Name:        "flex-flash",
					Description: `Use of FlexFlash (SD) card as sub type for SD card boot device.`,
				},
				resource.Attribute{
					Name:        "SDCARD",
					Description: `Use of SD card as sub type for the SD Card boot device. ### [boot.UefiShell](#argument-reference) Device type used when booting from UEFI shell. ### [boot.Usb](#argument-reference) Device type used when booting from USB device.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `No sub type for USB boot device.`,
				},
				resource.Attribute{
					Name:        "usb-cd",
					Description: `Use of Compact Disk (CD) as sub-type for the USB boot device.`,
				},
				resource.Attribute{
					Name:        "usb-fdd",
					Description: `Use of Floppy Disk Drive (FDD) as sub-type for the USB boot device.`,
				},
				resource.Attribute{
					Name:        "usb-hdd",
					Description: `Use of Hard Disk Drive (HDD) as sub-type for the USB boot device. ### [boot.VirtualMedia](#argument-reference) Device type used when booting from Virtual Media device.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `No sub type for virtual media.`,
				},
				resource.Attribute{
					Name:        "cimc-mapped-dvd",
					Description: `The virtual media device is mapped to a virtual DVD device.`,
				},
				resource.Attribute{
					Name:        "cimc-mapped-hdd",
					Description: `The virtual media device is mapped to a virtual HDD device.`,
				},
				resource.Attribute{
					Name:        "kvm-mapped-dvd",
					Description: `A KVM mapped DVD virtual media device.`,
				},
				resource.Attribute{
					Name:        "kvm-mapped-hdd",
					Description: `A KVM mapped HDD virtual media device.`,
				},
				resource.Attribute{
					Name:        "kvm-mapped-fdd",
					Description: `A KVM mapped FDD virtual media device.`,
				},
			},
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
				resource.Attribute{
					Name:        "end-host",
					Description: `In end-host mode, the fabric interconnects appear to the upstream devices as end hosts with multiple links.In this mode, the switch does not run Spanning Tree Protocol and avoids loops by following a set of rules for traffic forwarding.In case of ethernet switching mode - Ethernet end-host mode is also known as Ethernet host virtualizer.`,
				},
				resource.Attribute{
					Name:        "switch",
					Description: `In switch mode, the switch runs Spanning Tree Protocol to avoid loops, and broadcast and multicast packets are handled in the traditional way.This is the traditional switch mode. + ` + "`" + `vp_compression_supported` + "`" + `:(bool) VP Compression support on this switch.`,
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Import ` + "`" + `intersight_chassis_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_chassis_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The vethernet will remain up even if a suitable uplink is not pinned. ## Import ` + "`" + `intersight_fabric_eth_network_control_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_eth_network_control_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Admin configured Enabled State. ## Import ` + "`" + `intersight_fabric_fcoe_uplink_pc_role` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_fcoe_uplink_pc_role.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Admin configured Enabled State. ## Import ` + "`" + `intersight_fabric_fcoe_uplink_role` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_fcoe_uplink_role.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Import ` + "`" + `intersight_fabric_switch_cluster_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_switch_cluster_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Default",
					Description: `This option sets the default MAC address aging time to 14500 seconds for End Host mode.`,
				},
				resource.Attribute{
					Name:        "Custom",
					Description: `This option allows the the user to configure the MAC address aging time on the switch. For Switch Model UCS-FI-6454 or higher, the valid range is 120 to 918000 seconds and the switch will set the lower multiple of 5 of the given time.`,
				},
				resource.Attribute{
					Name:        "Never",
					Description: `This option disables the MAC address aging process and never allows the MAC address entries to get removed from the table. + ` + "`" + `mac_aging_time` + "`" + `:(int) Define the MAC address aging time in seconds. This field is valid when the \ Custom\ MAC address aging option is selected. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
				},
			},
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Import ` + "`" + `intersight_fabric_switch_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_switch_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Disabled",
					Description: `Admin configured Disabled State.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Admin configured Enabled State. + ` + "`" + `bandwidth_percent` + "`" + `:(int) Percentage of bandwidth received by the traffic tagged with this QoS. + ` + "`" + `cos` + "`" + `:(int) Class of service received by the traffic tagged with this QoS. + ` + "`" + `mtu` + "`" + `:(int) Maximum transmission unit (MTU) is the largest size packet or frame,that can be sent in a packet- or frame-based network such as the Internet.User can select from the following:1. Any value between 1500 and 92162. 'Normal' (default) mapping to a value of 1500.3. 'FC' mapping to a value of 2240. + ` + "`" + `multicast_optimize` + "`" + `:(bool) If enabled, this QoS class will be optimized to send multiple packets. + ` + "`" + `name` + "`" + `:(string) The 'name' of this QoS Class.`,
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
					Description: `QoS Priority for Bronze traffic. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `packet_drop` + "`" + `:(bool) If enabled, this QoS class will allow packet drops within an acceptable limit. + ` + "`" + `weight` + "`" + `:(int) The weight of the QoS Class controls the distribution of bandwidth between QoS Classes,with the same priority after the Guarantees for the QoS Classes are reached.`,
				},
			},
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
					Description: `Admin configured Enabled State. ## Import ` + "`" + `intersight_fabric_uplink_pc_role` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_uplink_pc_role.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Admin configured Enabled State. ## Import ` + "`" + `intersight_fabric_uplink_role` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_fabric_uplink_role.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "cisco",
					Description: `Image to be downloaded from Cisco software repository.`,
				},
				resource.Attribute{
					Name:        "localHttp",
					Description: `Image to be downloaded from a https server. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password as configured on the local https server. + ` + "`" + `upgradeoption` + "`" + `:(string) Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_mount_only",
					Description: `Direct upgrade SD upgrade mount only.`,
				},
				resource.Attribute{
					Name:        "sd_download_only",
					Description: `Direct upgrade SD download only.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_only",
					Description: `Direct upgrade SD upgrade only.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_full",
					Description: `Direct upgrade SD upgrade full.`,
				},
				resource.Attribute{
					Name:        "upgrade_full",
					Description: `The upgrade downloads or mounts the image, and reboots immediately for an upgrade.`,
				},
				resource.Attribute{
					Name:        "upgrade_mount_only",
					Description: `The upgrade downloads or mounts the image. The upgrade happens in next reboot.`,
				},
				resource.Attribute{
					Name:        "chassis_upgrade_full",
					Description: `Direct upgrade chassis upgrade full. + ` + "`" + `username` + "`" + `:(string) Username as configured on the local https server.`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `The default authentication protocol is decided by the endpoint.`,
				},
				resource.Attribute{
					Name:        "ntlm",
					Description: `The external CIFS server is configured with ntlm as the authentication protocol.`,
				},
				resource.Attribute{
					Name:        "ntlmi",
					Description: `Mount options of CIFS file server is ntlmi.`,
				},
				resource.Attribute{
					Name:        "ntlmv2",
					Description: `Mount options of CIFS file server is ntlmv2.`,
				},
				resource.Attribute{
					Name:        "ntlmv2i",
					Description: `Mount options of CIFS file server is ntlmv2i.`,
				},
				resource.Attribute{
					Name:        "ntlmssp",
					Description: `Mount options of CIFS file server is ntlmssp.`,
				},
				resource.Attribute{
					Name:        "ntlmsspi",
					Description: `Mount options of CIFS file server is ntlmsspi. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `remote_file` + "`" + `:(string)(Computed) Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso. + ` + "`" + `remote_ip` + "`" + `:(string)(Computed) CIFS Server Hostname or IP Address. For example:CIFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. + ` + "`" + `remote_share` + "`" + `:(string)(Computed) Directory where the image is stored. Example:share/subfolder. + ` + "`" + `http_server` + "`" + `:(HashMap) - HTTP (for WWW) file server option for network share upgrade. This complex property has following sub-properties: + ` + "`" + `location_link` + "`" + `:(string) HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. + ` + "`" + `mount_options` + "`" + `:(string) Mount option as configured on the HTTP server. Empty if nothing is configured. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `map_type` + "`" + `:(string) File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image.`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `File server protocol used is NFS.`,
				},
				resource.Attribute{
					Name:        "cifs",
					Description: `File server protocol used is CIFS.`,
				},
				resource.Attribute{
					Name:        "www",
					Description: `File server protocol used is WWW. + ` + "`" + `nfs_server` + "`" + `:(HashMap) - NFS file server option for network share upgrade. This complex property has following sub-properties: + ` + "`" + `file_location` + "`" + `:(string) The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. + ` + "`" + `mount_options` + "`" + `:(string) Mount option as configured on the NFS Server. For example:nolock. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `remote_file` + "`" + `:(string)(Computed) Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso. + ` + "`" + `remote_ip` + "`" + `:(string)(Computed) NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. + ` + "`" + `remote_share` + "`" + `:(string)(Computed) Directory where the image is stored. For example:/share/subfolder. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password as configured on the file server. + ` + "`" + `upgradeoption` + "`" + `:(string) Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade.`,
				},
				resource.Attribute{
					Name:        "nw_upgrade_full",
					Description: `Network upgrade option for full upgrade.`,
				},
				resource.Attribute{
					Name:        "nw_upgrade_mount_only",
					Description: `Network upgrade mount only. + ` + "`" + `username` + "`" + `:(string) Username as configured on the file server.`,
				},
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
					Description: `Upgrade mode is network upgrade. ## Import ` + "`" + `intersight_firmware_chassis_upgrade` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_firmware_chassis_upgrade.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "cisco",
					Description: `Image to be downloaded from Cisco software repository.`,
				},
				resource.Attribute{
					Name:        "localHttp",
					Description: `Image to be downloaded from a https server. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password as configured on the local https server. + ` + "`" + `upgradeoption` + "`" + `:(string) Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_mount_only",
					Description: `Direct upgrade SD upgrade mount only.`,
				},
				resource.Attribute{
					Name:        "sd_download_only",
					Description: `Direct upgrade SD download only.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_only",
					Description: `Direct upgrade SD upgrade only.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_full",
					Description: `Direct upgrade SD upgrade full.`,
				},
				resource.Attribute{
					Name:        "upgrade_full",
					Description: `The upgrade downloads or mounts the image, and reboots immediately for an upgrade.`,
				},
				resource.Attribute{
					Name:        "upgrade_mount_only",
					Description: `The upgrade downloads or mounts the image. The upgrade happens in next reboot.`,
				},
				resource.Attribute{
					Name:        "chassis_upgrade_full",
					Description: `Direct upgrade chassis upgrade full. + ` + "`" + `username` + "`" + `:(string) Username as configured on the local https server.`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `The default authentication protocol is decided by the endpoint.`,
				},
				resource.Attribute{
					Name:        "ntlm",
					Description: `The external CIFS server is configured with ntlm as the authentication protocol.`,
				},
				resource.Attribute{
					Name:        "ntlmi",
					Description: `Mount options of CIFS file server is ntlmi.`,
				},
				resource.Attribute{
					Name:        "ntlmv2",
					Description: `Mount options of CIFS file server is ntlmv2.`,
				},
				resource.Attribute{
					Name:        "ntlmv2i",
					Description: `Mount options of CIFS file server is ntlmv2i.`,
				},
				resource.Attribute{
					Name:        "ntlmssp",
					Description: `Mount options of CIFS file server is ntlmssp.`,
				},
				resource.Attribute{
					Name:        "ntlmsspi",
					Description: `Mount options of CIFS file server is ntlmsspi. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `remote_file` + "`" + `:(string)(Computed) Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso. + ` + "`" + `remote_ip` + "`" + `:(string)(Computed) CIFS Server Hostname or IP Address. For example:CIFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. + ` + "`" + `remote_share` + "`" + `:(string)(Computed) Directory where the image is stored. Example:share/subfolder. + ` + "`" + `http_server` + "`" + `:(HashMap) - HTTP (for WWW) file server option for network share upgrade. This complex property has following sub-properties: + ` + "`" + `location_link` + "`" + `:(string) HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. + ` + "`" + `mount_options` + "`" + `:(string) Mount option as configured on the HTTP server. Empty if nothing is configured. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `map_type` + "`" + `:(string) File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image.`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `File server protocol used is NFS.`,
				},
				resource.Attribute{
					Name:        "cifs",
					Description: `File server protocol used is CIFS.`,
				},
				resource.Attribute{
					Name:        "www",
					Description: `File server protocol used is WWW. + ` + "`" + `nfs_server` + "`" + `:(HashMap) - NFS file server option for network share upgrade. This complex property has following sub-properties: + ` + "`" + `file_location` + "`" + `:(string) The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. + ` + "`" + `mount_options` + "`" + `:(string) Mount option as configured on the NFS Server. For example:nolock. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `remote_file` + "`" + `:(string)(Computed) Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso. + ` + "`" + `remote_ip` + "`" + `:(string)(Computed) NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. + ` + "`" + `remote_share` + "`" + `:(string)(Computed) Directory where the image is stored. For example:/share/subfolder. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password as configured on the file server. + ` + "`" + `upgradeoption` + "`" + `:(string) Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade.`,
				},
				resource.Attribute{
					Name:        "nw_upgrade_full",
					Description: `Network upgrade option for full upgrade.`,
				},
				resource.Attribute{
					Name:        "nw_upgrade_mount_only",
					Description: `Network upgrade mount only. + ` + "`" + `username` + "`" + `:(string) Username as configured on the file server.`,
				},
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
					Description: `Upgrade mode is network upgrade. ## Import ` + "`" + `intersight_firmware_switch_upgrade` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_firmware_switch_upgrade.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "crc",
					Description: `A CRC hash as definded by RFC 3385. Generated with the IEEE polynomial.`,
				},
				resource.Attribute{
					Name:        "sha256",
					Description: `A SHA256 hash as defined by RFC 4634. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
				},
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
					Description: `Upgrade status is failed when upgrade has failed. ## Import ` + "`" + `intersight_firmware_unsupported_version_upgrade` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_firmware_unsupported_version_upgrade.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "cisco",
					Description: `Image to be downloaded from Cisco software repository.`,
				},
				resource.Attribute{
					Name:        "localHttp",
					Description: `Image to be downloaded from a https server. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password as configured on the local https server. + ` + "`" + `upgradeoption` + "`" + `:(string) Option to control the upgrade, e.g., sd_upgrade_mount_only - download the image into sd and upgrade wait for the server on-next boot.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_mount_only",
					Description: `Direct upgrade SD upgrade mount only.`,
				},
				resource.Attribute{
					Name:        "sd_download_only",
					Description: `Direct upgrade SD download only.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_only",
					Description: `Direct upgrade SD upgrade only.`,
				},
				resource.Attribute{
					Name:        "sd_upgrade_full",
					Description: `Direct upgrade SD upgrade full.`,
				},
				resource.Attribute{
					Name:        "upgrade_full",
					Description: `The upgrade downloads or mounts the image, and reboots immediately for an upgrade.`,
				},
				resource.Attribute{
					Name:        "upgrade_mount_only",
					Description: `The upgrade downloads or mounts the image. The upgrade happens in next reboot.`,
				},
				resource.Attribute{
					Name:        "chassis_upgrade_full",
					Description: `Direct upgrade chassis upgrade full. + ` + "`" + `username` + "`" + `:(string) Username as configured on the local https server.`,
				},
				resource.Attribute{
					Name:        "none",
					Description: `The default authentication protocol is decided by the endpoint.`,
				},
				resource.Attribute{
					Name:        "ntlm",
					Description: `The external CIFS server is configured with ntlm as the authentication protocol.`,
				},
				resource.Attribute{
					Name:        "ntlmi",
					Description: `Mount options of CIFS file server is ntlmi.`,
				},
				resource.Attribute{
					Name:        "ntlmv2",
					Description: `Mount options of CIFS file server is ntlmv2.`,
				},
				resource.Attribute{
					Name:        "ntlmv2i",
					Description: `Mount options of CIFS file server is ntlmv2i.`,
				},
				resource.Attribute{
					Name:        "ntlmssp",
					Description: `Mount options of CIFS file server is ntlmssp.`,
				},
				resource.Attribute{
					Name:        "ntlmsspi",
					Description: `Mount options of CIFS file server is ntlmsspi. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `remote_file` + "`" + `:(string)(Computed) Filename of the image in the remote share location. Example:ucs-c220m5-huu-3.1.2c.iso. + ` + "`" + `remote_ip` + "`" + `:(string)(Computed) CIFS Server Hostname or IP Address. For example:CIFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. + ` + "`" + `remote_share` + "`" + `:(string)(Computed) Directory where the image is stored. Example:share/subfolder. + ` + "`" + `http_server` + "`" + `:(HashMap) - HTTP (for WWW) file server option for network share upgrade. This complex property has following sub-properties: + ` + "`" + `location_link` + "`" + `:(string) HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices. + ` + "`" + `mount_options` + "`" + `:(string) Mount option as configured on the HTTP server. Empty if nothing is configured. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `map_type` + "`" + `:(string) File server protocols such as CIFS, NFS, WWW for HTTP (S) that hosts the image.`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `File server protocol used is NFS.`,
				},
				resource.Attribute{
					Name:        "cifs",
					Description: `File server protocol used is CIFS.`,
				},
				resource.Attribute{
					Name:        "www",
					Description: `File server protocol used is WWW. + ` + "`" + `nfs_server` + "`" + `:(HashMap) - NFS file server option for network share upgrade. This complex property has following sub-properties: + ` + "`" + `file_location` + "`" + `:(string) The location to the image file. The accepted format is IP-or-hostname/folder1/folder2/.../imageFile. + ` + "`" + `mount_options` + "`" + `:(string) Mount option as configured on the NFS Server. For example:nolock. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `remote_file` + "`" + `:(string)(Computed) Filename of the image in the remote share location. For example:ucs-c220m5-huu-3.1.2c.iso. + ` + "`" + `remote_ip` + "`" + `:(string)(Computed) NFS Server Hostname or IP Address. For example:NFS-server-hostname or 10.10.8.7. The remote share server should have network connectivity with the CIMC of the selected devices for a successful upgrade. + ` + "`" + `remote_share` + "`" + `:(string)(Computed) Directory where the image is stored. For example:/share/subfolder. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password as configured on the file server. + ` + "`" + `upgradeoption` + "`" + `:(string) Option to control the upgrade operation. Some examples, 1) nw_upgrade_mount_only - mount the image from a file server and run the upgrade on the next server boot and 2) nw_upgrade_full - mount the image and immediately run the upgrade.`,
				},
				resource.Attribute{
					Name:        "nw_upgrade_full",
					Description: `Network upgrade option for full upgrade.`,
				},
				resource.Attribute{
					Name:        "nw_upgrade_mount_only",
					Description: `Network upgrade mount only. + ` + "`" + `username` + "`" + `:(string) Username as configured on the file server.`,
				},
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
					Description: `Upgrade mode is network upgrade. ## Import ` + "`" + `intersight_firmware_upgrade` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_firmware_upgrade.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The hypervisor running on the HyperFlex cluster is not known. + ` + "`" + `mgmt_platform` + "`" + `:(string) The supported management platform for the HyperFlex Cluster.`,
				},
				resource.Attribute{
					Name:        "FI",
					Description: `The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.`,
				},
				resource.Attribute{
					Name:        "EDGE",
					Description: `The host servers used in the cluster deployment are standalone severs. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `server_model` + "`" + `:(string) The supported server models in regex format. + ` + "`" + `name` + "`" + `:(string) The application setting identifier. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) The application setting value.`,
				},
			},
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
					Description: `The hypervisor running on the HyperFlex cluster is not known. + ` + "`" + `mgmt_platform` + "`" + `:(string) The supported management platform for the HyperFlex Cluster.`,
				},
				resource.Attribute{
					Name:        "FI",
					Description: `The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.`,
				},
				resource.Attribute{
					Name:        "EDGE",
					Description: `The host servers used in the cluster deployment are standalone severs. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `server_model` + "`" + `:(string) The supported server models in regex format. + ` + "`" + `name` + "`" + `:(string) The application setting identifier. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) The application setting value.`,
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Import ` + "`" + `intersight_hyperflex_node_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_hyperflex_node_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The hypervisor running on the HyperFlex cluster is not known. + ` + "`" + `mgmt_platform` + "`" + `:(string) The supported management platform for the HyperFlex Cluster.`,
				},
				resource.Attribute{
					Name:        "FI",
					Description: `The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.`,
				},
				resource.Attribute{
					Name:        "EDGE",
					Description: `The host servers used in the cluster deployment are standalone severs. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `server_model` + "`" + `:(string) The supported server models in regex format.`,
				},
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
					Description: `The hypervisor running on the HyperFlex cluster is not known. + ` + "`" + `mgmt_platform` + "`" + `:(string) The supported management platform for the HyperFlex Cluster.`,
				},
				resource.Attribute{
					Name:        "FI",
					Description: `The host servers used in the cluster deployment are managed by a UCS Fabric Interconnect.`,
				},
				resource.Attribute{
					Name:        "EDGE",
					Description: `The host servers used in the cluster deployment are standalone severs. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `server_model` + "`" + `:(string) The supported server models in regex format. + ` + "`" + `name` + "`" + `:(string) The application setting identifier. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) The application setting value.`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "M5",
					Description: `M5 generation of UCS server.`,
				},
				resource.Attribute{
					Name:        "M4",
					Description: `M4 generation of UCS server. + ` + "`" + `nr_version` + "`" + `:(string) The server firmware bundle version.`,
				},
			},
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
					Description: `## Import ` + "`" + `intersight_hyperflex_sys_config_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_hyperflex_sys_config_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "IWO",
					Description: `Intersight Workflow Optimizer.`,
				},
				resource.Attribute{
					Name:        "Hitachi",
					Description: `Support to claim Hitachi Storage arrays using the Intersight Orchestrator framework.`,
				},
				resource.Attribute{
					Name:        "Kubernetes",
					Description: `Enables ability to create and manage Kubernetes clusters.`,
				},
				resource.Attribute{
					Name:        "NetAppIO",
					Description: `Support to claim NetApp Storage arrays as IO targets.`,
				},
				resource.Attribute{
					Name:        "IvsPublicCloud",
					Description: `Enables virtualization service for public clouds.`,
				},
				resource.Attribute{
					Name:        "TerraformCloud",
					Description: `Enables an ability to create Terraform Cloud. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
				},
			},
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
					Name:        "RSA",
					Description: `Key pairs should be generated by the RSA algorithm. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type.`,
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
				resource.Attribute{
					Name:        "P256",
					Description: `P256 returns a Curve which implements P-256, as defined in FIPS 186-4, section D.2.3.`,
				},
				resource.Attribute{
					Name:        "P224",
					Description: `P224 returns a Curve which implements P-224, as defined in FIPS 186-4, section D.2.2.`,
				},
				resource.Attribute{
					Name:        "P384",
					Description: `P384 returns a Curve which implements P-384, as defined in FIPS 186-4, section D.2.4.`,
				},
				resource.Attribute{
					Name:        "P521",
					Description: `P521 returns a Curve which implements P-521, as defined in FIPS 186-4, section D.2.5. ### [pkix.EddsaKeySpec](#argument-reference) The key pair is generated using Edwards-Curve Digital Signature Algorithm (EdDSA). The Edwards-curve Digital Signature Algorithm (EdDSA) is a variant of Schnorr's signature system with (possibly twisted) Edwards curves.`,
				},
				resource.Attribute{
					Name:        "Ed25519",
					Description: `The edwards25519 curve, as defined in RFC 8032 section 5.1.`,
				},
				resource.Attribute{
					Name:        "Ed25519ph",
					Description: `The edwards25519 curve for the PureEdDSA variant.`,
				},
				resource.Attribute{
					Name:        "Ed25519ctx",
					Description: `The edwards25519 curve for the HashEdDSA variant. ### [pkix.RsaAlgorithm](#argument-reference) The key pair is generated using the RSA algorithm and specified parameters.`,
				},
				resource.Attribute{
					Name:        "2048",
					Description: `A key length of 2048 bits.`,
				},
				resource.Attribute{
					Name:        "2560",
					Description: `A key length of 2560 bits.`,
				},
				resource.Attribute{
					Name:        "3072",
					Description: `A key length of 3072 bits.`,
				},
				resource.Attribute{
					Name:        "3584",
					Description: `A key length of 3584 bits.`,
				},
				resource.Attribute{
					Name:        "4096",
					Description: `A key length of 4096 bits.`,
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
					Name:        "LoginCredentials",
					Description: `Requires the user credentials. If the bind process fails, then user is denied access.`,
				},
				resource.Attribute{
					Name:        "Anonymous",
					Description: `Requires no username and password. If this option is selected and the LDAP server is configured for Anonymous logins, then the user gains access.`,
				},
				resource.Attribute{
					Name:        "ConfiguredCredentials",
					Description: `Requires a known set of credentials to be specified for the initial bind process. If the initial bind process succeeds, then the distinguished name (DN) of the user name is queried and re-used for the re-binding process. If the re-binding process fails, then the user is denied access. + ` + "`" + `domain` + "`" + `:(string) The IPv4 domain that all users must be in. + ` + "`" + `enable_encryption` + "`" + `:(bool) If enabled, the endpoint encrypts all information it sends to the LDAP server. + ` + "`" + `enable_group_authorization` + "`" + `:(bool) If enabled, user authorization is also done at the group level for LDAP users not in the local user database. + ` + "`" + `filter` + "`" + `:(string) Criteria to identify entries in search requests. + ` + "`" + `group_attribute` + "`" + `:(string) Groups to which an LDAP entry belongs. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `nested_group_search_depth` + "`" + `:(int) Search depth to look for a nested LDAP group in an LDAP group map. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) The password of the user for initial bind process. It can be any string that adheres to the following constraints. It can have character except spaces, tabs, line breaks. It cannot be more than 254 characters. + ` + "`" + `timeout` + "`" + `:(int) LDAP authentication timeout duration, in seconds.`,
				},
				resource.Attribute{
					Name:        "Extracted",
					Description: `The domain name extracted-domain from the login ID.`,
				},
				resource.Attribute{
					Name:        "Configured",
					Description: `The configured-search domain.`,
				},
				resource.Attribute{
					Name:        "ConfiguredExtracted",
					Description: `The domain name extracted from the login ID than the configured-search domain.`,
				},
				resource.Attribute{
					Name:        "LocalUserDb",
					Description: `Precedence is given to local user database while searching.`,
				},
				resource.Attribute{
					Name:        "LDAPUserDb",
					Description: `Precedence is given to LADP user database while searching. ## Usage Example ### Resource Creation ` + "`" + `` + "`" + `` + "`" + `hcl resource "intersight_iam_ldap_policy" "ldap1" { name = "ldap1" description = "test policy" enabled = true enable_dns = true user_search_precedence = "LocalUserDb" organization { object_type = "organization.Organization" moid = var.organization } base_properties { attribute = "CiscoAvPair" base_dn = "DC=QATCSLABTPI02,DC=cisco,DC=com" bind_dn = "CN=administrator,CN=Users,DC=QATCSLABTPI02,DC=cisco,DC=com" bind_method = "Anonymous" domain = "QATCSLABTPI02.cisco.com" enable_encryption = true enable_group_authorization = true filter = "sAMAccountName" group_attribute = "memberOf" nested_group_search_depth = 128 timeout = 180 } dns_parameters { nr_source = "Extracted" search_forest = "xyz" search_domain = "abc" } } ` + "`" + `` + "`" + `` + "`" + ` ## Import ` + "`" + `intersight_iam_ldap_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_iam_ldap_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "RSA",
					Description: `Key pairs should be generated by the RSA algorithm. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type.`,
				},
				resource.Attribute{
					Name:        "P256",
					Description: `P256 returns a Curve which implements P-256, as defined in FIPS 186-4, section D.2.3.`,
				},
				resource.Attribute{
					Name:        "P224",
					Description: `P224 returns a Curve which implements P-224, as defined in FIPS 186-4, section D.2.2.`,
				},
				resource.Attribute{
					Name:        "P384",
					Description: `P384 returns a Curve which implements P-384, as defined in FIPS 186-4, section D.2.4.`,
				},
				resource.Attribute{
					Name:        "P521",
					Description: `P521 returns a Curve which implements P-521, as defined in FIPS 186-4, section D.2.5. ### [pkix.EddsaKeySpec](#argument-reference) The key pair is generated using Edwards-Curve Digital Signature Algorithm (EdDSA). The Edwards-curve Digital Signature Algorithm (EdDSA) is a variant of Schnorr's signature system with (possibly twisted) Edwards curves.`,
				},
				resource.Attribute{
					Name:        "Ed25519",
					Description: `The edwards25519 curve, as defined in RFC 8032 section 5.1.`,
				},
				resource.Attribute{
					Name:        "Ed25519ph",
					Description: `The edwards25519 curve for the PureEdDSA variant.`,
				},
				resource.Attribute{
					Name:        "Ed25519ctx",
					Description: `The edwards25519 curve for the HashEdDSA variant. ### [pkix.RsaAlgorithm](#argument-reference) The key pair is generated using the RSA algorithm and specified parameters.`,
				},
				resource.Attribute{
					Name:        "2048",
					Description: `A key length of 2048 bits.`,
				},
				resource.Attribute{
					Name:        "2560",
					Description: `A key length of 2560 bits.`,
				},
				resource.Attribute{
					Name:        "3072",
					Description: `A key length of 3072 bits.`,
				},
				resource.Attribute{
					Name:        "3584",
					Description: `A key length of 3584 bits.`,
				},
				resource.Attribute{
					Name:        "4096",
					Description: `A key length of 4096 bits.`,
				},
			},
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
					Description: `Always remove older release and reinstall. ## Import ` + "`" + `intersight_kubernetes_addon` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_kubernetes_addon.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "None",
					Description: `A place holder for the default value.`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `Action triggered on the resource is still running.`,
				},
				resource.Attribute{
					Name:        "Success",
					Description: `Action triggered on the resource is completed successfully.`,
				},
				resource.Attribute{
					Name:        "Failure",
					Description: `Action triggered on the resource is failed.`,
				},
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Import ` + "`" + `intersight_kubernetes_node_group_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_kubernetes_node_group_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
					Description: `## Import ` + "`" + `intersight_kubernetes_sys_config_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_kubernetes_sys_config_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "app-direct",
					Description: `The App Direct interleaved Persistent Memory type.`,
				},
				resource.Attribute{
					Name:        "app-direct-non-interleaved",
					Description: `The App Direct non-interleaved Persistent Memory type. + ` + "`" + `socket_id` + "`" + `:(string) CPU Socket ID to which this goal will be applied.`,
				},
				resource.Attribute{
					Name:        "All Sockets",
					Description: `All the CPU socket IDs in a server.`,
				},
				resource.Attribute{
					Name:        "raw",
					Description: `The raw mode of Persistent Memory Namespace.`,
				},
				resource.Attribute{
					Name:        "block",
					Description: `The block mode of Persistent Memory Namespace. + ` + "`" + `name` + "`" + `:(string) Name of this Namespace to be created on the server. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `socket_id` + "`" + `:(int) Socket ID of the region on which this Namespace has to be created or modified.`,
				},
				resource.Attribute{
					Name:        "1",
					Description: `The first CPU socket in a server.`,
				},
				resource.Attribute{
					Name:        "2",
					Description: `The second CPU socket in a server.`,
				},
				resource.Attribute{
					Name:        "3",
					Description: `The third CPU socket in a server.`,
				},
				resource.Attribute{
					Name:        "4",
					Description: `The fourth CPU socket in a server. + ` + "`" + `socket_memory_id` + "`" + `:(string) Socket Memory ID of the region on which this Namespace has to be created or modified.`,
				},
				resource.Attribute{
					Name:        "Not Applicable",
					Description: `The socket memory ID is not applicable if app-direct persistent memory type is selected in the goal.`,
				},
				resource.Attribute{
					Name:        "2",
					Description: `The second socket memory ID within a socket in a server.`,
				},
				resource.Attribute{
					Name:        "4",
					Description: `The fourth socket memory ID within a socket in a server.`,
				},
				resource.Attribute{
					Name:        "6",
					Description: `The sixth socket memory ID within a socket in a server.`,
				},
				resource.Attribute{
					Name:        "8",
					Description: `The eighth socket memory ID within a socket in a server.`,
				},
				resource.Attribute{
					Name:        "10",
					Description: `The tenth socket memory ID within a socket in a server.`,
				},
				resource.Attribute{
					Name:        "12",
					Description: `The twelfth socket memory ID within a socket in a server.`,
				},
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
					Description: `## Usage Example ### Resource Creation ` + "`" + `` + "`" + `` + "`" + `hcl resource "intersight_ntp_policy" "ntp1" { name = "ntp1" enabled = true ntp_servers = [ "ntp.esl.cisco.com", "time-a-g.nist.gov", "time-b-g.nist.gov" ] organization { object_type = "organization.Organization" moid = var.organization } } ` + "`" + `` + "`" + `` + "`" + ` ## Import ` + "`" + `intersight_ntp_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_ntp_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `properties` + "`" + `:(HashMap) - Primitive data type properties. This complex property has following sub-properties: + ` + "`" + `constraints` + "`" + `:(HashMap) - Constraints that must be applied to the parameter value supplied for this data type. This complex property has following sub-properties: + ` + "`" + `enum_list` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `label` + "`" + `:(string) Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_). + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_). + ` + "`" + `max` + "`" + `:(float) Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `Enum to specify a string data type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `Enum to specify an integer32 data type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `Enum to specify a float64 data type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `Enum to specify a boolean data type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Enum to specify a json data type.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `Enum to specify a enum data type which is a list of pre-defined strings. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. + ` + "`" + `value` + "`" + `: Value for placeholder provided by user.`,
				},
			},
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
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `properties` + "`" + `:(HashMap) - Primitive data type properties. This complex property has following sub-properties: + ` + "`" + `constraints` + "`" + `:(HashMap) - Constraints that must be applied to the parameter value supplied for this data type. This complex property has following sub-properties: + ` + "`" + `enum_list` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `label` + "`" + `:(string) Label for the enum value. A user friendly short string to identify the enum value. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), single quote ('), forward slash (/), or an underscore (_). + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) Enum value for this enum entry. Value will be passed to the workflow as string type for execution. Value can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), period (.), colon (:), space ( ), forward slash (/), or an underscore (_). + ` + "`" + `max` + "`" + `:(float) Allowed maximum value of the parameter if parameter is integer/float or maximum length of the parameter if the parameter is string. When max and min are set to 0, then the limits are not checked. The maximum number supported is 1.797693134862315708145274237317043567981e+308 or (2`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `Enum to specify a string data type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `Enum to specify an integer32 data type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `Enum to specify a float64 data type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `Enum to specify a boolean data type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Enum to specify a json data type.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `Enum to specify a enum data type which is a list of pre-defined strings. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. + ` + "`" + `value` + "`" + `: Value for placeholder provided by user.`,
				},
				resource.Attribute{
					Name:        "static",
					Description: `In case of static IP configuraton, provide the details such as IP address, netmask, and gateway.`,
				},
				resource.Attribute{
					Name:        "DHCP",
					Description: `In case of dynamic IP configuration, the IP address, netmask and gateway detailsare obtained from DHCP. + ` + "`" + `ip_configuration` + "`" + `:(HashMap) - In case of static IP configuration, IP address, netmask and gateway details areprovided. This complex property has following sub-properties: + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `is_answer_file_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'answerFile' property has been set. + ` + "`" + `is_root_password_crypted` + "`" + `:(bool) Enable to indicate Root Password provided is encrypted. + ` + "`" + `is_root_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'rootPassword' property has been set. + ` + "`" + `nameserver` + "`" + `:(string) IP address of the name server to be configured in the OS. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `product_key` + "`" + `:(string) The product key to be used for a specific version of Windows installation. + ` + "`" + `root_password` + "`" + `:(string) Password configured for the root / administrator user in the OS. You can enter a plain text or an encrypted password.Intersight encrypts the plaintext password. Enable the Encrypted Password option to provide an encrypted password.For more details on encrypting passwords, see Help Center. + ` + "`" + `nr_source` + "`" + `:(string) Answer values can be provided from three sources - Embedded in OS image, static file,or as placeholder values for an answer file template.Source of the answers is given as value, Embedded/File/Template.'Embedded' option indicates that the answer file is embedded within the OS Image. 'File'option indicates that the answers are provided as a file. 'Template' indicates that theplaceholders in the selected os.ConfigurationFile MO are replaced with values providedas os.Answers MO.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that answers is not sent and values must be populated from Install Template.`,
				},
				resource.Attribute{
					Name:        "Embedded",
					Description: `Indicates that the answer file is embedded within OS image.`,
				},
				resource.Attribute{
					Name:        "File",
					Description: `Indicates that the answer file is a static content that has all thevalues populated.`,
				},
				resource.Attribute{
					Name:        "Template",
					Description: `Indicates that the given answers are used to populate the answer filetemplate. The template allows the users to refer some server specificanswers as fields/placeholders and replace these placeholders with theactual values for each Server during OS installation using 'Answers' and'AdditionalParameters' properties in os.Install MO.The answer file templates can be created by users as os.ConfigurationFile objects.`,
				},
				resource.Attribute{
					Name:        "vMedia",
					Description: `OS image is mounted as vMedia in target server for OS installation.`,
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Import ` + "`" + `intersight_recovery_backup_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_recovery_backup_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Daily",
					Description: `Allows the user to run the backup daily at a given time.`,
				},
				resource.Attribute{
					Name:        "Periodic",
					Description: `Allows the user to run the backup after a certain number of hours. + ` + "`" + `hours` + "`" + `:(int) The frequency, in hours, at which the backup schedule runs.`,
				},
				resource.Attribute{
					Name:        "20",
					Description: `+ ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "OS",
					Description: `This partition contains virtual drives where user can install operating system.`,
				},
				resource.Attribute{
					Name:        "Utility",
					Description: `This partition contains virtual drives for utilities such as SCU, HUU, Drivers and Diagnostics. + ` + "`" + `virtual_drives` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `additional_properties` + "`" + `:(JSON) - Additional Properties as per object type, can be added as JSON using ` + "`" + `jsonencode()` + "`" + `. Allowed Types are: [sdcard.Diagnostics](#sdcardDiagnostics) [sdcard.Drivers](#sdcardDrivers) [sdcard.HostUpgradeUtility](#sdcardHostUpgradeUtility) [sdcard.OperatingSystem](#sdcardOperatingSystem) [sdcard.ServerConfigurationUtility](#sdcardServerConfigurationUtility) [sdcard.UserPartition](#sdcardUserPartition) + ` + "`" + `enable` + "`" + `:(bool) Enable the respective virtual drive to be available to the host. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type.`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "WAN",
					Description: `Port-group being added is used for WAN traffic.`,
				},
				resource.Attribute{
					Name:        "LAN",
					Description: `Port-group being added is used for LAN traffic.`,
				},
				resource.Attribute{
					Name:        "Management",
					Description: `Port-group being added is used for Management traffic. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `port_group` + "`" + `:(string) Name of the Port Group to create. + ` + "`" + `vlan` + "`" + `:(int) VLAN to be added to the Port Group.`,
				},
			},
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
					Description: `Dually terminated WANs are configured on all the SD-WAN routers. A given WAN connection is available on multiple router nodes. ## Import ` + "`" + `intersight_sdwan_router_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_sdwan_router_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `The profile defines the configuration for a specific instance of a target. ## Usage Example ### Resource Creation ` + "`" + `` + "`" + `` + "`" + `hcl resource "intersight_server_profile" "server1" { name = "server1" action = "No-op" tags { key = "server" value = "demo" } organization { object_type = "organization.Organization" moid = var.organization } } ` + "`" + `` + "`" + `` + "`" + ` ## Import ` + "`" + `intersight_server_profile` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_server_profile.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
				resource.Attribute{
					Name:        "Trap",
					Description: `Do not receive notifications when trap is sent to the destination.`,
				},
				resource.Attribute{
					Name:        "Inform",
					Description: `Receive notifications when trap is sent to the destination. This option is valid only for V2 users. + ` + "`" + `user` + "`" + `:(string) SNMP user for the trap. Applicable only to SNMPv3. + ` + "`" + `nr_version` + "`" + `:(string) SNMP version used for the trap.`,
				},
				resource.Attribute{
					Name:        "V3",
					Description: `SNMP v3 trap version notifications.`,
				},
				resource.Attribute{
					Name:        "V2",
					Description: `SNMP v2 trap version notifications.`,
				},
				resource.Attribute{
					Name:        "NA",
					Description: `Authentication protocol is not applicable.`,
				},
				resource.Attribute{
					Name:        "MD5",
					Description: `MD5 protocol is used to authenticate SNMP user.`,
				},
				resource.Attribute{
					Name:        "SHA",
					Description: `SHA protocol is used to authenticate SNMP user. + ` + "`" + `is_auth_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'authPassword' property has been set. + ` + "`" + `is_privacy_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'privacyPassword' property has been set. + ` + "`" + `name` + "`" + `:(string) SNMP username. Must have a minimum of 1 and and a maximum of 31 characters. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `privacy_password` + "`" + `:(string) Privacy password for the user. + ` + "`" + `privacy_type` + "`" + `:(string) Privacy protocol for the user.`,
				},
				resource.Attribute{
					Name:        "NA",
					Description: `Privacy protocol is not applicable.`,
				},
				resource.Attribute{
					Name:        "DES",
					Description: `DES privacy protocol is used for SNMP user.`,
				},
				resource.Attribute{
					Name:        "AES",
					Description: `AES privacy protocol is used for SNMP user. + ` + "`" + `security_level` + "`" + `:(string) Security mechanism used for communication between agent and manager.`,
				},
				resource.Attribute{
					Name:        "AuthPriv",
					Description: `The user requires both an authorization password and a privacy password.`,
				},
				resource.Attribute{
					Name:        "NoAuthNoPriv",
					Description: `The user does not require an authorization or privacy password.`,
				},
				resource.Attribute{
					Name:        "AuthNoPriv",
					Description: `The user requires an authorization password but not a privacy password.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
				},
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
					Name:        "ALL",
					Description: `This represents all the components.`,
				},
				resource.Attribute{
					Name:        "ALL,HDD",
					Description: `This represents all the components plus the HDDs.`,
				},
				resource.Attribute{
					Name:        "Drive-U.2",
					Description: `This represents the U.2 drives that are SFF/LFF drives (mostly all the drives will fall under this category).`,
				},
				resource.Attribute{
					Name:        "Storage",
					Description: `This represents the storage controller components.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `This represents none of the components.`,
				},
				resource.Attribute{
					Name:        "NXOS",
					Description: `This represents NXOS components.`,
				},
				resource.Attribute{
					Name:        "IOM",
					Description: `This represents IOM components.`,
				},
				resource.Attribute{
					Name:        "PSU",
					Description: `This represents PSU components.`,
				},
				resource.Attribute{
					Name:        "CIMC",
					Description: `This represents CIMC components.`,
				},
				resource.Attribute{
					Name:        "BIOS",
					Description: `This represents BIOS components.`,
				},
				resource.Attribute{
					Name:        "PCIE",
					Description: `This represents PCIE components.`,
				},
				resource.Attribute{
					Name:        "Drive",
					Description: `This represents Drive components.`,
				},
				resource.Attribute{
					Name:        "DIMM",
					Description: `This represents DIMM components.`,
				},
				resource.Attribute{
					Name:        "BoardController",
					Description: `This represents Board Controller components.`,
				},
				resource.Attribute{
					Name:        "StorageController",
					Description: `This represents Storage Controller components.`,
				},
				resource.Attribute{
					Name:        "HBA",
					Description: `This represents HBA components.`,
				},
				resource.Attribute{
					Name:        "GPU",
					Description: `This represents GPU components.`,
				},
				resource.Attribute{
					Name:        "SasExpander",
					Description: `This represents SasExpander components.`,
				},
				resource.Attribute{
					Name:        "MSwitch",
					Description: `This represents mSwitch components.`,
				},
				resource.Attribute{
					Name:        "CMC",
					Description: `This represents CMC components. + ` + "`" + `description` + "`" + `:(string) This shows the description of component image within the distributable. + ` + "`" + `disruption` + "`" + `:(string) The type of disruption on each component. For example, host reboot, automatic power cycle, and manual power cycle.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Indicates that the component did not receive a disruption request.`,
				},
				resource.Attribute{
					Name:        "HostReboot",
					Description: `Indicates that the component received a host reboot request.`,
				},
				resource.Attribute{
					Name:        "EndpointReboot",
					Description: `Indicates that the component received an end point reboot request.`,
				},
				resource.Attribute{
					Name:        "ManualPowerCycle",
					Description: `Indicates that the component received a manual power cycle request.`,
				},
				resource.Attribute{
					Name:        "AutomaticPowerCycle",
					Description: `Indicates that the component received an automatic power cycle request. + ` + "`" + `image_path` + "`" + `:(string) This shows the path of component image within the distributable. + ` + "`" + `is_oob_supported` + "`" + `:(bool) If set, the component can be updated through out-of-band management, else, is updated through host service utility boot. + ` + "`" + `model` + "`" + `:(string) The model of the component image in the distributable. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `oob_manageability` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `packed_version` + "`" + `:(string) The image version of components packaged in the distributable. + ` + "`" + `redfish_url` + "`" + `:(string) The redfish target for each component. + ` + "`" + `vendor` + "`" + `:(string) The version of component image in the distributable.`,
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
				resource.Attribute{
					Name:        "Default",
					Description: `Use platform default access mode.`,
				},
				resource.Attribute{
					Name:        "ReadWrite",
					Description: `Enables host to perform read-write on the VD.`,
				},
				resource.Attribute{
					Name:        "ReadOnly",
					Description: `Host can only read from the VD.`,
				},
				resource.Attribute{
					Name:        "Blocked",
					Description: `Host can neither read nor write to the VD. + ` + "`" + `boot_drive` + "`" + `:(bool) The flag enables the use of this virtual drive as a boot drive. + ` + "`" + `disk_group_name` + "`" + `:(string)(Computed) Disk group policy that has the disk group in which this virtual drive needs to be created. + ` + "`" + `disk_group_policy` + "`" + `:(string) Disk group policy that has the disk group in which this virtual drive needs to be created. + ` + "`" + `drive_cache` + "`" + `:(string) Drive Cache property expect disk cache policy.`,
				},
				resource.Attribute{
					Name:        "Default",
					Description: `Use platform default drive cache mode.`,
				},
				resource.Attribute{
					Name:        "NoChange",
					Description: `Drive cache policy is unchanged.`,
				},
				resource.Attribute{
					Name:        "Enable",
					Description: `Enables IO caching on the drive.`,
				},
				resource.Attribute{
					Name:        "Disable",
					Description: `Disables IO caching on the drive. + ` + "`" + `expand_to_available` + "`" + `:(bool) The flag enables this virtual drive to use all the available space in the disk group. When this flag is configured, the size property is ignored. + ` + "`" + `io_policy` + "`" + `:(string) Desired IO mode - direct IO or cached IO.`,
				},
				resource.Attribute{
					Name:        "Default",
					Description: `Use platform default IO mode.`,
				},
				resource.Attribute{
					Name:        "Direct",
					Description: `Use direct IO for writing directly into the disk.`,
				},
				resource.Attribute{
					Name:        "Cached",
					Description: `Use cached IO for writing into cache and then to disk. + ` + "`" + `name` + "`" + `:(string) The name of the virtual drive. The name can be between 1 and 15 alphanumeric characters. Spaces or any special characters other than - (hyphen), _ (underscore), : (colon), and . (period) are not allowed. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `read_policy` + "`" + `:(string) Read ahead mode to be used to read data from this virtual drive.`,
				},
				resource.Attribute{
					Name:        "Default",
					Description: `Use platform default read ahead mode.`,
				},
				resource.Attribute{
					Name:        "ReadAhead",
					Description: `Use read ahead mode for the policy.`,
				},
				resource.Attribute{
					Name:        "NoReadAhead",
					Description: `Do not use read ahead mode for the policy. + ` + "`" + `size` + "`" + `:(int) Virtual drive size in MB. Size is mandatory field unless the 'Expand to Available' option is enabled. + ` + "`" + `strip_size` + "`" + `:(string) The strip size is the portion of a stripe that resides on a single drive in the drive group. The stripe consists of the data segments that the RAID controller writes across multiple drives, not including parity drives.`,
				},
				resource.Attribute{
					Name:        "Default",
					Description: `Use platform default strip size for a virtual drive.`,
				},
				resource.Attribute{
					Name:        "32k",
					Description: `Enables a strip size of 32k for a virtual drive.`,
				},
				resource.Attribute{
					Name:        "64k",
					Description: `Enables a strip size of 64k for a virtual drive.`,
				},
				resource.Attribute{
					Name:        "128k",
					Description: `Enables a strip size of 128k for a virtual drive.`,
				},
				resource.Attribute{
					Name:        "256k",
					Description: `Enables a strip size of 256k for a virtual drive.`,
				},
				resource.Attribute{
					Name:        "512k",
					Description: `Enables a strip size of 512k for a virtual drive.`,
				},
				resource.Attribute{
					Name:        "1024k",
					Description: `Enables a strip size of 1024k for a virtual drive. + ` + "`" + `vdid` + "`" + `:(string)(Computed) Unique Id of the Virtual Drive to be used to identify Virtual Drive when name is empty. + ` + "`" + `write_policy` + "`" + `:(string) Write mode to be used to write data to this virtual drive.`,
				},
				resource.Attribute{
					Name:        "Default",
					Description: `Use platform default write mode.`,
				},
				resource.Attribute{
					Name:        "WriteThrough",
					Description: `Data is written through the cache and to the physical drives. Performance is improved, because subsequent reads of that data can be satisfied from the cache.`,
				},
				resource.Attribute{
					Name:        "WriteBackGoodBbu",
					Description: `Data is stored in the cache, and is only written to the physical drives when space in the cache is needed. Virtual drives requesting this policy fall back to Write Through caching when the battery backup unit (BBU) cannot guarantee the safety of the cache in the event of a power failure.`,
				},
				resource.Attribute{
					Name:        "AlwaysWriteBack",
					Description: `With this policy, write caching remains Write Back even if the battery backup unit is defective or discharged. ## Usage Example ### Resource Creation ` + "`" + `` + "`" + `` + "`" + `hcl resource "intersight_storage_storage_policy" "storage_storage1" { name = "storage_storage_policy1" description = "storage policy test" retain_policy_virtual_drives = true unused_disks_state = "UnconfiguredGood" virtual_drives { object_type = "storage.VirtualDriveConfig" boot_drive = true drive_cache = "NoChange" expand_to_available = false io_policy = "Direct" name = "RAID0_1" access_policy = "ReadWrite" disk_group_policy = intersight_storage_disk_group_policy.storage_disk_group1.id read_policy = "NoReadAhead" size = 285148 write_policy = "WriteThrough" } organization { object_type = "organization.Organization" moid = var.organization } } ` + "`" + `` + "`" + `` + "`" + ` ## Import ` + "`" + `intersight_storage_storage_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_storage_storage_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "warning",
					Description: `Use logging level warning for logs classified as warning.`,
				},
				resource.Attribute{
					Name:        "emergency",
					Description: `Use logging level emergency for logs classified as emergency.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Use logging level alert for logs classified as alert.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `Use logging level critical for logs classified as critical.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Use logging level error for logs classified as error.`,
				},
				resource.Attribute{
					Name:        "notice",
					Description: `Use logging level notice for logs classified as notice.`,
				},
				resource.Attribute{
					Name:        "informational",
					Description: `Use logging level informational for logs classified as informational.`,
				},
				resource.Attribute{
					Name:        "debug",
					Description: `Use logging level debug for logs classified as debug. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type.`,
				},
				resource.Attribute{
					Name:        "warning",
					Description: `Use logging level warning for logs classified as warning.`,
				},
				resource.Attribute{
					Name:        "emergency",
					Description: `Use logging level emergency for logs classified as emergency.`,
				},
				resource.Attribute{
					Name:        "alert",
					Description: `Use logging level alert for logs classified as alert.`,
				},
				resource.Attribute{
					Name:        "critical",
					Description: `Use logging level critical for logs classified as critical.`,
				},
				resource.Attribute{
					Name:        "error",
					Description: `Use logging level error for logs classified as error.`,
				},
				resource.Attribute{
					Name:        "notice",
					Description: `Use logging level notice for logs classified as notice.`,
				},
				resource.Attribute{
					Name:        "informational",
					Description: `Use logging level informational for logs classified as informational.`,
				},
				resource.Attribute{
					Name:        "debug",
					Description: `Use logging level debug for logs classified as debug. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `port` + "`" + `:(int) Port number used for logging on syslog server. + ` + "`" + `protocol` + "`" + `:(string) Transport layer protocol for transmission of log messages to syslog server.`,
				},
				resource.Attribute{
					Name:        "udp",
					Description: `Use User Datagram Protocol (UDP) for syslog remote server connection.`,
				},
				resource.Attribute{
					Name:        "tcp",
					Description: `Use Transmission Control Protocol (TCP) for syslog remote server connection.`,
				},
			},
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
					Name:        "psirt",
					Description: `Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).`,
				},
				resource.Attribute{
					Name:        "fieldNotice",
					Description: `Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). + ` + "`" + `identifiers` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `name` + "`" + `:(string) Name of the filter paramter. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) Value of the filter paramter. + ` + "`" + `name` + "`" + `:(string) Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `operation_type` + "`" + `:(string) Operation type for the alert action. An action is used to carry out the process of \ reacting\ to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Create an instance of AdvisoryInstance.`,
				},
				resource.Attribute{
					Name:        "remove",
					Description: `Remove an instance of AdvisoryInstance. + ` + "`" + `queries` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `name` + "`" + `:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `priority` + "`" + `:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. + ` + "`" + `query` + "`" + `:(string) A SparkSQL query to be used on a given data source. + ` + "`" + `type` + "`" + `:(string) Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.`,
				},
				resource.Attribute{
					Name:        "restApi",
					Description: `Repesents the use of REST API for carrying out alert actions.`,
				},
				resource.Attribute{
					Name:        "nxos",
					Description: `Collector type for this data collection is NXOS.`,
				},
				resource.Attribute{
					Name:        "intersightApi",
					Description: `Collector type for this data collection is Intersight APIs.`,
				},
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
					Name:        "psirt",
					Description: `Respresents the psirt alert type (https://tools.cisco.com/security/center/publicationListing.x).`,
				},
				resource.Attribute{
					Name:        "fieldNotice",
					Description: `Respresents the field notice alert type (https://www.cisco.com/c/en/us/support/web/tsd-products-field-notice-summary.html). + ` + "`" + `identifiers` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `name` + "`" + `:(string) Name of the filter paramter. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) Value of the filter paramter. + ` + "`" + `name` + "`" + `:(string) Uniquely identifies a given action among the set of actions corresponding to an advisory. Primarily used to store and compare results of subsequent iterations corresponding to the action queries. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `operation_type` + "`" + `:(string) Operation type for the alert action. An action is used to carry out the process of \ reacting\ to an alert condition. For e.g.in case of a fieldNotice alert, the intention may be to create a new alert (if the condition matches and there is no existing alert) or to remove an existing alert when the alert condition has been remedied.`,
				},
				resource.Attribute{
					Name:        "create",
					Description: `Create an instance of AdvisoryInstance.`,
				},
				resource.Attribute{
					Name:        "remove",
					Description: `Remove an instance of AdvisoryInstance. + ` + "`" + `queries` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `name` + "`" + `:(string) Name is used to unique identify and result of the given query which can be used by subsequent queries as input data source. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `priority` + "`" + `:(int) An integer value depicting the priority of the query among the queries that are part of the same QueryEntry collection. + ` + "`" + `query` + "`" + `:(string) A SparkSQL query to be used on a given data source. + ` + "`" + `type` + "`" + `:(string) Type of Intersight alert. An alert in Intersight could be one of several kinds (FieldNotice, PSIRT etc.). Primarily used for filtering alerts based on the type.`,
				},
				resource.Attribute{
					Name:        "restApi",
					Description: `Repesents the use of REST API for carrying out alert actions.`,
				},
				resource.Attribute{
					Name:        "nxos",
					Description: `Collector type for this data collection is NXOS.`,
				},
				resource.Attribute{
					Name:        "intersightApi",
					Description: `Collector type for this data collection is Intersight APIs.`,
				},
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
					Description: `Disable techsupport collection. ## Import ` + "`" + `intersight_techsupportmanagement_collection_control_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_techsupportmanagement_collection_control_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
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
					Name:        "None",
					Description: `A place holder for the default value.`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `Previous action triggered on the resource is still running.`,
				},
				resource.Attribute{
					Name:        "Success",
					Description: `Previous action triggered on the resource has completed successfully.`,
				},
				resource.Attribute{
					Name:        "Failure",
					Description: `Previous action triggered on the resource has failed.`,
				},
				resource.Attribute{
					Name:        "NoCloudSource",
					Description: `Allows the user to provide user-data to the instance without running a network service.`,
				},
				resource.Attribute{
					Name:        "CloudConfigDrive",
					Description: `Allows the user to provide user-data and network-data from cloud. + ` + "`" + `network_data` + "`" + `:(string) Network configuration data for a virtual machine. + ` + "`" + `network_data_base64_encoded` + "`" + `:(bool) Set to true, if the cloud init network data is in base64 format. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `user_data` + "`" + `:(string) User configuration data for a virtual machine such as adding user, group etc. + ` + "`" + `user_data_base64_encoded` + "`" + `:(bool) Set to true, if the cloud init user data is in base64 format.`,
				},
				resource.Attribute{
					Name:        "virtio",
					Description: `Disk uses the same paths as a bare-metal system. This simplifies physical-to-virtual and virtual-to-virtual migration.`,
				},
				resource.Attribute{
					Name:        "sata",
					Description: `Serial ATA (SATA, abbreviated from Serial AT Attachment) is a computer bus interface that connects host bus adapters to mass storage devices such as hard disk drives, optical drives, and solid-state drives.`,
				},
				resource.Attribute{
					Name:        "scsi",
					Description: `SCSI (Small Computer System Interface) bus used.. + ` + "`" + `name` + "`" + `:(string) Virtual machine network bridge name. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `order` + "`" + `:(int) Priority order of the disk. + ` + "`" + `type` + "`" + `:(string) Disk type hdd or cdrom for a virtual machine.`,
				},
				resource.Attribute{
					Name:        "hdd",
					Description: `Allows the virtual machine to mount disk from hard disk drive (hdd) image.`,
				},
				resource.Attribute{
					Name:        "cdrom",
					Description: `Allows the virtual machine to mount disk from compact disk (cd) image. + ` + "`" + `virtual_disk` + "`" + `:(HashMap) - Virtual disk configuration. This complex property has following sub-properties: + ` + "`" + `capacity` + "`" + `:(string) Disk capacity to be provided with units example - 10Gi. + ` + "`" + `mode` + "`" + `:(string) File mode of the disk, example - Filesystem, Block.`,
				},
				resource.Attribute{
					Name:        "Block",
					Description: `It is a Block virtual disk.`,
				},
				resource.Attribute{
					Name:        "Filesystem",
					Description: `It is a File system virtual disk. + ` + "`" + `name` + "`" + `:(string) Name of the virtual disk. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `source_certs` + "`" + `:(string) Base64 encoded CA certificates of the https source to check against. + ` + "`" + `source_disk_to_clone` + "`" + `:(string) Source disk name from where the clone is done. + ` + "`" + `source_file_path` + "`" + `:(string) Disk image source for the virtual machine. + ` + "`" + `virtual_disk_reference` + "`" + `:(string) Name of the existing virtual disk to be attached to the Virtual Machine.`,
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
					Name:        "Unknown",
					Description: `The type of the network adaptor type is unknown.`,
				},
				resource.Attribute{
					Name:        "E1000",
					Description: `Emulated version of the Intel 82545EM Gigabit Ethernet NIC.`,
				},
				resource.Attribute{
					Name:        "SRIOV",
					Description: `Representation of a virtual function (VF) on a physical NIC with SR-IOV support.`,
				},
				resource.Attribute{
					Name:        "VMXNET2",
					Description: `VMXNET 2 (Enhanced) is available only for some guest operating systems on ESX/ESXi 3.5 and later.`,
				},
				resource.Attribute{
					Name:        "VMXNET3",
					Description: `VMXNET 3 offers all the features available in VMXNET 2 and adds several new features. + ` + "`" + `bridge` + "`" + `:(string) Virtual machine network bridge name. + ` + "`" + `connect_at_power_on` + "`" + `:(bool) Connect the adaptor at virtual machine power on. + ` + "`" + `direct_path_io` + "`" + `:(bool) Enable the direct path I/O. + ` + "`" + `mac_address` + "`" + `:(string) Virtual machine network mac address. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "none",
					Description: `No authentication is used.`,
				},
				resource.Attribute{
					Name:        "ntlm",
					Description: `NT LAN Manager (NTLM) security protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2.`,
				},
				resource.Attribute{
					Name:        "ntlmi",
					Description: `NTLMi security protocol. Use this option only when you enable Digital Signing in the CIFS Windows server.`,
				},
				resource.Attribute{
					Name:        "ntlmv2",
					Description: `NTLMv2 security protocol. Use this option only with Samba Linux.`,
				},
				resource.Attribute{
					Name:        "ntlmv2i",
					Description: `NTLMv2i security protocol. Use this option only with Samba Linux.`,
				},
				resource.Attribute{
					Name:        "ntlmssp",
					Description: `NT LAN Manager Security Support Provider (NTLMSSP) protocol. Use this option only with Windows 2008 R2 and Windows 2012 R2.`,
				},
				resource.Attribute{
					Name:        "ntlmsspi",
					Description: `NTLMSSPi protocol. Use this option only when you enable Digital Signing in the CIFS Windows server. + ` + "`" + `device_type` + "`" + `:(string) Type of remote Virtual Media device.`,
				},
				resource.Attribute{
					Name:        "cdd",
					Description: `Uses compact disc drive as the virtual media mount device.`,
				},
				resource.Attribute{
					Name:        "hdd",
					Description: `Uses hard disk drive as the virtual media mount device. + ` + "`" + `file_location` + "`" + `:(string) Remote location of image. Preferred format is 'hostname/filePath/fileName'. + ` + "`" + `host_name` + "`" + `:(string) IP address or hostname of the remote server. + ` + "`" + `is_password_set` + "`" + `:(bool)(Computed) Indicates whether the value of the 'password' property has been set. + ` + "`" + `mount_options` + "`" + `:(string) Mount options for the Virtual Media mapping. The field can be left blank or filled in a comma separated list with the following options.\ For NFS, supported options are ro, rw, nolock, noexec, soft, port=VALUE, timeo=VALUE, retry=VALUE.\ For CIFS, supported options are soft, nounix, noserverino, guest.\ For CIFS version < 3.0, vers=VALUE is mandatory. e.g. vers=2.0\ For HTTP/HTTPS, the only supported option is noauto. + ` + "`" + `mount_protocol` + "`" + `:(string) Protocol to use to communicate with the remote server.`,
				},
				resource.Attribute{
					Name:        "nfs",
					Description: `NFS protocol for vmedia mount.`,
				},
				resource.Attribute{
					Name:        "cifs",
					Description: `CIFS protocol for vmedia mount.`,
				},
				resource.Attribute{
					Name:        "http",
					Description: `HTTP protocol for vmedia mount.`,
				},
				resource.Attribute{
					Name:        "https",
					Description: `HTTPS protocol for vmedia mount. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `password` + "`" + `:(string) Password associated with the username. + ` + "`" + `remote_file` + "`" + `:(string) The remote file location path for the virtual media mapping. Accepted formats are:HDD for CIFS/NFS: hostname-or-IP/filePath/fileName.img.CDD for CIFS/NFS: hostname-or-IP/filePath/fileName.iso.HDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.img.CDD for HTTP/S: http[s]://hostname-or-IP/filePath/fileName.iso. + ` + "`" + `remote_path` + "`" + `:(string) URL path to the location of the image on the remote server. The preferred format is '/path'. + ` + "`" + `sanitized_file_location` + "`" + `:(string)(Computed) File Location in standard format 'hostname/filePath/fileName'. This field should be used to calculate config drift. User input format may vary while inventory will return data in format in compliance with mount option for the mount. Both will be converged to this standard format for comparison. + ` + "`" + `username` + "`" + `:(string) Username to log in to the remote server. + ` + "`" + `volume_name` + "`" + `:(string) Identity of the image for Virtual Media mapping.`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "MIN",
					Description: `The system waits for the time specified in the Coalescing Time field before sending another interrupt event.`,
				},
				resource.Attribute{
					Name:        "IDLE",
					Description: `The system does not send an interrupt until there is a period of no activity lasting as least as long as the time specified in the Coalescing Time field. + ` + "`" + `nr_count` + "`" + `:(int) The number of interrupt resources to allocate. Typical value is be equal to the number of completion queue resources. + ` + "`" + `mode` + "`" + `:(string) Preferred driver interrupt mode. This can be one of the following:- MSIx - Message Signaled Interrupts (MSI) with the optional extension. MSI - MSI only. INTx - PCI INTx interrupts. MSIx is the recommended option.`,
				},
				resource.Attribute{
					Name:        "MSIx",
					Description: `Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option.`,
				},
				resource.Attribute{
					Name:        "MSI",
					Description: `Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts.`,
				},
				resource.Attribute{
					Name:        "INTx",
					Description: `Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
				},
				resource.Attribute{
					Name:        "5",
					Description: `RDMA CoS Service Level 5.`,
				},
				resource.Attribute{
					Name:        "1",
					Description: `RDMA CoS Service Level 1.`,
				},
				resource.Attribute{
					Name:        "2",
					Description: `RDMA CoS Service Level 2.`,
				},
				resource.Attribute{
					Name:        "4",
					Description: `RDMA CoS Service Level 4.`,
				},
				resource.Attribute{
					Name:        "6",
					Description: `RDMA CoS Service Level 6. + ` + "`" + `enabled` + "`" + `:(bool) If enabled sets RDMA over Converged Ethernet (RoCE) on this virtual interface. + ` + "`" + `memory_regions` + "`" + `:(int) The number of memory regions per adapter. Recommended value = integer power of 2. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `queue_pairs` + "`" + `:(int) The number of queue pairs per adapter. Recommended value = integer power of 2. + ` + "`" + `resource_groups` + "`" + `:(int) The number of resource groups per adapter. Recommended value = be an integer power of 2 greater than or equal to the number of CPU cores on the system for optimum performance. + ` + "`" + `nr_version` + "`" + `:(int) Configure RDMA over Converged Ethernet (RoCE) version on the virtual interface. Only RoCEv1 is supported on Cisco VIC 13xx series adapters and only RoCEv2 is supported on Cisco VIC 14xx series adapters.`,
				},
				resource.Attribute{
					Name:        "1",
					Description: `RDMA over Converged Ethernet Protocol Version 1.`,
				},
				resource.Attribute{
					Name:        "2",
					Description: `RDMA over Converged Ethernet Protocol Version 2.`,
				},
			},
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
					Name:        "vnic",
					Description: `Source of the CDN is the same as the vNIC name.`,
				},
				resource.Attribute{
					Name:        "user",
					Description: `Source of the CDN is specified by the user. + ` + "`" + `value` + "`" + `:(string) The CDN value entered in case of user defined mode.`,
				},
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
				resource.Attribute{
					Name:        "None",
					Description: `Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.`,
				},
				resource.Attribute{
					Name:        "A",
					Description: `Fabric A of the FI cluster.`,
				},
				resource.Attribute{
					Name:        "B",
					Description: `Fabric B of the FI cluster. + ` + "`" + `uplink` + "`" + `:(int) Adapter port on which the virtual interface will be created.`,
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
				resource.Attribute{
					Name:        "ACCESS",
					Description: `An access port carries traffic only for a single VLAN on the interface.`,
				},
				resource.Attribute{
					Name:        "TRUNK",
					Description: `A trunk port can have two or more VLANs configured on the interface. It can carry traffic for several VLANs simultaneously. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. ## Usage Example ### Resource Creation ` + "`" + `` + "`" + `` + "`" + `hcl resource "intersight_vnic_eth_network_policy" "v_eth_network1" { name = "v_eth_network1" organization { object_type = "organization.Organization" moid = var.organization } vlan_settings { object_type = "vnic.VlanSettings" default_vlan = 1 mode = "ACCESS" } } ` + "`" + `` + "`" + `` + "`" + ` ## Import ` + "`" + `intersight_vnic_eth_network_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_vnic_eth_network_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "MSIx",
					Description: `Message Signaled Interrupt (MSI) mechanism with the optional extension (MSIx). MSIx is the recommended and default option.`,
				},
				resource.Attribute{
					Name:        "MSI",
					Description: `Message Signaled Interrupt (MSI) mechanism that treats messages as interrupts.`,
				},
				resource.Attribute{
					Name:        "INTx",
					Description: `Line-based interrupt (INTx) mechanism similar to the one used in Legacy systems. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.`,
				},
			},
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
					Name:        "None",
					Description: `Fabric Id is not set to either A or B for the standalone case where the server is not connected to Fabric Interconnects. The value 'None' should be used.`,
				},
				resource.Attribute{
					Name:        "A",
					Description: `Fabric A of the FI cluster.`,
				},
				resource.Attribute{
					Name:        "B",
					Description: `Fabric B of the FI cluster. + ` + "`" + `uplink` + "`" + `:(int) Adapter port on which the virtual interface will be created.`,
				},
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
					Description: `Type indicates that the system selects the target interface automatically during iSCSI boot. ## Import ` + "`" + `intersight_vnic_iscsi_boot_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_vnic_iscsi_boot_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Description: `Servers which are connected to a Fabric Interconnect that is managed by Intersight. ## Usage Example ### Resource Creation ` + "`" + `` + "`" + `` + "`" + `hcl resource "intersight_vnic_lan_connectivity_policy" "vnic_lan1" { name = "vnic_lan1" organization { object_type = "organization.Organization" moid = var.organization } profiles { moid = intersight_server_profile.server1.id object_type = "server.Profile" } } ` + "`" + `` + "`" + `` + "`" + ` ## Import ` + "`" + `intersight_vnic_lan_connectivity_policy` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_vnic_lan_connectivity_policy.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "NonInteractive",
					Description: `The CLI command is not an interactive command.`,
				},
				resource.Attribute{
					Name:        "Interactive",
					Description: `The CLI command is executed in interactive mode and the command must provide the expects andanswers. ### [workflow.FileOperations](#argument-reference) This models a single File Operation request within a batch of requests that get executed within a single workflow task.`,
				},
				resource.Attribute{
					Name:        "FileDownload",
					Description: `The API is executed in a remote device connected to the Intersightthrough its device connector. This operation is to download the filefrom specified storage bucket to the specific path on the device.`,
				},
				resource.Attribute{
					Name:        "FileTemplatize",
					Description: `Populates data driven template file with input values to generate textual output.Inputs - the path of the template file on the device and json values to populate.An error will be returned if the file does not exists or if there is an error whileexecuting the template. ### [workflow.SshSession](#argument-reference) This models a single SSH session from Intersight connected endpoint to a remote server. Multiple SSH operations can be run sequentially over a single SSH session.`,
				},
				resource.Attribute{
					Name:        "ExecuteCommand",
					Description: `Execute a SSH command on the remote server.`,
				},
				resource.Attribute{
					Name:        "NewSession",
					Description: `Open a new SSH connection to the remote server.`,
				},
				resource.Attribute{
					Name:        "FileTransfer",
					Description: `Transfer a file from Intersight connected device to the remote server.`,
				},
				resource.Attribute{
					Name:        "CloseSession",
					Description: `Close the SSH connection to the remote server.`,
				},
				resource.Attribute{
					Name:        "NonInteractiveCmd",
					Description: `Execute a non-interactive SSH command on the remote server.`,
				},
				resource.Attribute{
					Name:        "InteractiveCmd",
					Description: `Execute an interactive SSH command on the remote server. + ` + "`" + `encrypted_aes_key` + "`" + `:(string) The secure properties that have large text content as value can be encrypted using AES key. In these cases, the AES key needs to be encrypted using the device connector public key and passed as the value for this property.The secure properties that are encrypted using the AES key are mapped against the property name with prefix 'AES' in SecureProperties dictionary. + ` + "`" + `encryption_key` + "`" + `:(string) The public key that was used to encrypt the values present in SecureProperties dictionary.If the given public key is not same as device connector's public key, an error reponse with appropriate error message is thrown back. + ` + "`" + `expect_prompts` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `expect` + "`" + `:(string) The regex of the expect prompt of the interactive command. + ` + "`" + `expect_timeout` + "`" + `:(int) The timeout for the expect prompt while executing interactive command. If timeout is not set a default of 60 seconds will be used. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `send` + "`" + `:(string) The answer string to the expect prompt. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `secure_properties` + "`" + `: A dictionary of encrypted secure values mapped against the secure property name. The values that are encrypted using AES key must be mapped against the secure property name with a 'AES' prefixDevice connector expects the message body to be a golang template and the template can use the secure property names as placeholders. + ` + "`" + `shell_prompt` + "`" + `:(string) Regex of the remote server's shell prompt. + ` + "`" + `shell_prompt_timeout` + "`" + `:(int) Expect timeout value in seconds for the shell prompt.`,
				},
				resource.Attribute{
					Name:        "Internal",
					Description: `The endpoint API executed is an internal request handled by the device connector plugin.`,
				},
				resource.Attribute{
					Name:        "External",
					Description: `The endpoint API request is passed through by the device connector.`,
				},
			},
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
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "eq",
					Description: `Checks if the values of the two parameters are equal.`,
				},
				resource.Attribute{
					Name:        "ne",
					Description: `Checks if the values of the two parameters are not equal.`,
				},
				resource.Attribute{
					Name:        "contains",
					Description: `Checks if the second parameter string value is a substring of the first parameter string value.`,
				},
				resource.Attribute{
					Name:        "matchesPattern",
					Description: `Checks if a string matches a regular expression. + ` + "`" + `control_parameter` + "`" + `:(string) Name of the controlling entity, whose value will be used for evaluating the parameter set. + ` + "`" + `enable_parameters` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `name` + "`" + `:(string) Name for the parameter set. Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) The controlling parameter will be evaluated against this 'value'.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. ## Import ` + "`" + `intersight_workflow_custom_data_type_definition` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_workflow_custom_data_type_definition.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ` ## Allowed Types in ` + "`" + `AdditionalProperties` + "`" + ` ### [workflow.ArrayDataType](#argument-reference) This data type represents an array of a given type. It can be an array of primitive data or of custom data.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `Enum to specify a string data type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `Enum to specify an integer32 data type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `Enum to specify a float64 data type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `Enum to specify a boolean data type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Enum to specify a json data type.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `Enum to specify a enum data type which is a list of pre-defined strings. ### [workflow.TargetDataType](#argument-reference) Data type to capture a target endpoint or device.`,
				},
			},
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
					Name:        "simple",
					Description: `The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The parameter value to be extracted is of the string type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `The parameter value to be extracted is of the number type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `The parameter value to be extracted is of the float number type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `The parameter value to be extracted is of the boolean type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.`,
				},
				resource.Attribute{
					Name:        "complex",
					Description: `The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.`,
				},
				resource.Attribute{
					Name:        "collection",
					Description: `The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. + ` + "`" + `name` + "`" + `:(string) The name of the parameter. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `path` + "`" + `:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. + ` + "`" + `type` + "`" + `:(string) The type of the parameter. Accepted values are simple, complex,collection.`,
				},
				resource.Attribute{
					Name:        "simple",
					Description: `The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The parameter value to be extracted is of the string type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `The parameter value to be extracted is of the number type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `The parameter value to be extracted is of the float number type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `The parameter value to be extracted is of the boolean type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.`,
				},
				resource.Attribute{
					Name:        "complex",
					Description: `The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.`,
				},
				resource.Attribute{
					Name:        "collection",
					Description: `The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type.`,
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
					Name:        "simple",
					Description: `The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The parameter value to be extracted is of the string type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `The parameter value to be extracted is of the number type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `The parameter value to be extracted is of the float number type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `The parameter value to be extracted is of the boolean type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.`,
				},
				resource.Attribute{
					Name:        "complex",
					Description: `The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.`,
				},
				resource.Attribute{
					Name:        "collection",
					Description: `The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. + ` + "`" + `name` + "`" + `:(string) The name of the parameter. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `path` + "`" + `:(string) The content specific path information that identifies the parametervalue within the content. The value is usually a XPath or JSONPath or aregular expression in case of text content. + ` + "`" + `type` + "`" + `:(string) The type of the parameter. Accepted values are simple, complex,collection.`,
				},
				resource.Attribute{
					Name:        "simple",
					Description: `The parameter value to be extracted is of the type simple. All the common scalar typessuch as int, bool, string, etc are represented by the simple enum.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `The parameter value to be extracted is of the string type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `The parameter value to be extracted is of the number type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `The parameter value to be extracted is of the float number type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `The parameter value to be extracted is of the boolean type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `The parameter values to be extracted is of the generic JSON literal. JSON type is applicable only if the content to be parsed is of JSON type.`,
				},
				resource.Attribute{
					Name:        "complex",
					Description: `The parameter value to be extracted is a complex parameter that itself isanother collection of simple/complex parameters.`,
				},
				resource.Attribute{
					Name:        "collection",
					Description: `The parameter value to be extracted is a collection parameter whose item typeshall be either simple type or complex type. ## Import ` + "`" + `intersight_workflow_error_response_handler` + "`" + ` can be imported using the Moid of the object, e.g. ` + "`" + `` + "`" + `` + "`" + ` $ terraform import intersight_workflow_error_response_handler.example 1234567890987654321abcde ` + "`" + `` + "`" + `` + "`" + ``,
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
					Name:        "NotStarted",
					Description: `Status of rollback task when it is not started rollback.`,
				},
				resource.Attribute{
					Name:        "NotSupported",
					Description: `Status of task when it is not supporting rollback.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Status of rollback task once execution is successful.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Status of rollback task when it is failed. + ` + "`" + `task_info_moid` + "`" + `:(string) Moid of TaskInfo that supports rollback operation. + ` + "`" + `task_path` + "`" + `:(string)(Computed) Path of rollback task if it is inside sub-workflow.`,
				},
				resource.Attribute{
					Name:        "NotStarted",
					Description: `Status of rollback task when it is not started rollback.`,
				},
				resource.Attribute{
					Name:        "NotSupported",
					Description: `Status of task when it is not supporting rollback.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Status of rollback task once execution is successful.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `Status of rollback task when it is failed. + ` + "`" + `task_info_moid` + "`" + `:(string) Moid of TaskInfo that supports rollback operation. + ` + "`" + `task_path` + "`" + `:(string)(Computed) Path of rollback task if it is inside sub-workflow.`,
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
				resource.Attribute{
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `output_definition` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `additional_properties` + "`" + `:(JSON) - Additional Properties as per object type, can be added as JSON using ` + "`" + `jsonencode()` + "`" + `. Allowed Types are: [workflow.ArrayDataType](#workflowArrayDataType) [workflow.CustomDataType](#workflowCustomDataType) [workflow.MoReferenceDataType](#workflowMoReferenceDataType) [workflow.PrimitiveDataType](#workflowPrimitiveDataType) [workflow.TargetDataType](#workflowTargetDataType) + ` + "`" + `default` + "`" + `:(HashMap) - Default value for the data type. If default value was provided and the input was required the default value will be used as the input. This complex property has following sub-properties: + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `override` + "`" + `:(bool) Override the default value provided for the data type. When true, allow the user to enter value for the data type. + ` + "`" + `value` + "`" + `: Default value for the data type. If default value was provided and the input was required the default value will be used as the input. + ` + "`" + `description` + "`" + `:(string) Provide a detailed description of the data type. + ` + "`" + `display_meta` + "`" + `:(HashMap) - Captures the meta data needed for displaying workflow data types in Intersight User Interface. This complex property has following sub-properties: + ` + "`" + `inventory_selector` + "`" + `:(bool) Inventory selector specified for primitive data property should be used in Intersight User Interface. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `widget_type` + "`" + `:(string) Specify the widget type for data display.`,
				},
				resource.Attribute{
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow. + ` + "`" + `retry_count` + "`" + `:(int) The number of times a task should be tried before marking as failed. + ` + "`" + `retry_delay` + "`" + `:(int) The delay in seconds after which the the task is re-tried. + ` + "`" + `retry_policy` + "`" + `:(string) The retry policy for the task.`,
				},
				resource.Attribute{
					Name:        "Fixed",
					Description: `The enum specifies the option as Fixed where the task retry happens after fixed time specified by RetryDelay. + ` + "`" + `support_status` + "`" + `:(string) Supported status of the definition.`,
				},
				resource.Attribute{
					Name:        "Supported",
					Description: `The definition is a supported version and there will be no changes to the mandatory inputs or outputs.`,
				},
				resource.Attribute{
					Name:        "Beta",
					Description: `The definition is a Beta version and this version can under go changes until the version is marked supported.`,
				},
				resource.Attribute{
					Name:        "Deprecated",
					Description: `The version of definition is deprecated and typically there will be a higher version of the same definition that has been added. + ` + "`" + `timeout` + "`" + `:(int) The timeout value in seconds after which task will be marked as timed out. Max allowed value is 7 days. + ` + "`" + `timeout_policy` + "`" + `:(string) The timeout policy for the task.`,
				},
				resource.Attribute{
					Name:        "Timeout",
					Description: `The enum specifies the option as Timeout where task will be timed out after the specified time in Timeout property.`,
				},
				resource.Attribute{
					Name:        "Retry",
					Description: `The enum specifies the option as Retry where task will be re-tried.`,
				},
			},
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
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow.`,
				},
				resource.Attribute{
					Name:        "eq",
					Description: `Checks if the values of the two parameters are equal.`,
				},
				resource.Attribute{
					Name:        "ne",
					Description: `Checks if the values of the two parameters are not equal.`,
				},
				resource.Attribute{
					Name:        "contains",
					Description: `Checks if the second parameter string value is a substring of the first parameter string value.`,
				},
				resource.Attribute{
					Name:        "matchesPattern",
					Description: `Checks if a string matches a regular expression. + ` + "`" + `control_parameter` + "`" + `:(string) Name of the controlling entity, whose value will be used for evaluating the parameter set. + ` + "`" + `enable_parameters` + "`" + `: (Array of schema.TypeString) - + ` + "`" + `name` + "`" + `:(string) Name for the parameter set. Limited to 64 alphanumeric characters (upper and lower case), and special characters '-' and '_'. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `value` + "`" + `:(string) The controlling parameter will be evaluated against this 'value'.`,
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
				resource.Attribute{
					Name:        "None",
					Description: `Display none of the widget types.`,
				},
				resource.Attribute{
					Name:        "Radio",
					Description: `Display the widget as a radio button.`,
				},
				resource.Attribute{
					Name:        "Dropdown",
					Description: `Display the widget as a dropdown.`,
				},
				resource.Attribute{
					Name:        "GridSelector",
					Description: `Display the widget as a selector.`,
				},
				resource.Attribute{
					Name:        "DrawerSelector",
					Description: `Display the widget as a selector. + ` + "`" + `input_parameters` + "`" + `: JSON formatted mapping from other property of the definition to the current property. Input parameter mapping is supported only for custom data type property in workflow definition and custom data type definition. The format to specify mapping ina workflow definition when source property is of scalar types is '${workflow.input.property}'. The format to specify mapping when the source property is of object reference and mapping needs to be made to the property of the object is '${workflow.input.property.subproperty}'. The format to specify mapping in a custom data type definition is '${datatype.type.property}'. When the current property is of non-scalar type like composite custom data type, then mapping can be provided to the individual property of the custom data type like 'cdt_property:${workflow.input.property}'. + ` + "`" + `label` + "`" + `:(string) Descriptive label for the data type. Label can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-), space ( ) or an underscore (_). The first and last character in label must be an alphanumeric character. + ` + "`" + `name` + "`" + `:(string) Descriptive name for the data type. Name can only contain letters (a-z, A-Z), numbers (0-9), hyphen (-) or an underscore (_). The first and last character in name must be an alphanumeric character. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property.The enum values provides the list of concrete types that can be instantiated from this abstract type. + ` + "`" + `required` + "`" + `:(bool) Specifies whether this parameter is required. The field is applicable for task and workflow.`,
				},
				resource.Attribute{
					Name:        "Supported",
					Description: `The definition is a supported version and there will be no changes to the mandatory inputs or outputs.`,
				},
				resource.Attribute{
					Name:        "Beta",
					Description: `The definition is a Beta version and this version can under go changes until the version is marked supported.`,
				},
				resource.Attribute{
					Name:        "Deprecated",
					Description: `The version of definition is deprecated and typically there will be a higher version of the same definition that has been added.`,
				},
				resource.Attribute{
					Name:        "NotValidated",
					Description: `The state when workflow definition has not been validated.`,
				},
				resource.Attribute{
					Name:        "Valid",
					Description: `The state when workflow definition is valid.`,
				},
				resource.Attribute{
					Name:        "Invalid",
					Description: `The state when workflow definition is invalid. + ` + "`" + `validation_error` + "`" + `:(Array) This complex property has following sub-properties: + ` + "`" + `error_log` + "`" + `:(string)(Computed) Description of the error. + ` + "`" + `field` + "`" + `:(string)(Computed) When populated this refers to the input or output field within the workflow or task. + ` + "`" + `object_type` + "`" + `:(string) The fully-qualified name of the instantiated, concrete type.The value should be the same as the 'ClassId' property. + ` + "`" + `task_name` + "`" + `:(string)(Computed) The task name on which the error is found, when empty the error applies to the top level workflow. + ` + "`" + `transition_name` + "`" + `:(string)(Computed) When populated this refers to the transition connection that has a problem. When this field has value OnSuccess it means the transition connection OnSuccess for the task has an issue.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `Enum to specify a string data type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `Enum to specify an integer32 data type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `Enum to specify a float64 data type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `Enum to specify a boolean data type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Enum to specify a json data type.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `Enum to specify a enum data type which is a list of pre-defined strings. ### [workflow.TargetDataType](#argument-reference) Data type to capture a target endpoint or device.`,
				},
				resource.Attribute{
					Name:        "string",
					Description: `Enum to specify a string data type.`,
				},
				resource.Attribute{
					Name:        "integer",
					Description: `Enum to specify an integer32 data type.`,
				},
				resource.Attribute{
					Name:        "float",
					Description: `Enum to specify a float64 data type.`,
				},
				resource.Attribute{
					Name:        "boolean",
					Description: `Enum to specify a boolean data type.`,
				},
				resource.Attribute{
					Name:        "json",
					Description: `Enum to specify a json data type.`,
				},
				resource.Attribute{
					Name:        "enum",
					Description: `Enum to specify a enum data type which is a list of pre-defined strings. ### [workflow.TargetDataType](#argument-reference) Data type to capture a target endpoint or device.`,
				},
				resource.Attribute{
					Name:        "Scheduled",
					Description: `The enum represents the status when task is in scheduled state.`,
				},
				resource.Attribute{
					Name:        "InProgress",
					Description: `The enum represents the status when task is in-progress state.`,
				},
				resource.Attribute{
					Name:        "NoOp",
					Description: `The enum represents the status when task is a noop.`,
				},
				resource.Attribute{
					Name:        "Timeout",
					Description: `The enum represents the status when task has timed out.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `The enum represents the status when task has completed.`,
				},
				resource.Attribute{
					Name:        "Failed",
					Description: `The enum represents the status when task has failed.`,
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
					Name:        "Info",
					Description: `The enum represents the log level to be used to convey info message.`,
				},
				resource.Attribute{
					Name:        "Warning",
					Description: `The enum represents the log level to be used to convey warning message.`,
				},
				resource.Attribute{
					Name:        "Debug",
					Description: `The enum represents the log level to be used to convey debug message.`,
				},
				resource.Attribute{
					Name:        "Error",
					Description: `The enum represents the log level to be used to convey error message.`,
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
					Name:        "Disabled",
					Description: `Status of the rollback action when workflow is disabled for rollback.`,
				},
				resource.Attribute{
					Name:        "Enabled",
					Description: `Status of the rollback action when workflow is enabled for rollback.`,
				},
				resource.Attribute{
					Name:        "Completed",
					Description: `Status of the rollback action once workflow completes the rollback for all eligiable tasks.`,
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
	}

	resourcesMap = map[string]int{

		"intersight_access_policy":                                           0,
		"intersight_adapter_config_policy":                                   1,
		"intersight_appliance_backup":                                        2,
		"intersight_appliance_backup_policy":                                 3,
		"intersight_appliance_data_export_policy":                            4,
		"intersight_appliance_device_claim":                                  5,
		"intersight_appliance_diag_setting":                                  6,
		"intersight_appliance_remote_file_import":                            7,
		"intersight_appliance_restore":                                       8,
		"intersight_asset_target":                                            9,
		"intersight_bios_policy":                                             10,
		"intersight_boot_precision_policy":                                   11,
		"intersight_capability_adapter_unit_descriptor":                      12,
		"intersight_capability_chassis_descriptor":                           13,
		"intersight_capability_chassis_manufacturing_def":                    14,
		"intersight_capability_cimc_firmware_descriptor":                     15,
		"intersight_capability_equipment_physical_def":                       16,
		"intersight_capability_equipment_slot_array":                         17,
		"intersight_capability_fan_module_descriptor":                        18,
		"intersight_capability_fan_module_manufacturing_def":                 19,
		"intersight_capability_io_card_capability_def":                       20,
		"intersight_capability_io_card_descriptor":                           21,
		"intersight_capability_io_card_manufacturing_def":                    22,
		"intersight_capability_port_group_aggregation_def":                   23,
		"intersight_capability_psu_descriptor":                               24,
		"intersight_capability_psu_manufacturing_def":                        25,
		"intersight_capability_sioc_module_capability_def":                   26,
		"intersight_capability_sioc_module_descriptor":                       27,
		"intersight_capability_sioc_module_manufacturing_def":                28,
		"intersight_capability_switch_capability":                            29,
		"intersight_capability_switch_descriptor":                            30,
		"intersight_capability_switch_manufacturing_def":                     31,
		"intersight_certificatemanagement_policy":                            32,
		"intersight_chassis_config_import":                                   33,
		"intersight_chassis_profile":                                         34,
		"intersight_comm_http_proxy_policy":                                  35,
		"intersight_config_exporter":                                         36,
		"intersight_config_importer":                                         37,
		"intersight_connectorpack_connector_pack_upgrade":                    38,
		"intersight_deviceconnector_policy":                                  39,
		"intersight_externalsite_authorization":                              40,
		"intersight_fabric_appliance_pc_role":                                41,
		"intersight_fabric_appliance_role":                                   42,
		"intersight_fabric_eth_network_control_policy":                       43,
		"intersight_fabric_eth_network_group_policy":                         44,
		"intersight_fabric_eth_network_policy":                               45,
		"intersight_fabric_fc_network_policy":                                46,
		"intersight_fabric_fc_uplink_pc_role":                                47,
		"intersight_fabric_fc_uplink_role":                                   48,
		"intersight_fabric_fcoe_uplink_pc_role":                              49,
		"intersight_fabric_fcoe_uplink_role":                                 50,
		"intersight_fabric_multicast_policy":                                 51,
		"intersight_fabric_pc_operation":                                     52,
		"intersight_fabric_port_mode":                                        53,
		"intersight_fabric_port_operation":                                   54,
		"intersight_fabric_port_policy":                                      55,
		"intersight_fabric_server_role":                                      56,
		"intersight_fabric_switch_cluster_profile":                           57,
		"intersight_fabric_switch_control_policy":                            58,
		"intersight_fabric_switch_profile":                                   59,
		"intersight_fabric_system_qos_policy":                                60,
		"intersight_fabric_uplink_pc_role":                                   61,
		"intersight_fabric_uplink_role":                                      62,
		"intersight_fabric_vlan":                                             63,
		"intersight_fabric_vsan":                                             64,
		"intersight_fcpool_pool":                                             65,
		"intersight_firmware_bios_descriptor":                                66,
		"intersight_firmware_board_controller_descriptor":                    67,
		"intersight_firmware_chassis_upgrade":                                68,
		"intersight_firmware_cimc_descriptor":                                69,
		"intersight_firmware_dimm_descriptor":                                70,
		"intersight_firmware_distributable":                                  71,
		"intersight_firmware_drive_descriptor":                               72,
		"intersight_firmware_driver_distributable":                           73,
		"intersight_firmware_eula":                                           74,
		"intersight_firmware_gpu_descriptor":                                 75,
		"intersight_firmware_hba_descriptor":                                 76,
		"intersight_firmware_iom_descriptor":                                 77,
		"intersight_firmware_mswitch_descriptor":                             78,
		"intersight_firmware_nxos_descriptor":                                79,
		"intersight_firmware_pcie_descriptor":                                80,
		"intersight_firmware_psu_descriptor":                                 81,
		"intersight_firmware_sas_expander_descriptor":                        82,
		"intersight_firmware_server_configuration_utility_distributable":     83,
		"intersight_firmware_storage_controller_descriptor":                  84,
		"intersight_firmware_switch_upgrade":                                 85,
		"intersight_firmware_unsupported_version_upgrade":                    86,
		"intersight_firmware_upgrade":                                        87,
		"intersight_hcl_hyperflex_software_compatibility_info":               88,
		"intersight_hyperflex_app_catalog":                                   89,
		"intersight_hyperflex_auto_support_policy":                           90,
		"intersight_hyperflex_capability_info":                               91,
		"intersight_hyperflex_cisco_hypervisor_manager":                      92,
		"intersight_hyperflex_cluster_backup_policy":                         93,
		"intersight_hyperflex_cluster_backup_policy_deployment":              94,
		"intersight_hyperflex_cluster_network_policy":                        95,
		"intersight_hyperflex_cluster_profile":                               96,
		"intersight_hyperflex_cluster_replication_network_policy":            97,
		"intersight_hyperflex_cluster_replication_network_policy_deployment": 98,
		"intersight_hyperflex_cluster_storage_policy":                        99,
		"intersight_hyperflex_ext_fc_storage_policy":                         100,
		"intersight_hyperflex_ext_iscsi_storage_policy":                      101,
		"intersight_hyperflex_feature_limit_external":                        102,
		"intersight_hyperflex_feature_limit_internal":                        103,
		"intersight_hyperflex_health_check_definition":                       104,
		"intersight_hyperflex_health_check_package_checksum":                 105,
		"intersight_hyperflex_hxap_datacenter":                               106,
		"intersight_hyperflex_hxdp_version":                                  107,
		"intersight_hyperflex_local_credential_policy":                       108,
		"intersight_hyperflex_node_config_policy":                            109,
		"intersight_hyperflex_node_profile":                                  110,
		"intersight_hyperflex_proxy_setting_policy":                          111,
		"intersight_hyperflex_server_firmware_version":                       112,
		"intersight_hyperflex_server_firmware_version_entry":                 113,
		"intersight_hyperflex_server_model":                                  114,
		"intersight_hyperflex_software_version_policy":                       115,
		"intersight_hyperflex_sys_config_policy":                             116,
		"intersight_hyperflex_ucsm_config_policy":                            117,
		"intersight_hyperflex_vcenter_config_policy":                         118,
		"intersight_hyperflex_vm_restore_operation":                          119,
		"intersight_iam_account":                                             120,
		"intersight_iam_account_experience":                                  121,
		"intersight_iam_api_key":                                             122,
		"intersight_iam_app_registration":                                    123,
		"intersight_iam_certificate":                                         124,
		"intersight_iam_certificate_request":                                 125,
		"intersight_iam_end_point_user":                                      126,
		"intersight_iam_end_point_user_policy":                               127,
		"intersight_iam_end_point_user_role":                                 128,
		"intersight_iam_idp":                                                 129,
		"intersight_iam_ip_access_management":                                130,
		"intersight_iam_ip_address":                                          131,
		"intersight_iam_ldap_group":                                          132,
		"intersight_iam_ldap_policy":                                         133,
		"intersight_iam_ldap_provider":                                       134,
		"intersight_iam_permission":                                          135,
		"intersight_iam_private_key_spec":                                    136,
		"intersight_iam_qualifier":                                           137,
		"intersight_iam_resource_roles":                                      138,
		"intersight_iam_session_limits":                                      139,
		"intersight_iam_trust_point":                                         140,
		"intersight_iam_user":                                                141,
		"intersight_iam_user_group":                                          142,
		"intersight_ipmioverlan_policy":                                      143,
		"intersight_ippool_pool":                                             144,
		"intersight_iqnpool_pool":                                            145,
		"intersight_kubernetes_aci_cni_apic":                                 146,
		"intersight_kubernetes_aci_cni_profile":                              147,
		"intersight_kubernetes_aci_cni_tenant_cluster_allocation":            148,
		"intersight_kubernetes_addon":                                        149,
		"intersight_kubernetes_addon_definition":                             150,
		"intersight_kubernetes_addon_policy":                                 151,
		"intersight_kubernetes_addon_repository":                             152,
		"intersight_kubernetes_cluster":                                      153,
		"intersight_kubernetes_cluster_profile":                              154,
		"intersight_kubernetes_container_runtime_policy":                     155,
		"intersight_kubernetes_network_policy":                               156,
		"intersight_kubernetes_node_group_profile":                           157,
		"intersight_kubernetes_sys_config_policy":                            158,
		"intersight_kubernetes_trusted_registries_policy":                    159,
		"intersight_kubernetes_version":                                      160,
		"intersight_kubernetes_version_policy":                               161,
		"intersight_kubernetes_virtual_machine_infrastructure_provider":      162,
		"intersight_kubernetes_virtual_machine_instance_type":                163,
		"intersight_kubernetes_virtual_machine_node_profile":                 164,
		"intersight_kvm_policy":                                              165,
		"intersight_kvm_session":                                             166,
		"intersight_kvm_tunnel":                                              167,
		"intersight_license_iwo_license_count":                               168,
		"intersight_license_license_info":                                    169,
		"intersight_license_license_reservation_op":                          170,
		"intersight_macpool_pool":                                            171,
		"intersight_memory_persistent_memory_policy":                         172,
		"intersight_networkconfig_policy":                                    173,
		"intersight_ntp_policy":                                              174,
		"intersight_organization_organization":                               175,
		"intersight_os_configuration_file":                                   176,
		"intersight_os_install":                                              177,
		"intersight_recovery_backup_config_policy":                           178,
		"intersight_recovery_backup_profile":                                 179,
		"intersight_recovery_on_demand_backup":                               180,
		"intersight_recovery_restore":                                        181,
		"intersight_recovery_schedule_config_policy":                         182,
		"intersight_resource_group":                                          183,
		"intersight_sdcard_policy":                                           184,
		"intersight_sdwan_profile":                                           185,
		"intersight_sdwan_router_node":                                       186,
		"intersight_sdwan_router_policy":                                     187,
		"intersight_sdwan_vmanage_account_policy":                            188,
		"intersight_server_config_import":                                    189,
		"intersight_server_profile":                                          190,
		"intersight_smtp_policy":                                             191,
		"intersight_snmp_policy":                                             192,
		"intersight_software_appliance_distributable":                        193,
		"intersight_software_hcl_meta":                                       194,
		"intersight_software_hyperflex_bundle_distributable":                 195,
		"intersight_software_hyperflex_distributable":                        196,
		"intersight_software_solution_distributable":                         197,
		"intersight_software_ucsd_bundle_distributable":                      198,
		"intersight_software_ucsd_distributable":                             199,
		"intersight_softwarerepository_authorization":                        200,
		"intersight_softwarerepository_category_mapper":                      201,
		"intersight_softwarerepository_category_mapper_model":                202,
		"intersight_softwarerepository_category_support_constraint":          203,
		"intersight_softwarerepository_operating_system_file":                204,
		"intersight_softwarerepository_release":                              205,
		"intersight_sol_policy":                                              206,
		"intersight_ssh_policy":                                              207,
		"intersight_storage_disk_group_policy":                               208,
		"intersight_storage_storage_policy":                                  209,
		"intersight_syslog_policy":                                           210,
		"intersight_tam_advisory_count":                                      211,
		"intersight_tam_advisory_definition":                                 212,
		"intersight_tam_advisory_info":                                       213,
		"intersight_tam_advisory_instance":                                   214,
		"intersight_tam_security_advisory":                                   215,
		"intersight_techsupportmanagement_collection_control_policy":         216,
		"intersight_techsupportmanagement_tech_support_bundle":               217,
		"intersight_uuidpool_pool":                                           218,
		"intersight_virtualization_virtual_disk":                             219,
		"intersight_virtualization_virtual_machine":                          220,
		"intersight_vmedia_policy":                                           221,
		"intersight_vmrc_console":                                            222,
		"intersight_vnic_eth_adapter_policy":                                 223,
		"intersight_vnic_eth_if":                                             224,
		"intersight_vnic_eth_network_policy":                                 225,
		"intersight_vnic_eth_qos_policy":                                     226,
		"intersight_vnic_fc_adapter_policy":                                  227,
		"intersight_vnic_fc_if":                                              228,
		"intersight_vnic_fc_network_policy":                                  229,
		"intersight_vnic_fc_qos_policy":                                      230,
		"intersight_vnic_iscsi_adapter_policy":                               231,
		"intersight_vnic_iscsi_boot_policy":                                  232,
		"intersight_vnic_iscsi_static_target_policy":                         233,
		"intersight_vnic_lan_connectivity_policy":                            234,
		"intersight_vnic_san_connectivity_policy":                            235,
		"intersight_vrf_vrf":                                                 236,
		"intersight_workflow_batch_api_executor":                             237,
		"intersight_workflow_custom_data_type_definition":                    238,
		"intersight_workflow_error_response_handler":                         239,
		"intersight_workflow_rollback_workflow":                              240,
		"intersight_workflow_task_definition":                                241,
		"intersight_workflow_workflow_definition":                            242,
		"intersight_workflow_workflow_info":                                  243,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
