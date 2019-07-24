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
			Category:         "Host and Cluster Management Resources",
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
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) A map of custom attribute ids to attribute value strings to set for the datastore cluster. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "host_system_ids",
					Description: `(Optional) The [managed object IDs][docs-about-morefs] of the hosts to put in the cluster.`,
				},
				resource.Attribute{
					Name:        "host_cluster_exit_timeout",
					Description: `The timeout for each host maintenance mode operation when removing hosts from a cluster. The value is specified in seconds. Default: ` + "`" + `3600` + "`" + ` (1 hour).`,
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
					Description: `(Optional) When ` + "`" + `true` + "`" + `, enables DRS to use data from [vRealize Operations Manager][ref-vsphere-vro] to make proactive DRS recommendations. <sup>[\`,
				},
				resource.Attribute{
					Name:        "drs_advanced_options",
					Description: `(Optional) A key/value map that specifies advanced options for DRS and [DPM](#dpm-options). #### DPM options The following settings control the [Distributed Power Management][ref-vsphere-dpm] (DPM) settings for the cluster. DPM allows the cluster to manage host capacity on-demand depending on the needs of the cluster, powering on hosts when capacity is needed, and placing hosts in standby when there is excess capacity in the cluster. [ref-vsphere-dpm]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.resmgmt.doc/GUID-5E5E349A-4644-4C9C-B434-1C0243EBDC80.html#GUID-5E5E349A-4644-4C9C-B434-1C0243EBDC80`,
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
					Description: `(Optional) Additional delay in seconds after ready condition is met. A VM is considered ready at this point. Default: ` + "`" + `0` + "`" + ` (no delay). <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_vm_restart_timeout",
					Description: `(Optional) The maximum time, in seconds, that vSphere HA will wait for virtual machines in one priority to be ready before proceeding with the next priority. Default: ` + "`" + `600` + "`" + ` (10 minutes). <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_host_isolation_response",
					Description: `(Optional) The action to take on virtual machines when a host has detected that it has been isolated from the rest of the cluster. Can be one of ` + "`" + `none` + "`" + `, ` + "`" + `powerOff` + "`" + `, or ` + "`" + `shutdown` + "`" + `. Default: ` + "`" + `none` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_advanced_options",
					Description: `(Optional) A key/value map that specifies advanced options for vSphere HA. #### HA Virtual Machine Component Protection settings The following settings control Virtual Machine Component Protection (VMCP) in vSphere HA. VMCP gives vSphere HA the ability to monitor a host for datastore accessibility failures, and automate recovery for affected virtual machines. ->`,
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
					Description: `(Optional) Controls the delay in minutes to wait after an APD timeout event to execute the response action defined in [` + "`" + `ha_datastore_apd_response` + "`" + `](#ha_datastore_apd_response). Default: ` + "`" + `3` + "`" + ` minutes. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_vm_monitoring",
					Description: `(Optional) The type of virtual machine monitoring to use when HA is enabled in the cluster. Can be one of ` + "`" + `vmMonitoringDisabled` + "`" + `, ` + "`" + `vmMonitoringOnly` + "`" + `, or ` + "`" + `vmAndAppMonitoring` + "`" + `. Default: ` + "`" + `vmMonitoringDisabled` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_failure_interval",
					Description: `(Optional) If a heartbeat from a virtual machine is not received within this configured interval, the virtual machine is marked as failed. The value is in seconds. Default: ` + "`" + `30` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_vm_minimum_uptime",
					Description: `(Optional) The time, in seconds, that HA waits after powering on a virtual machine before monitoring for heartbeats. Default: ` + "`" + `120` + "`" + ` (2 minutes).`,
				},
				resource.Attribute{
					Name:        "ha_vm_maximum_resets",
					Description: `(Optional) The maximum number of resets that HA will perform to a virtual machine when responding to a failure event. Default: ` + "`" + `3` + "`" + ``,
				},
				resource.Attribute{
					Name:        "ha_vm_maximum_failure_window",
					Description: `(Optional) The length of the reset window in which [` + "`" + `ha_vm_maximum_resets` + "`" + `](#ha_vm_maximum_resets) can operate. When this window expires, no more resets are attempted regardless of the setting configured in ` + "`" + `ha_vm_maximum_resets` + "`" + `. ` + "`" + `-1` + "`" + ` means no window, meaning an unlimited reset time is allotted. The value is specified in seconds. Default: ` + "`" + `-1` + "`" + ` (no window). #### vSphere HA Admission Control settings The following settings control vSphere HA Admission Control, which controls whether or not specific VM operations are permitted in the cluster in order to protect the reliability of the cluster. Based on the constraints defined in these settings, operations such as power on or migration operations may be blocked to ensure that enough capacity remains to react to host failures. #### Admission control modes The [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy) parameter controls the specific mode that Admission Control uses. What settings are available depends on the admission control mode:`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_policy",
					Description: `(Optional) The type of admission control policy to use with vSphere HA. Can be one of ` + "`" + `resourcePercentage` + "`" + `, ` + "`" + `slotPolicy` + "`" + `, ` + "`" + `failoverHosts` + "`" + `, or ` + "`" + `disabled` + "`" + `. Default: ` + "`" + `resourcePercentage` + "`" + `. #### Common Admission Control settings The following settings are available for all Admission Control modes, but will infer different meanings in each mode.`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_host_failure_tolerance",
					Description: `(Optional) The maximum number of failed hosts that admission control tolerates when making decisions on whether to permit virtual machine operations. The maximum is one less than the number of hosts in the cluster. Default: ` + "`" + `1` + "`" + `. <sup>[\`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_performance_tolerance",
					Description: `(Optional) The percentage of resource reduction that a cluster of virtual machines can tolerate in case of a failover. A value of 0 produces warnings only, whereas a value of 100 disables the setting. Default: ` + "`" + `100` + "`" + ` (disabled). #### Admission Control settings for resource percentage mode The following settings control specific settings for Admission Control when ` + "`" + `resourcePercentage` + "`" + ` is selected in [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy).`,
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
					Description: `(Optional) Controls the user-defined percentage of memory resources in the cluster to reserve for failover. Default: ` + "`" + `100` + "`" + `. #### Admission Control settings for slot policy mode The following settings control specific settings for Admission Control when ` + "`" + `slotPolicy` + "`" + ` is selected in [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy).`,
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
					Description: `(Optional) Controls the user-defined memory slot size, in MB. Default: ` + "`" + `100` + "`" + `. #### Admission Control settings for dedicated failover host mode The following settings control specific settings for Admission Control when ` + "`" + `failoverHosts` + "`" + ` is selected in [` + "`" + `ha_admission_control_policy` + "`" + `](#ha_admission_control_policy).`,
				},
				resource.Attribute{
					Name:        "ha_admission_control_failover_host_system_ids",
					Description: `(Optional) Defines the [managed object IDs][docs-about-morefs] of hosts to use as dedicated failover hosts. These hosts are kept as available as possible - admission control will block access to the host, and DRS will ignore the host when making recommendations. #### vSphere HA datastore settings vSphere HA uses datastore heartbeating to determine the health of a particular host. Depending on how your datastores are configured, the settings below may need to be altered to ensure that specific datastores are used over others. If you require a user-defined list of datastores, ensure you select either ` + "`" + `userSelectedDs` + "`" + ` (for user selected only) or ` + "`" + `allFeasibleDsWithUserPreference` + "`" + ` (for automatic selection with preferred overrides) for the [` + "`" + `ha_heartbeat_datastore_policy` + "`" + `](#ha_heartbeat_datastore_policy) setting.`,
				},
				resource.Attribute{
					Name:        "ha_heartbeat_datastore_policy",
					Description: `(Optional) The selection policy for HA heartbeat datastores. Can be one of ` + "`" + `allFeasibleDs` + "`" + `, ` + "`" + `userSelectedDs` + "`" + `, or ` + "`" + `allFeasibleDsWithUserPreference` + "`" + `. Default: ` + "`" + `allFeasibleDsWithUserPreference` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ha_heartbeat_datastore_ids",
					Description: `(Optional) The list of managed object IDs for preferred datastores to use for HA heartbeating. This setting is only useful when [` + "`" + `ha_heartbeat_datastore_policy` + "`" + `](#ha_heartbeat_datastore_policy) is set to either ` + "`" + `userSelectedDs` + "`" + ` or ` + "`" + `allFeasibleDsWithUserPreference` + "`" + `. #### Proactive HA settings The following settings pertain to [Proactive HA][ref-vsphere-proactive-ha], an advanced feature of vSphere HA that allows the cluster to get data from external providers and make decisions based on the data reported. Working with Proactive HA is outside the scope of this document. For more details, see the referenced link in the above paragraph. [ref-vsphere-proactive-ha]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.avail.doc/GUID-3E3B18CC-8574-46FA-9170-CF549B8E55B8.html`,
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
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_compute_cluster_host_group",
			Category:         "Host and Cluster Management Resources",
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
			Category:         "Host and Cluster Management Resources",
			ShortDescription: `Provides a VMware vSphere cluster virtual machine affinity rule. This can be used to manage rules to tell virtual machines to run on the same host.`,
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
			Category:         "Host and Cluster Management Resources",
			ShortDescription: `Provides a VMware vSphere cluster virtual machine anti-affinity rule. This can be used to manage rules to tell virtual machines to run on separate hosts.`,
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
			Category:         "Host and Cluster Management Resources",
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
			Category:         "Host and Cluster Management Resources",
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
			Category:         "Host and Cluster Management Resources",
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
			Type:             "vsphere_custom_attribute",
			Category:         "Inventory Resources",
			ShortDescription: `Provides a vSphere custom attribute resource. This can be used to manage custom attributes in vSphere.`,
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
					Description: `(Optional) The object type that this attribute may be applied to. If not set, the custom attribute may be applied to any object type. For a full list, click [here](#managed-object-types). Forces a new resource if changed. ## Managed Object Types The following table will help you determine what value you need to enter for the managed object type you want the attribute to apply to. Note that if you want a attribute to apply to all objects, leave the type unspecified. <table> <tr><th>Type</th><th>Value</th></tr> <tr><td>Folders</td><td>` + "`" + `Folder` + "`" + `</td></tr> <tr><td>Clusters</td><td>` + "`" + `ClusterComputeResource` + "`" + `</td></tr> <tr><td>Datacenters</td><td>` + "`" + `Datacenter` + "`" + `</td></tr> <tr><td>Datastores</td><td>` + "`" + `Datastore` + "`" + `</td></tr> <tr><td>Datastore Clusters</td><td>` + "`" + `StoragePod` + "`" + `</td></tr> <tr><td>DVS Portgroups</td><td>` + "`" + `DistributedVirtualPortgroup` + "`" + `</td></tr> <tr><td>Distributed vSwitches</td><td>` + "`" + `DistributedVirtualSwitch` + "`" + `<br>` + "`" + `VmwareDistributedVirtualSwitch` + "`" + `</td></tr> <tr><td>Hosts</td><td>` + "`" + `HostSystem` + "`" + `</td></tr> <tr><td>Content Libraries</td><td>` + "`" + `com.vmware.content.Library` + "`" + `</td></tr> <tr><td>Content Library Items</td><td>` + "`" + `com.vmware.content.library.Item` + "`" + `</td></tr> <tr><td>Networks</td><td>` + "`" + `HostNetwork` + "`" + `<br>` + "`" + `Network` + "`" + `<br>` + "`" + `OpaqueNetwork` + "`" + `</td></tr> <tr><td>Resource Pools</td><td>` + "`" + `ResourcePool` + "`" + `</td></tr> <tr><td>vApps</td><td>` + "`" + `VirtualApp` + "`" + `</td></tr> <tr><td>Virtual Machines</td><td>` + "`" + `VirtualMachine` + "`" + `</td></tr> </table> ## Attribute Reference This resource only exports the ` + "`" + `id` + "`" + ` attribute for the vSphere custom attribute. ## Importing An existing custom attribute can be [imported][docs-import] into this resource via its name, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_custom_attribute.attribute terraform-test-attribute ` + "`" + `` + "`" + `` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_datacenter",
			Category:         "Inventory Resources",
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
			Category:         "Storage Resources",
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
					Description: `(Optional) The reservable IOPS threshold setting to use, ` + "`" + `sdrs_io_reservable_percent_threshold` + "`" + ` in the event of ` + "`" + `automatic` + "`" + `, or ` + "`" + `sdrs_io_reservable_iops_threshold` + "`" + ` in the event of ` + "`" + `manual` + "`" + `. Default: ` + "`" + `automatic` + "`" + `. ### Storage DRS disk space load balancing settings The following options control disk space load balancing for Storage DRS on the datastore cluster. ~>`,
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
			Category:         "Storage Resources",
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
			Category:         "Networking Resources",
			ShortDescription: `Provides a vSphere distributed virtual portgroup resource. This can be used to create and manage portgroups on a distributed virtual switch.`,
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
					Description: `(Required) The ID of the DVS to add the port group to. Forces a new resource if changed.`,
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
					Description: `(Optional) An optional formatting policy for naming of the ports in this port group. See the ` + "`" + `portNameFormat` + "`" + ` attribute listed [here][ext-vsphere-portname-format] for details on the format syntax. [ext-vsphere-portname-format]: https://code.vmware.com/apis/196/vsphere#/doc/vim.dvs.DistributedVirtualPortgroup.ConfigInfo.html#portNameFormat`,
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
			Category:         "Networking Resources",
			ShortDescription: `Provides a vSphere distributed virtual switch resource. This can be used to create and manage DVS resources in vCenter.`,
			Description:      ``,
			Keywords: []string{
				"networking",
				"distributed",
				"virtual",
				"switch",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the distributed virtual switch.`,
				},
				resource.Attribute{
					Name:        "datacenter_id",
					Description: `(Required) The ID of the datacenter where the distributed virtual switch will be created. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The folder to create the DVS in. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `(Optional) A detailed description for the DVS.`,
				},
				resource.Attribute{
					Name:        "contact_name",
					Description: `(Optional) The name of the person who is responsible for the DVS.`,
				},
				resource.Attribute{
					Name:        "contact_detail",
					Description: `(Optional) The detailed contact information for the person who is responsible for the DVS.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) An IPv4 address to identify the switch. This is mostly useful when used with the [Netflow arguments](#netflow-arguments) found below.`,
				},
				resource.Attribute{
					Name:        "lacp_api_version",
					Description: `(Optional) The Link Aggregation Control Protocol group version to use with the switch. Possible values are ` + "`" + `singleLag` + "`" + ` and ` + "`" + `multipleLag` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "link_discovery_operation",
					Description: `(Optional) Whether to ` + "`" + `advertise` + "`" + ` or ` + "`" + `listen` + "`" + ` for link discovery traffic.`,
				},
				resource.Attribute{
					Name:        "link_discovery_protocol",
					Description: `(Optional) The discovery protocol type. Valid types are ` + "`" + `cdp` + "`" + ` and ` + "`" + `lldp` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "max_mtu",
					Description: `(Optional) The maximum transmission unit (MTU) for the virtual switch.`,
				},
				resource.Attribute{
					Name:        "multicast_filtering_mode",
					Description: `(Optional) The multicast filtering mode to use with the switch. Can be one of ` + "`" + `legacyFiltering` + "`" + ` or ` + "`" + `snooping` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) - The version of the DVS to create. The default is to create the DVS at the latest version supported by the version of vSphere being used. A DVS can be upgraded to another version, but cannot be downgraded.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) Map of custom attribute ids to attribute value strings to set for virtual switch. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "uplinks",
					Description: `(Optional) A list of strings that uniquely identifies the names of the uplinks on the DVS across hosts. The number of items in this list controls the number of uplinks that exist on the DVS, in addition to the names. See [here](#uplink-name-and-count-control) for an example on how to use this option. ### Host management arguments`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `(Optional) Use the ` + "`" + `host` + "`" + ` block to declare a host specification. The options are:`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Required) The host system ID of the host to add to the DVS.`,
				},
				resource.Attribute{
					Name:        "devices",
					Description: `(Required) The list of NIC devices to map to uplinks on the DVS, added in order they are specified. ### Netflow arguments The following options control settings that you can use to configure Netflow on the DVS:`,
				},
				resource.Attribute{
					Name:        "netflow_active_flow_timeout",
					Description: `(Optional) The number of seconds after which active flows are forced to be exported to the collector. Allowed range is ` + "`" + `60` + "`" + ` to ` + "`" + `3600` + "`" + `. Default: ` + "`" + `60` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "netflow_collector_ip_address",
					Description: `(Optional) IP address for the Netflow collector, using IPv4 or IPv6. IPv6 is supported in vSphere Distributed Switch Version 6.0 or later. Must be set before Netflow can be enabled.`,
				},
				resource.Attribute{
					Name:        "netflow_collector_port",
					Description: `(Optional) Port for the Netflow collector. This must be set before Netflow can be enabled.`,
				},
				resource.Attribute{
					Name:        "netflow_idle_flow_timeout",
					Description: `(Optional) The number of seconds after which idle flows are forced to be exported to the collector. Allowed range is ` + "`" + `10` + "`" + ` to ` + "`" + `600` + "`" + `. Default: ` + "`" + `15` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "netflow_internal_flows_only",
					Description: `(Optional) Whether to limit analysis to traffic that has both source and destination served by the same host. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "netflow_observation_domain_id",
					Description: `(Optional) The observation domain ID for the Netflow collector.`,
				},
				resource.Attribute{
					Name:        "netflow_sampling_rate",
					Description: `(Optional) The ratio of total number of packets to the number of packets analyzed. The default is ` + "`" + `0` + "`" + `, which indicates that the switch should analyze all packets. The maximum value is ` + "`" + `1000` + "`" + `, which indicates an analysis rate of 0.001%. ### Network I/O control arguments The following arguments manage network I/O control. Network I/O control (also known as network resource control) can be used to set up advanced traffic shaping for the DVS, allowing control of various classes of traffic in a fashion similar to how resource pools work for virtual machines. Configuration of network I/O control is also a requirement for the use of network resource pools, if their use is so desired. #### General network I/O control arguments`,
				},
				resource.Attribute{
					Name:        "network_resource_control_enabled",
					Description: `(Optional) Set to ` + "`" + `true` + "`" + ` to enable network I/O control. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "network_resource_control_version",
					Description: `(Optional) The version of network I/O control to use. Can be one of ` + "`" + `version2` + "`" + ` or ` + "`" + `version3` + "`" + `. Default: ` + "`" + `version2` + "`" + `. #### Network I/O control traffic classes There are currently 9 traffic classes that can be used for network I/O control - they are below. Each of these classes has 4 options that can be tuned that are discussed in the next section. <table> <tr><th>Type</th><th>Class Name</th></tr> <tr><td>Fault Tolerance (FT) Traffic</td><td>` + "`" + `faulttolerance` + "`" + `</td></tr> <tr><td>vSphere Replication (VR) Traffic</td><td>` + "`" + `hbr` + "`" + `</td></tr> <tr><td>iSCSI Traffic</td><td>` + "`" + `iscsi` + "`" + `</td></tr> <tr><td>Management Traffic</td><td>` + "`" + `management` + "`" + `</td></tr> <tr><td>NFS Traffic</td><td>` + "`" + `nfs` + "`" + `</td></tr> <tr><td>vSphere Data Protection</td><td>` + "`" + `vdp` + "`" + `</td></tr> <tr><td>Virtual Machine Traffic</td><td>` + "`" + `virtualmachine` + "`" + `</td></tr> <tr><td>vMotion Traffic</td><td>` + "`" + `vmotion` + "`" + `</td></tr> <tr><td>VSAN Traffic</td><td>` + "`" + `vsan` + "`" + `</td></tr> </table> #### Traffic class resource options There are 4 traffic resource options for each class, prefixed with the name of the traffic classes seen above. For example, to set the traffic class resource options for virtual machine traffic, see the example below: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_distributed_virtual_switch" "dvs" { ... virtualmachine_share_level = "custom" virtualmachine_share_count = 150 virtualmachine_maximum_mbit = 200 virtualmachine_reservation_mbit = 20 } ` + "`" + `` + "`" + `` + "`" + ` The options are:`,
				},
				resource.Attribute{
					Name:        "share_level",
					Description: `(Optional) A pre-defined share level that can be assigned to this resource class. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "share_count",
					Description: `(Optional) The number of shares for a custom level. This is ignored if ` + "`" + `share_level` + "`" + ` is not ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "maximum_mbit",
					Description: `(Optional) The maximum amount of bandwidth allowed for this traffic class in Mbits/sec.`,
				},
				resource.Attribute{
					Name:        "reservation_mbit",
					Description: `(Optional) The guaranteed amount of bandwidth for this traffic class in Mbits/sec. ### Default port group policy arguments The following arguments are shared with the [` + "`" + `vsphere_distributed_port_group` + "`" + `][distributed-port-group] resource. Setting them here defines a default policy here that will be inherited by other port groups on this switch that do not have these values otherwise overridden. Not defining these options in a DVS will infer defaults that can be seen in the Terraform state after the initial apply. Of particular note to a DVS are the [HA policy options](#ha-policy-options), which is where the ` + "`" + `active_uplinks` + "`" + ` and ` + "`" + `standby_uplinks` + "`" + ` options are controlled, allowing the ability to create a NIC failover policy that applies to the entire DVS and all portgroups within it that don't override the policy. #### VLAN options The following options control the VLAN behaviour of the port groups the port policy applies to. One one of these 3 options may be set:`,
				},
				resource.Attribute{
					Name:        "vlan",
					Description: `(Optional) The member VLAN for the ports this policy applies to. A value of ` + "`" + `0` + "`" + ` means no VLAN.`,
				},
				resource.Attribute{
					Name:        "vlan_range",
					Description: `(Optional) Used to denote VLAN trunking. Use the ` + "`" + `min_vlan` + "`" + ` and ` + "`" + `max_vlan` + "`" + ` sub-arguments to define the tagged VLAN range. Multiple ` + "`" + `vlan_range` + "`" + ` definitions are allowed, but they must not overlap. Example below: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_distributed_virtual_switch" "dvs" { ... vlan_range { min_vlan = 1 max_vlan = 1000 } vlan_range { min_vlan = 2000 max_vlan = 4094 } } ` + "`" + `` + "`" + `` + "`" + ``,
				},
				resource.Attribute{
					Name:        "port_private_secondary_vlan_id",
					Description: `(Optional) Used to define a secondary VLAN ID when using private VLANs. #### HA policy options The following options control HA policy for ports that this policy applies to:`,
				},
				resource.Attribute{
					Name:        "active_uplinks",
					Description: `(Optional) A list of active uplinks to be used in load balancing. These uplinks need to match the definitions in the [` + "`" + `uplinks` + "`" + `](#uplinks) DVS argument. See [here](#uplink-name-and-count-control) for more details.`,
				},
				resource.Attribute{
					Name:        "standby_uplinks",
					Description: `(Optional) A list of standby uplinks to be used in failover. These uplinks need to match the definitions in the [` + "`" + `uplinks` + "`" + `](#uplinks) DVS argument. See [here](#uplink-name-and-count-control) for more details.`,
				},
				resource.Attribute{
					Name:        "check_beacon",
					Description: `(Optional) Enables beacon probing as an additional measure to detect NIC failure. ~>`,
				},
				resource.Attribute{
					Name:        "failback",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the teaming policy will re-activate failed uplinks higher in precedence when they come back up.`,
				},
				resource.Attribute{
					Name:        "notify_switches",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, the teaming policy will notify the broadcast network of an uplink failover, triggering cache updates.`,
				},
				resource.Attribute{
					Name:        "teaming_policy",
					Description: `(Optional) The uplink teaming policy. Can be one of ` + "`" + `loadbalance_ip` + "`" + `, ` + "`" + `loadbalance_srcmac` + "`" + `, ` + "`" + `loadbalance_srcid` + "`" + `, or ` + "`" + `failover_explicit` + "`" + `. #### LACP options The following options allow the use of LACP for NIC teaming for ports that this policy applies to. ~>`,
				},
				resource.Attribute{
					Name:        "lacp_enabled",
					Description: `(Optional) Enables LACP for the ports that this policy applies to.`,
				},
				resource.Attribute{
					Name:        "lacp_mode",
					Description: `(Optional) The LACP mode. Can be one of ` + "`" + `active` + "`" + ` or ` + "`" + `passive` + "`" + `. #### Security options The following options control security settings for the ports that this policy applies to:`,
				},
				resource.Attribute{
					Name:        "allow_forged_transmits",
					Description: `(Optional) Controls whether or not a virtual network adapter is allowed to send network traffic with a different MAC address than that of its own.`,
				},
				resource.Attribute{
					Name:        "allow_mac_changes",
					Description: `(Optional) Controls whether or not the Media Access Control (MAC) address can be changed.`,
				},
				resource.Attribute{
					Name:        "allow_promiscuous",
					Description: `(Optional) Enable promiscuous mode on the network. This flag indicates whether or not all traffic is seen on a given port. #### Traffic shaping options The following options control traffic shaping settings for the ports that this policy applies to:`,
				},
				resource.Attribute{
					Name:        "ingress_shaping_enabled",
					Description: `(Optional) ` + "`" + `true` + "`" + ` if the traffic shaper is enabled on the port for ingress traffic.`,
				},
				resource.Attribute{
					Name:        "ingress_shaping_average_bandwidth",
					Description: `(Optional) The average bandwidth in bits per second if ingress traffic shaping is enabled on the port.`,
				},
				resource.Attribute{
					Name:        "ingress_shaping_peak_bandwidth",
					Description: `(Optional) The peak bandwidth during bursts in bits per second if ingress traffic shaping is enabled on the port.`,
				},
				resource.Attribute{
					Name:        "ingress_shaping_burst_size",
					Description: `(Optional) The maximum burst size allowed in bytes if ingress traffic shaping is enabled on the port.`,
				},
				resource.Attribute{
					Name:        "egress_shaping_enabled",
					Description: `(Optional) ` + "`" + `true` + "`" + ` if the traffic shaper is enabled on the port for egress traffic.`,
				},
				resource.Attribute{
					Name:        "egress_shaping_average_bandwidth",
					Description: `(Optional) The average bandwidth in bits per second if egress traffic shaping is enabled on the port.`,
				},
				resource.Attribute{
					Name:        "egress_shaping_peak_bandwidth",
					Description: `(Optional) The peak bandwidth during bursts in bits per second if egress traffic shaping is enabled on the port.`,
				},
				resource.Attribute{
					Name:        "egress_shaping_burst_size",
					Description: `(Optional) The maximum burst size allowed in bytes if egress traffic shaping is enabled on the port. #### Miscellaneous options The following are some general options that also affect ports that this policy applies to:`,
				},
				resource.Attribute{
					Name:        "block_all_ports",
					Description: `(Optional) Shuts down all ports in the port groups that this policy applies to, effectively blocking all network access to connected virtual devices.`,
				},
				resource.Attribute{
					Name:        "netflow_enabled",
					Description: `(Optional) Enables Netflow on all ports that this policy applies to.`,
				},
				resource.Attribute{
					Name:        "tx_uplink",
					Description: `(Optional) Forward all traffic transmitted by ports for which this policy applies to its DVS uplinks.`,
				},
				resource.Attribute{
					Name:        "directpath_gen2_allowed",
					Description: `(Optional) Allow VMDirectPath Gen2 for the ports for which this policy applies to. ## Attribute Reference The following attributes are exported:`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_dpm_host_override",
			Category:         "Host and Cluster Management Resources",
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
			Category:         "Host and Cluster Management Resources",
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
			Category:         "Storage Resources",
			ShortDescription: `Provides a VMware vSphere virtual machine file resource. This can be used to upload files (e.g. vmdk disks) from the Terraform host machine to a remote vSphere or copy fields within vSphere.`,
			Description:      ``,
			Keywords: []string{
				"storage",
				"file",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "source_file",
					Description: `(Required) The path to the file being uploaded from the Terraform host to vSphere or copied within vSphere. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "destination_file",
					Description: `(Required) The path to where the file should be uploaded or copied to on vSphere.`,
				},
				resource.Attribute{
					Name:        "source_datacenter",
					Description: `(Optional) The name of a datacenter in which the file will be copied from. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "datacenter",
					Description: `(Optional) The name of a datacenter in which the file will be uploaded to.`,
				},
				resource.Attribute{
					Name:        "source_datastore",
					Description: `(Optional) The name of the datastore in which file will be copied from. Forces a new resource if changed.`,
				},
				resource.Attribute{
					Name:        "datastore",
					Description: `(Required) The name of the datastore in which to upload the file to.`,
				},
				resource.Attribute{
					Name:        "create_directories",
					Description: `(Optional) Create directories in ` + "`" + `destination_file` + "`" + ` path parameter if any missing for copy operation. ~>`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_folder",
			Category:         "Inventory Resources",
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
			Category:         "Host and Cluster Management Resources",
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
					Description: `(Optional) The restart priority for the virtual machine when vSphere detects a host failure. Can be one of ` + "`" + `clusterRestartPriority` + "`" + `, ` + "`" + `lowest` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `medium` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `highest` + "`" + `. Default: ` + "`" + `clusterRestartPriority` + "`" + `.`,
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
					Description: `(Optional) Controls the delay in minutes to wait after an APD timeout event to execute the response action defined in [` + "`" + `ha_datastore_apd_response` + "`" + `](#ha_datastore_apd_response). Use ` + "`" + `-1` + "`" + ` to use the cluster default. Default: ` + "`" + `-1` + "`" + `. <sup>[\`,
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
			Type:             "vsphere_host_port_group",
			Category:         "Networking Resources",
			ShortDescription: `Provides a vSphere Host Port Group Resource. This can be used to configure port groups direct on an ESXi host.`,
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
					Description: `An ID unique to Terraform for this port group. The convention is a prefix, the host system ID, and the port group name. An example would be ` + "`" + `tf-HostPortGroup:host-10:PGTerraformTest` + "`" + `.`,
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
					Description: `A list of ports that currently exist and are used on this port group.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `An ID unique to Terraform for this port group. The convention is a prefix, the host system ID, and the port group name. An example would be ` + "`" + `tf-HostPortGroup:host-10:PGTerraformTest` + "`" + `.`,
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
					Description: `A list of ports that currently exist and are used on this port group.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_host_virtual_switch",
			Category:         "Networking Resources",
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
					Description: `(Required) The list of standby network adapters used for failover.`,
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
			Category:         "Administration Resources",
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
			Category:         "Storage Resources",
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
			Category:         "Host and Cluster Management Resources",
			ShortDescription: `Provides a vSphere resource pool resource. This can be used to create and manage resource pools.`,
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
					Description: `(Required) The [managed object ID][docs-about-morefs] of the parent resource pool. This can be the root resource pool for a cluster or standalone host, or a resource pool itself. When moving a resource pool from one parent resource pool to another, both must share a common root resource pool or the move will fail.`,
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
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is the [managed object ID][docs-about-morefs] of the resource pool. ## Importing An existing resource pool can be [imported][docs-import] into this resource via the path to the resource pool, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_resource_pool.resource_pool /dc1/host/compute-cluster1/Resources/resource-pool1 ` + "`" + `` + "`" + `` + "`" + ` The above would import the resource pool named ` + "`" + `resource-pool1` + "`" + ` that is located in the compute cluster ` + "`" + `compute-cluster1` + "`" + ` in the ` + "`" + `dc1` + "`" + ` datacenter.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_storage_drs_vm_override",
			Category:         "Storage Resources",
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
			Category:         "Inventory Resources",
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
			Category:         "Inventory Resources",
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
			Category:         "Virtual Machine Resources",
			ShortDescription: `Provides a vSphere vApp container resource. This can be used to create and manage vApp container.`,
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
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ## Attribute Reference The only attribute this resource exports is the ` + "`" + `id` + "`" + ` of the resource, which is the [managed object ID][docs-about-morefs] of the resource pool. ## Importing An existing vApp container can be [imported][docs-import] into this resource via the path to the vApp container, using the following command: [docs-import]: https://www.terraform.io/docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_vapp_container.vapp_container /default-dc/host/cluster1/Resources/resource_pool1/vapp_container1 ` + "`" + `` + "`" + `` + "`" + ` The above would import the vApp container named ` + "`" + `vapp-container1` + "`" + ` that is located in the resource pool ` + "`" + `resource-pool1` + "`" + ` that is part of the host cluster ` + "`" + `cluster1` + "`" + ` in the ` + "`" + `dc1` + "`" + ` datacenter.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_vapp_entity",
			Category:         "Virtual Machine Resources",
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
			Category:         "Virtual Machine Resources",
			ShortDescription: `Provides a VMware virtual disk resource. This can be used to create and delete virtual disks.`,
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
					Description: `(Optional) The type of disk to create. Can be one of ` + "`" + `eagerZeroedThick` + "`" + `, ` + "`" + `lazy` + "`" + `, or ` + "`" + `thin` + "`" + `. Default: ` + "`" + `eagerZeroedThick` + "`" + `. For information on what each kind of disk provisioning policy means, click [here][docs-vmware-vm-disk-provisioning]. [docs-vmware-vm-disk-provisioning]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-4C0F4D73-82F2-4B81-8AA7-1DD752A8A5AC.html`,
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
			Category:         "Virtual Machine Resources",
			ShortDescription: `Provides a VMware vSphere virtual machine resource. This can be used to create, modify, and delete virtual machines.`,
			Description:      ``,
			Keywords: []string{
				"virtual",
				"machine",
			},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "resource_pool_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the resource pool to put this virtual machine in. See the section on [virtual machine migration](#virtual-machine-migration) for details on changing this value. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ~>`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the virtual machine's datastore. The virtual machine configuration is placed here, along with any virtual disks that are created where a datastore is not explicitly specified. See the section on [virtual machine migration](#virtual-machine-migration) for details on changing this value.`,
				},
				resource.Attribute{
					Name:        "datastore_cluster_id",
					Description: `(Optional) The [managed object reference ID][docs-about-morefs] of the datastore cluster ID to use. This setting applies to entire virtual machine and implies that you wish to use Storage DRS with this virtual machine. See the section on [virtual machine migration](#virtual-machine-migration) for details on changing this value. ~>`,
				},
				resource.Attribute{
					Name:        "folder",
					Description: `(Optional) The path to the folder to put this virtual machine in, relative to the datacenter that the resource pool is in.`,
				},
				resource.Attribute{
					Name:        "host_system_id",
					Description: `(Optional) An optional [managed object reference ID][docs-about-morefs] of a host to put this virtual machine on. See the section on [virtual machine migration](#virtual-machine-migration) for details on changing this value. If a ` + "`" + `host_system_id` + "`" + ` is not supplied, vSphere will select a host in the resource pool to place the virtual machine, according to any defaults or DRS policies in place.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `(Required) A specification for a virtual disk device on this virtual machine. See [disk options](#disk-options) below.`,
				},
				resource.Attribute{
					Name:        "network_interface",
					Description: `(Required) A specification for a virtual NIC on this virtual machine. See [network interface options](#network-interface-options) below.`,
				},
				resource.Attribute{
					Name:        "cdrom",
					Description: `(Optional) A specification for a CDROM device on this virtual machine. See [CDROM options](#cdrom-options) below.`,
				},
				resource.Attribute{
					Name:        "clone",
					Description: `(Optional) When specified, the VM will be created as a clone of a specified template. Optional customization options can be submitted as well. See [creating a virtual machine from a template](#creating-a-virtual-machine-from-a-template) for more details. ~>`,
				},
				resource.Attribute{
					Name:        "vapp",
					Description: `(Optional) Optional vApp configuration. The only sub-key available is ` + "`" + `properties` + "`" + `, which is a key/value map of properties for virtual machines imported from OVF or OVA files. See [Using vApp properties to supply OVF/OVA configuration](#using-vapp-properties-to-supply-ovf-ova-configuration) for more details.`,
				},
				resource.Attribute{
					Name:        "guest_id",
					Description: `(Optional) The guest ID for the operating system type. For a full list of possible values, see [here][vmware-docs-guest-ids]. Default: ` + "`" + `other-64` + "`" + `. [vmware-docs-guest-ids]: https://pubs.vmware.com/vsphere-6-5/topic/com.vmware.wssdk.apiref.doc/vim.vm.GuestOsDescriptor.GuestOsIdentifier.html`,
				},
				resource.Attribute{
					Name:        "alternate_guest_name",
					Description: `(Optional) The guest name for the operating system when ` + "`" + `guest_id` + "`" + ` is ` + "`" + `other` + "`" + ` or ` + "`" + `other-64` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "annotation",
					Description: `(Optional) A user-provided description of the virtual machine. The default is no annotation.`,
				},
				resource.Attribute{
					Name:        "firmware",
					Description: `(Optional) The firmware interface to use on the virtual machine. Can be one of ` + "`" + `bios` + "`" + ` or ` + "`" + `EFI` + "`" + `. Default: ` + "`" + `bios` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "extra_config",
					Description: `(Optional) Extra configuration data for this virtual machine. Can be used to supply advanced parameters not normally in configuration, such as instance metadata. ~>`,
				},
				resource.Attribute{
					Name:        "scsi_type",
					Description: `(Optional) The type of SCSI bus this virtual machine will have. Can be one of lsilogic (LSI Logic Parallel), lsilogic-sas (LSI Logic SAS) or pvscsi (VMware Paravirtual). Defualt: ` + "`" + `pvscsi` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scsi_bus_sharing",
					Description: `(Optional) Mode for sharing the SCSI bus. The modes are physicalSharing, virtualSharing, and noSharing. Default: ` + "`" + `noSharing` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `(Optional) The IDs of any tags to attach to this resource. See [here][docs-applying-tags] for a reference on how to apply tags. [docs-applying-tags]: /docs/providers/vsphere/r/tag.html#using-tags-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "custom_attributes",
					Description: `(Optional) Map of custom attribute ids to attribute value strings to set for virtual machine. See [here][docs-setting-custom-attributes] for a reference on how to set values for custom attributes. [docs-setting-custom-attributes]: /docs/providers/vsphere/r/custom_attribute.html#using-custom-attributes-in-a-supported-resource ~>`,
				},
				resource.Attribute{
					Name:        "num_cpus",
					Description: `(Optional) The number of virtual processors to assign to this virtual machine. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "num_cores_per_socket",
					Description: `(Optional) The number of cores to distribute among the CPUs in this virtual machine. If specified, the value supplied to ` + "`" + `num_cpus` + "`" + ` must be evenly divisible by this value. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_add_enabled",
					Description: `(Optional) Allow CPUs to be added to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "cpu_hot_remove_enabled",
					Description: `(Optional) Allow CPUs to be removed to this virtual machine while it is running.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `(Optional) The size of the virtual machine's memory, in MB. Default: ` + "`" + `1024` + "`" + ` (1 GB).`,
				},
				resource.Attribute{
					Name:        "memory_hot_add_enabled",
					Description: `(Optional) Allow memory to be added to this virtual machine while it is running. ~>`,
				},
				resource.Attribute{
					Name:        "boot_delay",
					Description: `(Optional) The number of milliseconds to wait before starting the boot sequence. The default is no delay.`,
				},
				resource.Attribute{
					Name:        "efi_secure_boot_enabled",
					Description: `(Optional) When the ` + "`" + `firmware` + "`" + ` type is set to is ` + "`" + `efi` + "`" + `, this enables EFI secure boot. Default: ` + "`" + `false` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "boot_retry_delay",
					Description: `(Optional) The number of milliseconds to wait before retrying the boot sequence. This only valid if ` + "`" + `boot_retry_enabled` + "`" + ` is true. Default: ` + "`" + `10000` + "`" + ` (10 seconds).`,
				},
				resource.Attribute{
					Name:        "boot_retry_enabled",
					Description: `(Optional) If set to true, a virtual machine that fails to boot will try again after the delay defined in ` + "`" + `boot_retry_delay` + "`" + `. Default: ` + "`" + `false` + "`" + `. ### VMware Tools options The following options control VMware tools options on the virtual machine:`,
				},
				resource.Attribute{
					Name:        "sync_time_with_host",
					Description: `(Optional) Enable guest clock synchronization with the host. Requires VMware tools to be installed. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_after_power_on",
					Description: `(Optional) Enable the execution of post-power-on scripts when VMware tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_after_resume",
					Description: `(Optional) Enable the execution of post-resume scripts when VMware tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_before_guest_reboot",
					Description: `(Optional) Enable the execution of pre-reboot scripts when VMware tools is installed. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_before_guest_shutdown",
					Description: `(Optional) Enable the execution of pre-shutdown scripts when VMware tools is installed. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "run_tools_scripts_before_guest_standby",
					Description: `(Optional) Enable the execution of pre-standby scripts when VMware tools is installed. Default: ` + "`" + `true` + "`" + `. ### Resource allocation options The following options allow control over CPU and memory allocation on the virtual machine. Note that the resource pool that this VM is in may affect these options.`,
				},
				resource.Attribute{
					Name:        "cpu_limit",
					Description: `(Optional) The maximum amount of CPU (in MHz) that this virtual machine can consume, regardless of available resources. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "cpu_reservation",
					Description: `(Optional) The amount of CPU (in MHz) that this virtual machine is guaranteed. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "cpu_share_level",
					Description: `(Optional) The allocation level for CPU resources. Can be one of ` + "`" + `high` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_share_count",
					Description: `(Optional) The number of CPU shares allocated to the virtual machine when the ` + "`" + `cpu_share_level` + "`" + ` is ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_limit",
					Description: `(Optional) The maximum amount of memory (in MB) that this virtual machine can consume, regardless of available resources. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "memory_reservation",
					Description: `(Optional) The amount of memory (in MB) that this virtual machine is guaranteed. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "memory_share_level",
					Description: `(Optional) The allocation level for memory resources. Can be one of ` + "`" + `high` + "`" + `, ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `custom` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "memory_share_count",
					Description: `(Optional) The number of memory shares allocated to the virtual machine when the ` + "`" + `memory_share_level` + "`" + ` is ` + "`" + `custom` + "`" + `. ### Advanced options The following options control advanced operation of the virtual machine, or control various parts of Terraform workflow, and should not need to be modified during basic operation of the resource. Only change these options if they are explicitly required, or if you are having trouble with Terraform's default behavior.`,
				},
				resource.Attribute{
					Name:        "enable_disk_uuid",
					Description: `(Optional) Expose the UUIDs of attached virtual disks to the virtual machine, allowing access to them in the guest. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "hv_mode",
					Description: `(Optional) The (non-nested) hardware virtualization setting for this virtual machine. Can be one of ` + "`" + `hvAuto` + "`" + `, ` + "`" + `hvOn` + "`" + `, or ` + "`" + `hvOff` + "`" + `. Default: ` + "`" + `hvAuto` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "ept_rvi_mode",
					Description: `(Optional) The EPT/RVI (hardware memory virtualization) setting for this virtual machine. Can be one of ` + "`" + `automatic` + "`" + `, ` + "`" + `on` + "`" + `, or ` + "`" + `off` + "`" + `. Default: ` + "`" + `automatic` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "nested_hv_enabled",
					Description: `(Optional) Enable nested hardware virtualization on this virtual machine, facilitating nested virtualization in the guest. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "enable_logging",
					Description: `(Optional) Enable logging of virtual machine events to a log file stored in the virtual machine directory. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "cpu_performance_counters_enabled",
					Description: `(Optional) Enable CPU performance counters on this virtual machine. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "swap_placement_policy",
					Description: `(Optional) The swap file placement policy for this virtual machine. Can be one of ` + "`" + `inherit` + "`" + `, ` + "`" + `hostLocal` + "`" + `, or ` + "`" + `vmDirectory` + "`" + `. Default: ` + "`" + `inherit` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "latency_sensitivity",
					Description: `(Optional) Controls the scheduling delay of the virtual machine. Use a higher sensitivity for applications that require lower latency, such as VOIP, media player applications, or applications that require frequent access to mouse or keyboard devices. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `medium` + "`" + `, or ` + "`" + `high` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "wait_for_guest_net_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for an available IP address on this virtual machine's NICs. Older versions of VMware Tools do not populate this property. In those cases, this waiter can be disabled and the [` + "`" + `wait_for_guest_ip_timeout` + "`" + `](#wait_for_guest_ip_timeout) waiter can be used instead. A value less than 1 disables the waiter. Default: 5 minutes.`,
				},
				resource.Attribute{
					Name:        "wait_for_guest_net_routable",
					Description: `(Optional) Controls whether or not the guest network waiter waits for a routable address. When ` + "`" + `false` + "`" + `, the waiter does not wait for a default gateway, nor are IP addresses checked against any discovered default gateways as part of its success criteria. This property is ignored if the [` + "`" + `wait_for_guest_ip_timeout` + "`" + `](#wait_for_guest_ip_timeout) waiter is used. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "wait_for_guest_ip_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for an available guest IP address on this virtual machine. This should only be used if your version of VMware Tools does not allow the [` + "`" + `wait_for_guest_net_timeout` + "`" + `](#wait_for_guest_net_timeout) waiter to be used. A value less than 1 disables the waiter. Default: 0.`,
				},
				resource.Attribute{
					Name:        "ignored_guest_ips",
					Description: `(Optional) List of IP addresses to ignore while waiting for an available IP address using either of the waiters. Any IP addresses in this list will be ignored if they show up so that the waiter will continue to wait for a real IP address. Default: [].`,
				},
				resource.Attribute{
					Name:        "shutdown_wait_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for a graceful guest shutdown when making necessary updates to the virtual machine. If ` + "`" + `force_power_off` + "`" + ` is set to true, the VM will be force powered-off after this timeout, otherwise an error is returned. Default: 3 minutes.`,
				},
				resource.Attribute{
					Name:        "migrate_wait_timeout",
					Description: `(Optional) The amount of time, in minutes, to wait for a virtual machine migration to complete before failing. Default: 10 minutes. Also see the section on [virtual machine migration](#virtual-machine-migration).`,
				},
				resource.Attribute{
					Name:        "force_power_off",
					Description: `(Optional) If a guest shutdown failed or timed out while updating or destroying (see [` + "`" + `shutdown_wait_timeout` + "`" + `](#shutdown_wait_timeout)), force the power-off of the virtual machine. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "scsi_controller_count",
					Description: `(Optional) The number of SCSI controllers that Terraform manages on this virtual machine. This directly affects the amount of disks you can add to the virtual machine and the maximum disk unit number. Note that lowering this value does not remove controllers. Default: ` + "`" + `1` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "label",
					Description: `(Required) A label for the disk. Forces a new disk if changed. ~>`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) An alias for both ` + "`" + `label` + "`" + ` and ` + "`" + `path` + "`" + `, the latter when using ` + "`" + `attach` + "`" + `. Required if not using ` + "`" + `label` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `(Required) The size of the disk, in GiB.`,
				},
				resource.Attribute{
					Name:        "unit_number",
					Description: `(Optional) The disk number on the SCSI bus. The maximum value for this setting is the value of [` + "`" + `scsi_controller_count` + "`" + `](#scsi_controller_count) times 15, minus 1 (so ` + "`" + `14` + "`" + `, ` + "`" + `29` + "`" + `, ` + "`" + `44` + "`" + `, and ` + "`" + `59` + "`" + `, for 1-4 controllers respectively). The default is ` + "`" + `0` + "`" + `, for which one disk must be set to. Duplicate unit numbers are not allowed.`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) A [managed object reference ID][docs-about-morefs] to the datastore for this virtual disk. The default is to use the datastore of the virtual machine. See the section on [virtual machine migration](#virtual-machine-migration) for details on changing this value. ~>`,
				},
				resource.Attribute{
					Name:        "attach",
					Description: `(Optional) Attach an external disk instead of creating a new one. Implies and conflicts with ` + "`" + `keep_on_remove` + "`" + `. If set, you cannot set ` + "`" + `size` + "`" + `, ` + "`" + `eagerly_scrub` + "`" + `, or ` + "`" + `thin_provisioned` + "`" + `. Must set ` + "`" + `path` + "`" + ` if used. ~>`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) When using ` + "`" + `attach` + "`" + `, this parameter controls the path of a virtual disk to attach externally. Otherwise, it is a computed attribute that contains the virtual disk's current filename.`,
				},
				resource.Attribute{
					Name:        "keep_on_remove",
					Description: `(Optional) Keep this disk when removing the device or destroying the virtual machine. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_mode",
					Description: `(Optional) The mode of this this virtual disk for purposes of writes and snapshotting. Can be one of ` + "`" + `append` + "`" + `, ` + "`" + `independent_nonpersistent` + "`" + `, ` + "`" + `independent_persistent` + "`" + `, ` + "`" + `nonpersistent` + "`" + `, ` + "`" + `persistent` + "`" + `, or ` + "`" + `undoable` + "`" + `. Default: ` + "`" + `persistent` + "`" + `. For an explanation of options, click [here][vmware-docs-disk-mode]. [vmware-docs-disk-mode]: https://pubs.vmware.com/vsphere-6-5/topic/com.vmware.wssdk.apiref.doc/vim.vm.device.VirtualDiskOption.DiskMode.html`,
				},
				resource.Attribute{
					Name:        "eagerly_scrub",
					Description: `(Optional) If set to ` + "`" + `true` + "`" + `, the disk space is zeroed out on VM creation. This will delay the creation of the disk or virtual machine. Cannot be set to ` + "`" + `true` + "`" + ` when ` + "`" + `thin_provisioned` + "`" + ` is ` + "`" + `true` + "`" + `. See the section on [picking a disk type](#picking-a-disk-type). Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "thin_provisioned",
					Description: `(Optional) If ` + "`" + `true` + "`" + `, this disk is thin provisioned, with space for the file being allocated on an as-needed basis. Cannot be set to ` + "`" + `true` + "`" + ` when ` + "`" + `eagerly_scrub` + "`" + ` is ` + "`" + `true` + "`" + `. See the section on [picking a disk type](#picking-a-disk-type). Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "disk_sharing",
					Description: `(Optional) The sharing mode of this virtual disk. Can be one of ` + "`" + `sharingMultiWriter` + "`" + ` or ` + "`" + `sharingNone` + "`" + `. Default: ` + "`" + `sharingNone` + "`" + `. ~>`,
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
					Description: `(Optional) The I/O reservation (guarantee) that this disk has, in IOPS. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "io_share_level",
					Description: `(Optional) The share allocation level for this disk. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "io_share_count",
					Description: `(Optional) The share count for this disk when the share level is ` + "`" + `custom` + "`" + `. #### Computed disk attributes`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the virtual disk's VMDK file. This is used to track the virtual disk on the virtual machine. #### Picking a disk type The ` + "`" + `eagerly_scrub` + "`" + ` and ` + "`" + `thin_provisioned` + "`" + ` options control the space allocation type of a virtual disk. These show up in the vSphere console as a unified enumeration of options, the equivalents of which are explained below. The defaults in Terraform are the equivalent of thin provisioning.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Required) The [managed object reference ID][docs-about-morefs] of the network to connect this interface to.`,
				},
				resource.Attribute{
					Name:        "adapter_type",
					Description: `(Optional) The network interface type. Can be one of ` + "`" + `e1000` + "`" + `, ` + "`" + `e1000e` + "`" + `, or ` + "`" + `vmxnet3` + "`" + `. Default: ` + "`" + `vmxnet3` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "use_static_mac",
					Description: `(Optional) If true, the ` + "`" + `mac_address` + "`" + ` field is treated as a static MAC address and set accordingly. Setting this to ` + "`" + `true` + "`" + ` requires ` + "`" + `mac_address` + "`" + ` to be set. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of this network interface. Can only be manually set if ` + "`" + `use_static_mac` + "`" + ` is true, otherwise this is a computed value that gives the current MAC address of this interface.`,
				},
				resource.Attribute{
					Name:        "bandwidth_limit",
					Description: `(Optional) The upper bandwidth limit of this network interface, in Mbits/sec. The default is no limit.`,
				},
				resource.Attribute{
					Name:        "bandwidth_reservation",
					Description: `(Optional) The bandwidth reservation of this network interface, in Mbits/sec. The default is no reservation.`,
				},
				resource.Attribute{
					Name:        "bandwidth_share_level",
					Description: `(Optional) The bandwidth share allocation level for this interface. Can be one of ` + "`" + `low` + "`" + `, ` + "`" + `normal` + "`" + `, ` + "`" + `high` + "`" + `, or ` + "`" + `custom` + "`" + `. Default: ` + "`" + `normal` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "bandwidth_share_count",
					Description: `(Optional) The share count for this network interface when the share level is ` + "`" + `custom` + "`" + `. ### CDROM options A single virtual CDROM device can be created and attached to the virtual machine. The resource supports attaching a CDROM from a datastore ISO or using a remote client device. An example is below: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { ... cdrom { datastore_id = "${data.vsphere_datastore.iso_datastore.id}" path = "ISOs/os-livecd.iso" } } ` + "`" + `` + "`" + `` + "`" + ` The options are:`,
				},
				resource.Attribute{
					Name:        "client_device",
					Description: `(Optional) Indicates whether the device should be backed by remote client device. Conflicts with ` + "`" + `datastore_id` + "`" + ` and ` + "`" + `path` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "datastore_id",
					Description: `(Optional) The datastore ID that the ISO is located in. Requried for using a datastore ISO. Conflicts with ` + "`" + `client_device` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "path",
					Description: `(Optional) The path to the ISO file. Requried for using a datastore ISO. Conflicts with ` + "`" + `client_device` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `The ID of the device within the virtual machine.`,
				},
				resource.Attribute{
					Name:        "device_address",
					Description: `An address internal to Terraform that helps locate the device when ` + "`" + `key` + "`" + ` is unavailable. This follows a convention of ` + "`" + `CONTROLLER_TYPE:BUS_NUMBER:UNIT_NUMBER` + "`" + `. Example: ` + "`" + `scsi:0:1` + "`" + ` means device unit 1 on SCSI bus 0. ## Creating a Virtual Machine from a Template The ` + "`" + `clone` + "`" + ` block can be used to create a new virtual machine from an existing virtual machine or template. The resource supports both making a complete copy of a virtual machine, or cloning from a snapshot (otherwise known as a linked clone). See the [cloning and customization example](#cloning-and-customization-example) for a usage synopsis. ~>`,
				},
				resource.Attribute{
					Name:        "template_uuid",
					Description: `(Required) The UUID of the source virtual machine or template.`,
				},
				resource.Attribute{
					Name:        "linked_clone",
					Description: `(Optional) Clone this virtual machine from a snapshot. Templates must have a single snapshot only in order to be eligible. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The timeout, in minutes, to wait for the virtual machine clone to complete. Default: 30 minutes.`,
				},
				resource.Attribute{
					Name:        "customize",
					Description: `(Optional) The customization spec for this clone. This allows the user to configure the virtual machine post-clone. For more details, see [virtual machine customization](#virtual-machine-customization). ### Virtual machine customization As part of the ` + "`" + `clone` + "`" + ` operation, a virtual machine can be [customized][vmware-docs-customize] to configure host, network, or licensing settings. [vmware-docs-customize]: https://docs.vmware.com/en/VMware-vSphere/6.5/com.vmware.vsphere.vm_admin.doc/GUID-58E346FF-83AE-42B8-BE58-253641D257BC.html To perform virtual machine customization as a part of the clone process, specify the ` + "`" + `customize` + "`" + ` block with the respective customization options, nested within the ` + "`" + `clone` + "`" + ` block. See the [cloning and customization example](#cloning-and-customization-example) for a usage synopsis. The settings for ` + "`" + `customize` + "`" + ` are as follows: #### Customization timeout settings`,
				},
				resource.Attribute{
					Name:        "timeout",
					Description: `(Optional) The time, in minutes that Terraform waits for customization to complete before failing. The default is 10 minutes, and setting the value to 0 or a negative value disables the waiter altogether. #### Network interface settings These settings, which should be specified in nested ` + "`" + `network_interface` + "`" + ` blocks within [` + "`" + `customize` + "`" + `](#virtual-machine-customization), configure network interfaces on a per-interface basis and are matched up to [` + "`" + `network_interface` + "`" + `](#network-interface-options) devices in the order they are declared. Given the following example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { ... network_interface { network_id = "${data.vsphere_network.public.id}" } network_interface { network_id = "${data.vsphere_network.private.id}" } clone { ... customize { ... network_interface { ipv4_address = "10.0.0.10" ipv4_netmask = 24 } network_interface { ipv4_address = "172.16.0.10" ipv4_netmask = 24 } ipv4_gateway = "10.0.0.1" } } } ` + "`" + `` + "`" + `` + "`" + ` The first set of ` + "`" + `network_interface` + "`" + ` data would be assigned to the ` + "`" + `public` + "`" + ` interface, and the second to the ` + "`" + `private` + "`" + ` interface. To use DHCP, declare an empty ` + "`" + `network_interface` + "`" + ` block for each interface being configured. So the above example would look like: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { ... network_interface { network_id = "${data.vsphere_network.public.id}" } network_interface { network_id = "${data.vsphere_network.private.id}" } clone { ... customize { ... network_interface {} network_interface {} } } } ` + "`" + `` + "`" + `` + "`" + ` The options are:`,
				},
				resource.Attribute{
					Name:        "dns_server_list",
					Description: `(Optional) Network interface-specific DNS server settings for Windows operating systems. Ignored on Linux and possibly other operating systems - for those systems, please see the [global DNS settings](#global-dns-settings) section.`,
				},
				resource.Attribute{
					Name:        "dns_domain",
					Description: `(Optional) Network interface-specific DNS search domain for Windows operating systems. Ignored on Linux and possibly other operating systems - for those systems, please see the [global DNS settings](#global-dns-settings) section.`,
				},
				resource.Attribute{
					Name:        "ipv4_address",
					Description: `(Optional) The IPv4 address assigned to this network adapter. If left blank or not included, DHCP is used.`,
				},
				resource.Attribute{
					Name:        "ipv6_address",
					Description: `(Optional) The IPv6 address assigned to this network adapter. If left blank or not included, auto-configuration is used.`,
				},
				resource.Attribute{
					Name:        "ipv6_netmask",
					Description: `(Optional) The IPv6 subnet mask, in bits (example: ` + "`" + `32` + "`" + `). ~>`,
				},
				resource.Attribute{
					Name:        "ipv4_gateway",
					Description: `(Optional) The IPv4 default gateway when using ` + "`" + `network_interface` + "`" + ` customization on the virtual machine.`,
				},
				resource.Attribute{
					Name:        "ipv6_gateway",
					Description: `(Optional) The IPv6 default gateway when using ` + "`" + `network_interface` + "`" + ` customization on the virtual machine. #### Global DNS settings The following settings configure DNS globally, generally for Linux systems. For Windows systems, this is done per-interface, see [network interface settings](#network-interface-settings).`,
				},
				resource.Attribute{
					Name:        "dns_server_list",
					Description: `The list of DNS servers to configure on a virtual machine.`,
				},
				resource.Attribute{
					Name:        "dns_suffix_list",
					Description: `A list of DNS search domains to add to the DNS configuration on the virtual machine. #### Linux customization options The settings in the ` + "`" + `linux_options` + "`" + ` block pertain to Linux guest OS customization. If you are customizing a Linux operating system, this section must be included. Example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { ... clone { ... customize { ... linux_options { host_name = "terraform-test" domain = "test.internal" } } } } ` + "`" + `` + "`" + `` + "`" + ` The options are:`,
				},
				resource.Attribute{
					Name:        "host_name",
					Description: `(Required) The host name for this machine. This, along with ` + "`" + `domain` + "`" + `, make up the FQDN of this virtual machine.`,
				},
				resource.Attribute{
					Name:        "domain",
					Description: `(Required) The domain name for this machine. This, along with ` + "`" + `host_name` + "`" + `, make up the FQDN of this virtual machine.`,
				},
				resource.Attribute{
					Name:        "hw_clock_utc",
					Description: `(Optional) Tells the operating system that the hardware clock is set to UTC. Default: ` + "`" + `true` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) Sets the time zone. For a list of possible combinations, click [here][vmware-docs-valid-linux-tzs]. The default is UTC. [vmware-docs-valid-linux-tzs]: https://pubs.vmware.com/vsphere-6-5/topic/com.vmware.wssdk.apiref.doc/timezone.html #### Windows customization options The settings in the ` + "`" + `windows_options` + "`" + ` block pertain to Windows guest OS customization. If you are customizing a Windows operating system, this section must be included. Example: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { ... clone { ... customize { ... windows_options { computer_name = "terraform-test" workgroup = "test" admin_password = "VMw4re" } } } } ` + "`" + `` + "`" + `` + "`" + ` The options are:`,
				},
				resource.Attribute{
					Name:        "computer_name",
					Description: `(Required) The computer name of this virtual machine.`,
				},
				resource.Attribute{
					Name:        "admin_password",
					Description: `(Optional) The administrator password for this virtual machine. ~>`,
				},
				resource.Attribute{
					Name:        "workgroup",
					Description: `(Optional) The workgroup name for this virtual machine. One of this or ` + "`" + `join_domain` + "`" + ` must be included.`,
				},
				resource.Attribute{
					Name:        "join_domain",
					Description: `(Optional) The domain to join for this virtual machine. One of this or ` + "`" + `workgroup` + "`" + ` must be included.`,
				},
				resource.Attribute{
					Name:        "domain_admin_user",
					Description: `(Optional) The user of the domain administrator used to join this virtual machine to the domain. Required if you are setting ` + "`" + `join_domain` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "domain_admin_password",
					Description: `(Optional) The password of the domain administrator used to join this virtual machine to the domain. Required if you are setting ` + "`" + `join_domain` + "`" + `. ~>`,
				},
				resource.Attribute{
					Name:        "full_name",
					Description: `(Optional) The full name of the user of this virtual machine. This populates the "user" field in the general Windows system information. Default: ` + "`" + `Administrator` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "organization_name",
					Description: `(Optional) The organization name this virtual machine is being installed for. This populates the "organization" field in the general Windows system information. Default: ` + "`" + `Managed by Terraform` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "product_key",
					Description: `(Optional) The product key for this virtual machine. The default is no key.`,
				},
				resource.Attribute{
					Name:        "run_once_command_list",
					Description: `(Optional) A list of commands to run at first user logon, after guest customization. Each command is limited by the API to 260 characters.`,
				},
				resource.Attribute{
					Name:        "auto_logon",
					Description: `(Optional) Specifies whether or not the VM automatically logs on as Administrator. Default: ` + "`" + `false` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "auto_logon_count",
					Description: `(Optional) Specifies how many times the VM should auto-logon the Administrator account when ` + "`" + `auto_logon` + "`" + ` is true. This should be set accordingly to ensure that all of your commands that run in ` + "`" + `run_once_command_list` + "`" + ` can log in to run. Default: ` + "`" + `1` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "time_zone",
					Description: `(Optional) The new time zone for the virtual machine. This is a numeric, sysprep-dictated, timezone code. For a list of codes, click [here][ms-docs-valid-sysprep-tzs]. The default is ` + "`" + `85` + "`" + ` (GMT/UTC). [ms-docs-valid-sysprep-tzs]: https://msdn.microsoft.com/en-us/library/ms912391(v=winembedded.11).aspx #### Supplying your own SysPrep file Alternative to the ` + "`" + `windows_options` + "`" + ` supplied above, you can instead supply your own ` + "`" + `sysprep.inf` + "`" + ` file contents via the ` + "`" + `windows_sysprep_text` + "`" + ` option. This allows full control of the customization process out-of-band of vSphere. Example below: ` + "`" + `` + "`" + `` + "`" + `hcl resource "vsphere_virtual_machine" "vm" { ... clone { ... customize { ... windows_sysprep_text = "${file("${path.module}/sysprep.inf")}" } } } ` + "`" + `` + "`" + `` + "`" + ` Note this option is mutually exclusive to ` + "`" + `windows_options` + "`" + ` - one must not be included if the other is specified. ### Using vApp properties to supply OVF/OVA configuration Alternative to the settings in ` + "`" + `customize` + "`" + `, one can use the settings in the ` + "`" + `properties` + "`" + ` section of the ` + "`" + `vapp` + "`" + ` block to supply configuration parameters to a virtual machine cloned from a template that came from an imported OVF or OVA file. Both GuestInfo and ISO transport methods are supported. For templates that use ISO transport, a CDROM backed by client device is required. See [CDROM options](#cdrom-options) for details. ~>`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "reboot_required",
					Description: `Value internal to Terraform used to determine if a configuration set change requires a reboot. This value is only useful during an update process and gets reset on refresh.`,
				},
				resource.Attribute{
					Name:        "vmware_tools_status",
					Description: `The state of VMware tools in the guest. This will determine the proper course of action for some device operations.`,
				},
				resource.Attribute{
					Name:        "vmx_path",
					Description: `The path of the virtual machine's configuration file in the VM's datastore.`,
				},
				resource.Attribute{
					Name:        "imported",
					Description: `This is flagged if the virtual machine has been imported, or the state has been migrated from a previous version of the resource. It influences the behavior of the first post-import apply operation. See the section on [importing](#importing) below.`,
				},
				resource.Attribute{
					Name:        "change_version",
					Description: `A unique identifier for a given version of the last configuration applied, such the timestamp of the last update to the configuration.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the virtual machine. Also exposed as the ` + "`" + `id` + "`" + ` of the resource.`,
				},
				resource.Attribute{
					Name:        "default_ip_address",
					Description: `The IP address selected by Terraform to be used with any [provisioners][tf-docs-provisioners] configured on this resource. Whenever possible, this is the first IPv4 address that is reachable through the default gateway configured on the machine, then the first reachable IPv6 address, and then the first general discovered address if neither exist. If VMware tools is not running on the virtual machine, or if the VM is powered off, this value will be blank.`,
				},
				resource.Attribute{
					Name:        "guest_ip_addresses",
					Description: `The current list of IP addresses on this machine, including the value of ` + "`" + `default_ip_address` + "`" + `. If VMware tools is not running on the virtual machine, or if the VM is powered off, this list will be empty.`,
				},
				resource.Attribute{
					Name:        "vapp_transport",
					Description: `Computed value which is only valid for cloned virtual machines. A list of vApp transport methods supported by the source virtual machine or template. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Importing An existing virtual machine can be [imported][docs-import] into this resource via supplying the full path to the virtual machine. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_virtual_machine.vm /dc1/vm/srv1 ` + "`" + `` + "`" + `` + "`" + ` The above would import the virtual machine named ` + "`" + `srv1` + "`" + ` that is located in the ` + "`" + `dc1` + "`" + ` datacenter. ### Additional requirements and notes for importing Many of the same requirements for [cloning](#additional-requirements-and-notes-for-cloning) apply to importing, although since importing writes directly to state, a lot of these rules cannot be enforced at import time, so every effort should be made to ensure the correctness of the configuration before the import. In addition to these rules, the following extra rules apply to importing:`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `The UUID of the virtual machine.`,
				},
				resource.Attribute{
					Name:        "reboot_required",
					Description: `Value internal to Terraform used to determine if a configuration set change requires a reboot. This value is only useful during an update process and gets reset on refresh.`,
				},
				resource.Attribute{
					Name:        "vmware_tools_status",
					Description: `The state of VMware tools in the guest. This will determine the proper course of action for some device operations.`,
				},
				resource.Attribute{
					Name:        "vmx_path",
					Description: `The path of the virtual machine's configuration file in the VM's datastore.`,
				},
				resource.Attribute{
					Name:        "imported",
					Description: `This is flagged if the virtual machine has been imported, or the state has been migrated from a previous version of the resource. It influences the behavior of the first post-import apply operation. See the section on [importing](#importing) below.`,
				},
				resource.Attribute{
					Name:        "change_version",
					Description: `A unique identifier for a given version of the last configuration applied, such the timestamp of the last update to the configuration.`,
				},
				resource.Attribute{
					Name:        "uuid",
					Description: `The UUID of the virtual machine. Also exposed as the ` + "`" + `id` + "`" + ` of the resource.`,
				},
				resource.Attribute{
					Name:        "default_ip_address",
					Description: `The IP address selected by Terraform to be used with any [provisioners][tf-docs-provisioners] configured on this resource. Whenever possible, this is the first IPv4 address that is reachable through the default gateway configured on the machine, then the first reachable IPv6 address, and then the first general discovered address if neither exist. If VMware tools is not running on the virtual machine, or if the VM is powered off, this value will be blank.`,
				},
				resource.Attribute{
					Name:        "guest_ip_addresses",
					Description: `The current list of IP addresses on this machine, including the value of ` + "`" + `default_ip_address` + "`" + `. If VMware tools is not running on the virtual machine, or if the VM is powered off, this list will be empty.`,
				},
				resource.Attribute{
					Name:        "vapp_transport",
					Description: `Computed value which is only valid for cloned virtual machines. A list of vApp transport methods supported by the source virtual machine or template. [docs-about-morefs]: /docs/providers/vsphere/index.html#use-of-managed-object-references-by-the-vsphere-provider ## Importing An existing virtual machine can be [imported][docs-import] into this resource via supplying the full path to the virtual machine. An example is below: [docs-import]: /docs/import/index.html ` + "`" + `` + "`" + `` + "`" + ` terraform import vsphere_virtual_machine.vm /dc1/vm/srv1 ` + "`" + `` + "`" + `` + "`" + ` The above would import the virtual machine named ` + "`" + `srv1` + "`" + ` that is located in the ` + "`" + `dc1` + "`" + ` datacenter. ### Additional requirements and notes for importing Many of the same requirements for [cloning](#additional-requirements-and-notes-for-cloning) apply to importing, although since importing writes directly to state, a lot of these rules cannot be enforced at import time, so every effort should be made to ensure the correctness of the configuration before the import. In addition to these rules, the following extra rules apply to importing:`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "vsphere_virtual_machine_snapshot",
			Category:         "Virtual Machine Resources",
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
			Type:             "vsphere_vmfs_datastore",
			Category:         "Storage Resources",
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
	}

	resourcesMap = map[string]int{

		"vsphere_compute_cluster":                         0,
		"vsphere_compute_cluster_host_group":              1,
		"vsphere_compute_cluster_vm_affinity_rule":        2,
		"vsphere_compute_cluster_vm_anti_affinity_rule":   3,
		"vsphere_compute_cluster_vm_dependency_rule":      4,
		"vsphere_compute_cluster_vm_group":                5,
		"vsphere_compute_cluster_vm_host_rule":            6,
		"vsphere_custom_attribute":                        7,
		"vsphere_datacenter":                              8,
		"vsphere_datastore_cluster":                       9,
		"vsphere_datastore_cluster_vm_anti_affinity_rule": 10,
		"vsphere_distributed_port_group":                  11,
		"vsphere_distributed_virtual_switch":              12,
		"vsphere_dpm_host_override":                       13,
		"vsphere_drs_vm_override":                         14,
		"vsphere_file":                                    15,
		"vsphere_folder":                                  16,
		"vsphere_ha_vm_override":                          17,
		"vsphere_host_port_group":                         18,
		"vsphere_host_virtual_switch":                     19,
		"vsphere_license":                                 20,
		"vsphere_nas_datastore":                           21,
		"vsphere_resource_pool":                           22,
		"vsphere_storage_drs_vm_override":                 23,
		"vsphere_tag":                                     24,
		"vsphere_tag_category":                            25,
		"vsphere_vapp_container":                          26,
		"vsphere_vapp_entity":                             27,
		"vsphere_virtual_disk":                            28,
		"vsphere_virtual_machine":                         29,
		"vsphere_virtual_machine_snapshot":                30,
		"vsphere_vmfs_datastore":                          31,
	}
)

func GetResource(r string) (*resource.Resource, error) {
	rs, ok := resourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("resource %q not found", r)
	}
	return Resources[rs], nil
}
