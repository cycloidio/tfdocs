package vsphere

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	Resources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a vSphere cluster resource. This can be used to create and manage clusters of hosts.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the cluster.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the datacenter to create the cluster in. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The relative path to a folder to put this cluster in. This is a path relative to the datacenter you are deploying the cluster to. Example: for the ` + "`" + `dc1` + "`" + ` datacenter, and a provided ` + "`" + `folder` + "`" + ` of ` + "`" + `foo/bar` + "`" + `, Terraform will place a cluster named ` + "`" + `terraform-compute-cluster-test` + "`" + ` in a host folder located at ` + "`" + `/dc1/host/foo/bar` + "`" + `, with the final inventory path being ` + "`" + `/dc1/host/foo/bar/terraform-datastore-cluster-test` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) A map of custom attribute ids to attribute value strings to set for the datastore cluster. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "host_system_ids",
					Description: `(Optional) The [managed object IDs][docs-about-morefs] of the hosts to put in the cluster. Conflicts with: ` + "`" + `host_managed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_managed",
					Description: `(Optional) Can be set to ` + "`" + `true` + "`" + ` if compute cluster membership will be managed through the ` + "`" + `host` + "`" + ` resource rather than the ` + "`" + `compute_cluster` + "`" + ` resource. Conflicts with: ` + "`" + `host_system_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "host_cluster_exit_timeout",
					Description: `The timeout, in seconds, for each host maintenance mode operation when removing hosts from a cluster. Default: ` + "`" + `3600` + "`" + ` seconds (1 hour).`,
				},
				resource.Attribute{
					Name:        "force_evacuate_on_destroy",
					Description: `When destroying the resource, setting this to ` + "`" + `true` + "`" + ` will auto-remove any hosts that are currently a member of the cluster, as if they were removed by taking their entry out of ` + "`" + `host_system_ids` + "`" + ` (see [below](#how-terraform-removes-hosts-from-clusters)). This is an advanced option and should only be used for testing. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "drs_enabled",
					Description: `(Optional) Enable DRS for this cluster. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drs_migration_threshold",
					Description: `(Optional) A value between ` + "`" + `1` + "`" + ` and ` + "`" + `5` + "`" + ` indicating the threshold of imbalance tolerated between hosts. A lower setting will tolerate more imbalance while a higher setting will tolerate less. Default: ` + "`" + `3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drs_enable_vm_overrides",
					Description: `(Optional) Allow individual DRS overrides to be set for virtual machines in the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drs_enable_predictive_drs",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, enables DRS to use data from [vRealize Operations Manager][ref-vsphere-vrops] to make proactive DRS recommendations. <sup>[\`,
				},
				resource.Attribute{
					Name:        "drs_scale_descendants_shares",
					Description: `(Optional) Enable scalable shares for all resource pools in the cluster. Can be one of ` + "`" + `disabled` + "`" + ` or ` + "`" + `scaleCpuAndMemoryShares` + "`" + `. Default: ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drs_advanced_options",
					Description: `(Optional) A key/value map that specifies advanced options for DRS and [DPM](#dpm-options). #### DPM Options The following settings control the [Distributed Power Management][ref-vsphere-dpm] (DPM) settings for the cluster. DPM allows the cluster to manage host capacity on-demand depending on the needs of the cluster, powering on hosts when capacity is needed, and placing hosts in standby when there is excess capacity in the cluster. [ref-vsphere-dpm]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.resmgmt.doc/GUID-5E5E349A-4644-4C9C-B434-1C0243EBDC80.html`,
				},
				resource.Attribute{
					Name:        "dpm_enabled",
					Description: `(Optional) Enable DPM support for DRS in this cluster. Requires [` + "`" + `drs_enabled` + "`" + `](#drs_enabled) to be ` + "`" + `true` + "`" + ` in order to be effective. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dpm_automation_level",
					Description: `(Optional) The automation level for host power operations in this cluster. Can be one of ` + "`" + `manual` + "`" + ` or ` + "`" + `automated` + "`" + `. Default: ` + "`" + `manual` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dpm_threshold",
					Description: `(Optional) A value between ` + "`" + `1` + "`" + ` and ` + "`" + `5` + "`" + ` indicating the threshold of load within the cluster that influences host power operations. This affects both power on and power off operations - a lower setting will tolerate more of a surplus/deficit than a higher setting. Default: ` + "`" + `3` + "`" + `. ### vSphere HA Options The following settings control the [vSphere HA][ref-vsphere-ha-clusters] settings for the cluster. ~>`,
				},
				resource.Attribute{
					Name:        "ha_enabled",
					Description: `(Optional) Enable vSphere HA for this cluster. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_host_monitoring",
					Description: `(Optional) Global setting that controls whether vSphere HA remediates virtual machines on host failure. Can be one of ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `. Default: ` + "`" + `enabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_restart_priority",
					Description: `(Optional) The default restart priority for affected virtual machines when vSphere detects a host failure. Can be one of ` + "`" + `lowest` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `highest` + "`" + `. Default: ` + "`" + `medium` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_dependency_restart_condition",
					Description: `(Optional) The condition used to determine whether or not virtual machines in a certain restart priority class are online, allowing HA to move on to restarting virtual machines on the next priority. Can be one of ` + "`" + `none` + "`" + `, ` + "`" + `poweredOn` + "`" + `, ` + "`" + `guestHbStatusGreen` + "`" + `, or ` + "`" + `appHbStatusGreen` + "`" + `. The default is ` + "`" + `none` + "`" + `, which means that a virtual machine is considered ready immediately after a host is found to start it on. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_vm_restart_additional_delay",
					Description: `(Optional) Additional delay, in seconds, after ready condition is met. A VM is considered ready at this point. Default: ` + "`" + `0` + "`" + ` seconds (no delay). <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_vm_restart_timeout",
					Description: `(Optional) The maximum time, in seconds, that vSphere HA will wait for virtual machines in one priority to be ready before proceeding with the next priority. Default: ` + "`" + `600` + "`" + ` seconds (10 minutes). <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_host_isolation_response",
					Description: `(Optional) The action to take on virtual machines when a host has detected that it has been isolated from the rest of the cluster. Can be one of ` + "`" + `none` + "`" + `, ` + "`" + `powerOff` + "`" + `, or ` + "`" + `shutdown` + "`" + `. Default: ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_advanced_options",
					Description: `(Optional) A key/value map that specifies advanced options for vSphere HA. #### HA Virtual Machine Component Protection Settings The following settings control Virtual Machine Component Protection (VMCP) in vSphere HA. VMCP gives vSphere HA the ability to monitor a host for datastore accessibility failures, and automate recovery for affected virtual machines. ->`,
				},
				resource.Attribute{
					Name:        "ha_vm_component_protection",
					Description: `(Optional) Controls vSphere VM component protection for virtual machines in this cluster. Can be one of ` + "`" + `enabled` + "`" + ` or ` + "`" + `disabled` + "`" + `. Default: ` + "`" + `enabled` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_pdl_response",
					Description: `(Optional) Controls the action to take on virtual machines when the cluster has detected a permanent device loss to a relevant datastore. Can be one of ` + "`" + `disabled` + "`" + `, ` + "`" + `warning` + "`" + `, or ` + "`" + `restartAggressive` + "`" + `. Default: ` + "`" + `disabled` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_apd_response",
					Description: `(Optional) Controls the action to take on virtual machines when the cluster has detected loss to all paths to a relevant datastore. Can be one of ` + "`" + `disabled` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `restartConservative` + "`" + `, or ` + "`" + `restartAggressive` + "`" + `. Default: ` + "`" + `disabled` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_apd_recovery_action",
					Description: `(Optional) Controls the action to take on virtual machines if an APD status on an affected datastore clears in the middle of an APD event. Can be one of ` + "`" + `none` + "`" + ` or ` + "`" + `reset` + "`" + `. Default: ` + "`" + `none` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_apd_response_delay",
					Description: `(Optional) The time, in seconds, to wait after an APD timeout event to run the response action defined in [` + "`" + `ha_datastore_apd_response` + "`" + `](#ha_datastore_apd_response). Default: ` + "`" + `180` + "`" + ` seconds (3 minutes). <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_vm_monitoring",
					Description: `(Optional) The type of virtual machine monitoring to use when HA is enabled in the cluster. Can be one of ` + "`" + `vmMonitoringDisabled` + "`" + `, ` + "`" + `vmMonitoringOnly` + "`" + `, or ` + "`" + `vmAndAppMonitoring` + "`" + `. Default: ` + "`" + `vmMonitoringDisabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_failure_interval",
					Description: `(Optional) The time interval, in seconds, a heartbeat from a virtual machine is not received within this configured interval, the virtual machine is marked as failed. Default: ` + "`" + `30` + "`" + ` seconds.`,
				},
				resource.Attribute{
					Name:        "ha_vm_minimum_uptime",
					Description: `(Optional) The time, in seconds, that HA waits after powering on a virtual machine before monitoring for heartbeats. Default: ` + "`" + `120` + "`" + ` seconds (2 minutes).`,
				},
				resource.Attribute{
					Name:        "ha_vm_maximum_resets",
					Description: `(Optional) The maximum number of resets that HA will perform to a virtual machine when responding to a failure event. Default: ` + "`" + `3` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ha_vm_maximum_failure_window",
					Description: `(Optional) The time, in seconds, for the reset window in which [` + "`" + `ha_vm_maximum_resets` + "`" + `](#ha_vm_maximum_resets) can operate. When this window expires, no more resets are attempted regardless of the setting configured in ` + "`" + `ha_vm_maximum_resets` + "`" + `. ` + "`" + `-1` + "`" + ` means no window, meaning an unlimited reset time is allotted. Default: ` + "`" + `-1` + "`" + ` (no window). #### vSphere HA Admission Control Settings The following settings control vSphere HA Admission Control, which controls whether or not specific VM operations are permitted in the cluster in order to protect the reliability of the cluster. Based on the constraints defined in these settings, operations such as power on or migration operations may be blocked to ensure that enough capacity remains to react to host failures. #### Admission Control Modes The [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy) parameter controls the specific mode that Admission Control uses. What settings are available depends on the admission control mode:`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_policy",
					Description: `(Optional) The type of admission control policy to use with vSphere HA. Can be one of ` + "`" + `resourcePercentage` + "`" + `, ` + "`" + `slotPolicy` + "`" + `, ` + "`" + `failoverHosts` + "`" + `, or ` + "`" + `disabled` + "`" + `. Default: ` + "`" + `resourcePercentage` + "`" + `. #### Common Admission Control Settings The following settings are available for all Admission Control modes, but will infer different meanings in each mode.`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_host_failure_tolerance",
					Description: `(Optional) The maximum number of failed hosts that admission control tolerates when making decisions on whether to permit virtual machine operations. The maximum is one less than the number of hosts in the cluster. Default: ` + "`" + `1` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_performance_tolerance",
					Description: `(Optional) The percentage of resource reduction that a cluster of virtual machines can tolerate in case of a failover. A value of 0 produces warnings only, whereas a value of 100 disables the setting. Default: ` + "`" + `100` + "`" + ` (disabled). #### Admission Control Settings for Resource Percentage Mode The following settings control specific settings for Admission Control when ` + "`" + `resourcePercentage` + "`" + ` is selected in [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy).`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_resource_percentage_auto_compute",
					Description: `(Optional) Automatically determine available resource percentages by subtracting the average number of host resources represented by the [` + "`" + `ha_admission_control_host_failure_tolerance` + "`" + `](#ha_admission_control_host_failure_tolerance) setting from the total amount of resources in the cluster. Disable to supply user-defined values. Default: ` + "`" + `true` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_resource_percentage_cpu",
					Description: `(Optional) Controls the user-defined percentage of CPU resources in the cluster to reserve for failover. Default: ` + "`" + `100` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_resource_percentage_memory",
					Description: `(Optional) Controls the user-defined percentage of memory resources in the cluster to reserve for failover. Default: ` + "`" + `100` + "`" + `. #### Admission Control Settings for Slot Policy Mode The following settings control specific settings for Admission Control when ` + "`" + `slotPolicy` + "`" + ` is selected in [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy).`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_slot_policy_use_explicit_size",
					Description: `(Optional) Controls whether or not you wish to supply explicit values to CPU and memory slot sizes. The default is ` + "`" + `false` + "`" + `, which tells vSphere to gather a automatic average based on all powered-on virtual machines currently in the cluster.`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_slot_policy_explicit_cpu",
					Description: `(Optional) Controls the user-defined CPU slot size, in MHz. Default: ` + "`" + `32` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_slot_policy_explicit_memory",
					Description: `(Optional) Controls the user-defined memory slot size, in MB. Default: ` + "`" + `100` + "`" + `. #### Admission Control Settings for Dedicated Failover Hosts Mode The following settings control specific settings for Admission Control when ` + "`" + `failoverHosts` + "`" + ` is selected in [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy).`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_failover_host_system_ids",
					Description: `(Optional) Defines the [managed object IDs][docs-about-morefs] of hosts to use as dedicated failover hosts. These hosts are kept as available as possible - admission control will block access to the host, and DRS will ignore the host when making recommendations. #### vSphere HA Datastore Settings vSphere HA uses datastore heartbeating to determine the health of a particular host. Depending on how your datastores are configured, the settings below may need to be altered to ensure that specific datastores are used over others. If you require a user-defined list of datastores, ensure you select either ` + "`" + `userSelectedDs` + "`" + ` (for user selected only) or ` + "`" + `allFeasibleDsWithUserPreference` + "`" + ` (for automatic selection with preferred overrides) for the [` + "`" + `ha_heartbeat_datastore_policy` + "`" + `](#ha_heartbeat_datastore_policy) setting.`,
				},
				resource.Attribute{
					Name:        "ha_heartbeat_datastore_policy",
					Description: `(Optional) The selection policy for HA heartbeat datastores. Can be one of ` + "`" + `allFeasibleDs` + "`" + `, ` + "`" + `userSelectedDs` + "`" + `, or ` + "`" + `allFeasibleDsWithUserPreference` + "`" + `. Default: ` + "`" + `allFeasibleDsWithUserPreference` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_heartbeat_datastore_ids",
					Description: `(Optional) The list of managed object IDs for preferred datastores to use for HA heartbeating. This setting is only useful when [` + "`" + `ha_heartbeat_datastore_policy` + "`" + `](#ha_heartbeat_datastore_policy) is set to either ` + "`" + `userSelectedDs` + "`" + ` or ` + "`" + `allFeasibleDsWithUserPreference` + "`" + `. #### Proactive HA Settings The following settings pertain to [Proactive HA][ref-vsphere-proactive-ha], an advanced feature of vSphere HA that allows the cluster to get data from external providers and make decisions based on the data reported. Working with Proactive HA is outside the scope of this document. For more details, see the referenced link in the above paragraph. [ref-vsphere-proactive-ha]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.avail.doc/GUID-3E3B18CC-8574-46FA-9170-CF549B8E55B8.html`,
				},
				resource.Attribute{
					Name:        "proactive_ha_enabled",
					Description: `(Optional) Enables Proactive HA. Default: ` + "`" + `false` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "proactive_ha_automation_level",
					Description: `(Optional) Determines how the host quarantine, maintenance mode, or virtual machine migration recommendations made by proactive HA are to be handled. Can be one of ` + "`" + `Automated` + "`" + ` or ` + "`" + `Manual` + "`" + `. Default: ` + "`" + `Manual` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "proactive_ha_moderate_remediation",
					Description: `(Optional) The configured remediation for moderately degraded hosts. Can be one of ` + "`" + `MaintenanceMode` + "`" + ` or ` + "`" + `QuarantineMode` + "`" + `. Note that this cannot be set to ` + "`" + `MaintenanceMode` + "`" + ` when [` + "`" + `proactive_ha_severe_remediation` + "`" + `](#proactive_ha_severe_remediation) is set to ` + "`" + `QuarantineMode` + "`" + `. Default: ` + "`" + `QuarantineMode` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "proactive_ha_severe_remediation",
					Description: `(Optional) The configured remediation for severely degraded hosts. Can be one of ` + "`" + `MaintenanceMode` + "`" + ` or ` + "`" + `QuarantineMode` + "`" + `. Note that this cannot be set to ` + "`" + `QuarantineMode` + "`" + ` when [` + "`" + `proactive_ha_moderate_remediation` + "`" + `](#proactive_ha_moderate_remediation) is set to ` + "`" + `MaintenanceMode` + "`" + `. Default: ` + "`" + `QuarantineMode` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "proactive_ha_provider_ids",
					Description: `(Optional) The list of IDs for health update providers configured for this cluster. <sup>[\`,
				},
				resource.Attribute{
					Name:        "vsan_enabled",
					Description: `(Optional) Enables vSAN on the cluster.`,
				},
				resource.Attribute{
					Name:        "vsan_dedup_enabled",
					Description: `(Optional) Enables vSAN deduplication on the cluster. Cannot be independently set to true. When vSAN deduplication is enabled, vSAN compression must also be enabled.`,
				},
				resource.Attribute{
					Name:        "vsan_compression_enabled",
					Description: `(Optional) Enables vSAN compression on the cluster.`,
				},
				resource.Attribute{
					Name:        "vsan_performance_enabled",
					Description: `(Optional) Enables vSAN performance service on the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsan_verbose_mode_enabled",
					Description: `(Optional) Enables verbose mode for vSAN performance service on the cluster.`,
				},
				resource.Attribute{
					Name:        "vsan_network_diagnostic_mode_enabled",
					Description: `(Optional) Enables network diagnostic mode for vSAN performance service on the cluster.`,
				},
				resource.Attribute{
					Name:        "vsan_unmap_enabled",
					Description: `(Optional) Enables vSAN unmap on the cluster.`,
				},
				resource.Attribute{
					Name:        "vsan_remote_datastore_ids",
					Description: `(Optional) The remote vSAN datastore IDs to be mounted to this cluster. Conflicts with ` + "`" + `vsan_dit_encryption_enabled` + "`" + ` and ` + "`" + `vsan_dit_rekey_interval` + "`" + `, i.e., vSAN HCI Mesh feature cannot be enabled with data-in-transit encryption feature at the same time.`,
				},
				resource.Attribute{
					Name:        "vsan_dit_encryption_enabled",
					Description: `(Optional) Enables vSAN data-in-transit encryption on the cluster. Conflicts with ` + "`" + `vsan_remote_datastore_ids` + "`" + `, i.e., vSAN data-in-transit feature cannot be enabled with the vSAN HCI Mesh feature at the same time.`,
				},
				resource.Attribute{
					Name:        "vsan_dit_rekey_interval",
					Description: `(Optional) Indicates the rekey interval in minutes for data-in-transit encryption. The valid rekey interval is 30 to 10800 (feature defaults to 1440). Conflicts with ` + "`" + `vsan_remote_datastore_ids` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vsan_disk_group",
					Description: `(Optional) Represents the configuration of a host disk group in the cluster.`,
				},
				resource.Attribute{
					Name:        "cache",
					Description: `The canonical name of the disk to use for vSAN cache.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `An array of disk canonical names for vSAN storage. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_host_group",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere cluster virtual machine group. This can be used to manage groups of virtual machines for relevant rules in a cluster.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the host group. This must be unique in the cluster. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "host_system_ids",
					Description: `(Optional) The [managed object IDs][docs-about-morefs] of the hosts to put in the cluster. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_vm_affinity_rule",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides the VMware vSphere Distributed Resource Scheduler affinity rule resource.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
				"vm",
				"affinity",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule. This must be unique in the cluster.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `(Required) The UUIDs of the virtual machines to run on the same host together.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable this rule in the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `(Optional) When this value is ` + "`" + `true` + "`" + `, prevents any virtual machine operations that may violate this rule. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_vm_anti_affinity_rule",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides the VMware vSphere Distributed Resource Scheduler anti-affinity rule resource.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
				"vm",
				"anti",
				"affinity",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule. This must be unique in the cluster.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `(Required) The UUIDs of the virtual machines to run on hosts different from each other.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable this rule in the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `(Optional) When this value is ` + "`" + `true` + "`" + `, prevents any virtual machine operations that may violate this rule. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_vm_dependency_rule",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere cluster VM dependency rule. This can be used to manage VM dependency rules for vSphere HA.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
				"vm",
				"dependency",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule. This must be unique in the cluster.`,
				},
				resource.Attribute{
					Name:        "dependency_vm_group_name",
					Description: `(Required) The name of the VM group that this rule depends on. The VMs defined in the group specified by [` + "`" + `vm_group_name` + "`" + `](#vm_group_name) will not be started until the VMs in this group are started.`,
				},
				resource.Attribute{
					Name:        "vm_group_name",
					Description: `(Required) The name of the VM group that is the subject of this rule. The VMs defined in this group will not be started until the VMs in the group specified by [` + "`" + `dependency_vm_group_name` + "`" + `](#dependency_vm_group_name) are started.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable this rule in the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `(Optional) When this value is ` + "`" + `true` + "`" + `, prevents any virtual machine operations that may violate this rule. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_vm_group",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere cluster virtual machine group. This can be used to manage groups of virtual machines for relevant rules in a cluster.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
				"vm",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the VM group. This must be unique in the cluster. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `(Required) The UUIDs of the virtual machines in this group. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_vm_host_rule",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere cluster VM/host rule. This can be used to manage VM-to-host affinity and anti-affinity rules.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"compute",
				"vm",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule. This must be unique in the cluster.`,
				},
				resource.Attribute{
					Name:        "vm_group_name",
					Description: `(Required) The name of the virtual machine group to use with this rule.`,
				},
				resource.Attribute{
					Name:        "affinity_host_group_name",
					Description: `(Optional) When this field is used, the virtual machines defined in [` + "`" + `vm_group_name` + "`" + `](#vm_group_name) will be run on the hosts defined in this host group.`,
				},
				resource.Attribute{
					Name:        "anti_affinity_host_group_name",
					Description: `(Optional) When this field is used, the virtual machines defined in [` + "`" + `vm_group_name` + "`" + `](#vm_group_name) will _not_ be run on the hosts defined in this host group.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable this rule in the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `(Optional) When this value is ` + "`" + `true` + "`" + `, prevents any virtual machine operations that may violate this rule. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_content_library",
			Category:         "Virtual Machine",
			ShortDescription: `Provides a vSphere content cibrary. Content libraries allow you to manage and share virtual machines, vApp templates, and other types of files. Content libraries enable you to share content across vCenter Server instances in the same or different locations.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
				"content",
				"library",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the content library.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the content library.`,
				},
				resource.Attribute{
					Name:        "storage_backing",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the datastore on which to store the content library items.`,
				},
				resource.Attribute{
					Name:        "publication",
					Description: `(Optional) Options to publish a local content library.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Method to authenticate users. Must be ` + "`" + `NONE` + "`" + ` or ` + "`" + `BASIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username used by subscribers to authenticate. Currently can only be ` + "`" + `vcsp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password used by subscribers to authenticate.`,
				},
				resource.Attribute{
					Name:        "published",
					Description: `(Optional) Publish the content library. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "subscription",
					Description: `(Optional) Options subscribe to a published content library.`,
				},
				resource.Attribute{
					Name:        "subscription_url",
					Description: `(Required) URL of the published content library.`,
				},
				resource.Attribute{
					Name:        "authentication_method",
					Description: `(Optional) Authentication method to connect ro a published content library. Must be ` + "`" + `NONE` + "`" + ` or ` + "`" + `BASIC` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Optional) Username used for authentication.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Optional) Password used for authentication.`,
				},
				resource.Attribute{
					Name:        "automatic_sync",
					Description: `(Optional) Enable automatic synchronization with the published library. Default ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "on_demand",
					Description: `(Optional) Download the library from a content only when needed. Default ` + "`" + `true` + "`" + `. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "publish_url",
					Description: `The URL of the published content library. ## Importing An existing content library can be [imported][docs-import] into this resource by supplying the content library ID. For example: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_content_library publisher_content_library f42a4b25-844a-44ec-9063-a3a5e9cc88c7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "publish_url",
					Description: `The URL of the published content library. ## Importing An existing content library can be [imported][docs-import] into this resource by supplying the content library ID. For example: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_content_library publisher_content_library f42a4b25-844a-44ec-9063-a3a5e9cc88c7 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_content_library_item",
			Category:         "Virtual Machine",
			ShortDescription: `Creates an item in a vSphere content library. Each item can contain multiple files.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
				"content",
				"library",
				"item",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the item to be created in the content library.`,
				},
				resource.Attribute{
					Name:        "library_id",
					Description: `(Required) The ID of the content library in which to create the item.`,
				},
				resource.Attribute{
					Name:        "file_url",
					Description: `(Optional) File to import as the content library item.`,
				},
				resource.Attribute{
					Name:        "source_uuid",
					Description: `(Optional) Virtual machine UUID to clone to content library.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the content library item.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) Type of content library item. One of "ovf", "iso", or "vm-template". Default: ` + "`" + `ovf` + "`" + `. ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is a combination of the [managed object reference ID][docs-about-morefs] of the cluster, and the name of the virtual machine group. ## Importing An existing content library item can be [imported][docs-import] into this resource by supplying the content library ID. An example is below: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_content_library_item iso-linux-ubuntu-server-lts xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_custom_attribute",
			Category:         "Inventory",
			ShortDescription: `Provides a VMware vSphere custom attribute resource. This can be used to manage custom attributes in vSphere.`,
			Description:      ``,
			Keywords: []string{
				"inventory",
				"custom",
				"attribute",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the custom attribute.`,
				},
				resource.Attribute{
					Name:        "managed_object_type",
					Description: `(Optional) The object type that this attribute may be applied to. If not set, the custom attribute may be applied to any object type. For a full list, review the [Managed Object Types](#managed-object-types). Forces a new resource if changed. ## Managed Object Types The following table provides the managed object values for which an attribute may apply. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datacenter",
			Category:         "Inventory",
			ShortDescription: `Provides a VMware vSphere datacenter resource. This can be used as the primary container of inventory objects such as hosts and virtual machines.`,
			Description:      ``,
			Keywords: []string{
				"inventory",
				"datacenter",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the datacenter. This name needs to be unique within the folder. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The folder where the datacenter should be created. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) Map of custom attribute ids to value strings to set for datacenter resource. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The name of this datacenter. This will be changed to the [managed object ID][docs-about-morefs] in v2.0.`,
				},
				resource.Attribute{
					Name:        "moid",
					Description: `[Managed object ID][docs-about-morefs] of this datacenter. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Importing An existing datacenter can be [imported][docs-import] into this resource via supplying the full path to the datacenter. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_datacenter.dc /dc1 ` + "`" + `` + "`" + `` + "`" + ` The above would import the datacenter named ` + "`" + `dc1` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The name of this datacenter. This will be changed to the [managed object ID][docs-about-morefs] in v2.0.`,
				},
				resource.Attribute{
					Name:        "moid",
					Description: `[Managed object ID][docs-about-morefs] of this datacenter. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Importing An existing datacenter can be [imported][docs-import] into this resource via supplying the full path to the datacenter. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_datacenter.dc /dc1 ` + "`" + `` + "`" + `` + "`" + ` The above would import the datacenter named ` + "`" + `dc1` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datastore_cluster",
			Category:         "Storage",
			ShortDescription: `Provides a vSphere datastore cluster resource. This can be used to create and manage datastore clusters.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"datastore",
				"cluster",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the datastore cluster.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the datacenter to create the datastore cluster in. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The relative path to a folder to put this datastore cluster in. This is a path relative to the datacenter you are deploying the datastore to. Example: for the ` + "`" + `dc1` + "`" + ` datacenter, and a provided ` + "`" + `folder` + "`" + ` of ` + "`" + `foo/bar` + "`" + `, Terraform will place a datastore cluster named ` + "`" + `terraform-datastore-cluster-test` + "`" + ` in a datastore folder located at ` + "`" + `/dc1/datastore/foo/bar` + "`" + `, with the final inventory path being ` + "`" + `/dc1/datastore/foo/bar/terraform-datastore-cluster-test` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sdrs_enabled",
					Description: `(Optional) Enable Storage DRS for this datastore cluster. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) A map of custom attribute ids to attribute value strings to set for the datastore cluster. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "sdrs_automation_level",
					Description: `(Optional) The global automation level for all virtual machines in this datastore cluster. Default: ` + "`" + `manual` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sdrs_space_balance_automation_level",
					Description: `(Optional) Overrides the default automation settings when correcting disk space imbalances.`,
				},
				resource.Attribute{
					Name:        "sdrs_io_balance_automation_level",
					Description: `(Optional) Overrides the default automation settings when correcting I/O load imbalances.`,
				},
				resource.Attribute{
					Name:        "sdrs_rule_enforcement_automation_level",
					Description: `(Optional) Overrides the default automation settings when correcting affinity rule violations.`,
				},
				resource.Attribute{
					Name:        "sdrs_policy_enforcement_automation_level",
					Description: `(Optional) Overrides the default automation settings when correcting storage and VM policy violations.`,
				},
				resource.Attribute{
					Name:        "sdrs_vm_evacuation_automation_level",
					Description: `(Optional) Overrides the default automation settings when generating recommendations for datastore evacuation. ### Storage DRS I/O load balancing settings The following options control I/O load balancing for Storage DRS on the datastore cluster. ~>`,
				},
				resource.Attribute{
					Name:        "sdrs_io_load_balance_enabled",
					Description: `(Optional) Enable I/O load balancing for this datastore cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sdrs_io_latency_threshold",
					Description: `(Optional) The I/O latency threshold, in milliseconds, that storage DRS uses to make recommendations to move disks from this datastore. Default: ` + "`" + `15` + "`" + ` seconds.`,
				},
				resource.Attribute{
					Name:        "sdrs_io_load_imbalance_threshold",
					Description: `(Optional) The difference between load in datastores in the cluster before storage DRS makes recommendations to balance the load. Default: ` + "`" + `5` + "`" + ` percent.`,
				},
				resource.Attribute{
					Name:        "sdrs_io_reservable_iops_threshold",
					Description: `(Optional) The threshold of reservable IOPS of all virtual machines on the datastore before storage DRS makes recommendations to move VMs off of a datastore. Note that this setting should only be set if ` + "`" + `sdrs_io_reservable_percent_threshold` + "`" + ` cannot make an accurate estimate of the capacity of the datastores in your cluster, and should be set to roughly 50-60% of the worst case peak performance of the backing LUNs.`,
				},
				resource.Attribute{
					Name:        "sdrs_io_reservable_percent_threshold",
					Description: `(Optional) The threshold, in percent, of actual estimated performance of the datastore (in IOPS) that storage DRS uses to make recommendations to move VMs off of a datastore when the total reservable IOPS exceeds the threshold. Default: ` + "`" + `60` + "`" + ` percent.`,
				},
				resource.Attribute{
					Name:        "sdrs_io_reservable_threshold_mode",
					Description: `(Optional) The reservable IOPS threshold setting to use, ` + "`" + `sdrs_io_reservable_percent_threshold` + "`" + ` in the event of ` + "`" + `automatic` + "`" + `, or ` + "`" + `sdrs_io_reservable_iops_threshold` + "`" + ` in the event of ` + "`" + `manual` + "`" + `. Default: ` + "`" + `automatic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sdrs_space_utilization_threshold",
					Description: `(Optional) Runtime thresholds govern when Storage DRS performs or recommends migrations (based on the selected automation level). Default: ` + "`" + `80` + "`" + ` percent. ### Storage DRS disk space load balancing settings The following options control disk space load balancing for Storage DRS on the datastore cluster. ~>`,
				},
				resource.Attribute{
					Name:        "sdrs_free_space_utilization_difference",
					Description: `(Optional) The threshold, in percent of used space, that storage DRS uses to make decisions to migrate VMs out of a datastore. Default: ` + "`" + `80` + "`" + ` percent.`,
				},
				resource.Attribute{
					Name:        "sdrs_free_space_utilization_difference",
					Description: `(Optional) The threshold, in percent, of difference between space utilization in datastores before storage DRS makes decisions to balance the space. Default: ` + "`" + `5` + "`" + ` percent.`,
				},
				resource.Attribute{
					Name:        "sdrs_free_space_threshold",
					Description: `(Optional) The threshold, in GB, that storage DRS uses to make decisions to migrate VMs out of a datastore. Default: ` + "`" + `50` + "`" + ` GB.`,
				},
				resource.Attribute{
					Name:        "sdrs_free_space_threshold",
					Description: `(Optional) The free space threshold to use. When set to ` + "`" + `utilization` + "`" + `, ` + "`" + `drs_space_utilization_threshold` + "`" + ` is used, and when set to ` + "`" + `freeSpace` + "`" + `, ` + "`" + `drs_free_space_threshold` + "`" + ` is used. Default: ` + "`" + `utilization` + "`" + `. ### Storage DRS advanced settings The following options control advanced parts of Storage DRS that may not require changing during basic operation:`,
				},
				resource.Attribute{
					Name:        "sdrs_default_intra_vm_affinity",
					Description: `(Optional) When ` + "`" + `true` + "`" + `, all disks in a single virtual machine will be kept on the same datastore. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sdrs_load_balance_interval",
					Description: `(Optional) The storage DRS poll interval, in minutes. Default: ` + "`" + `480` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "sdrs_advanced_options",
					Description: `(Optional) A key/value map of advanced Storage DRS settings that are not exposed via Terraform or the vSphere client. ## Attribute Reference The only computed attribute that is exported by this resource is the resource ` + "`" + `id` + "`" + `, which is the the [managed object reference ID][docs-about-morefs] of the datastore cluster. ## Importing An existing datastore cluster can be [imported][docs-import] into this resource via the path to the cluster, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_datastore_cluster.datastore_cluster /dc1/datastore/ds-cluster ` + "`" + `` + "`" + `` + "`" + ` The above would import the datastore cluster named ` + "`" + `ds-cluster` + "`" + ` that is located in the ` + "`" + `dc1` + "`" + ` datacenter.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datastore_cluster_vm_anti_affinity_rule",
			Category:         "Storage",
			ShortDescription: `Provides a VMware vSphere datastore cluster virtual machine anti-affinity rule. This can be used to manage rules to tell virtual machines to run on separate datastores.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"datastore",
				"cluster",
				"vm",
				"anti",
				"affinity",
				"rule",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datastore_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the datastore cluster to put the group in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the rule. This must be unique in the cluster.`,
				},
				resource.Attribute{
					Name:        "virtual_machine_ids",
					Description: `(Required) The UUIDs of the virtual machines to run on different datastores from each other. ~>`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `(Optional) Enable this rule in the cluster. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mandatory",
					Description: `(Optional) When this value is ` + "`" + `true` + "`" + `, prevents any virtual machine operations that may violate this rule. Default: ` + "`" + `false` + "`" + `. ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is a combination of the [managed object reference ID][docs-about-morefs] of the cluster, and the rule's key within the cluster configuration. ## Importing An existing rule can be [imported][docs-import] into this resource by supplying both the path to the cluster, and the name the rule. If the name or cluster is not found, or if the rule is of a different type, an error will be returned. An example is below: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_datastore_cluster_vm_anti_affinity_rule.cluster_vm_anti_affinity_rule \ '{"compute_cluster_path": "/dc1/datastore/cluster1", \ "name": "terraform-test-datastore-cluster-vm-anti-affinity-rule"}' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_distributed_port_group",
			Category:         "Networking",
			ShortDescription: `Provides a vSphere distributed port group resource. This can be used to create and manage port groups on a vSphere Distributed Switch.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"distributed",
				"port",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the port group.`,
				},
				resource.Attribute{
					Name:        "distributed_virtual_switch_uuid",
					Description: `(Required) The ID of the VDS to add the port group to. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The port group type. Can be one of ` + "`" + `earlyBinding` + "`" + ` (static binding) or ` + "`" + `ephemeral` + "`" + `. Default: ` + "`" + `earlyBinding` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) An optional description for the port group.`,
				},
				resource.Attribute{
					Name:        "number_of_ports",
					Description: `(Optional) The number of ports available on this port group. Cannot be decreased below the amount of used ports on the port group.`,
				},
				resource.Attribute{
					Name:        "auto_expand",
					Description: `(Optional) Allows the port group to create additional ports past the limit specified in ` + "`" + `number_of_ports` + "`" + ` if necessary. Default: ` + "`" + `true` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "port_name_format",
					Description: `(Optional) An optional formatting policy for naming of the ports in this port group. See the ` + "`" + `portNameFormat` + "`" + ` attribute listed [here][ext-vsphere-portname-format] for details on the format syntax. [ext-vsphere-portname-format]: https://vdc-download.vmware.com/vmwb-repository/dcr-public/b50dcbbf-051d-4204-a3e7-e1b618c1e384/538cf2ec-b34f-4bae-a332-3820ef9e7773/vim.dvs.DistributedVirtualPortgroup.ConfigInfo.html#portNameFormat`,
				},
				resource.Attribute{
					Name:        "network_resource_pool_key",
					Description: `(Optional) The key of a network resource pool to associate with this port group. The default is ` + "`" + `-1` + "`" + `, which implies no association.`,
				},
				resource.Attribute{
					Name:        "block_override_allowed",
					Description: `(Optional) Allow the [port shutdown policy][port-shutdown-policy] to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "live_port_moving_allowed",
					Description: `(Optional) Allow a port in this port group to be moved to another port group while it is connected.`,
				},
				resource.Attribute{
					Name:        "netflow_override_allowed",
					Description: `(Optional) Allow the [Netflow policy][netflow-policy] on this port group to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "network_resource_pool_override_allowed",
					Description: `(Optional) Allow the network resource pool set on this port group to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "port_config_reset_at_disconnect",
					Description: `(Optional) Reset a port's settings to the settings defined on this port group policy when the port disconnects.`,
				},
				resource.Attribute{
					Name:        "security_policy_override_allowed",
					Description: `(Optional) Allow the [security policy settings][sec-policy-settings] defined in this port group policy to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "shaping_override_allowed",
					Description: `(Optional) Allow the [traffic shaping options][traffic-shaping-settings] on this port group policy to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "traffic_filter_override_allowed",
					Description: `(Optional) Allow any traffic filters on this port group to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "uplink_teaming_override_allowed",
					Description: `(Optional) Allow the [uplink teaming options][uplink-teaming-settings] on this port group to be overridden on an individual port.`,
				},
				resource.Attribute{
					Name:        "vlan_override_allowed",
					Description: `(Optional) Allow the [VLAN settings][vlan-settings] on this port group to be overridden on an individual port. [port-shutdown-policy]: /docs/providers/vsphere/r/distributed_virtual_switch.html#block_all_ports [netflow-policy]: /docs/providers/vsphere/r/distributed_virtual_switch.html#netflow_enabled [sec-policy-settings]: /docs/providers/vsphere/r/distributed_virtual_switch.html#security-options [traffic-shaping-settings]: /docs/providers/vsphere/r/distributed_virtual_switch.html#traffic-shaping-options [uplink-teaming-settings]: /docs/providers/vsphere/r/distributed_virtual_switch.html#ha-policy-options [vlan-settings]: /docs/providers/vsphere/r/distributed_virtual_switch.html#vlan-options ## Attribute Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_distributed_virtual_switch",
			Category:         "Networking",
			ShortDescription: `Provides a vSphere Distributed Switch resource. This can be used to create and manage the vSphere Distributed Switch resources in vCenter Server.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"distributed",
				"virtual",
				"switch",
			},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_dpm_host_override",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere DPM host override resource. This can be used to override power management settings for a host in a cluster.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"dpm",
				"override",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the override in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "host_system_ids",
					Description: `(Optional) The [managed object ID][docs-about-morefs] of the host to create the override for.`,
				},
				resource.Attribute{
					Name:        "dpm_enabled",
					Description: `(Optional) Enable DPM support for this host. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "dpm_automation_level",
					Description: `(Optional) The automation level for host power operations on this host. Can be one of ` + "`" + `manual` + "`" + ` or ` + "`" + `automated` + "`" + `. Default: ` + "`" + `manual` + "`" + `. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_drs_vm_override",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere DRS virtual machine override resource. This can be used to override DRS settings in a cluster.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"drs",
				"vm",
				"override",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the override in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The UUID of the virtual machine to create the override for. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "drs_enabled",
					Description: `(Optional) Overrides the default DRS setting for this virtual machine. Can be either ` + "`" + `true` + "`" + ` or ` + "`" + `false` + "`" + `. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "drs_automation_level",
					Description: `(Optional) Overrides the automation level for this virtual machine in the cluster. Can be one of ` + "`" + `manual` + "`" + `, ` + "`" + `partiallyAutomated` + "`" + `, or ` + "`" + `fullyAutomated` + "`" + `. Default: ` + "`" + `manual` + "`" + `. ->`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_file",
			Category:         "Storage",
			ShortDescription: `Provides a VMware vSphere file resource. This can be used to upload files (e.g. .iso and .vmdk) from the Terraform host machine to a remote vSphere or copy files within vSphere.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_file",
					Description: `(Required) The path to the file being uploaded from the Terraform host to the vSphere environment or copied within vSphere environment. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "destination_file",
					Description: `(Required) The path to where the file should be uploaded or copied to on the destination ` + "`" + `datastore` + "`" + ` in vSphere.`,
				},
				resource.Attribute{
					Name:        "source_datacenter",
					Description: `(Optional) The name of a datacenter from which the file will be copied. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The name of a datacenter to which the file will be uploaded.`,
				},
				resource.Attribute{
					Name:        "source_datastore",
					Description: `(Optional) The name of the datastore from which file will be copied. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) The name of the datastore to which to upload the file.`,
				},
				resource.Attribute{
					Name:        "create_directories",
					Description: `(Optional) Create directories in ` + "`" + `destination_file` + "`" + ` path parameter on first apply if any are missing for copy operation. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_folder",
			Category:         "Inventory",
			ShortDescription: `Provides a VMware vSphere folder resource. This can be used to manage vSphere inventory folders.`,
			Description:      ``,
			Keywords: []string{
				"inventory",
				"folder",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "path",
					Description: `(Required) The path of the folder to be created. This is relative to the root of the type of folder you are creating, and the supplied datacenter. For example, given a default datacenter of ` + "`" + `default-dc` + "`" + `, a folder of type ` + "`" + `vm` + "`" + ` (denoting a virtual machine folder), and a supplied folder of ` + "`" + `terraform-test-folder` + "`" + `, the resulting path would be ` + "`" + `/default-dc/vm/terraform-test-folder` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Required) The type of folder to create. Allowed options are ` + "`" + `datacenter` + "`" + ` for datacenter folders, ` + "`" + `host` + "`" + ` for host and cluster folders, ` + "`" + `vm` + "`" + ` for virtual machine folders, ` + "`" + `datastore` + "`" + ` for datastore folders, and ` + "`" + `network` + "`" + ` for network folders. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `The ID of the datacenter the folder will be created in. Required for all folder types except for datacenter folders. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) Map of custom attribute ids to attribute value strings to set for folder. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_ha_vm_override",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere HA virtual machine override resource. This can be used to override high availability settings in a cluster.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"ha",
				"vm",
				"override",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "compute_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the cluster to put the override in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The UUID of the virtual machine to create the override for. Forces a new resource if changed. ### vSphere HA Options The following settings work nearly in the same fashion as their counterparts in the [` + "`" + `vsphere_compute_cluster` + "`" + `][tf-vsphere-cluster-resource] resource, with the exception that some options also allow settings that denote the use of cluster defaults. See the individual settings below for more details. [tf-vsphere-cluster-resource]: /docs/providers/vsphere/r/compute_cluster.html ~>`,
				},
				resource.Attribute{
					Name:        "ha_vm_restart_priority",
					Description: `(Optional) The restart priority for the virtual machine when vSphere detects a host failure. Can be one of ` + "`" + `clusterRestartPriority` + "`" + `, ` + "`" + `lowest` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `, ` + "`" + `highest` + "`" + `, or ` + "`" + `disabled` + "`" + `. Default: ` + "`" + `clusterRestartPriority` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_restart_timeout",
					Description: `(Optional) The maximum time, in seconds, that vSphere HA will wait for this virtual machine to be ready. Use ` + "`" + `-1` + "`" + ` to specify the cluster default. Default: ` + "`" + `-1` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_host_isolation_response",
					Description: `(Optional) The action to take on this virtual machine when a host has detected that it has been isolated from the rest of the cluster. Can be one of ` + "`" + `clusterIsolationResponse` + "`" + `, ` + "`" + `none` + "`" + `, ` + "`" + `powerOff` + "`" + `, or ` + "`" + `shutdown` + "`" + `. Default: ` + "`" + `clusterIsolationResponse` + "`" + `. #### HA Virtual Machine Component Protection settings The following settings control Virtual Machine Component Protection (VMCP) overrides.`,
				},
				resource.Attribute{
					Name:        "ha_datastore_pdl_response",
					Description: `(Optional) Controls the action to take on this virtual machine when the cluster has detected a permanent device loss to a relevant datastore. Can be one of ` + "`" + `clusterDefault` + "`" + `, ` + "`" + `disabled` + "`" + `, ` + "`" + `warning` + "`" + `, or ` + "`" + `restartAggressive` + "`" + `. Default: ` + "`" + `clusterDefault` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_apd_response",
					Description: `(Optional) Controls the action to take on this virtual machine when the cluster has detected loss to all paths to a relevant datastore. Can be one of ` + "`" + `clusterDefault` + "`" + `, ` + "`" + `disabled` + "`" + `, ` + "`" + `warning` + "`" + `, ` + "`" + `restartConservative` + "`" + `, or ` + "`" + `restartAggressive` + "`" + `. Default: ` + "`" + `clusterDefault` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_apd_recovery_action",
					Description: `(Optional) Controls the action to take on this virtual machine if an APD status on an affected datastore clears in the middle of an APD event. Can be one of ` + "`" + `useClusterDefault` + "`" + `, ` + "`" + `none` + "`" + ` or ` + "`" + `reset` + "`" + `. Default: ` + "`" + `useClusterDefault` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_datastore_apd_response_delay",
					Description: `(Optional) Controls the delay in seconds to wait after an APD timeout event to execute the response action defined in [` + "`" + `ha_datastore_apd_response` + "`" + `](#ha_datastore_apd_response). Use ` + "`" + `-1` + "`" + ` to use the cluster default. Default: ` + "`" + `-1` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_vm_monitoring_use_cluster_defaults",
					Description: `(Optional) Determines whether or not the cluster's default settings or the VM override settings specified in this resource are used for virtual machine monitoring. The default is ` + "`" + `true` + "`" + ` (use cluster defaults) - set to ` + "`" + `false` + "`" + ` to have overrides take effect.`,
				},
				resource.Attribute{
					Name:        "ha_vm_monitoring",
					Description: `(Optional) The type of virtual machine monitoring to use when HA is enabled in the cluster. Can be one of ` + "`" + `vmMonitoringDisabled` + "`" + `, ` + "`" + `vmMonitoringOnly` + "`" + `, or ` + "`" + `vmAndAppMonitoring` + "`" + `. Default: ` + "`" + `vmMonitoringDisabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_failure_interval",
					Description: `(Optional) If a heartbeat from this virtual machine is not received within this configured interval, the virtual machine is marked as failed. The value is in seconds. Default: ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_minimum_uptime",
					Description: `(Optional) The time, in seconds, that HA waits after powering on this virtual machine before monitoring for heartbeats. Default: ` + "`" + `120` + "`" + ` (2 minutes).`,
				},
				resource.Attribute{
					Name:        "ha_vm_maximum_resets",
					Description: `(Optional) The maximum number of resets that HA will perform to this virtual machine when responding to a failure event. Default: ` + "`" + `3` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ha_vm_maximum_failure_window",
					Description: `(Optional) The length of the reset window in which [` + "`" + `ha_vm_maximum_resets` + "`" + `](#ha_vm_maximum_resets) can operate. When this window expires, no more resets are attempted regardless of the setting configured in ` + "`" + `ha_vm_maximum_resets` + "`" + `. ` + "`" + `-1` + "`" + ` means no window, meaning an unlimited reset time is allotted. The value is specified in seconds. Default: ` + "`" + `-1` + "`" + ` (no window). ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is a combination of the [managed object reference ID][docs-about-morefs] of the cluster, and the UUID of the virtual machine. This is used to look up the override on subsequent plan and apply operations after the override has been created. ## Importing An existing override can be [imported][docs-import] into this resource by supplying both the path to the cluster, and the path to the virtual machine, to ` + "`" + `terraform import` + "`" + `. If no override exists, an error will be given. An example is below: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_ha_vm_override.ha_vm_override \ '{"compute_cluster_path": "/dc1/host/cluster1", \ "virtual_machine_path": "/dc1/vm/srv1"}' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_host",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere host resource. This represents an ESXi host that can be used as a member of a cluster or as a standalone host.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "hostname",
					Description: `(Required) FQDN or IP address of the host to be added.`,
				},
				resource.Attribute{
					Name:        "username",
					Description: `(Required) Username that will be used by vSphere to authenticate to the host.`,
				},
				resource.Attribute{
					Name:        "password",
					Description: `(Required) Password that will be used by vSphere to authenticate to the host.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The ID of the datacenter this host should be added to. This should not be set if ` + "`" + `cluster` + "`" + ` is set.`,
				},
				resource.Attribute{
					Name:        "cluster",
					Description: `(Optional) The ID of the Compute Cluster this host should be added to. This should not be set if ` + "`" + `datacenter` + "`" + ` is set. Conflicts with: ` + "`" + `cluster_managed` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cluster_managed",
					Description: `(Optional) Can be set to ` + "`" + `true` + "`" + ` if compute cluster membership will be managed through the ` + "`" + `compute_cluster` + "`" + ` resource rather than the` + "`" + `host` + "`" + ` resource. Conflicts with: ` + "`" + `cluster` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "thumbprint",
					Description: `(Optional) Host's certificate SHA-1 thumbprint. If not set the CA that signed the host's certificate should be trusted. If the CA is not trusted and no thumbprint is set then the operation will fail. See data source [` + "`" + `vsphere_host_thumbprint` + "`" + `][docs-host-thumbprint-data-source].`,
				},
				resource.Attribute{
					Name:        "license",
					Description: `(Optional) The license key that will be applied to the host. The license key is expected to be present in vSphere.`,
				},
				resource.Attribute{
					Name:        "force",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + ` then it will force the host to be added, even if the host is already connected to a different vCenter Server instance. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "connected",
					Description: `(Optional) If set to false then the host will be disconnected. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maintenance",
					Description: `(Optional) Set the management state of the host. Default is ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "lockdown",
					Description: `(Optional) Set the lockdown state of the host. Valid options are ` + "`" + `disabled` + "`" + `, ` + "`" + `normal` + "`" + `, and ` + "`" + `strict` + "`" + `. Default is ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. Please refer to the ` + "`" + `vsphere_tag` + "`" + ` resource for more information on applying tags to resources. ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) A map of custom attribute IDs and string values to apply to the resource. Please refer to the ` + "`" + `vsphere_custom_attributes` + "`" + ` resource for more information on applying tags to resources. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the host. ## Importing An existing host can be [imported][docs-import] into this resource by supplying the host's ID. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_host.esx-01 host-123 ` + "`" + `` + "`" + `` + "`" + ` The above would import the host with ID ` + "`" + `host-123` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the host. ## Importing An existing host can be [imported][docs-import] into this resource by supplying the host's ID. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_host.esx-01 host-123 ` + "`" + `` + "`" + `` + "`" + ` The above would import the host with ID ` + "`" + `host-123` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_host_port_group",
			Category:         "Networking",
			ShortDescription: `Provides a vSphere port group resource to manage port groups on ESXi hosts.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"host",
				"port",
				"group",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the port group. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the host to set the port group up on. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "virtual_switch_name",
					Description: `(Required) The name of the virtual switch to bind this port group to. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "vlan_id",
					Description: `(Optional) The VLAN ID/trunk mode for this port group. An ID of ` + "`" + `0` + "`" + ` denotes no tagging, an ID of ` + "`" + `1` + "`" + `-` + "`" + `4094` + "`" + ` tags with the specific ID, and an ID of ` + "`" + `4095` + "`" + ` enables trunk mode, allowing the guest to manage its own tagging. Default: ` + "`" + `0` + "`" + `. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ### Policy Options In addition to the above options, you can configure any policy option that is available under the [` + "`" + `vsphere_host_virtual_switch` + "`" + ` policy options section][host-vswitch-policy-options]. Any policy option that is not set is`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `An ID for the port group that is _unique_ to Terraform. The convention is a prefix, the host system ID, and the port group name. For example,` + "`" + `tf-HostPortGroup:host-10:portgroup-01` + "`" + `. Tracking a port group on a standard switch, which can be created with or without a vCenter Server, is different than then a dvPortGroup which is tracked as a managed object ID in vCemter Server versus a key on a host.`,
				},
				resource.Attribute{
					Name:        "computed_policy",
					Description: `A map with a full set of the [policy options][host-vswitch-policy-options] computed from defaults and overrides, explaining the effective policy for this port group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key for this port group as returned from the vSphere API.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A list of ports that currently exist and are used on this port group. ## Importing An existing host port group can be [imported][docs-import] into this resource using the host port group's ID. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_host_port_group.management tf-HostPortGroup:host-123:management ` + "`" + `` + "`" + `` + "`" + ` The above would import the ` + "`" + `management` + "`" + ` host port group from host with ID ` + "`" + `host-123` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An ID for the port group that is _unique_ to Terraform. The convention is a prefix, the host system ID, and the port group name. For example,` + "`" + `tf-HostPortGroup:host-10:portgroup-01` + "`" + `. Tracking a port group on a standard switch, which can be created with or without a vCenter Server, is different than then a dvPortGroup which is tracked as a managed object ID in vCemter Server versus a key on a host.`,
				},
				resource.Attribute{
					Name:        "computed_policy",
					Description: `A map with a full set of the [policy options][host-vswitch-policy-options] computed from defaults and overrides, explaining the effective policy for this port group.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The key for this port group as returned from the vSphere API.`,
				},
				resource.Attribute{
					Name:        "ports",
					Description: `A list of ports that currently exist and are used on this port group. ## Importing An existing host port group can be [imported][docs-import] into this resource using the host port group's ID. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_host_port_group.management tf-HostPortGroup:host-123:management ` + "`" + `` + "`" + `` + "`" + ` The above would import the ` + "`" + `management` + "`" + ` host port group from host with ID ` + "`" + `host-123` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_host_virtual_switch",
			Category:         "Networking",
			ShortDescription: `Provides a vSphere Host Virtual Switch Resource. This can be used to configure vSwitches direct on an ESXi host.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"host",
				"virtual",
				"switch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual switch. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the host to set the virtual switch up on. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) The maximum transmission unit (MTU) for the virtual switch. Default: ` + "`" + `1500` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "number_of_ports",
					Description: `(Optional) The number of ports to create with this virtual switch. Default: ` + "`" + `128` + "`" + `. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "network_adapters",
					Description: `(Required) The network interfaces to bind to the bridge.`,
				},
				resource.Attribute{
					Name:        "beacon_interval",
					Description: `(Optional) The interval, in seconds, that a NIC beacon packet is sent out. This can be used with [` + "`" + `check_beacon` + "`" + `](#check_beacon) to offer link failure capability beyond link status only. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_discovery_operation",
					Description: `(Optional) Whether to ` + "`" + `advertise` + "`" + ` or ` + "`" + `listen` + "`" + ` for link discovery traffic. Default: ` + "`" + `listen` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_discovery_protocol",
					Description: `(Optional) The discovery protocol type. Valid types are ` + "`" + `cpd` + "`" + ` and ` + "`" + `lldp` + "`" + `. Default: ` + "`" + `cdp` + "`" + `. ### Policy Options The following options relate to how network traffic is handled on this virtual switch. It also controls the NIC failover order. This subset of options is shared with the [` + "`" + `vsphere_host_port_group` + "`" + `][host-port-group] resource, in which options can be omitted to ensure options are inherited from the switch configuration here. #### NIC Teaming Options ~>`,
				},
				resource.Attribute{
					Name:        "active_nics",
					Description: `(Required) The list of active network adapters used for load balancing.`,
				},
				resource.Attribute{
					Name:        "standby_nics",
					Description: `(Optional) The list of standby network adapters used for failover.`,
				},
				resource.Attribute{
					Name:        "check_beacon",
					Description: `(Optional) Enable beacon probing - this requires that the [` + "`" + `beacon_interval` + "`" + `](#beacon_interval) option has been set in the bridge options. If this is set to ` + "`" + `false` + "`" + `, only link status is used to check for failed NICs. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "teaming_policy",
					Description: `(Optional) The network adapter teaming policy. Can be one of ` + "`" + `loadbalance_ip` + "`" + `, ` + "`" + `loadbalance_srcmac` + "`" + `, ` + "`" + `loadbalance_srcid` + "`" + `, or ` + "`" + `failover_explicit` + "`" + `. Default: ` + "`" + `loadbalance_srcid` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "notify_switches",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the teaming policy will notify the broadcast network of a NIC failover, triggering cache updates. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "failback",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the teaming policy will re-activate failed interfaces higher in precedence when they come back up. Default: ` + "`" + `true` + "`" + `. #### Security Policy Options`,
				},
				resource.Attribute{
					Name:        "allow_promiscuous",
					Description: `(Optional) Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_forged_transmits",
					Description: `(Optional) Controls whether or not the virtual network adapter is allowed to send network traffic with a different MAC address than that of its own. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "allow_mac_changes",
					Description: `(Optional) Controls whether or not the Media Access Control (MAC) address can be changed. Default: ` + "`" + `true` + "`" + `. #### Traffic Shaping Options`,
				},
				resource.Attribute{
					Name:        "shaping_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable the traffic shaper for ports managed by this virtual switch. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shaping_average_bandwidth",
					Description: `(Optional) The average bandwidth in bits per second if traffic shaping is enabled. Default: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "shaping_peak_bandwidth",
					Description: `(Optional) The peak bandwidth during bursts in bits per second if traffic shaping is enabled. Default: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "shaping_burst_size",
					Description: `(Optional) The maximum burst size allowed in bytes if shaping is enabled. Default: ` + "`" + `0` + "`" + ` ## Attribute Reference The only exported attribute, other than the attributes above, is the ` + "`" + `id` + "`" + ` of the resource. This is set to an ID value unique to Terraform - the convention is a prefix, the host system ID, and the virtual switch name. An example would be ` + "`" + `tf-HostVirtualSwitch:host-10:vSwitchTerraformTest` + "`" + `. ## Importing An existing vSwitch can be [imported][docs-import] into this resource by its ID. The convention of the id is a prefix, the host system [managed objectID][docs-about-morefs], and the virtual switch name. An example would be ` + "`" + `tf-HostVirtualSwitch:host-10:vSwitchTerraformTest` + "`" + `. Import can the be done via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_host_virtual_switch.switch tf-HostVirtualSwitch:host-10:vSwitchTerraformTest ` + "`" + `` + "`" + `` + "`" + ` The above would import the vSwtich named ` + "`" + `vSwitchTerraformTest` + "`" + ` that is located in the ` + "`" + `host-10` + "`" + ` vSphere host.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_license",
			Category:         "Administration",
			ShortDescription: `Provides a VMware vSphere license resource. This can be used to add and remove license keys.`,
			Description:      ``,
			Keywords: []string{
				"administration",
				"license",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "license_key",
					Description: `(Required) The license key to add.`,
				},
				resource.Attribute{
					Name:        "labels",
					Description: `(Optional) A map of key/value pairs to be attached as labels (tags) to the license key. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "edition_key",
					Description: `The product edition of the license key.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `Total number of units (example: CPUs) contained in the license.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `The number of units (example: CPUs) assigned to this license.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name for the license.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "edition_key",
					Description: `The product edition of the license key.`,
				},
				resource.Attribute{
					Name:        "total",
					Description: `Total number of units (example: CPUs) contained in the license.`,
				},
				resource.Attribute{
					Name:        "used",
					Description: `The number of units (example: CPUs) assigned to this license.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The display name for the license.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_nas_datastore",
			Category:         "Storage",
			ShortDescription: `Provides a vSphere NAS datastore resource. This can be used to mount a NFS share as a datastore on a host.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"nas",
				"datastore",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the datastore. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "host_system_ids",
					Description: `(Required) The [managed object IDs][docs-about-morefs] of the hosts to mount the datastore on.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of NAS volume. Can be one of ` + "`" + `NFS` + "`" + ` (to denote v3) or ` + "`" + `NFS41` + "`" + ` (to denote NFS v4.1). Default: ` + "`" + `NFS` + "`" + `. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "remote_hosts",
					Description: `(Required) The hostnames or IP addresses of the remote server or servers. Only one element should be present for NFS v3 but multiple can be present for NFS v4.1. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "remote_path",
					Description: `(Required) The remote path of the mount point. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "access_mode",
					Description: `(Optional) Access mode for the mount point. Can be one of ` + "`" + `readOnly` + "`" + ` or ` + "`" + `readWrite` + "`" + `. Note that ` + "`" + `readWrite` + "`" + ` does not necessarily mean that the datastore will be read-write depending on the permissions of the actual share. Default: ` + "`" + `readWrite` + "`" + `. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "security_type",
					Description: `(Optional) The security type to use when using NFS v4.1. Can be one of ` + "`" + `AUTH_SYS` + "`" + `, ` + "`" + `SEC_KRB5` + "`" + `, or ` + "`" + `SEC_KRB5I` + "`" + `. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The relative path to a folder to put this datastore in. This is a path relative to the datacenter you are deploying the datastore to. Example: for the ` + "`" + `dc1` + "`" + ` datacenter, and a provided ` + "`" + `folder` + "`" + ` of ` + "`" + `foo/bar` + "`" + `, Terraform will place a datastore named ` + "`" + `terraform-test` + "`" + ` in a datastore folder located at ` + "`" + `/dc1/datastore/foo/bar` + "`" + `, with the final inventory path being ` + "`" + `/dc1/datastore/foo/bar/terraform-test` + "`" + `. Conflicts with ` + "`" + `datastore_cluster_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datastore_cluster_id",
					Description: `(Optional) The [managed object ID][docs-about-morefs] of a datastore cluster to put this datastore in. Conflicts with ` + "`" + `folder` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) Map of custom attribute ids to attribute value strings to set on datasource resource. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The [managed object reference ID][docs-about-morefs] of the datastore.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `The connectivity status of the datastore. If this is ` + "`" + `false` + "`" + `, some other computed attributes may be out of date.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Maximum capacity of the datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "free_space",
					Description: `Available space of this datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "maintenance_mode",
					Description: `The current maintenance mode state of the datastore.`,
				},
				resource.Attribute{
					Name:        "multiple_host_access",
					Description: `If ` + "`" + `true` + "`" + `, more than one host in the datacenter has been configured with access to the datastore.`,
				},
				resource.Attribute{
					Name:        "uncommitted_space",
					Description: `Total additional storage space, in megabytes, potentially used by all virtual machines on this datastore.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The unique locator for the datastore.`,
				},
				resource.Attribute{
					Name:        "protocol_endpoint",
					Description: `Indicates that this NAS volume is a protocol endpoint. This field is only populated if the host supports virtual datastores. ## Importing An existing NAS datastore can be [imported][docs-import] into this resource via its managed object ID, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_nas_datastore.datastore datastore-123 ` + "`" + `` + "`" + `` + "`" + ` You need a tool like [` + "`" + `govc` + "`" + `][ext-govc] that can display managed object IDs. [ext-govc]: https://github.com/vmware/govmomi/tree/master/govc In the case of govc, you can locate a managed object ID from an inventory path by doing the following: ` + "`" + `` + "`" + `` + "`" + ` $ govc ls -i /dc/datastore/terraform-test Datastore:datastore-123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The [managed object reference ID][docs-about-morefs] of the datastore.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `The connectivity status of the datastore. If this is ` + "`" + `false` + "`" + `, some other computed attributes may be out of date.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Maximum capacity of the datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "free_space",
					Description: `Available space of this datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "maintenance_mode",
					Description: `The current maintenance mode state of the datastore.`,
				},
				resource.Attribute{
					Name:        "multiple_host_access",
					Description: `If ` + "`" + `true` + "`" + `, more than one host in the datacenter has been configured with access to the datastore.`,
				},
				resource.Attribute{
					Name:        "uncommitted_space",
					Description: `Total additional storage space, in megabytes, potentially used by all virtual machines on this datastore.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The unique locator for the datastore.`,
				},
				resource.Attribute{
					Name:        "protocol_endpoint",
					Description: `Indicates that this NAS volume is a protocol endpoint. This field is only populated if the host supports virtual datastores. ## Importing An existing NAS datastore can be [imported][docs-import] into this resource via its managed object ID, via the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_nas_datastore.datastore datastore-123 ` + "`" + `` + "`" + `` + "`" + ` You need a tool like [` + "`" + `govc` + "`" + `][ext-govc] that can display managed object IDs. [ext-govc]: https://github.com/vmware/govmomi/tree/master/govc In the case of govc, you can locate a managed object ID from an inventory path by doing the following: ` + "`" + `` + "`" + `` + "`" + ` $ govc ls -i /dc/datastore/terraform-test Datastore:datastore-123 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_resource_pool",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a resource for VMware vSphere resource pools. This can be used to create and manage resource pools.`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"resource",
				"pool",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the resource pool.`,
				},
				resource.Attribute{
					Name:        "parent_resource_pool_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the parent resource pool. This can be the root resource pool for a cluster or standalone host, or a resource pool itself. When moving a resource pool from one parent resource pool to another, both must share a common root resource pool.`,
				},
				resource.Attribute{
					Name:        "cpu_share_level",
					Description: `(Optional) The CPU allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. When ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `high` + "`" + ` are specified values in ` + "`" + `cpu_shares` + "`" + ` will be ignored. Default: ` + "`" + `normal` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_shares",
					Description: `(Optional) The number of shares allocated for CPU. Used to determine resource allocation in case of resource contention. If this is set, ` + "`" + `cpu_share_level` + "`" + ` must be ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_reservation",
					Description: `(Optional) Amount of CPU (MHz) that is guaranteed available to the resource pool. Default: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_expandable",
					Description: `(Optional) Determines if the reservation on a resource pool can grow beyond the specified value if the parent resource pool has unreserved resources. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `(Optional) The CPU utilization of a resource pool will not exceed this limit, even if there are available resources. Set to ` + "`" + `-1` + "`" + ` for unlimited. Default: ` + "`" + `-1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_share_level",
					Description: `(Optional) The CPU allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. When ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `high` + "`" + ` are specified values in ` + "`" + `memory_shares` + "`" + ` will be ignored. Default: ` + "`" + `normal` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_shares",
					Description: `(Optional) The number of shares allocated for CPU. Used to determine resource allocation in case of resource contention. If this is set, ` + "`" + `memory_share_level` + "`" + ` must be ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `(Optional) Amount of CPU (MHz) that is guaranteed available to the resource pool. Default: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_expandable",
					Description: `(Optional) Determines if the reservation on a resource pool can grow beyond the specified value if the parent resource pool has unreserved resources. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `(Optional) The CPU utilization of a resource pool will not exceed this limit, even if there are available resources. Set to ` + "`" + `-1` + "`" + ` for unlimited. Default: ` + "`" + `-1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "scale_descendants_shares",
					Description: `(Optional) Determines if the shares of all descendants of the resource pool are scaled up or down when the shares of the resource pool are scaled up or down. Can be one of ` + "`" + `disabled` + "`" + ` or ` + "`" + `scaleCpuAndMemoryShares` + "`" + `. Default: ` + "`" + `disabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is the [managed object ID][docs-about-morefs] of the resource pool. ## Importing An existing resource pool can be [imported][docs-import] into this resource via the path to the resource pool, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_resource_pool.resource_pool /dc-01/host/cluster-01/Resources/resource-pool-01 ` + "`" + `` + "`" + `` + "`" + ` The above would import the resource pool named ` + "`" + `resource-pool-01` + "`" + ` that is located in the compute cluster ` + "`" + `cluster-01` + "`" + ` in the ` + "`" + `dc-01` + "`" + ` datacenter. ### Settings that Require vSphere 7.0 or higher These settings require vSphere 7.0 or higher:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_storage_drs_vm_override",
			Category:         "Storage",
			ShortDescription: `Provides a VMware vSphere Storage DRS virtual machine override resource. This can be used to override Storage DRS settings in a datastore cluster.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"drs",
				"vm",
				"override",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "datastore_cluster_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the datastore cluster to put the override in. Forces a new resource if changed. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
				resource.Attribute{
					Name:        "virtual_machine_id",
					Description: `(Required) The UUID of the virtual machine to create the override for. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "sdrs_enabled",
					Description: `(Optional) Overrides the default Storage DRS setting for this virtual machine. When not specified, the datastore cluster setting is used.`,
				},
				resource.Attribute{
					Name:        "sdrs_automation_level",
					Description: `(Optional) Overrides any Storage DRS automation levels for this virtual machine. Can be one of ` + "`" + `automated` + "`" + ` or ` + "`" + `manual` + "`" + `. When not specified, the datastore cluster's settings are used according to the [specific SDRS subsystem][tf-vsphere-datastore-cluster-sdrs-levels]. [tf-vsphere-datastore-cluster-sdrs-levels]: /docs/providers/vsphere/r/datastore_cluster.html#storage-drs-automation-options`,
				},
				resource.Attribute{
					Name:        "sdrs_intra_vm_affinity",
					Description: `(Optional) Overrides the intra-VM affinity setting for this virtual machine. When ` + "`" + `true` + "`" + `, all disks for this virtual machine will be kept on the same datastore. When ` + "`" + `false` + "`" + `, Storage DRS may locate individual disks on different datastores if it helps satisfy cluster requirements. When not specified, the datastore cluster's settings are used. ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is a combination of the [managed object reference ID][docs-about-morefs] of the datastore cluster, and the UUID of the virtual machine. This is used to look up the override on subsequent plan and apply operations after the override has been created. ## Importing An existing override can be [imported][docs-import] into this resource by supplying both the path to the datastore cluster and the path to the virtual machine to ` + "`" + `terraform import` + "`" + `. If no override exists, an error will be given. An example is below: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_storage_drs_vm_override.drs_vm_override \ '{"datastore_cluster_path": "/dc1/datastore/ds-cluster", \ "virtual_machine_path": "/dc1/vm/srv1"}' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_tag",
			Category:         "Inventory",
			ShortDescription: `Provides a vSphere tag resource. This can be used to manage tags in vSphere.`,
			Description:      ``,
			Keywords: []string{
				"inventory",
				"tag",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The display name of the tag. The name must be unique within its category.`,
				},
				resource.Attribute{
					Name:        "category_id",
					Description: `(Required) The unique identifier of the parent category in which this tag will be created. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the tag. ## Attribute Reference The only attribute that is exported for this resource is the ` + "`" + `id` + "`" + `, which is the uniform resource name (URN) of this tag. ## Importing An existing tag can be [imported][docs-import] into this resource by supplying both the tag's category name and the name of the tag as a JSON string to ` + "`" + `terraform import` + "`" + `, as per the example below: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_tag.tag \ '{"category_name": "terraform-test-category", "tag_name": "terraform-test-tag"}' ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_tag_category",
			Category:         "Inventory",
			ShortDescription: `Provides a vSphere tag category resource. This can be used to manage tag categories in vSphere.`,
			Description:      ``,
			Keywords: []string{
				"inventory",
				"tag",
				"category",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the category.`,
				},
				resource.Attribute{
					Name:        "cardinality",
					Description: `(Required) The number of tags that can be assigned from this category to a single object at once. Can be one of ` + "`" + `SINGLE` + "`" + ` (object can only be assigned one tag in this category), to ` + "`" + `MULTIPLE` + "`" + ` (object can be assigned multiple tags in this category). Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "associable_types",
					Description: `(Required) A list object types that this category is valid to be assigned to. For a full list, click [here](#associable-object-types).`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A description for the category. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vapp_container",
			Category:         "Virtual Machine",
			ShortDescription: `Provides a VMware vSphere vApp container resource. This can be used to create and manage vApp container.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
				"vapp",
				"container",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the vApp container.`,
				},
				resource.Attribute{
					Name:        "parent_resource_pool_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the parent resource pool. This can be the root resource pool for a cluster or standalone host, or a resource pool itself. When moving a vApp container from one parent resource pool to another, both must share a common root resource pool or the move will fail.`,
				},
				resource.Attribute{
					Name:        "parent_folder_id",
					Description: `(Optional) The [managed object ID][docs-about-morefs] of the vApp container's parent folder.`,
				},
				resource.Attribute{
					Name:        "cpu_share_level",
					Description: `(Optional) The CPU allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. When ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `high` + "`" + ` are specified values in ` + "`" + `cpu_shares` + "`" + ` will be ignored. Default: ` + "`" + `normal` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_shares",
					Description: `(Optional) The number of shares allocated for CPU. Used to determine resource allocation in case of resource contention. If this is set, ` + "`" + `cpu_share_level` + "`" + ` must be ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_reservation",
					Description: `(Optional) Amount of CPU (MHz) that is guaranteed available to the vApp container. Default: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_expandable",
					Description: `(Optional) Determines if the reservation on a vApp container can grow beyond the specified value if the parent resource pool has unreserved resources. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `(Optional) The CPU utilization of a vApp container will not exceed this limit, even if there are available resources. Set to ` + "`" + `-1` + "`" + ` for unlimited. Default: ` + "`" + `-1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_share_level",
					Description: `(Optional) The CPU allocation level. The level is a simplified view of shares. Levels map to a pre-determined set of numeric values for shares. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. When ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `high` + "`" + ` are specified values in ` + "`" + `memory_shares` + "`" + ` will be ignored. Default: ` + "`" + `normal` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_shares",
					Description: `(Optional) The number of shares allocated for CPU. Used to determine resource allocation in case of resource contention. If this is set, ` + "`" + `memory_share_level` + "`" + ` must be ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `(Optional) Amount of CPU (MHz) that is guaranteed available to the vApp container. Default: ` + "`" + `0` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_expandable",
					Description: `(Optional) Determines if the reservation on a vApp container can grow beyond the specified value if the parent resource pool has unreserved resources. Default: ` + "`" + `true` + "`" + ``,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `(Optional) The CPU utilization of a vApp container will not exceed this limit, even if there are available resources. Set to ` + "`" + `-1` + "`" + ` for unlimited. Default: ` + "`" + `-1` + "`" + ``,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is the [managed object ID][docs-about-morefs] of the resource pool. ## Importing An existing vApp container can be [imported][docs-import] into this resource via the path to the vApp container, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html Example: ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vapp_container.vapp_container /dc-01/host/cluster-01/Resources/resource-pool-01/vapp-01 ` + "`" + `` + "`" + `` + "`" + ` The example above would import the vApp container named ` + "`" + `vapp-01` + "`" + ` that is located in the resource pool ` + "`" + `resource-pool-01` + "`" + ` that is part of the host cluster ` + "`" + `cluster-01` + "`" + ` in the ` + "`" + `dc-01` + "`" + ` datacenter.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vapp_entity",
			Category:         "Virtual Machine",
			ShortDescription: `Provides a vSphere vApp entity resource. This can be used to describe the behavior of an entity (virtual machine or sub-vApp container) in a vApp container.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
				"vapp",
				"entity",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "target_id",
					Description: `(Required) [Managed object ID|docs-about-morefs] of the entity to power on or power off. This can be a virtual machine or a vApp.`,
				},
				resource.Attribute{
					Name:        "container_id",
					Description: `(Required) [Managed object ID|docs-about-morefs] of the vApp container the entity is a member of.`,
				},
				resource.Attribute{
					Name:        "start_order",
					Description: `(Optional) Order to start and stop target in vApp. Default: 1`,
				},
				resource.Attribute{
					Name:        "start_action",
					Description: `(Optional) How to start the entity. Valid settings are none or powerOn. If set to none, then the entity does not participate in auto-start. Default: powerOn`,
				},
				resource.Attribute{
					Name:        "start_delay",
					Description: `(Optional) Delay in seconds before continuing with the next entity in the order of entities to be started. Default: 120`,
				},
				resource.Attribute{
					Name:        "stop_action",
					Description: `(Optional) Defines the stop action for the entity. Can be set to none, powerOff, guestShutdown, or suspend. If set to none, then the entity does not participate in auto-stop. Default: powerOff`,
				},
				resource.Attribute{
					Name:        "stop_delay",
					Description: `(Optional) Delay in seconds before continuing with the next entity in the order sequence. This is only used if the stopAction is guestShutdown. Default: 120`,
				},
				resource.Attribute{
					Name:        "wait_for_guest",
					Description: `(Optional) Determines if the VM should be marked as being started when VMware Tools are ready instead of waiting for ` + "`" + `start_delay` + "`" + `. This property has no effect for vApps. Default: false [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is the vApp entity's [managed object ID][docs-about-morefs] separated from the virtual machines [managed object ID][docs-about-morefs] by a colon. ## Importing An existing vApp entity can be [imported][docs-import] into this resource via the ID of the vApp Entity. [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vapp_entity.vapp_entity vm-123:res-456 ` + "`" + `` + "`" + `` + "`" + ` The above would import the vApp entity that governs the behavior of the virtual machine with a [managed object ID][docs-about-morefs] of vm-123 in the vApp container with the [managed object ID][docs-about-morefs] res-456.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_virtual_disk",
			Category:         "Virtual Machine",
			ShortDescription: `Provides a vSphere virtual disk resource. This can be used to create and delete virtual disks.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
				"disk",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "vmdk_path",
					Description: `(Required) The path, including filename, of the virtual disk to be created. This needs to end in ` + "`" + `.vmdk` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) The name of the datastore in which to create the disk.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) Size of the disk (in GB).`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The name of the datacenter in which to create the disk. Can be omitted when when ESXi or if there is only one datacenter in your infrastructure.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `(Optional) The type of disk to create. Can be one of ` + "`" + `eagerZeroedThick` + "`" + `, ` + "`" + `lazy` + "`" + `, or ` + "`" + `thin` + "`" + `. Default: ` + "`" + `eagerZeroedThick` + "`" + `. For information on what each kind of disk provisioning policy means, click [here][docs-vmware-vm-disk-provisioning]. [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html`,
				},
				resource.Attribute{
					Name:        "adapter_type",
					Description: `(Optional) The adapter type for this virtual disk. Can be one of ` + "`" + `ide` + "`" + `, ` + "`" + `lsiLogic` + "`" + `, or ` + "`" + `busLogic` + "`" + `. Default: ` + "`" + `lsiLogic` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "create_directories",
					Description: `(Optional) Tells the resource to create any directories that are a part of the ` + "`" + `vmdk_path` + "`" + ` parameter if they are missing. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_virtual_machine",
			Category:         "Virtual Machine",
			ShortDescription: `Provides a resource for VMware vSphere virtual machines. This resource can be used to create, modify, and delete virtual machines.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "alternate_guest_name",
					Description: `(Optional) The guest name for the operating system when ` + "`" + `guest_id` + "`" + ` is ` + "`" + `otherGuest` + "`" + ` or ` + "`" + `otherGuest64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) A user-provided description of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "cdrom",
					Description: `(Optional) A specification for a CD-ROM device on the virtual machine. See [CD-ROM options](#cd-rom-options) for more information.`,
				},
				resource.Attribute{
					Name:        "clone",
					Description: `(Optional) When specified, the virtual machine will be created as a clone of a specified template. Optional customization options can be submitted for the resource. See [creating a virtual machine from a template](#creating-a-virtual-machine-from-a-template) for more information.`,
				},
				resource.Attribute{
					Name:        "extra_config_reboot_required",
					Description: `(Optional) Allow the virtual machine to be rebooted when a change to ` + "`" + `extra_config` + "`" + ` occurs. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) Map of custom attribute ids to attribute value strings to set for virtual machine. Please refer to the [` + "`" + `vsphere_custom_attributes` + "`" + `][docs-setting-custom-attributes] resource for more information on setting custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datastore in which to place the virtual machine. The virtual machine configuration files is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on [virtual machine migration](#virtual-machine-migration) for more information on modifying this value.`,
				},
				resource.Attribute{
					Name:        "datastore_cluster_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datastore cluster in which to place the virtual machine. This setting applies to entire virtual machine and implies that you wish to use vSphere Storage DRS with the virtual machine. See the section on [virtual machine migration](#virtual-machine-migration) for more information on modifying this value. ~>`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Optional) The datacenter ID. Required only when deploying an OVF/OVA template.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Required) A specification for a virtual disk device on the virtual machine. See [disk options](#disk-options) for more information.`,
				},
				resource.Attribute{
					Name:        "extra_config",
					Description: `(Optional) Extra configuration data for the virtual machine. Can be used to supply advanced parameters not normally in configuration, such as instance metadata and userdata. ~>`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `(Optional) The firmware for the virtual machine. One of ` + "`" + `bios` + "`" + ` or ` + "`" + `efi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The path to the virtual machine folder in which to place the virtual machine, relative to the datacenter path (` + "`" + `/<datacenter-name>/vm` + "`" + `). For example, ` + "`" + `/dc-01/vm/foo` + "`" + ``,
				},
				resource.Attribute{
					Name:        "guest_id",
					Description: `(Optional) The guest ID for the operating system type. For a full list of possible values, see [here][vmware-docs-guest-ids]. Default: ` + "`" + `otherGuest64` + "`" + `. [vmware-docs-guest-ids]: https://vdc-download.vmware.com/vmwb-repository/dcr-public/b50dcbbf-051d-4204-a3e7-e1b618c1e384/538cf2ec-b34f-4bae-a332-3820ef9e7773/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html`,
				},
				resource.Attribute{
					Name:        "hardware_version",
					Description: `(Optional) The hardware version number. Valid range is from 4 to 19. The hardware version cannot be downgraded. See [virtual machine hardware compatibility][virtual-machine-hardware-compatibility] for more information. [virtual-machine-hardware-compatibility]: https://kb.vmware.com/s/article/2007240`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of a host on which to place the virtual machine. See the section on [virtual machine migration](#virtual-machine-migration) for more information on modifying this value. When using a vSphere cluster, if a ` + "`" + `host_system_id` + "`" + ` is not supplied, vSphere will select a host in the cluster to place the virtual machine, according to any defaults or vSphere DRS placement policies.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) A specification for a virtual NIC on the virtual machine. See [network interface options](#network-interface-options) for more information.`,
				},
				resource.Attribute{
					Name:        "pci_device_id",
					Description: `(Optional) List of host PCI device IDs in which to create PCI passthroughs. ~>`,
				},
				resource.Attribute{
					Name:        "ovf_deploy",
					Description: `(Optional) When specified, the virtual machine will be deployed from the provided OVF/OVA template. See [creating a virtual machine from an OVF/OVA template](#creating-a-virtual-machine-from-an-ovf-ova-template) for more information.`,
				},
				resource.Attribute{
					Name:        "replace_trigger",
					Description: `(Optional) Triggers replacement of resource whenever it changes. For example, ` + "`" + `replace_trigger = sha256(format("%s-%s",data.template_file.cloud_init_metadata.rendered,data.template_file.cloud_init_userdata.rendered))` + "`" + ` will fingerprint the changes in cloud-init metadata and userdata templates. This will enable a replacement of the resource whenever the dependant template renders a new configuration. (Forces a replacement.)`,
				},
				resource.Attribute{
					Name:        "resource_pool_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the resource pool in which to place the virtual machine. See the [Virtual Machine Migration](#virtual-machine-migration) section for more information on modifying this value. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "scsi_type",
					Description: `(Optional) The SCSI controller type for the virtual machine. One of ` + "`" + `lsilogic` + "`" + ` (LSI Logic Parallel), ` + "`" + `lsilogic-sas` + "`" + ` (LSI Logic SAS) or ` + "`" + `pvscsi` + "`" + ` (VMware Paravirtual). Default: ` + "`" + `pvscsi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scsi_bus_sharing",
					Description: `(Optional) The type of SCSI bus sharing for the virtual machine SCSI controller. One of ` + "`" + `physicalSharing` + "`" + `, ` + "`" + `virtualSharing` + "`" + `, and ` + "`" + `noSharing` + "`" + `. Default: ` + "`" + `noSharing` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_policy_id",
					Description: `(Optional) The ID of the storage policy to assign to the home directory of a virtual machine.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. Please refer to the [` + "`" + `vsphere_tag` + "`" + `][docs-applying-tags] resource for more information on applying tags to virtual machine resources. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "vapp",
					Description: `(Optional) Used for vApp configurations. The only sub-key available is ` + "`" + `properties` + "`" + `, which is a key/value map of properties for virtual machines imported from and OVF/OVA. See [Using vApp Properties for OVF/OVA Configuration](#using-vapp-properties-for-ovf-ova-configuration) for more information. ### CPU and Memory Options The following options control CPU and memory settings on a virtual machine:`,
				},
				resource.Attribute{
					Name:        "cpu_hot_add_enabled",
					Description: `(Optional) Allow CPUs to be added to the virtual machine while it is powered on.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_remove_enabled",
					Description: `(Optional) Allow CPUs to be removed to the virtual machine while it is powered on.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The memory size to assign to the virtual machine, in MB. Default: ` + "`" + `1024` + "`" + ` (1 GB).`,
				},
				resource.Attribute{
					Name:        "memory_hot_add_enabled",
					Description: `(Optional) Allow memory to be added to the virtual machine while it is powered on. ~>`,
				},
				resource.Attribute{
					Name:        "num_cores_per_socket",
					Description: `(Optional) The number of cores per socket in the virtual machine. The number of vCPUs on the virtual machine will be ` + "`" + `num_cpus` + "`" + ` divided by ` + "`" + `num_cores_per_socket` + "`" + `. If specified, the value supplied to ` + "`" + `num_cpus` + "`" + ` must be evenly divisible by this value. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_cpus",
					Description: `(Optional) The total number of virtual processor cores to assign to the virtual machine. Default: ` + "`" + `1` + "`" + `. ### Boot Options The following options control boot settings on a virtual machine:`,
				},
				resource.Attribute{
					Name:        "boot_delay",
					Description: `(Optional) The number of milliseconds to wait before starting the boot sequence. The default is no delay.`,
				},
				resource.Attribute{
					Name:        "boot_retry_delay",
					Description: `(Optional) The number of milliseconds to wait before retrying the boot sequence. This option is only valid if ` + "`" + `boot_retry_enabled` + "`" + ` is ` + "`" + `true` + "`" + `. Default: ` + "`" + `10000` + "`" + ` (10 seconds).`,
				},
				resource.Attribute{
					Name:        "boot_retry_enabled",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, a virtual machine that fails to boot will try again after the delay defined in ` + "`" + `boot_retry_delay` + "`" + `. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "efi_secure_boot_enabled",
					Description: `(Optional) Use this option to enable EFI secure boot when the ` + "`" + `firmware` + "`" + ` type is set to is ` + "`" + `efi` + "`" + `. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_after_power_on",
					Description: `(Optional) Enable post-power-on scripts to run when VMware Tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_after_resume",
					Description: `(Optional) Enable ost-resume scripts to run when VMware Tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_before_guest_reboot",
					Description: `(Optional) Enable pre-reboot scripts to run when VMware Tools is installed. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_before_guest_shutdown",
					Description: `(Optional) Enable pre-shutdown scripts to run when VMware Tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_before_guest_standby",
					Description: `(Optional) Enable pre-standby scripts to run when VMware Tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sync_time_with_host",
					Description: `(Optional) Enable the guest operating system to synchronization its clock with the host when the virtual machine is powered on or resumed. Requires vSphere 7.0 Update 1 and later. Requires VMware Tools to be installed. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sync_time_with_host_periodically",
					Description: `(Optional) Enable the guest operating system to periodically synchronize its clock with the host. Requires vSphere 7.0 Update 1 and later. On previous versions, setting ` + "`" + `sync_time_with_host` + "`" + ` is will enable periodic synchronization. Requires VMware Tools to be installed. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tools_upgrade_policy",
					Description: `(Optional) Enable automatic upgrade of the VMware Tools version when the virtual machine is rebooted. If necessary, VMware Tools is upgraded to the latest version supported by the host on which the virtual machine is running. Requires VMware Tools to be installed. One of ` + "`" + `manual` + "`" + ` or ` + "`" + `upgradeAtPowerCycle` + "`" + `. Default: ` + "`" + `manual` + "`" + `. ### Resource Allocation Options The following options control CPU and memory allocation on the virtual machine. Please note that the resource pool in which a virtual machine is placed may affect these options. The options are:`,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `(Optional) The maximum amount of CPU (in MHz) that the virtual machine can consume, regardless of available resources. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "cpu_reservation",
					Description: `(Optional) The amount of CPU (in MHz) that the virtual machine is guaranteed. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "cpu_share_level",
					Description: `(Optional) The allocation level for the virtual machine CPU resources. One of ` + "`" + `high` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_share_count",
					Description: `(Optional) The number of CPU shares allocated to the virtual machine when the ` + "`" + `cpu_share_level` + "`" + ` is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `(Optional) The maximum amount of memory (in MB) that th virtual machine can consume, regardless of available resources. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `(Optional) The amount of memory (in MB) that the virtual machine is guaranteed. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "memory_share_level",
					Description: `(Optional) The allocation level for the virtual machine memory resources. One of ` + "`" + `high` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_share_count",
					Description: `(Optional) The number of memory shares allocated to the virtual machine when the ` + "`" + `memory_share_level` + "`" + ` is ` + "`" + `custom` + "`" + `. ### Advanced Options The following options control advanced operation of the virtual machine, or control various parts of Terraform workflow, and should not need to be modified during basic operation of the resource. Only change these options if they are explicitly required, or if you are having trouble with Terraform's default behavior. The options are:`,
				},
				resource.Attribute{
					Name:        "cpu_performance_counters_enabled",
					Description: `(Optional) Enable CPU performance counters on the virtual machine. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_disk_uuid",
					Description: `(Optional) Expose the UUIDs of attached virtual disks to the virtual machine, allowing access to them in the guest. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `(Optional) Enable logging of virtual machine events to a log file stored in the virtual machine directory. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ept_rvi_mode",
					Description: `(Optional) The EPT/RVI (hardware memory virtualization) setting for the virtual machine. One of ` + "`" + `automatic` + "`" + `, ` + "`" + `on` + "`" + `, or ` + "`" + `off` + "`" + `. Default: ` + "`" + `automatic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "force_power_off",
					Description: `(Optional) If a guest shutdown failed or times out while updating or destroying (see [` + "`" + `shutdown_wait_timeout` + "`" + `](#shutdown_wait_timeout)), force the power-off of the virtual machine. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hv_mode",
					Description: `(Optional) The hardware virtualization (non-nested) setting for the virtual machine. One of ` + "`" + `hvAuto` + "`" + `, ` + "`" + `hvOn` + "`" + `, or ` + "`" + `hvOff` + "`" + `. Default: ` + "`" + `hvAuto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ide_controller_count",
					Description: `(Optional) The number of IDE controllers that the virtual machine. This directly affects the number of disks you can add to the virtual machine and the maximum disk unit number. Note that lowering this value does not remove controllers. Default: ` + "`" + `2` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ignored_guest_ips",
					Description: `(Optional) List of IP addresses and CIDR networks to ignore while waiting for an available IP address using either of the waiters. Any IP addresses in this list will be ignored so that the waiter will continue to wait for a valid IP address. Default: ` + "`" + `[]` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "latency_sensitivity",
					Description: `(Optional) Controls the scheduling delay of the virtual machine. Use a higher sensitivity for applications that require lower latency, such as VOIP, media player applications, or applications that require frequent access to mouse or keyboard devices. One of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `medium` + "`" + `, or ` + "`" + `high` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "migrate_wait_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for a virtual machine migration to complete before failing. Default: ` + "`" + `10` + "`" + ` minutes. See the section on [virtual machine migration](#virtual-machine-migration) for more information.`,
				},
				resource.Attribute{
					Name:        "nested_hv_enabled",
					Description: `(Optional) Enable nested hardware virtualization on the virtual machine, facilitating nested virtualization in the guest operating system. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "shutdown_wait_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for a graceful guest shutdown when making necessary updates to the virtual machine. If ` + "`" + `force_power_off` + "`" + ` is set to ` + "`" + `true` + "`" + `, the virtual machine will be forced to power-off after the timeout, otherwise an error is returned. Default: ` + "`" + `3` + "`" + ` minutes.`,
				},
				resource.Attribute{
					Name:        "swap_placement_policy",
					Description: `(Optional) The swap file placement policy for the virtual machine. One of ` + "`" + `inherit` + "`" + `, ` + "`" + `hostLocal` + "`" + `, or ` + "`" + `vmDirectory` + "`" + `. Default: ` + "`" + `inherit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vbs_enabled",
					Description: `(Optional) Enable Virtualization Based Security. Requires ` + "`" + `firmware` + "`" + ` to be ` + "`" + `efi` + "`" + `. In addition, ` + "`" + `vvtd_enabled` + "`" + `, ` + "`" + `nested_hv_enabled` + "`" + `, and ` + "`" + `efi_secure_boot_enabled` + "`" + ` must all have a value of ` + "`" + `true` + "`" + `. Supported on vSphere 6.7 and later. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "vvtd_enabled",
					Description: `(Optional) Enable Intel Virtualization Technology for Directed I/O for the virtual machine (_I/O MMU_ in the vSphere Client). Supported on vSphere 6.7 and later. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait_for_guest_ip_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for an available guest IP address on the virtual machine. This should only be used if the version VMware Tools does not allow the [` + "`" + `wait_for_guest_net_timeout` + "`" + `](#wait_for_guest_net_timeout) waiter to be used. A value less than ` + "`" + `1` + "`" + ` disables the waiter. Default: ` + "`" + `0` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait_for_guest_net_routable",
					Description: `(Optional) Controls whether or not the guest network waiter waits for a routable address. When ` + "`" + `false` + "`" + `, the waiter does not wait for a default gateway, nor are IP addresses checked against any discovered default gateways as part of its success criteria. This property is ignored if the [` + "`" + `wait_for_guest_ip_timeout` + "`" + `](#wait_for_guest_ip_timeout) waiter is used. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait_for_guest_net_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for an available guest IP address on the virtual machine. Older versions of VMware Tools do not populate this property. In those cases, this waiter can be disabled and the [` + "`" + `wait_for_guest_ip_timeout` + "`" + `](#wait_for_guest_ip_timeout) waiter can be used instead. A value less than ` + "`" + `1` + "`" + ` disables the waiter. Default: ` + "`" + `5` + "`" + ` minutes. ### Disk Options Virtual disks are managed by adding one or more instance of the ` + "`" + `disk` + "`" + ` block. At a minimum, both the ` + "`" + `label` + "`" + ` and ` + "`" + `size` + "`" + ` attributes must be provided. The ` + "`" + `unit_number` + "`" + ` is required for any disk other than the first, and there must be at least one resource with the implicit number of ` + "`" + `0` + "`" + `. The following example demonstrates and abridged multi-disk configuration:`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) A label for the virtual disk. Forces a new disk, if changed. ~>`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the disk, in GB. Must be a whole number.`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(Optional) The disk number on the storage bus. The maximum value for this setting is the value of the controller count times the controller capacity (15 for SCSI, 30 for SATA, and 2 for IDE). Duplicate unit numbers are not allowed. Default ` + "`" + `0` + "`" + `, for which one disk must be set to.`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] for the datastore on which the virtual disk is placed. The default is to use the datastore of the virtual machine. See the section on [virtual machine migration](#virtual-machine-migration) for information on modifying this value. ~>`,
				},
				resource.Attribute{
					Name:        "attach",
					Description: `(Optional) Attach an external disk instead of creating a new one. Implies and conflicts with ` + "`" + `keep_on_remove` + "`" + `. If set, you cannot set ` + "`" + `size` + "`" + `, ` + "`" + `eagerly_scrub` + "`" + `, or ` + "`" + `thin_provisioned` + "`" + `. Must set ` + "`" + `path` + "`" + ` if used. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) When using ` + "`" + `attach` + "`" + `, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk filename.`,
				},
				resource.Attribute{
					Name:        "keep_on_remove",
					Description: `(Optional) Keep this disk when removing the device or destroying the virtual machine. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_mode",
					Description: `(Optional) The mode of this this virtual disk for purposes of writes and snapshots. One of ` + "`" + `append` + "`" + `, ` + "`" + `independent_nonpersistent` + "`" + `, ` + "`" + `independent_persistent` + "`" + `, ` + "`" + `nonpersistent` + "`" + `, ` + "`" + `persistent` + "`" + `, or ` + "`" + `undoable` + "`" + `. Default: ` + "`" + `persistent` + "`" + `. For more information on these option, please refer to the [product documentation][vmware-docs-disk-mode]. [vmware-docs-disk-mode]: https://vdc-download.vmware.com/vmwb-repository/dcr-public/da47f910-60ac-438b-8b9b-6122f4d14524/16b7274a-bf8b-4b4c-a05e-746f2aa93c8c/doc/vim.vm.device.VirtualDiskOption.DiskMode.html`,
				},
				resource.Attribute{
					Name:        "eagerly_scrub",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the disk space is zeroed out when the virtual machine is created. This will delay the creation of the virtual disk. Cannot be set to ` + "`" + `true` + "`" + ` when ` + "`" + `thin_provisioned` + "`" + ` is ` + "`" + `true` + "`" + `. See the section on [picking a disk type](#picking-a-disk-type) for more information. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to ` + "`" + `true` + "`" + ` when ` + "`" + `eagerly_scrub` + "`" + ` is ` + "`" + `true` + "`" + `. See the section on [selecting a disk type](#selecting-a-disk-type) for more information. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_sharing",
					Description: `(Optional) The sharing mode of this virtual disk. One of ` + "`" + `sharingMultiWriter` + "`" + ` or ` + "`" + `sharingNone` + "`" + `. Default: ` + "`" + `sharingNone` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "write_through",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, writes for this disk are sent directly to the filesystem immediately instead of being buffered. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "io_limit",
					Description: `(Optional) The upper limit of IOPS that this disk can use. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "io_reservation",
					Description: `(Optional) The I/O reservation (guarantee) for the virtual disk has, in IOPS. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "io_share_level",
					Description: `(Optional) The share allocation level for the virtual disk. One of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "io_share_count",
					Description: `(Optional) The share count for the virtual disk when the share level is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "storage_policy_id",
					Description: `(Optional) The UUID of the storage policy to assign to the virtual disk.`,
				},
				resource.Attribute{
					Name:        "controller_type",
					Description: `(Optional) The type of storage controller to attach the disk to. Can be ` + "`" + `scsi` + "`" + `, ` + "`" + `sata` + "`" + `, or ` + "`" + `ide` + "`" + `. You must have the appropriate number of controllers enabled for the selected type. Default ` + "`" + `scsi` + "`" + `. #### Computed Disk Attributes`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the virtual disk VMDK file. This is used to track the virtual disk on the virtual machine. #### Virtual Disk Provisioning Policies The ` + "`" + `eagerly_scrub` + "`" + ` and ` + "`" + `thin_provisioned` + "`" + ` options control the virtual disk space allocation type. These appear in vSphere as a unified enumeration of options, the equivalents of which are explained below. The provider configuration defaults to thin provision. The options are:`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the network on which to connect the virtual machine network interface.`,
				},
				resource.Attribute{
					Name:        "adapter_type",
					Description: `(Optional) The network interface type. One of ` + "`" + `e1000` + "`" + `, ` + "`" + `e1000e` + "`" + `, or ` + "`" + `vmxnet3` + "`" + `. Default: ` + "`" + `vmxnet3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_static_mac",
					Description: `(Optional) If true, the ` + "`" + `mac_address` + "`" + ` field is treated as a static MAC address and set accordingly. Setting this to ` + "`" + `true` + "`" + ` requires ` + "`" + `mac_address` + "`" + ` to be set. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the network interface. Can only be manually set if ` + "`" + `use_static_mac` + "`" + ` is ` + "`" + `true` + "`" + `. Otherwise, the value is computed and presents the assigned MAC address for the interface.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `(Optional) The upper bandwidth limit of the network interface, in Mbits/sec. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "bandwidth_reservation",
					Description: `(Optional) The bandwidth reservation of the network interface, in Mbits/sec. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "bandwidth_share_level",
					Description: `(Optional) The bandwidth share allocation level for the network interface. One of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth_share_count",
					Description: `(Optional) The share count for the network interface when the share level is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ovf_mapping",
					Description: `(Optional) Specifies which NIC in an OVF/OVA the ` + "`" + `network_interface` + "`" + ` should be associated. Only applies at creation when deploying from an OVF/OVA. ### CD-ROM Options A CD-ROM device is managed by adding an instance of the ` + "`" + `cdrom` + "`" + ` block. Up to two virtual CD-ROM devices can be created and attached to the virtual machine. If adding multiple CD-ROM devices, add each device as a separate ` + "`" + `cdrom` + "`" + ` block. The resource supports attaching a CD-ROM from either a datastore ISO or using a remote client device.`,
				},
				resource.Attribute{
					Name:        "client_device",
					Description: `(Optional) Indicates whether the device should be backed by remote client device. Conflicts with ` + "`" + `datastore_id` + "`" + ` and ` + "`" + `path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) The datastore ID that on which the ISO is located. Required for using a datastore ISO. Conflicts with ` + "`" + `client_device` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to the ISO file. Required for using a datastore ISO. Conflicts with ` + "`" + `client_device` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The ID of the device within the virtual machine.`,
				},
				resource.Attribute{
					Name:        "device_address",
					Description: `An address internal to Terraform that helps locate the device when ` + "`" + `key` + "`" + ` is unavailable. This follows a convention of ` + "`" + `CONTROLLER_TYPE:BUS_NUMBER:UNIT_NUMBER` + "`" + `. Example: ` + "`" + `scsi:0:1` + "`" + ` means device unit ` + "`" + `1` + "`" + ` on SCSI bus ` + "`" + `0` + "`" + `. ## Creating a Virtual Machine from a Template The ` + "`" + `clone` + "`" + ` block can be used to create a new virtual machine from an existing virtual machine or template. The resource supports both making a complete copy of a virtual machine, or cloning from a snapshot (also known as a linked clone). See the section on [cloning and customization](#cloning-and-customization) for more information. ~>`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `(Required) The UUID of the source virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "linked_clone",
					Description: `(Optional) Clone the virtual machine from a snapshot or a template. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout, in minutes, to wait for the cloning process to complete. Default: 30 minutes.`,
				},
				resource.Attribute{
					Name:        "customize",
					Description: `(Optional) The customization spec for this clone. This allows the user to configure the virtual machine post-clone. For more details, see [virtual machine customizations](#virtual-machine-customizations). ### Virtual Machine Customizations As part of the ` + "`" + `clone` + "`" + ` operation, a virtual machine can be [customized][vmware-docs-customize] to configure host, network, or licensing settings. [vmware-docs-customize]: https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-58E346FF-83AE-42B8-BE58-253641D257BC.html To perform virtual machine customization as a part of the clone process, specify the ` + "`" + `customize` + "`" + ` block with the respective customization options, nested within the ` + "`" + `clone` + "`" + ` block. Windows guests are customized using Sysprep, which will result in the machine SID being reset. Before using customization, check is that your source virtual machine meets the [requirements](https://docs.vmware.com/en/VMware-vSphere/7.0/com.vmware.vsphere.vm_admin.doc/GUID-58E346FF-83AE-42B8-BE58-253641D257BC.html) for guest OS customization on vSphere. See the section on [cloning and customization](#cloning-and-customization) for a usage synopsis. The settings for ` + "`" + `customize` + "`" + ` are as follows: #### Customization Timeout Settings`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The time, in minutes, that the provider waits for customization to complete before failing. The default is ` + "`" + `10` + "`" + ` minutes. Setting the value to ` + "`" + `0` + "`" + ` or a negative value disables the waiter. #### Network Interface Settings These settings, which should be specified in nested ` + "`" + `network_interface` + "`" + ` blocks within [` + "`" + `customize` + "`" + `](#virtual-machine-customization) block, configure network interfaces on a per-interface basis and are matched up to [` + "`" + `network_interface` + "`" + `](#network-interface-options) devices in the order declared. Static IP Address Example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { # ... other configuration ... network_interface { network_id = data.vsphere_network.public.id } network_interface { network_id = data.vsphere_network.private.id } clone { # ... other configuration ... customize { # ... other configuration ... network_interface { ipv4_address = "10.0.0.10" ipv4_netmask = 24 } network_interface { ipv4_address = "172.16.0.10" ipv4_netmask = 24 } ipv4_gateway = "10.0.0.1" } } } ` + "`" + `` + "`" + `` + "`" + ` The first ` + "`" + `network_interface` + "`" + ` would be assigned to the ` + "`" + `public` + "`" + ` interface, and the second to the ` + "`" + `private` + "`" + ` interface. To use DHCP, declare an empty ` + "`" + `network_interface` + "`" + ` block for each interface.`,
				},
				resource.Attribute{
					Name:        "dns_server_list",
					Description: `(Optional) DNS servers for the network interface. Used by Windows guest operating systems, but ignored by Linux distribution guest operating systems. For Linux, please refer to the section on the [global DNS settings](#global-dns-settings).`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `(Optional) DNS search domain for the network interface. Used by Windows guest operating systems, but ignored by Linux distribution guest operating systems. For Linux, please refer to the section on the [global DNS settings](#global-dns-settings).`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) The IPv4 address assigned to the network adapter. If blank or not included, DHCP is used.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) The IPv6 address assigned to the network adapter. If blank or not included, auto-configuration is used.`,
				},
				resource.Attribute{
					Name:        "ipv6_netmask",
					Description: `(Optional) The IPv6 subnet mask, in bits (_e.g._ ` + "`" + `32` + "`" + `). ~>`,
				},
				resource.Attribute{
					Name:        "ipv4_gateway",
					Description: `(Optional) The IPv4 default gateway when using ` + "`" + `network_interface` + "`" + ` customization on the virtual machine.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `(Optional) The IPv6 default gateway when using ` + "`" + `network_interface` + "`" + ` customization on the virtual machine. #### Global DNS Settings The following settings configure DNS globally, generally for Linux distribution guest operating systems. For Windows guest operating systems, this is performer per-interface. See the section on [network interface settings](#network-interface-settings) for more information.`,
				},
				resource.Attribute{
					Name:        "dns_server_list",
					Description: `The list of DNS servers to configure on the virtual machine.`,
				},
				resource.Attribute{
					Name:        "dns_suffix_list",
					Description: `A list of DNS search domains to add to the DNS configuration on the virtual machine. #### Linux Customization Options The settings in the ` + "`" + `linux_options` + "`" + ` block pertain to Linux distribution guest operating system customization. If you are customizing a Linux guest operating system, this section must be included.`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Required) The host name for this machine. This, along with ` + "`" + `domain` + "`" + `, make up the FQDN of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name for this machine. This, along with ` + "`" + `host_name` + "`" + `, make up the FQDN of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "hw_clock_utc",
					Description: `(Optional) Tells the operating system that the hardware clock is set to UTC. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "script_text",
					Description: `(Optional) The customization script for the virtual machine that will be applied before and / or after guest customization. For more information on enabling and using a customization script, please refer to [VMware KB 74880][vmware-kb-74880]. The [Heredoc style][tf-heredoc-strings] of string literal is recommended. [vmware-kb-74880]: https://kb.vmware.com/s/article/74880 [tf-heredoc-strings]: https://www.terraform.io/language/expressions/strings#heredoc-strings`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) Sets the time zone. For a list of possible combinations, please refer to [VMware KB 2145518][vmware-kb-2145518]. The default is UTC. [vmware-kb-2145518]: https://kb.vmware.com/s/article/2145518 #### Windows Customization Options The settings in the ` + "`" + `windows_options` + "`" + ` block pertain to Windows guest OS customization. If you are customizing a Windows operating system, this section must be included.`,
				},
				resource.Attribute{
					Name:        "computer_name",
					Description: `(Required) The computer name of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Optional) The administrator password for the virtual machine. ~>`,
				},
				resource.Attribute{
					Name:        "workgroup",
					Description: `(Optional) The workgroup name for the virtual machine. One of this or ` + "`" + `join_domain` + "`" + ` must be included.`,
				},
				resource.Attribute{
					Name:        "join_domain",
					Description: `(Optional) The domain name in which to join the virtual machine. One of this or ` + "`" + `workgroup` + "`" + ` must be included.`,
				},
				resource.Attribute{
					Name:        "domain_admin_user",
					Description: `(Optional) The user account with administrative privileges to use to join the guest operating system to the domain. Required if setting ` + "`" + `join_domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_admin_password",
					Description: `(Optional) The password user account with administrative privileges used to join the virtual machine to the domain. Required if setting ` + "`" + `join_domain` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Optional) The full name of the organization owner of the virtual machine. This populates the "user" field in the general Windows system information. Default: ` + "`" + `Administrator` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `(Optional) The name of the organization for the virtual machine. This option populates the "organization" field in the general Windows system information. Default: ` + "`" + `Managed by Terraform` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "product_key",
					Description: `(Optional) The product key for the virtual machine Windows guest operating system. The default is no key.`,
				},
				resource.Attribute{
					Name:        "run_once_command_list",
					Description: `(Optional) A list of commands to run at first user logon, after guest customization. Each run once command is limited by the API to 260 characters.`,
				},
				resource.Attribute{
					Name:        "auto_logon",
					Description: `(Optional) Specifies whether or not the virtual machine automatically logs on as Administrator. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_logon_count",
					Description: `(Optional) Specifies how many times the virtual machine should auto-logon the Administrator account when ` + "`" + `auto_logon` + "`" + ` is ` + "`" + `true` + "`" + `. This option should be set accordingly to ensure that all of your commands that run in ` + "`" + `run_once_command_list` + "`" + ` can log in to run. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) The time zone for the virtual machine. For a list of supported codes, please refer to the [MIcrosoft documentation][ms-docs-valid-sysprep-tzs]. The default is ` + "`" + `85` + "`" + ` (GMT/UTC). [ms-docs-valid-sysprep-tzs]: https://msdn.microsoft.com/en-us/library/ms912391(v=winembedded.11).aspx ##### Using a Windows SysPrep Configuration An alternative to the ` + "`" + `windows_options` + "`" + ` demonstrated above, you can provide a SysPrep ` + "`" + `.XML` + "`" + ` file using the ` + "`" + `windows_sysprep_text` + "`" + ` option. This option allows full control of the Windows guest operating system customization process.`,
				},
				resource.Attribute{
					Name:        "allow_unverified_ssl_cert",
					Description: `(Optional) Allow unverified SSL certificates while deploying OVF/OVA from a URL. Defaults ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_hidden_properties",
					Description: `(Optional) Allow properties with ` + "`" + `ovf:userConfigurable=false` + "`" + ` to be set. Defaults ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "local_ovf_path",
					Description: `(Optional) The absolute path to the OVF/OVA file on the local system. When deploying from an OVF, ensure the necessary files, such as ` + "`" + `.vmdk` + "`" + ` and ` + "`" + `.mf` + "`" + ` files are also in the same directory as the ` + "`" + `.ovf` + "`" + ` file.`,
				},
				resource.Attribute{
					Name:        "remote_ovf_url",
					Description: `(Optional) URL to the OVF/OVA file. ~>`,
				},
				resource.Attribute{
					Name:        "ip_allocation_policy",
					Description: `(Optional) The IP allocation policy.`,
				},
				resource.Attribute{
					Name:        "ip_protocol",
					Description: `(Optional) The IP protocol.`,
				},
				resource.Attribute{
					Name:        "disk_provisioning",
					Description: `(Optional) The disk provisioning policy. If set, all the disks included in the OVF/OVA will have the same specified policy. One of ` + "`" + `thin` + "`" + `, ` + "`" + `flat` + "`" + `, ` + "`" + `thick` + "`" + `, or ` + "`" + `sameAsSource` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "deployment_option",
					Description: `(Optional) The key for the deployment option. If empty, the default option is selected.`,
				},
				resource.Attribute{
					Name:        "ovf_network_map",
					Description: `(Optional) The mapping of network identifiers from the OVF descriptor to a network UUID. ### Using vApp Properties for OVF/OVA Configuration You can use the ` + "`" + `properties` + "`" + ` section of the ` + "`" + `vapp` + "`" + ` block to supply configuration parameters to a virtual machine cloned from a template that originated from an imported OVF/OVA file. Both GuestInfo and ISO transport methods are supported. For templates that use ISO transport, a CD-ROM backed by a client device must be included.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `When reducing the memory size, or when increasing the memory size and ` + "`" + `memory_hot_add_enabled` + "`" + ` is set to ` + "`" + `false` + "`" + ``,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `When deleting a network interface and VMware Tools is not running.`,
				},
				resource.Attribute{
					Name:        "network_interface.adapter_type",
					Description: `When VMware Tools is not running.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "reboot_required",
					Description: `Value internal to Terraform used to determine if a configuration set change requires a reboot. This value is most useful during an update process and gets reset on refresh.`,
				},
				resource.Attribute{
					Name:        "vmware_tools_status",
					Description: `The state of VMware Tools in the guest. This will determine the proper course of action for some device operations.`,
				},
				resource.Attribute{
					Name:        "vmx_path",
					Description: `The path of the virtual machine configuration file on the datastore in which the virtual machine is placed.`,
				},
				resource.Attribute{
					Name:        "imported",
					Description: `Indicates if the virtual machine resource has been imported, or if the state has been migrated from a previous version of the resource. It influences the behavior of the first post-import apply operation. See the section on [importing](#importing) below.`,
				},
				resource.Attribute{
					Name:        "change_version",
					Description: `A unique identifier for a given version of the last configuration was applied.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the virtual machine. Also exposed as the ` + "`" + `id` + "`" + ` of the resource.`,
				},
				resource.Attribute{
					Name:        "default_ip_address",
					Description: `The IP address selected by Terraform to be used with any [provisioners][tf-docs-provisioners] configured on this resource. When possible, this is the first IPv4 address that is reachable through the default gateway configured on the machine, then the first reachable IPv6 address, and then the first general discovered address if neither exists. If VMware Tools is not running on the virtual machine, or if the virtual machine is powered off, this value will be blank.`,
				},
				resource.Attribute{
					Name:        "guest_ip_addresses",
					Description: `The current list of IP addresses on this machine, including the value of ` + "`" + `default_ip_address` + "`" + `. If VMware Tools is not running on the virtual machine, or if the virtul machine is powered off, this list will be empty.`,
				},
				resource.Attribute{
					Name:        "vapp_transport",
					Description: `Computed value which is only valid for cloned virtual machines. A list of vApp transport methods supported by the source virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `A computed value for the current power state of the virtual machine. One of ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `, or ` + "`" + `suspended` + "`" + `. [docs-about-morefs]: https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs#use-of-managed-object-references-by-the-vsphere-provider ## Importing An existing virtual machine can be [imported][docs-import] into the Terraform state by providing the full path to the virtual machine. [docs-import]: /docs/import/index.html`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "reboot_required",
					Description: `Value internal to Terraform used to determine if a configuration set change requires a reboot. This value is most useful during an update process and gets reset on refresh.`,
				},
				resource.Attribute{
					Name:        "vmware_tools_status",
					Description: `The state of VMware Tools in the guest. This will determine the proper course of action for some device operations.`,
				},
				resource.Attribute{
					Name:        "vmx_path",
					Description: `The path of the virtual machine configuration file on the datastore in which the virtual machine is placed.`,
				},
				resource.Attribute{
					Name:        "imported",
					Description: `Indicates if the virtual machine resource has been imported, or if the state has been migrated from a previous version of the resource. It influences the behavior of the first post-import apply operation. See the section on [importing](#importing) below.`,
				},
				resource.Attribute{
					Name:        "change_version",
					Description: `A unique identifier for a given version of the last configuration was applied.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the virtual machine. Also exposed as the ` + "`" + `id` + "`" + ` of the resource.`,
				},
				resource.Attribute{
					Name:        "default_ip_address",
					Description: `The IP address selected by Terraform to be used with any [provisioners][tf-docs-provisioners] configured on this resource. When possible, this is the first IPv4 address that is reachable through the default gateway configured on the machine, then the first reachable IPv6 address, and then the first general discovered address if neither exists. If VMware Tools is not running on the virtual machine, or if the virtual machine is powered off, this value will be blank.`,
				},
				resource.Attribute{
					Name:        "guest_ip_addresses",
					Description: `The current list of IP addresses on this machine, including the value of ` + "`" + `default_ip_address` + "`" + `. If VMware Tools is not running on the virtual machine, or if the virtul machine is powered off, this list will be empty.`,
				},
				resource.Attribute{
					Name:        "vapp_transport",
					Description: `Computed value which is only valid for cloned virtual machines. A list of vApp transport methods supported by the source virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "power_state",
					Description: `A computed value for the current power state of the virtual machine. One of ` + "`" + `on` + "`" + `, ` + "`" + `off` + "`" + `, or ` + "`" + `suspended` + "`" + `. [docs-about-morefs]: https://registry.terraform.io/providers/hashicorp/vsphere/latest/docs#use-of-managed-object-references-by-the-vsphere-provider ## Importing An existing virtual machine can be [imported][docs-import] into the Terraform state by providing the full path to the virtual machine. [docs-import]: /docs/import/index.html`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_virtual_machine_snapshot",
			Category:         "Virtual Machine",
			ShortDescription: `Provides a VMware vSphere virtual machine snapshot resource. This can be used to create and delete virtual machine snapshots.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
				"snapshot",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "virtual_machine_uuid",
					Description: `(Required) The virtual machine UUID.`,
				},
				resource.Attribute{
					Name:        "snapshot_name",
					Description: `(Required) The name of the snapshot.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Required) A description for the snapshot.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Required) If set to ` + "`" + `true` + "`" + `, a dump of the internal state of the virtual machine is included in the snapshot.`,
				},
				resource.Attribute{
					Name:        "quiesce",
					Description: `(Required) If set to ` + "`" + `true` + "`" + `, and the virtual machine is powered on when the snapshot is taken, VMware Tools is used to quiesce the file system in the virtual machine.`,
				},
				resource.Attribute{
					Name:        "remove_children",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the entire snapshot subtree is removed when this resource is destroyed.`,
				},
				resource.Attribute{
					Name:        "consolidate",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the delta disks involved in this snapshot will be consolidated into the parent when this resource is destroyed. ## Attribute Reference The only attribute this resource exports is the resource ` + "`" + `id` + "`" + `, which is set to the [managed object reference ID][docs-about-morefs] of the snapshot. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vm_storage_policy",
			Category:         "Storage",
			ShortDescription: `Storage policies can select the most appropriate datastore for the virtual machine and enforce the required level of service.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"vm",
				"policy",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the storage policy.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) Description of the storage policy.`,
				},
				resource.Attribute{
					Name:        "tag_rules",
					Description: `(Required) List of tag rules. The tag category and tags to be associated to this storage policy.`,
				},
				resource.Attribute{
					Name:        "tag_category",
					Description: `(Required) Name of the tag category.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Required) List of Name of tags to select from the given category.`,
				},
				resource.Attribute{
					Name:        "include_datastores_with_tags",
					Description: `(Optional) Include datastores with the given tags or exclude. Default ` + "`" + `true` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vmfs_datastore",
			Category:         "Storage",
			ShortDescription: `Provides a vSphere VMFS datastore resource. This can be used to configure a VMFS datastore on a host or set of hosts.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"vmfs",
				"datastore",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the datastore. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Required) The [managed object ID][docs-about-morefs] of the host to set the datastore up on. Note that this is not necessarily the only host that the datastore will be set up on - see [here](#auto-mounting-of-datastores-within-vcenter) for more info. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "disks",
					Description: `(Required) The disks to use with the datastore.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The relative path to a folder to put this datastore in. This is a path relative to the datacenter you are deploying the datastore to. Example: for the ` + "`" + `dc1` + "`" + ` datacenter, and a provided ` + "`" + `folder` + "`" + ` of ` + "`" + `foo/bar` + "`" + `, Terraform will place a datastore named ` + "`" + `terraform-test` + "`" + ` in a datastore folder located at ` + "`" + `/dc1/datastore/foo/bar` + "`" + `, with the final inventory path being ` + "`" + `/dc1/datastore/foo/bar/terraform-test` + "`" + `. Conflicts with ` + "`" + `datastore_cluster_id` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datastore_cluster_id",
					Description: `(Optional) The [managed object ID][docs-about-morefs] of a datastore cluster to put this datastore in. Conflicts with ` + "`" + `folder` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The [managed object reference ID][docs-about-morefs] of the datastore.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `The connectivity status of the datastore. If this is ` + "`" + `false` + "`" + `, some other computed attributes may be out of date.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Maximum capacity of the datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "free_space",
					Description: `Available space of this datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "maintenance_mode",
					Description: `The current maintenance mode state of the datastore.`,
				},
				resource.Attribute{
					Name:        "multiple_host_access",
					Description: `If ` + "`" + `true` + "`" + `, more than one host in the datacenter has been configured with access to the datastore.`,
				},
				resource.Attribute{
					Name:        "uncommitted_space",
					Description: `Total additional storage space, in megabytes, potentially used by all virtual machines on this datastore.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The unique locator for the datastore. ## Importing An existing VMFS datastore can be [imported][docs-import] into this resource via its managed object ID, via the command below. You also need the host system ID. [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vmfs_datastore.datastore datastore-123:host-10 ` + "`" + `` + "`" + `` + "`" + ` You need a tool like [` + "`" + `govc` + "`" + `][ext-govc] that can display managed object IDs. [ext-govc]: https://github.com/vmware/govmomi/tree/master/govc In the case of govc, you can locate a managed object ID from an inventory path by doing the following: ` + "`" + `` + "`" + `` + "`" + ` $ govc ls -i /dc/datastore/terraform-test Datastore:datastore-123 ` + "`" + `` + "`" + `` + "`" + ` To locate host IDs, it might be a good idea to supply the ` + "`" + `-l` + "`" + ` flag as well so that you can line up the names with the IDs: ` + "`" + `` + "`" + `` + "`" + ` $ govc ls -l -i /dc/host/cluster1 ResourcePool:resgroup-10 /dc/host/cluster1/Resources HostSystem:host-10 /dc/host/cluster1/esxi1 HostSystem:host-11 /dc/host/cluster1/esxi2 HostSystem:host-12 /dc/host/cluster1/esxi3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The [managed object reference ID][docs-about-morefs] of the datastore.`,
				},
				resource.Attribute{
					Name:        "accessible",
					Description: `The connectivity status of the datastore. If this is ` + "`" + `false` + "`" + `, some other computed attributes may be out of date.`,
				},
				resource.Attribute{
					Name:        "capacity",
					Description: `Maximum capacity of the datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "free_space",
					Description: `Available space of this datastore, in megabytes.`,
				},
				resource.Attribute{
					Name:        "maintenance_mode",
					Description: `The current maintenance mode state of the datastore.`,
				},
				resource.Attribute{
					Name:        "multiple_host_access",
					Description: `If ` + "`" + `true` + "`" + `, more than one host in the datacenter has been configured with access to the datastore.`,
				},
				resource.Attribute{
					Name:        "uncommitted_space",
					Description: `Total additional storage space, in megabytes, potentially used by all virtual machines on this datastore.`,
				},
				resource.Attribute{
					Name:        "url",
					Description: `The unique locator for the datastore. ## Importing An existing VMFS datastore can be [imported][docs-import] into this resource via its managed object ID, via the command below. You also need the host system ID. [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vmfs_datastore.datastore datastore-123:host-10 ` + "`" + `` + "`" + `` + "`" + ` You need a tool like [` + "`" + `govc` + "`" + `][ext-govc] that can display managed object IDs. [ext-govc]: https://github.com/vmware/govmomi/tree/master/govc In the case of govc, you can locate a managed object ID from an inventory path by doing the following: ` + "`" + `` + "`" + `` + "`" + ` $ govc ls -i /dc/datastore/terraform-test Datastore:datastore-123 ` + "`" + `` + "`" + `` + "`" + ` To locate host IDs, it might be a good idea to supply the ` + "`" + `-l` + "`" + ` flag as well so that you can line up the names with the IDs: ` + "`" + `` + "`" + `` + "`" + ` $ govc ls -l -i /dc/host/cluster1 ResourcePool:resgroup-10 /dc/host/cluster1/Resources HostSystem:host-10 /dc/host/cluster1/esxi1 HostSystem:host-11 /dc/host/cluster1/esxi2 HostSystem:host-12 /dc/host/cluster1/esxi3 ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vnic",
			Category:         "Host and Cluster Management",
			ShortDescription: `Provides a VMware vSphere vnic resource..`,
			Description:      ``,
			Keywords: []string{
				"host",
				"and",
				"cluster",
				"management",
				"vnic",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "portgroup",
					Description: `(Optional) Portgroup to attach the nic to. Do not set if you set distributed_switch_port.`,
				},
				resource.Attribute{
					Name:        "distributed_switch_port",
					Description: `(Optional) UUID of the DVSwitch the nic will be attached to. Do not set if you set portgroup.`,
				},
				resource.Attribute{
					Name:        "distributed_port_group",
					Description: `(Optional) Key of the distributed portgroup the nic will connect to.`,
				},
				resource.Attribute{
					Name:        "ipv4",
					Description: `(Optional) IPv4 settings. Either this or ` + "`" + `ipv6` + "`" + ` needs to be set. See [ipv4 options](#ipv4-options) below.`,
				},
				resource.Attribute{
					Name:        "ipv6",
					Description: `(Optional) IPv6 settings. Either this or ` + "`" + `ipv6` + "`" + ` needs to be set. See [ipv6 options](#ipv6-options) below.`,
				},
				resource.Attribute{
					Name:        "mac",
					Description: `(Optional) MAC address of the interface.`,
				},
				resource.Attribute{
					Name:        "mtu",
					Description: `(Optional) MTU of the interface.`,
				},
				resource.Attribute{
					Name:        "netstack",
					Description: `(Optional) TCP/IP stack setting for this interface. Possible values are 'defaultTcpipStack', 'vmotion', 'vSphereProvisioning'. Changing this will force the creation of a new interface since it's not possible to change the stack once it gets created. (Default: ` + "`" + `defaultTcpipStack` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "services",
					Description: `(Optional) Enabled services setting for this interface. Current possible values are 'vmotion', 'management', and 'vsan'. ### ipv4 options Configures the IPv4 settings of the network interface. Either DHCP or Static IP has to be set.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure the interface's IPv4 stack.`,
				},
				resource.Attribute{
					Name:        "ip",
					Description: `Address of the interface, if DHCP is not set.`,
				},
				resource.Attribute{
					Name:        "netmask",
					Description: `Netmask of the interface, if DHCP is not set.`,
				},
				resource.Attribute{
					Name:        "gw",
					Description: `IP address of the default gateway, if DHCP is not set. ### ipv6 options Configures the IPv6 settings of the network interface. Either DHCP or Autoconfig or Static IP has to be set.`,
				},
				resource.Attribute{
					Name:        "dhcp",
					Description: `Use DHCP to configure the interface's IPv4 stack.`,
				},
				resource.Attribute{
					Name:        "autoconfig",
					Description: `Use IPv6 Autoconfiguration (RFC2462).`,
				},
				resource.Attribute{
					Name:        "addresses",
					Description: `List of IPv6 addresses`,
				},
				resource.Attribute{
					Name:        "gw",
					Description: `IP address of the default gateway, if DHCP or autoconfig is not set. ## Attribute Reference`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vNic. ## Importing An existing vNic can be [imported][docs-import] into this resource via supplying the vNic's ID. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vnic.v1 host-123_vmk2 ` + "`" + `` + "`" + `` + "`" + ` The above would import the vnic ` + "`" + `vmk2` + "`" + ` from host with ID ` + "`" + `host-123` + "`" + `.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the vNic. ## Importing An existing vNic can be [imported][docs-import] into this resource via supplying the vNic's ID. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vnic.v1 host-123_vmk2 ` + "`" + `` + "`" + `` + "`" + ` The above would import the vnic ` + "`" + `vmk2` + "`" + ` from host with ID ` + "`" + `host-123` + "`" + `.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_entity_permissions",
			Category:         "Security",
			ShortDescription: `Provides CRUD operations on a vsphere entity permissions. Permissions can be created on an entity for a given user or group with the specified roles.`,
			Description:      ``,
			Keywords: []string{
				"security",
				"entity",
				"permissions",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "entity_type",
					Description: `(Required) The managed object type, types can be found in the managed object type section [here](https://developer.vmware.com/apis/968/vsphere).`,
				},
				resource.Attribute{
					Name:        "user_or_group",
					Description: `(Required) The user/group getting the permission.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_role",
			Category:         "Security",
			ShortDescription: `Provides CRUD operations on a vsphere role. A role can be created and privileges can be associated with it,`,
			Description:      ``,
			Keywords: []string{
				"security",
				"role",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the role.`,
				},
				resource.Attribute{
					Name:        "role_privileges",
					Description: `(Optional) The privileges to be associated with this role. ## Importing An existing role can be imported into this resource by supplying the role id. An example is below: ` + "`" + `` + "`" + `` + "`" + `hcl terraform import vsphere_role.role1 -709298051 ` + "`" + `` + "`" + `` + "`" + ` ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
	}

	resourcesMap = map[string]int{

		"vsphere_compute_cluster":                         0,
		"vsphere_compute_cluster_host_group":              1,
		"vsphere_compute_cluster_vm_affinity_rule":        2,
		"vsphere_compute_cluster_vm_anti_affinity_rule":   3,
		"vsphere_compute_cluster_vm_dependency_rule":      4,
		"vsphere_compute_cluster_vm_group":                5,
		"vsphere_compute_cluster_vm_host_rule":            6,
		"vsphere_content_library":                         7,
		"vsphere_content_library_item":                    8,
		"vsphere_custom_attribute":                        9,
		"vsphere_datacenter":                              10,
		"vsphere_datastore_cluster":                       11,
		"vsphere_datastore_cluster_vm_anti_affinity_rule": 12,
		"vsphere_distributed_port_group":                  13,
		"vsphere_distributed_virtual_switch":              14,
		"vsphere_dpm_host_override":                       15,
		"vsphere_drs_vm_override":                         16,
		"vsphere_file":                                    17,
		"vsphere_folder":                                  18,
		"vsphere_ha_vm_override":                          19,
		"vsphere_host":                                    20,
		"vsphere_host_port_group":                         21,
		"vsphere_host_virtual_switch":                     22,
		"vsphere_license":                                 23,
		"vsphere_nas_datastore":                           24,
		"vsphere_resource_pool":                           25,
		"vsphere_storage_drs_vm_override":                 26,
		"vsphere_tag":                                     27,
		"vsphere_tag_category":                            28,
		"vsphere_vapp_container":                          29,
		"vsphere_vapp_entity":                             30,
		"vsphere_virtual_disk":                            31,
		"vsphere_virtual_machine":                         32,
		"vsphere_virtual_machine_snapshot":                33,
		"vsphere_vm_storage_policy":                       34,
		"vsphere_vmfs_datastore":                          35,
		"vsphere_vnic":                                    36,
		"vsphere_entity_permissions":                      37,
		"vsphere_role":                                    38,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
