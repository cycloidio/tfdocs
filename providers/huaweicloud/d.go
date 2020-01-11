package huaweicloud

import (
	"fmt"

	"github.com/cycloidio/tfdocs/resource"
)

var (
	DataSources = []*resource.Resource{

		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_antiddos_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides status of a specific EIP.`,
			Description: `

The HuaweiCloud Antiddos data source allows to query the status of EIP, regardless whether an EIP has been bound to an Elastic Cloud Server (ECS) or not.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "floating_ip_id",
					Description: `(Optional) The Elastic IP ID.`,
				},
				resource.Attribute{
					Name:        "floating_ip_address",
					Description: `(Optional) The Elastic IP address.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The defense status. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "network_type",
					Description: `The EIP type.`,
				},
				resource.Attribute{
					Name:        "period_start",
					Description: `The Start time.`,
				},
				resource.Attribute{
					Name:        "bps_attack",
					Description: `The Attack traffic in (bit/s).`,
				},
				resource.Attribute{
					Name:        "bps_in",
					Description: `The inbound traffic in (bit/s).`,
				},
				resource.Attribute{
					Name:        "total_bps",
					Description: `The total traffic.`,
				},
				resource.Attribute{
					Name:        "pps_in",
					Description: `The inbound packet rate (number of packets per second).`,
				},
				resource.Attribute{
					Name:        "pps_attack",
					Description: `The attack packet rate (number of packets per second).`,
				},
				resource.Attribute{
					Name:        "total_pps",
					Description: `The total packet rate.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start time of cleaning and blackhole event.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of cleaning and blackhole event.`,
				},
				resource.Attribute{
					Name:        "traffic_cleaning_status",
					Description: `The traffic cleaning status.`,
				},
				resource.Attribute{
					Name:        "trigger_bps",
					Description: `The traffic at the triggering point.`,
				},
				resource.Attribute{
					Name:        "trigger_pps",
					Description: `The packet rate at the triggering point.`,
				},
				resource.Attribute{
					Name:        "trigger_http_pps",
					Description: `The HTTP request rate at the triggering point.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "network_type",
					Description: `The EIP type.`,
				},
				resource.Attribute{
					Name:        "period_start",
					Description: `The Start time.`,
				},
				resource.Attribute{
					Name:        "bps_attack",
					Description: `The Attack traffic in (bit/s).`,
				},
				resource.Attribute{
					Name:        "bps_in",
					Description: `The inbound traffic in (bit/s).`,
				},
				resource.Attribute{
					Name:        "total_bps",
					Description: `The total traffic.`,
				},
				resource.Attribute{
					Name:        "pps_in",
					Description: `The inbound packet rate (number of packets per second).`,
				},
				resource.Attribute{
					Name:        "pps_attack",
					Description: `The attack packet rate (number of packets per second).`,
				},
				resource.Attribute{
					Name:        "total_pps",
					Description: `The total packet rate.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `The start time of cleaning and blackhole event.`,
				},
				resource.Attribute{
					Name:        "end_time",
					Description: `The end time of cleaning and blackhole event.`,
				},
				resource.Attribute{
					Name:        "traffic_cleaning_status",
					Description: `The traffic cleaning status.`,
				},
				resource.Attribute{
					Name:        "trigger_bps",
					Description: `The traffic at the triggering point.`,
				},
				resource.Attribute{
					Name:        "trigger_pps",
					Description: `The packet rate at the triggering point.`,
				},
				resource.Attribute{
					Name:        "trigger_http_pps",
					Description: `The HTTP request rate at the triggering point.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cce_cluster_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on Cloud Container Engine Cluster (CCE).`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional)The Name of the cluster resource.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The ID of container cluster.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The state of the cluster.`,
				},
				resource.Attribute{
					Name:        "cluster_type",
					Description: `(Optional) Type of the cluster. Possible values: VirtualMachine, BareMetal or Windows ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference:`,
				},
				resource.Attribute{
					Name:        "billingMode",
					Description: `Charging mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cluster in string format.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The cluster specification in string format.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of cluster in string format.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `The container network segment.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `The container network type: overlay_l2 , underlay_ipvlan or vpc-router.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet used to create the node.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `The ID of the high speed network used to create bare metal nodes.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `The address accessed within the user's subnet.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Public network access address.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "billingMode",
					Description: `Charging mode of the cluster.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Cluster description.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cluster in string format.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the cluster.`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The cluster specification in string format.`,
				},
				resource.Attribute{
					Name:        "cluster_version",
					Description: `The version of cluster in string format.`,
				},
				resource.Attribute{
					Name:        "container_network_cidr",
					Description: `The container network segment.`,
				},
				resource.Attribute{
					Name:        "container_network_type",
					Description: `The container network type: overlay_l2 , underlay_ipvlan or vpc-router.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `The ID of the subnet used to create the node.`,
				},
				resource.Attribute{
					Name:        "highway_subnet_id",
					Description: `The ID of the high speed network used to create bare metal nodes.`,
				},
				resource.Attribute{
					Name:        "internal",
					Description: `The address accessed within the user's subnet.`,
				},
				resource.Attribute{
					Name:        "external",
					Description: `Public network access address.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cce_nodes_v3",
			Category:         "Data Sources",
			ShortDescription: `To get the specified node in a cluster.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "Cluster_id",
					Description: `(Required) The id of container cluster.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) - Name of the node.`,
				},
				resource.Attribute{
					Name:        "node_id",
					Description: `(Optional) - The id of the node.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) - The state of the node. ## Attributes Reference All above argument parameters can be exported as attribute parameters along with attribute reference:`,
				},
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The flavor id to be used.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Available partitions where the node is located.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `Key pair name when logging in to select the key pair mode.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `Node's billing mode: The value is 0 (on demand).`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `Bandwidth billing type.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `Bandwidth (Mbit/s), in the range of [1, 2000].`,
				},
				resource.Attribute{
					Name:        "extendparam",
					Description: `Extended parameters.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes in batch creation.`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `List of existing elastic IP IDs.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The node's virtual machine ID in ECS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Elastic IP parameters of the node.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the node`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `Elastic IP address type.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The bandwidth sharing type. NOTE: This parameter is mandatory when share_type is set to PER and is optional when share_type is set to WHOLE with an ID specified. Enumerated values: PER (indicates exclusive bandwidth) and WHOLE (indicates sharing)`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `Disk type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flavor_id",
					Description: `The flavor id to be used.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `Available partitions where the node is located.`,
				},
				resource.Attribute{
					Name:        "key_pair",
					Description: `Key pair name when logging in to select the key pair mode.`,
				},
				resource.Attribute{
					Name:        "billing_mode",
					Description: `Node's billing mode: The value is 0 (on demand).`,
				},
				resource.Attribute{
					Name:        "charge_mode",
					Description: `Bandwidth billing type.`,
				},
				resource.Attribute{
					Name:        "bandwidth_size",
					Description: `Bandwidth (Mbit/s), in the range of [1, 2000].`,
				},
				resource.Attribute{
					Name:        "extendparam",
					Description: `Extended parameters.`,
				},
				resource.Attribute{
					Name:        "node_count",
					Description: `The number of nodes in batch creation.`,
				},
				resource.Attribute{
					Name:        "eip_ids",
					Description: `List of existing elastic IP IDs.`,
				},
				resource.Attribute{
					Name:        "server_id",
					Description: `The node's virtual machine ID in ECS.`,
				},
				resource.Attribute{
					Name:        "public_ip",
					Description: `Elastic IP parameters of the node.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `Private IP of the node`,
				},
				resource.Attribute{
					Name:        "ip_type",
					Description: `Elastic IP address type.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The bandwidth sharing type. NOTE: This parameter is mandatory when share_type is set to PER and is optional when share_type is set to WHOLE with an ID specified. Enumerated values: PER (indicates exclusive bandwidth) and WHOLE (indicates sharing)`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `Disk type.`,
				},
				resource.Attribute{
					Name:        "disk_size",
					Description: `Disk size in GB.`,
				},
				resource.Attribute{
					Name:        "volumetype",
					Description: `Disk type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cdm_flavors_v1",
			Category:         "Data Sources",
			ShortDescription: `Get the flavor information on a HuaweiCloud cdm service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments:        []resource.Attribute{},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "version",
					Description: `The version of the flavor.`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the cdm flavor.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The id of the cdm flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_compute_availability_zones_v2",
			Category:         "Data Sources",
			ShortDescription: `Get a list of availability zones from HuaweiCloud`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "state",
					Description: `(Optional) The ` + "`" + `state` + "`" + ` of the availability zones to match, default ("available"). ## Attributes Reference ` + "`" + `id` + "`" + ` is set to hash of the returned zone list. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + ``,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "names",
					Description: `The names of the availability zones, ordered alphanumerically, that match the queried ` + "`" + `state` + "`" + ``,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_csbs_backup_policy_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Backup Policy.`,
			Description: `

The HuaweiCloud CSBS Backup Policy data source allows access of backup Policy resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Specifies the ID of backup policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) Specifies the backup policy name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specifies the backup policy status. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the backup policy description.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `Provides the Backup provider ID.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Specifies the parameters of a backup policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies Scheduling period name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies Scheduling period description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether the scheduling period is enabled.`,
				},
				resource.Attribute{
					Name:        "max_backups",
					Description: `Specifies maximum number of backups that can be automatically created for a backup object.`,
				},
				resource.Attribute{
					Name:        "retention_duration_days",
					Description: `Specifies duration of retaining a backup, in days.`,
				},
				resource.Attribute{
					Name:        "permanent",
					Description: `Specifies whether backups are permanently retained.`,
				},
				resource.Attribute{
					Name:        "trigger_pattern",
					Description: `Specifies Scheduling policy of the scheduler.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `Specifies Operation type, which can be backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Scheduling period ID.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `Specifies Scheduler ID.`,
				},
				resource.Attribute{
					Name:        "trigger_name",
					Description: `Specifies Scheduler name.`,
				},
				resource.Attribute{
					Name:        "trigger_type",
					Description: `Specifies Scheduler type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the ID of the object to be backed up.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Entity object type of the backup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies backup object name.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Specifies the backup policy description.`,
				},
				resource.Attribute{
					Name:        "provider_id",
					Description: `Provides the Backup provider ID.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Specifies the parameters of a backup policy.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies Scheduling period name.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Specifies Scheduling period description.`,
				},
				resource.Attribute{
					Name:        "enabled",
					Description: `Specifies whether the scheduling period is enabled.`,
				},
				resource.Attribute{
					Name:        "max_backups",
					Description: `Specifies maximum number of backups that can be automatically created for a backup object.`,
				},
				resource.Attribute{
					Name:        "retention_duration_days",
					Description: `Specifies duration of retaining a backup, in days.`,
				},
				resource.Attribute{
					Name:        "permanent",
					Description: `Specifies whether backups are permanently retained.`,
				},
				resource.Attribute{
					Name:        "trigger_pattern",
					Description: `Specifies Scheduling policy of the scheduler.`,
				},
				resource.Attribute{
					Name:        "operation_type",
					Description: `Specifies Operation type, which can be backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Scheduling period ID.`,
				},
				resource.Attribute{
					Name:        "trigger_id",
					Description: `Specifies Scheduler ID.`,
				},
				resource.Attribute{
					Name:        "trigger_name",
					Description: `Specifies Scheduler name.`,
				},
				resource.Attribute{
					Name:        "trigger_type",
					Description: `Specifies Scheduler type.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies the ID of the object to be backed up.`,
				},
				resource.Attribute{
					Name:        "type",
					Description: `Entity object type of the backup object.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies backup object name.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_csbs_backup_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific Backup.`,
			Description: `

The HuaweiCloud CSBS Backup data source allows access of backup resources.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) Specifies the ID of backup.`,
				},
				resource.Attribute{
					Name:        "backup_name",
					Description: `(Optional) Specifies the backup name.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Specifies the backup status.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) Specifies the backup object name.`,
				},
				resource.Attribute{
					Name:        "backup_record_id",
					Description: `(Optional) Specifies the backup record ID.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) Specifies the type of backup objects.`,
				},
				resource.Attribute{
					Name:        "resource_id",
					Description: `(Optional) Specifies the backup object ID.`,
				},
				resource.Attribute{
					Name:        "policy_id",
					Description: `(Optional) Specifies the Policy Id.`,
				},
				resource.Attribute{
					Name:        "vm_ip",
					Description: `(Optional) Specifies the ip of VM. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Provides the backup description.`,
				},
				resource.Attribute{
					Name:        "auto_trigger",
					Description: `Specifies whether automatic trigger is enabled.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies average speed.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies the backup capacity.`,
				},
				resource.Attribute{
					Name:        "space_saving_ratio",
					Description: `Specifies the space saving rate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of backup Volume.`,
				},
				resource.Attribute{
					Name:        "space_saving_ratio",
					Description: `Specifies space saving rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `It gives EVS disk backup name.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Specifies whether the disk is bootable.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies the average speed.`,
				},
				resource.Attribute{
					Name:        "source_volume_size",
					Description: `Shows source volume size in GB.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `It specifies source volume ID.`,
				},
				resource.Attribute{
					Name:        "incremental",
					Description: `Shows whether incremental backup is used.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of snapshot.`,
				},
				resource.Attribute{
					Name:        "source_volume_name",
					Description: `Specifies source volume name.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `It specifies backup. The default value is backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Cinder backup ID.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies accumulated size (MB) of backups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of backup data.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Specifies elastic IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "cloud_service_type",
					Description: `Specifies ECS type.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Specifies memory size of the ECS, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Specifies CPU cores corresponding to the ECS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `It specifies internal IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Shows system disk size corresponding to the ECS specifications.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Specifies image type.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `Provides the backup description.`,
				},
				resource.Attribute{
					Name:        "auto_trigger",
					Description: `Specifies whether automatic trigger is enabled.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies average speed.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies the backup capacity.`,
				},
				resource.Attribute{
					Name:        "space_saving_ratio",
					Description: `Specifies the space saving rate.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Status of backup Volume.`,
				},
				resource.Attribute{
					Name:        "space_saving_ratio",
					Description: `Specifies space saving rate.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `It gives EVS disk backup name.`,
				},
				resource.Attribute{
					Name:        "bootable",
					Description: `Specifies whether the disk is bootable.`,
				},
				resource.Attribute{
					Name:        "average_speed",
					Description: `Specifies the average speed.`,
				},
				resource.Attribute{
					Name:        "source_volume_size",
					Description: `Shows source volume size in GB.`,
				},
				resource.Attribute{
					Name:        "source_volume_id",
					Description: `It specifies source volume ID.`,
				},
				resource.Attribute{
					Name:        "incremental",
					Description: `Shows whether incremental backup is used.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `ID of snapshot.`,
				},
				resource.Attribute{
					Name:        "source_volume_name",
					Description: `Specifies source volume name.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `It specifies backup. The default value is backup.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `Specifies Cinder backup ID.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `Specifies accumulated size (MB) of backups.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Name of backup data.`,
				},
				resource.Attribute{
					Name:        "eip",
					Description: `Specifies elastic IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "cloud_service_type",
					Description: `Specifies ECS type.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `Specifies memory size of the ECS, in MB.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Specifies CPU cores corresponding to the ECS.`,
				},
				resource.Attribute{
					Name:        "private_ip",
					Description: `It specifies internal IP address of the ECS.`,
				},
				resource.Attribute{
					Name:        "disk",
					Description: `Shows system disk size corresponding to the ECS specifications.`,
				},
				resource.Attribute{
					Name:        "image_type",
					Description: `Specifies image type.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_cts_tracker_v1",
			Category:         "Data Sources",
			ShortDescription: `CTS tracker allows you to collect, store, and query cloud resource operation records and use these records for security analysis, compliance auditing, resource tracking, and fault locating.`,
			Description: `

CTS Tracker data source allows access of Cloud Tracker.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "tracker_name",
					Description: `(Optional) The tracker name.`,
				},
				resource.Attribute{
					Name:        "bucket_name",
					Description: `(Optional) The OBS bucket name for a tracker.`,
				},
				resource.Attribute{
					Name:        "file_prefix_name",
					Description: `(Optional) The prefix of a log that needs to be stored in an OBS bucket.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) Status of a tracker. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "topic_id",
					Description: `The theme of the SMN service.`,
				},
				resource.Attribute{
					Name:        "is_send_all_key_operation",
					Description: `Specifies Typical or All operations for Trigger Condition.`,
				},
				resource.Attribute{
					Name:        "need_notify_user_list",
					Description: `The users using the login function.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "topic_id",
					Description: `The theme of the SMN service.`,
				},
				resource.Attribute{
					Name:        "is_send_all_key_operation",
					Description: `Specifies Typical or All operations for Trigger Condition.`,
				},
				resource.Attribute{
					Name:        "need_notify_user_list",
					Description: `The users using the login function.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dcs_az_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Huaweicloud dcs az.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the name of an AZ.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Optional) Indicates the code of an AZ.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Indicates the port number of an AZ. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found az. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dcs_maintainwindow_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Huaweicloud dcs maintainwindow.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "seq",
					Description: `(Required) Indicates the sequential number of a maintenance time window.`,
				},
				resource.Attribute{
					Name:        "begin",
					Description: `(Optional) Indicates the time at which a maintenance time window starts.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) Indicates the time at which a maintenance time window ends.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Indicates whether a maintenance time window is set to the default time segment. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found maintainwindow. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dcs_product_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Flexibleengine dcs product.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) Indicates the name of a message engine.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Indicates the version of a message engine.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Indicates an instance type. Options: "single" and "cluster"`,
				},
				resource.Attribute{
					Name:        "vm_specification",
					Description: `(Optional) Indicates VM specifications.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Indicates the message storage space.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Indicates the baseline bandwidth of a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `(Optional) Indicates the maximum number of topics that can be created for a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `(Optional) Indicates an I/O specification.`,
				},
				resource.Attribute{
					Name:        "io_type",
					Description: `(Optional) Indicates an I/O type.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `(Optional) Indicates the number of nodes in a cluster. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found product. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vm_specification",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "io_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vm_specification",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "io_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dis_partition_v2",
			Category:         "Data Sources",
			ShortDescription: `Get all the partitions of a stream`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "stream_name",
					Description: `(Required) Name of the DIS stream. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "partitions",
					Description: `The information of stream partitions. Structure is documented below. The ` + "`" + `partitions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the partition.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the partition.`,
				},
				resource.Attribute{
					Name:        "hash_range",
					Description: `Possible value range of the hash key used by each partition.`,
				},
				resource.Attribute{
					Name:        "sequence_number_range",
					Description: `Sequence number range of each partition..`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "partitions",
					Description: `The information of stream partitions. Structure is documented below. The ` + "`" + `partitions` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `The ID of the partition.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the partition.`,
				},
				resource.Attribute{
					Name:        "hash_range",
					Description: `Possible value range of the hash key used by each partition.`,
				},
				resource.Attribute{
					Name:        "sequence_number_range",
					Description: `Sequence number range of each partition..`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dms_az_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud dms az.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) Indicates the name of an AZ.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `(Optional) Indicates the code of an AZ.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `(Required) Indicates the port number of an AZ. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found az. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dms_maintainwindow_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud dms maintainwindow.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "seq",
					Description: `(Required) Indicates the sequential number of a maintenance time window.`,
				},
				resource.Attribute{
					Name:        "begin",
					Description: `(Optional) Indicates the time at which a maintenance time window starts.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `(Required) Indicates the time at which a maintenance time window ends.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `(Required) Indicates whether a maintenance time window is set to the default time segment. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found maintainwindow. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "begin",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "end",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_dms_product_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud dms product.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `(Required) Indicates the name of a message engine.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `(Optional) Indicates the version of a message engine.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `(Required) Indicates an instance type. Options: "single" and "cluster"`,
				},
				resource.Attribute{
					Name:        "vm_specification",
					Description: `(Optional) Indicates VM specifications.`,
				},
				resource.Attribute{
					Name:        "storage",
					Description: `(Optional) Indicates the message storage space.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `(Optional) Indicates the baseline bandwidth of a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `(Optional) Indicates the maximum number of topics that can be created for a Kafka instance.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `(Optional) Indicates an I/O specification.`,
				},
				resource.Attribute{
					Name:        "io_type",
					Description: `(Optional) Indicates an I/O type.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `(Optional) Indicates the number of nodes in a cluster. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found product. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vm_specification",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "io_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "engine",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "instance_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "vm_specification",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "bandwidth",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "partition_num",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "storage_spec_code",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "io_type",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "node_num",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_identity_role_v3",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud Role.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `The name of the role.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) The domain the role belongs to. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found role. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_images_image_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on a HuaweiCloud Image.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Glance client. A Glance client is needed to create an Image that can be used with a compute instance. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "most_recent",
					Description: `(Optional) If more than one result is returned, use the most recent image.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the image.`,
				},
				resource.Attribute{
					Name:        "owner",
					Description: `(Optional) The owner (UUID) of the image.`,
				},
				resource.Attribute{
					Name:        "size_min",
					Description: `(Optional) The minimum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "size_max",
					Description: `(Optional) The maximum size (in bytes) of the image to return.`,
				},
				resource.Attribute{
					Name:        "sort_direction",
					Description: `(Optional) Order the results in either ` + "`" + `asc` + "`" + ` or ` + "`" + `desc` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "sort_key",
					Description: `(Optional) Sort images based on a certain key. Must be one of "name", "container_format", "disk_format", "status", "id" or "size". Defaults to ` + "`" + `name` + "`" + `.`,
				},
				resource.Attribute{
					Name:        "tag",
					Description: `(Optional) Search for images with a specific tag.`,
				},
				resource.Attribute{
					Name:        "visibility",
					Description: `(Optional) The visibility of the image. Must be one of "public", "private", "community", or "shared". Defaults to ` + "`" + `private` + "`" + `. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found image. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `The minimum amount of disk space required to use the image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `The minimum amount of ram required to use the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Freeform information about the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "checksum",
					Description: `The checksum of the data associated with the image.`,
				},
				resource.Attribute{
					Name:        "created_at",
					Description: `The date the image was created.`,
				},
				resource.Attribute{
					Name:        "file",
					Description: `the trailing path after the glance endpoint that represent the location of the image or the path to retrieve it.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `The metadata associated with the image. Image metadata allow for meaningfully define the image properties and tags.`,
				},
				resource.Attribute{
					Name:        "min_disk_gb",
					Description: `The minimum amount of disk space required to use the image.`,
				},
				resource.Attribute{
					Name:        "min_ram_mb",
					Description: `The minimum amount of ram required to use the image.`,
				},
				resource.Attribute{
					Name:        "properties",
					Description: `Freeform information about the image.`,
				},
				resource.Attribute{
					Name:        "protected",
					Description: `Whether or not the image is protected.`,
				},
				resource.Attribute{
					Name:        "schema",
					Description: `The path to the JSON-schema that represent the image or image`,
				},
				resource.Attribute{
					Name:        "size_bytes",
					Description: `The size of the image (in bytes).`,
				},
				resource.Attribute{
					Name:        "tags",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "update_at",
					Description: `The date the image was last updated.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_kms_data_key_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud KMS data encryption key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_id",
					Description: `(Required) The globally unique identifier for the key. Changing this gets the new data encryption key.`,
				},
				resource.Attribute{
					Name:        "encryption_context",
					Description: `(Optional) The value of this parameter must be a series of "key:value" pairs used to record resource context information. The value of this parameter must not contain sensitive information and must be within 8192 characters in length. Example: {"Key1":"Value1","Key2":"Value2"}`,
				},
				resource.Attribute{
					Name:        "datakey_length",
					Description: `(Required) Number of bits in the length of a DEK (data encryption keys). The maximum number is 512. Changing this gets the new data encryption key. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the date of the found data key. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "plain_text",
					Description: `The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
				resource.Attribute{
					Name:        "cipher_text",
					Description: `The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "plain_text",
					Description: `The plaintext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
				resource.Attribute{
					Name:        "cipher_text",
					Description: `The ciphertext of a DEK is expressed in hexadecimal format, and two characters indicate one byte.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_kms_key_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud KMS Key.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key_alias",
					Description: `(Optional) The alias in which to create the key. It is required when we create a new key. Changing this gets the new key.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `(Optional) The description of the key as viewed in Huawei console. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `(Optional) Region where a key resides. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `(Optional) The globally unique identifier for the key. Changing this gets the new key.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `(Optional) Identification of a Master Key. The value "1" indicates a Default Master Key, and the value "0" indicates a key. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `(Optional) The state of a key. "1" indicates that the key is waiting to be activated. "2" indicates that the key is enabled. "3" indicates that the key is disabled. "4" indicates that the key is scheduled for deletion. Changing this gets a new key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `(Optional) - ID of a user domain for the key. Changing this gets a new key. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found key. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "key_alias",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "key_alias",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_description",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "realm",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "key_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "default_key_flag",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "scheduled_deletion_date",
					Description: `Scheduled deletion time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "domain_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "expiration_time",
					Description: `Expiration time.`,
				},
				resource.Attribute{
					Name:        "creation_date",
					Description: `Creation time (time stamp) of a key.`,
				},
				resource.Attribute{
					Name:        "key_state",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_network_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud Network.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve networks ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the network.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the network.`,
				},
				resource.Attribute{
					Name:        "matching_subnet_cidr",
					Description: `(Optional) The CIDR of a subnet within the network.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the network. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found network. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the network.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `(Optional) Specifies whether the network resource can be accessed by any tenant or not.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_port_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information of an HuaweiCloud Port.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve port ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `(Optional) The owner of the port.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `(Optional) The ID of the port.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the port.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `(Optional) The administrative state of the port.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the port belongs to.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `(Optional) The device owner of the port.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `(Optional) The MAC address of the port.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `(Optional) The ID of the device the port belongs to.`,
				},
				resource.Attribute{
					Name:        "fixed_ip",
					Description: `(Optional) The port IP address filter.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the port.`,
				},
				resource.Attribute{
					Name:        "security_group_ids",
					Description: `(Optional) The list of port security group IDs to filter. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found port. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The set of security group IDs applied on the port.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "project_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "port_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "admin_state_up",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_owner",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "mac_address",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "device_id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "all_fixed_ips",
					Description: `The collection of Fixed IP addresses on the port in the order returned by the Network v2 API.`,
				},
				resource.Attribute{
					Name:        "all_security_group_ids",
					Description: `The set of security group IDs applied on the port.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_secgroup_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud Security Group.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve security groups ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "secgroup_id",
					Description: `(Optional) The ID of the security group.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the security group.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the security group. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found security group. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_networking_subnet_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud Subnet.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V2 Neutron client. A Neutron client is needed to retrieve subnet ids. If omitted, the ` + "`" + `region` + "`" + ` argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enabled",
					Description: `(Optional) If the subnet has DHCP enabled.`,
				},
				resource.Attribute{
					Name:        "dhcp_disabled",
					Description: `(Optional) If the subnet has DHCP disabled.`,
				},
				resource.Attribute{
					Name:        "ip_version",
					Description: `(Optional) The IP version of the subnet (either 4 or 6).`,
				},
				resource.Attribute{
					Name:        "gateway_ip",
					Description: `(Optional) The IP of the subnet's gateway.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The CIDR of the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `(Optional) The ID of the subnet.`,
				},
				resource.Attribute{
					Name:        "network_id",
					Description: `(Optional) The ID of the network the subnet belongs to.`,
				},
				resource.Attribute{
					Name:        "tenant_id",
					Description: `(Optional) The owner of the subnet. ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found subnet. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "allocation_pools",
					Description: `Allocation pools of the subnet.`,
				},
				resource.Attribute{
					Name:        "enable_dhcp",
					Description: `Whether the subnet has DHCP enabled or not.`,
				},
				resource.Attribute{
					Name:        "dns_nameservers",
					Description: `DNS Nameservers of the subnet.`,
				},
				resource.Attribute{
					Name:        "host_routes",
					Description: `Host Routes of the subnet.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rds_flavors_v1",
			Category:         "Data Sources",
			ShortDescription: `Get the flavor information on an HuaweiCloud rds service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Required) The region in which to obtain the V1 rds client.`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `(Required) The datastore name of the rds.`,
				},
				resource.Attribute{
					Name:        "datastore_version",
					Description: `(Required) The datastore version of the rds.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `(Optional) The spec code of a rds flavor. ## Available value for attributes datastore_name | datastore_version | speccode ---- | --- | --- PostgreSQL | 9.5.5 <br> 9.6.3 | ha = True: <br> rds.pg.m1.2xlarge.ha rds.pg.c2.large.ha rds.pg.s1.2xlarge.ha rds.pg.c2.xlarge.ha rds.pg.s1.xlarge.ha rds.pg.m1.xlarge.ha rds.pg.m1.large.ha rds.pg.c2.medium.ha rds.pg.s1.medium.ha rds.pg.s1.large.ha <br> ha = False: <br> rds.pg.s1.xlarge rds.pg.m1.2xlarge rds.pg.c2.xlarge rds.pg.s1.medium rds.pg.c2.medium rds.pg.s1.large rds.pg.c2.large rds.pg.m1.large rds.pg.s1.2xlarge rds.pg.m1.xlarge MySQL| 5.6.33 <br>5.6.30 <br>5.6.34 <br>5.6.35 <br>5.6.36 <br>5.7.17| ha = True: <br> rds.mysql.s1.medium.ha rds.mysql.s1.large.ha rds.mysql.s1.xlarge.ha rds.mysql.s1.2xlarge.ha rds.mysql.s1.8xlarge.ha rds.mysql.s1.4xlarge.ha rds.mysql.m1.2xlarge.ha rds.mysql.c2.medium.ha rds.mysql.c2.large.ha rds.mysql.c2.xlarge.ha rds.mysql.c2.2xlarge.ha rds.mysql.c2.4xlarge.ha rds.mysql.c2.8xlarge.ha rds.mysql.m1.medium.ha rds.mysql.m1.large.ha rds.mysql.m1.xlarge.ha rds.mysql.m1.4xlarge.ha <br> ha = False: <br> rds.mysql.s1.medium rds.mysql.s1.large rds.mysql.s1.xlarge rds.mysql.s1.2xlarge rds.mysql.s1.8xlarge rds.mysql.s1.4xlarge rds.mysql.m1.2xlarge rds.mysql.c2.medium rds.mysql.c2.large rds.mysql.c2.xlarge rds.mysql.c2.2xlarge rds.mysql.c2.4xlarge rds.mysql.c2.8xlarge rds.mysql.m1.medium rds.mysql.m1.large rds.mysql.m1.xlarge rds.mysql.m1.4xlarge SQLServer| 2014 SP2 SE | <br> ha = True: <br> rds.mssql.m1.2xlarge.ha rds.mssql.m1.xlarge.ha rds.mssql.m1.4xlarge.ha rds.mssql.s1.xlarge.ha rds.mssql.c2.xlarge.ha rds.mssql.s1.2xlarge.ha <br> ha = False: <br> rds.mssql.m1.2xlarge rds.mssql.m1.xlarge rds.mssql.m1.4xlarge rds.mssql.s1.xlarge rds.mssql.c2.xlarge rds.mssql.s1.2xlarge ## Attributes Reference ` + "`" + `id` + "`" + ` is set to the ID of the found rds flavor. In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The name of the rds flavor.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "datastore_version",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "speccode",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "ram",
					Description: `The name of the rds flavor.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rds_flavors_v3",
			Category:         "Data Sources",
			ShortDescription: `Get the flavor information on an HuaweiCloud rds service.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "db_type",
					Description: `(Required) Specifies the DB engine. Value: MySQL, PostgreSQL, SQLServer.`,
				},
				resource.Attribute{
					Name:        "db_version",
					Description: `(Required) Specifies the database version. Available value: type | version ---- | --- MySQL| 5.6 <br>5.7 <br>8.0 PostgreSQL | 9.5 <br> 9.6 <br>10 <br>11 SQLServer| 2008_R2_EE <br>2008_R2_WEB <br>2012_SE <br>2014_SE <br>2016_SE <br>2017_SE <br>2012_EE <br>2014_EE <br>2016_EE <br>2017_EE <br>2012_WEB <br>2014_WEB <br>2016_WEB <br>2017_WEB`,
				},
				resource.Attribute{
					Name:        "instance_mode",
					Description: `(Required) The mode of instance. Value: ha(indicates primary/standby instance), single(indicates single instance) ## Attributes Reference In addition, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Indicates the CPU size.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Indicates the memory size in GB.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See 'instance_mode' above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "flavors",
					Description: `Indicates the flavors information. Structure is documented below. The ` + "`" + `flavors` + "`" + ` block contains:`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `The name of the rds flavor.`,
				},
				resource.Attribute{
					Name:        "vcpus",
					Description: `Indicates the CPU size.`,
				},
				resource.Attribute{
					Name:        "memory",
					Description: `Indicates the memory size in GB.`,
				},
				resource.Attribute{
					Name:        "mode",
					Description: `See 'instance_mode' above.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rts_software_config_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific RTS Software Config.`,
			Description: `

The RTS Software Config data source provides details about a specific RTS Software Config.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the software configuration.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the software configuration. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "group",
					Description: `The namespace that groups this software configuration by when it is delivered to a server.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `A list of software configuration inputs.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of software configuration outputs.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The software configuration code.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `The software configuration options.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "group",
					Description: `The namespace that groups this software configuration by when it is delivered to a server.`,
				},
				resource.Attribute{
					Name:        "inputs",
					Description: `A list of software configuration inputs.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of software configuration outputs.`,
				},
				resource.Attribute{
					Name:        "config",
					Description: `The software configuration code.`,
				},
				resource.Attribute{
					Name:        "options",
					Description: `The software configuration options.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rts_stack_resource_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata of an RTS stack resource`,
			Description: `

The HuaweiCloud RTS Stack Resource data source allows access to stack resource metadata.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "stack_name",
					Description: `(Required) The unique stack name.`,
				},
				resource.Attribute{
					Name:        "resource_name",
					Description: `(Optional) The name of a resource in the stack.`,
				},
				resource.Attribute{
					Name:        "physical_resource_id",
					Description: `(Optional) The physical resource ID.`,
				},
				resource.Attribute{
					Name:        "resource_type",
					Description: `(Optional) The resource type. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "logical_resource_id",
					Description: `The logical resource ID.`,
				},
				resource.Attribute{
					Name:        "resource_status",
					Description: `The status of the resource.`,
				},
				resource.Attribute{
					Name:        "resource_status_reason",
					Description: `The resource operation reason.`,
				},
				resource.Attribute{
					Name:        "required_by",
					Description: `Specifies the resource dependency.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "logical_resource_id",
					Description: `The logical resource ID.`,
				},
				resource.Attribute{
					Name:        "resource_status",
					Description: `The status of the resource.`,
				},
				resource.Attribute{
					Name:        "resource_status_reason",
					Description: `The resource operation reason.`,
				},
				resource.Attribute{
					Name:        "required_by",
					Description: `Specifies the resource dependency.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_rts_stack_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an HuaweiCloud RTS.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Required) The name of the stack. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describes the stack.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `Specifies whether to perform a rollback if the update fails.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of stack outputs.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Specifies the stack parameters.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `Structure containing the template body.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `Specifies the timeout duration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the stack name.`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `Specifies the description of the stack operation.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "capabilities",
					Description: `List of stack capabilities for stack.`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `Describes the stack.`,
				},
				resource.Attribute{
					Name:        "disable_rollback",
					Description: `Specifies whether to perform a rollback if the update fails.`,
				},
				resource.Attribute{
					Name:        "outputs",
					Description: `A list of stack outputs.`,
				},
				resource.Attribute{
					Name:        "parameters",
					Description: `Specifies the stack parameters.`,
				},
				resource.Attribute{
					Name:        "template_body",
					Description: `Structure containing the template body.`,
				},
				resource.Attribute{
					Name:        "timeout_mins",
					Description: `Specifies the timeout duration.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `Specifies the stack status.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `Specifies the stack name.`,
				},
				resource.Attribute{
					Name:        "status_reason",
					Description: `Specifies the description of the stack operation.`,
				},
				resource.Attribute{
					Name:        "notification_topics",
					Description: `List of notification topics for stack.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_s3_bucket_object",
			Category:         "Data Sources",
			ShortDescription: `Provides metadata and optionally content of an S3 object`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "bucket",
					Description: `(Required) The name of the bucket to read the object from`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `(Required) The full path to the object inside the bucket`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `(Optional) Specific version ID of the object returned (defaults to latest version) ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.huaweicloud.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `[Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for ` + "`" + `Standard` + "`" + ` storage class objects.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "body",
					Description: `Object data (see`,
				},
				resource.Attribute{
					Name:        "cache_control",
					Description: `Specifies caching behavior along the request/reply chain.`,
				},
				resource.Attribute{
					Name:        "content_disposition",
					Description: `Specifies presentational information for the object.`,
				},
				resource.Attribute{
					Name:        "content_encoding",
					Description: `Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field.`,
				},
				resource.Attribute{
					Name:        "content_language",
					Description: `The language the content is in.`,
				},
				resource.Attribute{
					Name:        "content_length",
					Description: `Size of the body in bytes.`,
				},
				resource.Attribute{
					Name:        "content_type",
					Description: `A standard MIME type describing the format of the object data.`,
				},
				resource.Attribute{
					Name:        "etag",
					Description: `[ETag](https://en.wikipedia.org/wiki/HTTP_ETag) generated for the object (an MD5 sum of the object content in case it's not encrypted)`,
				},
				resource.Attribute{
					Name:        "expiration",
					Description: `If the object expiration is configured (see [object lifecycle management](http://docs.huaweicloud.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html)), the field includes this header. It includes the expiry-date and rule-id key value pairs providing object expiration information. The value of the rule-id is URL encoded.`,
				},
				resource.Attribute{
					Name:        "expires",
					Description: `The date and time at which the object is no longer cacheable.`,
				},
				resource.Attribute{
					Name:        "last_modified",
					Description: `Last modified date of the object in RFC1123 format (e.g. ` + "`" + `Mon, 02 Jan 2006 15:04:05 MST` + "`" + `)`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `A map of metadata stored with the object in S3`,
				},
				resource.Attribute{
					Name:        "server_side_encryption",
					Description: `If the object is stored using server-side encryption (KMS or Amazon S3-managed encryption key), this field includes the chosen encryption and algorithm used.`,
				},
				resource.Attribute{
					Name:        "sse_kms_key_id",
					Description: `If present, specifies the ID of the Key Management Service (KMS) master encryption key that was used for the object.`,
				},
				resource.Attribute{
					Name:        "storage_class",
					Description: `[Storage class](http://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html) information of the object. Available for all objects except for ` + "`" + `Standard` + "`" + ` storage class objects.`,
				},
				resource.Attribute{
					Name:        "version_id",
					Description: `The latest version ID of the object returned.`,
				},
				resource.Attribute{
					Name:        "website_redirect_location",
					Description: `If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. Amazon S3 stores the value of this header in the object metadata.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_sfs_file_sharing_v2",
			Category:         "Data Sources",
			ShortDescription: `Get information on an Huawei Cloud shared file system.`,
			Description: `

Provides information about an Shared File System (SFS).

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The UUID of the shared file system.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the shared file system. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size (GB) of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type for the shared file system, such as high-performance storage (composed of SSDs) or large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `The level of visibility for the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The protocol for sharing file systems.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key and value pairs as a dictionary of strings.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The path for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `The level of the access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `The type of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `The UUID of the mount location of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_instance_id",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `Identifies which mount locations are most efficient and are used preferentially when multiple mount locations exist.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The availability zone name.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size (GB) of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_type",
					Description: `The storage service type for the shared file system, such as high-performance storage (composed of SSDs) or large-capacity storage (composed of SATA disks).`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `The status of the shared file system.`,
				},
				resource.Attribute{
					Name:        "host",
					Description: `The host name of the shared file system.`,
				},
				resource.Attribute{
					Name:        "is_public",
					Description: `The level of visibility for the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_proto",
					Description: `The protocol for sharing file systems.`,
				},
				resource.Attribute{
					Name:        "volume_type",
					Description: `The volume type.`,
				},
				resource.Attribute{
					Name:        "metadata",
					Description: `Metadata key and value pairs as a dictionary of strings.`,
				},
				resource.Attribute{
					Name:        "export_location",
					Description: `The path for accessing the shared file system.`,
				},
				resource.Attribute{
					Name:        "access_level",
					Description: `The level of the access rule.`,
				},
				resource.Attribute{
					Name:        "access_rules_status",
					Description: `The status of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_type",
					Description: `The type of the share access rule.`,
				},
				resource.Attribute{
					Name:        "access_to",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "share_access_id",
					Description: `The UUID of the share access rule.`,
				},
				resource.Attribute{
					Name:        "mount_id",
					Description: `The UUID of the mount location of the shared file system.`,
				},
				resource.Attribute{
					Name:        "share_instance_id",
					Description: `The access that the back end grants or denies.`,
				},
				resource.Attribute{
					Name:        "preferred",
					Description: `Identifies which mount locations are most efficient and are used preferentially when multiple mount locations exist.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vbs_backup_policy_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VBS backup policy.`,
			Description: `

The VBS Backup Policy data source provides details about a specific VBS backup policy.


`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "key",
					Description: `(Required) Specifies the tag key. Tag keys must be unique.`,
				},
				resource.Attribute{
					Name:        "values",
					Description: `(Required) Specifies the List of tag values. This list can have a maximum of 10 values and all be unique. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Specifies the start time of the backup job.The value is in the HH:mm format.`,
				},
				resource.Attribute{
					Name:        "retain_first_backup",
					Description: `Specifies whether to retain the first backup in the current month.`,
				},
				resource.Attribute{
					Name:        "rentention_num",
					Description: `Specifies number of retained backups.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Specifies the backup interval. The value is in the range of 1 to 14 days.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Specifies the tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the tag value.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "start_time",
					Description: `Specifies the start time of the backup job.The value is in the HH:mm format.`,
				},
				resource.Attribute{
					Name:        "retain_first_backup",
					Description: `Specifies whether to retain the first backup in the current month.`,
				},
				resource.Attribute{
					Name:        "rentention_num",
					Description: `Specifies number of retained backups.`,
				},
				resource.Attribute{
					Name:        "frequency",
					Description: `Specifies the backup interval. The value is in the range of 1 to 14 days.`,
				},
				resource.Attribute{
					Name:        "key",
					Description: `Specifies the tag key.`,
				},
				resource.Attribute{
					Name:        "value",
					Description: `Specifies the tag value.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vbs_backup_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VBS Backup.`,
			Description: `

The VBS Backup data source provides details about a specific VBS Backup.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) The name of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "volume_id",
					Description: `(Optional) The source volume ID of the backup.`,
				},
				resource.Attribute{
					Name:        "snapshot_id",
					Description: `(Optional) ID of the snapshot associated with the backup.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The status of the VBS backup. ## Attributes Reference In addition to all arguments above, the following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "description",
					Description: `The description of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ where the backup resides.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `The container of the backup.`,
				},
				resource.Attribute{
					Name:        "service_metadata",
					Description: `The metadata of the vbs backup.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "description",
					Description: `The description of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "availability_zone",
					Description: `The AZ where the backup resides.`,
				},
				resource.Attribute{
					Name:        "size",
					Description: `The size of the vbs backup.`,
				},
				resource.Attribute{
					Name:        "container",
					Description: `The container of the backup.`,
				},
				resource.Attribute{
					Name:        "service_metadata",
					Description: `The metadata of the vbs backup.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_peering_connection_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC peering connection.`,
			Description: `

The VPC Peering Connection data source provides details about a specific VPC peering connection.


`,
			Keywords:   []string{},
			Arguments:  []resource.Attribute{},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_route_ids_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of route Ids for a VPC`,
			Description: `

` + "`" + `huaweicloud_vpc_route_ids_v2` + "`" + ` provides a list of route ids for a vpc_id.

This resource can be useful for getting back a list of route ids for a vpc.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the route ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the route ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_route_v2",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC Route.`,
			Description: `

` + "`" + `huaweicloud_vpc_route_v2` + "`" + ` provides details about a specific VPC route.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "nexthop",
					Description: `The next hop of the route. If the route type is peering, it will provide VPC peering connection ID.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "nexthop",
					Description: `The next hop of the route. If the route type is peering, it will provide VPC peering connection ID.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_subnet_ids_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides a list of subnet Ids for a VPC`,
			Description: `

` + "`" + `huaweicloud_vpc_subnet_ids_v1` + "`" + ` provides a list of subnet ids for a vpc_id

This resource can be useful for getting back a list of subnet ids for a vpc.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the subnet ids found. This data source will fail if none are found.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "ids",
					Description: `A list of all the subnet ids found. This data source will fail if none are found.`,
				},
			},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_subnet_v1",
			Category:         "Data Sources",
			ShortDescription: `Provides details about a specific VPC subnet`,
			Description: `

` + "`" + `huaweicloud_vpc_subnet_v1` + "`" + ` provides details about a specific VPC subnet.

This resource can prove useful when a module accepts a subnet id as
an input variable and needs to, for example, determine the id of the
VPC that the subnet belongs to.

`,
			Keywords: []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) - Specifies a resource ID in UUID format.`,
				},
				resource.Attribute{
					Name:        "dns_list",
					Description: `The IP address list of DNS servers on the subnet.`,
				},
				resource.Attribute{
					Name:        "dhcp_enable",
					Description: `DHCP function for the subnet.`,
				},
				resource.Attribute{
					Name:        "subnet_id",
					Description: `Specifies the subnet (Native OpenStack API) ID.`,
				},
			},
			Attributes: []resource.Attribute{},
		},
		&resource.Resource{
			Name:             "",
			Type:             "huaweicloud_vpc_v1",
			Category:         "Data Sources",
			ShortDescription: `Get information on an huaweicloud VPC.`,
			Description:      ``,
			Keywords:         []string{},
			Arguments: []resource.Attribute{
				resource.Attribute{
					Name:        "region",
					Description: `(Optional) The region in which to obtain the V1 VPC client. A VPC client is needed to retrieve VPCs. If omitted, the region argument of the provider is used.`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `(Optional) The id of the specific VPC to retrieve.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `(Optional) The current status of the desired VPC. Can be either CREATING, OK, DOWN, PENDING_UPDATE, PENDING_DELETE, or ERROR.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `(Optional) A unique name for the VPC. The name must be unique for a tenant. The value is a string of no more than 64 characters and can contain digits, letters, underscores (_), and hyphens (-).`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `(Optional) The cidr block of the desired VPC. ## Attributes Reference The following attributes are exported:`,
				},
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The list of route information with destination and nexthop fields.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
			Attributes: []resource.Attribute{
				resource.Attribute{
					Name:        "id",
					Description: `ID of the VPC.`,
				},
				resource.Attribute{
					Name:        "name",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "status",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "cidr",
					Description: `See Argument Reference above.`,
				},
				resource.Attribute{
					Name:        "routes",
					Description: `The list of route information with destination and nexthop fields.`,
				},
				resource.Attribute{
					Name:        "shared",
					Description: `Specifies whether the cross-tenant sharing is supported.`,
				},
				resource.Attribute{
					Name:        "region",
					Description: `See Argument Reference above.`,
				},
			},
		},
	}

	dataSourcesMap = map[string]int{

		"huaweicloud_antiddos_v1":                   0,
		"huaweicloud_cce_cluster_v3":                1,
		"huaweicloud_cce_nodes_v3":                  2,
		"huaweicloud_cdm_flavors_v1":                3,
		"huaweicloud_compute_availability_zones_v2": 4,
		"huaweicloud_csbs_backup_policy_v1":         5,
		"huaweicloud_csbs_backup_v1":                6,
		"huaweicloud_cts_tracker_v1":                7,
		"huaweicloud_dcs_az_v1":                     8,
		"huaweicloud_dcs_maintainwindow_v1":         9,
		"huaweicloud_dcs_product_v1":                10,
		"huaweicloud_dis_partition_v2":              11,
		"huaweicloud_dms_az_v1":                     12,
		"huaweicloud_dms_maintainwindow_v1":         13,
		"huaweicloud_dms_product_v1":                14,
		"huaweicloud_identity_role_v3":              15,
		"huaweicloud_images_image_v2":               16,
		"huaweicloud_kms_data_key_v1":               17,
		"huaweicloud_kms_key_v1":                    18,
		"huaweicloud_networking_network_v2":         19,
		"huaweicloud_networking_port_v2":            20,
		"huaweicloud_networking_secgroup_v2":        21,
		"huaweicloud_networking_subnet_v2":          22,
		"huaweicloud_rds_flavors_v1":                23,
		"huaweicloud_rds_flavors_v3":                24,
		"huaweicloud_rts_software_config_v1":        25,
		"huaweicloud_rts_stack_resource_v1":         26,
		"huaweicloud_rts_stack_v1":                  27,
		"huaweicloud_s3_bucket_object":              28,
		"huaweicloud_sfs_file_sharing_v2":           29,
		"huaweicloud_vbs_backup_policy_v2":          30,
		"huaweicloud_vbs_backup_v2":                 31,
		"huaweicloud_vpc_peering_connection_v2":     32,
		"huaweicloud_vpc_route_ids_v2":              33,
		"huaweicloud_vpc_route_v2":                  34,
		"huaweicloud_vpc_subnet_ids_v1":             35,
		"huaweicloud_vpc_subnet_v1":                 36,
		"huaweicloud_vpc_v1":                        37,
	}
)

func GetDataSource(r string) (*resource.Resource, error) {
	rs, ok := dataSourcesMap[r]
	if !ok {
		return nil, fmt.Errorf("datasource %q not found", r)
	}
	return DataSources[rs], nil
}
